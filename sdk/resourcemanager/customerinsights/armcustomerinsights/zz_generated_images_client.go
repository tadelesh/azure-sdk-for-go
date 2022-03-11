//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// ImagesClient contains the methods for the Images group.
// Don't use this type directly, use NewImagesClient() instead.
type ImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewImagesClient creates a new instance of ImagesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ImagesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ImagesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetUploadURLForData - Gets data image upload URL.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// parameters - Parameters supplied to the GetUploadUrlForData operation.
// options - ImagesClientGetUploadURLForDataOptions contains the optional parameters for the ImagesClient.GetUploadURLForData
// method.
func (client *ImagesClient) GetUploadURLForData(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput, options *ImagesClientGetUploadURLForDataOptions) (ImagesClientGetUploadURLForDataResponse, error) {
	req, err := client.getUploadURLForDataCreateRequest(ctx, resourceGroupName, hubName, parameters, options)
	if err != nil {
		return ImagesClientGetUploadURLForDataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImagesClientGetUploadURLForDataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImagesClientGetUploadURLForDataResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUploadURLForDataHandleResponse(resp)
}

// getUploadURLForDataCreateRequest creates the GetUploadURLForData request.
func (client *ImagesClient) getUploadURLForDataCreateRequest(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput, options *ImagesClientGetUploadURLForDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/images/getDataImageUploadUrl"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getUploadURLForDataHandleResponse handles the GetUploadURLForData response.
func (client *ImagesClient) getUploadURLForDataHandleResponse(resp *http.Response) (ImagesClientGetUploadURLForDataResponse, error) {
	result := ImagesClientGetUploadURLForDataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageDefinition); err != nil {
		return ImagesClientGetUploadURLForDataResponse{}, err
	}
	return result, nil
}

// GetUploadURLForEntityType - Gets entity type (profile or interaction) image upload URL.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// parameters - Parameters supplied to the GetUploadUrlForEntityType operation.
// options - ImagesClientGetUploadURLForEntityTypeOptions contains the optional parameters for the ImagesClient.GetUploadURLForEntityType
// method.
func (client *ImagesClient) GetUploadURLForEntityType(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput, options *ImagesClientGetUploadURLForEntityTypeOptions) (ImagesClientGetUploadURLForEntityTypeResponse, error) {
	req, err := client.getUploadURLForEntityTypeCreateRequest(ctx, resourceGroupName, hubName, parameters, options)
	if err != nil {
		return ImagesClientGetUploadURLForEntityTypeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImagesClientGetUploadURLForEntityTypeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImagesClientGetUploadURLForEntityTypeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUploadURLForEntityTypeHandleResponse(resp)
}

// getUploadURLForEntityTypeCreateRequest creates the GetUploadURLForEntityType request.
func (client *ImagesClient) getUploadURLForEntityTypeCreateRequest(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput, options *ImagesClientGetUploadURLForEntityTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/images/getEntityTypeImageUploadUrl"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getUploadURLForEntityTypeHandleResponse handles the GetUploadURLForEntityType response.
func (client *ImagesClient) getUploadURLForEntityTypeHandleResponse(resp *http.Response) (ImagesClientGetUploadURLForEntityTypeResponse, error) {
	result := ImagesClientGetUploadURLForEntityTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageDefinition); err != nil {
		return ImagesClientGetUploadURLForEntityTypeResponse{}, err
	}
	return result, nil
}
