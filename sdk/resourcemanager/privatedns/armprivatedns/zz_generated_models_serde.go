//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

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

// MarshalJSON implements the json.Marshaller interface for type PrivateZone.
func (p PrivateZone) MarshalJSON() ([]byte, error) {
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

// MarshalJSON implements the json.Marshaller interface for type PrivateZoneListResult.
func (p PrivateZoneListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecordSet.
func (r RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecordSetListResult.
func (r RecordSetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecordSetProperties.
func (r RecordSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aRecords", r.ARecords)
	populate(objectMap, "aaaaRecords", r.AaaaRecords)
	populate(objectMap, "cnameRecord", r.CnameRecord)
	populate(objectMap, "fqdn", r.Fqdn)
	populate(objectMap, "isAutoRegistered", r.IsAutoRegistered)
	populate(objectMap, "metadata", r.Metadata)
	populate(objectMap, "mxRecords", r.MxRecords)
	populate(objectMap, "ptrRecords", r.PtrRecords)
	populate(objectMap, "soaRecord", r.SoaRecord)
	populate(objectMap, "srvRecords", r.SrvRecords)
	populate(objectMap, "ttl", r.TTL)
	populate(objectMap, "txtRecords", r.TxtRecords)
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

// MarshalJSON implements the json.Marshaller interface for type TxtRecord.
func (t TxtRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLink.
func (v VirtualNetworkLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", v.Etag)
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkListResult.
func (v VirtualNetworkLinkListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
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
