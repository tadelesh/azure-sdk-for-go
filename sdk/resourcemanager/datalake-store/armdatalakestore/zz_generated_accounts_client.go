//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

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

// AccountsClient contains the methods for the Accounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccountsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AccountsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Checks whether the specified account name is available or taken.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The resource location without whitespace.
// parameters - Parameters supplied to check the Data Lake Store account name availability.
// options - AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
func (client *AccountsClient) CheckNameAvailability(ctx context.Context, location string, parameters CheckNameAvailabilityParameters, options *AccountsClientCheckNameAvailabilityOptions) (AccountsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *AccountsClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, parameters CheckNameAvailabilityParameters, options *AccountsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *AccountsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (AccountsClientCheckNameAvailabilityResponse, error) {
	result := AccountsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityInformation); err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Store account.
// parameters - Parameters supplied to create the Data Lake Store account.
// options - AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
func (client *AccountsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, parameters CreateDataLakeStoreAccountParameters, options *AccountsClientBeginCreateOptions) (AccountsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return AccountsClientCreatePollerResponse{}, err
	}
	result := AccountsClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("AccountsClient.Create", "", resp, client.pl)
	if err != nil {
		return AccountsClientCreatePollerResponse{}, err
	}
	result.Poller = &AccountsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountsClient) create(ctx context.Context, resourceGroupName string, accountName string, parameters CreateDataLakeStoreAccountParameters, options *AccountsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
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
func (client *AccountsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters CreateDataLakeStoreAccountParameters, options *AccountsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Store account.
// options - AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
func (client *AccountsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (AccountsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientDeletePollerResponse{}, err
	}
	result := AccountsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("AccountsClient.Delete", "", resp, client.pl)
	if err != nil {
		return AccountsClientDeletePollerResponse{}, err
	}
	result.Poller = &AccountsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (*http.Response, error) {
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
func (client *AccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// EnableKeyVault - Attempts to enable a user managed Key Vault for encryption of the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Store account.
// options - AccountsClientEnableKeyVaultOptions contains the optional parameters for the AccountsClient.EnableKeyVault method.
func (client *AccountsClient) EnableKeyVault(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientEnableKeyVaultOptions) (AccountsClientEnableKeyVaultResponse, error) {
	req, err := client.enableKeyVaultCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientEnableKeyVaultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientEnableKeyVaultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientEnableKeyVaultResponse{}, runtime.NewResponseError(resp)
	}
	return AccountsClientEnableKeyVaultResponse{}, nil
}

// enableKeyVaultCreateRequest creates the EnableKeyVault request.
func (client *AccountsClient) enableKeyVaultCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientEnableKeyVaultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/enableKeyVault"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Store account.
// options - AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
func (client *AccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientGetOptions) (AccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}"
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
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountsClient) getHandleResponse(resp *http.Response) (AccountsClientGetResponse, error) {
	result := AccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the Data Lake Store accounts within the subscription. The response includes a link to the next page of results,
// if any.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AccountsClientListOptions contains the optional parameters for the AccountsClient.List method.
func (client *AccountsClient) List(options *AccountsClientListOptions) *AccountsClientListPager {
	return &AccountsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AccountsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccountListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccountsClient) listCreateRequest(ctx context.Context, options *AccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountsClient) listHandleResponse(resp *http.Response) (AccountsClientListResponse, error) {
	result := AccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists the Data Lake Store accounts within a specific resource group. The response includes a link
// to the next page of results, if any.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// options - AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup
// method.
func (client *AccountsClient) ListByResourceGroup(resourceGroupName string, options *AccountsClientListByResourceGroupOptions) *AccountsClientListByResourceGroupPager {
	return &AccountsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AccountsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccountListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountsClientListByResourceGroupResponse, error) {
	result := AccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the specified Data Lake Store account information.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Store account.
// parameters - Parameters supplied to update the Data Lake Store account.
// options - AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
func (client *AccountsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters UpdateDataLakeStoreAccountParameters, options *AccountsClientBeginUpdateOptions) (AccountsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return AccountsClientUpdatePollerResponse{}, err
	}
	result := AccountsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("AccountsClient.Update", "", resp, client.pl)
	if err != nil {
		return AccountsClientUpdatePollerResponse{}, err
	}
	result.Poller = &AccountsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the specified Data Lake Store account information.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AccountsClient) update(ctx context.Context, resourceGroupName string, accountName string, parameters UpdateDataLakeStoreAccountParameters, options *AccountsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters UpdateDataLakeStoreAccountParameters, options *AccountsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
