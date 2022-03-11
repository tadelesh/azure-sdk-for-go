//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// RegistrationAssignmentsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type RegistrationAssignmentsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RegistrationAssignmentsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RegistrationAssignmentsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RegistrationAssignmentsClientCreateOrUpdateResponse will be returned.
func (p *RegistrationAssignmentsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (RegistrationAssignmentsClientCreateOrUpdateResponse, error) {
	respType := RegistrationAssignmentsClientCreateOrUpdateResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.RegistrationAssignment)
	if err != nil {
		return RegistrationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RegistrationAssignmentsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RegistrationAssignmentsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type RegistrationAssignmentsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RegistrationAssignmentsClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RegistrationAssignmentsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RegistrationAssignmentsClientDeleteResponse will be returned.
func (p *RegistrationAssignmentsClientDeletePoller) FinalResponse(ctx context.Context) (RegistrationAssignmentsClientDeleteResponse, error) {
	respType := RegistrationAssignmentsClientDeleteResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RegistrationAssignmentsClientDeleteResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RegistrationAssignmentsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RegistrationDefinitionsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type RegistrationDefinitionsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RegistrationDefinitionsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RegistrationDefinitionsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RegistrationDefinitionsClientCreateOrUpdateResponse will be returned.
func (p *RegistrationDefinitionsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (RegistrationDefinitionsClientCreateOrUpdateResponse, error) {
	respType := RegistrationDefinitionsClientCreateOrUpdateResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.RegistrationDefinition)
	if err != nil {
		return RegistrationDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RegistrationDefinitionsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
