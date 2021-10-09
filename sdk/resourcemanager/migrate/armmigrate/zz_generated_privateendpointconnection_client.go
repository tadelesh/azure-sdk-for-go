//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// PrivateEndpointConnectionClient contains the methods for the PrivateEndpointConnection group.
// Don't use this type directly, use NewPrivateEndpointConnectionClient() instead.
type PrivateEndpointConnectionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient with the specified values.
func NewPrivateEndpointConnectionClient(con *arm.Connection, subscriptionID string) *PrivateEndpointConnectionClient {
	return &PrivateEndpointConnectionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Delete - Delete the private endpoint connection from the project. T.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateEndpointConnectionClient) Delete(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionDeleteOptions) (PrivateEndpointConnectionDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PrivateEndpointConnectionDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *PrivateEndpointConnectionClient) deleteHandleResponse(resp *http.Response) (PrivateEndpointConnectionDeleteResponse, error) {
	result := PrivateEndpointConnectionDeleteResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateEndpointConnectionClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get information related to a specific private endpoint connection in the project. Returns a json object of type 'privateEndpointConnections' as
// specified in the models section.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateEndpointConnectionClient) Get(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionGetOptions) (PrivateEndpointConnectionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionGetResponse, error) {
	result := PrivateEndpointConnectionGetResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointConnectionClient) getHandleError(resp *http.Response) error {
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

// ListByProject - Get all private endpoint connections in the project. Returns a json array of objects of type 'privateEndpointConnections' as specified
// in the Models section.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateEndpointConnectionClient) ListByProject(ctx context.Context, resourceGroupName string, projectName string, options *PrivateEndpointConnectionListByProjectOptions) (PrivateEndpointConnectionListByProjectResponse, error) {
	req, err := client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return PrivateEndpointConnectionListByProjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionListByProjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionListByProjectResponse{}, client.listByProjectHandleError(resp)
	}
	return client.listByProjectHandleResponse(resp)
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *PrivateEndpointConnectionClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *PrivateEndpointConnectionListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *PrivateEndpointConnectionClient) listByProjectHandleResponse(resp *http.Response) (PrivateEndpointConnectionListByProjectResponse, error) {
	result := PrivateEndpointConnectionListByProjectResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionCollection); err != nil {
		return PrivateEndpointConnectionListByProjectResponse{}, err
	}
	return result, nil
}

// listByProjectHandleError handles the ListByProject error response.
func (client *PrivateEndpointConnectionClient) listByProjectHandleError(resp *http.Response) error {
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

// Update - Update a specific private endpoint connection in the project.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateEndpointConnectionClient) Update(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionUpdateOptions) (PrivateEndpointConnectionUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return PrivateEndpointConnectionUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PrivateEndpointConnectionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.PrivateEndpointConnectionBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.PrivateEndpointConnectionBody)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PrivateEndpointConnectionClient) updateHandleResponse(resp *http.Response) (PrivateEndpointConnectionUpdateResponse, error) {
	result := PrivateEndpointConnectionUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *PrivateEndpointConnectionClient) updateHandleError(resp *http.Response) error {
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
