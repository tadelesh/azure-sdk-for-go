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

// RegisteredPrefixesClient contains the methods for the RegisteredPrefixes group.
// Don't use this type directly, use NewRegisteredPrefixesClient() instead.
type RegisteredPrefixesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegisteredPrefixesClient creates a new instance of RegisteredPrefixesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegisteredPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegisteredPrefixesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RegisteredPrefixesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates a new registered prefix with the specified name under the given subscription, resource group and
// peering.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// registeredPrefixName - The name of the registered prefix.
// registeredPrefix - The properties needed to create a registered prefix.
// options - RegisteredPrefixesClientCreateOrUpdateOptions contains the optional parameters for the RegisteredPrefixesClient.CreateOrUpdate
// method.
func (client *RegisteredPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix, options *RegisteredPrefixesClientCreateOrUpdateOptions) (RegisteredPrefixesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, registeredPrefix, options)
	if err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegisteredPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix, options *RegisteredPrefixesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
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
	return req, runtime.MarshalAsJSON(req, registeredPrefix)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RegisteredPrefixesClient) createOrUpdateHandleResponse(resp *http.Response) (RegisteredPrefixesClientCreateOrUpdateResponse, error) {
	result := RegisteredPrefixesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefix); err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing registered prefix with the specified name under the given subscription, resource group and
// peering.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// registeredPrefixName - The name of the registered prefix.
// options - RegisteredPrefixesClientDeleteOptions contains the optional parameters for the RegisteredPrefixesClient.Delete
// method.
func (client *RegisteredPrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientDeleteOptions) (RegisteredPrefixesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, options)
	if err != nil {
		return RegisteredPrefixesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegisteredPrefixesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RegisteredPrefixesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RegisteredPrefixesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegisteredPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
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

// Get - Gets an existing registered prefix with the specified name under the given subscription, resource group and peering.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// registeredPrefixName - The name of the registered prefix.
// options - RegisteredPrefixesClientGetOptions contains the optional parameters for the RegisteredPrefixesClient.Get method.
func (client *RegisteredPrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientGetOptions) (RegisteredPrefixesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, options)
	if err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegisteredPrefixesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegisteredPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
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
func (client *RegisteredPrefixesClient) getHandleResponse(resp *http.Response) (RegisteredPrefixesClientGetResponse, error) {
	result := RegisteredPrefixesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefix); err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	return result, nil
}

// ListByPeering - Lists all registered prefixes under the given subscription, resource group and peering.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// options - RegisteredPrefixesClientListByPeeringOptions contains the optional parameters for the RegisteredPrefixesClient.ListByPeering
// method.
func (client *RegisteredPrefixesClient) ListByPeering(resourceGroupName string, peeringName string, options *RegisteredPrefixesClientListByPeeringOptions) *RegisteredPrefixesClientListByPeeringPager {
	return &RegisteredPrefixesClientListByPeeringPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPeeringCreateRequest(ctx, resourceGroupName, peeringName, options)
		},
		advancer: func(ctx context.Context, resp RegisteredPrefixesClientListByPeeringResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RegisteredPrefixListResult.NextLink)
		},
	}
}

// listByPeeringCreateRequest creates the ListByPeering request.
func (client *RegisteredPrefixesClient) listByPeeringCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *RegisteredPrefixesClientListByPeeringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// listByPeeringHandleResponse handles the ListByPeering response.
func (client *RegisteredPrefixesClient) listByPeeringHandleResponse(resp *http.Response) (RegisteredPrefixesClientListByPeeringResponse, error) {
	result := RegisteredPrefixesClientListByPeeringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefixListResult); err != nil {
		return RegisteredPrefixesClientListByPeeringResponse{}, err
	}
	return result, nil
}
