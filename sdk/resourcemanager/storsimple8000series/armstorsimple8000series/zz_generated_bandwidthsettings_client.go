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

// BandwidthSettingsClient contains the methods for the BandwidthSettings group.
// Don't use this type directly, use NewBandwidthSettingsClient() instead.
type BandwidthSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBandwidthSettingsClient creates a new instance of BandwidthSettingsClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBandwidthSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BandwidthSettingsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BandwidthSettingsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the bandwidth setting
// If the operation fails it returns an *azcore.ResponseError type.
// bandwidthSettingName - The bandwidth setting name.
// resourceGroupName - The resource group name
// managerName - The manager name
// parameters - The bandwidth setting to be added or updated.
// options - BandwidthSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the BandwidthSettingsClient.BeginCreateOrUpdate
// method.
func (client *BandwidthSettingsClient) BeginCreateOrUpdate(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, parameters BandwidthSetting, options *BandwidthSettingsClientBeginCreateOrUpdateOptions) (BandwidthSettingsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, bandwidthSettingName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return BandwidthSettingsClientCreateOrUpdatePollerResponse{}, err
	}
	result := BandwidthSettingsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("BandwidthSettingsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return BandwidthSettingsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BandwidthSettingsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the bandwidth setting
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BandwidthSettingsClient) createOrUpdate(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, parameters BandwidthSetting, options *BandwidthSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, bandwidthSettingName, resourceGroupName, managerName, parameters, options)
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
func (client *BandwidthSettingsClient) createOrUpdateCreateRequest(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, parameters BandwidthSetting, options *BandwidthSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}"
	urlPath = strings.ReplaceAll(urlPath, "{bandwidthSettingName}", bandwidthSettingName)
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

// BeginDelete - Deletes the bandwidth setting
// If the operation fails it returns an *azcore.ResponseError type.
// bandwidthSettingName - The name of the bandwidth setting.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BandwidthSettingsClientBeginDeleteOptions contains the optional parameters for the BandwidthSettingsClient.BeginDelete
// method.
func (client *BandwidthSettingsClient) BeginDelete(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, options *BandwidthSettingsClientBeginDeleteOptions) (BandwidthSettingsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, bandwidthSettingName, resourceGroupName, managerName, options)
	if err != nil {
		return BandwidthSettingsClientDeletePollerResponse{}, err
	}
	result := BandwidthSettingsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("BandwidthSettingsClient.Delete", "", resp, client.pl)
	if err != nil {
		return BandwidthSettingsClientDeletePollerResponse{}, err
	}
	result.Poller = &BandwidthSettingsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the bandwidth setting
// If the operation fails it returns an *azcore.ResponseError type.
func (client *BandwidthSettingsClient) deleteOperation(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, options *BandwidthSettingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, bandwidthSettingName, resourceGroupName, managerName, options)
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
func (client *BandwidthSettingsClient) deleteCreateRequest(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, options *BandwidthSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}"
	urlPath = strings.ReplaceAll(urlPath, "{bandwidthSettingName}", bandwidthSettingName)
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

// Get - Returns the properties of the specified bandwidth setting name.
// If the operation fails it returns an *azcore.ResponseError type.
// bandwidthSettingName - The name of bandwidth setting to be fetched.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BandwidthSettingsClientGetOptions contains the optional parameters for the BandwidthSettingsClient.Get method.
func (client *BandwidthSettingsClient) Get(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, options *BandwidthSettingsClientGetOptions) (BandwidthSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, bandwidthSettingName, resourceGroupName, managerName, options)
	if err != nil {
		return BandwidthSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BandwidthSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BandwidthSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BandwidthSettingsClient) getCreateRequest(ctx context.Context, bandwidthSettingName string, resourceGroupName string, managerName string, options *BandwidthSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}"
	urlPath = strings.ReplaceAll(urlPath, "{bandwidthSettingName}", bandwidthSettingName)
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
func (client *BandwidthSettingsClient) getHandleResponse(resp *http.Response) (BandwidthSettingsClientGetResponse, error) {
	result := BandwidthSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSetting); err != nil {
		return BandwidthSettingsClientGetResponse{}, err
	}
	return result, nil
}

// ListByManager - Retrieves all the bandwidth setting in a manager.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - BandwidthSettingsClientListByManagerOptions contains the optional parameters for the BandwidthSettingsClient.ListByManager
// method.
func (client *BandwidthSettingsClient) ListByManager(resourceGroupName string, managerName string, options *BandwidthSettingsClientListByManagerOptions) *BandwidthSettingsClientListByManagerPager {
	return &BandwidthSettingsClientListByManagerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
		},
	}
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *BandwidthSettingsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *BandwidthSettingsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings"
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

// listByManagerHandleResponse handles the ListByManager response.
func (client *BandwidthSettingsClient) listByManagerHandleResponse(resp *http.Response) (BandwidthSettingsClientListByManagerResponse, error) {
	result := BandwidthSettingsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSettingList); err != nil {
		return BandwidthSettingsClientListByManagerResponse{}, err
	}
	return result, nil
}
