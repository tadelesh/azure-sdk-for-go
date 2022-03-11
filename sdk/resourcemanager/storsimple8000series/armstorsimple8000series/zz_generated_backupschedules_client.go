//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// BackupSchedulesClient contains the methods for the BackupSchedules group.
// Don't use this type directly, use NewBackupSchedulesClient() instead.
type BackupSchedulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupSchedulesClient creates a new instance of BackupSchedulesClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupSchedulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BackupSchedulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the backup schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The backup policy name.
// backupScheduleName - The backup schedule name.
// resourceGroupName - The resource group name
// managerName - The manager name
// parameters - The backup schedule.
// options - BackupSchedulesClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupSchedulesClient.BeginCreateOrUpdate
// method.
func (client *BackupSchedulesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, parameters BackupSchedule, options *BackupSchedulesClientBeginCreateOrUpdateOptions) (BackupSchedulesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return BackupSchedulesClientCreateOrUpdatePollerResponse{}, err
	}
	result := BackupSchedulesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("BackupSchedulesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return BackupSchedulesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BackupSchedulesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the backup schedule.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BackupSchedulesClient) createOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, parameters BackupSchedule, options *BackupSchedulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, parameters, options)
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
func (client *BackupSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, parameters BackupSchedule, options *BackupSchedulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
	urlPath = strings.ReplaceAll(urlPath, "{backupScheduleName}", backupScheduleName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the backup schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The backup policy name.
// backupScheduleName - The name the backup schedule.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupSchedulesClientBeginDeleteOptions contains the optional parameters for the BackupSchedulesClient.BeginDelete
// method.
func (client *BackupSchedulesClient) BeginDelete(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *BackupSchedulesClientBeginDeleteOptions) (BackupSchedulesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupSchedulesClientDeletePollerResponse{}, err
	}
	result := BackupSchedulesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("BackupSchedulesClient.Delete", "", resp, client.pl)
	if err != nil {
		return BackupSchedulesClientDeletePollerResponse{}, err
	}
	result.Poller = &BackupSchedulesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the backup schedule.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BackupSchedulesClient) deleteOperation(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *BackupSchedulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, options)
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
func (client *BackupSchedulesClient) deleteCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *BackupSchedulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
	urlPath = strings.ReplaceAll(urlPath, "{backupScheduleName}", backupScheduleName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the properties of the specified backup schedule name.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The backup policy name.
// backupScheduleName - The name of the backup schedule to be fetched
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupSchedulesClientGetOptions contains the optional parameters for the BackupSchedulesClient.Get method.
func (client *BackupSchedulesClient) Get(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *BackupSchedulesClientGetOptions) (BackupSchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupSchedulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupSchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupSchedulesClient) getCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *BackupSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
	urlPath = strings.ReplaceAll(urlPath, "{backupScheduleName}", backupScheduleName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *BackupSchedulesClient) getHandleResponse(resp *http.Response) (BackupSchedulesClientGetResponse, error) {
	result := BackupSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupSchedule); err != nil {
		return BackupSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByBackupPolicy - Gets all the backup schedules in a backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The backup policy name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupSchedulesClientListByBackupPolicyOptions contains the optional parameters for the BackupSchedulesClient.ListByBackupPolicy
// method.
func (client *BackupSchedulesClient) ListByBackupPolicy(deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupSchedulesClientListByBackupPolicyOptions) *BackupSchedulesClientListByBackupPolicyPager {
	return &BackupSchedulesClientListByBackupPolicyPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBackupPolicyCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
		},
	}
}

// listByBackupPolicyCreateRequest creates the ListByBackupPolicy request.
func (client *BackupSchedulesClient) listByBackupPolicyCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupSchedulesClientListByBackupPolicyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBackupPolicyHandleResponse handles the ListByBackupPolicy response.
func (client *BackupSchedulesClient) listByBackupPolicyHandleResponse(resp *http.Response) (BackupSchedulesClientListByBackupPolicyResponse, error) {
	result := BackupSchedulesClientListByBackupPolicyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupScheduleList); err != nil {
		return BackupSchedulesClientListByBackupPolicyResponse{}, err
	}
	return result, nil
}
