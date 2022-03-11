//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

import "time"

// AccessKeys - The Account access keys.
type AccessKeys struct {
	// Gets or sets the primary connection string.
	AtlasKafkaPrimaryEndpoint *string `json:"atlasKafkaPrimaryEndpoint,omitempty"`

	// Gets or sets the secondary connection string.
	AtlasKafkaSecondaryEndpoint *string `json:"atlasKafkaSecondaryEndpoint,omitempty"`
}

// Account resource
type Account struct {
	// Identity Info on the tracked resource
	Identity *Identity `json:"identity,omitempty"`

	// Gets or sets the location.
	Location *string `json:"location,omitempty"`

	// Gets or sets the properties.
	Properties *AccountProperties `json:"properties,omitempty"`

	// Tags on the azure resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Gets or sets the identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the Sku.
	SKU *AccountSKU `json:"sku,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *TrackedResourceSystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountEndpoints - The account endpoints
type AccountEndpoints struct {
	// READ-ONLY; Gets the catalog endpoint.
	Catalog *string `json:"catalog,omitempty" azure:"ro"`

	// READ-ONLY; Gets the guardian endpoint.
	Guardian *string `json:"guardian,omitempty" azure:"ro"`

	// READ-ONLY; Gets the scan endpoint.
	Scan *string `json:"scan,omitempty" azure:"ro"`
}

// AccountList - Paged list of account resources
type AccountList struct {
	// REQUIRED; Collection of items of type results.
	Value []*Account `json:"value,omitempty"`

	// Total item count.
	Count *int64 `json:"count,omitempty"`

	// The Url of next result page.
	NextLink *string `json:"nextLink,omitempty"`
}

// AccountProperties - The account properties
type AccountProperties struct {
	// Cloud connectors. External cloud identifier used as part of scanning configuration.
	CloudConnectors *CloudConnectors `json:"cloudConnectors,omitempty"`

	// Gets or sets the managed resource group name
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty"`

	// Gets or sets the public network access.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// READ-ONLY; Gets the time at which the entity was created.
	CreatedAt *time.Time `json:"createdAt,omitempty" azure:"ro"`

	// READ-ONLY; Gets the creator of the entity.
	CreatedBy *string `json:"createdBy,omitempty" azure:"ro"`

	// READ-ONLY; Gets the creators of the entity's object id.
	CreatedByObjectID *string `json:"createdByObjectId,omitempty" azure:"ro"`

	// READ-ONLY; The URIs that are the public endpoints of the account.
	Endpoints *AccountPropertiesEndpoints `json:"endpoints,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the friendly name.
	FriendlyName *string `json:"friendlyName,omitempty" azure:"ro"`

	// READ-ONLY; Gets the resource identifiers of the managed resources.
	ManagedResources *AccountPropertiesManagedResources `json:"managedResources,omitempty" azure:"ro"`

	// READ-ONLY; Gets the private endpoint connections information.
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the state of the provisioning.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// AccountPropertiesEndpoints - The URIs that are the public endpoints of the account.
type AccountPropertiesEndpoints struct {
	// READ-ONLY; Gets the catalog endpoint.
	Catalog *string `json:"catalog,omitempty" azure:"ro"`

	// READ-ONLY; Gets the guardian endpoint.
	Guardian *string `json:"guardian,omitempty" azure:"ro"`

	// READ-ONLY; Gets the scan endpoint.
	Scan *string `json:"scan,omitempty" azure:"ro"`
}

// AccountPropertiesManagedResources - Gets the resource identifiers of the managed resources.
type AccountPropertiesManagedResources struct {
	// READ-ONLY; Gets the managed event hub namespace resource identifier.
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" azure:"ro"`

	// READ-ONLY; Gets the managed resource group resource identifier. This resource group will host resource dependencies for
	// the account.
	ResourceGroup *string `json:"resourceGroup,omitempty" azure:"ro"`

	// READ-ONLY; Gets the managed storage account resource identifier.
	StorageAccount *string `json:"storageAccount,omitempty" azure:"ro"`
}

// AccountSKU - Gets or sets the Sku.
type AccountSKU struct {
	// Gets or sets the sku capacity.
	Capacity *int32 `json:"capacity,omitempty"`

	// Gets or sets the sku name.
	Name *Name `json:"name,omitempty"`
}

