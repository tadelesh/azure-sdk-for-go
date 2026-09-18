// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package streaming

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"iter"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Frame is a single Server-Sent Event frame parsed from a text/event-stream
// body. It exposes only the wire-level SSE envelope; the strongly-typed payload
// is produced by a generated per-stream decoder that consumes a Frame.
type Frame struct {
	// Type is the value of the SSE "event" field. An empty value means the
	// default "message" event type.
	Type string

	// Data is the concatenated "data" field(s) for the event with the single
	// trailing newline removed. For JSON events this is the raw JSON document;
	// for @data/text events it is the raw text payload.
	Data []byte

	// ID is the value of the SSE "id" field. It carries over to subsequent
	// events until changed, per the SSE specification.
	ID string

	// Retry is the reconnection delay in milliseconds from the SSE "retry"
	// field, or -1 when unset or invalid.
	Retry int
}

// Event provides typed, forward-only iteration over a Server-Sent Events
// response body. T is the generated event union for the operation.
//
// The zero value is a valid, already-exhausted stream: iteration yields no
// events and Close is a no-op.
type Event[T any] struct {
	body      io.ReadCloser
	scanner   *sseScanner
	decode    func(Frame) (T, bool, error)
	connect   func(ctx context.Context, lastEventID string) (*http.Response, error)
	ctx       context.Context // governs the stream lifetime; used for reconnect
	reconnect bool
	done      bool
	lastID    string
	retry     int
	segFrames int // events delivered since the current segment connected
}

// EventStreamHandler supplies the typed decoder and the connection factory used
// to open, and on an unexpected disconnect reopen, an SSE stream.
type EventStreamHandler[T any] struct {
	// Decode maps a wire-level Frame to the typed union value; it returns
	// terminal=true when the event signals the end of the stream.
	Decode func(frame Frame) (value T, terminal bool, err error)

	// Connect opens the stream. lastEventID is empty on the initial connect and
	// carries the most recent event id on a reconnect so the server can resume.
	Connect func(ctx context.Context, lastEventID string) (*http.Response, error)

	// Reconnect enables transparent reconnection after an unexpected mid-stream
	// disconnect. Reconnection resumes from the last seen event id, so it only
	// takes effect once the stream has emitted an id-bearing frame: a stream
	// whose frames carry no id cannot be resumed and is treated as complete at a
	// clean end of body rather than replayed. It must remain false for streams
	// that signal completion with a clean end of body rather than a terminal event.
	Reconnect bool
}

// NewEvent opens an SSE stream via handler.Connect and returns a typed reader
// over it. The provided ctx governs the lifetime of the whole stream, including
// any reconnect attempts made by Next after an unexpected disconnect; canceling
// ctx aborts an in-progress reconnect and ends the stream.
func NewEvent[T any](ctx context.Context, handler EventStreamHandler[T]) (*Event[T], error) {
	resp, err := handler.Connect(ctx, "")
	if err != nil {
		return nil, err
	}
	return &Event[T]{
		body:      resp.Body,
		scanner:   newSSEScanner(resp.Body),
		decode:    handler.Decode,
		connect:   handler.Connect,
		ctx:       ctx,
		reconnect: handler.Reconnect,
		retry:     -1,
	}, nil
}

func newSSEScanner(body io.ReadCloser) *sseScanner {
	sc := bufio.NewScanner(body)
	sc.Split(scanSSELines)
	// SSE payloads can carry large JSON documents; grow the buffer accordingly.
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	return &sseScanner{scanner: sc}
}

// Next returns the next typed event in the stream. It returns io.EOF when the
// stream is complete, including when a terminal event is reached. When the
// Event was created with reconnect support, an unexpected end of the current
// segment triggers a reconnect (honoring the server retry delay, bound by the
// stream's context) before Next reports io.EOF.
func (s *Event[T]) Next() (T, error) {
	var zero T
	// a nil scanner is an empty stream: nothing to consume.
	if s.done || s.scanner == nil {
		return zero, io.EOF
	}
	for {
		frame, err := s.scanner.next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// Reconnect only from a resumable position: SSE resumes via
				// Last-Event-ID, so a stream that never emitted an event id cannot
				// be resumed and reconnecting would only replay it indefinitely.
				// The segFrames guard stops a resumed-but-empty segment from looping.
				if s.reconnect && s.lastID != "" && s.segFrames > 0 {
					if rerr := s.reopen(); rerr != nil {
						s.done = true
						return zero, rerr
					}
					continue
				}
				s.done = true
			}
			return zero, err
		}
		s.lastID = frame.ID
		if frame.Retry >= 0 {
			s.retry = frame.Retry
		}
		value, terminal, derr := s.decode(frame)
		if derr != nil {
			return zero, derr
		}
		if terminal {
			s.done = true
			return zero, io.EOF
		}
		s.segFrames++
		return value, nil
	}
}

