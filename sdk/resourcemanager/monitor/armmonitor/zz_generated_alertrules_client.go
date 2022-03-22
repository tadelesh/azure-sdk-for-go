//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// AlertRulesClient contains the methods for the AlertRules group.
// Don't use this type directly, use NewAlertRulesClient() instead.
type AlertRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertRulesClient creates a new instance of AlertRulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AlertRulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AlertRulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates a classic metric alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// parameters - The parameters of the rule to create or update.
// options - AlertRulesClientCreateOrUpdateOptions contains the optional parameters for the AlertRulesClient.CreateOrUpdate
// method.
func (client *AlertRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters AlertRuleResource, options *AlertRulesClientCreateOrUpdateOptions) (AlertRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AlertRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AlertRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters AlertRuleResource, options *AlertRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/alertrules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AlertRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AlertRulesClientCreateOrUpdateResponse, error) {
	result := AlertRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResource); err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a classic metric alert rule
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// options - AlertRulesClientDeleteOptions contains the optional parameters for the AlertRulesClient.Delete method.
func (client *AlertRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesClientDeleteOptions) (AlertRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return AlertRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AlertRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AlertRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AlertRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/alertrules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a classic metric alert rule
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// options - AlertRulesClientGetOptions contains the optional parameters for the AlertRulesClient.Get method.
func (client *AlertRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesClientGetOptions) (AlertRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/alertrules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertRulesClient) getHandleResponse(resp *http.Response) (AlertRulesClientGetResponse, error) {
	result := AlertRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResource); err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List the classic metric alert rules within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - AlertRulesClientListByResourceGroupOptions contains the optional parameters for the AlertRulesClient.ListByResourceGroup
// method.
func (client *AlertRulesClient) ListByResourceGroup(resourceGroupName string, options *AlertRulesClientListByResourceGroupOptions) *runtime.Pager[AlertRulesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[AlertRulesClientListByResourceGroupResponse]{
		More: func(page AlertRulesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AlertRulesClientListByResourceGroupResponse) (AlertRulesClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return AlertRulesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertRulesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertRulesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AlertRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AlertRulesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/alertrules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AlertRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (AlertRulesClientListByResourceGroupResponse, error) {
	result := AlertRulesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResourceCollection); err != nil {
		return AlertRulesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - List the classic metric alert rules within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AlertRulesClientListBySubscriptionOptions contains the optional parameters for the AlertRulesClient.ListBySubscription
// method.
func (client *AlertRulesClient) ListBySubscription(options *AlertRulesClientListBySubscriptionOptions) *runtime.Pager[AlertRulesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[AlertRulesClientListBySubscriptionResponse]{
		More: func(page AlertRulesClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AlertRulesClientListBySubscriptionResponse) (AlertRulesClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return AlertRulesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertRulesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertRulesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AlertRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AlertRulesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/alertrules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AlertRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (AlertRulesClientListBySubscriptionResponse, error) {
	result := AlertRulesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResourceCollection); err != nil {
		return AlertRulesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing classic metric AlertRuleResource. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// alertRulesResource - Parameters supplied to the operation.
// options - AlertRulesClientUpdateOptions contains the optional parameters for the AlertRulesClient.Update method.
func (client *AlertRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource AlertRuleResourcePatch, options *AlertRulesClientUpdateOptions) (AlertRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleName, alertRulesResource, options)
	if err != nil {
		return AlertRulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AlertRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AlertRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource AlertRuleResourcePatch, options *AlertRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/alertrules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, alertRulesResource)
}

// updateHandleResponse handles the Update response.
func (client *AlertRulesClient) updateHandleResponse(resp *http.Response) (AlertRulesClientUpdateResponse, error) {
	result := AlertRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResource); err != nil {
		return AlertRulesClientUpdateResponse{}, err
	}
	return result, nil
}
