//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	host string
	subscriptionID string
	pl runtime.Pipeline
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceSKUsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceSKUsClient{
		subscriptionID: subscriptionID,
		host: ep,
pl: pl,
	}
	return client, nil
}

// NewListPager - Gets the list of Microsoft.Compute SKUs available for your Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ResourceSKUsClientListOptions contains the optional parameters for the ResourceSKUsClient.List method.
func (client *ResourceSKUsClient) NewListPager(options *ResourceSKUsClientListOptions) (*runtime.Pager[ResourceSKUsClientListResponse]) {
	return runtime.NewPager(runtime.PageProcessor[ResourceSKUsClientListResponse]{
		More: func(page ResourceSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceSKUsClientListResponse) (ResourceSKUsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceSKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceSKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceSKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ResourceSKUsClient) listCreateRequest(ctx context.Context, options *ResourceSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.IncludeExtendedLocations != nil {
		reqQP.Set("includeExtendedLocations", *options.IncludeExtendedLocations)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceSKUsClient) listHandleResponse(resp *http.Response) (ResourceSKUsClientListResponse, error) {
	result := ResourceSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUsResult); err != nil {
		return ResourceSKUsClientListResponse{}, err
	}
	return result, nil
}

