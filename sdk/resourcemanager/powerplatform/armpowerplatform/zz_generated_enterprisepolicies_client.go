//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerplatform

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

// EnterprisePoliciesClient contains the methods for the EnterprisePolicies group.
// Don't use this type directly, use NewEnterprisePoliciesClient() instead.
type EnterprisePoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEnterprisePoliciesClient creates a new instance of EnterprisePoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEnterprisePoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EnterprisePoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &EnterprisePoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates an EnterprisePolicy
// If the operation fails it returns an *azcore.ResponseError type.
// enterprisePolicyName - Name of the EnterprisePolicy.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// parameters - Parameters supplied to create or update EnterprisePolicy.
// options - EnterprisePoliciesClientCreateOrUpdateOptions contains the optional parameters for the EnterprisePoliciesClient.CreateOrUpdate
// method.
func (client *EnterprisePoliciesClient) CreateOrUpdate(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters EnterprisePolicy, options *EnterprisePoliciesClientCreateOrUpdateOptions) (EnterprisePoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, enterprisePolicyName, resourceGroupName, parameters, options)
	if err != nil {
		return EnterprisePoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnterprisePoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EnterprisePoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EnterprisePoliciesClient) createOrUpdateCreateRequest(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters EnterprisePolicy, options *EnterprisePoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}"
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EnterprisePoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (EnterprisePoliciesClientCreateOrUpdateResponse, error) {
	result := EnterprisePoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePolicy); err != nil {
		return EnterprisePoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an EnterprisePolicy
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// enterprisePolicyName - Name of the EnterprisePolicy
// options - EnterprisePoliciesClientDeleteOptions contains the optional parameters for the EnterprisePoliciesClient.Delete
// method.
func (client *EnterprisePoliciesClient) Delete(ctx context.Context, resourceGroupName string, enterprisePolicyName string, options *EnterprisePoliciesClientDeleteOptions) (EnterprisePoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, enterprisePolicyName, options)
	if err != nil {
		return EnterprisePoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnterprisePoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return EnterprisePoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return EnterprisePoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnterprisePoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, enterprisePolicyName string, options *EnterprisePoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get information about an EnterprisePolicy
// If the operation fails it returns an *azcore.ResponseError type.
// enterprisePolicyName - The EnterprisePolicy name.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - EnterprisePoliciesClientGetOptions contains the optional parameters for the EnterprisePoliciesClient.Get method.
func (client *EnterprisePoliciesClient) Get(ctx context.Context, enterprisePolicyName string, resourceGroupName string, options *EnterprisePoliciesClientGetOptions) (EnterprisePoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, enterprisePolicyName, resourceGroupName, options)
	if err != nil {
		return EnterprisePoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnterprisePoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnterprisePoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnterprisePoliciesClient) getCreateRequest(ctx context.Context, enterprisePolicyName string, resourceGroupName string, options *EnterprisePoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}"
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
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
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnterprisePoliciesClient) getHandleResponse(resp *http.Response) (EnterprisePoliciesClientGetResponse, error) {
	result := EnterprisePoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePolicy); err != nil {
		return EnterprisePoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieve a list of EnterprisePolicies within a given resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - EnterprisePoliciesClientListByResourceGroupOptions contains the optional parameters for the EnterprisePoliciesClient.ListByResourceGroup
// method.
func (client *EnterprisePoliciesClient) ListByResourceGroup(resourceGroupName string, options *EnterprisePoliciesClientListByResourceGroupOptions) *EnterprisePoliciesClientListByResourceGroupPager {
	return &EnterprisePoliciesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp EnterprisePoliciesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EnterprisePolicyList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *EnterprisePoliciesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *EnterprisePoliciesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies"
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
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *EnterprisePoliciesClient) listByResourceGroupHandleResponse(resp *http.Response) (EnterprisePoliciesClientListByResourceGroupResponse, error) {
	result := EnterprisePoliciesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePolicyList); err != nil {
		return EnterprisePoliciesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieve a list of EnterprisePolicies within a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - EnterprisePoliciesClientListBySubscriptionOptions contains the optional parameters for the EnterprisePoliciesClient.ListBySubscription
// method.
func (client *EnterprisePoliciesClient) ListBySubscription(options *EnterprisePoliciesClientListBySubscriptionOptions) *EnterprisePoliciesClientListBySubscriptionPager {
	return &EnterprisePoliciesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp EnterprisePoliciesClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EnterprisePolicyList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *EnterprisePoliciesClient) listBySubscriptionCreateRequest(ctx context.Context, options *EnterprisePoliciesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerPlatform/enterprisePolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *EnterprisePoliciesClient) listBySubscriptionHandleResponse(resp *http.Response) (EnterprisePoliciesClientListBySubscriptionResponse, error) {
	result := EnterprisePoliciesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePolicyList); err != nil {
		return EnterprisePoliciesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates an EnterprisePolicy
// If the operation fails it returns an *azcore.ResponseError type.
// enterprisePolicyName - Name of the EnterprisePolicy.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// parameters - Parameters supplied to update EnterprisePolicy.
// options - EnterprisePoliciesClientUpdateOptions contains the optional parameters for the EnterprisePoliciesClient.Update
// method.
func (client *EnterprisePoliciesClient) Update(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters PatchEnterprisePolicy, options *EnterprisePoliciesClientUpdateOptions) (EnterprisePoliciesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, enterprisePolicyName, resourceGroupName, parameters, options)
	if err != nil {
		return EnterprisePoliciesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnterprisePoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnterprisePoliciesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EnterprisePoliciesClient) updateCreateRequest(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters PatchEnterprisePolicy, options *EnterprisePoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}"
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *EnterprisePoliciesClient) updateHandleResponse(resp *http.Response) (EnterprisePoliciesClientUpdateResponse, error) {
	result := EnterprisePoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePolicy); err != nil {
		return EnterprisePoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
