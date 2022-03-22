//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

const (
	moduleName    = "armconfluent"
	moduleVersion = "v0.4.0"
)

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

// ProvisionState - Provision states for confluent RP
type ProvisionState string

const (
	ProvisionStateAccepted     ProvisionState = "Accepted"
	ProvisionStateCanceled     ProvisionState = "Canceled"
	ProvisionStateCreating     ProvisionState = "Creating"
	ProvisionStateDeleted      ProvisionState = "Deleted"
	ProvisionStateDeleting     ProvisionState = "Deleting"
	ProvisionStateFailed       ProvisionState = "Failed"
	ProvisionStateNotSpecified ProvisionState = "NotSpecified"
	ProvisionStateSucceeded    ProvisionState = "Succeeded"
	ProvisionStateUpdating     ProvisionState = "Updating"
)

// PossibleProvisionStateValues returns the possible values for the ProvisionState const type.
func PossibleProvisionStateValues() []ProvisionState {
	return []ProvisionState{
		ProvisionStateAccepted,
		ProvisionStateCanceled,
		ProvisionStateCreating,
		ProvisionStateDeleted,
		ProvisionStateDeleting,
		ProvisionStateFailed,
		ProvisionStateNotSpecified,
		ProvisionStateSucceeded,
		ProvisionStateUpdating,
	}
}

// ToPtr returns a *ProvisionState pointing to the current value.
func (c ProvisionState) ToPtr() *ProvisionState {
	return &c
}

// SaaSOfferStatus - SaaS Offer Status for confluent RP
type SaaSOfferStatus string

const (
	SaaSOfferStatusFailed                  SaaSOfferStatus = "Failed"
	SaaSOfferStatusInProgress              SaaSOfferStatus = "InProgress"
	SaaSOfferStatusPendingFulfillmentStart SaaSOfferStatus = "PendingFulfillmentStart"
	SaaSOfferStatusReinstated              SaaSOfferStatus = "Reinstated"
	SaaSOfferStatusStarted                 SaaSOfferStatus = "Started"
	SaaSOfferStatusSubscribed              SaaSOfferStatus = "Subscribed"
	SaaSOfferStatusSucceeded               SaaSOfferStatus = "Succeeded"
	SaaSOfferStatusSuspended               SaaSOfferStatus = "Suspended"
	SaaSOfferStatusUnsubscribed            SaaSOfferStatus = "Unsubscribed"
	SaaSOfferStatusUpdating                SaaSOfferStatus = "Updating"
)

// PossibleSaaSOfferStatusValues returns the possible values for the SaaSOfferStatus const type.
func PossibleSaaSOfferStatusValues() []SaaSOfferStatus {
	return []SaaSOfferStatus{
		SaaSOfferStatusFailed,
		SaaSOfferStatusInProgress,
		SaaSOfferStatusPendingFulfillmentStart,
		SaaSOfferStatusReinstated,
		SaaSOfferStatusStarted,
		SaaSOfferStatusSubscribed,
		SaaSOfferStatusSucceeded,
		SaaSOfferStatusSuspended,
		SaaSOfferStatusUnsubscribed,
		SaaSOfferStatusUpdating,
	}
}

// ToPtr returns a *SaaSOfferStatus pointing to the current value.
func (c SaaSOfferStatus) ToPtr() *SaaSOfferStatus {
	return &c
}
