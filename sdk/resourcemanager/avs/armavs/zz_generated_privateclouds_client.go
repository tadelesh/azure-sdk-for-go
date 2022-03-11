//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// PrivateCloudsClient contains the methods for the PrivateClouds group.
// Don't use this type directly, use NewPrivateCloudsClient() instead.
type PrivateCloudsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateCloudsClient creates a new instance of PrivateCloudsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateCloudsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateCloudsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateCloudsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// privateCloud - The private cloud
// options - PrivateCloudsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateCloudsClient.BeginCreateOrUpdate
// method.
func (client *PrivateCloudsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (PrivateCloudsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, privateCloud, options)
	if err != nil {
		return PrivateCloudsClientCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateCloudsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("PrivateCloudsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return PrivateCloudsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateCloudsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateCloudsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloud, options)
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
func (client *PrivateCloudsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateCloud)
}

// BeginDelete - Delete a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - PrivateCloudsClientBeginDeleteOptions contains the optional parameters for the PrivateCloudsClient.BeginDelete
// method.
func (client *PrivateCloudsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (PrivateCloudsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientDeletePollerResponse{}, err
	}
	result := PrivateCloudsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("PrivateCloudsClient.Delete", "", resp, client.pl)
	if err != nil {
		return PrivateCloudsClientDeletePollerResponse{}, err
	}
	result.Poller = &PrivateCloudsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateCloudsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, options)
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
func (client *PrivateCloudsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - PrivateCloudsClientGetOptions contains the optional parameters for the PrivateCloudsClient.Get method.
func (client *PrivateCloudsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientGetOptions) (PrivateCloudsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateCloudsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateCloudsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateCloudsClient) getHandleResponse(resp *http.Response) (PrivateCloudsClientGetResponse, error) {
	result := PrivateCloudsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloud); err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	return result, nil
}

// List - List private clouds in a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PrivateCloudsClientListOptions contains the optional parameters for the PrivateCloudsClient.List method.
func (client *PrivateCloudsClient) List(resourceGroupName string, options *PrivateCloudsClientListOptions) *PrivateCloudsClientListPager {
	return &PrivateCloudsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PrivateCloudsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateCloudList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateCloudsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateCloudsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds"
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
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateCloudsClient) listHandleResponse(resp *http.Response) (PrivateCloudsClientListResponse, error) {
	result := PrivateCloudsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloudList); err != nil {
		return PrivateCloudsClientListResponse{}, err
	}
	return result, nil
}

// ListAdminCredentials - List the admin credentials for the private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - PrivateCloudsClientListAdminCredentialsOptions contains the optional parameters for the PrivateCloudsClient.ListAdminCredentials
// method.
func (client *PrivateCloudsClient) ListAdminCredentials(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientListAdminCredentialsOptions) (PrivateCloudsClientListAdminCredentialsResponse, error) {
	req, err := client.listAdminCredentialsCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateCloudsClientListAdminCredentialsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAdminCredentialsHandleResponse(resp)
}

// listAdminCredentialsCreateRequest creates the ListAdminCredentials request.
func (client *PrivateCloudsClient) listAdminCredentialsCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientListAdminCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/listAdminCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAdminCredentialsHandleResponse handles the ListAdminCredentials response.
func (client *PrivateCloudsClient) listAdminCredentialsHandleResponse(resp *http.Response) (PrivateCloudsClientListAdminCredentialsResponse, error) {
	result := PrivateCloudsClientListAdminCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminCredentials); err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	return result, nil
}

// ListInSubscription - List private clouds in a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateCloudsClientListInSubscriptionOptions contains the optional parameters for the PrivateCloudsClient.ListInSubscription
// method.
func (client *PrivateCloudsClient) ListInSubscription(options *PrivateCloudsClientListInSubscriptionOptions) *PrivateCloudsClientListInSubscriptionPager {
	return &PrivateCloudsClientListInSubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listInSubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PrivateCloudsClientListInSubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateCloudList.NextLink)
		},
	}
}

