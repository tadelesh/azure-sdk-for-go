//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armresourceconnector

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

// AppliancesClient contains the methods for the Appliances group.
// Don't use this type directly, use NewAppliancesClient() instead.
type AppliancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAppliancesClient creates a new instance of AppliancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAppliancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AppliancesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AppliancesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an Appliance in the specified Subscription and Resource Group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Appliances name.
// parameters - Parameters supplied to create or update an Appliance.
// options - AppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the AppliancesClient.BeginCreateOrUpdate
// method.
func (client *AppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters Appliance, options *AppliancesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[AppliancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AppliancesClientCreateOrUpdateResponse]("AppliancesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AppliancesClientCreateOrUpdateResponse]("AppliancesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an Appliance in the specified Subscription and Resource Group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AppliancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters Appliance, options *AppliancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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
func (client *AppliancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters Appliance, options *AppliancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an Appliance with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Appliances name.
// options - AppliancesClientBeginDeleteOptions contains the optional parameters for the AppliancesClient.BeginDelete method.
func (client *AppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientBeginDeleteOptions) (*armruntime.Poller[AppliancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AppliancesClientDeleteResponse]("AppliancesClient.Delete", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AppliancesClientDeleteResponse]("AppliancesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an Appliance with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AppliancesClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
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
func (client *AppliancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the details of an Appliance with a specified resource group and name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Appliances name.
// options - AppliancesClientGetOptions contains the optional parameters for the AppliancesClient.Get method.
func (client *AppliancesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientGetOptions) (AppliancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return AppliancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppliancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppliancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AppliancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AppliancesClient) getHandleResponse(resp *http.Response) (AppliancesClientGetResponse, error) {
	result := AppliancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Appliance); err != nil {
		return AppliancesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets a list of Appliances in the specified subscription and resource group. The operation returns
// properties of each Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - AppliancesClientListByResourceGroupOptions contains the optional parameters for the AppliancesClient.ListByResourceGroup
// method.
func (client *AppliancesClient) ListByResourceGroup(resourceGroupName string, options *AppliancesClientListByResourceGroupOptions) *runtime.Pager[AppliancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppliancesClientListByResourceGroupResponse]{
		More: func(page AppliancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppliancesClientListByResourceGroupResponse) (AppliancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppliancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppliancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppliancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AppliancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AppliancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances"
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
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AppliancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AppliancesClientListByResourceGroupResponse, error) {
	result := AppliancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplianceListResult); err != nil {
		return AppliancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets a list of Appliances in the specified subscription. The operation returns properties of each
// Appliance
// If the operation fails it returns an *azcore.ResponseError type.
// options - AppliancesClientListBySubscriptionOptions contains the optional parameters for the AppliancesClient.ListBySubscription
// method.
func (client *AppliancesClient) ListBySubscription(options *AppliancesClientListBySubscriptionOptions) *runtime.Pager[AppliancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppliancesClientListBySubscriptionResponse]{
		More: func(page AppliancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppliancesClientListBySubscriptionResponse) (AppliancesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppliancesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppliancesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppliancesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AppliancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AppliancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceConnector/appliances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AppliancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AppliancesClientListBySubscriptionResponse, error) {
	result := AppliancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplianceListResult); err != nil {
		return AppliancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListClusterUserCredential - Returns the cluster user credentials for the dedicated appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Appliances name.
// options - AppliancesClientListClusterUserCredentialOptions contains the optional parameters for the AppliancesClient.ListClusterUserCredential
// method.
func (client *AppliancesClient) ListClusterUserCredential(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientListClusterUserCredentialOptions) (AppliancesClientListClusterUserCredentialResponse, error) {
	req, err := client.listClusterUserCredentialCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return AppliancesClientListClusterUserCredentialResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppliancesClientListClusterUserCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppliancesClientListClusterUserCredentialResponse{}, runtime.NewResponseError(resp)
	}
	return client.listClusterUserCredentialHandleResponse(resp)
}

// listClusterUserCredentialCreateRequest creates the ListClusterUserCredential request.
func (client *AppliancesClient) listClusterUserCredentialCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AppliancesClientListClusterUserCredentialOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances/{resourceName}/listClusterUserCredential"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listClusterUserCredentialHandleResponse handles the ListClusterUserCredential response.
func (client *AppliancesClient) listClusterUserCredentialHandleResponse(resp *http.Response) (AppliancesClientListClusterUserCredentialResponse, error) {
	result := AppliancesClientListClusterUserCredentialResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplianceListCredentialResults); err != nil {
		return AppliancesClientListClusterUserCredentialResponse{}, err
	}
	return result, nil
}

// ListOperations - Lists all available Appliances operations.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AppliancesClientListOperationsOptions contains the optional parameters for the AppliancesClient.ListOperations
// method.
func (client *AppliancesClient) ListOperations(options *AppliancesClientListOperationsOptions) *runtime.Pager[AppliancesClientListOperationsResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppliancesClientListOperationsResponse]{
		More: func(page AppliancesClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppliancesClientListOperationsResponse) (AppliancesClientListOperationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOperationsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppliancesClientListOperationsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppliancesClientListOperationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppliancesClientListOperationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOperationsHandleResponse(resp)
		},
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *AppliancesClient) listOperationsCreateRequest(ctx context.Context, options *AppliancesClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceConnector/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *AppliancesClient) listOperationsHandleResponse(resp *http.Response) (AppliancesClientListOperationsResponse, error) {
	result := AppliancesClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplianceOperationsList); err != nil {
		return AppliancesClientListOperationsResponse{}, err
	}
	return result, nil
}

// Update - Updates an Appliance with the specified Resource Name in the specified Resource Group and Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Appliances name.
// parameters - The updatable fields of an existing Appliance.
// options - AppliancesClientUpdateOptions contains the optional parameters for the AppliancesClient.Update method.
func (client *AppliancesClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableAppliance, options *AppliancesClientUpdateOptions) (AppliancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return AppliancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppliancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppliancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AppliancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableAppliance, options *AppliancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceConnector/appliances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AppliancesClient) updateHandleResponse(resp *http.Response) (AppliancesClientUpdateResponse, error) {
	result := AppliancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Appliance); err != nil {
		return AppliancesClientUpdateResponse{}, err
	}
	return result, nil
}
