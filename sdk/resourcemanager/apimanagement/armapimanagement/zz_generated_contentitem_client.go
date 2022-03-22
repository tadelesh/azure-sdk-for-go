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
	"strings"
)

// ContentItemClient contains the methods for the ContentItem group.
// Don't use this type directly, use NewContentItemClient() instead.
type ContentItemClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContentItemClient creates a new instance of ContentItemClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContentItemClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContentItemClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ContentItemClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates a new developer portal's content item specified by the provided content type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// contentItemID - Content item identifier.
// options - ContentItemClientCreateOrUpdateOptions contains the optional parameters for the ContentItemClient.CreateOrUpdate
// method.
func (client *ContentItemClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientCreateOrUpdateOptions) (ContentItemClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentItemClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ContentItemClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContentItemClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
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
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContentItemClient) createOrUpdateHandleResponse(resp *http.Response) (ContentItemClientCreateOrUpdateResponse, error) {
	result := ContentItemClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentItemContract); err != nil {
		return ContentItemClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the specified developer portal's content item.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// contentItemID - Content item identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - ContentItemClientDeleteOptions contains the optional parameters for the ContentItemClient.Delete method.
func (client *ContentItemClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, ifMatch string, options *ContentItemClientDeleteOptions) (ContentItemClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, ifMatch, options)
	if err != nil {
		return ContentItemClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentItemClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ContentItemClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ContentItemClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentItemClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, ifMatch string, options *ContentItemClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
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

// Get - Returns the developer portal's content item specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// contentItemID - Content item identifier.
// options - ContentItemClientGetOptions contains the optional parameters for the ContentItemClient.Get method.
func (client *ContentItemClient) Get(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientGetOptions) (ContentItemClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentItemClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContentItemClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContentItemClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
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
func (client *ContentItemClient) getHandleResponse(resp *http.Response) (ContentItemClientGetResponse, error) {
	result := ContentItemClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentItemContract); err != nil {
		return ContentItemClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Returns the entity state (ETag) version of the developer portal's content item specified by its identifier.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// contentItemID - Content item identifier.
// options - ContentItemClientGetEntityTagOptions contains the optional parameters for the ContentItemClient.GetEntityTag
// method.
func (client *ContentItemClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientGetEntityTagOptions) (ContentItemClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentItemClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *ContentItemClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
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
func (client *ContentItemClient) getEntityTagHandleResponse(resp *http.Response) (ContentItemClientGetEntityTagResponse, error) {
	result := ContentItemClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists developer portal's content items specified by the provided content type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// options - ContentItemClientListByServiceOptions contains the optional parameters for the ContentItemClient.ListByService
// method.
func (client *ContentItemClient) ListByService(resourceGroupName string, serviceName string, contentTypeID string, options *ContentItemClientListByServiceOptions) *runtime.Pager[ContentItemClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[ContentItemClientListByServiceResponse]{
		More: func(page ContentItemClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContentItemClientListByServiceResponse) (ContentItemClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContentItemClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContentItemClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContentItemClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *ContentItemClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentItemClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
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

// listByServiceHandleResponse handles the ListByService response.
func (client *ContentItemClient) listByServiceHandleResponse(resp *http.Response) (ContentItemClientListByServiceResponse, error) {
	result := ContentItemClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentItemCollection); err != nil {
		return ContentItemClientListByServiceResponse{}, err
	}
	return result, nil
}
