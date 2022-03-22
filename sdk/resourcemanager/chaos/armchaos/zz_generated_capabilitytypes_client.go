//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

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

// CapabilityTypesClient contains the methods for the CapabilityTypes group.
// Don't use this type directly, use NewCapabilityTypesClient() instead.
type CapabilityTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCapabilityTypesClient creates a new instance of CapabilityTypesClient with the specified values.
// subscriptionID - GUID that represents an Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCapabilityTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CapabilityTypesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CapabilityTypesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get a Capability Type resource for given Target Type and location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - String that represents a Location resource name.
// targetTypeName - String that represents a Target Type resource name.
// capabilityTypeName - String that represents a Capability Type resource name.
// options - CapabilityTypesClientGetOptions contains the optional parameters for the CapabilityTypesClient.Get method.
func (client *CapabilityTypesClient) Get(ctx context.Context, locationName string, targetTypeName string, capabilityTypeName string, options *CapabilityTypesClientGetOptions) (CapabilityTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, targetTypeName, capabilityTypeName, options)
	if err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapabilityTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CapabilityTypesClient) getCreateRequest(ctx context.Context, locationName string, targetTypeName string, capabilityTypeName string, options *CapabilityTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}/capabilityTypes/{capabilityTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if targetTypeName == "" {
		return nil, errors.New("parameter targetTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetTypeName}", url.PathEscape(targetTypeName))
	if capabilityTypeName == "" {
		return nil, errors.New("parameter capabilityTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityTypeName}", url.PathEscape(capabilityTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapabilityTypesClient) getHandleResponse(resp *http.Response) (CapabilityTypesClientGetResponse, error) {
	result := CapabilityTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityType); err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	return result, nil
}

// List - Get a list of Capability Type resources for given Target Type and location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - String that represents a Location resource name.
// targetTypeName - String that represents a Target Type resource name.
// options - CapabilityTypesClientListOptions contains the optional parameters for the CapabilityTypesClient.List method.
func (client *CapabilityTypesClient) List(locationName string, targetTypeName string, options *CapabilityTypesClientListOptions) *runtime.Pager[CapabilityTypesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[CapabilityTypesClientListResponse]{
		More: func(page CapabilityTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapabilityTypesClientListResponse) (CapabilityTypesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, targetTypeName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CapabilityTypesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CapabilityTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CapabilityTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CapabilityTypesClient) listCreateRequest(ctx context.Context, locationName string, targetTypeName string, options *CapabilityTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}/capabilityTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if targetTypeName == "" {
		return nil, errors.New("parameter targetTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetTypeName}", url.PathEscape(targetTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapabilityTypesClient) listHandleResponse(resp *http.Response) (CapabilityTypesClientListResponse, error) {
	result := CapabilityTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityTypeListResult); err != nil {
		return CapabilityTypesClientListResponse{}, err
	}
	return result, nil
}
