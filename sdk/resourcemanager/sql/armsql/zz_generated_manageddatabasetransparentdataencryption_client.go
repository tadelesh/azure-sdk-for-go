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

// ManagedDatabaseTransparentDataEncryptionClient contains the methods for the ManagedDatabaseTransparentDataEncryption group.
// Don't use this type directly, use NewManagedDatabaseTransparentDataEncryptionClient() instead.
type ManagedDatabaseTransparentDataEncryptionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedDatabaseTransparentDataEncryptionClient creates a new instance of ManagedDatabaseTransparentDataEncryptionClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedDatabaseTransparentDataEncryptionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedDatabaseTransparentDataEncryptionClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedDatabaseTransparentDataEncryptionClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Updates a database's transparent data encryption configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the managed database for which the security alert policy is defined.
// tdeName - The name of the transparent data encryption configuration.
// parameters - The database transparent data encryption.
// options - ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateOptions contains the optional parameters for the
// ManagedDatabaseTransparentDataEncryptionClient.CreateOrUpdate method.
func (client *ManagedDatabaseTransparentDataEncryptionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, tdeName TransparentDataEncryptionName, parameters ManagedTransparentDataEncryption, options *ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateOptions) (ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, tdeName, parameters, options)
	if err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedDatabaseTransparentDataEncryptionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, tdeName TransparentDataEncryptionName, parameters ManagedTransparentDataEncryption, options *ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/transparentDataEncryption/{tdeName}"
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
	if tdeName == "" {
		return nil, errors.New("parameter tdeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tdeName}", url.PathEscape(string(tdeName)))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedDatabaseTransparentDataEncryptionClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse, error) {
	result := ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedTransparentDataEncryption); err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a managed database's transparent data encryption.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the managed database for which the transparent data encryption is defined.
// tdeName - The name of the transparent data encryption configuration.
// options - ManagedDatabaseTransparentDataEncryptionClientGetOptions contains the optional parameters for the ManagedDatabaseTransparentDataEncryptionClient.Get
// method.
func (client *ManagedDatabaseTransparentDataEncryptionClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, tdeName TransparentDataEncryptionName, options *ManagedDatabaseTransparentDataEncryptionClientGetOptions) (ManagedDatabaseTransparentDataEncryptionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, tdeName, options)
	if err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseTransparentDataEncryptionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseTransparentDataEncryptionClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, tdeName TransparentDataEncryptionName, options *ManagedDatabaseTransparentDataEncryptionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/transparentDataEncryption/{tdeName}"
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
	if tdeName == "" {
		return nil, errors.New("parameter tdeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tdeName}", url.PathEscape(string(tdeName)))
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
func (client *ManagedDatabaseTransparentDataEncryptionClient) getHandleResponse(resp *http.Response) (ManagedDatabaseTransparentDataEncryptionClientGetResponse, error) {
	result := ManagedDatabaseTransparentDataEncryptionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedTransparentDataEncryption); err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientGetResponse{}, err
	}
	return result, nil
}

// ListByDatabase - Gets a list of managed database's transparent data encryptions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the managed database for which the transparent data encryption is defined.
// options - ManagedDatabaseTransparentDataEncryptionClientListByDatabaseOptions contains the optional parameters for the
// ManagedDatabaseTransparentDataEncryptionClient.ListByDatabase method.
func (client *ManagedDatabaseTransparentDataEncryptionClient) ListByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseTransparentDataEncryptionClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse]{
		More: func(page ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse) (ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseTransparentDataEncryptionClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseTransparentDataEncryptionClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/transparentDataEncryption"
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

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseTransparentDataEncryptionClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse, error) {
	result := ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedTransparentDataEncryptionListResult); err != nil {
		return ManagedDatabaseTransparentDataEncryptionClientListByDatabaseResponse{}, err
	}
	return result, nil
}
