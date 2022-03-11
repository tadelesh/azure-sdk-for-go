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

// ReplicationVaultHealthClient contains the methods for the ReplicationVaultHealth group.
// Don't use this type directly, use NewReplicationVaultHealthClient() instead.
type ReplicationVaultHealthClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationVaultHealthClient creates a new instance of ReplicationVaultHealthClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationVaultHealthClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationVaultHealthClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ReplicationVaultHealthClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              string(ep),
		pl:                armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the health details of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ReplicationVaultHealthClientGetOptions contains the optional parameters for the ReplicationVaultHealthClient.Get
// method.
func (client *ReplicationVaultHealthClient) Get(ctx context.Context, options *ReplicationVaultHealthClientGetOptions) (ReplicationVaultHealthClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ReplicationVaultHealthClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationVaultHealthClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationVaultHealthClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationVaultHealthClient) getCreateRequest(ctx context.Context, options *ReplicationVaultHealthClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth"
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
func (client *ReplicationVaultHealthClient) getHandleResponse(resp *http.Response) (ReplicationVaultHealthClientGetResponse, error) {
	result := ReplicationVaultHealthClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultHealthDetails); err != nil {
		return ReplicationVaultHealthClientGetResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refreshes health summary of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ReplicationVaultHealthClientBeginRefreshOptions contains the optional parameters for the ReplicationVaultHealthClient.BeginRefresh
// method.
func (client *ReplicationVaultHealthClient) BeginRefresh(ctx context.Context, options *ReplicationVaultHealthClientBeginRefreshOptions) (ReplicationVaultHealthClientRefreshPollerResponse, error) {
	resp, err := client.refresh(ctx, options)
	if err != nil {
		return ReplicationVaultHealthClientRefreshPollerResponse{}, err
	}
	result := ReplicationVaultHealthClientRefreshPollerResponse{}
	pt, err := armruntime.NewPoller("ReplicationVaultHealthClient.Refresh", "", resp, client.pl)
	if err != nil {
		return ReplicationVaultHealthClientRefreshPollerResponse{}, err
	}
	result.Poller = &ReplicationVaultHealthClientRefreshPoller{
		pt: pt,
	}
	return result, nil
}

// Refresh - Refreshes health summary of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ReplicationVaultHealthClient) refresh(ctx context.Context, options *ReplicationVaultHealthClientBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, options)
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

// refreshCreateRequest creates the Refresh request.
func (client *ReplicationVaultHealthClient) refreshCreateRequest(ctx context.Context, options *ReplicationVaultHealthClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth/default/refresh"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
