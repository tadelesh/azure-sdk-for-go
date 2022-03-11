//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

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

// SharedPrivateLinkResourcesClient contains the methods for the WebPubSubSharedPrivateLinkResources group.
// Don't use this type directly, use NewSharedPrivateLinkResourcesClient() instead.
type SharedPrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient with the specified values.
// subscriptionID - Gets subscription Id which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSharedPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SharedPrivateLinkResourcesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SharedPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// parameters - The shared private link resource
// options - SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginCreateOrUpdate
// method.
func (client *SharedPrivateLinkResourcesClient) BeginCreateOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (SharedPrivateLinkResourcesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientCreateOrUpdatePollerResponse{}, err
	}
	result := SharedPrivateLinkResourcesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("SharedPrivateLinkResourcesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return SharedPrivateLinkResourcesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SharedPrivateLinkResourcesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SharedPrivateLinkResourcesClient) createOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SharedPrivateLinkResourcesClient) createOrUpdateCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientBeginDeleteOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginDelete
// method.
func (client *SharedPrivateLinkResourcesClient) BeginDelete(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (SharedPrivateLinkResourcesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientDeletePollerResponse{}, err
	}
	result := SharedPrivateLinkResourcesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("SharedPrivateLinkResourcesClient.Delete", "", resp, client.pl)
	if err != nil {
		return SharedPrivateLinkResourcesClientDeletePollerResponse{}, err
	}
	result.Poller = &SharedPrivateLinkResourcesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SharedPrivateLinkResourcesClient) deleteOperation(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SharedPrivateLinkResourcesClient) deleteCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientGetOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.Get
// method.
func (client *SharedPrivateLinkResourcesClient) Get(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientGetOptions) (SharedPrivateLinkResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedPrivateLinkResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedPrivateLinkResourcesClient) getCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedPrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientGetResponse, error) {
	result := SharedPrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResource); err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// List - List shared private link resources
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientListOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.List
// method.
func (client *SharedPrivateLinkResourcesClient) List(resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientListOptions) *SharedPrivateLinkResourcesClientListPager {
	return &SharedPrivateLinkResourcesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp SharedPrivateLinkResourcesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SharedPrivateLinkResourceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SharedPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientListResponse, error) {
	result := SharedPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResourceList); err != nil {
		return SharedPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
