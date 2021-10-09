//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// VirtualMachineTemplatesClient contains the methods for the VirtualMachineTemplates group.
// Don't use this type directly, use NewVirtualMachineTemplatesClient() instead.
type VirtualMachineTemplatesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineTemplatesClient creates a new instance of VirtualMachineTemplatesClient with the specified values.
func NewVirtualMachineTemplatesClient(con *arm.Connection, subscriptionID string) *VirtualMachineTemplatesClient {
	return &VirtualMachineTemplatesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Returns virtual machine templates by its name
// If the operation fails it returns the *CSRPError error type.
func (client *VirtualMachineTemplatesClient) Get(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesGetOptions) (VirtualMachineTemplatesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, virtualMachineTemplateName, options)
	if err != nil {
		return VirtualMachineTemplatesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineTemplatesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineTemplatesClient) getCreateRequest(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineTemplatesClient) getHandleResponse(resp *http.Response) (VirtualMachineTemplatesGetResponse, error) {
	result := VirtualMachineTemplatesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplate); err != nil {
		return VirtualMachineTemplatesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineTemplatesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CSRPError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Returns list of virtual machine templates in region for private cloud
// If the operation fails it returns the *CSRPError error type.
func (client *VirtualMachineTemplatesClient) List(pcName string, regionID string, resourcePoolName string, options *VirtualMachineTemplatesListOptions) *VirtualMachineTemplatesListPager {
	return &VirtualMachineTemplatesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, pcName, regionID, resourcePoolName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineTemplatesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineTemplateListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineTemplatesClient) listCreateRequest(ctx context.Context, pcName string, regionID string, resourcePoolName string, options *VirtualMachineTemplatesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	reqQP.Set("resourcePoolName", resourcePoolName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineTemplatesClient) listHandleResponse(resp *http.Response) (VirtualMachineTemplatesListResponse, error) {
	result := VirtualMachineTemplatesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplateListResponse); err != nil {
		return VirtualMachineTemplatesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineTemplatesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CSRPError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
