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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectionMonitorTestsClient contains the methods for the ConnectionMonitorTests group.
// Don't use this type directly, use NewConnectionMonitorTestsClient() instead.
type ConnectionMonitorTestsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectionMonitorTestsClient creates a new instance of ConnectionMonitorTestsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectionMonitorTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConnectionMonitorTestsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ConnectionMonitorTestsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates a connection monitor test with the specified name under the given subscription, resource
// group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// connectionMonitorTestName - The name of the connection monitor test
// connectionMonitorTest - The properties needed to create a connection monitor test
// options - ConnectionMonitorTestsClientCreateOrUpdateOptions contains the optional parameters for the ConnectionMonitorTestsClient.CreateOrUpdate
// method.
func (client *ConnectionMonitorTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, connectionMonitorTest ConnectionMonitorTest, options *ConnectionMonitorTestsClientCreateOrUpdateOptions) (ConnectionMonitorTestsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, connectionMonitorTest, options)
	if err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionMonitorTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, connectionMonitorTest ConnectionMonitorTest, options *ConnectionMonitorTestsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectionMonitorTest)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionMonitorTestsClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientCreateOrUpdateResponse, error) {
	result := ConnectionMonitorTestsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTest); err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing connection monitor test with the specified name under the given subscription, resource group
// and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// connectionMonitorTestName - The name of the connection monitor test
// options - ConnectionMonitorTestsClientDeleteOptions contains the optional parameters for the ConnectionMonitorTestsClient.Delete
// method.
func (client *ConnectionMonitorTestsClient) Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientDeleteOptions) (ConnectionMonitorTestsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, options)
	if err != nil {
		return ConnectionMonitorTestsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectionMonitorTestsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConnectionMonitorTestsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionMonitorTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an existing connection monitor test with the specified name under the given subscription, resource group and
// peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// connectionMonitorTestName - The name of the connection monitor test
// options - ConnectionMonitorTestsClientGetOptions contains the optional parameters for the ConnectionMonitorTestsClient.Get
// method.
func (client *ConnectionMonitorTestsClient) Get(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientGetOptions) (ConnectionMonitorTestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, options)
	if err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionMonitorTestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionMonitorTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *ConnectionMonitorTestsClient) getHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientGetResponse, error) {
	result := ConnectionMonitorTestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTest); err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	return result, nil
}

// ListByPeeringService - Lists all connection monitor tests under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// options - ConnectionMonitorTestsClientListByPeeringServiceOptions contains the optional parameters for the ConnectionMonitorTestsClient.ListByPeeringService
// method.
func (client *ConnectionMonitorTestsClient) ListByPeeringService(resourceGroupName string, peeringServiceName string, options *ConnectionMonitorTestsClientListByPeeringServiceOptions) *ConnectionMonitorTestsClientListByPeeringServicePager {
	return &ConnectionMonitorTestsClientListByPeeringServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPeeringServiceCreateRequest(ctx, resourceGroupName, peeringServiceName, options)
		},
		advancer: func(ctx context.Context, resp ConnectionMonitorTestsClientListByPeeringServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConnectionMonitorTestListResult.NextLink)
		},
	}
}

// listByPeeringServiceCreateRequest creates the ListByPeeringService request.
func (client *ConnectionMonitorTestsClient) listByPeeringServiceCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, options *ConnectionMonitorTestsClientListByPeeringServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPeeringServiceHandleResponse handles the ListByPeeringService response.
func (client *ConnectionMonitorTestsClient) listByPeeringServiceHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientListByPeeringServiceResponse, error) {
	result := ConnectionMonitorTestsClientListByPeeringServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTestListResult); err != nil {
		return ConnectionMonitorTestsClientListByPeeringServiceResponse{}, err
	}
	return result, nil
}
