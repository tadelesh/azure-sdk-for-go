//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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
	"strconv"
	"strings"
)

// WorkflowRunsClient contains the methods for the WorkflowRuns group.
// Don't use this type directly, use NewWorkflowRunsClient() instead.
type WorkflowRunsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunsClient creates a new instance of WorkflowRunsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowRunsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkflowRunsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Cancel - Cancels a workflow run.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// options - WorkflowRunsClientCancelOptions contains the optional parameters for the WorkflowRunsClient.Cancel method.
func (client *WorkflowRunsClient) Cancel(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsClientCancelOptions) (WorkflowRunsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
	if err != nil {
		return WorkflowRunsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return WorkflowRunsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *WorkflowRunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a workflow run.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// options - WorkflowRunsClientGetOptions contains the optional parameters for the WorkflowRunsClient.Get method.
func (client *WorkflowRunsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsClientGetOptions) (WorkflowRunsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
	if err != nil {
		return WorkflowRunsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunsClient) getHandleResponse(resp *http.Response) (WorkflowRunsClientGetResponse, error) {
	result := WorkflowRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRun); err != nil {
		return WorkflowRunsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of workflow runs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// options - WorkflowRunsClientListOptions contains the optional parameters for the WorkflowRunsClient.List method.
func (client *WorkflowRunsClient) List(resourceGroupName string, workflowName string, options *WorkflowRunsClientListOptions) *WorkflowRunsClientListPager {
	return &WorkflowRunsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowRunsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkflowRunListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunsClient) listHandleResponse(resp *http.Response) (WorkflowRunsClientListResponse, error) {
	result := WorkflowRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunListResult); err != nil {
		return WorkflowRunsClientListResponse{}, err
	}
	return result, nil
}
