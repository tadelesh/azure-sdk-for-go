//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListQueryKeysResult.
func (l ListQueryKeysResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSet.
func (n NetworkRuleSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "ipRules", n.IPRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
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

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	populate(objectMap, "shareablePrivateLinkResourceTypes", p.ShareablePrivateLinkResourceTypes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourcesResult.
func (p PrivateLinkResourcesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Service.
func (s Service) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceListResult.
func (s ServiceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceProperties.
func (s ServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hostingMode", s.HostingMode)
	populate(objectMap, "networkRuleSet", s.NetworkRuleSet)
	populate(objectMap, "partitionCount", s.PartitionCount)
	populate(objectMap, "privateEndpointConnections", s.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "replicaCount", s.ReplicaCount)
	populate(objectMap, "sharedPrivateLinkResources", s.SharedPrivateLinkResources)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceUpdate.
func (s ServiceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SharedPrivateLinkResourceListResult.
func (s SharedPrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
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
