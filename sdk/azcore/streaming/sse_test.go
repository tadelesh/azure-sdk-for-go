// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package streaming_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
)

// frameView is a decoded SSE frame used only by these tests. It captures the
// wire envelope plus the parsed JSON payload so the runtime can be exercised
// without depending on any generated event union.
type frameView struct {
	eventType string
	data      string
	retry     int
	fields    map[string]any
}

// decodeView is a generic decoder: JSON object payloads are unmarshaled into
// fields, the "[DONE]" sentinel is treated as terminal, and everything else is
// surfaced as raw text.
func decodeView(f streaming.EventFrame) (frameView, bool, error) {
	v := frameView{eventType: f.Type, data: string(f.Data), retry: f.Retry}
	if v.data == "[DONE]" {
		return v, true, nil
	}
	if len(f.Data) > 0 && f.Data[0] == '{' {
		if err := json.Unmarshal(f.Data, &v.fields); err != nil {
			return frameView{}, false, err
		}
	}
	return v, false, nil
}

func bodyConnect(body string) func(context.Context, string) (*http.Response, error) {
	return func(context.Context, string) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body))}, nil
	}
}

func collect[T any](t *testing.T, body string, decode func(streaming.EventFrame) (T, bool, error)) ([]T, *streaming.EventReader[T]) {
	t.Helper()
	s, err := streaming.NewEventReader(context.Background(), streaming.EventHandler[T]{
		Decode:  decode,
		Connect: bodyConnect(body),
	}, nil)
	if err != nil {
		t.Fatalf("NewEventReader: %v", err)
	}
	var got []T
	for {
		v, err := s.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got = append(got, v)
	}
	return got, s
}

func TestUnnamedEvents(t *testing.T) {
	body := "data: {\"desc\": \"one\"}\n\ndata: {\"desc\": \"two\"}\n\ndata: {\"desc\": \"three\"}\n\n"
	got, _ := collect(t, body, decodeView)
	want := []string{"one", "two", "three"}
	if len(got) != len(want) {
		t.Fatalf("got %d events, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i].fields["desc"] != w {
			t.Fatalf("event %d = %+v, want desc %q", i, got[i], w)
		}
	}
}

func TestNamedEventsWithTerminal(t *testing.T) {
	body := "event: responseCreated\ndata: {\"id\": \"resp_1\"}\n\n" +
		"event: responseDelta\ndata: {\"delta\": \"Hello\"}\n\n" +
		"event: responseDelta\ndata: {\"delta\": \" world\"}\n\n" +
		"data: [DONE]\n\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 3 {
		t.Fatalf("got %d events, want 3 (terminal [DONE] ends the stream)", len(got))
	}
	if got[0].eventType != "responseCreated" || got[0].fields["id"] != "resp_1" {
		t.Fatalf("event 0 = %+v", got[0])
	}
	if got[1].fields["delta"] != "Hello" {
		t.Fatalf("event 1 = %+v", got[1])
	}
	if got[2].fields["delta"] != " world" {
		t.Fatalf("event 2 = %+v", got[2])
	}
}

func TestProtocolEnvelopeMetadata(t *testing.T) {
	body := "id: event-1\nevent: message\ndata: {\"message\": \"hello\"}\n\n"
	got, s := collect(t, body, decodeView)
	if len(got) != 1 || got[0].fields["message"] != "hello" {
		t.Fatalf("got %+v", got)
	}
	if s.LastEventID() != "event-1" {
		t.Fatalf("LastEventID = %q, want event-1", s.LastEventID())
	}
}

func TestProtocolRetryValid(t *testing.T) {
	body := "retry: 1000\nevent: message\ndata: {\"message\": \"hello\"}\n\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].retry != 1000 {
		t.Fatalf("retry = %d, want 1000", got[0].retry)
	}
}

func TestProtocolRetryInvalidIgnored(t *testing.T) {
	body := "retry: not-a-number\nevent: message\ndata: {\"message\": \"hello\"}\n\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].retry != -1 {
		t.Fatalf("retry = %d, want -1 for invalid retry", got[0].retry)
	}
}

func TestProtocolInvalidIDIgnored(t *testing.T) {
	body := "id: invalid\x00id\nevent: message\ndata: {\"message\": \"hello\"}\n\n"
	_, s := collect(t, body, decodeView)
	if s.LastEventID() != "" {
		t.Fatalf("LastEventID = %q, want empty for id containing NUL", s.LastEventID())
	}
}

func TestDataWithEnvelopeNoTrailingBlank(t *testing.T) {
	body := "event: withEnvelope\ndata: hello"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].data != "hello" {
		t.Fatalf("got %+v", got)
	}
}

