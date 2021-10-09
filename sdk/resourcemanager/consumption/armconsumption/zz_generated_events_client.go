//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

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

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
func NewEventsClient(con *arm.Connection) *EventsClient {
	return &EventsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// ListByBillingAccount - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment for a billing account or a billing profile
// for a given start and end date.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventsClient) ListByBillingAccount(billingAccountID string, options *EventsListByBillingAccountOptions) *EventsListByBillingAccountPager {
	return &EventsListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountID, options)
		},
		advancer: func(ctx context.Context, resp EventsListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Events.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *EventsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountID string, options *EventsListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *EventsClient) listByBillingAccountHandleResponse(resp *http.Response) (EventsListByBillingAccountResponse, error) {
	result := EventsListByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsListByBillingAccountResponse{}, err
	}
	return result, nil
}

// listByBillingAccountHandleError handles the ListByBillingAccount error response.
func (client *EventsClient) listByBillingAccountHandleError(resp *http.Response) error {
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

// ListByBillingProfile - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment for a billing account or a billing profile
// for a given start and end date.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventsClient) ListByBillingProfile(billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsListByBillingProfileOptions) *EventsListByBillingProfilePager {
	return &EventsListByBillingProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingProfileCreateRequest(ctx, billingAccountID, billingProfileID, startDate, endDate, options)
		},
		advancer: func(ctx context.Context, resp EventsListByBillingProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Events.NextLink)
		},
	}
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *EventsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *EventsClient) listByBillingProfileHandleResponse(resp *http.Response) (EventsListByBillingProfileResponse, error) {
	result := EventsListByBillingProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsListByBillingProfileResponse{}, err
	}
	return result, nil
}

// listByBillingProfileHandleError handles the ListByBillingProfile error response.
func (client *EventsClient) listByBillingProfileHandleError(resp *http.Response) error {
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
