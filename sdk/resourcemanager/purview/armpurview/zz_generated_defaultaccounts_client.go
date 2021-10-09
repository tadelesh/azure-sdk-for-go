//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DefaultAccountsClient contains the methods for the DefaultAccounts group.
// Don't use this type directly, use NewDefaultAccountsClient() instead.
type DefaultAccountsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewDefaultAccountsClient creates a new instance of DefaultAccountsClient with the specified values.
func NewDefaultAccountsClient(con *arm.Connection) *DefaultAccountsClient {
	return &DefaultAccountsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get the default account for the scope.
// If the operation fails it returns the *ErrorResponseModel error type.
func (client *DefaultAccountsClient) Get(ctx context.Context, scopeTenantID string, scopeType ScopeType, options *DefaultAccountsGetOptions) (DefaultAccountsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scopeTenantID, scopeType, options)
	if err != nil {
		return DefaultAccountsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultAccountsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefaultAccountsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DefaultAccountsClient) getCreateRequest(ctx context.Context, scopeTenantID string, scopeType ScopeType, options *DefaultAccountsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Purview/getDefaultAccount"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("scopeTenantId", scopeTenantID)
	reqQP.Set("scopeType", string(scopeType))
	if options != nil && options.Scope != nil {
		reqQP.Set("scope", *options.Scope)
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefaultAccountsClient) getHandleResponse(resp *http.Response) (DefaultAccountsGetResponse, error) {
	result := DefaultAccountsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultAccountPayload); err != nil {
		return DefaultAccountsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DefaultAccountsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Remove - Removes the default account from the scope.
// If the operation fails it returns the *ErrorResponseModel error type.
func (client *DefaultAccountsClient) Remove(ctx context.Context, scopeTenantID string, scopeType ScopeType, options *DefaultAccountsRemoveOptions) (DefaultAccountsRemoveResponse, error) {
	req, err := client.removeCreateRequest(ctx, scopeTenantID, scopeType, options)
	if err != nil {
		return DefaultAccountsRemoveResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultAccountsRemoveResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DefaultAccountsRemoveResponse{}, client.removeHandleError(resp)
	}
	return DefaultAccountsRemoveResponse{RawResponse: resp}, nil
}

// removeCreateRequest creates the Remove request.
func (client *DefaultAccountsClient) removeCreateRequest(ctx context.Context, scopeTenantID string, scopeType ScopeType, options *DefaultAccountsRemoveOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Purview/removeDefaultAccount"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("scopeTenantId", scopeTenantID)
	reqQP.Set("scopeType", string(scopeType))
	if options != nil && options.Scope != nil {
		reqQP.Set("scope", *options.Scope)
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// removeHandleError handles the Remove error response.
func (client *DefaultAccountsClient) removeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Set - Sets the default account for the scope.
// If the operation fails it returns the *ErrorResponseModel error type.
func (client *DefaultAccountsClient) Set(ctx context.Context, defaultAccountPayload DefaultAccountPayload, options *DefaultAccountsSetOptions) (DefaultAccountsSetResponse, error) {
	req, err := client.setCreateRequest(ctx, defaultAccountPayload, options)
	if err != nil {
		return DefaultAccountsSetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultAccountsSetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefaultAccountsSetResponse{}, client.setHandleError(resp)
	}
	return client.setHandleResponse(resp)
}

// setCreateRequest creates the Set request.
func (client *DefaultAccountsClient) setCreateRequest(ctx context.Context, defaultAccountPayload DefaultAccountPayload, options *DefaultAccountsSetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Purview/setDefaultAccount"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, defaultAccountPayload)
}

// setHandleResponse handles the Set response.
func (client *DefaultAccountsClient) setHandleResponse(resp *http.Response) (DefaultAccountsSetResponse, error) {
	result := DefaultAccountsSetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultAccountPayload); err != nil {
		return DefaultAccountsSetResponse{}, err
	}
	return result, nil
}

// setHandleError handles the Set error response.
func (client *DefaultAccountsClient) setHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
