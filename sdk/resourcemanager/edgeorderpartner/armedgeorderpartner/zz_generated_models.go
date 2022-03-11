//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

import "time"

// APISClientBeginManageInventoryMetadataOptions contains the optional parameters for the APISClient.BeginManageInventoryMetadata
// method.
type APISClientBeginManageInventoryMetadataOptions struct {
	// placeholder for future optional parameters
}

// APISClientListOperationsPartnerOptions contains the optional parameters for the APISClient.ListOperationsPartner method.
type APISClientListOperationsPartnerOptions struct {
	// placeholder for future optional parameters
}

// APISClientManageLinkOptions contains the optional parameters for the APISClient.ManageLink method.
type APISClientManageLinkOptions struct {
	// placeholder for future optional parameters
}

// APISClientSearchInventoriesOptions contains the optional parameters for the APISClient.SearchInventories method.
type APISClientSearchInventoriesOptions struct {
	// placeholder for future optional parameters
}

type AdditionalErrorInfo struct {
	// Anything
	Info interface{} `json:"info,omitempty"`
	Type *string     `json:"type,omitempty"`
}

// AdditionalInventoryDetails - Contains additional data about inventory in dictionary format
type AdditionalInventoryDetails struct {
	// READ-ONLY; Additional Data
	AdditionalData map[string]*string `json:"additionalData,omitempty" azure:"ro"`
}

// AdditionalOrderItemDetails - Contains additional order item details
type AdditionalOrderItemDetails struct {
	// READ-ONLY; Order item status
	Status *StageDetails `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; Subscription details
	Subscription *SubscriptionDetails `json:"subscription,omitempty" azure:"ro"`
}

// BillingDetails - Contains billing details for the inventory
type BillingDetails struct {
	// READ-ONLY; Billing type for the inventory
	BillingType *string `json:"billingType,omitempty" azure:"ro"`

	// READ-ONLY; Billing status for the inventory
	Status *string `json:"status,omitempty" azure:"ro"`
}

type CloudError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`

	// READ-ONLY
	AdditionalInfo []*AdditionalErrorInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY
	Details []*CloudError `json:"details,omitempty" azure:"ro"`
}

// ConfigurationData - Contains information about inventory configuration
type ConfigurationData struct {
	// READ-ONLY; Configuration identifier of inventory
	ConfigurationIdentifier *string `json:"configurationIdentifier,omitempty" azure:"ro"`

	// READ-ONLY; Configuration identifier on device - this is used in case of any mismatch between actual configuration on inventory
	// and configuration stored in service
	ConfigurationIdentifierOnDevice *string `json:"configurationIdentifierOnDevice,omitempty" azure:"ro"`

	// READ-ONLY; Family identifier of inventory
	FamilyIdentifier *string `json:"familyIdentifier,omitempty" azure:"ro"`

	// READ-ONLY; Product identifier of inventory
	ProductIdentifier *string `json:"productIdentifier,omitempty" azure:"ro"`

	// READ-ONLY; Product Line identifier of inventory
	ProductLineIdentifier *string `json:"productLineIdentifier,omitempty" azure:"ro"`
}

// ConfigurationDetails - Contains additional configuration details about inventory
type ConfigurationDetails struct {
	// READ-ONLY; Collection of specification details about the inventory
	Specifications []*SpecificationDetails `json:"specifications,omitempty" azure:"ro"`
}

