//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualNetworksClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Return virtual network by its name
// If the operation fails it returns an *azcore.ResponseError type.
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// virtualNetworkName - virtual network id (vsphereId)
// options - VirtualNetworksClientGetOptions contains the optional parameters for the VirtualNetworksClient.Get method.
func (client *VirtualNetworksClient) Get(ctx context.Context, regionID string, pcName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (VirtualNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, regionID string, pcName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksClientGetResponse, error) {
	result := VirtualNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// List - Return list of virtual networks in location for private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// resourcePoolName - Resource pool used to derive vSphere cluster which contains virtual networks
// options - VirtualNetworksClientListOptions contains the optional parameters for the VirtualNetworksClient.List method.
func (client *VirtualNetworksClient) List(regionID string, pcName string, resourcePoolName string, options *VirtualNetworksClientListOptions) *VirtualNetworksClientListPager {
	return &VirtualNetworksClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, regionID, pcName, resourcePoolName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworksClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualNetworksClient) listCreateRequest(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *VirtualNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	reqQP.Set("resourcePoolName", resourcePoolName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworksClient) listHandleResponse(resp *http.Response) (VirtualNetworksClientListResponse, error) {
	result := VirtualNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResponse); err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	return result, nil
}
