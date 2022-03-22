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

// ProductSubscriptionsClient contains the methods for the ProductSubscriptions group.
// Don't use this type directly, use NewProductSubscriptionsClient() instead.
type ProductSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProductSubscriptionsClient creates a new instance of ProductSubscriptionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProductSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProductSubscriptionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ProductSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists the collection of subscriptions to the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// options - ProductSubscriptionsClientListOptions contains the optional parameters for the ProductSubscriptionsClient.List
// method.
func (client *ProductSubscriptionsClient) List(resourceGroupName string, serviceName string, productID string, options *ProductSubscriptionsClientListOptions) *runtime.Pager[ProductSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ProductSubscriptionsClientListResponse]{
		More: func(page ProductSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductSubscriptionsClientListResponse) (ProductSubscriptionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProductSubscriptionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProductSubscriptionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProductSubscriptionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProductSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/subscriptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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

// listHandleResponse handles the List response.
func (client *ProductSubscriptionsClient) listHandleResponse(resp *http.Response) (ProductSubscriptionsClientListResponse, error) {
	result := ProductSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionCollection); err != nil {
		return ProductSubscriptionsClientListResponse{}, err
	}
	return result, nil
}
