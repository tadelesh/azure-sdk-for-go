//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ManagementClientCreateAddressPoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientCreateAddressPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientCreateAddressPoller) Done() bool {
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
func (p *ManagementClientCreateAddressPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientCreateAddressResponse will be returned.
func (p *ManagementClientCreateAddressPoller) FinalResponse(ctx context.Context) (ManagementClientCreateAddressResponse, error) {
	respType := ManagementClientCreateAddressResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.AddressResource)
	if err != nil {
		return ManagementClientCreateAddressResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientCreateAddressPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientCreateOrderItemPoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientCreateOrderItemPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientCreateOrderItemPoller) Done() bool {
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
func (p *ManagementClientCreateOrderItemPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientCreateOrderItemResponse will be returned.
func (p *ManagementClientCreateOrderItemPoller) FinalResponse(ctx context.Context) (ManagementClientCreateOrderItemResponse, error) {
	respType := ManagementClientCreateOrderItemResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.OrderItemResource)
	if err != nil {
		return ManagementClientCreateOrderItemResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientCreateOrderItemPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientDeleteAddressByNamePoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientDeleteAddressByNamePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientDeleteAddressByNamePoller) Done() bool {
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
func (p *ManagementClientDeleteAddressByNamePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientDeleteAddressByNameResponse will be returned.
func (p *ManagementClientDeleteAddressByNamePoller) FinalResponse(ctx context.Context) (ManagementClientDeleteAddressByNameResponse, error) {
	respType := ManagementClientDeleteAddressByNameResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ManagementClientDeleteAddressByNameResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientDeleteAddressByNamePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientDeleteOrderItemByNamePoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientDeleteOrderItemByNamePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientDeleteOrderItemByNamePoller) Done() bool {
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
func (p *ManagementClientDeleteOrderItemByNamePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientDeleteOrderItemByNameResponse will be returned.
func (p *ManagementClientDeleteOrderItemByNamePoller) FinalResponse(ctx context.Context) (ManagementClientDeleteOrderItemByNameResponse, error) {
	respType := ManagementClientDeleteOrderItemByNameResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ManagementClientDeleteOrderItemByNameResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientDeleteOrderItemByNamePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientReturnOrderItemPoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientReturnOrderItemPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientReturnOrderItemPoller) Done() bool {
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
func (p *ManagementClientReturnOrderItemPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientReturnOrderItemResponse will be returned.
func (p *ManagementClientReturnOrderItemPoller) FinalResponse(ctx context.Context) (ManagementClientReturnOrderItemResponse, error) {
	respType := ManagementClientReturnOrderItemResponse{}
	_, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ManagementClientReturnOrderItemResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientReturnOrderItemPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientUpdateAddressPoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientUpdateAddressPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientUpdateAddressPoller) Done() bool {
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
func (p *ManagementClientUpdateAddressPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientUpdateAddressResponse will be returned.
func (p *ManagementClientUpdateAddressPoller) FinalResponse(ctx context.Context) (ManagementClientUpdateAddressResponse, error) {
	respType := ManagementClientUpdateAddressResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.AddressResource)
	if err != nil {
		return ManagementClientUpdateAddressResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientUpdateAddressPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ManagementClientUpdateOrderItemPoller provides polling facilities until the operation reaches a terminal state.
type ManagementClientUpdateOrderItemPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ManagementClientUpdateOrderItemPoller) Done() bool {
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
func (p *ManagementClientUpdateOrderItemPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ManagementClientUpdateOrderItemResponse will be returned.
func (p *ManagementClientUpdateOrderItemPoller) FinalResponse(ctx context.Context) (ManagementClientUpdateOrderItemResponse, error) {
	respType := ManagementClientUpdateOrderItemResponse{}
	_, err := p.pt.FinalResponse(ctx, &respType.OrderItemResource)
	if err != nil {
		return ManagementClientUpdateOrderItemResponse{}, err
	}
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ManagementClientUpdateOrderItemPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