// reopen waits the server-suggested retry delay, then reconnects via the
// factory and swaps in the new segment's body, forwarding the last event id.
// It is bound by the stream's context.
func (s *Event[T]) reopen() error {
	if s.retry > 0 {
		timer := time.NewTimer(time.Duration(s.retry) * time.Millisecond)
		defer timer.Stop()
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		case <-timer.C:
		}
	}
	resp, err := s.connect(s.ctx, s.lastID)
	if err != nil {
		return err
	}
	if s.body != nil {
		s.body.Close()
	}
	s.body = resp.Body
	s.scanner = newSSEScanner(resp.Body)
	s.segFrames = 0
	return nil
}

// Events returns a range-over-func iterator over the stream. Iteration ends at
// end of stream (io.EOF is not yielded) or after the first error is yielded.
func (s *Event[T]) Events() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			value, err := s.Next()
			if errors.Is(err, io.EOF) {
				return
			}
			if !yield(value, err) {
				return
			}
			if err != nil {
				return
			}
		}
	}
}

// LastEventID returns the id of the most recently received event. On reconnect
// this value is sent in the Last-Event-ID request header.
func (s *Event[T]) LastEventID() string { return s.lastID }

// Close closes the underlying response body, if any.
func (s *Event[T]) Close() error {
	if s.body == nil {
		return nil
	}
	return s.body.Close()
}

// sseScanner turns a stream of SSE lines into discrete Frame values following
// the WHATWG event stream parsing rules.
type sseScanner struct {
	scanner *bufio.Scanner
	lastID  string
}

func (s *sseScanner) next() (Frame, error) {
	var (
		frame    Frame
		dataBuf  bytes.Buffer
		haveData bool
		haveAny  bool
	)
	frame.Retry = -1
	frame.ID = s.lastID
	for s.scanner.Scan() {
		line := s.scanner.Text()
		if line == "" {
			if !haveAny {
				continue // ignore leading blank lines between events
			}
			return finishFrame(&frame, &dataBuf, haveData), nil
		}
		haveAny = true
		if strings.HasPrefix(line, ":") {
			continue // comment line
		}
		field, value := splitSSEField(line)
		switch field {
		case "event":
			frame.Type = value
		case "data":
			dataBuf.WriteString(value)
			dataBuf.WriteByte('\n')
			haveData = true
		case "id":
			if !strings.ContainsRune(value, 0) { // ignore ids containing U+0000 NULL
				frame.ID = value
				s.lastID = value
			}
		case "retry":
			if isASCIIDigits(value) {
				if n, err := strconv.Atoi(value); err == nil {
					frame.Retry = n
				}
			}
		}
	}
	if err := s.scanner.Err(); err != nil {
		return Frame{}, err
	}
	// EOF: dispatch a final event when the body ended without a trailing blank line.
	if haveAny {
		return finishFrame(&frame, &dataBuf, haveData), nil
	}
	return Frame{}, io.EOF
}

func finishFrame(frame *Frame, dataBuf *bytes.Buffer, haveData bool) Frame {
	if haveData {
		d := dataBuf.Bytes()
		if n := len(d); n > 0 && d[n-1] == '\n' {
			d = d[:n-1] // strip the single trailing newline added during accumulation
		}
		frame.Data = append([]byte(nil), d...)
	}
	return *frame
}

func splitSSEField(line string) (field, value string) {
	if i := strings.IndexByte(line, ':'); i >= 0 {
		field = line[:i]
		value = strings.TrimPrefix(line[i+1:], " ")
		return field, value
	}
	return line, ""
}

func isASCIIDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// scanSSELines is a bufio.SplitFunc that splits on the SSE line terminators
// CRLF, LF, and a lone CR, none of which are included in the returned token.
func scanSSELines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '\n':
			return i + 1, data[:i], nil
		case '\r':
			if i+1 < len(data) {
				if data[i+1] == '\n' {
					return i + 2, data[:i], nil
				}
				return i + 1, data[:i], nil
			}
			if atEOF {
				return i + 1, data[:i], nil
			}
			// A trailing CR might be the first half of a CRLF split across reads.
			return 0, nil, nil
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil // request more data
}
