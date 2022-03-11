//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ServiceTagInformationClient contains the methods for the ServiceTagInformation group.
// Don't use this type directly, use NewServiceTagInformationClient() instead.
type ServiceTagInformationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceTagInformationClient creates a new instance of ServiceTagInformationClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceTagInformationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceTagInformationClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServiceTagInformationClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Gets a list of service tag information resources with pagination.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location that will be used as a reference for cloud (not as a filter based on location, you will get the
// list of service tags with prefix details across all regions but limited to the cloud that
// your subscription belongs to).
// options - ServiceTagInformationClientListOptions contains the optional parameters for the ServiceTagInformationClient.List
// method.
func (client *ServiceTagInformationClient) List(location string, options *ServiceTagInformationClientListOptions) *ServiceTagInformationClientListPager {
	return &ServiceTagInformationClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp ServiceTagInformationClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceTagInformationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceTagInformationClient) listCreateRequest(ctx context.Context, location string, options *ServiceTagInformationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTagDetails"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	if options != nil && options.NoAddressPrefixes != nil {
		reqQP.Set("noAddressPrefixes", strconv.FormatBool(*options.NoAddressPrefixes))
	}
	if options != nil && options.TagName != nil {
		reqQP.Set("tagName", *options.TagName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceTagInformationClient) listHandleResponse(resp *http.Response) (ServiceTagInformationClientListResponse, error) {
	result := ServiceTagInformationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceTagInformationListResult); err != nil {
		return ServiceTagInformationClientListResponse{}, err
	}
	return result, nil
}
