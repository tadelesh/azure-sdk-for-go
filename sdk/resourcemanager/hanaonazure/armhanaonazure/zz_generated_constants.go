//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhanaonazure

const (
	moduleName    = "armhanaonazure"
	moduleVersion = "v0.4.0"
)

// HanaProvisioningStatesEnum - State of provisioning of the HanaInstance
type HanaProvisioningStatesEnum string

const (
	HanaProvisioningStatesEnumAccepted  HanaProvisioningStatesEnum = "Accepted"
	HanaProvisioningStatesEnumCreating  HanaProvisioningStatesEnum = "Creating"
	HanaProvisioningStatesEnumDeleting  HanaProvisioningStatesEnum = "Deleting"
	HanaProvisioningStatesEnumFailed    HanaProvisioningStatesEnum = "Failed"
	HanaProvisioningStatesEnumMigrating HanaProvisioningStatesEnum = "Migrating"
	HanaProvisioningStatesEnumSucceeded HanaProvisioningStatesEnum = "Succeeded"
	HanaProvisioningStatesEnumUpdating  HanaProvisioningStatesEnum = "Updating"
)

// PossibleHanaProvisioningStatesEnumValues returns the possible values for the HanaProvisioningStatesEnum const type.
func PossibleHanaProvisioningStatesEnumValues() []HanaProvisioningStatesEnum {
	return []HanaProvisioningStatesEnum{
		HanaProvisioningStatesEnumAccepted,
		HanaProvisioningStatesEnumCreating,
		HanaProvisioningStatesEnumDeleting,
		HanaProvisioningStatesEnumFailed,
		HanaProvisioningStatesEnumMigrating,
		HanaProvisioningStatesEnumSucceeded,
		HanaProvisioningStatesEnumUpdating,
	}
}

// ToPtr returns a *HanaProvisioningStatesEnum pointing to the current value.
func (c HanaProvisioningStatesEnum) ToPtr() *HanaProvisioningStatesEnum {
	return &c
}
