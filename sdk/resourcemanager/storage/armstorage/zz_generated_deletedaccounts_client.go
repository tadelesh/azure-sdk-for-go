//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// DeletedAccountsClient contains the methods for the DeletedAccounts group.
// Don't use this type directly, use NewDeletedAccountsClient() instead.
type DeletedAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDeletedAccountsClient creates a new instance of DeletedAccountsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeletedAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DeletedAccountsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DeletedAccountsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get properties of specified deleted account resource.
// If the operation fails it returns an *azcore.ResponseError type.
// deletedAccountName - Name of the deleted storage account.
// location - The location of the deleted storage account.
// options - DeletedAccountsClientGetOptions contains the optional parameters for the DeletedAccountsClient.Get method.
func (client *DeletedAccountsClient) Get(ctx context.Context, deletedAccountName string, location string, options *DeletedAccountsClientGetOptions) (DeletedAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deletedAccountName, location, options)
	if err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeletedAccountsClient) getCreateRequest(ctx context.Context, deletedAccountName string, location string, options *DeletedAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/deletedAccounts/{deletedAccountName}"
	if deletedAccountName == "" {
		return nil, errors.New("parameter deletedAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedAccountName}", url.PathEscape(deletedAccountName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeletedAccountsClient) getHandleResponse(resp *http.Response) (DeletedAccountsClientGetResponse, error) {
	result := DeletedAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedAccount); err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists deleted accounts under the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DeletedAccountsClientListOptions contains the optional parameters for the DeletedAccountsClient.List method.
func (client *DeletedAccountsClient) List(options *DeletedAccountsClientListOptions) *runtime.Pager[DeletedAccountsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[DeletedAccountsClientListResponse]{
		More: func(page DeletedAccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedAccountsClientListResponse) (DeletedAccountsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeletedAccountsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DeletedAccountsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeletedAccountsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeletedAccountsClient) listCreateRequest(ctx context.Context, options *DeletedAccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/deletedAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedAccountsClient) listHandleResponse(resp *http.Response) (DeletedAccountsClientListResponse, error) {
	result := DeletedAccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedAccountListResult); err != nil {
		return DeletedAccountsClientListResponse{}, err
	}
	return result, nil
}
