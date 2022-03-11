//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

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

// IscsiTargetsClient contains the methods for the IscsiTargets group.
// Don't use this type directly, use NewIscsiTargetsClient() instead.
type IscsiTargetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIscsiTargetsClient creates a new instance of IscsiTargetsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIscsiTargetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IscsiTargetsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IscsiTargetsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// iscsiTargetName - The name of the iSCSI Target.
// iscsiTargetCreatePayload - Request payload for iSCSI Target create operation.
// options - IscsiTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginCreateOrUpdate
// method.
func (client *IscsiTargetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, options *IscsiTargetsClientBeginCreateOrUpdateOptions) (IscsiTargetsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetCreatePayload, options)
	if err != nil {
		return IscsiTargetsClientCreateOrUpdatePollerResponse{}, err
	}
	result := IscsiTargetsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("IscsiTargetsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return IscsiTargetsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IscsiTargetsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiTargetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, options *IscsiTargetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetCreatePayload, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IscsiTargetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, options *IscsiTargetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iscsiTargetCreatePayload)
}

// BeginDelete - Delete an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// iscsiTargetName - The name of the iSCSI Target.
// options - IscsiTargetsClientBeginDeleteOptions contains the optional parameters for the IscsiTargetsClient.BeginDelete
// method.
func (client *IscsiTargetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientBeginDeleteOptions) (IscsiTargetsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
	if err != nil {
		return IscsiTargetsClientDeletePollerResponse{}, err
	}
	result := IscsiTargetsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("IscsiTargetsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return IscsiTargetsClientDeletePollerResponse{}, err
	}
	result.Poller = &IscsiTargetsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiTargetsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IscsiTargetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// iscsiTargetName - The name of the iSCSI Target.
// options - IscsiTargetsClientGetOptions contains the optional parameters for the IscsiTargetsClient.Get method.
func (client *IscsiTargetsClient) Get(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientGetOptions) (IscsiTargetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
	if err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IscsiTargetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IscsiTargetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IscsiTargetsClient) getHandleResponse(resp *http.Response) (IscsiTargetsClientGetResponse, error) {
	result := IscsiTargetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IscsiTarget); err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	return result, nil
}

// ListByDiskPool - Get iSCSI Targets in a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - IscsiTargetsClientListByDiskPoolOptions contains the optional parameters for the IscsiTargetsClient.ListByDiskPool
// method.
func (client *IscsiTargetsClient) ListByDiskPool(resourceGroupName string, diskPoolName string, options *IscsiTargetsClientListByDiskPoolOptions) *IscsiTargetsClientListByDiskPoolPager {
	return &IscsiTargetsClientListByDiskPoolPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDiskPoolCreateRequest(ctx, resourceGroupName, diskPoolName, options)
		},
		advancer: func(ctx context.Context, resp IscsiTargetsClientListByDiskPoolResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IscsiTargetList.NextLink)
		},
	}
}

// listByDiskPoolCreateRequest creates the ListByDiskPool request.
func (client *IscsiTargetsClient) listByDiskPoolCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *IscsiTargetsClientListByDiskPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDiskPoolHandleResponse handles the ListByDiskPool response.
func (client *IscsiTargetsClient) listByDiskPoolHandleResponse(resp *http.Response) (IscsiTargetsClientListByDiskPoolResponse, error) {
	result := IscsiTargetsClientListByDiskPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IscsiTargetList); err != nil {
		return IscsiTargetsClientListByDiskPoolResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// iscsiTargetName - The name of the iSCSI Target.
// iscsiTargetUpdatePayload - Request payload for iSCSI Target update operation.
// options - IscsiTargetsClientBeginUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginUpdate
// method.
func (client *IscsiTargetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, options *IscsiTargetsClientBeginUpdateOptions) (IscsiTargetsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetUpdatePayload, options)
	if err != nil {
		return IscsiTargetsClientUpdatePollerResponse{}, err
	}
	result := IscsiTargetsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("IscsiTargetsClient.Update", "azure-async-operation", resp, client.pl)
	if err != nil {
		return IscsiTargetsClientUpdatePollerResponse{}, err
	}
	result.Poller = &IscsiTargetsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiTargetsClient) update(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, options *IscsiTargetsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetUpdatePayload, options)
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

// updateCreateRequest creates the Update request.
func (client *IscsiTargetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, options *IscsiTargetsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iscsiTargetUpdatePayload)
}
