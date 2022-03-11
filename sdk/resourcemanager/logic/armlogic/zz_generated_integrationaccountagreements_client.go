//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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
	"strconv"
	"strings"
)

// IntegrationAccountAgreementsClient contains the methods for the IntegrationAccountAgreements group.
// Don't use this type directly, use NewIntegrationAccountAgreementsClient() instead.
type IntegrationAccountAgreementsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationAccountAgreementsClient creates a new instance of IntegrationAccountAgreementsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationAccountAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationAccountAgreementsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IntegrationAccountAgreementsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// agreementName - The integration account agreement name.
// agreement - The integration account agreement.
// options - IntegrationAccountAgreementsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountAgreementsClient.CreateOrUpdate
// method.
func (client *IntegrationAccountAgreementsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement, options *IntegrationAccountAgreementsClientCreateOrUpdateOptions) (IntegrationAccountAgreementsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, agreement, options)
	if err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountAgreementsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement, options *IntegrationAccountAgreementsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agreement)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountAgreementsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountAgreementsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// agreementName - The integration account agreement name.
// options - IntegrationAccountAgreementsClientDeleteOptions contains the optional parameters for the IntegrationAccountAgreementsClient.Delete
// method.
func (client *IntegrationAccountAgreementsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientDeleteOptions) (IntegrationAccountAgreementsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountAgreementsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationAccountAgreementsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountAgreementsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// agreementName - The integration account agreement name.
// options - IntegrationAccountAgreementsClientGetOptions contains the optional parameters for the IntegrationAccountAgreementsClient.Get
// method.
func (client *IntegrationAccountAgreementsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientGetOptions) (IntegrationAccountAgreementsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAgreementsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountAgreementsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountAgreementsClient) getHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientGetResponse, error) {
	result := IntegrationAccountAgreementsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of integration account agreements.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// options - IntegrationAccountAgreementsClientListOptions contains the optional parameters for the IntegrationAccountAgreementsClient.List
// method.
func (client *IntegrationAccountAgreementsClient) List(resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsClientListOptions) *IntegrationAccountAgreementsClientListPager {
	return &IntegrationAccountAgreementsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationAccountAgreementsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationAccountAgreementListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountAgreementsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountAgreementsClient) listHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientListResponse, error) {
	result := IntegrationAccountAgreementsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreementListResult); err != nil {
		return IntegrationAccountAgreementsClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// agreementName - The integration account agreement name.
// options - IntegrationAccountAgreementsClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountAgreementsClient.ListContentCallbackURL
// method.
func (client *IntegrationAccountAgreementsClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountAgreementsClientListContentCallbackURLOptions) (IntegrationAccountAgreementsClientListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, listContentCallbackURL, options)
	if err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountAgreementsClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, listContentCallbackURL)
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountAgreementsClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
