//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

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

// QuotaRequestStatusClient contains the methods for the QuotaRequestStatus group.
// Don't use this type directly, use NewQuotaRequestStatusClient() instead.
type QuotaRequestStatusClient struct {
	host string
	pl   runtime.Pipeline
}

// NewQuotaRequestStatusClient creates a new instance of QuotaRequestStatusClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewQuotaRequestStatusClient(credential azcore.TokenCredential, options *arm.ClientOptions) *QuotaRequestStatusClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &QuotaRequestStatusClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - For the specified Azure region (location), get the details and status of the quota request by the quota request ID
// for the resources of the resource provider. The PUT request for the quota (service
// limit) returns a response with the requestId parameter.
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - Azure subscription ID.
// providerID - Azure resource provider ID.
// location - Azure region.
// id - Quota Request ID.
// options - QuotaRequestStatusClientGetOptions contains the optional parameters for the QuotaRequestStatusClient.Get method.
func (client *QuotaRequestStatusClient) Get(ctx context.Context, subscriptionID string, providerID string, location string, id string, options *QuotaRequestStatusClientGetOptions) (QuotaRequestStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, subscriptionID, providerID, location, id, options)
	if err != nil {
		return QuotaRequestStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotaRequestStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotaRequestStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QuotaRequestStatusClient) getCreateRequest(ctx context.Context, subscriptionID string, providerID string, location string, id string, options *QuotaRequestStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests/{id}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if providerID == "" {
		return nil, errors.New("parameter providerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerId}", url.PathEscape(providerID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QuotaRequestStatusClient) getHandleResponse(resp *http.Response) (QuotaRequestStatusClientGetResponse, error) {
	result := QuotaRequestStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaRequestDetails); err != nil {
		return QuotaRequestStatusClientGetResponse{}, err
	}
	return result, nil
}

// List - For the specified Azure region (location), subscription, and resource provider, get the history of the quota requests
// for the past year. To select specific quota requests, use the oData filter.
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - Azure subscription ID.
// providerID - Azure resource provider ID.
// location - Azure region.
// options - QuotaRequestStatusClientListOptions contains the optional parameters for the QuotaRequestStatusClient.List method.
func (client *QuotaRequestStatusClient) List(subscriptionID string, providerID string, location string, options *QuotaRequestStatusClientListOptions) *runtime.Pager[QuotaRequestStatusClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[QuotaRequestStatusClientListResponse]{
		More: func(page QuotaRequestStatusClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotaRequestStatusClientListResponse) (QuotaRequestStatusClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, subscriptionID, providerID, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QuotaRequestStatusClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return QuotaRequestStatusClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QuotaRequestStatusClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *QuotaRequestStatusClient) listCreateRequest(ctx context.Context, subscriptionID string, providerID string, location string, options *QuotaRequestStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if providerID == "" {
		return nil, errors.New("parameter providerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerId}", url.PathEscape(providerID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-25")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QuotaRequestStatusClient) listHandleResponse(resp *http.Response) (QuotaRequestStatusClientListResponse, error) {
	result := QuotaRequestStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaRequestDetailsList); err != nil {
		return QuotaRequestStatusClientListResponse{}, err
	}
	return result, nil
}
