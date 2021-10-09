//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ProviderShareSubscriptionsClient contains the methods for the ProviderShareSubscriptions group.
// Don't use this type directly, use NewProviderShareSubscriptionsClient() instead.
type ProviderShareSubscriptionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProviderShareSubscriptionsClient creates a new instance of ProviderShareSubscriptionsClient with the specified values.
func NewProviderShareSubscriptionsClient(con *arm.Connection, subscriptionID string) *ProviderShareSubscriptionsClient {
	return &ProviderShareSubscriptionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Adjust - Adjust a share subscription's expiration date in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) Adjust(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsAdjustOptions) (ProviderShareSubscriptionsAdjustResponse, error) {
	req, err := client.adjustCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, providerShareSubscription, options)
	if err != nil {
		return ProviderShareSubscriptionsAdjustResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsAdjustResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsAdjustResponse{}, client.adjustHandleError(resp)
	}
	return client.adjustHandleResponse(resp)
}

// adjustCreateRequest creates the Adjust request.
func (client *ProviderShareSubscriptionsClient) adjustCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsAdjustOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/adjust"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, providerShareSubscription)
}

// adjustHandleResponse handles the Adjust response.
func (client *ProviderShareSubscriptionsClient) adjustHandleResponse(resp *http.Response) (ProviderShareSubscriptionsAdjustResponse, error) {
	result := ProviderShareSubscriptionsAdjustResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsAdjustResponse{}, err
	}
	return result, nil
}

// adjustHandleError handles the Adjust error response.
func (client *ProviderShareSubscriptionsClient) adjustHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByShare - Get share subscription in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) GetByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsGetByShareOptions) (ProviderShareSubscriptionsGetByShareResponse, error) {
	req, err := client.getByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
	if err != nil {
		return ProviderShareSubscriptionsGetByShareResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsGetByShareResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsGetByShareResponse{}, client.getByShareHandleError(resp)
	}
	return client.getByShareHandleResponse(resp)
}

// getByShareCreateRequest creates the GetByShare request.
func (client *ProviderShareSubscriptionsClient) getByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsGetByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByShareHandleResponse handles the GetByShare response.
func (client *ProviderShareSubscriptionsClient) getByShareHandleResponse(resp *http.Response) (ProviderShareSubscriptionsGetByShareResponse, error) {
	result := ProviderShareSubscriptionsGetByShareResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsGetByShareResponse{}, err
	}
	return result, nil
}

// getByShareHandleError handles the GetByShare error response.
func (client *ProviderShareSubscriptionsClient) getByShareHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByShare - List share subscriptions in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) ListByShare(resourceGroupName string, accountName string, shareName string, options *ProviderShareSubscriptionsListByShareOptions) *ProviderShareSubscriptionsListBySharePager {
	return &ProviderShareSubscriptionsListBySharePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
		},
		advancer: func(ctx context.Context, resp ProviderShareSubscriptionsListByShareResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProviderShareSubscriptionList.NextLink)
		},
	}
}

// listByShareCreateRequest creates the ListByShare request.
func (client *ProviderShareSubscriptionsClient) listByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *ProviderShareSubscriptionsListByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByShareHandleResponse handles the ListByShare response.
func (client *ProviderShareSubscriptionsClient) listByShareHandleResponse(resp *http.Response) (ProviderShareSubscriptionsListByShareResponse, error) {
	result := ProviderShareSubscriptionsListByShareResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscriptionList); err != nil {
		return ProviderShareSubscriptionsListByShareResponse{}, err
	}
	return result, nil
}

// listByShareHandleError handles the ListByShare error response.
func (client *ProviderShareSubscriptionsClient) listByShareHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Reinstate - Reinstate share subscription in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) Reinstate(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsReinstateOptions) (ProviderShareSubscriptionsReinstateResponse, error) {
	req, err := client.reinstateCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, providerShareSubscription, options)
	if err != nil {
		return ProviderShareSubscriptionsReinstateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderShareSubscriptionsReinstateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderShareSubscriptionsReinstateResponse{}, client.reinstateHandleError(resp)
	}
	return client.reinstateHandleResponse(resp)
}

// reinstateCreateRequest creates the Reinstate request.
func (client *ProviderShareSubscriptionsClient) reinstateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, providerShareSubscription ProviderShareSubscription, options *ProviderShareSubscriptionsReinstateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/reinstate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, providerShareSubscription)
}

// reinstateHandleResponse handles the Reinstate response.
func (client *ProviderShareSubscriptionsClient) reinstateHandleResponse(resp *http.Response) (ProviderShareSubscriptionsReinstateResponse, error) {
	result := ProviderShareSubscriptionsReinstateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderShareSubscription); err != nil {
		return ProviderShareSubscriptionsReinstateResponse{}, err
	}
	return result, nil
}

// reinstateHandleError handles the Reinstate error response.
func (client *ProviderShareSubscriptionsClient) reinstateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRevoke - Revoke share subscription in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) BeginRevoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsBeginRevokeOptions) (ProviderShareSubscriptionsRevokePollerResponse, error) {
	resp, err := client.revoke(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
	if err != nil {
		return ProviderShareSubscriptionsRevokePollerResponse{}, err
	}
	result := ProviderShareSubscriptionsRevokePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ProviderShareSubscriptionsClient.Revoke", "azure-async-operation", resp, client.pl, client.revokeHandleError)
	if err != nil {
		return ProviderShareSubscriptionsRevokePollerResponse{}, err
	}
	result.Poller = &ProviderShareSubscriptionsRevokePoller{
		pt: pt,
	}
	return result, nil
}

// Revoke - Revoke share subscription in a provider share
// If the operation fails it returns the *DataShareError error type.
func (client *ProviderShareSubscriptionsClient) revoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsBeginRevokeOptions) (*http.Response, error) {
	req, err := client.revokeCreateRequest(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.revokeHandleError(resp)
	}
	return resp, nil
}

// revokeCreateRequest creates the Revoke request.
func (client *ProviderShareSubscriptionsClient) revokeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string, options *ProviderShareSubscriptionsBeginRevokeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/revoke"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if providerShareSubscriptionID == "" {
		return nil, errors.New("parameter providerShareSubscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerShareSubscriptionId}", url.PathEscape(providerShareSubscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// revokeHandleError handles the Revoke error response.
func (client *ProviderShareSubscriptionsClient) revokeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
