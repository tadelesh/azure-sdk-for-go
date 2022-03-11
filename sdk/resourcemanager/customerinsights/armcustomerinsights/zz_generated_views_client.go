//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// ViewsClient contains the methods for the Views group.
// Don't use this type directly, use NewViewsClient() instead.
type ViewsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewViewsClient creates a new instance of ViewsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewViewsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ViewsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ViewsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates a view or updates an existing view in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// viewName - The name of the view.
// parameters - Parameters supplied to the CreateOrUpdate View operation.
// options - ViewsClientCreateOrUpdateOptions contains the optional parameters for the ViewsClient.CreateOrUpdate method.
func (client *ViewsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, viewName string, parameters ViewResourceFormat, options *ViewsClientCreateOrUpdateOptions) (ViewsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, viewName, parameters, options)
	if err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ViewsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ViewsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, viewName string, parameters ViewResourceFormat, options *ViewsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ViewsClient) createOrUpdateHandleResponse(resp *http.Response) (ViewsClientCreateOrUpdateResponse, error) {
	result := ViewsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViewResourceFormat); err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a view in the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// viewName - The name of the view.
// userID - The user ID. Use * to retrieve hub level view.
// options - ViewsClientDeleteOptions contains the optional parameters for the ViewsClient.Delete method.
func (client *ViewsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string, options *ViewsClientDeleteOptions) (ViewsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, viewName, userID, options)
	if err != nil {
		return ViewsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ViewsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ViewsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ViewsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ViewsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string, options *ViewsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	reqQP.Set("userId", userID)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a view in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// viewName - The name of the view.
// userID - The user ID. Use * to retrieve hub level view.
// options - ViewsClientGetOptions contains the optional parameters for the ViewsClient.Get method.
func (client *ViewsClient) Get(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string, options *ViewsClientGetOptions) (ViewsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, viewName, userID, options)
	if err != nil {
		return ViewsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ViewsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ViewsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ViewsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string, options *ViewsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	reqQP.Set("userId", userID)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ViewsClient) getHandleResponse(resp *http.Response) (ViewsClientGetResponse, error) {
	result := ViewsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViewResourceFormat); err != nil {
		return ViewsClientGetResponse{}, err
	}
	return result, nil
}

// ListByHub - Gets all available views for given user in the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// userID - The user ID. Use * to retrieve hub level views.
// options - ViewsClientListByHubOptions contains the optional parameters for the ViewsClient.ListByHub method.
func (client *ViewsClient) ListByHub(resourceGroupName string, hubName string, userID string, options *ViewsClientListByHubOptions) *ViewsClientListByHubPager {
	return &ViewsClientListByHubPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHubCreateRequest(ctx, resourceGroupName, hubName, userID, options)
		},
		advancer: func(ctx context.Context, resp ViewsClientListByHubResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ViewListResult.NextLink)
		},
	}
}

// listByHubCreateRequest creates the ListByHub request.
func (client *ViewsClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, userID string, options *ViewsClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	reqQP.Set("userId", userID)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *ViewsClient) listByHubHandleResponse(resp *http.Response) (ViewsClientListByHubResponse, error) {
	result := ViewsClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViewListResult); err != nil {
		return ViewsClientListByHubResponse{}, err
	}
	return result, nil
}
