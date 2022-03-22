//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// GatewayCustomDomainsClient contains the methods for the GatewayCustomDomains group.
// Don't use this type directly, use NewGatewayCustomDomainsClient() instead.
type GatewayCustomDomainsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGatewayCustomDomainsClient creates a new instance of GatewayCustomDomainsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGatewayCustomDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GatewayCustomDomainsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &GatewayCustomDomainsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or update the Spring Cloud Gateway custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// domainName - The name of the Spring Cloud Gateway custom domain.
// gatewayCustomDomainResource - The gateway custom domain resource for the create or update operation
// options - GatewayCustomDomainsClientBeginCreateOrUpdateOptions contains the optional parameters for the GatewayCustomDomainsClient.BeginCreateOrUpdate
// method.
func (client *GatewayCustomDomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource GatewayCustomDomainResource, options *GatewayCustomDomainsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[GatewayCustomDomainsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, gatewayName, domainName, gatewayCustomDomainResource, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GatewayCustomDomainsClientCreateOrUpdateResponse]("GatewayCustomDomainsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GatewayCustomDomainsClientCreateOrUpdateResponse]("GatewayCustomDomainsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the Spring Cloud Gateway custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GatewayCustomDomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource GatewayCustomDomainResource, options *GatewayCustomDomainsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, domainName, gatewayCustomDomainResource, options)
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
func (client *GatewayCustomDomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource GatewayCustomDomainResource, options *GatewayCustomDomainsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, gatewayCustomDomainResource)
}

// BeginDelete - Delete the Spring Cloud Gateway custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// domainName - The name of the Spring Cloud Gateway custom domain.
// options - GatewayCustomDomainsClientBeginDeleteOptions contains the optional parameters for the GatewayCustomDomainsClient.BeginDelete
// method.
func (client *GatewayCustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, options *GatewayCustomDomainsClientBeginDeleteOptions) (*armruntime.Poller[GatewayCustomDomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, gatewayName, domainName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GatewayCustomDomainsClientDeleteResponse]("GatewayCustomDomainsClient.Delete", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GatewayCustomDomainsClientDeleteResponse]("GatewayCustomDomainsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the Spring Cloud Gateway custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GatewayCustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, options *GatewayCustomDomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, domainName, options)
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
func (client *GatewayCustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, options *GatewayCustomDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the Spring Cloud Gateway custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// domainName - The name of the Spring Cloud Gateway custom domain.
// options - GatewayCustomDomainsClientGetOptions contains the optional parameters for the GatewayCustomDomainsClient.Get
// method.
func (client *GatewayCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, options *GatewayCustomDomainsClientGetOptions) (GatewayCustomDomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, domainName, options)
	if err != nil {
		return GatewayCustomDomainsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCustomDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayCustomDomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayCustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, options *GatewayCustomDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewayCustomDomainsClient) getHandleResponse(resp *http.Response) (GatewayCustomDomainsClientGetResponse, error) {
	result := GatewayCustomDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCustomDomainResource); err != nil {
		return GatewayCustomDomainsClientGetResponse{}, err
	}
	return result, nil
}

// List - Handle requests to list all Spring Cloud Gateway custom domains.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// options - GatewayCustomDomainsClientListOptions contains the optional parameters for the GatewayCustomDomainsClient.List
// method.
func (client *GatewayCustomDomainsClient) List(resourceGroupName string, serviceName string, gatewayName string, options *GatewayCustomDomainsClientListOptions) *runtime.Pager[GatewayCustomDomainsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[GatewayCustomDomainsClientListResponse]{
		More: func(page GatewayCustomDomainsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewayCustomDomainsClientListResponse) (GatewayCustomDomainsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GatewayCustomDomainsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GatewayCustomDomainsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GatewayCustomDomainsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GatewayCustomDomainsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewayCustomDomainsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GatewayCustomDomainsClient) listHandleResponse(resp *http.Response) (GatewayCustomDomainsClientListResponse, error) {
	result := GatewayCustomDomainsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCustomDomainResourceCollection); err != nil {
		return GatewayCustomDomainsClientListResponse{}, err
	}
	return result, nil
}
