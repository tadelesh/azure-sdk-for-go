//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicemap

// ClientGroupsClientGetMembersCountResponse contains the response from method ClientGroupsClient.GetMembersCount.
type ClientGroupsClientGetMembersCountResponse struct {
	ClientGroupMembersCount
}

// ClientGroupsClientGetResponse contains the response from method ClientGroupsClient.Get.
type ClientGroupsClientGetResponse struct {
	ClientGroup
}

// ClientGroupsClientListMembersResponse contains the response from method ClientGroupsClient.ListMembers.
type ClientGroupsClientListMembersResponse struct {
	ClientGroupMembersCollection
}

// MachineGroupsClientCreateResponse contains the response from method MachineGroupsClient.Create.
type MachineGroupsClientCreateResponse struct {
	MachineGroup
}

// MachineGroupsClientDeleteResponse contains the response from method MachineGroupsClient.Delete.
type MachineGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineGroupsClientGetResponse contains the response from method MachineGroupsClient.Get.
type MachineGroupsClientGetResponse struct {
	MachineGroup
}

// MachineGroupsClientListByWorkspaceResponse contains the response from method MachineGroupsClient.ListByWorkspace.
type MachineGroupsClientListByWorkspaceResponse struct {
	MachineGroupCollection
}

// MachineGroupsClientUpdateResponse contains the response from method MachineGroupsClient.Update.
type MachineGroupsClientUpdateResponse struct {
	MachineGroup
}

// MachinesClientGetLivenessResponse contains the response from method MachinesClient.GetLiveness.
type MachinesClientGetLivenessResponse struct {
	Liveness
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	Machine
}

// MachinesClientListByWorkspaceResponse contains the response from method MachinesClient.ListByWorkspace.
type MachinesClientListByWorkspaceResponse struct {
	MachineCollection
}

// MachinesClientListConnectionsResponse contains the response from method MachinesClient.ListConnections.
type MachinesClientListConnectionsResponse struct {
	ConnectionCollection
}

// MachinesClientListMachineGroupMembershipResponse contains the response from method MachinesClient.ListMachineGroupMembership.
type MachinesClientListMachineGroupMembershipResponse struct {
	MachineGroupCollection
}

// MachinesClientListPortsResponse contains the response from method MachinesClient.ListPorts.
type MachinesClientListPortsResponse struct {
	PortCollection
}

// MachinesClientListProcessesResponse contains the response from method MachinesClient.ListProcesses.
type MachinesClientListProcessesResponse struct {
	ProcessCollection
}

// MapsClientGenerateResponse contains the response from method MapsClient.Generate.
type MapsClientGenerateResponse struct {
	MapResponse
}

// PortsClientGetLivenessResponse contains the response from method PortsClient.GetLiveness.
type PortsClientGetLivenessResponse struct {
	Liveness
}

// PortsClientGetResponse contains the response from method PortsClient.Get.
type PortsClientGetResponse struct {
	Port
}

// PortsClientListAcceptingProcessesResponse contains the response from method PortsClient.ListAcceptingProcesses.
type PortsClientListAcceptingProcessesResponse struct {
	ProcessCollection
}

// PortsClientListConnectionsResponse contains the response from method PortsClient.ListConnections.
type PortsClientListConnectionsResponse struct {
	ConnectionCollection
}

// ProcessesClientGetLivenessResponse contains the response from method ProcessesClient.GetLiveness.
type ProcessesClientGetLivenessResponse struct {
	Liveness
}

// ProcessesClientGetResponse contains the response from method ProcessesClient.Get.
type ProcessesClientGetResponse struct {
	Process
}

// ProcessesClientListAcceptingPortsResponse contains the response from method ProcessesClient.ListAcceptingPorts.
type ProcessesClientListAcceptingPortsResponse struct {
	PortCollection
}

// ProcessesClientListConnectionsResponse contains the response from method ProcessesClient.ListConnections.
type ProcessesClientListConnectionsResponse struct {
	ConnectionCollection
}

// SummariesClientGetMachinesResponse contains the response from method SummariesClient.GetMachines.
type SummariesClientGetMachinesResponse struct {
	MachinesSummary
}

