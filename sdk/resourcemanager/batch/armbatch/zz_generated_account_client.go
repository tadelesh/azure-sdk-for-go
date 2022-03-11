//go:build go1.16
// +build go1.16

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
	"strings"
)

// AccountClient contains the methods for the BatchAccount group.
// Don't use this type directly, use NewAccountClient() instead.
type AccountClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountClient creates a new instance of AccountClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccountClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AccountClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Creates a new Batch account with the specified parameters. Existing accounts cannot be updated with this
// API and should instead be updated with the Update Batch Account API.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - A name for the Batch account which must be unique within the region. Batch account names must be between
// 3 and 24 characters in length and must use only numbers and lowercase letters. This name is
// used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For
// example: http://accountname.region.batch.azure.com/.
// parameters - Additional parameters for account creation.
// options - AccountClientBeginCreateOptions contains the optional parameters for the AccountClient.BeginCreate method.
func (client *AccountClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters, options *AccountClientBeginCreateOptions) (AccountClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return AccountClientCreatePollerResponse{}, err
	}
	result := AccountClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("AccountClient.Create", "location", resp, client.pl)
	if err != nil {
		return AccountClientCreatePollerResponse{}, err
	}
	result.Poller = &AccountClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new Batch account with the specified parameters. Existing accounts cannot be updated with this API and
// should instead be updated with the Update Batch Account API.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountClient) create(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters, options *AccountClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
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

// createCreateRequest creates the Create request.
func (client *AccountClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters, options *AccountClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Batch account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - AccountClientBeginDeleteOptions contains the optional parameters for the AccountClient.BeginDelete method.
func (client *AccountClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientBeginDeleteOptions) (AccountClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountClientDeletePollerResponse{}, err
	}
	result := AccountClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("AccountClient.Delete", "location", resp, client.pl)
	if err != nil {
		return AccountClientDeletePollerResponse{}, err
	}
	result.Poller = &AccountClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified Batch account.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, options)
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
func (client *AccountClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}"
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

// Get - Gets information about the specified Batch account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - AccountClientGetOptions contains the optional parameters for the AccountClient.Get method.
func (client *AccountClient) Get(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientGetOptions) (AccountClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountClient) getHandleResponse(resp *http.Response) (AccountClientGetResponse, error) {
	result := AccountClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountClientGetResponse{}, err
	}
	return result, nil
}

// GetKeys - This operation applies only to Batch accounts with allowedAuthenticationModes containing 'SharedKey'. If the
// Batch account doesn't contain 'SharedKey' in its allowedAuthenticationMode, clients cannot
// use shared keys to authenticate, and must use another allowedAuthenticationModes instead. In this case, getting the keys
// will fail.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - AccountClientGetKeysOptions contains the optional parameters for the AccountClient.GetKeys method.
func (client *AccountClient) GetKeys(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientGetKeysOptions) (AccountClientGetKeysResponse, error) {
	req, err := client.getKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountClientGetKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountClientGetKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountClientGetKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.getKeysHandleResponse(resp)
}

// getKeysCreateRequest creates the GetKeys request.
func (client *AccountClient) getKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientGetKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/listKeys"
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

// getKeysHandleResponse handles the GetKeys response.
func (client *AccountClient) getKeysHandleResponse(resp *http.Response) (AccountClientGetKeysResponse, error) {
	result := AccountClientGetKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return AccountClientGetKeysResponse{}, err
	}
	return result, nil
}

