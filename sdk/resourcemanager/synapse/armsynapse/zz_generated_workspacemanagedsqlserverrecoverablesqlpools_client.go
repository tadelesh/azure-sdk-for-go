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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// WorkspaceManagedSQLServerRecoverableSQLPoolsClient contains the methods for the WorkspaceManagedSQLServerRecoverableSQLPools group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient() instead.
type WorkspaceManagedSQLServerRecoverableSQLPoolsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient creates a new instance of WorkspaceManagedSQLServerRecoverableSQLPoolsClient with the specified values.
func NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient(con *arm.Connection, subscriptionID string) *WorkspaceManagedSQLServerRecoverableSQLPoolsClient {
	return &WorkspaceManagedSQLServerRecoverableSQLPoolsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get recoverable sql pools for workspace managed sql server.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsGetOptions) (WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse, error) {
	result := WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableSQLPool); err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Get list of recoverable sql pools for workspace managed sql server.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) List(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsListOptions) *WorkspaceManagedSQLServerRecoverableSQLPoolsListPager {
	return &WorkspaceManagedSQLServerRecoverableSQLPoolsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerRecoverableSQLPoolsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecoverableSQLPoolListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerRecoverableSQLPoolsListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerRecoverableSQLPoolsListResponse, error) {
	result := WorkspaceManagedSQLServerRecoverableSQLPoolsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableSQLPoolListResult); err != nil {
		return WorkspaceManagedSQLServerRecoverableSQLPoolsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkspaceManagedSQLServerRecoverableSQLPoolsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
