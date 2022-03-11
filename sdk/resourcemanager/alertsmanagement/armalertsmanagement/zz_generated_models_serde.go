//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// GetAction implements the ActionClassification interface for type Action.
func (a *Action) GetAction() *Action { return a }

// GetAction implements the ActionClassification interface for type AddActionGroups.
func (a *AddActionGroups) GetAction() *Action {
	return &Action{
		ActionType: a.ActionType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type AddActionGroups.
func (a AddActionGroups) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionGroupIds", a.ActionGroupIDs)
	objectMap["actionType"] = ActionTypeAddActionGroups
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddActionGroups.
func (a *AddActionGroups) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionGroupIds":
			err = unpopulate(val, &a.ActionGroupIDs)
			delete(rawMsg, key)
		case "actionType":
			err = unpopulate(val, &a.ActionType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AlertModificationProperties.
func (a AlertModificationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alertId", a.AlertID)
	populate(objectMap, "modifications", a.Modifications)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AlertProcessingRule.
func (a AlertProcessingRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AlertProcessingRuleProperties.
func (a AlertProcessingRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", a.Actions)
	populate(objectMap, "conditions", a.Conditions)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "enabled", a.Enabled)
	populate(objectMap, "schedule", a.Schedule)
	populate(objectMap, "scopes", a.Scopes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertProcessingRuleProperties.
func (a *AlertProcessingRuleProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			a.Actions, err = unmarshalActionClassificationArray(val)
			delete(rawMsg, key)
		case "conditions":
			err = unpopulate(val, &a.Conditions)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &a.Description)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, &a.Enabled)
			delete(rawMsg, key)
		case "schedule":
			err = unpopulate(val, &a.Schedule)
			delete(rawMsg, key)
		case "scopes":
			err = unpopulate(val, &a.Scopes)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AlertProcessingRulesList.
func (a AlertProcessingRulesList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AlertsList.
func (a AlertsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AlertsMetaData.
func (a AlertsMetaData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertsMetaData.
func (a *AlertsMetaData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			a.Properties, err = unmarshalAlertsMetaDataPropertiesClassification(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAlertsMetaDataProperties implements the AlertsMetaDataPropertiesClassification interface for type AlertsMetaDataProperties.
func (a *AlertsMetaDataProperties) GetAlertsMetaDataProperties() *AlertsMetaDataProperties { return a }

// MarshalJSON implements the json.Marshaller interface for type AlertsSummaryGroup.
func (a AlertsSummaryGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupedby", a.Groupedby)
	populate(objectMap, "smartGroupsCount", a.SmartGroupsCount)
	populate(objectMap, "total", a.Total)
	populate(objectMap, "values", a.Values)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AlertsSummaryGroupItem.
func (a AlertsSummaryGroupItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", a.Count)
	populate(objectMap, "groupedby", a.Groupedby)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "values", a.Values)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Condition.
func (c Condition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "field", c.Field)
	populate(objectMap, "operator", c.Operator)
	populate(objectMap, "values", c.Values)
	return json.Marshal(objectMap)
}

// GetRecurrence implements the RecurrenceClassification interface for type DailyRecurrence.
func (d *DailyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		RecurrenceType: d.RecurrenceType,
		StartTime:      d.StartTime,
		EndTime:        d.EndTime,
	}
}

// MarshalJSON implements the json.Marshaller interface for type DailyRecurrence.
func (d DailyRecurrence) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endTime", d.EndTime)
	objectMap["recurrenceType"] = RecurrenceTypeDaily
	populate(objectMap, "startTime", d.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DailyRecurrence.
func (d *DailyRecurrence) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulate(val, &d.EndTime)
			delete(rawMsg, key)
		case "recurrenceType":
			err = unpopulate(val, &d.RecurrenceType)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulate(val, &d.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseBody.
func (e ErrorResponseBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseBodyAutoGenerated.
func (e ErrorResponseBodyAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseBodyAutoGenerated2.
func (e ErrorResponseBodyAutoGenerated2) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Essentials.
func (e Essentials) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionStatus", e.ActionStatus)
	populate(objectMap, "alertRule", e.AlertRule)
	populate(objectMap, "alertState", e.AlertState)
	populate(objectMap, "description", e.Description)
	populateTimeRFC3339(objectMap, "lastModifiedDateTime", e.LastModifiedDateTime)
	populate(objectMap, "lastModifiedUserName", e.LastModifiedUserName)
	populate(objectMap, "monitorCondition", e.MonitorCondition)
	populateTimeRFC3339(objectMap, "monitorConditionResolvedDateTime", e.MonitorConditionResolvedDateTime)
	populate(objectMap, "monitorService", e.MonitorService)
	populate(objectMap, "severity", e.Severity)
	populate(objectMap, "signalType", e.SignalType)
	populate(objectMap, "smartGroupId", e.SmartGroupID)
	populate(objectMap, "smartGroupingReason", e.SmartGroupingReason)
	populate(objectMap, "sourceCreatedId", e.SourceCreatedID)
	populateTimeRFC3339(objectMap, "startDateTime", e.StartDateTime)
	populate(objectMap, "targetResource", e.TargetResource)
	populate(objectMap, "targetResourceGroup", e.TargetResourceGroup)
	populate(objectMap, "targetResourceName", e.TargetResourceName)
	populate(objectMap, "targetResourceType", e.TargetResourceType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Essentials.
func (e *Essentials) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionStatus":
			err = unpopulate(val, &e.ActionStatus)
			delete(rawMsg, key)
		case "alertRule":
			err = unpopulate(val, &e.AlertRule)
			delete(rawMsg, key)
		case "alertState":
			err = unpopulate(val, &e.AlertState)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &e.Description)
			delete(rawMsg, key)
		case "lastModifiedDateTime":
			err = unpopulateTimeRFC3339(val, &e.LastModifiedDateTime)
			delete(rawMsg, key)
		case "lastModifiedUserName":
			err = unpopulate(val, &e.LastModifiedUserName)
			delete(rawMsg, key)
		case "monitorCondition":
			err = unpopulate(val, &e.MonitorCondition)
			delete(rawMsg, key)
		case "monitorConditionResolvedDateTime":
			err = unpopulateTimeRFC3339(val, &e.MonitorConditionResolvedDateTime)
			delete(rawMsg, key)
		case "monitorService":
			err = unpopulate(val, &e.MonitorService)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, &e.Severity)
			delete(rawMsg, key)
		case "signalType":
			err = unpopulate(val, &e.SignalType)
			delete(rawMsg, key)
		case "smartGroupId":
			err = unpopulate(val, &e.SmartGroupID)
			delete(rawMsg, key)
		case "smartGroupingReason":
			err = unpopulate(val, &e.SmartGroupingReason)
			delete(rawMsg, key)
		case "sourceCreatedId":
			err = unpopulate(val, &e.SourceCreatedID)
			delete(rawMsg, key)
		case "startDateTime":
			err = unpopulateTimeRFC3339(val, &e.StartDateTime)
			delete(rawMsg, key)
		case "targetResource":
			err = unpopulate(val, &e.TargetResource)
			delete(rawMsg, key)
		case "targetResourceGroup":
			err = unpopulate(val, &e.TargetResourceGroup)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &e.TargetResourceName)
			delete(rawMsg, key)
		case "targetResourceType":
			err = unpopulate(val, &e.TargetResourceType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedResource.
func (m ManagedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// GetAlertsMetaDataProperties implements the AlertsMetaDataPropertiesClassification interface for type MonitorServiceList.
func (m *MonitorServiceList) GetAlertsMetaDataProperties() *AlertsMetaDataProperties {
	return &AlertsMetaDataProperties{
		MetadataIdentifier: m.MetadataIdentifier,
	}
}

// MarshalJSON implements the json.Marshaller interface for type MonitorServiceList.
func (m MonitorServiceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "data", m.Data)
	objectMap["metadataIdentifier"] = MetadataIdentifierMonitorServiceList
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorServiceList.
func (m *MonitorServiceList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, &m.Data)
			delete(rawMsg, key)
		case "metadataIdentifier":
			err = unpopulate(val, &m.MetadataIdentifier)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetRecurrence implements the RecurrenceClassification interface for type MonthlyRecurrence.
func (m *MonthlyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		RecurrenceType: m.RecurrenceType,
		StartTime:      m.StartTime,
		EndTime:        m.EndTime,
	}
}

// MarshalJSON implements the json.Marshaller interface for type MonthlyRecurrence.
func (m MonthlyRecurrence) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "daysOfMonth", m.DaysOfMonth)
	populate(objectMap, "endTime", m.EndTime)
	objectMap["recurrenceType"] = RecurrenceTypeMonthly
	populate(objectMap, "startTime", m.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonthlyRecurrence.
func (m *MonthlyRecurrence) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "daysOfMonth":
			err = unpopulate(val, &m.DaysOfMonth)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulate(val, &m.EndTime)
			delete(rawMsg, key)
		case "recurrenceType":
			err = unpopulate(val, &m.RecurrenceType)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulate(val, &m.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationsList.
func (o OperationsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PatchObject.
func (p PatchObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// GetRecurrence implements the RecurrenceClassification interface for type Recurrence.
func (r *Recurrence) GetRecurrence() *Recurrence { return r }

// GetAction implements the ActionClassification interface for type RemoveAllActionGroups.
func (r *RemoveAllActionGroups) GetAction() *Action {
	return &Action{
		ActionType: r.ActionType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type RemoveAllActionGroups.
func (r RemoveAllActionGroups) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["actionType"] = ActionTypeRemoveAllActionGroups
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RemoveAllActionGroups.
func (r *RemoveAllActionGroups) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionType":
			err = unpopulate(val, &r.ActionType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Schedule.
func (s Schedule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "effectiveFrom", s.EffectiveFrom)
	populate(objectMap, "effectiveUntil", s.EffectiveUntil)
	populate(objectMap, "recurrences", s.Recurrences)
	populate(objectMap, "timeZone", s.TimeZone)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Schedule.
func (s *Schedule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "effectiveFrom":
			err = unpopulate(val, &s.EffectiveFrom)
			delete(rawMsg, key)
		case "effectiveUntil":
			err = unpopulate(val, &s.EffectiveUntil)
			delete(rawMsg, key)
		case "recurrences":
			s.Recurrences, err = unmarshalRecurrenceClassificationArray(val)
			delete(rawMsg, key)
		case "timeZone":
			err = unpopulate(val, &s.TimeZone)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SmartGroupModificationProperties.
func (s SmartGroupModificationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "modifications", s.Modifications)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "smartGroupId", s.SmartGroupID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SmartGroupProperties.
func (s SmartGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alertSeverities", s.AlertSeverities)
	populate(objectMap, "alertStates", s.AlertStates)
	populate(objectMap, "alertsCount", s.AlertsCount)
	populateTimeRFC3339(objectMap, "lastModifiedDateTime", s.LastModifiedDateTime)
	populate(objectMap, "lastModifiedUserName", s.LastModifiedUserName)
	populate(objectMap, "monitorConditions", s.MonitorConditions)
	populate(objectMap, "monitorServices", s.MonitorServices)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "resourceGroups", s.ResourceGroups)
	populate(objectMap, "resourceTypes", s.ResourceTypes)
	populate(objectMap, "resources", s.Resources)
	populate(objectMap, "severity", s.Severity)
	populate(objectMap, "smartGroupState", s.SmartGroupState)
	populateTimeRFC3339(objectMap, "startDateTime", s.StartDateTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartGroupProperties.
func (s *SmartGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alertSeverities":
			err = unpopulate(val, &s.AlertSeverities)
			delete(rawMsg, key)
		case "alertStates":
			err = unpopulate(val, &s.AlertStates)
			delete(rawMsg, key)
		case "alertsCount":
			err = unpopulate(val, &s.AlertsCount)
			delete(rawMsg, key)
		case "lastModifiedDateTime":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedDateTime)
			delete(rawMsg, key)
		case "lastModifiedUserName":
			err = unpopulate(val, &s.LastModifiedUserName)
			delete(rawMsg, key)
		case "monitorConditions":
			err = unpopulate(val, &s.MonitorConditions)
			delete(rawMsg, key)
		case "monitorServices":
			err = unpopulate(val, &s.MonitorServices)
			delete(rawMsg, key)
		case "nextLink":
			err = unpopulate(val, &s.NextLink)
			delete(rawMsg, key)
		case "resourceGroups":
			err = unpopulate(val, &s.ResourceGroups)
			delete(rawMsg, key)
		case "resourceTypes":
			err = unpopulate(val, &s.ResourceTypes)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, &s.Resources)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, &s.Severity)
			delete(rawMsg, key)
		case "smartGroupState":
			err = unpopulate(val, &s.SmartGroupState)
			delete(rawMsg, key)
		case "startDateTime":
			err = unpopulateTimeRFC3339(val, &s.StartDateTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SmartGroupsList.
func (s SmartGroupsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
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

// GetRecurrence implements the RecurrenceClassification interface for type WeeklyRecurrence.
func (w *WeeklyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		RecurrenceType: w.RecurrenceType,
		StartTime:      w.StartTime,
		EndTime:        w.EndTime,
	}
}

// MarshalJSON implements the json.Marshaller interface for type WeeklyRecurrence.
func (w WeeklyRecurrence) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "daysOfWeek", w.DaysOfWeek)
	populate(objectMap, "endTime", w.EndTime)
	objectMap["recurrenceType"] = RecurrenceTypeWeekly
	populate(objectMap, "startTime", w.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WeeklyRecurrence.
func (w *WeeklyRecurrence) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "daysOfWeek":
			err = unpopulate(val, &w.DaysOfWeek)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulate(val, &w.EndTime)
			delete(rawMsg, key)
		case "recurrenceType":
			err = unpopulate(val, &w.RecurrenceType)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulate(val, &w.StartTime)
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
