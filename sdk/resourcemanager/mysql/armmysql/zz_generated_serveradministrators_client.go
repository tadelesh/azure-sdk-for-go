//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// ServerAdministratorsClient contains the methods for the ServerAdministrators group.
// Don't use this type directly, use NewServerAdministratorsClient() instead.
type ServerAdministratorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerAdministratorsClient creates a new instance of ServerAdministratorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerAdministratorsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServerAdministratorsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite
// the existing administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// properties - The required parameters for creating or updating an AAD server administrator.
// options - ServerAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerAdministratorsClient.BeginCreateOrUpdate
// method.
func (client *ServerAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ServerAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, properties, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerAdministratorsClientCreateOrUpdateResponse]("ServerAdministratorsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerAdministratorsClientCreateOrUpdateResponse]("ServerAdministratorsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite
// the existing administrator.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, properties, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// BeginDelete - Deletes server active directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - ServerAdministratorsClientBeginDeleteOptions contains the optional parameters for the ServerAdministratorsClient.BeginDelete
// method.
func (client *ServerAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*armruntime.Poller[ServerAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerAdministratorsClientDeleteResponse]("ServerAdministratorsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerAdministratorsClientDeleteResponse]("ServerAdministratorsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes server active directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
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
func (client *ServerAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about a AAD server administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - ServerAdministratorsClientGetOptions contains the optional parameters for the ServerAdministratorsClient.Get
// method.
func (client *ServerAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientGetOptions) (ServerAdministratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdministratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAdministratorsClient) getHandleResponse(resp *http.Response) (ServerAdministratorsClientGetResponse, error) {
	result := ServerAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResource); err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns a list of server Administrators.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - ServerAdministratorsClientListOptions contains the optional parameters for the ServerAdministratorsClient.List
// method.
func (client *ServerAdministratorsClient) List(resourceGroupName string, serverName string, options *ServerAdministratorsClientListOptions) *runtime.Pager[ServerAdministratorsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServerAdministratorsClientListResponse]{
		More: func(page ServerAdministratorsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerAdministratorsClientListResponse) (ServerAdministratorsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ServerAdministratorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerAdministratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerAdministratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServerAdministratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerAdministratorsClient) listHandleResponse(resp *http.Response) (ServerAdministratorsClientListResponse, error) {
	result := ServerAdministratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResourceListResult); err != nil {
		return ServerAdministratorsClientListResponse{}, err
	}
	return result, nil
}
