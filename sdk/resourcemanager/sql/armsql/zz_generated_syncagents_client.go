//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SyncAgentsClient contains the methods for the SyncAgents group.
// Don't use this type directly, use NewSyncAgentsClient() instead.
type SyncAgentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSyncAgentsClient creates a new instance of SyncAgentsClient with the specified values.
func NewSyncAgentsClient(con *arm.Connection, subscriptionID string) *SyncAgentsClient {
	return &SyncAgentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsBeginCreateOrUpdateOptions) (SyncAgentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, syncAgentName, parameters, options)
	if err != nil {
		return SyncAgentsCreateOrUpdatePollerResponse{}, err
	}
	result := SyncAgentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SyncAgentsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SyncAgentsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SyncAgentsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SyncAgentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SyncAgentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsBeginDeleteOptions) (SyncAgentsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return SyncAgentsDeletePollerResponse{}, err
	}
	result := SyncAgentsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SyncAgentsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SyncAgentsDeletePollerResponse{}, err
	}
	result.Poller = &SyncAgentsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SyncAgentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SyncAgentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GenerateKey - Generates a sync agent key.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) GenerateKey(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsGenerateKeyOptions) (SyncAgentsGenerateKeyResponse, error) {
	req, err := client.generateKeyCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return SyncAgentsGenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncAgentsGenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncAgentsGenerateKeyResponse{}, client.generateKeyHandleError(resp)
	}
	return client.generateKeyHandleResponse(resp)
}

// generateKeyCreateRequest creates the GenerateKey request.
func (client *SyncAgentsClient) generateKeyCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsGenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/generateKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateKeyHandleResponse handles the GenerateKey response.
func (client *SyncAgentsClient) generateKeyHandleResponse(resp *http.Response) (SyncAgentsGenerateKeyResponse, error) {
	result := SyncAgentsGenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentKeyProperties); err != nil {
		return SyncAgentsGenerateKeyResponse{}, err
	}
	return result, nil
}

// generateKeyHandleError handles the GenerateKey error response.
func (client *SyncAgentsClient) generateKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) Get(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsGetOptions) (SyncAgentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return SyncAgentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncAgentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncAgentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SyncAgentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SyncAgentsClient) getHandleResponse(resp *http.Response) (SyncAgentsGetResponse, error) {
	result := SyncAgentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgent); err != nil {
		return SyncAgentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SyncAgentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Lists sync agents in a server.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) ListByServer(resourceGroupName string, serverName string, options *SyncAgentsListByServerOptions) *SyncAgentsListByServerPager {
	return &SyncAgentsListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp SyncAgentsListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SyncAgentListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *SyncAgentsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *SyncAgentsListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *SyncAgentsClient) listByServerHandleResponse(resp *http.Response) (SyncAgentsListByServerResponse, error) {
	result := SyncAgentsListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentListResult); err != nil {
		return SyncAgentsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *SyncAgentsClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListLinkedDatabases - Lists databases linked to a sync agent.
// If the operation fails it returns a generic error.
func (client *SyncAgentsClient) ListLinkedDatabases(resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsListLinkedDatabasesOptions) *SyncAgentsListLinkedDatabasesPager {
	return &SyncAgentsListLinkedDatabasesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listLinkedDatabasesCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
		},
		advancer: func(ctx context.Context, resp SyncAgentsListLinkedDatabasesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SyncAgentLinkedDatabaseListResult.NextLink)
		},
	}
}

// listLinkedDatabasesCreateRequest creates the ListLinkedDatabases request.
func (client *SyncAgentsClient) listLinkedDatabasesCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsListLinkedDatabasesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/linkedDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listLinkedDatabasesHandleResponse handles the ListLinkedDatabases response.
func (client *SyncAgentsClient) listLinkedDatabasesHandleResponse(resp *http.Response) (SyncAgentsListLinkedDatabasesResponse, error) {
	result := SyncAgentsListLinkedDatabasesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentLinkedDatabaseListResult); err != nil {
		return SyncAgentsListLinkedDatabasesResponse{}, err
	}
	return result, nil
}

// listLinkedDatabasesHandleError handles the ListLinkedDatabases error response.
func (client *SyncAgentsClient) listLinkedDatabasesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
