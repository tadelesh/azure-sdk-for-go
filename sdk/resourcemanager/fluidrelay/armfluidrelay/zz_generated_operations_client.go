//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armfluidrelay

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsClient contains the methods for the FluidRelayOperations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *OperationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - List all operations provided by Microsoft.FluidRelay.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
func (client *OperationsClient) List(options *OperationsClientListOptions) *OperationsClientListPager {
	return &OperationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp OperationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OperationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.FluidRelay/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}
