//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

// DiagnosticSettingsCategoryClientListResponse contains the response from method DiagnosticSettingsCategoryClient.List.
type DiagnosticSettingsCategoryClientListResponse struct {
	DiagnosticSettingsCategoryResourceCollection
}

// DiagnosticSettingsClientCreateOrUpdateResponse contains the response from method DiagnosticSettingsClient.CreateOrUpdate.
type DiagnosticSettingsClientCreateOrUpdateResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientDeleteResponse contains the response from method DiagnosticSettingsClient.Delete.
type DiagnosticSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiagnosticSettingsClientGetResponse contains the response from method DiagnosticSettingsClient.Get.
type DiagnosticSettingsClientGetResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientListResponse contains the response from method DiagnosticSettingsClient.List.
type DiagnosticSettingsClientListResponse struct {
	DiagnosticSettingsResourceCollection
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsDiscoveryCollection
}
