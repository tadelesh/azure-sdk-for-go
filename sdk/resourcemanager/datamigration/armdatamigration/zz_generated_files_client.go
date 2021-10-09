//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// FilesClient contains the methods for the Files group.
// Don't use this type directly, use NewFilesClient() instead.
type FilesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewFilesClient creates a new instance of FilesClient with the specified values.
func NewFilesClient(con *arm.Connection, subscriptionID string) *FilesClient {
	return &FilesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - The PUT method creates a new file or updates an existing one.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) CreateOrUpdate(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesCreateOrUpdateOptions) (FilesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, projectName, fileName, parameters, options)
	if err != nil {
		return FilesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FilesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FilesClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FilesClient) createOrUpdateHandleResponse(resp *http.Response) (FilesCreateOrUpdateResponse, error) {
	result := FilesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FilesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - This method deletes a file.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesDeleteOptions) (FilesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FilesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return FilesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FilesClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *FilesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - The files resource is a nested, proxy-only resource representing a file stored under the project resource. This method retrieves information about
// a file.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesGetOptions) (FilesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FilesClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FilesClient) getHandleResponse(resp *http.Response) (FilesGetResponse, error) {
	result := FilesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FilesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - The project resource is a nested resource representing a stored migration project. This method returns a list of files owned by a project resource.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) List(groupName string, serviceName string, projectName string, options *FilesListOptions) *FilesListPager {
	return &FilesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, groupName, serviceName, projectName, options)
		},
		advancer: func(ctx context.Context, resp FilesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FileList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FilesClient) listCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, options *FilesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FilesClient) listHandleResponse(resp *http.Response) (FilesListResponse, error) {
	result := FilesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileList); err != nil {
		return FilesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FilesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Read - This method is used for requesting storage information using which contents of the file can be downloaded.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) Read(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesReadOptions) (FilesReadResponse, error) {
	req, err := client.readCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesReadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesReadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesReadResponse{}, client.readHandleError(resp)
	}
	return client.readHandleResponse(resp)
}

// readCreateRequest creates the Read request.
func (client *FilesClient) readCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesReadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}/read"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// readHandleResponse handles the Read response.
func (client *FilesClient) readHandleResponse(resp *http.Response) (FilesReadResponse, error) {
	result := FilesReadResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileStorageInfo); err != nil {
		return FilesReadResponse{}, err
	}
	return result, nil
}

// readHandleError handles the Read error response.
func (client *FilesClient) readHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ReadWrite - This method is used for requesting information for reading and writing the file content.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) ReadWrite(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesReadWriteOptions) (FilesReadWriteResponse, error) {
	req, err := client.readWriteCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesReadWriteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesReadWriteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesReadWriteResponse{}, client.readWriteHandleError(resp)
	}
	return client.readWriteHandleResponse(resp)
}

// readWriteCreateRequest creates the ReadWrite request.
func (client *FilesClient) readWriteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesReadWriteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}/readwrite"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// readWriteHandleResponse handles the ReadWrite response.
func (client *FilesClient) readWriteHandleResponse(resp *http.Response) (FilesReadWriteResponse, error) {
	result := FilesReadWriteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileStorageInfo); err != nil {
		return FilesReadWriteResponse{}, err
	}
	return result, nil
}

// readWriteHandleError handles the ReadWrite error response.
func (client *FilesClient) readWriteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - This method updates an existing file.
// If the operation fails it returns the *APIError error type.
func (client *FilesClient) Update(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesUpdateOptions) (FilesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, projectName, fileName, parameters, options)
	if err != nil {
		return FilesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FilesClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *FilesClient) updateHandleResponse(resp *http.Response) (FilesUpdateResponse, error) {
	result := FilesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *FilesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
