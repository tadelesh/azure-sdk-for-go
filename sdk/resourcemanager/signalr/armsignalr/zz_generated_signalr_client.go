//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SignalRClient contains the methods for the SignalR group.
// Don't use this type directly, use NewSignalRClient() instead.
type SignalRClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSignalRClient creates a new instance of SignalRClient with the specified values.
func NewSignalRClient(con *arm.Connection, subscriptionID string) *SignalRClient {
	return &SignalRClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckNameAvailability - Checks that the resource name is valid and is not already in use.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) CheckNameAvailability(ctx context.Context, location string, parameters NameAvailabilityParameters, options *SignalRCheckNameAvailabilityOptions) (SignalRCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return SignalRCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SignalRCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SignalRCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *SignalRClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, parameters NameAvailabilityParameters, options *SignalRCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/locations/{location}/checkNameAvailability"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *SignalRClient) checkNameAvailabilityHandleResponse(resp *http.Response) (SignalRCheckNameAvailabilityResponse, error) {
	result := SignalRCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return SignalRCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *SignalRClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Create or update a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginCreateOrUpdateOptions) (SignalRCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SignalRCreateOrUpdatePollerResponse{}, err
	}
	result := SignalRCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SignalRClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SignalRCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SignalRCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SignalRClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SignalRClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Operation to delete a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginDeleteOptions) (SignalRDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SignalRDeletePollerResponse{}, err
	}
	result := SignalRDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SignalRClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SignalRDeletePollerResponse{}, err
	}
	result.Poller = &SignalRDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SignalRClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SignalRClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the resource and its properties.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRGetOptions) (SignalRGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SignalRGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SignalRGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SignalRGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SignalRClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SignalRClient) getHandleResponse(resp *http.Response) (SignalRGetResponse, error) {
	result := SignalRGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignalRResource); err != nil {
		return SignalRGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SignalRClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Handles requests to list all resources in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) ListByResourceGroup(resourceGroupName string, options *SignalRListByResourceGroupOptions) *SignalRListByResourceGroupPager {
	return &SignalRListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SignalRListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SignalRResourceList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SignalRClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SignalRListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SignalRClient) listByResourceGroupHandleResponse(resp *http.Response) (SignalRListByResourceGroupResponse, error) {
	result := SignalRListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignalRResourceList); err != nil {
		return SignalRListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SignalRClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Handles requests to list all resources in a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) ListBySubscription(options *SignalRListBySubscriptionOptions) *SignalRListBySubscriptionPager {
	return &SignalRListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SignalRListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SignalRResourceList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SignalRClient) listBySubscriptionCreateRequest(ctx context.Context, options *SignalRListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/signalR"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SignalRClient) listBySubscriptionHandleResponse(resp *http.Response) (SignalRListBySubscriptionResponse, error) {
	result := SignalRListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignalRResourceList); err != nil {
		return SignalRListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SignalRClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListKeys - Get the access keys of the resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) ListKeys(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRListKeysOptions) (SignalRListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SignalRListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SignalRListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SignalRListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *SignalRClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *SignalRClient) listKeysHandleResponse(resp *http.Response) (SignalRListKeysResponse, error) {
	result := SignalRListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignalRKeys); err != nil {
		return SignalRListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *SignalRClient) listKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSKUs - List all available skus of the resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) ListSKUs(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRListSKUsOptions) (SignalRListSKUsResponse, error) {
	req, err := client.listSKUsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SignalRListSKUsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SignalRListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SignalRListSKUsResponse{}, client.listSKUsHandleError(resp)
	}
	return client.listSKUsHandleResponse(resp)
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *SignalRClient) listSKUsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *SignalRClient) listSKUsHandleResponse(resp *http.Response) (SignalRListSKUsResponse, error) {
	result := SignalRListSKUsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUList); err != nil {
		return SignalRListSKUsResponse{}, err
	}
	return result, nil
}

// listSKUsHandleError handles the ListSKUs error response.
func (client *SignalRClient) listSKUsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRegenerateKey - Regenerate the access key for the resource. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) BeginRegenerateKey(ctx context.Context, resourceGroupName string, resourceName string, parameters RegenerateKeyParameters, options *SignalRBeginRegenerateKeyOptions) (SignalRRegenerateKeyPollerResponse, error) {
	resp, err := client.regenerateKey(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SignalRRegenerateKeyPollerResponse{}, err
	}
	result := SignalRRegenerateKeyPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SignalRClient.RegenerateKey", "azure-async-operation", resp, client.pl, client.regenerateKeyHandleError)
	if err != nil {
		return SignalRRegenerateKeyPollerResponse{}, err
	}
	result.Poller = &SignalRRegenerateKeyPoller{
		pt: pt,
	}
	return result, nil
}

// RegenerateKey - Regenerate the access key for the resource. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) regenerateKey(ctx context.Context, resourceGroupName string, resourceName string, parameters RegenerateKeyParameters, options *SignalRBeginRegenerateKeyOptions) (*http.Response, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.regenerateKeyHandleError(resp)
	}
	return resp, nil
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *SignalRClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters RegenerateKeyParameters, options *SignalRBeginRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleError handles the RegenerateKey error response.
func (client *SignalRClient) regenerateKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRestart - Operation to restart a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) BeginRestart(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginRestartOptions) (SignalRRestartPollerResponse, error) {
	resp, err := client.restart(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SignalRRestartPollerResponse{}, err
	}
	result := SignalRRestartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SignalRClient.Restart", "azure-async-operation", resp, client.pl, client.restartHandleError)
	if err != nil {
		return SignalRRestartPollerResponse{}, err
	}
	result.Poller = &SignalRRestartPoller{
		pt: pt,
	}
	return result, nil
}

// Restart - Operation to restart a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) restart(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.restartHandleError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *SignalRClient) restartCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SignalRBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *SignalRClient) restartHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Operation to update an exiting resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginUpdateOptions) (SignalRUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SignalRUpdatePollerResponse{}, err
	}
	result := SignalRUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SignalRClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return SignalRUpdatePollerResponse{}, err
	}
	result.Poller = &SignalRUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update an exiting resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SignalRClient) update(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SignalRClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters SignalRResource, options *SignalRBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *SignalRClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
