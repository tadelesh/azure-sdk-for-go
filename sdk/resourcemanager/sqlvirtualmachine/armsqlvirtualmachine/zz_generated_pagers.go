//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AvailabilityGroupListenersClientListByGroupPager provides operations for iterating over paged responses.
type AvailabilityGroupListenersClientListByGroupPager struct {
	client    *AvailabilityGroupListenersClient
	current   AvailabilityGroupListenersClientListByGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AvailabilityGroupListenersClientListByGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AvailabilityGroupListenersClientListByGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailabilityGroupListenerListResult.NextLink == nil || len(*p.current.AvailabilityGroupListenerListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AvailabilityGroupListenersClientListByGroupPager) NextPage(ctx context.Context) (AvailabilityGroupListenersClientListByGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AvailabilityGroupListenersClientListByGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AvailabilityGroupListenersClientListByGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AvailabilityGroupListenersClientListByGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AvailabilityGroupListenersClientListByGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByGroupHandleResponse(resp)
	if err != nil {
		return AvailabilityGroupListenersClientListByGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// GroupsClientListByResourceGroupPager provides operations for iterating over paged responses.
type GroupsClientListByResourceGroupPager struct {
	client    *GroupsClient
	current   GroupsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GroupsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *GroupsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GroupListResult.NextLink == nil || len(*p.current.GroupListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *GroupsClientListByResourceGroupPager) NextPage(ctx context.Context) (GroupsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return GroupsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return GroupsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return GroupsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return GroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return GroupsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// GroupsClientListPager provides operations for iterating over paged responses.
type GroupsClientListPager struct {
	client    *GroupsClient
	current   GroupsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GroupsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *GroupsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GroupListResult.NextLink == nil || len(*p.current.GroupListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *GroupsClientListPager) NextPage(ctx context.Context) (GroupsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return GroupsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return GroupsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return GroupsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return GroupsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return GroupsClientListResponse{}, err
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

// SQLVirtualMachinesClientListByResourceGroupPager provides operations for iterating over paged responses.
type SQLVirtualMachinesClientListByResourceGroupPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *SQLVirtualMachinesClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *SQLVirtualMachinesClientListByResourceGroupPager) NextPage(ctx context.Context) (SQLVirtualMachinesClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return SQLVirtualMachinesClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return SQLVirtualMachinesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// SQLVirtualMachinesClientListBySQLVMGroupPager provides operations for iterating over paged responses.
type SQLVirtualMachinesClientListBySQLVMGroupPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesClientListBySQLVMGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesClientListBySQLVMGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *SQLVirtualMachinesClientListBySQLVMGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *SQLVirtualMachinesClientListBySQLVMGroupPager) NextPage(ctx context.Context) (SQLVirtualMachinesClientListBySQLVMGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBySQLVMGroupHandleResponse(resp)
	if err != nil {
		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// SQLVirtualMachinesClientListPager provides operations for iterating over paged responses.
type SQLVirtualMachinesClientListPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *SQLVirtualMachinesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *SQLVirtualMachinesClientListPager) NextPage(ctx context.Context) (SQLVirtualMachinesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return SQLVirtualMachinesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return SQLVirtualMachinesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return SQLVirtualMachinesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return SQLVirtualMachinesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return SQLVirtualMachinesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
