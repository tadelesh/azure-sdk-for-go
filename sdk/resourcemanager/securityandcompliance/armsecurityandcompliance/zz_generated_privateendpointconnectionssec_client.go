//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityandcompliance

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

// PrivateEndpointConnectionsSecClient contains the methods for the PrivateEndpointConnectionsSec group.
// Don't use this type directly, use NewPrivateEndpointConnectionsSecClient() instead.
type PrivateEndpointConnectionsSecClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionsSecClient creates a new instance of PrivateEndpointConnectionsSecClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsSecClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsSecClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateEndpointConnectionsSecClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Update the state of the specified private endpoint connection associated with the service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
// properties - The private endpoint connection properties.
// options - PrivateEndpointConnectionsSecClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsSecClient.BeginCreateOrUpdate
// method.
func (client *PrivateEndpointConnectionsSecClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecClientBeginCreateOrUpdateOptions) (*armruntime.Poller[PrivateEndpointConnectionsSecClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionsSecClientCreateOrUpdateResponse]("PrivateEndpointConnectionsSecClient.CreateOrUpdate", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionsSecClientCreateOrUpdateResponse]("PrivateEndpointConnectionsSecClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Update the state of the specified private endpoint connection associated with the service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsSecClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionsSecClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// BeginDelete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
// options - PrivateEndpointConnectionsSecClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsSecClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsSecClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecClientBeginDeleteOptions) (*armruntime.Poller[PrivateEndpointConnectionsSecClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionsSecClientDeleteResponse]("PrivateEndpointConnectionsSecClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionsSecClientDeleteResponse]("PrivateEndpointConnectionsSecClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsSecClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsSecClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified private endpoint connection associated with the service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
// options - PrivateEndpointConnectionsSecClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsSecClient.Get
// method.
func (client *PrivateEndpointConnectionsSecClient) Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecClientGetOptions) (PrivateEndpointConnectionsSecClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsSecClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsSecClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsSecClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsSecClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsSecClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsSecClientGetResponse, error) {
	result := PrivateEndpointConnectionsSecClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsSecClientGetResponse{}, err
	}
	return result, nil
}

// ListByService - Lists all private endpoint connections for a service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateEndpointConnectionsSecClientListByServiceOptions contains the optional parameters for the PrivateEndpointConnectionsSecClient.ListByService
// method.
func (client *PrivateEndpointConnectionsSecClient) ListByService(resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsSecClientListByServiceOptions) *runtime.Pager[PrivateEndpointConnectionsSecClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateEndpointConnectionsSecClientListByServiceResponse]{
		More: func(page PrivateEndpointConnectionsSecClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsSecClientListByServiceResponse) (PrivateEndpointConnectionsSecClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateEndpointConnectionsSecClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateEndpointConnectionsSecClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionsSecClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PrivateEndpointConnectionsSecClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsSecClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PrivateEndpointConnectionsSecClient) listByServiceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsSecClientListByServiceResponse, error) {
	result := PrivateEndpointConnectionsSecClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsSecClientListByServiceResponse{}, err
	}
	return result, nil
}
