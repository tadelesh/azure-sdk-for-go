//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

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

// AssignmentsClient contains the methods for the GuestConfigurationAssignments group.
// Don't use this type directly, use NewAssignmentsClient() instead.
type AssignmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAssignmentsClient creates a new instance of AssignmentsClient with the specified values.
// subscriptionID - Subscription ID which uniquely identify Microsoft Azure subscription. The subscription ID forms part of
// the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AssignmentsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AssignmentsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates an association between a VM and guest configuration
// If the operation fails it returns an *azcore.ResponseError type.
// guestConfigurationAssignmentName - Name of the guest configuration assignment.
// resourceGroupName - The resource group name.
// vmName - The name of the virtual machine.
// parameters - Parameters supplied to the create or update guest configuration assignment.
// options - AssignmentsClientCreateOrUpdateOptions contains the optional parameters for the AssignmentsClient.CreateOrUpdate
// method.
func (client *AssignmentsClient) CreateOrUpdate(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, vmName string, parameters Assignment, options *AssignmentsClientCreateOrUpdateOptions) (AssignmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, guestConfigurationAssignmentName, resourceGroupName, vmName, parameters, options)
	if err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssignmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, vmName string, parameters Assignment, options *AssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
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
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (AssignmentsClientCreateOrUpdateResponse, error) {
	result := AssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a guest configuration assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// guestConfigurationAssignmentName - Name of the guest configuration assignment
// vmName - The name of the virtual machine.
// options - AssignmentsClientDeleteOptions contains the optional parameters for the AssignmentsClient.Delete method.
func (client *AssignmentsClient) Delete(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, vmName string, options *AssignmentsClientDeleteOptions) (AssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, vmName, options)
	if err != nil {
		return AssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, vmName string, options *AssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
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
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get information about a guest configuration assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// guestConfigurationAssignmentName - The guest configuration assignment name.
// vmName - The name of the virtual machine.
// options - AssignmentsClientGetOptions contains the optional parameters for the AssignmentsClient.Get method.
func (client *AssignmentsClient) Get(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, vmName string, options *AssignmentsClientGetOptions) (AssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, vmName, options)
	if err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, vmName string, options *AssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
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
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssignmentsClient) getHandleResponse(resp *http.Response) (AssignmentsClientGetResponse, error) {
	result := AssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all guest configuration assignments for a virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// vmName - The name of the virtual machine.
// options - AssignmentsClientListOptions contains the optional parameters for the AssignmentsClient.List method.
func (client *AssignmentsClient) List(resourceGroupName string, vmName string, options *AssignmentsClientListOptions) *runtime.Pager[AssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AssignmentsClientListResponse]{
		More: func(page AssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AssignmentsClientListResponse) (AssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, vmName, options)
			if err != nil {
				return AssignmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *AssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssignmentsClient) listHandleResponse(resp *http.Response) (AssignmentsClientListResponse, error) {
	result := AssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return AssignmentsClientListResponse{}, err
	}
	return result, nil
}

// RGList - List all guest configuration assignments for a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// options - AssignmentsClientRGListOptions contains the optional parameters for the AssignmentsClient.RGList method.
func (client *AssignmentsClient) RGList(resourceGroupName string, options *AssignmentsClientRGListOptions) *runtime.Pager[AssignmentsClientRGListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AssignmentsClientRGListResponse]{
		More: func(page AssignmentsClientRGListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AssignmentsClientRGListResponse) (AssignmentsClientRGListResponse, error) {
			req, err := client.rgListCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return AssignmentsClientRGListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssignmentsClientRGListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssignmentsClientRGListResponse{}, runtime.NewResponseError(resp)
			}
			return client.rgListHandleResponse(resp)
		},
	})
}

// rgListCreateRequest creates the RGList request.
func (client *AssignmentsClient) rgListCreateRequest(ctx context.Context, resourceGroupName string, options *AssignmentsClientRGListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
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
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// rgListHandleResponse handles the RGList response.
func (client *AssignmentsClient) rgListHandleResponse(resp *http.Response) (AssignmentsClientRGListResponse, error) {
	result := AssignmentsClientRGListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return AssignmentsClientRGListResponse{}, err
	}
	return result, nil
}

// SubscriptionList - List all guest configuration assignments for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AssignmentsClientSubscriptionListOptions contains the optional parameters for the AssignmentsClient.SubscriptionList
// method.
func (client *AssignmentsClient) SubscriptionList(options *AssignmentsClientSubscriptionListOptions) *runtime.Pager[AssignmentsClientSubscriptionListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AssignmentsClientSubscriptionListResponse]{
		More: func(page AssignmentsClientSubscriptionListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AssignmentsClientSubscriptionListResponse) (AssignmentsClientSubscriptionListResponse, error) {
			req, err := client.subscriptionListCreateRequest(ctx, options)
			if err != nil {
				return AssignmentsClientSubscriptionListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssignmentsClientSubscriptionListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssignmentsClientSubscriptionListResponse{}, runtime.NewResponseError(resp)
			}
			return client.subscriptionListHandleResponse(resp)
		},
	})
}

// subscriptionListCreateRequest creates the SubscriptionList request.
func (client *AssignmentsClient) subscriptionListCreateRequest(ctx context.Context, options *AssignmentsClientSubscriptionListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// subscriptionListHandleResponse handles the SubscriptionList response.
func (client *AssignmentsClient) subscriptionListHandleResponse(resp *http.Response) (AssignmentsClientSubscriptionListResponse, error) {
	result := AssignmentsClientSubscriptionListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return AssignmentsClientSubscriptionListResponse{}, err
	}
	return result, nil
}
