//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

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

// GroupsClient contains the methods for the SQLVirtualMachineGroups group.
// Don't use this type directly, use NewGroupsClient() instead.
type GroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGroupsClient creates a new instance of GroupsClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GroupsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &GroupsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// parameters - The SQL virtual machine group.
// options - GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate
// method.
func (client *GroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (GroupsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
	if err != nil {
		return GroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result := GroupsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("GroupsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return GroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &GroupsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
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
func (client *GroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// options - GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
func (client *GroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (GroupsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
	if err != nil {
		return GroupsClientDeletePollerResponse{}, err
	}
	result := GroupsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("GroupsClient.Delete", "", resp, client.pl)
	if err != nil {
		return GroupsClientDeletePollerResponse{}, err
	}
	result.Poller = &GroupsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
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
func (client *GroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// options - GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
func (client *GroupsClient) Get(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientGetOptions) (GroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GroupsClient) getHandleResponse(resp *http.Response) (GroupsClientGetResponse, error) {
	result := GroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Group); err != nil {
		return GroupsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all SQL virtual machine groups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - GroupsClientListOptions contains the optional parameters for the GroupsClient.List method.
func (client *GroupsClient) List(options *GroupsClientListOptions) *GroupsClientListPager {
	return &GroupsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp GroupsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *GroupsClient) listCreateRequest(ctx context.Context, options *GroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GroupsClient) listHandleResponse(resp *http.Response) (GroupsClientListResponse, error) {
	result := GroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all SQL virtual machine groups in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// options - GroupsClientListByResourceGroupOptions contains the optional parameters for the GroupsClient.ListByResourceGroup
// method.
func (client *GroupsClient) ListByResourceGroup(resourceGroupName string, options *GroupsClientListByResourceGroupOptions) *GroupsClientListByResourceGroupPager {
	return &GroupsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp GroupsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"
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
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (GroupsClientListByResourceGroupResponse, error) {
	result := GroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates SQL virtual machine group tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// parameters - The SQL virtual machine group.
// options - GroupsClientBeginUpdateOptions contains the optional parameters for the GroupsClient.BeginUpdate method.
func (client *GroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (GroupsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
	if err != nil {
		return GroupsClientUpdatePollerResponse{}, err
	}
	result := GroupsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("GroupsClient.Update", "", resp, client.pl)
	if err != nil {
		return GroupsClientUpdatePollerResponse{}, err
	}
	result.Poller = &GroupsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates SQL virtual machine group tags.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GroupsClient) update(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
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
func (client *GroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
