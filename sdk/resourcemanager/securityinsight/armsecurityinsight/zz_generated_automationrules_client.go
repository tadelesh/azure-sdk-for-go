//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsight

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

// AutomationRulesClient contains the methods for the AutomationRules group.
// Don't use this type directly, use NewAutomationRulesClient() instead.
type AutomationRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAutomationRulesClient creates a new instance of AutomationRulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAutomationRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AutomationRulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AutomationRulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates the automation rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// automationRuleID - Automation rule ID
// automationRule - The automation rule
// options - AutomationRulesClientCreateOrUpdateOptions contains the optional parameters for the AutomationRulesClient.CreateOrUpdate
// method.
func (client *AutomationRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, automationRule AutomationRule, options *AutomationRulesClientCreateOrUpdateOptions) (AutomationRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, automationRule, options)
	if err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AutomationRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AutomationRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, automationRule AutomationRule, options *AutomationRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, automationRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AutomationRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AutomationRulesClientCreateOrUpdateResponse, error) {
	result := AutomationRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRule); err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the automation rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// automationRuleID - Automation rule ID
// options - AutomationRulesClientDeleteOptions contains the optional parameters for the AutomationRulesClient.Delete method.
func (client *AutomationRulesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientDeleteOptions) (AutomationRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, options)
	if err != nil {
		return AutomationRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutomationRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AutomationRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AutomationRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AutomationRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the automation rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// automationRuleID - Automation rule ID
// options - AutomationRulesClientGetOptions contains the optional parameters for the AutomationRulesClient.Get method.
func (client *AutomationRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientGetOptions) (AutomationRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, options)
	if err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutomationRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AutomationRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AutomationRulesClient) getHandleResponse(resp *http.Response) (AutomationRulesClientGetResponse, error) {
	result := AutomationRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRule); err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all automation rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - AutomationRulesClientListOptions contains the optional parameters for the AutomationRulesClient.List method.
func (client *AutomationRulesClient) List(resourceGroupName string, workspaceName string, options *AutomationRulesClientListOptions) *runtime.Pager[AutomationRulesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AutomationRulesClientListResponse]{
		More: func(page AutomationRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AutomationRulesClientListResponse) (AutomationRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AutomationRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AutomationRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AutomationRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AutomationRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AutomationRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AutomationRulesClient) listHandleResponse(resp *http.Response) (AutomationRulesClientListResponse, error) {
	result := AutomationRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRulesList); err != nil {
		return AutomationRulesClientListResponse{}, err
	}
	return result, nil
}
