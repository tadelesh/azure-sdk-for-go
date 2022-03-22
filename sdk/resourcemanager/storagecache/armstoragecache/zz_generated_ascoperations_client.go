//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

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

// AscOperationsClient contains the methods for the AscOperations group.
// Don't use this type directly, use NewAscOperationsClient() instead.
type AscOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAscOperationsClient creates a new instance of AscOperationsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAscOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AscOperationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AscOperationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the status of an asynchronous operation for the Azure HPC Cache
// If the operation fails it returns an *azcore.ResponseError type.
// location - The name of the region used to look up the operation.
// operationID - The operation id which uniquely identifies the asynchronous operation.
// options - AscOperationsClientGetOptions contains the optional parameters for the AscOperationsClient.Get method.
func (client *AscOperationsClient) Get(ctx context.Context, location string, operationID string, options *AscOperationsClientGetOptions) (AscOperationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return AscOperationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AscOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AscOperationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AscOperationsClient) getCreateRequest(ctx context.Context, location string, operationID string, options *AscOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/locations/{location}/ascOperations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AscOperationsClient) getHandleResponse(resp *http.Response) (AscOperationsClientGetResponse, error) {
	result := AscOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AscOperation); err != nil {
		return AscOperationsClientGetResponse{}, err
	}
	return result, nil
}
