//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsClientListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccountsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccountsClientListByResourceGroupPager) NextPage(ctx context.Context) (AccountsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccountsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccountsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// AccountsClientListPager provides operations for iterating over paged responses.
type AccountsClientListPager struct {
	client    *AccountsClient
	current   AccountsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccountsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccountsClientListPager) NextPage(ctx context.Context) (AccountsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccountsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccountsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccountsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccountsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return AccountsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// FirewallRulesClientListByAccountPager provides operations for iterating over paged responses.
type FirewallRulesClientListByAccountPager struct {
	client    *FirewallRulesClient
	current   FirewallRulesClientListByAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FirewallRulesClientListByAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *FirewallRulesClientListByAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FirewallRuleListResult.NextLink == nil || len(*p.current.FirewallRuleListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *FirewallRulesClientListByAccountPager) NextPage(ctx context.Context) (FirewallRulesClientListByAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return FirewallRulesClientListByAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return FirewallRulesClientListByAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return FirewallRulesClientListByAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return FirewallRulesClientListByAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		return FirewallRulesClientListByAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LocationsClientGetUsagePager provides operations for iterating over paged responses.
type LocationsClientGetUsagePager struct {
	client    *LocationsClient
	current   LocationsClientGetUsageResponse
	requester func(context.Context) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LocationsClientGetUsagePager) More() bool {
	return reflect.ValueOf(p.current).IsZero()
}

// NextPage advances the pager to the next page.
func (p *LocationsClientGetUsagePager) NextPage(ctx context.Context) (LocationsClientGetUsageResponse, error) {
	var req *policy.Request
	var err error
	if !p.More() {
		return LocationsClientGetUsageResponse{}, errors.New("no more pages")
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LocationsClientGetUsageResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LocationsClientGetUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LocationsClientGetUsageResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getUsageHandleResponse(resp)
	if err != nil {
		return LocationsClientGetUsageResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// TrustedIDProvidersClientListByAccountPager provides operations for iterating over paged responses.
type TrustedIDProvidersClientListByAccountPager struct {
	client    *TrustedIDProvidersClient
	current   TrustedIDProvidersClientListByAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TrustedIDProvidersClientListByAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *TrustedIDProvidersClientListByAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TrustedIDProviderListResult.NextLink == nil || len(*p.current.TrustedIDProviderListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *TrustedIDProvidersClientListByAccountPager) NextPage(ctx context.Context) (TrustedIDProvidersClientListByAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return TrustedIDProvidersClientListByAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return TrustedIDProvidersClientListByAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return TrustedIDProvidersClientListByAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return TrustedIDProvidersClientListByAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		return TrustedIDProvidersClientListByAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualNetworkRulesClientListByAccountPager provides operations for iterating over paged responses.
type VirtualNetworkRulesClientListByAccountPager struct {
	client    *VirtualNetworkRulesClient
	current   VirtualNetworkRulesClientListByAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworkRulesClientListByAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualNetworkRulesClientListByAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualNetworkRulesClientListByAccountPager) NextPage(ctx context.Context) (VirtualNetworkRulesClientListByAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualNetworkRulesClientListByAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualNetworkRulesClientListByAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualNetworkRulesClientListByAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualNetworkRulesClientListByAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		return VirtualNetworkRulesClientListByAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}
