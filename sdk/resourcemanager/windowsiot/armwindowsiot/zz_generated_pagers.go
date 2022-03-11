//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsiot

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

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

// ServicesClientListByResourceGroupPager provides operations for iterating over paged responses.
type ServicesClientListByResourceGroupPager struct {
	client    *ServicesClient
	current   ServicesClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ServicesClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceServiceDescriptionListResult.NextLink == nil || len(*p.current.DeviceServiceDescriptionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ServicesClientListByResourceGroupPager) NextPage(ctx context.Context) (ServicesClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServicesClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ServicesClientListPager provides operations for iterating over paged responses.
type ServicesClientListPager struct {
	client    *ServicesClient
	current   ServicesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ServicesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceServiceDescriptionListResult.NextLink == nil || len(*p.current.DeviceServiceDescriptionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ServicesClientListPager) NextPage(ctx context.Context) (ServicesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServicesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServicesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServicesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServicesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ServicesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
