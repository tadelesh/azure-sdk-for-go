//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

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

// MonitorClient contains the methods for the Monitor group.
// Don't use this type directly, use NewMonitorClient() instead.
type MonitorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMonitorClient creates a new instance of MonitorClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMonitorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MonitorClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &MonitorClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListVMHostUpdate - Sending request to update the collection when Logz.io agent has been installed on a VM for a given monitor.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorClientListVMHostUpdateOptions contains the optional parameters for the MonitorClient.ListVMHostUpdate
// method.
func (client *MonitorClient) ListVMHostUpdate(resourceGroupName string, monitorName string, options *MonitorClientListVMHostUpdateOptions) *MonitorClientListVMHostUpdatePager {
	return &MonitorClientListVMHostUpdatePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listVMHostUpdateCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorClientListVMHostUpdateResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VMResourcesListResponse.NextLink)
		},
	}
}

// listVMHostUpdateCreateRequest creates the ListVMHostUpdate request.
func (client *MonitorClient) listVMHostUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorClientListVMHostUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/vmHostUpdate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// listVMHostUpdateHandleResponse handles the ListVMHostUpdate response.
func (client *MonitorClient) listVMHostUpdateHandleResponse(resp *http.Response) (MonitorClientListVMHostUpdateResponse, error) {
	result := MonitorClientListVMHostUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMResourcesListResponse); err != nil {
		return MonitorClientListVMHostUpdateResponse{}, err
	}
	return result, nil
}

// ListVMHosts - List the compute resources currently being monitored by the Logz main account resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorClientListVMHostsOptions contains the optional parameters for the MonitorClient.ListVMHosts method.
func (client *MonitorClient) ListVMHosts(resourceGroupName string, monitorName string, options *MonitorClientListVMHostsOptions) *MonitorClientListVMHostsPager {
	return &MonitorClientListVMHostsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listVMHostsCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorClientListVMHostsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VMResourcesListResponse.NextLink)
		},
	}
}

// listVMHostsCreateRequest creates the ListVMHosts request.
func (client *MonitorClient) listVMHostsCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorClientListVMHostsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/listVMHosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listVMHostsHandleResponse handles the ListVMHosts response.
func (client *MonitorClient) listVMHostsHandleResponse(resp *http.Response) (MonitorClientListVMHostsResponse, error) {
	result := MonitorClientListVMHostsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMResourcesListResponse); err != nil {
		return MonitorClientListVMHostsResponse{}, err
	}
	return result, nil
}

// VMHostPayload - Returns the payload that needs to be passed in the request body for installing Logz.io agent on a VM.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorClientVMHostPayloadOptions contains the optional parameters for the MonitorClient.VMHostPayload method.
func (client *MonitorClient) VMHostPayload(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorClientVMHostPayloadOptions) (MonitorClientVMHostPayloadResponse, error) {
	req, err := client.vmHostPayloadCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorClientVMHostPayloadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorClientVMHostPayloadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorClientVMHostPayloadResponse{}, runtime.NewResponseError(resp)
	}
	return client.vmHostPayloadHandleResponse(resp)
}

// vmHostPayloadCreateRequest creates the VMHostPayload request.
func (client *MonitorClient) vmHostPayloadCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorClientVMHostPayloadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/vmHostPayload"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// vmHostPayloadHandleResponse handles the VMHostPayload response.
func (client *MonitorClient) vmHostPayloadHandleResponse(resp *http.Response) (MonitorClientVMHostPayloadResponse, error) {
	result := MonitorClientVMHostPayloadResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMExtensionPayload); err != nil {
		return MonitorClientVMHostPayloadResponse{}, err
	}
	return result, nil
}
