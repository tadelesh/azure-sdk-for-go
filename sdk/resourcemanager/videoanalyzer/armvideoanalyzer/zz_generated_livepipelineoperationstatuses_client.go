//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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

// LivePipelineOperationStatusesClient contains the methods for the LivePipelineOperationStatuses group.
// Don't use this type directly, use NewLivePipelineOperationStatusesClient() instead.
type LivePipelineOperationStatusesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLivePipelineOperationStatusesClient creates a new instance of LivePipelineOperationStatusesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLivePipelineOperationStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LivePipelineOperationStatusesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LivePipelineOperationStatusesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get the operation status of a live pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// operationID - The operation ID.
// options - LivePipelineOperationStatusesClientGetOptions contains the optional parameters for the LivePipelineOperationStatusesClient.Get
// method.
func (client *LivePipelineOperationStatusesClient) Get(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, operationID string, options *LivePipelineOperationStatusesClientGetOptions) (LivePipelineOperationStatusesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, operationID, options)
	if err != nil {
		return LivePipelineOperationStatusesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LivePipelineOperationStatusesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LivePipelineOperationStatusesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LivePipelineOperationStatusesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, operationID string, options *LivePipelineOperationStatusesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}/operationStatuses/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LivePipelineOperationStatusesClient) getHandleResponse(resp *http.Response) (LivePipelineOperationStatusesClientGetResponse, error) {
	result := LivePipelineOperationStatusesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivePipelineOperationStatus); err != nil {
		return LivePipelineOperationStatusesClientGetResponse{}, err
	}
	return result, nil
}
