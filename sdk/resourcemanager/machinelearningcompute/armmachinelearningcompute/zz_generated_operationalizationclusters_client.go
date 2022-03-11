//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningcompute

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
	"strconv"
	"strings"
)

// OperationalizationClustersClient contains the methods for the OperationalizationClusters group.
// Don't use this type directly, use NewOperationalizationClustersClient() instead.
type OperationalizationClustersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationalizationClustersClient creates a new instance of OperationalizationClustersClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationalizationClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationalizationClustersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationalizationClustersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckSystemServicesUpdatesAvailable - Checks if updates are available for system services in the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// options - OperationalizationClustersClientCheckSystemServicesUpdatesAvailableOptions contains the optional parameters for
// the OperationalizationClustersClient.CheckSystemServicesUpdatesAvailable method.
func (client *OperationalizationClustersClient) CheckSystemServicesUpdatesAvailable(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientCheckSystemServicesUpdatesAvailableOptions) (OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse, error) {
	req, err := client.checkSystemServicesUpdatesAvailableCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkSystemServicesUpdatesAvailableHandleResponse(resp)
}

// checkSystemServicesUpdatesAvailableCreateRequest creates the CheckSystemServicesUpdatesAvailable request.
func (client *OperationalizationClustersClient) checkSystemServicesUpdatesAvailableCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientCheckSystemServicesUpdatesAvailableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}/checkSystemServicesUpdatesAvailable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkSystemServicesUpdatesAvailableHandleResponse handles the CheckSystemServicesUpdatesAvailable response.
func (client *OperationalizationClustersClient) checkSystemServicesUpdatesAvailableHandleResponse(resp *http.Response) (OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse, error) {
	result := OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckSystemServicesUpdatesAvailableResponse); err != nil {
		return OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update an operationalization cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// parameters - Parameters supplied to create or update an Operationalization cluster.
// options - OperationalizationClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the OperationalizationClustersClient.BeginCreateOrUpdate
// method.
func (client *OperationalizationClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters OperationalizationCluster, options *OperationalizationClustersClientBeginCreateOrUpdateOptions) (OperationalizationClustersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return OperationalizationClustersClientCreateOrUpdatePollerResponse{}, err
	}
	result := OperationalizationClustersClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("OperationalizationClustersClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return OperationalizationClustersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &OperationalizationClustersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update an operationalization cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OperationalizationClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters OperationalizationCluster, options *OperationalizationClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
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
func (client *OperationalizationClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters OperationalizationCluster, options *OperationalizationClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// options - OperationalizationClustersClientBeginDeleteOptions contains the optional parameters for the OperationalizationClustersClient.BeginDelete
// method.
func (client *OperationalizationClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginDeleteOptions) (OperationalizationClustersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return OperationalizationClustersClientDeletePollerResponse{}, err
	}
	result := OperationalizationClustersClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("OperationalizationClustersClient.Delete", "", resp, client.pl)
	if err != nil {
		return OperationalizationClustersClientDeletePollerResponse{}, err
	}
	result.Poller = &OperationalizationClustersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OperationalizationClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OperationalizationClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	if options != nil && options.DeleteAll != nil {
		reqQP.Set("deleteAll", strconv.FormatBool(*options.DeleteAll))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the operationalization cluster resource view. Note that the credentials are not returned by this call. Call
// ListKeys to get them.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// options - OperationalizationClustersClientGetOptions contains the optional parameters for the OperationalizationClustersClient.Get
// method.
func (client *OperationalizationClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientGetOptions) (OperationalizationClustersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return OperationalizationClustersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationalizationClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationalizationClustersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationalizationClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationalizationClustersClient) getHandleResponse(resp *http.Response) (OperationalizationClustersClientGetResponse, error) {
	result := OperationalizationClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationalizationCluster); err != nil {
		return OperationalizationClustersClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets the clusters in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// options - OperationalizationClustersClientListByResourceGroupOptions contains the optional parameters for the OperationalizationClustersClient.ListByResourceGroup
// method.
func (client *OperationalizationClustersClient) ListByResourceGroup(resourceGroupName string, options *OperationalizationClustersClientListByResourceGroupOptions) *OperationalizationClustersClientListByResourceGroupPager {
	return &OperationalizationClustersClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp OperationalizationClustersClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PaginatedOperationalizationClustersList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OperationalizationClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OperationalizationClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters"
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
	reqQP.Set("api-version", "2017-08-01-preview")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OperationalizationClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (OperationalizationClustersClientListByResourceGroupResponse, error) {
	result := OperationalizationClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedOperationalizationClustersList); err != nil {
		return OperationalizationClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscriptionID - Gets the operationalization clusters in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OperationalizationClustersClientListBySubscriptionIDOptions contains the optional parameters for the OperationalizationClustersClient.ListBySubscriptionID
// method.
func (client *OperationalizationClustersClient) ListBySubscriptionID(options *OperationalizationClustersClientListBySubscriptionIDOptions) *OperationalizationClustersClientListBySubscriptionIDPager {
	return &OperationalizationClustersClientListBySubscriptionIDPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionIDCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp OperationalizationClustersClientListBySubscriptionIDResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PaginatedOperationalizationClustersList.NextLink)
		},
	}
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *OperationalizationClustersClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *OperationalizationClustersClientListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningCompute/operationalizationClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *OperationalizationClustersClient) listBySubscriptionIDHandleResponse(resp *http.Response) (OperationalizationClustersClientListBySubscriptionIDResponse, error) {
	result := OperationalizationClustersClientListBySubscriptionIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedOperationalizationClustersList); err != nil {
		return OperationalizationClustersClientListBySubscriptionIDResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets the credentials for the specified cluster such as Storage, ACR and ACS credentials. This is a long running
// operation because it fetches keys from dependencies.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// options - OperationalizationClustersClientListKeysOptions contains the optional parameters for the OperationalizationClustersClient.ListKeys
// method.
func (client *OperationalizationClustersClient) ListKeys(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientListKeysOptions) (OperationalizationClustersClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return OperationalizationClustersClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationalizationClustersClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationalizationClustersClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *OperationalizationClustersClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *OperationalizationClustersClient) listKeysHandleResponse(resp *http.Response) (OperationalizationClustersClientListKeysResponse, error) {
	result := OperationalizationClustersClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationalizationClusterCredentials); err != nil {
		return OperationalizationClustersClientListKeysResponse{}, err
	}
	return result, nil
}

// Update - The PATCH operation can be used to update only the tags for a cluster. Use PUT operation to update other properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// parameters - The parameters supplied to patch the cluster.
// options - OperationalizationClustersClientUpdateOptions contains the optional parameters for the OperationalizationClustersClient.Update
// method.
func (client *OperationalizationClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters OperationalizationClusterUpdateParameters, options *OperationalizationClustersClientUpdateOptions) (OperationalizationClustersClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return OperationalizationClustersClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationalizationClustersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationalizationClustersClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *OperationalizationClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters OperationalizationClusterUpdateParameters, options *OperationalizationClustersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *OperationalizationClustersClient) updateHandleResponse(resp *http.Response) (OperationalizationClustersClientUpdateResponse, error) {
	result := OperationalizationClustersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationalizationCluster); err != nil {
		return OperationalizationClustersClientUpdateResponse{}, err
	}
	return result, nil
}

// BeginUpdateSystemServices - Updates system services in a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group in which the cluster is located.
// clusterName - The name of the cluster.
// options - OperationalizationClustersClientBeginUpdateSystemServicesOptions contains the optional parameters for the OperationalizationClustersClient.BeginUpdateSystemServices
// method.
func (client *OperationalizationClustersClient) BeginUpdateSystemServices(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginUpdateSystemServicesOptions) (OperationalizationClustersClientUpdateSystemServicesPollerResponse, error) {
	resp, err := client.updateSystemServices(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return OperationalizationClustersClientUpdateSystemServicesPollerResponse{}, err
	}
	result := OperationalizationClustersClientUpdateSystemServicesPollerResponse{}
	pt, err := armruntime.NewPoller("OperationalizationClustersClient.UpdateSystemServices", "", resp, client.pl)
	if err != nil {
		return OperationalizationClustersClientUpdateSystemServicesPollerResponse{}, err
	}
	result.Poller = &OperationalizationClustersClientUpdateSystemServicesPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateSystemServices - Updates system services in a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OperationalizationClustersClient) updateSystemServices(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginUpdateSystemServicesOptions) (*http.Response, error) {
	req, err := client.updateSystemServicesCreateRequest(ctx, resourceGroupName, clusterName, options)
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

// updateSystemServicesCreateRequest creates the UpdateSystemServices request.
func (client *OperationalizationClustersClient) updateSystemServicesCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *OperationalizationClustersClientBeginUpdateSystemServicesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}/updateSystemServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
