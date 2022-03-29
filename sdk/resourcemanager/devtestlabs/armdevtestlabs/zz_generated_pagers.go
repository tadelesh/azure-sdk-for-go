//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ArmTemplatesClientListPager provides operations for iterating over paged responses.
type ArmTemplatesClientListPager struct {
	client    *ArmTemplatesClient
	current   ArmTemplatesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ArmTemplatesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ArmTemplatesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ArmTemplatesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArmTemplateList.NextLink == nil || len(*p.current.ArmTemplateList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ArmTemplatesClientListResponse page.
func (p *ArmTemplatesClientListPager) PageResponse() ArmTemplatesClientListResponse {
	return p.current
}

// ArtifactSourcesClientListPager provides operations for iterating over paged responses.
type ArtifactSourcesClientListPager struct {
	client    *ArtifactSourcesClient
	current   ArtifactSourcesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ArtifactSourcesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ArtifactSourcesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ArtifactSourcesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArtifactSourceList.NextLink == nil || len(*p.current.ArtifactSourceList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ArtifactSourcesClientListResponse page.
func (p *ArtifactSourcesClientListPager) PageResponse() ArtifactSourcesClientListResponse {
	return p.current
}

// ArtifactsClientListPager provides operations for iterating over paged responses.
type ArtifactsClientListPager struct {
	client    *ArtifactsClient
	current   ArtifactsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ArtifactsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ArtifactsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ArtifactsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArtifactList.NextLink == nil || len(*p.current.ArtifactList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ArtifactsClientListResponse page.
func (p *ArtifactsClientListPager) PageResponse() ArtifactsClientListResponse {
	return p.current
}

// CustomImagesClientListPager provides operations for iterating over paged responses.
type CustomImagesClientListPager struct {
	client    *CustomImagesClient
	current   CustomImagesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CustomImagesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CustomImagesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CustomImagesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CustomImageList.NextLink == nil || len(*p.current.CustomImageList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current CustomImagesClientListResponse page.
func (p *CustomImagesClientListPager) PageResponse() CustomImagesClientListResponse {
	return p.current
}

// DisksClientListPager provides operations for iterating over paged responses.
type DisksClientListPager struct {
	client    *DisksClient
	current   DisksClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DisksClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DisksClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DisksClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DiskList.NextLink == nil || len(*p.current.DiskList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current DisksClientListResponse page.
func (p *DisksClientListPager) PageResponse() DisksClientListResponse {
	return p.current
}

// EnvironmentsClientListPager provides operations for iterating over paged responses.
type EnvironmentsClientListPager struct {
	client    *EnvironmentsClient
	current   EnvironmentsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EnvironmentsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EnvironmentsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EnvironmentsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DtlEnvironmentList.NextLink == nil || len(*p.current.DtlEnvironmentList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current EnvironmentsClientListResponse page.
func (p *EnvironmentsClientListPager) PageResponse() EnvironmentsClientListResponse {
	return p.current
}

// FormulasClientListPager provides operations for iterating over paged responses.
type FormulasClientListPager struct {
	client    *FormulasClient
	current   FormulasClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FormulasClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FormulasClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FormulasClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FormulaList.NextLink == nil || len(*p.current.FormulaList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current FormulasClientListResponse page.
func (p *FormulasClientListPager) PageResponse() FormulasClientListResponse {
	return p.current
}

// GalleryImagesClientListPager provides operations for iterating over paged responses.
type GalleryImagesClientListPager struct {
	client    *GalleryImagesClient
	current   GalleryImagesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GalleryImagesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *GalleryImagesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *GalleryImagesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GalleryImageList.NextLink == nil || len(*p.current.GalleryImageList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current GalleryImagesClientListResponse page.
func (p *GalleryImagesClientListPager) PageResponse() GalleryImagesClientListResponse {
	return p.current
}

// GlobalSchedulesClientListByResourceGroupPager provides operations for iterating over paged responses.
type GlobalSchedulesClientListByResourceGroupPager struct {
	client    *GlobalSchedulesClient
	current   GlobalSchedulesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GlobalSchedulesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *GlobalSchedulesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *GlobalSchedulesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current GlobalSchedulesClientListByResourceGroupResponse page.
func (p *GlobalSchedulesClientListByResourceGroupPager) PageResponse() GlobalSchedulesClientListByResourceGroupResponse {
	return p.current
}

// GlobalSchedulesClientListBySubscriptionPager provides operations for iterating over paged responses.
type GlobalSchedulesClientListBySubscriptionPager struct {
	client    *GlobalSchedulesClient
	current   GlobalSchedulesClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GlobalSchedulesClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *GlobalSchedulesClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *GlobalSchedulesClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current GlobalSchedulesClientListBySubscriptionResponse page.
func (p *GlobalSchedulesClientListBySubscriptionPager) PageResponse() GlobalSchedulesClientListBySubscriptionResponse {
	return p.current
}

// LabsClientListByResourceGroupPager provides operations for iterating over paged responses.
type LabsClientListByResourceGroupPager struct {
	client    *LabsClient
	current   LabsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LabsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LabsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LabsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LabList.NextLink == nil || len(*p.current.LabList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current LabsClientListByResourceGroupResponse page.
func (p *LabsClientListByResourceGroupPager) PageResponse() LabsClientListByResourceGroupResponse {
	return p.current
}

// LabsClientListBySubscriptionPager provides operations for iterating over paged responses.
type LabsClientListBySubscriptionPager struct {
	client    *LabsClient
	current   LabsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LabsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LabsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LabsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LabList.NextLink == nil || len(*p.current.LabList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current LabsClientListBySubscriptionResponse page.
func (p *LabsClientListBySubscriptionPager) PageResponse() LabsClientListBySubscriptionResponse {
	return p.current
}

// LabsClientListVhdsPager provides operations for iterating over paged responses.
type LabsClientListVhdsPager struct {
	client    *LabsClient
	current   LabsClientListVhdsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LabsClientListVhdsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LabsClientListVhdsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LabsClientListVhdsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LabVhdList.NextLink == nil || len(*p.current.LabVhdList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listVhdsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LabsClientListVhdsResponse page.
func (p *LabsClientListVhdsPager) PageResponse() LabsClientListVhdsResponse {
	return p.current
}

// NotificationChannelsClientListPager provides operations for iterating over paged responses.
type NotificationChannelsClientListPager struct {
	client    *NotificationChannelsClient
	current   NotificationChannelsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NotificationChannelsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NotificationChannelsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NotificationChannelsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotificationChannelList.NextLink == nil || len(*p.current.NotificationChannelList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current NotificationChannelsClientListResponse page.
func (p *NotificationChannelsClientListPager) PageResponse() NotificationChannelsClientListResponse {
	return p.current
}

// PoliciesClientListPager provides operations for iterating over paged responses.
type PoliciesClientListPager struct {
	client    *PoliciesClient
	current   PoliciesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PoliciesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PoliciesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PoliciesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PolicyList.NextLink == nil || len(*p.current.PolicyList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current PoliciesClientListResponse page.
func (p *PoliciesClientListPager) PageResponse() PoliciesClientListResponse {
	return p.current
}

// ProviderOperationsClientListPager provides operations for iterating over paged responses.
type ProviderOperationsClientListPager struct {
	client    *ProviderOperationsClient
	current   ProviderOperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProviderOperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ProviderOperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ProviderOperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProviderOperationResult.NextLink == nil || len(*p.current.ProviderOperationResult.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ProviderOperationsClientListResponse page.
func (p *ProviderOperationsClientListPager) PageResponse() ProviderOperationsClientListResponse {
	return p.current
}

// SchedulesClientListApplicablePager provides operations for iterating over paged responses.
type SchedulesClientListApplicablePager struct {
	client    *SchedulesClient
	current   SchedulesClientListApplicableResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SchedulesClientListApplicableResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SchedulesClientListApplicablePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SchedulesClientListApplicablePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listApplicableHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SchedulesClientListApplicableResponse page.
func (p *SchedulesClientListApplicablePager) PageResponse() SchedulesClientListApplicableResponse {
	return p.current
}

// SchedulesClientListPager provides operations for iterating over paged responses.
type SchedulesClientListPager struct {
	client    *SchedulesClient
	current   SchedulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SchedulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SchedulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SchedulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current SchedulesClientListResponse page.
func (p *SchedulesClientListPager) PageResponse() SchedulesClientListResponse {
	return p.current
}

// SecretsClientListPager provides operations for iterating over paged responses.
type SecretsClientListPager struct {
	client    *SecretsClient
	current   SecretsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SecretsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SecretsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SecretsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretList.NextLink == nil || len(*p.current.SecretList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current SecretsClientListResponse page.
func (p *SecretsClientListPager) PageResponse() SecretsClientListResponse {
	return p.current
}

// ServiceFabricSchedulesClientListPager provides operations for iterating over paged responses.
type ServiceFabricSchedulesClientListPager struct {
	client    *ServiceFabricSchedulesClient
	current   ServiceFabricSchedulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceFabricSchedulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServiceFabricSchedulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServiceFabricSchedulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ServiceFabricSchedulesClientListResponse page.
func (p *ServiceFabricSchedulesClientListPager) PageResponse() ServiceFabricSchedulesClientListResponse {
	return p.current
}

// ServiceFabricsClientListPager provides operations for iterating over paged responses.
type ServiceFabricsClientListPager struct {
	client    *ServiceFabricsClient
	current   ServiceFabricsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceFabricsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServiceFabricsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServiceFabricsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServiceFabricList.NextLink == nil || len(*p.current.ServiceFabricList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current ServiceFabricsClientListResponse page.
func (p *ServiceFabricsClientListPager) PageResponse() ServiceFabricsClientListResponse {
	return p.current
}

// UsersClientListPager provides operations for iterating over paged responses.
type UsersClientListPager struct {
	client    *UsersClient
	current   UsersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UsersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *UsersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *UsersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.UserList.NextLink == nil || len(*p.current.UserList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current UsersClientListResponse page.
func (p *UsersClientListPager) PageResponse() UsersClientListResponse {
	return p.current
}

// VirtualMachineSchedulesClientListPager provides operations for iterating over paged responses.
type VirtualMachineSchedulesClientListPager struct {
	client    *VirtualMachineSchedulesClient
	current   VirtualMachineSchedulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachineSchedulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualMachineSchedulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualMachineSchedulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScheduleList.NextLink == nil || len(*p.current.ScheduleList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current VirtualMachineSchedulesClientListResponse page.
func (p *VirtualMachineSchedulesClientListPager) PageResponse() VirtualMachineSchedulesClientListResponse {
	return p.current
}

// VirtualMachinesClientListPager provides operations for iterating over paged responses.
type VirtualMachinesClientListPager struct {
	client    *VirtualMachinesClient
	current   VirtualMachinesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachinesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualMachinesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualMachinesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LabVirtualMachineList.NextLink == nil || len(*p.current.LabVirtualMachineList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current VirtualMachinesClientListResponse page.
func (p *VirtualMachinesClientListPager) PageResponse() VirtualMachinesClientListResponse {
	return p.current
}

// VirtualNetworksClientListPager provides operations for iterating over paged responses.
type VirtualNetworksClientListPager struct {
	client    *VirtualNetworksClient
	current   VirtualNetworksClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworksClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualNetworksClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualNetworksClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkList.NextLink == nil || len(*p.current.VirtualNetworkList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current VirtualNetworksClientListResponse page.
func (p *VirtualNetworksClientListPager) PageResponse() VirtualNetworksClientListResponse {
	return p.current
}