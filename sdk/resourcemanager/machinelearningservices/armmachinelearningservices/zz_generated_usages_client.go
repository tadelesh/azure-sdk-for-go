//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

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

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UsagesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Gets the current usage information as well as limits for AML resources for given subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location for which resource usage is queried.
// options - UsagesClientListOptions contains the optional parameters for the UsagesClient.List method.
func (client *UsagesClient) List(location string, options *UsagesClientListOptions) *UsagesClientListPager {
	return &UsagesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp UsagesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListUsagesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UsagesClient) listCreateRequest(ctx context.Context, location string, options *UsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsagesClient) listHandleResponse(resp *http.Response) (UsagesClientListResponse, error) {
	result := UsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUsagesResult); err != nil {
		return UsagesClientListResponse{}, err
	}
	return result, nil
}
