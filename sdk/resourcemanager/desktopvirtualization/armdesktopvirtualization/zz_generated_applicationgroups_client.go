//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// ApplicationGroupsClient contains the methods for the ApplicationGroups group.
// Don't use this type directly, use NewApplicationGroupsClient() instead.
type ApplicationGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationGroupsClient creates a new instance of ApplicationGroupsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationGroupsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ApplicationGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationGroupName - The name of the application group
// applicationGroup - Object containing ApplicationGroup definitions.
// options - ApplicationGroupsClientCreateOrUpdateOptions contains the optional parameters for the ApplicationGroupsClient.CreateOrUpdate
// method.
func (client *ApplicationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsClientCreateOrUpdateOptions) (ApplicationGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationGroupName, applicationGroup, options)
	if err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, applicationGroup)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationGroupsClientCreateOrUpdateResponse, error) {
	result := ApplicationGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationGroupName - The name of the application group
// options - ApplicationGroupsClientDeleteOptions contains the optional parameters for the ApplicationGroupsClient.Delete
// method.
func (client *ApplicationGroupsClient) Delete(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientDeleteOptions) (ApplicationGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApplicationGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get an application group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationGroupName - The name of the application group
// options - ApplicationGroupsClientGetOptions contains the optional parameters for the ApplicationGroupsClient.Get method.
func (client *ApplicationGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientGetOptions) (ApplicationGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGroupsClient) getHandleResponse(resp *http.Response) (ApplicationGroupsClientGetResponse, error) {
	result := ApplicationGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List applicationGroups.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ApplicationGroupsClientListByResourceGroupOptions contains the optional parameters for the ApplicationGroupsClient.ListByResourceGroup
// method.
func (client *ApplicationGroupsClient) ListByResourceGroup(resourceGroupName string, options *ApplicationGroupsClientListByResourceGroupOptions) *ApplicationGroupsClientListByResourceGroupPager {
	return &ApplicationGroupsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ApplicationGroupsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationGroupList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups"
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
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationGroupsClientListByResourceGroupResponse, error) {
	result := ApplicationGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - List applicationGroups in subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ApplicationGroupsClientListBySubscriptionOptions contains the optional parameters for the ApplicationGroupsClient.ListBySubscription
// method.
func (client *ApplicationGroupsClient) ListBySubscription(options *ApplicationGroupsClientListBySubscriptionOptions) *ApplicationGroupsClientListBySubscriptionPager {
	return &ApplicationGroupsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ApplicationGroupsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationGroupList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/applicationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationGroupsClientListBySubscriptionResponse, error) {
	result := ApplicationGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationGroupName - The name of the application group
// options - ApplicationGroupsClientUpdateOptions contains the optional parameters for the ApplicationGroupsClient.Update
// method.
func (client *ApplicationGroupsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientUpdateOptions) (ApplicationGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.ApplicationGroup != nil {
		return req, runtime.MarshalAsJSON(req, *options.ApplicationGroup)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ApplicationGroupsClient) updateHandleResponse(resp *http.Response) (ApplicationGroupsClientUpdateResponse, error) {
	result := ApplicationGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
