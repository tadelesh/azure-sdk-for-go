//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

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

// CustomResourceProviderClient contains the methods for the CustomResourceProvider group.
// Don't use this type directly, use NewCustomResourceProviderClient() instead.
type CustomResourceProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCustomResourceProviderClient creates a new instance of CustomResourceProviderClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCustomResourceProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CustomResourceProviderClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CustomResourceProviderClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the custom resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceProviderName - The name of the resource provider.
// resourceProvider - The parameters required to create or update a custom resource provider definition.
// options - CustomResourceProviderClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomResourceProviderClient.BeginCreateOrUpdate
// method.
func (client *CustomResourceProviderClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderClientBeginCreateOrUpdateOptions) (CustomResourceProviderClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceProviderName, resourceProvider, options)
	if err != nil {
		return CustomResourceProviderClientCreateOrUpdatePollerResponse{}, err
	}
	result := CustomResourceProviderClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("CustomResourceProviderClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return CustomResourceProviderClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomResourceProviderClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the custom resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomResourceProviderClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceProviderName, resourceProvider, options)
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
func (client *CustomResourceProviderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceProvider)
}

// BeginDelete - Deletes the custom resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceProviderName - The name of the resource provider.
// options - CustomResourceProviderClientBeginDeleteOptions contains the optional parameters for the CustomResourceProviderClient.BeginDelete
// method.
func (client *CustomResourceProviderClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderClientBeginDeleteOptions) (CustomResourceProviderClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceProviderName, options)
	if err != nil {
		return CustomResourceProviderClientDeletePollerResponse{}, err
	}
	result := CustomResourceProviderClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("CustomResourceProviderClient.Delete", "", resp, client.pl)
	if err != nil {
		return CustomResourceProviderClientDeletePollerResponse{}, err
	}
	result.Poller = &CustomResourceProviderClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the custom resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomResourceProviderClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceProviderName, options)
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
func (client *CustomResourceProviderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the custom resource provider manifest.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceProviderName - The name of the resource provider.
// options - CustomResourceProviderClientGetOptions contains the optional parameters for the CustomResourceProviderClient.Get
// method.
func (client *CustomResourceProviderClient) Get(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderClientGetOptions) (CustomResourceProviderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceProviderName, options)
	if err != nil {
		return CustomResourceProviderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomResourceProviderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomResourceProviderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomResourceProviderClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomResourceProviderClient) getHandleResponse(resp *http.Response) (CustomResourceProviderClientGetResponse, error) {
	result := CustomResourceProviderClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRPManifest); err != nil {
		return CustomResourceProviderClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all the custom resource providers within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - CustomResourceProviderClientListByResourceGroupOptions contains the optional parameters for the CustomResourceProviderClient.ListByResourceGroup
// method.
func (client *CustomResourceProviderClient) ListByResourceGroup(resourceGroupName string, options *CustomResourceProviderClientListByResourceGroupOptions) *CustomResourceProviderClientListByResourceGroupPager {
	return &CustomResourceProviderClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomResourceProviderClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListByCustomRPManifest.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomResourceProviderClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomResourceProviderClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomResourceProviderClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomResourceProviderClientListByResourceGroupResponse, error) {
	result := CustomResourceProviderClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListByCustomRPManifest); err != nil {
		return CustomResourceProviderClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets all the custom resource providers within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - CustomResourceProviderClientListBySubscriptionOptions contains the optional parameters for the CustomResourceProviderClient.ListBySubscription
// method.
func (client *CustomResourceProviderClient) ListBySubscription(options *CustomResourceProviderClientListBySubscriptionOptions) *CustomResourceProviderClientListBySubscriptionPager {
	return &CustomResourceProviderClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomResourceProviderClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListByCustomRPManifest.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomResourceProviderClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomResourceProviderClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomProviders/resourceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomResourceProviderClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomResourceProviderClientListBySubscriptionResponse, error) {
	result := CustomResourceProviderClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListByCustomRPManifest); err != nil {
		return CustomResourceProviderClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing custom resource provider. The only value that can be updated via PATCH currently is the tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceProviderName - The name of the resource provider.
// patchableResource - The updatable fields of a custom resource provider.
// options - CustomResourceProviderClientUpdateOptions contains the optional parameters for the CustomResourceProviderClient.Update
// method.
func (client *CustomResourceProviderClient) Update(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate, options *CustomResourceProviderClientUpdateOptions) (CustomResourceProviderClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceProviderName, patchableResource, options)
	if err != nil {
		return CustomResourceProviderClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomResourceProviderClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomResourceProviderClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CustomResourceProviderClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate, options *CustomResourceProviderClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, patchableResource)
}

// updateHandleResponse handles the Update response.
func (client *CustomResourceProviderClient) updateHandleResponse(resp *http.Response) (CustomResourceProviderClientUpdateResponse, error) {
	result := CustomResourceProviderClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRPManifest); err != nil {
		return CustomResourceProviderClientUpdateResponse{}, err
	}
	return result, nil
}
