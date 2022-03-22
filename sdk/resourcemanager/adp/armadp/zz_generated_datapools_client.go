//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadp

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

// DataPoolsClient contains the methods for the DataPools group.
// Don't use this type directly, use NewDataPoolsClient() instead.
type DataPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataPoolsClient creates a new instance of DataPoolsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataPoolsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DataPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the ADP account
// dataPoolName - The name of the Data Pool
// options - DataPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the DataPoolsClient.BeginCreateOrUpdate
// method.
func (client *DataPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DataPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, dataPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DataPoolsClientCreateOrUpdateResponse]("DataPoolsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DataPoolsClientCreateOrUpdateResponse]("DataPoolsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, dataPoolName, options)
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
func (client *DataPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/accounts/{accountName}/dataPools/{dataPoolName}"
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
	if dataPoolName == "" {
		return nil, errors.New("parameter dataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPoolName}", url.PathEscape(dataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// BeginDelete - Deletes a Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the ADP account
// dataPoolName - The name of the Data Pool
// options - DataPoolsClientBeginDeleteOptions contains the optional parameters for the DataPoolsClient.BeginDelete method.
func (client *DataPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginDeleteOptions) (*armruntime.Poller[DataPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, dataPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DataPoolsClientDeleteResponse]("DataPoolsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DataPoolsClientDeleteResponse]("DataPoolsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, dataPoolName, options)
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
func (client *DataPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/accounts/{accountName}/dataPools/{dataPoolName}"
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
	if dataPoolName == "" {
		return nil, errors.New("parameter dataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPoolName}", url.PathEscape(dataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the properties of a Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the ADP account
// dataPoolName - The name of the Data Pool
// options - DataPoolsClientGetOptions contains the optional parameters for the DataPoolsClient.Get method.
func (client *DataPoolsClient) Get(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientGetOptions) (DataPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, dataPoolName, options)
	if err != nil {
		return DataPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/accounts/{accountName}/dataPools/{dataPoolName}"
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
	if dataPoolName == "" {
		return nil, errors.New("parameter dataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPoolName}", url.PathEscape(dataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataPoolsClient) getHandleResponse(resp *http.Response) (DataPoolsClientGetResponse, error) {
	result := DataPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPool); err != nil {
		return DataPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the data pools under the ADP account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the ADP account
// options - DataPoolsClientListOptions contains the optional parameters for the DataPoolsClient.List method.
func (client *DataPoolsClient) List(resourceGroupName string, accountName string, options *DataPoolsClientListOptions) *runtime.Pager[DataPoolsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataPoolsClientListResponse]{
		More: func(page DataPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataPoolsClientListResponse) (DataPoolsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataPoolsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataPoolsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataPoolsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DataPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *DataPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/accounts/{accountName}/dataPools"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataPoolsClient) listHandleResponse(resp *http.Response) (DataPoolsClientListResponse, error) {
	result := DataPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPoolList); err != nil {
		return DataPoolsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the properties of an existing Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the ADP account
// dataPoolName - The name of the Data Pool
// options - DataPoolsClientBeginUpdateOptions contains the optional parameters for the DataPoolsClient.BeginUpdate method.
func (client *DataPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginUpdateOptions) (*armruntime.Poller[DataPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, dataPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DataPoolsClientUpdateResponse]("DataPoolsClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DataPoolsClientUpdateResponse]("DataPoolsClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the properties of an existing Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataPoolsClient) update(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, dataPoolName, options)
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

// updateCreateRequest creates the Update request.
func (client *DataPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, options *DataPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/accounts/{accountName}/dataPools/{dataPoolName}"
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
	if dataPoolName == "" {
		return nil, errors.New("parameter dataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPoolName}", url.PathEscape(dataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}
