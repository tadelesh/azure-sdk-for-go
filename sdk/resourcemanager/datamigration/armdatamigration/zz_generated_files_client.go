//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// FilesClient contains the methods for the Files group.
// Don't use this type directly, use NewFilesClient() instead.
type FilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFilesClient creates a new instance of FilesClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FilesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &FilesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - The PUT method creates a new file or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// parameters - Information about the file
// options - FilesClientCreateOrUpdateOptions contains the optional parameters for the FilesClient.CreateOrUpdate method.
func (client *FilesClient) CreateOrUpdate(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesClientCreateOrUpdateOptions) (FilesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, projectName, fileName, parameters, options)
	if err != nil {
		return FilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FilesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FilesClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesClientCreateOrUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FilesClient) createOrUpdateHandleResponse(resp *http.Response) (FilesClientCreateOrUpdateResponse, error) {
	result := FilesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - This method deletes a file.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// options - FilesClientDeleteOptions contains the optional parameters for the FilesClient.Delete method.
func (client *FilesClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientDeleteOptions) (FilesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FilesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FilesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FilesClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientDeleteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - The files resource is a nested, proxy-only resource representing a file stored under the project resource. This method
// retrieves information about a file.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// options - FilesClientGetOptions contains the optional parameters for the FilesClient.Get method.
func (client *FilesClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientGetOptions) (FilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FilesClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FilesClient) getHandleResponse(resp *http.Response) (FilesClientGetResponse, error) {
	result := FilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesClientGetResponse{}, err
	}
	return result, nil
}

// List - The project resource is a nested resource representing a stored migration project. This method returns a list of
// files owned by a project resource.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// options - FilesClientListOptions contains the optional parameters for the FilesClient.List method.
func (client *FilesClient) List(groupName string, serviceName string, projectName string, options *FilesClientListOptions) *runtime.Pager[FilesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[FilesClientListResponse]{
		More: func(page FilesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FilesClientListResponse) (FilesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, groupName, serviceName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FilesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FilesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FilesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FilesClient) listCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, options *FilesClientListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FilesClient) listHandleResponse(resp *http.Response) (FilesClientListResponse, error) {
	result := FilesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileList); err != nil {
		return FilesClientListResponse{}, err
	}
	return result, nil
}

// Read - This method is used for requesting storage information using which contents of the file can be downloaded.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// options - FilesClientReadOptions contains the optional parameters for the FilesClient.Read method.
func (client *FilesClient) Read(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientReadOptions) (FilesClientReadResponse, error) {
	req, err := client.readCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesClientReadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientReadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientReadResponse{}, runtime.NewResponseError(resp)
	}
	return client.readHandleResponse(resp)
}

// readCreateRequest creates the Read request.
func (client *FilesClient) readCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientReadOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// readHandleResponse handles the Read response.
func (client *FilesClient) readHandleResponse(resp *http.Response) (FilesClientReadResponse, error) {
	result := FilesClientReadResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileStorageInfo); err != nil {
		return FilesClientReadResponse{}, err
	}
	return result, nil
}

// ReadWrite - This method is used for requesting information for reading and writing the file content.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// options - FilesClientReadWriteOptions contains the optional parameters for the FilesClient.ReadWrite method.
func (client *FilesClient) ReadWrite(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientReadWriteOptions) (FilesClientReadWriteResponse, error) {
	req, err := client.readWriteCreateRequest(ctx, groupName, serviceName, projectName, fileName, options)
	if err != nil {
		return FilesClientReadWriteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientReadWriteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientReadWriteResponse{}, runtime.NewResponseError(resp)
	}
	return client.readWriteHandleResponse(resp)
}

// readWriteCreateRequest creates the ReadWrite request.
func (client *FilesClient) readWriteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, options *FilesClientReadWriteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// readWriteHandleResponse handles the ReadWrite response.
func (client *FilesClient) readWriteHandleResponse(resp *http.Response) (FilesClientReadWriteResponse, error) {
	result := FilesClientReadWriteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileStorageInfo); err != nil {
		return FilesClientReadWriteResponse{}, err
	}
	return result, nil
}

// Update - This method updates an existing file.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// fileName - Name of the File
// parameters - Information about the file
// options - FilesClientUpdateOptions contains the optional parameters for the FilesClient.Update method.
func (client *FilesClient) Update(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesClientUpdateOptions) (FilesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, projectName, fileName, parameters, options)
	if err != nil {
		return FilesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FilesClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, fileName string, parameters ProjectFile, options *FilesClientUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *FilesClient) updateHandleResponse(resp *http.Response) (FilesClientUpdateResponse, error) {
	result := FilesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectFile); err != nil {
		return FilesClientUpdateResponse{}, err
	}
	return result, nil
}
