//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LegacyPeeringsListPager provides operations for iterating over paged responses.
type LegacyPeeringsListPager struct {
	client    *LegacyPeeringsClient
	current   LegacyPeeringsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LegacyPeeringsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LegacyPeeringsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LegacyPeeringsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringListResult.NextLink == nil || len(*p.current.PeeringListResult.NextLink) == 0 {
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

// PageResponse returns the current LegacyPeeringsListResponse page.
func (p *LegacyPeeringsListPager) PageResponse() LegacyPeeringsListResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
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

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// PeerAsnsListBySubscriptionPager provides operations for iterating over paged responses.
type PeerAsnsListBySubscriptionPager struct {
	client    *PeerAsnsClient
	current   PeerAsnsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeerAsnsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeerAsnsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeerAsnsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeerAsnListResult.NextLink == nil || len(*p.current.PeerAsnListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
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

// PageResponse returns the current PeerAsnsListBySubscriptionResponse page.
func (p *PeerAsnsListBySubscriptionPager) PageResponse() PeerAsnsListBySubscriptionResponse {
	return p.current
}

// PeeringLocationsListPager provides operations for iterating over paged responses.
type PeeringLocationsListPager struct {
	client    *PeeringLocationsClient
	current   PeeringLocationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringLocationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringLocationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringLocationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringLocationListResult.NextLink == nil || len(*p.current.PeeringLocationListResult.NextLink) == 0 {
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

// PageResponse returns the current PeeringLocationsListResponse page.
func (p *PeeringLocationsListPager) PageResponse() PeeringLocationsListResponse {
	return p.current
}

// PeeringServiceLocationsListPager provides operations for iterating over paged responses.
type PeeringServiceLocationsListPager struct {
	client    *PeeringServiceLocationsClient
	current   PeeringServiceLocationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringServiceLocationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringServiceLocationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringServiceLocationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringServiceLocationListResult.NextLink == nil || len(*p.current.PeeringServiceLocationListResult.NextLink) == 0 {
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

// PageResponse returns the current PeeringServiceLocationsListResponse page.
func (p *PeeringServiceLocationsListPager) PageResponse() PeeringServiceLocationsListResponse {
	return p.current
}

// PeeringServiceProvidersListPager provides operations for iterating over paged responses.
type PeeringServiceProvidersListPager struct {
	client    *PeeringServiceProvidersClient
	current   PeeringServiceProvidersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringServiceProvidersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringServiceProvidersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringServiceProvidersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringServiceProviderListResult.NextLink == nil || len(*p.current.PeeringServiceProviderListResult.NextLink) == 0 {
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

// PageResponse returns the current PeeringServiceProvidersListResponse page.
func (p *PeeringServiceProvidersListPager) PageResponse() PeeringServiceProvidersListResponse {
	return p.current
}

// PeeringServicesListByResourceGroupPager provides operations for iterating over paged responses.
type PeeringServicesListByResourceGroupPager struct {
	client    *PeeringServicesClient
	current   PeeringServicesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringServicesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringServicesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringServicesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringServiceListResult.NextLink == nil || len(*p.current.PeeringServiceListResult.NextLink) == 0 {
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

// PageResponse returns the current PeeringServicesListByResourceGroupResponse page.
func (p *PeeringServicesListByResourceGroupPager) PageResponse() PeeringServicesListByResourceGroupResponse {
	return p.current
}

// PeeringServicesListBySubscriptionPager provides operations for iterating over paged responses.
type PeeringServicesListBySubscriptionPager struct {
	client    *PeeringServicesClient
	current   PeeringServicesListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringServicesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringServicesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringServicesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringServiceListResult.NextLink == nil || len(*p.current.PeeringServiceListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
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

// PageResponse returns the current PeeringServicesListBySubscriptionResponse page.
func (p *PeeringServicesListBySubscriptionPager) PageResponse() PeeringServicesListBySubscriptionResponse {
	return p.current
}

// PeeringsListByResourceGroupPager provides operations for iterating over paged responses.
type PeeringsListByResourceGroupPager struct {
	client    *PeeringsClient
	current   PeeringsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringListResult.NextLink == nil || len(*p.current.PeeringListResult.NextLink) == 0 {
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

// PageResponse returns the current PeeringsListByResourceGroupResponse page.
func (p *PeeringsListByResourceGroupPager) PageResponse() PeeringsListByResourceGroupResponse {
	return p.current
}

// PeeringsListBySubscriptionPager provides operations for iterating over paged responses.
type PeeringsListBySubscriptionPager struct {
	client    *PeeringsClient
	current   PeeringsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PeeringsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PeeringsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PeeringsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringListResult.NextLink == nil || len(*p.current.PeeringListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
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

// PageResponse returns the current PeeringsListBySubscriptionResponse page.
func (p *PeeringsListBySubscriptionPager) PageResponse() PeeringsListBySubscriptionResponse {
	return p.current
}

// PrefixesListByPeeringServicePager provides operations for iterating over paged responses.
type PrefixesListByPeeringServicePager struct {
	client    *PrefixesClient
	current   PrefixesListByPeeringServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrefixesListByPeeringServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrefixesListByPeeringServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrefixesListByPeeringServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PeeringServicePrefixListResult.NextLink == nil || len(*p.current.PeeringServicePrefixListResult.NextLink) == 0 {
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
		p.err = p.client.listByPeeringServiceHandleError(resp)
		return false
	}
	result, err := p.client.listByPeeringServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrefixesListByPeeringServiceResponse page.
func (p *PrefixesListByPeeringServicePager) PageResponse() PrefixesListByPeeringServiceResponse {
	return p.current
}
