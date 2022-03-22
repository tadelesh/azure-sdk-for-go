//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

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

// VirtualMachineImageTemplatesClient contains the methods for the VirtualMachineImageTemplates group.
// Don't use this type directly, use NewVirtualMachineImageTemplatesClient() instead.
type VirtualMachineImageTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineImageTemplatesClient creates a new instance of VirtualMachineImageTemplatesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription Id forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineImageTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineImageTemplatesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualMachineImageTemplatesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCancel - Cancel the long running image build based on the image template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// options - VirtualMachineImageTemplatesClientBeginCancelOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.BeginCancel
// method.
func (client *VirtualMachineImageTemplatesClient) BeginCancel(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginCancelOptions) (*armruntime.Poller[VirtualMachineImageTemplatesClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, imageTemplateName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachineImageTemplatesClientCancelResponse]("VirtualMachineImageTemplatesClient.Cancel", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachineImageTemplatesClientCancelResponse]("VirtualMachineImageTemplatesClient.Cancel", options.ResumeToken, client.pl, nil)
	}
}

// Cancel - Cancel the long running image build based on the image template
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineImageTemplatesClient) cancel(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, imageTemplateName, options)
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

// cancelCreateRequest creates the Cancel request.
func (client *VirtualMachineImageTemplatesClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginCreateOrUpdate - Create or update a virtual machine image template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// parameters - Parameters supplied to the CreateImageTemplate operation
// options - VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.BeginCreateOrUpdate
// method.
func (client *VirtualMachineImageTemplatesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplate, options *VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[VirtualMachineImageTemplatesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, imageTemplateName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachineImageTemplatesClientCreateOrUpdateResponse]("VirtualMachineImageTemplatesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachineImageTemplatesClientCreateOrUpdateResponse]("VirtualMachineImageTemplatesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a virtual machine image template
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineImageTemplatesClient) createOrUpdate(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplate, options *VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, imageTemplateName, parameters, options)
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
func (client *VirtualMachineImageTemplatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplate, options *VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a virtual machine image template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// options - VirtualMachineImageTemplatesClientBeginDeleteOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.BeginDelete
// method.
func (client *VirtualMachineImageTemplatesClient) BeginDelete(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginDeleteOptions) (*armruntime.Poller[VirtualMachineImageTemplatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, imageTemplateName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachineImageTemplatesClientDeleteResponse]("VirtualMachineImageTemplatesClient.Delete", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachineImageTemplatesClientDeleteResponse]("VirtualMachineImageTemplatesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a virtual machine image template
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineImageTemplatesClient) deleteOperation(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, imageTemplateName, options)
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
func (client *VirtualMachineImageTemplatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get information about a virtual machine image template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// options - VirtualMachineImageTemplatesClientGetOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.Get
// method.
func (client *VirtualMachineImageTemplatesClient) Get(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientGetOptions) (VirtualMachineImageTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, imageTemplateName, options)
	if err != nil {
		return VirtualMachineImageTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImageTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImageTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImageTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineImageTemplatesClient) getHandleResponse(resp *http.Response) (VirtualMachineImageTemplatesClientGetResponse, error) {
	result := VirtualMachineImageTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageTemplate); err != nil {
		return VirtualMachineImageTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// GetRunOutput - Get the specified run output for the specified image template resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// runOutputName - The name of the run output
// options - VirtualMachineImageTemplatesClientGetRunOutputOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.GetRunOutput
// method.
func (client *VirtualMachineImageTemplatesClient) GetRunOutput(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string, options *VirtualMachineImageTemplatesClientGetRunOutputOptions) (VirtualMachineImageTemplatesClientGetRunOutputResponse, error) {
	req, err := client.getRunOutputCreateRequest(ctx, resourceGroupName, imageTemplateName, runOutputName, options)
	if err != nil {
		return VirtualMachineImageTemplatesClientGetRunOutputResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImageTemplatesClientGetRunOutputResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImageTemplatesClientGetRunOutputResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRunOutputHandleResponse(resp)
}

// getRunOutputCreateRequest creates the GetRunOutput request.
func (client *VirtualMachineImageTemplatesClient) getRunOutputCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string, options *VirtualMachineImageTemplatesClientGetRunOutputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs/{runOutputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	if runOutputName == "" {
		return nil, errors.New("parameter runOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runOutputName}", url.PathEscape(runOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getRunOutputHandleResponse handles the GetRunOutput response.
func (client *VirtualMachineImageTemplatesClient) getRunOutputHandleResponse(resp *http.Response) (VirtualMachineImageTemplatesClientGetRunOutputResponse, error) {
	result := VirtualMachineImageTemplatesClientGetRunOutputResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunOutput); err != nil {
		return VirtualMachineImageTemplatesClientGetRunOutputResponse{}, err
	}
	return result, nil
}

// List - Gets information about the VM image templates associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VirtualMachineImageTemplatesClientListOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.List
// method.
func (client *VirtualMachineImageTemplatesClient) List(options *VirtualMachineImageTemplatesClientListOptions) *runtime.Pager[VirtualMachineImageTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualMachineImageTemplatesClientListResponse]{
		More: func(page VirtualMachineImageTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineImageTemplatesClientListResponse) (VirtualMachineImageTemplatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineImageTemplatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineImageTemplatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineImageTemplatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImageTemplatesClient) listCreateRequest(ctx context.Context, options *VirtualMachineImageTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VirtualMachineImages/imageTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImageTemplatesClient) listHandleResponse(resp *http.Response) (VirtualMachineImageTemplatesClientListResponse, error) {
	result := VirtualMachineImageTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageTemplateListResult); err != nil {
		return VirtualMachineImageTemplatesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets information about the VM image templates associated with the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - VirtualMachineImageTemplatesClientListByResourceGroupOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.ListByResourceGroup
// method.
func (client *VirtualMachineImageTemplatesClient) ListByResourceGroup(resourceGroupName string, options *VirtualMachineImageTemplatesClientListByResourceGroupOptions) *runtime.Pager[VirtualMachineImageTemplatesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualMachineImageTemplatesClientListByResourceGroupResponse]{
		More: func(page VirtualMachineImageTemplatesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineImageTemplatesClientListByResourceGroupResponse) (VirtualMachineImageTemplatesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineImageTemplatesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineImageTemplatesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineImageTemplatesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualMachineImageTemplatesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualMachineImageTemplatesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates"
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
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualMachineImageTemplatesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualMachineImageTemplatesClientListByResourceGroupResponse, error) {
	result := VirtualMachineImageTemplatesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageTemplateListResult); err != nil {
		return VirtualMachineImageTemplatesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListRunOutputs - List all run outputs for the specified Image Template resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// options - VirtualMachineImageTemplatesClientListRunOutputsOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.ListRunOutputs
// method.
func (client *VirtualMachineImageTemplatesClient) ListRunOutputs(resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientListRunOutputsOptions) *runtime.Pager[VirtualMachineImageTemplatesClientListRunOutputsResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualMachineImageTemplatesClientListRunOutputsResponse]{
		More: func(page VirtualMachineImageTemplatesClientListRunOutputsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineImageTemplatesClientListRunOutputsResponse) (VirtualMachineImageTemplatesClientListRunOutputsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRunOutputsCreateRequest(ctx, resourceGroupName, imageTemplateName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineImageTemplatesClientListRunOutputsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineImageTemplatesClientListRunOutputsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineImageTemplatesClientListRunOutputsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRunOutputsHandleResponse(resp)
		},
	})
}

// listRunOutputsCreateRequest creates the ListRunOutputs request.
func (client *VirtualMachineImageTemplatesClient) listRunOutputsCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientListRunOutputsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRunOutputsHandleResponse handles the ListRunOutputs response.
func (client *VirtualMachineImageTemplatesClient) listRunOutputsHandleResponse(resp *http.Response) (VirtualMachineImageTemplatesClientListRunOutputsResponse, error) {
	result := VirtualMachineImageTemplatesClientListRunOutputsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunOutputCollection); err != nil {
		return VirtualMachineImageTemplatesClientListRunOutputsResponse{}, err
	}
	return result, nil
}

// BeginRun - Create artifacts from a existing image template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// options - VirtualMachineImageTemplatesClientBeginRunOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.BeginRun
// method.
func (client *VirtualMachineImageTemplatesClient) BeginRun(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginRunOptions) (*armruntime.Poller[VirtualMachineImageTemplatesClientRunResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.run(ctx, resourceGroupName, imageTemplateName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachineImageTemplatesClientRunResponse]("VirtualMachineImageTemplatesClient.Run", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachineImageTemplatesClientRunResponse]("VirtualMachineImageTemplatesClient.Run", options.ResumeToken, client.pl, nil)
	}
}

// Run - Create artifacts from a existing image template
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineImageTemplatesClient) run(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginRunOptions) (*http.Response, error) {
	req, err := client.runCreateRequest(ctx, resourceGroupName, imageTemplateName, options)
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

// runCreateRequest creates the Run request.
func (client *VirtualMachineImageTemplatesClient) runCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, options *VirtualMachineImageTemplatesClientBeginRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/run"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Update the tags for this Virtual Machine Image Template
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageTemplateName - The name of the image Template
// parameters - Additional parameters for Image Template update.
// options - VirtualMachineImageTemplatesClientBeginUpdateOptions contains the optional parameters for the VirtualMachineImageTemplatesClient.BeginUpdate
// method.
func (client *VirtualMachineImageTemplatesClient) BeginUpdate(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplateUpdateParameters, options *VirtualMachineImageTemplatesClientBeginUpdateOptions) (*armruntime.Poller[VirtualMachineImageTemplatesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, imageTemplateName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualMachineImageTemplatesClientUpdateResponse]("VirtualMachineImageTemplatesClient.Update", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualMachineImageTemplatesClientUpdateResponse]("VirtualMachineImageTemplatesClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the tags for this Virtual Machine Image Template
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineImageTemplatesClient) update(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplateUpdateParameters, options *VirtualMachineImageTemplatesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, imageTemplateName, parameters, options)
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
func (client *VirtualMachineImageTemplatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters ImageTemplateUpdateParameters, options *VirtualMachineImageTemplatesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageTemplateName == "" {
		return nil, errors.New("parameter imageTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageTemplateName}", url.PathEscape(imageTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
