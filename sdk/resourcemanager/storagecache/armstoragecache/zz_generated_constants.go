//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

const (
	moduleName    = "armstoragecache"
	moduleVersion = "v0.4.0"
)

// CacheIdentityType - The type of identity used for the cache
type CacheIdentityType string

const (
	CacheIdentityTypeSystemAssigned             CacheIdentityType = "SystemAssigned"
	CacheIdentityTypeUserAssigned               CacheIdentityType = "UserAssigned"
	CacheIdentityTypeSystemAssignedUserAssigned CacheIdentityType = "SystemAssigned, UserAssigned"
	CacheIdentityTypeNone                       CacheIdentityType = "None"
)

// PossibleCacheIdentityTypeValues returns the possible values for the CacheIdentityType const type.
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return []CacheIdentityType{
		CacheIdentityTypeSystemAssigned,
		CacheIdentityTypeUserAssigned,
		CacheIdentityTypeSystemAssignedUserAssigned,
		CacheIdentityTypeNone,
	}
}

// ToPtr returns a *CacheIdentityType pointing to the current value.
func (c CacheIdentityType) ToPtr() *CacheIdentityType {
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

// DomainJoinedType - True if the HPC Cache is joined to the Active Directory domain.
type DomainJoinedType string

const (
	DomainJoinedTypeError DomainJoinedType = "Error"
	DomainJoinedTypeNo    DomainJoinedType = "No"
	DomainJoinedTypeYes   DomainJoinedType = "Yes"
)

// PossibleDomainJoinedTypeValues returns the possible values for the DomainJoinedType const type.
func PossibleDomainJoinedTypeValues() []DomainJoinedType {
	return []DomainJoinedType{
		DomainJoinedTypeError,
		DomainJoinedTypeNo,
		DomainJoinedTypeYes,
	}
}

// ToPtr returns a *DomainJoinedType pointing to the current value.
func (c DomainJoinedType) ToPtr() *DomainJoinedType {
	return &c
}

// FirmwareStatusType - True if there is a firmware update ready to install on this Cache. The firmware will automatically
// be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
type FirmwareStatusType string

const (
	FirmwareStatusTypeAvailable   FirmwareStatusType = "available"
	FirmwareStatusTypeUnavailable FirmwareStatusType = "unavailable"
)

// PossibleFirmwareStatusTypeValues returns the possible values for the FirmwareStatusType const type.
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return []FirmwareStatusType{
		FirmwareStatusTypeAvailable,
		FirmwareStatusTypeUnavailable,
	}
}

// ToPtr returns a *FirmwareStatusType pointing to the current value.
func (c FirmwareStatusType) ToPtr() *FirmwareStatusType {
	return &c
}

// HealthStateType - List of Cache health states.
type HealthStateType string

const (
	HealthStateTypeDegraded      HealthStateType = "Degraded"
	HealthStateTypeDown          HealthStateType = "Down"
	HealthStateTypeFlushing      HealthStateType = "Flushing"
	HealthStateTypeHealthy       HealthStateType = "Healthy"
	HealthStateTypeStopped       HealthStateType = "Stopped"
	HealthStateTypeStopping      HealthStateType = "Stopping"
	HealthStateTypeTransitioning HealthStateType = "Transitioning"
	HealthStateTypeUnknown       HealthStateType = "Unknown"
	HealthStateTypeUpgrading     HealthStateType = "Upgrading"
)

// PossibleHealthStateTypeValues returns the possible values for the HealthStateType const type.
func PossibleHealthStateTypeValues() []HealthStateType {
	return []HealthStateType{
		HealthStateTypeDegraded,
		HealthStateTypeDown,
		HealthStateTypeFlushing,
		HealthStateTypeHealthy,
		HealthStateTypeStopped,
		HealthStateTypeStopping,
		HealthStateTypeTransitioning,
		HealthStateTypeUnknown,
		HealthStateTypeUpgrading,
	}
}

// ToPtr returns a *HealthStateType pointing to the current value.
func (c HealthStateType) ToPtr() *HealthStateType {
	return &c
}

type MetricAggregationType string

const (
	MetricAggregationTypeAverage      MetricAggregationType = "Average"
	MetricAggregationTypeCount        MetricAggregationType = "Count"
	MetricAggregationTypeMaximum      MetricAggregationType = "Maximum"
	MetricAggregationTypeMinimum      MetricAggregationType = "Minimum"
	MetricAggregationTypeNone         MetricAggregationType = "None"
	MetricAggregationTypeNotSpecified MetricAggregationType = "NotSpecified"
	MetricAggregationTypeTotal        MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns the possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{
		MetricAggregationTypeAverage,
		MetricAggregationTypeCount,
		MetricAggregationTypeMaximum,
		MetricAggregationTypeMinimum,
		MetricAggregationTypeNone,
		MetricAggregationTypeNotSpecified,
		MetricAggregationTypeTotal,
	}
}

// ToPtr returns a *MetricAggregationType pointing to the current value.
func (c MetricAggregationType) ToPtr() *MetricAggregationType {
	return &c
}

// NfsAccessRuleAccess - Access allowed by this rule.
type NfsAccessRuleAccess string

const (
	NfsAccessRuleAccessNo NfsAccessRuleAccess = "no"
	NfsAccessRuleAccessRo NfsAccessRuleAccess = "ro"
	NfsAccessRuleAccessRw NfsAccessRuleAccess = "rw"
)

// PossibleNfsAccessRuleAccessValues returns the possible values for the NfsAccessRuleAccess const type.
func PossibleNfsAccessRuleAccessValues() []NfsAccessRuleAccess {
	return []NfsAccessRuleAccess{
		NfsAccessRuleAccessNo,
		NfsAccessRuleAccessRo,
		NfsAccessRuleAccessRw,
	}
}

