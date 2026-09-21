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

// EventFrame is a single Server-Sent Event frame parsed from a text/event-stream
// body. It exposes only the wire-level SSE envelope; the strongly-typed payload
// is produced by a generated per-stream decoder that consumes a EventFrame.
type EventFrame struct {
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

// EventReader provides typed, forward-only iteration over a Server-Sent Events
// response body. T is the generated event union for the operation.
//
// The stream is bound to the context of the request that opened it: canceling
// that context ends the stream and fails in-progress reads. Call Close to release
// the stream early.
//
// The zero value is a valid, already-exhausted stream: iteration yields no
// events and Close is a no-op.
type EventReader[T any] struct {
	body      io.ReadCloser
	scanner   *sseScanner
	decode    func(EventFrame) (T, bool, error)
	connect   func(ctx context.Context, lastEventID string) (*http.Response, error)
	ctx       context.Context // governs the stream lifetime; used for reconnect
	reconnect bool
	done      bool
	lastID    string
	retry     int
	segFrames int // events delivered since the current segment connected
}

// EventHandler supplies the typed decoder and the connection factory used
// to reopen an SSE stream after an unexpected disconnect.
type EventHandler[T any] struct {
	// Decode maps a wire-level EventFrame to the typed union value; it returns
	// terminal=true when the event signals the end of the stream.
	Decode func(frame EventFrame) (value T, terminal bool, err error)

	// Connect reopens the stream after an unexpected disconnect; the initial
	// connection is made by the caller, not through this factory. lastEventID
	// carries the most recent event id so the server can resume from it.
	Connect func(ctx context.Context, lastEventID string) (*http.Response, error)

	// Reconnect enables transparent reconnection after an unexpected mid-stream
	// disconnect. Reconnection resumes from the last seen event id, so it only
	// takes effect once the stream has emitted an id-bearing frame: a stream
	// whose frames carry no id cannot be resumed and is treated as complete at a
	// clean end of body rather than replayed. It must remain false for streams
	// that signal completion with a clean end of body rather than a terminal event.
	Reconnect bool
}

// EventReaderOptions contains the optional values when constructing an EventReader.
type EventReaderOptions struct {
	// for future expansion
}

// NewEventReader wraps an already-open Server-Sent Events response in a typed
// reader. The caller makes the initial connection so it can inspect the response
// (status, headers) first; handler.Connect is used only to reopen the stream on
// an unexpected mid-stream disconnect. The response's request context governs the
// stream lifetime, including reconnect attempts made by Next; canceling it aborts
// an in-progress reconnect and ends the stream.
func NewEventReader[T any](resp *http.Response, handler EventHandler[T], _ *EventReaderOptions) (*EventReader[T], error) {
	if resp == nil || resp.Body == nil {
		return nil, errors.New("streaming: response and its body must not be nil")
	}
	ctx := context.Background()
	if resp.Request != nil {
		ctx = resp.Request.Context()
	}
	return &EventReader[T]{
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
// EventReader was created with reconnect support, an unexpected end of the current
// segment triggers a reconnect (honoring the server retry delay, bound by the
// stream's context) before Next reports io.EOF.
func (e *EventReader[T]) Next() (T, error) {
	var zero T
	// a nil scanner is an empty stream: nothing to consume.
	if e.done || e.scanner == nil {
		return zero, io.EOF
	}
	for {
		frame, err := e.scanner.next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// Reconnect only from a resumable position: SSE resumes via
				// Last-Event-ID, so a stream that never emitted an event id cannot
				// be resumed and reconnecting would only replay it indefinitely.
				// The segFrames guard stops a resumed-but-empty segment from looping.
				if e.reconnect && e.lastID != "" && e.segFrames > 0 {
					if rerr := e.reopen(); rerr != nil {
						e.done = true
						return zero, rerr
					}
					continue
				}
				e.done = true
			}
			return zero, err
		}
		e.lastID = frame.ID
		if frame.Retry >= 0 {
			e.retry = frame.Retry
		}
		value, terminal, derr := e.decode(frame)
		if derr != nil {
			return zero, derr
		}
		if terminal {
			e.done = true
			return zero, io.EOF
		}
		e.segFrames++
		return value, nil
	}
}

// reopen waits the server-suggested retry delay, then reconnects via the
// factory and swaps in the new segment's body, forwarding the last event id.
// It is bound by the stream's context.
func (e *EventReader[T]) reopen() error {
	if e.retry > 0 {
		timer := time.NewTimer(time.Duration(e.retry) * time.Millisecond)
		defer timer.Stop()
		select {
		case <-e.ctx.Done():
			return e.ctx.Err()
		case <-timer.C:
		}
	}
	resp, err := e.connect(e.ctx, e.lastID)
	if err != nil {
		return err
	}
	if e.body != nil {
		e.body.Close()
	}
	e.body = resp.Body
	e.scanner = newSSEScanner(resp.Body)
	e.segFrames = 0
	return nil
}

// Events returns a range-over-func iterator over the stream. Iteration ends at
// end of stream (io.EOF is not yielded) or after the first error is yielded.
func (e *EventReader[T]) Events() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			value, err := e.Next()
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
func (e *EventReader[T]) LastEventID() string { return e.lastID }

// Close closes the underlying response body, if any.
func (e *EventReader[T]) Close() error {
	if e.body == nil {
		return nil
	}
	return e.body.Close()
}

// sseScanner turns a stream of SSE lines into discrete EventFrame values following
// the WHATWG event stream parsing rulee.
type sseScanner struct {
	scanner *bufio.Scanner
	lastID  string
}

func (e *sseScanner) next() (EventFrame, error) {
	var (
		frame    EventFrame
		dataBuf  bytes.Buffer
		haveData bool
		haveAny  bool
	)
	frame.Retry = -1
	frame.ID = e.lastID
	for e.scanner.Scan() {
		line := e.scanner.Text()
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
				e.lastID = value
			}
		case "retry":
			if isASCIIDigits(value) {
				if n, err := strconv.Atoi(value); err == nil {
					frame.Retry = n
				}
			}
		}
	}
	if err := e.scanner.Err(); err != nil {
		return EventFrame{}, err
	}
	// EOF: dispatch a final event when the body ended without a trailing blank line.
	if haveAny {
		return finishFrame(&frame, &dataBuf, haveData), nil
	}
	return EventFrame{}, io.EOF
}

func finishFrame(frame *EventFrame, dataBuf *bytes.Buffer, haveData bool) EventFrame {
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
			// A trailing CR might be the first half of a CRLF split across reade.
			return 0, nil, nil
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil // request more data
}
