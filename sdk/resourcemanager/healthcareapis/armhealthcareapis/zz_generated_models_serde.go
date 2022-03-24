//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DicomService.
func (d DicomService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DicomServiceAuthenticationConfiguration.
func (d DicomServiceAuthenticationConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "audiences", d.Audiences)
	populate(objectMap, "authority", d.Authority)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DicomServiceCollection.
func (d DicomServiceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DicomServicePatchResource.
func (d DicomServicePatchResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirService.
func (f FhirService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", f.Etag)
	populate(objectMap, "id", f.ID)
	populate(objectMap, "identity", f.Identity)
	populate(objectMap, "kind", f.Kind)
	populate(objectMap, "location", f.Location)
	populate(objectMap, "name", f.Name)
	populate(objectMap, "properties", f.Properties)
	populate(objectMap, "systemData", f.SystemData)
	populate(objectMap, "tags", f.Tags)
	populate(objectMap, "type", f.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirServiceAcrConfiguration.
func (f FhirServiceAcrConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "loginServers", f.LoginServers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirServiceCollection.
func (f FhirServiceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirServiceCorsConfiguration.
func (f FhirServiceCorsConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowCredentials", f.AllowCredentials)
	populate(objectMap, "headers", f.Headers)
	populate(objectMap, "maxAge", f.MaxAge)
	populate(objectMap, "methods", f.Methods)
	populate(objectMap, "origins", f.Origins)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirServicePatchResource.
func (f FhirServicePatchResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", f.Identity)
	populate(objectMap, "tags", f.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FhirServiceProperties.
func (f FhirServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessPolicies", f.AccessPolicies)
	populate(objectMap, "acrConfiguration", f.AcrConfiguration)
	populate(objectMap, "authenticationConfiguration", f.AuthenticationConfiguration)
	populate(objectMap, "corsConfiguration", f.CorsConfiguration)
	populate(objectMap, "exportConfiguration", f.ExportConfiguration)
	populate(objectMap, "provisioningState", f.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotConnector.
func (i IotConnector) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", i.Etag)
	populate(objectMap, "id", i.ID)
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "tags", i.Tags)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotConnectorCollection.
func (i IotConnectorCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotConnectorPatchResource.
func (i IotConnectorPatchResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "tags", i.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotFhirDestinationCollection.
func (i IotFhirDestinationCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListOperations.
func (l ListOperations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResultDescription.
func (p PrivateEndpointConnectionListResultDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResultDescription.
func (p PrivateLinkResourceListResultDescription) MarshalJSON() ([]byte, error) {
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

// MarshalJSON implements the json.Marshaller interface for type ResourceTags.
func (r ResourceTags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", r.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceAcrConfigurationInfo.
func (s ServiceAcrConfigurationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "loginServers", s.LoginServers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceCorsConfigurationInfo.
func (s ServiceCorsConfigurationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowCredentials", s.AllowCredentials)
	populate(objectMap, "headers", s.Headers)
	populate(objectMap, "maxAge", s.MaxAge)
	populate(objectMap, "methods", s.Methods)
	populate(objectMap, "origins", s.Origins)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicesDescription.
func (s ServicesDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "kind", s.Kind)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicesDescriptionListResult.
func (s ServicesDescriptionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicesPatchDescription.
func (s ServicesPatchDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicesProperties.
func (s ServicesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessPolicies", s.AccessPolicies)
	populate(objectMap, "acrConfiguration", s.AcrConfiguration)
	populate(objectMap, "authenticationConfiguration", s.AuthenticationConfiguration)
	populate(objectMap, "corsConfiguration", s.CorsConfiguration)
	populate(objectMap, "cosmosDbConfiguration", s.CosmosDbConfiguration)
	populate(objectMap, "exportConfiguration", s.ExportConfiguration)
	populate(objectMap, "privateEndpointConnections", s.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicesResource.
func (s ServicesResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "kind", s.Kind)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
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

// MarshalJSON implements the json.Marshaller interface for type TaggedResource.
func (t TaggedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", t.Etag)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", w.Etag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceList.
func (w WorkspaceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspacePatchResource.
func (w WorkspacePatchResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", w.Tags)
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