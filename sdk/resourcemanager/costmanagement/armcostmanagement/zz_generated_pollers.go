//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// GenerateDetailedCostReportClientCreateOperationPoller provides polling facilities until the operation reaches a terminal state.
type GenerateDetailedCostReportClientCreateOperationPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *GenerateDetailedCostReportClientCreateOperationPoller) Done() bool {
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
func (p *GenerateDetailedCostReportClientCreateOperationPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final GenerateDetailedCostReportClientCreateOperationResponse will be returned.
func (p *GenerateDetailedCostReportClientCreateOperationPoller) FinalResponse(ctx context.Context) (GenerateDetailedCostReportClientCreateOperationResponse, error) {
	respType := GenerateDetailedCostReportClientCreateOperationResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.GenerateDetailedCostReportOperationResult)
	if err != nil {
		return GenerateDetailedCostReportClientCreateOperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *GenerateDetailedCostReportClientCreateOperationPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}