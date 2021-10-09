//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// RulesEnginesClient contains the methods for the RulesEngines group.
// Don't use this type directly, use NewRulesEnginesClient() instead.
type RulesEnginesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRulesEnginesClient creates a new instance of RulesEnginesClient with the specified values.
func NewRulesEnginesClient(con *arm.Connection, subscriptionID string) *RulesEnginesClient {
	return &RulesEnginesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesBeginCreateOrUpdateOptions) (RulesEnginesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, frontDoorName, rulesEngineName, rulesEngineParameters, options)
	if err != nil {
		return RulesEnginesCreateOrUpdatePollerResponse{}, err
	}
	result := RulesEnginesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RulesEnginesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return RulesEnginesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RulesEnginesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) createOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, rulesEngineParameters, options)
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
func (client *RulesEnginesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, rulesEngineParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RulesEnginesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes an existing Rules Engine Configuration with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) BeginDelete(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesBeginDeleteOptions) (RulesEnginesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
	if err != nil {
		return RulesEnginesDeletePollerResponse{}, err
	}
	result := RulesEnginesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RulesEnginesClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return RulesEnginesDeletePollerResponse{}, err
	}
	result.Poller = &RulesEnginesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing Rules Engine Configuration with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) deleteOperation(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RulesEnginesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RulesEnginesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets a Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesGetOptions) (RulesEnginesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
	if err != nil {
		return RulesEnginesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesEnginesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RulesEnginesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RulesEnginesClient) getCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RulesEnginesClient) getHandleResponse(resp *http.Response) (RulesEnginesGetResponse, error) {
	result := RulesEnginesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RulesEngine); err != nil {
		return RulesEnginesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RulesEnginesClient) getHandleError(resp *http.Response) error {
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

// ListByFrontDoor - Lists all of the Rules Engine Configurations within a Front Door.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RulesEnginesClient) ListByFrontDoor(resourceGroupName string, frontDoorName string, options *RulesEnginesListByFrontDoorOptions) *RulesEnginesListByFrontDoorPager {
	return &RulesEnginesListByFrontDoorPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFrontDoorCreateRequest(ctx, resourceGroupName, frontDoorName, options)
		},
		advancer: func(ctx context.Context, resp RulesEnginesListByFrontDoorResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RulesEngineListResult.NextLink)
		},
	}
}

// listByFrontDoorCreateRequest creates the ListByFrontDoor request.
func (client *RulesEnginesClient) listByFrontDoorCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *RulesEnginesListByFrontDoorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFrontDoorHandleResponse handles the ListByFrontDoor response.
func (client *RulesEnginesClient) listByFrontDoorHandleResponse(resp *http.Response) (RulesEnginesListByFrontDoorResponse, error) {
	result := RulesEnginesListByFrontDoorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RulesEngineListResult); err != nil {
		return RulesEnginesListByFrontDoorResponse{}, err
	}
	return result, nil
}

// listByFrontDoorHandleError handles the ListByFrontDoor error response.
func (client *RulesEnginesClient) listByFrontDoorHandleError(resp *http.Response) error {
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
