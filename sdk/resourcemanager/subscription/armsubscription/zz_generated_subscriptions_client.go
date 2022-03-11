//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

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

// SubscriptionsClient contains the methods for the Subscriptions group.
// Don't use this type directly, use NewSubscriptionsClient() instead.
type SubscriptionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *SubscriptionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SubscriptionsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets details about a specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - The ID of the target subscription.
// options - SubscriptionsClientGetOptions contains the optional parameters for the SubscriptionsClient.Get method.
func (client *SubscriptionsClient) Get(ctx context.Context, subscriptionID string, options *SubscriptionsClientGetOptions) (SubscriptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionsClient) getCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionsClient) getHandleResponse(resp *http.Response) (SubscriptionsClientGetResponse, error) {
	result := SubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subscription); err != nil {
		return SubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all subscriptions for a tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionsClientListOptions contains the optional parameters for the SubscriptionsClient.List method.
func (client *SubscriptionsClient) List(options *SubscriptionsClientListOptions) *SubscriptionsClientListPager {
	return &SubscriptionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SubscriptionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SubscriptionsClient) listCreateRequest(ctx context.Context, options *SubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionsClient) listHandleResponse(resp *http.Response) (SubscriptionsClientListResponse, error) {
	result := SubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SubscriptionsClientListResponse{}, err
	}
	return result, nil
}

// ListLocations - This operation provides all the locations that are available for resource providers; however, each resource
// provider may support a subset of this list.
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - The ID of the target subscription.
// options - SubscriptionsClientListLocationsOptions contains the optional parameters for the SubscriptionsClient.ListLocations
// method.
func (client *SubscriptionsClient) ListLocations(subscriptionID string, options *SubscriptionsClientListLocationsOptions) *SubscriptionsClientListLocationsPager {
	return &SubscriptionsClientListLocationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listLocationsCreateRequest(ctx, subscriptionID, options)
		},
	}
}

// listLocationsCreateRequest creates the ListLocations request.
func (client *SubscriptionsClient) listLocationsCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionsClientListLocationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/locations"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listLocationsHandleResponse handles the ListLocations response.
func (client *SubscriptionsClient) listLocationsHandleResponse(resp *http.Response) (SubscriptionsClientListLocationsResponse, error) {
	result := SubscriptionsClientListLocationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationListResult); err != nil {
		return SubscriptionsClientListLocationsResponse{}, err
	}
	return result, nil
}
