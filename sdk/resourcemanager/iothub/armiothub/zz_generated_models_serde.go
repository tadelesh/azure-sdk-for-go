//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ArmIdentity.
func (a ArmIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", a.PrincipalID)
	populate(objectMap, "tenantId", a.TenantID)
	populate(objectMap, "type", a.Type)
	populate(objectMap, "userAssignedIdentities", a.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CertificateListDescription.
func (c CertificateListDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CertificateProperties.
func (c CertificateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "certificate", c.Certificate)
	populateTimeRFC1123(objectMap, "created", c.Created)
	populateTimeRFC1123(objectMap, "expiry", c.Expiry)
	populate(objectMap, "isVerified", c.IsVerified)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populateTimeRFC1123(objectMap, "updated", c.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateProperties.
func (c *CertificateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = unpopulate(val, &c.Certificate)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC1123(val, &c.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulateTimeRFC1123(val, &c.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, &c.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, &c.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, &c.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeRFC1123(val, &c.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CertificatePropertiesWithNonce.
func (c CertificatePropertiesWithNonce) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "certificate", c.Certificate)
	populateTimeRFC1123(objectMap, "created", c.Created)
	populateTimeRFC1123(objectMap, "expiry", c.Expiry)
	populate(objectMap, "isVerified", c.IsVerified)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populateTimeRFC1123(objectMap, "updated", c.Updated)
	populate(objectMap, "verificationCode", c.VerificationCode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificatePropertiesWithNonce.
func (c *CertificatePropertiesWithNonce) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = unpopulate(val, &c.Certificate)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC1123(val, &c.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulateTimeRFC1123(val, &c.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, &c.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, &c.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, &c.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeRFC1123(val, &c.Updated)
			delete(rawMsg, key)
		case "verificationCode":
			err = unpopulate(val, &c.VerificationCode)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Description.
func (d Description) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "identity", d.Identity)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DescriptionListResult.
func (d DescriptionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointHealthData.
func (e EndpointHealthData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointId", e.EndpointID)
	populate(objectMap, "healthStatus", e.HealthStatus)
	populate(objectMap, "lastKnownError", e.LastKnownError)
	populateTimeRFC1123(objectMap, "lastKnownErrorTime", e.LastKnownErrorTime)
	populateTimeRFC1123(objectMap, "lastSendAttemptTime", e.LastSendAttemptTime)
	populateTimeRFC1123(objectMap, "lastSuccessfulSendAttemptTime", e.LastSuccessfulSendAttemptTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointHealthData.
func (e *EndpointHealthData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endpointId":
			err = unpopulate(val, &e.EndpointID)
			delete(rawMsg, key)
		case "healthStatus":
			err = unpopulate(val, &e.HealthStatus)
			delete(rawMsg, key)
		case "lastKnownError":
			err = unpopulate(val, &e.LastKnownError)
			delete(rawMsg, key)
		case "lastKnownErrorTime":
			err = unpopulateTimeRFC1123(val, &e.LastKnownErrorTime)
			delete(rawMsg, key)
		case "lastSendAttemptTime":
			err = unpopulateTimeRFC1123(val, &e.LastSendAttemptTime)
			delete(rawMsg, key)
		case "lastSuccessfulSendAttemptTime":
			err = unpopulateTimeRFC1123(val, &e.LastSuccessfulSendAttemptTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EndpointHealthDataListResult.
func (e EndpointHealthDataListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EnrichmentProperties.
func (e EnrichmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointNames", e.EndpointNames)
	populate(objectMap, "key", e.Key)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EventHubConsumerGroupInfo.
func (e EventHubConsumerGroupInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", e.Etag)
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EventHubConsumerGroupsListResult.
func (e EventHubConsumerGroupsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EventHubProperties.
func (e EventHubProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpoint", e.Endpoint)
	populate(objectMap, "partitionCount", e.PartitionCount)
	populate(objectMap, "partitionIds", e.PartitionIDs)
	populate(objectMap, "path", e.Path)
	populate(objectMap, "retentionTimeInDays", e.RetentionTimeInDays)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FallbackRouteProperties.
func (f FallbackRouteProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "condition", f.Condition)
	populate(objectMap, "endpointNames", f.EndpointNames)
	populate(objectMap, "isEnabled", f.IsEnabled)
	populate(objectMap, "name", f.Name)
	populate(objectMap, "source", f.Source)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GroupIDInformationProperties.
func (g GroupIDInformationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", g.GroupID)
	populate(objectMap, "requiredMembers", g.RequiredMembers)
	populate(objectMap, "requiredZoneNames", g.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobResponse.
func (j JobResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC1123(objectMap, "endTimeUtc", j.EndTimeUTC)
	populate(objectMap, "failureReason", j.FailureReason)
	populate(objectMap, "jobId", j.JobID)
	populate(objectMap, "parentJobId", j.ParentJobID)
	populateTimeRFC1123(objectMap, "startTimeUtc", j.StartTimeUTC)
	populate(objectMap, "status", j.Status)
	populate(objectMap, "statusMessage", j.StatusMessage)
	populate(objectMap, "type", j.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobResponse.
func (j *JobResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTimeUtc":
			err = unpopulateTimeRFC1123(val, &j.EndTimeUTC)
			delete(rawMsg, key)
		case "failureReason":
			err = unpopulate(val, &j.FailureReason)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, &j.JobID)
			delete(rawMsg, key)
		case "parentJobId":
			err = unpopulate(val, &j.ParentJobID)
			delete(rawMsg, key)
		case "startTimeUtc":
			err = unpopulateTimeRFC1123(val, &j.StartTimeUTC)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &j.Status)
			delete(rawMsg, key)
		case "statusMessage":
			err = unpopulate(val, &j.StatusMessage)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &j.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobResponseListResult.
func (j JobResponseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", j.NextLink)
	populate(objectMap, "value", j.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSetProperties.
func (n NetworkRuleSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applyToBuiltInEventHubEndpoint", n.ApplyToBuiltInEventHubEndpoint)
	populate(objectMap, "defaultAction", n.DefaultAction)
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

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResources.
func (p PrivateLinkResources) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedFqdnList", p.AllowedFqdnList)
	populate(objectMap, "authorizationPolicies", p.AuthorizationPolicies)
	populate(objectMap, "cloudToDevice", p.CloudToDevice)
	populate(objectMap, "comments", p.Comments)
	populate(objectMap, "disableDeviceSAS", p.DisableDeviceSAS)
	populate(objectMap, "disableLocalAuth", p.DisableLocalAuth)
	populate(objectMap, "disableModuleSAS", p.DisableModuleSAS)
	populate(objectMap, "enableDataResidency", p.EnableDataResidency)
	populate(objectMap, "enableFileUploadNotifications", p.EnableFileUploadNotifications)
	populate(objectMap, "eventHubEndpoints", p.EventHubEndpoints)
	populate(objectMap, "features", p.Features)
	populate(objectMap, "hostName", p.HostName)
	populate(objectMap, "ipFilterRules", p.IPFilterRules)
	populate(objectMap, "locations", p.Locations)
	populate(objectMap, "messagingEndpoints", p.MessagingEndpoints)
	populate(objectMap, "minTlsVersion", p.MinTLSVersion)
	populate(objectMap, "networkRuleSets", p.NetworkRuleSets)
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	populate(objectMap, "restrictOutboundNetworkAccess", p.RestrictOutboundNetworkAccess)
	populate(objectMap, "routing", p.Routing)
	populate(objectMap, "state", p.State)
	populate(objectMap, "storageEndpoints", p.StorageEndpoints)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type QuotaMetricInfoListResult.
func (q QuotaMetricInfoListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", q.NextLink)
	populate(objectMap, "value", q.Value)
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

// MarshalJSON implements the json.Marshaller interface for type RouteProperties.
func (r RouteProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "condition", r.Condition)
	populate(objectMap, "endpointNames", r.EndpointNames)
	populate(objectMap, "isEnabled", r.IsEnabled)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "source", r.Source)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RoutingEndpoints.
func (r RoutingEndpoints) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eventHubs", r.EventHubs)
	populate(objectMap, "serviceBusQueues", r.ServiceBusQueues)
	populate(objectMap, "serviceBusTopics", r.ServiceBusTopics)
	populate(objectMap, "storageContainers", r.StorageContainers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RoutingMessage.
func (r RoutingMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "appProperties", r.AppProperties)
	populate(objectMap, "body", r.Body)
	populate(objectMap, "systemProperties", r.SystemProperties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RoutingProperties.
func (r RoutingProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpoints", r.Endpoints)
	populate(objectMap, "enrichments", r.Enrichments)
	populate(objectMap, "fallbackRoute", r.FallbackRoute)
	populate(objectMap, "routes", r.Routes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUDescriptionListResult.
func (s SKUDescriptionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SharedAccessSignatureAuthorizationRuleListResult.
func (s SharedAccessSignatureAuthorizationRuleListResult) MarshalJSON() ([]byte, error) {
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

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TestAllRoutesResult.
func (t TestAllRoutesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "routes", t.Routes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TestRouteResultDetails.
func (t TestRouteResultDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "compilationErrors", t.CompilationErrors)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UserSubscriptionQuotaListResult.
func (u UserSubscriptionQuotaListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
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
