//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LabPlansClient contains the methods for the LabPlans group.
// Don't use this type directly, use NewLabPlansClient() instead.
type LabPlansClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLabPlansClient creates a new instance of LabPlansClient with the specified values.
func NewLabPlansClient(con *arm.Connection, subscriptionID string) *LabPlansClient {
	return &LabPlansClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Operation to create or update a Lab Plan resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlan, options *LabPlansBeginCreateOrUpdateOptions) (LabPlansCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return LabPlansCreateOrUpdatePollerResponse{}, err
	}
	result := LabPlansCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LabPlansClient.CreateOrUpdate", "original-uri", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return LabPlansCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &LabPlansCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Operation to create or update a Lab Plan resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) createOrUpdate(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlan, options *LabPlansBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LabPlansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlan, options *LabPlansBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans/{labPlanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labPlanName == "" {
		return nil, errors.New("parameter labPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labPlanName}", url.PathEscape(labPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LabPlansClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Operation to delete a Lab Plan resource. Deleting a lab plan does not delete labs associated with a lab plan, nor does it delete shared
// images added to a gallery via the lab plan permission container.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, labPlanName string, options *LabPlansBeginDeleteOptions) (LabPlansDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, labPlanName, options)
	if err != nil {
		return LabPlansDeletePollerResponse{}, err
	}
	result := LabPlansDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LabPlansClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return LabPlansDeletePollerResponse{}, err
	}
	result.Poller = &LabPlansDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a Lab Plan resource. Deleting a lab plan does not delete labs associated with a lab plan, nor does it delete shared images
// added to a gallery via the lab plan permission container.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) deleteOperation(ctx context.Context, resourceGroupName string, labPlanName string, options *LabPlansBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labPlanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LabPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labPlanName string, options *LabPlansBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans/{labPlanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labPlanName == "" {
		return nil, errors.New("parameter labPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labPlanName}", url.PathEscape(labPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LabPlansClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the properties of a Lab Plan.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) Get(ctx context.Context, resourceGroupName string, labPlanName string, options *LabPlansGetOptions) (LabPlansGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labPlanName, options)
	if err != nil {
		return LabPlansGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LabPlansGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LabPlansGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LabPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, labPlanName string, options *LabPlansGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans/{labPlanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labPlanName == "" {
		return nil, errors.New("parameter labPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labPlanName}", url.PathEscape(labPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LabPlansClient) getHandleResponse(resp *http.Response) (LabPlansGetResponse, error) {
	result := LabPlansGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LabPlan); err != nil {
		return LabPlansGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LabPlansClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Returns a list of all lab plans for a subscription and resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) ListByResourceGroup(resourceGroupName string, options *LabPlansListByResourceGroupOptions) *LabPlansListByResourceGroupPager {
	return &LabPlansListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LabPlansListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PagedLabPlans.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LabPlansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LabPlansListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans"
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
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LabPlansClient) listByResourceGroupHandleResponse(resp *http.Response) (LabPlansListByResourceGroupResponse, error) {
	result := LabPlansListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedLabPlans); err != nil {
		return LabPlansListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *LabPlansClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Returns a list of all lab plans within a subscription
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) ListBySubscription(options *LabPlansListBySubscriptionOptions) *LabPlansListBySubscriptionPager {
	return &LabPlansListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LabPlansListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PagedLabPlans.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LabPlansClient) listBySubscriptionCreateRequest(ctx context.Context, options *LabPlansListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LabPlansClient) listBySubscriptionHandleResponse(resp *http.Response) (LabPlansListBySubscriptionResponse, error) {
	result := LabPlansListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedLabPlans); err != nil {
		return LabPlansListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *LabPlansClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginSaveImage - Saves an image from a lab VM to the attached shared image gallery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) BeginSaveImage(ctx context.Context, resourceGroupName string, labPlanName string, body SaveImageBody, options *LabPlansBeginSaveImageOptions) (LabPlansSaveImagePollerResponse, error) {
	resp, err := client.saveImage(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return LabPlansSaveImagePollerResponse{}, err
	}
	result := LabPlansSaveImagePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LabPlansClient.SaveImage", "location", resp, client.pl, client.saveImageHandleError)
	if err != nil {
		return LabPlansSaveImagePollerResponse{}, err
	}
	result.Poller = &LabPlansSaveImagePoller{
		pt: pt,
	}
	return result, nil
}

// SaveImage - Saves an image from a lab VM to the attached shared image gallery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) saveImage(ctx context.Context, resourceGroupName string, labPlanName string, body SaveImageBody, options *LabPlansBeginSaveImageOptions) (*http.Response, error) {
	req, err := client.saveImageCreateRequest(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.saveImageHandleError(resp)
	}
	return resp, nil
}

// saveImageCreateRequest creates the SaveImage request.
func (client *LabPlansClient) saveImageCreateRequest(ctx context.Context, resourceGroupName string, labPlanName string, body SaveImageBody, options *LabPlansBeginSaveImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans/{labPlanName}/saveImage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labPlanName == "" {
		return nil, errors.New("parameter labPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labPlanName}", url.PathEscape(labPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// saveImageHandleError handles the SaveImage error response.
func (client *LabPlansClient) saveImageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Operation to update a Lab Plan resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) BeginUpdate(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlanUpdate, options *LabPlansBeginUpdateOptions) (LabPlansUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return LabPlansUpdatePollerResponse{}, err
	}
	result := LabPlansUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LabPlansClient.Update", "location", resp, client.pl, client.updateHandleError)
	if err != nil {
		return LabPlansUpdatePollerResponse{}, err
	}
	result.Poller = &LabPlansUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update a Lab Plan resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LabPlansClient) update(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlanUpdate, options *LabPlansBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labPlanName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *LabPlansClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labPlanName string, body LabPlanUpdate, options *LabPlansBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labPlans/{labPlanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labPlanName == "" {
		return nil, errors.New("parameter labPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labPlanName}", url.PathEscape(labPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleError handles the Update error response.
func (client *LabPlansClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
