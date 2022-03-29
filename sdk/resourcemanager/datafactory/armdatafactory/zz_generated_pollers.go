//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DataFlowDebugSessionClientCreatePoller provides polling facilities until the operation reaches a terminal state.
type DataFlowDebugSessionClientCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DataFlowDebugSessionClientCreatePoller) Done() bool {
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
func (p *DataFlowDebugSessionClientCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DataFlowDebugSessionClientCreateResponse will be returned.
func (p *DataFlowDebugSessionClientCreatePoller) FinalResponse(ctx context.Context) (DataFlowDebugSessionClientCreateResponse, error) {
	respType := DataFlowDebugSessionClientCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.CreateDataFlowDebugSessionResponse)
	if err != nil {
		return DataFlowDebugSessionClientCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DataFlowDebugSessionClientCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DataFlowDebugSessionClientExecuteCommandPoller provides polling facilities until the operation reaches a terminal state.
type DataFlowDebugSessionClientExecuteCommandPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DataFlowDebugSessionClientExecuteCommandPoller) Done() bool {
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
func (p *DataFlowDebugSessionClientExecuteCommandPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DataFlowDebugSessionClientExecuteCommandResponse will be returned.
func (p *DataFlowDebugSessionClientExecuteCommandPoller) FinalResponse(ctx context.Context) (DataFlowDebugSessionClientExecuteCommandResponse, error) {
	respType := DataFlowDebugSessionClientExecuteCommandResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DataFlowDebugCommandResponse)
	if err != nil {
		return DataFlowDebugSessionClientExecuteCommandResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DataFlowDebugSessionClientExecuteCommandPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IntegrationRuntimeObjectMetadataClientRefreshPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimeObjectMetadataClientRefreshPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IntegrationRuntimeObjectMetadataClientRefreshPoller) Done() bool {
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
func (p *IntegrationRuntimeObjectMetadataClientRefreshPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IntegrationRuntimeObjectMetadataClientRefreshResponse will be returned.
func (p *IntegrationRuntimeObjectMetadataClientRefreshPoller) FinalResponse(ctx context.Context) (IntegrationRuntimeObjectMetadataClientRefreshResponse, error) {
	respType := IntegrationRuntimeObjectMetadataClientRefreshResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SsisObjectMetadataStatusResponse)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientRefreshResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IntegrationRuntimeObjectMetadataClientRefreshPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IntegrationRuntimesClientStartPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimesClientStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IntegrationRuntimesClientStartPoller) Done() bool {
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
func (p *IntegrationRuntimesClientStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IntegrationRuntimesClientStartResponse will be returned.
func (p *IntegrationRuntimesClientStartPoller) FinalResponse(ctx context.Context) (IntegrationRuntimesClientStartResponse, error) {
	respType := IntegrationRuntimesClientStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.IntegrationRuntimeStatusResponse)
	if err != nil {
		return IntegrationRuntimesClientStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IntegrationRuntimesClientStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IntegrationRuntimesClientStopPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimesClientStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IntegrationRuntimesClientStopPoller) Done() bool {
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
func (p *IntegrationRuntimesClientStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IntegrationRuntimesClientStopResponse will be returned.
func (p *IntegrationRuntimesClientStopPoller) FinalResponse(ctx context.Context) (IntegrationRuntimesClientStopResponse, error) {
	respType := IntegrationRuntimesClientStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return IntegrationRuntimesClientStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IntegrationRuntimesClientStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// TriggersClientStartPoller provides polling facilities until the operation reaches a terminal state.
type TriggersClientStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *TriggersClientStartPoller) Done() bool {
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
func (p *TriggersClientStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final TriggersClientStartResponse will be returned.
func (p *TriggersClientStartPoller) FinalResponse(ctx context.Context) (TriggersClientStartResponse, error) {
	respType := TriggersClientStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return TriggersClientStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *TriggersClientStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// TriggersClientStopPoller provides polling facilities until the operation reaches a terminal state.
type TriggersClientStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *TriggersClientStopPoller) Done() bool {
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
func (p *TriggersClientStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final TriggersClientStopResponse will be returned.
func (p *TriggersClientStopPoller) FinalResponse(ctx context.Context) (TriggersClientStopResponse, error) {
	respType := TriggersClientStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return TriggersClientStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *TriggersClientStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// TriggersClientSubscribeToEventsPoller provides polling facilities until the operation reaches a terminal state.
type TriggersClientSubscribeToEventsPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *TriggersClientSubscribeToEventsPoller) Done() bool {
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
func (p *TriggersClientSubscribeToEventsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final TriggersClientSubscribeToEventsResponse will be returned.
func (p *TriggersClientSubscribeToEventsPoller) FinalResponse(ctx context.Context) (TriggersClientSubscribeToEventsResponse, error) {
	respType := TriggersClientSubscribeToEventsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersClientSubscribeToEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *TriggersClientSubscribeToEventsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// TriggersClientUnsubscribeFromEventsPoller provides polling facilities until the operation reaches a terminal state.
type TriggersClientUnsubscribeFromEventsPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *TriggersClientUnsubscribeFromEventsPoller) Done() bool {
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
func (p *TriggersClientUnsubscribeFromEventsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final TriggersClientUnsubscribeFromEventsResponse will be returned.
func (p *TriggersClientUnsubscribeFromEventsPoller) FinalResponse(ctx context.Context) (TriggersClientUnsubscribeFromEventsResponse, error) {
	respType := TriggersClientUnsubscribeFromEventsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersClientUnsubscribeFromEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *TriggersClientUnsubscribeFromEventsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}