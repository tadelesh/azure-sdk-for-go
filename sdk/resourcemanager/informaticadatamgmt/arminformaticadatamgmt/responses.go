//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arminformaticadatamgmt

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// OrganizationsClientCreateOrUpdateResponse contains the response from method OrganizationsClient.BeginCreateOrUpdate.
type OrganizationsClientCreateOrUpdateResponse struct {
	// An Organization Resource by Informatica.
	InformaticaOrganizationResource
}

// OrganizationsClientDeleteResponse contains the response from method OrganizationsClient.BeginDelete.
type OrganizationsClientDeleteResponse struct {
	// placeholder for future response values
}

// OrganizationsClientGetAllServerlessRuntimesResponse contains the response from method OrganizationsClient.GetAllServerlessRuntimes.
type OrganizationsClientGetAllServerlessRuntimesResponse struct {
	// A list of serverless runtime resources as fetched using the informatica APIs
	InformaticaServerlessRuntimeResourceList
}

// OrganizationsClientGetResponse contains the response from method OrganizationsClient.Get.
type OrganizationsClientGetResponse struct {
	// An Organization Resource by Informatica.
	InformaticaOrganizationResource
}

// OrganizationsClientGetServerlessMetadataResponse contains the response from method OrganizationsClient.GetServerlessMetadata.
type OrganizationsClientGetServerlessMetadataResponse struct {
	// Serverless Runtime environment Metadata response.
	ServerlessMetadataResponse
}

// OrganizationsClientListByResourceGroupResponse contains the response from method OrganizationsClient.NewListByResourceGroupPager.
type OrganizationsClientListByResourceGroupResponse struct {
	// The response of a InformaticaOrganizationResource list operation.
	InformaticaOrganizationResourceListResult
}

// OrganizationsClientListBySubscriptionResponse contains the response from method OrganizationsClient.NewListBySubscriptionPager.
type OrganizationsClientListBySubscriptionResponse struct {
	// The response of a InformaticaOrganizationResource list operation.
	InformaticaOrganizationResourceListResult
}

// OrganizationsClientUpdateResponse contains the response from method OrganizationsClient.Update.
type OrganizationsClientUpdateResponse struct {
	// An Organization Resource by Informatica.
	InformaticaOrganizationResource
}

// ServerlessRuntimesClientCheckDependenciesResponse contains the response from method ServerlessRuntimesClient.CheckDependencies.
type ServerlessRuntimesClientCheckDependenciesResponse struct {
	// Model for the check dependencies API for an informatica serverless runtime resource
	CheckDependenciesResponse
}

// ServerlessRuntimesClientCreateOrUpdateResponse contains the response from method ServerlessRuntimesClient.BeginCreateOrUpdate.
type ServerlessRuntimesClientCreateOrUpdateResponse struct {
	// A Serverless Runtime environment resource by Informatica.
	InformaticaServerlessRuntimeResource
}

// ServerlessRuntimesClientDeleteResponse contains the response from method ServerlessRuntimesClient.BeginDelete.
type ServerlessRuntimesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerlessRuntimesClientGetResponse contains the response from method ServerlessRuntimesClient.Get.
type ServerlessRuntimesClientGetResponse struct {
	// A Serverless Runtime environment resource by Informatica.
	InformaticaServerlessRuntimeResource
}

// ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse contains the response from method ServerlessRuntimesClient.NewListByInformaticaOrganizationResourcePager.
type ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse struct {
	// The response of a InformaticaServerlessRuntimeResource list operation.
	InformaticaServerlessRuntimeResourceListResult
}

// ServerlessRuntimesClientServerlessResourceByIDResponse contains the response from method ServerlessRuntimesClient.ServerlessResourceByID.
type ServerlessRuntimesClientServerlessResourceByIDResponse struct {
	// A Serverless Runtime environment resource by Informatica.
	InformaticaServerlessRuntimeResource
}

// ServerlessRuntimesClientStartFailedServerlessRuntimeResponse contains the response from method ServerlessRuntimesClient.StartFailedServerlessRuntime.
type ServerlessRuntimesClientStartFailedServerlessRuntimeResponse struct {
	// placeholder for future response values
}

// ServerlessRuntimesClientUpdateResponse contains the response from method ServerlessRuntimesClient.Update.
type ServerlessRuntimesClientUpdateResponse struct {
	// A Serverless Runtime environment resource by Informatica.
	InformaticaServerlessRuntimeResource
}
