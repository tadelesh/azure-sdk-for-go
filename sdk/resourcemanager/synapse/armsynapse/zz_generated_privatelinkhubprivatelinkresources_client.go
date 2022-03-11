//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// PrivateLinkHubPrivateLinkResourcesClient contains the methods for the PrivateLinkHubPrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkHubPrivateLinkResourcesClient() instead.
type PrivateLinkHubPrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkHubPrivateLinkResourcesClient creates a new instance of PrivateLinkHubPrivateLinkResourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkHubPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkHubPrivateLinkResourcesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateLinkHubPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get private link resource in private link hub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - The name of the private link hub
// privateLinkResourceName - The name of the private link resource
// options - PrivateLinkHubPrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkHubPrivateLinkResourcesClient.Get
// method.
func (client *PrivateLinkHubPrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkResourceName string, options *PrivateLinkHubPrivateLinkResourcesClientGetOptions) (PrivateLinkHubPrivateLinkResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateLinkHubName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkHubPrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkHubPrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkHubPrivateLinkResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkHubPrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkResourceName string, options *PrivateLinkHubPrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateLinkResources/{privateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
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
func (client *PrivateLinkHubPrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkHubPrivateLinkResourcesClientGetResponse, error) {
	result := PrivateLinkHubPrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkHubPrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all private link resources for a private link hub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - The name of the private link hub
// options - PrivateLinkHubPrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkHubPrivateLinkResourcesClient.List
// method.
func (client *PrivateLinkHubPrivateLinkResourcesClient) List(resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubPrivateLinkResourcesClientListOptions) *PrivateLinkHubPrivateLinkResourcesClientListPager {
	return &PrivateLinkHubPrivateLinkResourcesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateLinkHubName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkHubPrivateLinkResourcesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkResourceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateLinkHubPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
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
func (client *PrivateLinkHubPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (PrivateLinkHubPrivateLinkResourcesClientListResponse, error) {
	result := PrivateLinkHubPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkHubPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
