//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Update the state of the specified private endpoint connection associated with the configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// privateEndpointConnectionName - Private endpoint connection name
// privateEndpointConnection - The private endpoint connection properties.
// options - PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
func (client *PrivateEndpointConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[PrivateEndpointConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, configStoreName, privateEndpointConnectionName, privateEndpointConnection, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionsClientCreateOrUpdateResponse]("PrivateEndpointConnectionsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientCreateOrUpdateResponse]("PrivateEndpointConnectionsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Update the state of the specified private endpoint connection associated with the configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, configStoreName, privateEndpointConnectionName, privateEndpointConnection, options)
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
func (client *PrivateEndpointConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateEndpointConnection)
}

// BeginDelete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// privateEndpointConnectionName - Private endpoint connection name
// options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*armruntime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, configStoreName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionsClientDeleteResponse]("PrivateEndpointConnectionsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientDeleteResponse]("PrivateEndpointConnectionsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configStoreName, privateEndpointConnectionName, options)
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
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified private endpoint connection associated with the configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// privateEndpointConnectionName - Private endpoint connection name
// options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByConfigurationStore - Lists all private endpoint connections for a configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// options - PrivateEndpointConnectionsClientListByConfigurationStoreOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByConfigurationStore
// method.
func (client *PrivateEndpointConnectionsClient) ListByConfigurationStore(resourceGroupName string, configStoreName string, options *PrivateEndpointConnectionsClientListByConfigurationStoreOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByConfigurationStoreResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateEndpointConnectionsClientListByConfigurationStoreResponse]{
		More: func(page PrivateEndpointConnectionsClientListByConfigurationStoreResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByConfigurationStoreResponse) (PrivateEndpointConnectionsClientListByConfigurationStoreResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByConfigurationStoreCreateRequest(ctx, resourceGroupName, configStoreName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateEndpointConnectionsClientListByConfigurationStoreResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateEndpointConnectionsClientListByConfigurationStoreResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionsClientListByConfigurationStoreResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationStoreHandleResponse(resp)
		},
	})
}

// listByConfigurationStoreCreateRequest creates the ListByConfigurationStore request.
func (client *PrivateEndpointConnectionsClient) listByConfigurationStoreCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *PrivateEndpointConnectionsClientListByConfigurationStoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByConfigurationStoreHandleResponse handles the ListByConfigurationStore response.
func (client *PrivateEndpointConnectionsClient) listByConfigurationStoreHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByConfigurationStoreResponse, error) {
	result := PrivateEndpointConnectionsClientListByConfigurationStoreResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByConfigurationStoreResponse{}, err
	}
	return result, nil
}
