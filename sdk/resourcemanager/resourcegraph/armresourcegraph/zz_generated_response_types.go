//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import "net/http"

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

// ResourceGraphClientResourcesHistoryResponse contains the response from method ResourceGraphClient.ResourcesHistory.
type ResourceGraphClientResourcesHistoryResponse struct {
	ResourceGraphClientResourcesHistoryResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceGraphClientResourcesHistoryResult contains the result from method ResourceGraphClient.ResourcesHistory.
type ResourceGraphClientResourcesHistoryResult struct {
	// Any object
	Object map[string]interface{}
}

// ResourceGraphClientResourcesResponse contains the response from method ResourceGraphClient.Resources.
type ResourceGraphClientResourcesResponse struct {
	ResourceGraphClientResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceGraphClientResourcesResult contains the result from method ResourceGraphClient.Resources.
type ResourceGraphClientResourcesResult struct {
	QueryResponse
}
