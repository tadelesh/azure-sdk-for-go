//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

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

// InputsClient contains the methods for the Inputs group.
// Don't use this type directly, use NewInputsClient() instead.
type InputsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInputsClient creates a new instance of InputsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInputsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *InputsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &InputsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrReplace - Creates an input or replaces an already existing input under an existing streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// inputName - The name of the input.
// input - The definition of the input that will be used to create a new input or replace the existing one under the streaming
// job.
// options - InputsClientCreateOrReplaceOptions contains the optional parameters for the InputsClient.CreateOrReplace method.
func (client *InputsClient) CreateOrReplace(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsClientCreateOrReplaceOptions) (InputsClientCreateOrReplaceResponse, error) {
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, jobName, inputName, input, options)
	if err != nil {
		return InputsClientCreateOrReplaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsClientCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return InputsClientCreateOrReplaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrReplaceHandleResponse(resp)
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *InputsClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsClientCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *InputsClient) createOrReplaceHandleResponse(resp *http.Response) (InputsClientCreateOrReplaceResponse, error) {
	result := InputsClientCreateOrReplaceResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsClientCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an input from the streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// inputName - The name of the input.
// options - InputsClientDeleteOptions contains the optional parameters for the InputsClient.Delete method.
func (client *InputsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientDeleteOptions) (InputsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return InputsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return InputsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InputsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets details about the specified input.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// inputName - The name of the input.
// options - InputsClientGetOptions contains the optional parameters for the InputsClient.Get method.
func (client *InputsClient) Get(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientGetOptions) (InputsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InputsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InputsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InputsClient) getHandleResponse(resp *http.Response) (InputsClientGetResponse, error) {
	result := InputsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsClientGetResponse{}, err
	}
	return result, nil
}

// ListByStreamingJob - Lists all of the inputs under the specified streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - InputsClientListByStreamingJobOptions contains the optional parameters for the InputsClient.ListByStreamingJob
// method.
func (client *InputsClient) ListByStreamingJob(resourceGroupName string, jobName string, options *InputsClientListByStreamingJobOptions) *InputsClientListByStreamingJobPager {
	return &InputsClientListByStreamingJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByStreamingJobCreateRequest(ctx, resourceGroupName, jobName, options)
		},
		advancer: func(ctx context.Context, resp InputsClientListByStreamingJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InputListResult.NextLink)
		},
	}
}

// listByStreamingJobCreateRequest creates the ListByStreamingJob request.
func (client *InputsClient) listByStreamingJobCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *InputsClientListByStreamingJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByStreamingJobHandleResponse handles the ListByStreamingJob response.
func (client *InputsClient) listByStreamingJobHandleResponse(resp *http.Response) (InputsClientListByStreamingJobResponse, error) {
	result := InputsClientListByStreamingJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InputListResult); err != nil {
		return InputsClientListByStreamingJobResponse{}, err
	}
	return result, nil
}

// BeginTest - Tests whether an input’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// inputName - The name of the input.
// options - InputsClientBeginTestOptions contains the optional parameters for the InputsClient.BeginTest method.
func (client *InputsClient) BeginTest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientBeginTestOptions) (InputsClientTestPollerResponse, error) {
	resp, err := client.test(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsClientTestPollerResponse{}, err
	}
	result := InputsClientTestPollerResponse{}
	pt, err := armruntime.NewPoller("InputsClient.Test", "", resp, client.pl)
	if err != nil {
		return InputsClientTestPollerResponse{}, err
	}
	result.Poller = &InputsClientTestPoller{
		pt: pt,
	}
	return result, nil
}

// Test - Tests whether an input’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InputsClient) test(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientBeginTestOptions) (*http.Response, error) {
	req, err := client.testCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
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

// testCreateRequest creates the Test request.
func (client *InputsClient) testCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsClientBeginTestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}/test"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, runtime.MarshalAsJSON(req, *options.Input)
	}
	return req, nil
}

// Update - Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one
// or two properties) an input without affecting the rest the job or input definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// inputName - The name of the input.
// input - An Input object. The properties specified here will overwrite the corresponding properties in the existing input
// (ie. Those properties will be updated). Any properties that are set to null here will
// mean that the corresponding property in the existing input will remain the same and not change as a result of this PATCH
// operation.
// options - InputsClientUpdateOptions contains the optional parameters for the InputsClient.Update method.
func (client *InputsClient) Update(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsClientUpdateOptions) (InputsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, inputName, input, options)
	if err != nil {
		return InputsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InputsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *InputsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// updateHandleResponse handles the Update response.
func (client *InputsClient) updateHandleResponse(resp *http.Response) (InputsClientUpdateResponse, error) {
	result := InputsClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsClientUpdateResponse{}, err
	}
	return result, nil
}
