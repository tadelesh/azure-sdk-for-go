//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// AdminKeysClient contains the methods for the AdminKeys group.
// Don't use this type directly, use NewAdminKeysClient() instead.
type AdminKeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAdminKeysClient creates a new instance of AdminKeysClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
// Manager API or the portal.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAdminKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AdminKeysClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AdminKeysClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
// method.
// options - AdminKeysClientGetOptions contains the optional parameters for the AdminKeysClient.Get method.
func (client *AdminKeysClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientGetOptions) (AdminKeysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
	if err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AdminKeysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AdminKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *searchManagementRequestOptions.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdminKeysClient) getHandleResponse(resp *http.Response) (AdminKeysClientGetResponse, error) {
	result := AdminKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminKeyResult); err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	return result, nil
}

// Regenerate - Regenerates either the primary or secondary admin API key. You can only regenerate one key at a time.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// keyKind - Specifies which key to regenerate. Valid values include 'primary' and 'secondary'.
// SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
// method.
// options - AdminKeysClientRegenerateOptions contains the optional parameters for the AdminKeysClient.Regenerate method.
func (client *AdminKeysClient) Regenerate(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientRegenerateOptions) (AdminKeysClientRegenerateResponse, error) {
	req, err := client.regenerateCreateRequest(ctx, resourceGroupName, searchServiceName, keyKind, searchManagementRequestOptions, options)
	if err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AdminKeysClientRegenerateResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateHandleResponse(resp)
}

// regenerateCreateRequest creates the Regenerate request.
func (client *AdminKeysClient) regenerateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientRegenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/regenerateAdminKey/{keyKind}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if keyKind == "" {
		return nil, errors.New("parameter keyKind cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyKind}", url.PathEscape(string(keyKind)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *searchManagementRequestOptions.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// regenerateHandleResponse handles the Regenerate response.
func (client *AdminKeysClient) regenerateHandleResponse(resp *http.Response) (AdminKeysClientRegenerateResponse, error) {
	result := AdminKeysClientRegenerateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminKeyResult); err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	return result, nil
}
