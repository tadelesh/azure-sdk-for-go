//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type APIKey.
func (a APIKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectionString", a.ConnectionString)
	populate(objectMap, "id", a.ID)
	populateTimeRFC3339(objectMap, "lastModified", a.LastModified)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "readOnly", a.ReadOnly)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type APIKey.
func (a *APIKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionString":
			err = unpopulate(val, &a.ConnectionString)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &a.ID)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, &a.LastModified)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &a.Name)
			delete(rawMsg, key)
		case "readOnly":
			err = unpopulate(val, &a.ReadOnly)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &a.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type APIKeyListResult.
func (a APIKeyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStore.
func (c ConfigurationStore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreListResult.
func (c ConfigurationStoreListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreProperties.
func (c ConfigurationStoreProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "createMode", c.CreateMode)
	populateTimeRFC3339(objectMap, "creationDate", c.CreationDate)
	populate(objectMap, "disableLocalAuth", c.DisableLocalAuth)
	populate(objectMap, "enablePurgeProtection", c.EnablePurgeProtection)
	populate(objectMap, "encryption", c.Encryption)
	populate(objectMap, "endpoint", c.Endpoint)
	populate(objectMap, "privateEndpointConnections", c.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", c.PublicNetworkAccess)
	populate(objectMap, "softDeleteRetentionInDays", c.SoftDeleteRetentionInDays)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigurationStoreProperties.
func (c *ConfigurationStoreProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createMode":
			err = unpopulate(val, &c.CreateMode)
			delete(rawMsg, key)
		case "creationDate":
			err = unpopulateTimeRFC3339(val, &c.CreationDate)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, &c.DisableLocalAuth)
			delete(rawMsg, key)
		case "enablePurgeProtection":
			err = unpopulate(val, &c.EnablePurgeProtection)
			delete(rawMsg, key)
		case "encryption":
			err = unpopulate(val, &c.Encryption)
			delete(rawMsg, key)
		case "endpoint":
			err = unpopulate(val, &c.Endpoint)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &c.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &c.PublicNetworkAccess)
			delete(rawMsg, key)
		case "softDeleteRetentionInDays":
			err = unpopulate(val, &c.SoftDeleteRetentionInDays)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreUpdateParameters.
func (c ConfigurationStoreUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DeletedConfigurationStoreListResult.
func (d DeletedConfigurationStoreListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DeletedConfigurationStoreProperties.
func (d DeletedConfigurationStoreProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configurationStoreId", d.ConfigurationStoreID)
	populateTimeRFC3339(objectMap, "deletionDate", d.DeletionDate)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "purgeProtectionEnabled", d.PurgeProtectionEnabled)
	populateTimeRFC3339(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedConfigurationStoreProperties.
func (d *DeletedConfigurationStoreProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "configurationStoreId":
			err = unpopulate(val, &d.ConfigurationStoreID)
			delete(rawMsg, key)
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
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetails.
func (e ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyValueListResult.
func (k KeyValueListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", k.NextLink)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyValueProperties.
func (k KeyValueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", k.ContentType)
	populate(objectMap, "eTag", k.ETag)
	populate(objectMap, "key", k.Key)
	populate(objectMap, "label", k.Label)
	populateTimeRFC3339(objectMap, "lastModified", k.LastModified)
	populate(objectMap, "locked", k.Locked)
	populate(objectMap, "tags", k.Tags)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyValueProperties.
func (k *KeyValueProperties) UnmarshalJSON(data []byte) error {
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
		case "eTag":
			err = unpopulate(val, &k.ETag)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, &k.Key)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &k.Label)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, &k.LastModified)
			delete(rawMsg, key)
		case "locked":
			err = unpopulate(val, &k.Locked)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &k.Tags)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
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
	populate(objectMap, "name", m.Name)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationDefinitionListResult.
func (o OperationDefinitionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
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

// MarshalJSON implements the json.Marshaller interface for type ResourceIdentity.
func (r ResourceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", r.PrincipalID)
	populate(objectMap, "tenantId", r.TenantID)
	populate(objectMap, "type", r.Type)
	populate(objectMap, "userAssignedIdentities", r.UserAssignedIdentities)
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

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
