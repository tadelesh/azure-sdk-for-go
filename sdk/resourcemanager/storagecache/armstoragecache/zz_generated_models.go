//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import "time"

// APIOperation - REST API operation description: see https://github.com/Azure/azure-rest-api-specs/blob/master/documentation/openapi-authoring-automated-guidelines.md#r3023-operationsapiimplementation
type APIOperation struct {
	// The object that represents the operation.
	Display *APIOperationDisplay `json:"display,omitempty"`

	// The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// Origin of the operation.
	Origin *string `json:"origin,omitempty"`

	// Additional details about an operation.
	Properties *APIOperationProperties `json:"properties,omitempty"`
}

// APIOperationDisplay - The object that represents the operation.
type APIOperationDisplay struct {
	// The description of the operation
	Description *string `json:"description,omitempty"`

	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.StorageCache
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Cache, etc.
	Resource *string `json:"resource,omitempty"`
}

// APIOperationListResult - Result of the request to list Resource Provider operations. It contains a list of operations and
// a URL link to get the next set of results.
type APIOperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Resource Provider operations supported by the Microsoft.StorageCache resource provider.
	Value []*APIOperation `json:"value,omitempty"`
}

// APIOperationProperties - Additional details about an operation.
type APIOperationProperties struct {
	// Specification of the all the metrics provided for a resource type.
	ServiceSpecification *APIOperationPropertiesServiceSpecification `json:"serviceSpecification,omitempty"`
}

// APIOperationPropertiesServiceSpecification - Specification of the all the metrics provided for a resource type.
type APIOperationPropertiesServiceSpecification struct {
	// Details about operations related to metrics.
	MetricSpecifications []*MetricSpecification `json:"metricSpecifications,omitempty"`
}

// AscOperation - The status of operation.
type AscOperation struct {
	// The end time of the operation.
	EndTime *string `json:"endTime,omitempty"`

	// The error detail of the operation if any.
	Error *ErrorResponse `json:"error,omitempty"`

	// The operation Id.
	ID *string `json:"id,omitempty"`

	// The operation name.
	Name *string `json:"name,omitempty"`

	// Additional operation-specific properties
	Properties *AscOperationProperties `json:"properties,omitempty"`

	// The start time of the operation.
	StartTime *string `json:"startTime,omitempty"`

	// The status of the operation.
	Status *string `json:"status,omitempty"`
}

// AscOperationProperties - Additional operation-specific output.
type AscOperationProperties struct {
	// Additional operation-specific output.
	Output map[string]interface{} `json:"output,omitempty"`
}

// AscOperationsClientGetOptions contains the optional parameters for the AscOperationsClient.Get method.
type AscOperationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BlobNfsTarget - Properties pertaining to the BlobNfsTarget.
type BlobNfsTarget struct {
	// Resource ID of the storage container.
	Target *string `json:"target,omitempty"`

	// Identifies the StorageCache usage model to be used for this storage target.
	UsageModel *string `json:"usageModel,omitempty"`
}

// Cache - A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type Cache struct {
	// The identity of the cache, if configured.
	Identity *CacheIdentity `json:"identity,omitempty"`

	// Region name string.
	Location *string `json:"location,omitempty"`

	// Properties of the Cache.
	Properties *CacheProperties `json:"properties,omitempty"`

	// SKU for the Cache.
	SKU *CacheSKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource ID of the Cache.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of Cache.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the Cache; Microsoft.StorageCache/Cache
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CacheActiveDirectorySettings - Active Directory settings used to join a cache to a domain.
type CacheActiveDirectorySettings struct {
	// REQUIRED; The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server. Length must
	// 1-15 characters from the class [-0-9a-zA-Z].
	CacheNetBiosName *string `json:"cacheNetBiosName,omitempty"`

	// REQUIRED; The fully qualified domain name of the Active Directory domain controller.
	DomainName *string `json:"domainName,omitempty"`

	// REQUIRED; The Active Directory domain's NetBIOS name.
	DomainNetBiosName *string `json:"domainNetBiosName,omitempty"`

	// REQUIRED; Primary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
	PrimaryDNSIPAddress *string `json:"primaryDnsIpAddress,omitempty"`

	// Active Directory admin credentials used to join the HPC Cache to a domain.
	Credentials *CacheActiveDirectorySettingsCredentials `json:"credentials,omitempty"`

	// Secondary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
	SecondaryDNSIPAddress *string `json:"secondaryDnsIpAddress,omitempty"`

	// READ-ONLY; True if the HPC Cache is joined to the Active Directory domain.
	DomainJoined *DomainJoinedType `json:"domainJoined,omitempty" azure:"ro"`
}

