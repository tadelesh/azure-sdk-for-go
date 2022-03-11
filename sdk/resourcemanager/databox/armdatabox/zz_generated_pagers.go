//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// JobsClientListByResourceGroupPager provides operations for iterating over paged responses.
type JobsClientListByResourceGroupPager struct {
	client    *JobsClient
	current   JobsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobResourceList.NextLink == nil || len(*p.current.JobResourceList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *JobsClientListByResourceGroupPager) NextPage(ctx context.Context) (JobsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return JobsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return JobsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// JobsClientListCredentialsPager provides operations for iterating over paged responses.
type JobsClientListCredentialsPager struct {
	client    *JobsClient
	current   JobsClientListCredentialsResponse
	requester func(context.Context) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobsClientListCredentialsPager) More() bool {
	return reflect.ValueOf(p.current).IsZero()
}

// NextPage advances the pager to the next page.
func (p *JobsClientListCredentialsPager) NextPage(ctx context.Context) (JobsClientListCredentialsResponse, error) {
	var req *policy.Request
	var err error
	if !p.More() {
		return JobsClientListCredentialsResponse{}, errors.New("no more pages")
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobsClientListCredentialsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobsClientListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobsClientListCredentialsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listCredentialsHandleResponse(resp)
	if err != nil {
		return JobsClientListCredentialsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// JobsClientListPager provides operations for iterating over paged responses.
type JobsClientListPager struct {
	client    *JobsClient
	current   JobsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobResourceList.NextLink == nil || len(*p.current.JobResourceList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *JobsClientListPager) NextPage(ctx context.Context) (JobsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return JobsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return JobsClientListResponse{}, err
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
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
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

// ServiceClientListAvailableSKUsByResourceGroupPager provides operations for iterating over paged responses.
type ServiceClientListAvailableSKUsByResourceGroupPager struct {
	client    *ServiceClient
	current   ServiceClientListAvailableSKUsByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceClientListAvailableSKUsByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ServiceClientListAvailableSKUsByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailableSKUsResult.NextLink == nil || len(*p.current.AvailableSKUsResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ServiceClientListAvailableSKUsByResourceGroupPager) NextPage(ctx context.Context) (ServiceClientListAvailableSKUsByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServiceClientListAvailableSKUsByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServiceClientListAvailableSKUsByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServiceClientListAvailableSKUsByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServiceClientListAvailableSKUsByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listAvailableSKUsByResourceGroupHandleResponse(resp)
	if err != nil {
		return ServiceClientListAvailableSKUsByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}
