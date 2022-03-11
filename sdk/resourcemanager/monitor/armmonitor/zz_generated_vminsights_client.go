//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// VMInsightsClient contains the methods for the VMInsights group.
// Don't use this type directly, use NewVMInsightsClient() instead.
type VMInsightsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewVMInsightsClient creates a new instance of VMInsightsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVMInsightsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *VMInsightsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VMInsightsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetOnboardingStatus - Retrieves the VM Insights onboarding status for the specified resource or resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The fully qualified Azure Resource manager identifier of the resource, or scope, whose status to retrieve.
// options - VMInsightsClientGetOnboardingStatusOptions contains the optional parameters for the VMInsightsClient.GetOnboardingStatus
// method.
func (client *VMInsightsClient) GetOnboardingStatus(ctx context.Context, resourceURI string, options *VMInsightsClientGetOnboardingStatusOptions) (VMInsightsClientGetOnboardingStatusResponse, error) {
	req, err := client.getOnboardingStatusCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return VMInsightsClientGetOnboardingStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMInsightsClientGetOnboardingStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VMInsightsClientGetOnboardingStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOnboardingStatusHandleResponse(resp)
}

// getOnboardingStatusCreateRequest creates the GetOnboardingStatus request.
func (client *VMInsightsClient) getOnboardingStatusCreateRequest(ctx context.Context, resourceURI string, options *VMInsightsClientGetOnboardingStatusOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOnboardingStatusHandleResponse handles the GetOnboardingStatus response.
func (client *VMInsightsClient) getOnboardingStatusHandleResponse(resp *http.Response) (VMInsightsClientGetOnboardingStatusResponse, error) {
	result := VMInsightsClientGetOnboardingStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMInsightsOnboardingStatus); err != nil {
		return VMInsightsClientGetOnboardingStatusResponse{}, err
	}
	return result, nil
}
