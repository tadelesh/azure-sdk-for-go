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

// WorkspaceManagedSQLServerRecoverableSQLPoolsClient contains the methods for the WorkspaceManagedSQLServerRecoverableSQLPools group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient() instead.
type WorkspaceManagedSQLServerRecoverableSQLPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient creates a new instance of WorkspaceManagedSQLServerRecoverableSQLPoolsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceManagedSQLServerRecoverableSQLPoolsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkspaceManagedSQLServerRecoverableSQLPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get recoverable sql pools for workspace managed sql server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - The name of the sql pool
// options - WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetOptions contains the optional parameters for the WorkspaceManagedSQLServerRecoverableSQLPoolsClient.Get
// method.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetOptions) (WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/recoverableSqlPools/{sqlPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
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
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse, error) {
	result := WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableSQLPool); err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - Get list of recoverable sql pools for workspace managed sql server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceManagedSQLServerRecoverableSQLPoolsClientListOptions contains the optional parameters for the WorkspaceManagedSQLServerRecoverableSQLPoolsClient.List
// method.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) List(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsClientListOptions) *WorkspaceManagedSQLServerRecoverableSQLPoolsClientListPager {
	return &WorkspaceManagedSQLServerRecoverableSQLPoolsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerRecoverableSQLPoolsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecoverableSQLPoolListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/recoverableSqlPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerRecoverableSQLPoolsClientListResponse, error) {
	result := WorkspaceManagedSQLServerRecoverableSQLPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableSQLPoolListResult); err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsClientListResponse{}, err
	}
	return result, nil
}
