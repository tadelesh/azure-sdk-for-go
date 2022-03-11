//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// RoleAssignmentsClient contains the methods for the RoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RoleAssignmentsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RoleAssignmentsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a role assignment in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// assignmentName - The assignment name
// parameters - Parameters supplied to the CreateOrUpdate RoleAssignment operation.
// options - RoleAssignmentsClientBeginCreateOrUpdateOptions contains the optional parameters for the RoleAssignmentsClient.BeginCreateOrUpdate
// method.
func (client *RoleAssignmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, parameters RoleAssignmentResourceFormat, options *RoleAssignmentsClientBeginCreateOrUpdateOptions) (RoleAssignmentsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, assignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentsClientCreateOrUpdatePollerResponse{}, err
	}
	result := RoleAssignmentsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("RoleAssignmentsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return RoleAssignmentsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RoleAssignmentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a role assignment in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoleAssignmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, parameters RoleAssignmentResourceFormat, options *RoleAssignmentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, assignmentName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RoleAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, parameters RoleAssignmentResourceFormat, options *RoleAssignmentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes the role assignment in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// assignmentName - The name of the role assignment.
// options - RoleAssignmentsClientDeleteOptions contains the optional parameters for the RoleAssignmentsClient.Delete method.
func (client *RoleAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, options *RoleAssignmentsClientDeleteOptions) (RoleAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, assignmentName, options)
	if err != nil {
		return RoleAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return RoleAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RoleAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, options *RoleAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the role assignment in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// assignmentName - The name of the role assignment.
// options - RoleAssignmentsClientGetOptions contains the optional parameters for the RoleAssignmentsClient.Get method.
func (client *RoleAssignmentsClient) Get(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, options *RoleAssignmentsClientGetOptions) (RoleAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, assignmentName, options)
	if err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, options *RoleAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentsClient) getHandleResponse(resp *http.Response) (RoleAssignmentsClientGetResponse, error) {
	result := RoleAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentResourceFormat); err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// ListByHub - Gets all the role assignments for the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - RoleAssignmentsClientListByHubOptions contains the optional parameters for the RoleAssignmentsClient.ListByHub
// method.
func (client *RoleAssignmentsClient) ListByHub(resourceGroupName string, hubName string, options *RoleAssignmentsClientListByHubOptions) *RoleAssignmentsClientListByHubPager {
	return &RoleAssignmentsClientListByHubPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsClientListByHubResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listByHubCreateRequest creates the ListByHub request.
func (client *RoleAssignmentsClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *RoleAssignmentsClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *RoleAssignmentsClient) listByHubHandleResponse(resp *http.Response) (RoleAssignmentsClientListByHubResponse, error) {
	result := RoleAssignmentsClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListByHubResponse{}, err
	}
	return result, nil
}
