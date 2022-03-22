//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnotificationhubs

const (
	moduleName    = "armnotificationhubs"
	moduleVersion = "v0.5.0"
)

type AccessRights string

const (
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
	AccessRightsListen AccessRights = "Listen"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsManage,
		AccessRightsSend,
		AccessRightsListen,
	}
}

// ToPtr returns a *AccessRights pointing to the current value.
func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// NamespaceType - The namespace type.
type NamespaceType string

const (
	NamespaceTypeMessaging       NamespaceType = "Messaging"
	NamespaceTypeNotificationHub NamespaceType = "NotificationHub"
)

// PossibleNamespaceTypeValues returns the possible values for the NamespaceType const type.
func PossibleNamespaceTypeValues() []NamespaceType {
	return []NamespaceType{
		NamespaceTypeMessaging,
		NamespaceTypeNotificationHub,
	}
}

// ToPtr returns a *NamespaceType pointing to the current value.
func (c NamespaceType) ToPtr() *NamespaceType {
	return &c
}

// SKUName - Name of the notification hub sku
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameFree     SKUName = "Free"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameFree,
		SKUNameStandard,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}
