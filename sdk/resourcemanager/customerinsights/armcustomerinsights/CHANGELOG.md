# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*InteractionsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(InteractionsClientListByHubResponse, error)`
- Function `*RelationshipLinksClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(RelationshipLinksClientListByHubResponse, error)`
- Function `*LinksClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(LinksClientListByHubResponse, error)`
- Function `*PredictionsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(PredictionsClientListByHubResponse, error)`
- Function `*AuthorizationPoliciesClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(AuthorizationPoliciesClientListByHubResponse, error)`
- Function `*KpiClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(KpiClientListByHubResponse, error)`
- Function `*RoleAssignmentsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(RoleAssignmentsClientListByHubResponse, error)`
- Function `*HubsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(HubsClientListResponse, error)`
- Function `*ConnectorMappingsClientListByConnectorPager.NextPage` return value(s) have been changed from `(bool)` to `(ConnectorMappingsClientListByConnectorResponse, error)`
- Function `*RolesClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(RolesClientListByHubResponse, error)`
- Function `*WidgetTypesClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(WidgetTypesClientListByHubResponse, error)`
- Function `*RelationshipsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(RelationshipsClientListByHubResponse, error)`
- Function `*ViewsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(ViewsClientListByHubResponse, error)`
- Function `*HubsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(HubsClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*ProfilesClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(ProfilesClientListByHubResponse, error)`
- Function `*ConnectorsClientListByHubPager.NextPage` return value(s) have been changed from `(bool)` to `(ConnectorsClientListByHubResponse, error)`
- Function `*ViewsClientListByHubPager.Err` has been removed
- Function `*ConnectorMappingsClientListByConnectorPager.Err` has been removed
- Function `*PredictionsClientListByHubPager.PageResponse` has been removed
- Function `*HubsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*ProfilesClientListByHubPager.PageResponse` has been removed
- Function `*RoleAssignmentsClientListByHubPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*ConnectorMappingsClientListByConnectorPager.PageResponse` has been removed
- Function `*AuthorizationPoliciesClientListByHubPager.PageResponse` has been removed
- Function `*RelationshipsClientListByHubPager.Err` has been removed
- Function `*LinksClientListByHubPager.PageResponse` has been removed
- Function `*RelationshipLinksClientListByHubPager.Err` has been removed
- Function `*ProfilesClientListByHubPager.Err` has been removed
- Function `*InteractionsClientListByHubPager.Err` has been removed
- Function `*RolesClientListByHubPager.Err` has been removed
- Function `*LinksClientListByHubPager.Err` has been removed
- Function `*WidgetTypesClientListByHubPager.Err` has been removed
- Function `*WidgetTypesClientListByHubPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*RelationshipLinksClientListByHubPager.PageResponse` has been removed
- Function `*ConnectorsClientListByHubPager.PageResponse` has been removed
- Function `*ViewsClientListByHubPager.PageResponse` has been removed
- Function `*RolesClientListByHubPager.PageResponse` has been removed
- Function `*AuthorizationPoliciesClientListByHubPager.Err` has been removed
- Function `*InteractionsClientListByHubPager.PageResponse` has been removed
- Function `*HubsClientListByResourceGroupPager.Err` has been removed
- Function `*RelationshipsClientListByHubPager.PageResponse` has been removed
- Function `*HubsClientListPager.Err` has been removed
- Function `*KpiClientListByHubPager.PageResponse` has been removed
- Function `*ConnectorsClientListByHubPager.Err` has been removed
- Function `*PredictionsClientListByHubPager.Err` has been removed
- Function `*KpiClientListByHubPager.Err` has been removed
- Function `*RoleAssignmentsClientListByHubPager.Err` has been removed
- Function `*HubsClientListPager.PageResponse` has been removed
- Struct `AuthorizationPoliciesClientCreateOrUpdateResult` has been removed
- Struct `AuthorizationPoliciesClientGetResult` has been removed
- Struct `AuthorizationPoliciesClientListByHubResult` has been removed
- Struct `AuthorizationPoliciesClientRegeneratePrimaryKeyResult` has been removed
- Struct `AuthorizationPoliciesClientRegenerateSecondaryKeyResult` has been removed
- Struct `ConnectorMappingsClientCreateOrUpdateResult` has been removed
- Struct `ConnectorMappingsClientGetResult` has been removed
- Struct `ConnectorMappingsClientListByConnectorResult` has been removed
- Struct `ConnectorsClientCreateOrUpdateResult` has been removed
- Struct `ConnectorsClientGetResult` has been removed
- Struct `ConnectorsClientListByHubResult` has been removed
- Struct `HubsClientCreateOrUpdateResult` has been removed
- Struct `HubsClientGetResult` has been removed
- Struct `HubsClientListByResourceGroupResult` has been removed
- Struct `HubsClientListResult` has been removed
- Struct `HubsClientUpdateResult` has been removed
- Struct `ImagesClientGetUploadURLForDataResult` has been removed
- Struct `ImagesClientGetUploadURLForEntityTypeResult` has been removed
- Struct `InteractionsClientCreateOrUpdateResult` has been removed
- Struct `InteractionsClientGetResult` has been removed
- Struct `InteractionsClientListByHubResult` has been removed
- Struct `InteractionsClientSuggestRelationshipLinksResult` has been removed
- Struct `KpiClientCreateOrUpdateResult` has been removed
- Struct `KpiClientGetResult` has been removed
- Struct `KpiClientListByHubResult` has been removed
- Struct `LinksClientCreateOrUpdateResult` has been removed
- Struct `LinksClientGetResult` has been removed
- Struct `LinksClientListByHubResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `PredictionsClientCreateOrUpdateResult` has been removed
- Struct `PredictionsClientGetModelStatusResult` has been removed
- Struct `PredictionsClientGetResult` has been removed
- Struct `PredictionsClientGetTrainingResultsResult` has been removed
- Struct `PredictionsClientListByHubResult` has been removed
- Struct `ProfilesClientCreateOrUpdateResult` has been removed
- Struct `ProfilesClientGetEnrichingKpisResult` has been removed
- Struct `ProfilesClientGetResult` has been removed
- Struct `ProfilesClientListByHubResult` has been removed
- Struct `RelationshipLinksClientCreateOrUpdateResult` has been removed
- Struct `RelationshipLinksClientGetResult` has been removed
- Struct `RelationshipLinksClientListByHubResult` has been removed
- Struct `RelationshipsClientCreateOrUpdateResult` has been removed
- Struct `RelationshipsClientGetResult` has been removed
- Struct `RelationshipsClientListByHubResult` has been removed
- Struct `RoleAssignmentsClientCreateOrUpdateResult` has been removed
- Struct `RoleAssignmentsClientGetResult` has been removed
- Struct `RoleAssignmentsClientListByHubResult` has been removed
- Struct `RolesClientListByHubResult` has been removed
- Struct `ViewsClientCreateOrUpdateResult` has been removed
- Struct `ViewsClientGetResult` has been removed
- Struct `ViewsClientListByHubResult` has been removed
- Struct `WidgetTypesClientGetResult` has been removed
- Struct `WidgetTypesClientListByHubResult` has been removed
- Field `ProfilesClientListByHubResult` of struct `ProfilesClientListByHubResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientListByHubResponse` has been removed
- Field `AuthorizationPoliciesClientRegeneratePrimaryKeyResult` of struct `AuthorizationPoliciesClientRegeneratePrimaryKeyResponse` has been removed
- Field `RawResponse` of struct `AuthorizationPoliciesClientRegeneratePrimaryKeyResponse` has been removed
- Field `RawResponse` of struct `ViewsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `RoleAssignmentsClientDeleteResponse` has been removed
- Field `KpiClientListByHubResult` of struct `KpiClientListByHubResponse` has been removed
- Field `RawResponse` of struct `KpiClientListByHubResponse` has been removed
- Field `ConnectorMappingsClientListByConnectorResult` of struct `ConnectorMappingsClientListByConnectorResponse` has been removed
- Field `RawResponse` of struct `ConnectorMappingsClientListByConnectorResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientCreateOrUpdatePollerResponse` has been removed
- Field `ConnectorMappingsClientCreateOrUpdateResult` of struct `ConnectorMappingsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ConnectorMappingsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientDeletePollerResponse` has been removed
- Field `ConnectorMappingsClientGetResult` of struct `ConnectorMappingsClientGetResponse` has been removed
- Field `RawResponse` of struct `ConnectorMappingsClientGetResponse` has been removed
- Field `RawResponse` of struct `HubsClientDeletePollerResponse` has been removed
- Field `RoleAssignmentsClientCreateOrUpdateResult` of struct `RoleAssignmentsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `RoleAssignmentsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `LinksClientCreateOrUpdatePollerResponse` has been removed
- Field `AuthorizationPoliciesClientListByHubResult` of struct `AuthorizationPoliciesClientListByHubResponse` has been removed
- Field `RawResponse` of struct `AuthorizationPoliciesClientListByHubResponse` has been removed
- Field `ConnectorsClientListByHubResult` of struct `ConnectorsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientListByHubResponse` has been removed
- Field `RelationshipsClientListByHubResult` of struct `RelationshipsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientDeletePollerResponse` has been removed
- Field `KpiClientGetResult` of struct `KpiClientGetResponse` has been removed
- Field `RawResponse` of struct `KpiClientGetResponse` has been removed
- Field `RawResponse` of struct `KpiClientDeletePollerResponse` has been removed
- Field `ProfilesClientCreateOrUpdateResult` of struct `ProfilesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientCreateOrUpdateResponse` has been removed
- Field `PredictionsClientGetResult` of struct `PredictionsClientGetResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientGetResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `ConnectorMappingsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `KpiClientReprocessResponse` has been removed
- Field `RelationshipsClientGetResult` of struct `RelationshipsClientGetResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientGetResponse` has been removed
- Field `RawResponse` of struct `InteractionsClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientModelStatusResponse` has been removed
- Field `HubsClientCreateOrUpdateResult` of struct `HubsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `HubsClientCreateOrUpdateResponse` has been removed
- Field `LinksClientListByHubResult` of struct `LinksClientListByHubResponse` has been removed
- Field `RawResponse` of struct `LinksClientListByHubResponse` has been removed
- Field `ImagesClientGetUploadURLForDataResult` of struct `ImagesClientGetUploadURLForDataResponse` has been removed
- Field `RawResponse` of struct `ImagesClientGetUploadURLForDataResponse` has been removed
- Field `HubsClientGetResult` of struct `HubsClientGetResponse` has been removed
- Field `RawResponse` of struct `HubsClientGetResponse` has been removed
- Field `HubsClientListByResourceGroupResult` of struct `HubsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `HubsClientListByResourceGroupResponse` has been removed
- Field `ProfilesClientGetEnrichingKpisResult` of struct `ProfilesClientGetEnrichingKpisResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientGetEnrichingKpisResponse` has been removed
- Field `LinksClientGetResult` of struct `LinksClientGetResponse` has been removed
- Field `RawResponse` of struct `LinksClientGetResponse` has been removed
- Field `InteractionsClientListByHubResult` of struct `InteractionsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `InteractionsClientListByHubResponse` has been removed
- Field `ConnectorsClientCreateOrUpdateResult` of struct `ConnectorsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientCreateOrUpdatePollerResponse` has been removed
- Field `AuthorizationPoliciesClientGetResult` of struct `AuthorizationPoliciesClientGetResponse` has been removed
- Field `RawResponse` of struct `AuthorizationPoliciesClientGetResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientDeleteResponse` has been removed
- Field `RelationshipsClientCreateOrUpdateResult` of struct `RelationshipsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientCreateOrUpdateResponse` has been removed
- Field `WidgetTypesClientGetResult` of struct `WidgetTypesClientGetResponse` has been removed
- Field `RawResponse` of struct `WidgetTypesClientGetResponse` has been removed
- Field `WidgetTypesClientListByHubResult` of struct `WidgetTypesClientListByHubResponse` has been removed
- Field `RawResponse` of struct `WidgetTypesClientListByHubResponse` has been removed
- Field `ViewsClientListByHubResult` of struct `ViewsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `ViewsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientDeleteResponse` has been removed
- Field `KpiClientCreateOrUpdateResult` of struct `KpiClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `KpiClientCreateOrUpdateResponse` has been removed
- Field `ConnectorsClientGetResult` of struct `ConnectorsClientGetResponse` has been removed
- Field `RawResponse` of struct `ConnectorsClientGetResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientDeleteResponse` has been removed
- Field `LinksClientCreateOrUpdateResult` of struct `LinksClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `LinksClientCreateOrUpdateResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `InteractionsClientGetResult` of struct `InteractionsClientGetResponse` has been removed
- Field `RawResponse` of struct `InteractionsClientGetResponse` has been removed
- Field `InteractionsClientSuggestRelationshipLinksResult` of struct `InteractionsClientSuggestRelationshipLinksResponse` has been removed
- Field `RawResponse` of struct `InteractionsClientSuggestRelationshipLinksResponse` has been removed
- Field `RawResponse` of struct `HubsClientDeleteResponse` has been removed
- Field `PredictionsClientListByHubResult` of struct `PredictionsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientListByHubResponse` has been removed
- Field `ViewsClientGetResult` of struct `ViewsClientGetResponse` has been removed
- Field `RawResponse` of struct `ViewsClientGetResponse` has been removed
- Field `InteractionsClientCreateOrUpdateResult` of struct `InteractionsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `InteractionsClientCreateOrUpdateResponse` has been removed
- Field `AuthorizationPoliciesClientCreateOrUpdateResult` of struct `AuthorizationPoliciesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AuthorizationPoliciesClientCreateOrUpdateResponse` has been removed
- Field `ProfilesClientGetResult` of struct `ProfilesClientGetResponse` has been removed
- Field `RawResponse` of struct `ProfilesClientGetResponse` has been removed
- Field `ViewsClientCreateOrUpdateResult` of struct `ViewsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ViewsClientCreateOrUpdateResponse` has been removed
- Field `AuthorizationPoliciesClientRegenerateSecondaryKeyResult` of struct `AuthorizationPoliciesClientRegenerateSecondaryKeyResponse` has been removed
- Field `RawResponse` of struct `AuthorizationPoliciesClientRegenerateSecondaryKeyResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `KpiClientDeleteResponse` has been removed
- Field `RolesClientListByHubResult` of struct `RolesClientListByHubResponse` has been removed
- Field `RawResponse` of struct `RolesClientListByHubResponse` has been removed
- Field `PredictionsClientGetTrainingResultsResult` of struct `PredictionsClientGetTrainingResultsResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientGetTrainingResultsResponse` has been removed
- Field `HubsClientListResult` of struct `HubsClientListResponse` has been removed
- Field `RawResponse` of struct `HubsClientListResponse` has been removed
- Field `RelationshipLinksClientGetResult` of struct `RelationshipLinksClientGetResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientGetResponse` has been removed
- Field `RawResponse` of struct `KpiClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `RoleAssignmentsClientCreateOrUpdatePollerResponse` has been removed
- Field `RelationshipLinksClientListByHubResult` of struct `RelationshipLinksClientListByHubResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientListByHubResponse` has been removed
- Field `HubsClientUpdateResult` of struct `HubsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `HubsClientUpdateResponse` has been removed
- Field `ImagesClientGetUploadURLForEntityTypeResult` of struct `ImagesClientGetUploadURLForEntityTypeResponse` has been removed
- Field `RawResponse` of struct `ImagesClientGetUploadURLForEntityTypeResponse` has been removed
- Field `RoleAssignmentsClientGetResult` of struct `RoleAssignmentsClientGetResponse` has been removed
- Field `RawResponse` of struct `RoleAssignmentsClientGetResponse` has been removed
- Field `PredictionsClientCreateOrUpdateResult` of struct `PredictionsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `RelationshipsClientCreateOrUpdatePollerResponse` has been removed
- Field `RoleAssignmentsClientListByHubResult` of struct `RoleAssignmentsClientListByHubResponse` has been removed
- Field `RawResponse` of struct `RoleAssignmentsClientListByHubResponse` has been removed
- Field `RelationshipLinksClientCreateOrUpdateResult` of struct `RelationshipLinksClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `RelationshipLinksClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `LinksClientDeleteResponse` has been removed
- Field `PredictionsClientGetModelStatusResult` of struct `PredictionsClientGetModelStatusResponse` has been removed
- Field `RawResponse` of struct `PredictionsClientGetModelStatusResponse` has been removed

