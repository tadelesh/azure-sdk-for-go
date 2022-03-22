//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// NetworkFunctionVendorsClient contains the methods for the NetworkFunctionVendors group.
// Don't use this type directly, use NewNetworkFunctionVendorsClient() instead.
type NetworkFunctionVendorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkFunctionVendorsClient creates a new instance of NetworkFunctionVendorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNetworkFunctionVendorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NetworkFunctionVendorsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &NetworkFunctionVendorsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists all the available vendor and sku information.
// If the operation fails it returns an *azcore.ResponseError type.
// options - NetworkFunctionVendorsClientListOptions contains the optional parameters for the NetworkFunctionVendorsClient.List
// method.
func (client *NetworkFunctionVendorsClient) List(options *NetworkFunctionVendorsClientListOptions) *runtime.Pager[NetworkFunctionVendorsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[NetworkFunctionVendorsClientListResponse]{
		More: func(page NetworkFunctionVendorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkFunctionVendorsClientListResponse) (NetworkFunctionVendorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkFunctionVendorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkFunctionVendorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkFunctionVendorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *NetworkFunctionVendorsClient) listCreateRequest(ctx context.Context, options *NetworkFunctionVendorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctionVendors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkFunctionVendorsClient) listHandleResponse(resp *http.Response) (NetworkFunctionVendorsClientListResponse, error) {
	result := NetworkFunctionVendorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionVendorListResult); err != nil {
		return NetworkFunctionVendorsClientListResponse{}, err
	}
	return result, nil
}