// CacheActiveDirectorySettingsCredentials - Active Directory admin credentials used to join the HPC Cache to a domain.
type CacheActiveDirectorySettingsCredentials struct {
	// REQUIRED; Plain text password of the Active Directory domain administrator. This value is stored encrypted and not returned
	// on response.
	Password *string `json:"password,omitempty"`

	// REQUIRED; Username of the Active Directory domain administrator. This value is stored encrypted and not returned on response.
	Username *string `json:"username,omitempty"`
}

// CacheDirectorySettings - Cache Directory Services settings.
type CacheDirectorySettings struct {
	// Specifies settings for joining the HPC Cache to an Active Directory domain.
	ActiveDirectory *CacheActiveDirectorySettings `json:"activeDirectory,omitempty"`

	// Specifies settings for Extended Groups. Extended Groups allows users to be members of more than 16 groups.
	UsernameDownload *CacheUsernameDownloadSettings `json:"usernameDownload,omitempty"`
}

// CacheEncryptionSettings - Cache encryption settings.
type CacheEncryptionSettings struct {
	// Specifies the location of the key encryption key in Key Vault.
	KeyEncryptionKey *KeyVaultKeyReference `json:"keyEncryptionKey,omitempty"`

	// Specifies whether the service will automatically rotate to the newest version of the key in the Key Vault.
	RotationToLatestKeyVersionEnabled *bool `json:"rotationToLatestKeyVersionEnabled,omitempty"`
}

// CacheHealth - An indication of Cache health. Gives more information about health than just that related to provisioning.
type CacheHealth struct {
	// List of Cache health states.
	State *HealthStateType `json:"state,omitempty"`

	// Describes explanation of state.
	StatusDescription *string `json:"statusDescription,omitempty"`

	// READ-ONLY; Outstanding conditions that need to be investigated and resolved.
	Conditions []*Condition `json:"conditions,omitempty" azure:"ro"`
}

// CacheIdentity - Cache identity properties.
type CacheIdentity struct {
	// The type of identity used for the cache
	Type *CacheIdentityType `json:"type,omitempty"`

	// A dictionary where each key is a user assigned identity resource ID, and each key's value is an empty dictionary.
	UserAssignedIdentities map[string]*UserAssignedIdentitiesValue `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The principal ID for the system-assigned identity of the cache.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID associated with the cache.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// CacheNetworkSettings - Cache network settings.
type CacheNetworkSettings struct {
	// DNS search domain
	DNSSearchDomain *string `json:"dnsSearchDomain,omitempty"`

	// DNS servers for the cache to use. It will be set from the network configuration if no value is provided.
	DNSServers []*string `json:"dnsServers,omitempty"`

	// The IPv4 maximum transmission unit configured for the subnet.
	Mtu *int32 `json:"mtu,omitempty"`

	// NTP server IP Address or FQDN for the cache to use. The default is time.windows.com.
	NtpServer *string `json:"ntpServer,omitempty"`

	// READ-ONLY; Array of additional IP addresses used by this Cache.
	UtilityAddresses []*string `json:"utilityAddresses,omitempty" azure:"ro"`
}

// CacheProperties - Properties of the Cache.
type CacheProperties struct {
	// The size of this Cache, in GB.
	CacheSizeGB *int32 `json:"cacheSizeGB,omitempty"`

	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings *CacheDirectorySettings `json:"directoryServicesSettings,omitempty"`

	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettings `json:"encryptionSettings,omitempty"`

	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettings `json:"networkSettings,omitempty"`

	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettings `json:"securitySettings,omitempty"`

	// Subnet used for the Cache.
	Subnet *string `json:"subnet,omitempty"`

	// READ-ONLY; Health of the Cache.
	Health *CacheHealth `json:"health,omitempty" azure:"ro"`

	// READ-ONLY; Array of IP addresses that can be used by clients mounting this Cache.
	MountAddresses []*string `json:"mountAddresses,omitempty" azure:"ro"`

	// READ-ONLY; ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *ProvisioningStateType `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Upgrade status of the Cache.
	UpgradeStatus *CacheUpgradeStatus `json:"upgradeStatus,omitempty" azure:"ro"`
}

// CacheSKU - SKU for the Cache.
type CacheSKU struct {
	// SKU name for this Cache.
	Name *string `json:"name,omitempty"`
}

