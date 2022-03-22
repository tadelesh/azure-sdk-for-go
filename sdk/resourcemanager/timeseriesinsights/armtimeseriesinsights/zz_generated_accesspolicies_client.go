//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtimeseriesinsights

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

// AccessPoliciesClient contains the methods for the AccessPolicies group.
// Don't use this type directly, use NewAccessPoliciesClient() instead.
type AccessPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessPoliciesClient creates a new instance of AccessPoliciesClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccessPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AccessPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update an access policy in the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// environmentName - The name of the Time Series Insights environment associated with the specified resource group.
// accessPolicyName - Name of the access policy.
// parameters - Parameters for creating an access policy.
// options - AccessPoliciesClientCreateOrUpdateOptions contains the optional parameters for the AccessPoliciesClient.CreateOrUpdate
// method.
func (client *AccessPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, parameters AccessPolicyCreateOrUpdateParameters, options *AccessPoliciesClientCreateOrUpdateOptions) (AccessPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, accessPolicyName, parameters, options)
	if err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AccessPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccessPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, parameters AccessPolicyCreateOrUpdateParameters, options *AccessPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AccessPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (AccessPoliciesClientCreateOrUpdateResponse, error) {
	result := AccessPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyResource); err != nil {
		return AccessPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the access policy with the specified name in the specified subscription, resource group, and environment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// environmentName - The name of the Time Series Insights environment associated with the specified resource group.
// accessPolicyName - The name of the Time Series Insights access policy associated with the specified environment.
// options - AccessPoliciesClientDeleteOptions contains the optional parameters for the AccessPoliciesClient.Delete method.
func (client *AccessPoliciesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, options *AccessPoliciesClientDeleteOptions) (AccessPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, accessPolicyName, options)
	if err != nil {
		return AccessPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccessPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AccessPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, options *AccessPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the access policy with the specified name in the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// environmentName - The name of the Time Series Insights environment associated with the specified resource group.
// accessPolicyName - The name of the Time Series Insights access policy associated with the specified environment.
// options - AccessPoliciesClientGetOptions contains the optional parameters for the AccessPoliciesClient.Get method.
func (client *AccessPoliciesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, options *AccessPoliciesClientGetOptions) (AccessPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, accessPolicyName, options)
	if err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccessPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, options *AccessPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessPoliciesClient) getHandleResponse(resp *http.Response) (AccessPoliciesClientGetResponse, error) {
	result := AccessPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyResource); err != nil {
		return AccessPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByEnvironment - Lists all the available access policies associated with the environment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// environmentName - The name of the Time Series Insights environment associated with the specified resource group.
// options - AccessPoliciesClientListByEnvironmentOptions contains the optional parameters for the AccessPoliciesClient.ListByEnvironment
// method.
func (client *AccessPoliciesClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string, options *AccessPoliciesClientListByEnvironmentOptions) (AccessPoliciesClientListByEnvironmentResponse, error) {
	req, err := client.listByEnvironmentCreateRequest(ctx, resourceGroupName, environmentName, options)
	if err != nil {
		return AccessPoliciesClientListByEnvironmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessPoliciesClientListByEnvironmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByEnvironmentHandleResponse(resp)
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *AccessPoliciesClient) listByEnvironmentCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *AccessPoliciesClientListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *AccessPoliciesClient) listByEnvironmentHandleResponse(resp *http.Response) (AccessPoliciesClientListByEnvironmentResponse, error) {
	result := AccessPoliciesClientListByEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyListResponse); err != nil {
		return AccessPoliciesClientListByEnvironmentResponse{}, err
	}
	return result, nil
}

// Update - Updates the access policy with the specified name in the specified subscription, resource group, and environment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// environmentName - The name of the Time Series Insights environment associated with the specified resource group.
// accessPolicyName - The name of the Time Series Insights access policy associated with the specified environment.
// accessPolicyUpdateParameters - Request object that contains the updated information for the access policy.
// options - AccessPoliciesClientUpdateOptions contains the optional parameters for the AccessPoliciesClient.Update method.
func (client *AccessPoliciesClient) Update(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, accessPolicyUpdateParameters AccessPolicyUpdateParameters, options *AccessPoliciesClientUpdateOptions) (AccessPoliciesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, environmentName, accessPolicyName, accessPolicyUpdateParameters, options)
	if err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessPoliciesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AccessPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, accessPolicyUpdateParameters AccessPolicyUpdateParameters, options *AccessPoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, accessPolicyUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *AccessPoliciesClient) updateHandleResponse(resp *http.Response) (AccessPoliciesClientUpdateResponse, error) {
	result := AccessPoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyResource); err != nil {
		return AccessPoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
