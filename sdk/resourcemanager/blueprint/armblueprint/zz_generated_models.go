//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

import "time"

// ArtifactClassification provides polymorphic access to related types.
// Call the interface's GetArtifact() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Artifact, *PolicyAssignmentArtifact, *RoleAssignmentArtifact, *TemplateArtifact
type ArtifactClassification interface {
	// GetArtifact returns the Artifact content of the underlying type.
	GetArtifact() *Artifact
}

// Artifact - Represents a blueprint artifact.
type Artifact struct {
	// REQUIRED; Specifies the kind of blueprint artifact.
	Kind *ArtifactKind `json:"kind,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ArtifactList - List of blueprint artifacts.
type ArtifactList struct {
	// List of blueprint artifacts.
	Value []ArtifactClassification `json:"value,omitempty"`

	// READ-ONLY; Link to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ArtifactPropertiesBase - Common properties shared by different artifacts.
type ArtifactPropertiesBase struct {
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []*string `json:"dependsOn,omitempty"`
}

// ArtifactsClientCreateOrUpdateOptions contains the optional parameters for the ArtifactsClient.CreateOrUpdate method.
type ArtifactsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ArtifactsClientDeleteOptions contains the optional parameters for the ArtifactsClient.Delete method.
type ArtifactsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ArtifactsClientGetOptions contains the optional parameters for the ArtifactsClient.Get method.
type ArtifactsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ArtifactsClientListOptions contains the optional parameters for the ArtifactsClient.List method.
type ArtifactsClientListOptions struct {
	// placeholder for future optional parameters
}

// Assignment - Represents a blueprint assignment.
type Assignment struct {
	// REQUIRED; Managed identity for this blueprint assignment.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// REQUIRED; The location of this blueprint assignment.
	Location *string `json:"location,omitempty"`

	// REQUIRED; Properties for blueprint assignment object.
	Properties *AssignmentProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AssignmentDeploymentJob - Represents individual job in given blueprint assignment operation.
type AssignmentDeploymentJob struct {
	// Name of the action performed in this job.
	Action *string `json:"action,omitempty"`

	// Result of this deployment job for each retry.
	History []*AssignmentDeploymentJobResult `json:"history,omitempty"`

	// Id of this job.
	JobID *string `json:"jobId,omitempty"`

	// State of this job.
	JobState *string `json:"jobState,omitempty"`

	// Kind of job.
	Kind *string `json:"kind,omitempty"`

	// Reference to deployment job resource id.
	RequestURI *string `json:"requestUri,omitempty"`

	// Deployment job result.
	Result *AssignmentDeploymentJobResult `json:"result,omitempty"`
}

// AssignmentDeploymentJobResult - Result of each individual deployment in a blueprint assignment.
type AssignmentDeploymentJobResult struct {
	// Contains error details if deployment job failed.
	Error *AzureResourceManagerError `json:"error,omitempty"`

	// Resources created as result of the deployment job.
	Resources []*AssignmentJobCreatedResource `json:"resources,omitempty"`
}

// AssignmentJobCreatedResource - Azure resource created from deployment job.
type AssignmentJobCreatedResource struct {
	// Additional properties in a dictionary.
	Properties map[string]*string `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AssignmentList - List of blueprint assignments
type AssignmentList struct {
	// List of blueprint assignments.
	Value []*Assignment `json:"value,omitempty"`

	// READ-ONLY; Link to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AssignmentLockSettings - Defines how resources deployed by a blueprint assignment are locked.
type AssignmentLockSettings struct {
	// List of management operations that are excluded from blueprint locks. Up to 200 actions are permitted. If the lock mode
	// is set to 'AllResourcesReadOnly', then the following actions are automatically
	// appended to 'excludedActions': '*/read', 'Microsoft.Network/virtualNetworks/subnets/join/action' and 'Microsoft.Authorization/locks/delete'.
	// If the lock mode is set to 'AllResourcesDoNotDelete', then
	// the following actions are automatically appended to 'excludedActions': 'Microsoft.Authorization/locks/delete'. Duplicate
	// actions will get removed.
	ExcludedActions []*string `json:"excludedActions,omitempty"`

	// List of AAD principals excluded from blueprint locks. Up to 5 principals are permitted.
	ExcludedPrincipals []*string `json:"excludedPrincipals,omitempty"`

	// Lock mode.
	Mode *AssignmentLockMode `json:"mode,omitempty"`
}

// AssignmentOperation - Represents underlying deployment detail for each update to the blueprint assignment.
type AssignmentOperation struct {
	// Properties for AssignmentOperation.
	Properties *AssignmentOperationProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AssignmentOperationList - List of AssignmentOperation.
type AssignmentOperationList struct {
	// List of AssignmentOperation.
	Value []*AssignmentOperation `json:"value,omitempty"`

	// READ-ONLY; Link to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AssignmentOperationProperties - Properties of AssignmentOperation.
type AssignmentOperationProperties struct {
	// State of this blueprint assignment operation.
	AssignmentState *string `json:"assignmentState,omitempty"`

	// The published version of the blueprint definition used for the blueprint assignment operation.
	BlueprintVersion *string `json:"blueprintVersion,omitempty"`

	// List of jobs in this blueprint assignment operation.
	Deployments []*AssignmentDeploymentJob `json:"deployments,omitempty"`

	// Create time of this blueprint assignment operation.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// Finish time of the overall underlying deployments.
	TimeFinished *string `json:"timeFinished,omitempty"`

	// Start time of the underlying deployment.
	TimeStarted *string `json:"timeStarted,omitempty"`
}

// AssignmentOperationsClientGetOptions contains the optional parameters for the AssignmentOperationsClient.Get method.
type AssignmentOperationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssignmentOperationsClientListOptions contains the optional parameters for the AssignmentOperationsClient.List method.
type AssignmentOperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// AssignmentProperties - Detailed properties for a blueprint assignment.
type AssignmentProperties struct {
	// REQUIRED; Blueprint assignment parameter values.
	Parameters map[string]*ParameterValue `json:"parameters,omitempty"`

	// REQUIRED; Names and locations of resource group placeholders.
	ResourceGroups map[string]*ResourceGroupValue `json:"resourceGroups,omitempty"`

	// ID of the published version of a blueprint definition.
	BlueprintID *string `json:"blueprintId,omitempty"`

	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Defines how resources deployed by a blueprint assignment are locked.
	Locks *AssignmentLockSettings `json:"locks,omitempty"`

	// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group
	// level assignments, the property is required.
	Scope *string `json:"scope,omitempty"`

	// READ-ONLY; State of the blueprint assignment.
	ProvisioningState *AssignmentProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of blueprint assignment. This field is readonly.
	Status *AssignmentStatus `json:"status,omitempty" azure:"ro"`
}

// AssignmentStatus - The status of a blueprint assignment. This field is readonly.
type AssignmentStatus struct {
	// READ-ONLY; Last modified time of this blueprint definition.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; List of resources that were created by the blueprint assignment.
	ManagedResources []*string `json:"managedResources,omitempty" azure:"ro"`

	// READ-ONLY; Creation time of this blueprint definition.
	TimeCreated *time.Time `json:"timeCreated,omitempty" azure:"ro"`
}

// AssignmentsClientCreateOrUpdateOptions contains the optional parameters for the AssignmentsClient.CreateOrUpdate method.
type AssignmentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AssignmentsClientDeleteOptions contains the optional parameters for the AssignmentsClient.Delete method.
type AssignmentsClientDeleteOptions struct {
	// When deleteBehavior=all, the resources that were created by the blueprint assignment will be deleted.
	DeleteBehavior *AssignmentDeleteBehavior
}

// AssignmentsClientGetOptions contains the optional parameters for the AssignmentsClient.Get method.
type AssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssignmentsClientListOptions contains the optional parameters for the AssignmentsClient.List method.
type AssignmentsClientListOptions struct {
	// placeholder for future optional parameters
}

// AssignmentsClientWhoIsBlueprintOptions contains the optional parameters for the AssignmentsClient.WhoIsBlueprint method.
type AssignmentsClientWhoIsBlueprintOptions struct {
	// placeholder for future optional parameters
}

// AzureResourceBase - Common properties for all Azure resources.
type AzureResourceBase struct {
	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureResourceManagerError - Error code and message
type AzureResourceManagerError struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message.
	Message *string `json:"message,omitempty"`
}

// Blueprint - Represents a Blueprint definition.
type Blueprint struct {
	// REQUIRED; Detailed properties for blueprint definition.
	Properties *Properties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// BlueprintsClientCreateOrUpdateOptions contains the optional parameters for the BlueprintsClient.CreateOrUpdate method.
type BlueprintsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// BlueprintsClientDeleteOptions contains the optional parameters for the BlueprintsClient.Delete method.
type BlueprintsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// BlueprintsClientGetOptions contains the optional parameters for the BlueprintsClient.Get method.
type BlueprintsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BlueprintsClientListOptions contains the optional parameters for the BlueprintsClient.List method.
type BlueprintsClientListOptions struct {
	// placeholder for future optional parameters
}

type CloudError struct {
	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows
	// the OData error response format.)
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.)
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorResponse `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// KeyVaultReference - Specifies the link to a Key Vault.
type KeyVaultReference struct {
	// REQUIRED; Azure resource ID of the Key Vault.
	ID *string `json:"id,omitempty"`
}

// List of blueprint definitions.
type List struct {
	// List of blueprint definitions.
	Value []*Blueprint `json:"value,omitempty"`

	// READ-ONLY; Link to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ManagedServiceIdentity - Managed identity generic object.
type ManagedServiceIdentity struct {
	// REQUIRED; Type of the managed identity.
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// Azure Active Directory principal ID associated with this Identity.
	PrincipalID *string `json:"principalId,omitempty"`

	// ID of the Azure Active Directory.
	TenantID *string `json:"tenantId,omitempty"`

	// The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed
	// identity.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`
}

// ParameterDefinition - Represent a parameter with constrains and metadata.
type ParameterDefinition struct {
	// REQUIRED; Allowed data types for Resource Manager template parameters.
	Type *TemplateParameterType `json:"type,omitempty"`

	// Array of allowed values for this parameter.
	AllowedValues []interface{} `json:"allowedValues,omitempty"`

	// Default Value for this parameter.
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// User-friendly properties for this parameter.
	Metadata *ParameterDefinitionMetadata `json:"metadata,omitempty"`
}

// ParameterDefinitionMetadata - User-friendly properties for this parameter.
type ParameterDefinitionMetadata struct {
	// Description of this parameter/resourceGroup.
	Description *string `json:"description,omitempty"`

	// DisplayName of this parameter/resourceGroup.
	DisplayName *string `json:"displayName,omitempty"`

	// StrongType for UI to render rich experience during blueprint assignment. Supported strong types are resourceType, principalId
	// and location.
	StrongType *string `json:"strongType,omitempty"`
}

// ParameterValue - Value for the specified parameter. Can be either 'value' or 'reference' but not both.
type ParameterValue struct {
	// Parameter value as reference type.
	Reference *SecretValueReference `json:"reference,omitempty"`

	// Parameter value. Any valid JSON value is allowed including objects, arrays, strings, numbers and booleans.
	Value interface{} `json:"value,omitempty"`
}

// PolicyAssignmentArtifact - Blueprint artifact that applies a Policy assignment.
type PolicyAssignmentArtifact struct {
	// REQUIRED; Specifies the kind of blueprint artifact.
	Kind *ArtifactKind `json:"kind,omitempty"`

	// REQUIRED; properties for policyAssignment Artifact
	Properties *PolicyAssignmentArtifactProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PolicyAssignmentArtifactProperties - Properties of a Policy assignment blueprint artifact.
type PolicyAssignmentArtifactProperties struct {
	// REQUIRED; Parameter values for the policy definition.
	Parameters map[string]*ParameterValue `json:"parameters,omitempty"`

	// REQUIRED; Azure resource ID of the policy definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []*string `json:"dependsOn,omitempty"`

	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
}

// Properties - Schema for blueprint definition properties.
type Properties struct {
	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Parameters required by this blueprint definition.
	Parameters map[string]*ParameterDefinition `json:"parameters,omitempty"`

	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]*ResourceGroupDefinition `json:"resourceGroups,omitempty"`

	// The scope where this blueprint definition can be assigned.
	TargetScope *BlueprintTargetScope `json:"targetScope,omitempty"`

	// Published versions of this blueprint definition.
	Versions interface{} `json:"versions,omitempty"`

	// READ-ONLY; Layout view of the blueprint definition for UI reference.
	Layout interface{} `json:"layout,omitempty" azure:"ro"`

	// READ-ONLY; Status of the blueprint. This field is readonly.
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// PublishedArtifactsClientGetOptions contains the optional parameters for the PublishedArtifactsClient.Get method.
type PublishedArtifactsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PublishedArtifactsClientListOptions contains the optional parameters for the PublishedArtifactsClient.List method.
type PublishedArtifactsClientListOptions struct {
	// placeholder for future optional parameters
}

// PublishedBlueprint - Represents a published blueprint.
type PublishedBlueprint struct {
	// REQUIRED; Detailed properties for published blueprint.
	Properties *PublishedBlueprintProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PublishedBlueprintList - List of published blueprint definitions.
type PublishedBlueprintList struct {
	// List of published blueprint definitions.
	Value []*PublishedBlueprint `json:"value,omitempty"`

	// READ-ONLY; Link to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// PublishedBlueprintProperties - Schema for published blueprint definition properties.
type PublishedBlueprintProperties struct {
	// Name of the published blueprint definition.
	BlueprintName *string `json:"blueprintName,omitempty"`

	// Version-specific change notes.
	ChangeNotes *string `json:"changeNotes,omitempty"`

	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Parameters required by this blueprint definition.
	Parameters map[string]*ParameterDefinition `json:"parameters,omitempty"`

	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]*ResourceGroupDefinition `json:"resourceGroups,omitempty"`

	// The scope where this blueprint definition can be assigned.
	TargetScope *BlueprintTargetScope `json:"targetScope,omitempty"`

	// READ-ONLY; Status of the blueprint. This field is readonly.
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// PublishedBlueprintsClientCreateOptions contains the optional parameters for the PublishedBlueprintsClient.Create method.
type PublishedBlueprintsClientCreateOptions struct {
	// Published Blueprint to create or update.
	PublishedBlueprint *PublishedBlueprint
}

// PublishedBlueprintsClientDeleteOptions contains the optional parameters for the PublishedBlueprintsClient.Delete method.
type PublishedBlueprintsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PublishedBlueprintsClientGetOptions contains the optional parameters for the PublishedBlueprintsClient.Get method.
type PublishedBlueprintsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PublishedBlueprintsClientListOptions contains the optional parameters for the PublishedBlueprintsClient.List method.
type PublishedBlueprintsClientListOptions struct {
	// placeholder for future optional parameters
}

// ResourceGroupDefinition - Represents an Azure resource group in a blueprint definition.
type ResourceGroupDefinition struct {
	// Artifacts which need to be deployed before this resource group.
	DependsOn []*string `json:"dependsOn,omitempty"`

	// Location of this resourceGroup. Leave empty if the resource group location will be specified during the blueprint assignment.
	Location *string `json:"location,omitempty"`

	// User-friendly properties for this resource group.
	Metadata *ParameterDefinitionMetadata `json:"metadata,omitempty"`

	// Name of this resourceGroup. Leave empty if the resource group name will be specified during the blueprint assignment.
	Name *string `json:"name,omitempty"`

	// Tags to be assigned to this resource group.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ResourceGroupValue - Represents an Azure resource group.
type ResourceGroupValue struct {
	// Location of the resource group.
	Location *string `json:"location,omitempty"`

	// Name of the resource group.
	Name *string `json:"name,omitempty"`
}

// ResourcePropertiesBase - Shared properties between all blueprint resources.
type ResourcePropertiesBase struct {
	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`
}

// ResourceProviderOperation - Supported operations of this resource provider.
type ResourceProviderOperation struct {
	// Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`

	// Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// ResourceProviderOperationDisplay - Display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Description of this operation.
	Description *string `json:"description,omitempty"`

	// Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Resource provider: Microsoft Blueprint.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// ResourceProviderOperationList - Results of the request to list operations.
type ResourceProviderOperationList struct {
	// List of operations supported by this resource provider.
	Value []*ResourceProviderOperation `json:"value,omitempty"`
}

// ResourceStatusBase - Shared status properties between all blueprint resources.
type ResourceStatusBase struct {
	// READ-ONLY; Last modified time of this blueprint definition.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; Creation time of this blueprint definition.
	TimeCreated *time.Time `json:"timeCreated,omitempty" azure:"ro"`
}

// RoleAssignmentArtifact - Blueprint artifact that applies a Role assignment.
type RoleAssignmentArtifact struct {
	// REQUIRED; Specifies the kind of blueprint artifact.
	Kind *ArtifactKind `json:"kind,omitempty"`

	// REQUIRED; Properties for a Role assignment blueprint artifact.
	Properties *RoleAssignmentArtifactProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RoleAssignmentArtifactProperties - Properties of a Role assignment blueprint artifact.
type RoleAssignmentArtifactProperties struct {
	// REQUIRED; Array of user or group identities in Azure Active Directory. The roleDefinition will apply to each identity.
	PrincipalIDs interface{} `json:"principalIds,omitempty"`

	// REQUIRED; Azure resource ID of the RoleDefinition.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []*string `json:"dependsOn,omitempty"`

	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// RoleAssignment will be scope to this resourceGroup. If empty, it scopes to the subscription.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
}

// SecretValueReference - Reference to a Key Vault secret.
type SecretValueReference struct {
	// REQUIRED; Specifies the reference to a given Azure Key Vault.
	KeyVault *KeyVaultReference `json:"keyVault,omitempty"`

	// REQUIRED; Name of the secret.
	SecretName *string `json:"secretName,omitempty"`

	// The version of the secret to use. If left blank, the latest version of the secret is used.
	SecretVersion *string `json:"secretVersion,omitempty"`
}

// SharedBlueprintProperties - Shared Schema for both blueprintProperties and publishedBlueprintProperties.
type SharedBlueprintProperties struct {
	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Parameters required by this blueprint definition.
	Parameters map[string]*ParameterDefinition `json:"parameters,omitempty"`

	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]*ResourceGroupDefinition `json:"resourceGroups,omitempty"`

	// The scope where this blueprint definition can be assigned.
	TargetScope *BlueprintTargetScope `json:"targetScope,omitempty"`

	// READ-ONLY; Status of the blueprint. This field is readonly.
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// Status - The status of the blueprint. This field is readonly.
type Status struct {
	// READ-ONLY; Last modified time of this blueprint definition.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; Creation time of this blueprint definition.
	TimeCreated *time.Time `json:"timeCreated,omitempty" azure:"ro"`
}

// TemplateArtifact - Blueprint artifact that deploys a Resource Manager template.
type TemplateArtifact struct {
	// REQUIRED; Specifies the kind of blueprint artifact.
	Kind *ArtifactKind `json:"kind,omitempty"`

	// REQUIRED; Properties for a Resource Manager template blueprint artifact.
	Properties *TemplateArtifactProperties `json:"properties,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TemplateArtifactProperties - Properties of a Resource Manager template blueprint artifact.
type TemplateArtifactProperties struct {
	// REQUIRED; Resource Manager template blueprint artifact parameter values.
	Parameters map[string]*ParameterValue `json:"parameters,omitempty"`

	// REQUIRED; The Resource Manager template blueprint artifact body.
	Template interface{} `json:"template,omitempty"`

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []*string `json:"dependsOn,omitempty"`

	// Multi-line explain this resource.
	Description *string `json:"description,omitempty"`

	// One-liner string explain this resource.
	DisplayName *string `json:"displayName,omitempty"`

	// If applicable, the name of the resource group placeholder to which the Resource Manager template blueprint artifact will
	// be deployed.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
}

// TrackedResource - Common properties for all Azure tracked resources.
type TrackedResource struct {
	// REQUIRED; The location of this blueprint assignment.
	Location *string `json:"location,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// UserAssignedIdentity - User-assigned managed identity.
type UserAssignedIdentity struct {
	// Client App Id associated with this identity.
	ClientID *string `json:"clientId,omitempty"`

	// Azure Active Directory principal ID associated with this Identity.
	PrincipalID *string `json:"principalId,omitempty"`
}

// WhoIsBlueprintContract - Response schema for querying the Azure Blueprints service principal in the tenant.
type WhoIsBlueprintContract struct {
	// AAD object Id of the Azure Blueprints service principal in the tenant.
	ObjectID *string `json:"objectId,omitempty"`
}
