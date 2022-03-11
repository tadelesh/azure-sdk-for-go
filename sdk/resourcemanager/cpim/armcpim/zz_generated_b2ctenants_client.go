//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcpim

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

// B2CTenantsClient contains the methods for the B2CTenants group.
// Don't use this type directly, use NewB2CTenantsClient() instead.
type B2CTenantsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewB2CTenantsClient creates a new instance of B2CTenantsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewB2CTenantsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *B2CTenantsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &B2CTenantsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Checks the availability and validity of a domain name for the tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - B2CTenantsClientCheckNameAvailabilityOptions contains the optional parameters for the B2CTenantsClient.CheckNameAvailability
// method.
func (client *B2CTenantsClient) CheckNameAvailability(ctx context.Context, options *B2CTenantsClientCheckNameAvailabilityOptions) (B2CTenantsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, options)
	if err != nil {
		return B2CTenantsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return B2CTenantsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return B2CTenantsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *B2CTenantsClient) checkNameAvailabilityCreateRequest(ctx context.Context, options *B2CTenantsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureActiveDirectory/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.CheckNameAvailabilityRequestBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.CheckNameAvailabilityRequestBody)
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *B2CTenantsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (B2CTenantsClientCheckNameAvailabilityResponse, error) {
	result := B2CTenantsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityResponse); err != nil {
		return B2CTenantsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Initiates an async request to create both the Azure AD B2C tenant and the corresponding Azure resource linked
// to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The initial domain name of the B2C tenant.
// options - B2CTenantsClientBeginCreateOptions contains the optional parameters for the B2CTenantsClient.BeginCreate method.
func (client *B2CTenantsClient) BeginCreate(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginCreateOptions) (B2CTenantsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return B2CTenantsClientCreatePollerResponse{}, err
	}
	result := B2CTenantsClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("B2CTenantsClient.Create", "location", resp, client.pl)
	if err != nil {
		return B2CTenantsClientCreatePollerResponse{}, err
	}
	result.Poller = &B2CTenantsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Initiates an async request to create both the Azure AD B2C tenant and the corresponding Azure resource linked
// to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *B2CTenantsClient) create(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, options)
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

// createCreateRequest creates the Create request.
func (client *B2CTenantsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/b2cDirectories/{resourceName}"
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
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.CreateTenantRequestBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.CreateTenantRequestBody)
	}
	return req, nil
}

// BeginDelete - Initiates an async operation to delete the Azure AD B2C tenant and Azure resource. The resource deletion
// can only happen as the last step in the tenant deletion process
// [https://aka.ms/deleteB2Ctenant].
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The initial domain name of the B2C tenant.
// options - B2CTenantsClientBeginDeleteOptions contains the optional parameters for the B2CTenantsClient.BeginDelete method.
func (client *B2CTenantsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginDeleteOptions) (B2CTenantsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return B2CTenantsClientDeletePollerResponse{}, err
	}
	result := B2CTenantsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("B2CTenantsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return B2CTenantsClientDeletePollerResponse{}, err
	}
	result.Poller = &B2CTenantsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Initiates an async operation to delete the Azure AD B2C tenant and Azure resource. The resource deletion can only
// happen as the last step in the tenant deletion process
// [https://aka.ms/deleteB2Ctenant].
// If the operation fails it returns an *azcore.ResponseError type.
func (client *B2CTenantsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
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
func (client *B2CTenantsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/b2cDirectories/{resourceName}"
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
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the Azure AD B2C tenant resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The initial domain name of the B2C tenant.
// options - B2CTenantsClientGetOptions contains the optional parameters for the B2CTenantsClient.Get method.
func (client *B2CTenantsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientGetOptions) (B2CTenantsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return B2CTenantsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return B2CTenantsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return B2CTenantsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *B2CTenantsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/b2cDirectories/{resourceName}"
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
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *B2CTenantsClient) getHandleResponse(resp *http.Response) (B2CTenantsClientGetResponse, error) {
	result := B2CTenantsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.B2CTenantResource); err != nil {
		return B2CTenantsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all the Azure AD B2C tenant resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - B2CTenantsClientListByResourceGroupOptions contains the optional parameters for the B2CTenantsClient.ListByResourceGroup
// method.
func (client *B2CTenantsClient) ListByResourceGroup(resourceGroupName string, options *B2CTenantsClientListByResourceGroupOptions) *B2CTenantsClientListByResourceGroupPager {
	return &B2CTenantsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *B2CTenantsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *B2CTenantsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/b2cDirectories"
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
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *B2CTenantsClient) listByResourceGroupHandleResponse(resp *http.Response) (B2CTenantsClientListByResourceGroupResponse, error) {
	result := B2CTenantsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.B2CTenantResourceList); err != nil {
		return B2CTenantsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get all the Azure AD B2C tenant resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - B2CTenantsClientListBySubscriptionOptions contains the optional parameters for the B2CTenantsClient.ListBySubscription
// method.
func (client *B2CTenantsClient) ListBySubscription(options *B2CTenantsClientListBySubscriptionOptions) *B2CTenantsClientListBySubscriptionPager {
	return &B2CTenantsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *B2CTenantsClient) listBySubscriptionCreateRequest(ctx context.Context, options *B2CTenantsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureActiveDirectory/b2cDirectories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *B2CTenantsClient) listBySubscriptionHandleResponse(resp *http.Response) (B2CTenantsClientListBySubscriptionResponse, error) {
	result := B2CTenantsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.B2CTenantResourceList); err != nil {
		return B2CTenantsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update the Azure AD B2C tenant resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The initial domain name of the B2C tenant.
// options - B2CTenantsClientUpdateOptions contains the optional parameters for the B2CTenantsClient.Update method.
func (client *B2CTenantsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientUpdateOptions) (B2CTenantsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return B2CTenantsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return B2CTenantsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return B2CTenantsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *B2CTenantsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *B2CTenantsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/b2cDirectories/{resourceName}"
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
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.UpdateTenantRequestBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.UpdateTenantRequestBody)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *B2CTenantsClient) updateHandleResponse(resp *http.Response) (B2CTenantsClientUpdateResponse, error) {
	result := B2CTenantsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.B2CTenantResource); err != nil {
		return B2CTenantsClientUpdateResponse{}, err
	}
	return result, nil
}