func TestDataWithoutEnvelope(t *testing.T) {
	body := "event: withoutEnvelope\ndata: {\"metadata\": {\"source\": \"test\"}, \"contents\": \"world\"}\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].fields["contents"] != "world" {
		t.Fatalf("got %+v", got)
	}
	meta, ok := got[0].fields["metadata"].(map[string]any)
	if !ok || meta["source"] != "test" {
		t.Fatalf("metadata = %+v", got[0].fields["metadata"])
	}
}

func TestCRLFLineEndings(t *testing.T) {
	body := "event: responseDelta\r\ndata: {\"delta\": \"crlf\"}\r\n\r\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].fields["delta"] != "crlf" {
		t.Fatalf("got %+v", got)
	}
}

func TestMultilineData(t *testing.T) {
	// Two data lines concatenate with a newline; the JSON stays valid.
	body := "event: responseDelta\ndata: {\"delta\":\ndata: \"multi\"}\n\n"
	got, _ := collect(t, body, decodeView)
	if len(got) != 1 || got[0].fields["delta"] != "multi" {
		t.Fatalf("got %+v", got)
	}
}

func TestEventsIterator(t *testing.T) {
	body := "data: {\"desc\": \"a\"}\n\ndata: {\"desc\": \"b\"}\n\n"
	s, err := streaming.NewEventReader(context.Background(), streaming.EventHandler[frameView]{
		Decode:  decodeView,
		Connect: bodyConnect(body),
	}, nil)
	if err != nil {
		t.Fatalf("NewEventReader: %v", err)
	}
	var descs []string
	for ev, err := range s.Events() {
		if err != nil {
			t.Fatalf("iterator error: %v", err)
		}
		descs = append(descs, ev.fields["desc"].(string))
	}
	if strings.Join(descs, ",") != "a,b" {
		t.Fatalf("descs = %v", descs)
	}
}

func TestNewEventConnectError(t *testing.T) {
	handler := streaming.EventHandler[frameView]{
		Decode: decodeView,
		Connect: func(context.Context, string) (*http.Response, error) {
			return nil, errors.New("boom")
		},
	}
	if _, err := streaming.NewEventReader(context.Background(), handler, nil); err == nil {
		t.Fatal("expected initial connect error")
	}
}

func TestReconnectResumesWithLastEventID(t *testing.T) {
	segments := []string{
		"id: 1\ndata: {\"desc\": \"a\"}\n\nid: 2\ndata: {\"desc\": \"b\"}\n\n",
		"id: 3\ndata: {\"desc\": \"c\"}\n\n",
		"", // an empty segment signals the stream is exhausted
	}
	var lastIDs []string
	call := 0
	handler := streaming.EventHandler[frameView]{
		Decode:    decodeView,
		Reconnect: true,
		Connect: func(_ context.Context, lastEventID string) (*http.Response, error) {
			lastIDs = append(lastIDs, lastEventID)
			body := segments[call]
			call++
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}
	s, err := streaming.NewEventReader(context.Background(), handler, nil)
	if err != nil {
		t.Fatalf("NewEventReader: %v", err)
	}
	var got []string
	for {
		v, err := s.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got = append(got, v.fields["desc"].(string))
	}
	if strings.Join(got, ",") != "a,b,c" {
		t.Fatalf("got %v, want a,b,c", got)
	}
	// initial connect ("") then reconnect after id 2, then after id 3.
	if strings.Join(lastIDs, ",") != ",2,3" {
		t.Fatalf("lastIDs = %v, want [\"\" 2 3]", lastIDs)
	}
}

// A reconnect-enabled stream whose frames carry no event id has no resumable
// position, so it must end at a clean end of body instead of reconnecting and
// replaying the same body forever (the missing-terminal-event footgun).
func TestReconnectWithoutEventIDDoesNotReplay(t *testing.T) {
	const body = "data: {\"desc\": \"a\"}\n\ndata: {\"desc\": \"b\"}\n\n"
	connects := 0
	handler := streaming.EventHandler[frameView]{
		Decode:    decodeView,
		Reconnect: true,
		Connect: func(context.Context, string) (*http.Response, error) {
			connects++
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}
	s, err := streaming.NewEventReader(context.Background(), handler, nil)
	if err != nil {
		t.Fatalf("NewEventReader: %v", err)
	}
	var got []string
	for {
		v, err := s.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got = append(got, v.fields["desc"].(string))
	}
	if strings.Join(got, ",") != "a,b" {
		t.Fatalf("got %v, want a,b", got)
	}
	if connects != 1 {
		t.Fatalf("connects = %d, want 1 (no replay)", connects)
	}
}
