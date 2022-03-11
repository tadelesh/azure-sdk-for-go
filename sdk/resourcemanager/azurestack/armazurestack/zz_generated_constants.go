//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

const (
	moduleName    = "armazurestack"
	moduleVersion = "v0.1.0"
)

// Category - Identity system of the device.
type Category string

const (
	CategoryADFS    Category = "ADFS"
	CategoryAzureAD Category = "AzureAD"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryADFS,
		CategoryAzureAD,
	}
}

// ToPtr returns a *Category pointing to the current value.
func (c Category) ToPtr() *Category {
	return &c
}

// CompatibilityIssue - Compatibility issue
type CompatibilityIssue string

const (
	CompatibilityIssueADFSIdentitySystemRequired      CompatibilityIssue = "ADFSIdentitySystemRequired"
	CompatibilityIssueAzureADIdentitySystemRequired   CompatibilityIssue = "AzureADIdentitySystemRequired"
	CompatibilityIssueCapacityBillingModelRequired    CompatibilityIssue = "CapacityBillingModelRequired"
	CompatibilityIssueConnectionToAzureRequired       CompatibilityIssue = "ConnectionToAzureRequired"
	CompatibilityIssueConnectionToInternetRequired    CompatibilityIssue = "ConnectionToInternetRequired"
	CompatibilityIssueDevelopmentBillingModelRequired CompatibilityIssue = "DevelopmentBillingModelRequired"
	CompatibilityIssueDisconnectedEnvironmentRequired CompatibilityIssue = "DisconnectedEnvironmentRequired"
	CompatibilityIssueHigherDeviceVersionRequired     CompatibilityIssue = "HigherDeviceVersionRequired"
	CompatibilityIssueLowerDeviceVersionRequired      CompatibilityIssue = "LowerDeviceVersionRequired"
	CompatibilityIssuePayAsYouGoBillingModelRequired  CompatibilityIssue = "PayAsYouGoBillingModelRequired"
)

// PossibleCompatibilityIssueValues returns the possible values for the CompatibilityIssue const type.
func PossibleCompatibilityIssueValues() []CompatibilityIssue {
	return []CompatibilityIssue{
		CompatibilityIssueADFSIdentitySystemRequired,
		CompatibilityIssueAzureADIdentitySystemRequired,
		CompatibilityIssueCapacityBillingModelRequired,
		CompatibilityIssueConnectionToAzureRequired,
		CompatibilityIssueConnectionToInternetRequired,
		CompatibilityIssueDevelopmentBillingModelRequired,
		CompatibilityIssueDisconnectedEnvironmentRequired,
		CompatibilityIssueHigherDeviceVersionRequired,
		CompatibilityIssueLowerDeviceVersionRequired,
		CompatibilityIssuePayAsYouGoBillingModelRequired,
	}
}

// ToPtr returns a *CompatibilityIssue pointing to the current value.
func (c CompatibilityIssue) ToPtr() *CompatibilityIssue {
	return &c
}

// ComputeRole - Compute role type (IaaS or PaaS).
type ComputeRole string

const (
	ComputeRoleIaaS ComputeRole = "IaaS"
	ComputeRoleNone ComputeRole = "None"
	ComputeRolePaaS ComputeRole = "PaaS"
)

// PossibleComputeRoleValues returns the possible values for the ComputeRole const type.
func PossibleComputeRoleValues() []ComputeRole {
	return []ComputeRole{
		ComputeRoleIaaS,
		ComputeRoleNone,
		ComputeRolePaaS,
	}
}

// ToPtr returns a *ComputeRole pointing to the current value.
func (c ComputeRole) ToPtr() *ComputeRole {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// Location - Location of the resource.
type Location string

const (
	LocationGlobal Location = "global"
)

// PossibleLocationValues returns the possible values for the Location const type.
func PossibleLocationValues() []Location {
	return []Location{
		LocationGlobal,
	}
}

// ToPtr returns a *Location pointing to the current value.
func (c Location) ToPtr() *Location {
	return &c
}

// OperatingSystem - Operating system type (Windows or Linux).
type OperatingSystem string

const (
	OperatingSystemLinux   OperatingSystem = "Linux"
	OperatingSystemNone    OperatingSystem = "None"
	OperatingSystemWindows OperatingSystem = "Windows"
)

// PossibleOperatingSystemValues returns the possible values for the OperatingSystem const type.
func PossibleOperatingSystemValues() []OperatingSystem {
	return []OperatingSystem{
		OperatingSystemLinux,
		OperatingSystemNone,
		OperatingSystemWindows,
	}
}

// ToPtr returns a *OperatingSystem pointing to the current value.
func (c OperatingSystem) ToPtr() *OperatingSystem {
	return &c
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateCanceled,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}
