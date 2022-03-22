//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// FailoverGroupsClient contains the methods for the FailoverGroups group.
// Don't use this type directly, use NewFailoverGroupsClient() instead.
type FailoverGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFailoverGroupsClient creates a new instance of FailoverGroupsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFailoverGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FailoverGroupsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &FailoverGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// parameters - The failover group parameters.
// options - FailoverGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the FailoverGroupsClient.BeginCreateOrUpdate
// method.
func (client *FailoverGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup, options *FailoverGroupsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[FailoverGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, failoverGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FailoverGroupsClientCreateOrUpdateResponse]("FailoverGroupsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FailoverGroupsClientCreateOrUpdateResponse]("FailoverGroupsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FailoverGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup, options *FailoverGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FailoverGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup, options *FailoverGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// options - FailoverGroupsClientBeginDeleteOptions contains the optional parameters for the FailoverGroupsClient.BeginDelete
// method.
func (client *FailoverGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*armruntime.Poller[FailoverGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, failoverGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FailoverGroupsClientDeleteResponse]("FailoverGroupsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FailoverGroupsClientDeleteResponse]("FailoverGroupsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FailoverGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, options)
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
func (client *FailoverGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Fails over from the current primary server to this server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// options - FailoverGroupsClientBeginFailoverOptions contains the optional parameters for the FailoverGroupsClient.BeginFailover
// method.
func (client *FailoverGroupsClient) BeginFailover(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginFailoverOptions) (*armruntime.Poller[FailoverGroupsClientFailoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.failover(ctx, resourceGroupName, serverName, failoverGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FailoverGroupsClientFailoverResponse]("FailoverGroupsClient.Failover", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FailoverGroupsClientFailoverResponse]("FailoverGroupsClient.Failover", options.ResumeToken, client.pl, nil)
	}
}

// Failover - Fails over from the current primary server to this server.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FailoverGroupsClient) failover(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, options)
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

// failoverCreateRequest creates the Failover request.
func (client *FailoverGroupsClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginForceFailoverAllowDataLoss - Fails over from the current primary server to this server. This operation might result
// in data loss.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// options - FailoverGroupsClientBeginForceFailoverAllowDataLossOptions contains the optional parameters for the FailoverGroupsClient.BeginForceFailoverAllowDataLoss
// method.
func (client *FailoverGroupsClient) BeginForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (*armruntime.Poller[FailoverGroupsClientForceFailoverAllowDataLossResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.forceFailoverAllowDataLoss(ctx, resourceGroupName, serverName, failoverGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FailoverGroupsClientForceFailoverAllowDataLossResponse]("FailoverGroupsClient.ForceFailoverAllowDataLoss", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FailoverGroupsClientForceFailoverAllowDataLossResponse]("FailoverGroupsClient.ForceFailoverAllowDataLoss", options.ResumeToken, client.pl, nil)
	}
}

// ForceFailoverAllowDataLoss - Fails over from the current primary server to this server. This operation might result in
// data loss.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FailoverGroupsClient) forceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (*http.Response, error) {
	req, err := client.forceFailoverAllowDataLossCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, options)
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

// forceFailoverAllowDataLossCreateRequest creates the ForceFailoverAllowDataLoss request.
func (client *FailoverGroupsClient) forceFailoverAllowDataLossCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientBeginForceFailoverAllowDataLossOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// options - FailoverGroupsClientGetOptions contains the optional parameters for the FailoverGroupsClient.Get method.
func (client *FailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, options)
	if err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FailoverGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FailoverGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, options *FailoverGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FailoverGroupsClient) getHandleResponse(resp *http.Response) (FailoverGroupsClientGetResponse, error) {
	result := FailoverGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FailoverGroup); err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Lists the failover groups in a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// options - FailoverGroupsClientListByServerOptions contains the optional parameters for the FailoverGroupsClient.ListByServer
// method.
func (client *FailoverGroupsClient) ListByServer(resourceGroupName string, serverName string, options *FailoverGroupsClientListByServerOptions) *runtime.Pager[FailoverGroupsClientListByServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[FailoverGroupsClientListByServerResponse]{
		More: func(page FailoverGroupsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FailoverGroupsClientListByServerResponse) (FailoverGroupsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FailoverGroupsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FailoverGroupsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FailoverGroupsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *FailoverGroupsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *FailoverGroupsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *FailoverGroupsClient) listByServerHandleResponse(resp *http.Response) (FailoverGroupsClientListByServerResponse, error) {
	result := FailoverGroupsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FailoverGroupListResult); err != nil {
		return FailoverGroupsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server containing the failover group.
// failoverGroupName - The name of the failover group.
// parameters - The failover group parameters.
// options - FailoverGroupsClientBeginUpdateOptions contains the optional parameters for the FailoverGroupsClient.BeginUpdate
// method.
func (client *FailoverGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate, options *FailoverGroupsClientBeginUpdateOptions) (*armruntime.Poller[FailoverGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serverName, failoverGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FailoverGroupsClientUpdateResponse]("FailoverGroupsClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FailoverGroupsClientUpdateResponse]("FailoverGroupsClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a failover group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FailoverGroupsClient) update(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate, options *FailoverGroupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, failoverGroupName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *FailoverGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate, options *FailoverGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
