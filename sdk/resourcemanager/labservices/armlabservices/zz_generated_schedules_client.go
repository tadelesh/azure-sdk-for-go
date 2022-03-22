//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// SchedulesClient contains the methods for the Schedules group.
// Don't use this type directly, use NewSchedulesClient() instead.
type SchedulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSchedulesClient creates a new instance of SchedulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SchedulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SchedulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Operation to create or update a lab schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// scheduleName - The name of the schedule that uniquely identifies it within containing lab. Used in resource URIs.
// body - The request body.
// options - SchedulesClientCreateOrUpdateOptions contains the optional parameters for the SchedulesClient.CreateOrUpdate
// method.
func (client *SchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body Schedule, options *SchedulesClientCreateOrUpdateOptions) (SchedulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, scheduleName, body, options)
	if err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SchedulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body Schedule, options *SchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (SchedulesClientCreateOrUpdateResponse, error) {
	result := SchedulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Operation to delete a schedule resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// scheduleName - The name of the schedule that uniquely identifies it within containing lab. Used in resource URIs.
// options - SchedulesClientBeginDeleteOptions contains the optional parameters for the SchedulesClient.BeginDelete method.
func (client *SchedulesClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*armruntime.Poller[SchedulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, scheduleName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SchedulesClientDeleteResponse]("SchedulesClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SchedulesClientDeleteResponse]("SchedulesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Operation to delete a schedule resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SchedulesClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, scheduleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of a lab Schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// scheduleName - The name of the schedule that uniquely identifies it within containing lab. Used in resource URIs.
// options - SchedulesClientGetOptions contains the optional parameters for the SchedulesClient.Get method.
func (client *SchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesClientGetOptions) (SchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, scheduleName, options)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SchedulesClient) getHandleResponse(resp *http.Response) (SchedulesClientGetResponse, error) {
	result := SchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByLab - Returns a list of all schedules for a lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - SchedulesClientListByLabOptions contains the optional parameters for the SchedulesClient.ListByLab method.
func (client *SchedulesClient) ListByLab(resourceGroupName string, labName string, options *SchedulesClientListByLabOptions) *runtime.Pager[SchedulesClientListByLabResponse] {
	return runtime.NewPager(runtime.PageProcessor[SchedulesClientListByLabResponse]{
		More: func(page SchedulesClientListByLabResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SchedulesClientListByLabResponse) (SchedulesClientListByLabResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLabCreateRequest(ctx, resourceGroupName, labName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SchedulesClientListByLabResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SchedulesClientListByLabResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SchedulesClientListByLabResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLabHandleResponse(resp)
		},
	})
}

// listByLabCreateRequest creates the ListByLab request.
func (client *SchedulesClient) listByLabCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *SchedulesClientListByLabOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLabHandleResponse handles the ListByLab response.
func (client *SchedulesClient) listByLabHandleResponse(resp *http.Response) (SchedulesClientListByLabResponse, error) {
	result := SchedulesClientListByLabResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedSchedules); err != nil {
		return SchedulesClientListByLabResponse{}, err
	}
	return result, nil
}

// Update - Operation to update a lab schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// scheduleName - The name of the schedule that uniquely identifies it within containing lab. Used in resource URIs.
// body - The request body.
// options - SchedulesClientUpdateOptions contains the optional parameters for the SchedulesClient.Update method.
func (client *SchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body ScheduleUpdate, options *SchedulesClientUpdateOptions) (SchedulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, scheduleName, body, options)
	if err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body ScheduleUpdate, options *SchedulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *SchedulesClient) updateHandleResponse(resp *http.Response) (SchedulesClientUpdateResponse, error) {
	result := SchedulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	return result, nil
}
