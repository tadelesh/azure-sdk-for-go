//go:build go1.18
// +build go1.18

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

// ExportConfigurationsClient contains the methods for the ExportConfigurations group.
// Don't use this type directly, use NewExportConfigurationsClient() instead.
type ExportConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExportConfigurationsClient creates a new instance of ExportConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExportConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExportConfigurationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExportConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Create a Continuous Export configuration of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// exportProperties - Properties that need to be specified to create a Continuous Export configuration of a Application Insights
// component.
// options - ExportConfigurationsClientCreateOptions contains the optional parameters for the ExportConfigurationsClient.Create
// method.
func (client *ExportConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ComponentExportRequest, options *ExportConfigurationsClientCreateOptions) (ExportConfigurationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, exportProperties, options)
	if err != nil {
		return ExportConfigurationsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ExportConfigurationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ComponentExportRequest, options *ExportConfigurationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration"
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
	return req, runtime.MarshalAsJSON(req, exportProperties)
}

// createHandleResponse handles the Create response.
func (client *ExportConfigurationsClient) createHandleResponse(resp *http.Response) (ExportConfigurationsClientCreateResponse, error) {
	result := ExportConfigurationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentExportConfigurationArray); err != nil {
		return ExportConfigurationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Continuous Export configuration of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// exportID - The Continuous Export configuration ID. This is unique within a Application Insights component.
// options - ExportConfigurationsClientDeleteOptions contains the optional parameters for the ExportConfigurationsClient.Delete
// method.
func (client *ExportConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsClientDeleteOptions) (ExportConfigurationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, exportID, options)
	if err != nil {
		return ExportConfigurationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ExportConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
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
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ExportConfigurationsClient) deleteHandleResponse(resp *http.Response) (ExportConfigurationsClientDeleteResponse, error) {
	result := ExportConfigurationsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentExportConfiguration); err != nil {
		return ExportConfigurationsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get the Continuous Export configuration for this export id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// exportID - The Continuous Export configuration ID. This is unique within a Application Insights component.
// options - ExportConfigurationsClientGetOptions contains the optional parameters for the ExportConfigurationsClient.Get
// method.
func (client *ExportConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsClientGetOptions) (ExportConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, exportID, options)
	if err != nil {
		return ExportConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExportConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, options *ExportConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
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
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
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

// getHandleResponse handles the Get response.
func (client *ExportConfigurationsClient) getHandleResponse(resp *http.Response) (ExportConfigurationsClientGetResponse, error) {
	result := ExportConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentExportConfiguration); err != nil {
		return ExportConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of Continuous Export configuration of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ExportConfigurationsClientListOptions contains the optional parameters for the ExportConfigurationsClient.List
// method.
func (client *ExportConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *ExportConfigurationsClientListOptions) (ExportConfigurationsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ExportConfigurationsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ExportConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ExportConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration"
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
func (client *ExportConfigurationsClient) listHandleResponse(resp *http.Response) (ExportConfigurationsClientListResponse, error) {
	result := ExportConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentExportConfigurationArray); err != nil {
		return ExportConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update the Continuous Export configuration for this export id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// exportID - The Continuous Export configuration ID. This is unique within a Application Insights component.
// exportProperties - Properties that need to be specified to update the Continuous Export configuration.
// options - ExportConfigurationsClientUpdateOptions contains the optional parameters for the ExportConfigurationsClient.Update
// method.
func (client *ExportConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ComponentExportRequest, options *ExportConfigurationsClientUpdateOptions) (ExportConfigurationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, exportID, exportProperties, options)
	if err != nil {
		return ExportConfigurationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportConfigurationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportConfigurationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ExportConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ComponentExportRequest, options *ExportConfigurationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}"
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
	if exportID == "" {
		return nil, errors.New("parameter exportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportId}", url.PathEscape(exportID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, exportProperties)
}

// updateHandleResponse handles the Update response.
func (client *ExportConfigurationsClient) updateHandleResponse(resp *http.Response) (ExportConfigurationsClientUpdateResponse, error) {
	result := ExportConfigurationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentExportConfiguration); err != nil {
		return ExportConfigurationsClientUpdateResponse{}, err
	}
	return result, nil
}
