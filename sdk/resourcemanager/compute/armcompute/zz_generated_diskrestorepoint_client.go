//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiskRestorePointClient contains the methods for the DiskRestorePoint group.
// Don't use this type directly, use NewDiskRestorePointClient() instead.
type DiskRestorePointClient struct {
	host string
	subscriptionID string
	pl runtime.Pipeline
}

// NewDiskRestorePointClient creates a new instance of DiskRestorePointClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiskRestorePointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiskRestorePointClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DiskRestorePointClient{
		subscriptionID: subscriptionID,
		host: ep,
pl: pl,
	}
	return client, nil
}

// Get - Get disk restorePoint resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection that the disk restore point belongs.
// vmRestorePointName - The name of the vm restore point that the disk disk restore point belongs.
// diskRestorePointName - The name of the disk restore point created.
// options - DiskRestorePointClientGetOptions contains the optional parameters for the DiskRestorePointClient.Get method.
func (client *DiskRestorePointClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointClientGetOptions) (DiskRestorePointClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
	if err != nil {
		return DiskRestorePointClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskRestorePointClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskRestorePointClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskRestorePointClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskRestorePointClient) getHandleResponse(resp *http.Response) (DiskRestorePointClientGetResponse, error) {
	result := DiskRestorePointClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskRestorePoint); err != nil {
		return DiskRestorePointClientGetResponse{}, err
	}
	return result, nil
}

// BeginGrantAccess - Grants access to a diskRestorePoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection that the disk restore point belongs.
// vmRestorePointName - The name of the vm restore point that the disk disk restore point belongs.
// diskRestorePointName - The name of the disk restore point created.
// grantAccessData - Access data object supplied in the body of the get disk access operation.
// options - DiskRestorePointClientBeginGrantAccessOptions contains the optional parameters for the DiskRestorePointClient.BeginGrantAccess
// method.
func (client *DiskRestorePointClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointClientBeginGrantAccessOptions) (*armruntime.Poller[DiskRestorePointClientGrantAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.grantAccess(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, grantAccessData, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskRestorePointClientGrantAccessResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskRestorePointClientGrantAccessResponse](options.ResumeToken, client.pl, nil)
	}
}

// GrantAccess - Grants access to a diskRestorePoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskRestorePointClient) grantAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointClientBeginGrantAccessOptions) (*http.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, grantAccessData, options)
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

// grantAccessCreateRequest creates the GrantAccess request.
func (client *DiskRestorePointClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointClientBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, grantAccessData)
}

// NewListByRestorePointPager - Lists diskRestorePoints under a vmRestorePoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection that the disk restore point belongs.
// vmRestorePointName - The name of the vm restore point that the disk disk restore point belongs.
// options - DiskRestorePointClientListByRestorePointOptions contains the optional parameters for the DiskRestorePointClient.ListByRestorePoint
// method.
func (client *DiskRestorePointClient) NewListByRestorePointPager(resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointClientListByRestorePointOptions) (*runtime.Pager[DiskRestorePointClientListByRestorePointResponse]) {
	return runtime.NewPager(runtime.PageProcessor[DiskRestorePointClientListByRestorePointResponse]{
		More: func(page DiskRestorePointClientListByRestorePointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskRestorePointClientListByRestorePointResponse) (DiskRestorePointClientListByRestorePointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRestorePointCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskRestorePointClientListByRestorePointResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskRestorePointClientListByRestorePointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskRestorePointClientListByRestorePointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRestorePointHandleResponse(resp)
		},
	})
}

// listByRestorePointCreateRequest creates the ListByRestorePoint request.
func (client *DiskRestorePointClient) listByRestorePointCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointClientListByRestorePointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRestorePointHandleResponse handles the ListByRestorePoint response.
func (client *DiskRestorePointClient) listByRestorePointHandleResponse(resp *http.Response) (DiskRestorePointClientListByRestorePointResponse, error) {
	result := DiskRestorePointClientListByRestorePointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskRestorePointList); err != nil {
		return DiskRestorePointClientListByRestorePointResponse{}, err
	}
	return result, nil
}

// BeginRevokeAccess - Revokes access to a diskRestorePoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection that the disk restore point belongs.
// vmRestorePointName - The name of the vm restore point that the disk disk restore point belongs.
// diskRestorePointName - The name of the disk restore point created.
// options - DiskRestorePointClientBeginRevokeAccessOptions contains the optional parameters for the DiskRestorePointClient.BeginRevokeAccess
// method.
func (client *DiskRestorePointClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointClientBeginRevokeAccessOptions) (*armruntime.Poller[DiskRestorePointClientRevokeAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revokeAccess(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskRestorePointClientRevokeAccessResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskRestorePointClientRevokeAccessResponse](options.ResumeToken, client.pl, nil)
	}
}

// RevokeAccess - Revokes access to a diskRestorePoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskRestorePointClient) revokeAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointClientBeginRevokeAccessOptions) (*http.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
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

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *DiskRestorePointClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointClientBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