// CacheSecuritySettings - Cache security settings.
type CacheSecuritySettings struct {
	// NFS access policies defined for this cache.
	AccessPolicies []*NfsAccessPolicy `json:"accessPolicies,omitempty"`
}

// CacheUpgradeStatus - Properties describing the software upgrade state of the Cache.
type CacheUpgradeStatus struct {
	// READ-ONLY; Version string of the firmware currently installed on this Cache.
	CurrentFirmwareVersion *string `json:"currentFirmwareVersion,omitempty" azure:"ro"`

	// READ-ONLY; Time at which the pending firmware update will automatically be installed on the Cache.
	FirmwareUpdateDeadline *time.Time `json:"firmwareUpdateDeadline,omitempty" azure:"ro"`

	// READ-ONLY; True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed
	// after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
	FirmwareUpdateStatus *FirmwareStatusType `json:"firmwareUpdateStatus,omitempty" azure:"ro"`

	// READ-ONLY; Time of the last successful firmware update.
	LastFirmwareUpdate *time.Time `json:"lastFirmwareUpdate,omitempty" azure:"ro"`

	// READ-ONLY; When firmwareUpdateAvailable is true, this field holds the version string for the update.
	PendingFirmwareVersion *string `json:"pendingFirmwareVersion,omitempty" azure:"ro"`
}

// CacheUsernameDownloadSettings - Settings for Extended Groups username and group download.
type CacheUsernameDownloadSettings struct {
	// Determines if the certificate should be automatically downloaded. This applies to 'caCertificateURI' only if 'requireValidCertificate'
	// is true.
	AutoDownloadCertificate *bool `json:"autoDownloadCertificate,omitempty"`

	// The URI of the CA certificate to validate the LDAP secure connection. This field must be populated when 'requireValidCertificate'
	// is set to true.
	CaCertificateURI *string `json:"caCertificateURI,omitempty"`

	// When present, these are the credentials for the secure LDAP connection.
	Credentials *CacheUsernameDownloadSettingsCredentials `json:"credentials,omitempty"`

	// Whether or not the LDAP connection should be encrypted.
	EncryptLdapConnection *bool `json:"encryptLdapConnection,omitempty"`

	// Whether or not Extended Groups is enabled.
	ExtendedGroups *bool `json:"extendedGroups,omitempty"`

	// The URI of the file containing group information (in /etc/group file format). This field must be populated when 'usernameSource'
	// is set to 'File'.
	GroupFileURI *string `json:"groupFileURI,omitempty"`

	// The base distinguished name for the LDAP domain.
	LdapBaseDN *string `json:"ldapBaseDN,omitempty"`

	// The fully qualified domain name or IP address of the LDAP server to use.
	LdapServer *string `json:"ldapServer,omitempty"`

	// Determines if the certificates must be validated by a certificate authority. When true, caCertificateURI must be provided.
	RequireValidCertificate *bool `json:"requireValidCertificate,omitempty"`

	// The URI of the file containing user information (in /etc/passwd file format). This field must be populated when 'usernameSource'
	// is set to 'File'.
	UserFileURI *string `json:"userFileURI,omitempty"`

	// This setting determines how the cache gets username and group names for clients.
	UsernameSource *UsernameSource `json:"usernameSource,omitempty"`

	// READ-ONLY; Indicates whether or not the HPC Cache has performed the username download successfully.
	UsernameDownloaded *UsernameDownloadedType `json:"usernameDownloaded,omitempty" azure:"ro"`
}

// CacheUsernameDownloadSettingsCredentials - When present, these are the credentials for the secure LDAP connection.
type CacheUsernameDownloadSettingsCredentials struct {
	// The Bind Distinguished Name identity to be used in the secure LDAP connection. This value is stored encrypted and not returned
	// on response.
	BindDn *string `json:"bindDn,omitempty"`

	// The Bind password to be used in the secure LDAP connection. This value is stored encrypted and not returned on response.
	BindPassword *string `json:"bindPassword,omitempty"`
}

