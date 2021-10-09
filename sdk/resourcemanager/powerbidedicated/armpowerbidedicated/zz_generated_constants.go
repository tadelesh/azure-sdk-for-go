//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbidedicated

const (
	module  = "armpowerbidedicated"
	version = "v0.1.0"
)

// CapacityProvisioningState - The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
type CapacityProvisioningState string

const (
	CapacityProvisioningStateDeleting     CapacityProvisioningState = "Deleting"
	CapacityProvisioningStateFailed       CapacityProvisioningState = "Failed"
	CapacityProvisioningStatePaused       CapacityProvisioningState = "Paused"
	CapacityProvisioningStatePausing      CapacityProvisioningState = "Pausing"
	CapacityProvisioningStatePreparing    CapacityProvisioningState = "Preparing"
	CapacityProvisioningStateProvisioning CapacityProvisioningState = "Provisioning"
	CapacityProvisioningStateResuming     CapacityProvisioningState = "Resuming"
	CapacityProvisioningStateScaling      CapacityProvisioningState = "Scaling"
	CapacityProvisioningStateSucceeded    CapacityProvisioningState = "Succeeded"
	CapacityProvisioningStateSuspended    CapacityProvisioningState = "Suspended"
	CapacityProvisioningStateSuspending   CapacityProvisioningState = "Suspending"
	CapacityProvisioningStateUpdating     CapacityProvisioningState = "Updating"
)

// PossibleCapacityProvisioningStateValues returns the possible values for the CapacityProvisioningState const type.
func PossibleCapacityProvisioningStateValues() []CapacityProvisioningState {
	return []CapacityProvisioningState{
		CapacityProvisioningStateDeleting,
		CapacityProvisioningStateFailed,
		CapacityProvisioningStatePaused,
		CapacityProvisioningStatePausing,
		CapacityProvisioningStatePreparing,
		CapacityProvisioningStateProvisioning,
		CapacityProvisioningStateResuming,
		CapacityProvisioningStateScaling,
		CapacityProvisioningStateSucceeded,
		CapacityProvisioningStateSuspended,
		CapacityProvisioningStateSuspending,
		CapacityProvisioningStateUpdating,
	}
}

// ToPtr returns a *CapacityProvisioningState pointing to the current value.
func (c CapacityProvisioningState) ToPtr() *CapacityProvisioningState {
	return &c
}

// CapacitySKUTier - The name of the Azure pricing tier to which the SKU applies.
type CapacitySKUTier string

const (
	CapacitySKUTierAutoPremiumHost CapacitySKUTier = "AutoPremiumHost"
	CapacitySKUTierPBIEAzure       CapacitySKUTier = "PBIE_Azure"
	CapacitySKUTierPremium         CapacitySKUTier = "Premium"
)

// PossibleCapacitySKUTierValues returns the possible values for the CapacitySKUTier const type.
func PossibleCapacitySKUTierValues() []CapacitySKUTier {
	return []CapacitySKUTier{
		CapacitySKUTierAutoPremiumHost,
		CapacitySKUTierPBIEAzure,
		CapacitySKUTierPremium,
	}
}

// ToPtr returns a *CapacitySKUTier pointing to the current value.
func (c CapacitySKUTier) ToPtr() *CapacitySKUTier {
	return &c
}

// IdentityType - The type of identity that created/modified the resource.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// Mode - The capacity mode.
type Mode string

const (
	ModeGen1 Mode = "Gen1"
	ModeGen2 Mode = "Gen2"
)

// PossibleModeValues returns the possible values for the Mode const type.
func PossibleModeValues() []Mode {
	return []Mode{
		ModeGen1,
		ModeGen2,
	}
}

// ToPtr returns a *Mode pointing to the current value.
func (c Mode) ToPtr() *Mode {
	return &c
}

// State - The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
type State string

const (
	StateDeleting     State = "Deleting"
	StateFailed       State = "Failed"
	StatePaused       State = "Paused"
	StatePausing      State = "Pausing"
	StatePreparing    State = "Preparing"
	StateProvisioning State = "Provisioning"
	StateResuming     State = "Resuming"
	StateScaling      State = "Scaling"
	StateSucceeded    State = "Succeeded"
	StateSuspended    State = "Suspended"
	StateSuspending   State = "Suspending"
	StateUpdating     State = "Updating"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDeleting,
		StateFailed,
		StatePaused,
		StatePausing,
		StatePreparing,
		StateProvisioning,
		StateResuming,
		StateScaling,
		StateSucceeded,
		StateSuspended,
		StateSuspending,
		StateUpdating,
	}
}

// ToPtr returns a *State pointing to the current value.
func (c State) ToPtr() *State {
	return &c
}

// VCoreProvisioningState - The current deployment state of an auto scale v-core resource. The provisioningState is to indicate states for resource provisioning.
type VCoreProvisioningState string

const (
	VCoreProvisioningStateSucceeded VCoreProvisioningState = "Succeeded"
)

// PossibleVCoreProvisioningStateValues returns the possible values for the VCoreProvisioningState const type.
func PossibleVCoreProvisioningStateValues() []VCoreProvisioningState {
	return []VCoreProvisioningState{
		VCoreProvisioningStateSucceeded,
	}
}

// ToPtr returns a *VCoreProvisioningState pointing to the current value.
func (c VCoreProvisioningState) ToPtr() *VCoreProvisioningState {
	return &c
}

// VCoreSKUTier - The name of the Azure pricing tier to which the SKU applies.
type VCoreSKUTier string

const (
	VCoreSKUTierAutoScale VCoreSKUTier = "AutoScale"
)

// PossibleVCoreSKUTierValues returns the possible values for the VCoreSKUTier const type.
func PossibleVCoreSKUTierValues() []VCoreSKUTier {
	return []VCoreSKUTier{
		VCoreSKUTierAutoScale,
	}
}

// ToPtr returns a *VCoreSKUTier pointing to the current value.
func (c VCoreSKUTier) ToPtr() *VCoreSKUTier {
	return &c
}
