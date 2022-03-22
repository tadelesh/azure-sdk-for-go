//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// VirtualMachinesClient contains the methods for the VirtualMachines group.
// Don't use this type directly, use NewVirtualMachinesClient() instead.
type VirtualMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachinesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualMachinesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Returns the properties for a lab virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// options - VirtualMachinesClientGetOptions contains the optional parameters for the VirtualMachinesClient.Get method.
func (client *VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (VirtualMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, options)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachinesClient) getHandleResponse(resp *http.Response) (VirtualMachinesClientGetResponse, error) {
	result := VirtualMachinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachine); err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// ListByLab - Returns a list of all virtual machines for a lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - VirtualMachinesClientListByLabOptions contains the optional parameters for the VirtualMachinesClient.ListByLab
// method.
func (client *VirtualMachinesClient) ListByLab(resourceGroupName string, labName string, options *VirtualMachinesClientListByLabOptions) *runtime.Pager[VirtualMachinesClientListByLabResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualMachinesClientListByLabResponse]{
		More: func(page VirtualMachinesClientListByLabResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachinesClientListByLabResponse) (VirtualMachinesClientListByLabResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLabCreateRequest(ctx, resourceGroupName, labName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachinesClientListByLabResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachinesClientListByLabResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachinesClientListByLabResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLabHandleResponse(resp)
		},
	})
}

// listByLabCreateRequest creates the ListByLab request.
func (client *VirtualMachinesClient) listByLabCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *VirtualMachinesClientListByLabOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLabHandleResponse handles the ListByLab response.
func (client *VirtualMachinesClient) listByLabHandleResponse(resp *http.Response) (VirtualMachinesClientListByLabResponse, error) {
	result := VirtualMachinesClientListByLabResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedVirtualMachines); err != nil {
		return VirtualMachinesClientListByLabResponse{}, err
	}
	return result, nil
}

// BeginRedeploy - Action to redeploy a lab virtual machine to a different compute node. For troubleshooting connectivity.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// options - VirtualMachinesClientBeginRedeployOptions contains the optional parameters for the VirtualMachinesClient.BeginRedeploy
// method.
func (client *VirtualMachinesClient) BeginRedeploy(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginRedeployOptions) (*armruntime.Poller[VirtualMachinesClientRedeployResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.redeploy(ctx, resourceGroupName, labName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachinesClientRedeployResponse]("VirtualMachinesClient.Redeploy", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachinesClientRedeployResponse]("VirtualMachinesClient.Redeploy", options.ResumeToken, client.pl, nil)
	}
}

// Redeploy - Action to redeploy a lab virtual machine to a different compute node. For troubleshooting connectivity.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) redeploy(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginRedeployOptions) (*http.Response, error) {
	req, err := client.redeployCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, options)
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

// redeployCreateRequest creates the Redeploy request.
func (client *VirtualMachinesClient) redeployCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginRedeployOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}/redeploy"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginReimage - Re-image a lab virtual machine. The virtual machine will be deleted and recreated using the latest published
// snapshot of the reference environment of the lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// options - VirtualMachinesClientBeginReimageOptions contains the optional parameters for the VirtualMachinesClient.BeginReimage
// method.
func (client *VirtualMachinesClient) BeginReimage(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*armruntime.Poller[VirtualMachinesClientReimageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reimage(ctx, resourceGroupName, labName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachinesClientReimageResponse]("VirtualMachinesClient.Reimage", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachinesClientReimageResponse]("VirtualMachinesClient.Reimage", options.ResumeToken, client.pl, nil)
	}
}

// Reimage - Re-image a lab virtual machine. The virtual machine will be deleted and recreated using the latest published
// snapshot of the reference environment of the lab.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) reimage(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*http.Response, error) {
	req, err := client.reimageCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, options)
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

// reimageCreateRequest creates the Reimage request.
func (client *VirtualMachinesClient) reimageCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}/reimage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginResetPassword - Resets a lab virtual machine password.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// body - The request body.
// options - VirtualMachinesClientBeginResetPasswordOptions contains the optional parameters for the VirtualMachinesClient.BeginResetPassword
// method.
func (client *VirtualMachinesClient) BeginResetPassword(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, body ResetPasswordBody, options *VirtualMachinesClientBeginResetPasswordOptions) (*armruntime.Poller[VirtualMachinesClientResetPasswordResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resetPassword(ctx, resourceGroupName, labName, virtualMachineName, body, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachinesClientResetPasswordResponse]("VirtualMachinesClient.ResetPassword", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachinesClientResetPasswordResponse]("VirtualMachinesClient.ResetPassword", options.ResumeToken, client.pl, nil)
	}
}

// ResetPassword - Resets a lab virtual machine password.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) resetPassword(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, body ResetPasswordBody, options *VirtualMachinesClientBeginResetPasswordOptions) (*http.Response, error) {
	req, err := client.resetPasswordCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, body, options)
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

// resetPasswordCreateRequest creates the ResetPassword request.
func (client *VirtualMachinesClient) resetPasswordCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, body ResetPasswordBody, options *VirtualMachinesClientBeginResetPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}/resetPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginStart - Action to start a lab virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// options - VirtualMachinesClientBeginStartOptions contains the optional parameters for the VirtualMachinesClient.BeginStart
// method.
func (client *VirtualMachinesClient) BeginStart(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*armruntime.Poller[VirtualMachinesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, labName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachinesClientStartResponse]("VirtualMachinesClient.Start", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachinesClientStartResponse]("VirtualMachinesClient.Start", options.ResumeToken, client.pl, nil)
	}
}

// Start - Action to start a lab virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) start(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, options)
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

// startCreateRequest creates the Start request.
func (client *VirtualMachinesClient) startCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - Action to stop a lab virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// virtualMachineName - The ID of the virtual machine that uniquely identifies it within the containing lab. Used in resource
// URIs.
// options - VirtualMachinesClientBeginStopOptions contains the optional parameters for the VirtualMachinesClient.BeginStop
// method.
func (client *VirtualMachinesClient) BeginStop(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*armruntime.Poller[VirtualMachinesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, labName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachinesClientStopResponse]("VirtualMachinesClient.Stop", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachinesClientStopResponse]("VirtualMachinesClient.Stop", options.ResumeToken, client.pl, nil)
	}
}

// Stop - Action to stop a lab virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) stop(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, labName, virtualMachineName, options)
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

// stopCreateRequest creates the Stop request.
func (client *VirtualMachinesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/virtualMachines/{virtualMachineName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
