//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ExtensionsClientListByFarmBeatsPager provides operations for iterating over paged responses.
type ExtensionsClientListByFarmBeatsPager struct {
	client *ExtensionsClient
	current ExtensionsClientListByFarmBeatsResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, ExtensionsClientListByFarmBeatsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExtensionsClientListByFarmBeatsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExtensionsClientListByFarmBeatsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExtensionListResponse.NextLink == nil || len(*p.current.ExtensionListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByFarmBeatsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExtensionsClientListByFarmBeatsResponse page.
func (p *ExtensionsClientListByFarmBeatsPager) PageResponse() ExtensionsClientListByFarmBeatsResponse {
	return p.current
}

// FarmBeatsExtensionsClientListPager provides operations for iterating over paged responses.
type FarmBeatsExtensionsClientListPager struct {
	client *FarmBeatsExtensionsClient
	current FarmBeatsExtensionsClientListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, FarmBeatsExtensionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsExtensionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsExtensionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsExtensionListResponse.NextLink == nil || len(*p.current.FarmBeatsExtensionListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsExtensionsClientListResponse page.
func (p *FarmBeatsExtensionsClientListPager) PageResponse() FarmBeatsExtensionsClientListResponse {
	return p.current
}

// FarmBeatsModelsClientListByResourceGroupPager provides operations for iterating over paged responses.
type FarmBeatsModelsClientListByResourceGroupPager struct {
	client *FarmBeatsModelsClient
	current FarmBeatsModelsClientListByResourceGroupResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, FarmBeatsModelsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsModelsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsModelsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsListResponse.NextLink == nil || len(*p.current.FarmBeatsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsModelsClientListByResourceGroupResponse page.
func (p *FarmBeatsModelsClientListByResourceGroupPager) PageResponse() FarmBeatsModelsClientListByResourceGroupResponse {
	return p.current
}

// FarmBeatsModelsClientListBySubscriptionPager provides operations for iterating over paged responses.
type FarmBeatsModelsClientListBySubscriptionPager struct {
	client *FarmBeatsModelsClient
	current FarmBeatsModelsClientListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, FarmBeatsModelsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsModelsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsModelsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsListResponse.NextLink == nil || len(*p.current.FarmBeatsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsModelsClientListBySubscriptionResponse page.
func (p *FarmBeatsModelsClientListBySubscriptionPager) PageResponse() FarmBeatsModelsClientListBySubscriptionResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client *OperationsClient
	current OperationsClientListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

