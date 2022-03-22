//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSnapshotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SnapshotsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SnapshotsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Create the specified snapshot within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// snapshotName - The name of the snapshot
// body - Snapshot object supplied in the body of the operation.
// options - SnapshotsClientBeginCreateOptions contains the optional parameters for the SnapshotsClient.BeginCreate method.
func (client *SnapshotsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body Snapshot, options *SnapshotsClientBeginCreateOptions) (*armruntime.Poller[SnapshotsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, body, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SnapshotsClientCreateResponse]("SnapshotsClient.Create", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SnapshotsClientCreateResponse]("SnapshotsClient.Create", options.ResumeToken, client.pl, nil)
	}
}

// Create - Create the specified snapshot within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body Snapshot, options *SnapshotsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *SnapshotsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body Snapshot, options *SnapshotsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}"
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
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete snapshot
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// snapshotName - The name of the snapshot
// options - SnapshotsClientBeginDeleteOptions contains the optional parameters for the SnapshotsClient.BeginDelete method.
func (client *SnapshotsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*armruntime.Poller[SnapshotsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SnapshotsClientDeleteResponse]("SnapshotsClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SnapshotsClientDeleteResponse]("SnapshotsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete snapshot
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, options)
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
func (client *SnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}"
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
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get details of the specified snapshot
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// snapshotName - The name of the snapshot
// options - SnapshotsClientGetOptions contains the optional parameters for the SnapshotsClient.Get method.
func (client *SnapshotsClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, options *SnapshotsClientGetOptions) (SnapshotsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, options)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, options *SnapshotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}"
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
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
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
func (client *SnapshotsClient) getHandleResponse(resp *http.Response) (SnapshotsClientGetResponse, error) {
	result := SnapshotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all snapshots associated with the volume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// options - SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.List method.
func (client *SnapshotsClient) List(resourceGroupName string, accountName string, poolName string, volumeName string, options *SnapshotsClientListOptions) *runtime.Pager[SnapshotsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SnapshotsClientListResponse]{
		More: func(page SnapshotsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListResponse) (SnapshotsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SnapshotsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *SnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots"
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

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *http.Response) (SnapshotsClientListResponse, error) {
	result := SnapshotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotsList); err != nil {
		return SnapshotsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a snapshot
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// snapshotName - The name of the snapshot
// body - Snapshot object supplied in the body of the operation.
// options - SnapshotsClientBeginUpdateOptions contains the optional parameters for the SnapshotsClient.BeginUpdate method.
func (client *SnapshotsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body interface{}, options *SnapshotsClientBeginUpdateOptions) (*armruntime.Poller[SnapshotsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, body, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SnapshotsClientUpdateResponse]("SnapshotsClient.Update", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SnapshotsClientUpdateResponse]("SnapshotsClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a snapshot
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body interface{}, options *SnapshotsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, snapshotName, body, options)
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
func (client *SnapshotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string, body interface{}, options *SnapshotsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}"
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
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}
