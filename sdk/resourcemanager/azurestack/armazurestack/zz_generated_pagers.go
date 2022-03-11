//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// CustomerSubscriptionsClientListPager provides operations for iterating over paged responses.
type CustomerSubscriptionsClientListPager struct {
	client    *CustomerSubscriptionsClient
	current   CustomerSubscriptionsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CustomerSubscriptionsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *CustomerSubscriptionsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CustomerSubscriptionList.NextLink == nil || len(*p.current.CustomerSubscriptionList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *CustomerSubscriptionsClientListPager) NextPage(ctx context.Context) (CustomerSubscriptionsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return CustomerSubscriptionsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return CustomerSubscriptionsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return CustomerSubscriptionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return CustomerSubscriptionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return CustomerSubscriptionsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LinkedSubscriptionsClientListByResourceGroupPager provides operations for iterating over paged responses.
type LinkedSubscriptionsClientListByResourceGroupPager struct {
	client    *LinkedSubscriptionsClient
	current   LinkedSubscriptionsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LinkedSubscriptionsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LinkedSubscriptionsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedSubscriptionsList.NextLink == nil || len(*p.current.LinkedSubscriptionsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *LinkedSubscriptionsClientListByResourceGroupPager) NextPage(ctx context.Context) (LinkedSubscriptionsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return LinkedSubscriptionsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LinkedSubscriptionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// LinkedSubscriptionsClientListBySubscriptionPager provides operations for iterating over paged responses.
type LinkedSubscriptionsClientListBySubscriptionPager struct {
	client    *LinkedSubscriptionsClient
	current   LinkedSubscriptionsClientListBySubscriptionResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LinkedSubscriptionsClientListBySubscriptionResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *LinkedSubscriptionsClientListBySubscriptionPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedSubscriptionsList.NextLink == nil || len(*p.current.LinkedSubscriptionsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *LinkedSubscriptionsClientListBySubscriptionPager) NextPage(ctx context.Context) (LinkedSubscriptionsClientListBySubscriptionResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return LinkedSubscriptionsClientListBySubscriptionResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return LinkedSubscriptionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
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

// ProductsClientListPager provides operations for iterating over paged responses.
type ProductsClientListPager struct {
	client    *ProductsClient
	current   ProductsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProductsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ProductsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductList.NextLink == nil || len(*p.current.ProductList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ProductsClientListPager) NextPage(ctx context.Context) (ProductsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ProductsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ProductsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ProductsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ProductsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ProductsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// RegistrationsClientListBySubscriptionPager provides operations for iterating over paged responses.
type RegistrationsClientListBySubscriptionPager struct {
	client    *RegistrationsClient
	current   RegistrationsClientListBySubscriptionResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RegistrationsClientListBySubscriptionResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *RegistrationsClientListBySubscriptionPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RegistrationList.NextLink == nil || len(*p.current.RegistrationList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *RegistrationsClientListBySubscriptionPager) NextPage(ctx context.Context) (RegistrationsClientListBySubscriptionResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return RegistrationsClientListBySubscriptionResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return RegistrationsClientListBySubscriptionResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return RegistrationsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return RegistrationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		return RegistrationsClientListBySubscriptionResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// RegistrationsClientListPager provides operations for iterating over paged responses.
type RegistrationsClientListPager struct {
	client    *RegistrationsClient
	current   RegistrationsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RegistrationsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *RegistrationsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RegistrationList.NextLink == nil || len(*p.current.RegistrationList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *RegistrationsClientListPager) NextPage(ctx context.Context) (RegistrationsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return RegistrationsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return RegistrationsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return RegistrationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return RegistrationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return RegistrationsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
