//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// PacketCapturesClient contains the methods for the PacketCaptures group.
// Don't use this type directly, use NewPacketCapturesClient() instead.
type PacketCapturesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient with the specified values.
func NewPacketCapturesClient(con *arm.Connection, subscriptionID string) *PacketCapturesClient {
	return &PacketCapturesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Create and start a packet capture on the specified VM.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (PacketCapturesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	result := PacketCapturesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	result.Poller = &PacketCapturesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create and start a packet capture on the specified VM.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PacketCapturesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *PacketCapturesClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (PacketCapturesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	result := PacketCapturesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	result.Poller = &PacketCapturesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PacketCapturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PacketCapturesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets a packet capture session by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesGetOptions) (PacketCapturesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCapturesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCapturesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCapturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PacketCapturesClient) getHandleResponse(resp *http.Response) (PacketCapturesGetResponse, error) {
	result := PacketCapturesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureResult); err != nil {
		return PacketCapturesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PacketCapturesClient) getHandleError(resp *http.Response) error {
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

// BeginGetStatus - Query the status of a running packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginGetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (PacketCapturesGetStatusPollerResponse, error) {
	resp, err := client.getStatus(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	result := PacketCapturesGetStatusPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.GetStatus", "location", resp, client.pl, client.getStatusHandleError)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	result.Poller = &PacketCapturesGetStatusPoller{
		pt: pt,
	}
	return result, nil
}

// GetStatus - Query the status of a running packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) getStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (*http.Response, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.getStatusHandleError(resp)
	}
	return resp, nil
}

// getStatusCreateRequest creates the GetStatus request.
func (client *PacketCapturesClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStatusHandleError handles the GetStatus error response.
func (client *PacketCapturesClient) getStatusHandleError(resp *http.Response) error {
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

// List - Lists all packet capture sessions within the specified resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesListOptions) (PacketCapturesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
	if err != nil {
		return PacketCapturesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCapturesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCapturesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *PacketCapturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PacketCapturesClient) listHandleResponse(resp *http.Response) (PacketCapturesListResponse, error) {
	result := PacketCapturesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureListResult); err != nil {
		return PacketCapturesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PacketCapturesClient) listHandleError(resp *http.Response) error {
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

// BeginStop - Stops a specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (PacketCapturesStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	result := PacketCapturesStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Stop", "location", resp, client.pl, client.stopHandleError)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	result.Poller = &PacketCapturesStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stops a specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *PacketCapturesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *PacketCapturesClient) stopHandleError(resp *http.Response) error {
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
