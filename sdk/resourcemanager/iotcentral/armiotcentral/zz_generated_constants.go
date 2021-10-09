//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotcentral

const (
	module  = "armiotcentral"
	version = "v0.1.0"
)

// AppSKU - The name of the SKU.
type AppSKU string

const (
	AppSKUST0 AppSKU = "ST0"
	AppSKUST1 AppSKU = "ST1"
	AppSKUST2 AppSKU = "ST2"
)

// PossibleAppSKUValues returns the possible values for the AppSKU const type.
func PossibleAppSKUValues() []AppSKU {
	return []AppSKU{
		AppSKUST0,
		AppSKUST1,
		AppSKUST2,
	}
}

// ToPtr returns a *AppSKU pointing to the current value.
func (c AppSKU) ToPtr() *AppSKU {
	return &c
}

// AppState - The current state of the application.
type AppState string

const (
	AppStateCreated   AppState = "created"
	AppStateSuspended AppState = "suspended"
)

// PossibleAppStateValues returns the possible values for the AppState const type.
func PossibleAppStateValues() []AppState {
	return []AppState{
		AppStateCreated,
		AppStateSuspended,
	}
}

// ToPtr returns a *AppState pointing to the current value.
func (c AppState) ToPtr() *AppState {
	return &c
}
