//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnc

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// ControllerClientCreatePollerResponse contains the response from method ControllerClient.Create.
type ControllerClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ControllerClientCreatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ControllerClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ControllerClientCreateResponse, error) {
	respType := ControllerClientCreateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DelegatedController)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ControllerClientCreatePollerResponse from the provided client and resume token.
func (l *ControllerClientCreatePollerResponse) Resume(ctx context.Context, client *ControllerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ControllerClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ControllerClientCreatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ControllerClientCreateResponse contains the response from method ControllerClient.Create.
type ControllerClientCreateResponse struct {
	DelegatedController
}

// ControllerClientDeletePollerResponse contains the response from method ControllerClient.Delete.
type ControllerClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ControllerClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ControllerClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ControllerClientDeleteResponse, error) {
	respType := ControllerClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ControllerClientDeletePollerResponse from the provided client and resume token.
func (l *ControllerClientDeletePollerResponse) Resume(ctx context.Context, client *ControllerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ControllerClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ControllerClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ControllerClientDeleteResponse contains the response from method ControllerClient.Delete.
type ControllerClientDeleteResponse struct {
	// placeholder for future response values
}

// ControllerClientGetDetailsResponse contains the response from method ControllerClient.GetDetails.
type ControllerClientGetDetailsResponse struct {
	DelegatedController
}

// ControllerClientPatchResponse contains the response from method ControllerClient.Patch.
type ControllerClientPatchResponse struct {
	DelegatedController
}

// DelegatedNetworkClientListByResourceGroupResponse contains the response from method DelegatedNetworkClient.ListByResourceGroup.
type DelegatedNetworkClientListByResourceGroupResponse struct {
	DelegatedControllers
}

// DelegatedNetworkClientListBySubscriptionResponse contains the response from method DelegatedNetworkClient.ListBySubscription.
type DelegatedNetworkClientListBySubscriptionResponse struct {
	DelegatedControllers
}

// DelegatedSubnetServiceClientDeleteDetailsPollerResponse contains the response from method DelegatedSubnetServiceClient.DeleteDetails.
type DelegatedSubnetServiceClientDeleteDetailsPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DelegatedSubnetServiceClientDeleteDetailsPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DelegatedSubnetServiceClientDeleteDetailsPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DelegatedSubnetServiceClientDeleteDetailsResponse, error) {
	respType := DelegatedSubnetServiceClientDeleteDetailsResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a DelegatedSubnetServiceClientDeleteDetailsPollerResponse from the provided client and resume token.
func (l *DelegatedSubnetServiceClientDeleteDetailsPollerResponse) Resume(ctx context.Context, client *DelegatedSubnetServiceClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DelegatedSubnetServiceClient.DeleteDetails", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DelegatedSubnetServiceClientDeleteDetailsPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// DelegatedSubnetServiceClientDeleteDetailsResponse contains the response from method DelegatedSubnetServiceClient.DeleteDetails.
type DelegatedSubnetServiceClientDeleteDetailsResponse struct {
	// placeholder for future response values
}

// DelegatedSubnetServiceClientGetDetailsResponse contains the response from method DelegatedSubnetServiceClient.GetDetails.
type DelegatedSubnetServiceClientGetDetailsResponse struct {
	DelegatedSubnet
}

// DelegatedSubnetServiceClientListByResourceGroupResponse contains the response from method DelegatedSubnetServiceClient.ListByResourceGroup.
type DelegatedSubnetServiceClientListByResourceGroupResponse struct {
	DelegatedSubnets
}

// DelegatedSubnetServiceClientListBySubscriptionResponse contains the response from method DelegatedSubnetServiceClient.ListBySubscription.
type DelegatedSubnetServiceClientListBySubscriptionResponse struct {
	DelegatedSubnets
}

// DelegatedSubnetServiceClientPatchDetailsPollerResponse contains the response from method DelegatedSubnetServiceClient.PatchDetails.
type DelegatedSubnetServiceClientPatchDetailsPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DelegatedSubnetServiceClientPatchDetailsPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DelegatedSubnetServiceClientPatchDetailsPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DelegatedSubnetServiceClientPatchDetailsResponse, error) {
	respType := DelegatedSubnetServiceClientPatchDetailsResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DelegatedSubnet)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a DelegatedSubnetServiceClientPatchDetailsPollerResponse from the provided client and resume token.
func (l *DelegatedSubnetServiceClientPatchDetailsPollerResponse) Resume(ctx context.Context, client *DelegatedSubnetServiceClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DelegatedSubnetServiceClient.PatchDetails", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DelegatedSubnetServiceClientPatchDetailsPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// DelegatedSubnetServiceClientPatchDetailsResponse contains the response from method DelegatedSubnetServiceClient.PatchDetails.
type DelegatedSubnetServiceClientPatchDetailsResponse struct {
	DelegatedSubnet
}

// DelegatedSubnetServiceClientPutDetailsPollerResponse contains the response from method DelegatedSubnetServiceClient.PutDetails.
type DelegatedSubnetServiceClientPutDetailsPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DelegatedSubnetServiceClientPutDetailsPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DelegatedSubnetServiceClientPutDetailsPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DelegatedSubnetServiceClientPutDetailsResponse, error) {
	respType := DelegatedSubnetServiceClientPutDetailsResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DelegatedSubnet)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a DelegatedSubnetServiceClientPutDetailsPollerResponse from the provided client and resume token.
func (l *DelegatedSubnetServiceClientPutDetailsPollerResponse) Resume(ctx context.Context, client *DelegatedSubnetServiceClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DelegatedSubnetServiceClient.PutDetails", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DelegatedSubnetServiceClientPutDetailsPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// DelegatedSubnetServiceClientPutDetailsResponse contains the response from method DelegatedSubnetServiceClient.PutDetails.
type DelegatedSubnetServiceClientPutDetailsResponse struct {
	DelegatedSubnet
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// OrchestratorInstanceServiceClientCreatePollerResponse contains the response from method OrchestratorInstanceServiceClient.Create.
type OrchestratorInstanceServiceClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OrchestratorInstanceServiceClientCreatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l OrchestratorInstanceServiceClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OrchestratorInstanceServiceClientCreateResponse, error) {
	respType := OrchestratorInstanceServiceClientCreateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Orchestrator)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a OrchestratorInstanceServiceClientCreatePollerResponse from the provided client and resume token.
func (l *OrchestratorInstanceServiceClientCreatePollerResponse) Resume(ctx context.Context, client *OrchestratorInstanceServiceClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OrchestratorInstanceServiceClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &OrchestratorInstanceServiceClientCreatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// OrchestratorInstanceServiceClientCreateResponse contains the response from method OrchestratorInstanceServiceClient.Create.
type OrchestratorInstanceServiceClientCreateResponse struct {
	Orchestrator
}

// OrchestratorInstanceServiceClientDeletePollerResponse contains the response from method OrchestratorInstanceServiceClient.Delete.
type OrchestratorInstanceServiceClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OrchestratorInstanceServiceClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l OrchestratorInstanceServiceClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OrchestratorInstanceServiceClientDeleteResponse, error) {
	respType := OrchestratorInstanceServiceClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a OrchestratorInstanceServiceClientDeletePollerResponse from the provided client and resume token.
func (l *OrchestratorInstanceServiceClientDeletePollerResponse) Resume(ctx context.Context, client *OrchestratorInstanceServiceClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OrchestratorInstanceServiceClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &OrchestratorInstanceServiceClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// OrchestratorInstanceServiceClientDeleteResponse contains the response from method OrchestratorInstanceServiceClient.Delete.
type OrchestratorInstanceServiceClientDeleteResponse struct {
	// placeholder for future response values
}

// OrchestratorInstanceServiceClientGetDetailsResponse contains the response from method OrchestratorInstanceServiceClient.GetDetails.
type OrchestratorInstanceServiceClientGetDetailsResponse struct {
	Orchestrator
}

// OrchestratorInstanceServiceClientListByResourceGroupResponse contains the response from method OrchestratorInstanceServiceClient.ListByResourceGroup.
type OrchestratorInstanceServiceClientListByResourceGroupResponse struct {
	Orchestrators
}

// OrchestratorInstanceServiceClientListBySubscriptionResponse contains the response from method OrchestratorInstanceServiceClient.ListBySubscription.
type OrchestratorInstanceServiceClientListBySubscriptionResponse struct {
	Orchestrators
}

// OrchestratorInstanceServiceClientPatchResponse contains the response from method OrchestratorInstanceServiceClient.Patch.
type OrchestratorInstanceServiceClientPatchResponse struct {
	Orchestrator
}
