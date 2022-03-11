//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// IPGroupsClient contains the methods for the IPGroups group.
// Don't use this type directly, use NewIPGroupsClient() instead.
type IPGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIPGroupsClient creates a new instance of IPGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIPGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IPGroupsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IPGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ipGroupsName - The name of the ipGroups.
// parameters - Parameters supplied to the create or update IpGroups operation.
// options - IPGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the IPGroupsClient.BeginCreateOrUpdate
// method.
func (client *IPGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (IPGroupsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return IPGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result := IPGroupsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("IPGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return IPGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IPGroupsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IPGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
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
func (client *IPGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ipGroupsName - The name of the ipGroups.
// options - IPGroupsClientBeginDeleteOptions contains the optional parameters for the IPGroupsClient.BeginDelete method.
func (client *IPGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (IPGroupsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return IPGroupsClientDeletePollerResponse{}, err
	}
	result := IPGroupsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("IPGroupsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return IPGroupsClientDeletePollerResponse{}, err
	}
	result.Poller = &IPGroupsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IPGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
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
func (client *IPGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ipGroupsName - The name of the ipGroups.
// options - IPGroupsClientGetOptions contains the optional parameters for the IPGroupsClient.Get method.
func (client *IPGroupsClient) Get(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientGetOptions) (IPGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPGroupsClient) getHandleResponse(resp *http.Response) (IPGroupsClientGetResponse, error) {
	result := IPGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all IpGroups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - IPGroupsClientListOptions contains the optional parameters for the IPGroupsClient.List method.
func (client *IPGroupsClient) List(options *IPGroupsClientListOptions) *IPGroupsClientListPager {
	return &IPGroupsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IPGroupsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IPGroupsClient) listCreateRequest(ctx context.Context, options *IPGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ipGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IPGroupsClient) listHandleResponse(resp *http.Response) (IPGroupsClientListResponse, error) {
	result := IPGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all IpGroups in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - IPGroupsClientListByResourceGroupOptions contains the optional parameters for the IPGroupsClient.ListByResourceGroup
// method.
func (client *IPGroupsClient) ListByResourceGroup(resourceGroupName string, options *IPGroupsClientListByResourceGroupOptions) *IPGroupsClientListByResourceGroupPager {
	return &IPGroupsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IPGroupsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (IPGroupsClientListByResourceGroupResponse, error) {
	result := IPGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateGroups - Updates tags of an IpGroups resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ipGroupsName - The name of the ipGroups.
// parameters - Parameters supplied to the update ipGroups operation.
// options - IPGroupsClientUpdateGroupsOptions contains the optional parameters for the IPGroupsClient.UpdateGroups method.
func (client *IPGroupsClient) UpdateGroups(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsClientUpdateGroupsOptions) (IPGroupsClientUpdateGroupsResponse, error) {
	req, err := client.updateGroupsCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPGroupsClientUpdateGroupsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateGroupsHandleResponse(resp)
}

// updateGroupsCreateRequest creates the UpdateGroups request.
func (client *IPGroupsClient) updateGroupsCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsClientUpdateGroupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateGroupsHandleResponse handles the UpdateGroups response.
func (client *IPGroupsClient) updateGroupsHandleResponse(resp *http.Response) (IPGroupsClientUpdateGroupsResponse, error) {
	result := IPGroupsClientUpdateGroupsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	return result, nil
}
