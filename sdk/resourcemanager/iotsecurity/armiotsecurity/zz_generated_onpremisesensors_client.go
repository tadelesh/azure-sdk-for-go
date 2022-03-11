//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// OnPremiseSensorsClient contains the methods for the OnPremiseSensors group.
// Don't use this type directly, use NewOnPremiseSensorsClient() instead.
type OnPremiseSensorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOnPremiseSensorsClient creates a new instance of OnPremiseSensorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOnPremiseSensorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OnPremiseSensorsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OnPremiseSensorsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update on-premise IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
// onPremiseSensorName - Name of the on-premise IoT sensor
// options - OnPremiseSensorsClientCreateOrUpdateOptions contains the optional parameters for the OnPremiseSensorsClient.CreateOrUpdate
// method.
func (client *OnPremiseSensorsClient) CreateOrUpdate(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientCreateOrUpdateOptions) (OnPremiseSensorsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, onPremiseSensorName, options)
	if err != nil {
		return OnPremiseSensorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return OnPremiseSensorsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OnPremiseSensorsClient) createOrUpdateCreateRequest(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors/{onPremiseSensorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if onPremiseSensorName == "" {
		return nil, errors.New("parameter onPremiseSensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{onPremiseSensorName}", url.PathEscape(onPremiseSensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *OnPremiseSensorsClient) createOrUpdateHandleResponse(resp *http.Response) (OnPremiseSensorsClientCreateOrUpdateResponse, error) {
	result := OnPremiseSensorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnPremiseSensor); err != nil {
		return OnPremiseSensorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete on-premise IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
// onPremiseSensorName - Name of the on-premise IoT sensor
// options - OnPremiseSensorsClientDeleteOptions contains the optional parameters for the OnPremiseSensorsClient.Delete method.
func (client *OnPremiseSensorsClient) Delete(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientDeleteOptions) (OnPremiseSensorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, onPremiseSensorName, options)
	if err != nil {
		return OnPremiseSensorsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return OnPremiseSensorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return OnPremiseSensorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OnPremiseSensorsClient) deleteCreateRequest(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors/{onPremiseSensorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if onPremiseSensorName == "" {
		return nil, errors.New("parameter onPremiseSensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{onPremiseSensorName}", url.PathEscape(onPremiseSensorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// DownloadActivation - Download sensor activation file
// If the operation fails it returns an *azcore.ResponseError type.
// onPremiseSensorName - Name of the on-premise IoT sensor
// options - OnPremiseSensorsClientDownloadActivationOptions contains the optional parameters for the OnPremiseSensorsClient.DownloadActivation
// method.
func (client *OnPremiseSensorsClient) DownloadActivation(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientDownloadActivationOptions) (OnPremiseSensorsClientDownloadActivationResponse, error) {
	req, err := client.downloadActivationCreateRequest(ctx, onPremiseSensorName, options)
	if err != nil {
		return OnPremiseSensorsClientDownloadActivationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientDownloadActivationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnPremiseSensorsClientDownloadActivationResponse{}, runtime.NewResponseError(resp)
	}
	return OnPremiseSensorsClientDownloadActivationResponse{Body: resp.Body}, nil
}

// downloadActivationCreateRequest creates the DownloadActivation request.
func (client *OnPremiseSensorsClient) downloadActivationCreateRequest(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientDownloadActivationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors/{onPremiseSensorName}/downloadActivation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if onPremiseSensorName == "" {
		return nil, errors.New("parameter onPremiseSensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{onPremiseSensorName}", url.PathEscape(onPremiseSensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/zip")
	return req, nil
}

// DownloadResetPassword - Download file for reset password of the sensor
// If the operation fails it returns an *azcore.ResponseError type.
// onPremiseSensorName - Name of the on-premise IoT sensor
// body - Input for reset password.
// options - OnPremiseSensorsClientDownloadResetPasswordOptions contains the optional parameters for the OnPremiseSensorsClient.DownloadResetPassword
// method.
func (client *OnPremiseSensorsClient) DownloadResetPassword(ctx context.Context, onPremiseSensorName string, body ResetPasswordInput, options *OnPremiseSensorsClientDownloadResetPasswordOptions) (OnPremiseSensorsClientDownloadResetPasswordResponse, error) {
	req, err := client.downloadResetPasswordCreateRequest(ctx, onPremiseSensorName, body, options)
	if err != nil {
		return OnPremiseSensorsClientDownloadResetPasswordResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientDownloadResetPasswordResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnPremiseSensorsClientDownloadResetPasswordResponse{}, runtime.NewResponseError(resp)
	}
	return OnPremiseSensorsClientDownloadResetPasswordResponse{Body: resp.Body}, nil
}

// downloadResetPasswordCreateRequest creates the DownloadResetPassword request.
func (client *OnPremiseSensorsClient) downloadResetPasswordCreateRequest(ctx context.Context, onPremiseSensorName string, body ResetPasswordInput, options *OnPremiseSensorsClientDownloadResetPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors/{onPremiseSensorName}/downloadResetPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if onPremiseSensorName == "" {
		return nil, errors.New("parameter onPremiseSensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{onPremiseSensorName}", url.PathEscape(onPremiseSensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/zip")
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Get on-premise IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
// onPremiseSensorName - Name of the on-premise IoT sensor
// options - OnPremiseSensorsClientGetOptions contains the optional parameters for the OnPremiseSensorsClient.Get method.
func (client *OnPremiseSensorsClient) Get(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientGetOptions) (OnPremiseSensorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, onPremiseSensorName, options)
	if err != nil {
		return OnPremiseSensorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnPremiseSensorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OnPremiseSensorsClient) getCreateRequest(ctx context.Context, onPremiseSensorName string, options *OnPremiseSensorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors/{onPremiseSensorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if onPremiseSensorName == "" {
		return nil, errors.New("parameter onPremiseSensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{onPremiseSensorName}", url.PathEscape(onPremiseSensorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OnPremiseSensorsClient) getHandleResponse(resp *http.Response) (OnPremiseSensorsClientGetResponse, error) {
	result := OnPremiseSensorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnPremiseSensor); err != nil {
		return OnPremiseSensorsClientGetResponse{}, err
	}
	return result, nil
}

// List - List on-premise IoT sensors
// If the operation fails it returns an *azcore.ResponseError type.
// options - OnPremiseSensorsClientListOptions contains the optional parameters for the OnPremiseSensorsClient.List method.
func (client *OnPremiseSensorsClient) List(ctx context.Context, options *OnPremiseSensorsClientListOptions) (OnPremiseSensorsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OnPremiseSensorsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OnPremiseSensorsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnPremiseSensorsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OnPremiseSensorsClient) listCreateRequest(ctx context.Context, options *OnPremiseSensorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/onPremiseSensors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OnPremiseSensorsClient) listHandleResponse(resp *http.Response) (OnPremiseSensorsClientListResponse, error) {
	result := OnPremiseSensorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnPremiseSensorsList); err != nil {
		return OnPremiseSensorsClientListResponse{}, err
	}
	return result, nil
}
