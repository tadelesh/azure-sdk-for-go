//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsoftwareplan

// ClientRegisterResponse contains the response from method Client.Register.
type ClientRegisterResponse struct {
	// placeholder for future response values
}

// HybridUseBenefitClientCreateResponse contains the response from method HybridUseBenefitClient.Create.
type HybridUseBenefitClientCreateResponse struct {
	HybridUseBenefitModel
}

// HybridUseBenefitClientDeleteResponse contains the response from method HybridUseBenefitClient.Delete.
type HybridUseBenefitClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridUseBenefitClientGetResponse contains the response from method HybridUseBenefitClient.Get.
type HybridUseBenefitClientGetResponse struct {
	HybridUseBenefitModel
}

// HybridUseBenefitClientListResponse contains the response from method HybridUseBenefitClient.List.
type HybridUseBenefitClientListResponse struct {
	HybridUseBenefitListResult
}

// HybridUseBenefitClientUpdateResponse contains the response from method HybridUseBenefitClient.Update.
type HybridUseBenefitClientUpdateResponse struct {
	HybridUseBenefitModel
}

// HybridUseBenefitRevisionClientListResponse contains the response from method HybridUseBenefitRevisionClient.List.
type HybridUseBenefitRevisionClientListResponse struct {
	HybridUseBenefitListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationList
}
