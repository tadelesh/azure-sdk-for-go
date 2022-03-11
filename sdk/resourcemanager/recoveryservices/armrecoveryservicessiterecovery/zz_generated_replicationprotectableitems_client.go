//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationProtectableItemsClient contains the methods for the ReplicationProtectableItems group.
// Don't use this type directly, use NewReplicationProtectableItemsClient() instead.
type ReplicationProtectableItemsClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationProtectableItemsClient creates a new instance of ReplicationProtectableItemsClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationProtectableItemsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationProtectableItemsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ReplicationProtectableItemsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              string(ep),
		pl:                armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - The operation to get the details of a protectable item.
// If the operation fails it returns an *azcore.ResponseError type.
// fabricName - Fabric name.
// protectionContainerName - Protection container name.
// protectableItemName - Protectable item name.
// options - ReplicationProtectableItemsClientGetOptions contains the optional parameters for the ReplicationProtectableItemsClient.Get
// method.
func (client *ReplicationProtectableItemsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, protectableItemName string, options *ReplicationProtectableItemsClientGetOptions) (ReplicationProtectableItemsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, protectionContainerName, protectableItemName, options)
	if err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationProtectableItemsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectableItemsClient) getCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, protectableItemName string, options *ReplicationProtectableItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if protectableItemName == "" {
		return nil, errors.New("parameter protectableItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectableItemName}", url.PathEscape(protectableItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationProtectableItemsClient) getHandleResponse(resp *http.Response) (ReplicationProtectableItemsClientGetResponse, error) {
	result := ReplicationProtectableItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectableItem); err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	return result, nil
}

// ListByReplicationProtectionContainers - Lists the protectable items in a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
// fabricName - Fabric name.
// protectionContainerName - Protection container name.
// options - ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions contains the optional parameters
// for the ReplicationProtectableItemsClient.ListByReplicationProtectionContainers method.
func (client *ReplicationProtectableItemsClient) ListByReplicationProtectionContainers(fabricName string, protectionContainerName string, options *ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions) *ReplicationProtectableItemsClientListByReplicationProtectionContainersPager {
	return &ReplicationProtectableItemsClientListByReplicationProtectionContainersPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByReplicationProtectionContainersCreateRequest(ctx, fabricName, protectionContainerName, options)
		},
		advancer: func(ctx context.Context, resp ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProtectableItemCollection.NextLink)
		},
	}
}

// listByReplicationProtectionContainersCreateRequest creates the ListByReplicationProtectionContainers request.
func (client *ReplicationProtectableItemsClient) listByReplicationProtectionContainersCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, options *ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Take != nil {
		reqQP.Set("$take", *options.Take)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByReplicationProtectionContainersHandleResponse handles the ListByReplicationProtectionContainers response.
func (client *ReplicationProtectableItemsClient) listByReplicationProtectionContainersHandleResponse(resp *http.Response) (ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse, error) {
	result := ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectableItemCollection); err != nil {
		return ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse{}, err
	}
	return result, nil
}
