//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerAutomaticTuningClient contains the methods for the ServerAutomaticTuning group.
// Don't use this type directly, use NewServerAutomaticTuningClient() instead.
type ServerAutomaticTuningClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerAutomaticTuningClient creates a new instance of ServerAutomaticTuningClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerAutomaticTuningClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerAutomaticTuningClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServerAutomaticTuningClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Retrieves server automatic tuning options.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ServerAutomaticTuningClientGetOptions contains the optional parameters for the ServerAutomaticTuningClient.Get
// method.
func (client *ServerAutomaticTuningClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAutomaticTuningClientGetOptions) (ServerAutomaticTuningClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAutomaticTuningClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAutomaticTuningClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAutomaticTuningClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAutomaticTuningClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAutomaticTuningClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/automaticTuning/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAutomaticTuningClient) getHandleResponse(resp *http.Response) (ServerAutomaticTuningClientGetResponse, error) {
	result := ServerAutomaticTuningClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAutomaticTuning); err != nil {
		return ServerAutomaticTuningClientGetResponse{}, err
	}
	return result, nil
}

// Update - Update automatic tuning options on server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// parameters - The requested automatic tuning resource state.
// options - ServerAutomaticTuningClientUpdateOptions contains the optional parameters for the ServerAutomaticTuningClient.Update
// method.
func (client *ServerAutomaticTuningClient) Update(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAutomaticTuning, options *ServerAutomaticTuningClientUpdateOptions) (ServerAutomaticTuningClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return ServerAutomaticTuningClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAutomaticTuningClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAutomaticTuningClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServerAutomaticTuningClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAutomaticTuning, options *ServerAutomaticTuningClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/automaticTuning/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ServerAutomaticTuningClient) updateHandleResponse(resp *http.Response) (ServerAutomaticTuningClientUpdateResponse, error) {
	result := ServerAutomaticTuningClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAutomaticTuning); err != nil {
		return ServerAutomaticTuningClientUpdateResponse{}, err
	}
	return result, nil
}
