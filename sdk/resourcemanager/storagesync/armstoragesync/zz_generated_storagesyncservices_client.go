//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

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

// StorageSyncServicesClient contains the methods for the StorageSyncServices group.
// Don't use this type directly, use NewStorageSyncServicesClient() instead.
type StorageSyncServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewStorageSyncServicesClient creates a new instance of StorageSyncServicesClient with the specified values.
func NewStorageSyncServicesClient(con *arm.Connection, subscriptionID string) *StorageSyncServicesClient {
	return &StorageSyncServicesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckNameAvailability - Check the give namespace name availability.
// If the operation fails it returns a generic error.
func (client *StorageSyncServicesClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *StorageSyncServicesCheckNameAvailabilityOptions) (StorageSyncServicesCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return StorageSyncServicesCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSyncServicesCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageSyncServicesCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *StorageSyncServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *StorageSyncServicesCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/locations/{locationName}/checkNameAvailability"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *StorageSyncServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (StorageSyncServicesCheckNameAvailabilityResponse, error) {
	result := StorageSyncServicesCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return StorageSyncServicesCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *StorageSyncServicesClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginCreate - Create a new StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) BeginCreate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters StorageSyncServiceCreateParameters, options *StorageSyncServicesBeginCreateOptions) (StorageSyncServicesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, storageSyncServiceName, parameters, options)
	if err != nil {
		return StorageSyncServicesCreatePollerResponse{}, err
	}
	result := StorageSyncServicesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageSyncServicesClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return StorageSyncServicesCreatePollerResponse{}, err
	}
	result.Poller = &StorageSyncServicesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a new StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters StorageSyncServiceCreateParameters, options *StorageSyncServicesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, storageSyncServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *StorageSyncServicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters StorageSyncServiceCreateParameters, options *StorageSyncServicesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *StorageSyncServicesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a given StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginDeleteOptions) (StorageSyncServicesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return StorageSyncServicesDeletePollerResponse{}, err
	}
	result := StorageSyncServicesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageSyncServicesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return StorageSyncServicesDeletePollerResponse{}, err
	}
	result.Poller = &StorageSyncServicesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a given StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
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
func (client *StorageSyncServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *StorageSyncServicesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a given StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesGetOptions) (StorageSyncServicesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return StorageSyncServicesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSyncServicesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageSyncServicesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageSyncServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageSyncServicesClient) getHandleResponse(resp *http.Response) (StorageSyncServicesGetResponse, error) {
	result := StorageSyncServicesGetResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSyncService); err != nil {
		return StorageSyncServicesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *StorageSyncServicesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Get a StorageSyncService list by Resource group name.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *StorageSyncServicesListByResourceGroupOptions) (StorageSyncServicesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return StorageSyncServicesListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSyncServicesListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageSyncServicesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *StorageSyncServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *StorageSyncServicesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices"
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
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *StorageSyncServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (StorageSyncServicesListByResourceGroupResponse, error) {
	result := StorageSyncServicesListByResourceGroupResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSyncServiceArray); err != nil {
		return StorageSyncServicesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *StorageSyncServicesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Get a StorageSyncService list by subscription.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) ListBySubscription(ctx context.Context, options *StorageSyncServicesListBySubscriptionOptions) (StorageSyncServicesListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return StorageSyncServicesListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSyncServicesListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageSyncServicesListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *StorageSyncServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *StorageSyncServicesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/storageSyncServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *StorageSyncServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (StorageSyncServicesListBySubscriptionResponse, error) {
	result := StorageSyncServicesListBySubscriptionResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSyncServiceArray); err != nil {
		return StorageSyncServicesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *StorageSyncServicesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Patch a given StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginUpdateOptions) (StorageSyncServicesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return StorageSyncServicesUpdatePollerResponse{}, err
	}
	result := StorageSyncServicesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageSyncServicesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return StorageSyncServicesUpdatePollerResponse{}, err
	}
	result.Poller = &StorageSyncServicesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch a given StorageSyncService.
// If the operation fails it returns the *StorageSyncError error type.
func (client *StorageSyncServicesClient) update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
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
func (client *StorageSyncServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *StorageSyncServicesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleError handles the Update error response.
func (client *StorageSyncServicesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := StorageSyncError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
