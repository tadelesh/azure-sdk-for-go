//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

// AuthorizationPoliciesClientCreateOrUpdateResponse contains the response from method AuthorizationPoliciesClient.CreateOrUpdate.
type AuthorizationPoliciesClientCreateOrUpdateResponse struct {
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientGetResponse contains the response from method AuthorizationPoliciesClient.Get.
type AuthorizationPoliciesClientGetResponse struct {
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientListByHubResponse contains the response from method AuthorizationPoliciesClient.ListByHub.
type AuthorizationPoliciesClientListByHubResponse struct {
	AuthorizationPolicyListResult
}

// AuthorizationPoliciesClientRegeneratePrimaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegeneratePrimaryKey.
type AuthorizationPoliciesClientRegeneratePrimaryKeyResponse struct {
	AuthorizationPolicy
}

// AuthorizationPoliciesClientRegenerateSecondaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegenerateSecondaryKey.
type AuthorizationPoliciesClientRegenerateSecondaryKeyResponse struct {
	AuthorizationPolicy
}

// ConnectorMappingsClientCreateOrUpdateResponse contains the response from method ConnectorMappingsClient.CreateOrUpdate.
type ConnectorMappingsClientCreateOrUpdateResponse struct {
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientDeleteResponse contains the response from method ConnectorMappingsClient.Delete.
type ConnectorMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorMappingsClientGetResponse contains the response from method ConnectorMappingsClient.Get.
type ConnectorMappingsClientGetResponse struct {
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientListByConnectorResponse contains the response from method ConnectorMappingsClient.ListByConnector.
type ConnectorMappingsClientListByConnectorResponse struct {
	ConnectorMappingListResult
}

// ConnectorsClientCreateOrUpdateResponse contains the response from method ConnectorsClient.CreateOrUpdate.
type ConnectorsClientCreateOrUpdateResponse struct {
	ConnectorResourceFormat
}

// ConnectorsClientDeleteResponse contains the response from method ConnectorsClient.Delete.
type ConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorsClientGetResponse contains the response from method ConnectorsClient.Get.
type ConnectorsClientGetResponse struct {
	ConnectorResourceFormat
}

// ConnectorsClientListByHubResponse contains the response from method ConnectorsClient.ListByHub.
type ConnectorsClientListByHubResponse struct {
	ConnectorListResult
}

// HubsClientCreateOrUpdateResponse contains the response from method HubsClient.CreateOrUpdate.
type HubsClientCreateOrUpdateResponse struct {
	Hub
}

// HubsClientDeleteResponse contains the response from method HubsClient.Delete.
type HubsClientDeleteResponse struct {
	// placeholder for future response values
}

// HubsClientGetResponse contains the response from method HubsClient.Get.
type HubsClientGetResponse struct {
	Hub
}

// HubsClientListByResourceGroupResponse contains the response from method HubsClient.ListByResourceGroup.
type HubsClientListByResourceGroupResponse struct {
	HubListResult
}

// HubsClientListResponse contains the response from method HubsClient.List.
type HubsClientListResponse struct {
	HubListResult
}

// HubsClientUpdateResponse contains the response from method HubsClient.Update.
type HubsClientUpdateResponse struct {
	Hub
}

// ImagesClientGetUploadURLForDataResponse contains the response from method ImagesClient.GetUploadURLForData.
type ImagesClientGetUploadURLForDataResponse struct {
	ImageDefinition
}

// ImagesClientGetUploadURLForEntityTypeResponse contains the response from method ImagesClient.GetUploadURLForEntityType.
type ImagesClientGetUploadURLForEntityTypeResponse struct {
	ImageDefinition
}

// InteractionsClientCreateOrUpdateResponse contains the response from method InteractionsClient.CreateOrUpdate.
type InteractionsClientCreateOrUpdateResponse struct {
	InteractionResourceFormat
}

// InteractionsClientGetResponse contains the response from method InteractionsClient.Get.
type InteractionsClientGetResponse struct {
	InteractionResourceFormat
}

// InteractionsClientListByHubResponse contains the response from method InteractionsClient.ListByHub.
type InteractionsClientListByHubResponse struct {
	InteractionListResult
}

// InteractionsClientSuggestRelationshipLinksResponse contains the response from method InteractionsClient.SuggestRelationshipLinks.
type InteractionsClientSuggestRelationshipLinksResponse struct {
	SuggestRelationshipLinksResponse
}

// KpiClientCreateOrUpdateResponse contains the response from method KpiClient.CreateOrUpdate.
type KpiClientCreateOrUpdateResponse struct {
	KpiResourceFormat
}

// KpiClientDeleteResponse contains the response from method KpiClient.Delete.
type KpiClientDeleteResponse struct {
	// placeholder for future response values
}

// KpiClientGetResponse contains the response from method KpiClient.Get.
type KpiClientGetResponse struct {
	KpiResourceFormat
}

// KpiClientListByHubResponse contains the response from method KpiClient.ListByHub.
type KpiClientListByHubResponse struct {
	KpiListResult
}

// KpiClientReprocessResponse contains the response from method KpiClient.Reprocess.
type KpiClientReprocessResponse struct {
	// placeholder for future response values
}

// LinksClientCreateOrUpdateResponse contains the response from method LinksClient.CreateOrUpdate.
type LinksClientCreateOrUpdateResponse struct {
	LinkResourceFormat
}

// LinksClientDeleteResponse contains the response from method LinksClient.Delete.
type LinksClientDeleteResponse struct {
	// placeholder for future response values
}

// LinksClientGetResponse contains the response from method LinksClient.Get.
type LinksClientGetResponse struct {
	LinkResourceFormat
}

// LinksClientListByHubResponse contains the response from method LinksClient.ListByHub.
type LinksClientListByHubResponse struct {
	LinkListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PredictionsClientCreateOrUpdateResponse contains the response from method PredictionsClient.CreateOrUpdate.
type PredictionsClientCreateOrUpdateResponse struct {
	PredictionResourceFormat
}

// PredictionsClientDeleteResponse contains the response from method PredictionsClient.Delete.
type PredictionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PredictionsClientGetModelStatusResponse contains the response from method PredictionsClient.GetModelStatus.
type PredictionsClientGetModelStatusResponse struct {
	PredictionModelStatus
}

// PredictionsClientGetResponse contains the response from method PredictionsClient.Get.
type PredictionsClientGetResponse struct {
	PredictionResourceFormat
}

// PredictionsClientGetTrainingResultsResponse contains the response from method PredictionsClient.GetTrainingResults.
type PredictionsClientGetTrainingResultsResponse struct {
	PredictionTrainingResults
}

// PredictionsClientListByHubResponse contains the response from method PredictionsClient.ListByHub.
type PredictionsClientListByHubResponse struct {
	PredictionListResult
}

// PredictionsClientModelStatusResponse contains the response from method PredictionsClient.ModelStatus.
type PredictionsClientModelStatusResponse struct {
	// placeholder for future response values
}

// ProfilesClientCreateOrUpdateResponse contains the response from method ProfilesClient.CreateOrUpdate.
type ProfilesClientCreateOrUpdateResponse struct {
	ProfileResourceFormat
}

// ProfilesClientDeleteResponse contains the response from method ProfilesClient.Delete.
type ProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProfilesClientGetEnrichingKpisResponse contains the response from method ProfilesClient.GetEnrichingKpis.
type ProfilesClientGetEnrichingKpisResponse struct {
	// Array of KpiDefinition
	KpiDefinitionArray []*KpiDefinition
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	ProfileResourceFormat
}

// ProfilesClientListByHubResponse contains the response from method ProfilesClient.ListByHub.
type ProfilesClientListByHubResponse struct {
	ProfileListResult
}

// RelationshipLinksClientCreateOrUpdateResponse contains the response from method RelationshipLinksClient.CreateOrUpdate.
type RelationshipLinksClientCreateOrUpdateResponse struct {
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientDeleteResponse contains the response from method RelationshipLinksClient.Delete.
type RelationshipLinksClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipLinksClientGetResponse contains the response from method RelationshipLinksClient.Get.
type RelationshipLinksClientGetResponse struct {
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientListByHubResponse contains the response from method RelationshipLinksClient.ListByHub.
type RelationshipLinksClientListByHubResponse struct {
	RelationshipLinkListResult
}

// RelationshipsClientCreateOrUpdateResponse contains the response from method RelationshipsClient.CreateOrUpdate.
type RelationshipsClientCreateOrUpdateResponse struct {
	RelationshipResourceFormat
}

// RelationshipsClientDeleteResponse contains the response from method RelationshipsClient.Delete.
type RelationshipsClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipsClientGetResponse contains the response from method RelationshipsClient.Get.
type RelationshipsClientGetResponse struct {
	RelationshipResourceFormat
}

// RelationshipsClientListByHubResponse contains the response from method RelationshipsClient.ListByHub.
type RelationshipsClientListByHubResponse struct {
	RelationshipListResult
}

// RoleAssignmentsClientCreateOrUpdateResponse contains the response from method RoleAssignmentsClient.CreateOrUpdate.
type RoleAssignmentsClientCreateOrUpdateResponse struct {
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientDeleteResponse contains the response from method RoleAssignmentsClient.Delete.
type RoleAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// RoleAssignmentsClientGetResponse contains the response from method RoleAssignmentsClient.Get.
type RoleAssignmentsClientGetResponse struct {
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientListByHubResponse contains the response from method RoleAssignmentsClient.ListByHub.
type RoleAssignmentsClientListByHubResponse struct {
	RoleAssignmentListResult
}

// RolesClientListByHubResponse contains the response from method RolesClient.ListByHub.
type RolesClientListByHubResponse struct {
	RoleListResult
}

// ViewsClientCreateOrUpdateResponse contains the response from method ViewsClient.CreateOrUpdate.
type ViewsClientCreateOrUpdateResponse struct {
	ViewResourceFormat
}

// ViewsClientDeleteResponse contains the response from method ViewsClient.Delete.
type ViewsClientDeleteResponse struct {
	// placeholder for future response values
}

// ViewsClientGetResponse contains the response from method ViewsClient.Get.
type ViewsClientGetResponse struct {
	ViewResourceFormat
}

// ViewsClientListByHubResponse contains the response from method ViewsClient.ListByHub.
type ViewsClientListByHubResponse struct {
	ViewListResult
}

// WidgetTypesClientGetResponse contains the response from method WidgetTypesClient.Get.
type WidgetTypesClientGetResponse struct {
	WidgetTypeResourceFormat
}

// WidgetTypesClientListByHubResponse contains the response from method WidgetTypesClient.ListByHub.
type WidgetTypesClientListByHubResponse struct {
	WidgetTypeListResult
}
