//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

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

// DataStoresClient contains the methods for the DataStores group.
// Don't use this type directly, use NewDataStoresClient() instead.
type DataStoresClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataStoresClient creates a new instance of DataStoresClient with the specified values.
// subscriptionID - The Subscription Id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataStoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataStoresClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DataStoresClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the data store/repository in the data manager.
// If the operation fails it returns an *azcore.ResponseError type.
// dataStoreName - The data store/repository name to be created or updated.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// dataStore - The data store/repository object to be created or updated.
// options - DataStoresClientBeginCreateOrUpdateOptions contains the optional parameters for the DataStoresClient.BeginCreateOrUpdate
// method.
func (client *DataStoresClient) BeginCreateOrUpdate(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, dataStore DataStore, options *DataStoresClientBeginCreateOrUpdateOptions) (DataStoresClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, dataStoreName, resourceGroupName, dataManagerName, dataStore, options)
	if err != nil {
		return DataStoresClientCreateOrUpdatePollerResponse{}, err
	}
	result := DataStoresClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("DataStoresClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return DataStoresClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DataStoresClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the data store/repository in the data manager.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataStoresClient) createOrUpdate(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, dataStore DataStore, options *DataStoresClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, dataStoreName, resourceGroupName, dataManagerName, dataStore, options)
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
func (client *DataStoresClient) createOrUpdateCreateRequest(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, dataStore DataStore, options *DataStoresClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}"
	if dataStoreName == "" {
		return nil, errors.New("parameter dataStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataStoreName}", url.PathEscape(dataStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataStore)
}

// BeginDelete - This method deletes the given data store/repository.
// If the operation fails it returns an *azcore.ResponseError type.
// dataStoreName - The data store/repository name to be deleted.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - DataStoresClientBeginDeleteOptions contains the optional parameters for the DataStoresClient.BeginDelete method.
func (client *DataStoresClient) BeginDelete(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, options *DataStoresClientBeginDeleteOptions) (DataStoresClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, dataStoreName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return DataStoresClientDeletePollerResponse{}, err
	}
	result := DataStoresClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("DataStoresClient.Delete", "", resp, client.pl)
	if err != nil {
		return DataStoresClientDeletePollerResponse{}, err
	}
	result.Poller = &DataStoresClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - This method deletes the given data store/repository.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataStoresClient) deleteOperation(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, options *DataStoresClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, dataStoreName, resourceGroupName, dataManagerName, options)
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
func (client *DataStoresClient) deleteCreateRequest(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, options *DataStoresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}"
	if dataStoreName == "" {
		return nil, errors.New("parameter dataStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataStoreName}", url.PathEscape(dataStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - This method gets the data store/repository by name.
// If the operation fails it returns an *azcore.ResponseError type.
// dataStoreName - The data store/repository name queried.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - DataStoresClientGetOptions contains the optional parameters for the DataStoresClient.Get method.
func (client *DataStoresClient) Get(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, options *DataStoresClientGetOptions) (DataStoresClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, dataStoreName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return DataStoresClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataStoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataStoresClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataStoresClient) getCreateRequest(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string, options *DataStoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}"
	if dataStoreName == "" {
		return nil, errors.New("parameter dataStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataStoreName}", url.PathEscape(dataStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataStoresClient) getHandleResponse(resp *http.Response) (DataStoresClientGetResponse, error) {
	result := DataStoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataStore); err != nil {
		return DataStoresClientGetResponse{}, err
	}
	return result, nil
}

// ListByDataManager - Gets all the data stores/repositories in the given resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - DataStoresClientListByDataManagerOptions contains the optional parameters for the DataStoresClient.ListByDataManager
// method.
func (client *DataStoresClient) ListByDataManager(resourceGroupName string, dataManagerName string, options *DataStoresClientListByDataManagerOptions) *DataStoresClientListByDataManagerPager {
	return &DataStoresClientListByDataManagerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
		},
		advancer: func(ctx context.Context, resp DataStoresClientListByDataManagerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DataStoreList.NextLink)
		},
	}
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *DataStoresClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *DataStoresClientListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *DataStoresClient) listByDataManagerHandleResponse(resp *http.Response) (DataStoresClientListByDataManagerResponse, error) {
	result := DataStoresClientListByDataManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataStoreList); err != nil {
		return DataStoresClientListByDataManagerResponse{}, err
	}
	return result, nil
}
