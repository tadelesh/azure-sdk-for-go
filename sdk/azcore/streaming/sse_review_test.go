// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package streaming_test

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"testing/synctest"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func reviewReader[T any](t *testing.T, body io.ReadCloser, decode func(streaming.EventFrame) (T, bool, error)) *streaming.EventReader[T] {
	t.Helper()
	reader, err := streaming.NewEventReader(context.Background(), func(context.Context, string) (io.ReadCloser, error) {
		return body, nil
	}, streaming.EventHandler[T]{Decode: decode}, nil)
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, reader.Close()) })
	return reader
}

func TestReviewNamedTerminalPayloadIsDelivered(t *testing.T) {
	type completion struct {
		Total int `json:"total"`
	}
	body := io.NopCloser(strings.NewReader("event: complete\ndata: {\"total\":7}\n\n"))
	reader := reviewReader(t, body, func(frame streaming.EventFrame) (completion, bool, error) {
		var value completion
		err := json.Unmarshal(frame.Data, &value)
		return value, frame.Type == "complete", err
	})
	// A terminal model carries useful data that should be delivered once before EOF.
	// The current reader discards that data as soon as Decode reports terminal=true.
	event, err := reader.Next()
	require.NoError(t, err, "a terminal model event must be delivered before EOF")
	assert.Equal(t, 7, event.Total)
	_, err = reader.Next()
	require.ErrorIs(t, err, io.EOF)
}

func TestReviewLoneCRDispatchesWithoutAnotherRead(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		read, write := io.Pipe()
		defer write.Close()
		reader := reviewReader(t, read, func(frame streaming.EventFrame) (string, bool, error) {
			return string(frame.Data), false, nil
		})
		defer reader.Close()
		type result struct {
			event string
			err   error
		}
		next := make(chan result, 1)
		go func() {
			event, err := reader.Next()
			next <- result{event: event, err: err}
		}()
		// CR (\r, carriage return) is a valid SSE line ending: these two CRs end the
		// data line and its terminating blank line. Keep the pipe open to expose the
		// splitter waiting for a possible LF (\n), rather than letting EOF hide it.
		_, err := io.WriteString(write, "data: hello\r\r")
		require.NoError(t, err)
		// Synchronize with Next so the nonblocking result check is deterministic.
		synctest.Wait()
		select {
		case got := <-next:
			require.NoError(t, got.err)
			assert.Equal(t, "hello", got.event)
		default:
			assert.Fail(t, "a complete CR-delimited event is blocked until another byte arrives")
			require.NoError(t, reader.Close())
			<-next
		}
	})
}
