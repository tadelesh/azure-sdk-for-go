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

// SQLPoolSecurityAlertPoliciesClient contains the methods for the SQLPoolSecurityAlertPolicies group.
// Don't use this type directly, use NewSQLPoolSecurityAlertPoliciesClient() instead.
type SQLPoolSecurityAlertPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolSecurityAlertPoliciesClient creates a new instance of SQLPoolSecurityAlertPoliciesClient with the specified values.
func NewSQLPoolSecurityAlertPoliciesClient(con *arm.Connection, subscriptionID string) *SQLPoolSecurityAlertPoliciesClient {
	return &SQLPoolSecurityAlertPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a Sql pool's security alert policy.
// If the operation fails it returns a generic error.
func (client *SQLPoolSecurityAlertPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, securityAlertPolicyName SecurityAlertPolicyName, parameters SQLPoolSecurityAlertPolicy, options *SQLPoolSecurityAlertPoliciesCreateOrUpdateOptions) (SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, securityAlertPolicyName SecurityAlertPolicyName, parameters SQLPoolSecurityAlertPolicy, options *SQLPoolSecurityAlertPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SQLPoolSecurityAlertPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse, error) {
	result := SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolSecurityAlertPolicy); err != nil {
		return SQLPoolSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SQLPoolSecurityAlertPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a Sql pool's security alert policy.
// If the operation fails it returns a generic error.
func (client *SQLPoolSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, securityAlertPolicyName SecurityAlertPolicyName, options *SQLPoolSecurityAlertPoliciesGetOptions) (SQLPoolSecurityAlertPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, securityAlertPolicyName, options)
	if err != nil {
		return SQLPoolSecurityAlertPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSecurityAlertPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolSecurityAlertPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, securityAlertPolicyName SecurityAlertPolicyName, options *SQLPoolSecurityAlertPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
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
func (client *SQLPoolSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (SQLPoolSecurityAlertPoliciesGetResponse, error) {
	result := SQLPoolSecurityAlertPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolSecurityAlertPolicy); err != nil {
		return SQLPoolSecurityAlertPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLPoolSecurityAlertPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Get a list of Sql pool's security alert policies.
// If the operation fails it returns a generic error.
func (client *SQLPoolSecurityAlertPoliciesClient) List(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSecurityAlertPoliciesListOptions) *SQLPoolSecurityAlertPoliciesListPager {
	return &SQLPoolSecurityAlertPoliciesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolSecurityAlertPoliciesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListSQLPoolSecurityAlertPolicies.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SQLPoolSecurityAlertPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSecurityAlertPoliciesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies"
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

// listHandleResponse handles the List response.
func (client *SQLPoolSecurityAlertPoliciesClient) listHandleResponse(resp *http.Response) (SQLPoolSecurityAlertPoliciesListResponse, error) {
	result := SQLPoolSecurityAlertPoliciesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListSQLPoolSecurityAlertPolicies); err != nil {
		return SQLPoolSecurityAlertPoliciesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SQLPoolSecurityAlertPoliciesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
