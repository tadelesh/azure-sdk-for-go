//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

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
	"strconv"
	"strings"
)

// PoolClient contains the methods for the Pool group.
// Don't use this type directly, use NewPoolClient() instead.
type PoolClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPoolClient creates a new instance of PoolClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPoolClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PoolClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PoolClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Creates a new pool inside the specified account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// parameters - Additional parameters for pool creation.
// options - PoolClientCreateOptions contains the optional parameters for the PoolClient.Create method.
func (client *PoolClient) Create(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, options *PoolClientCreateOptions) (PoolClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, parameters, options)
	if err != nil {
		return PoolClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoolClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoolClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *PoolClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, options *PoolClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *PoolClient) createHandleResponse(resp *http.Response) (PoolClientCreateResponse, error) {
	result := PoolClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// options - PoolClientBeginDeleteOptions contains the optional parameters for the PoolClient.BeginDelete method.
func (client *PoolClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientBeginDeleteOptions) (*armruntime.Poller[PoolClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PoolClientDeleteResponse]("PoolClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PoolClientDeleteResponse]("PoolClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PoolClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, options)
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
func (client *PoolClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// DisableAutoScale - Disables automatic scaling for a pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// options - PoolClientDisableAutoScaleOptions contains the optional parameters for the PoolClient.DisableAutoScale method.
func (client *PoolClient) DisableAutoScale(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientDisableAutoScaleOptions) (PoolClientDisableAutoScaleResponse, error) {
	req, err := client.disableAutoScaleCreateRequest(ctx, resourceGroupName, accountName, poolName, options)
	if err != nil {
		return PoolClientDisableAutoScaleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoolClientDisableAutoScaleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoolClientDisableAutoScaleResponse{}, runtime.NewResponseError(resp)
	}
	return client.disableAutoScaleHandleResponse(resp)
}

// disableAutoScaleCreateRequest creates the DisableAutoScale request.
func (client *PoolClient) disableAutoScaleCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientDisableAutoScaleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/disableAutoScale"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// disableAutoScaleHandleResponse handles the DisableAutoScale response.
func (client *PoolClient) disableAutoScaleHandleResponse(resp *http.Response) (PoolClientDisableAutoScaleResponse, error) {
	result := PoolClientDisableAutoScaleResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolClientDisableAutoScaleResponse{}, err
	}
	return result, nil
}

// Get - Gets information about the specified pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// options - PoolClientGetOptions contains the optional parameters for the PoolClient.Get method.
func (client *PoolClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientGetOptions) (PoolClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, options)
	if err != nil {
		return PoolClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoolClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoolClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PoolClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *PoolClient) getHandleResponse(resp *http.Response) (PoolClientGetResponse, error) {
	result := PoolClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolClientGetResponse{}, err
	}
	return result, nil
}

// ListByBatchAccount - Lists all of the pools in the specified account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - PoolClientListByBatchAccountOptions contains the optional parameters for the PoolClient.ListByBatchAccount method.
func (client *PoolClient) ListByBatchAccount(resourceGroupName string, accountName string, options *PoolClientListByBatchAccountOptions) *runtime.Pager[PoolClientListByBatchAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[PoolClientListByBatchAccountResponse]{
		More: func(page PoolClientListByBatchAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PoolClientListByBatchAccountResponse) (PoolClientListByBatchAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBatchAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PoolClientListByBatchAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PoolClientListByBatchAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PoolClientListByBatchAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBatchAccountHandleResponse(resp)
		},
	})
}

// listByBatchAccountCreateRequest creates the ListByBatchAccount request.
func (client *PoolClient) listByBatchAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PoolClientListByBatchAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBatchAccountHandleResponse handles the ListByBatchAccount response.
func (client *PoolClient) listByBatchAccountHandleResponse(resp *http.Response) (PoolClientListByBatchAccountResponse, error) {
	result := PoolClientListByBatchAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListPoolsResult); err != nil {
		return PoolClientListByBatchAccountResponse{}, err
	}
	return result, nil
}

// StopResize - This does not restore the pool to its previous state before the resize operation: it only stops any further
// changes being made, and the pool maintains its current state. After stopping, the pool
// stabilizes at the number of nodes it was at when the stop operation was done. During the stop operation, the pool allocation
// state changes first to stopping and then to steady. A resize operation need
// not be an explicit resize pool request; this API can also be used to halt the initial sizing of the pool when it is created.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// options - PoolClientStopResizeOptions contains the optional parameters for the PoolClient.StopResize method.
func (client *PoolClient) StopResize(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientStopResizeOptions) (PoolClientStopResizeResponse, error) {
	req, err := client.stopResizeCreateRequest(ctx, resourceGroupName, accountName, poolName, options)
	if err != nil {
		return PoolClientStopResizeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoolClientStopResizeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoolClientStopResizeResponse{}, runtime.NewResponseError(resp)
	}
	return client.stopResizeHandleResponse(resp)
}

// stopResizeCreateRequest creates the StopResize request.
func (client *PoolClient) stopResizeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *PoolClientStopResizeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/stopResize"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopResizeHandleResponse handles the StopResize response.
func (client *PoolClient) stopResizeHandleResponse(resp *http.Response) (PoolClientStopResizeResponse, error) {
	result := PoolClientStopResizeResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolClientStopResizeResponse{}, err
	}
	return result, nil
}

// Update - Updates the properties of an existing pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// poolName - The pool name. This must be unique within the account.
// parameters - Pool properties that should be updated. Properties that are supplied will be updated, any property not supplied
// will be unchanged.
// options - PoolClientUpdateOptions contains the optional parameters for the PoolClient.Update method.
func (client *PoolClient) Update(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, options *PoolClientUpdateOptions) (PoolClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, parameters, options)
	if err != nil {
		return PoolClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoolClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoolClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PoolClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, options *PoolClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PoolClient) updateHandleResponse(resp *http.Response) (PoolClientUpdateResponse, error) {
	result := PoolClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolClientUpdateResponse{}, err
	}
	return result, nil
}