// List - Gets information about the Batch accounts associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AccountClientListOptions contains the optional parameters for the AccountClient.List method.
func (client *AccountClient) List(options *AccountClientListOptions) *AccountClientListPager {
	return &AccountClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AccountClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccountListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccountClient) listCreateRequest(ctx context.Context, options *AccountClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts"
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

// listHandleResponse handles the List response.
func (client *AccountClient) listHandleResponse(resp *http.Response) (AccountClientListResponse, error) {
	result := AccountClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets information about the Batch accounts associated with the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// options - AccountClientListByResourceGroupOptions contains the optional parameters for the AccountClient.ListByResourceGroup
// method.
func (client *AccountClient) ListByResourceGroup(resourceGroupName string, options *AccountClientListByResourceGroupOptions) *AccountClientListByResourceGroupPager {
	return &AccountClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AccountClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccountListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccountClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccountClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountClientListByResourceGroupResponse, error) {
	result := AccountClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListOutboundNetworkDependenciesEndpoints - Lists the endpoints that a Batch Compute Node under this Batch Account may call
// as part of Batch service administration. If you are deploying a Pool inside of a virtual network that you specify, you
// must make sure your network allows outbound access to these endpoints. Failure to allow access to these endpoints may cause
// Batch to mark the affected nodes as unusable. For more information about
// creating a pool inside of a virtual network, see https://docs.microsoft.com/en-us/azure/batch/batch-virtual-network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - AccountClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the AccountClient.ListOutboundNetworkDependenciesEndpoints
// method.
func (client *AccountClient) ListOutboundNetworkDependenciesEndpoints(resourceGroupName string, accountName string, options *AccountClientListOutboundNetworkDependenciesEndpointsOptions) *AccountClientListOutboundNetworkDependenciesEndpointsPager {
	return &AccountClientListOutboundNetworkDependenciesEndpointsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp AccountClientListOutboundNetworkDependenciesEndpointsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OutboundEnvironmentEndpointCollection.NextLink)
		},
	}
}

// listOutboundNetworkDependenciesEndpointsCreateRequest creates the ListOutboundNetworkDependenciesEndpoints request.
func (client *AccountClient) listOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientListOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/outboundNetworkDependenciesEndpoints"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOutboundNetworkDependenciesEndpointsHandleResponse handles the ListOutboundNetworkDependenciesEndpoints response.
func (client *AccountClient) listOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (AccountClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	result := AccountClientListOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointCollection); err != nil {
		return AccountClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// RegenerateKey - This operation applies only to Batch accounts with allowedAuthenticationModes containing 'SharedKey'. If
// the Batch account doesn't contain 'SharedKey' in its allowedAuthenticationMode, clients cannot
// use shared keys to authenticate, and must use another allowedAuthenticationModes instead. In this case, regenerating the
// keys will fail.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// parameters - The type of key to regenerate.
// options - AccountClientRegenerateKeyOptions contains the optional parameters for the AccountClient.RegenerateKey method.
func (client *AccountClient) RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters AccountRegenerateKeyParameters, options *AccountClientRegenerateKeyOptions) (AccountClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return AccountClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *AccountClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters AccountRegenerateKeyParameters, options *AccountClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/regenerateKeys"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *AccountClient) regenerateKeyHandleResponse(resp *http.Response) (AccountClientRegenerateKeyResponse, error) {
	result := AccountClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return AccountClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// SynchronizeAutoStorageKeys - Synchronizes access keys for the auto-storage account configured for the specified Batch account,
// only if storage key authentication is being used.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - AccountClientSynchronizeAutoStorageKeysOptions contains the optional parameters for the AccountClient.SynchronizeAutoStorageKeys
// method.
func (client *AccountClient) SynchronizeAutoStorageKeys(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientSynchronizeAutoStorageKeysOptions) (AccountClientSynchronizeAutoStorageKeysResponse, error) {
	req, err := client.synchronizeAutoStorageKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountClientSynchronizeAutoStorageKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountClientSynchronizeAutoStorageKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AccountClientSynchronizeAutoStorageKeysResponse{}, runtime.NewResponseError(resp)
	}
	return AccountClientSynchronizeAutoStorageKeysResponse{}, nil
}

// synchronizeAutoStorageKeysCreateRequest creates the SynchronizeAutoStorageKeys request.
func (client *AccountClient) synchronizeAutoStorageKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountClientSynchronizeAutoStorageKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/syncAutoStorageKeys"
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

// Update - Updates the properties of an existing Batch account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// parameters - Additional parameters for account update.
// options - AccountClientUpdateOptions contains the optional parameters for the AccountClient.Update method.
func (client *AccountClient) Update(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters, options *AccountClientUpdateOptions) (AccountClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return AccountClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AccountClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters, options *AccountClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AccountClient) updateHandleResponse(resp *http.Response) (AccountClientUpdateResponse, error) {
	result := AccountClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountClientUpdateResponse{}, err
	}
	return result, nil
}
