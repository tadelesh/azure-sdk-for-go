//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import "net/http"

// ExtensionsClientCreateResponse contains the response from method ExtensionsClient.Create.
type ExtensionsClientCreateResponse struct {
	ExtensionsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsClientCreateResult contains the result from method ExtensionsClient.Create.
type ExtensionsClientCreateResult struct {
	Extension
}

// ExtensionsClientDeleteResponse contains the response from method ExtensionsClient.Delete.
type ExtensionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsClientGetResponse contains the response from method ExtensionsClient.Get.
type ExtensionsClientGetResponse struct {
	ExtensionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsClientGetResult contains the result from method ExtensionsClient.Get.
type ExtensionsClientGetResult struct {
	Extension
}

// ExtensionsClientListByFarmBeatsResponse contains the response from method ExtensionsClient.ListByFarmBeats.
type ExtensionsClientListByFarmBeatsResponse struct {
	ExtensionsClientListByFarmBeatsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsClientListByFarmBeatsResult contains the result from method ExtensionsClient.ListByFarmBeats.
type ExtensionsClientListByFarmBeatsResult struct {
	ExtensionListResponse
}

// ExtensionsClientUpdateResponse contains the response from method ExtensionsClient.Update.
type ExtensionsClientUpdateResponse struct {
	ExtensionsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsClientUpdateResult contains the result from method ExtensionsClient.Update.
type ExtensionsClientUpdateResult struct {
	Extension
}

// FarmBeatsExtensionsClientGetResponse contains the response from method FarmBeatsExtensionsClient.Get.
type FarmBeatsExtensionsClientGetResponse struct {
	FarmBeatsExtensionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsExtensionsClientGetResult contains the result from method FarmBeatsExtensionsClient.Get.
type FarmBeatsExtensionsClientGetResult struct {
	FarmBeatsExtension
}

// FarmBeatsExtensionsClientListResponse contains the response from method FarmBeatsExtensionsClient.List.
type FarmBeatsExtensionsClientListResponse struct {
	FarmBeatsExtensionsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsExtensionsClientListResult contains the result from method FarmBeatsExtensionsClient.List.
type FarmBeatsExtensionsClientListResult struct {
	FarmBeatsExtensionListResponse
}

// FarmBeatsModelsClientCreateOrUpdateResponse contains the response from method FarmBeatsModelsClient.CreateOrUpdate.
type FarmBeatsModelsClientCreateOrUpdateResponse struct {
	FarmBeatsModelsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientCreateOrUpdateResult contains the result from method FarmBeatsModelsClient.CreateOrUpdate.
type FarmBeatsModelsClientCreateOrUpdateResult struct {
	FarmBeats
}

// FarmBeatsModelsClientDeleteResponse contains the response from method FarmBeatsModelsClient.Delete.
type FarmBeatsModelsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientGetResponse contains the response from method FarmBeatsModelsClient.Get.
type FarmBeatsModelsClientGetResponse struct {
	FarmBeatsModelsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientGetResult contains the result from method FarmBeatsModelsClient.Get.
type FarmBeatsModelsClientGetResult struct {
	FarmBeats
}

// FarmBeatsModelsClientListByResourceGroupResponse contains the response from method FarmBeatsModelsClient.ListByResourceGroup.
type FarmBeatsModelsClientListByResourceGroupResponse struct {
	FarmBeatsModelsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientListByResourceGroupResult contains the result from method FarmBeatsModelsClient.ListByResourceGroup.
type FarmBeatsModelsClientListByResourceGroupResult struct {
	FarmBeatsListResponse
}

// FarmBeatsModelsClientListBySubscriptionResponse contains the response from method FarmBeatsModelsClient.ListBySubscription.
type FarmBeatsModelsClientListBySubscriptionResponse struct {
	FarmBeatsModelsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientListBySubscriptionResult contains the result from method FarmBeatsModelsClient.ListBySubscription.
type FarmBeatsModelsClientListBySubscriptionResult struct {
	FarmBeatsListResponse
}

// FarmBeatsModelsClientUpdateResponse contains the response from method FarmBeatsModelsClient.Update.
type FarmBeatsModelsClientUpdateResponse struct {
	FarmBeatsModelsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FarmBeatsModelsClientUpdateResult contains the result from method FarmBeatsModelsClient.Update.
type FarmBeatsModelsClientUpdateResult struct {
	FarmBeats
}

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	LocationsClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationsClientCheckNameAvailabilityResult contains the result from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

