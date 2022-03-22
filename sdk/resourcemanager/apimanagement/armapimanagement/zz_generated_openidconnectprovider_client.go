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

// OpenIDConnectProviderClient contains the methods for the OpenIDConnectProvider group.
// Don't use this type directly, use NewOpenIDConnectProviderClient() instead.
type OpenIDConnectProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOpenIDConnectProviderClient creates a new instance of OpenIDConnectProviderClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOpenIDConnectProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OpenIDConnectProviderClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OpenIDConnectProviderClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates the OpenID Connect Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// parameters - Create parameters.
// options - OpenIDConnectProviderClientCreateOrUpdateOptions contains the optional parameters for the OpenIDConnectProviderClient.CreateOrUpdate
// method.
func (client *OpenIDConnectProviderClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderContract, options *OpenIDConnectProviderClientCreateOrUpdateOptions) (OpenIDConnectProviderClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, opid, parameters, options)
	if err != nil {
		return OpenIDConnectProviderClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return OpenIDConnectProviderClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OpenIDConnectProviderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderContract, options *OpenIDConnectProviderClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
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
func (client *OpenIDConnectProviderClient) createOrUpdateHandleResponse(resp *http.Response) (OpenIDConnectProviderClientCreateOrUpdateResponse, error) {
	result := OpenIDConnectProviderClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenidConnectProviderContract); err != nil {
		return OpenIDConnectProviderClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific OpenID Connect Provider of the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - OpenIDConnectProviderClientDeleteOptions contains the optional parameters for the OpenIDConnectProviderClient.Delete
// method.
func (client *OpenIDConnectProviderClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, options *OpenIDConnectProviderClientDeleteOptions) (OpenIDConnectProviderClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, opid, ifMatch, options)
	if err != nil {
		return OpenIDConnectProviderClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return OpenIDConnectProviderClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return OpenIDConnectProviderClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OpenIDConnectProviderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, options *OpenIDConnectProviderClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
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

// Get - Gets specific OpenID Connect Provider without secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// options - OpenIDConnectProviderClientGetOptions contains the optional parameters for the OpenIDConnectProviderClient.Get
// method.
func (client *OpenIDConnectProviderClient) Get(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientGetOptions) (OpenIDConnectProviderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, opid, options)
	if err != nil {
		return OpenIDConnectProviderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OpenIDConnectProviderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OpenIDConnectProviderClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
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
func (client *OpenIDConnectProviderClient) getHandleResponse(resp *http.Response) (OpenIDConnectProviderClientGetResponse, error) {
	result := OpenIDConnectProviderClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenidConnectProviderContract); err != nil {
		return OpenIDConnectProviderClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the openIdConnectProvider specified by its identifier.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// options - OpenIDConnectProviderClientGetEntityTagOptions contains the optional parameters for the OpenIDConnectProviderClient.GetEntityTag
// method.
func (client *OpenIDConnectProviderClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientGetEntityTagOptions) (OpenIDConnectProviderClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, opid, options)
	if err != nil {
		return OpenIDConnectProviderClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *OpenIDConnectProviderClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
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
func (client *OpenIDConnectProviderClient) getEntityTagHandleResponse(resp *http.Response) (OpenIDConnectProviderClientGetEntityTagResponse, error) {
	result := OpenIDConnectProviderClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists of all the OpenId Connect Providers.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - OpenIDConnectProviderClientListByServiceOptions contains the optional parameters for the OpenIDConnectProviderClient.ListByService
// method.
func (client *OpenIDConnectProviderClient) ListByService(resourceGroupName string, serviceName string, options *OpenIDConnectProviderClientListByServiceOptions) *runtime.Pager[OpenIDConnectProviderClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[OpenIDConnectProviderClientListByServiceResponse]{
		More: func(page OpenIDConnectProviderClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenIDConnectProviderClientListByServiceResponse) (OpenIDConnectProviderClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OpenIDConnectProviderClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OpenIDConnectProviderClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OpenIDConnectProviderClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *OpenIDConnectProviderClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *OpenIDConnectProviderClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders"
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
func (client *OpenIDConnectProviderClient) listByServiceHandleResponse(resp *http.Response) (OpenIDConnectProviderClientListByServiceResponse, error) {
	result := OpenIDConnectProviderClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenIDConnectProviderCollection); err != nil {
		return OpenIDConnectProviderClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListSecrets - Gets the client secret details of the OpenID Connect Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// options - OpenIDConnectProviderClientListSecretsOptions contains the optional parameters for the OpenIDConnectProviderClient.ListSecrets
// method.
func (client *OpenIDConnectProviderClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientListSecretsOptions) (OpenIDConnectProviderClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, opid, options)
	if err != nil {
		return OpenIDConnectProviderClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OpenIDConnectProviderClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *OpenIDConnectProviderClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *OpenIDConnectProviderClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *OpenIDConnectProviderClient) listSecretsHandleResponse(resp *http.Response) (OpenIDConnectProviderClientListSecretsResponse, error) {
	result := OpenIDConnectProviderClientListSecretsResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientSecretContract); err != nil {
		return OpenIDConnectProviderClientListSecretsResponse{}, err
	}
	return result, nil
}

// Update - Updates the specific OpenID Connect Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// opid - Identifier of the OpenID Connect Provider.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update parameters.
// options - OpenIDConnectProviderClientUpdateOptions contains the optional parameters for the OpenIDConnectProviderClient.Update
// method.
func (client *OpenIDConnectProviderClient) Update(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, parameters OpenidConnectProviderUpdateContract, options *OpenIDConnectProviderClientUpdateOptions) (OpenIDConnectProviderClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, opid, ifMatch, parameters, options)
	if err != nil {
		return OpenIDConnectProviderClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenIDConnectProviderClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OpenIDConnectProviderClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *OpenIDConnectProviderClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, parameters OpenidConnectProviderUpdateContract, options *OpenIDConnectProviderClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if opid == "" {
		return nil, errors.New("parameter opid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{opid}", url.PathEscape(opid))
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
func (client *OpenIDConnectProviderClient) updateHandleResponse(resp *http.Response) (OpenIDConnectProviderClientUpdateResponse, error) {
	result := OpenIDConnectProviderClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenidConnectProviderContract); err != nil {
		return OpenIDConnectProviderClientUpdateResponse{}, err
	}
	return result, nil
}
