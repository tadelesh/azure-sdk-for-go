//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// GatewayClient contains the methods for the Gateway group.
// Don't use this type directly, use NewGatewayClient() instead.
type GatewayClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGatewayClient creates a new instance of GatewayClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGatewayClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GatewayClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &GatewayClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Creates a gateway resource with the specified name, description and properties. If a gateway resource with the
// same name exists, then it is updated with the specified description and properties. Use
// gateway resources to create a gateway for public connectivity for services within your application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// gatewayResourceName - The identity of the gateway.
// gatewayResourceDescription - Description for creating a Gateway resource.
// options - GatewayClientCreateOptions contains the optional parameters for the GatewayClient.Create method.
func (client *GatewayClient) Create(ctx context.Context, resourceGroupName string, gatewayResourceName string, gatewayResourceDescription GatewayResourceDescription, options *GatewayClientCreateOptions) (GatewayClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, gatewayResourceName, gatewayResourceDescription, options)
	if err != nil {
		return GatewayClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return GatewayClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *GatewayClient) createCreateRequest(ctx context.Context, resourceGroupName string, gatewayResourceName string, gatewayResourceDescription GatewayResourceDescription, options *GatewayClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayResourceName}", gatewayResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, gatewayResourceDescription)
}

// createHandleResponse handles the Create response.
func (client *GatewayClient) createHandleResponse(resp *http.Response) (GatewayClientCreateResponse, error) {
	result := GatewayClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceDescription); err != nil {
		return GatewayClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the gateway resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// gatewayResourceName - The identity of the gateway.
// options - GatewayClientDeleteOptions contains the optional parameters for the GatewayClient.Delete method.
func (client *GatewayClient) Delete(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *GatewayClientDeleteOptions) (GatewayClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayResourceName, options)
	if err != nil {
		return GatewayClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return GatewayClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return GatewayClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *GatewayClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayResourceName}", gatewayResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the information about the gateway resource with the given name. The information include the description and
// other properties of the gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// gatewayResourceName - The identity of the gateway.
// options - GatewayClientGetOptions contains the optional parameters for the GatewayClient.Get method.
func (client *GatewayClient) Get(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *GatewayClientGetOptions) (GatewayClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayResourceName, options)
	if err != nil {
		return GatewayClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *GatewayClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayResourceName}", gatewayResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewayClient) getHandleResponse(resp *http.Response) (GatewayClientGetResponse, error) {
	result := GatewayClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceDescription); err != nil {
		return GatewayClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets the information about all gateway resources in a given resource group. The information include
// the description and other properties of the Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// options - GatewayClientListByResourceGroupOptions contains the optional parameters for the GatewayClient.ListByResourceGroup
// method.
func (client *GatewayClient) ListByResourceGroup(resourceGroupName string, options *GatewayClientListByResourceGroupOptions) *GatewayClientListByResourceGroupPager {
	return &GatewayClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp GatewayClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GatewayResourceDescriptionList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GatewayClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GatewayClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GatewayClient) listByResourceGroupHandleResponse(resp *http.Response) (GatewayClientListByResourceGroupResponse, error) {
	result := GatewayClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceDescriptionList); err != nil {
		return GatewayClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets the information about all gateway resources in a given resource group. The information include
// the description and other properties of the gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// options - GatewayClientListBySubscriptionOptions contains the optional parameters for the GatewayClient.ListBySubscription
// method.
func (client *GatewayClient) ListBySubscription(options *GatewayClientListBySubscriptionOptions) *GatewayClientListBySubscriptionPager {
	return &GatewayClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp GatewayClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GatewayResourceDescriptionList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *GatewayClient) listBySubscriptionCreateRequest(ctx context.Context, options *GatewayClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/gateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *GatewayClient) listBySubscriptionHandleResponse(resp *http.Response) (GatewayClientListBySubscriptionResponse, error) {
	result := GatewayClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceDescriptionList); err != nil {
		return GatewayClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
