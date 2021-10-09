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

// BackupPoliciesClient contains the methods for the BackupPolicies group.
// Don't use this type directly, use NewBackupPoliciesClient() instead.
type BackupPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient with the specified values.
func NewBackupPoliciesClient(con *arm.Connection, subscriptionID string) *BackupPoliciesClient {
	return &BackupPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginBackupNow - Backup the backup policy now.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) BeginBackupNow(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesBeginBackupNowOptions) (BackupPoliciesBackupNowPollerResponse, error) {
	resp, err := client.backupNow(ctx, deviceName, backupPolicyName, backupType, resourceGroupName, managerName, options)
	if err != nil {
		return BackupPoliciesBackupNowPollerResponse{}, err
	}
	result := BackupPoliciesBackupNowPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupPoliciesClient.BackupNow", "", resp, client.pl, client.backupNowHandleError)
	if err != nil {
		return BackupPoliciesBackupNowPollerResponse{}, err
	}
	result.Poller = &BackupPoliciesBackupNowPoller{
		pt: pt,
	}
	return result, nil
}

// BackupNow - Backup the backup policy now.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) backupNow(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesBeginBackupNowOptions) (*http.Response, error) {
	req, err := client.backupNowCreateRequest(ctx, deviceName, backupPolicyName, backupType, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.backupNowHandleError(resp)
	}
	return resp, nil
}

// backupNowCreateRequest creates the BackupNow request.
func (client *BackupPoliciesClient) backupNowCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesBeginBackupNowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/backup"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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
	unencodedParams := []string{req.Raw().URL.RawQuery}
	unencodedParams = append(unencodedParams, "backupType="+backupType)
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	return req, nil
}

// backupNowHandleError handles the BackupNow error response.
func (client *BackupPoliciesClient) backupNowHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginCreateOrUpdate - Creates or updates the backup policy.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesBeginCreateOrUpdateOptions) (BackupPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return BackupPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := BackupPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupPoliciesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return BackupPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BackupPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the backup policy.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) createOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *BackupPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes the backup policy.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) BeginDelete(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesBeginDeleteOptions) (BackupPoliciesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupPoliciesDeletePollerResponse{}, err
	}
	result := BackupPoliciesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupPoliciesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return BackupPoliciesDeletePollerResponse{}, err
	}
	result.Poller = &BackupPoliciesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the backup policy.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) deleteOperation(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
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
func (client *BackupPoliciesClient) deleteCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BackupPoliciesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the properties of the specified backup policy name.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) Get(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesGetOptions) (BackupPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupPoliciesClient) getCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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

// getHandleResponse handles the Get response.
func (client *BackupPoliciesClient) getHandleResponse(resp *http.Response) (BackupPoliciesGetResponse, error) {
	result := BackupPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupPolicy); err != nil {
		return BackupPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BackupPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDevice - Gets all the backup policies in a device.
// If the operation fails it returns a generic error.
func (client *BackupPoliciesClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *BackupPoliciesListByDeviceOptions) (BackupPoliciesListByDeviceResponse, error) {
	req, err := client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupPoliciesListByDeviceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesListByDeviceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupPoliciesListByDeviceResponse{}, client.listByDeviceHandleError(resp)
	}
	return client.listByDeviceHandleResponse(resp)
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *BackupPoliciesClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *BackupPoliciesListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies"
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
func (client *BackupPoliciesClient) listByDeviceHandleResponse(resp *http.Response) (BackupPoliciesListByDeviceResponse, error) {
	result := BackupPoliciesListByDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupPolicyList); err != nil {
		return BackupPoliciesListByDeviceResponse{}, err
	}
	return result, nil
}

// listByDeviceHandleError handles the ListByDevice error response.
func (client *BackupPoliciesClient) listByDeviceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
