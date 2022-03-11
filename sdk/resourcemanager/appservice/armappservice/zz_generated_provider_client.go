//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// ProviderClient contains the methods for the Provider group.
// Don't use this type directly, use NewProviderClient() instead.
type ProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProviderClient creates a new instance of ProviderClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProviderClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ProviderClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetAvailableStacks - Description for Get available application frameworks and their versions
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProviderClientGetAvailableStacksOptions contains the optional parameters for the ProviderClient.GetAvailableStacks
// method.
func (client *ProviderClient) GetAvailableStacks(options *ProviderClientGetAvailableStacksOptions) *ProviderClientGetAvailableStacksPager {
	return &ProviderClientGetAvailableStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getAvailableStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetAvailableStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationStackCollection.NextLink)
		},
	}
}

// getAvailableStacksCreateRequest creates the GetAvailableStacks request.
func (client *ProviderClient) getAvailableStacksCreateRequest(ctx context.Context, options *ProviderClientGetAvailableStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/availableStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAvailableStacksHandleResponse handles the GetAvailableStacks response.
func (client *ProviderClient) getAvailableStacksHandleResponse(resp *http.Response) (ProviderClientGetAvailableStacksResponse, error) {
	result := ProviderClientGetAvailableStacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderClientGetAvailableStacksResponse{}, err
	}
	return result, nil
}

// GetAvailableStacksOnPrem - Description for Get available application frameworks and their versions
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProviderClientGetAvailableStacksOnPremOptions contains the optional parameters for the ProviderClient.GetAvailableStacksOnPrem
// method.
func (client *ProviderClient) GetAvailableStacksOnPrem(options *ProviderClientGetAvailableStacksOnPremOptions) *ProviderClientGetAvailableStacksOnPremPager {
	return &ProviderClientGetAvailableStacksOnPremPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getAvailableStacksOnPremCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetAvailableStacksOnPremResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationStackCollection.NextLink)
		},
	}
}

// getAvailableStacksOnPremCreateRequest creates the GetAvailableStacksOnPrem request.
func (client *ProviderClient) getAvailableStacksOnPremCreateRequest(ctx context.Context, options *ProviderClientGetAvailableStacksOnPremOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAvailableStacksOnPremHandleResponse handles the GetAvailableStacksOnPrem response.
func (client *ProviderClient) getAvailableStacksOnPremHandleResponse(resp *http.Response) (ProviderClientGetAvailableStacksOnPremResponse, error) {
	result := ProviderClientGetAvailableStacksOnPremResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderClientGetAvailableStacksOnPremResponse{}, err
	}
	return result, nil
}

// GetFunctionAppStacks - Description for Get available Function app frameworks and their versions
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProviderClientGetFunctionAppStacksOptions contains the optional parameters for the ProviderClient.GetFunctionAppStacks
// method.
func (client *ProviderClient) GetFunctionAppStacks(options *ProviderClientGetFunctionAppStacksOptions) *ProviderClientGetFunctionAppStacksPager {
	return &ProviderClientGetFunctionAppStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getFunctionAppStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetFunctionAppStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FunctionAppStackCollection.NextLink)
		},
	}
}

// getFunctionAppStacksCreateRequest creates the GetFunctionAppStacks request.
func (client *ProviderClient) getFunctionAppStacksCreateRequest(ctx context.Context, options *ProviderClientGetFunctionAppStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/functionAppStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFunctionAppStacksHandleResponse handles the GetFunctionAppStacks response.
func (client *ProviderClient) getFunctionAppStacksHandleResponse(resp *http.Response) (ProviderClientGetFunctionAppStacksResponse, error) {
	result := ProviderClientGetFunctionAppStacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FunctionAppStackCollection); err != nil {
		return ProviderClientGetFunctionAppStacksResponse{}, err
	}
	return result, nil
}

// GetFunctionAppStacksForLocation - Description for Get available Function app frameworks and their versions for location
// If the operation fails it returns an *azcore.ResponseError type.
// location - Function App stack location.
// options - ProviderClientGetFunctionAppStacksForLocationOptions contains the optional parameters for the ProviderClient.GetFunctionAppStacksForLocation
// method.
func (client *ProviderClient) GetFunctionAppStacksForLocation(location string, options *ProviderClientGetFunctionAppStacksForLocationOptions) *ProviderClientGetFunctionAppStacksForLocationPager {
	return &ProviderClientGetFunctionAppStacksForLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getFunctionAppStacksForLocationCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetFunctionAppStacksForLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FunctionAppStackCollection.NextLink)
		},
	}
}

