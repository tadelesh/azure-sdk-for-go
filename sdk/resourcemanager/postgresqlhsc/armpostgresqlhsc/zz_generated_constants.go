//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlhsc

const (
	moduleName    = "armpostgresqlhsc"
	moduleVersion = "v0.3.0"
)

// CitusVersion - The Citus version.
type CitusVersion string

const (
	CitusVersionEight3 CitusVersion = "8.3"
	CitusVersionNine0  CitusVersion = "9.0"
	CitusVersionNine1  CitusVersion = "9.1"
	CitusVersionNine2  CitusVersion = "9.2"
	CitusVersionNine3  CitusVersion = "9.3"
	CitusVersionNine4  CitusVersion = "9.4"
	CitusVersionNine5  CitusVersion = "9.5"
)

// PossibleCitusVersionValues returns the possible values for the CitusVersion const type.
func PossibleCitusVersionValues() []CitusVersion {
	return []CitusVersion{
		CitusVersionEight3,
		CitusVersionNine0,
		CitusVersionNine1,
		CitusVersionNine2,
		CitusVersionNine3,
		CitusVersionNine4,
		CitusVersionNine5,
	}
}

// ToPtr returns a *CitusVersion pointing to the current value.
func (c CitusVersion) ToPtr() *CitusVersion {
	return &c
}

// ConfigurationDataType - Data type of the configuration.
type ConfigurationDataType string

const (
	ConfigurationDataTypeBoolean     ConfigurationDataType = "Boolean"
	ConfigurationDataTypeEnumeration ConfigurationDataType = "Enumeration"
	ConfigurationDataTypeInteger     ConfigurationDataType = "Integer"
	ConfigurationDataTypeNumeric     ConfigurationDataType = "Numeric"
)

// PossibleConfigurationDataTypeValues returns the possible values for the ConfigurationDataType const type.
func PossibleConfigurationDataTypeValues() []ConfigurationDataType {
	return []ConfigurationDataType{
		ConfigurationDataTypeBoolean,
		ConfigurationDataTypeEnumeration,
		ConfigurationDataTypeInteger,
		ConfigurationDataTypeNumeric,
	}
}

// ToPtr returns a *ConfigurationDataType pointing to the current value.
func (c ConfigurationDataType) ToPtr() *ConfigurationDataType {
	return &c
}

// CreateMode - The mode to create a new server group.
type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReadReplica        CreateMode = "ReadReplica"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModePointInTimeRestore,
		CreateModeReadReplica,
	}
}

// ToPtr returns a *CreateMode pointing to the current value.
func (c CreateMode) ToPtr() *CreateMode {
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

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	OperationOriginSystem       OperationOrigin = "system"
	OperationOriginUser         OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginNotSpecified,
		OperationOriginSystem,
		OperationOriginUser,
	}
}

// ToPtr returns a *OperationOrigin pointing to the current value.
func (c OperationOrigin) ToPtr() *OperationOrigin {
	return &c
}

// PostgreSQLVersion - The PostgreSQL version.
type PostgreSQLVersion string

const (
	PostgreSQLVersionEleven PostgreSQLVersion = "11"
	PostgreSQLVersionTwelve PostgreSQLVersion = "12"
)

// PossiblePostgreSQLVersionValues returns the possible values for the PostgreSQLVersion const type.
func PossiblePostgreSQLVersionValues() []PostgreSQLVersion {
	return []PostgreSQLVersion{
		PostgreSQLVersionEleven,
		PostgreSQLVersionTwelve,
	}
}

// ToPtr returns a *PostgreSQLVersion pointing to the current value.
func (c PostgreSQLVersion) ToPtr() *PostgreSQLVersion {
	return &c
}

// ResourceProviderType - The resource provider type of server group.
type ResourceProviderType string

const (
	ResourceProviderTypeMarlin ResourceProviderType = "Marlin"
	ResourceProviderTypeMeru   ResourceProviderType = "Meru"
)

