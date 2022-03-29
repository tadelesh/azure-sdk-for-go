//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DeviceGroupsClientListPager provides operations for iterating over paged responses.
type DeviceGroupsClientListPager struct {
	client    *DeviceGroupsClient
	current   DeviceGroupsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DeviceGroupsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DeviceGroupsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DeviceGroupsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceGroupList.NextLink == nil || len(*p.current.DeviceGroupList.NextLink) == 0 {
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

// PageResponse returns the current DeviceGroupsClientListResponse page.
func (p *DeviceGroupsClientListPager) PageResponse() DeviceGroupsClientListResponse {
	return p.current
}

// DevicesClientListPager provides operations for iterating over paged responses.
type DevicesClientListPager struct {
	client    *DevicesClient
	current   DevicesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DevicesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DevicesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DevicesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceList.NextLink == nil || len(*p.current.DeviceList.NextLink) == 0 {
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

// PageResponse returns the current DevicesClientListResponse page.
func (p *DevicesClientListPager) PageResponse() DevicesClientListResponse {
	return p.current
}

// IotAlertsClientListPager provides operations for iterating over paged responses.
type IotAlertsClientListPager struct {
	client    *IotAlertsClient
	current   IotAlertsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotAlertsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotAlertsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotAlertsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AlertListModel.NextLink == nil || len(*p.current.AlertListModel.NextLink) == 0 {
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

// PageResponse returns the current IotAlertsClientListResponse page.
func (p *IotAlertsClientListPager) PageResponse() IotAlertsClientListResponse {
	return p.current
}

// IotDeviceVulnerabilityClientListPager provides operations for iterating over paged responses.
type IotDeviceVulnerabilityClientListPager struct {
	client    *IotDeviceVulnerabilityClient
	current   IotDeviceVulnerabilityClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDeviceVulnerabilityClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDeviceVulnerabilityClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDeviceVulnerabilityClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceVulnerabilityListModel.NextLink == nil || len(*p.current.DeviceVulnerabilityListModel.NextLink) == 0 {
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

// PageResponse returns the current IotDeviceVulnerabilityClientListResponse page.
func (p *IotDeviceVulnerabilityClientListPager) PageResponse() IotDeviceVulnerabilityClientListResponse {
	return p.current
}

// IotRecommendationsClientListPager provides operations for iterating over paged responses.
type IotRecommendationsClientListPager struct {
	client    *IotRecommendationsClient
	current   IotRecommendationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotRecommendationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotRecommendationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotRecommendationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecommendationListModel.NextLink == nil || len(*p.current.RecommendationListModel.NextLink) == 0 {
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

// PageResponse returns the current IotRecommendationsClientListResponse page.
func (p *IotRecommendationsClientListPager) PageResponse() IotRecommendationsClientListResponse {
	return p.current
}

// LocationsClientListPager provides operations for iterating over paged responses.
type LocationsClientListPager struct {
	client    *LocationsClient
	current   LocationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LocationList.NextLink == nil || len(*p.current.LocationList.NextLink) == 0 {
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

// PageResponse returns the current LocationsClientListResponse page.
func (p *LocationsClientListPager) PageResponse() LocationsClientListResponse {
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
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
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