// getFunctionAppStacksForLocationCreateRequest creates the GetFunctionAppStacksForLocation request.
func (client *ProviderClient) getFunctionAppStacksForLocationCreateRequest(ctx context.Context, location string, options *ProviderClientGetFunctionAppStacksForLocationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/locations/{location}/functionAppStacks"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFunctionAppStacksForLocationHandleResponse handles the GetFunctionAppStacksForLocation response.
func (client *ProviderClient) getFunctionAppStacksForLocationHandleResponse(resp *http.Response) (ProviderClientGetFunctionAppStacksForLocationResponse, error) {
	result := ProviderClientGetFunctionAppStacksForLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FunctionAppStackCollection); err != nil {
		return ProviderClientGetFunctionAppStacksForLocationResponse{}, err
	}
	return result, nil
}

// GetWebAppStacks - Description for Get available Web app frameworks and their versions
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProviderClientGetWebAppStacksOptions contains the optional parameters for the ProviderClient.GetWebAppStacks
// method.
func (client *ProviderClient) GetWebAppStacks(options *ProviderClientGetWebAppStacksOptions) *ProviderClientGetWebAppStacksPager {
	return &ProviderClientGetWebAppStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getWebAppStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetWebAppStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebAppStackCollection.NextLink)
		},
	}
}

// getWebAppStacksCreateRequest creates the GetWebAppStacks request.
func (client *ProviderClient) getWebAppStacksCreateRequest(ctx context.Context, options *ProviderClientGetWebAppStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/webAppStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWebAppStacksHandleResponse handles the GetWebAppStacks response.
func (client *ProviderClient) getWebAppStacksHandleResponse(resp *http.Response) (ProviderClientGetWebAppStacksResponse, error) {
	result := ProviderClientGetWebAppStacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppStackCollection); err != nil {
		return ProviderClientGetWebAppStacksResponse{}, err
	}
	return result, nil
}

// GetWebAppStacksForLocation - Description for Get available Web app frameworks and their versions for location
// If the operation fails it returns an *azcore.ResponseError type.
// location - Web App stack location.
// options - ProviderClientGetWebAppStacksForLocationOptions contains the optional parameters for the ProviderClient.GetWebAppStacksForLocation
// method.
func (client *ProviderClient) GetWebAppStacksForLocation(location string, options *ProviderClientGetWebAppStacksForLocationOptions) *ProviderClientGetWebAppStacksForLocationPager {
	return &ProviderClientGetWebAppStacksForLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getWebAppStacksForLocationCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientGetWebAppStacksForLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebAppStackCollection.NextLink)
		},
	}
}

// getWebAppStacksForLocationCreateRequest creates the GetWebAppStacksForLocation request.
func (client *ProviderClient) getWebAppStacksForLocationCreateRequest(ctx context.Context, location string, options *ProviderClientGetWebAppStacksForLocationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/locations/{location}/webAppStacks"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWebAppStacksForLocationHandleResponse handles the GetWebAppStacksForLocation response.
func (client *ProviderClient) getWebAppStacksForLocationHandleResponse(resp *http.Response) (ProviderClientGetWebAppStacksForLocationResponse, error) {
	result := ProviderClientGetWebAppStacksForLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppStackCollection); err != nil {
		return ProviderClientGetWebAppStacksForLocationResponse{}, err
	}
	return result, nil
}

// ListOperations - Description for Gets all available operations for the Microsoft.Web resource provider. Also exposes resource
// metric definitions
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProviderClientListOperationsOptions contains the optional parameters for the ProviderClient.ListOperations method.
func (client *ProviderClient) ListOperations(options *ProviderClientListOperationsOptions) *ProviderClientListOperationsPager {
	return &ProviderClientListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderClientListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CsmOperationCollection.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *ProviderClient) listOperationsCreateRequest(ctx context.Context, options *ProviderClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *ProviderClient) listOperationsHandleResponse(resp *http.Response) (ProviderClientListOperationsResponse, error) {
	result := ProviderClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return ProviderClientListOperationsResponse{}, err
	}
	return result, nil
}
