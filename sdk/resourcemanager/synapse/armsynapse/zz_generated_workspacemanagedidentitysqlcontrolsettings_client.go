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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// WorkspaceManagedIdentitySQLControlSettingsClient contains the methods for the WorkspaceManagedIdentitySQLControlSettings group.
// Don't use this type directly, use NewWorkspaceManagedIdentitySQLControlSettingsClient() instead.
type WorkspaceManagedIdentitySQLControlSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspaceManagedIdentitySQLControlSettingsClient creates a new instance of WorkspaceManagedIdentitySQLControlSettingsClient with the specified values.
func NewWorkspaceManagedIdentitySQLControlSettingsClient(con *arm.Connection, subscriptionID string) *WorkspaceManagedIdentitySQLControlSettingsClient {
	return &WorkspaceManagedIdentitySQLControlSettingsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update Managed Identity Sql Control Settings
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsBeginCreateOrUpdateOptions) (WorkspaceManagedIdentitySQLControlSettingsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, managedIdentitySQLControlSettings, options)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceManagedIdentitySQLControlSettingsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedIdentitySQLControlSettingsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedIdentitySQLControlSettingsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update Managed Identity Sql Control Settings
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, managedIdentitySQLControlSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, managedIdentitySQLControlSettings)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Get - Get Managed Identity Sql Control Settings
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedIdentitySQLControlSettingsGetOptions) (WorkspaceManagedIdentitySQLControlSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedIdentitySQLControlSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedIdentitySQLControlSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) getHandleResponse(resp *http.Response) (WorkspaceManagedIdentitySQLControlSettingsGetResponse, error) {
	result := WorkspaceManagedIdentitySQLControlSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedIdentitySQLControlSettingsModel); err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) getHandleError(resp *http.Response) error {
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
