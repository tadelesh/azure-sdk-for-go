//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExperimentsListByProfilePager provides operations for iterating over paged responses.
type ExperimentsListByProfilePager struct {
	client    *ExperimentsClient
	current   ExperimentsListByProfileResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExperimentsListByProfileResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExperimentsListByProfilePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExperimentsListByProfilePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExperimentList.NextLink == nil || len(*p.current.ExperimentList.NextLink) == 0 {
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
		p.err = p.client.listByProfileHandleError(resp)
		return false
	}
	result, err := p.client.listByProfileHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExperimentsListByProfileResponse page.
func (p *ExperimentsListByProfilePager) PageResponse() ExperimentsListByProfileResponse {
	return p.current
}

// FrontDoorsListByResourceGroupPager provides operations for iterating over paged responses.
type FrontDoorsListByResourceGroupPager struct {
	client    *FrontDoorsClient
	current   FrontDoorsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FrontDoorsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FrontDoorsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FrontDoorsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FrontDoorListResult.NextLink == nil || len(*p.current.FrontDoorListResult.NextLink) == 0 {
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
		p.err = p.client.listByResourceGroupHandleError(resp)
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

// PageResponse returns the current FrontDoorsListByResourceGroupResponse page.
func (p *FrontDoorsListByResourceGroupPager) PageResponse() FrontDoorsListByResourceGroupResponse {
	return p.current
}

// FrontDoorsListPager provides operations for iterating over paged responses.
type FrontDoorsListPager struct {
	client    *FrontDoorsClient
	current   FrontDoorsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FrontDoorsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FrontDoorsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FrontDoorsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FrontDoorListResult.NextLink == nil || len(*p.current.FrontDoorListResult.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current FrontDoorsListResponse page.
func (p *FrontDoorsListPager) PageResponse() FrontDoorsListResponse {
	return p.current
}

// FrontendEndpointsListByFrontDoorPager provides operations for iterating over paged responses.
type FrontendEndpointsListByFrontDoorPager struct {
	client    *FrontendEndpointsClient
	current   FrontendEndpointsListByFrontDoorResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FrontendEndpointsListByFrontDoorResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FrontendEndpointsListByFrontDoorPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FrontendEndpointsListByFrontDoorPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FrontendEndpointsListResult.NextLink == nil || len(*p.current.FrontendEndpointsListResult.NextLink) == 0 {
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
		p.err = p.client.listByFrontDoorHandleError(resp)
		return false
	}
	result, err := p.client.listByFrontDoorHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FrontendEndpointsListByFrontDoorResponse page.
func (p *FrontendEndpointsListByFrontDoorPager) PageResponse() FrontendEndpointsListByFrontDoorResponse {
	return p.current
}

// ManagedRuleSetsListPager provides operations for iterating over paged responses.
type ManagedRuleSetsListPager struct {
	client    *ManagedRuleSetsClient
	current   ManagedRuleSetsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedRuleSetsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedRuleSetsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedRuleSetsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedRuleSetDefinitionList.NextLink == nil || len(*p.current.ManagedRuleSetDefinitionList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current ManagedRuleSetsListResponse page.
func (p *ManagedRuleSetsListPager) PageResponse() ManagedRuleSetsListResponse {
	return p.current
}

// NetworkExperimentProfilesListByResourceGroupPager provides operations for iterating over paged responses.
type NetworkExperimentProfilesListByResourceGroupPager struct {
	client    *NetworkExperimentProfilesClient
	current   NetworkExperimentProfilesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NetworkExperimentProfilesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NetworkExperimentProfilesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NetworkExperimentProfilesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProfileList.NextLink == nil || len(*p.current.ProfileList.NextLink) == 0 {
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
		p.err = p.client.listByResourceGroupHandleError(resp)
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

// PageResponse returns the current NetworkExperimentProfilesListByResourceGroupResponse page.
func (p *NetworkExperimentProfilesListByResourceGroupPager) PageResponse() NetworkExperimentProfilesListByResourceGroupResponse {
	return p.current
}

// NetworkExperimentProfilesListPager provides operations for iterating over paged responses.
type NetworkExperimentProfilesListPager struct {
	client    *NetworkExperimentProfilesClient
	current   NetworkExperimentProfilesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NetworkExperimentProfilesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NetworkExperimentProfilesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NetworkExperimentProfilesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProfileList.NextLink == nil || len(*p.current.ProfileList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current NetworkExperimentProfilesListResponse page.
func (p *NetworkExperimentProfilesListPager) PageResponse() NetworkExperimentProfilesListResponse {
	return p.current
}

// PoliciesListPager provides operations for iterating over paged responses.
type PoliciesListPager struct {
	client    *PoliciesClient
	current   PoliciesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PoliciesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PoliciesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PoliciesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WebApplicationFirewallPolicyList.NextLink == nil || len(*p.current.WebApplicationFirewallPolicyList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current PoliciesListResponse page.
func (p *PoliciesListPager) PageResponse() PoliciesListResponse {
	return p.current
}

// PreconfiguredEndpointsListPager provides operations for iterating over paged responses.
type PreconfiguredEndpointsListPager struct {
	client    *PreconfiguredEndpointsClient
	current   PreconfiguredEndpointsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PreconfiguredEndpointsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PreconfiguredEndpointsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PreconfiguredEndpointsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PreconfiguredEndpointList.NextLink == nil || len(*p.current.PreconfiguredEndpointList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current PreconfiguredEndpointsListResponse page.
func (p *PreconfiguredEndpointsListPager) PageResponse() PreconfiguredEndpointsListResponse {
	return p.current
}

// RulesEnginesListByFrontDoorPager provides operations for iterating over paged responses.
type RulesEnginesListByFrontDoorPager struct {
	client    *RulesEnginesClient
	current   RulesEnginesListByFrontDoorResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RulesEnginesListByFrontDoorResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RulesEnginesListByFrontDoorPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RulesEnginesListByFrontDoorPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RulesEngineListResult.NextLink == nil || len(*p.current.RulesEngineListResult.NextLink) == 0 {
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
		p.err = p.client.listByFrontDoorHandleError(resp)
		return false
	}
	result, err := p.client.listByFrontDoorHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RulesEnginesListByFrontDoorResponse page.
func (p *RulesEnginesListByFrontDoorPager) PageResponse() RulesEnginesListByFrontDoorResponse {
	return p.current
}
