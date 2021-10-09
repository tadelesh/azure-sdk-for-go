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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SyncGroupsClient contains the methods for the SyncGroups group.
// Don't use this type directly, use NewSyncGroupsClient() instead.
type SyncGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSyncGroupsClient creates a new instance of SyncGroupsClient with the specified values.
func NewSyncGroupsClient(con *arm.Connection, subscriptionID string) *SyncGroupsClient {
	return &SyncGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Create a new SyncGroup.
// If the operation fails it returns the *StorageSyncError error type.
func (client *SyncGroupsClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters, options *SyncGroupsCreateOptions) (SyncGroupsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, parameters, options)
	if err != nil {
		return SyncGroupsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SyncGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters, options *SyncGroupsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
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

// createHandleResponse handles the Create response.
func (client *SyncGroupsClient) createHandleResponse(resp *http.Response) (SyncGroupsCreateResponse, error) {
	result := SyncGroupsCreateResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroup); err != nil {
		return SyncGroupsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *SyncGroupsClient) createHandleError(resp *http.Response) error {
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

// Delete - Delete a given SyncGroup.
// If the operation fails it returns the *StorageSyncError error type.
func (client *SyncGroupsClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsDeleteOptions) (SyncGroupsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, options)
	if err != nil {
		return SyncGroupsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SyncGroupsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *SyncGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
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

// deleteHandleResponse handles the Delete response.
func (client *SyncGroupsClient) deleteHandleResponse(resp *http.Response) (SyncGroupsDeleteResponse, error) {
	result := SyncGroupsDeleteResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *SyncGroupsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a given SyncGroup.
// If the operation fails it returns the *StorageSyncError error type.
func (client *SyncGroupsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsGetOptions) (SyncGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, options)
	if err != nil {
		return SyncGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SyncGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
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
func (client *SyncGroupsClient) getHandleResponse(resp *http.Response) (SyncGroupsGetResponse, error) {
	result := SyncGroupsGetResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroup); err != nil {
		return SyncGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SyncGroupsClient) getHandleError(resp *http.Response) error {
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

// ListByStorageSyncService - Get a SyncGroup List.
// If the operation fails it returns the *StorageSyncError error type.
func (client *SyncGroupsClient) ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *SyncGroupsListByStorageSyncServiceOptions) (SyncGroupsListByStorageSyncServiceResponse, error) {
	req, err := client.listByStorageSyncServiceCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return SyncGroupsListByStorageSyncServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsListByStorageSyncServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsListByStorageSyncServiceResponse{}, client.listByStorageSyncServiceHandleError(resp)
	}
	return client.listByStorageSyncServiceHandleResponse(resp)
}

// listByStorageSyncServiceCreateRequest creates the ListByStorageSyncService request.
func (client *SyncGroupsClient) listByStorageSyncServiceCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *SyncGroupsListByStorageSyncServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups"
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

// listByStorageSyncServiceHandleResponse handles the ListByStorageSyncService response.
func (client *SyncGroupsClient) listByStorageSyncServiceHandleResponse(resp *http.Response) (SyncGroupsListByStorageSyncServiceResponse, error) {
	result := SyncGroupsListByStorageSyncServiceResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroupArray); err != nil {
		return SyncGroupsListByStorageSyncServiceResponse{}, err
	}
	return result, nil
}

// listByStorageSyncServiceHandleError handles the ListByStorageSyncService error response.
func (client *SyncGroupsClient) listByStorageSyncServiceHandleError(resp *http.Response) error {
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
