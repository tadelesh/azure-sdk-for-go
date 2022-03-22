//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerlockbox

const (
	moduleName    = "armcustomerlockbox"
	moduleVersion = "v0.4.0"
)

// Status - The status of the request.
type Status string

const (
	StatusApprove      Status = "Approve"
	StatusApproved     Status = "Approved"
	StatusApproving    Status = "Approving"
	StatusCompleted    Status = "Completed"
	StatusCompleting   Status = "Completing"
	StatusDenied       Status = "Denied"
	StatusDeny         Status = "Deny"
	StatusDenying      Status = "Denying"
	StatusError        Status = "Error"
	StatusExpired      Status = "Expired"
	StatusInitializing Status = "Initializing"
	StatusPending      Status = "Pending"
	StatusRevoked      Status = "Revoked"
	StatusRevoking     Status = "Revoking"
	StatusUnknown      Status = "Unknown"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusApprove,
		StatusApproved,
		StatusApproving,
		StatusCompleted,
		StatusCompleting,
		StatusDenied,
		StatusDeny,
		StatusDenying,
		StatusError,
		StatusExpired,
		StatusInitializing,
		StatusPending,
		StatusRevoked,
		StatusRevoking,
		StatusUnknown,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}
