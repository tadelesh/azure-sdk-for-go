//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorSimple1200Series

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

// AlertsClient contains the methods for the Alerts group.
// Don't use this type directly, use NewAlertsClient() instead.
type AlertsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertsClient creates a new instance of AlertsClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AlertsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AlertsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Clear - Clear the alerts.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name
// managerName - The manager name
// request - The clear alert request.
// options - AlertsClientClearOptions contains the optional parameters for the AlertsClient.Clear method.
func (client *AlertsClient) Clear(ctx context.Context, resourceGroupName string, managerName string, request ClearAlertRequest, options *AlertsClientClearOptions) (AlertsClientClearResponse, error) {
	req, err := client.clearCreateRequest(ctx, resourceGroupName, managerName, request, options)
	if err != nil {
		return AlertsClientClearResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientClearResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AlertsClientClearResponse{}, runtime.NewResponseError(resp)
	}
	return AlertsClientClearResponse{}, nil
}

// clearCreateRequest creates the Clear request.
func (client *AlertsClient) clearCreateRequest(ctx context.Context, resourceGroupName string, managerName string, request ClearAlertRequest, options *AlertsClientClearOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/clearAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// ListByManager - Retrieves all the alerts in a manager.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - AlertsClientListByManagerOptions contains the optional parameters for the AlertsClient.ListByManager method.
func (client *AlertsClient) ListByManager(resourceGroupName string, managerName string, options *AlertsClientListByManagerOptions) *runtime.Pager[AlertsClientListByManagerResponse] {
	return runtime.NewPager(runtime.PageProcessor[AlertsClientListByManagerResponse]{
		More: func(page AlertsClientListByManagerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertsClientListByManagerResponse) (AlertsClientListByManagerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertsClientListByManagerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertsClientListByManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertsClientListByManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagerHandleResponse(resp)
		},
	})
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *AlertsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *AlertsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *AlertsClient) listByManagerHandleResponse(resp *http.Response) (AlertsClientListByManagerResponse, error) {
	result := AlertsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertList); err != nil {
		return AlertsClientListByManagerResponse{}, err
	}
	return result, nil
}

// SendTestEmail - Sends a test alert email.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name
// managerName - The manager name
// request - The send test alert email request.
// options - AlertsClientSendTestEmailOptions contains the optional parameters for the AlertsClient.SendTestEmail method.
func (client *AlertsClient) SendTestEmail(ctx context.Context, deviceName string, resourceGroupName string, managerName string, request SendTestAlertEmailRequest, options *AlertsClientSendTestEmailOptions) (AlertsClientSendTestEmailResponse, error) {
	req, err := client.sendTestEmailCreateRequest(ctx, deviceName, resourceGroupName, managerName, request, options)
	if err != nil {
		return AlertsClientSendTestEmailResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientSendTestEmailResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AlertsClientSendTestEmailResponse{}, runtime.NewResponseError(resp)
	}
	return AlertsClientSendTestEmailResponse{}, nil
}

// sendTestEmailCreateRequest creates the SendTestEmail request.
func (client *AlertsClient) sendTestEmailCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, request SendTestAlertEmailRequest, options *AlertsClientSendTestEmailOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/sendTestAlertEmail"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}
