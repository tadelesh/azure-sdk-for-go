//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// OperationStatusesClient contains the methods for the OperationStatuses group.
// Don't use this type directly, use NewOperationStatusesClient() instead.
type OperationStatusesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationStatusesClient creates a new instance of OperationStatusesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationStatusesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationStatusesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get the status of a long running azure asynchronous operation.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The region name of operation.
// asyncOperationID - The operation Id.
// options - OperationStatusesClientGetOptions contains the optional parameters for the OperationStatusesClient.Get method.
func (client *OperationStatusesClient) Get(ctx context.Context, location string, asyncOperationID string, options *OperationStatusesClientGetOptions) (OperationStatusesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, asyncOperationID, options)
	if err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationStatusesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationStatusesClient) getCreateRequest(ctx context.Context, location string, asyncOperationID string, options *OperationStatusesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/locations/{location}/operationStatuses/{asyncOperationId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if asyncOperationID == "" {
		return nil, errors.New("parameter asyncOperationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{asyncOperationId}", url.PathEscape(asyncOperationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationStatusesClient) getHandleResponse(resp *http.Response) (OperationStatusesClientGetResponse, error) {
	result := OperationStatusesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	return result, nil
}
