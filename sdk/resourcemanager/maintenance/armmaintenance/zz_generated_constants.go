//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

const (
	moduleName    = "armmaintenance"
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

// ImpactType - The impact type
type ImpactType string

const (
	// ImpactTypeFreeze - Pending updates can freeze network or disk io operation on resource.
	ImpactTypeFreeze ImpactType = "Freeze"
	// ImpactTypeNone - Pending updates has no impact on resource.
	ImpactTypeNone ImpactType = "None"
	// ImpactTypeRedeploy - Pending updates can redeploy resource.
	ImpactTypeRedeploy ImpactType = "Redeploy"
	// ImpactTypeRestart - Pending updates can cause resource to restart.
	ImpactTypeRestart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns the possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{
		ImpactTypeFreeze,
		ImpactTypeNone,
		ImpactTypeRedeploy,
		ImpactTypeRestart,
	}
}

// ToPtr returns a *ImpactType pointing to the current value.
func (c ImpactType) ToPtr() *ImpactType {
	return &c
}

// MaintenanceScope - Gets or sets maintenanceScope of the configuration
type MaintenanceScope string

const (
	// MaintenanceScopeExtension - This maintenance scope controls extension installation on VM/VMSS
	MaintenanceScopeExtension MaintenanceScope = "Extension"
	// MaintenanceScopeHost - This maintenance scope controls installation of azure platform updates i.e. services on physical
	// nodes hosting customer VMs.
	MaintenanceScopeHost MaintenanceScope = "Host"
	// MaintenanceScopeInGuestPatch - This maintenance scope controls installation of windows and linux packages on VM/VMSS
	MaintenanceScopeInGuestPatch MaintenanceScope = "InGuestPatch"
	// MaintenanceScopeOSImage - This maintenance scope controls os image installation on VM/VMSS
	MaintenanceScopeOSImage MaintenanceScope = "OSImage"
	// MaintenanceScopeSQLDB - This maintenance scope controls installation of SQL server platform updates.
	MaintenanceScopeSQLDB MaintenanceScope = "SQLDB"
	// MaintenanceScopeSQLManagedInstance - This maintenance scope controls installation of SQL managed instance platform update.
	MaintenanceScopeSQLManagedInstance MaintenanceScope = "SQLManagedInstance"
)

// PossibleMaintenanceScopeValues returns the possible values for the MaintenanceScope const type.
func PossibleMaintenanceScopeValues() []MaintenanceScope {
	return []MaintenanceScope{
		MaintenanceScopeExtension,
		MaintenanceScopeHost,
		MaintenanceScopeInGuestPatch,
		MaintenanceScopeOSImage,
		MaintenanceScopeSQLDB,
		MaintenanceScopeSQLManagedInstance,
	}
}

// ToPtr returns a *MaintenanceScope pointing to the current value.
func (c MaintenanceScope) ToPtr() *MaintenanceScope {
	return &c
}

// RebootOptions - Possible reboot preference as defined by the user based on which it would be decided to reboot the machine
// or not after the patch operation is completed.
type RebootOptions string

const (
	RebootOptionsAlways     RebootOptions = "Always"
	RebootOptionsIfRequired RebootOptions = "IfRequired"
	RebootOptionsNever      RebootOptions = "Never"
)

// PossibleRebootOptionsValues returns the possible values for the RebootOptions const type.
func PossibleRebootOptionsValues() []RebootOptions {
	return []RebootOptions{
		RebootOptionsAlways,
		RebootOptionsIfRequired,
		RebootOptionsNever,
	}
}

// ToPtr returns a *RebootOptions pointing to the current value.
func (c RebootOptions) ToPtr() *RebootOptions {
	return &c
}

// TaskScope - Global Task execute once when schedule trigger. Resource task execute for each VM.
type TaskScope string

const (
	TaskScopeGlobal   TaskScope = "Global"
	TaskScopeResource TaskScope = "Resource"
)

// PossibleTaskScopeValues returns the possible values for the TaskScope const type.
func PossibleTaskScopeValues() []TaskScope {
	return []TaskScope{
		TaskScopeGlobal,
		TaskScopeResource,
	}
}

// ToPtr returns a *TaskScope pointing to the current value.
func (c TaskScope) ToPtr() *TaskScope {
	return &c
}

// UpdateStatus - The status
type UpdateStatus string

const (
	// UpdateStatusCompleted - All updates are successfully applied.
	UpdateStatusCompleted UpdateStatus = "Completed"
	// UpdateStatusInProgress - Updates installation are in progress.
	UpdateStatusInProgress UpdateStatus = "InProgress"
	// UpdateStatusPending - There are pending updates to be installed.
	UpdateStatusPending UpdateStatus = "Pending"
	// UpdateStatusRetryLater - Updates installation failed and should be retried later.
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	// UpdateStatusRetryNow - Updates installation failed but are ready to retry again.
	UpdateStatusRetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns the possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{
		UpdateStatusCompleted,
		UpdateStatusInProgress,
		UpdateStatusPending,
		UpdateStatusRetryLater,
		UpdateStatusRetryNow,
	}
}

// ToPtr returns a *UpdateStatus pointing to the current value.
func (c UpdateStatus) ToPtr() *UpdateStatus {
	return &c
}

// Visibility - Gets or sets the visibility of the configuration. The default value is 'Custom'
type Visibility string

const (
	// VisibilityCustom - Only visible to users with permissions.
	VisibilityCustom Visibility = "Custom"
	// VisibilityPublic - Visible to all users.
	VisibilityPublic Visibility = "Public"
)

// PossibleVisibilityValues returns the possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{
		VisibilityCustom,
		VisibilityPublic,
	}
}

// ToPtr returns a *Visibility pointing to the current value.
func (c Visibility) ToPtr() *Visibility {
	return &c
}
