//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// SharedGalleryImagesClient contains the methods for the SharedGalleryImages group.
// Don't use this type directly, use NewSharedGalleryImagesClient() instead.
type SharedGalleryImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSharedGalleryImagesClient creates a new instance of SharedGalleryImagesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSharedGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SharedGalleryImagesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SharedGalleryImagesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get a shared gallery image by subscription id or tenant id.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Resource location.
// galleryUniqueName - The unique name of the Shared Gallery.
// galleryImageName - The name of the Shared Gallery Image Definition from which the Image Versions are to be listed.
// options - SharedGalleryImagesClientGetOptions contains the optional parameters for the SharedGalleryImagesClient.Get method.
func (client *SharedGalleryImagesClient) Get(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImagesClientGetOptions) (SharedGalleryImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, galleryUniqueName, galleryImageName, options)
	if err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedGalleryImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedGalleryImagesClient) getCreateRequest(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedGalleryImagesClient) getHandleResponse(resp *http.Response) (SharedGalleryImagesClientGetResponse, error) {
	result := SharedGalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImage); err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

// List - List shared gallery images by subscription id or tenant id.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Resource location.
// galleryUniqueName - The unique name of the Shared Gallery.
// options - SharedGalleryImagesClientListOptions contains the optional parameters for the SharedGalleryImagesClient.List
// method.
func (client *SharedGalleryImagesClient) List(location string, galleryUniqueName string, options *SharedGalleryImagesClientListOptions) *SharedGalleryImagesClientListPager {
	return &SharedGalleryImagesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, galleryUniqueName, options)
		},
		advancer: func(ctx context.Context, resp SharedGalleryImagesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SharedGalleryImageList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SharedGalleryImagesClient) listCreateRequest(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleryImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.SharedTo != nil {
		reqQP.Set("sharedTo", string(*options.SharedTo))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedGalleryImagesClient) listHandleResponse(resp *http.Response) (SharedGalleryImagesClientListResponse, error) {
	result := SharedGalleryImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImageList); err != nil {
		return SharedGalleryImagesClientListResponse{}, err
	}
	return result, nil
}
