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

// ManagedInstanceEncryptionProtectorsClient contains the methods for the ManagedInstanceEncryptionProtectors group.
// Don't use this type directly, use NewManagedInstanceEncryptionProtectorsClient() instead.
type ManagedInstanceEncryptionProtectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedInstanceEncryptionProtectorsClient creates a new instance of ManagedInstanceEncryptionProtectorsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedInstanceEncryptionProtectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedInstanceEncryptionProtectorsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedInstanceEncryptionProtectorsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// encryptionProtectorName - The name of the encryption protector to be updated.
// parameters - The requested encryption protector resource state.
// options - ManagedInstanceEncryptionProtectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the
// ManagedInstanceEncryptionProtectorsClient.BeginCreateOrUpdate method.
func (client *ManagedInstanceEncryptionProtectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse]("ManagedInstanceEncryptionProtectorsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse]("ManagedInstanceEncryptionProtectorsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedInstanceEncryptionProtectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, parameters, options)
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
func (client *ManagedInstanceEncryptionProtectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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

// Get - Gets a managed instance encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// encryptionProtectorName - The name of the encryption protector to be retrieved.
// options - ManagedInstanceEncryptionProtectorsClientGetOptions contains the optional parameters for the ManagedInstanceEncryptionProtectorsClient.Get
// method.
func (client *ManagedInstanceEncryptionProtectorsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsClientGetOptions) (ManagedInstanceEncryptionProtectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceEncryptionProtectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceEncryptionProtectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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
func (client *ManagedInstanceEncryptionProtectorsClient) getHandleResponse(resp *http.Response) (ManagedInstanceEncryptionProtectorsClientGetResponse, error) {
	result := ManagedInstanceEncryptionProtectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceEncryptionProtector); err != nil {
		return ManagedInstanceEncryptionProtectorsClientGetResponse{}, err
	}
	return result, nil
}

// ListByInstance - Gets a list of managed instance encryption protectors
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstanceEncryptionProtectorsClientListByInstanceOptions contains the optional parameters for the ManagedInstanceEncryptionProtectorsClient.ListByInstance
// method.
func (client *ManagedInstanceEncryptionProtectorsClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *ManagedInstanceEncryptionProtectorsClientListByInstanceOptions) *runtime.Pager[ManagedInstanceEncryptionProtectorsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedInstanceEncryptionProtectorsClientListByInstanceResponse]{
		More: func(page ManagedInstanceEncryptionProtectorsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceEncryptionProtectorsClientListByInstanceResponse) (ManagedInstanceEncryptionProtectorsClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstanceEncryptionProtectorsClientListByInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstanceEncryptionProtectorsClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstanceEncryptionProtectorsClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceEncryptionProtectorsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceEncryptionProtectorsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector"
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
func (client *ManagedInstanceEncryptionProtectorsClient) listByInstanceHandleResponse(resp *http.Response) (ManagedInstanceEncryptionProtectorsClientListByInstanceResponse, error) {
	result := ManagedInstanceEncryptionProtectorsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceEncryptionProtectorListResult); err != nil {
		return ManagedInstanceEncryptionProtectorsClientListByInstanceResponse{}, err
	}
	return result, nil
}

// BeginRevalidate - Revalidates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// encryptionProtectorName - The name of the encryption protector to be updated.
// options - ManagedInstanceEncryptionProtectorsClientBeginRevalidateOptions contains the optional parameters for the ManagedInstanceEncryptionProtectorsClient.BeginRevalidate
// method.
func (client *ManagedInstanceEncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsClientBeginRevalidateOptions) (*armruntime.Poller[ManagedInstanceEncryptionProtectorsClientRevalidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revalidate(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagedInstanceEncryptionProtectorsClientRevalidateResponse]("ManagedInstanceEncryptionProtectorsClient.Revalidate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagedInstanceEncryptionProtectorsClientRevalidateResponse]("ManagedInstanceEncryptionProtectorsClient.Revalidate", options.ResumeToken, client.pl, nil)
	}
}

// Revalidate - Revalidates an existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagedInstanceEncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsClientBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
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

// revalidateCreateRequest creates the Revalidate request.
func (client *ManagedInstanceEncryptionProtectorsClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsClientBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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
	return req, nil
}
