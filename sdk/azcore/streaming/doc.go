// Copyright 2017 Microsoft Corporation. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

// Package streaming contains helpers for streaming IO operations and progress reporting.
//
// NewEventReader eagerly opens an SSE connection through an EventConnector and
// reads a single response by default. Continuous operations can opt into
// reconnection with EventHandler.Reconnect; the same connector opens each fresh
// connection through the client pipeline and returns its body. A connector
// io.EOF signals service completion, such as HTTP 204, including an initially
// exhausted reader. SSE reconnection delays
// govern when to open a new stream; the pipeline's retry policy governs attempts
// to establish that connection. A final callback error stops the reader.
//
// EventReader.Next blocks waiting for a complete data event and has no built-in
// idle timeout. An open stream can wait indefinitely, including when it receives
// only heartbeat comments or metadata. Call Close to interrupt the wait, or pass
// an operation context with cancellation or a deadline to NewEventReader.
//
// Pass the original operation context to NewEventReader, not the response's
// per-attempt context. The connector must use the context passed to it for each
// request. An empty EventReaderOptions.LastEventID means no checkpoint; the
// service determines the starting position. Always close the reader when
// finished, including when stopping early.
package streaming
