//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PowerBIResourcesCreateResponse contains the response from method PowerBIResources.Create.
type PowerBIResourcesCreateResponse struct {
	PowerBIResourcesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PowerBIResourcesCreateResult contains the result from method PowerBIResources.Create.
type PowerBIResourcesCreateResult struct {
	TenantResource
}

// PowerBIResourcesDeleteResponse contains the response from method PowerBIResources.Delete.
type PowerBIResourcesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PowerBIResourcesListByResourceNameResponse contains the response from method PowerBIResources.ListByResourceName.
type PowerBIResourcesListByResourceNameResponse struct {
	PowerBIResourcesListByResourceNameResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PowerBIResourcesListByResourceNameResult contains the result from method PowerBIResources.ListByResourceName.
type PowerBIResourcesListByResourceNameResult struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}

// PowerBIResourcesUpdateResponse contains the response from method PowerBIResources.Update.
type PowerBIResourcesUpdateResponse struct {
	PowerBIResourcesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PowerBIResourcesUpdateResult contains the result from method PowerBIResources.Update.
type PowerBIResourcesUpdateResult struct {
	TenantResource
}

// PrivateEndpointConnectionsCreateResponse contains the response from method PrivateEndpointConnections.Create.
type PrivateEndpointConnectionsCreateResponse struct {
	PrivateEndpointConnectionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateResult contains the result from method PrivateEndpointConnections.Create.
type PrivateEndpointConnectionsCreateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListByResourceResponse contains the response from method PrivateEndpointConnections.ListByResource.
type PrivateEndpointConnectionsListByResourceResponse struct {
	PrivateEndpointConnectionsListByResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListByResourceResult contains the result from method PrivateEndpointConnections.ListByResource.
type PrivateEndpointConnectionsListByResourceResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	PrivateLinkResource
}

// PrivateLinkResourcesListByResourceResponse contains the response from method PrivateLinkResources.ListByResource.
type PrivateLinkResourcesListByResourceResponse struct {
	PrivateLinkResourcesListByResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListByResourceResult contains the result from method PrivateLinkResources.ListByResource.
type PrivateLinkResourcesListByResourceResult struct {
	PrivateLinkResourcesListResult
}

// PrivateLinkServiceResourceOperationResultsGetPollerResponse contains the response from method PrivateLinkServiceResourceOperationResults.Get.
type PrivateLinkServiceResourceOperationResultsGetPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateLinkServiceResourceOperationResultsGetPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateLinkServiceResourceOperationResultsGetPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateLinkServiceResourceOperationResultsGetResponse, error) {
	respType := PrivateLinkServiceResourceOperationResultsGetResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.AsyncOperationDetail)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateLinkServiceResourceOperationResultsGetPollerResponse from the provided client and resume token.
func (l *PrivateLinkServiceResourceOperationResultsGetPollerResponse) Resume(ctx context.Context, client *PrivateLinkServiceResourceOperationResultsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateLinkServiceResourceOperationResultsClient.Get", token, client.pl, client.getHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateLinkServiceResourceOperationResultsGetPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateLinkServiceResourceOperationResultsGetResponse contains the response from method PrivateLinkServiceResourceOperationResults.Get.
type PrivateLinkServiceResourceOperationResultsGetResponse struct {
	PrivateLinkServiceResourceOperationResultsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkServiceResourceOperationResultsGetResult contains the result from method PrivateLinkServiceResourceOperationResults.Get.
type PrivateLinkServiceResourceOperationResultsGetResult struct {
	AsyncOperationDetail
}

// PrivateLinkServicesForPowerBIListBySubscriptionIDResponse contains the response from method PrivateLinkServicesForPowerBI.ListBySubscriptionID.
type PrivateLinkServicesForPowerBIListBySubscriptionIDResponse struct {
	PrivateLinkServicesForPowerBIListBySubscriptionIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkServicesForPowerBIListBySubscriptionIDResult contains the result from method PrivateLinkServicesForPowerBI.ListBySubscriptionID.
type PrivateLinkServicesForPowerBIListBySubscriptionIDResult struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}

// PrivateLinkServicesListByResourceGroupResponse contains the response from method PrivateLinkServices.ListByResourceGroup.
type PrivateLinkServicesListByResourceGroupResponse struct {
	PrivateLinkServicesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkServicesListByResourceGroupResult contains the result from method PrivateLinkServices.ListByResourceGroup.
type PrivateLinkServicesListByResourceGroupResult struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}
