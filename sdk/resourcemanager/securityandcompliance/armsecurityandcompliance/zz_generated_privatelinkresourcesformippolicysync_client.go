//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityandcompliance

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

// PrivateLinkResourcesForMIPPolicySyncClient contains the methods for the PrivateLinkResourcesForMIPPolicySync group.
// Don't use this type directly, use NewPrivateLinkResourcesForMIPPolicySyncClient() instead.
type PrivateLinkResourcesForMIPPolicySyncClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesForMIPPolicySyncClient creates a new instance of PrivateLinkResourcesForMIPPolicySyncClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourcesForMIPPolicySyncClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkResourcesForMIPPolicySyncClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateLinkResourcesForMIPPolicySyncClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets a private link resource that need to be created for a service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// groupName - The name of the private link resource group.
// options - PrivateLinkResourcesForMIPPolicySyncClientGetOptions contains the optional parameters for the PrivateLinkResourcesForMIPPolicySyncClient.Get
// method.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) Get(ctx context.Context, resourceGroupName string, resourceName string, groupName string, options *PrivateLinkResourcesForMIPPolicySyncClientGetOptions) (PrivateLinkResourcesForMIPPolicySyncClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, groupName, options)
	if err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesForMIPPolicySyncClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, groupName string, options *PrivateLinkResourcesForMIPPolicySyncClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesForMIPPolicySyncClientGetResponse, error) {
	result := PrivateLinkResourcesForMIPPolicySyncClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientGetResponse{}, err
	}
	return result, nil
}

// ListByService - Gets the private link resources that need to be created for a service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkResourcesForMIPPolicySyncClientListByServiceOptions contains the optional parameters for the PrivateLinkResourcesForMIPPolicySyncClient.ListByService
// method.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) ListByService(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkResourcesForMIPPolicySyncClientListByServiceOptions) (PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkResourcesForMIPPolicySyncClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PrivateLinkResourcesForMIPPolicySyncClient) listByServiceHandleResponse(resp *http.Response) (PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse, error) {
	result := PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesForMIPPolicySyncClientListByServiceResponse{}, err
	}
	return result, nil
}
