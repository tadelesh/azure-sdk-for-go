//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PeeringServiceProvidersClient contains the methods for the PeeringServiceProviders group.
// Don't use this type directly, use NewPeeringServiceProvidersClient() instead.
type PeeringServiceProvidersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPeeringServiceProvidersClient creates a new instance of PeeringServiceProvidersClient with the specified values.
func NewPeeringServiceProvidersClient(con *arm.Connection, subscriptionID string) *PeeringServiceProvidersClient {
	return &PeeringServiceProvidersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Lists all of the available peering service locations for the specified kind of peering.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PeeringServiceProvidersClient) List(options *PeeringServiceProvidersListOptions) *PeeringServiceProvidersListPager {
	return &PeeringServiceProvidersListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PeeringServiceProvidersListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PeeringServiceProviderListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PeeringServiceProvidersClient) listCreateRequest(ctx context.Context, options *PeeringServiceProvidersListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PeeringServiceProvidersClient) listHandleResponse(resp *http.Response) (PeeringServiceProvidersListResponse, error) {
	result := PeeringServiceProvidersListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeeringServiceProviderListResult); err != nil {
		return PeeringServiceProvidersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PeeringServiceProvidersClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
