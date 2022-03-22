//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// MetricNamespacesClient contains the methods for the MetricNamespaces group.
// Don't use this type directly, use NewMetricNamespacesClient() instead.
type MetricNamespacesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewMetricNamespacesClient creates a new instance of MetricNamespacesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMetricNamespacesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *MetricNamespacesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &MetricNamespacesClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists the metric namespaces for the resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// options - MetricNamespacesClientListOptions contains the optional parameters for the MetricNamespacesClient.List method.
func (client *MetricNamespacesClient) List(resourceURI string, options *MetricNamespacesClientListOptions) *runtime.Pager[MetricNamespacesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[MetricNamespacesClientListResponse]{
		More: func(page MetricNamespacesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricNamespacesClientListResponse) (MetricNamespacesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceURI, options)
			if err != nil {
				return MetricNamespacesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MetricNamespacesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricNamespacesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MetricNamespacesClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricNamespacesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/metricNamespaces"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01-preview")
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricNamespacesClient) listHandleResponse(resp *http.Response) (MetricNamespacesClientListResponse, error) {
	result := MetricNamespacesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricNamespaceCollection); err != nil {
		return MetricNamespacesClientListResponse{}, err
	}
	return result, nil
}
