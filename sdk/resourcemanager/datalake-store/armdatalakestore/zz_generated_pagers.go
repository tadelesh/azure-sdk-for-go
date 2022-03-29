//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsClientListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
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

// PageResponse returns the current AccountsClientListByResourceGroupResponse page.
func (p *AccountsClientListByResourceGroupPager) PageResponse() AccountsClientListByResourceGroupResponse {
	return p.current
}

// AccountsClientListPager provides operations for iterating over paged responses.
type AccountsClientListPager struct {
	client    *AccountsClient
	current   AccountsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
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

// PageResponse returns the current AccountsClientListResponse page.
func (p *AccountsClientListPager) PageResponse() AccountsClientListResponse {
	return p.current
}

// FirewallRulesClientListByAccountPager provides operations for iterating over paged responses.
type FirewallRulesClientListByAccountPager struct {
	client    *FirewallRulesClient
	current   FirewallRulesClientListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FirewallRulesClientListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FirewallRulesClientListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FirewallRulesClientListByAccountPager) NextPage(ctx context.Context) bool {
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
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FirewallRulesClientListByAccountResponse page.
func (p *FirewallRulesClientListByAccountPager) PageResponse() FirewallRulesClientListByAccountResponse {
	return p.current
}

// TrustedIDProvidersClientListByAccountPager provides operations for iterating over paged responses.
type TrustedIDProvidersClientListByAccountPager struct {
	client    *TrustedIDProvidersClient
	current   TrustedIDProvidersClientListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TrustedIDProvidersClientListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TrustedIDProvidersClientListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TrustedIDProvidersClientListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TrustedIDProviderListResult.NextLink == nil || len(*p.current.TrustedIDProviderListResult.NextLink) == 0 {
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
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TrustedIDProvidersClientListByAccountResponse page.
func (p *TrustedIDProvidersClientListByAccountPager) PageResponse() TrustedIDProvidersClientListByAccountResponse {
	return p.current
}

// VirtualNetworkRulesClientListByAccountPager provides operations for iterating over paged responses.
type VirtualNetworkRulesClientListByAccountPager struct {
	client    *VirtualNetworkRulesClient
	current   VirtualNetworkRulesClientListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworkRulesClientListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualNetworkRulesClientListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualNetworkRulesClientListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
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
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VirtualNetworkRulesClientListByAccountResponse page.
func (p *VirtualNetworkRulesClientListByAccountPager) PageResponse() VirtualNetworkRulesClientListByAccountResponse {
	return p.current
}