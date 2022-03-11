//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
// Manager API or the portal.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkResourcesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListSupported - Gets a list of all supported private link resource types for the given service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// options - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get method.
func (client *PrivateLinkResourcesClient) ListSupported(resourceGroupName string, searchServiceName string, options *SearchManagementRequestOptions) *PrivateLinkResourcesClientListSupportedPager {
	return &PrivateLinkResourcesClientListSupportedPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSupportedCreateRequest(ctx, resourceGroupName, searchServiceName, options)
		},
	}
}

// listSupportedCreateRequest creates the ListSupported request.
func (client *PrivateLinkResourcesClient) listSupportedCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, options *SearchManagementRequestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSupportedHandleResponse handles the ListSupported response.
func (client *PrivateLinkResourcesClient) listSupportedHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListSupportedResponse, error) {
	result := PrivateLinkResourcesClientListSupportedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourcesResult); err != nil {
		return PrivateLinkResourcesClientListSupportedResponse{}, err
	}
	return result, nil
}
