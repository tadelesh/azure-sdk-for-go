//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// ExtensionTypeVersionsClient contains the methods for the ExtensionTypeVersions group.
// Don't use this type directly, use NewExtensionTypeVersionsClient() instead.
type ExtensionTypeVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionTypeVersionsClient creates a new instance of ExtensionTypeVersionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtensionTypeVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExtensionTypeVersionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExtensionTypeVersionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - List available versions for an Extension Type
// If the operation fails it returns an *azcore.ResponseError type.
// location - extension location
// extensionTypeName - Extension type name
// options - ExtensionTypeVersionsClientListOptions contains the optional parameters for the ExtensionTypeVersionsClient.List
// method.
func (client *ExtensionTypeVersionsClient) List(location string, extensionTypeName string, options *ExtensionTypeVersionsClientListOptions) *ExtensionTypeVersionsClientListPager {
	return &ExtensionTypeVersionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, extensionTypeName, options)
		},
		advancer: func(ctx context.Context, resp ExtensionTypeVersionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtensionVersionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExtensionTypeVersionsClient) listCreateRequest(ctx context.Context, location string, extensionTypeName string, options *ExtensionTypeVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KubernetesConfiguration/locations/{location}/extensionTypes/{extensionTypeName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if extensionTypeName == "" {
		return nil, errors.New("parameter extensionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionTypeName}", url.PathEscape(extensionTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExtensionTypeVersionsClient) listHandleResponse(resp *http.Response) (ExtensionTypeVersionsClientListResponse, error) {
	result := ExtensionTypeVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionVersionList); err != nil {
		return ExtensionTypeVersionsClientListResponse{}, err
	}
	return result, nil
}
