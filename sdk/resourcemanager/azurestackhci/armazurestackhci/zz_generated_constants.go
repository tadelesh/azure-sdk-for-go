//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

const (
	moduleName    = "armazurestackhci"
	moduleVersion = "v0.3.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// ArcSettingAggregateState - Aggregate state of Arc agent across the nodes in this HCI cluster.
type ArcSettingAggregateState string

const (
	ArcSettingAggregateStateCanceled           ArcSettingAggregateState = "Canceled"
	ArcSettingAggregateStateConnected          ArcSettingAggregateState = "Connected"
	ArcSettingAggregateStateCreating           ArcSettingAggregateState = "Creating"
	ArcSettingAggregateStateDeleted            ArcSettingAggregateState = "Deleted"
	ArcSettingAggregateStateDeleting           ArcSettingAggregateState = "Deleting"
	ArcSettingAggregateStateDisconnected       ArcSettingAggregateState = "Disconnected"
	ArcSettingAggregateStateError              ArcSettingAggregateState = "Error"
	ArcSettingAggregateStateFailed             ArcSettingAggregateState = "Failed"
	ArcSettingAggregateStateInProgress         ArcSettingAggregateState = "InProgress"
	ArcSettingAggregateStateMoving             ArcSettingAggregateState = "Moving"
	ArcSettingAggregateStateNotSpecified       ArcSettingAggregateState = "NotSpecified"
	ArcSettingAggregateStatePartiallyConnected ArcSettingAggregateState = "PartiallyConnected"
	ArcSettingAggregateStatePartiallySucceeded ArcSettingAggregateState = "PartiallySucceeded"
	ArcSettingAggregateStateSucceeded          ArcSettingAggregateState = "Succeeded"
	ArcSettingAggregateStateUpdating           ArcSettingAggregateState = "Updating"
)

// PossibleArcSettingAggregateStateValues returns the possible values for the ArcSettingAggregateState const type.
func PossibleArcSettingAggregateStateValues() []ArcSettingAggregateState {
	return []ArcSettingAggregateState{
		ArcSettingAggregateStateCanceled,
		ArcSettingAggregateStateConnected,
		ArcSettingAggregateStateCreating,
		ArcSettingAggregateStateDeleted,
		ArcSettingAggregateStateDeleting,
		ArcSettingAggregateStateDisconnected,
		ArcSettingAggregateStateError,
		ArcSettingAggregateStateFailed,
		ArcSettingAggregateStateInProgress,
		ArcSettingAggregateStateMoving,
		ArcSettingAggregateStateNotSpecified,
		ArcSettingAggregateStatePartiallyConnected,
		ArcSettingAggregateStatePartiallySucceeded,
		ArcSettingAggregateStateSucceeded,
		ArcSettingAggregateStateUpdating,
	}
}

// ToPtr returns a *ArcSettingAggregateState pointing to the current value.
func (c ArcSettingAggregateState) ToPtr() *ArcSettingAggregateState {
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

// DiagnosticLevel - Desired level of diagnostic data emitted by the cluster.
type DiagnosticLevel string

const (
	DiagnosticLevelBasic    DiagnosticLevel = "Basic"
	DiagnosticLevelEnhanced DiagnosticLevel = "Enhanced"
	DiagnosticLevelOff      DiagnosticLevel = "Off"
)

// PossibleDiagnosticLevelValues returns the possible values for the DiagnosticLevel const type.
func PossibleDiagnosticLevelValues() []DiagnosticLevel {
	return []DiagnosticLevel{
		DiagnosticLevelBasic,
		DiagnosticLevelEnhanced,
		DiagnosticLevelOff,
	}
}

// ToPtr returns a *DiagnosticLevel pointing to the current value.
func (c DiagnosticLevel) ToPtr() *DiagnosticLevel {
	return &c
}

// ExtensionAggregateState - Aggregate state of Arc Extensions across the nodes in this HCI cluster.
type ExtensionAggregateState string

const (
	ExtensionAggregateStateCanceled           ExtensionAggregateState = "Canceled"
	ExtensionAggregateStateConnected          ExtensionAggregateState = "Connected"
	ExtensionAggregateStateCreating           ExtensionAggregateState = "Creating"
	ExtensionAggregateStateDeleted            ExtensionAggregateState = "Deleted"
	ExtensionAggregateStateDeleting           ExtensionAggregateState = "Deleting"
	ExtensionAggregateStateDisconnected       ExtensionAggregateState = "Disconnected"
	ExtensionAggregateStateError              ExtensionAggregateState = "Error"
	ExtensionAggregateStateFailed             ExtensionAggregateState = "Failed"
	ExtensionAggregateStateInProgress         ExtensionAggregateState = "InProgress"
	ExtensionAggregateStateMoving             ExtensionAggregateState = "Moving"
	ExtensionAggregateStateNotSpecified       ExtensionAggregateState = "NotSpecified"
	ExtensionAggregateStatePartiallyConnected ExtensionAggregateState = "PartiallyConnected"
	ExtensionAggregateStatePartiallySucceeded ExtensionAggregateState = "PartiallySucceeded"
	ExtensionAggregateStateSucceeded          ExtensionAggregateState = "Succeeded"
	ExtensionAggregateStateUpdating           ExtensionAggregateState = "Updating"
)

// PossibleExtensionAggregateStateValues returns the possible values for the ExtensionAggregateState const type.
func PossibleExtensionAggregateStateValues() []ExtensionAggregateState {
	return []ExtensionAggregateState{
		ExtensionAggregateStateCanceled,
		ExtensionAggregateStateConnected,
		ExtensionAggregateStateCreating,
		ExtensionAggregateStateDeleted,
		ExtensionAggregateStateDeleting,
		ExtensionAggregateStateDisconnected,
		ExtensionAggregateStateError,
		ExtensionAggregateStateFailed,
		ExtensionAggregateStateInProgress,
		ExtensionAggregateStateMoving,
		ExtensionAggregateStateNotSpecified,
		ExtensionAggregateStatePartiallyConnected,
		ExtensionAggregateStatePartiallySucceeded,
		ExtensionAggregateStateSucceeded,
		ExtensionAggregateStateUpdating,
	}
}

// ToPtr returns a *ExtensionAggregateState pointing to the current value.
func (c ExtensionAggregateState) ToPtr() *ExtensionAggregateState {
	return &c
}

// ImdsAttestation - IMDS attestation status of the cluster.
type ImdsAttestation string

const (
	ImdsAttestationDisabled ImdsAttestation = "Disabled"
	ImdsAttestationEnabled  ImdsAttestation = "Enabled"
)

// PossibleImdsAttestationValues returns the possible values for the ImdsAttestation const type.
func PossibleImdsAttestationValues() []ImdsAttestation {
	return []ImdsAttestation{
		ImdsAttestationDisabled,
		ImdsAttestationEnabled,
	}
}

// ToPtr returns a *ImdsAttestation pointing to the current value.
func (c ImdsAttestation) ToPtr() *ImdsAttestation {
	return &c
}

// NodeArcState - State of Arc agent in this node.
type NodeArcState string

const (
	NodeArcStateCanceled     NodeArcState = "Canceled"
	NodeArcStateConnected    NodeArcState = "Connected"
	NodeArcStateCreating     NodeArcState = "Creating"
	NodeArcStateDeleted      NodeArcState = "Deleted"
	NodeArcStateDeleting     NodeArcState = "Deleting"
	NodeArcStateDisconnected NodeArcState = "Disconnected"
	NodeArcStateError        NodeArcState = "Error"
	NodeArcStateFailed       NodeArcState = "Failed"
	NodeArcStateMoving       NodeArcState = "Moving"
	NodeArcStateNotSpecified NodeArcState = "NotSpecified"
	NodeArcStateSucceeded    NodeArcState = "Succeeded"
	NodeArcStateUpdating     NodeArcState = "Updating"
)

// PossibleNodeArcStateValues returns the possible values for the NodeArcState const type.
func PossibleNodeArcStateValues() []NodeArcState {
	return []NodeArcState{
		NodeArcStateCanceled,
		NodeArcStateConnected,
		NodeArcStateCreating,
		NodeArcStateDeleted,
		NodeArcStateDeleting,
		NodeArcStateDisconnected,
		NodeArcStateError,
		NodeArcStateFailed,
		NodeArcStateMoving,
		NodeArcStateNotSpecified,
		NodeArcStateSucceeded,
		NodeArcStateUpdating,
	}
}

// ToPtr returns a *NodeArcState pointing to the current value.
func (c NodeArcState) ToPtr() *NodeArcState {
	return &c
}

// NodeExtensionState - State of Arc Extension in this node.
type NodeExtensionState string

const (
	NodeExtensionStateCanceled     NodeExtensionState = "Canceled"
	NodeExtensionStateConnected    NodeExtensionState = "Connected"
	NodeExtensionStateCreating     NodeExtensionState = "Creating"
	NodeExtensionStateDeleted      NodeExtensionState = "Deleted"
	NodeExtensionStateDeleting     NodeExtensionState = "Deleting"
	NodeExtensionStateDisconnected NodeExtensionState = "Disconnected"
	NodeExtensionStateError        NodeExtensionState = "Error"
	NodeExtensionStateFailed       NodeExtensionState = "Failed"
	NodeExtensionStateMoving       NodeExtensionState = "Moving"
	NodeExtensionStateNotSpecified NodeExtensionState = "NotSpecified"
	NodeExtensionStateSucceeded    NodeExtensionState = "Succeeded"
	NodeExtensionStateUpdating     NodeExtensionState = "Updating"
)

// PossibleNodeExtensionStateValues returns the possible values for the NodeExtensionState const type.
func PossibleNodeExtensionStateValues() []NodeExtensionState {
	return []NodeExtensionState{
		NodeExtensionStateCanceled,
		NodeExtensionStateConnected,
		NodeExtensionStateCreating,
		NodeExtensionStateDeleted,
		NodeExtensionStateDeleting,
		NodeExtensionStateDisconnected,
		NodeExtensionStateError,
		NodeExtensionStateFailed,
		NodeExtensionStateMoving,
		NodeExtensionStateNotSpecified,
		NodeExtensionStateSucceeded,
		NodeExtensionStateUpdating,
	}
}

// ToPtr returns a *NodeExtensionState pointing to the current value.
func (c NodeExtensionState) ToPtr() *NodeExtensionState {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// ProvisioningState - Provisioning state of the ArcSetting proxy resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// Status - Status of the cluster agent.
type Status string

const (
	StatusConnectedRecently    Status = "ConnectedRecently"
	StatusDisconnected         Status = "Disconnected"
	StatusError                Status = "Error"
	StatusNotConnectedRecently Status = "NotConnectedRecently"
	StatusNotYetRegistered     Status = "NotYetRegistered"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusConnectedRecently,
		StatusDisconnected,
		StatusError,
		StatusNotConnectedRecently,
		StatusNotYetRegistered,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}

// WindowsServerSubscription - Desired state of Windows Server Subscription.
type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled WindowsServerSubscription = "Disabled"
	WindowsServerSubscriptionEnabled  WindowsServerSubscription = "Enabled"
)

// PossibleWindowsServerSubscriptionValues returns the possible values for the WindowsServerSubscription const type.
func PossibleWindowsServerSubscriptionValues() []WindowsServerSubscription {
	return []WindowsServerSubscription{
		WindowsServerSubscriptionDisabled,
		WindowsServerSubscriptionEnabled,
	}
}

// ToPtr returns a *WindowsServerSubscription pointing to the current value.
func (c WindowsServerSubscription) ToPtr() *WindowsServerSubscription {
	return &c
}
