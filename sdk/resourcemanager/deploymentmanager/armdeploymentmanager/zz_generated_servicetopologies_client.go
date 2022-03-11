//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

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

// ServiceTopologiesClient contains the methods for the ServiceTopologies group.
// Don't use this type directly, use NewServiceTopologiesClient() instead.
type ServiceTopologiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceTopologiesClient creates a new instance of ServiceTopologiesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceTopologiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceTopologiesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServiceTopologiesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Synchronously creates a new service topology or updates an existing service topology.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// serviceTopologyInfo - Source topology object defines the resource.
// options - ServiceTopologiesClientCreateOrUpdateOptions contains the optional parameters for the ServiceTopologiesClient.CreateOrUpdate
// method.
func (client *ServiceTopologiesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceTopologyInfo ServiceTopologyResource, options *ServiceTopologiesClientCreateOrUpdateOptions) (ServiceTopologiesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceTopologyName, serviceTopologyInfo, options)
	if err != nil {
		return ServiceTopologiesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceTopologiesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ServiceTopologiesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceTopologiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceTopologyInfo ServiceTopologyResource, options *ServiceTopologiesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, serviceTopologyInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceTopologiesClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceTopologiesClientCreateOrUpdateResponse, error) {
	result := ServiceTopologiesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceTopologyResource); err != nil {
		return ServiceTopologiesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the service topology.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// options - ServiceTopologiesClientDeleteOptions contains the optional parameters for the ServiceTopologiesClient.Delete
// method.
func (client *ServiceTopologiesClient) Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServiceTopologiesClientDeleteOptions) (ServiceTopologiesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceTopologyName, options)
	if err != nil {
		return ServiceTopologiesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceTopologiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServiceTopologiesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ServiceTopologiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceTopologiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServiceTopologiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the service topology.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// options - ServiceTopologiesClientGetOptions contains the optional parameters for the ServiceTopologiesClient.Get method.
func (client *ServiceTopologiesClient) Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServiceTopologiesClientGetOptions) (ServiceTopologiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceTopologyName, options)
	if err != nil {
		return ServiceTopologiesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceTopologiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceTopologiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceTopologiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServiceTopologiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceTopologiesClient) getHandleResponse(resp *http.Response) (ServiceTopologiesClientGetResponse, error) {
	result := ServiceTopologiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceTopologyResource); err != nil {
		return ServiceTopologiesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the service topologies in the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ServiceTopologiesClientListOptions contains the optional parameters for the ServiceTopologiesClient.List method.
func (client *ServiceTopologiesClient) List(ctx context.Context, resourceGroupName string, options *ServiceTopologiesClientListOptions) (ServiceTopologiesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ServiceTopologiesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceTopologiesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceTopologiesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServiceTopologiesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ServiceTopologiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceTopologiesClient) listHandleResponse(resp *http.Response) (ServiceTopologiesClientListResponse, error) {
	result := ServiceTopologiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceTopologyResourceArray); err != nil {
		return ServiceTopologiesClientListResponse{}, err
	}
	return result, nil
}