### Features Added

- New function `*OperationsClientListPager.More() bool`
- New function `*InteractionsClientListByHubPager.More() bool`
- New function `*ConnectorMappingsClientListByConnectorPager.More() bool`
- New function `*AuthorizationPoliciesClientListByHubPager.More() bool`
- New function `SalesforceConnectorProperties.MarshalJSON() ([]byte, error)`
- New function `*PredictionsClientListByHubPager.More() bool`
- New function `*RelationshipLinksClientListByHubPager.More() bool`
- New function `*KpiClientListByHubPager.More() bool`
- New function `*RelationshipsClientListByHubPager.More() bool`
- New function `*HubsClientListPager.More() bool`
- New function `*LinksClientListByHubPager.More() bool`
- New function `*RolesClientListByHubPager.More() bool`
- New function `*HubsClientListByResourceGroupPager.More() bool`
- New function `*ConnectorsClientListByHubPager.More() bool`
- New function `CrmConnectorProperties.MarshalJSON() ([]byte, error)`
- New function `*RoleAssignmentsClientListByHubPager.More() bool`
- New function `*ViewsClientListByHubPager.More() bool`
- New function `*WidgetTypesClientListByHubPager.More() bool`
- New function `*ProfilesClientListByHubPager.More() bool`
- New struct `AzureBlobConnectorProperties`
- New struct `CrmConnectorEntities`
- New struct `CrmConnectorProperties`
- New struct `SalesforceConnectorProperties`
- New struct `SalesforceDiscoverSetting`
- New struct `SalesforceTable`
- New anonymous field `KpiListResult` in struct `KpiClientListByHubResponse`
- New anonymous field `ImageDefinition` in struct `ImagesClientGetUploadURLForDataResponse`
- New anonymous field `PredictionModelStatus` in struct `PredictionsClientGetModelStatusResponse`
- New anonymous field `SuggestRelationshipLinksResponse` in struct `InteractionsClientSuggestRelationshipLinksResponse`
- New anonymous field `ConnectorMappingListResult` in struct `ConnectorMappingsClientListByConnectorResponse`
- New anonymous field `Hub` in struct `HubsClientGetResponse`
- New anonymous field `AuthorizationPolicy` in struct `AuthorizationPoliciesClientRegeneratePrimaryKeyResponse`
- New anonymous field `InteractionResourceFormat` in struct `InteractionsClientCreateOrUpdateResponse`
- New anonymous field `ViewResourceFormat` in struct `ViewsClientGetResponse`
- New anonymous field `InteractionListResult` in struct `InteractionsClientListByHubResponse`
- New anonymous field `ConnectorResourceFormat` in struct `ConnectorsClientGetResponse`
- New anonymous field `RelationshipLinkResourceFormat` in struct `RelationshipLinksClientGetResponse`
- New anonymous field `WidgetTypeResourceFormat` in struct `WidgetTypesClientGetResponse`
- New anonymous field `Hub` in struct `HubsClientCreateOrUpdateResponse`
- New anonymous field `RelationshipResourceFormat` in struct `RelationshipsClientGetResponse`
- New anonymous field `ProfileResourceFormat` in struct `ProfilesClientGetResponse`
- New anonymous field `LinkResourceFormat` in struct `LinksClientGetResponse`
- New anonymous field `RoleAssignmentListResult` in struct `RoleAssignmentsClientListByHubResponse`
- New anonymous field `KpiResourceFormat` in struct `KpiClientCreateOrUpdateResponse`
- New anonymous field `RoleListResult` in struct `RolesClientListByHubResponse`
- New anonymous field `ConnectorResourceFormat` in struct `ConnectorsClientCreateOrUpdateResponse`
- New anonymous field `RelationshipLinkListResult` in struct `RelationshipLinksClientListByHubResponse`
- New anonymous field `ConnectorListResult` in struct `ConnectorsClientListByHubResponse`
- New anonymous field `ViewResourceFormat` in struct `ViewsClientCreateOrUpdateResponse`
- New anonymous field `RelationshipListResult` in struct `RelationshipsClientListByHubResponse`
- New anonymous field `ConnectorMappingResourceFormat` in struct `ConnectorMappingsClientCreateOrUpdateResponse`
- New anonymous field `ProfileResourceFormat` in struct `ProfilesClientCreateOrUpdateResponse`
- New anonymous field `PredictionResourceFormat` in struct `PredictionsClientGetResponse`
- New anonymous field `AuthorizationPolicyResourceFormat` in struct `AuthorizationPoliciesClientCreateOrUpdateResponse`
- New anonymous field `WidgetTypeListResult` in struct `WidgetTypesClientListByHubResponse`
- New anonymous field `AuthorizationPolicyListResult` in struct `AuthorizationPoliciesClientListByHubResponse`
- New anonymous field `ImageDefinition` in struct `ImagesClientGetUploadURLForEntityTypeResponse`
- New anonymous field `RoleAssignmentResourceFormat` in struct `RoleAssignmentsClientCreateOrUpdateResponse`
- New field `KpiDefinitionArray` in struct `ProfilesClientGetEnrichingKpisResponse`
- New anonymous field `RelationshipLinkResourceFormat` in struct `RelationshipLinksClientCreateOrUpdateResponse`
- New anonymous field `AuthorizationPolicy` in struct `AuthorizationPoliciesClientRegenerateSecondaryKeyResponse`
- New anonymous field `LinkListResult` in struct `LinksClientListByHubResponse`
- New anonymous field `LinkResourceFormat` in struct `LinksClientCreateOrUpdateResponse`
- New anonymous field `KpiResourceFormat` in struct `KpiClientGetResponse`
- New anonymous field `HubListResult` in struct `HubsClientListByResourceGroupResponse`
- New anonymous field `PredictionListResult` in struct `PredictionsClientListByHubResponse`
- New anonymous field `PredictionTrainingResults` in struct `PredictionsClientGetTrainingResultsResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `AuthorizationPolicyResourceFormat` in struct `AuthorizationPoliciesClientGetResponse`
- New anonymous field `InteractionResourceFormat` in struct `InteractionsClientGetResponse`
- New anonymous field `RoleAssignmentResourceFormat` in struct `RoleAssignmentsClientGetResponse`
- New anonymous field `Hub` in struct `HubsClientUpdateResponse`
- New anonymous field `ProfileListResult` in struct `ProfilesClientListByHubResponse`
- New anonymous field `RelationshipResourceFormat` in struct `RelationshipsClientCreateOrUpdateResponse`
- New anonymous field `PredictionResourceFormat` in struct `PredictionsClientCreateOrUpdateResponse`
- New anonymous field `HubListResult` in struct `HubsClientListResponse`
- New anonymous field `ViewListResult` in struct `ViewsClientListByHubResponse`
- New anonymous field `ConnectorMappingResourceFormat` in struct `ConnectorMappingsClientGetResponse`


## 0.1.0 (2022-03-10)

- Init release.