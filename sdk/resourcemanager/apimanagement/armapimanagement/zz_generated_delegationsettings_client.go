//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// DelegationSettingsClient contains the methods for the DelegationSettings group.
// Don't use this type directly, use NewDelegationSettingsClient() instead.
type DelegationSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDelegationSettingsClient creates a new instance of DelegationSettingsClient with the specified values.
func NewDelegationSettingsClient(con *arm.Connection, subscriptionID string) *DelegationSettingsClient {
	return &DelegationSettingsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or Update Delegation settings.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DelegationSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalDelegationSettings, options *DelegationSettingsCreateOrUpdateOptions) (DelegationSettingsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
	if err != nil {
		return DelegationSettingsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DelegationSettingsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DelegationSettingsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DelegationSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalDelegationSettings, options *DelegationSettingsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DelegationSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DelegationSettingsCreateOrUpdateResponse, error) {
	result := DelegationSettingsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalDelegationSettings); err != nil {
		return DelegationSettingsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DelegationSettingsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get Delegation Settings for the Portal.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DelegationSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsGetOptions) (DelegationSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DelegationSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DelegationSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DelegationSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DelegationSettingsClient) getHandleResponse(resp *http.Response) (DelegationSettingsGetResponse, error) {
	result := DelegationSettingsGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalDelegationSettings); err != nil {
		return DelegationSettingsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DelegationSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetEntityTag - Gets the entity state (Etag) version of the DelegationSettings.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DelegationSettingsClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsGetEntityTagOptions) (DelegationSettingsGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DelegationSettingsGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *DelegationSettingsClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *DelegationSettingsClient) getEntityTagHandleResponse(resp *http.Response) (DelegationSettingsGetEntityTagResponse, error) {
	result := DelegationSettingsGetEntityTagResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListSecrets - Gets the secret validation key of the DelegationSettings.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DelegationSettingsClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsListSecretsOptions) (DelegationSettingsListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DelegationSettingsListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DelegationSettingsListSecretsResponse{}, client.listSecretsHandleError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *DelegationSettingsClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *DelegationSettingsClient) listSecretsHandleResponse(resp *http.Response) (DelegationSettingsListSecretsResponse, error) {
	result := DelegationSettingsListSecretsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalSettingValidationKeyContract); err != nil {
		return DelegationSettingsListSecretsResponse{}, err
	}
	return result, nil
}

// listSecretsHandleError handles the ListSecrets error response.
func (client *DelegationSettingsClient) listSecretsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update Delegation settings.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DelegationSettingsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, ifMatch string, parameters PortalDelegationSettings, options *DelegationSettingsUpdateOptions) (DelegationSettingsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, ifMatch, parameters, options)
	if err != nil {
		return DelegationSettingsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DelegationSettingsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return DelegationSettingsUpdateResponse{}, client.updateHandleError(resp)
	}
	return DelegationSettingsUpdateResponse{RawResponse: resp}, nil
}

// updateCreateRequest creates the Update request.
func (client *DelegationSettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, ifMatch string, parameters PortalDelegationSettings, options *DelegationSettingsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *DelegationSettingsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
