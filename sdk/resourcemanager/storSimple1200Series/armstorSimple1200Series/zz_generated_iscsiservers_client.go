//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorSimple1200Series

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

// IscsiServersClient contains the methods for the IscsiServers group.
// Don't use this type directly, use NewIscsiServersClient() instead.
type IscsiServersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIscsiServersClient creates a new instance of IscsiServersClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIscsiServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IscsiServersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IscsiServersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginBackupNow - Backup the iSCSI server now.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientBeginBackupNowOptions contains the optional parameters for the IscsiServersClient.BeginBackupNow
// method.
func (client *IscsiServersClient) BeginBackupNow(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginBackupNowOptions) (IscsiServersClientBackupNowPollerResponse, error) {
	resp, err := client.backupNow(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
	if err != nil {
		return IscsiServersClientBackupNowPollerResponse{}, err
	}
	result := IscsiServersClientBackupNowPollerResponse{}
	pt, err := armruntime.NewPoller("IscsiServersClient.BackupNow", "", resp, client.pl)
	if err != nil {
		return IscsiServersClientBackupNowPollerResponse{}, err
	}
	result.Poller = &IscsiServersClientBackupNowPoller{
		pt: pt,
	}
	return result, nil
}

// BackupNow - Backup the iSCSI server now.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiServersClient) backupNow(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginBackupNowOptions) (*http.Response, error) {
	req, err := client.backupNowCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// backupNowCreateRequest creates the BackupNow request.
func (client *IscsiServersClient) backupNowCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginBackupNowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/backup"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginCreateOrUpdate - Creates or updates the iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// iscsiServer - The iSCSI server.
// options - IscsiServersClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiServersClient.BeginCreateOrUpdate
// method.
func (client *IscsiServersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, iscsiServer ISCSIServer, options *IscsiServersClientBeginCreateOrUpdateOptions) (IscsiServersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, iscsiServer, options)
	if err != nil {
		return IscsiServersClientCreateOrUpdatePollerResponse{}, err
	}
	result := IscsiServersClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("IscsiServersClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return IscsiServersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IscsiServersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiServersClient) createOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, iscsiServer ISCSIServer, options *IscsiServersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, iscsiServer, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IscsiServersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, iscsiServer ISCSIServer, options *IscsiServersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iscsiServer)
}

// BeginDelete - Deletes the iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientBeginDeleteOptions contains the optional parameters for the IscsiServersClient.BeginDelete
// method.
func (client *IscsiServersClient) BeginDelete(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginDeleteOptions) (IscsiServersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
	if err != nil {
		return IscsiServersClientDeletePollerResponse{}, err
	}
	result := IscsiServersClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("IscsiServersClient.Delete", "", resp, client.pl)
	if err != nil {
		return IscsiServersClientDeletePollerResponse{}, err
	}
	result.Poller = &IscsiServersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiServersClient) deleteOperation(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IscsiServersClient) deleteCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of the specified iSCSI server name.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientGetOptions contains the optional parameters for the IscsiServersClient.Get method.
func (client *IscsiServersClient) Get(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientGetOptions) (IscsiServersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
	if err != nil {
		return IscsiServersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IscsiServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IscsiServersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IscsiServersClient) getCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IscsiServersClient) getHandleResponse(resp *http.Response) (IscsiServersClientGetResponse, error) {
	result := IscsiServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIServer); err != nil {
		return IscsiServersClientGetResponse{}, err
	}
	return result, nil
}

// ListByDevice - Retrieves all the iSCSI in a device.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientListByDeviceOptions contains the optional parameters for the IscsiServersClient.ListByDevice
// method.
func (client *IscsiServersClient) ListByDevice(deviceName string, resourceGroupName string, managerName string, options *IscsiServersClientListByDeviceOptions) *IscsiServersClientListByDevicePager {
	return &IscsiServersClientListByDevicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
		},
	}
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *IscsiServersClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *IscsiServersClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *IscsiServersClient) listByDeviceHandleResponse(resp *http.Response) (IscsiServersClientListByDeviceResponse, error) {
	result := IscsiServersClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIServerList); err != nil {
		return IscsiServersClientListByDeviceResponse{}, err
	}
	return result, nil
}

// ListByManager - Retrieves all the iSCSI servers in a manager.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientListByManagerOptions contains the optional parameters for the IscsiServersClient.ListByManager
// method.
func (client *IscsiServersClient) ListByManager(resourceGroupName string, managerName string, options *IscsiServersClientListByManagerOptions) *IscsiServersClientListByManagerPager {
	return &IscsiServersClientListByManagerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
		},
	}
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *IscsiServersClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *IscsiServersClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/iscsiservers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *IscsiServersClient) listByManagerHandleResponse(resp *http.Response) (IscsiServersClientListByManagerResponse, error) {
	result := IscsiServersClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIServerList); err != nil {
		return IscsiServersClientListByManagerResponse{}, err
	}
	return result, nil
}

// ListMetricDefinition - Retrieves metric definitions for all metrics aggregated at iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientListMetricDefinitionOptions contains the optional parameters for the IscsiServersClient.ListMetricDefinition
// method.
func (client *IscsiServersClient) ListMetricDefinition(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientListMetricDefinitionOptions) *IscsiServersClientListMetricDefinitionPager {
	return &IscsiServersClientListMetricDefinitionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listMetricDefinitionCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
		},
	}
}

// listMetricDefinitionCreateRequest creates the ListMetricDefinition request.
func (client *IscsiServersClient) listMetricDefinitionCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientListMetricDefinitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metricsDefinitions"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricDefinitionHandleResponse handles the ListMetricDefinition response.
func (client *IscsiServersClient) listMetricDefinitionHandleResponse(resp *http.Response) (IscsiServersClientListMetricDefinitionResponse, error) {
	result := IscsiServersClientListMetricDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionList); err != nil {
		return IscsiServersClientListMetricDefinitionResponse{}, err
	}
	return result, nil
}

// ListMetrics - Gets the iSCSI server metrics
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiServersClientListMetricsOptions contains the optional parameters for the IscsiServersClient.ListMetrics
// method.
func (client *IscsiServersClient) ListMetrics(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientListMetricsOptions) *IscsiServersClientListMetricsPager {
	return &IscsiServersClientListMetricsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listMetricsCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
		},
	}
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *IscsiServersClient) listMetricsCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiServersClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metrics"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *IscsiServersClient) listMetricsHandleResponse(resp *http.Response) (IscsiServersClientListMetricsResponse, error) {
	result := IscsiServersClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricList); err != nil {
		return IscsiServersClientListMetricsResponse{}, err
	}
	return result, nil
}
