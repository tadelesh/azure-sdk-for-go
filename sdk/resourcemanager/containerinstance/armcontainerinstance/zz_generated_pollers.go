//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ContainerGroupsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ContainerGroupsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ContainerGroupsClientCreateOrUpdatePoller) Done() bool {
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
func (p *ContainerGroupsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ContainerGroupsClientCreateOrUpdateResponse will be returned.
func (p *ContainerGroupsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ContainerGroupsClientCreateOrUpdateResponse, error) {
	respType := ContainerGroupsClientCreateOrUpdateResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.ContainerGroup)
	if err != nil {
		return ContainerGroupsClientCreateOrUpdateResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ContainerGroupsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ContainerGroupsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ContainerGroupsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ContainerGroupsClientDeletePoller) Done() bool {
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
func (p *ContainerGroupsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ContainerGroupsClientDeleteResponse will be returned.
func (p *ContainerGroupsClientDeletePoller) FinalResponse(ctx context.Context) (ContainerGroupsClientDeleteResponse, error) {
	respType := ContainerGroupsClientDeleteResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.ContainerGroup)
	if err != nil {
		return ContainerGroupsClientDeleteResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ContainerGroupsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ContainerGroupsClientRestartPoller provides polling facilities until the operation reaches a terminal state.
type ContainerGroupsClientRestartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ContainerGroupsClientRestartPoller) Done() bool {
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
func (p *ContainerGroupsClientRestartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ContainerGroupsClientRestartResponse will be returned.
func (p *ContainerGroupsClientRestartPoller) FinalResponse(ctx context.Context) (ContainerGroupsClientRestartResponse, error) {
	respType := ContainerGroupsClientRestartResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ContainerGroupsClientRestartResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ContainerGroupsClientRestartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ContainerGroupsClientStartPoller provides polling facilities until the operation reaches a terminal state.
type ContainerGroupsClientStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ContainerGroupsClientStartPoller) Done() bool {
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
func (p *ContainerGroupsClientStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ContainerGroupsClientStartResponse will be returned.
func (p *ContainerGroupsClientStartPoller) FinalResponse(ctx context.Context) (ContainerGroupsClientStartResponse, error) {
	respType := ContainerGroupsClientStartResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ContainerGroupsClientStartResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ContainerGroupsClientStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
