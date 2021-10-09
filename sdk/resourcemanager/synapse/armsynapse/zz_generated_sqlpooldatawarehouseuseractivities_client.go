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

// SQLPoolDataWarehouseUserActivitiesClient contains the methods for the SQLPoolDataWarehouseUserActivities group.
// Don't use this type directly, use NewSQLPoolDataWarehouseUserActivitiesClient() instead.
type SQLPoolDataWarehouseUserActivitiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolDataWarehouseUserActivitiesClient creates a new instance of SQLPoolDataWarehouseUserActivitiesClient with the specified values.
func NewSQLPoolDataWarehouseUserActivitiesClient(con *arm.Connection, subscriptionID string) *SQLPoolDataWarehouseUserActivitiesClient {
	return &SQLPoolDataWarehouseUserActivitiesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets the user activities of a SQL pool which includes running and suspended queries
// If the operation fails it returns a generic error.
func (client *SQLPoolDataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *SQLPoolDataWarehouseUserActivitiesGetOptions) (SQLPoolDataWarehouseUserActivitiesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, dataWarehouseUserActivityName, options)
	if err != nil {
		return SQLPoolDataWarehouseUserActivitiesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolDataWarehouseUserActivitiesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolDataWarehouseUserActivitiesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolDataWarehouseUserActivitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *SQLPoolDataWarehouseUserActivitiesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}"
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
	if dataWarehouseUserActivityName == "" {
		return nil, errors.New("parameter dataWarehouseUserActivityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataWarehouseUserActivityName}", url.PathEscape(string(dataWarehouseUserActivityName)))
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
func (client *SQLPoolDataWarehouseUserActivitiesClient) getHandleResponse(resp *http.Response) (SQLPoolDataWarehouseUserActivitiesGetResponse, error) {
	result := SQLPoolDataWarehouseUserActivitiesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataWarehouseUserActivities); err != nil {
		return SQLPoolDataWarehouseUserActivitiesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLPoolDataWarehouseUserActivitiesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
