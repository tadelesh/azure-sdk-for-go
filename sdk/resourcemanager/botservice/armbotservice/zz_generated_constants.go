//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

const (
	moduleName    = "armbotservice"
	moduleVersion = "v0.3.0"
)

type ChannelName string

const (
	ChannelNameAlexaChannel            ChannelName = "AlexaChannel"
	ChannelNameFacebookChannel         ChannelName = "FacebookChannel"
	ChannelNameEmailChannel            ChannelName = "EmailChannel"
	ChannelNameKikChannel              ChannelName = "KikChannel"
	ChannelNameTelegramChannel         ChannelName = "TelegramChannel"
	ChannelNameSlackChannel            ChannelName = "SlackChannel"
	ChannelNameMsTeamsChannel          ChannelName = "MsTeamsChannel"
	ChannelNameSkypeChannel            ChannelName = "SkypeChannel"
	ChannelNameWebChatChannel          ChannelName = "WebChatChannel"
	ChannelNameDirectLineChannel       ChannelName = "DirectLineChannel"
	ChannelNameSmsChannel              ChannelName = "SmsChannel"
	ChannelNameLineChannel             ChannelName = "LineChannel"
	ChannelNameDirectLineSpeechChannel ChannelName = "DirectLineSpeechChannel"
)

// PossibleChannelNameValues returns the possible values for the ChannelName const type.
func PossibleChannelNameValues() []ChannelName {
	return []ChannelName{
		ChannelNameAlexaChannel,
		ChannelNameFacebookChannel,
		ChannelNameEmailChannel,
		ChannelNameKikChannel,
		ChannelNameTelegramChannel,
		ChannelNameSlackChannel,
		ChannelNameMsTeamsChannel,
		ChannelNameSkypeChannel,
		ChannelNameWebChatChannel,
		ChannelNameDirectLineChannel,
		ChannelNameSmsChannel,
		ChannelNameLineChannel,
		ChannelNameDirectLineSpeechChannel,
	}
}

// ToPtr returns a *ChannelName pointing to the current value.
func (c ChannelName) ToPtr() *ChannelName {
	return &c
}

// Key - Determines which key is to be regenerated
type Key string

const (
	KeyKey1 Key = "key1"
	KeyKey2 Key = "key2"
)

// PossibleKeyValues returns the possible values for the Key const type.
func PossibleKeyValues() []Key {
	return []Key{
		KeyKey1,
		KeyKey2,
	}
}

// ToPtr returns a *Key pointing to the current value.
func (c Key) ToPtr() *Key {
	return &c
}

// Kind - Indicates the type of bot service
type Kind string

const (
	KindAzurebot Kind = "azurebot"
	KindBot      Kind = "bot"
	KindDesigner Kind = "designer"
	KindFunction Kind = "function"
	KindSdk      Kind = "sdk"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindAzurebot,
		KindBot,
		KindDesigner,
		KindFunction,
		KindSdk,
	}
}

// ToPtr returns a *Kind pointing to the current value.
func (c Kind) ToPtr() *Kind {
	return &c
}

// MsaAppType - Microsoft App Type for the bot
type MsaAppType string

const (
	MsaAppTypeMultiTenant     MsaAppType = "MultiTenant"
	MsaAppTypeSingleTenant    MsaAppType = "SingleTenant"
	MsaAppTypeUserAssignedMSI MsaAppType = "UserAssignedMSI"
)

// PossibleMsaAppTypeValues returns the possible values for the MsaAppType const type.
func PossibleMsaAppTypeValues() []MsaAppType {
	return []MsaAppType{
		MsaAppTypeMultiTenant,
		MsaAppTypeSingleTenant,
		MsaAppTypeUserAssignedMSI,
	}
}

// ToPtr returns a *MsaAppType pointing to the current value.
func (c MsaAppType) ToPtr() *MsaAppType {
	return &c
}

// OperationResultStatus - The status of the operation being performed.
type OperationResultStatus string

const (
	OperationResultStatusCanceled  OperationResultStatus = "Canceled"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusRequested OperationResultStatus = "Requested"
	OperationResultStatusRunning   OperationResultStatus = "Running"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
)

// PossibleOperationResultStatusValues returns the possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{
		OperationResultStatusCanceled,
		OperationResultStatusFailed,
		OperationResultStatusRequested,
		OperationResultStatusRunning,
		OperationResultStatusSucceeded,
	}
}

// ToPtr returns a *OperationResultStatus pointing to the current value.
func (c OperationResultStatus) ToPtr() *OperationResultStatus {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// ToPtr returns a *PrivateEndpointConnectionProvisioningState pointing to the current value.
func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateEndpointServiceConnectionStatus pointing to the current value.
func (c PrivateEndpointServiceConnectionStatus) ToPtr() *PrivateEndpointServiceConnectionStatus {
	return &c
}

// PublicNetworkAccess - Whether the bot is in an isolated network
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

type RegenerateKeysChannelName string

const (
	RegenerateKeysChannelNameWebChatChannel    RegenerateKeysChannelName = "WebChatChannel"
	RegenerateKeysChannelNameDirectLineChannel RegenerateKeysChannelName = "DirectLineChannel"
)

// PossibleRegenerateKeysChannelNameValues returns the possible values for the RegenerateKeysChannelName const type.
func PossibleRegenerateKeysChannelNameValues() []RegenerateKeysChannelName {
	return []RegenerateKeysChannelName{
		RegenerateKeysChannelNameWebChatChannel,
		RegenerateKeysChannelNameDirectLineChannel,
	}
}

// ToPtr returns a *RegenerateKeysChannelName pointing to the current value.
func (c RegenerateKeysChannelName) ToPtr() *RegenerateKeysChannelName {
	return &c
}

// SKUName - The name of SKU.
type SKUName string

const (
	SKUNameF0 SKUName = "F0"
	SKUNameS1 SKUName = "S1"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameF0,
		SKUNameS1,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - Gets the sku tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierStandard,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}
