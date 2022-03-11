//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

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

// ServiceClientListByResourceGroupPager provides operations for iterating over paged responses.
type ServiceClientListByResourceGroupPager struct {
	client    *ServiceClient
	current   ServiceClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ServiceClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServiceResourceList.NextLink == nil || len(*p.current.ServiceResourceList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ServiceClientListByResourceGroupPager) NextPage(ctx context.Context) (ServiceClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServiceClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServiceClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServiceClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServiceClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return ServiceClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ServiceClientListBySubscriptionPager provides operations for iterating over paged responses.
type ServiceClientListBySubscriptionPager struct {
	client    *ServiceClient
	current   ServiceClientListBySubscriptionResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceClientListBySubscriptionResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ServiceClientListBySubscriptionPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServiceResourceList.NextLink == nil || len(*p.current.ServiceResourceList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ServiceClientListBySubscriptionPager) NextPage(ctx context.Context) (ServiceClientListBySubscriptionResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServiceClientListBySubscriptionResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServiceClientListBySubscriptionResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServiceClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServiceClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		return ServiceClientListBySubscriptionResponse{}, err
	}
	p.current = result
	return p.current, nil
}
