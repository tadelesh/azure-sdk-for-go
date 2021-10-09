//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SparkConfigurationClient contains the methods for the SparkConfiguration group.
// Don't use this type directly, use NewSparkConfigurationClient() instead.
type SparkConfigurationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSparkConfigurationClient creates a new instance of SparkConfigurationClient with the specified values.
func NewSparkConfigurationClient(con *arm.Connection, subscriptionID string) *SparkConfigurationClient {
	return &SparkConfigurationClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get SparkConfiguration by name in a workspace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SparkConfigurationClient) Get(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string, options *SparkConfigurationGetOptions) (SparkConfigurationGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sparkConfigurationName, workspaceName, options)
	if err != nil {
		return SparkConfigurationGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SparkConfigurationGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkConfigurationGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SparkConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string, options *SparkConfigurationGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sparkconfigurations/{sparkConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SparkConfigurationClient) getHandleResponse(resp *http.Response) (SparkConfigurationGetResponse, error) {
	result := SparkConfigurationGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkConfigurationResource); err != nil {
		return SparkConfigurationGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SparkConfigurationClient) getHandleError(resp *http.Response) error {
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
