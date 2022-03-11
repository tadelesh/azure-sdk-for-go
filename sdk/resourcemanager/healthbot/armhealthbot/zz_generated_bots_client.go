//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthbot

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

// BotsClient contains the methods for the Bots group.
// Don't use this type directly, use NewBotsClient() instead.
type BotsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBotsClient creates a new instance of BotsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BotsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BotsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Create a new Azure Health Bot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Bot resource group in the user subscription.
// botName - The name of the Bot resource.
// parameters - The parameters to provide for the created Azure Health Bot.
// options - BotsClientBeginCreateOptions contains the optional parameters for the BotsClient.BeginCreate method.
func (client *BotsClient) BeginCreate(ctx context.Context, resourceGroupName string, botName string, parameters HealthBot, options *BotsClientBeginCreateOptions) (BotsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, botName, parameters, options)
	if err != nil {
		return BotsClientCreatePollerResponse{}, err
	}
	result := BotsClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("BotsClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return BotsClientCreatePollerResponse{}, err
	}
	result.Poller = &BotsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a new Azure Health Bot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BotsClient) create(ctx context.Context, resourceGroupName string, botName string, parameters HealthBot, options *BotsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, botName, parameters, options)
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
func (client *BotsClient) createCreateRequest(ctx context.Context, resourceGroupName string, botName string, parameters HealthBot, options *BotsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthBot/healthBots/{botName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if botName == "" {
		return nil, errors.New("parameter botName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{botName}", url.PathEscape(botName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a HealthBot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Bot resource group in the user subscription.
// botName - The name of the Bot resource.
// options - BotsClientBeginDeleteOptions contains the optional parameters for the BotsClient.BeginDelete method.
func (client *BotsClient) BeginDelete(ctx context.Context, resourceGroupName string, botName string, options *BotsClientBeginDeleteOptions) (BotsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, botName, options)
	if err != nil {
		return BotsClientDeletePollerResponse{}, err
	}
	result := BotsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("BotsClient.Delete", "", resp, client.pl)
	if err != nil {
		return BotsClientDeletePollerResponse{}, err
	}
	result.Poller = &BotsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a HealthBot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BotsClient) deleteOperation(ctx context.Context, resourceGroupName string, botName string, options *BotsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, botName, options)
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
func (client *BotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, botName string, options *BotsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthBot/healthBots/{botName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if botName == "" {
		return nil, errors.New("parameter botName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{botName}", url.PathEscape(botName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a HealthBot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Bot resource group in the user subscription.
// botName - The name of the Bot resource.
// options - BotsClientGetOptions contains the optional parameters for the BotsClient.Get method.
func (client *BotsClient) Get(ctx context.Context, resourceGroupName string, botName string, options *BotsClientGetOptions) (BotsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, botName, options)
	if err != nil {
		return BotsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, botName string, options *BotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthBot/healthBots/{botName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if botName == "" {
		return nil, errors.New("parameter botName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{botName}", url.PathEscape(botName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BotsClient) getHandleResponse(resp *http.Response) (BotsClientGetResponse, error) {
	result := BotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthBot); err != nil {
		return BotsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns all the resources of a particular type belonging to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - BotsClientListOptions contains the optional parameters for the BotsClient.List method.
func (client *BotsClient) List(options *BotsClientListOptions) *BotsClientListPager {
	return &BotsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp BotsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BotResponseList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BotsClient) listCreateRequest(ctx context.Context, options *BotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HealthBot/healthBots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BotsClient) listHandleResponse(resp *http.Response) (BotsClientListResponse, error) {
	result := BotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotResponseList); err != nil {
		return BotsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Returns all the resources of a particular type belonging to a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Bot resource group in the user subscription.
// options - BotsClientListByResourceGroupOptions contains the optional parameters for the BotsClient.ListByResourceGroup
// method.
func (client *BotsClient) ListByResourceGroup(resourceGroupName string, options *BotsClientListByResourceGroupOptions) *BotsClientListByResourceGroupPager {
	return &BotsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp BotsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BotResponseList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *BotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthBot/healthBots"
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
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BotsClient) listByResourceGroupHandleResponse(resp *http.Response) (BotsClientListByResourceGroupResponse, error) {
	result := BotsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotResponseList); err != nil {
		return BotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Patch a HealthBot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Bot resource group in the user subscription.
// botName - The name of the Bot resource.
// parameters - The parameters to provide for the required Azure Health Bot.
// options - BotsClientUpdateOptions contains the optional parameters for the BotsClient.Update method.
func (client *BotsClient) Update(ctx context.Context, resourceGroupName string, botName string, parameters UpdateParameters, options *BotsClientUpdateOptions) (BotsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, botName, parameters, options)
	if err != nil {
		return BotsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *BotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, botName string, parameters UpdateParameters, options *BotsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthBot/healthBots/{botName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if botName == "" {
		return nil, errors.New("parameter botName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{botName}", url.PathEscape(botName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *BotsClient) updateHandleResponse(resp *http.Response) (BotsClientUpdateResponse, error) {
	result := BotsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthBot); err != nil {
		return BotsClientUpdateResponse{}, err
	}
	return result, nil
}
