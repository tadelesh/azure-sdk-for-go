//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armEnterpriseKnowledgeGraph

// ClientCreateOptions contains the optional parameters for the Client.Create method.
type ClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteOptions contains the optional parameters for the Client.Delete method.
type ClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
type ClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.List method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// ClientUpdateOptions contains the optional parameters for the Client.Update method.
type ClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnterpriseKnowledgeGraph resource definition
type EnterpriseKnowledgeGraph struct {
	// Specifies the location of the resource.
	Location *string `json:"location,omitempty"`

	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties *Properties `json:"properties,omitempty"`

	// Gets or sets the SKU of the resource.
	SKU *SKU `json:"sku,omitempty"`

	// Contains resource tags defined as key/value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Specifies the resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Error - EnterpriseKnowledgeGraph Service error object.
type Error struct {
	// The error body.
	Error *ErrorBody `json:"error,omitempty"`
}

// ErrorBody - EnterpriseKnowledgeGraph Service error body.
type ErrorBody struct {
	// REQUIRED; error code
	Code *string `json:"code,omitempty"`

	// REQUIRED; error message
	Message *string `json:"message,omitempty"`
}

// OperationDisplayInfo - The operation supported by EnterpriseKnowledgeGraph Service Management.
type OperationDisplayInfo struct {
	// The description of the operation.
	Description *string `json:"description,omitempty"`

	// The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft EnterpriseKnowledgeGraph Service.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity - The operations supported by EnterpriseKnowledgeGraph Service Management.
type OperationEntity struct {
	// The operation supported by EnterpriseKnowledgeGraph Service Management.
	Display *OperationDisplayInfo `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`

	// The origin of the operation.
	Origin *string `json:"origin,omitempty"`

	// Additional properties.
	Properties interface{} `json:"properties,omitempty"`
}

// OperationEntityListResult - The list of EnterpriseKnowledgeGraph service operation response.
type OperationEntityListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of operations.
	Value []*OperationEntity `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Properties - The parameters to provide for the EnterpriseKnowledgeGraph.
type Properties struct {
	// The description of the EnterpriseKnowledgeGraph
	Description *string `json:"description,omitempty"`

	// Specifies the metadata of the resource.
	Metadata interface{} `json:"metadata,omitempty"`

	// The state of EnterpriseKnowledgeGraph provisioning
	ProvisioningState *EnterpriseKnowledgeGraphPropertiesProvisioningState `json:"provisioningState,omitempty"`
}

// Resource - Azure resource
type Resource struct {
	// Specifies the location of the resource.
	Location *string `json:"location,omitempty"`

	// Gets or sets the SKU of the resource.
	SKU *SKU `json:"sku,omitempty"`

	// Contains resource tags defined as key/value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Specifies the resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResponseList - The list of EnterpriseKnowledgeGraph service operation response.
type ResponseList struct {
	// The link used to get the next page of EnterpriseKnowledgeGraph service resources.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; Gets the list of EnterpriseKnowledgeGraph service results and their properties.
	Value []*EnterpriseKnowledgeGraph `json:"value,omitempty" azure:"ro"`
}

// SKU - The SKU of the EnterpriseKnowledgeGraph service account.
type SKU struct {
	// REQUIRED; The sku name
	Name *SKUName `json:"name,omitempty"`
}
