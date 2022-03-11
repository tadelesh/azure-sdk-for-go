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

// ExtendedSQLPoolBlobAuditingPoliciesClient contains the methods for the ExtendedSQLPoolBlobAuditingPolicies group.
// Don't use this type directly, use NewExtendedSQLPoolBlobAuditingPoliciesClient() instead.
type ExtendedSQLPoolBlobAuditingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtendedSQLPoolBlobAuditingPoliciesClient creates a new instance of ExtendedSQLPoolBlobAuditingPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtendedSQLPoolBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExtendedSQLPoolBlobAuditingPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExtendedSQLPoolBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates an extended Sql pool's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// parameters - The extended Sql pool blob auditing policy.
// options - ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ExtendedSQLPoolBlobAuditingPoliciesClient.CreateOrUpdate
// method.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters ExtendedSQLPoolBlobAuditingPolicy, options *ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions) (ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, parameters, options)
	if err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters ExtendedSQLPoolBlobAuditingPolicy, options *ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings/{blobAuditingPolicyName}"
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
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
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
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse, error) {
	result := ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedSQLPoolBlobAuditingPolicy); err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets an extended Sql pool's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// options - ExtendedSQLPoolBlobAuditingPoliciesClientGetOptions contains the optional parameters for the ExtendedSQLPoolBlobAuditingPoliciesClient.Get
// method.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *ExtendedSQLPoolBlobAuditingPoliciesClientGetOptions) (ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *ExtendedSQLPoolBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings/{blobAuditingPolicyName}"
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
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape("default"))
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
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse, error) {
	result := ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedSQLPoolBlobAuditingPolicy); err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListBySQLPool - Lists extended auditing settings of a Sql pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// options - ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolOptions contains the optional parameters for the ExtendedSQLPoolBlobAuditingPoliciesClient.ListBySQLPool
// method.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPool(resourceGroupName string, workspaceName string, sqlPoolName string, options *ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolOptions) *ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolPager {
	return &ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySQLPoolCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
		},
		advancer: func(ctx context.Context, resp ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtendedSQLPoolBlobAuditingPolicyListResult.NextLink)
		},
	}
}

// listBySQLPoolCreateRequest creates the ListBySQLPool request.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) listBySQLPoolCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings"
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

// listBySQLPoolHandleResponse handles the ListBySQLPool response.
func (client *ExtendedSQLPoolBlobAuditingPoliciesClient) listBySQLPoolHandleResponse(resp *http.Response) (ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolResponse, error) {
	result := ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedSQLPoolBlobAuditingPolicyListResult); err != nil {
		return ExtendedSQLPoolBlobAuditingPoliciesClientListBySQLPoolResponse{}, err
	}
	return result, nil
}
