//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

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

// ConfigurationProfileAssignmentsClient contains the methods for the ConfigurationProfileAssignments group.
// Don't use this type directly, use NewConfigurationProfileAssignmentsClient() instead.
type ConfigurationProfileAssignmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationProfileAssignmentsClient creates a new instance of ConfigurationProfileAssignmentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationProfileAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationProfileAssignmentsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ConfigurationProfileAssignmentsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates an association between a VM and Automanage configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
// configurationProfileAssignmentName - Name of the configuration profile assignment. Only default is supported.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// vmName - The name of the virtual machine.
// parameters - Parameters supplied to the create or update configuration profile assignment.
// options - ConfigurationProfileAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.CreateOrUpdate
// method.
func (client *ConfigurationProfileAssignmentsClient) CreateOrUpdate(ctx context.Context, configurationProfileAssignmentName string, resourceGroupName string, vmName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileAssignmentsClientCreateOrUpdateOptions) (ConfigurationProfileAssignmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, configurationProfileAssignmentName, resourceGroupName, vmName, parameters, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationProfileAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, configurationProfileAssignmentName string, resourceGroupName string, vmName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationProfileAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientCreateOrUpdateResponse, error) {
	result := ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - Name of the configuration profile assignment
// vmName - The name of the virtual machine.
// options - ConfigurationProfileAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.Delete
// method.
func (client *ConfigurationProfileAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientDeleteOptions) (ConfigurationProfileAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, vmName, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConfigurationProfileAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationProfileAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get information about a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - The configuration profile assignment name.
// vmName - The name of the virtual machine.
// options - ConfigurationProfileAssignmentsClientGetOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.Get
// method.
func (client *ConfigurationProfileAssignmentsClient) Get(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientGetOptions) (ConfigurationProfileAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, vmName, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationProfileAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationProfileAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationProfileAssignmentsClient) getHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientGetResponse, error) {
	result := ConfigurationProfileAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Get list of configuration profile assignments
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ConfigurationProfileAssignmentsClientListOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.List
// method.
func (client *ConfigurationProfileAssignmentsClient) List(resourceGroupName string, options *ConfigurationProfileAssignmentsClientListOptions) *runtime.Pager[ConfigurationProfileAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConfigurationProfileAssignmentsClientListResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListResponse) (ConfigurationProfileAssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationProfileAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationProfileAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfileAssignments"
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
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationProfileAssignmentsClient) listHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListResponse, error) {
	result := ConfigurationProfileAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get list of configuration profile assignments under a given subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConfigurationProfileAssignmentsClientListBySubscriptionOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.ListBySubscription
// method.
func (client *ConfigurationProfileAssignmentsClient) ListBySubscription(options *ConfigurationProfileAssignmentsClientListBySubscriptionOptions) *runtime.Pager[ConfigurationProfileAssignmentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConfigurationProfileAssignmentsClientListBySubscriptionResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListBySubscriptionResponse) (ConfigurationProfileAssignmentsClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationProfileAssignmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationProfileAssignmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfileAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationProfileAssignmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListBySubscriptionResponse, error) {
	result := ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
