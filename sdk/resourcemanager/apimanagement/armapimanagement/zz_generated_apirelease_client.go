//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// APIReleaseClient contains the methods for the APIRelease group.
// Don't use this type directly, use NewAPIReleaseClient() instead.
type APIReleaseClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIReleaseClient creates a new instance of APIReleaseClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIReleaseClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *APIReleaseClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &APIReleaseClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates a new Release for the API.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// releaseID - Release identifier within an API. Must be unique in the current API Management service instance.
// parameters - Create parameters.
// options - APIReleaseClientCreateOrUpdateOptions contains the optional parameters for the APIReleaseClient.CreateOrUpdate
// method.
func (client *APIReleaseClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, parameters APIReleaseContract, options *APIReleaseClientCreateOrUpdateOptions) (APIReleaseClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, parameters, options)
	if err != nil {
		return APIReleaseClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIReleaseClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return APIReleaseClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIReleaseClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, parameters APIReleaseContract, options *APIReleaseClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIReleaseClient) createOrUpdateHandleResponse(resp *http.Response) (APIReleaseClientCreateOrUpdateResponse, error) {
	result := APIReleaseClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIReleaseContract); err != nil {
		return APIReleaseClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified release in the API.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// releaseID - Release identifier within an API. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - APIReleaseClientDeleteOptions contains the optional parameters for the APIReleaseClient.Delete method.
func (client *APIReleaseClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, options *APIReleaseClientDeleteOptions) (APIReleaseClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, ifMatch, options)
	if err != nil {
		return APIReleaseClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIReleaseClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APIReleaseClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return APIReleaseClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIReleaseClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, options *APIReleaseClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the details of an API release.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// releaseID - Release identifier within an API. Must be unique in the current API Management service instance.
// options - APIReleaseClientGetOptions contains the optional parameters for the APIReleaseClient.Get method.
func (client *APIReleaseClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseClientGetOptions) (APIReleaseClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, options)
	if err != nil {
		return APIReleaseClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIReleaseClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIReleaseClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIReleaseClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIReleaseClient) getHandleResponse(resp *http.Response) (APIReleaseClientGetResponse, error) {
	result := APIReleaseClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIReleaseContract); err != nil {
		return APIReleaseClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Returns the etag of an API release.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// releaseID - Release identifier within an API. Must be unique in the current API Management service instance.
// options - APIReleaseClientGetEntityTagOptions contains the optional parameters for the APIReleaseClient.GetEntityTag method.
func (client *APIReleaseClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseClientGetEntityTagOptions) (APIReleaseClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, options)
	if err != nil {
		return APIReleaseClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIReleaseClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIReleaseClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIReleaseClient) getEntityTagHandleResponse(resp *http.Response) (APIReleaseClientGetEntityTagResponse, error) {
	result := APIReleaseClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists all releases of an API. An API release is created when making an API Revision current. Releases are
// also used to rollback to previous revisions. Results will be paged and can be constrained by
// the $top and $skip parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// options - APIReleaseClientListByServiceOptions contains the optional parameters for the APIReleaseClient.ListByService
// method.
func (client *APIReleaseClient) ListByService(resourceGroupName string, serviceName string, apiID string, options *APIReleaseClientListByServiceOptions) *runtime.Pager[APIReleaseClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[APIReleaseClientListByServiceResponse]{
		More: func(page APIReleaseClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIReleaseClientListByServiceResponse) (APIReleaseClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIReleaseClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIReleaseClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIReleaseClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIReleaseClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIReleaseClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIReleaseClient) listByServiceHandleResponse(resp *http.Response) (APIReleaseClientListByServiceResponse, error) {
	result := APIReleaseClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIReleaseCollection); err != nil {
		return APIReleaseClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the release of the API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// releaseID - Release identifier within an API. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - API Release Update parameters.
// options - APIReleaseClientUpdateOptions contains the optional parameters for the APIReleaseClient.Update method.
func (client *APIReleaseClient) Update(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, parameters APIReleaseContract, options *APIReleaseClientUpdateOptions) (APIReleaseClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, ifMatch, parameters, options)
	if err != nil {
		return APIReleaseClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIReleaseClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIReleaseClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *APIReleaseClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, parameters APIReleaseContract, options *APIReleaseClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *APIReleaseClient) updateHandleResponse(resp *http.Response) (APIReleaseClientUpdateResponse, error) {
	result := APIReleaseClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIReleaseContract); err != nil {
		return APIReleaseClientUpdateResponse{}, err
	}
	return result, nil
}
