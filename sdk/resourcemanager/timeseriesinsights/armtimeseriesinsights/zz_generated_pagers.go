//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtimeseriesinsights

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

// PrivateLinkResourcesClientListSupportedPager provides operations for iterating over paged responses.
type PrivateLinkResourcesClientListSupportedPager struct {
	client    *PrivateLinkResourcesClient
	current   PrivateLinkResourcesClientListSupportedResponse
	requester func(context.Context) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PrivateLinkResourcesClientListSupportedPager) More() bool {
	return reflect.ValueOf(p.current).IsZero()
}

// NextPage advances the pager to the next page.
func (p *PrivateLinkResourcesClientListSupportedPager) NextPage(ctx context.Context) (PrivateLinkResourcesClientListSupportedResponse, error) {
	var req *policy.Request
	var err error
	if !p.More() {
		return PrivateLinkResourcesClientListSupportedResponse{}, errors.New("no more pages")
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PrivateLinkResourcesClientListSupportedResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesClientListSupportedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PrivateLinkResourcesClientListSupportedResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listSupportedHandleResponse(resp)
	if err != nil {
		return PrivateLinkResourcesClientListSupportedResponse{}, err
	}
	p.current = result
	return p.current, nil
}
