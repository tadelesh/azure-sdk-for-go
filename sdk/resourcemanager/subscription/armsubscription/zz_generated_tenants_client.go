//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// TenantsClient contains the methods for the Tenants group.
// Don't use this type directly, use NewTenantsClient() instead.
type TenantsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewTenantsClient creates a new instance of TenantsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTenantsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TenantsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TenantsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Gets the tenants for your account.
// If the operation fails it returns an *azcore.ResponseError type.
// options - TenantsClientListOptions contains the optional parameters for the TenantsClient.List method.
func (client *TenantsClient) List(options *TenantsClientListOptions) *runtime.Pager[TenantsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TenantsClientListResponse]{
		More: func(page TenantsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantsClientListResponse) (TenantsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TenantsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TenantsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TenantsClient) listCreateRequest(ctx context.Context, options *TenantsClientListOptions) (*policy.Request, error) {
	urlPath := "/tenants"
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
func (client *TenantsClient) listHandleResponse(resp *http.Response) (TenantsClientListResponse, error) {
	result := TenantsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantListResult); err != nil {
		return TenantsClientListResponse{}, err
	}
	return result, nil
}
