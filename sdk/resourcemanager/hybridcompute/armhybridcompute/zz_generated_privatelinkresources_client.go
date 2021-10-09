//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
func NewPrivateLinkResourcesClient(con *arm.Connection, subscriptionID string) *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesGetOptions) (PrivateLinkResourcesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, groupName, options)
	if err != nil {
		return PrivateLinkResourcesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesGetResponse, error) {
	result := PrivateLinkResourcesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateLinkResourcesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByPrivateLinkScope - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateLinkResourcesClient) ListByPrivateLinkScope(resourceGroupName string, scopeName string, options *PrivateLinkResourcesListByPrivateLinkScopeOptions) *PrivateLinkResourcesListByPrivateLinkScopePager {
	return &PrivateLinkResourcesListByPrivateLinkScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkResourcesListByPrivateLinkScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkResourceListResult.NextLink)
		},
	}
}

// listByPrivateLinkScopeCreateRequest creates the ListByPrivateLinkScope request.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkResourcesListByPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPrivateLinkScopeHandleResponse handles the ListByPrivateLinkScope response.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeHandleResponse(resp *http.Response) (PrivateLinkResourcesListByPrivateLinkScopeResponse, error) {
	result := PrivateLinkResourcesListByPrivateLinkScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesListByPrivateLinkScopeResponse{}, err
	}
	return result, nil
}

// listByPrivateLinkScopeHandleError handles the ListByPrivateLinkScope error response.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
