//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armattestation

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

// AttestationProvidersClient contains the methods for the AttestationProviders group.
// Don't use this type directly, use NewAttestationProvidersClient() instead.
type AttestationProvidersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAttestationProvidersClient creates a new instance of AttestationProvidersClient with the specified values.
func NewAttestationProvidersClient(con *arm.Connection, subscriptionID string) *AttestationProvidersClient {
	return &AttestationProvidersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Creates a new Attestation Provider.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) Create(ctx context.Context, resourceGroupName string, providerName string, creationParams AttestationServiceCreationParams, options *AttestationProvidersCreateOptions) (AttestationProvidersCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, providerName, creationParams, options)
	if err != nil {
		return AttestationProvidersCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AttestationProvidersCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AttestationProvidersClient) createCreateRequest(ctx context.Context, resourceGroupName string, providerName string, creationParams AttestationServiceCreationParams, options *AttestationProvidersCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, creationParams)
}

// createHandleResponse handles the Create response.
func (client *AttestationProvidersClient) createHandleResponse(resp *http.Response) (AttestationProvidersCreateResponse, error) {
	result := AttestationProvidersCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProvider); err != nil {
		return AttestationProvidersCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *AttestationProvidersClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete Attestation Service.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) Delete(ctx context.Context, resourceGroupName string, providerName string, options *AttestationProvidersDeleteOptions) (AttestationProvidersDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, providerName, options)
	if err != nil {
		return AttestationProvidersDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return AttestationProvidersDeleteResponse{}, client.deleteHandleError(resp)
	}
	return AttestationProvidersDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AttestationProvidersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, providerName string, options *AttestationProvidersDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AttestationProvidersClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the status of Attestation Provider.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) Get(ctx context.Context, resourceGroupName string, providerName string, options *AttestationProvidersGetOptions) (AttestationProvidersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, providerName, options)
	if err != nil {
		return AttestationProvidersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AttestationProvidersClient) getCreateRequest(ctx context.Context, resourceGroupName string, providerName string, options *AttestationProvidersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AttestationProvidersClient) getHandleResponse(resp *http.Response) (AttestationProvidersGetResponse, error) {
	result := AttestationProvidersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProvider); err != nil {
		return AttestationProvidersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AttestationProvidersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetDefaultByLocation - Get the default provider by location.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) GetDefaultByLocation(ctx context.Context, location string, options *AttestationProvidersGetDefaultByLocationOptions) (AttestationProvidersGetDefaultByLocationResponse, error) {
	req, err := client.getDefaultByLocationCreateRequest(ctx, location, options)
	if err != nil {
		return AttestationProvidersGetDefaultByLocationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersGetDefaultByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersGetDefaultByLocationResponse{}, client.getDefaultByLocationHandleError(resp)
	}
	return client.getDefaultByLocationHandleResponse(resp)
}

// getDefaultByLocationCreateRequest creates the GetDefaultByLocation request.
func (client *AttestationProvidersClient) getDefaultByLocationCreateRequest(ctx context.Context, location string, options *AttestationProvidersGetDefaultByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Attestation/locations/{location}/defaultProvider"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDefaultByLocationHandleResponse handles the GetDefaultByLocation response.
func (client *AttestationProvidersClient) getDefaultByLocationHandleResponse(resp *http.Response) (AttestationProvidersGetDefaultByLocationResponse, error) {
	result := AttestationProvidersGetDefaultByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProvider); err != nil {
		return AttestationProvidersGetDefaultByLocationResponse{}, err
	}
	return result, nil
}

// getDefaultByLocationHandleError handles the GetDefaultByLocation error response.
func (client *AttestationProvidersClient) getDefaultByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Returns a list of attestation providers in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) List(ctx context.Context, options *AttestationProvidersListOptions) (AttestationProvidersListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return AttestationProvidersListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AttestationProvidersClient) listCreateRequest(ctx context.Context, options *AttestationProvidersListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Attestation/attestationProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AttestationProvidersClient) listHandleResponse(resp *http.Response) (AttestationProvidersListResponse, error) {
	result := AttestationProvidersListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProviderListResult); err != nil {
		return AttestationProvidersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AttestationProvidersClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Returns attestation providers list in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *AttestationProvidersListByResourceGroupOptions) (AttestationProvidersListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return AttestationProvidersListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AttestationProvidersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AttestationProvidersListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AttestationProvidersClient) listByResourceGroupHandleResponse(resp *http.Response) (AttestationProvidersListByResourceGroupResponse, error) {
	result := AttestationProvidersListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProviderListResult); err != nil {
		return AttestationProvidersListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AttestationProvidersClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListDefault - Get the default provider
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) ListDefault(ctx context.Context, options *AttestationProvidersListDefaultOptions) (AttestationProvidersListDefaultResponse, error) {
	req, err := client.listDefaultCreateRequest(ctx, options)
	if err != nil {
		return AttestationProvidersListDefaultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersListDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersListDefaultResponse{}, client.listDefaultHandleError(resp)
	}
	return client.listDefaultHandleResponse(resp)
}

// listDefaultCreateRequest creates the ListDefault request.
func (client *AttestationProvidersClient) listDefaultCreateRequest(ctx context.Context, options *AttestationProvidersListDefaultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Attestation/defaultProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listDefaultHandleResponse handles the ListDefault response.
func (client *AttestationProvidersClient) listDefaultHandleResponse(resp *http.Response) (AttestationProvidersListDefaultResponse, error) {
	result := AttestationProvidersListDefaultResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProviderListResult); err != nil {
		return AttestationProvidersListDefaultResponse{}, err
	}
	return result, nil
}

// listDefaultHandleError handles the ListDefault error response.
func (client *AttestationProvidersClient) listDefaultHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the Attestation Provider.
// If the operation fails it returns the *CloudError error type.
func (client *AttestationProvidersClient) Update(ctx context.Context, resourceGroupName string, providerName string, updateParams AttestationServicePatchParams, options *AttestationProvidersUpdateOptions) (AttestationProvidersUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, providerName, updateParams, options)
	if err != nil {
		return AttestationProvidersUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationProvidersUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationProvidersUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AttestationProvidersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, providerName string, updateParams AttestationServicePatchParams, options *AttestationProvidersUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, updateParams)
}

// updateHandleResponse handles the Update response.
func (client *AttestationProvidersClient) updateHandleResponse(resp *http.Response) (AttestationProvidersUpdateResponse, error) {
	result := AttestationProvidersUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationProvider); err != nil {
		return AttestationProvidersUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *AttestationProvidersClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
