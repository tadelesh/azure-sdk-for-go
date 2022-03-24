//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningexperimentation

import "time"

// Account - An object that represents a machine learning team account.
type Account struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties of the machine learning team account.
	Properties *AccountProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountListResult - The result of a request to list machine learning team accounts.
type AccountListResult struct {
	// The URI that can be used to request the next list of machine learning team accounts.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of machine learning team accounts. Since this list may be incomplete, the nextLink field should be used to request
	// the next list of machine learning team accounts.
	Value []*Account `json:"value,omitempty"`
}

// AccountProperties - The properties of a machine learning team account.
type AccountProperties struct {
	// REQUIRED; The fully qualified arm id of the user key vault.
	KeyVaultID *string `json:"keyVaultId,omitempty"`

	// REQUIRED; The properties of the storage account for the machine learning team account.
	StorageAccount *StorageAccountProperties `json:"storageAccount,omitempty"`

	// REQUIRED; The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountID *string `json:"vsoAccountId,omitempty"`

	// The description of this workspace.
	Description *string `json:"description,omitempty"`

	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `json:"seats,omitempty"`

	// READ-ONLY; The immutable id associated with this team account.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; The creation date of the machine learning team account in ISO8601 format.
	CreationDate *time.Time `json:"creationDate,omitempty" azure:"ro"`

	// READ-ONLY; The uri for this machine learning team account.
	DiscoveryURI *string `json:"discoveryUri,omitempty" azure:"ro"`

	// READ-ONLY; The current deployment state of team account resource. The provisioningState is to indicate states for resource
	// provisioning.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// AccountPropertiesUpdateParameters - The parameters for updating the properties of a machine learning team account.
type AccountPropertiesUpdateParameters struct {
	// The description of this workspace.
	Description *string `json:"description,omitempty"`

	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `json:"seats,omitempty"`

	// The key for storage account associated with this team account
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`
}

// AccountUpdateParameters - The parameters for updating a machine learning team account.
type AccountUpdateParameters struct {
	// The properties that the machine learning team account will be updated with.
	Properties *AccountPropertiesUpdateParameters `json:"properties,omitempty"`

	// The resource tags for the machine learning team account.
	Tags map[string]*string `json:"tags,omitempty"`
}

// AccountsClientCreateOrUpdateOptions contains the optional parameters for the AccountsClient.CreateOrUpdate method.
type AccountsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientDeleteOptions contains the optional parameters for the AccountsClient.Delete method.
type AccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup method.
type AccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListOptions contains the optional parameters for the AccountsClient.List method.
type AccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientUpdateOptions contains the optional parameters for the AccountsClient.Update method.
type AccountsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ErrorResponse - The error response send when an operation fails.
type ErrorResponse struct {
	// REQUIRED; error code
	Code *string `json:"code,omitempty"`

	// REQUIRED; error message
	Message *string `json:"message,omitempty"`
}

// Operation - Azure Machine Learning team account REST API operation
type Operation struct {
	// Display name of operation
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - Display name of operation
type OperationDisplay struct {
	// The description for the operation.
	Description *string `json:"description,omitempty"`

	// The operation that users can perform.
	Operation *string `json:"operation,omitempty"`

	// The resource provider name: Microsoft.MachineLearningExperimentation
	Provider *string `json:"provider,omitempty"`

	// The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - An array of operations supported by the resource provider.
type OperationListResult struct {
	// List of AML team account operations supported by the AML team account resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Project - An object that represents a machine learning project.
type Project struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties of the Project.
	Properties *ProjectProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ProjectListResult - The result of a request to list projects.
type ProjectListResult struct {
	// The URI that can be used to request the next list of projects.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of projects. Since this list may be incomplete, the nextLink field should be used to request the next list of
	// projects.
	Value []*Project `json:"value,omitempty"`
}

// ProjectProperties - The properties of a machine learning project.
type ProjectProperties struct {
	// REQUIRED; The friendly name for this project.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The description of this project.
	Description *string `json:"description,omitempty"`

	// The reference to git repo for this project.
	Gitrepo *string `json:"gitrepo,omitempty"`

	// READ-ONLY; The immutable id of the team account which contains this project.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; The creation date of the project in ISO8601 format.
	CreationDate *time.Time `json:"creationDate,omitempty" azure:"ro"`

	// READ-ONLY; The immutable id of this project.
	ProjectID *string `json:"projectId,omitempty" azure:"ro"`

	// READ-ONLY; The current deployment state of project resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The immutable id of the workspace which contains this project.
	WorkspaceID *string `json:"workspaceId,omitempty" azure:"ro"`
}

// ProjectPropertiesUpdateParameters - The parameters for updating the properties of a project.
type ProjectPropertiesUpdateParameters struct {
	// The description of this project.
	Description *string `json:"description,omitempty"`

	// The friendly name for this project.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The reference to git repo for this project.
	Gitrepo *string `json:"gitrepo,omitempty"`
}

// ProjectUpdateParameters - The parameters for updating a machine learning project.
type ProjectUpdateParameters struct {
	// The properties that the project will be updated with.
	Properties *ProjectPropertiesUpdateParameters `json:"properties,omitempty"`

	// The resource tags for the machine learning project.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ProjectsClientCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.CreateOrUpdate method.
type ProjectsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientDeleteOptions contains the optional parameters for the ProjectsClient.Delete method.
type ProjectsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
type ProjectsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientListByWorkspaceOptions contains the optional parameters for the ProjectsClient.ListByWorkspace method.
type ProjectsClientListByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
type ProjectsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Resource - An Azure resource.
type Resource struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// StorageAccountProperties - The properties of a storage account for a machine learning team account.
type StorageAccountProperties struct {
	// REQUIRED; The access key to the storage account.
	AccessKey *string `json:"accessKey,omitempty"`

	// REQUIRED; The fully qualified arm Id of the storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
}

// Workspace - An object that represents a machine learning team account workspace.
type Workspace struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties of the machine learning team account workspace.
	Properties *WorkspaceProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// WorkspaceListResult - The result of a request to list machine learning team account workspaces.
type WorkspaceListResult struct {
	// The URI that can be used to request the next list of machine learning workspaces.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of machine learning team account workspaces. Since this list may be incomplete, the nextLink field should be used
	// to request the next list of machine learning team accounts.
	Value []*Workspace `json:"value,omitempty"`
}

// WorkspaceProperties - The properties of a machine learning team account workspace.
type WorkspaceProperties struct {
	// REQUIRED; The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object
	// gets created
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The description of this workspace.
	Description *string `json:"description,omitempty"`

	// READ-ONLY; The immutable id of the team account which contains this workspace.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; The creation date of the machine learning workspace in ISO8601 format.
	CreationDate *time.Time `json:"creationDate,omitempty" azure:"ro"`

	// READ-ONLY; The current deployment state of team account workspace resource. The provisioningState is to indicate states
	// for resource provisioning.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The immutable id of this workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" azure:"ro"`
}

// WorkspacePropertiesUpdateParameters - The parameters for updating the properties of a machine learning team account workspace.
type WorkspacePropertiesUpdateParameters struct {
	// Description for this workspace.
	Description *string `json:"description,omitempty"`

	// Friendly name of this workspace.
	FriendlyName *string `json:"friendlyName,omitempty"`
}

// WorkspaceUpdateParameters - The parameters for updating a machine learning team account workspace.
type WorkspaceUpdateParameters struct {
	// The properties that the machine learning workspace will be updated with.
	Properties *WorkspacePropertiesUpdateParameters `json:"properties,omitempty"`

	// The resource tags for the machine learning team account workspace.
	Tags map[string]*string `json:"tags,omitempty"`
}

// WorkspacesClientCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.CreateOrUpdate method.
type WorkspacesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientDeleteOptions contains the optional parameters for the WorkspacesClient.Delete method.
type WorkspacesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
type WorkspacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListByAccountsOptions contains the optional parameters for the WorkspacesClient.ListByAccounts method.
type WorkspacesClientListByAccountsOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientUpdateOptions contains the optional parameters for the WorkspacesClient.Update method.
type WorkspacesClientUpdateOptions struct {
	// placeholder for future optional parameters
}