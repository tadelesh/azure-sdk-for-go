//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Authorization.
func (a Authorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "delegatedRoleDefinitionIds", a.DelegatedRoleDefinitionIDs)
	populate(objectMap, "principalId", a.PrincipalID)
	populate(objectMap, "principalIdDisplayName", a.PrincipalIDDisplayName)
	populate(objectMap, "roleDefinitionId", a.RoleDefinitionID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDefinition.
func (e ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JustInTimeAccessPolicy.
func (j JustInTimeAccessPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managedByTenantApprovers", j.ManagedByTenantApprovers)
	populate(objectMap, "maximumActivationDuration", j.MaximumActivationDuration)
	populate(objectMap, "multiFactorAuthProvider", j.MultiFactorAuthProvider)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceRegistrationDefinitionList.
func (m MarketplaceRegistrationDefinitionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceRegistrationDefinitionProperties.
func (m MarketplaceRegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", m.Authorizations)
	populate(objectMap, "eligibleAuthorizations", m.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", m.ManagedByTenantID)
	populate(objectMap, "offerDisplayName", m.OfferDisplayName)
	populate(objectMap, "planDisplayName", m.PlanDisplayName)
	populate(objectMap, "publisherDisplayName", m.PublisherDisplayName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentList.
func (r RegistrationAssignmentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentPropertiesRegistrationDefinitionProperties.
func (r RegistrationAssignmentPropertiesRegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", r.Authorizations)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "eligibleAuthorizations", r.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", r.ManagedByTenantID)
	populate(objectMap, "managedByTenantName", r.ManagedByTenantName)
	populate(objectMap, "manageeTenantId", r.ManageeTenantID)
	populate(objectMap, "manageeTenantName", r.ManageeTenantName)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "registrationDefinitionName", r.RegistrationDefinitionName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationDefinitionList.
func (r RegistrationDefinitionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationDefinitionProperties.
func (r RegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", r.Authorizations)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "eligibleAuthorizations", r.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", r.ManagedByTenantID)
	populate(objectMap, "managedByTenantName", r.ManagedByTenantName)
	populate(objectMap, "manageeTenantId", r.ManageeTenantID)
	populate(objectMap, "manageeTenantName", r.ManageeTenantName)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "registrationDefinitionName", r.RegistrationDefinitionName)
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