// ToPtr returns a *NfsAccessRuleAccess pointing to the current value.
func (c NfsAccessRuleAccess) ToPtr() *NfsAccessRuleAccess {
	return &c
}

// NfsAccessRuleScope - Scope for this rule. The scope and filter determine which clients match the rule.
type NfsAccessRuleScope string

const (
	NfsAccessRuleScopeDefault NfsAccessRuleScope = "default"
	NfsAccessRuleScopeHost    NfsAccessRuleScope = "host"
	NfsAccessRuleScopeNetwork NfsAccessRuleScope = "network"
)

// PossibleNfsAccessRuleScopeValues returns the possible values for the NfsAccessRuleScope const type.
func PossibleNfsAccessRuleScopeValues() []NfsAccessRuleScope {
	return []NfsAccessRuleScope{
		NfsAccessRuleScopeDefault,
		NfsAccessRuleScopeHost,
		NfsAccessRuleScopeNetwork,
	}
}

// ToPtr returns a *NfsAccessRuleScope pointing to the current value.
func (c NfsAccessRuleScope) ToPtr() *NfsAccessRuleScope {
	return &c
}

// OperationalStateType - Storage target operational state.
type OperationalStateType string

const (
	OperationalStateTypeBusy      OperationalStateType = "Busy"
	OperationalStateTypeFlushing  OperationalStateType = "Flushing"
	OperationalStateTypeReady     OperationalStateType = "Ready"
	OperationalStateTypeSuspended OperationalStateType = "Suspended"
)

// PossibleOperationalStateTypeValues returns the possible values for the OperationalStateType const type.
func PossibleOperationalStateTypeValues() []OperationalStateType {
	return []OperationalStateType{
		OperationalStateTypeBusy,
		OperationalStateTypeFlushing,
		OperationalStateTypeReady,
		OperationalStateTypeSuspended,
	}
}

// ToPtr returns a *OperationalStateType pointing to the current value.
func (c OperationalStateType) ToPtr() *OperationalStateType {
	return &c
}

// ProvisioningStateType - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
type ProvisioningStateType string

const (
	ProvisioningStateTypeCancelled ProvisioningStateType = "Cancelled"
	ProvisioningStateTypeCreating  ProvisioningStateType = "Creating"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
	ProvisioningStateTypeUpdating  ProvisioningStateType = "Updating"
)

// PossibleProvisioningStateTypeValues returns the possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{
		ProvisioningStateTypeCancelled,
		ProvisioningStateTypeCreating,
		ProvisioningStateTypeDeleting,
		ProvisioningStateTypeFailed,
		ProvisioningStateTypeSucceeded,
		ProvisioningStateTypeUpdating,
	}
}

// ToPtr returns a *ProvisioningStateType pointing to the current value.
func (c ProvisioningStateType) ToPtr() *ProvisioningStateType {
	return &c
}

// ReasonCode - The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". "QuotaId"
// is set when the SKU has requiredQuotas parameter as the subscription does not belong to that
// quota. "NotAvailableForSubscription" is related to capacity at the datacenter.
type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaID                     ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns the possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeQuotaID,
	}
}

// ToPtr returns a *ReasonCode pointing to the current value.
func (c ReasonCode) ToPtr() *ReasonCode {
	return &c
}

// StorageTargetType - Type of the Storage Target.
type StorageTargetType string

const (
	StorageTargetTypeBlobNfs StorageTargetType = "blobNfs"
	StorageTargetTypeClfs    StorageTargetType = "clfs"
	StorageTargetTypeNfs3    StorageTargetType = "nfs3"
	StorageTargetTypeUnknown StorageTargetType = "unknown"
)

// PossibleStorageTargetTypeValues returns the possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{
		StorageTargetTypeBlobNfs,
		StorageTargetTypeClfs,
		StorageTargetTypeNfs3,
		StorageTargetTypeUnknown,
	}
}

// ToPtr returns a *StorageTargetType pointing to the current value.
func (c StorageTargetType) ToPtr() *StorageTargetType {
	return &c
}

// UsernameDownloadedType - Indicates whether or not the HPC Cache has performed the username download successfully.
type UsernameDownloadedType string

const (
	UsernameDownloadedTypeError UsernameDownloadedType = "Error"
	UsernameDownloadedTypeNo    UsernameDownloadedType = "No"
	UsernameDownloadedTypeYes   UsernameDownloadedType = "Yes"
)

// PossibleUsernameDownloadedTypeValues returns the possible values for the UsernameDownloadedType const type.
func PossibleUsernameDownloadedTypeValues() []UsernameDownloadedType {
	return []UsernameDownloadedType{
		UsernameDownloadedTypeError,
		UsernameDownloadedTypeNo,
		UsernameDownloadedTypeYes,
	}
}

// ToPtr returns a *UsernameDownloadedType pointing to the current value.
func (c UsernameDownloadedType) ToPtr() *UsernameDownloadedType {
	return &c
}

// UsernameSource - This setting determines how the cache gets username and group names for clients.
type UsernameSource string

const (
	UsernameSourceAD   UsernameSource = "AD"
	UsernameSourceFile UsernameSource = "File"
	UsernameSourceLDAP UsernameSource = "LDAP"
	UsernameSourceNone UsernameSource = "None"
)

// PossibleUsernameSourceValues returns the possible values for the UsernameSource const type.
func PossibleUsernameSourceValues() []UsernameSource {
	return []UsernameSource{
		UsernameSourceAD,
		UsernameSourceFile,
		UsernameSourceLDAP,
		UsernameSourceNone,
	}
}

// ToPtr returns a *UsernameSource pointing to the current value.
func (c UsernameSource) ToPtr() *UsernameSource {
	return &c
}
