//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerlockbox

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

// RequestsClient contains the methods for the Requests group.
// Don't use this type directly, use NewRequestsClient() instead.
type RequestsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRequestsClient creates a new instance of RequestsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RequestsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RequestsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get Customer Lockbox request
// If the operation fails it returns an *azcore.ResponseError type.
// requestID - The Lockbox request ID.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// options - RequestsClientGetOptions contains the optional parameters for the RequestsClient.Get method.
func (client *RequestsClient) Get(ctx context.Context, requestID string, subscriptionID string, options *RequestsClientGetOptions) (RequestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, requestID, subscriptionID, options)
	if err != nil {
		return RequestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RequestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RequestsClient) getCreateRequest(ctx context.Context, requestID string, subscriptionID string, options *RequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests/{requestId}"
	if requestID == "" {
		return nil, errors.New("parameter requestID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{requestId}", url.PathEscape(requestID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-28-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RequestsClient) getHandleResponse(resp *http.Response) (RequestsClientGetResponse, error) {
	result := RequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LockboxRequestResponse); err != nil {
		return RequestsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all of the Lockbox requests in the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// options - RequestsClientListOptions contains the optional parameters for the RequestsClient.List method.
func (client *RequestsClient) List(subscriptionID string, options *RequestsClientListOptions) *RequestsClientListPager {
	return &RequestsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, subscriptionID, options)
		},
		advancer: func(ctx context.Context, resp RequestsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RequestListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RequestsClient) listCreateRequest(ctx context.Context, subscriptionID string, options *RequestsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RequestsClient) listHandleResponse(resp *http.Response) (RequestsClientListResponse, error) {
	result := RequestsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestListResult); err != nil {
		return RequestsClientListResponse{}, err
	}
	return result, nil
}

// UpdateStatus - Update Customer Lockbox request approval status API
// If the operation fails it returns an *azcore.ResponseError type.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// requestID - The Lockbox request ID.
// approval - The approval object to update request status.
// options - RequestsClientUpdateStatusOptions contains the optional parameters for the RequestsClient.UpdateStatus method.
func (client *RequestsClient) UpdateStatus(ctx context.Context, subscriptionID string, requestID string, approval Approval, options *RequestsClientUpdateStatusOptions) (RequestsClientUpdateStatusResponse, error) {
	req, err := client.updateStatusCreateRequest(ctx, subscriptionID, requestID, approval, options)
	if err != nil {
		return RequestsClientUpdateStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RequestsClientUpdateStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RequestsClientUpdateStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateStatusHandleResponse(resp)
}

// updateStatusCreateRequest creates the UpdateStatus request.
func (client *RequestsClient) updateStatusCreateRequest(ctx context.Context, subscriptionID string, requestID string, approval Approval, options *RequestsClientUpdateStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests/{requestId}/updateApproval"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if requestID == "" {
		return nil, errors.New("parameter requestID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{requestId}", url.PathEscape(requestID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-28-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, approval)
}

// updateStatusHandleResponse handles the UpdateStatus response.
func (client *RequestsClient) updateStatusHandleResponse(resp *http.Response) (RequestsClientUpdateStatusResponse, error) {
	result := RequestsClientUpdateStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Approval); err != nil {
		return RequestsClientUpdateStatusResponse{}, err
	}
	return result, nil
}
