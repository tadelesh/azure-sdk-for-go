//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DNSForwardingRulesetsClientListByResourceGroupPager provides operations for iterating over paged responses.
type DNSForwardingRulesetsClientListByResourceGroupPager struct {
	client    *DNSForwardingRulesetsClient
	current   DNSForwardingRulesetsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSForwardingRulesetsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSForwardingRulesetsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSForwardingRulesetsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DNSForwardingRulesetListResult.NextLink == nil || len(*p.current.DNSForwardingRulesetListResult.NextLink) == 0 {
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

// PageResponse returns the current DNSForwardingRulesetsClientListByResourceGroupResponse page.
func (p *DNSForwardingRulesetsClientListByResourceGroupPager) PageResponse() DNSForwardingRulesetsClientListByResourceGroupResponse {
	return p.current
}

// DNSForwardingRulesetsClientListByVirtualNetworkPager provides operations for iterating over paged responses.
type DNSForwardingRulesetsClientListByVirtualNetworkPager struct {
	client    *DNSForwardingRulesetsClient
	current   DNSForwardingRulesetsClientListByVirtualNetworkResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSForwardingRulesetsClientListByVirtualNetworkResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSForwardingRulesetsClientListByVirtualNetworkPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSForwardingRulesetsClientListByVirtualNetworkPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkDNSForwardingRulesetListResult.NextLink == nil || len(*p.current.VirtualNetworkDNSForwardingRulesetListResult.NextLink) == 0 {
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
	result, err := p.client.listByVirtualNetworkHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DNSForwardingRulesetsClientListByVirtualNetworkResponse page.
func (p *DNSForwardingRulesetsClientListByVirtualNetworkPager) PageResponse() DNSForwardingRulesetsClientListByVirtualNetworkResponse {
	return p.current
}

// DNSForwardingRulesetsClientListPager provides operations for iterating over paged responses.
type DNSForwardingRulesetsClientListPager struct {
	client    *DNSForwardingRulesetsClient
	current   DNSForwardingRulesetsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSForwardingRulesetsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSForwardingRulesetsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSForwardingRulesetsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DNSForwardingRulesetListResult.NextLink == nil || len(*p.current.DNSForwardingRulesetListResult.NextLink) == 0 {
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

// PageResponse returns the current DNSForwardingRulesetsClientListResponse page.
func (p *DNSForwardingRulesetsClientListPager) PageResponse() DNSForwardingRulesetsClientListResponse {
	return p.current
}

// DNSResolversClientListByResourceGroupPager provides operations for iterating over paged responses.
type DNSResolversClientListByResourceGroupPager struct {
	client    *DNSResolversClient
	current   DNSResolversClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSResolversClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSResolversClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSResolversClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
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

// PageResponse returns the current DNSResolversClientListByResourceGroupResponse page.
func (p *DNSResolversClientListByResourceGroupPager) PageResponse() DNSResolversClientListByResourceGroupResponse {
	return p.current
}

// DNSResolversClientListByVirtualNetworkPager provides operations for iterating over paged responses.
type DNSResolversClientListByVirtualNetworkPager struct {
	client    *DNSResolversClient
	current   DNSResolversClientListByVirtualNetworkResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSResolversClientListByVirtualNetworkResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSResolversClientListByVirtualNetworkPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSResolversClientListByVirtualNetworkPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SubResourceListResult.NextLink == nil || len(*p.current.SubResourceListResult.NextLink) == 0 {
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
	result, err := p.client.listByVirtualNetworkHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DNSResolversClientListByVirtualNetworkResponse page.
func (p *DNSResolversClientListByVirtualNetworkPager) PageResponse() DNSResolversClientListByVirtualNetworkResponse {
	return p.current
}

// DNSResolversClientListPager provides operations for iterating over paged responses.
type DNSResolversClientListPager struct {
	client    *DNSResolversClient
	current   DNSResolversClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DNSResolversClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DNSResolversClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DNSResolversClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
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

// PageResponse returns the current DNSResolversClientListResponse page.
func (p *DNSResolversClientListPager) PageResponse() DNSResolversClientListResponse {
	return p.current
}

// ForwardingRulesClientListPager provides operations for iterating over paged responses.
type ForwardingRulesClientListPager struct {
	client    *ForwardingRulesClient
	current   ForwardingRulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ForwardingRulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ForwardingRulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ForwardingRulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ForwardingRuleListResult.NextLink == nil || len(*p.current.ForwardingRuleListResult.NextLink) == 0 {
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

// PageResponse returns the current ForwardingRulesClientListResponse page.
func (p *ForwardingRulesClientListPager) PageResponse() ForwardingRulesClientListResponse {
	return p.current
}

// InboundEndpointsClientListPager provides operations for iterating over paged responses.
type InboundEndpointsClientListPager struct {
	client    *InboundEndpointsClient
	current   InboundEndpointsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, InboundEndpointsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *InboundEndpointsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *InboundEndpointsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.InboundEndpointListResult.NextLink == nil || len(*p.current.InboundEndpointListResult.NextLink) == 0 {
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

// PageResponse returns the current InboundEndpointsClientListResponse page.
func (p *InboundEndpointsClientListPager) PageResponse() InboundEndpointsClientListResponse {
	return p.current
}

// OutboundEndpointsClientListPager provides operations for iterating over paged responses.
type OutboundEndpointsClientListPager struct {
	client    *OutboundEndpointsClient
	current   OutboundEndpointsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OutboundEndpointsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OutboundEndpointsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OutboundEndpointsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEndpointListResult.NextLink == nil || len(*p.current.OutboundEndpointListResult.NextLink) == 0 {
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

// PageResponse returns the current OutboundEndpointsClientListResponse page.
func (p *OutboundEndpointsClientListPager) PageResponse() OutboundEndpointsClientListResponse {
	return p.current
}

// VirtualNetworkLinksClientListPager provides operations for iterating over paged responses.
type VirtualNetworkLinksClientListPager struct {
	client    *VirtualNetworkLinksClient
	current   VirtualNetworkLinksClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworkLinksClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualNetworkLinksClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualNetworkLinksClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkLinkListResult.NextLink == nil || len(*p.current.VirtualNetworkLinkListResult.NextLink) == 0 {
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

// PageResponse returns the current VirtualNetworkLinksClientListResponse page.
func (p *VirtualNetworkLinksClientListPager) PageResponse() VirtualNetworkLinksClientListResponse {
	return p.current
}