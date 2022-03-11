//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

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

// AccountsClient contains the methods for the Accounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
// subscriptionID - The ID of the target subscription.
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

// CreateOrUpdate - Create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// mapsAccount - The new or updated parameters for the Maps Account.
// options - AccountsClientCreateOrUpdateOptions contains the optional parameters for the AccountsClient.CreateOrUpdate method.
func (client *AccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, mapsAccount Account, options *AccountsClientCreateOrUpdateOptions) (AccountsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, mapsAccount, options)
	if err != nil {
		return AccountsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AccountsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, mapsAccount Account, options *AccountsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, mapsAccount)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AccountsClient) createOrUpdateHandleResponse(resp *http.Response) (AccountsClientCreateOrUpdateResponse, error) {
	result := AccountsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Maps Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// options - AccountsClientDeleteOptions contains the optional parameters for the AccountsClient.Delete method.
func (client *AccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientDeleteOptions) (AccountsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccountsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a Maps Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
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
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}"
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
	reqQP.Set("api-version", "2021-12-01-preview")
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

// ListByResourceGroup - Get all Maps Accounts in a Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup
// method.
func (client *AccountsClient) ListByResourceGroup(resourceGroupName string, options *AccountsClientListByResourceGroupOptions) *AccountsClientListByResourceGroupPager {
	return &AccountsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AccountsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Accounts.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountsClientListByResourceGroupResponse, error) {
	result := AccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Accounts); err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get all Maps Accounts in a Subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.ListBySubscription
// method.
func (client *AccountsClient) ListBySubscription(options *AccountsClientListBySubscriptionOptions) *AccountsClientListBySubscriptionPager {
	return &AccountsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AccountsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Accounts.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maps/accounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (AccountsClientListBySubscriptionResponse, error) {
	result := AccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Accounts); err != nil {
		return AccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST
// APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// options - AccountsClientListKeysOptions contains the optional parameters for the AccountsClient.ListKeys method.
func (client *AccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientListKeysOptions) (AccountsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *AccountsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/listKeys"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *AccountsClient) listKeysHandleResponse(resp *http.Response) (AccountsClientListKeysResponse, error) {
	result := AccountsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	return result, nil
}

// ListSas - Create and list an account shared access signature token. Use this SAS token for authentication to Azure Maps
// REST APIs through various Azure Maps SDKs. As prerequisite to create a SAS Token.
// Prerequisites:
// 1. Create or have an existing User Assigned Managed Identity in the same Azure region as the account.
// 2. Create or update an Azure Map account with the same Azure region as the User Assigned Managed Identity is placed.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// mapsAccountSasParameters - The updated parameters for the Maps Account.
// options - AccountsClientListSasOptions contains the optional parameters for the AccountsClient.ListSas method.
func (client *AccountsClient) ListSas(ctx context.Context, resourceGroupName string, accountName string, mapsAccountSasParameters AccountSasParameters, options *AccountsClientListSasOptions) (AccountsClientListSasResponse, error) {
	req, err := client.listSasCreateRequest(ctx, resourceGroupName, accountName, mapsAccountSasParameters, options)
	if err != nil {
		return AccountsClientListSasResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientListSasResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientListSasResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSasHandleResponse(resp)
}

// listSasCreateRequest creates the ListSas request.
func (client *AccountsClient) listSasCreateRequest(ctx context.Context, resourceGroupName string, accountName string, mapsAccountSasParameters AccountSasParameters, options *AccountsClientListSasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/listSas"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, mapsAccountSasParameters)
}

// listSasHandleResponse handles the ListSas response.
func (client *AccountsClient) listSasHandleResponse(resp *http.Response) (AccountsClientListSasResponse, error) {
	result := AccountsClientListSasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountSasToken); err != nil {
		return AccountsClientListSasResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerate either the primary or secondary key for use with the Maps APIs. The old key will stop working
// immediately.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// keySpecification - Which key to regenerate: primary or secondary.
// options - AccountsClientRegenerateKeysOptions contains the optional parameters for the AccountsClient.RegenerateKeys method.
func (client *AccountsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, keySpecification KeySpecification, options *AccountsClientRegenerateKeysOptions) (AccountsClientRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, accountName, keySpecification, options)
	if err != nil {
		return AccountsClientRegenerateKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientRegenerateKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *AccountsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, keySpecification KeySpecification, options *AccountsClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/regenerateKey"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, keySpecification)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *AccountsClient) regenerateKeysHandleResponse(resp *http.Response) (AccountsClientRegenerateKeysResponse, error) {
	result := AccountsClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return AccountsClientRegenerateKeysResponse{}, err
	}
	return result, nil
}

// Update - Updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// mapsAccountUpdateParameters - The updated parameters for the Maps Account.
// options - AccountsClientUpdateOptions contains the optional parameters for the AccountsClient.Update method.
func (client *AccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, mapsAccountUpdateParameters AccountUpdateParameters, options *AccountsClientUpdateOptions) (AccountsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, mapsAccountUpdateParameters, options)
	if err != nil {
		return AccountsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, mapsAccountUpdateParameters AccountUpdateParameters, options *AccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, mapsAccountUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *AccountsClient) updateHandleResponse(resp *http.Response) (AccountsClientUpdateResponse, error) {
	result := AccountsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountsClientUpdateResponse{}, err
	}
	return result, nil
}
