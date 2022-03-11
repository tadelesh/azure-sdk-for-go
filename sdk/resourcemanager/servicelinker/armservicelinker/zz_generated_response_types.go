//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// LinkerClientCreateOrUpdatePollerResponse contains the response from method LinkerClient.CreateOrUpdate.
type LinkerClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *LinkerClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l LinkerClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (LinkerClientCreateOrUpdateResponse, error) {
	respType := LinkerClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.LinkerResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a LinkerClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *LinkerClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *LinkerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("LinkerClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &LinkerClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// LinkerClientCreateOrUpdateResponse contains the response from method LinkerClient.CreateOrUpdate.
type LinkerClientCreateOrUpdateResponse struct {
	LinkerResource
}

// LinkerClientDeletePollerResponse contains the response from method LinkerClient.Delete.
type LinkerClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *LinkerClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l LinkerClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (LinkerClientDeleteResponse, error) {
	respType := LinkerClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a LinkerClientDeletePollerResponse from the provided client and resume token.
func (l *LinkerClientDeletePollerResponse) Resume(ctx context.Context, client *LinkerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("LinkerClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &LinkerClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// LinkerClientDeleteResponse contains the response from method LinkerClient.Delete.
type LinkerClientDeleteResponse struct {
	// placeholder for future response values
}

// LinkerClientGetResponse contains the response from method LinkerClient.Get.
type LinkerClientGetResponse struct {
	LinkerResource
}

// LinkerClientListConfigurationsResponse contains the response from method LinkerClient.ListConfigurations.
type LinkerClientListConfigurationsResponse struct {
	SourceConfigurationResult
}

// LinkerClientListResponse contains the response from method LinkerClient.List.
type LinkerClientListResponse struct {
	LinkerList
}

// LinkerClientUpdatePollerResponse contains the response from method LinkerClient.Update.
type LinkerClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *LinkerClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l LinkerClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (LinkerClientUpdateResponse, error) {
	respType := LinkerClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.LinkerResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a LinkerClientUpdatePollerResponse from the provided client and resume token.
func (l *LinkerClientUpdatePollerResponse) Resume(ctx context.Context, client *LinkerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("LinkerClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &LinkerClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// LinkerClientUpdateResponse contains the response from method LinkerClient.Update.
type LinkerClientUpdateResponse struct {
	LinkerResource
}

// LinkerClientValidatePollerResponse contains the response from method LinkerClient.Validate.
type LinkerClientValidatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *LinkerClientValidatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l LinkerClientValidatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (LinkerClientValidateResponse, error) {
	respType := LinkerClientValidateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ValidateResult)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a LinkerClientValidatePollerResponse from the provided client and resume token.
func (l *LinkerClientValidatePollerResponse) Resume(ctx context.Context, client *LinkerClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("LinkerClient.Validate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &LinkerClientValidatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// LinkerClientValidateResponse contains the response from method LinkerClient.Validate.
type LinkerClientValidateResponse struct {
	ValidateResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}
