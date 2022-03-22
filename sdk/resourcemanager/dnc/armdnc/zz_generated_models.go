//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnc

// ControllerClientBeginCreateOptions contains the optional parameters for the ControllerClient.BeginCreate method.
type ControllerClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ControllerClientBeginDeleteOptions contains the optional parameters for the ControllerClient.BeginDelete method.
type ControllerClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ControllerClientGetDetailsOptions contains the optional parameters for the ControllerClient.GetDetails method.
type ControllerClientGetDetailsOptions struct {
	// placeholder for future optional parameters
}

// ControllerClientPatchOptions contains the optional parameters for the ControllerClient.Patch method.
type ControllerClientPatchOptions struct {
	// placeholder for future optional parameters
}

// ControllerDetails - controller details
type ControllerDetails struct {
	// controller arm resource id
	ID *string `json:"id,omitempty"`
}

// ControllerResource - Represents an instance of a resource.
type ControllerResource struct {
	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ControllerResourceUpdateParameters - Parameters for updating a resource.
type ControllerResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// DelegatedController - Represents an instance of a DNC controller.
type DelegatedController struct {
	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Properties of the provision operation request.
	Properties *DelegatedControllerProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DelegatedControllerProperties - Properties of Delegated controller resource.
type DelegatedControllerProperties struct {
	// READ-ONLY; dnc application id should be used by customer to authenticate with dnc gateway.
	DncAppID *string `json:"dncAppId,omitempty" azure:"ro"`

	// READ-ONLY; dnc endpoint url that customers can use to connect to
	DncEndpoint *string `json:"dncEndpoint,omitempty" azure:"ro"`

	// READ-ONLY; tenant id of dnc application id
	DncTenantID *string `json:"dncTenantId,omitempty" azure:"ro"`

	// READ-ONLY; The current state of dnc controller resource.
	ProvisioningState *ControllerState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource guid.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// DelegatedControllers - An array of Delegated controller resources.
type DelegatedControllers struct {
	// REQUIRED; An array of Delegated controller resources.
	Value []*DelegatedController `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of controllers.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// DelegatedNetworkClientListByResourceGroupOptions contains the optional parameters for the DelegatedNetworkClient.ListByResourceGroup
// method.
type DelegatedNetworkClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DelegatedNetworkClientListBySubscriptionOptions contains the optional parameters for the DelegatedNetworkClient.ListBySubscription
// method.
type DelegatedNetworkClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnet - Represents an instance of a orchestrator.
type DelegatedSubnet struct {
	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// Properties of the provision operation request.
	Properties *DelegatedSubnetProperties `json:"properties,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DelegatedSubnetProperties - Properties of delegated subnet
type DelegatedSubnetProperties struct {
	// Properties of the controller.
	ControllerDetails *ControllerDetails `json:"controllerDetails,omitempty"`

	// subnet details
	SubnetDetails *SubnetDetails `json:"subnetDetails,omitempty"`

	// READ-ONLY; The current state of dnc delegated subnet resource.
	ProvisioningState *DelegatedSubnetState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource guid.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// DelegatedSubnetResource - Represents an instance of a resource.
type DelegatedSubnetResource struct {
	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DelegatedSubnetServiceClientBeginDeleteDetailsOptions contains the optional parameters for the DelegatedSubnetServiceClient.BeginDeleteDetails
// method.
type DelegatedSubnetServiceClientBeginDeleteDetailsOptions struct {
	// Force delete resource
	ForceDelete *bool
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DelegatedSubnetServiceClientBeginPatchDetailsOptions contains the optional parameters for the DelegatedSubnetServiceClient.BeginPatchDetails
// method.
type DelegatedSubnetServiceClientBeginPatchDetailsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DelegatedSubnetServiceClientBeginPutDetailsOptions contains the optional parameters for the DelegatedSubnetServiceClient.BeginPutDetails
// method.
type DelegatedSubnetServiceClientBeginPutDetailsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DelegatedSubnetServiceClientGetDetailsOptions contains the optional parameters for the DelegatedSubnetServiceClient.GetDetails
// method.
type DelegatedSubnetServiceClientGetDetailsOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnetServiceClientListByResourceGroupOptions contains the optional parameters for the DelegatedSubnetServiceClient.ListByResourceGroup
// method.
type DelegatedSubnetServiceClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnetServiceClientListBySubscriptionOptions contains the optional parameters for the DelegatedSubnetServiceClient.ListBySubscription
// method.
type DelegatedSubnetServiceClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnets - An array of DelegatedSubnet resources.
type DelegatedSubnets struct {
	// REQUIRED; An array of DelegatedSubnet resources.
	Value []*DelegatedSubnet `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of DelegatedSubnet resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Orchestrator - Represents an instance of a orchestrator.
type Orchestrator struct {
	// REQUIRED; The kind of workbook. Choices are user and shared.
	Kind *OrchestratorKind `json:"kind,omitempty"`

	// The identity of the orchestrator
	Identity *OrchestratorIdentity `json:"identity,omitempty"`

	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// Properties of the provision operation request.
	Properties *OrchestratorResourceProperties `json:"properties,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

type OrchestratorIdentity struct {
	// The type of identity used for orchestrator cluster. Type 'SystemAssigned' will use an implicitly created identity orchestrator
	// clusters
	Type *ResourceIdentityType `json:"type,omitempty"`

	// READ-ONLY; The principal id of the system assigned identity which is used by orchestrator.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant id of the system assigned identity which is used by orchestrator.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// OrchestratorInstanceServiceClientBeginCreateOptions contains the optional parameters for the OrchestratorInstanceServiceClient.BeginCreate
// method.
type OrchestratorInstanceServiceClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OrchestratorInstanceServiceClientBeginDeleteOptions contains the optional parameters for the OrchestratorInstanceServiceClient.BeginDelete
// method.
type OrchestratorInstanceServiceClientBeginDeleteOptions struct {
	// Force delete resource
	ForceDelete *bool
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OrchestratorInstanceServiceClientGetDetailsOptions contains the optional parameters for the OrchestratorInstanceServiceClient.GetDetails
// method.
type OrchestratorInstanceServiceClientGetDetailsOptions struct {
	// placeholder for future optional parameters
}

// OrchestratorInstanceServiceClientListByResourceGroupOptions contains the optional parameters for the OrchestratorInstanceServiceClient.ListByResourceGroup
// method.
type OrchestratorInstanceServiceClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// OrchestratorInstanceServiceClientListBySubscriptionOptions contains the optional parameters for the OrchestratorInstanceServiceClient.ListBySubscription
// method.
type OrchestratorInstanceServiceClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// OrchestratorInstanceServiceClientPatchOptions contains the optional parameters for the OrchestratorInstanceServiceClient.Patch
// method.
type OrchestratorInstanceServiceClientPatchOptions struct {
	// placeholder for future optional parameters
}

// OrchestratorResource - Represents an instance of a resource.
type OrchestratorResource struct {
	// REQUIRED; The kind of workbook. Choices are user and shared.
	Kind *OrchestratorKind `json:"kind,omitempty"`

	// The identity of the orchestrator
	Identity *OrchestratorIdentity `json:"identity,omitempty"`

	// Location of the resource.
	Location *string `json:"location,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; An identifier that represents the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OrchestratorResourceProperties - Properties of orchestrator
type OrchestratorResourceProperties struct {
	// REQUIRED; Properties of the controller.
	ControllerDetails *ControllerDetails `json:"controllerDetails,omitempty"`

	// K8s APIServer url. Either one of apiServerEndpoint or privateLinkResourceId can be specified
	APIServerEndpoint *string `json:"apiServerEndpoint,omitempty"`

	// RootCA certificate of kubernetes cluster base64 encoded
	ClusterRootCA *string `json:"clusterRootCA,omitempty"`

	// AAD ID used with apiserver
	OrchestratorAppID *string `json:"orchestratorAppId,omitempty"`

	// TenantID of server App ID
	OrchestratorTenantID *string `json:"orchestratorTenantId,omitempty"`

	// private link arm resource id. Either one of apiServerEndpoint or privateLinkResourceId can be specified
	PrivateLinkResourceID *string `json:"privateLinkResourceId,omitempty"`

	// READ-ONLY; The current state of orchestratorInstance resource.
	ProvisioningState *OrchestratorInstanceState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource guid.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// OrchestratorResourceUpdateParameters - Parameters for updating a resource.
type OrchestratorResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// Orchestrators - An array of OrchestratorInstance resources.
type Orchestrators struct {
	// REQUIRED; An array of OrchestratorInstance resources.
	Value []*Orchestrator `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of orchestrators.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ResourceUpdateParameters - Parameters for updating a resource.
type ResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// SubnetDetails - Properties of orchestrator
type SubnetDetails struct {
	// subnet arm resource id
	ID *string `json:"id,omitempty"`
}
