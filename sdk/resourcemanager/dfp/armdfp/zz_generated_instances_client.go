//go:build go1.16
// +build go1.16

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

package armdfp

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

// InstancesClient contains the methods for the Instances group.
// Don't use this type directly, use NewInstancesClient() instead.
type InstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInstancesClient creates a new instance of InstancesClient with the specified values.
// subscriptionID - A unique identifier for a Microsoft Azure subscription. The subscription ID forms part of the URI for
// every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *InstancesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &InstancesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Check the name availability in the target location.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The region name which the operation will lookup into.
// instanceParameters - The name of the instance.
// options - InstancesClientCheckNameAvailabilityOptions contains the optional parameters for the InstancesClient.CheckNameAvailability
// method.
func (client *InstancesClient) CheckNameAvailability(ctx context.Context, location string, instanceParameters CheckInstanceNameAvailabilityParameters, options *InstancesClientCheckNameAvailabilityOptions) (InstancesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, instanceParameters, options)
	if err != nil {
		return InstancesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstancesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstancesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *InstancesClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, instanceParameters CheckInstanceNameAvailabilityParameters, options *InstancesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Dynamics365FraudProtection/locations/{location}/checkNameAvailability"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, instanceParameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *InstancesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (InstancesClientCheckNameAvailabilityResponse, error) {
	result := InstancesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckInstanceNameAvailabilityResult); err != nil {
		return InstancesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Provisions the specified DFP instance based on the configuration specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure Resource group of which a given DFP instance is part. This name must be at least
// 1 character in length, and no more than 90.
// instanceName - The name of the DFP instances. It must be a minimum of 3 characters, and a maximum of 63.
// instanceParameters - Contains the information used to provision the DFP instance.
// options - InstancesClientBeginCreateOptions contains the optional parameters for the InstancesClient.BeginCreate method.
func (client *InstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, instanceName string, instanceParameters Instance, options *InstancesClientBeginCreateOptions) (InstancesClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, instanceName, instanceParameters, options)
	if err != nil {
		return InstancesClientCreatePollerResponse{}, err
	}
	result := InstancesClientCreatePollerResponse{}
	pt, err := armruntime.NewPoller("InstancesClient.Create", "", resp, client.pl)
	if err != nil {
		return InstancesClientCreatePollerResponse{}, err
	}
	result.Poller = &InstancesClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Provisions the specified DFP instance based on the configuration specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstancesClient) create(ctx context.Context, resourceGroupName string, instanceName string, instanceParameters Instance, options *InstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, instanceName, instanceParameters, options)
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
func (client *InstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, instanceParameters Instance, options *InstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Dynamics365FraudProtection/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, instanceParameters)
}

// BeginDelete - Deletes the specified DFP instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure Resource group of which a given DFP instance is part. This name must be at least
// 1 character in length, and no more than 90.
// instanceName - The name of the DFP instance. It must be at least 3 characters in length, and no more than 63.
// options - InstancesClientBeginDeleteOptions contains the optional parameters for the InstancesClient.BeginDelete method.
func (client *InstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, instanceName string, options *InstancesClientBeginDeleteOptions) (InstancesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, instanceName, options)
	if err != nil {
		return InstancesClientDeletePollerResponse{}, err
	}
	result := InstancesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("InstancesClient.Delete", "", resp, client.pl)
	if err != nil {
		return InstancesClientDeletePollerResponse{}, err
	}
	result.Poller = &InstancesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified DFP instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, instanceName, options)
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
func (client *InstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Dynamics365FraudProtection/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetDetails - Gets details about the specified instances.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure Resource group of which a given DFP instance is part. This name must be at least
// 1 character in length, and no more than 90.
// instanceName - The name of the instance. It must be a minimum of 3 characters, and a maximum of 63.
// options - InstancesClientGetDetailsOptions contains the optional parameters for the InstancesClient.GetDetails method.
func (client *InstancesClient) GetDetails(ctx context.Context, resourceGroupName string, instanceName string, options *InstancesClientGetDetailsOptions) (InstancesClientGetDetailsResponse, error) {
	req, err := client.getDetailsCreateRequest(ctx, resourceGroupName, instanceName, options)
	if err != nil {
		return InstancesClientGetDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstancesClientGetDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstancesClientGetDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDetailsHandleResponse(resp)
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *InstancesClient) getDetailsCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, options *InstancesClientGetDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Dynamics365FraudProtection/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDetailsHandleResponse handles the GetDetails response.
func (client *InstancesClient) getDetailsHandleResponse(resp *http.Response) (InstancesClientGetDetailsResponse, error) {
	result := InstancesClientGetDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instance); err != nil {
		return InstancesClientGetDetailsResponse{}, err
	}
	return result, nil
}

// List - Lists all the Dedicated instances for the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - InstancesClientListOptions contains the optional parameters for the InstancesClient.List method.
func (client *InstancesClient) List(options *InstancesClientListOptions) *InstancesClientListPager {
	return &InstancesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp InstancesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Instances.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *InstancesClient) listCreateRequest(ctx context.Context, options *InstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Dynamics365FraudProtection/instances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InstancesClient) listHandleResponse(resp *http.Response) (InstancesClientListResponse, error) {
	result := InstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instances); err != nil {
		return InstancesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all the Dedicated instance for the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure Resource group of which a given DFP instance is part. This name must be at least
// 1 character in length, and no more than 90.
// options - InstancesClientListByResourceGroupOptions contains the optional parameters for the InstancesClient.ListByResourceGroup
// method.
func (client *InstancesClient) ListByResourceGroup(resourceGroupName string, options *InstancesClientListByResourceGroupOptions) *InstancesClientListByResourceGroupPager {
	return &InstancesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp InstancesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Instances.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *InstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *InstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Dynamics365FraudProtection/instances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *InstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (InstancesClientListByResourceGroupResponse, error) {
	result := InstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instances); err != nil {
		return InstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the current state of the specified DFP instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure Resource group of which a given DFP instance is part. This name must be at least
// 1 character in length, and no more than 90.
// instanceName - The name of the DFP instance. It must be at least 3 characters in length, and no more than 63.
// instanceUpdateParameters - Request object that contains the updated information for the instance.
// options - InstancesClientBeginUpdateOptions contains the optional parameters for the InstancesClient.BeginUpdate method.
func (client *InstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, instanceName string, instanceUpdateParameters InstanceUpdateParameters, options *InstancesClientBeginUpdateOptions) (InstancesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, instanceName, instanceUpdateParameters, options)
	if err != nil {
		return InstancesClientUpdatePollerResponse{}, err
	}
	result := InstancesClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("InstancesClient.Update", "", resp, client.pl)
	if err != nil {
		return InstancesClientUpdatePollerResponse{}, err
	}
	result.Poller = &InstancesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the current state of the specified DFP instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstancesClient) update(ctx context.Context, resourceGroupName string, instanceName string, instanceUpdateParameters InstanceUpdateParameters, options *InstancesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, instanceName, instanceUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *InstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, instanceUpdateParameters InstanceUpdateParameters, options *InstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Dynamics365FraudProtection/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, instanceUpdateParameters)
}
