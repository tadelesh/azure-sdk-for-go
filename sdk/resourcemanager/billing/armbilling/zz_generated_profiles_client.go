//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// ProfilesClient contains the methods for the BillingProfiles group.
// Don't use this type directly, use NewProfilesClient() instead.
type ProfilesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewProfilesClient creates a new instance of ProfilesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProfilesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ProfilesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ProfilesClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a billing profile. The operation is supported for billing accounts with agreement
// type Microsoft Customer Agreement or Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// parameters - The new or updated billing profile.
// options - ProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the ProfilesClient.BeginCreateOrUpdate
// method.
func (client *ProfilesClient) BeginCreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile, options *ProfilesClientBeginCreateOrUpdateOptions) (ProfilesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return ProfilesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ProfilesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ProfilesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ProfilesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ProfilesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a billing profile. The operation is supported for billing accounts with agreement type
// Microsoft Customer Agreement or Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ProfilesClient) createOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile, options *ProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProfilesClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile, options *ProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets a billing profile by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer
// Agreement or Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// options - ProfilesClientGetOptions contains the optional parameters for the ProfilesClient.Get method.
func (client *ProfilesClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, options *ProfilesClientGetOptions) (ProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProfilesClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *ProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProfilesClient) getHandleResponse(resp *http.Response) (ProfilesClientGetResponse, error) {
	result := ProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientGetResponse{}, err
	}
	return result, nil
}

// ListByBillingAccount - Lists the billing profiles that a user has access to. The operation is supported for billing accounts
// with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// options - ProfilesClientListByBillingAccountOptions contains the optional parameters for the ProfilesClient.ListByBillingAccount
// method.
func (client *ProfilesClient) ListByBillingAccount(billingAccountName string, options *ProfilesClientListByBillingAccountOptions) *ProfilesClientListByBillingAccountPager {
	return &ProfilesClientListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
		},
		advancer: func(ctx context.Context, resp ProfilesClientListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProfileListResult.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *ProfilesClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *ProfilesClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *ProfilesClient) listByBillingAccountHandleResponse(resp *http.Response) (ProfilesClientListByBillingAccountResponse, error) {
	result := ProfilesClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListByBillingAccountResponse{}, err
	}
	return result, nil
}
