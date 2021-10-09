//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AccountsListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountList.NextLink == nil || len(*p.current.AccountList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsListByResourceGroupResponse page.
func (p *AccountsListByResourceGroupPager) PageResponse() AccountsListByResourceGroupResponse {
	return p.current
}

// AccountsListBySubscriptionPager provides operations for iterating over paged responses.
type AccountsListBySubscriptionPager struct {
	client    *AccountsClient
	current   AccountsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountList.NextLink == nil || len(*p.current.AccountList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsListBySubscriptionResponse page.
func (p *AccountsListBySubscriptionPager) PageResponse() AccountsListBySubscriptionResponse {
	return p.current
}

// ConsumerInvitationsListInvitationsPager provides operations for iterating over paged responses.
type ConsumerInvitationsListInvitationsPager struct {
	client    *ConsumerInvitationsClient
	current   ConsumerInvitationsListInvitationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConsumerInvitationsListInvitationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConsumerInvitationsListInvitationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConsumerInvitationsListInvitationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConsumerInvitationList.NextLink == nil || len(*p.current.ConsumerInvitationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listInvitationsHandleError(resp)
		return false
	}
	result, err := p.client.listInvitationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConsumerInvitationsListInvitationsResponse page.
func (p *ConsumerInvitationsListInvitationsPager) PageResponse() ConsumerInvitationsListInvitationsResponse {
	return p.current
}

// ConsumerSourceDataSetsListByShareSubscriptionPager provides operations for iterating over paged responses.
type ConsumerSourceDataSetsListByShareSubscriptionPager struct {
	client    *ConsumerSourceDataSetsClient
	current   ConsumerSourceDataSetsListByShareSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConsumerSourceDataSetsListByShareSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConsumerSourceDataSetsListByShareSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConsumerSourceDataSetsListByShareSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConsumerSourceDataSetList.NextLink == nil || len(*p.current.ConsumerSourceDataSetList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareSubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listByShareSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConsumerSourceDataSetsListByShareSubscriptionResponse page.
func (p *ConsumerSourceDataSetsListByShareSubscriptionPager) PageResponse() ConsumerSourceDataSetsListByShareSubscriptionResponse {
	return p.current
}

// DataSetMappingsListByShareSubscriptionPager provides operations for iterating over paged responses.
type DataSetMappingsListByShareSubscriptionPager struct {
	client    *DataSetMappingsClient
	current   DataSetMappingsListByShareSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataSetMappingsListByShareSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataSetMappingsListByShareSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataSetMappingsListByShareSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataSetMappingList.NextLink == nil || len(*p.current.DataSetMappingList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareSubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listByShareSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataSetMappingsListByShareSubscriptionResponse page.
func (p *DataSetMappingsListByShareSubscriptionPager) PageResponse() DataSetMappingsListByShareSubscriptionResponse {
	return p.current
}

// DataSetsListBySharePager provides operations for iterating over paged responses.
type DataSetsListBySharePager struct {
	client    *DataSetsClient
	current   DataSetsListByShareResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataSetsListByShareResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataSetsListBySharePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataSetsListBySharePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataSetList.NextLink == nil || len(*p.current.DataSetList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareHandleError(resp)
		return false
	}
	result, err := p.client.listByShareHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataSetsListByShareResponse page.
func (p *DataSetsListBySharePager) PageResponse() DataSetsListByShareResponse {
	return p.current
}

// InvitationsListBySharePager provides operations for iterating over paged responses.
type InvitationsListBySharePager struct {
	client    *InvitationsClient
	current   InvitationsListByShareResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, InvitationsListByShareResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *InvitationsListBySharePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *InvitationsListBySharePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.InvitationList.NextLink == nil || len(*p.current.InvitationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareHandleError(resp)
		return false
	}
	result, err := p.client.listByShareHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current InvitationsListByShareResponse page.
func (p *InvitationsListBySharePager) PageResponse() InvitationsListByShareResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// ProviderShareSubscriptionsListBySharePager provides operations for iterating over paged responses.
type ProviderShareSubscriptionsListBySharePager struct {
	client    *ProviderShareSubscriptionsClient
	current   ProviderShareSubscriptionsListByShareResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProviderShareSubscriptionsListByShareResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ProviderShareSubscriptionsListBySharePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ProviderShareSubscriptionsListBySharePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProviderShareSubscriptionList.NextLink == nil || len(*p.current.ProviderShareSubscriptionList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareHandleError(resp)
		return false
	}
	result, err := p.client.listByShareHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ProviderShareSubscriptionsListByShareResponse page.
func (p *ProviderShareSubscriptionsListBySharePager) PageResponse() ProviderShareSubscriptionsListByShareResponse {
	return p.current
}

// ShareSubscriptionsListByAccountPager provides operations for iterating over paged responses.
type ShareSubscriptionsListByAccountPager struct {
	client    *ShareSubscriptionsClient
	current   ShareSubscriptionsListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ShareSubscriptionsListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ShareSubscriptionsListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ShareSubscriptionsListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ShareSubscriptionList.NextLink == nil || len(*p.current.ShareSubscriptionList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ShareSubscriptionsListByAccountResponse page.
func (p *ShareSubscriptionsListByAccountPager) PageResponse() ShareSubscriptionsListByAccountResponse {
	return p.current
}

// ShareSubscriptionsListSourceShareSynchronizationSettingsPager provides operations for iterating over paged responses.
type ShareSubscriptionsListSourceShareSynchronizationSettingsPager struct {
	client    *ShareSubscriptionsClient
	current   ShareSubscriptionsListSourceShareSynchronizationSettingsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ShareSubscriptionsListSourceShareSynchronizationSettingsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ShareSubscriptionsListSourceShareSynchronizationSettingsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ShareSubscriptionsListSourceShareSynchronizationSettingsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SourceShareSynchronizationSettingList.NextLink == nil || len(*p.current.SourceShareSynchronizationSettingList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSourceShareSynchronizationSettingsHandleError(resp)
		return false
	}
	result, err := p.client.listSourceShareSynchronizationSettingsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ShareSubscriptionsListSourceShareSynchronizationSettingsResponse page.
func (p *ShareSubscriptionsListSourceShareSynchronizationSettingsPager) PageResponse() ShareSubscriptionsListSourceShareSynchronizationSettingsResponse {
	return p.current
}

// ShareSubscriptionsListSynchronizationDetailsPager provides operations for iterating over paged responses.
type ShareSubscriptionsListSynchronizationDetailsPager struct {
	client    *ShareSubscriptionsClient
	current   ShareSubscriptionsListSynchronizationDetailsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ShareSubscriptionsListSynchronizationDetailsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ShareSubscriptionsListSynchronizationDetailsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ShareSubscriptionsListSynchronizationDetailsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SynchronizationDetailsList.NextLink == nil || len(*p.current.SynchronizationDetailsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSynchronizationDetailsHandleError(resp)
		return false
	}
	result, err := p.client.listSynchronizationDetailsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ShareSubscriptionsListSynchronizationDetailsResponse page.
func (p *ShareSubscriptionsListSynchronizationDetailsPager) PageResponse() ShareSubscriptionsListSynchronizationDetailsResponse {
	return p.current
}

// ShareSubscriptionsListSynchronizationsPager provides operations for iterating over paged responses.
type ShareSubscriptionsListSynchronizationsPager struct {
	client    *ShareSubscriptionsClient
	current   ShareSubscriptionsListSynchronizationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ShareSubscriptionsListSynchronizationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ShareSubscriptionsListSynchronizationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ShareSubscriptionsListSynchronizationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ShareSubscriptionSynchronizationList.NextLink == nil || len(*p.current.ShareSubscriptionSynchronizationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSynchronizationsHandleError(resp)
		return false
	}
	result, err := p.client.listSynchronizationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ShareSubscriptionsListSynchronizationsResponse page.
func (p *ShareSubscriptionsListSynchronizationsPager) PageResponse() ShareSubscriptionsListSynchronizationsResponse {
	return p.current
}

// SharesListByAccountPager provides operations for iterating over paged responses.
type SharesListByAccountPager struct {
	client    *SharesClient
	current   SharesListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SharesListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SharesListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SharesListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ShareList.NextLink == nil || len(*p.current.ShareList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SharesListByAccountResponse page.
func (p *SharesListByAccountPager) PageResponse() SharesListByAccountResponse {
	return p.current
}

// SharesListSynchronizationDetailsPager provides operations for iterating over paged responses.
type SharesListSynchronizationDetailsPager struct {
	client    *SharesClient
	current   SharesListSynchronizationDetailsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SharesListSynchronizationDetailsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SharesListSynchronizationDetailsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SharesListSynchronizationDetailsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SynchronizationDetailsList.NextLink == nil || len(*p.current.SynchronizationDetailsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSynchronizationDetailsHandleError(resp)
		return false
	}
	result, err := p.client.listSynchronizationDetailsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SharesListSynchronizationDetailsResponse page.
func (p *SharesListSynchronizationDetailsPager) PageResponse() SharesListSynchronizationDetailsResponse {
	return p.current
}

// SharesListSynchronizationsPager provides operations for iterating over paged responses.
type SharesListSynchronizationsPager struct {
	client    *SharesClient
	current   SharesListSynchronizationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SharesListSynchronizationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SharesListSynchronizationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SharesListSynchronizationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ShareSynchronizationList.NextLink == nil || len(*p.current.ShareSynchronizationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSynchronizationsHandleError(resp)
		return false
	}
	result, err := p.client.listSynchronizationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SharesListSynchronizationsResponse page.
func (p *SharesListSynchronizationsPager) PageResponse() SharesListSynchronizationsResponse {
	return p.current
}

// SynchronizationSettingsListBySharePager provides operations for iterating over paged responses.
type SynchronizationSettingsListBySharePager struct {
	client    *SynchronizationSettingsClient
	current   SynchronizationSettingsListByShareResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SynchronizationSettingsListByShareResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SynchronizationSettingsListBySharePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SynchronizationSettingsListBySharePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SynchronizationSettingList.NextLink == nil || len(*p.current.SynchronizationSettingList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareHandleError(resp)
		return false
	}
	result, err := p.client.listByShareHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SynchronizationSettingsListByShareResponse page.
func (p *SynchronizationSettingsListBySharePager) PageResponse() SynchronizationSettingsListByShareResponse {
	return p.current
}

// TriggersListByShareSubscriptionPager provides operations for iterating over paged responses.
type TriggersListByShareSubscriptionPager struct {
	client    *TriggersClient
	current   TriggersListByShareSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TriggersListByShareSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TriggersListByShareSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TriggersListByShareSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TriggerList.NextLink == nil || len(*p.current.TriggerList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByShareSubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listByShareSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TriggersListByShareSubscriptionResponse page.
func (p *TriggersListByShareSubscriptionPager) PageResponse() TriggersListByShareSubscriptionResponse {
	return p.current
}
