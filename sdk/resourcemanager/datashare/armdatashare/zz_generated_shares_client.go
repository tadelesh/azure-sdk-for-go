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

// SharesClient contains the methods for the Shares group.
// Don't use this type directly, use NewSharesClient() instead.
type SharesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSharesClient creates a new instance of SharesClient with the specified values.
func NewSharesClient(con *arm.Connection, subscriptionID string) *SharesClient {
	return &SharesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Create a share
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share, options *SharesCreateOptions) (SharesCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareName, share, options)
	if err != nil {
		return SharesCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharesCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SharesCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SharesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share, options *SharesCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, share)
}

// createHandleResponse handles the Create response.
func (client *SharesClient) createHandleResponse(resp *http.Response) (SharesCreateResponse, error) {
	result := SharesCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Share); err != nil {
		return SharesCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *SharesClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Delete a share
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesBeginDeleteOptions) (SharesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return SharesDeletePollerResponse{}, err
	}
	result := SharesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SharesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SharesDeletePollerResponse{}, err
	}
	result.Poller = &SharesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a share
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SharesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SharesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a share
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesGetOptions) (SharesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return SharesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharesClient) getHandleResponse(resp *http.Response) (SharesGetResponse, error) {
	result := SharesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Share); err != nil {
		return SharesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SharesClient) getHandleError(resp *http.Response) error {
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

// ListByAccount - List shares in an account
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) ListByAccount(resourceGroupName string, accountName string, options *SharesListByAccountOptions) *SharesListByAccountPager {
	return &SharesListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp SharesListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ShareList.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *SharesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SharesListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *SharesClient) listByAccountHandleResponse(resp *http.Response) (SharesListByAccountResponse, error) {
	result := SharesListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShareList); err != nil {
		return SharesListByAccountResponse{}, err
	}
	return result, nil
}

// listByAccountHandleError handles the ListByAccount error response.
func (client *SharesClient) listByAccountHandleError(resp *http.Response) error {
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

// ListSynchronizationDetails - List synchronization details
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) ListSynchronizationDetails(resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, options *SharesListSynchronizationDetailsOptions) *SharesListSynchronizationDetailsPager {
	return &SharesListSynchronizationDetailsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSynchronizationDetailsCreateRequest(ctx, resourceGroupName, accountName, shareName, shareSynchronization, options)
		},
		advancer: func(ctx context.Context, resp SharesListSynchronizationDetailsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SynchronizationDetailsList.NextLink)
		},
	}
}

// listSynchronizationDetailsCreateRequest creates the ListSynchronizationDetails request.
func (client *SharesClient) listSynchronizationDetailsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, options *SharesListSynchronizationDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizationDetails"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, shareSynchronization)
}

// listSynchronizationDetailsHandleResponse handles the ListSynchronizationDetails response.
func (client *SharesClient) listSynchronizationDetailsHandleResponse(resp *http.Response) (SharesListSynchronizationDetailsResponse, error) {
	result := SharesListSynchronizationDetailsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SynchronizationDetailsList); err != nil {
		return SharesListSynchronizationDetailsResponse{}, err
	}
	return result, nil
}

// listSynchronizationDetailsHandleError handles the ListSynchronizationDetails error response.
func (client *SharesClient) listSynchronizationDetailsHandleError(resp *http.Response) error {
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

// ListSynchronizations - List synchronizations of a share
// If the operation fails it returns the *DataShareError error type.
func (client *SharesClient) ListSynchronizations(resourceGroupName string, accountName string, shareName string, options *SharesListSynchronizationsOptions) *SharesListSynchronizationsPager {
	return &SharesListSynchronizationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSynchronizationsCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
		},
		advancer: func(ctx context.Context, resp SharesListSynchronizationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ShareSynchronizationList.NextLink)
		},
	}
}

// listSynchronizationsCreateRequest creates the ListSynchronizations request.
func (client *SharesClient) listSynchronizationsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesListSynchronizationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizations"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSynchronizationsHandleResponse handles the ListSynchronizations response.
func (client *SharesClient) listSynchronizationsHandleResponse(resp *http.Response) (SharesListSynchronizationsResponse, error) {
	result := SharesListSynchronizationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShareSynchronizationList); err != nil {
		return SharesListSynchronizationsResponse{}, err
	}
	return result, nil
}

// listSynchronizationsHandleError handles the ListSynchronizations error response.
func (client *SharesClient) listSynchronizationsHandleError(resp *http.Response) error {
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
