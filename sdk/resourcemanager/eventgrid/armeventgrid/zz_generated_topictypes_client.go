//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// TopicTypesClient contains the methods for the TopicTypes group.
// Don't use this type directly, use NewTopicTypesClient() instead.
type TopicTypesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewTopicTypesClient creates a new instance of TopicTypesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTopicTypesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TopicTypesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TopicTypesClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get information about a topic type.
// If the operation fails it returns an *azcore.ResponseError type.
// topicTypeName - Name of the topic type.
// options - TopicTypesClientGetOptions contains the optional parameters for the TopicTypesClient.Get method.
func (client *TopicTypesClient) Get(ctx context.Context, topicTypeName string, options *TopicTypesClientGetOptions) (TopicTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, topicTypeName, options)
	if err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TopicTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TopicTypesClient) getCreateRequest(ctx context.Context, topicTypeName string, options *TopicTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}"
	if topicTypeName == "" {
		return nil, errors.New("parameter topicTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicTypeName}", url.PathEscape(topicTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicTypesClient) getHandleResponse(resp *http.Response) (TopicTypesClientGetResponse, error) {
	result := TopicTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicTypeInfo); err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	return result, nil
}

// List - List all registered topic types.
// If the operation fails it returns an *azcore.ResponseError type.
// options - TopicTypesClientListOptions contains the optional parameters for the TopicTypesClient.List method.
func (client *TopicTypesClient) List(options *TopicTypesClientListOptions) *runtime.Pager[TopicTypesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TopicTypesClientListResponse]{
		More: func(page TopicTypesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TopicTypesClientListResponse) (TopicTypesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return TopicTypesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TopicTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TopicTypesClient) listCreateRequest(ctx context.Context, options *TopicTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TopicTypesClient) listHandleResponse(resp *http.Response) (TopicTypesClientListResponse, error) {
	result := TopicTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicTypesListResult); err != nil {
		return TopicTypesClientListResponse{}, err
	}
	return result, nil
}

// ListEventTypes - List event types for a topic type.
// If the operation fails it returns an *azcore.ResponseError type.
// topicTypeName - Name of the topic type.
// options - TopicTypesClientListEventTypesOptions contains the optional parameters for the TopicTypesClient.ListEventTypes
// method.
func (client *TopicTypesClient) ListEventTypes(topicTypeName string, options *TopicTypesClientListEventTypesOptions) *runtime.Pager[TopicTypesClientListEventTypesResponse] {
	return runtime.NewPager(runtime.PageProcessor[TopicTypesClientListEventTypesResponse]{
		More: func(page TopicTypesClientListEventTypesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TopicTypesClientListEventTypesResponse) (TopicTypesClientListEventTypesResponse, error) {
			req, err := client.listEventTypesCreateRequest(ctx, topicTypeName, options)
			if err != nil {
				return TopicTypesClientListEventTypesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TopicTypesClientListEventTypesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicTypesClientListEventTypesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listEventTypesHandleResponse(resp)
		},
	})
}

// listEventTypesCreateRequest creates the ListEventTypes request.
func (client *TopicTypesClient) listEventTypesCreateRequest(ctx context.Context, topicTypeName string, options *TopicTypesClientListEventTypesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes"
	if topicTypeName == "" {
		return nil, errors.New("parameter topicTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicTypeName}", url.PathEscape(topicTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listEventTypesHandleResponse handles the ListEventTypes response.
func (client *TopicTypesClient) listEventTypesHandleResponse(resp *http.Response) (TopicTypesClientListEventTypesResponse, error) {
	result := TopicTypesClientListEventTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventTypesListResult); err != nil {
		return TopicTypesClientListEventTypesResponse{}, err
	}
	return result, nil
}