// CachesClientBeginCreateOrUpdateOptions contains the optional parameters for the CachesClient.BeginCreateOrUpdate method.
type CachesClientBeginCreateOrUpdateOptions struct {
	// Object containing the user-selectable properties of the new Cache. If read-only properties are included, they must match
	// the existing values of those properties.
	Cache *Cache
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginDebugInfoOptions contains the optional parameters for the CachesClient.BeginDebugInfo method.
type CachesClientBeginDebugInfoOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginDeleteOptions contains the optional parameters for the CachesClient.BeginDelete method.
type CachesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginFlushOptions contains the optional parameters for the CachesClient.BeginFlush method.
type CachesClientBeginFlushOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginStartOptions contains the optional parameters for the CachesClient.BeginStart method.
type CachesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginStopOptions contains the optional parameters for the CachesClient.BeginStop method.
type CachesClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginUpgradeFirmwareOptions contains the optional parameters for the CachesClient.BeginUpgradeFirmware method.
type CachesClientBeginUpgradeFirmwareOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientGetOptions contains the optional parameters for the CachesClient.Get method.
type CachesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CachesClientListByResourceGroupOptions contains the optional parameters for the CachesClient.ListByResourceGroup method.
type CachesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CachesClientListOptions contains the optional parameters for the CachesClient.List method.
type CachesClientListOptions struct {
	// placeholder for future optional parameters
}

// CachesClientUpdateOptions contains the optional parameters for the CachesClient.Update method.
type CachesClientUpdateOptions struct {
	// Object containing the user-selectable properties of the Cache. If read-only properties are included, they must match the
	// existing values of those properties.
	Cache *Cache
}

// CachesListResult - Result of the request to list Caches. It contains a list of Caches and a URL link to get the next set
// of results.
type CachesListResult struct {
	// URL to get the next set of Cache list results, if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Caches.
	Value []*Cache `json:"value,omitempty"`
}

// ClfsTarget - Properties pertaining to the ClfsTarget
type ClfsTarget struct {
	// Resource ID of storage container.
	Target *string `json:"target,omitempty"`
}

// CloudError - An error response.
type CloudError struct {
	// The body of the error.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// Condition - Outstanding conditions that will need to be resolved.
type Condition struct {
	// READ-ONLY; The issue requiring attention.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The time when the condition was raised.
	Timestamp *time.Time `json:"timestamp,omitempty" azure:"ro"`
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// KeyVaultKeyReference - Describes a reference to Key Vault Key.
type KeyVaultKeyReference struct {
	// REQUIRED; The URL referencing a key encryption key in Key Vault.
	KeyURL *string `json:"keyUrl,omitempty"`

	// REQUIRED; Describes a resource Id to source Key Vault.
	SourceVault *KeyVaultKeyReferenceSourceVault `json:"sourceVault,omitempty"`
}

// KeyVaultKeyReferenceSourceVault - Describes a resource Id to source Key Vault.
type KeyVaultKeyReferenceSourceVault struct {
	// Resource Id.
	ID *string `json:"id,omitempty"`
}

// MetricDimension - Specifications of the Dimension of metrics.
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string `json:"displayName,omitempty"`

	// Internal name of the dimension.
	InternalName *string `json:"internalName,omitempty"`

	// Name of the dimension
	Name *string `json:"name,omitempty"`

	// To be exported to shoe box.
	ToBeExportedForShoebox *bool `json:"toBeExportedForShoebox,omitempty"`
}

// MetricSpecification - Details about operation related to metrics.
type MetricSpecification struct {
	// The type of metric aggregation.
	AggregationType *string `json:"aggregationType,omitempty"`

	// Dimensions of the metric
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`

	// The description of the metric.
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// Localized display name of the metric.
	DisplayName *string `json:"displayName,omitempty"`

	// Type of metrics.
	MetricClass *string `json:"metricClass,omitempty"`

	// The name of the metric.
	Name *string `json:"name,omitempty"`

	// Support metric aggregation type.
	SupportedAggregationTypes []*MetricAggregationType `json:"supportedAggregationTypes,omitempty"`

	// The unit that the metric is measured in.
	Unit *string `json:"unit,omitempty"`
}

// NamespaceJunction - A namespace junction.
type NamespaceJunction struct {
	// Namespace path on a Cache for a Storage Target.
	NamespacePath *string `json:"namespacePath,omitempty"`

	// Name of the access policy applied to this junction.
	NfsAccessPolicy *string `json:"nfsAccessPolicy,omitempty"`

	// NFS export where targetPath exists.
	NfsExport *string `json:"nfsExport,omitempty"`

	// Path in Storage Target to which namespacePath points.
	TargetPath *string `json:"targetPath,omitempty"`
}

// Nfs3Target - Properties pertaining to the Nfs3Target
type Nfs3Target struct {
	// IP address or host name of an NFSv3 host (e.g., 10.0.44.44).
	Target *string `json:"target,omitempty"`

	// Identifies the StorageCache usage model to be used for this storage target.
	UsageModel *string `json:"usageModel,omitempty"`
}

// NfsAccessPolicy - A set of rules describing access policies applied to NFSv3 clients of the cache.
type NfsAccessPolicy struct {
	// REQUIRED; The set of rules describing client accesses allowed under this policy.
	AccessRules []*NfsAccessRule `json:"accessRules,omitempty"`

	// REQUIRED; Name identifying this policy. Access Policy names are not case sensitive.
	Name *string `json:"name,omitempty"`
}

// NfsAccessRule - Rule to place restrictions on portions of the cache namespace being presented to clients.
type NfsAccessRule struct {
	// REQUIRED; Access allowed by this rule.
	Access *NfsAccessRuleAccess `json:"access,omitempty"`

	// REQUIRED; Scope for this rule. The scope and filter determine which clients match the rule.
	Scope *NfsAccessRuleScope `json:"scope,omitempty"`

	// GID value that replaces 0 when rootSquash is true. This will use the value of anonymousUID if not provided.
	AnonymousGID *string `json:"anonymousGID,omitempty"`

	// UID value that replaces 0 when rootSquash is true. 65534 will be used if not provided.
	AnonymousUID *string `json:"anonymousUID,omitempty"`

	// Filter applied to the scope for this rule. The filter's format depends on its scope. 'default' scope matches all clients
	// and has no filter value. 'network' scope takes a filter in CIDR format (for
	// example, 10.99.1.0/24). 'host' takes an IP address or fully qualified domain name as filter. If a client does not match
	// any filter rule and there is no default rule, access is denied.
	Filter *string `json:"filter,omitempty"`

	// Map root accesses to anonymousUID and anonymousGID.
	RootSquash *bool `json:"rootSquash,omitempty"`

	// For the default policy, allow access to subdirectories under the root export. If this is set to no, clients can only mount
	// the path '/'. If set to yes, clients can mount a deeper path, like '/a/b'.
	SubmountAccess *bool `json:"submountAccess,omitempty"`

	// Allow SUID semantics.
	Suid *bool `json:"suid,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ResourceSKU - A resource SKU.
type ResourceSKU struct {
	// A list of capabilities of this SKU, such as throughput or ops/sec.
	Capabilities []*ResourceSKUCapabilities `json:"capabilities,omitempty"`

	// The set of locations where the SKU is available.
	LocationInfo []*ResourceSKULocationInfo `json:"locationInfo,omitempty"`

	// The name of this SKU.
	Name *string `json:"name,omitempty"`

	// The restrictions preventing this SKU from being used. This is empty if there are no restrictions.
	Restrictions []*Restriction `json:"restrictions,omitempty"`

	// READ-ONLY; The set of locations where the SKU is available. This is the supported and registered Azure Geo Regions (e.g.,
	// West US, East US, Southeast Asia, etc.).
	Locations []*string `json:"locations,omitempty" azure:"ro"`

	// READ-ONLY; The type of resource the SKU applies to.
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`
}

// ResourceSKUCapabilities - A resource SKU capability.
type ResourceSKUCapabilities struct {
	// Name of a capability, such as ops/sec.
	Name *string `json:"name,omitempty"`

	// Quantity, if the capability is measured by quantity.
	Value *string `json:"value,omitempty"`
}

// ResourceSKULocationInfo - Resource SKU location information.
type ResourceSKULocationInfo struct {
	// Location where this SKU is available.
	Location *string `json:"location,omitempty"`

	// Zones if any.
	Zones []*string `json:"zones,omitempty"`
}

// ResourceSKUsResult - The response from the List Cache SKUs operation.
type ResourceSKUsResult struct {
	// The URI to fetch the next page of Cache SKUs.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; The list of SKUs available for the subscription.
	Value []*ResourceSKU `json:"value,omitempty" azure:"ro"`
}

// Restriction - The restrictions preventing this SKU from being used.
type Restriction struct {
	// The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". "QuotaId" is set when
	// the SKU has requiredQuotas parameter as the subscription does not belong to that
	// quota. "NotAvailableForSubscription" is related to capacity at the datacenter.
	ReasonCode *ReasonCode `json:"reasonCode,omitempty"`

	// READ-ONLY; The type of restrictions. In this version, the only possible value for this is location.
	Type *string `json:"type,omitempty" azure:"ro"`

	// READ-ONLY; The value of restrictions. If the restriction type is set to location, then this would be the different locations
	// where the SKU is restricted.
	Values []*string `json:"values,omitempty" azure:"ro"`
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.List method.
type SKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// StorageTarget - Type of the Storage Target.
type StorageTarget struct {
	// StorageTarget properties
	Properties *StorageTargetProperties `json:"properties,omitempty"`

	// READ-ONLY; Resource ID of the Storage Target.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Region name string.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Name of the Storage Target.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type *string `json:"type,omitempty" azure:"ro"`
}

// StorageTargetClientBeginFlushOptions contains the optional parameters for the StorageTargetClient.BeginFlush method.
type StorageTargetClientBeginFlushOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetClientBeginResumeOptions contains the optional parameters for the StorageTargetClient.BeginResume method.
type StorageTargetClientBeginResumeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetClientBeginSuspendOptions contains the optional parameters for the StorageTargetClient.BeginSuspend method.
type StorageTargetClientBeginSuspendOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetProperties - Properties of the Storage Target.
type StorageTargetProperties struct {
	// REQUIRED; Type of the Storage Target.
	TargetType *StorageTargetType `json:"targetType,omitempty"`

	// Properties when targetType is blobNfs.
	BlobNfs *BlobNfsTarget `json:"blobNfs,omitempty"`

	// Properties when targetType is clfs.
	Clfs *ClfsTarget `json:"clfs,omitempty"`

	// List of Cache namespace junctions to target for namespace associations.
	Junctions []*NamespaceJunction `json:"junctions,omitempty"`

	// Properties when targetType is nfs3.
	Nfs3 *Nfs3Target `json:"nfs3,omitempty"`

	// Storage target operational state.
	State *OperationalStateType `json:"state,omitempty"`

	// Properties when targetType is unknown.
	Unknown *UnknownTarget `json:"unknown,omitempty"`

	// READ-ONLY; ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *ProvisioningStateType `json:"provisioningState,omitempty" azure:"ro"`
}

// StorageTargetResource - Resource used by a Cache.
type StorageTargetResource struct {
	// READ-ONLY; Resource ID of the Storage Target.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Region name string.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Name of the Storage Target.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type *string `json:"type,omitempty" azure:"ro"`
}

// StorageTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageTargetsClient.BeginCreateOrUpdate
// method.
type StorageTargetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
	// Object containing the definition of a Storage Target.
	Storagetarget *StorageTarget
}

// StorageTargetsClientBeginDNSRefreshOptions contains the optional parameters for the StorageTargetsClient.BeginDNSRefresh
// method.
type StorageTargetsClientBeginDNSRefreshOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientBeginDeleteOptions contains the optional parameters for the StorageTargetsClient.BeginDelete method.
type StorageTargetsClientBeginDeleteOptions struct {
	// Boolean value requesting the force delete operation for a storage target. Force delete discards unwritten-data in the cache
	// instead of flushing it to back-end storage.
	Force *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientGetOptions contains the optional parameters for the StorageTargetsClient.Get method.
type StorageTargetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageTargetsClientListByCacheOptions contains the optional parameters for the StorageTargetsClient.ListByCache method.
type StorageTargetsClientListByCacheOptions struct {
	// placeholder for future optional parameters
}

// StorageTargetsResult - A list of Storage Targets.
type StorageTargetsResult struct {
	// The URI to fetch the next page of Storage Targets.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of Storage Targets defined for the Cache.
	Value []*StorageTarget `json:"value,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// UnknownTarget - Properties pertaining to the UnknownTarget
type UnknownTarget struct {
	// Dictionary of string->string pairs containing information about the Storage Target.
	Attributes map[string]*string `json:"attributes,omitempty"`
}

// UsageModel - A usage model.
type UsageModel struct {
	// Localized information describing this usage model.
	Display *UsageModelDisplay `json:"display,omitempty"`

	// Non-localized keyword name for this usage model.
	ModelName *string `json:"modelName,omitempty"`

	// The type of Storage Target to which this model is applicable (only nfs3 as of this version).
	TargetType *string `json:"targetType,omitempty"`
}

// UsageModelDisplay - Localized information describing this usage model.
type UsageModelDisplay struct {
	// String to display for this usage model.
	Description *string `json:"description,omitempty"`
}

// UsageModelsClientListOptions contains the optional parameters for the UsageModelsClient.List method.
type UsageModelsClientListOptions struct {
	// placeholder for future optional parameters
}

// UsageModelsResult - A list of Cache usage models.
type UsageModelsResult struct {
	// The URI to fetch the next page of Cache usage models.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of usage models available for the subscription.
	Value []*UsageModel `json:"value,omitempty"`
}

type UserAssignedIdentitiesValue struct {
	// READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}
