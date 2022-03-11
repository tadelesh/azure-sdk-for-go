//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// ClusterExtensionTypeClientGetResponse contains the response from method ClusterExtensionTypeClient.Get.
type ClusterExtensionTypeClientGetResponse struct {
	ExtensionType
}

// ClusterExtensionTypesClientListResponse contains the response from method ClusterExtensionTypesClient.List.
type ClusterExtensionTypesClientListResponse struct {
	ExtensionTypeList
}

// ExtensionTypeVersionsClientListResponse contains the response from method ExtensionTypeVersionsClient.List.
type ExtensionTypeVersionsClientListResponse struct {
	ExtensionVersionList
}

// ExtensionsClientCreatePollerResponse contains the response from method ExtensionsClient.Create.
type ExtensionsClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsClientCreatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsClientCreateResponse, error) {
	respType := ExtensionsClientCreateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Extension)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExtensionsClientCreatePollerResponse from the provided client and resume token.
func (l *ExtensionsClientCreatePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExtensionsClientCreatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExtensionsClientCreateResponse contains the response from method ExtensionsClient.Create.
type ExtensionsClientCreateResponse struct {
	Extension
}

// ExtensionsClientDeletePollerResponse contains the response from method ExtensionsClient.Delete.
type ExtensionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsClientDeleteResponse, error) {
	respType := ExtensionsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExtensionsClientDeletePollerResponse from the provided client and resume token.
func (l *ExtensionsClientDeletePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExtensionsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExtensionsClientDeleteResponse contains the response from method ExtensionsClient.Delete.
type ExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExtensionsClientGetResponse contains the response from method ExtensionsClient.Get.
type ExtensionsClientGetResponse struct {
	Extension
}

// ExtensionsClientListResponse contains the response from method ExtensionsClient.List.
type ExtensionsClientListResponse struct {
	ExtensionsList
}

// ExtensionsClientUpdatePollerResponse contains the response from method ExtensionsClient.Update.
type ExtensionsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsClientUpdateResponse, error) {
	respType := ExtensionsClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Extension)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExtensionsClientUpdatePollerResponse from the provided client and resume token.
func (l *ExtensionsClientUpdatePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExtensionsClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExtensionsClientUpdateResponse contains the response from method ExtensionsClient.Update.
type ExtensionsClientUpdateResponse struct {
	Extension
}

// FluxConfigOperationStatusClientGetResponse contains the response from method FluxConfigOperationStatusClient.Get.
type FluxConfigOperationStatusClientGetResponse struct {
	OperationStatusResult
}

// FluxConfigurationsClientCreateOrUpdatePollerResponse contains the response from method FluxConfigurationsClient.CreateOrUpdate.
type FluxConfigurationsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FluxConfigurationsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FluxConfigurationsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FluxConfigurationsClientCreateOrUpdateResponse, error) {
	respType := FluxConfigurationsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.FluxConfiguration)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FluxConfigurationsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *FluxConfigurationsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *FluxConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FluxConfigurationsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FluxConfigurationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FluxConfigurationsClientCreateOrUpdateResponse contains the response from method FluxConfigurationsClient.CreateOrUpdate.
type FluxConfigurationsClientCreateOrUpdateResponse struct {
	FluxConfiguration
}

// FluxConfigurationsClientDeletePollerResponse contains the response from method FluxConfigurationsClient.Delete.
type FluxConfigurationsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FluxConfigurationsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FluxConfigurationsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FluxConfigurationsClientDeleteResponse, error) {
	respType := FluxConfigurationsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FluxConfigurationsClientDeletePollerResponse from the provided client and resume token.
func (l *FluxConfigurationsClientDeletePollerResponse) Resume(ctx context.Context, client *FluxConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FluxConfigurationsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FluxConfigurationsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FluxConfigurationsClientDeleteResponse contains the response from method FluxConfigurationsClient.Delete.
type FluxConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// FluxConfigurationsClientGetResponse contains the response from method FluxConfigurationsClient.Get.
type FluxConfigurationsClientGetResponse struct {
	FluxConfiguration
}

// FluxConfigurationsClientListResponse contains the response from method FluxConfigurationsClient.List.
type FluxConfigurationsClientListResponse struct {
	FluxConfigurationsList
}

// FluxConfigurationsClientUpdatePollerResponse contains the response from method FluxConfigurationsClient.Update.
type FluxConfigurationsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FluxConfigurationsClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FluxConfigurationsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FluxConfigurationsClientUpdateResponse, error) {
	respType := FluxConfigurationsClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.FluxConfiguration)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FluxConfigurationsClientUpdatePollerResponse from the provided client and resume token.
func (l *FluxConfigurationsClientUpdatePollerResponse) Resume(ctx context.Context, client *FluxConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FluxConfigurationsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FluxConfigurationsClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FluxConfigurationsClientUpdateResponse contains the response from method FluxConfigurationsClient.Update.
type FluxConfigurationsClientUpdateResponse struct {
	FluxConfiguration
}

// LocationExtensionTypesClientListResponse contains the response from method LocationExtensionTypesClient.List.
type LocationExtensionTypesClientListResponse struct {
	ExtensionTypeList
}

// OperationStatusClientGetResponse contains the response from method OperationStatusClient.Get.
type OperationStatusClientGetResponse struct {
	OperationStatusResult
}

// OperationStatusClientListResponse contains the response from method OperationStatusClient.List.
type OperationStatusClientListResponse struct {
	OperationStatusList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ResourceProviderOperationList
}

// SourceControlConfigurationsClientCreateOrUpdateResponse contains the response from method SourceControlConfigurationsClient.CreateOrUpdate.
type SourceControlConfigurationsClientCreateOrUpdateResponse struct {
	SourceControlConfiguration
}

// SourceControlConfigurationsClientDeletePollerResponse contains the response from method SourceControlConfigurationsClient.Delete.
type SourceControlConfigurationsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SourceControlConfigurationsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SourceControlConfigurationsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SourceControlConfigurationsClientDeleteResponse, error) {
	respType := SourceControlConfigurationsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a SourceControlConfigurationsClientDeletePollerResponse from the provided client and resume token.
func (l *SourceControlConfigurationsClientDeletePollerResponse) Resume(ctx context.Context, client *SourceControlConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SourceControlConfigurationsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SourceControlConfigurationsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// SourceControlConfigurationsClientDeleteResponse contains the response from method SourceControlConfigurationsClient.Delete.
type SourceControlConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// SourceControlConfigurationsClientGetResponse contains the response from method SourceControlConfigurationsClient.Get.
type SourceControlConfigurationsClientGetResponse struct {
	SourceControlConfiguration
}

// SourceControlConfigurationsClientListResponse contains the response from method SourceControlConfigurationsClient.List.
type SourceControlConfigurationsClientListResponse struct {
	SourceControlConfigurationList
}
