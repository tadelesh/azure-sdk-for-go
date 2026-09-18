// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package streaming_test

import (
	"encoding/json"
	"errors"
	"io"
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
	fields    map[string]any
}

// decodeView is a generic decoder: JSON object payloads are unmarshaled into
// fields, the "[DONE]" sentinel is treated as terminal, and everything else is
// surfaced as raw text.
func decodeView(f streaming.Frame) (frameView, bool, error) {
	v := frameView{eventType: f.Type, data: string(f.Data)}
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

func collect[T any](t *testing.T, body string, decode func(streaming.Frame) (T, bool, error)) ([]T, *streaming.Event[T]) {
	t.Helper()
	s := streaming.NewEvent(io.NopCloser(strings.NewReader(body)), decode)
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
	_, s := collect(t, body, decodeView)
	if s.RetryAfter() != 1000 {
		t.Fatalf("RetryAfter = %d, want 1000", s.RetryAfter())
	}
}

func TestProtocolRetryInvalidIgnored(t *testing.T) {
	body := "retry: not-a-number\nevent: message\ndata: {\"message\": \"hello\"}\n\n"
	_, s := collect(t, body, decodeView)
	if s.RetryAfter() != -1 {
		t.Fatalf("RetryAfter = %d, want -1 for invalid retry", s.RetryAfter())
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
	s := streaming.NewEvent(io.NopCloser(strings.NewReader(body)), decodeView)
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
