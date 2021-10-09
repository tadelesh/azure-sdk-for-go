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

// VMwareCollectorsClient contains the methods for the VMwareCollectors group.
// Don't use this type directly, use NewVMwareCollectorsClient() instead.
type VMwareCollectorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVMwareCollectorsClient creates a new instance of VMwareCollectorsClient with the specified values.
func NewVMwareCollectorsClient(con *arm.Connection, subscriptionID string) *VMwareCollectorsClient {
	return &VMwareCollectorsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Create or Update VMware collector
// If the operation fails it returns the *CloudError error type.
func (client *VMwareCollectorsClient) Create(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsCreateOptions) (VMwareCollectorsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, vmWareCollectorName, options)
	if err != nil {
		return VMwareCollectorsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMwareCollectorsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return VMwareCollectorsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *VMwareCollectorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}"
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
	if vmWareCollectorName == "" {
		return nil, errors.New("parameter vmWareCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmWareCollectorName}", url.PathEscape(vmWareCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.CollectorBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.CollectorBody)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *VMwareCollectorsClient) createHandleResponse(resp *http.Response) (VMwareCollectorsCreateResponse, error) {
	result := VMwareCollectorsCreateResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMwareCollector); err != nil {
		return VMwareCollectorsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *VMwareCollectorsClient) createHandleError(resp *http.Response) error {
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

// Delete - Delete a VMware collector from the project.
// If the operation fails it returns the *CloudError error type.
func (client *VMwareCollectorsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsDeleteOptions) (VMwareCollectorsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, vmWareCollectorName, options)
	if err != nil {
		return VMwareCollectorsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMwareCollectorsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VMwareCollectorsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *VMwareCollectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}"
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
	if vmWareCollectorName == "" {
		return nil, errors.New("parameter vmWareCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmWareCollectorName}", url.PathEscape(vmWareCollectorName))
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
func (client *VMwareCollectorsClient) deleteHandleResponse(resp *http.Response) (VMwareCollectorsDeleteResponse, error) {
	result := VMwareCollectorsDeleteResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *VMwareCollectorsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a VMware collector.
// If the operation fails it returns the *CloudError error type.
func (client *VMwareCollectorsClient) Get(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsGetOptions) (VMwareCollectorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, vmWareCollectorName, options)
	if err != nil {
		return VMwareCollectorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMwareCollectorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VMwareCollectorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VMwareCollectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, vmWareCollectorName string, options *VMwareCollectorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}"
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
	if vmWareCollectorName == "" {
		return nil, errors.New("parameter vmWareCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmWareCollectorName}", url.PathEscape(vmWareCollectorName))
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
func (client *VMwareCollectorsClient) getHandleResponse(resp *http.Response) (VMwareCollectorsGetResponse, error) {
	result := VMwareCollectorsGetResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMwareCollector); err != nil {
		return VMwareCollectorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VMwareCollectorsClient) getHandleError(resp *http.Response) error {
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

// ListByProject - Get a list of VMware collector.
// If the operation fails it returns the *CloudError error type.
func (client *VMwareCollectorsClient) ListByProject(ctx context.Context, resourceGroupName string, projectName string, options *VMwareCollectorsListByProjectOptions) (VMwareCollectorsListByProjectResponse, error) {
	req, err := client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return VMwareCollectorsListByProjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMwareCollectorsListByProjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VMwareCollectorsListByProjectResponse{}, client.listByProjectHandleError(resp)
	}
	return client.listByProjectHandleResponse(resp)
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *VMwareCollectorsClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *VMwareCollectorsListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors"
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
func (client *VMwareCollectorsClient) listByProjectHandleResponse(resp *http.Response) (VMwareCollectorsListByProjectResponse, error) {
	result := VMwareCollectorsListByProjectResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMwareCollectorList); err != nil {
		return VMwareCollectorsListByProjectResponse{}, err
	}
	return result, nil
}

// listByProjectHandleError handles the ListByProject error response.
func (client *VMwareCollectorsClient) listByProjectHandleError(resp *http.Response) error {
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
