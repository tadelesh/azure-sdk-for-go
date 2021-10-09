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

// IntegrationRuntimeStatusClient contains the methods for the IntegrationRuntimeStatus group.
// Don't use this type directly, use NewIntegrationRuntimeStatusClient() instead.
type IntegrationRuntimeStatusClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationRuntimeStatusClient creates a new instance of IntegrationRuntimeStatusClient with the specified values.
func NewIntegrationRuntimeStatusClient(con *arm.Connection, subscriptionID string) *IntegrationRuntimeStatusClient {
	return &IntegrationRuntimeStatusClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get the integration runtime status
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationRuntimeStatusClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeStatusGetOptions) (IntegrationRuntimeStatusGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeStatusGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeStatusGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeStatusGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeStatusGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/getStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
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
func (client *IntegrationRuntimeStatusClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeStatusGetResponse, error) {
	result := IntegrationRuntimeStatusGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeStatusResponse); err != nil {
		return IntegrationRuntimeStatusGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationRuntimeStatusClient) getHandleError(resp *http.Response) error {
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
