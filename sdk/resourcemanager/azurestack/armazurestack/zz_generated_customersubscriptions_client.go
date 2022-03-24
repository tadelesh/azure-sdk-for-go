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

// CustomerSubscriptionsClient contains the methods for the CustomerSubscriptions group.
// Don't use this type directly, use NewCustomerSubscriptionsClient() instead.
type CustomerSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCustomerSubscriptionsClient creates a new instance of CustomerSubscriptionsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCustomerSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CustomerSubscriptionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CustomerSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Creates a new customer subscription under a registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// customerSubscriptionName - Name of the product.
// customerCreationParameters - Parameters use to create a customer subscription.
// options - CustomerSubscriptionsClientCreateOptions contains the optional parameters for the CustomerSubscriptionsClient.Create
// method.
func (client *CustomerSubscriptionsClient) Create(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, customerCreationParameters CustomerSubscription, options *CustomerSubscriptionsClientCreateOptions) (CustomerSubscriptionsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroup, registrationName, customerSubscriptionName, customerCreationParameters, options)
	if err != nil {
		return CustomerSubscriptionsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomerSubscriptionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomerSubscriptionsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *CustomerSubscriptionsClient) createCreateRequest(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, customerCreationParameters CustomerSubscription, options *CustomerSubscriptionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}"
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
	if customerSubscriptionName == "" {
		return nil, errors.New("parameter customerSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerSubscriptionName}", url.PathEscape(customerSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customerCreationParameters)
}

// createHandleResponse handles the Create response.
func (client *CustomerSubscriptionsClient) createHandleResponse(resp *http.Response) (CustomerSubscriptionsClientCreateResponse, error) {
	result := CustomerSubscriptionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerSubscription); err != nil {
		return CustomerSubscriptionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a customer subscription under a registration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// customerSubscriptionName - Name of the product.
// options - CustomerSubscriptionsClientDeleteOptions contains the optional parameters for the CustomerSubscriptionsClient.Delete
// method.
func (client *CustomerSubscriptionsClient) Delete(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, options *CustomerSubscriptionsClientDeleteOptions) (CustomerSubscriptionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, registrationName, customerSubscriptionName, options)
	if err != nil {
		return CustomerSubscriptionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomerSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CustomerSubscriptionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CustomerSubscriptionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomerSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, options *CustomerSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}"
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
	if customerSubscriptionName == "" {
		return nil, errors.New("parameter customerSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerSubscriptionName}", url.PathEscape(customerSubscriptionName))
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

// Get - Returns the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// customerSubscriptionName - Name of the product.
// options - CustomerSubscriptionsClientGetOptions contains the optional parameters for the CustomerSubscriptionsClient.Get
// method.
func (client *CustomerSubscriptionsClient) Get(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, options *CustomerSubscriptionsClientGetOptions) (CustomerSubscriptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, registrationName, customerSubscriptionName, options)
	if err != nil {
		return CustomerSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomerSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomerSubscriptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomerSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, options *CustomerSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}"
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
	if customerSubscriptionName == "" {
		return nil, errors.New("parameter customerSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerSubscriptionName}", url.PathEscape(customerSubscriptionName))
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
func (client *CustomerSubscriptionsClient) getHandleResponse(resp *http.Response) (CustomerSubscriptionsClientGetResponse, error) {
	result := CustomerSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerSubscription); err != nil {
		return CustomerSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns a list of products.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - Name of the resource group.
// registrationName - Name of the Azure Stack registration.
// options - CustomerSubscriptionsClientListOptions contains the optional parameters for the CustomerSubscriptionsClient.List
// method.
func (client *CustomerSubscriptionsClient) List(resourceGroup string, registrationName string, options *CustomerSubscriptionsClientListOptions) *runtime.Pager[CustomerSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[CustomerSubscriptionsClientListResponse]{
		More: func(page CustomerSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomerSubscriptionsClientListResponse) (CustomerSubscriptionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroup, registrationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomerSubscriptionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CustomerSubscriptionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomerSubscriptionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CustomerSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroup string, registrationName string, options *CustomerSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions"
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

// listHandleResponse handles the List response.
func (client *CustomerSubscriptionsClient) listHandleResponse(resp *http.Response) (CustomerSubscriptionsClientListResponse, error) {
	result := CustomerSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerSubscriptionList); err != nil {
		return CustomerSubscriptionsClientListResponse{}, err
	}
	return result, nil
}