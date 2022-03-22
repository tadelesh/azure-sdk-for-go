//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// CheckNameAvailabilityClient contains the methods for the CheckNameAvailability group.
// Don't use this type directly, use NewCheckNameAvailabilityClient() instead.
type CheckNameAvailabilityClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCheckNameAvailabilityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CheckNameAvailabilityClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CheckNameAvailabilityClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Execute - Check the availability of name for resource
// If the operation fails it returns an *azcore.ResponseError type.
// nameAvailabilityRequest - The required parameters for checking if resource name is available.
// options - CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
// method.
func (client *CheckNameAvailabilityClient) Execute(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityClientExecuteOptions) (CheckNameAvailabilityClientExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CheckNameAvailabilityClientExecuteResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *CheckNameAvailabilityClient) executeCreateRequest(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, nameAvailabilityRequest)
}

// executeHandleResponse handles the Execute response.
func (client *CheckNameAvailabilityClient) executeHandleResponse(resp *http.Response) (CheckNameAvailabilityClientExecuteResponse, error) {
	result := CheckNameAvailabilityClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	return result, nil
}
