//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextendedlocation

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

// CustomLocationsClient contains the methods for the CustomLocations group.
// Don't use this type directly, use NewCustomLocationsClient() instead.
type CustomLocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCustomLocationsClient creates a new instance of CustomLocationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCustomLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CustomLocationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CustomLocationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Custom Locations name.
// parameters - Parameters supplied to create or update a Custom Location.
// options - CustomLocationsClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomLocationsClient.BeginCreateOrUpdate
// method.
func (client *CustomLocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (CustomLocationsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsClientCreateOrUpdatePollerResponse{}, err
	}
	result := CustomLocationsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("CustomLocationsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return CustomLocationsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomLocationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomLocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
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
func (client *CustomLocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Custom Locations name.
// options - CustomLocationsClientBeginDeleteOptions contains the optional parameters for the CustomLocationsClient.BeginDelete
// method.
func (client *CustomLocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (CustomLocationsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CustomLocationsClientDeletePollerResponse{}, err
	}
	result := CustomLocationsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("CustomLocationsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return CustomLocationsClientDeletePollerResponse{}, err
	}
	result.Poller = &CustomLocationsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CustomLocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (*http.Response, error) {
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
func (client *CustomLocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the details of the customLocation with a specified resource group and name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Custom Locations name.
// options - CustomLocationsClientGetOptions contains the optional parameters for the CustomLocationsClient.Get method.
func (client *CustomLocationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientGetOptions) (CustomLocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomLocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomLocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomLocationsClient) getHandleResponse(resp *http.Response) (CustomLocationsClientGetResponse, error) {
	result := CustomLocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets a list of Custom Locations in the specified subscription and resource group. The operation returns
// properties of each Custom Location.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - CustomLocationsClientListByResourceGroupOptions contains the optional parameters for the CustomLocationsClient.ListByResourceGroup
// method.
func (client *CustomLocationsClient) ListByResourceGroup(resourceGroupName string, options *CustomLocationsClientListByResourceGroupOptions) *CustomLocationsClientListByResourceGroupPager {
	return &CustomLocationsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomLocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomLocationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomLocationsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomLocationsClientListByResourceGroupResponse, error) {
	result := CustomLocationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets a list of Custom Locations in the specified subscription. The operation returns properties of
// each Custom Location
// If the operation fails it returns an *azcore.ResponseError type.
// options - CustomLocationsClientListBySubscriptionOptions contains the optional parameters for the CustomLocationsClient.ListBySubscription
// method.
func (client *CustomLocationsClient) ListBySubscription(options *CustomLocationsClientListBySubscriptionOptions) *CustomLocationsClientListBySubscriptionPager {
	return &CustomLocationsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomLocationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomLocationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ExtendedLocation/customLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomLocationsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomLocationsClientListBySubscriptionResponse, error) {
	result := CustomLocationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListEnabledResourceTypes - Gets the list of the Enabled Resource Types.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Custom Locations name.
// options - CustomLocationsClientListEnabledResourceTypesOptions contains the optional parameters for the CustomLocationsClient.ListEnabledResourceTypes
// method.
func (client *CustomLocationsClient) ListEnabledResourceTypes(resourceGroupName string, resourceName string, options *CustomLocationsClientListEnabledResourceTypesOptions) *CustomLocationsClientListEnabledResourceTypesPager {
	return &CustomLocationsClientListEnabledResourceTypesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listEnabledResourceTypesCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsClientListEnabledResourceTypesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EnabledResourceTypesListResult.NextLink)
		},
	}
}

// listEnabledResourceTypesCreateRequest creates the ListEnabledResourceTypes request.
func (client *CustomLocationsClient) listEnabledResourceTypesCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientListEnabledResourceTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/enabledResourceTypes"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listEnabledResourceTypesHandleResponse handles the ListEnabledResourceTypes response.
func (client *CustomLocationsClient) listEnabledResourceTypesHandleResponse(resp *http.Response) (CustomLocationsClientListEnabledResourceTypesResponse, error) {
	result := CustomLocationsClientListEnabledResourceTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnabledResourceTypesListResult); err != nil {
		return CustomLocationsClientListEnabledResourceTypesResponse{}, err
	}
	return result, nil
}

// ListOperations - Lists all available Custom Locations operations.
// If the operation fails it returns an *azcore.ResponseError type.
// options - CustomLocationsClientListOperationsOptions contains the optional parameters for the CustomLocationsClient.ListOperations
// method.
func (client *CustomLocationsClient) ListOperations(options *CustomLocationsClientListOperationsOptions) *CustomLocationsClientListOperationsPager {
	return &CustomLocationsClientListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsClientListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationOperationsList.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *CustomLocationsClient) listOperationsCreateRequest(ctx context.Context, options *CustomLocationsClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ExtendedLocation/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *CustomLocationsClient) listOperationsHandleResponse(resp *http.Response) (CustomLocationsClientListOperationsResponse, error) {
	result := CustomLocationsClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationOperationsList); err != nil {
		return CustomLocationsClientListOperationsResponse{}, err
	}
	return result, nil
}

// Update - Updates a Custom Location with the specified Resource Name in the specified Resource Group and Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - Custom Locations name.
// parameters - The updatable fields of an existing Custom Location.
// options - CustomLocationsClientUpdateOptions contains the optional parameters for the CustomLocationsClient.Update method.
func (client *CustomLocationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsClientUpdateOptions) (CustomLocationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomLocationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CustomLocationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CustomLocationsClient) updateHandleResponse(resp *http.Response) (CustomLocationsClientUpdateResponse, error) {
	result := CustomLocationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	return result, nil
}
