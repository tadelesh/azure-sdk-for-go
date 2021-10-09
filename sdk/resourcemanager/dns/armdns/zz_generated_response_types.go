//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// DNSResourceReferenceGetByTargetResourcesResponse contains the response from method DNSResourceReference.GetByTargetResources.
type DNSResourceReferenceGetByTargetResourcesResponse struct {
	DNSResourceReferenceGetByTargetResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DNSResourceReferenceGetByTargetResourcesResult contains the result from method DNSResourceReference.GetByTargetResources.
type DNSResourceReferenceGetByTargetResourcesResult struct {
	DNSResourceReferenceResult
}

// RecordSetsCreateOrUpdateResponse contains the response from method RecordSets.CreateOrUpdate.
type RecordSetsCreateOrUpdateResponse struct {
	RecordSetsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsCreateOrUpdateResult contains the result from method RecordSets.CreateOrUpdate.
type RecordSetsCreateOrUpdateResult struct {
	RecordSet
}

// RecordSetsDeleteResponse contains the response from method RecordSets.Delete.
type RecordSetsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsGetResponse contains the response from method RecordSets.Get.
type RecordSetsGetResponse struct {
	RecordSetsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsGetResult contains the result from method RecordSets.Get.
type RecordSetsGetResult struct {
	RecordSet
}

// RecordSetsListAllByDNSZoneResponse contains the response from method RecordSets.ListAllByDNSZone.
type RecordSetsListAllByDNSZoneResponse struct {
	RecordSetsListAllByDNSZoneResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsListAllByDNSZoneResult contains the result from method RecordSets.ListAllByDNSZone.
type RecordSetsListAllByDNSZoneResult struct {
	RecordSetListResult
}

// RecordSetsListByDNSZoneResponse contains the response from method RecordSets.ListByDNSZone.
type RecordSetsListByDNSZoneResponse struct {
	RecordSetsListByDNSZoneResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsListByDNSZoneResult contains the result from method RecordSets.ListByDNSZone.
type RecordSetsListByDNSZoneResult struct {
	RecordSetListResult
}

// RecordSetsListByTypeResponse contains the response from method RecordSets.ListByType.
type RecordSetsListByTypeResponse struct {
	RecordSetsListByTypeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsListByTypeResult contains the result from method RecordSets.ListByType.
type RecordSetsListByTypeResult struct {
	RecordSetListResult
}

// RecordSetsUpdateResponse contains the response from method RecordSets.Update.
type RecordSetsUpdateResponse struct {
	RecordSetsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsUpdateResult contains the result from method RecordSets.Update.
type RecordSetsUpdateResult struct {
	RecordSet
}

// ZonesCreateOrUpdateResponse contains the response from method Zones.CreateOrUpdate.
type ZonesCreateOrUpdateResponse struct {
	ZonesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesCreateOrUpdateResult contains the result from method Zones.CreateOrUpdate.
type ZonesCreateOrUpdateResult struct {
	Zone
}

// ZonesDeletePollerResponse contains the response from method Zones.Delete.
type ZonesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ZonesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ZonesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ZonesDeleteResponse, error) {
	respType := ZonesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ZonesDeletePollerResponse from the provided client and resume token.
func (l *ZonesDeletePollerResponse) Resume(ctx context.Context, client *ZonesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ZonesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ZonesDeletePoller{
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

// ZonesDeleteResponse contains the response from method Zones.Delete.
type ZonesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesGetResponse contains the response from method Zones.Get.
type ZonesGetResponse struct {
	ZonesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesGetResult contains the result from method Zones.Get.
type ZonesGetResult struct {
	Zone
}

// ZonesListByResourceGroupResponse contains the response from method Zones.ListByResourceGroup.
type ZonesListByResourceGroupResponse struct {
	ZonesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesListByResourceGroupResult contains the result from method Zones.ListByResourceGroup.
type ZonesListByResourceGroupResult struct {
	ZoneListResult
}

// ZonesListResponse contains the response from method Zones.List.
type ZonesListResponse struct {
	ZonesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesListResult contains the result from method Zones.List.
type ZonesListResult struct {
	ZoneListResult
}

// ZonesUpdateResponse contains the response from method Zones.Update.
type ZonesUpdateResponse struct {
	ZonesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ZonesUpdateResult contains the result from method Zones.Update.
type ZonesUpdateResult struct {
	Zone
}
