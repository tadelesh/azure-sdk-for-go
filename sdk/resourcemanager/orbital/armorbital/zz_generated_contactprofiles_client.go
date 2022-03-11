//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

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

// ContactProfilesClient contains the methods for the ContactProfiles group.
// Don't use this type directly, use NewContactProfilesClient() instead.
type ContactProfilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContactProfilesClient creates a new instance of ContactProfilesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContactProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContactProfilesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ContactProfilesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a contact profile
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// contactProfileName - Contact Profile Name
// parameters - The parameters to provide for the created Contact Profile.
// options - ContactProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the ContactProfilesClient.BeginCreateOrUpdate
// method.
func (client *ContactProfilesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, contactProfileName string, parameters ContactProfile, options *ContactProfilesClientBeginCreateOrUpdateOptions) (ContactProfilesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, contactProfileName, parameters, options)
	if err != nil {
		return ContactProfilesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ContactProfilesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ContactProfilesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ContactProfilesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ContactProfilesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a contact profile
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ContactProfilesClient) createOrUpdate(ctx context.Context, resourceGroupName string, contactProfileName string, parameters ContactProfile, options *ContactProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, contactProfileName, parameters, options)
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
func (client *ContactProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, contactProfileName string, parameters ContactProfile, options *ContactProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles/{contactProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if contactProfileName == "" {
		return nil, errors.New("parameter contactProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactProfileName}", url.PathEscape(contactProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a specified contact profile resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// contactProfileName - Contact Profile Name
// options - ContactProfilesClientBeginDeleteOptions contains the optional parameters for the ContactProfilesClient.BeginDelete
// method.
func (client *ContactProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, contactProfileName string, options *ContactProfilesClientBeginDeleteOptions) (ContactProfilesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, contactProfileName, options)
	if err != nil {
		return ContactProfilesClientDeletePollerResponse{}, err
	}
	result := ContactProfilesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ContactProfilesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return ContactProfilesClientDeletePollerResponse{}, err
	}
	result.Poller = &ContactProfilesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a specified contact profile resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ContactProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, contactProfileName string, options *ContactProfilesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, contactProfileName, options)
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
func (client *ContactProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, contactProfileName string, options *ContactProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles/{contactProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if contactProfileName == "" {
		return nil, errors.New("parameter contactProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactProfileName}", url.PathEscape(contactProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified contact Profile in a specified resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// contactProfileName - Contact Profile Name
// options - ContactProfilesClientGetOptions contains the optional parameters for the ContactProfilesClient.Get method.
func (client *ContactProfilesClient) Get(ctx context.Context, resourceGroupName string, contactProfileName string, options *ContactProfilesClientGetOptions) (ContactProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, contactProfileName, options)
	if err != nil {
		return ContactProfilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContactProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContactProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContactProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, contactProfileName string, options *ContactProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles/{contactProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if contactProfileName == "" {
		return nil, errors.New("parameter contactProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactProfileName}", url.PathEscape(contactProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContactProfilesClient) getHandleResponse(resp *http.Response) (ContactProfilesClientGetResponse, error) {
	result := ContactProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactProfile); err != nil {
		return ContactProfilesClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns list of contact profiles
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ContactProfilesClientListOptions contains the optional parameters for the ContactProfilesClient.List method.
func (client *ContactProfilesClient) List(resourceGroupName string, options *ContactProfilesClientListOptions) *ContactProfilesClientListPager {
	return &ContactProfilesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ContactProfilesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ContactProfilesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles"
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
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContactProfilesClient) listHandleResponse(resp *http.Response) (ContactProfilesClientListResponse, error) {
	result := ContactProfilesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactProfileListResult); err != nil {
		return ContactProfilesClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Returns list of contact profiles
// If the operation fails it returns an *azcore.ResponseError type.
// options - ContactProfilesClientListBySubscriptionOptions contains the optional parameters for the ContactProfilesClient.ListBySubscription
// method.
func (client *ContactProfilesClient) ListBySubscription(options *ContactProfilesClientListBySubscriptionOptions) *ContactProfilesClientListBySubscriptionPager {
	return &ContactProfilesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ContactProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ContactProfilesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Orbital/contactProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ContactProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (ContactProfilesClientListBySubscriptionResponse, error) {
	result := ContactProfilesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactProfileListResult); err != nil {
		return ContactProfilesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates the specified contact profile tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// contactProfileName - Contact Profile Name
// parameters - Parameters supplied to update contact profile tags.
// options - ContactProfilesClientUpdateTagsOptions contains the optional parameters for the ContactProfilesClient.UpdateTags
// method.
func (client *ContactProfilesClient) UpdateTags(ctx context.Context, resourceGroupName string, contactProfileName string, parameters TagsObject, options *ContactProfilesClientUpdateTagsOptions) (ContactProfilesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, contactProfileName, parameters, options)
	if err != nil {
		return ContactProfilesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContactProfilesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContactProfilesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ContactProfilesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, contactProfileName string, parameters TagsObject, options *ContactProfilesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles/{contactProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if contactProfileName == "" {
		return nil, errors.New("parameter contactProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactProfileName}", url.PathEscape(contactProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ContactProfilesClient) updateTagsHandleResponse(resp *http.Response) (ContactProfilesClientUpdateTagsResponse, error) {
	result := ContactProfilesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactProfile); err != nil {
		return ContactProfilesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
