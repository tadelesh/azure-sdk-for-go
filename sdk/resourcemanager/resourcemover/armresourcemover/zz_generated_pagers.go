//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// MoveCollectionsClientListMoveCollectionsByResourceGroupPager provides operations for iterating over paged responses.
type MoveCollectionsClientListMoveCollectionsByResourceGroupPager struct {
	client    *MoveCollectionsClient
	current   MoveCollectionsClientListMoveCollectionsByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MoveCollectionsClientListMoveCollectionsByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *MoveCollectionsClientListMoveCollectionsByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MoveCollectionResultList.NextLink == nil || len(*p.current.MoveCollectionResultList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *MoveCollectionsClientListMoveCollectionsByResourceGroupPager) NextPage(ctx context.Context) (MoveCollectionsClientListMoveCollectionsByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return MoveCollectionsClientListMoveCollectionsByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return MoveCollectionsClientListMoveCollectionsByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listMoveCollectionsByResourceGroupHandleResponse(resp)
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// MoveCollectionsClientListMoveCollectionsBySubscriptionPager provides operations for iterating over paged responses.
type MoveCollectionsClientListMoveCollectionsBySubscriptionPager struct {
	client    *MoveCollectionsClient
	current   MoveCollectionsClientListMoveCollectionsBySubscriptionResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MoveCollectionsClientListMoveCollectionsBySubscriptionResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *MoveCollectionsClientListMoveCollectionsBySubscriptionPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MoveCollectionResultList.NextLink == nil || len(*p.current.MoveCollectionResultList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *MoveCollectionsClientListMoveCollectionsBySubscriptionPager) NextPage(ctx context.Context) (MoveCollectionsClientListMoveCollectionsBySubscriptionResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return MoveCollectionsClientListMoveCollectionsBySubscriptionResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsBySubscriptionResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return MoveCollectionsClientListMoveCollectionsBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listMoveCollectionsBySubscriptionHandleResponse(resp)
	if err != nil {
		return MoveCollectionsClientListMoveCollectionsBySubscriptionResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// MoveResourcesClientListPager provides operations for iterating over paged responses.
type MoveResourcesClientListPager struct {
	client    *MoveResourcesClient
	current   MoveResourcesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MoveResourcesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *MoveResourcesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MoveResourceCollection.NextLink == nil || len(*p.current.MoveResourceCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *MoveResourcesClientListPager) NextPage(ctx context.Context) (MoveResourcesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return MoveResourcesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return MoveResourcesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return MoveResourcesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return MoveResourcesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return MoveResourcesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// UnresolvedDependenciesClientGetPager provides operations for iterating over paged responses.
type UnresolvedDependenciesClientGetPager struct {
	client    *UnresolvedDependenciesClient
	current   UnresolvedDependenciesClientGetResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UnresolvedDependenciesClientGetResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *UnresolvedDependenciesClientGetPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.UnresolvedDependencyCollection.NextLink == nil || len(*p.current.UnresolvedDependencyCollection.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *UnresolvedDependenciesClientGetPager) NextPage(ctx context.Context) (UnresolvedDependenciesClientGetResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return UnresolvedDependenciesClientGetResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return UnresolvedDependenciesClientGetResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return UnresolvedDependenciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return UnresolvedDependenciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getHandleResponse(resp)
	if err != nil {
		return UnresolvedDependenciesClientGetResponse{}, err
	}
	p.current = result
	return p.current, nil
}
