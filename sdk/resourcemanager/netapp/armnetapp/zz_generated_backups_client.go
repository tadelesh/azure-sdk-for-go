//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
func NewBackupsClient(con *arm.Connection, subscriptionID string) *BackupsClient {
	return &BackupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Create a backup for the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsBeginCreateOptions) (BackupsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
	if err != nil {
		return BackupsCreatePollerResponse{}, err
	}
	result := BackupsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupsClient.Create", "location", resp, client.pl, client.createHandleError)
	if err != nil {
		return BackupsCreatePollerResponse{}, err
	}
	result.Poller = &BackupsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a backup for the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *BackupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleError handles the Create error response.
func (client *BackupsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete a backup of the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginDeleteOptions) (BackupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
	if err != nil {
		return BackupsDeletePollerResponse{}, err
	}
	result := BackupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return BackupsDeletePollerResponse{}, err
	}
	result.Poller = &BackupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a backup of the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
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
func (client *BackupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BackupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the specified backup of the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsGetOptions) (BackupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
	if err != nil {
		return BackupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupsClient) getHandleResponse(resp *http.Response) (BackupsGetResponse, error) {
	result := BackupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Backup); err != nil {
		return BackupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BackupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetStatus - Get the status of the backup for a volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) GetStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsGetStatusOptions) (BackupsGetStatusResponse, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsGetStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsGetStatusResponse{}, client.getStatusHandleError(resp)
	}
	return client.getStatusHandleResponse(resp)
}

// getStatusCreateRequest creates the GetStatus request.
func (client *BackupsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backupStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *BackupsClient) getStatusHandleResponse(resp *http.Response) (BackupsGetStatusResponse, error) {
	result := BackupsGetStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupStatus); err != nil {
		return BackupsGetStatusResponse{}, err
	}
	return result, nil
}

// getStatusHandleError handles the GetStatus error response.
func (client *BackupsClient) getStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetVolumeRestoreStatus - Get the status of the restore for a volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) GetVolumeRestoreStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsGetVolumeRestoreStatusOptions) (BackupsGetVolumeRestoreStatusResponse, error) {
	req, err := client.getVolumeRestoreStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsGetVolumeRestoreStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsGetVolumeRestoreStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsGetVolumeRestoreStatusResponse{}, client.getVolumeRestoreStatusHandleError(resp)
	}
	return client.getVolumeRestoreStatusHandleResponse(resp)
}

// getVolumeRestoreStatusCreateRequest creates the GetVolumeRestoreStatus request.
func (client *BackupsClient) getVolumeRestoreStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsGetVolumeRestoreStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/restoreStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getVolumeRestoreStatusHandleResponse handles the GetVolumeRestoreStatus response.
func (client *BackupsClient) getVolumeRestoreStatusHandleResponse(resp *http.Response) (BackupsGetVolumeRestoreStatusResponse, error) {
	result := BackupsGetVolumeRestoreStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestoreStatus); err != nil {
		return BackupsGetVolumeRestoreStatusResponse{}, err
	}
	return result, nil
}

// getVolumeRestoreStatusHandleError handles the GetVolumeRestoreStatus error response.
func (client *BackupsClient) getVolumeRestoreStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - List all backups for a volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsListOptions) (BackupsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BackupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupsClient) listHandleResponse(resp *http.Response) (BackupsListResponse, error) {
	result := BackupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupsList); err != nil {
		return BackupsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *BackupsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Patch a backup for the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginUpdateOptions) (BackupsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
	if err != nil {
		return BackupsUpdatePollerResponse{}, err
	}
	result := BackupsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BackupsClient.Update", "location", resp, client.pl, client.updateHandleError)
	if err != nil {
		return BackupsUpdatePollerResponse{}, err
	}
	result.Poller = &BackupsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch a backup for the volume
// If the operation fails it returns a generic error.
func (client *BackupsClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
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
func (client *BackupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleError handles the Update error response.
func (client *BackupsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
