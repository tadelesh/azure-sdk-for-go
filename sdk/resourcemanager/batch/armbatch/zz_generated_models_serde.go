//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountCreateParameters.
func (a AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountCreateProperties.
func (a AccountCreateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAuthenticationModes", a.AllowedAuthenticationModes)
	populate(objectMap, "autoStorage", a.AutoStorage)
	populate(objectMap, "encryption", a.Encryption)
	populate(objectMap, "keyVaultReference", a.KeyVaultReference)
	populate(objectMap, "poolAllocationMode", a.PoolAllocationMode)
	populate(objectMap, "publicNetworkAccess", a.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountIdentity.
func (a AccountIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", a.PrincipalID)
	populate(objectMap, "tenantId", a.TenantID)
	populate(objectMap, "type", a.Type)
	populate(objectMap, "userAssignedIdentities", a.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountListResult.
func (a AccountListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountProperties.
func (a AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountEndpoint", a.AccountEndpoint)
	populate(objectMap, "activeJobAndJobScheduleQuota", a.ActiveJobAndJobScheduleQuota)
	populate(objectMap, "allowedAuthenticationModes", a.AllowedAuthenticationModes)
	populate(objectMap, "autoStorage", a.AutoStorage)
	populate(objectMap, "dedicatedCoreQuota", a.DedicatedCoreQuota)
	populate(objectMap, "dedicatedCoreQuotaPerVMFamily", a.DedicatedCoreQuotaPerVMFamily)
	populate(objectMap, "dedicatedCoreQuotaPerVMFamilyEnforced", a.DedicatedCoreQuotaPerVMFamilyEnforced)
	populate(objectMap, "encryption", a.Encryption)
	populate(objectMap, "keyVaultReference", a.KeyVaultReference)
	populate(objectMap, "lowPriorityCoreQuota", a.LowPriorityCoreQuota)
	populate(objectMap, "poolAllocationMode", a.PoolAllocationMode)
	populate(objectMap, "poolQuota", a.PoolQuota)
	populate(objectMap, "privateEndpointConnections", a.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", a.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountUpdateParameters.
func (a AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountUpdateProperties.
func (a AccountUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAuthenticationModes", a.AllowedAuthenticationModes)
	populate(objectMap, "autoStorage", a.AutoStorage)
	populate(objectMap, "encryption", a.Encryption)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Application.
func (a Application) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", a.Etag)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationPackageProperties.
func (a ApplicationPackageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "format", a.Format)
	populateTimeRFC3339(objectMap, "lastActivationTime", a.LastActivationTime)
	populate(objectMap, "state", a.State)
	populate(objectMap, "storageUrl", a.StorageURL)
	populateTimeRFC3339(objectMap, "storageUrlExpiry", a.StorageURLExpiry)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplicationPackageProperties.
func (a *ApplicationPackageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "format":
			err = unpopulate(val, &a.Format)
			delete(rawMsg, key)
		case "lastActivationTime":
			err = unpopulateTimeRFC3339(val, &a.LastActivationTime)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, &a.State)
			delete(rawMsg, key)
		case "storageUrl":
			err = unpopulate(val, &a.StorageURL)
			delete(rawMsg, key)
		case "storageUrlExpiry":
			err = unpopulateTimeRFC3339(val, &a.StorageURLExpiry)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AutoScaleRun.
func (a AutoScaleRun) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", a.Error)
	populateTimeRFC3339(objectMap, "evaluationTime", a.EvaluationTime)
	populate(objectMap, "results", a.Results)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutoScaleRun.
func (a *AutoScaleRun) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, &a.Error)
			delete(rawMsg, key)
		case "evaluationTime":
			err = unpopulateTimeRFC3339(val, &a.EvaluationTime)
			delete(rawMsg, key)
		case "results":
			err = unpopulate(val, &a.Results)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AutoScaleRunError.
func (a AutoScaleRunError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", a.Code)
	populate(objectMap, "details", a.Details)
	populate(objectMap, "message", a.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AutoStorageProperties.
func (a AutoStorageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authenticationMode", a.AuthenticationMode)
	populateTimeRFC3339(objectMap, "lastKeySync", a.LastKeySync)
	populate(objectMap, "nodeIdentityReference", a.NodeIdentityReference)
	populate(objectMap, "storageAccountId", a.StorageAccountID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutoStorageProperties.
func (a *AutoStorageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticationMode":
			err = unpopulate(val, &a.AuthenticationMode)
			delete(rawMsg, key)
		case "lastKeySync":
			err = unpopulateTimeRFC3339(val, &a.LastKeySync)
			delete(rawMsg, key)
		case "nodeIdentityReference":
			err = unpopulate(val, &a.NodeIdentityReference)
			delete(rawMsg, key)
		case "storageAccountId":
			err = unpopulate(val, &a.StorageAccountID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CertificateCreateOrUpdateParameters.
func (c CertificateCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CertificateProperties.
func (c CertificateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "deleteCertificateError", c.DeleteCertificateError)
	populate(objectMap, "format", c.Format)
	populate(objectMap, "previousProvisioningState", c.PreviousProvisioningState)
	populateTimeRFC3339(objectMap, "previousProvisioningStateTransitionTime", c.PreviousProvisioningStateTransitionTime)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populateTimeRFC3339(objectMap, "provisioningStateTransitionTime", c.ProvisioningStateTransitionTime)
	populate(objectMap, "publicData", c.PublicData)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populate(objectMap, "thumbprintAlgorithm", c.ThumbprintAlgorithm)
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
		case "deleteCertificateError":
			err = unpopulate(val, &c.DeleteCertificateError)
			delete(rawMsg, key)
		case "format":
			err = unpopulate(val, &c.Format)
			delete(rawMsg, key)
		case "previousProvisioningState":
			err = unpopulate(val, &c.PreviousProvisioningState)
			delete(rawMsg, key)
		case "previousProvisioningStateTransitionTime":
			err = unpopulateTimeRFC3339(val, &c.PreviousProvisioningStateTransitionTime)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "provisioningStateTransitionTime":
			err = unpopulateTimeRFC3339(val, &c.ProvisioningStateTransitionTime)
			delete(rawMsg, key)
		case "publicData":
			err = unpopulate(val, &c.PublicData)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, &c.Thumbprint)
			delete(rawMsg, key)
		case "thumbprintAlgorithm":
			err = unpopulate(val, &c.ThumbprintAlgorithm)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CertificateReference.
func (c CertificateReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "storeLocation", c.StoreLocation)
	populate(objectMap, "storeName", c.StoreName)
	populate(objectMap, "visibility", c.Visibility)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContainerConfiguration.
func (c ContainerConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "containerImageNames", c.ContainerImageNames)
	populate(objectMap, "containerRegistries", c.ContainerRegistries)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DeleteCertificateError.
func (d DeleteCertificateError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", d.Code)
	populate(objectMap, "details", d.Details)
	populate(objectMap, "message", d.Message)
	populate(objectMap, "target", d.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskEncryptionConfiguration.
func (d DiskEncryptionConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "targets", d.Targets)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointDependency.
func (e EndpointDependency) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", e.Description)
	populate(objectMap, "domainName", e.DomainName)
	populate(objectMap, "endpointDetails", e.EndpointDetails)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InboundNatPool.
func (i InboundNatPool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "backendPort", i.BackendPort)
	populate(objectMap, "frontendPortRangeEnd", i.FrontendPortRangeEnd)
	populate(objectMap, "frontendPortRangeStart", i.FrontendPortRangeStart)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "networkSecurityGroupRules", i.NetworkSecurityGroupRules)
	populate(objectMap, "protocol", i.Protocol)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListApplicationPackagesResult.
func (l ListApplicationPackagesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListApplicationsResult.
func (l ListApplicationsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListCertificatesResult.
func (l ListCertificatesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListPoolsResult.
func (l ListPoolsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListPrivateEndpointConnectionsResult.
func (l ListPrivateEndpointConnectionsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListPrivateLinkResourcesResult.
func (l ListPrivateLinkResourcesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityGroupRule.
func (n NetworkSecurityGroupRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "access", n.Access)
	populate(objectMap, "priority", n.Priority)
	populate(objectMap, "sourceAddressPrefix", n.SourceAddressPrefix)
	populate(objectMap, "sourcePortRanges", n.SourcePortRanges)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpoint.
func (o OutboundEnvironmentEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", o.Category)
	populate(objectMap, "endpoints", o.Endpoints)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpointCollection.
func (o OutboundEnvironmentEndpointCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Pool.
func (p Pool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "identity", p.Identity)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PoolEndpointConfiguration.
func (p PoolEndpointConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inboundNatPools", p.InboundNatPools)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PoolIdentity.
func (p PoolIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "type", p.Type)
	populate(objectMap, "userAssignedIdentities", p.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PoolProperties.
func (p PoolProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocationState", p.AllocationState)
	populateTimeRFC3339(objectMap, "allocationStateTransitionTime", p.AllocationStateTransitionTime)
	populate(objectMap, "applicationLicenses", p.ApplicationLicenses)
	populate(objectMap, "applicationPackages", p.ApplicationPackages)
	populate(objectMap, "autoScaleRun", p.AutoScaleRun)
	populate(objectMap, "certificates", p.Certificates)
	populateTimeRFC3339(objectMap, "creationTime", p.CreationTime)
	populate(objectMap, "currentDedicatedNodes", p.CurrentDedicatedNodes)
	populate(objectMap, "currentLowPriorityNodes", p.CurrentLowPriorityNodes)
	populate(objectMap, "deploymentConfiguration", p.DeploymentConfiguration)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "interNodeCommunication", p.InterNodeCommunication)
	populateTimeRFC3339(objectMap, "lastModified", p.LastModified)
	populate(objectMap, "metadata", p.Metadata)
	populate(objectMap, "mountConfiguration", p.MountConfiguration)
	populate(objectMap, "networkConfiguration", p.NetworkConfiguration)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populateTimeRFC3339(objectMap, "provisioningStateTransitionTime", p.ProvisioningStateTransitionTime)
	populate(objectMap, "resizeOperationStatus", p.ResizeOperationStatus)
	populate(objectMap, "scaleSettings", p.ScaleSettings)
	populate(objectMap, "startTask", p.StartTask)
	populate(objectMap, "taskSchedulingPolicy", p.TaskSchedulingPolicy)
	populate(objectMap, "taskSlotsPerNode", p.TaskSlotsPerNode)
	populate(objectMap, "userAccounts", p.UserAccounts)
	populate(objectMap, "vmSize", p.VMSize)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PoolProperties.
func (p *PoolProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "allocationState":
			err = unpopulate(val, &p.AllocationState)
			delete(rawMsg, key)
		case "allocationStateTransitionTime":
			err = unpopulateTimeRFC3339(val, &p.AllocationStateTransitionTime)
			delete(rawMsg, key)
		case "applicationLicenses":
			err = unpopulate(val, &p.ApplicationLicenses)
			delete(rawMsg, key)
		case "applicationPackages":
			err = unpopulate(val, &p.ApplicationPackages)
			delete(rawMsg, key)
		case "autoScaleRun":
			err = unpopulate(val, &p.AutoScaleRun)
			delete(rawMsg, key)
		case "certificates":
			err = unpopulate(val, &p.Certificates)
			delete(rawMsg, key)
		case "creationTime":
			err = unpopulateTimeRFC3339(val, &p.CreationTime)
			delete(rawMsg, key)
		case "currentDedicatedNodes":
			err = unpopulate(val, &p.CurrentDedicatedNodes)
			delete(rawMsg, key)
		case "currentLowPriorityNodes":
			err = unpopulate(val, &p.CurrentLowPriorityNodes)
			delete(rawMsg, key)
		case "deploymentConfiguration":
			err = unpopulate(val, &p.DeploymentConfiguration)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &p.DisplayName)
			delete(rawMsg, key)
		case "interNodeCommunication":
			err = unpopulate(val, &p.InterNodeCommunication)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, &p.LastModified)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, &p.Metadata)
			delete(rawMsg, key)
		case "mountConfiguration":
			err = unpopulate(val, &p.MountConfiguration)
			delete(rawMsg, key)
		case "networkConfiguration":
			err = unpopulate(val, &p.NetworkConfiguration)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &p.ProvisioningState)
			delete(rawMsg, key)
		case "provisioningStateTransitionTime":
			err = unpopulateTimeRFC3339(val, &p.ProvisioningStateTransitionTime)
			delete(rawMsg, key)
		case "resizeOperationStatus":
			err = unpopulate(val, &p.ResizeOperationStatus)
			delete(rawMsg, key)
		case "scaleSettings":
			err = unpopulate(val, &p.ScaleSettings)
			delete(rawMsg, key)
		case "startTask":
			err = unpopulate(val, &p.StartTask)
			delete(rawMsg, key)
		case "taskSchedulingPolicy":
			err = unpopulate(val, &p.TaskSchedulingPolicy)
			delete(rawMsg, key)
		case "taskSlotsPerNode":
			err = unpopulate(val, &p.TaskSlotsPerNode)
			delete(rawMsg, key)
		case "userAccounts":
			err = unpopulate(val, &p.UserAccounts)
			delete(rawMsg, key)
		case "vmSize":
			err = unpopulate(val, &p.VMSize)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnection.
func (p PrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "type", p.Type)
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

// MarshalJSON implements the json.Marshaller interface for type PublicIPAddressConfiguration.
func (p PublicIPAddressConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "ipAddressIds", p.IPAddressIDs)
	populate(objectMap, "provision", p.Provision)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResizeError.
func (r ResizeError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", r.Code)
	populate(objectMap, "details", r.Details)
	populate(objectMap, "message", r.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResizeOperationStatus.
func (r ResizeOperationStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errors", r.Errors)
	populate(objectMap, "nodeDeallocationOption", r.NodeDeallocationOption)
	populate(objectMap, "resizeTimeout", r.ResizeTimeout)
	populateTimeRFC3339(objectMap, "startTime", r.StartTime)
	populate(objectMap, "targetDedicatedNodes", r.TargetDedicatedNodes)
	populate(objectMap, "targetLowPriorityNodes", r.TargetLowPriorityNodes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResizeOperationStatus.
func (r *ResizeOperationStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errors":
			err = unpopulate(val, &r.Errors)
			delete(rawMsg, key)
		case "nodeDeallocationOption":
			err = unpopulate(val, &r.NodeDeallocationOption)
			delete(rawMsg, key)
		case "resizeTimeout":
			err = unpopulate(val, &r.ResizeTimeout)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &r.StartTime)
			delete(rawMsg, key)
		case "targetDedicatedNodes":
			err = unpopulate(val, &r.TargetDedicatedNodes)
			delete(rawMsg, key)
		case "targetLowPriorityNodes":
			err = unpopulate(val, &r.TargetLowPriorityNodes)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
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

// MarshalJSON implements the json.Marshaller interface for type StartTask.
func (s StartTask) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "commandLine", s.CommandLine)
	populate(objectMap, "containerSettings", s.ContainerSettings)
	populate(objectMap, "environmentSettings", s.EnvironmentSettings)
	populate(objectMap, "maxTaskRetryCount", s.MaxTaskRetryCount)
	populate(objectMap, "resourceFiles", s.ResourceFiles)
	populate(objectMap, "userIdentity", s.UserIdentity)
	populate(objectMap, "waitForSuccess", s.WaitForSuccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SupportedSKU.
func (s SupportedSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", s.Capabilities)
	populate(objectMap, "familyName", s.FamilyName)
	populate(objectMap, "name", s.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SupportedSKUsResult.
func (s SupportedSKUsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VMExtension.
func (v VMExtension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoUpgradeMinorVersion", v.AutoUpgradeMinorVersion)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "protectedSettings", &v.ProtectedSettings)
	populate(objectMap, "provisionAfterExtensions", v.ProvisionAfterExtensions)
	populate(objectMap, "publisher", v.Publisher)
	populate(objectMap, "settings", &v.Settings)
	populate(objectMap, "type", v.Type)
	populate(objectMap, "typeHandlerVersion", v.TypeHandlerVersion)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachineConfiguration.
func (v VirtualMachineConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "containerConfiguration", v.ContainerConfiguration)
	populate(objectMap, "dataDisks", v.DataDisks)
	populate(objectMap, "diskEncryptionConfiguration", v.DiskEncryptionConfiguration)
	populate(objectMap, "extensions", v.Extensions)
	populate(objectMap, "imageReference", v.ImageReference)
	populate(objectMap, "licenseType", v.LicenseType)
	populate(objectMap, "nodeAgentSkuId", v.NodeAgentSKUID)
	populate(objectMap, "nodePlacementConfiguration", v.NodePlacementConfiguration)
	populate(objectMap, "osDisk", v.OSDisk)
	populate(objectMap, "windowsConfiguration", v.WindowsConfiguration)
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
