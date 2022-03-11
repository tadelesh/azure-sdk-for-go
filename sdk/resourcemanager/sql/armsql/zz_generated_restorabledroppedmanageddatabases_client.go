//go:build go1.16
// +build go1.16

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

// RestorableDroppedManagedDatabasesClient contains the methods for the RestorableDroppedManagedDatabases group.
// Don't use this type directly, use NewRestorableDroppedManagedDatabasesClient() instead.
type RestorableDroppedManagedDatabasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableDroppedManagedDatabasesClient creates a new instance of RestorableDroppedManagedDatabasesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableDroppedManagedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableDroppedManagedDatabasesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RestorableDroppedManagedDatabasesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets a restorable dropped managed database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - RestorableDroppedManagedDatabasesClientGetOptions contains the optional parameters for the RestorableDroppedManagedDatabasesClient.Get
// method.
func (client *RestorableDroppedManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *RestorableDroppedManagedDatabasesClientGetOptions) (RestorableDroppedManagedDatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, options)
	if err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorableDroppedManagedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *RestorableDroppedManagedDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorableDroppedManagedDatabasesClient) getHandleResponse(resp *http.Response) (RestorableDroppedManagedDatabasesClientGetResponse, error) {
	result := RestorableDroppedManagedDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedManagedDatabase); err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// ListByInstance - Gets a list of restorable dropped managed databases.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - RestorableDroppedManagedDatabasesClientListByInstanceOptions contains the optional parameters for the RestorableDroppedManagedDatabasesClient.ListByInstance
// method.
func (client *RestorableDroppedManagedDatabasesClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *RestorableDroppedManagedDatabasesClientListByInstanceOptions) *RestorableDroppedManagedDatabasesClientListByInstancePager {
	return &RestorableDroppedManagedDatabasesClientListByInstancePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp RestorableDroppedManagedDatabasesClientListByInstanceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RestorableDroppedManagedDatabaseListResult.NextLink)
		},
	}
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *RestorableDroppedManagedDatabasesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *RestorableDroppedManagedDatabasesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *RestorableDroppedManagedDatabasesClient) listByInstanceHandleResponse(resp *http.Response) (RestorableDroppedManagedDatabasesClientListByInstanceResponse, error) {
	result := RestorableDroppedManagedDatabasesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedManagedDatabaseListResult); err != nil {
		return RestorableDroppedManagedDatabasesClientListByInstanceResponse{}, err
	}
	return result, nil
}
