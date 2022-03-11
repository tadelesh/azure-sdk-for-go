//go:build go1.16
// +build go1.16

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

// ConfigurationStoresClient contains the methods for the ConfigurationStores group.
// Don't use this type directly, use NewConfigurationStoresClient() instead.
type ConfigurationStoresClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationStoresClient creates a new instance of ConfigurationStoresClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationStoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationStoresClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ConfigurationStoresClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Creates a configuration store with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// configStoreCreationParameters - The parameters for creating a configuration store.
// options - ConfigurationStoresClientBeginCreateOptions contains the optional parameters for the ConfigurationStoresClient.BeginCreate
// method.
func (client *ConfigurationStoresClient) BeginCreate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresClientBeginCreateOptions) (ConfigurationStoresClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
	if err != nil {
		return ConfigurationStoresClientCreatePollerResponse{}, err
	}
	result := ConfigurationStoresClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Create", "", resp, client.pl)
	if err != nil {
		return ConfigurationStoresClientCreatePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a configuration store with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationStoresClient) create(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
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

// createCreateRequest creates the Create request.
func (client *ConfigurationStoresClient) createCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configStoreCreationParameters)
}

// BeginDelete - Deletes a configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// options - ConfigurationStoresClientBeginDeleteOptions contains the optional parameters for the ConfigurationStoresClient.BeginDelete
// method.
func (client *ConfigurationStoresClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientBeginDeleteOptions) (ConfigurationStoresClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return ConfigurationStoresClientDeletePollerResponse{}, err
	}
	result := ConfigurationStoresClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Delete", "", resp, client.pl)
	if err != nil {
		return ConfigurationStoresClientDeletePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationStoresClient) deleteOperation(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configStoreName, options)
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
func (client *ConfigurationStoresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
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

// Get - Gets the properties of the specified configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// options - ConfigurationStoresClientGetOptions contains the optional parameters for the ConfigurationStoresClient.Get method.
func (client *ConfigurationStoresClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientGetOptions) (ConfigurationStoresClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return ConfigurationStoresClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationStoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationStoresClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationStoresClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
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

// getHandleResponse handles the Get response.
func (client *ConfigurationStoresClient) getHandleResponse(resp *http.Response) (ConfigurationStoresClientGetResponse, error) {
	result := ConfigurationStoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStore); err != nil {
		return ConfigurationStoresClientGetResponse{}, err
	}
	return result, nil
}

// GetDeleted - Gets a deleted Azure app configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location in which uniqueness will be verified.
// configStoreName - The name of the configuration store.
// options - ConfigurationStoresClientGetDeletedOptions contains the optional parameters for the ConfigurationStoresClient.GetDeleted
// method.
func (client *ConfigurationStoresClient) GetDeleted(ctx context.Context, location string, configStoreName string, options *ConfigurationStoresClientGetDeletedOptions) (ConfigurationStoresClientGetDeletedResponse, error) {
	req, err := client.getDeletedCreateRequest(ctx, location, configStoreName, options)
	if err != nil {
		return ConfigurationStoresClientGetDeletedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationStoresClientGetDeletedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationStoresClientGetDeletedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedHandleResponse(resp)
}

// getDeletedCreateRequest creates the GetDeleted request.
func (client *ConfigurationStoresClient) getDeletedCreateRequest(ctx context.Context, location string, configStoreName string, options *ConfigurationStoresClientGetDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/locations/{location}/deletedConfigurationStores/{configStoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// getDeletedHandleResponse handles the GetDeleted response.
func (client *ConfigurationStoresClient) getDeletedHandleResponse(resp *http.Response) (ConfigurationStoresClientGetDeletedResponse, error) {
	result := ConfigurationStoresClientGetDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedConfigurationStore); err != nil {
		return ConfigurationStoresClientGetDeletedResponse{}, err
	}
	return result, nil
}

// List - Lists the configuration stores for a given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConfigurationStoresClientListOptions contains the optional parameters for the ConfigurationStoresClient.List
// method.
func (client *ConfigurationStoresClient) List(options *ConfigurationStoresClientListOptions) *ConfigurationStoresClientListPager {
	return &ConfigurationStoresClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ConfigurationStoresClient) listCreateRequest(ctx context.Context, options *ConfigurationStoresClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationStoresClient) listHandleResponse(resp *http.Response) (ConfigurationStoresClientListResponse, error) {
	result := ConfigurationStoresClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStoreListResult); err != nil {
		return ConfigurationStoresClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists the configuration stores for a given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// options - ConfigurationStoresClientListByResourceGroupOptions contains the optional parameters for the ConfigurationStoresClient.ListByResourceGroup
// method.
func (client *ConfigurationStoresClient) ListByResourceGroup(resourceGroupName string, options *ConfigurationStoresClientListByResourceGroupOptions) *ConfigurationStoresClientListByResourceGroupPager {
	return &ConfigurationStoresClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationStoresClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationStoresClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationStoresClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationStoresClientListByResourceGroupResponse, error) {
	result := ConfigurationStoresClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStoreListResult); err != nil {
		return ConfigurationStoresClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListDeleted - Gets information about the deleted configuration stores in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConfigurationStoresClientListDeletedOptions contains the optional parameters for the ConfigurationStoresClient.ListDeleted
// method.
func (client *ConfigurationStoresClient) ListDeleted(options *ConfigurationStoresClientListDeletedOptions) *ConfigurationStoresClientListDeletedPager {
	return &ConfigurationStoresClientListDeletedPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listDeletedCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresClientListDeletedResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeletedConfigurationStoreListResult.NextLink)
		},
	}
}

// listDeletedCreateRequest creates the ListDeleted request.
func (client *ConfigurationStoresClient) listDeletedCreateRequest(ctx context.Context, options *ConfigurationStoresClientListDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/deletedConfigurationStores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listDeletedHandleResponse handles the ListDeleted response.
func (client *ConfigurationStoresClient) listDeletedHandleResponse(resp *http.Response) (ConfigurationStoresClientListDeletedResponse, error) {
	result := ConfigurationStoresClientListDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedConfigurationStoreListResult); err != nil {
		return ConfigurationStoresClientListDeletedResponse{}, err
	}
	return result, nil
}

// ListKeys - Lists the access key for the specified configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// options - ConfigurationStoresClientListKeysOptions contains the optional parameters for the ConfigurationStoresClient.ListKeys
// method.
func (client *ConfigurationStoresClient) ListKeys(resourceGroupName string, configStoreName string, options *ConfigurationStoresClientListKeysOptions) *ConfigurationStoresClientListKeysPager {
	return &ConfigurationStoresClientListKeysPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listKeysCreateRequest(ctx, resourceGroupName, configStoreName, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresClientListKeysResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.APIKeyListResult.NextLink)
		},
	}
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ConfigurationStoresClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/listKeys"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ConfigurationStoresClient) listKeysHandleResponse(resp *http.Response) (ConfigurationStoresClientListKeysResponse, error) {
	result := ConfigurationStoresClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyListResult); err != nil {
		return ConfigurationStoresClientListKeysResponse{}, err
	}
	return result, nil
}