// listInSubscriptionCreateRequest creates the ListInSubscription request.
func (client *PrivateCloudsClient) listInSubscriptionCreateRequest(ctx context.Context, options *PrivateCloudsClientListInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AVS/privateClouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listInSubscriptionHandleResponse handles the ListInSubscription response.
func (client *PrivateCloudsClient) listInSubscriptionHandleResponse(resp *http.Response) (PrivateCloudsClientListInSubscriptionResponse, error) {
	result := PrivateCloudsClientListInSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloudList); err != nil {
		return PrivateCloudsClientListInSubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRotateNsxtPassword - Rotate the NSX-T Manager password
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - PrivateCloudsClientBeginRotateNsxtPasswordOptions contains the optional parameters for the PrivateCloudsClient.BeginRotateNsxtPassword
// method.
func (client *PrivateCloudsClient) BeginRotateNsxtPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (PrivateCloudsClientRotateNsxtPasswordPollerResponse, error) {
	resp, err := client.rotateNsxtPassword(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientRotateNsxtPasswordPollerResponse{}, err
	}
	result := PrivateCloudsClientRotateNsxtPasswordPollerResponse{}
	pt, err := armruntime.NewPoller("PrivateCloudsClient.RotateNsxtPassword", "", resp, client.pl)
	if err != nil {
		return PrivateCloudsClientRotateNsxtPasswordPollerResponse{}, err
	}
	result.Poller = &PrivateCloudsClientRotateNsxtPasswordPoller{
		pt: pt,
	}
	return result, nil
}

// RotateNsxtPassword - Rotate the NSX-T Manager password
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateCloudsClient) rotateNsxtPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (*http.Response, error) {
	req, err := client.rotateNsxtPasswordCreateRequest(ctx, resourceGroupName, privateCloudName, options)
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

// rotateNsxtPasswordCreateRequest creates the RotateNsxtPassword request.
func (client *PrivateCloudsClient) rotateNsxtPasswordCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/rotateNsxtPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginRotateVcenterPassword - Rotate the vCenter password
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - PrivateCloudsClientBeginRotateVcenterPasswordOptions contains the optional parameters for the PrivateCloudsClient.BeginRotateVcenterPassword
// method.
func (client *PrivateCloudsClient) BeginRotateVcenterPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (PrivateCloudsClientRotateVcenterPasswordPollerResponse, error) {
	resp, err := client.rotateVcenterPassword(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientRotateVcenterPasswordPollerResponse{}, err
	}
	result := PrivateCloudsClientRotateVcenterPasswordPollerResponse{}
	pt, err := armruntime.NewPoller("PrivateCloudsClient.RotateVcenterPassword", "", resp, client.pl)
	if err != nil {
		return PrivateCloudsClientRotateVcenterPasswordPollerResponse{}, err
	}
	result.Poller = &PrivateCloudsClientRotateVcenterPasswordPoller{
		pt: pt,
	}
	return result, nil
}

// RotateVcenterPassword - Rotate the vCenter password
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateCloudsClient) rotateVcenterPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (*http.Response, error) {
	req, err := client.rotateVcenterPasswordCreateRequest(ctx, resourceGroupName, privateCloudName, options)
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

// rotateVcenterPasswordCreateRequest creates the RotateVcenterPassword request.
func (client *PrivateCloudsClient) rotateVcenterPasswordCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/rotateVcenterPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// privateCloudUpdate - The private cloud properties to be updated
// options - PrivateCloudsClientBeginUpdateOptions contains the optional parameters for the PrivateCloudsClient.BeginUpdate
// method.
func (client *PrivateCloudsClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (PrivateCloudsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
	if err != nil {
		return PrivateCloudsClientUpdatePollerResponse{}, err
	}
	result := PrivateCloudsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("PrivateCloudsClient.Update", "", resp, client.pl)
	if err != nil {
		return PrivateCloudsClientUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateCloudsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateCloudsClient) update(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
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

// updateCreateRequest creates the Update request.
func (client *PrivateCloudsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateCloudUpdate)
}
