//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRouteCrossConnectionPeeringsClient contains the methods for the ExpressRouteCrossConnectionPeerings group.
// Don't use this type directly, use NewExpressRouteCrossConnectionPeeringsClient() instead.
type ExpressRouteCrossConnectionPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCrossConnectionPeeringsClient creates a new instance of ExpressRouteCrossConnectionPeeringsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCrossConnectionPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteCrossConnectionPeeringsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExpressRouteCrossConnectionPeeringsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// peeringParameters - Parameters supplied to the create or update ExpressRouteCrossConnection peering operation.
// options - ExpressRouteCrossConnectionPeeringsClientBeginCreateOrUpdateOptions contains the optional parameters for the
// ExpressRouteCrossConnectionPeeringsClient.BeginCreateOrUpdate method.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ExpressRouteCrossConnectionPeeringsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ExpressRouteCrossConnectionPeeringsClientCreateOrUpdateResponse]("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRouteCrossConnectionPeeringsClientCreateOrUpdateResponse]("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, peeringParameters)
}

// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// options - ExpressRouteCrossConnectionPeeringsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteCrossConnectionPeeringsClient.BeginDelete
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsClientBeginDeleteOptions) (*armruntime.Poller[ExpressRouteCrossConnectionPeeringsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, crossConnectionName, peeringName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ExpressRouteCrossConnectionPeeringsClientDeleteResponse]("ExpressRouteCrossConnectionPeeringsClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRouteCrossConnectionPeeringsClientDeleteResponse]("ExpressRouteCrossConnectionPeeringsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// options - ExpressRouteCrossConnectionPeeringsClientGetOptions contains the optional parameters for the ExpressRouteCrossConnectionPeeringsClient.Get
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsClientGetOptions) (ExpressRouteCrossConnectionPeeringsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionPeeringsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCrossConnectionPeeringsClient) getHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionPeeringsClientGetResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionPeering); err != nil {
		return ExpressRouteCrossConnectionPeeringsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// options - ExpressRouteCrossConnectionPeeringsClientListOptions contains the optional parameters for the ExpressRouteCrossConnectionPeeringsClient.List
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsClientListOptions) *runtime.Pager[ExpressRouteCrossConnectionPeeringsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ExpressRouteCrossConnectionPeeringsClientListResponse]{
		More: func(page ExpressRouteCrossConnectionPeeringsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCrossConnectionPeeringsClientListResponse) (ExpressRouteCrossConnectionPeeringsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCrossConnectionPeeringsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRouteCrossConnectionPeeringsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCrossConnectionPeeringsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCrossConnectionPeeringsClient) listHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionPeeringsClientListResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionPeeringList); err != nil {
		return ExpressRouteCrossConnectionPeeringsClientListResponse{}, err
	}
	return result, nil
}
