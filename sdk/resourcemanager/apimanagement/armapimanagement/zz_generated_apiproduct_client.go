//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// APIProductClient contains the methods for the APIProduct group.
// Don't use this type directly, use NewAPIProductClient() instead.
type APIProductClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIProductClient creates a new instance of APIProductClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIProductClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *APIProductClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &APIProductClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListByApis - Lists all Products, which the API is part of.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// options - APIProductClientListByApisOptions contains the optional parameters for the APIProductClient.ListByApis method.
func (client *APIProductClient) ListByApis(resourceGroupName string, serviceName string, apiID string, options *APIProductClientListByApisOptions) *runtime.Pager[APIProductClientListByApisResponse] {
	return runtime.NewPager(runtime.PageProcessor[APIProductClientListByApisResponse]{
		More: func(page APIProductClientListByApisResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIProductClientListByApisResponse) (APIProductClientListByApisResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByApisCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIProductClientListByApisResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIProductClientListByApisResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIProductClientListByApisResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByApisHandleResponse(resp)
		},
	})
}

// listByApisCreateRequest creates the ListByApis request.
func (client *APIProductClient) listByApisCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIProductClientListByApisOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByApisHandleResponse handles the ListByApis response.
func (client *APIProductClient) listByApisHandleResponse(resp *http.Response) (APIProductClientListByApisResponse, error) {
	result := APIProductClientListByApisResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductCollection); err != nil {
		return APIProductClientListByApisResponse{}, err
	}
	return result, nil
}
