//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementpartner

// OperationClientListResponse contains the response from method OperationClient.List.
type OperationClientListResponse struct {
	OperationList
}

// PartnerClientCreateResponse contains the response from method PartnerClient.Create.
type PartnerClientCreateResponse struct {
	PartnerResponse
}

// PartnerClientDeleteResponse contains the response from method PartnerClient.Delete.
type PartnerClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerClientGetResponse contains the response from method PartnerClient.Get.
type PartnerClientGetResponse struct {
	PartnerResponse
}

// PartnerClientUpdateResponse contains the response from method PartnerClient.Update.
type PartnerClientUpdateResponse struct {
	PartnerResponse
}

// PartnersClientGetResponse contains the response from method PartnersClient.Get.
type PartnersClientGetResponse struct {
	PartnerResponse
}
