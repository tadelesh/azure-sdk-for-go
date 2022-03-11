//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadmonitor

const (
	moduleName    = "armworkloadmonitor"
	moduleVersion = "v0.3.0"
)

// HealthState - One of health states - healthy, critical, warning, unknown, none, disabled.
type HealthState string

const (
	HealthStateCritical HealthState = "Critical"
	HealthStateDisabled HealthState = "Disabled"
	HealthStateHealthy  HealthState = "Healthy"
	HealthStateNone     HealthState = "None"
	HealthStateUnknown  HealthState = "Unknown"
	HealthStateWarning  HealthState = "Warning"
)

// PossibleHealthStateValues returns the possible values for the HealthState const type.
func PossibleHealthStateValues() []HealthState {
	return []HealthState{
		HealthStateCritical,
		HealthStateDisabled,
		HealthStateHealthy,
		HealthStateNone,
		HealthStateUnknown,
		HealthStateWarning,
	}
}

// ToPtr returns a *HealthState pointing to the current value.
func (c HealthState) ToPtr() *HealthState {
	return &c
}
