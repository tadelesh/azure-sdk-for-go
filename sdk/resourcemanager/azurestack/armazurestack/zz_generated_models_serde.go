//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Compatibility.
func (c Compatibility) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", c.Description)
	populate(objectMap, "isCompatible", c.IsCompatible)
	populate(objectMap, "issues", c.Issues)
	populate(objectMap, "message", c.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomerSubscriptionList.
func (c CustomerSubscriptionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtendedProductProperties.
func (e ExtendedProductProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "computeRole", e.ComputeRole)
	populate(objectMap, "dataDiskImages", e.DataDiskImages)
	populate(objectMap, "isSystemExtension", e.IsSystemExtension)
	populate(objectMap, "osDiskImage", e.OSDiskImage)
	populate(objectMap, "sourceBlob", e.SourceBlob)
	populate(objectMap, "supportMultipleExtensions", e.SupportMultipleExtensions)
	populate(objectMap, "vmOsType", e.VMOsType)
	populate(objectMap, "vmScaleSetEnabled", e.VMScaleSetEnabled)
	populate(objectMap, "version", e.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkedSubscription.
func (l LinkedSubscription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", l.Etag)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "kind", l.Kind)
	populate(objectMap, "location", l.Location)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkedSubscriptionParameter.
func (l LinkedSubscriptionParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", l.Location)
	populate(objectMap, "properties", l.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkedSubscriptionsList.
func (l LinkedSubscriptionsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProductList.
func (p ProductList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProductNestedProperties.
func (p ProductNestedProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billingPartNumber", p.BillingPartNumber)
	populate(objectMap, "compatibility", p.Compatibility)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "galleryItemIdentity", p.GalleryItemIdentity)
	populate(objectMap, "iconUris", p.IconUris)
	populate(objectMap, "legalTerms", p.LegalTerms)
	populate(objectMap, "links", p.Links)
	populate(objectMap, "offer", p.Offer)
	populate(objectMap, "offerVersion", p.OfferVersion)
	populate(objectMap, "payloadLength", p.PayloadLength)
	populate(objectMap, "privacyPolicy", p.PrivacyPolicy)
	populate(objectMap, "productKind", p.ProductKind)
	populate(objectMap, "productProperties", p.ProductProperties)
	populate(objectMap, "publisherDisplayName", p.PublisherDisplayName)
	populate(objectMap, "publisherIdentifier", p.PublisherIdentifier)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "vmExtensionType", p.VMExtensionType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Registration.
func (r Registration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "kind", r.Kind)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationList.
func (r RegistrationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationParameter.
func (r RegistrationParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", r.Location)
	populate(objectMap, "properties", r.Properties)
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
	populate(objectMap, "etag", t.Etag)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "kind", t.Kind)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachineProductProperties.
func (v VirtualMachineProductProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dataDiskImages", v.DataDiskImages)
	populate(objectMap, "osDiskImage", v.OSDiskImage)
	populate(objectMap, "version", v.Version)
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
