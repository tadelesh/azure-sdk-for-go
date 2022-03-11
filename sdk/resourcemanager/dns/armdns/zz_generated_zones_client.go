//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

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
	"strconv"
	"strings"
)

// ZonesClient contains the methods for the Zones group.
// Don't use this type directly, use NewZonesClient() instead.
type ZonesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewZonesClient creates a new instance of ZonesClient with the specified values.
// subscriptionID - Specifies the Azure subscription ID, which uniquely identifies the Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ZonesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ZonesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates a DNS zone. Does not modify DNS records within the zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// zoneName - The name of the DNS zone (without a terminating dot).
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - ZonesClientCreateOrUpdateOptions contains the optional parameters for the ZonesClient.CreateOrUpdate method.
func (client *ZonesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, parameters Zone, options *ZonesClientCreateOrUpdateOptions) (ZonesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, zoneName, parameters, options)
	if err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ZonesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ZonesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, parameters Zone, options *ZonesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ZonesClient) createOrUpdateHandleResponse(resp *http.Response) (ZonesClientCreateOrUpdateResponse, error) {
	result := ZonesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Zone); err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// zoneName - The name of the DNS zone (without a terminating dot).
// options - ZonesClientBeginDeleteOptions contains the optional parameters for the ZonesClient.BeginDelete method.
func (client *ZonesClient) BeginDelete(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (ZonesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, zoneName, options)
	if err != nil {
		return ZonesClientDeletePollerResponse{}, err
	}
	result := ZonesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ZonesClient.Delete", "", resp, client.pl)
	if err != nil {
		return ZonesClientDeletePollerResponse{}, err
	}
	result.Poller = &ZonesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ZonesClient) deleteOperation(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, zoneName, options)
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
func (client *ZonesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a DNS zone. Retrieves the zone properties, but not the record sets within the zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// zoneName - The name of the DNS zone (without a terminating dot).
// options - ZonesClientGetOptions contains the optional parameters for the ZonesClient.Get method.
func (client *ZonesClient) Get(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientGetOptions) (ZonesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, zoneName, options)
	if err != nil {
		return ZonesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ZonesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ZonesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ZonesClient) getCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ZonesClient) getHandleResponse(resp *http.Response) (ZonesClientGetResponse, error) {
	result := ZonesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Zone); err != nil {
		return ZonesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the DNS zones in all resource groups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ZonesClientListOptions contains the optional parameters for the ZonesClient.List method.
func (client *ZonesClient) List(options *ZonesClientListOptions) *ZonesClientListPager {
	return &ZonesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ZonesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ZoneListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ZonesClient) listCreateRequest(ctx context.Context, options *ZonesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ZonesClient) listHandleResponse(resp *http.Response) (ZonesClientListResponse, error) {
	result := ZonesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ZoneListResult); err != nil {
		return ZonesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists the DNS zones within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ZonesClientListByResourceGroupOptions contains the optional parameters for the ZonesClient.ListByResourceGroup
// method.
func (client *ZonesClient) ListByResourceGroup(resourceGroupName string, options *ZonesClientListByResourceGroupOptions) *ZonesClientListByResourceGroupPager {
	return &ZonesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ZonesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ZoneListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ZonesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ZonesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ZonesClient) listByResourceGroupHandleResponse(resp *http.Response) (ZonesClientListByResourceGroupResponse, error) {
	result := ZonesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ZoneListResult); err != nil {
		return ZonesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a DNS zone. Does not modify DNS records within the zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// zoneName - The name of the DNS zone (without a terminating dot).
// parameters - Parameters supplied to the Update operation.
// options - ZonesClientUpdateOptions contains the optional parameters for the ZonesClient.Update method.
func (client *ZonesClient) Update(ctx context.Context, resourceGroupName string, zoneName string, parameters ZoneUpdate, options *ZonesClientUpdateOptions) (ZonesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, zoneName, parameters, options)
	if err != nil {
		return ZonesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ZonesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ZonesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ZonesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, parameters ZoneUpdate, options *ZonesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ZonesClient) updateHandleResponse(resp *http.Response) (ZonesClientUpdateResponse, error) {
	result := ZonesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Zone); err != nil {
		return ZonesClientUpdateResponse{}, err
	}
	return result, nil
}
