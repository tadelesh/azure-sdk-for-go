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

// ManagedDatabasesClient contains the methods for the ManagedDatabases group.
// Don't use this type directly, use NewManagedDatabasesClient() instead.
type ManagedDatabasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedDatabasesClient creates a new instance of ManagedDatabasesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedDatabasesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedDatabasesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCompleteRestore - Completes the restore operation on a managed database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// parameters - The definition for completing the restore of this managed database.
// options - ManagedDatabasesClientBeginCompleteRestoreOptions contains the optional parameters for the ManagedDatabasesClient.BeginCompleteRestore
// method.
func (client *ManagedDatabasesClient) BeginCompleteRestore(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters CompleteDatabaseRestoreDefinition, options *ManagedDatabasesClientBeginCompleteRestoreOptions) (*armruntime.Poller[ManagedDatabasesClientCompleteRestoreResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.completeRestore(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedDatabasesClientCompleteRestoreResponse]("ManagedDatabasesClient.CompleteRestore", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedDatabasesClientCompleteRestoreResponse]("ManagedDatabasesClient.CompleteRestore", options.ResumeToken, client.pl, nil)
	}
}

// CompleteRestore - Completes the restore operation on a managed database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedDatabasesClient) completeRestore(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters CompleteDatabaseRestoreDefinition, options *ManagedDatabasesClientBeginCompleteRestoreOptions) (*http.Response, error) {
	req, err := client.completeRestoreCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
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

// completeRestoreCreateRequest creates the CompleteRestore request.
func (client *ManagedDatabasesClient) completeRestoreCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters CompleteDatabaseRestoreDefinition, options *ManagedDatabasesClientBeginCompleteRestoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/completeRestore"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdate - Creates a new database or updates an existing database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// parameters - The requested database resource state.
// options - ManagedDatabasesClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedDatabasesClient.BeginCreateOrUpdate
// method.
func (client *ManagedDatabasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabase, options *ManagedDatabasesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ManagedDatabasesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedDatabasesClientCreateOrUpdateResponse]("ManagedDatabasesClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedDatabasesClientCreateOrUpdateResponse]("ManagedDatabasesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new database or updates an existing database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedDatabasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabase, options *ManagedDatabasesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
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
func (client *ManagedDatabasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabase, options *ManagedDatabasesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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

// BeginDelete - Deletes a managed database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// options - ManagedDatabasesClientBeginDeleteOptions contains the optional parameters for the ManagedDatabasesClient.BeginDelete
// method.
func (client *ManagedDatabasesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabasesClientBeginDeleteOptions) (*armruntime.Poller[ManagedDatabasesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, databaseName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedDatabasesClientDeleteResponse]("ManagedDatabasesClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedDatabasesClientDeleteResponse]("ManagedDatabasesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a managed database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedDatabasesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabasesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
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
func (client *ManagedDatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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

// Get - Gets a managed database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// options - ManagedDatabasesClientGetOptions contains the optional parameters for the ManagedDatabasesClient.Get method.
func (client *ManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabasesClientGetOptions) (ManagedDatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
	if err != nil {
		return ManagedDatabasesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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
func (client *ManagedDatabasesClient) getHandleResponse(resp *http.Response) (ManagedDatabasesClientGetResponse, error) {
	result := ManagedDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabase); err != nil {
		return ManagedDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// ListByInstance - Gets a list of managed databases.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedDatabasesClientListByInstanceOptions contains the optional parameters for the ManagedDatabasesClient.ListByInstance
// method.
func (client *ManagedDatabasesClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *ManagedDatabasesClientListByInstanceOptions) *runtime.Pager[ManagedDatabasesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedDatabasesClientListByInstanceResponse]{
		More: func(page ManagedDatabasesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabasesClientListByInstanceResponse) (ManagedDatabasesClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabasesClientListByInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabasesClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabasesClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedDatabasesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedDatabasesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedDatabasesClient) listByInstanceHandleResponse(resp *http.Response) (ManagedDatabasesClientListByInstanceResponse, error) {
	result := ManagedDatabasesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseListResult); err != nil {
		return ManagedDatabasesClientListByInstanceResponse{}, err
	}
	return result, nil
}

// ListInaccessibleByInstance - Gets a list of inaccessible managed databases in a managed instance
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedDatabasesClientListInaccessibleByInstanceOptions contains the optional parameters for the ManagedDatabasesClient.ListInaccessibleByInstance
// method.
func (client *ManagedDatabasesClient) ListInaccessibleByInstance(resourceGroupName string, managedInstanceName string, options *ManagedDatabasesClientListInaccessibleByInstanceOptions) *runtime.Pager[ManagedDatabasesClientListInaccessibleByInstanceResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedDatabasesClientListInaccessibleByInstanceResponse]{
		More: func(page ManagedDatabasesClientListInaccessibleByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabasesClientListInaccessibleByInstanceResponse) (ManagedDatabasesClientListInaccessibleByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listInaccessibleByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabasesClientListInaccessibleByInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabasesClientListInaccessibleByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabasesClientListInaccessibleByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listInaccessibleByInstanceHandleResponse(resp)
		},
	})
}

// listInaccessibleByInstanceCreateRequest creates the ListInaccessibleByInstance request.
func (client *ManagedDatabasesClient) listInaccessibleByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedDatabasesClientListInaccessibleByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/inaccessibleManagedDatabases"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listInaccessibleByInstanceHandleResponse handles the ListInaccessibleByInstance response.
func (client *ManagedDatabasesClient) listInaccessibleByInstanceHandleResponse(resp *http.Response) (ManagedDatabasesClientListInaccessibleByInstanceResponse, error) {
	result := ManagedDatabasesClientListInaccessibleByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseListResult); err != nil {
		return ManagedDatabasesClientListInaccessibleByInstanceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// parameters - The requested database resource state.
// options - ManagedDatabasesClientBeginUpdateOptions contains the optional parameters for the ManagedDatabasesClient.BeginUpdate
// method.
func (client *ManagedDatabasesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabaseUpdate, options *ManagedDatabasesClientBeginUpdateOptions) (*armruntime.Poller[ManagedDatabasesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedDatabasesClientUpdateResponse]("ManagedDatabasesClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedDatabasesClientUpdateResponse]("ManagedDatabasesClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an existing database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedDatabasesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabaseUpdate, options *ManagedDatabasesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
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
func (client *ManagedDatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabaseUpdate, options *ManagedDatabasesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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
