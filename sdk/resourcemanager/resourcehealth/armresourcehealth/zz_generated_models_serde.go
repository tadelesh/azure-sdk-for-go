//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AvailabilityStatusListResult.
func (a AvailabilityStatusListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilityStatusProperties.
func (a AvailabilityStatusProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityState", a.AvailabilityState)
	populate(objectMap, "detailedStatus", a.DetailedStatus)
	populate(objectMap, "healthEventCategory", a.HealthEventCategory)
	populate(objectMap, "healthEventCause", a.HealthEventCause)
	populate(objectMap, "healthEventId", a.HealthEventID)
	populate(objectMap, "healthEventType", a.HealthEventType)
	populateTimeRFC3339(objectMap, "occuredTime", a.OccuredTime)
	populate(objectMap, "reasonChronicity", a.ReasonChronicity)
	populate(objectMap, "reasonType", a.ReasonType)
	populate(objectMap, "recentlyResolvedState", a.RecentlyResolvedState)
	populate(objectMap, "recommendedActions", a.RecommendedActions)
	populateTimeRFC3339(objectMap, "reportedTime", a.ReportedTime)
	populateTimeRFC3339(objectMap, "resolutionETA", a.ResolutionETA)
	populateTimeRFC3339(objectMap, "rootCauseAttributionTime", a.RootCauseAttributionTime)
	populate(objectMap, "serviceImpactingEvents", a.ServiceImpactingEvents)
	populate(objectMap, "summary", a.Summary)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AvailabilityStatusProperties.
func (a *AvailabilityStatusProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "availabilityState":
			err = unpopulate(val, &a.AvailabilityState)
			delete(rawMsg, key)
		case "detailedStatus":
			err = unpopulate(val, &a.DetailedStatus)
			delete(rawMsg, key)
		case "healthEventCategory":
			err = unpopulate(val, &a.HealthEventCategory)
			delete(rawMsg, key)
		case "healthEventCause":
			err = unpopulate(val, &a.HealthEventCause)
			delete(rawMsg, key)
		case "healthEventId":
			err = unpopulate(val, &a.HealthEventID)
			delete(rawMsg, key)
		case "healthEventType":
			err = unpopulate(val, &a.HealthEventType)
			delete(rawMsg, key)
		case "occuredTime":
			err = unpopulateTimeRFC3339(val, &a.OccuredTime)
			delete(rawMsg, key)
		case "reasonChronicity":
			err = unpopulate(val, &a.ReasonChronicity)
			delete(rawMsg, key)
		case "reasonType":
			err = unpopulate(val, &a.ReasonType)
			delete(rawMsg, key)
		case "recentlyResolvedState":
			err = unpopulate(val, &a.RecentlyResolvedState)
			delete(rawMsg, key)
		case "recommendedActions":
			err = unpopulate(val, &a.RecommendedActions)
			delete(rawMsg, key)
		case "reportedTime":
			err = unpopulateTimeRFC3339(val, &a.ReportedTime)
			delete(rawMsg, key)
		case "resolutionETA":
			err = unpopulateTimeRFC3339(val, &a.ResolutionETA)
			delete(rawMsg, key)
		case "rootCauseAttributionTime":
			err = unpopulateTimeRFC3339(val, &a.RootCauseAttributionTime)
			delete(rawMsg, key)
		case "serviceImpactingEvents":
			err = unpopulate(val, &a.ServiceImpactingEvents)
			delete(rawMsg, key)
		case "summary":
			err = unpopulate(val, &a.Summary)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilityStatusPropertiesRecentlyResolvedState.
func (a AvailabilityStatusPropertiesRecentlyResolvedState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "resolvedTime", a.ResolvedTime)
	populate(objectMap, "unavailabilitySummary", a.UnavailabilitySummary)
	populateTimeRFC3339(objectMap, "unavailableOccurredTime", a.UnavailableOccurredTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AvailabilityStatusPropertiesRecentlyResolvedState.
func (a *AvailabilityStatusPropertiesRecentlyResolvedState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resolvedTime":
			err = unpopulateTimeRFC3339(val, &a.ResolvedTime)
			delete(rawMsg, key)
		case "unavailabilitySummary":
			err = unpopulate(val, &a.UnavailabilitySummary)
			delete(rawMsg, key)
		case "unavailableOccurredTime":
			err = unpopulateTimeRFC3339(val, &a.UnavailableOccurredTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmergingIssue.
func (e EmergingIssue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "refreshTimestamp", e.RefreshTimestamp)
	populate(objectMap, "statusActiveEvents", e.StatusActiveEvents)
	populate(objectMap, "statusBanners", e.StatusBanners)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EmergingIssue.
func (e *EmergingIssue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "refreshTimestamp":
			err = unpopulateTimeRFC3339(val, &e.RefreshTimestamp)
			delete(rawMsg, key)
		case "statusActiveEvents":
			err = unpopulate(val, &e.StatusActiveEvents)
			delete(rawMsg, key)
		case "statusBanners":
			err = unpopulate(val, &e.StatusBanners)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EmergingIssueImpact.
func (e EmergingIssueImpact) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "regions", e.Regions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EmergingIssueListResult.
func (e EmergingIssueListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceImpactingEvent.
func (s ServiceImpactingEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "correlationId", s.CorrelationID)
	populateTimeRFC3339(objectMap, "eventStartTime", s.EventStartTime)
	populateTimeRFC3339(objectMap, "eventStatusLastModifiedTime", s.EventStatusLastModifiedTime)
	populate(objectMap, "incidentProperties", s.IncidentProperties)
	populate(objectMap, "status", s.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceImpactingEvent.
func (s *ServiceImpactingEvent) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "correlationId":
			err = unpopulate(val, &s.CorrelationID)
			delete(rawMsg, key)
		case "eventStartTime":
			err = unpopulateTimeRFC3339(val, &s.EventStartTime)
			delete(rawMsg, key)
		case "eventStatusLastModifiedTime":
			err = unpopulateTimeRFC3339(val, &s.EventStatusLastModifiedTime)
			delete(rawMsg, key)
		case "incidentProperties":
			err = unpopulate(val, &s.IncidentProperties)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatusActiveEvent.
func (s StatusActiveEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "cloud", s.Cloud)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "impacts", s.Impacts)
	populateTimeRFC3339(objectMap, "lastModifiedTime", s.LastModifiedTime)
	populate(objectMap, "published", s.Published)
	populate(objectMap, "severity", s.Severity)
	populate(objectMap, "stage", s.Stage)
	populateTimeRFC3339(objectMap, "startTime", s.StartTime)
	populate(objectMap, "title", s.Title)
	populate(objectMap, "trackingId", s.TrackingID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatusActiveEvent.
func (s *StatusActiveEvent) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cloud":
			err = unpopulate(val, &s.Cloud)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &s.Description)
			delete(rawMsg, key)
		case "impacts":
			err = unpopulate(val, &s.Impacts)
			delete(rawMsg, key)
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedTime)
			delete(rawMsg, key)
		case "published":
			err = unpopulate(val, &s.Published)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, &s.Severity)
			delete(rawMsg, key)
		case "stage":
			err = unpopulate(val, &s.Stage)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &s.StartTime)
			delete(rawMsg, key)
		case "title":
			err = unpopulate(val, &s.Title)
			delete(rawMsg, key)
		case "trackingId":
			err = unpopulate(val, &s.TrackingID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatusBanner.
func (s StatusBanner) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "cloud", s.Cloud)
	populateTimeRFC3339(objectMap, "lastModifiedTime", s.LastModifiedTime)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "title", s.Title)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatusBanner.
func (s *StatusBanner) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cloud":
			err = unpopulate(val, &s.Cloud)
			delete(rawMsg, key)
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &s.Message)
			delete(rawMsg, key)
		case "title":
			err = unpopulate(val, &s.Title)
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