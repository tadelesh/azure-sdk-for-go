//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AssessedMachineProperties.
func (a AssessedMachineProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bootType", a.BootType)
	populate(objectMap, "confidenceRatingInPercentage", a.ConfidenceRatingInPercentage)
	populateTimeRFC3339(objectMap, "createdTimestamp", a.CreatedTimestamp)
	populate(objectMap, "datacenterMachineArmId", a.DatacenterMachineArmID)
	populate(objectMap, "datacenterManagementServerArmId", a.DatacenterManagementServerArmID)
	populate(objectMap, "datacenterManagementServerName", a.DatacenterManagementServerName)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "disks", a.Disks)
	populate(objectMap, "displayName", a.DisplayName)
	populate(objectMap, "megabytesOfMemory", a.MegabytesOfMemory)
	populate(objectMap, "megabytesOfMemoryForRecommendedSize", a.MegabytesOfMemoryForRecommendedSize)
	populate(objectMap, "monthlyBandwidthCost", a.MonthlyBandwidthCost)
	populate(objectMap, "monthlyComputeCostForRecommendedSize", a.MonthlyComputeCostForRecommendedSize)
	populate(objectMap, "monthlyPremiumStorageCost", a.MonthlyPremiumStorageCost)
	populate(objectMap, "monthlyStandardSSDStorageCost", a.MonthlyStandardSSDStorageCost)
	populate(objectMap, "monthlyStorageCost", a.MonthlyStorageCost)
	populate(objectMap, "networkAdapters", a.NetworkAdapters)
	populate(objectMap, "numberOfCores", a.NumberOfCores)
	populate(objectMap, "numberOfCoresForRecommendedSize", a.NumberOfCoresForRecommendedSize)
	populate(objectMap, "operatingSystemName", a.OperatingSystemName)
	populate(objectMap, "operatingSystemType", a.OperatingSystemType)
	populate(objectMap, "operatingSystemVersion", a.OperatingSystemVersion)
	populate(objectMap, "percentageCoresUtilization", a.PercentageCoresUtilization)
	populate(objectMap, "percentageMemoryUtilization", a.PercentageMemoryUtilization)
	populate(objectMap, "recommendedSize", a.RecommendedSize)
	populate(objectMap, "suitability", a.Suitability)
	populate(objectMap, "suitabilityDetail", a.SuitabilityDetail)
	populate(objectMap, "suitabilityExplanation", a.SuitabilityExplanation)
	populateTimeRFC3339(objectMap, "updatedTimestamp", a.UpdatedTimestamp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssessedMachineProperties.
func (a *AssessedMachineProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bootType":
			err = unpopulate(val, &a.BootType)
			delete(rawMsg, key)
		case "confidenceRatingInPercentage":
			err = unpopulate(val, &a.ConfidenceRatingInPercentage)
			delete(rawMsg, key)
		case "createdTimestamp":
			err = unpopulateTimeRFC3339(val, &a.CreatedTimestamp)
			delete(rawMsg, key)
		case "datacenterMachineArmId":
			err = unpopulate(val, &a.DatacenterMachineArmID)
			delete(rawMsg, key)
		case "datacenterManagementServerArmId":
			err = unpopulate(val, &a.DatacenterManagementServerArmID)
			delete(rawMsg, key)
		case "datacenterManagementServerName":
			err = unpopulate(val, &a.DatacenterManagementServerName)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &a.Description)
			delete(rawMsg, key)
		case "disks":
			err = unpopulate(val, &a.Disks)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &a.DisplayName)
			delete(rawMsg, key)
		case "megabytesOfMemory":
			err = unpopulate(val, &a.MegabytesOfMemory)
			delete(rawMsg, key)
		case "megabytesOfMemoryForRecommendedSize":
			err = unpopulate(val, &a.MegabytesOfMemoryForRecommendedSize)
			delete(rawMsg, key)
		case "monthlyBandwidthCost":
			err = unpopulate(val, &a.MonthlyBandwidthCost)
			delete(rawMsg, key)
		case "monthlyComputeCostForRecommendedSize":
			err = unpopulate(val, &a.MonthlyComputeCostForRecommendedSize)
			delete(rawMsg, key)
		case "monthlyPremiumStorageCost":
			err = unpopulate(val, &a.MonthlyPremiumStorageCost)
			delete(rawMsg, key)
		case "monthlyStandardSSDStorageCost":
			err = unpopulate(val, &a.MonthlyStandardSSDStorageCost)
			delete(rawMsg, key)
		case "monthlyStorageCost":
			err = unpopulate(val, &a.MonthlyStorageCost)
			delete(rawMsg, key)
		case "networkAdapters":
			err = unpopulate(val, &a.NetworkAdapters)
			delete(rawMsg, key)
		case "numberOfCores":
			err = unpopulate(val, &a.NumberOfCores)
			delete(rawMsg, key)
		case "numberOfCoresForRecommendedSize":
			err = unpopulate(val, &a.NumberOfCoresForRecommendedSize)
			delete(rawMsg, key)
		case "operatingSystemName":
			err = unpopulate(val, &a.OperatingSystemName)
			delete(rawMsg, key)
		case "operatingSystemType":
			err = unpopulate(val, &a.OperatingSystemType)
			delete(rawMsg, key)
		case "operatingSystemVersion":
			err = unpopulate(val, &a.OperatingSystemVersion)
			delete(rawMsg, key)
		case "percentageCoresUtilization":
			err = unpopulate(val, &a.PercentageCoresUtilization)
			delete(rawMsg, key)
		case "percentageMemoryUtilization":
			err = unpopulate(val, &a.PercentageMemoryUtilization)
			delete(rawMsg, key)
		case "recommendedSize":
			err = unpopulate(val, &a.RecommendedSize)
			delete(rawMsg, key)
		case "suitability":
			err = unpopulate(val, &a.Suitability)
			delete(rawMsg, key)
		case "suitabilityDetail":
			err = unpopulate(val, &a.SuitabilityDetail)
			delete(rawMsg, key)
		case "suitabilityExplanation":
			err = unpopulate(val, &a.SuitabilityExplanation)
			delete(rawMsg, key)
		case "updatedTimestamp":
			err = unpopulateTimeRFC3339(val, &a.UpdatedTimestamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssessedMachineResultList.
func (a AssessedMachineResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssessedNetworkAdapter.
func (a AssessedNetworkAdapter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", a.DisplayName)
	populate(objectMap, "ipAddresses", a.IPAddresses)
	populate(objectMap, "macAddress", a.MacAddress)
	populate(objectMap, "megabytesPerSecondReceived", a.MegabytesPerSecondReceived)
	populate(objectMap, "megabytesPerSecondTransmitted", a.MegabytesPerSecondTransmitted)
	populate(objectMap, "monthlyBandwidthCosts", a.MonthlyBandwidthCosts)
	populate(objectMap, "netGigabytesTransmittedPerMonth", a.NetGigabytesTransmittedPerMonth)
	populate(objectMap, "suitability", a.Suitability)
	populate(objectMap, "suitabilityDetail", a.SuitabilityDetail)
	populate(objectMap, "suitabilityExplanation", a.SuitabilityExplanation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssessmentOptionsProperties.
func (a AssessmentOptionsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "reservedInstanceSupportedCurrencies", a.ReservedInstanceSupportedCurrencies)
	populate(objectMap, "reservedInstanceSupportedLocations", a.ReservedInstanceSupportedLocations)
	populate(objectMap, "reservedInstanceSupportedOffers", a.ReservedInstanceSupportedOffers)
	populate(objectMap, "reservedInstanceVmFamilies", a.ReservedInstanceVMFamilies)
	populate(objectMap, "vmFamilies", a.VMFamilies)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssessmentOptionsResultList.
func (a AssessmentOptionsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssessmentProperties.
func (a AssessmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "azureDiskType", a.AzureDiskType)
	populate(objectMap, "azureHybridUseBenefit", a.AzureHybridUseBenefit)
	populate(objectMap, "azureLocation", a.AzureLocation)
	populate(objectMap, "azureOfferCode", a.AzureOfferCode)
	populate(objectMap, "azurePricingTier", a.AzurePricingTier)
	populate(objectMap, "azureStorageRedundancy", a.AzureStorageRedundancy)
	populate(objectMap, "azureVmFamilies", a.AzureVMFamilies)
	populate(objectMap, "confidenceRatingInPercentage", a.ConfidenceRatingInPercentage)
	populateTimeRFC3339(objectMap, "createdTimestamp", a.CreatedTimestamp)
	populate(objectMap, "currency", a.Currency)
	populate(objectMap, "discountPercentage", a.DiscountPercentage)
	populate(objectMap, "eaSubscriptionId", a.EaSubscriptionID)
	populate(objectMap, "monthlyBandwidthCost", a.MonthlyBandwidthCost)
	populate(objectMap, "monthlyComputeCost", a.MonthlyComputeCost)
	populate(objectMap, "monthlyPremiumStorageCost", a.MonthlyPremiumStorageCost)
	populate(objectMap, "monthlyStandardSSDStorageCost", a.MonthlyStandardSSDStorageCost)
	populate(objectMap, "monthlyStorageCost", a.MonthlyStorageCost)
	populate(objectMap, "numberOfMachines", a.NumberOfMachines)
	populate(objectMap, "percentile", a.Percentile)
	populateTimeRFC3339(objectMap, "perfDataEndTime", a.PerfDataEndTime)
	populateTimeRFC3339(objectMap, "perfDataStartTime", a.PerfDataStartTime)
	populateTimeRFC3339(objectMap, "pricesTimestamp", a.PricesTimestamp)
	populate(objectMap, "reservedInstance", a.ReservedInstance)
	populate(objectMap, "scalingFactor", a.ScalingFactor)
	populate(objectMap, "sizingCriterion", a.SizingCriterion)
	populate(objectMap, "stage", a.Stage)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "timeRange", a.TimeRange)
	populateTimeRFC3339(objectMap, "updatedTimestamp", a.UpdatedTimestamp)
	populate(objectMap, "vmUptime", a.VMUptime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssessmentProperties.
func (a *AssessmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureDiskType":
			err = unpopulate(val, &a.AzureDiskType)
			delete(rawMsg, key)
		case "azureHybridUseBenefit":
			err = unpopulate(val, &a.AzureHybridUseBenefit)
			delete(rawMsg, key)
		case "azureLocation":
			err = unpopulate(val, &a.AzureLocation)
			delete(rawMsg, key)
		case "azureOfferCode":
			err = unpopulate(val, &a.AzureOfferCode)
			delete(rawMsg, key)
		case "azurePricingTier":
			err = unpopulate(val, &a.AzurePricingTier)
			delete(rawMsg, key)
		case "azureStorageRedundancy":
			err = unpopulate(val, &a.AzureStorageRedundancy)
			delete(rawMsg, key)
		case "azureVmFamilies":
			err = unpopulate(val, &a.AzureVMFamilies)
			delete(rawMsg, key)
		case "confidenceRatingInPercentage":
			err = unpopulate(val, &a.ConfidenceRatingInPercentage)
			delete(rawMsg, key)
		case "createdTimestamp":
			err = unpopulateTimeRFC3339(val, &a.CreatedTimestamp)
			delete(rawMsg, key)
		case "currency":
			err = unpopulate(val, &a.Currency)
			delete(rawMsg, key)
		case "discountPercentage":
			err = unpopulate(val, &a.DiscountPercentage)
			delete(rawMsg, key)
		case "eaSubscriptionId":
			err = unpopulate(val, &a.EaSubscriptionID)
			delete(rawMsg, key)
		case "monthlyBandwidthCost":
			err = unpopulate(val, &a.MonthlyBandwidthCost)
			delete(rawMsg, key)
		case "monthlyComputeCost":
			err = unpopulate(val, &a.MonthlyComputeCost)
			delete(rawMsg, key)
		case "monthlyPremiumStorageCost":
			err = unpopulate(val, &a.MonthlyPremiumStorageCost)
			delete(rawMsg, key)
		case "monthlyStandardSSDStorageCost":
			err = unpopulate(val, &a.MonthlyStandardSSDStorageCost)
			delete(rawMsg, key)
		case "monthlyStorageCost":
			err = unpopulate(val, &a.MonthlyStorageCost)
			delete(rawMsg, key)
		case "numberOfMachines":
			err = unpopulate(val, &a.NumberOfMachines)
			delete(rawMsg, key)
		case "percentile":
			err = unpopulate(val, &a.Percentile)
			delete(rawMsg, key)
		case "perfDataEndTime":
			err = unpopulateTimeRFC3339(val, &a.PerfDataEndTime)
			delete(rawMsg, key)
		case "perfDataStartTime":
			err = unpopulateTimeRFC3339(val, &a.PerfDataStartTime)
			delete(rawMsg, key)
		case "pricesTimestamp":
			err = unpopulateTimeRFC3339(val, &a.PricesTimestamp)
			delete(rawMsg, key)
		case "reservedInstance":
			err = unpopulate(val, &a.ReservedInstance)
			delete(rawMsg, key)
		case "scalingFactor":
			err = unpopulate(val, &a.ScalingFactor)
			delete(rawMsg, key)
		case "sizingCriterion":
			err = unpopulate(val, &a.SizingCriterion)
			delete(rawMsg, key)
		case "stage":
			err = unpopulate(val, &a.Stage)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &a.Status)
			delete(rawMsg, key)
		case "timeRange":
			err = unpopulate(val, &a.TimeRange)
			delete(rawMsg, key)
		case "updatedTimestamp":
			err = unpopulateTimeRFC3339(val, &a.UpdatedTimestamp)
			delete(rawMsg, key)
		case "vmUptime":
			err = unpopulate(val, &a.VMUptime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssessmentResultList.
func (a AssessmentResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
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

// MarshalJSON implements the json.Marshaller interface for type CollectorAgentProperties.
func (c CollectorAgentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populateTimeRFC3339(objectMap, "lastHeartbeatUtc", c.LastHeartbeatUTC)
	populate(objectMap, "spnDetails", c.SpnDetails)
	populate(objectMap, "version", c.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CollectorAgentProperties.
func (c *CollectorAgentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &c.ID)
			delete(rawMsg, key)
		case "lastHeartbeatUtc":
			err = unpopulateTimeRFC3339(val, &c.LastHeartbeatUTC)
			delete(rawMsg, key)
		case "spnDetails":
			err = unpopulate(val, &c.SpnDetails)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &c.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DownloadURL.
func (d DownloadURL) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assessmentReportUrl", d.AssessmentReportURL)
	populateTimeRFC3339(objectMap, "expirationTime", d.ExpirationTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DownloadURL.
func (d *DownloadURL) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assessmentReportUrl":
			err = unpopulate(val, &d.AssessmentReportURL)
			delete(rawMsg, key)
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, &d.ExpirationTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GroupBodyProperties.
func (g GroupBodyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "machines", g.Machines)
	populate(objectMap, "operationType", g.OperationType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GroupProperties.
func (g GroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "areAssessmentsRunning", g.AreAssessmentsRunning)
	populate(objectMap, "assessments", g.Assessments)
	populateTimeRFC3339(objectMap, "createdTimestamp", g.CreatedTimestamp)
	populate(objectMap, "groupStatus", g.GroupStatus)
	populate(objectMap, "groupType", g.GroupType)
	populate(objectMap, "machineCount", g.MachineCount)
	populateTimeRFC3339(objectMap, "updatedTimestamp", g.UpdatedTimestamp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GroupProperties.
func (g *GroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "areAssessmentsRunning":
			err = unpopulate(val, &g.AreAssessmentsRunning)
			delete(rawMsg, key)
		case "assessments":
			err = unpopulate(val, &g.Assessments)
			delete(rawMsg, key)
		case "createdTimestamp":
			err = unpopulateTimeRFC3339(val, &g.CreatedTimestamp)
			delete(rawMsg, key)
		case "groupStatus":
			err = unpopulate(val, &g.GroupStatus)
			delete(rawMsg, key)
		case "groupType":
			err = unpopulate(val, &g.GroupType)
			delete(rawMsg, key)
		case "machineCount":
			err = unpopulate(val, &g.MachineCount)
			delete(rawMsg, key)
		case "updatedTimestamp":
			err = unpopulateTimeRFC3339(val, &g.UpdatedTimestamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GroupResultList.
func (g GroupResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", g.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HyperVCollectorList.
func (h HyperVCollectorList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", h.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImportCollectorList.
func (i ImportCollectorList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineProperties.
func (m MachineProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bootType", m.BootType)
	populateTimeRFC3339(objectMap, "createdTimestamp", m.CreatedTimestamp)
	populate(objectMap, "datacenterManagementServerArmId", m.DatacenterManagementServerArmID)
	populate(objectMap, "datacenterManagementServerName", m.DatacenterManagementServerName)
	populate(objectMap, "description", m.Description)
	populate(objectMap, "discoveryMachineArmId", m.DiscoveryMachineArmID)
	populate(objectMap, "disks", m.Disks)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "groups", m.Groups)
	populate(objectMap, "megabytesOfMemory", m.MegabytesOfMemory)
	populate(objectMap, "networkAdapters", m.NetworkAdapters)
	populate(objectMap, "numberOfCores", m.NumberOfCores)
	populate(objectMap, "operatingSystemName", m.OperatingSystemName)
	populate(objectMap, "operatingSystemType", m.OperatingSystemType)
	populate(objectMap, "operatingSystemVersion", m.OperatingSystemVersion)
	populateTimeRFC3339(objectMap, "updatedTimestamp", m.UpdatedTimestamp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MachineProperties.
func (m *MachineProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bootType":
			err = unpopulate(val, &m.BootType)
			delete(rawMsg, key)
		case "createdTimestamp":
			err = unpopulateTimeRFC3339(val, &m.CreatedTimestamp)
			delete(rawMsg, key)
		case "datacenterManagementServerArmId":
			err = unpopulate(val, &m.DatacenterManagementServerArmID)
			delete(rawMsg, key)
		case "datacenterManagementServerName":
			err = unpopulate(val, &m.DatacenterManagementServerName)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &m.Description)
			delete(rawMsg, key)
		case "discoveryMachineArmId":
			err = unpopulate(val, &m.DiscoveryMachineArmID)
			delete(rawMsg, key)
		case "disks":
			err = unpopulate(val, &m.Disks)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &m.DisplayName)
			delete(rawMsg, key)
		case "groups":
			err = unpopulate(val, &m.Groups)
			delete(rawMsg, key)
		case "megabytesOfMemory":
			err = unpopulate(val, &m.MegabytesOfMemory)
			delete(rawMsg, key)
		case "networkAdapters":
			err = unpopulate(val, &m.NetworkAdapters)
			delete(rawMsg, key)
		case "numberOfCores":
			err = unpopulate(val, &m.NumberOfCores)
			delete(rawMsg, key)
		case "operatingSystemName":
			err = unpopulate(val, &m.OperatingSystemName)
			delete(rawMsg, key)
		case "operatingSystemType":
			err = unpopulate(val, &m.OperatingSystemType)
			delete(rawMsg, key)
		case "operatingSystemVersion":
			err = unpopulate(val, &m.OperatingSystemVersion)
			delete(rawMsg, key)
		case "updatedTimestamp":
			err = unpopulateTimeRFC3339(val, &m.UpdatedTimestamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MachineResultList.
func (m MachineResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkAdapter.
func (n NetworkAdapter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", n.DisplayName)
	populate(objectMap, "ipAddresses", n.IPAddresses)
	populate(objectMap, "macAddress", n.MacAddress)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationResultList.
func (o OperationResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionCollection.
func (p PrivateEndpointConnectionCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceCollection.
func (p PrivateLinkResourceCollection) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Project.
func (p Project) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eTag", p.ETag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", &p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProjectProperties.
func (p ProjectProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assessmentSolutionId", p.AssessmentSolutionID)
	populateTimeRFC3339(objectMap, "createdTimestamp", p.CreatedTimestamp)
	populate(objectMap, "customerStorageAccountArmId", p.CustomerStorageAccountArmID)
	populate(objectMap, "customerWorkspaceId", p.CustomerWorkspaceID)
	populate(objectMap, "customerWorkspaceLocation", p.CustomerWorkspaceLocation)
	populateTimeRFC3339(objectMap, "lastAssessmentTimestamp", p.LastAssessmentTimestamp)
	populate(objectMap, "numberOfAssessments", p.NumberOfAssessments)
	populate(objectMap, "numberOfGroups", p.NumberOfGroups)
	populate(objectMap, "numberOfMachines", p.NumberOfMachines)
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "projectStatus", p.ProjectStatus)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	populate(objectMap, "serviceEndpoint", p.ServiceEndpoint)
	populateTimeRFC3339(objectMap, "updatedTimestamp", p.UpdatedTimestamp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProjectProperties.
func (p *ProjectProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assessmentSolutionId":
			err = unpopulate(val, &p.AssessmentSolutionID)
			delete(rawMsg, key)
		case "createdTimestamp":
			err = unpopulateTimeRFC3339(val, &p.CreatedTimestamp)
			delete(rawMsg, key)
		case "customerStorageAccountArmId":
			err = unpopulate(val, &p.CustomerStorageAccountArmID)
			delete(rawMsg, key)
		case "customerWorkspaceId":
			err = unpopulate(val, &p.CustomerWorkspaceID)
			delete(rawMsg, key)
		case "customerWorkspaceLocation":
			err = unpopulate(val, &p.CustomerWorkspaceLocation)
			delete(rawMsg, key)
		case "lastAssessmentTimestamp":
			err = unpopulateTimeRFC3339(val, &p.LastAssessmentTimestamp)
			delete(rawMsg, key)
		case "numberOfAssessments":
			err = unpopulate(val, &p.NumberOfAssessments)
			delete(rawMsg, key)
		case "numberOfGroups":
			err = unpopulate(val, &p.NumberOfGroups)
			delete(rawMsg, key)
		case "numberOfMachines":
			err = unpopulate(val, &p.NumberOfMachines)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &p.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "projectStatus":
			err = unpopulate(val, &p.ProjectStatus)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &p.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &p.PublicNetworkAccess)
			delete(rawMsg, key)
		case "serviceEndpoint":
			err = unpopulate(val, &p.ServiceEndpoint)
			delete(rawMsg, key)
		case "updatedTimestamp":
			err = unpopulateTimeRFC3339(val, &p.UpdatedTimestamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProjectResultList.
func (p ProjectResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerCollectorList.
func (s ServerCollectorList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VMFamily.
func (v VMFamily) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", v.Category)
	populate(objectMap, "familyName", v.FamilyName)
	populate(objectMap, "targetLocations", v.TargetLocations)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VMwareCollectorList.
func (v VMwareCollectorList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
