//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountClientListByResourceGroupPager provides operations for iterating over paged responses.
type AccountClientListByResourceGroupPager struct {
	client    *AccountClient
	current   AccountClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccountClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccountClientListByResourceGroupPager) NextPage(ctx context.Context) (AccountClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccountClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccountClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccountClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccountClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return AccountClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// AccountClientListOutboundNetworkDependenciesEndpointsPager provides operations for iterating over paged responses.
type AccountClientListOutboundNetworkDependenciesEndpointsPager struct {
	client    *AccountClient
	current   AccountClientListOutboundNetworkDependenciesEndpointsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountClientListOutboundNetworkDependenciesEndpointsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccountClientListOutboundNetworkDependenciesEndpointsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEnvironmentEndpointCollection.NextLink == nil || len(*p.current.OutboundEnvironmentEndpointCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccountClientListOutboundNetworkDependenciesEndpointsPager) NextPage(ctx context.Context) (AccountClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
	if err != nil {
		return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// AccountClientListPager provides operations for iterating over paged responses.
type AccountClientListPager struct {
	client    *AccountClient
	current   AccountClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccountClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccountClientListPager) NextPage(ctx context.Context) (AccountClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccountClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccountClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccountClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccountClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return AccountClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ApplicationClientListPager provides operations for iterating over paged responses.
type ApplicationClientListPager struct {
	client    *ApplicationClient
	current   ApplicationClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ApplicationClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ApplicationClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListApplicationsResult.NextLink == nil || len(*p.current.ListApplicationsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ApplicationClientListPager) NextPage(ctx context.Context) (ApplicationClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ApplicationClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ApplicationClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ApplicationClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ApplicationClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ApplicationClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ApplicationPackageClientListPager provides operations for iterating over paged responses.
type ApplicationPackageClientListPager struct {
	client    *ApplicationPackageClient
	current   ApplicationPackageClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ApplicationPackageClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ApplicationPackageClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListApplicationPackagesResult.NextLink == nil || len(*p.current.ListApplicationPackagesResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ApplicationPackageClientListPager) NextPage(ctx context.Context) (ApplicationPackageClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ApplicationPackageClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ApplicationPackageClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ApplicationPackageClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ApplicationPackageClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ApplicationPackageClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// CertificateClientListByBatchAccountPager provides operations for iterating over paged responses.
type CertificateClientListByBatchAccountPager struct {
	client    *CertificateClient
	current   CertificateClientListByBatchAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CertificateClientListByBatchAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *CertificateClientListByBatchAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListCertificatesResult.NextLink == nil || len(*p.current.ListCertificatesResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *CertificateClientListByBatchAccountPager) NextPage(ctx context.Context) (CertificateClientListByBatchAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return CertificateClientListByBatchAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return CertificateClientListByBatchAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return CertificateClientListByBatchAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return CertificateClientListByBatchAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		return CertificateClientListByBatchAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LocationClientListSupportedCloudServiceSKUsPager provides operations for iterating over paged responses.
type LocationClientListSupportedCloudServiceSKUsPager struct {
	client    *LocationClient
	current   LocationClientListSupportedCloudServiceSKUsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationClientListSupportedCloudServiceSKUsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LocationClientListSupportedCloudServiceSKUsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SupportedSKUsResult.NextLink == nil || len(*p.current.SupportedSKUsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *LocationClientListSupportedCloudServiceSKUsPager) NextPage(ctx context.Context) (LocationClientListSupportedCloudServiceSKUsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return LocationClientListSupportedCloudServiceSKUsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LocationClientListSupportedCloudServiceSKUsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LocationClientListSupportedCloudServiceSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LocationClientListSupportedCloudServiceSKUsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listSupportedCloudServiceSKUsHandleResponse(resp)
	if err != nil {
		return LocationClientListSupportedCloudServiceSKUsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LocationClientListSupportedVirtualMachineSKUsPager provides operations for iterating over paged responses.
type LocationClientListSupportedVirtualMachineSKUsPager struct {
	client    *LocationClient
	current   LocationClientListSupportedVirtualMachineSKUsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationClientListSupportedVirtualMachineSKUsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LocationClientListSupportedVirtualMachineSKUsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SupportedSKUsResult.NextLink == nil || len(*p.current.SupportedSKUsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *LocationClientListSupportedVirtualMachineSKUsPager) NextPage(ctx context.Context) (LocationClientListSupportedVirtualMachineSKUsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return LocationClientListSupportedVirtualMachineSKUsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LocationClientListSupportedVirtualMachineSKUsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LocationClientListSupportedVirtualMachineSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LocationClientListSupportedVirtualMachineSKUsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listSupportedVirtualMachineSKUsHandleResponse(resp)
	if err != nil {
		return LocationClientListSupportedVirtualMachineSKUsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *OperationsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *OperationsClientListPager) NextPage(ctx context.Context) (OperationsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return OperationsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return OperationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// PoolClientListByBatchAccountPager provides operations for iterating over paged responses.
type PoolClientListByBatchAccountPager struct {
	client    *PoolClient
	current   PoolClientListByBatchAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PoolClientListByBatchAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PoolClientListByBatchAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPoolsResult.NextLink == nil || len(*p.current.ListPoolsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *PoolClientListByBatchAccountPager) NextPage(ctx context.Context) (PoolClientListByBatchAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return PoolClientListByBatchAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PoolClientListByBatchAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PoolClientListByBatchAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PoolClientListByBatchAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		return PoolClientListByBatchAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// PrivateEndpointConnectionClientListByBatchAccountPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionClientListByBatchAccountPager struct {
	client    *PrivateEndpointConnectionClient
	current   PrivateEndpointConnectionClientListByBatchAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionClientListByBatchAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PrivateEndpointConnectionClientListByBatchAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPrivateEndpointConnectionsResult.NextLink == nil || len(*p.current.ListPrivateEndpointConnectionsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *PrivateEndpointConnectionClientListByBatchAccountPager) NextPage(ctx context.Context) (PrivateEndpointConnectionClientListByBatchAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return PrivateEndpointConnectionClientListByBatchAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PrivateEndpointConnectionClientListByBatchAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionClientListByBatchAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PrivateEndpointConnectionClientListByBatchAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		return PrivateEndpointConnectionClientListByBatchAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// PrivateLinkResourceClientListByBatchAccountPager provides operations for iterating over paged responses.
type PrivateLinkResourceClientListByBatchAccountPager struct {
	client    *PrivateLinkResourceClient
	current   PrivateLinkResourceClientListByBatchAccountResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkResourceClientListByBatchAccountResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PrivateLinkResourceClientListByBatchAccountPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPrivateLinkResourcesResult.NextLink == nil || len(*p.current.ListPrivateLinkResourcesResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *PrivateLinkResourceClientListByBatchAccountPager) NextPage(ctx context.Context) (PrivateLinkResourceClientListByBatchAccountResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return PrivateLinkResourceClientListByBatchAccountResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PrivateLinkResourceClientListByBatchAccountResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourceClientListByBatchAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PrivateLinkResourceClientListByBatchAccountResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		return PrivateLinkResourceClientListByBatchAccountResponse{}, err
	}
	p.current = result
	return p.current, nil
}
