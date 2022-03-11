//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// ServiceRegistriesClient contains the methods for the ServiceRegistries group.
// Don't use this type directly, use NewServiceRegistriesClient() instead.
type ServiceRegistriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceRegistriesClient creates a new instance of ServiceRegistriesClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceRegistriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceRegistriesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServiceRegistriesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create the default Service Registry or update the existing Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// serviceRegistryName - The name of Service Registry.
// options - ServiceRegistriesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceRegistriesClient.BeginCreateOrUpdate
// method.
func (client *ServiceRegistriesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (ServiceRegistriesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return ServiceRegistriesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ServiceRegistriesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ServiceRegistriesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServiceRegistriesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServiceRegistriesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create the default Service Registry or update the existing Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServiceRegistriesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceRegistriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDelete - Disable the default Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// serviceRegistryName - The name of Service Registry.
// options - ServiceRegistriesClientBeginDeleteOptions contains the optional parameters for the ServiceRegistriesClient.BeginDelete
// method.
func (client *ServiceRegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (ServiceRegistriesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return ServiceRegistriesClientDeletePollerResponse{}, err
	}
	result := ServiceRegistriesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ServiceRegistriesClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServiceRegistriesClientDeletePollerResponse{}, err
	}
	result.Poller = &ServiceRegistriesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Disable the default Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServiceRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
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
func (client *ServiceRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the Service Registry and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// serviceRegistryName - The name of Service Registry.
// options - ServiceRegistriesClientGetOptions contains the optional parameters for the ServiceRegistriesClient.Get method.
func (client *ServiceRegistriesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientGetOptions) (ServiceRegistriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceRegistriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceRegistriesClient) getHandleResponse(resp *http.Response) (ServiceRegistriesClientGetResponse, error) {
	result := ServiceRegistriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRegistryResource); err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	return result, nil
}

// List - Handles requests to list all resources in a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServiceRegistriesClientListOptions contains the optional parameters for the ServiceRegistriesClient.List method.
func (client *ServiceRegistriesClient) List(resourceGroupName string, serviceName string, options *ServiceRegistriesClientListOptions) *ServiceRegistriesClientListPager {
	return &ServiceRegistriesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp ServiceRegistriesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceRegistryResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceRegistriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceRegistriesClient) listHandleResponse(resp *http.Response) (ServiceRegistriesClientListResponse, error) {
	result := ServiceRegistriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRegistryResourceCollection); err != nil {
		return ServiceRegistriesClientListResponse{}, err
	}
	return result, nil
}
