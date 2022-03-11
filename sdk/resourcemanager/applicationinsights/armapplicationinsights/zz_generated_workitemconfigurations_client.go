//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// WorkItemConfigurationsClient contains the methods for the WorkItemConfigurations group.
// Don't use this type directly, use NewWorkItemConfigurationsClient() instead.
type WorkItemConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkItemConfigurationsClient creates a new instance of WorkItemConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkItemConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkItemConfigurationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkItemConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Create a work item configuration for an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workItemConfigurationProperties - Properties that need to be specified to create a work item configuration of a Application
// Insights component.
// options - WorkItemConfigurationsClientCreateOptions contains the optional parameters for the WorkItemConfigurationsClient.Create
// method.
func (client *WorkItemConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration, options *WorkItemConfigurationsClientCreateOptions) (WorkItemConfigurationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, workItemConfigurationProperties, options)
	if err != nil {
		return WorkItemConfigurationsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkItemConfigurationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkItemConfigurationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *WorkItemConfigurationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration, options *WorkItemConfigurationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workItemConfigurationProperties)
}

// createHandleResponse handles the Create response.
func (client *WorkItemConfigurationsClient) createHandleResponse(resp *http.Response) (WorkItemConfigurationsClientCreateResponse, error) {
	result := WorkItemConfigurationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkItemConfiguration); err != nil {
		return WorkItemConfigurationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a work item configuration of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workItemConfigID - The unique work item configuration Id. This can be either friendly name of connector as defined in connector
// configuration
// options - WorkItemConfigurationsClientDeleteOptions contains the optional parameters for the WorkItemConfigurationsClient.Delete
// method.
func (client *WorkItemConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *WorkItemConfigurationsClientDeleteOptions) (WorkItemConfigurationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, workItemConfigID, options)
	if err != nil {
		return WorkItemConfigurationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkItemConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkItemConfigurationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkItemConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkItemConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *WorkItemConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if workItemConfigID == "" {
		return nil, errors.New("parameter workItemConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workItemConfigId}", url.PathEscape(workItemConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// GetDefault - Gets default work item configurations that exist for the application
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - WorkItemConfigurationsClientGetDefaultOptions contains the optional parameters for the WorkItemConfigurationsClient.GetDefault
// method.
func (client *WorkItemConfigurationsClient) GetDefault(ctx context.Context, resourceGroupName string, resourceName string, options *WorkItemConfigurationsClientGetDefaultOptions) (WorkItemConfigurationsClientGetDefaultResponse, error) {
	req, err := client.getDefaultCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return WorkItemConfigurationsClientGetDefaultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkItemConfigurationsClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkItemConfigurationsClientGetDefaultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDefaultHandleResponse(resp)
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *WorkItemConfigurationsClient) getDefaultCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WorkItemConfigurationsClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/DefaultWorkItemConfig"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *WorkItemConfigurationsClient) getDefaultHandleResponse(resp *http.Response) (WorkItemConfigurationsClientGetDefaultResponse, error) {
	result := WorkItemConfigurationsClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkItemConfiguration); err != nil {
		return WorkItemConfigurationsClientGetDefaultResponse{}, err
	}
	return result, nil
}

// GetItem - Gets specified work item configuration for an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workItemConfigID - The unique work item configuration Id. This can be either friendly name of connector as defined in connector
// configuration
// options - WorkItemConfigurationsClientGetItemOptions contains the optional parameters for the WorkItemConfigurationsClient.GetItem
// method.
func (client *WorkItemConfigurationsClient) GetItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *WorkItemConfigurationsClientGetItemOptions) (WorkItemConfigurationsClientGetItemResponse, error) {
	req, err := client.getItemCreateRequest(ctx, resourceGroupName, resourceName, workItemConfigID, options)
	if err != nil {
		return WorkItemConfigurationsClientGetItemResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkItemConfigurationsClientGetItemResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkItemConfigurationsClientGetItemResponse{}, runtime.NewResponseError(resp)
	}
	return client.getItemHandleResponse(resp)
}

// getItemCreateRequest creates the GetItem request.
func (client *WorkItemConfigurationsClient) getItemCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *WorkItemConfigurationsClientGetItemOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if workItemConfigID == "" {
		return nil, errors.New("parameter workItemConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workItemConfigId}", url.PathEscape(workItemConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getItemHandleResponse handles the GetItem response.
func (client *WorkItemConfigurationsClient) getItemHandleResponse(resp *http.Response) (WorkItemConfigurationsClientGetItemResponse, error) {
	result := WorkItemConfigurationsClientGetItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkItemConfiguration); err != nil {
		return WorkItemConfigurationsClientGetItemResponse{}, err
	}
	return result, nil
}

// List - Gets the list work item configurations that exist for the application
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - WorkItemConfigurationsClientListOptions contains the optional parameters for the WorkItemConfigurationsClient.List
// method.
func (client *WorkItemConfigurationsClient) List(resourceGroupName string, resourceName string, options *WorkItemConfigurationsClientListOptions) *WorkItemConfigurationsClientListPager {
	return &WorkItemConfigurationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkItemConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WorkItemConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkItemConfigurationsClient) listHandleResponse(resp *http.Response) (WorkItemConfigurationsClientListResponse, error) {
	result := WorkItemConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkItemConfigurationsListResult); err != nil {
		return WorkItemConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// UpdateItem - Update a work item configuration for an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workItemConfigID - The unique work item configuration Id. This can be either friendly name of connector as defined in connector
// configuration
// workItemConfigurationProperties - Properties that need to be specified to update a work item configuration for this Application
// Insights component.
// options - WorkItemConfigurationsClientUpdateItemOptions contains the optional parameters for the WorkItemConfigurationsClient.UpdateItem
// method.
func (client *WorkItemConfigurationsClient) UpdateItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties WorkItemCreateConfiguration, options *WorkItemConfigurationsClientUpdateItemOptions) (WorkItemConfigurationsClientUpdateItemResponse, error) {
	req, err := client.updateItemCreateRequest(ctx, resourceGroupName, resourceName, workItemConfigID, workItemConfigurationProperties, options)
	if err != nil {
		return WorkItemConfigurationsClientUpdateItemResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkItemConfigurationsClientUpdateItemResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkItemConfigurationsClientUpdateItemResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateItemHandleResponse(resp)
}

// updateItemCreateRequest creates the UpdateItem request.
func (client *WorkItemConfigurationsClient) updateItemCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties WorkItemCreateConfiguration, options *WorkItemConfigurationsClientUpdateItemOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if workItemConfigID == "" {
		return nil, errors.New("parameter workItemConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workItemConfigId}", url.PathEscape(workItemConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workItemConfigurationProperties)
}

// updateItemHandleResponse handles the UpdateItem response.
func (client *WorkItemConfigurationsClient) updateItemHandleResponse(resp *http.Response) (WorkItemConfigurationsClientUpdateItemResponse, error) {
	result := WorkItemConfigurationsClientUpdateItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkItemConfiguration); err != nil {
		return WorkItemConfigurationsClientUpdateItemResponse{}, err
	}
	return result, nil
}
