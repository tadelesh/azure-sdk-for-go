//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// DataSourcesClient contains the methods for the DataSources group.
// Don't use this type directly, use NewDataSourcesClient() instead.
type DataSourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataSourcesClient creates a new instance of DataSourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataSourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataSourcesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DataSourcesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update a data source.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceName - The name of the datasource resource.
// parameters - The parameters required to create or update a datasource.
// options - DataSourcesClientCreateOrUpdateOptions contains the optional parameters for the DataSourcesClient.CreateOrUpdate
// method.
func (client *DataSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters DataSource, options *DataSourcesClientCreateOrUpdateOptions) (DataSourcesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, parameters, options)
	if err != nil {
		return DataSourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataSourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataSourcesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataSourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters DataSource, options *DataSourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataSourcesClient) createOrUpdateHandleResponse(resp *http.Response) (DataSourcesClientCreateOrUpdateResponse, error) {
	result := DataSourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataSource); err != nil {
		return DataSourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a data source instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceName - Name of the datasource.
// options - DataSourcesClientDeleteOptions contains the optional parameters for the DataSourcesClient.Delete method.
func (client *DataSourcesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesClientDeleteOptions) (DataSourcesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, options)
	if err != nil {
		return DataSourcesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataSourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataSourcesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataSourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataSourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a datasource instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceName - Name of the datasource
// options - DataSourcesClientGetOptions contains the optional parameters for the DataSourcesClient.Get method.
func (client *DataSourcesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesClientGetOptions) (DataSourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, options)
	if err != nil {
		return DataSourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataSourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataSourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataSourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataSourcesClient) getHandleResponse(resp *http.Response) (DataSourcesClientGetResponse, error) {
	result := DataSourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataSource); err != nil {
		return DataSourcesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Gets the first page of data source instances in a workspace with the link to the next page.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// filter - The filter to apply on the operation.
// options - DataSourcesClientListByWorkspaceOptions contains the optional parameters for the DataSourcesClient.ListByWorkspace
// method.
func (client *DataSourcesClient) ListByWorkspace(resourceGroupName string, workspaceName string, filter string, options *DataSourcesClientListByWorkspaceOptions) *runtime.Pager[DataSourcesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataSourcesClientListByWorkspaceResponse]{
		More: func(page DataSourcesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataSourcesClientListByWorkspaceResponse) (DataSourcesClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataSourcesClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataSourcesClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataSourcesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *DataSourcesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, filter string, options *DataSourcesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *DataSourcesClient) listByWorkspaceHandleResponse(resp *http.Response) (DataSourcesClientListByWorkspaceResponse, error) {
	result := DataSourcesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataSourceListResult); err != nil {
		return DataSourcesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
