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

// BaselinesClient contains the methods for the Baselines group.
// Don't use this type directly, use NewBaselinesClient() instead.
type BaselinesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBaselinesClient creates a new instance of BaselinesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBaselinesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *BaselinesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BaselinesClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists the metric baseline values for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// options - BaselinesClientListOptions contains the optional parameters for the BaselinesClient.List method.
func (client *BaselinesClient) List(resourceURI string, options *BaselinesClientListOptions) *runtime.Pager[BaselinesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[BaselinesClientListResponse]{
		More: func(page BaselinesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BaselinesClientListResponse) (BaselinesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceURI, options)
			if err != nil {
				return BaselinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BaselinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BaselinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BaselinesClient) listCreateRequest(ctx context.Context, resourceURI string, options *BaselinesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricBaselines"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Sensitivities != nil {
		reqQP.Set("sensitivities", *options.Sensitivities)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2019-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BaselinesClient) listHandleResponse(resp *http.Response) (BaselinesClientListResponse, error) {
	result := BaselinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricBaselinesResponse); err != nil {
		return BaselinesClientListResponse{}, err
	}
	return result, nil
}
