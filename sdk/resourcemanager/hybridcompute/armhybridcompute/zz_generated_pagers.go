//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// MachineExtensionsClientListPager provides operations for iterating over paged responses.
type MachineExtensionsClientListPager struct {
	client    *MachineExtensionsClient
	current   MachineExtensionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MachineExtensionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MachineExtensionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MachineExtensionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MachineExtensionsListResult.NextLink == nil || len(*p.current.MachineExtensionsListResult.NextLink) == 0 {
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

// PageResponse returns the current MachineExtensionsClientListResponse page.
func (p *MachineExtensionsClientListPager) PageResponse() MachineExtensionsClientListResponse {
	return p.current
}

// MachinesClientListByResourceGroupPager provides operations for iterating over paged responses.
type MachinesClientListByResourceGroupPager struct {
	client    *MachinesClient
	current   MachinesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MachinesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MachinesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MachinesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MachineListResult.NextLink == nil || len(*p.current.MachineListResult.NextLink) == 0 {
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

// PageResponse returns the current MachinesClientListByResourceGroupResponse page.
func (p *MachinesClientListByResourceGroupPager) PageResponse() MachinesClientListByResourceGroupResponse {
	return p.current
}

// MachinesClientListBySubscriptionPager provides operations for iterating over paged responses.
type MachinesClientListBySubscriptionPager struct {
	client    *MachinesClient
	current   MachinesClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MachinesClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MachinesClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MachinesClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MachineListResult.NextLink == nil || len(*p.current.MachineListResult.NextLink) == 0 {
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

// PageResponse returns the current MachinesClientListBySubscriptionResponse page.
func (p *MachinesClientListBySubscriptionPager) PageResponse() MachinesClientListBySubscriptionResponse {
	return p.current
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopePager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsClientListByPrivateLinkScopePager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsClientListByPrivateLinkScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsClientListByPrivateLinkScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
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
	result, err := p.client.listByPrivateLinkScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse page.
func (p *PrivateEndpointConnectionsClientListByPrivateLinkScopePager) PageResponse() PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse {
	return p.current
}

// PrivateLinkResourcesClientListByPrivateLinkScopePager provides operations for iterating over paged responses.
type PrivateLinkResourcesClientListByPrivateLinkScopePager struct {
	client    *PrivateLinkResourcesClient
	current   PrivateLinkResourcesClientListByPrivateLinkScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkResourcesClientListByPrivateLinkScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkResourcesClientListByPrivateLinkScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkResourcesClientListByPrivateLinkScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkResourceListResult.NextLink == nil || len(*p.current.PrivateLinkResourceListResult.NextLink) == 0 {
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
	result, err := p.client.listByPrivateLinkScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateLinkResourcesClientListByPrivateLinkScopeResponse page.
func (p *PrivateLinkResourcesClientListByPrivateLinkScopePager) PageResponse() PrivateLinkResourcesClientListByPrivateLinkScopeResponse {
	return p.current
}

// PrivateLinkScopesClientListByResourceGroupPager provides operations for iterating over paged responses.
type PrivateLinkScopesClientListByResourceGroupPager struct {
	client    *PrivateLinkScopesClient
	current   PrivateLinkScopesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkScopesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkScopesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkScopesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkScopeListResult.NextLink == nil || len(*p.current.PrivateLinkScopeListResult.NextLink) == 0 {
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

// PageResponse returns the current PrivateLinkScopesClientListByResourceGroupResponse page.
func (p *PrivateLinkScopesClientListByResourceGroupPager) PageResponse() PrivateLinkScopesClientListByResourceGroupResponse {
	return p.current
}

// PrivateLinkScopesClientListPager provides operations for iterating over paged responses.
type PrivateLinkScopesClientListPager struct {
	client    *PrivateLinkScopesClient
	current   PrivateLinkScopesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkScopesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkScopesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkScopesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkScopeListResult.NextLink == nil || len(*p.current.PrivateLinkScopeListResult.NextLink) == 0 {
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

// PageResponse returns the current PrivateLinkScopesClientListResponse page.
func (p *PrivateLinkScopesClientListPager) PageResponse() PrivateLinkScopesClientListResponse {
	return p.current
}