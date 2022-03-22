//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationClient contains the methods for the QuotaOperation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationClient(credential azcore.TokenCredential, options *arm.ClientOptions) *OperationClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - List all the operations supported by the Microsoft.Quota resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OperationClientListOptions contains the optional parameters for the OperationClient.List method.
func (client *OperationClient) List(options *OperationClientListOptions) *runtime.Pager[OperationClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[OperationClientListResponse]{
		More: func(page OperationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationClientListResponse) (OperationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OperationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationClient) listCreateRequest(ctx context.Context, options *OperationClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Quota/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationClient) listHandleResponse(resp *http.Response) (OperationClientListResponse, error) {
	result := OperationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationList); err != nil {
		return OperationClientListResponse{}, err
	}
	return result, nil
}
