//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CommunityGalleryImagesClient contains the methods for the CommunityGalleryImages group.
// Don't use this type directly, use NewCommunityGalleryImagesClient() instead.
type CommunityGalleryImagesClient struct {
	host string
	subscriptionID string
	pl runtime.Pipeline
}

// NewCommunityGalleryImagesClient creates a new instance of CommunityGalleryImagesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCommunityGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunityGalleryImagesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CommunityGalleryImagesClient{
		subscriptionID: subscriptionID,
		host: ep,
pl: pl,
	}
	return client, nil
}

// Get - Get a community gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Resource location.
// publicGalleryName - The public name of the community gallery.
// galleryImageName - The name of the community gallery image definition.
// options - CommunityGalleryImagesClientGetOptions contains the optional parameters for the CommunityGalleryImagesClient.Get
// method.
func (client *CommunityGalleryImagesClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string, options *CommunityGalleryImagesClientGetOptions) (CommunityGalleryImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publicGalleryName, galleryImageName, options)
	if err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunityGalleryImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommunityGalleryImagesClient) getCreateRequest(ctx context.Context, location string, publicGalleryName string, galleryImageName string, options *CommunityGalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
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
func (client *CommunityGalleryImagesClient) getHandleResponse(resp *http.Response) (CommunityGalleryImagesClientGetResponse, error) {
	result := CommunityGalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGalleryImage); err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

