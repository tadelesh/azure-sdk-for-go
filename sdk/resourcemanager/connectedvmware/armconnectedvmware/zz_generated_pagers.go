//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armconnectedvmware

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClustersClientListByResourceGroupPager provides operations for iterating over paged responses.
type ClustersClientListByResourceGroupPager struct {
	client    *ClustersClient
	current   ClustersClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClustersClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClustersList.NextLink == nil || len(*p.current.ClustersList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClustersClientListByResourceGroupPager) NextPage(ctx context.Context) (ClustersClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClustersClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClustersClientListPager provides operations for iterating over paged responses.
type ClustersClientListPager struct {
	client    *ClustersClient
	current   ClustersClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClustersClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClustersList.NextLink == nil || len(*p.current.ClustersList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClustersClientListPager) NextPage(ctx context.Context) (ClustersClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClustersClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClustersClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClustersClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClustersClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ClustersClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// DatastoresClientListByResourceGroupPager provides operations for iterating over paged responses.
type DatastoresClientListByResourceGroupPager struct {
	client    *DatastoresClient
	current   DatastoresClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatastoresClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *DatastoresClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatastoresList.NextLink == nil || len(*p.current.DatastoresList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *DatastoresClientListByResourceGroupPager) NextPage(ctx context.Context) (DatastoresClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return DatastoresClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return DatastoresClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return DatastoresClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return DatastoresClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return DatastoresClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// DatastoresClientListPager provides operations for iterating over paged responses.
type DatastoresClientListPager struct {
	client    *DatastoresClient
	current   DatastoresClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatastoresClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *DatastoresClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatastoresList.NextLink == nil || len(*p.current.DatastoresList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *DatastoresClientListPager) NextPage(ctx context.Context) (DatastoresClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return DatastoresClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return DatastoresClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return DatastoresClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return DatastoresClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return DatastoresClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// GuestAgentsClientListByVMPager provides operations for iterating over paged responses.
type GuestAgentsClientListByVMPager struct {
	client    *GuestAgentsClient
	current   GuestAgentsClientListByVMResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GuestAgentsClientListByVMResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *GuestAgentsClientListByVMPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GuestAgentList.NextLink == nil || len(*p.current.GuestAgentList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *GuestAgentsClientListByVMPager) NextPage(ctx context.Context) (GuestAgentsClientListByVMResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return GuestAgentsClientListByVMResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return GuestAgentsClientListByVMResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return GuestAgentsClientListByVMResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return GuestAgentsClientListByVMResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByVMHandleResponse(resp)
	if err != nil {
		return GuestAgentsClientListByVMResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// HostsClientListByResourceGroupPager provides operations for iterating over paged responses.
type HostsClientListByResourceGroupPager struct {
	client    *HostsClient
	current   HostsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HostsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *HostsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HostsList.NextLink == nil || len(*p.current.HostsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *HostsClientListByResourceGroupPager) NextPage(ctx context.Context) (HostsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return HostsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return HostsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return HostsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return HostsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return HostsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// HostsClientListPager provides operations for iterating over paged responses.
type HostsClientListPager struct {
	client    *HostsClient
	current   HostsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HostsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *HostsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HostsList.NextLink == nil || len(*p.current.HostsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *HostsClientListPager) NextPage(ctx context.Context) (HostsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return HostsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return HostsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return HostsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return HostsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return HostsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// HybridIdentityMetadataClientListByVMPager provides operations for iterating over paged responses.
type HybridIdentityMetadataClientListByVMPager struct {
	client    *HybridIdentityMetadataClient
	current   HybridIdentityMetadataClientListByVMResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HybridIdentityMetadataClientListByVMResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *HybridIdentityMetadataClientListByVMPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HybridIdentityMetadataList.NextLink == nil || len(*p.current.HybridIdentityMetadataList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *HybridIdentityMetadataClientListByVMPager) NextPage(ctx context.Context) (HybridIdentityMetadataClientListByVMResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return HybridIdentityMetadataClientListByVMResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return HybridIdentityMetadataClientListByVMResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return HybridIdentityMetadataClientListByVMResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return HybridIdentityMetadataClientListByVMResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByVMHandleResponse(resp)
	if err != nil {
		return HybridIdentityMetadataClientListByVMResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// InventoryItemsClientListByVCenterPager provides operations for iterating over paged responses.
type InventoryItemsClientListByVCenterPager struct {
	client    *InventoryItemsClient
	current   InventoryItemsClientListByVCenterResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, InventoryItemsClientListByVCenterResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *InventoryItemsClientListByVCenterPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.InventoryItemsList.NextLink == nil || len(*p.current.InventoryItemsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *InventoryItemsClientListByVCenterPager) NextPage(ctx context.Context) (InventoryItemsClientListByVCenterResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return InventoryItemsClientListByVCenterResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return InventoryItemsClientListByVCenterResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return InventoryItemsClientListByVCenterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return InventoryItemsClientListByVCenterResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByVCenterHandleResponse(resp)
	if err != nil {
		return InventoryItemsClientListByVCenterResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// MachineExtensionsClientListPager provides operations for iterating over paged responses.
type MachineExtensionsClientListPager struct {
	client    *MachineExtensionsClient
	current   MachineExtensionsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MachineExtensionsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *MachineExtensionsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MachineExtensionsListResult.NextLink == nil || len(*p.current.MachineExtensionsListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *MachineExtensionsClientListPager) NextPage(ctx context.Context) (MachineExtensionsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return MachineExtensionsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return MachineExtensionsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return MachineExtensionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return MachineExtensionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return MachineExtensionsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *OperationsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationsList.NextLink == nil || len(*p.current.OperationsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *OperationsClientListPager) NextPage(ctx context.Context) (OperationsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return OperationsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return OperationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ResourcePoolsClientListByResourceGroupPager provides operations for iterating over paged responses.
type ResourcePoolsClientListByResourceGroupPager struct {
	client    *ResourcePoolsClient
	current   ResourcePoolsClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ResourcePoolsClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ResourcePoolsClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourcePoolsList.NextLink == nil || len(*p.current.ResourcePoolsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ResourcePoolsClientListByResourceGroupPager) NextPage(ctx context.Context) (ResourcePoolsClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ResourcePoolsClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ResourcePoolsClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ResourcePoolsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ResourcePoolsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return ResourcePoolsClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ResourcePoolsClientListPager provides operations for iterating over paged responses.
type ResourcePoolsClientListPager struct {
	client    *ResourcePoolsClient
	current   ResourcePoolsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ResourcePoolsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ResourcePoolsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourcePoolsList.NextLink == nil || len(*p.current.ResourcePoolsList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ResourcePoolsClientListPager) NextPage(ctx context.Context) (ResourcePoolsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ResourcePoolsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ResourcePoolsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ResourcePoolsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ResourcePoolsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return ResourcePoolsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VCentersClientListByResourceGroupPager provides operations for iterating over paged responses.
type VCentersClientListByResourceGroupPager struct {
	client    *VCentersClient
	current   VCentersClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VCentersClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VCentersClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VCentersList.NextLink == nil || len(*p.current.VCentersList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VCentersClientListByResourceGroupPager) NextPage(ctx context.Context) (VCentersClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VCentersClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VCentersClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VCentersClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VCentersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return VCentersClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VCentersClientListPager provides operations for iterating over paged responses.
type VCentersClientListPager struct {
	client    *VCentersClient
	current   VCentersClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VCentersClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VCentersClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VCentersList.NextLink == nil || len(*p.current.VCentersList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VCentersClientListPager) NextPage(ctx context.Context) (VCentersClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VCentersClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VCentersClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VCentersClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VCentersClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return VCentersClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualMachineTemplatesClientListByResourceGroupPager provides operations for iterating over paged responses.
type VirtualMachineTemplatesClientListByResourceGroupPager struct {
	client    *VirtualMachineTemplatesClient
	current   VirtualMachineTemplatesClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachineTemplatesClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualMachineTemplatesClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachineTemplatesList.NextLink == nil || len(*p.current.VirtualMachineTemplatesList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualMachineTemplatesClientListByResourceGroupPager) NextPage(ctx context.Context) (VirtualMachineTemplatesClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualMachineTemplatesClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualMachineTemplatesClientListPager provides operations for iterating over paged responses.
type VirtualMachineTemplatesClientListPager struct {
	client    *VirtualMachineTemplatesClient
	current   VirtualMachineTemplatesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachineTemplatesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualMachineTemplatesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachineTemplatesList.NextLink == nil || len(*p.current.VirtualMachineTemplatesList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualMachineTemplatesClientListPager) NextPage(ctx context.Context) (VirtualMachineTemplatesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualMachineTemplatesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualMachineTemplatesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualMachineTemplatesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return VirtualMachineTemplatesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualMachinesClientListByResourceGroupPager provides operations for iterating over paged responses.
type VirtualMachinesClientListByResourceGroupPager struct {
	client    *VirtualMachinesClient
	current   VirtualMachinesClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachinesClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualMachinesClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachinesList.NextLink == nil || len(*p.current.VirtualMachinesList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualMachinesClientListByResourceGroupPager) NextPage(ctx context.Context) (VirtualMachinesClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualMachinesClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualMachinesClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualMachinesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualMachinesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return VirtualMachinesClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualMachinesClientListPager provides operations for iterating over paged responses.
type VirtualMachinesClientListPager struct {
	client    *VirtualMachinesClient
	current   VirtualMachinesClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachinesClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualMachinesClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachinesList.NextLink == nil || len(*p.current.VirtualMachinesList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualMachinesClientListPager) NextPage(ctx context.Context) (VirtualMachinesClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualMachinesClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualMachinesClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualMachinesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualMachinesClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return VirtualMachinesClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualNetworksClientListByResourceGroupPager provides operations for iterating over paged responses.
type VirtualNetworksClientListByResourceGroupPager struct {
	client    *VirtualNetworksClient
	current   VirtualNetworksClientListByResourceGroupResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworksClientListByResourceGroupResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualNetworksClientListByResourceGroupPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworksList.NextLink == nil || len(*p.current.VirtualNetworksList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualNetworksClientListByResourceGroupPager) NextPage(ctx context.Context) (VirtualNetworksClientListByResourceGroupResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualNetworksClientListByResourceGroupResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualNetworksClientListByResourceGroupResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualNetworksClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualNetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		return VirtualNetworksClientListByResourceGroupResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// VirtualNetworksClientListPager provides operations for iterating over paged responses.
type VirtualNetworksClientListPager struct {
	client    *VirtualNetworksClient
	current   VirtualNetworksClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworksClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *VirtualNetworksClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworksList.NextLink == nil || len(*p.current.VirtualNetworksList.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *VirtualNetworksClientListPager) NextPage(ctx context.Context) (VirtualNetworksClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return VirtualNetworksClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return VirtualNetworksClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
