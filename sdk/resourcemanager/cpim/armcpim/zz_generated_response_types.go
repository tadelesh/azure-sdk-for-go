//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcpim

// B2CTenantsClientCheckNameAvailabilityResponse contains the response from method B2CTenantsClient.CheckNameAvailability.
type B2CTenantsClientCheckNameAvailabilityResponse struct {
	NameAvailabilityResponse
}

// B2CTenantsClientCreateResponse contains the response from method B2CTenantsClient.Create.
type B2CTenantsClientCreateResponse struct {
	B2CTenantResource
}

// B2CTenantsClientDeleteResponse contains the response from method B2CTenantsClient.Delete.
type B2CTenantsClientDeleteResponse struct {
	// placeholder for future response values
}

// B2CTenantsClientGetResponse contains the response from method B2CTenantsClient.Get.
type B2CTenantsClientGetResponse struct {
	B2CTenantResource
}

// B2CTenantsClientListByResourceGroupResponse contains the response from method B2CTenantsClient.ListByResourceGroup.
type B2CTenantsClientListByResourceGroupResponse struct {
	B2CTenantResourceList
}

// B2CTenantsClientListBySubscriptionResponse contains the response from method B2CTenantsClient.ListBySubscription.
type B2CTenantsClientListBySubscriptionResponse struct {
	B2CTenantResourceList
}

// B2CTenantsClientUpdateResponse contains the response from method B2CTenantsClient.Update.
type B2CTenantsClientUpdateResponse struct {
	B2CTenantResource
}

// OperationsClientGetAsyncStatusResponse contains the response from method OperationsClient.GetAsyncStatus.
type OperationsClientGetAsyncStatusResponse struct {
	AsyncOperationStatus
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}
