//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
func NewJobsClient(con *arm.Connection, subscriptionID string) *JobsClient {
	return &JobsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Provisions a new job or updates an existing job.
// If the operation fails it returns a generic error.
func (client *JobsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsCreateOrUpdateOptions) (JobsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, job, options)
	if err != nil {
		return JobsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, job)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobsClient) createOrUpdateHandleResponse(resp *http.Response) (JobsCreateOrUpdateResponse, error) {
	result := JobsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *JobsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes a job.
// If the operation fails it returns a generic error.
func (client *JobsClient) Delete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsDeleteOptions) (JobsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return JobsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *JobsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a job.
// If the operation fails it returns a generic error.
func (client *JobsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsGetOptions) (JobsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsGetResponse, error) {
	result := JobsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists all jobs under the specified job collection.
// If the operation fails it returns a generic error.
func (client *JobsClient) List(resourceGroupName string, jobCollectionName string, options *JobsListOptions) *JobsListPager {
	return &JobsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
		},
		advancer: func(ctx context.Context, resp JobsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *JobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JobsClient) listHandleResponse(resp *http.Response) (JobsListResponse, error) {
	result := JobsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobListResult); err != nil {
		return JobsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *JobsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListJobHistory - Lists job history.
// If the operation fails it returns a generic error.
func (client *JobsClient) ListJobHistory(resourceGroupName string, jobCollectionName string, jobName string, options *JobsListJobHistoryOptions) *JobsListJobHistoryPager {
	return &JobsListJobHistoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listJobHistoryCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
		},
		advancer: func(ctx context.Context, resp JobsListJobHistoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobHistoryListResult.NextLink)
		},
	}
}

// listJobHistoryCreateRequest creates the ListJobHistory request.
func (client *JobsClient) listJobHistoryCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsListJobHistoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listJobHistoryHandleResponse handles the ListJobHistory response.
func (client *JobsClient) listJobHistoryHandleResponse(resp *http.Response) (JobsListJobHistoryResponse, error) {
	result := JobsListJobHistoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobHistoryListResult); err != nil {
		return JobsListJobHistoryResponse{}, err
	}
	return result, nil
}

// listJobHistoryHandleError handles the ListJobHistory error response.
func (client *JobsClient) listJobHistoryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Patch - Patches an existing job.
// If the operation fails it returns a generic error.
func (client *JobsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsPatchOptions) (JobsPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, job, options)
	if err != nil {
		return JobsPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsPatchResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *JobsClient) patchCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, job)
}

// patchHandleResponse handles the Patch response.
func (client *JobsClient) patchHandleResponse(resp *http.Response) (JobsPatchResponse, error) {
	result := JobsPatchResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsPatchResponse{}, err
	}
	return result, nil
}

// patchHandleError handles the Patch error response.
func (client *JobsClient) patchHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Run - Runs a job.
// If the operation fails it returns a generic error.
func (client *JobsClient) Run(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsRunOptions) (JobsRunResponse, error) {
	req, err := client.runCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsRunResponse{}, client.runHandleError(resp)
	}
	return JobsRunResponse{RawResponse: resp}, nil
}

// runCreateRequest creates the Run request.
func (client *JobsClient) runCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/run"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// runHandleError handles the Run error response.
func (client *JobsClient) runHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
