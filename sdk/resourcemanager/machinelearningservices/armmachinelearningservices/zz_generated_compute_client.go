//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

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

// ComputeClient contains the methods for the Compute group.
// Don't use this type directly, use NewComputeClient() instead.
type ComputeClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewComputeClient creates a new instance of ComputeClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewComputeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ComputeClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ComputeClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation. If your intent is to create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// parameters - Payload with Machine Learning compute definition.
// options - ComputeClientBeginCreateOrUpdateOptions contains the optional parameters for the ComputeClient.BeginCreateOrUpdate
// method.
func (client *ComputeClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (ComputeClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return ComputeClientCreateOrUpdatePollerResponse{}, err
	}
	result := ComputeClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ComputeClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ComputeClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation. If your intent is to create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
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
func (client *ComputeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes specified Machine Learning compute.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// underlyingResourceAction - Delete the underlying compute if 'Delete', or detach the underlying compute from workspace if
// 'Detach'.
// options - ComputeClientBeginDeleteOptions contains the optional parameters for the ComputeClient.BeginDelete method.
func (client *ComputeClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (ComputeClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
	if err != nil {
		return ComputeClientDeletePollerResponse{}, err
	}
	result := ComputeClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.Delete", "", resp, client.pl)
	if err != nil {
		return ComputeClientDeletePollerResponse{}, err
	}
	result.Poller = &ComputeClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes specified Machine Learning compute.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
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

// deleteCreateRequest creates the Delete request.
func (client *ComputeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	reqQP.Set("underlyingResourceAction", string(underlyingResourceAction))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use
// 'keys' nested resource to get them.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientGetOptions contains the optional parameters for the ComputeClient.Get method.
func (client *ComputeClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientGetOptions) (ComputeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComputeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComputeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComputeClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComputeClient) getHandleResponse(resp *http.Response) (ComputeClientGetResponse, error) {
	result := ComputeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputeResource); err != nil {
		return ComputeClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets computes in specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// options - ComputeClientListOptions contains the optional parameters for the ComputeClient.List method.
func (client *ComputeClient) List(resourceGroupName string, workspaceName string, options *ComputeClientListOptions) *ComputeClientListPager {
	return &ComputeClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp ComputeClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PaginatedComputeResourcesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ComputeClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ComputeClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComputeClient) listHandleResponse(resp *http.Response) (ComputeClientListResponse, error) {
	result := ComputeClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedComputeResourcesList); err != nil {
		return ComputeClientListResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientListKeysOptions contains the optional parameters for the ComputeClient.ListKeys method.
func (client *ComputeClient) ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListKeysOptions) (ComputeClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComputeClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ComputeClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ComputeClient) listKeysHandleResponse(resp *http.Response) (ComputeClientListKeysResponse, error) {
	result := ComputeClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	return result, nil
}

// ListNodes - Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientListNodesOptions contains the optional parameters for the ComputeClient.ListNodes method.
func (client *ComputeClient) ListNodes(resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListNodesOptions) *ComputeClientListNodesPager {
	return &ComputeClientListNodesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listNodesCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
		},
		advancer: func(ctx context.Context, resp ComputeClientListNodesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AmlComputeNodesInformation.NextLink)
		},
	}
}

// listNodesCreateRequest creates the ListNodes request.
func (client *ComputeClient) listNodesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListNodesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listNodes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listNodesHandleResponse handles the ListNodes response.
func (client *ComputeClient) listNodesHandleResponse(resp *http.Response) (ComputeClientListNodesResponse, error) {
	result := ComputeClientListNodesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AmlComputeNodesInformation); err != nil {
		return ComputeClientListNodesResponse{}, err
	}
	return result, nil
}

// BeginRestart - Posts a restart action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientBeginRestartOptions contains the optional parameters for the ComputeClient.BeginRestart method.
func (client *ComputeClient) BeginRestart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (ComputeClientRestartPollerResponse, error) {
	resp, err := client.restart(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientRestartPollerResponse{}, err
	}
	result := ComputeClientRestartPollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.Restart", "", resp, client.pl)
	if err != nil {
		return ComputeClientRestartPollerResponse{}, err
	}
	result.Poller = &ComputeClientRestartPoller{
		pt: pt,
	}
	return result, nil
}

// Restart - Posts a restart action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) restart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ComputeClient) restartCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStart - Posts a start action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientBeginStartOptions contains the optional parameters for the ComputeClient.BeginStart method.
func (client *ComputeClient) BeginStart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (ComputeClientStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientStartPollerResponse{}, err
	}
	result := ComputeClientStartPollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.Start", "", resp, client.pl)
	if err != nil {
		return ComputeClientStartPollerResponse{}, err
	}
	result.Poller = &ComputeClientStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Posts a start action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) start(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ComputeClient) startCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - Posts a stop action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// options - ComputeClientBeginStopOptions contains the optional parameters for the ComputeClient.BeginStop method.
func (client *ComputeClient) BeginStop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (ComputeClientStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientStopPollerResponse{}, err
	}
	result := ComputeClientStopPollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.Stop", "", resp, client.pl)
	if err != nil {
		return ComputeClientStopPollerResponse{}, err
	}
	result.Poller = &ComputeClientStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Posts a stop action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) stop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ComputeClient) stopCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// computeName - Name of the Azure Machine Learning compute.
// parameters - Additional parameters for cluster update.
// options - ComputeClientBeginUpdateOptions contains the optional parameters for the ComputeClient.BeginUpdate method.
func (client *ComputeClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (ComputeClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return ComputeClientUpdatePollerResponse{}, err
	}
	result := ComputeClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ComputeClient.Update", "", resp, client.pl)
	if err != nil {
		return ComputeClientUpdatePollerResponse{}, err
	}
	result.Poller = &ComputeClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ComputeClient) update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ComputeClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