// ConfigurationOnDevice - Configuration parameters for ManageInventoryMetadata call
type ConfigurationOnDevice struct {
	// REQUIRED; Configuration identifier on device
	ConfigurationIdentifier *string `json:"configurationIdentifier,omitempty"`
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

// InventoryAdditionalDetails - Represents additional details about the partner inventory
type InventoryAdditionalDetails struct {
	// Represents additional details about the order item
	OrderItem *AdditionalOrderItemDetails `json:"orderItem,omitempty"`

	// READ-ONLY; Represents additional details about billing for the inventory
	Billing *BillingDetails `json:"billing,omitempty" azure:"ro"`

	// READ-ONLY; Represents additional details about the configuration
	Configuration *ConfigurationDetails `json:"configuration,omitempty" azure:"ro"`

	// READ-ONLY; Represents additional data about the inventory
	Inventory *AdditionalInventoryDetails `json:"inventory,omitempty" azure:"ro"`

	// READ-ONLY; Contains inventory metadata
	InventoryMetadata *string `json:"inventoryMetadata,omitempty" azure:"ro"`

	// READ-ONLY; Represents secrets on the inventory
	InventorySecrets map[string]*string `json:"inventorySecrets,omitempty" azure:"ro"`
}

// InventoryData - Contains basic information about inventory
type InventoryData struct {
	// READ-ONLY; Inventory location
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Boolean flag to indicate if registration is allowed
	RegistrationAllowed *bool `json:"registrationAllowed,omitempty" azure:"ro"`

	// READ-ONLY; Inventory status
	Status *string `json:"status,omitempty" azure:"ro"`
}

// InventoryProperties - Represents inventory properties
type InventoryProperties struct {
	// READ-ONLY; Represents basic configuration data.
	Configuration *ConfigurationData `json:"configuration,omitempty" azure:"ro"`

	// READ-ONLY; Represents additional details of inventory
	Details *InventoryAdditionalDetails `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Represents basic inventory data.
	Inventory *InventoryData `json:"inventory,omitempty" azure:"ro"`

	// READ-ONLY; Location of inventory
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Represents management resource data associated with inventory.
	ManagementResource *ManagementResourceData `json:"managementResource,omitempty" azure:"ro"`

	// READ-ONLY; Represents basic order item data.
	OrderItem *OrderItemData `json:"orderItem,omitempty" azure:"ro"`

	// READ-ONLY; Serial number of the device.
	SerialNumber *string `json:"serialNumber,omitempty" azure:"ro"`
}

// ManageInventoryMetadataRequest - Request body for ManageInventoryMetadata call
type ManageInventoryMetadataRequest struct {
	// REQUIRED; Inventory metadata to be updated
	InventoryMetadata *string `json:"inventoryMetadata,omitempty"`

	// Inventory configuration to be updated
	ConfigurationOnDevice *ConfigurationOnDevice `json:"configurationOnDevice,omitempty"`
}

// ManageLinkRequest - Request body for ManageLink call
type ManageLinkRequest struct {
	// REQUIRED; Arm Id of the management resource to which inventory is to be linked For unlink operation, enter empty string
	ManagementResourceArmID *string `json:"managementResourceArmId,omitempty"`

	// REQUIRED; Operation to be performed - Link, Unlink, Relink
	Operation *ManageLinkOperation `json:"operation,omitempty"`

	// REQUIRED; Tenant ID of management resource associated with inventory
	TenantID *string `json:"tenantId,omitempty"`
}

// ManagementResourceData - Contains information about management resource
type ManagementResourceData struct {
	// READ-ONLY; Arm ID of management resource associated with inventory
	ArmID *string `json:"armId,omitempty" azure:"ro"`

	// READ-ONLY; Tenant ID of management resource associated with inventory
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
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

// OrderItemData - Contains information about the order item to which inventory belongs
type OrderItemData struct {
	// READ-ONLY; Arm ID of order item
	ArmID *string `json:"armId,omitempty" azure:"ro"`

	// READ-ONLY; Order item type - purchase or rental
	OrderItemType *OrderItemType `json:"orderItemType,omitempty" azure:"ro"`
}

// PartnerInventory - Represents partner inventory contract
type PartnerInventory struct {
	// READ-ONLY; Inventory properties
	Properties *InventoryProperties `json:"properties,omitempty" azure:"ro"`
}

// PartnerInventoryList - Represents the list of partner inventories
type PartnerInventoryList struct {
	// Link for the next set of partner inventories.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; List of partner inventories
	Value []*PartnerInventory `json:"value,omitempty" azure:"ro"`
}

// SearchInventoriesRequest - Request body for SearchInventories call
type SearchInventoriesRequest struct {
	// REQUIRED; Family identifier for inventory
	FamilyIdentifier *string `json:"familyIdentifier,omitempty"`

	// REQUIRED; Serial number of the inventory
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// SpecificationDetails - Specification details for the inventory
type SpecificationDetails struct {
	// READ-ONLY; Name of the specification property
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Value of the specification property
	Value *string `json:"value,omitempty" azure:"ro"`
}

// StageDetails - Resource stage details.
type StageDetails struct {
	// READ-ONLY; Display name of the resource stage.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; Stage name
	StageName *StageName `json:"stageName,omitempty" azure:"ro"`

	// READ-ONLY; Stage status.
	StageStatus *StageStatus `json:"stageStatus,omitempty" azure:"ro"`

	// READ-ONLY; Stage start time
	StartTime *time.Time `json:"startTime,omitempty" azure:"ro"`
}

// SubscriptionDetails - Contains subscription details
type SubscriptionDetails struct {
	// READ-ONLY; Subscription Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Subscription QuotaId
	QuotaID *string `json:"quotaId,omitempty" azure:"ro"`

	// READ-ONLY; Subscription State
	State *string `json:"state,omitempty" azure:"ro"`
}
