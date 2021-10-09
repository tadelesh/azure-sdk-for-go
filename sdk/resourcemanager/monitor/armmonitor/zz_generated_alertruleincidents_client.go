//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AlertRuleIncidentsClient contains the methods for the AlertRuleIncidents group.
// Don't use this type directly, use NewAlertRuleIncidentsClient() instead.
type AlertRuleIncidentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAlertRuleIncidentsClient creates a new instance of AlertRuleIncidentsClient with the specified values.
func NewAlertRuleIncidentsClient(con *arm.Connection, subscriptionID string) *AlertRuleIncidentsClient {
	return &AlertRuleIncidentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets an incident associated to an alert rule
// If the operation fails it returns the *ErrorResponse error type.
func (client *AlertRuleIncidentsClient) Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *AlertRuleIncidentsGetOptions) (AlertRuleIncidentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, incidentName, options)
	if err != nil {
		return AlertRuleIncidentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRuleIncidentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertRuleIncidentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertRuleIncidentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *AlertRuleIncidentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents/{incidentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if incidentName == "" {
		return nil, errors.New("parameter incidentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentName}", url.PathEscape(incidentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *AlertRuleIncidentsClient) getHandleResponse(resp *http.Response) (AlertRuleIncidentsGetResponse, error) {
	result := AlertRuleIncidentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return AlertRuleIncidentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AlertRuleIncidentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByAlertRule - Gets a list of incidents associated to an alert rule
// If the operation fails it returns a generic error.
func (client *AlertRuleIncidentsClient) ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRuleIncidentsListByAlertRuleOptions) (AlertRuleIncidentsListByAlertRuleResponse, error) {
	req, err := client.listByAlertRuleCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return AlertRuleIncidentsListByAlertRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRuleIncidentsListByAlertRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertRuleIncidentsListByAlertRuleResponse{}, client.listByAlertRuleHandleError(resp)
	}
	return client.listByAlertRuleHandleResponse(resp)
}

// listByAlertRuleCreateRequest creates the ListByAlertRule request.
func (client *AlertRuleIncidentsClient) listByAlertRuleCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRuleIncidentsListByAlertRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAlertRuleHandleResponse handles the ListByAlertRule response.
func (client *AlertRuleIncidentsClient) listByAlertRuleHandleResponse(resp *http.Response) (AlertRuleIncidentsListByAlertRuleResponse, error) {
	result := AlertRuleIncidentsListByAlertRuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentListResult); err != nil {
		return AlertRuleIncidentsListByAlertRuleResponse{}, err
	}
	return result, nil
}

// listByAlertRuleHandleError handles the ListByAlertRule error response.
func (client *AlertRuleIncidentsClient) listByAlertRuleHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
