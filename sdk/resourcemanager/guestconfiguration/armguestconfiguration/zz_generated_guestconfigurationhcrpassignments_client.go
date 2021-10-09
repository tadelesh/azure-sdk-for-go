//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

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

// GuestConfigurationHCRPAssignmentsClient contains the methods for the GuestConfigurationHCRPAssignments group.
// Don't use this type directly, use NewGuestConfigurationHCRPAssignmentsClient() instead.
type GuestConfigurationHCRPAssignmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewGuestConfigurationHCRPAssignmentsClient creates a new instance of GuestConfigurationHCRPAssignmentsClient with the specified values.
func NewGuestConfigurationHCRPAssignmentsClient(con *arm.Connection, subscriptionID string) *GuestConfigurationHCRPAssignmentsClient {
	return &GuestConfigurationHCRPAssignmentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates an association between a ARC machine and guest configuration
// If the operation fails it returns the *ErrorResponse error type.
func (client *GuestConfigurationHCRPAssignmentsClient) CreateOrUpdate(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, machineName string, parameters GuestConfigurationAssignment, options *GuestConfigurationHCRPAssignmentsCreateOrUpdateOptions) (GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, guestConfigurationAssignmentName, resourceGroupName, machineName, parameters, options)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GuestConfigurationHCRPAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, machineName string, parameters GuestConfigurationAssignment, options *GuestConfigurationHCRPAssignmentsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GuestConfigurationHCRPAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse, error) {
	result := GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GuestConfigurationAssignment); err != nil {
		return GuestConfigurationHCRPAssignmentsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GuestConfigurationHCRPAssignmentsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete a guest configuration assignment
// If the operation fails it returns the *ErrorResponse error type.
func (client *GuestConfigurationHCRPAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *GuestConfigurationHCRPAssignmentsDeleteOptions) (GuestConfigurationHCRPAssignmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, machineName, options)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GuestConfigurationHCRPAssignmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return GuestConfigurationHCRPAssignmentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GuestConfigurationHCRPAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *GuestConfigurationHCRPAssignmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GuestConfigurationHCRPAssignmentsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get information about a guest configuration assignment
// If the operation fails it returns the *ErrorResponse error type.
func (client *GuestConfigurationHCRPAssignmentsClient) Get(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *GuestConfigurationHCRPAssignmentsGetOptions) (GuestConfigurationHCRPAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, machineName, options)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GuestConfigurationHCRPAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GuestConfigurationHCRPAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *GuestConfigurationHCRPAssignmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GuestConfigurationHCRPAssignmentsClient) getHandleResponse(resp *http.Response) (GuestConfigurationHCRPAssignmentsGetResponse, error) {
	result := GuestConfigurationHCRPAssignmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GuestConfigurationAssignment); err != nil {
		return GuestConfigurationHCRPAssignmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GuestConfigurationHCRPAssignmentsClient) getHandleError(resp *http.Response) error {
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

// List - List all guest configuration assignments for an ARC machine.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GuestConfigurationHCRPAssignmentsClient) List(ctx context.Context, resourceGroupName string, machineName string, options *GuestConfigurationHCRPAssignmentsListOptions) (GuestConfigurationHCRPAssignmentsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, machineName, options)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GuestConfigurationHCRPAssignmentsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GuestConfigurationHCRPAssignmentsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *GuestConfigurationHCRPAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *GuestConfigurationHCRPAssignmentsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GuestConfigurationHCRPAssignmentsClient) listHandleResponse(resp *http.Response) (GuestConfigurationHCRPAssignmentsListResponse, error) {
	result := GuestConfigurationHCRPAssignmentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GuestConfigurationAssignmentList); err != nil {
		return GuestConfigurationHCRPAssignmentsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *GuestConfigurationHCRPAssignmentsClient) listHandleError(resp *http.Response) error {
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
