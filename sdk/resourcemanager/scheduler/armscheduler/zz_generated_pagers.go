//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// JobCollectionsClientListByResourceGroupPager provides operations for iterating over paged responses.
type JobCollectionsClientListByResourceGroupPager struct {
	client    *JobCollectionsClient
	current   JobCollectionsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobCollectionsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobCollectionsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobCollectionListResult.NextLink == nil || len(*p.current.JobCollectionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *JobCollectionsClientListByResourceGroupPager) NextPage(ctx context.Context) (JobCollectionsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return JobCollectionsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobCollectionsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobCollectionsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobCollectionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return JobCollectionsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// JobCollectionsClientListBySubscriptionPager provides operations for iterating over paged responses.
type JobCollectionsClientListBySubscriptionPager struct {
	client    *JobCollectionsClient
	current   JobCollectionsClientListBySubscriptionResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobCollectionsClientListBySubscriptionResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobCollectionsClientListBySubscriptionPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobCollectionListResult.NextLink == nil || len(*p.current.JobCollectionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *JobCollectionsClientListBySubscriptionPager) NextPage(ctx context.Context) (JobCollectionsClientListBySubscriptionResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return JobCollectionsClientListBySubscriptionResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobCollectionsClientListBySubscriptionResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobCollectionsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobCollectionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		return JobCollectionsClientListBySubscriptionResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// JobsClientListJobHistoryPager provides operations for iterating over paged responses.
type JobsClientListJobHistoryPager struct {
	client    *JobsClient
	current   JobsClientListJobHistoryResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListJobHistoryResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *JobsClientListJobHistoryPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobHistoryListResult.NextLink == nil || len(*p.current.JobHistoryListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *JobsClientListJobHistoryPager) NextPage(ctx context.Context) (JobsClientListJobHistoryResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return JobsClientListJobHistoryResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return JobsClientListJobHistoryResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return JobsClientListJobHistoryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return JobsClientListJobHistoryResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listJobHistoryHandleResponse(resp)
	if err != nil {
		return JobsClientListJobHistoryResponse{}, err
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
		if p.current.JobListResult.NextLink == nil || len(*p.current.JobListResult.NextLink) == 0 {
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
