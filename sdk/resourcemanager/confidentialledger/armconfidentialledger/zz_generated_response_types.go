//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResponse
}

// LedgerClientCreateResponse contains the response from method LedgerClient.Create.
type LedgerClientCreateResponse struct {
	ConfidentialLedger
}

// LedgerClientDeleteResponse contains the response from method LedgerClient.Delete.
type LedgerClientDeleteResponse struct {
	// placeholder for future response values
}

// LedgerClientGetResponse contains the response from method LedgerClient.Get.
type LedgerClientGetResponse struct {
	ConfidentialLedger
}

// LedgerClientListByResourceGroupResponse contains the response from method LedgerClient.ListByResourceGroup.
type LedgerClientListByResourceGroupResponse struct {
	List
}

// LedgerClientListBySubscriptionResponse contains the response from method LedgerClient.ListBySubscription.
type LedgerClientListBySubscriptionResponse struct {
	List
}

// LedgerClientUpdateResponse contains the response from method LedgerClient.Update.
type LedgerClientUpdateResponse struct {
	ConfidentialLedger
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ResourceProviderOperationList
}
