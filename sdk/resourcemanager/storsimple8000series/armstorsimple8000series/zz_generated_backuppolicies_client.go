//go:build go1.18
// +build go1.18

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

// BackupPoliciesClient contains the methods for the BackupPolicies group.
// Don't use this type directly, use NewBackupPoliciesClient() instead.
type BackupPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BackupPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginBackupNow - Backup the backup policy now.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The backup policy name.
// backupType - The backup Type. This can be cloudSnapshot or localSnapshot.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupPoliciesClientBeginBackupNowOptions contains the optional parameters for the BackupPoliciesClient.BeginBackupNow
// method.
func (client *BackupPoliciesClient) BeginBackupNow(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginBackupNowOptions) (*armruntime.Poller[BackupPoliciesClientBackupNowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.backupNow(ctx, deviceName, backupPolicyName, backupType, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[BackupPoliciesClientBackupNowResponse]("BackupPoliciesClient.BackupNow", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[BackupPoliciesClientBackupNowResponse]("BackupPoliciesClient.BackupNow", options.ResumeToken, client.pl, nil)
	}
}

// BackupNow - Backup the backup policy now.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BackupPoliciesClient) backupNow(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginBackupNowOptions) (*http.Response, error) {
	req, err := client.backupNowCreateRequest(ctx, deviceName, backupPolicyName, backupType, resourceGroupName, managerName, options)
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
func (client *BackupPoliciesClient) backupNowCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginBackupNowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/backup"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
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

// BeginCreateOrUpdate - Creates or updates the backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The name of the backup policy to be created/updated.
// resourceGroupName - The resource group name
// managerName - The manager name
// parameters - The backup policy.
// options - BackupPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupPoliciesClient.BeginCreateOrUpdate
// method.
func (client *BackupPoliciesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[BackupPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[BackupPoliciesClientCreateOrUpdateResponse]("BackupPoliciesClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[BackupPoliciesClientCreateOrUpdateResponse]("BackupPoliciesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BackupPoliciesClient) createOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, parameters, options)
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
func (client *BackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, parameters BackupPolicy, options *BackupPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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

// BeginDelete - Deletes the backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The name of the backup policy.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupPoliciesClientBeginDeleteOptions contains the optional parameters for the BackupPoliciesClient.BeginDelete
// method.
func (client *BackupPoliciesClient) BeginDelete(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginDeleteOptions) (*armruntime.Poller[BackupPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[BackupPoliciesClientDeleteResponse]("BackupPoliciesClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[BackupPoliciesClientDeleteResponse]("BackupPoliciesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BackupPoliciesClient) deleteOperation(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
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
func (client *BackupPoliciesClient) deleteCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", backupPolicyName)
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

// Get - Gets the properties of the specified backup policy name.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// backupPolicyName - The name of backup policy to be fetched.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupPoliciesClientGetOptions contains the optional parameters for the BackupPoliciesClient.Get method.
func (client *BackupPoliciesClient) Get(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesClientGetOptions) (BackupPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, backupPolicyName, resourceGroupName, managerName, options)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupPoliciesClient) getCreateRequest(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *BackupPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}"
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

// getHandleResponse handles the Get response.
func (client *BackupPoliciesClient) getHandleResponse(resp *http.Response) (BackupPoliciesClientGetResponse, error) {
	result := BackupPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupPolicy); err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByDevice - Gets all the backup policies in a device.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BackupPoliciesClientListByDeviceOptions contains the optional parameters for the BackupPoliciesClient.ListByDevice
// method.
func (client *BackupPoliciesClient) ListByDevice(deviceName string, resourceGroupName string, managerName string, options *BackupPoliciesClientListByDeviceOptions) *runtime.Pager[BackupPoliciesClientListByDeviceResponse] {
	return runtime.NewPager(runtime.PageProcessor[BackupPoliciesClientListByDeviceResponse]{
		More: func(page BackupPoliciesClientListByDeviceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BackupPoliciesClientListByDeviceResponse) (BackupPoliciesClientListByDeviceResponse, error) {
			req, err := client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
			if err != nil {
				return BackupPoliciesClientListByDeviceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupPoliciesClientListByDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupPoliciesClientListByDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDeviceHandleResponse(resp)
		},
	})
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *BackupPoliciesClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *BackupPoliciesClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
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

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *BackupPoliciesClient) listByDeviceHandleResponse(resp *http.Response) (BackupPoliciesClientListByDeviceResponse, error) {
	result := BackupPoliciesClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupPolicyList); err != nil {
		return BackupPoliciesClientListByDeviceResponse{}, err
	}
	return result, nil
}
