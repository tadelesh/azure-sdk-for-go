//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// VendorsClient contains the methods for the Vendors group.
// Don't use this type directly, use NewVendorsClient() instead.
type VendorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVendorsClient creates a new instance of VendorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVendorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VendorsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VendorsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a vendor.
// If the operation fails it returns an *azcore.ResponseError type.
// vendorName - The name of the vendor.
// options - VendorsClientBeginCreateOrUpdateOptions contains the optional parameters for the VendorsClient.BeginCreateOrUpdate
// method.
func (client *VendorsClient) BeginCreateOrUpdate(ctx context.Context, vendorName string, options *VendorsClientBeginCreateOrUpdateOptions) (VendorsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, vendorName, options)
	if err != nil {
		return VendorsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VendorsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("VendorsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VendorsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VendorsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a vendor.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VendorsClient) createOrUpdate(ctx context.Context, vendorName string, options *VendorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vendorName, options)
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
func (client *VendorsClient) createOrUpdateCreateRequest(ctx context.Context, vendorName string, options *VendorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// BeginDelete - Deletes the specified vendor.
// If the operation fails it returns an *azcore.ResponseError type.
// vendorName - The name of the vendor.
// options - VendorsClientBeginDeleteOptions contains the optional parameters for the VendorsClient.BeginDelete method.
func (client *VendorsClient) BeginDelete(ctx context.Context, vendorName string, options *VendorsClientBeginDeleteOptions) (VendorsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, vendorName, options)
	if err != nil {
		return VendorsClientDeletePollerResponse{}, err
	}
	result := VendorsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("VendorsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VendorsClientDeletePollerResponse{}, err
	}
	result.Poller = &VendorsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified vendor.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VendorsClient) deleteOperation(ctx context.Context, vendorName string, options *VendorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, vendorName, options)
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
func (client *VendorsClient) deleteCreateRequest(ctx context.Context, vendorName string, options *VendorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified vendor.
// If the operation fails it returns an *azcore.ResponseError type.
// vendorName - The name of the vendor.
// options - VendorsClientGetOptions contains the optional parameters for the VendorsClient.Get method.
func (client *VendorsClient) Get(ctx context.Context, vendorName string, options *VendorsClientGetOptions) (VendorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vendorName, options)
	if err != nil {
		return VendorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VendorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VendorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VendorsClient) getCreateRequest(ctx context.Context, vendorName string, options *VendorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VendorsClient) getHandleResponse(resp *http.Response) (VendorsClientGetResponse, error) {
	result := VendorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Vendor); err != nil {
		return VendorsClientGetResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all the vendors in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VendorsClientListBySubscriptionOptions contains the optional parameters for the VendorsClient.ListBySubscription
// method.
func (client *VendorsClient) ListBySubscription(options *VendorsClientListBySubscriptionOptions) *VendorsClientListBySubscriptionPager {
	return &VendorsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VendorsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VendorListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VendorsClient) listBySubscriptionCreateRequest(ctx context.Context, options *VendorsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VendorsClient) listBySubscriptionHandleResponse(resp *http.Response) (VendorsClientListBySubscriptionResponse, error) {
	result := VendorsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VendorListResult); err != nil {
		return VendorsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