// PossibleResourceProviderTypeValues returns the possible values for the ResourceProviderType const type.
func PossibleResourceProviderTypeValues() []ResourceProviderType {
	return []ResourceProviderType{
		ResourceProviderTypeMarlin,
		ResourceProviderTypeMeru,
	}
}

// ToPtr returns a *ResourceProviderType pointing to the current value.
func (c ResourceProviderType) ToPtr() *ResourceProviderType {
	return &c
}

// ServerEdition - The edition of a server (default: GeneralPurpose).
type ServerEdition string

const (
	ServerEditionGeneralPurpose  ServerEdition = "GeneralPurpose"
	ServerEditionMemoryOptimized ServerEdition = "MemoryOptimized"
)

// PossibleServerEditionValues returns the possible values for the ServerEdition const type.
func PossibleServerEditionValues() []ServerEdition {
	return []ServerEdition{
		ServerEditionGeneralPurpose,
		ServerEditionMemoryOptimized,
	}
}

// ToPtr returns a *ServerEdition pointing to the current value.
func (c ServerEdition) ToPtr() *ServerEdition {
	return &c
}

// ServerHaState - A state of a server role group/server that is visible to user for HA feature.
type ServerHaState string

const (
	ServerHaStateCreatingStandby ServerHaState = "CreatingStandby"
	ServerHaStateFailingOver     ServerHaState = "FailingOver"
	ServerHaStateHealthy         ServerHaState = "Healthy"
	ServerHaStateNotEnabled      ServerHaState = "NotEnabled"
	ServerHaStateNotSync         ServerHaState = "NotSync"
	ServerHaStateRemovingStandby ServerHaState = "RemovingStandby"
	ServerHaStateReplicatingData ServerHaState = "ReplicatingData"
)

// PossibleServerHaStateValues returns the possible values for the ServerHaState const type.
func PossibleServerHaStateValues() []ServerHaState {
	return []ServerHaState{
		ServerHaStateCreatingStandby,
		ServerHaStateFailingOver,
		ServerHaStateHealthy,
		ServerHaStateNotEnabled,
		ServerHaStateNotSync,
		ServerHaStateRemovingStandby,
		ServerHaStateReplicatingData,
	}
}

// ToPtr returns a *ServerHaState pointing to the current value.
func (c ServerHaState) ToPtr() *ServerHaState {
	return &c
}

// ServerRole - The role of a server.
type ServerRole string

const (
	ServerRoleCoordinator ServerRole = "Coordinator"
	ServerRoleWorker      ServerRole = "Worker"
)

// PossibleServerRoleValues returns the possible values for the ServerRole const type.
func PossibleServerRoleValues() []ServerRole {
	return []ServerRole{
		ServerRoleCoordinator,
		ServerRoleWorker,
	}
}

// ToPtr returns a *ServerRole pointing to the current value.
func (c ServerRole) ToPtr() *ServerRole {
	return &c
}

// ServerState - A state of a server group/server that is visible to user.
type ServerState string

const (
	ServerStateDisabled     ServerState = "Disabled"
	ServerStateDropping     ServerState = "Dropping"
	ServerStateProvisioning ServerState = "Provisioning"
	ServerStateReady        ServerState = "Ready"
	ServerStateStarting     ServerState = "Starting"
	ServerStateStopped      ServerState = "Stopped"
	ServerStateStopping     ServerState = "Stopping"
	ServerStateUpdating     ServerState = "Updating"
)

// PossibleServerStateValues returns the possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{
		ServerStateDisabled,
		ServerStateDropping,
		ServerStateProvisioning,
		ServerStateReady,
		ServerStateStarting,
		ServerStateStopped,
		ServerStateStopping,
		ServerStateUpdating,
	}
}

// ToPtr returns a *ServerState pointing to the current value.
func (c ServerState) ToPtr() *ServerState {
	return &c
}
