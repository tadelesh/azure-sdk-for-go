//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

import "time"

// AzureBareMetalInstance - AzureBareMetal instance info on Azure (ARM properties and AzureBareMetal properties)
type AzureBareMetalInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// AzureBareMetal instance properties
	Properties *AzureBareMetalInstanceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureBareMetalInstanceProperties - Describes the properties of an AzureBareMetal instance.
type AzureBareMetalInstanceProperties struct {
	// Specifies the hardware settings for the AzureBareMetal instance.
	HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`

	// Specifies the network settings for the AzureBareMetal instance.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`

	// Specifies the operating system settings for the AzureBareMetal instance.
	OSProfile *OSProfile `json:"osProfile,omitempty"`

	// ARM ID of another AzureBareMetalInstance that will share a network with this AzureBareMetalInstance
	PartnerNodeID *string `json:"partnerNodeId,omitempty"`

	// Specifies the storage settings for the AzureBareMetal instance disks.
	StorageProfile *StorageProfile `json:"storageProfile,omitempty"`

	// READ-ONLY; Specifies the AzureBareMetal instance unique ID.
	AzureBareMetalInstanceID *string `json:"azureBareMetalInstanceId,omitempty" azure:"ro"`

	// READ-ONLY; Hardware revision of an AzureBareMetal instance
	HwRevision *string `json:"hwRevision,omitempty" azure:"ro"`

	// READ-ONLY; Resource power state
	PowerState *AzureBareMetalInstancePowerStateEnum `json:"powerState,omitempty" azure:"ro"`

	// READ-ONLY; State of provisioning of the AzureBareMetalInstance
	ProvisioningState *AzureBareMetalProvisioningStatesEnum `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource proximity placement group
	ProximityPlacementGroup *string `json:"proximityPlacementGroup,omitempty" azure:"ro"`
}

// AzureBareMetalInstancesClientGetOptions contains the optional parameters for the AzureBareMetalInstancesClient.Get method.
type AzureBareMetalInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalInstancesClient.ListByResourceGroup
// method.
type AzureBareMetalInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalInstancesClient.ListBySubscription
// method.
type AzureBareMetalInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalInstancesClient.Update
// method.
type AzureBareMetalInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesListResult - The response from the List AzureBareMetal Instances operation.
type AzureBareMetalInstancesListResult struct {
	// The URL to get the next set of AzureBareMetal instances.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of Azure BareMetal instances.
	Value []*AzureBareMetalInstance `json:"value,omitempty"`
}

// Disk - Specifies the disk information fo the AzureBareMetal instance
type Disk struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`

	// The disk name.
	Name *string `json:"name,omitempty"`

	// READ-ONLY; Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM
	// and therefore must be unique for each data disk attached to a VM.
	Lun *int32 `json:"lun,omitempty" azure:"ro"`
}

// Display - Detailed BareMetal operation information
type Display struct {
	// READ-ONLY; The localized friendly description for the operation as shown to the user.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name for the operation as shown to the user.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource type related to this action/operation.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// HardwareProfile - Specifies the hardware settings for the AzureBareMetal instance.
type HardwareProfile struct {
	// READ-ONLY; Specifies the AzureBareMetal instance SKU.
	AzureBareMetalInstanceSize *AzureBareMetalInstanceSizeNamesEnum `json:"azureBareMetalInstanceSize,omitempty" azure:"ro"`

	// READ-ONLY; Name of the hardware type (vendor and/or their product name)
	HardwareType *AzureBareMetalHardwareTypeNamesEnum `json:"hardwareType,omitempty" azure:"ro"`
}

// IPAddress - Specifies the IP address of the network interface.
type IPAddress struct {
	// Specifies the IP address of the network interface.
	IPAddress *string `json:"ipAddress,omitempty"`
}

// NetworkProfile - Specifies the network settings for the AzureBareMetal instance disks.
type NetworkProfile struct {
	// Specifies the network interfaces for the AzureBareMetal instance.
	NetworkInterfaces []*IPAddress `json:"networkInterfaces,omitempty"`

	// READ-ONLY; Specifies the circuit id for connecting to express route.
	CircuitID *string `json:"circuitId,omitempty" azure:"ro"`
}

// OSProfile - Specifies the operating system settings for the AzureBareMetal instance.
type OSProfile struct {
	// Specifies the host OS name of the AzureBareMetal instance.
	ComputerName *string `json:"computerName,omitempty"`

	// Specifies the SSH public key used to access the operating system.
	SSHPublicKey *string `json:"sshPublicKey,omitempty"`

	// READ-ONLY; This property allows you to specify the type of the OS.
	OSType *string `json:"osType,omitempty" azure:"ro"`

	// READ-ONLY; Specifies version of operating system.
	Version *string `json:"version,omitempty" azure:"ro"`
}

// Operation - AzureBareMetal operation information
type Operation struct {
	// Displayed AzureBareMetal operation information
	Display *Display `json:"display,omitempty"`

	// READ-ONLY; indicates whether an operation is a data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation being performed on this particular object. This name should match the action name
	// that appears in RBAC / the event service.
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationList - List of AzureBareMetal operations
type OperationList struct {
	// List of AzureBareMetal operations
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Result - Sample result definition
type Result struct {
	// Sample property of type string
	SampleProperty *string `json:"sampleProperty,omitempty"`
}

// StorageProfile - Specifies the storage settings for the AzureBareMetal instance disks.
type StorageProfile struct {
	// Specifies information about the operating system disk used by baremetal instance.
	OSDisks []*Disk `json:"osDisks,omitempty"`

	// READ-ONLY; IP Address to connect to storage.
	NfsIPAddress *string `json:"nfsIpAddress,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// Tags field of the AzureBareMetal instance.
type Tags struct {
	// Tags field of the AzureBareMetal instance.
	Tags map[string]*string `json:"tags,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}
