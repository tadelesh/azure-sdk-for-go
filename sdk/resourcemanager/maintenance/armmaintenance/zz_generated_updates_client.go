//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// UpdatesClient contains the methods for the Updates group.
// Don't use this type directly, use NewUpdatesClient() instead.
type UpdatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUpdatesClient creates a new instance of UpdatesClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUpdatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UpdatesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &UpdatesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Get updates to resources.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// providerName - Resource provider name
// resourceType - Resource type
// resourceName - Resource identifier
// options - UpdatesClientListOptions contains the optional parameters for the UpdatesClient.List method.
func (client *UpdatesClient) List(resourceGroupName string, providerName string, resourceType string, resourceName string, options *UpdatesClientListOptions) *UpdatesClientListPager {
	return &UpdatesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, options)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UpdatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, options *UpdatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UpdatesClient) listHandleResponse(resp *http.Response) (UpdatesClientListResponse, error) {
	result := UpdatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUpdatesResult); err != nil {
		return UpdatesClientListResponse{}, err
	}
	return result, nil
}

// ListParent - Get updates to resources.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name
// providerName - Resource provider name
// resourceParentType - Resource parent type
// resourceParentName - Resource parent identifier
// resourceType - Resource type
// resourceName - Resource identifier
// options - UpdatesClientListParentOptions contains the optional parameters for the UpdatesClient.ListParent method.
func (client *UpdatesClient) ListParent(resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *UpdatesClientListParentOptions) *UpdatesClientListParentPager {
	return &UpdatesClientListParentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, options)
		},
	}
}

// listParentCreateRequest creates the ListParent request.
func (client *UpdatesClient) listParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *UpdatesClientListParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listParentHandleResponse handles the ListParent response.
func (client *UpdatesClient) listParentHandleResponse(resp *http.Response) (UpdatesClientListParentResponse, error) {
	result := UpdatesClientListParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUpdatesResult); err != nil {
		return UpdatesClientListParentResponse{}, err
	}
	return result, nil
}
