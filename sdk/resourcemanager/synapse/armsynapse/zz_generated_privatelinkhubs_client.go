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

// PrivateLinkHubsClient contains the methods for the PrivateLinkHubs group.
// Don't use this type directly, use NewPrivateLinkHubsClient() instead.
type PrivateLinkHubsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkHubsClient creates a new instance of PrivateLinkHubsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkHubsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkHubsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateLinkHubsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates a privateLinkHub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - Name of the privateLinkHub
// privateLinkHubInfo - PrivateLinkHub create or update request properties
// options - PrivateLinkHubsClientCreateOrUpdateOptions contains the optional parameters for the PrivateLinkHubsClient.CreateOrUpdate
// method.
func (client *PrivateLinkHubsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkHubInfo PrivateLinkHub, options *PrivateLinkHubsClientCreateOrUpdateOptions) (PrivateLinkHubsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateLinkHubName, privateLinkHubInfo, options)
	if err != nil {
		return PrivateLinkHubsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkHubsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrivateLinkHubsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkHubsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkHubInfo PrivateLinkHub, options *PrivateLinkHubsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateLinkHubInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateLinkHubsClient) createOrUpdateHandleResponse(resp *http.Response) (PrivateLinkHubsClientCreateOrUpdateResponse, error) {
	result := PrivateLinkHubsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkHub); err != nil {
		return PrivateLinkHubsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a privateLinkHub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - Name of the privateLinkHub
// options - PrivateLinkHubsClientBeginDeleteOptions contains the optional parameters for the PrivateLinkHubsClient.BeginDelete
// method.
func (client *PrivateLinkHubsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubsClientBeginDeleteOptions) (PrivateLinkHubsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateLinkHubName, options)
	if err != nil {
		return PrivateLinkHubsClientDeletePollerResponse{}, err
	}
	result := PrivateLinkHubsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("PrivateLinkHubsClient.Delete", "", resp, client.pl)
	if err != nil {
		return PrivateLinkHubsClientDeletePollerResponse{}, err
	}
	result.Poller = &PrivateLinkHubsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a privateLinkHub
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkHubsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateLinkHubName, options)
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
func (client *PrivateLinkHubsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}"
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

// Get - Gets a privateLinkHub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - Name of the privateLinkHub
// options - PrivateLinkHubsClientGetOptions contains the optional parameters for the PrivateLinkHubsClient.Get method.
func (client *PrivateLinkHubsClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubsClientGetOptions) (PrivateLinkHubsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateLinkHubName, options)
	if err != nil {
		return PrivateLinkHubsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkHubsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkHubsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkHubsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateLinkHubsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}"
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

// getHandleResponse handles the Get response.
func (client *PrivateLinkHubsClient) getHandleResponse(resp *http.Response) (PrivateLinkHubsClientGetResponse, error) {
	result := PrivateLinkHubsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkHub); err != nil {
		return PrivateLinkHubsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns a list of privateLinkHubs in a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateLinkHubsClientListOptions contains the optional parameters for the PrivateLinkHubsClient.List method.
func (client *PrivateLinkHubsClient) List(options *PrivateLinkHubsClientListOptions) *PrivateLinkHubsClientListPager {
	return &PrivateLinkHubsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkHubsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkHubInfoListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateLinkHubsClient) listCreateRequest(ctx context.Context, options *PrivateLinkHubsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/privateLinkHubs"
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
func (client *PrivateLinkHubsClient) listHandleResponse(resp *http.Response) (PrivateLinkHubsClientListResponse, error) {
	result := PrivateLinkHubsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkHubInfoListResult); err != nil {
		return PrivateLinkHubsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Returns a list of privateLinkHubs in a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PrivateLinkHubsClientListByResourceGroupOptions contains the optional parameters for the PrivateLinkHubsClient.ListByResourceGroup
// method.
func (client *PrivateLinkHubsClient) ListByResourceGroup(resourceGroupName string, options *PrivateLinkHubsClientListByResourceGroupOptions) *PrivateLinkHubsClientListByResourceGroupPager {
	return &PrivateLinkHubsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkHubsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkHubInfoListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkHubsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkHubsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkHubsClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkHubsClientListByResourceGroupResponse, error) {
	result := PrivateLinkHubsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkHubInfoListResult); err != nil {
		return PrivateLinkHubsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a privateLinkHub
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateLinkHubName - Name of the privateLinkHub
// privateLinkHubPatchInfo - PrivateLinkHub patch request properties
// options - PrivateLinkHubsClientUpdateOptions contains the optional parameters for the PrivateLinkHubsClient.Update method.
func (client *PrivateLinkHubsClient) Update(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkHubPatchInfo PrivateLinkHubPatchInfo, options *PrivateLinkHubsClientUpdateOptions) (PrivateLinkHubsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateLinkHubName, privateLinkHubPatchInfo, options)
	if err != nil {
		return PrivateLinkHubsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkHubsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrivateLinkHubsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PrivateLinkHubsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkHubPatchInfo PrivateLinkHubPatchInfo, options *PrivateLinkHubsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateLinkHubPatchInfo)
}

// updateHandleResponse handles the Update response.
func (client *PrivateLinkHubsClient) updateHandleResponse(resp *http.Response) (PrivateLinkHubsClientUpdateResponse, error) {
	result := PrivateLinkHubsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkHub); err != nil {
		return PrivateLinkHubsClientUpdateResponse{}, err
	}
	return result, nil
}
