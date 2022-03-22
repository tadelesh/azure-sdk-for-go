//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// QuotaByCounterKeysClient contains the methods for the QuotaByCounterKeys group.
// Don't use this type directly, use NewQuotaByCounterKeysClient() instead.
type QuotaByCounterKeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewQuotaByCounterKeysClient creates a new instance of QuotaByCounterKeysClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewQuotaByCounterKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *QuotaByCounterKeysClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &QuotaByCounterKeysClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListByService - Lists a collection of current quota counter periods associated with the counter-key configured in the policy
// on the specified service instance. The api does not support paging yet.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// quotaCounterKey - Quota counter key identifier.This is the result of expression defined in counter-key attribute of the
// quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
// accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible by "ba" key
// options - QuotaByCounterKeysClientListByServiceOptions contains the optional parameters for the QuotaByCounterKeysClient.ListByService
// method.
func (client *QuotaByCounterKeysClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, options *QuotaByCounterKeysClientListByServiceOptions) (QuotaByCounterKeysClientListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, quotaCounterKey, options)
	if err != nil {
		return QuotaByCounterKeysClientListByServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotaByCounterKeysClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotaByCounterKeysClientListByServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *QuotaByCounterKeysClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, options *QuotaByCounterKeysClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if quotaCounterKey == "" {
		return nil, errors.New("parameter quotaCounterKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaCounterKey}", url.PathEscape(quotaCounterKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *QuotaByCounterKeysClient) listByServiceHandleResponse(resp *http.Response) (QuotaByCounterKeysClientListByServiceResponse, error) {
	result := QuotaByCounterKeysClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaCounterCollection); err != nil {
		return QuotaByCounterKeysClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates all the quota counter values specified with the existing quota counter key to a value in the specified
// service instance. This should be used for reset of the quota counter values.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// quotaCounterKey - Quota counter key identifier.This is the result of expression defined in counter-key attribute of the
// quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
// accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible by "ba" key
// parameters - The value of the quota counter to be applied to all quota counter periods.
// options - QuotaByCounterKeysClientUpdateOptions contains the optional parameters for the QuotaByCounterKeysClient.Update
// method.
func (client *QuotaByCounterKeysClient) Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, parameters QuotaCounterValueUpdateContract, options *QuotaByCounterKeysClientUpdateOptions) (QuotaByCounterKeysClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, quotaCounterKey, parameters, options)
	if err != nil {
		return QuotaByCounterKeysClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotaByCounterKeysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotaByCounterKeysClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *QuotaByCounterKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, parameters QuotaCounterValueUpdateContract, options *QuotaByCounterKeysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if quotaCounterKey == "" {
		return nil, errors.New("parameter quotaCounterKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaCounterKey}", url.PathEscape(quotaCounterKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *QuotaByCounterKeysClient) updateHandleResponse(resp *http.Response) (QuotaByCounterKeysClientUpdateResponse, error) {
	result := QuotaByCounterKeysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaCounterCollection); err != nil {
		return QuotaByCounterKeysClientUpdateResponse{}, err
	}
	return result, nil
}
