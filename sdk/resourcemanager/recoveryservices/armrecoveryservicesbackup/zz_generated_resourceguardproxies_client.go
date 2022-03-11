//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ResourceGuardProxiesClient contains the methods for the ResourceGuardProxies group.
// Don't use this type directly, use NewResourceGuardProxiesClient() instead.
type ResourceGuardProxiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceGuardProxiesClient creates a new instance of ResourceGuardProxiesClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceGuardProxiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceGuardProxiesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ResourceGuardProxiesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - List the ResourceGuardProxies under vault
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - ResourceGuardProxiesClientGetOptions contains the optional parameters for the ResourceGuardProxiesClient.Get
// method.
func (client *ResourceGuardProxiesClient) Get(vaultName string, resourceGroupName string, options *ResourceGuardProxiesClientGetOptions) *ResourceGuardProxiesClientGetPager {
	return &ResourceGuardProxiesClientGetPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getCreateRequest(ctx, vaultName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ResourceGuardProxiesClientGetResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceGuardProxyBaseResourceList.NextLink)
		},
	}
}

// getCreateRequest creates the Get request.
func (client *ResourceGuardProxiesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *ResourceGuardProxiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *ResourceGuardProxiesClient) getHandleResponse(resp *http.Response) (ResourceGuardProxiesClientGetResponse, error) {
	result := ResourceGuardProxiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResourceList); err != nil {
		return ResourceGuardProxiesClientGetResponse{}, err
	}
	return result, nil
}
