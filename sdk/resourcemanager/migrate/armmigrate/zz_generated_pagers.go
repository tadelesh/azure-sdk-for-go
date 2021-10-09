//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AssessedMachinesListByAssessmentPager provides operations for iterating over paged responses.
type AssessedMachinesListByAssessmentPager struct {
	client    *AssessedMachinesClient
	current   AssessedMachinesListByAssessmentResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AssessedMachinesListByAssessmentResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AssessedMachinesListByAssessmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AssessedMachinesListByAssessmentPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AssessedMachineResultList.NextLink == nil || len(*p.current.AssessedMachineResultList.NextLink) == 0 {
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
		p.err = p.client.listByAssessmentHandleError(resp)
		return false
	}
	result, err := p.client.listByAssessmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AssessedMachinesListByAssessmentResponse page.
func (p *AssessedMachinesListByAssessmentPager) PageResponse() AssessedMachinesListByAssessmentResponse {
	return p.current
}

// MachinesListByProjectPager provides operations for iterating over paged responses.
type MachinesListByProjectPager struct {
	client    *MachinesClient
	current   MachinesListByProjectResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MachinesListByProjectResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MachinesListByProjectPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MachinesListByProjectPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MachineResultList.NextLink == nil || len(*p.current.MachineResultList.NextLink) == 0 {
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
		p.err = p.client.listByProjectHandleError(resp)
		return false
	}
	result, err := p.client.listByProjectHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MachinesListByProjectResponse page.
func (p *MachinesListByProjectPager) PageResponse() MachinesListByProjectResponse {
	return p.current
}

// ProjectsListBySubscriptionPager provides operations for iterating over paged responses.
type ProjectsListBySubscriptionPager struct {
	client    *ProjectsClient
	current   ProjectsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProjectsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ProjectsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ProjectsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProjectResultList.NextLink == nil || len(*p.current.ProjectResultList.NextLink) == 0 {
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

// PageResponse returns the current ProjectsListBySubscriptionResponse page.
func (p *ProjectsListBySubscriptionPager) PageResponse() ProjectsListBySubscriptionResponse {
	return p.current
}

// ProjectsListPager provides operations for iterating over paged responses.
type ProjectsListPager struct {
	client    *ProjectsClient
	current   ProjectsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProjectsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ProjectsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ProjectsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProjectResultList.NextLink == nil || len(*p.current.ProjectResultList.NextLink) == 0 {
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

// PageResponse returns the current ProjectsListResponse page.
func (p *ProjectsListPager) PageResponse() ProjectsListResponse {
	return p.current
}
