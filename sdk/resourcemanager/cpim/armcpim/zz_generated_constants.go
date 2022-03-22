//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcpim

const (
	moduleName    = "armcpim"
	moduleVersion = "v0.2.0"
)

// B2CResourceSKUName - The name of the SKU for the tenant.
type B2CResourceSKUName string

const (
	// B2CResourceSKUNameStandard - Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users
	// (MAU) billing model.
	B2CResourceSKUNameStandard B2CResourceSKUName = "Standard"
	// B2CResourceSKUNamePremiumP1 - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	B2CResourceSKUNamePremiumP1 B2CResourceSKUName = "PremiumP1"
	// B2CResourceSKUNamePremiumP2 - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	B2CResourceSKUNamePremiumP2 B2CResourceSKUName = "PremiumP2"
)

// PossibleB2CResourceSKUNameValues returns the possible values for the B2CResourceSKUName const type.
func PossibleB2CResourceSKUNameValues() []B2CResourceSKUName {
	return []B2CResourceSKUName{
		B2CResourceSKUNameStandard,
		B2CResourceSKUNamePremiumP1,
		B2CResourceSKUNamePremiumP2,
	}
}

// ToPtr returns a *B2CResourceSKUName pointing to the current value.
func (c B2CResourceSKUName) ToPtr() *B2CResourceSKUName {
	return &c
}

// BillingType - The type of billing. Will be MAU for all new customers. If 'Auths', it can be updated to 'MAU'. Cannot be
// changed if value is 'MAU'. Learn more about Azure AD B2C billing at aka.ms/b2cBilling
// [https://aka.ms/b2cbilling].
type BillingType string

const (
	// BillingTypeMAU - Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users (MAU) billing
	// model.
	BillingTypeMAU BillingType = "MAU"
	// BillingTypeAuths - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based
	// billing.
	BillingTypeAuths BillingType = "Auths"
)

// PossibleBillingTypeValues returns the possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{
		BillingTypeMAU,
		BillingTypeAuths,
	}
}

// ToPtr returns a *BillingType pointing to the current value.
func (c BillingType) ToPtr() *BillingType {
	return &c
}

// NameAvailabilityReasonType - Describes the reason for the 'nameAvailable' value.
type NameAvailabilityReasonType string

const (
	// NameAvailabilityReasonTypeAlreadyExists - The name is already in use and is therefore unavailable.
	NameAvailabilityReasonTypeAlreadyExists NameAvailabilityReasonType = "AlreadyExists"
	// NameAvailabilityReasonTypeInvalid - The name provided does not match the resource provider’s naming requirements (incorrect
	// length, unsupported characters, etc.).
	NameAvailabilityReasonTypeInvalid NameAvailabilityReasonType = "Invalid"
)

// PossibleNameAvailabilityReasonTypeValues returns the possible values for the NameAvailabilityReasonType const type.
func PossibleNameAvailabilityReasonTypeValues() []NameAvailabilityReasonType {
	return []NameAvailabilityReasonType{
		NameAvailabilityReasonTypeAlreadyExists,
		NameAvailabilityReasonTypeInvalid,
	}
}

// ToPtr returns a *NameAvailabilityReasonType pointing to the current value.
func (c NameAvailabilityReasonType) ToPtr() *NameAvailabilityReasonType {
	return &c
}

// StatusType - The status of the operation.
type StatusType string

const (
	// StatusTypeSucceeded - The operation succeeded.
	StatusTypeSucceeded StatusType = "Succeeded"
	// StatusTypePending - The operation is pending.
	StatusTypePending StatusType = "Pending"
	// StatusTypeFailed - The operation failed.
	StatusTypeFailed StatusType = "Failed"
)

// PossibleStatusTypeValues returns the possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{
		StatusTypeSucceeded,
		StatusTypePending,
		StatusTypeFailed,
	}
}

// ToPtr returns a *StatusType pointing to the current value.
func (c StatusType) ToPtr() *StatusType {
	return &c
}
