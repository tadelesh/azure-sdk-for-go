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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// WorkspaceManagedSQLServerBlobAuditingPoliciesClient contains the methods for the WorkspaceManagedSQLServerBlobAuditingPolicies group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient() instead.
type WorkspaceManagedSQLServerBlobAuditingPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerBlobAuditingPoliciesClient with the specified values.
func NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient(con *arm.Connection, subscriptionID string) *WorkspaceManagedSQLServerBlobAuditingPoliciesClient {
	return &WorkspaceManagedSQLServerBlobAuditingPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (WorkspaceManagedSQLServerBlobAuditingPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedSQLServerBlobAuditingPoliciesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedSQLServerBlobAuditingPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a workspace managed sql server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesGetOptions) (WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
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
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicy); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByWorkspace - List workspace managed sql server's blob auditing policies.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceOptions) *WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspacePager {
	return &WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerBlobAuditingPolicyListResult.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings"
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

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesListByWorkspaceResponse{}, err
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
