//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MicrosoftSerialConsoleClient contains the methods for the MicrosoftSerialConsoleClient group.
// Don't use this type directly, use NewMicrosoftSerialConsoleClient() instead.
type MicrosoftSerialConsoleClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMicrosoftSerialConsoleClient creates a new instance of MicrosoftSerialConsoleClient with the specified values.
// subscriptionID - Subscription ID which uniquely identifies the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call requiring it.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMicrosoftSerialConsoleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MicrosoftSerialConsoleClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &MicrosoftSerialConsoleClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// DisableConsole - Disables the Serial Console service for all VMs and VM scale sets in the provided subscription
// If the operation fails it returns an *azcore.ResponseError type.
// defaultParam - Default parameter. Leave the value as "default".
// options - MicrosoftSerialConsoleClientDisableConsoleOptions contains the optional parameters for the MicrosoftSerialConsoleClient.DisableConsole
// method.
func (client *MicrosoftSerialConsoleClient) DisableConsole(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientDisableConsoleOptions) (MicrosoftSerialConsoleClientDisableConsoleResponse, error) {
	req, err := client.disableConsoleCreateRequest(ctx, defaultParam, options)
	if err != nil {
		return MicrosoftSerialConsoleClientDisableConsoleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MicrosoftSerialConsoleClientDisableConsoleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return MicrosoftSerialConsoleClientDisableConsoleResponse{}, runtime.NewResponseError(resp)
	}
	return client.disableConsoleHandleResponse(resp)
}

// disableConsoleCreateRequest creates the DisableConsole request.
func (client *MicrosoftSerialConsoleClient) disableConsoleCreateRequest(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientDisableConsoleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/disableConsole"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(defaultParam))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// disableConsoleHandleResponse handles the DisableConsole response.
func (client *MicrosoftSerialConsoleClient) disableConsoleHandleResponse(resp *http.Response) (MicrosoftSerialConsoleClientDisableConsoleResponse, error) {
	result := MicrosoftSerialConsoleClientDisableConsoleResponse{}
	switch resp.StatusCode {
	case http.StatusOK:
		var val DisableSerialConsoleResult
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientDisableConsoleResponse{}, err
		}
		result.Value = val
	case http.StatusNotFound:
		var val GetSerialConsoleSubscriptionNotFound
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientDisableConsoleResponse{}, err
		}
		result.Value = val
	default:
		return MicrosoftSerialConsoleClientDisableConsoleResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// EnableConsole - Enables the Serial Console service for all VMs and VM scale sets in the provided subscription
// If the operation fails it returns an *azcore.ResponseError type.
// defaultParam - Default parameter. Leave the value as "default".
// options - MicrosoftSerialConsoleClientEnableConsoleOptions contains the optional parameters for the MicrosoftSerialConsoleClient.EnableConsole
// method.
func (client *MicrosoftSerialConsoleClient) EnableConsole(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientEnableConsoleOptions) (MicrosoftSerialConsoleClientEnableConsoleResponse, error) {
	req, err := client.enableConsoleCreateRequest(ctx, defaultParam, options)
	if err != nil {
		return MicrosoftSerialConsoleClientEnableConsoleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MicrosoftSerialConsoleClientEnableConsoleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return MicrosoftSerialConsoleClientEnableConsoleResponse{}, runtime.NewResponseError(resp)
	}
	return client.enableConsoleHandleResponse(resp)
}

// enableConsoleCreateRequest creates the EnableConsole request.
func (client *MicrosoftSerialConsoleClient) enableConsoleCreateRequest(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientEnableConsoleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/enableConsole"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(defaultParam))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// enableConsoleHandleResponse handles the EnableConsole response.
func (client *MicrosoftSerialConsoleClient) enableConsoleHandleResponse(resp *http.Response) (MicrosoftSerialConsoleClientEnableConsoleResponse, error) {
	result := MicrosoftSerialConsoleClientEnableConsoleResponse{}
	switch resp.StatusCode {
	case http.StatusOK:
		var val EnableSerialConsoleResult
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientEnableConsoleResponse{}, err
		}
		result.Value = val
	case http.StatusNotFound:
		var val GetSerialConsoleSubscriptionNotFound
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientEnableConsoleResponse{}, err
		}
		result.Value = val
	default:
		return MicrosoftSerialConsoleClientEnableConsoleResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// GetConsoleStatus - Gets whether or not Serial Console is disabled for a given subscription
// If the operation fails it returns an *azcore.ResponseError type.
// defaultParam - Default parameter. Leave the value as "default".
// options - MicrosoftSerialConsoleClientGetConsoleStatusOptions contains the optional parameters for the MicrosoftSerialConsoleClient.GetConsoleStatus
// method.
func (client *MicrosoftSerialConsoleClient) GetConsoleStatus(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientGetConsoleStatusOptions) (MicrosoftSerialConsoleClientGetConsoleStatusResponse, error) {
	req, err := client.getConsoleStatusCreateRequest(ctx, defaultParam, options)
	if err != nil {
		return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getConsoleStatusHandleResponse(resp)
}

// getConsoleStatusCreateRequest creates the GetConsoleStatus request.
func (client *MicrosoftSerialConsoleClient) getConsoleStatusCreateRequest(ctx context.Context, defaultParam string, options *MicrosoftSerialConsoleClientGetConsoleStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(defaultParam))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getConsoleStatusHandleResponse handles the GetConsoleStatus response.
func (client *MicrosoftSerialConsoleClient) getConsoleStatusHandleResponse(resp *http.Response) (MicrosoftSerialConsoleClientGetConsoleStatusResponse, error) {
	result := MicrosoftSerialConsoleClientGetConsoleStatusResponse{}
	switch resp.StatusCode {
	case http.StatusOK:
		var val Status
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, err
		}
		result.Value = val
	case http.StatusNotFound:
		var val GetSerialConsoleSubscriptionNotFound
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, err
		}
		result.Value = val
	default:
		return MicrosoftSerialConsoleClientGetConsoleStatusResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// ListOperations - Gets a list of Serial Console API operations.
// If the operation fails it returns an *azcore.ResponseError type.
// options - MicrosoftSerialConsoleClientListOperationsOptions contains the optional parameters for the MicrosoftSerialConsoleClient.ListOperations
// method.
func (client *MicrosoftSerialConsoleClient) ListOperations(ctx context.Context, options *MicrosoftSerialConsoleClientListOperationsOptions) (MicrosoftSerialConsoleClientListOperationsResponse, error) {
	req, err := client.listOperationsCreateRequest(ctx, options)
	if err != nil {
		return MicrosoftSerialConsoleClientListOperationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MicrosoftSerialConsoleClientListOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MicrosoftSerialConsoleClientListOperationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listOperationsHandleResponse(resp)
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *MicrosoftSerialConsoleClient) listOperationsCreateRequest(ctx context.Context, options *MicrosoftSerialConsoleClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SerialConsole/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *MicrosoftSerialConsoleClient) listOperationsHandleResponse(resp *http.Response) (MicrosoftSerialConsoleClientListOperationsResponse, error) {
	result := MicrosoftSerialConsoleClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Operations); err != nil {
		return MicrosoftSerialConsoleClientListOperationsResponse{}, err
	}
	return result, nil
}
