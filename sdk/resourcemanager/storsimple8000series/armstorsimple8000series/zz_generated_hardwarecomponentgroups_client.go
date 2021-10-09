//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// HardwareComponentGroupsClient contains the methods for the HardwareComponentGroups group.
// Don't use this type directly, use NewHardwareComponentGroupsClient() instead.
type HardwareComponentGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewHardwareComponentGroupsClient creates a new instance of HardwareComponentGroupsClient with the specified values.
func NewHardwareComponentGroupsClient(con *arm.Connection, subscriptionID string) *HardwareComponentGroupsClient {
	return &HardwareComponentGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginChangeControllerPowerState - Changes the power state of the controller.
// If the operation fails it returns a generic error.
func (client *HardwareComponentGroupsClient) BeginChangeControllerPowerState(ctx context.Context, deviceName string, hardwareComponentGroupName string, resourceGroupName string, managerName string, parameters ControllerPowerStateChangeRequest, options *HardwareComponentGroupsBeginChangeControllerPowerStateOptions) (HardwareComponentGroupsChangeControllerPowerStatePollerResponse, error) {
	resp, err := client.changeControllerPowerState(ctx, deviceName, hardwareComponentGroupName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return HardwareComponentGroupsChangeControllerPowerStatePollerResponse{}, err
	}
	result := HardwareComponentGroupsChangeControllerPowerStatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("HardwareComponentGroupsClient.ChangeControllerPowerState", "", resp, client.pl, client.changeControllerPowerStateHandleError)
	if err != nil {
		return HardwareComponentGroupsChangeControllerPowerStatePollerResponse{}, err
	}
	result.Poller = &HardwareComponentGroupsChangeControllerPowerStatePoller{
		pt: pt,
	}
	return result, nil
}

// ChangeControllerPowerState - Changes the power state of the controller.
// If the operation fails it returns a generic error.
func (client *HardwareComponentGroupsClient) changeControllerPowerState(ctx context.Context, deviceName string, hardwareComponentGroupName string, resourceGroupName string, managerName string, parameters ControllerPowerStateChangeRequest, options *HardwareComponentGroupsBeginChangeControllerPowerStateOptions) (*http.Response, error) {
	req, err := client.changeControllerPowerStateCreateRequest(ctx, deviceName, hardwareComponentGroupName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.changeControllerPowerStateHandleError(resp)
	}
	return resp, nil
}

// changeControllerPowerStateCreateRequest creates the ChangeControllerPowerState request.
func (client *HardwareComponentGroupsClient) changeControllerPowerStateCreateRequest(ctx context.Context, deviceName string, hardwareComponentGroupName string, resourceGroupName string, managerName string, parameters ControllerPowerStateChangeRequest, options *HardwareComponentGroupsBeginChangeControllerPowerStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups/{hardwareComponentGroupName}/changeControllerPowerState"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if hardwareComponentGroupName == "" {
		return nil, errors.New("parameter hardwareComponentGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hardwareComponentGroupName}", hardwareComponentGroupName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// changeControllerPowerStateHandleError handles the ChangeControllerPowerState error response.
func (client *HardwareComponentGroupsClient) changeControllerPowerStateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDevice - Lists the hardware component groups at device-level.
// If the operation fails it returns a generic error.
func (client *HardwareComponentGroupsClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *HardwareComponentGroupsListByDeviceOptions) (HardwareComponentGroupsListByDeviceResponse, error) {
	req, err := client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return HardwareComponentGroupsListByDeviceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HardwareComponentGroupsListByDeviceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HardwareComponentGroupsListByDeviceResponse{}, client.listByDeviceHandleError(resp)
	}
	return client.listByDeviceHandleResponse(resp)
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *HardwareComponentGroupsClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *HardwareComponentGroupsListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *HardwareComponentGroupsClient) listByDeviceHandleResponse(resp *http.Response) (HardwareComponentGroupsListByDeviceResponse, error) {
	result := HardwareComponentGroupsListByDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HardwareComponentGroupList); err != nil {
		return HardwareComponentGroupsListByDeviceResponse{}, err
	}
	return result, nil
}

// listByDeviceHandleError handles the ListByDevice error response.
func (client *HardwareComponentGroupsClient) listByDeviceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
