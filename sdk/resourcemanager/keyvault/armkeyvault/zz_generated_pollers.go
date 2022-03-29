//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// MHSMPrivateEndpointConnectionsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type MHSMPrivateEndpointConnectionsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MHSMPrivateEndpointConnectionsClientDeletePoller) Done() bool {
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
func (p *MHSMPrivateEndpointConnectionsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MHSMPrivateEndpointConnectionsClientDeleteResponse will be returned.
func (p *MHSMPrivateEndpointConnectionsClientDeletePoller) FinalResponse(ctx context.Context) (MHSMPrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := MHSMPrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.MHSMPrivateEndpointConnection)
	if err != nil {
		return MHSMPrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MHSMPrivateEndpointConnectionsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagedHsmsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ManagedHsmsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagedHsmsClientCreateOrUpdatePoller) Done() bool {
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
func (p *ManagedHsmsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagedHsmsClientCreateOrUpdateResponse will be returned.
func (p *ManagedHsmsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ManagedHsmsClientCreateOrUpdateResponse, error) {
	respType := ManagedHsmsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ManagedHsm)
	if err != nil {
		return ManagedHsmsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagedHsmsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagedHsmsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ManagedHsmsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagedHsmsClientDeletePoller) Done() bool {
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
func (p *ManagedHsmsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagedHsmsClientDeleteResponse will be returned.
func (p *ManagedHsmsClientDeletePoller) FinalResponse(ctx context.Context) (ManagedHsmsClientDeleteResponse, error) {
	respType := ManagedHsmsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ManagedHsmsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagedHsmsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagedHsmsClientPurgeDeletedPoller provides polling facilities until the operation reaches a terminal state.
type ManagedHsmsClientPurgeDeletedPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagedHsmsClientPurgeDeletedPoller) Done() bool {
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
func (p *ManagedHsmsClientPurgeDeletedPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagedHsmsClientPurgeDeletedResponse will be returned.
func (p *ManagedHsmsClientPurgeDeletedPoller) FinalResponse(ctx context.Context) (ManagedHsmsClientPurgeDeletedResponse, error) {
	respType := ManagedHsmsClientPurgeDeletedResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ManagedHsmsClientPurgeDeletedResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagedHsmsClientPurgeDeletedPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagedHsmsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ManagedHsmsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagedHsmsClientUpdatePoller) Done() bool {
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
func (p *ManagedHsmsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagedHsmsClientUpdateResponse will be returned.
func (p *ManagedHsmsClientUpdatePoller) FinalResponse(ctx context.Context) (ManagedHsmsClientUpdateResponse, error) {
	respType := ManagedHsmsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ManagedHsm)
	if err != nil {
		return ManagedHsmsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagedHsmsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsClientDeletePoller) Done() bool {
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
func (p *PrivateEndpointConnectionsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsClientDeleteResponse will be returned.
func (p *PrivateEndpointConnectionsClientDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VaultsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VaultsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VaultsClientCreateOrUpdatePoller) Done() bool {
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
func (p *VaultsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VaultsClientCreateOrUpdateResponse will be returned.
func (p *VaultsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VaultsClientCreateOrUpdateResponse, error) {
	respType := VaultsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Vault)
	if err != nil {
		return VaultsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VaultsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VaultsClientPurgeDeletedPoller provides polling facilities until the operation reaches a terminal state.
type VaultsClientPurgeDeletedPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VaultsClientPurgeDeletedPoller) Done() bool {
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
func (p *VaultsClientPurgeDeletedPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VaultsClientPurgeDeletedResponse will be returned.
func (p *VaultsClientPurgeDeletedPoller) FinalResponse(ctx context.Context) (VaultsClientPurgeDeletedResponse, error) {
	respType := VaultsClientPurgeDeletedResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VaultsClientPurgeDeletedResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VaultsClientPurgeDeletedPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}