//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ServicesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ServicesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServicesClientCreateOrUpdatePoller) Done() bool {
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
func (p *ServicesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServicesClientCreateOrUpdateResponse will be returned.
func (p *ServicesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ServicesClientCreateOrUpdateResponse, error) {
	respType := ServicesClientCreateOrUpdateResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.Service)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServicesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SharedPrivateLinkResourcesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type SharedPrivateLinkResourcesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SharedPrivateLinkResourcesClientCreateOrUpdatePoller) Done() bool {
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
func (p *SharedPrivateLinkResourcesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SharedPrivateLinkResourcesClientCreateOrUpdateResponse will be returned.
func (p *SharedPrivateLinkResourcesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (SharedPrivateLinkResourcesClientCreateOrUpdateResponse, error) {
	respType := SharedPrivateLinkResourcesClientCreateOrUpdateResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.SharedPrivateLinkResource)
	if err != nil {
		return SharedPrivateLinkResourcesClientCreateOrUpdateResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SharedPrivateLinkResourcesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SharedPrivateLinkResourcesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type SharedPrivateLinkResourcesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SharedPrivateLinkResourcesClientDeletePoller) Done() bool {
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
func (p *SharedPrivateLinkResourcesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SharedPrivateLinkResourcesClientDeleteResponse will be returned.
func (p *SharedPrivateLinkResourcesClientDeletePoller) FinalResponse(ctx context.Context) (SharedPrivateLinkResourcesClientDeleteResponse, error) {
	respType := SharedPrivateLinkResourcesClientDeleteResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return SharedPrivateLinkResourcesClientDeleteResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SharedPrivateLinkResourcesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
