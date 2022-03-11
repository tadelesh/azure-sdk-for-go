//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// AFDProfilesClient contains the methods for the AFDProfiles group.
// Don't use this type directly, use NewAFDProfilesClient() instead.
type AFDProfilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAFDProfilesClient creates a new instance of AFDProfilesClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAFDProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AFDProfilesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AFDProfilesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckHostNameAvailability - Validates the custom domain mapping to ensure it maps to the correct CDN endpoint in DNS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the
// resource group.
// checkHostNameAvailabilityInput - Custom domain to be validated.
// options - AFDProfilesClientCheckHostNameAvailabilityOptions contains the optional parameters for the AFDProfilesClient.CheckHostNameAvailability
// method.
func (client *AFDProfilesClient) CheckHostNameAvailability(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput CheckHostNameAvailabilityInput, options *AFDProfilesClientCheckHostNameAvailabilityOptions) (AFDProfilesClientCheckHostNameAvailabilityResponse, error) {
	req, err := client.checkHostNameAvailabilityCreateRequest(ctx, resourceGroupName, profileName, checkHostNameAvailabilityInput, options)
	if err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkHostNameAvailabilityHandleResponse(resp)
}

// checkHostNameAvailabilityCreateRequest creates the CheckHostNameAvailability request.
func (client *AFDProfilesClient) checkHostNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput CheckHostNameAvailabilityInput, options *AFDProfilesClientCheckHostNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/checkHostNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, checkHostNameAvailabilityInput)
}

// checkHostNameAvailabilityHandleResponse handles the CheckHostNameAvailability response.
func (client *AFDProfilesClient) checkHostNameAvailabilityHandleResponse(resp *http.Response) (AFDProfilesClientCheckHostNameAvailabilityResponse, error) {
	result := AFDProfilesClientCheckHostNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	return result, nil
}

// ListResourceUsage - Checks the quota and actual usage of endpoints under the given CDN profile.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the
// resource group.
// options - AFDProfilesClientListResourceUsageOptions contains the optional parameters for the AFDProfilesClient.ListResourceUsage
// method.
func (client *AFDProfilesClient) ListResourceUsage(resourceGroupName string, profileName string, options *AFDProfilesClientListResourceUsageOptions) *AFDProfilesClientListResourceUsagePager {
	return &AFDProfilesClientListResourceUsagePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listResourceUsageCreateRequest(ctx, resourceGroupName, profileName, options)
		},
		advancer: func(ctx context.Context, resp AFDProfilesClientListResourceUsageResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UsagesListResult.NextLink)
		},
	}
}

// listResourceUsageCreateRequest creates the ListResourceUsage request.
func (client *AFDProfilesClient) listResourceUsageCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *AFDProfilesClientListResourceUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listResourceUsageHandleResponse handles the ListResourceUsage response.
func (client *AFDProfilesClient) listResourceUsageHandleResponse(resp *http.Response) (AFDProfilesClientListResourceUsageResponse, error) {
	result := AFDProfilesClientListResourceUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return AFDProfilesClientListResourceUsageResponse{}, err
	}
	return result, nil
}
