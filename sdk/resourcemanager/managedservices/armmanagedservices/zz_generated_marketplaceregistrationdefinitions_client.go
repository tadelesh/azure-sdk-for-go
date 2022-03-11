//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

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

// MarketplaceRegistrationDefinitionsClient contains the methods for the MarketplaceRegistrationDefinitions group.
// Don't use this type directly, use NewMarketplaceRegistrationDefinitionsClient() instead.
type MarketplaceRegistrationDefinitionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewMarketplaceRegistrationDefinitionsClient creates a new instance of MarketplaceRegistrationDefinitionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMarketplaceRegistrationDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *MarketplaceRegistrationDefinitionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &MarketplaceRegistrationDefinitionsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get the marketplace registration definition for the marketplace identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the resource.
// marketplaceIdentifier - The Azure Marketplace identifier. Expected formats: {publisher}.{product[-preview]}.{planName}.{version}
// or {publisher}.{product[-preview]}.{planName} or {publisher}.{product[-preview]} or
// {publisher}).
// options - MarketplaceRegistrationDefinitionsClientGetOptions contains the optional parameters for the MarketplaceRegistrationDefinitionsClient.Get
// method.
func (client *MarketplaceRegistrationDefinitionsClient) Get(ctx context.Context, scope string, marketplaceIdentifier string, options *MarketplaceRegistrationDefinitionsClientGetOptions) (MarketplaceRegistrationDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, marketplaceIdentifier, options)
	if err != nil {
		return MarketplaceRegistrationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceRegistrationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceRegistrationDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MarketplaceRegistrationDefinitionsClient) getCreateRequest(ctx context.Context, scope string, marketplaceIdentifier string, options *MarketplaceRegistrationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{marketplaceIdentifier}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if marketplaceIdentifier == "" {
		return nil, errors.New("parameter marketplaceIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceIdentifier}", url.PathEscape(marketplaceIdentifier))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MarketplaceRegistrationDefinitionsClient) getHandleResponse(resp *http.Response) (MarketplaceRegistrationDefinitionsClientGetResponse, error) {
	result := MarketplaceRegistrationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceRegistrationDefinition); err != nil {
		return MarketplaceRegistrationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of the marketplace registration definitions for the marketplace identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the resource.
// options - MarketplaceRegistrationDefinitionsClientListOptions contains the optional parameters for the MarketplaceRegistrationDefinitionsClient.List
// method.
func (client *MarketplaceRegistrationDefinitionsClient) List(scope string, options *MarketplaceRegistrationDefinitionsClientListOptions) *MarketplaceRegistrationDefinitionsClientListPager {
	return &MarketplaceRegistrationDefinitionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp MarketplaceRegistrationDefinitionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MarketplaceRegistrationDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MarketplaceRegistrationDefinitionsClient) listCreateRequest(ctx context.Context, scope string, options *MarketplaceRegistrationDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceRegistrationDefinitionsClient) listHandleResponse(resp *http.Response) (MarketplaceRegistrationDefinitionsClientListResponse, error) {
	result := MarketplaceRegistrationDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceRegistrationDefinitionList); err != nil {
		return MarketplaceRegistrationDefinitionsClientListResponse{}, err
	}
	return result, nil
}