// BeginPurgeDeleted - Permanently deletes the specified configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location in which uniqueness will be verified.
// configStoreName - The name of the configuration store.
// options - ConfigurationStoresClientBeginPurgeDeletedOptions contains the optional parameters for the ConfigurationStoresClient.BeginPurgeDeleted
// method.
func (client *ConfigurationStoresClient) BeginPurgeDeleted(ctx context.Context, location string, configStoreName string, options *ConfigurationStoresClientBeginPurgeDeletedOptions) (ConfigurationStoresClientPurgeDeletedPollerResponse, error) {
	resp, err := client.purgeDeleted(ctx, location, configStoreName, options)
	if err != nil {
		return ConfigurationStoresClientPurgeDeletedPollerResponse{}, err
	}
	result := ConfigurationStoresClientPurgeDeletedPollerResponse{}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.PurgeDeleted", "", resp, client.pl)
	if err != nil {
		return ConfigurationStoresClientPurgeDeletedPollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresClientPurgeDeletedPoller{
		pt: pt,
	}
	return result, nil
}

// PurgeDeleted - Permanently deletes the specified configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationStoresClient) purgeDeleted(ctx context.Context, location string, configStoreName string, options *ConfigurationStoresClientBeginPurgeDeletedOptions) (*http.Response, error) {
	req, err := client.purgeDeletedCreateRequest(ctx, location, configStoreName, options)
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

// purgeDeletedCreateRequest creates the PurgeDeleted request.
func (client *ConfigurationStoresClient) purgeDeletedCreateRequest(ctx context.Context, location string, configStoreName string, options *ConfigurationStoresClientBeginPurgeDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/locations/{location}/deletedConfigurationStores/{configStoreName}/purge"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// RegenerateKey - Regenerates an access key for the specified configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// regenerateKeyParameters - The parameters for regenerating an access key.
// options - ConfigurationStoresClientRegenerateKeyOptions contains the optional parameters for the ConfigurationStoresClient.RegenerateKey
// method.
func (client *ConfigurationStoresClient) RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresClientRegenerateKeyOptions) (ConfigurationStoresClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, configStoreName, regenerateKeyParameters, options)
	if err != nil {
		return ConfigurationStoresClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationStoresClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationStoresClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *ConfigurationStoresClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/regenerateKey"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateKeyParameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *ConfigurationStoresClient) regenerateKeyHandleResponse(resp *http.Response) (ConfigurationStoresClientRegenerateKeyResponse, error) {
	result := ConfigurationStoresClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKey); err != nil {
		return ConfigurationStoresClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a configuration store with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// configStoreUpdateParameters - The parameters for updating a configuration store.
// options - ConfigurationStoresClientBeginUpdateOptions contains the optional parameters for the ConfigurationStoresClient.BeginUpdate
// method.
func (client *ConfigurationStoresClient) BeginUpdate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresClientBeginUpdateOptions) (ConfigurationStoresClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
	if err != nil {
		return ConfigurationStoresClientUpdatePollerResponse{}, err
	}
	result := ConfigurationStoresClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Update", "", resp, client.pl)
	if err != nil {
		return ConfigurationStoresClientUpdatePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a configuration store with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationStoresClient) update(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
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
func (client *ConfigurationStoresClient) updateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configStoreUpdateParameters)
}
