//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccessPoliciesClientListPager provides operations for iterating over paged responses.
type AccessPoliciesClientListPager struct {
	client    *AccessPoliciesClient
	current   AccessPoliciesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessPoliciesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *AccessPoliciesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessPolicyEntityCollection.NextLink == nil || len(*p.current.AccessPolicyEntityCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *AccessPoliciesClientListPager) NextPage(ctx context.Context) (AccessPoliciesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return AccessPoliciesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return AccessPoliciesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return AccessPoliciesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return AccessPoliciesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// EdgeModulesClientListPager provides operations for iterating over paged responses.
type EdgeModulesClientListPager struct {
	client    *EdgeModulesClient
	current   EdgeModulesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EdgeModulesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *EdgeModulesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EdgeModuleEntityCollection.NextLink == nil || len(*p.current.EdgeModuleEntityCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *EdgeModulesClientListPager) NextPage(ctx context.Context) (EdgeModulesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return EdgeModulesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return EdgeModulesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return EdgeModulesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return EdgeModulesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return EdgeModulesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LivePipelinesClientListPager provides operations for iterating over paged responses.
type LivePipelinesClientListPager struct {
	client    *LivePipelinesClient
	current   LivePipelinesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LivePipelinesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LivePipelinesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LivePipelineCollection.NextLink == nil || len(*p.current.LivePipelineCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *LivePipelinesClientListPager) NextPage(ctx context.Context) (LivePipelinesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return LivePipelinesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LivePipelinesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LivePipelinesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LivePipelinesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return LivePipelinesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// PipelineJobsClientListPager provides operations for iterating over paged responses.
type PipelineJobsClientListPager struct {
	client    *PipelineJobsClient
	current   PipelineJobsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PipelineJobsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PipelineJobsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineJobCollection.NextLink == nil || len(*p.current.PipelineJobCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *PipelineJobsClientListPager) NextPage(ctx context.Context) (PipelineJobsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return PipelineJobsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PipelineJobsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PipelineJobsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PipelineJobsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return PipelineJobsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// PipelineTopologiesClientListPager provides operations for iterating over paged responses.
type PipelineTopologiesClientListPager struct {
	client    *PipelineTopologiesClient
	current   PipelineTopologiesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PipelineTopologiesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PipelineTopologiesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineTopologyCollection.NextLink == nil || len(*p.current.PipelineTopologyCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *PipelineTopologiesClientListPager) NextPage(ctx context.Context) (PipelineTopologiesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return PipelineTopologiesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PipelineTopologiesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PipelineTopologiesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PipelineTopologiesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return PipelineTopologiesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VideosClientListPager provides operations for iterating over paged responses.
type VideosClientListPager struct {
	client    *VideosClient
	current   VideosClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VideosClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VideosClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VideoEntityCollection.NextLink == nil || len(*p.current.VideoEntityCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VideosClientListPager) NextPage(ctx context.Context) (VideosClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VideosClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VideosClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VideosClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VideosClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return VideosClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
