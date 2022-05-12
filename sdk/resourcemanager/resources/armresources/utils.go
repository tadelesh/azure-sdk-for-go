//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armresources

import (
	"context"
	"errors"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type OperationWaiter[TResult any] struct {
	poller *armruntime.Poller[TResult]
	err    error
}

func (ow OperationWaiter[TResult]) Wait(ctx context.Context, freq time.Duration) (TResult, error) {
	if ow.err != nil {
		return *new(TResult), ow.err
	}
	return ow.poller.PollUntilDone(ctx, freq)
}

func NewOperationWaiter[TResult any](poller *armruntime.Poller[TResult], err error) OperationWaiter[TResult] {
	return OperationWaiter[TResult]{poller: poller, err: err}
}

type PageConstraint[TItem any] interface {
	Items() []*TItem
}

type Iterator[TItem any, TPage PageConstraint[TItem]] struct {
	pager *runtime.Pager[TPage]
	cur   []*TItem
	index int
}

func (iter *Iterator[TItem, TPage]) More() bool {
	return iter.pager.More() || iter.index < len(iter.cur)
}

func (iter *Iterator[TItem, TPage]) NextItem(ctx context.Context) (*TItem, error) {
	if iter.index == len(iter.cur) && !iter.pager.More() {
		return nil, errors.New("no more items")
	}
	if iter.cur == nil || iter.index == len(iter.cur) {
		// first page or page exhausted
		page, err := iter.pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		iter.cur = page.Items()
		iter.index = 0
	}
	item := iter.cur[iter.index]
	// advance item
	iter.index++
	return item, nil
}

func NewIterator[TItem any, TPage PageConstraint[TItem]](pager *runtime.Pager[TPage]) *Iterator[TItem, TPage] {
	return &Iterator[TItem, TPage]{
		pager: pager,
	}
}
