//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

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

// AzureReservationAPIClient contains the methods for the AzureReservationAPI group.
// Don't use this type directly, use NewAzureReservationAPIClient() instead.
type AzureReservationAPIClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAzureReservationAPIClient creates a new instance of AzureReservationAPIClient with the specified values.
func NewAzureReservationAPIClient(con *arm.Connection) *AzureReservationAPIClient {
	return &AzureReservationAPIClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// GetAppliedReservationList - Get applicable Reservations that are applied to this subscription or a resource group under this subscription.
// If the operation fails it returns the *Error error type.
func (client *AzureReservationAPIClient) GetAppliedReservationList(ctx context.Context, subscriptionID string, options *AzureReservationAPIGetAppliedReservationListOptions) (AzureReservationAPIGetAppliedReservationListResponse, error) {
	req, err := client.getAppliedReservationListCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return AzureReservationAPIGetAppliedReservationListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureReservationAPIGetAppliedReservationListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureReservationAPIGetAppliedReservationListResponse{}, client.getAppliedReservationListHandleError(resp)
	}
	return client.getAppliedReservationListHandleResponse(resp)
}

// getAppliedReservationListCreateRequest creates the GetAppliedReservationList request.
func (client *AzureReservationAPIClient) getAppliedReservationListCreateRequest(ctx context.Context, subscriptionID string, options *AzureReservationAPIGetAppliedReservationListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/appliedReservations"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAppliedReservationListHandleResponse handles the GetAppliedReservationList response.
func (client *AzureReservationAPIClient) getAppliedReservationListHandleResponse(resp *http.Response) (AzureReservationAPIGetAppliedReservationListResponse, error) {
	result := AzureReservationAPIGetAppliedReservationListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppliedReservations); err != nil {
		return AzureReservationAPIGetAppliedReservationListResponse{}, err
	}
	return result, nil
}

// getAppliedReservationListHandleError handles the GetAppliedReservationList error response.
func (client *AzureReservationAPIClient) getAppliedReservationListHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetCatalog - Get the regions and skus that are available for RI purchase for the specified Azure subscription.
// If the operation fails it returns the *Error error type.
func (client *AzureReservationAPIClient) GetCatalog(ctx context.Context, subscriptionID string, options *AzureReservationAPIGetCatalogOptions) (AzureReservationAPIGetCatalogResponse, error) {
	req, err := client.getCatalogCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return AzureReservationAPIGetCatalogResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureReservationAPIGetCatalogResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureReservationAPIGetCatalogResponse{}, client.getCatalogHandleError(resp)
	}
	return client.getCatalogHandleResponse(resp)
}

// getCatalogCreateRequest creates the GetCatalog request.
func (client *AzureReservationAPIClient) getCatalogCreateRequest(ctx context.Context, subscriptionID string, options *AzureReservationAPIGetCatalogOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/catalogs"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.ReservedResourceType != nil {
		reqQP.Set("reservedResourceType", *options.ReservedResourceType)
	}
	if options != nil && options.Location != nil {
		reqQP.Set("location", *options.Location)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getCatalogHandleResponse handles the GetCatalog response.
func (client *AzureReservationAPIClient) getCatalogHandleResponse(resp *http.Response) (AzureReservationAPIGetCatalogResponse, error) {
	result := AzureReservationAPIGetCatalogResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatalogArray); err != nil {
		return AzureReservationAPIGetCatalogResponse{}, err
	}
	return result, nil
}

// getCatalogHandleError handles the GetCatalog error response.
func (client *AzureReservationAPIClient) getCatalogHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
