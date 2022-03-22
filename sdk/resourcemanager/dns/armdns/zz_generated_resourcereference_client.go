//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

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

// ResourceReferenceClient contains the methods for the DNSResourceReference group.
// Don't use this type directly, use NewResourceReferenceClient() instead.
type ResourceReferenceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceReferenceClient creates a new instance of ResourceReferenceClient with the specified values.
// subscriptionID - Specifies the Azure subscription ID, which uniquely identifies the Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceReferenceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceReferenceClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ResourceReferenceClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetByTargetResources - Returns the DNS records specified by the referencing targetResourceIds.
// If the operation fails it returns an *azcore.ResponseError type.
// parameters - Properties for dns resource reference request.
// options - ResourceReferenceClientGetByTargetResourcesOptions contains the optional parameters for the ResourceReferenceClient.GetByTargetResources
// method.
func (client *ResourceReferenceClient) GetByTargetResources(ctx context.Context, parameters ResourceReferenceRequest, options *ResourceReferenceClientGetByTargetResourcesOptions) (ResourceReferenceClientGetByTargetResourcesResponse, error) {
	req, err := client.getByTargetResourcesCreateRequest(ctx, parameters, options)
	if err != nil {
		return ResourceReferenceClientGetByTargetResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceReferenceClientGetByTargetResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceReferenceClientGetByTargetResourcesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByTargetResourcesHandleResponse(resp)
}

// getByTargetResourcesCreateRequest creates the GetByTargetResources request.
func (client *ResourceReferenceClient) getByTargetResourcesCreateRequest(ctx context.Context, parameters ResourceReferenceRequest, options *ResourceReferenceClientGetByTargetResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/getDnsResourceReference"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getByTargetResourcesHandleResponse handles the GetByTargetResources response.
func (client *ResourceReferenceClient) getByTargetResourcesHandleResponse(resp *http.Response) (ResourceReferenceClientGetByTargetResourcesResponse, error) {
	result := ResourceReferenceClientGetByTargetResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceReferenceResult); err != nil {
		return ResourceReferenceClientGetByTargetResourcesResponse{}, err
	}
	return result, nil
}