// AccountSKUAutoGenerated - The Sku
type AccountSKUAutoGenerated struct {
	// Gets or sets the sku capacity.
	Capacity *int32 `json:"capacity,omitempty"`

	// Gets or sets the sku name.
	Name *Name `json:"name,omitempty"`
}

// AccountUpdateParameters - The account update properties.
type AccountUpdateParameters struct {
	// Identity related info to add/remove userAssignedIdentities.
	Identity *Identity `json:"identity,omitempty"`

	// The account properties.
	Properties *AccountProperties `json:"properties,omitempty"`

	// Tags on the azure resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// AccountsClientAddRootCollectionAdminOptions contains the optional parameters for the AccountsClient.AddRootCollectionAdmin
// method.
type AccountsClientAddRootCollectionAdminOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccountsClient.BeginCreateOrUpdate method.
type AccountsClientBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
type AccountsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup method.
type AccountsClientListByResourceGroupOptions struct {
	// The skip token.
	SkipToken *string
}

// AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.ListBySubscription method.
type AccountsClientListBySubscriptionOptions struct {
	// The skip token.
	SkipToken *string
}

// AccountsClientListKeysOptions contains the optional parameters for the AccountsClient.ListKeys method.
type AccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// CheckNameAvailabilityRequest - The request payload for CheckNameAvailability API
type CheckNameAvailabilityRequest struct {
	// Resource name to verify for availability
	Name *string `json:"name,omitempty"`

	// Fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResult - The response payload for CheckNameAvailability API
type CheckNameAvailabilityResult struct {
	// Error message
	Message *string `json:"message,omitempty"`

	// Indicates if name is valid and available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason the name is not available.
	Reason *Reason `json:"reason,omitempty"`
}

// CloudConnectors - External Cloud Service connectors
type CloudConnectors struct {
	// READ-ONLY; AWS external identifier. Configured in AWS to allow use of the role arn used for scanning
	AwsExternalID *string `json:"awsExternalId,omitempty" azure:"ro"`
}

// CollectionAdminUpdate - Collection administrator update.
type CollectionAdminUpdate struct {
	// Gets or sets the object identifier of the admin.
	ObjectID *string `json:"objectId,omitempty"`
}

// DefaultAccountPayload - Payload to get and set the default account in the given scope
type DefaultAccountPayload struct {
	// The name of the account that is set as the default.
	AccountName *string `json:"accountName,omitempty"`

	// The resource group name of the account that is set as the default.
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	// The scope object ID. For example, sub ID or tenant ID.
	Scope *string `json:"scope,omitempty"`

	// The scope tenant in which the default account is set.
	ScopeTenantID *string `json:"scopeTenantId,omitempty"`

	// The scope where the default account is set.
	ScopeType *ScopeType `json:"scopeType,omitempty"`

	// The subscription ID of the account that is set as the default.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// DefaultAccountsClientGetOptions contains the optional parameters for the DefaultAccountsClient.Get method.
type DefaultAccountsClientGetOptions struct {
	// The Id of the scope object, for example if the scope is "Subscription" then it is the ID of that subscription.
	Scope *string
}

// DefaultAccountsClientRemoveOptions contains the optional parameters for the DefaultAccountsClient.Remove method.
type DefaultAccountsClientRemoveOptions struct {
	// The Id of the scope object, for example if the scope is "Subscription" then it is the ID of that subscription.
	Scope *string
}

// DefaultAccountsClientSetOptions contains the optional parameters for the DefaultAccountsClient.Set method.
type DefaultAccountsClientSetOptions struct {
	// placeholder for future optional parameters
}

// DimensionProperties - properties for dimension
type DimensionProperties struct {
	// localized display name of the dimension to customer
	DisplayName *string `json:"displayName,omitempty"`

	// dimension name
	Name *string `json:"name,omitempty"`

	// flag indicating whether this dimension should be included to the customer in Azure Monitor logs (aka Shoebox)
	ToBeExportedForCustomer *bool `json:"toBeExportedForCustomer,omitempty"`
}

// ErrorModel - Default error model
type ErrorModel struct {
	// READ-ONLY; Gets or sets the code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the details.
	Details []*ErrorModel `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the messages.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponseModel - Default error response model
type ErrorResponseModel struct {
	// READ-ONLY; Gets or sets the error.
	Error *ErrorResponseModelError `json:"error,omitempty" azure:"ro"`
}

// ErrorResponseModelError - Gets or sets the error.
type ErrorResponseModelError struct {
	// READ-ONLY; Gets or sets the code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the details.
	Details []*ErrorModel `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the messages.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// Identity - The Managed Identity of the resource
type Identity struct {
	// Identity Type
	Type *Type `json:"type,omitempty"`

	// User Assigned Identities
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; Service principal object Id
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; Tenant Id
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ManagedResources - The managed resources in customer subscription.
type ManagedResources struct {
	// READ-ONLY; Gets the managed event hub namespace resource identifier.
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" azure:"ro"`

	// READ-ONLY; Gets the managed resource group resource identifier. This resource group will host resource dependencies for
	// the account.
	ResourceGroup *string `json:"resourceGroup,omitempty" azure:"ro"`

	// READ-ONLY; Gets the managed storage account resource identifier.
	StorageAccount *string `json:"storageAccount,omitempty" azure:"ro"`
}

// Operation resource
type Operation struct {
	// Properties on the operation
	Display *OperationDisplay `json:"display,omitempty"`

	// Whether operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name for display purposes
	Name *string `json:"name,omitempty"`

	// origin of the operation
	Origin *string `json:"origin,omitempty"`

	// properties for the operation meta info
	Properties *OperationProperties `json:"properties,omitempty"`
}

// OperationDisplay - The response model for get operation properties
type OperationDisplay struct {
	// Description of the operation for display purposes
	Description *string `json:"description,omitempty"`

	// Name of the operation for display purposes
	Operation *string `json:"operation,omitempty"`

	// Name of the provider for display purposes
	Provider *string `json:"provider,omitempty"`

	// Name of the resource type for display purposes
	Resource *string `json:"resource,omitempty"`
}

// OperationList - Paged list of operation resources
type OperationList struct {
	// REQUIRED; Collection of items of type results.
	Value []*Operation `json:"value,omitempty"`

	// Total item count.
	Count *int64 `json:"count,omitempty"`

	// The Url of next result page.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationMetaLogSpecification - log specifications for operation api
type OperationMetaLogSpecification struct {
	// blob duration of the log
	BlobDuration *string `json:"blobDuration,omitempty"`

	// localized name of the log category
	DisplayName *string `json:"displayName,omitempty"`

	// name of the log category
	Name *string `json:"name,omitempty"`
}

// OperationMetaMetricSpecification - metric specifications for the operation
type OperationMetaMetricSpecification struct {
	// aggregation type of metric
	AggregationType *string `json:"aggregationType,omitempty"`

	// properties for dimension
	Dimensions []*DimensionProperties `json:"dimensions,omitempty"`

	// description of the metric
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// localized name of the metric
	DisplayName *string `json:"displayName,omitempty"`

	// enable regional mdm account
	EnableRegionalMdmAccount *string `json:"enableRegionalMdmAccount,omitempty"`

	// internal metric name
	InternalMetricName *string `json:"internalMetricName,omitempty"`

	// name of the metric
	Name *string `json:"name,omitempty"`

	// dimension name use to replace resource id if specified
	ResourceIDDimensionNameOverride *string `json:"resourceIdDimensionNameOverride,omitempty"`

	// Metric namespace. Only set the namespace if different from the default value, leaving it empty makes it use the value from
	// the ARM manifest.
	SourceMdmNamespace *string `json:"sourceMdmNamespace,omitempty"`

	// supported aggregation types
	SupportedAggregationTypes []*string `json:"supportedAggregationTypes,omitempty"`

	// supported time grain types
	SupportedTimeGrainTypes []*string `json:"supportedTimeGrainTypes,omitempty"`

	// units for the metric
	Unit *string `json:"unit,omitempty"`
}

// OperationMetaServiceSpecification - The operation meta service specification
type OperationMetaServiceSpecification struct {
	// log specifications for the operation
	LogSpecifications []*OperationMetaLogSpecification `json:"logSpecifications,omitempty"`

	// metric specifications for the operation
	MetricSpecifications []*OperationMetaMetricSpecification `json:"metricSpecifications,omitempty"`
}

// OperationProperties - properties on meta info
type OperationProperties struct {
	// meta service specification
	ServiceSpecification *OperationMetaServiceSpecification `json:"serviceSpecification,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpoint - A private endpoint class.
type PrivateEndpoint struct {
	// The private endpoint identifier.
	ID *string `json:"id,omitempty"`
}

// PrivateEndpointConnection - A private endpoint connection class.
type PrivateEndpointConnection struct {
	// The connection identifier.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Gets or sets the identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionList - Paged list of private endpoint connections
type PrivateEndpointConnectionList struct {
	// REQUIRED; Collection of items of type results.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`

	// Total item count.
	Count *int64 `json:"count,omitempty"`

	// The Url of next result page.
	NextLink *string `json:"nextLink,omitempty"`
}

// PrivateEndpointConnectionProperties - A private endpoint connection properties class.
type PrivateEndpointConnectionProperties struct {
	// The private endpoint information.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// The private link service connection state.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByAccountOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByAccount
// method.
type PrivateEndpointConnectionsClientListByAccountOptions struct {
	// The skip token.
	SkipToken *string
}

// PrivateLinkResource - A privately linkable resource.
type PrivateLinkResource struct {
	// READ-ONLY; The private link resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceList - Paged list of private link resources
type PrivateLinkResourceList struct {
	// REQUIRED; Collection of items of type results.
	Value []*PrivateLinkResource `json:"value,omitempty"`

	// Total item count.
	Count *int64 `json:"count,omitempty"`

	// The Url of next result page.
	NextLink *string `json:"nextLink,omitempty"`
}

// PrivateLinkResourceProperties - A privately linkable resource properties.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The private link resource group identifier.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; This translates to how many Private IPs should be created for each privately linkable resource.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`

	// READ-ONLY; The required zone names for private link resource.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty" azure:"ro"`
}

// PrivateLinkResourcesClientGetByGroupIDOptions contains the optional parameters for the PrivateLinkResourcesClient.GetByGroupID
// method.
type PrivateLinkResourcesClientGetByGroupIDOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByAccountOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByAccount
// method.
type PrivateLinkResourcesClientListByAccountOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - The private link service connection state.
type PrivateLinkServiceConnectionState struct {
	// The required actions.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The description.
	Description *string `json:"description,omitempty"`

	// The status.
	Status *Status `json:"status,omitempty"`
}

// ProxyResource - Proxy Azure Resource
type ProxyResource struct {
	// READ-ONLY; Gets or sets the identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// READ-ONLY; The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty" azure:"ro"`

	// READ-ONLY; The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty" azure:"ro"`

	// READ-ONLY; The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last modification the resource (UTC).
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty" azure:"ro"`

	// READ-ONLY; The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" azure:"ro"`

	// READ-ONLY; The type of identity that last modified the resource.
	LastModifiedByType *LastModifiedByType `json:"lastModifiedByType,omitempty" azure:"ro"`
}

// TrackedResource - Azure ARM Tracked Resource
type TrackedResource struct {
	// Identity Info on the tracked resource
	Identity *Identity `json:"identity,omitempty"`

	// Gets or sets the location.
	Location *string `json:"location,omitempty"`

	// Tags on the azure resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Gets or sets the identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *TrackedResourceSystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TrackedResourceSystemData - Metadata pertaining to creation and last modification of the resource.
type TrackedResourceSystemData struct {
	// READ-ONLY; The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty" azure:"ro"`

	// READ-ONLY; The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty" azure:"ro"`

	// READ-ONLY; The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last modification the resource (UTC).
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty" azure:"ro"`

	// READ-ONLY; The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" azure:"ro"`

	// READ-ONLY; The type of identity that last modified the resource.
	LastModifiedByType *LastModifiedByType `json:"lastModifiedByType,omitempty" azure:"ro"`
}

// UserAssignedIdentity - Uses client ID and Principal ID
type UserAssignedIdentity struct {
	// READ-ONLY; Gets or Sets Client ID
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; Gets or Sets Principal ID
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}
