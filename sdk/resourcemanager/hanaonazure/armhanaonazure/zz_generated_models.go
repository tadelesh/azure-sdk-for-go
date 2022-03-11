//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhanaonazure

// Display - Detailed HANA operation information
type Display struct {
	// READ-ONLY; The localized friendly description for the operation as shown to the user. This description should be thorough,
	// yet concise. It will be used in tool-tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name for the operation as shown to the user. This name should be concise (to fit in drop
	// downs), but clear (self-documenting). Use Title Casing and include the entity/resource
	// to which it applies.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs
	// UX. Default value is 'user,system'
	Origin *string `json:"origin,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name. This form is also expected to include the publisher/company
	// responsible. Use Title Casing. Begin with "Microsoft" for 1st party services.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource type related to this action/operation. This form should match the
	// public documentation for the resource provider. Use Title Casing. For examples, refer to
	// the “name” section.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Describes the error object.
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError - Describes the error object.
type ErrorResponseError struct {
	// READ-ONLY; Error code
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// Operation - HANA operation information
type Operation struct {
	// Displayed HANA operation information
	Display *Display `json:"display,omitempty"`

	// READ-ONLY; The name of the operation being performed on this particular object. This name should match the action name
	// that appears in RBAC / the event service.
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationList - List of HANA operations
type OperationList struct {
	// List of HANA operations
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProviderInstance - A provider instance associated with a SAP monitor.
type ProviderInstance struct {
	// Provider Instance properties
	Properties *ProviderInstanceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ProviderInstanceListResult - The response from the List provider instances operation.
type ProviderInstanceListResult struct {
	// The URL to get the next set of provider instances.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of provider instances.
	Value []*ProviderInstance `json:"value,omitempty"`
}

// ProviderInstanceProperties - Describes the properties of a provider instance.
type ProviderInstanceProperties struct {
	// A JSON string containing metadata of the provider instance.
	Metadata *string `json:"metadata,omitempty"`

	// A JSON string containing the properties of the provider instance.
	Properties *string `json:"properties,omitempty"`

	// The type of provider instance.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; State of provisioning of the provider instance
	ProvisioningState *HanaProvisioningStatesEnum `json:"provisioningState,omitempty" azure:"ro"`
}

// ProviderInstancesClientBeginCreateOptions contains the optional parameters for the ProviderInstancesClient.BeginCreate
// method.
type ProviderInstancesClientBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// ProviderInstancesClientBeginDeleteOptions contains the optional parameters for the ProviderInstancesClient.BeginDelete
// method.
type ProviderInstancesClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ProviderInstancesClientGetOptions contains the optional parameters for the ProviderInstancesClient.Get method.
type ProviderInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProviderInstancesClientListOptions contains the optional parameters for the ProviderInstancesClient.List method.
type ProviderInstancesClientListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
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

// SapMonitor - SAP monitor info on Azure (ARM properties and SAP monitor properties)
type SapMonitor struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// SAP monitor properties
	Properties *SapMonitorProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SapMonitorListResult - The response from the List SAP monitors operation.
type SapMonitorListResult struct {
	// The URL to get the next set of SAP monitors.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of SAP monitors.
	Value []*SapMonitor `json:"value,omitempty"`
}

// SapMonitorProperties - Describes the properties of a SAP monitor.
type SapMonitorProperties struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `json:"enableCustomerAnalytics,omitempty"`

	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmID *string `json:"logAnalyticsWorkspaceArmId,omitempty"`

	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty"`

	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `json:"logAnalyticsWorkspaceSharedKey,omitempty"`

	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `json:"monitorSubnet,omitempty"`

	// READ-ONLY; The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" azure:"ro"`

	// READ-ONLY; State of provisioning of the HanaInstance
	ProvisioningState *HanaProvisioningStatesEnum `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The version of the payload running in the Collector VM
	SapMonitorCollectorVersion *string `json:"sapMonitorCollectorVersion,omitempty" azure:"ro"`
}

// SapMonitorsClientBeginCreateOptions contains the optional parameters for the SapMonitorsClient.BeginCreate method.
type SapMonitorsClientBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// SapMonitorsClientBeginDeleteOptions contains the optional parameters for the SapMonitorsClient.BeginDelete method.
type SapMonitorsClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SapMonitorsClientGetOptions contains the optional parameters for the SapMonitorsClient.Get method.
type SapMonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SapMonitorsClientListOptions contains the optional parameters for the SapMonitorsClient.List method.
type SapMonitorsClientListOptions struct {
	// placeholder for future optional parameters
}

// SapMonitorsClientUpdateOptions contains the optional parameters for the SapMonitorsClient.Update method.
type SapMonitorsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Tags field of the resource.
type Tags struct {
	// Tags field of the resource.
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
