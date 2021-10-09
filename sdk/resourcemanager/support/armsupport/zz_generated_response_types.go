//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// CommunicationsCheckNameAvailabilityResponse contains the response from method Communications.CheckNameAvailability.
type CommunicationsCheckNameAvailabilityResponse struct {
	CommunicationsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommunicationsCheckNameAvailabilityResult contains the result from method Communications.CheckNameAvailability.
type CommunicationsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityOutput
}

// CommunicationsCreatePollerResponse contains the response from method Communications.Create.
type CommunicationsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CommunicationsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l CommunicationsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CommunicationsCreateResponse, error) {
	respType := CommunicationsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.CommunicationDetails)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CommunicationsCreatePollerResponse from the provided client and resume token.
func (l *CommunicationsCreatePollerResponse) Resume(ctx context.Context, client *CommunicationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CommunicationsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &CommunicationsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CommunicationsCreateResponse contains the response from method Communications.Create.
type CommunicationsCreateResponse struct {
	CommunicationsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommunicationsCreateResult contains the result from method Communications.Create.
type CommunicationsCreateResult struct {
	CommunicationDetails
}

// CommunicationsGetResponse contains the response from method Communications.Get.
type CommunicationsGetResponse struct {
	CommunicationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommunicationsGetResult contains the result from method Communications.Get.
type CommunicationsGetResult struct {
	CommunicationDetails
}

// CommunicationsListResponse contains the response from method Communications.List.
type CommunicationsListResponse struct {
	CommunicationsListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommunicationsListResultEnvelope contains the result from method Communications.List.
type CommunicationsListResultEnvelope struct {
	CommunicationsListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResultEnvelope contains the result from method Operations.List.
type OperationsListResultEnvelope struct {
	OperationsListResult
}

// ProblemClassificationsGetResponse contains the response from method ProblemClassifications.Get.
type ProblemClassificationsGetResponse struct {
	ProblemClassificationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProblemClassificationsGetResult contains the result from method ProblemClassifications.Get.
type ProblemClassificationsGetResult struct {
	ProblemClassification
}

// ProblemClassificationsListResponse contains the response from method ProblemClassifications.List.
type ProblemClassificationsListResponse struct {
	ProblemClassificationsListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProblemClassificationsListResultEnvelope contains the result from method ProblemClassifications.List.
type ProblemClassificationsListResultEnvelope struct {
	ProblemClassificationsListResult
}

// ServicesGetResponse contains the response from method Services.Get.
type ServicesGetResponse struct {
	ServicesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesGetResult contains the result from method Services.Get.
type ServicesGetResult struct {
	Service
}

// ServicesListResponse contains the response from method Services.List.
type ServicesListResponse struct {
	ServicesListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesListResultEnvelope contains the result from method Services.List.
type ServicesListResultEnvelope struct {
	ServicesListResult
}

// SupportTicketsCheckNameAvailabilityResponse contains the response from method SupportTickets.CheckNameAvailability.
type SupportTicketsCheckNameAvailabilityResponse struct {
	SupportTicketsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SupportTicketsCheckNameAvailabilityResult contains the result from method SupportTickets.CheckNameAvailability.
type SupportTicketsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityOutput
}

// SupportTicketsCreatePollerResponse contains the response from method SupportTickets.Create.
type SupportTicketsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SupportTicketsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l SupportTicketsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SupportTicketsCreateResponse, error) {
	respType := SupportTicketsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SupportTicketDetails)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SupportTicketsCreatePollerResponse from the provided client and resume token.
func (l *SupportTicketsCreatePollerResponse) Resume(ctx context.Context, client *SupportTicketsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SupportTicketsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &SupportTicketsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SupportTicketsCreateResponse contains the response from method SupportTickets.Create.
type SupportTicketsCreateResponse struct {
	SupportTicketsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SupportTicketsCreateResult contains the result from method SupportTickets.Create.
type SupportTicketsCreateResult struct {
	SupportTicketDetails
}

// SupportTicketsGetResponse contains the response from method SupportTickets.Get.
type SupportTicketsGetResponse struct {
	SupportTicketsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SupportTicketsGetResult contains the result from method SupportTickets.Get.
type SupportTicketsGetResult struct {
	SupportTicketDetails
}

// SupportTicketsListResponse contains the response from method SupportTickets.List.
type SupportTicketsListResponse struct {
	SupportTicketsListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SupportTicketsListResultEnvelope contains the result from method SupportTickets.List.
type SupportTicketsListResultEnvelope struct {
	SupportTicketsListResult
}

// SupportTicketsUpdateResponse contains the response from method SupportTickets.Update.
type SupportTicketsUpdateResponse struct {
	SupportTicketsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SupportTicketsUpdateResult contains the result from method SupportTickets.Update.
type SupportTicketsUpdateResult struct {
	SupportTicketDetails
}
