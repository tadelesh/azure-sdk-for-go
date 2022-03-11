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
	"strings"
)

// WorkflowRunOperationsClient contains the methods for the WorkflowRunOperations group.
// Don't use this type directly, use NewWorkflowRunOperationsClient() instead.
type WorkflowRunOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunOperationsClient creates a new instance of WorkflowRunOperationsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowRunOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowRunOperationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkflowRunOperationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets an operation for a run.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// operationID - The workflow operation id.
// options - WorkflowRunOperationsClientGetOptions contains the optional parameters for the WorkflowRunOperationsClient.Get
// method.
func (client *WorkflowRunOperationsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, operationID string, options *WorkflowRunOperationsClientGetOptions) (WorkflowRunOperationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, operationID, options)
	if err != nil {
		return WorkflowRunOperationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunOperationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, operationID string, options *WorkflowRunOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/operations/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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
func (client *WorkflowRunOperationsClient) getHandleResponse(resp *http.Response) (WorkflowRunOperationsClientGetResponse, error) {
	result := WorkflowRunOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRun); err != nil {
		return WorkflowRunOperationsClientGetResponse{}, err
	}
	return result, nil
}
