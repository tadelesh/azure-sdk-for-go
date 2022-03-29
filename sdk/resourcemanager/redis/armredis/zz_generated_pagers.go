//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClientListByResourceGroupPager provides operations for iterating over paged responses.
type ClientListByResourceGroupPager struct {
	client    *Client
	current   ClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
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

// PageResponse returns the current ClientListByResourceGroupResponse page.
func (p *ClientListByResourceGroupPager) PageResponse() ClientListByResourceGroupResponse {
	return p.current
}

// ClientListBySubscriptionPager provides operations for iterating over paged responses.
type ClientListBySubscriptionPager struct {
	client    *Client
	current   ClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
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
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientListBySubscriptionResponse page.
func (p *ClientListBySubscriptionPager) PageResponse() ClientListBySubscriptionResponse {
	return p.current
}

// ClientListUpgradeNotificationsPager provides operations for iterating over paged responses.
type ClientListUpgradeNotificationsPager struct {
	client    *Client
	current   ClientListUpgradeNotificationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientListUpgradeNotificationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientListUpgradeNotificationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientListUpgradeNotificationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotificationListResponse.NextLink == nil || len(*p.current.NotificationListResponse.NextLink) == 0 {
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
	result, err := p.client.listUpgradeNotificationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientListUpgradeNotificationsResponse page.
func (p *ClientListUpgradeNotificationsPager) PageResponse() ClientListUpgradeNotificationsResponse {
	return p.current
}

// FirewallRulesClientListPager provides operations for iterating over paged responses.
type FirewallRulesClientListPager struct {
	client    *FirewallRulesClient
	current   FirewallRulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FirewallRulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FirewallRulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FirewallRulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FirewallRuleListResult.NextLink == nil || len(*p.current.FirewallRuleListResult.NextLink) == 0 {
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

// PageResponse returns the current FirewallRulesClientListResponse page.
func (p *FirewallRulesClientListPager) PageResponse() FirewallRulesClientListResponse {
	return p.current
}

// LinkedServerClientListPager provides operations for iterating over paged responses.
type LinkedServerClientListPager struct {
	client    *LinkedServerClient
	current   LinkedServerClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LinkedServerClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LinkedServerClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LinkedServerClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedServerWithPropertiesList.NextLink == nil || len(*p.current.LinkedServerWithPropertiesList.NextLink) == 0 {
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

// PageResponse returns the current LinkedServerClientListResponse page.
func (p *LinkedServerClientListPager) PageResponse() LinkedServerClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
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

// PatchSchedulesClientListByRedisResourcePager provides operations for iterating over paged responses.
type PatchSchedulesClientListByRedisResourcePager struct {
	client    *PatchSchedulesClient
	current   PatchSchedulesClientListByRedisResourceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PatchSchedulesClientListByRedisResourceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PatchSchedulesClientListByRedisResourcePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PatchSchedulesClientListByRedisResourcePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PatchScheduleListResult.NextLink == nil || len(*p.current.PatchScheduleListResult.NextLink) == 0 {
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
	result, err := p.client.listByRedisResourceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PatchSchedulesClientListByRedisResourceResponse page.
func (p *PatchSchedulesClientListByRedisResourcePager) PageResponse() PatchSchedulesClientListByRedisResourceResponse {
	return p.current
}