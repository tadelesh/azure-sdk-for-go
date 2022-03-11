//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Attributes.
func (a Attributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeUnix(objectMap, "created", a.Created)
	populate(objectMap, "enabled", a.Enabled)
	populateTimeUnix(objectMap, "exp", a.Expires)
	populateTimeUnix(objectMap, "nbf", a.NotBefore)
	populateTimeUnix(objectMap, "updated", a.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Attributes.
func (a *Attributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeUnix(val, &a.Created)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, &a.Enabled)
			delete(rawMsg, key)
		case "exp":
			err = unpopulateTimeUnix(val, &a.Expires)
			delete(rawMsg, key)
		case "nbf":
			err = unpopulateTimeUnix(val, &a.NotBefore)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeUnix(val, &a.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedManagedHsmListResult.
func (d DeletedManagedHsmListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DeletedManagedHsmProperties.
func (d DeletedManagedHsmProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "deletionDate", d.DeletionDate)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "mhsmId", d.MhsmID)
	populate(objectMap, "purgeProtectionEnabled", d.PurgeProtectionEnabled)
	populateTimeRFC3339(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedManagedHsmProperties.
func (d *DeletedManagedHsmProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deletionDate":
			err = unpopulateTimeRFC3339(val, &d.DeletionDate)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &d.Location)
			delete(rawMsg, key)
		case "mhsmId":
			err = unpopulate(val, &d.MhsmID)
			delete(rawMsg, key)
		case "purgeProtectionEnabled":
			err = unpopulate(val, &d.PurgeProtectionEnabled)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeRFC3339(val, &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &d.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedVaultListResult.
func (d DeletedVaultListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DeletedVaultProperties.
func (d DeletedVaultProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "deletionDate", d.DeletionDate)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "purgeProtectionEnabled", d.PurgeProtectionEnabled)
	populateTimeRFC3339(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "vaultId", d.VaultID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedVaultProperties.
func (d *DeletedVaultProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deletionDate":
			err = unpopulateTimeRFC3339(val, &d.DeletionDate)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &d.Location)
			delete(rawMsg, key)
		case "purgeProtectionEnabled":
			err = unpopulate(val, &d.PurgeProtectionEnabled)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeRFC3339(val, &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &d.Tags)
			delete(rawMsg, key)
		case "vaultId":
			err = unpopulate(val, &d.VaultID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Key.
func (k Key) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", k.ID)
	populate(objectMap, "location", k.Location)
	populate(objectMap, "name", k.Name)
	populate(objectMap, "properties", k.Properties)
	populate(objectMap, "tags", k.Tags)
	populate(objectMap, "type", k.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyCreateParameters.
func (k KeyCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", k.Properties)
	populate(objectMap, "tags", k.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyListResult.
func (k KeyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", k.NextLink)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyProperties.
func (k KeyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "attributes", k.Attributes)
	populate(objectMap, "curveName", k.CurveName)
	populate(objectMap, "keyOps", k.KeyOps)
	populate(objectMap, "keySize", k.KeySize)
	populate(objectMap, "keyUri", k.KeyURI)
	populate(objectMap, "keyUriWithVersion", k.KeyURIWithVersion)
	populate(objectMap, "kty", k.Kty)
	populate(objectMap, "release_policy", k.ReleasePolicy)
	populate(objectMap, "rotationPolicy", k.RotationPolicy)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyReleasePolicy.
func (k KeyReleasePolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", k.ContentType)
	populateByteArray(objectMap, "data", k.Data, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyReleasePolicy.
func (k *KeyReleasePolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, &k.ContentType)
			delete(rawMsg, key)
		case "data":
			err = runtime.DecodeByteArray(string(val), &k.Data, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MHSMNetworkRuleSet.
func (m MHSMNetworkRuleSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bypass", m.Bypass)
	populate(objectMap, "defaultAction", m.DefaultAction)
	populate(objectMap, "ipRules", m.IPRules)
	populate(objectMap, "virtualNetworkRules", m.VirtualNetworkRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MHSMPrivateEndpointConnection.
func (m MHSMPrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", m.Etag)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MHSMPrivateEndpointConnectionsListResult.
func (m MHSMPrivateEndpointConnectionsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MHSMPrivateLinkResource.
func (m MHSMPrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MHSMPrivateLinkResourceListResult.
func (m MHSMPrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MHSMPrivateLinkResourceProperties.
func (m MHSMPrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", m.GroupID)
	populate(objectMap, "requiredMembers", m.RequiredMembers)
	populate(objectMap, "requiredZoneNames", m.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedHsm.
func (m ManagedHsm) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedHsmListResult.
func (m ManagedHsmListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedHsmProperties.
func (m ManagedHsmProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "createMode", m.CreateMode)
	populate(objectMap, "enablePurgeProtection", m.EnablePurgeProtection)
	populate(objectMap, "enableSoftDelete", m.EnableSoftDelete)
	populate(objectMap, "hsmUri", m.HsmURI)
	populate(objectMap, "initialAdminObjectIds", m.InitialAdminObjectIDs)
	populate(objectMap, "networkAcls", m.NetworkACLs)
	populate(objectMap, "privateEndpointConnections", m.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", m.PublicNetworkAccess)
	populateTimeRFC3339(objectMap, "scheduledPurgeDate", m.ScheduledPurgeDate)
	populate(objectMap, "softDeleteRetentionInDays", m.SoftDeleteRetentionInDays)
	populate(objectMap, "statusMessage", m.StatusMessage)
	populate(objectMap, "tenantId", m.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagedHsmProperties.
func (m *ManagedHsmProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createMode":
			err = unpopulate(val, &m.CreateMode)
			delete(rawMsg, key)
		case "enablePurgeProtection":
			err = unpopulate(val, &m.EnablePurgeProtection)
			delete(rawMsg, key)
		case "enableSoftDelete":
			err = unpopulate(val, &m.EnableSoftDelete)
			delete(rawMsg, key)
		case "hsmUri":
			err = unpopulate(val, &m.HsmURI)
			delete(rawMsg, key)
		case "initialAdminObjectIds":
			err = unpopulate(val, &m.InitialAdminObjectIDs)
			delete(rawMsg, key)
		case "networkAcls":
			err = unpopulate(val, &m.NetworkACLs)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &m.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &m.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &m.PublicNetworkAccess)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeRFC3339(val, &m.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "softDeleteRetentionInDays":
			err = unpopulate(val, &m.SoftDeleteRetentionInDays)
			delete(rawMsg, key)
		case "statusMessage":
			err = unpopulate(val, &m.StatusMessage)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, &m.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedHsmResource.
func (m ManagedHsmResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecification.
func (m MetricSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "fillGapWithZero", m.FillGapWithZero)
	populate(objectMap, "internalMetricName", m.InternalMetricName)
	populate(objectMap, "lockAggregationType", m.LockAggregationType)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "supportedAggregationTypes", m.SupportedAggregationTypes)
	populate(objectMap, "supportedTimeGrainTypes", m.SupportedTimeGrainTypes)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSet.
func (n NetworkRuleSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bypass", n.Bypass)
	populate(objectMap, "defaultAction", n.DefaultAction)
	populate(objectMap, "ipRules", n.IPRules)
	populate(objectMap, "virtualNetworkRules", n.VirtualNetworkRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Permissions.
func (p Permissions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "certificates", p.Certificates)
	populate(objectMap, "keys", p.Keys)
	populate(objectMap, "secrets", p.Secrets)
	populate(objectMap, "storage", p.Storage)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnection.
func (p PrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResource.
func (p PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceListResult.
func (r ResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RotationPolicy.
func (r RotationPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "attributes", r.Attributes)
	populate(objectMap, "lifetimeActions", r.LifetimeActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Secret.
func (s Secret) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SecretAttributes.
func (s SecretAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeUnix(objectMap, "created", s.Created)
	populate(objectMap, "enabled", s.Enabled)
	populateTimeUnix(objectMap, "exp", s.Expires)
	populateTimeUnix(objectMap, "nbf", s.NotBefore)
	populateTimeUnix(objectMap, "updated", s.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretAttributes.
func (s *SecretAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeUnix(val, &s.Created)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, &s.Enabled)
			delete(rawMsg, key)
		case "exp":
			err = unpopulateTimeUnix(val, &s.Expires)
			delete(rawMsg, key)
		case "nbf":
			err = unpopulateTimeUnix(val, &s.NotBefore)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeUnix(val, &s.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecretCreateOrUpdateParameters.
func (s SecretCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SecretListResult.
func (s SecretListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SecretPatchParameters.
func (s SecretPatchParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceSpecification.
func (s ServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "logSpecifications", s.LogSpecifications)
	populate(objectMap, "metricSpecifications", s.MetricSpecifications)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Vault.
func (v Vault) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "systemData", v.SystemData)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultAccessPolicyProperties.
func (v VaultAccessPolicyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessPolicies", v.AccessPolicies)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultCreateOrUpdateParameters.
func (v VaultCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", v.Location)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultListResult.
func (v VaultListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultPatchParameters.
func (v VaultPatchParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultPatchProperties.
func (v VaultPatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessPolicies", v.AccessPolicies)
	populate(objectMap, "createMode", v.CreateMode)
	populate(objectMap, "enablePurgeProtection", v.EnablePurgeProtection)
	populate(objectMap, "enableRbacAuthorization", v.EnableRbacAuthorization)
	populate(objectMap, "enableSoftDelete", v.EnableSoftDelete)
	populate(objectMap, "enabledForDeployment", v.EnabledForDeployment)
	populate(objectMap, "enabledForDiskEncryption", v.EnabledForDiskEncryption)
	populate(objectMap, "enabledForTemplateDeployment", v.EnabledForTemplateDeployment)
	populate(objectMap, "networkAcls", v.NetworkACLs)
	populate(objectMap, "publicNetworkAccess", v.PublicNetworkAccess)
	populate(objectMap, "sku", v.SKU)
	populate(objectMap, "softDeleteRetentionInDays", v.SoftDeleteRetentionInDays)
	populate(objectMap, "tenantId", v.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VaultProperties.
func (v VaultProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessPolicies", v.AccessPolicies)
	populate(objectMap, "createMode", v.CreateMode)
	populate(objectMap, "enablePurgeProtection", v.EnablePurgeProtection)
	populate(objectMap, "enableRbacAuthorization", v.EnableRbacAuthorization)
	populate(objectMap, "enableSoftDelete", v.EnableSoftDelete)
	populate(objectMap, "enabledForDeployment", v.EnabledForDeployment)
	populate(objectMap, "enabledForDiskEncryption", v.EnabledForDiskEncryption)
	populate(objectMap, "enabledForTemplateDeployment", v.EnabledForTemplateDeployment)
	populate(objectMap, "hsmPoolResourceId", v.HsmPoolResourceID)
	populate(objectMap, "networkAcls", v.NetworkACLs)
	populate(objectMap, "privateEndpointConnections", v.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", v.PublicNetworkAccess)
	populate(objectMap, "sku", v.SKU)
	populate(objectMap, "softDeleteRetentionInDays", v.SoftDeleteRetentionInDays)
	populate(objectMap, "tenantId", v.TenantID)
	populate(objectMap, "vaultUri", v.VaultURI)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
