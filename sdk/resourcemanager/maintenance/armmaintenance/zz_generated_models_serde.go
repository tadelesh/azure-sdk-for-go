//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ApplyUpdateProperties.
func (a ApplyUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "lastUpdateTime", a.LastUpdateTime)
	populate(objectMap, "resourceId", a.ResourceID)
	populate(objectMap, "status", a.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplyUpdateProperties.
func (a *ApplyUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastUpdateTime":
			err = unpopulateTimeRFC3339(val, &a.LastUpdateTime)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, &a.ResourceID)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &a.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Configuration.
func (c Configuration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationProperties.
func (c ConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extensionProperties", c.ExtensionProperties)
	populate(objectMap, "installPatches", c.InstallPatches)
	populate(objectMap, "maintenanceScope", c.MaintenanceScope)
	populate(objectMap, "maintenanceWindow", c.MaintenanceWindow)
	populate(objectMap, "namespace", c.Namespace)
	populate(objectMap, "visibility", c.Visibility)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InputLinuxParameters.
func (i InputLinuxParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "classificationsToInclude", i.ClassificationsToInclude)
	populate(objectMap, "packageNameMasksToExclude", i.PackageNameMasksToExclude)
	populate(objectMap, "packageNameMasksToInclude", i.PackageNameMasksToInclude)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InputWindowsParameters.
func (i InputWindowsParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "classificationsToInclude", i.ClassificationsToInclude)
	populate(objectMap, "excludeKbsRequiringReboot", i.ExcludeKbsRequiringReboot)
	populate(objectMap, "kbNumbersToExclude", i.KbNumbersToExclude)
	populate(objectMap, "kbNumbersToInclude", i.KbNumbersToInclude)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListApplyUpdate.
func (l ListApplyUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListConfigurationAssignmentsResult.
func (l ListConfigurationAssignmentsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListMaintenanceConfigurationsResult.
func (l ListMaintenanceConfigurationsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListUpdatesResult.
func (l ListUpdatesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationsListResult.
func (o OperationsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SoftwareUpdateConfigurationTasks.
func (s SoftwareUpdateConfigurationTasks) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "postTasks", s.PostTasks)
	populate(objectMap, "preTasks", s.PreTasks)
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

// MarshalJSON implements the json.Marshaller interface for type TaskProperties.
func (t TaskProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "parameters", t.Parameters)
	populate(objectMap, "source", t.Source)
	populate(objectMap, "taskScope", t.TaskScope)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Update.
func (u Update) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "impactDurationInSec", u.ImpactDurationInSec)
	populate(objectMap, "impactType", u.ImpactType)
	populate(objectMap, "maintenanceScope", u.MaintenanceScope)
	populateTimeRFC3339(objectMap, "notBefore", u.NotBefore)
	populate(objectMap, "properties", u.Properties)
	populate(objectMap, "status", u.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Update.
func (u *Update) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "impactDurationInSec":
			err = unpopulate(val, &u.ImpactDurationInSec)
			delete(rawMsg, key)
		case "impactType":
			err = unpopulate(val, &u.ImpactType)
			delete(rawMsg, key)
		case "maintenanceScope":
			err = unpopulate(val, &u.MaintenanceScope)
			delete(rawMsg, key)
		case "notBefore":
			err = unpopulateTimeRFC3339(val, &u.NotBefore)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &u.Properties)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &u.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
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
