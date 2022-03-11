//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

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

// OrganizationClient contains the methods for the Organization group.
// Don't use this type directly, use NewOrganizationClient() instead.
type OrganizationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOrganizationClient creates a new instance of OrganizationClient with the specified values.
// subscriptionID - Microsoft Azure subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOrganizationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OrganizationClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OrganizationClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Create Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientBeginCreateOptions contains the optional parameters for the OrganizationClient.BeginCreate
// method.
func (client *OrganizationClient) BeginCreate(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (OrganizationClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientCreatePollerResponse{}, err
	}
	result := OrganizationClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("OrganizationClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return OrganizationClientCreatePollerResponse{}, err
	}
	result.Poller = &OrganizationClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OrganizationClient) create(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, organizationName, options)
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
func (client *OrganizationClient) createCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Delete Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientBeginDeleteOptions contains the optional parameters for the OrganizationClient.BeginDelete
// method.
func (client *OrganizationClient) BeginDelete(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (OrganizationClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientDeletePollerResponse{}, err
	}
	result := OrganizationClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("OrganizationClient.Delete", "location", resp, client.pl)
	if err != nil {
		return OrganizationClientDeletePollerResponse{}, err
	}
	result.Poller = &OrganizationClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OrganizationClient) deleteOperation(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, options)
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
func (client *OrganizationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the properties of a specific Organization resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientGetOptions contains the optional parameters for the OrganizationClient.Get method.
func (client *OrganizationClient) Get(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientGetOptions) (OrganizationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrganizationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrganizationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OrganizationClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrganizationClient) getHandleResponse(resp *http.Response) (OrganizationClientGetResponse, error) {
	result := OrganizationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResource); err != nil {
		return OrganizationClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List all Organizations under the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// options - OrganizationClientListByResourceGroupOptions contains the optional parameters for the OrganizationClient.ListByResourceGroup
// method.
func (client *OrganizationClient) ListByResourceGroup(resourceGroupName string, options *OrganizationClientListByResourceGroupOptions) *OrganizationClientListByResourceGroupPager {
	return &OrganizationClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp OrganizationClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OrganizationResourceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrganizationClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrganizationClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations"
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
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OrganizationClient) listByResourceGroupHandleResponse(resp *http.Response) (OrganizationClientListByResourceGroupResponse, error) {
	result := OrganizationClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResourceListResult); err != nil {
		return OrganizationClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - List all organizations under the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OrganizationClientListBySubscriptionOptions contains the optional parameters for the OrganizationClient.ListBySubscription
// method.
func (client *OrganizationClient) ListBySubscription(options *OrganizationClientListBySubscriptionOptions) *OrganizationClientListBySubscriptionPager {
	return &OrganizationClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp OrganizationClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OrganizationResourceListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrganizationClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrganizationClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/organizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OrganizationClient) listBySubscriptionHandleResponse(resp *http.Response) (OrganizationClientListBySubscriptionResponse, error) {
	result := OrganizationClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResourceListResult); err != nil {
		return OrganizationClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientUpdateOptions contains the optional parameters for the OrganizationClient.Update method.
func (client *OrganizationClient) Update(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientUpdateOptions) (OrganizationClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrganizationClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *OrganizationClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *OrganizationClient) updateHandleResponse(resp *http.Response) (OrganizationClientUpdateResponse, error) {
	result := OrganizationClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResource); err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	return result, nil
}
