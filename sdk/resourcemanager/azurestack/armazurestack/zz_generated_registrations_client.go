//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

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

// RegistrationsClient contains the methods for the Registrations group.
// Don't use this type directly, use NewRegistrationsClient() instead.
type RegistrationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegistrationsClient creates a new instance of RegistrationsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegistrationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RegistrationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update an Azure Stack registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// tokenParam - Registration token
// options - RegistrationsClientCreateOrUpdateOptions contains the optional parameters for the RegistrationsClient.CreateOrUpdate
// method.
func (client *RegistrationsClient) CreateOrUpdate(ctx context.Context, resourceGroup string, registrationName string, tokenParam RegistrationParameter, options *RegistrationsClientCreateOrUpdateOptions) (RegistrationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroup, registrationName, tokenParam, options)
	if err != nil {
		return RegistrationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RegistrationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroup string, registrationName string, tokenParam RegistrationParameter, options *RegistrationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, tokenParam)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RegistrationsClient) createOrUpdateHandleResponse(resp *http.Response) (RegistrationsClientCreateOrUpdateResponse, error) {
	result := RegistrationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Registration); err != nil {
		return RegistrationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the requested Azure Stack registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// options - RegistrationsClientDeleteOptions contains the optional parameters for the RegistrationsClient.Delete method.
func (client *RegistrationsClient) Delete(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientDeleteOptions) (RegistrationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, registrationName, options)
	if err != nil {
		return RegistrationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RegistrationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RegistrationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistrationsClient) deleteCreateRequest(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// EnableRemoteManagement - Enables remote management for device under the Azure Stack registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// options - RegistrationsClientEnableRemoteManagementOptions contains the optional parameters for the RegistrationsClient.EnableRemoteManagement
// method.
func (client *RegistrationsClient) EnableRemoteManagement(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientEnableRemoteManagementOptions) (RegistrationsClientEnableRemoteManagementResponse, error) {
	req, err := client.enableRemoteManagementCreateRequest(ctx, resourceGroup, registrationName, options)
	if err != nil {
		return RegistrationsClientEnableRemoteManagementResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientEnableRemoteManagementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationsClientEnableRemoteManagementResponse{}, runtime.NewResponseError(resp)
	}
	return RegistrationsClientEnableRemoteManagementResponse{}, nil
}

// enableRemoteManagementCreateRequest creates the EnableRemoteManagement request.
func (client *RegistrationsClient) enableRemoteManagementCreateRequest(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientEnableRemoteManagementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/enableRemoteManagement"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of an Azure Stack registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// options - RegistrationsClientGetOptions contains the optional parameters for the RegistrationsClient.Get method.
func (client *RegistrationsClient) Get(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientGetOptions) (RegistrationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, registrationName, options)
	if err != nil {
		return RegistrationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistrationsClient) getCreateRequest(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistrationsClient) getHandleResponse(resp *http.Response) (RegistrationsClientGetResponse, error) {
	result := RegistrationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Registration); err != nil {
		return RegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// GetActivationKey - Returns Azure Stack Activation Key.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// options - RegistrationsClientGetActivationKeyOptions contains the optional parameters for the RegistrationsClient.GetActivationKey
// method.
func (client *RegistrationsClient) GetActivationKey(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientGetActivationKeyOptions) (RegistrationsClientGetActivationKeyResponse, error) {
	req, err := client.getActivationKeyCreateRequest(ctx, resourceGroup, registrationName, options)
	if err != nil {
		return RegistrationsClientGetActivationKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientGetActivationKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationsClientGetActivationKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getActivationKeyHandleResponse(resp)
}

// getActivationKeyCreateRequest creates the GetActivationKey request.
func (client *RegistrationsClient) getActivationKeyCreateRequest(ctx context.Context, resourceGroup string, registrationName string, options *RegistrationsClientGetActivationKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/getactivationkey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getActivationKeyHandleResponse handles the GetActivationKey response.
func (client *RegistrationsClient) getActivationKeyHandleResponse(resp *http.Response) (RegistrationsClientGetActivationKeyResponse, error) {
	result := RegistrationsClientGetActivationKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivationKeyResult); err != nil {
		return RegistrationsClientGetActivationKeyResponse{}, err
	}
	return result, nil
}

// List - Returns a list of all registrations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// options - RegistrationsClientListOptions contains the optional parameters for the RegistrationsClient.List method.
func (client *RegistrationsClient) List(resourceGroup string, options *RegistrationsClientListOptions) *runtime.Pager[RegistrationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RegistrationsClientListResponse]{
		More: func(page RegistrationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistrationsClientListResponse) (RegistrationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroup, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistrationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RegistrationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistrationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistrationsClient) listCreateRequest(ctx context.Context, resourceGroup string, options *RegistrationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistrationsClient) listHandleResponse(resp *http.Response) (RegistrationsClientListResponse, error) {
	result := RegistrationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationList); err != nil {
		return RegistrationsClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Returns a list of all registrations under current subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RegistrationsClientListBySubscriptionOptions contains the optional parameters for the RegistrationsClient.ListBySubscription
// method.
func (client *RegistrationsClient) ListBySubscription(options *RegistrationsClientListBySubscriptionOptions) *runtime.Pager[RegistrationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[RegistrationsClientListBySubscriptionResponse]{
		More: func(page RegistrationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistrationsClientListBySubscriptionResponse) (RegistrationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistrationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RegistrationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistrationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *RegistrationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *RegistrationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStack/registrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *RegistrationsClient) listBySubscriptionHandleResponse(resp *http.Response) (RegistrationsClientListBySubscriptionResponse, error) {
	result := RegistrationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationList); err != nil {
		return RegistrationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patch an Azure Stack registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// tokenParam - Registration token
// options - RegistrationsClientUpdateOptions contains the optional parameters for the RegistrationsClient.Update method.
func (client *RegistrationsClient) Update(ctx context.Context, resourceGroup string, registrationName string, tokenParam RegistrationParameter, options *RegistrationsClientUpdateOptions) (RegistrationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroup, registrationName, tokenParam, options)
	if err != nil {
		return RegistrationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RegistrationsClient) updateCreateRequest(ctx context.Context, resourceGroup string, registrationName string, tokenParam RegistrationParameter, options *RegistrationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if registrationName == "" {
		return nil, errors.New("parameter registrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationName}", url.PathEscape(registrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, tokenParam)
}

// updateHandleResponse handles the Update response.
func (client *RegistrationsClient) updateHandleResponse(resp *http.Response) (RegistrationsClientUpdateResponse, error) {
	result := RegistrationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Registration); err != nil {
		return RegistrationsClientUpdateResponse{}, err
	}
	return result, nil
}
