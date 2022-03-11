//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

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

// ManagedPrivateEndpointsClient contains the methods for the ManagedPrivateEndpoints group.
// Don't use this type directly, use NewManagedPrivateEndpointsClient() instead.
type ManagedPrivateEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedPrivateEndpointsClient creates a new instance of ManagedPrivateEndpointsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedPrivateEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedPrivateEndpointsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedPrivateEndpointsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Checks that the managed private endpoints resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// resourceName - The name of the resource.
// options - ManagedPrivateEndpointsClientCheckNameAvailabilityOptions contains the optional parameters for the ManagedPrivateEndpointsClient.CheckNameAvailability
// method.
func (client *ManagedPrivateEndpointsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName ManagedPrivateEndpointsCheckNameRequest, options *ManagedPrivateEndpointsClientCheckNameAvailabilityOptions) (ManagedPrivateEndpointsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ManagedPrivateEndpointsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName ManagedPrivateEndpointsCheckNameRequest, options *ManagedPrivateEndpointsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpointsCheckNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ManagedPrivateEndpointsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientCheckNameAvailabilityResponse, error) {
	result := ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// managedPrivateEndpointName - The name of the managed private endpoint.
// parameters - The managed private endpoint parameters.
// options - ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginCreateOrUpdate
// method.
func (client *ManagedPrivateEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (ManagedPrivateEndpointsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
	if err != nil {
		return ManagedPrivateEndpointsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedPrivateEndpointsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ManagedPrivateEndpointsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ManagedPrivateEndpointsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedPrivateEndpointsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedPrivateEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedPrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// managedPrivateEndpointName - The name of the managed private endpoint.
// options - ManagedPrivateEndpointsClientBeginDeleteOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginDelete
// method.
func (client *ManagedPrivateEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (ManagedPrivateEndpointsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientDeletePollerResponse{}, err
	}
	result := ManagedPrivateEndpointsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ManagedPrivateEndpointsClient.Delete", "", resp, client.pl)
	if err != nil {
		return ManagedPrivateEndpointsClientDeletePollerResponse{}, err
	}
	result.Poller = &ManagedPrivateEndpointsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedPrivateEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedPrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// managedPrivateEndpointName - The name of the managed private endpoint.
// options - ManagedPrivateEndpointsClientGetOptions contains the optional parameters for the ManagedPrivateEndpointsClient.Get
// method.
func (client *ManagedPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (ManagedPrivateEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedPrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedPrivateEndpointsClient) getHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientGetResponse, error) {
	result := ManagedPrivateEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpoint); err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns the list of managed private endpoints.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// options - ManagedPrivateEndpointsClientListOptions contains the optional parameters for the ManagedPrivateEndpointsClient.List
// method.
func (client *ManagedPrivateEndpointsClient) List(resourceGroupName string, clusterName string, options *ManagedPrivateEndpointsClientListOptions) *ManagedPrivateEndpointsClientListPager {
	return &ManagedPrivateEndpointsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ManagedPrivateEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ManagedPrivateEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedPrivateEndpointsClient) listHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientListResponse, error) {
	result := ManagedPrivateEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointListResult); err != nil {
		return ManagedPrivateEndpointsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// managedPrivateEndpointName - The name of the managed private endpoint.
// parameters - The managed private endpoint parameters.
// options - ManagedPrivateEndpointsClientBeginUpdateOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginUpdate
// method.
func (client *ManagedPrivateEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (ManagedPrivateEndpointsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
	if err != nil {
		return ManagedPrivateEndpointsClientUpdatePollerResponse{}, err
	}
	result := ManagedPrivateEndpointsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ManagedPrivateEndpointsClient.Update", "", resp, client.pl)
	if err != nil {
		return ManagedPrivateEndpointsClientUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedPrivateEndpointsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedPrivateEndpointsClient) update(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedPrivateEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
