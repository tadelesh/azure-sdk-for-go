//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// SourceControlClient contains the methods for the SourceControl group.
// Don't use this type directly, use NewSourceControlClient() instead.
type SourceControlClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSourceControlClient creates a new instance of SourceControlClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSourceControlClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SourceControlClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SourceControlClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create a source control.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The source control name.
// parameters - The parameters supplied to the create or update source control operation.
// options - SourceControlClientCreateOrUpdateOptions contains the optional parameters for the SourceControlClient.CreateOrUpdate
// method.
func (client *SourceControlClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters, options *SourceControlClientCreateOrUpdateOptions) (SourceControlClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, parameters, options)
	if err != nil {
		return SourceControlClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SourceControlClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SourceControlClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters, options *SourceControlClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SourceControlClient) createOrUpdateHandleResponse(resp *http.Response) (SourceControlClientCreateOrUpdateResponse, error) {
	result := SourceControlClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControl); err != nil {
		return SourceControlClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the source control.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The name of source control.
// options - SourceControlClientDeleteOptions contains the optional parameters for the SourceControlClient.Delete method.
func (client *SourceControlClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlClientDeleteOptions) (SourceControlClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, options)
	if err != nil {
		return SourceControlClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SourceControlClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SourceControlClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SourceControlClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieve the source control identified by source control name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The name of source control.
// options - SourceControlClientGetOptions contains the optional parameters for the SourceControlClient.Get method.
func (client *SourceControlClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlClientGetOptions) (SourceControlClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, options)
	if err != nil {
		return SourceControlClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SourceControlClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SourceControlClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SourceControlClient) getHandleResponse(resp *http.Response) (SourceControlClientGetResponse, error) {
	result := SourceControlClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControl); err != nil {
		return SourceControlClientGetResponse{}, err
	}
	return result, nil
}

// ListByAutomationAccount - Retrieve a list of source controls.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - SourceControlClientListByAutomationAccountOptions contains the optional parameters for the SourceControlClient.ListByAutomationAccount
// method.
func (client *SourceControlClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *SourceControlClientListByAutomationAccountOptions) *runtime.Pager[SourceControlClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[SourceControlClientListByAutomationAccountResponse]{
		More: func(page SourceControlClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlClientListByAutomationAccountResponse) (SourceControlClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SourceControlClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SourceControlClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SourceControlClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *SourceControlClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SourceControlClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *SourceControlClient) listByAutomationAccountHandleResponse(resp *http.Response) (SourceControlClientListByAutomationAccountResponse, error) {
	result := SourceControlClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlListResult); err != nil {
		return SourceControlClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a source control.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The source control name.
// parameters - The parameters supplied to the update source control operation.
// options - SourceControlClientUpdateOptions contains the optional parameters for the SourceControlClient.Update method.
func (client *SourceControlClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, parameters SourceControlUpdateParameters, options *SourceControlClientUpdateOptions) (SourceControlClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, parameters, options)
	if err != nil {
		return SourceControlClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SourceControlClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SourceControlClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, parameters SourceControlUpdateParameters, options *SourceControlClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SourceControlClient) updateHandleResponse(resp *http.Response) (SourceControlClientUpdateResponse, error) {
	result := SourceControlClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControl); err != nil {
		return SourceControlClientUpdateResponse{}, err
	}
	return result, nil
}
