//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// GlobalSchedulesClient contains the methods for the GlobalSchedules group.
// Don't use this type directly, use NewGlobalSchedulesClient() instead.
type GlobalSchedulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewGlobalSchedulesClient creates a new instance of GlobalSchedulesClient with the specified values.
func NewGlobalSchedulesClient(con *arm.Connection, subscriptionID string) *GlobalSchedulesClient {
	return &GlobalSchedulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or replace an existing schedule.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, schedule Schedule, options *GlobalSchedulesCreateOrUpdateOptions) (GlobalSchedulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, schedule, options)
	if err != nil {
		return GlobalSchedulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalSchedulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GlobalSchedulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GlobalSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, schedule Schedule, options *GlobalSchedulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GlobalSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (GlobalSchedulesCreateOrUpdateResponse, error) {
	result := GlobalSchedulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return GlobalSchedulesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GlobalSchedulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete schedule.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) Delete(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesDeleteOptions) (GlobalSchedulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return GlobalSchedulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalSchedulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GlobalSchedulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return GlobalSchedulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GlobalSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GlobalSchedulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginExecute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) BeginExecute(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesBeginExecuteOptions) (GlobalSchedulesExecutePollerResponse, error) {
	resp, err := client.execute(ctx, resourceGroupName, name, options)
	if err != nil {
		return GlobalSchedulesExecutePollerResponse{}, err
	}
	result := GlobalSchedulesExecutePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GlobalSchedulesClient.Execute", "", resp, client.pl, client.executeHandleError)
	if err != nil {
		return GlobalSchedulesExecutePollerResponse{}, err
	}
	result.Poller = &GlobalSchedulesExecutePoller{
		pt: pt,
	}
	return result, nil
}

// Execute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) execute(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesBeginExecuteOptions) (*http.Response, error) {
	req, err := client.executeCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.executeHandleError(resp)
	}
	return resp, nil
}

// executeCreateRequest creates the Execute request.
func (client *GlobalSchedulesClient) executeCreateRequest(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesBeginExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}/execute"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeHandleError handles the Execute error response.
func (client *GlobalSchedulesClient) executeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get schedule.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) Get(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesGetOptions) (GlobalSchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return GlobalSchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalSchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalSchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GlobalSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *GlobalSchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GlobalSchedulesClient) getHandleResponse(resp *http.Response) (GlobalSchedulesGetResponse, error) {
	result := GlobalSchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return GlobalSchedulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GlobalSchedulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List schedules in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) ListByResourceGroup(resourceGroupName string, options *GlobalSchedulesListByResourceGroupOptions) *GlobalSchedulesListByResourceGroupPager {
	return &GlobalSchedulesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp GlobalSchedulesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScheduleList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GlobalSchedulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GlobalSchedulesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GlobalSchedulesClient) listByResourceGroupHandleResponse(resp *http.Response) (GlobalSchedulesListByResourceGroupResponse, error) {
	result := GlobalSchedulesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return GlobalSchedulesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *GlobalSchedulesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - List schedules in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) ListBySubscription(options *GlobalSchedulesListBySubscriptionOptions) *GlobalSchedulesListBySubscriptionPager {
	return &GlobalSchedulesListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp GlobalSchedulesListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScheduleList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *GlobalSchedulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *GlobalSchedulesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *GlobalSchedulesClient) listBySubscriptionHandleResponse(resp *http.Response) (GlobalSchedulesListBySubscriptionResponse, error) {
	result := GlobalSchedulesListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return GlobalSchedulesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *GlobalSchedulesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRetarget - Updates a schedule's target resource Id. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) BeginRetarget(ctx context.Context, resourceGroupName string, name string, retargetScheduleProperties RetargetScheduleProperties, options *GlobalSchedulesBeginRetargetOptions) (GlobalSchedulesRetargetPollerResponse, error) {
	resp, err := client.retarget(ctx, resourceGroupName, name, retargetScheduleProperties, options)
	if err != nil {
		return GlobalSchedulesRetargetPollerResponse{}, err
	}
	result := GlobalSchedulesRetargetPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GlobalSchedulesClient.Retarget", "", resp, client.pl, client.retargetHandleError)
	if err != nil {
		return GlobalSchedulesRetargetPollerResponse{}, err
	}
	result.Poller = &GlobalSchedulesRetargetPoller{
		pt: pt,
	}
	return result, nil
}

// Retarget - Updates a schedule's target resource Id. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) retarget(ctx context.Context, resourceGroupName string, name string, retargetScheduleProperties RetargetScheduleProperties, options *GlobalSchedulesBeginRetargetOptions) (*http.Response, error) {
	req, err := client.retargetCreateRequest(ctx, resourceGroupName, name, retargetScheduleProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.retargetHandleError(resp)
	}
	return resp, nil
}

// retargetCreateRequest creates the Retarget request.
func (client *GlobalSchedulesClient) retargetCreateRequest(ctx context.Context, resourceGroupName string, name string, retargetScheduleProperties RetargetScheduleProperties, options *GlobalSchedulesBeginRetargetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}/retarget"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, retargetScheduleProperties)
}

// retargetHandleError handles the Retarget error response.
func (client *GlobalSchedulesClient) retargetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Allows modifying tags of schedules. All other properties will be ignored.
// If the operation fails it returns the *CloudError error type.
func (client *GlobalSchedulesClient) Update(ctx context.Context, resourceGroupName string, name string, schedule ScheduleFragment, options *GlobalSchedulesUpdateOptions) (GlobalSchedulesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, schedule, options)
	if err != nil {
		return GlobalSchedulesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalSchedulesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalSchedulesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *GlobalSchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, schedule ScheduleFragment, options *GlobalSchedulesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// updateHandleResponse handles the Update response.
func (client *GlobalSchedulesClient) updateHandleResponse(resp *http.Response) (GlobalSchedulesUpdateResponse, error) {
	result := GlobalSchedulesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return GlobalSchedulesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *GlobalSchedulesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
