//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// KustoOperationsClient contains the methods for the KustoOperations group.
// Don't use this type directly, use NewKustoOperationsClient() instead.
type KustoOperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewKustoOperationsClient creates a new instance of KustoOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewKustoOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *KustoOperationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &KustoOperationsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists available operations for the Kusto sub-resources inside Microsoft.Synapse provider.
// If the operation fails it returns an *azcore.ResponseError type.
// options - KustoOperationsClientListOptions contains the optional parameters for the KustoOperationsClient.List method.
func (client *KustoOperationsClient) List(options *KustoOperationsClientListOptions) *KustoOperationsClientListPager {
	return &KustoOperationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp KustoOperationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OperationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *KustoOperationsClient) listCreateRequest(ctx context.Context, options *KustoOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Synapse/kustooperations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *KustoOperationsClient) listHandleResponse(resp *http.Response) (KustoOperationsClientListResponse, error) {
	result := KustoOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return KustoOperationsClientListResponse{}, err
	}
	return result, nil
}
