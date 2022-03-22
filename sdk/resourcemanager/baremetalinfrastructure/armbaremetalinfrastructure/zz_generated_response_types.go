//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

// AzureBareMetalInstancesClientGetResponse contains the response from method AzureBareMetalInstancesClient.Get.
type AzureBareMetalInstancesClientGetResponse struct {
	AzureBareMetalInstance
}

// AzureBareMetalInstancesClientListByResourceGroupResponse contains the response from method AzureBareMetalInstancesClient.ListByResourceGroup.
type AzureBareMetalInstancesClientListByResourceGroupResponse struct {
	AzureBareMetalInstancesListResult
}

// AzureBareMetalInstancesClientListBySubscriptionResponse contains the response from method AzureBareMetalInstancesClient.ListBySubscription.
type AzureBareMetalInstancesClientListBySubscriptionResponse struct {
	AzureBareMetalInstancesListResult
}

// AzureBareMetalInstancesClientUpdateResponse contains the response from method AzureBareMetalInstancesClient.Update.
type AzureBareMetalInstancesClientUpdateResponse struct {
	AzureBareMetalInstance
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationList
}
