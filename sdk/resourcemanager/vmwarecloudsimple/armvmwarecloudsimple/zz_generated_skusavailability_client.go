//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// SKUsAvailabilityClient contains the methods for the SKUsAvailability group.
// Don't use this type directly, use NewSKUsAvailabilityClient() instead.
type SKUsAvailabilityClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSKUsAvailabilityClient creates a new instance of SKUsAvailabilityClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSKUsAvailabilityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SKUsAvailabilityClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SKUsAvailabilityClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Returns list of available resources in region
// If the operation fails it returns an *azcore.ResponseError type.
// regionID - The region Id (westus, eastus)
// options - SKUsAvailabilityClientListOptions contains the optional parameters for the SKUsAvailabilityClient.List method.
func (client *SKUsAvailabilityClient) List(regionID string, options *SKUsAvailabilityClientListOptions) *runtime.Pager[SKUsAvailabilityClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SKUsAvailabilityClientListResponse]{
		More: func(page SKUsAvailabilityClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SKUsAvailabilityClientListResponse) (SKUsAvailabilityClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, regionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SKUsAvailabilityClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SKUsAvailabilityClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SKUsAvailabilityClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SKUsAvailabilityClient) listCreateRequest(ctx context.Context, regionID string, options *SKUsAvailabilityClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/availabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SKUID != nil {
		reqQP.Set("skuId", *options.SKUID)
	}
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SKUsAvailabilityClient) listHandleResponse(resp *http.Response) (SKUsAvailabilityClientListResponse, error) {
	result := SKUsAvailabilityClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUAvailabilityListResponse); err != nil {
		return SKUsAvailabilityClientListResponse{}, err
	}
	return result, nil
}
