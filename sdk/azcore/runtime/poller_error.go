//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

type PollingError[T any] struct {
	inner  error
	Poller *Poller[T]
}

func (e *PollingError[T]) Error() string {
	return e.inner.Error()
}

func (e *PollingError[T]) Unwrap() error {
	return e.inner
}

func NewPollingError[T any](err error, poller *Poller[T]) error {
	return &PollingError[T]{
		inner:  err,
		Poller: poller,
	}
}
