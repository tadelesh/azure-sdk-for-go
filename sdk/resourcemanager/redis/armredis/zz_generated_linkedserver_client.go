//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkedServerClient contains the methods for the LinkedServer group.
// Don't use this type directly, use NewLinkedServerClient() instead.
type LinkedServerClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLinkedServerClient creates a new instance of LinkedServerClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkedServerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LinkedServerClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LinkedServerClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the Redis cache.
// linkedServerName - The name of the linked server that is being added to the Redis cache.
// parameters - Parameters supplied to the Create Linked server operation.
// options - LinkedServerClientBeginCreateOptions contains the optional parameters for the LinkedServerClient.BeginCreate
// method.
func (client *LinkedServerClient) BeginCreate(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (LinkedServerClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, name, linkedServerName, parameters, options)
	if err != nil {
		return LinkedServerClientCreatePollerResponse{}, err
	}
	result := LinkedServerClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("LinkedServerClient.Create", "", resp, client.pl)
	if err != nil {
		return LinkedServerClientCreatePollerResponse{}, err
	}
	result.Poller = &LinkedServerClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LinkedServerClient) create(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, name, linkedServerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LinkedServerClient) createCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes the linked server from a redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the redis cache.
// linkedServerName - The name of the linked server that is being added to the Redis cache.
// options - LinkedServerClientDeleteOptions contains the optional parameters for the LinkedServerClient.Delete method.
func (client *LinkedServerClient) Delete(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientDeleteOptions) (LinkedServerClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return LinkedServerClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedServerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LinkedServerClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LinkedServerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedServerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the detailed information about a linked server of a redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the redis cache.
// linkedServerName - The name of the linked server.
// options - LinkedServerClientGetOptions contains the optional parameters for the LinkedServerClient.Get method.
func (client *LinkedServerClient) Get(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientGetOptions) (LinkedServerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedServerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedServerClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedServerClient) getHandleResponse(resp *http.Response) (LinkedServerClientGetResponse, error) {
	result := LinkedServerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServerWithProperties); err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of linked servers associated with this redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the redis cache.
// options - LinkedServerClientListOptions contains the optional parameters for the LinkedServerClient.List method.
func (client *LinkedServerClient) List(resourceGroupName string, name string, options *LinkedServerClientListOptions) *LinkedServerClientListPager {
	return &LinkedServerClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, name, options)
		},
		advancer: func(ctx context.Context, resp LinkedServerClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LinkedServerWithPropertiesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LinkedServerClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, options *LinkedServerClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LinkedServerClient) listHandleResponse(resp *http.Response) (LinkedServerClientListResponse, error) {
	result := LinkedServerClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServerWithPropertiesList); err != nil {
		return LinkedServerClientListResponse{}, err
	}
	return result, nil
}
