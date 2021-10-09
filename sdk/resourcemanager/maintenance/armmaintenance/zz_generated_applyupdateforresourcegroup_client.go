//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// ApplyUpdateForResourceGroupClient contains the methods for the ApplyUpdateForResourceGroup group.
// Don't use this type directly, use NewApplyUpdateForResourceGroupClient() instead.
type ApplyUpdateForResourceGroupClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewApplyUpdateForResourceGroupClient creates a new instance of ApplyUpdateForResourceGroupClient with the specified values.
func NewApplyUpdateForResourceGroupClient(con *arm.Connection, subscriptionID string) *ApplyUpdateForResourceGroupClient {
	return &ApplyUpdateForResourceGroupClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Get Configuration records within a subscription and resource group
// If the operation fails it returns the *MaintenanceError error type.
func (client *ApplyUpdateForResourceGroupClient) List(ctx context.Context, resourceGroupName string, options *ApplyUpdateForResourceGroupListOptions) (ApplyUpdateForResourceGroupListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ApplyUpdateForResourceGroupListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplyUpdateForResourceGroupListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplyUpdateForResourceGroupListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ApplyUpdateForResourceGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplyUpdateForResourceGroupListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maintenance/applyUpdates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplyUpdateForResourceGroupClient) listHandleResponse(resp *http.Response) (ApplyUpdateForResourceGroupListResponse, error) {
	result := ApplyUpdateForResourceGroupListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListApplyUpdate); err != nil {
		return ApplyUpdateForResourceGroupListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ApplyUpdateForResourceGroupClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := MaintenanceError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
