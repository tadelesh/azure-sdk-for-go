//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// EndpointsClientPurgeContentPollerResponse contains the response from method EndpointsClient.PurgeContent.
type EndpointsClientPurgeContentPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *EndpointsClientPurgeContentPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l EndpointsClientPurgeContentPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (EndpointsClientPurgeContentResponse, error) {
	respType := EndpointsClientPurgeContentResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a EndpointsClientPurgeContentPollerResponse from the provided client and resume token.
func (l *EndpointsClientPurgeContentPollerResponse) Resume(ctx context.Context, client *EndpointsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("EndpointsClient.PurgeContent", token, client.pl)
	if err != nil {
		return err
	}
	poller := &EndpointsClientPurgeContentPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// EndpointsClientPurgeContentResponse contains the response from method EndpointsClient.PurgeContent.
type EndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// ExperimentsClientCreateOrUpdatePollerResponse contains the response from method ExperimentsClient.CreateOrUpdate.
type ExperimentsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExperimentsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExperimentsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExperimentsClientCreateOrUpdateResponse, error) {
	respType := ExperimentsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Experiment)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExperimentsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ExperimentsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ExperimentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExperimentsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExperimentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExperimentsClientCreateOrUpdateResponse contains the response from method ExperimentsClient.CreateOrUpdate.
type ExperimentsClientCreateOrUpdateResponse struct {
	Experiment
}

// ExperimentsClientDeletePollerResponse contains the response from method ExperimentsClient.Delete.
type ExperimentsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExperimentsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExperimentsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExperimentsClientDeleteResponse, error) {
	respType := ExperimentsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExperimentsClientDeletePollerResponse from the provided client and resume token.
func (l *ExperimentsClientDeletePollerResponse) Resume(ctx context.Context, client *ExperimentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExperimentsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExperimentsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExperimentsClientDeleteResponse contains the response from method ExperimentsClient.Delete.
type ExperimentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExperimentsClientGetResponse contains the response from method ExperimentsClient.Get.
type ExperimentsClientGetResponse struct {
	Experiment
}

// ExperimentsClientListByProfileResponse contains the response from method ExperimentsClient.ListByProfile.
type ExperimentsClientListByProfileResponse struct {
	ExperimentList
}

// ExperimentsClientUpdatePollerResponse contains the response from method ExperimentsClient.Update.
type ExperimentsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExperimentsClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExperimentsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExperimentsClientUpdateResponse, error) {
	respType := ExperimentsClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Experiment)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ExperimentsClientUpdatePollerResponse from the provided client and resume token.
func (l *ExperimentsClientUpdatePollerResponse) Resume(ctx context.Context, client *ExperimentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExperimentsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ExperimentsClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ExperimentsClientUpdateResponse contains the response from method ExperimentsClient.Update.
type ExperimentsClientUpdateResponse struct {
	Experiment
}

// FrontDoorsClientCreateOrUpdatePollerResponse contains the response from method FrontDoorsClient.CreateOrUpdate.
type FrontDoorsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FrontDoorsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FrontDoorsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FrontDoorsClientCreateOrUpdateResponse, error) {
	respType := FrontDoorsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.FrontDoor)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FrontDoorsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *FrontDoorsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *FrontDoorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FrontDoorsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FrontDoorsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FrontDoorsClientCreateOrUpdateResponse contains the response from method FrontDoorsClient.CreateOrUpdate.
type FrontDoorsClientCreateOrUpdateResponse struct {
	FrontDoor
}

// FrontDoorsClientDeletePollerResponse contains the response from method FrontDoorsClient.Delete.
type FrontDoorsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FrontDoorsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FrontDoorsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FrontDoorsClientDeleteResponse, error) {
	respType := FrontDoorsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FrontDoorsClientDeletePollerResponse from the provided client and resume token.
func (l *FrontDoorsClientDeletePollerResponse) Resume(ctx context.Context, client *FrontDoorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FrontDoorsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FrontDoorsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FrontDoorsClientDeleteResponse contains the response from method FrontDoorsClient.Delete.
type FrontDoorsClientDeleteResponse struct {
	// placeholder for future response values
}

// FrontDoorsClientGetResponse contains the response from method FrontDoorsClient.Get.
type FrontDoorsClientGetResponse struct {
	FrontDoor
}

// FrontDoorsClientListByResourceGroupResponse contains the response from method FrontDoorsClient.ListByResourceGroup.
type FrontDoorsClientListByResourceGroupResponse struct {
	ListResult
}

// FrontDoorsClientListResponse contains the response from method FrontDoorsClient.List.
type FrontDoorsClientListResponse struct {
	ListResult
}

// FrontDoorsClientValidateCustomDomainResponse contains the response from method FrontDoorsClient.ValidateCustomDomain.
type FrontDoorsClientValidateCustomDomainResponse struct {
	ValidateCustomDomainOutput
}

// FrontendEndpointsClientDisableHTTPSPollerResponse contains the response from method FrontendEndpointsClient.DisableHTTPS.
type FrontendEndpointsClientDisableHTTPSPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FrontendEndpointsClientDisableHTTPSPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FrontendEndpointsClientDisableHTTPSPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FrontendEndpointsClientDisableHTTPSResponse, error) {
	respType := FrontendEndpointsClientDisableHTTPSResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FrontendEndpointsClientDisableHTTPSPollerResponse from the provided client and resume token.
func (l *FrontendEndpointsClientDisableHTTPSPollerResponse) Resume(ctx context.Context, client *FrontendEndpointsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FrontendEndpointsClient.DisableHTTPS", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FrontendEndpointsClientDisableHTTPSPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FrontendEndpointsClientDisableHTTPSResponse contains the response from method FrontendEndpointsClient.DisableHTTPS.
type FrontendEndpointsClientDisableHTTPSResponse struct {
	// placeholder for future response values
}

// FrontendEndpointsClientEnableHTTPSPollerResponse contains the response from method FrontendEndpointsClient.EnableHTTPS.
type FrontendEndpointsClientEnableHTTPSPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FrontendEndpointsClientEnableHTTPSPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FrontendEndpointsClientEnableHTTPSPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FrontendEndpointsClientEnableHTTPSResponse, error) {
	respType := FrontendEndpointsClientEnableHTTPSResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a FrontendEndpointsClientEnableHTTPSPollerResponse from the provided client and resume token.
func (l *FrontendEndpointsClientEnableHTTPSPollerResponse) Resume(ctx context.Context, client *FrontendEndpointsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FrontendEndpointsClient.EnableHTTPS", token, client.pl)
	if err != nil {
		return err
	}
	poller := &FrontendEndpointsClientEnableHTTPSPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// FrontendEndpointsClientEnableHTTPSResponse contains the response from method FrontendEndpointsClient.EnableHTTPS.
type FrontendEndpointsClientEnableHTTPSResponse struct {
	// placeholder for future response values
}

// FrontendEndpointsClientGetResponse contains the response from method FrontendEndpointsClient.Get.
type FrontendEndpointsClientGetResponse struct {
	FrontendEndpoint
}

// FrontendEndpointsClientListByFrontDoorResponse contains the response from method FrontendEndpointsClient.ListByFrontDoor.
type FrontendEndpointsClientListByFrontDoorResponse struct {
	FrontendEndpointsListResult
}

// ManagedRuleSetsClientListResponse contains the response from method ManagedRuleSetsClient.List.
type ManagedRuleSetsClientListResponse struct {
	ManagedRuleSetDefinitionList
}

// NameAvailabilityClientCheckResponse contains the response from method NameAvailabilityClient.Check.
type NameAvailabilityClientCheckResponse struct {
	CheckNameAvailabilityOutput
}

// NameAvailabilityWithSubscriptionClientCheckResponse contains the response from method NameAvailabilityWithSubscriptionClient.Check.
type NameAvailabilityWithSubscriptionClientCheckResponse struct {
	CheckNameAvailabilityOutput
}

// NetworkExperimentProfilesClientCreateOrUpdatePollerResponse contains the response from method NetworkExperimentProfilesClient.CreateOrUpdate.
type NetworkExperimentProfilesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NetworkExperimentProfilesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l NetworkExperimentProfilesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NetworkExperimentProfilesClientCreateOrUpdateResponse, error) {
	respType := NetworkExperimentProfilesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Profile)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a NetworkExperimentProfilesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *NetworkExperimentProfilesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *NetworkExperimentProfilesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NetworkExperimentProfilesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &NetworkExperimentProfilesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// NetworkExperimentProfilesClientCreateOrUpdateResponse contains the response from method NetworkExperimentProfilesClient.CreateOrUpdate.
type NetworkExperimentProfilesClientCreateOrUpdateResponse struct {
	Profile
}

// NetworkExperimentProfilesClientDeletePollerResponse contains the response from method NetworkExperimentProfilesClient.Delete.
type NetworkExperimentProfilesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NetworkExperimentProfilesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l NetworkExperimentProfilesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NetworkExperimentProfilesClientDeleteResponse, error) {
	respType := NetworkExperimentProfilesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a NetworkExperimentProfilesClientDeletePollerResponse from the provided client and resume token.
func (l *NetworkExperimentProfilesClientDeletePollerResponse) Resume(ctx context.Context, client *NetworkExperimentProfilesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NetworkExperimentProfilesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &NetworkExperimentProfilesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// NetworkExperimentProfilesClientDeleteResponse contains the response from method NetworkExperimentProfilesClient.Delete.
type NetworkExperimentProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkExperimentProfilesClientGetResponse contains the response from method NetworkExperimentProfilesClient.Get.
type NetworkExperimentProfilesClientGetResponse struct {
	Profile
}

// NetworkExperimentProfilesClientListByResourceGroupResponse contains the response from method NetworkExperimentProfilesClient.ListByResourceGroup.
type NetworkExperimentProfilesClientListByResourceGroupResponse struct {
	ProfileList
}

// NetworkExperimentProfilesClientListResponse contains the response from method NetworkExperimentProfilesClient.List.
type NetworkExperimentProfilesClientListResponse struct {
	ProfileList
}

// NetworkExperimentProfilesClientUpdatePollerResponse contains the response from method NetworkExperimentProfilesClient.Update.
type NetworkExperimentProfilesClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NetworkExperimentProfilesClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l NetworkExperimentProfilesClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NetworkExperimentProfilesClientUpdateResponse, error) {
	respType := NetworkExperimentProfilesClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Profile)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a NetworkExperimentProfilesClientUpdatePollerResponse from the provided client and resume token.
func (l *NetworkExperimentProfilesClientUpdatePollerResponse) Resume(ctx context.Context, client *NetworkExperimentProfilesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NetworkExperimentProfilesClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &NetworkExperimentProfilesClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// NetworkExperimentProfilesClientUpdateResponse contains the response from method NetworkExperimentProfilesClient.Update.
type NetworkExperimentProfilesClientUpdateResponse struct {
	Profile
}

// PoliciesClientCreateOrUpdatePollerResponse contains the response from method PoliciesClient.CreateOrUpdate.
type PoliciesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PoliciesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PoliciesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PoliciesClientCreateOrUpdateResponse, error) {
	respType := PoliciesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.WebApplicationFirewallPolicy)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PoliciesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PoliciesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PoliciesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PoliciesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PoliciesClientCreateOrUpdateResponse contains the response from method PoliciesClient.CreateOrUpdate.
type PoliciesClientCreateOrUpdateResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientDeletePollerResponse contains the response from method PoliciesClient.Delete.
type PoliciesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PoliciesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PoliciesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PoliciesClientDeleteResponse, error) {
	respType := PoliciesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PoliciesClientDeletePollerResponse from the provided client and resume token.
func (l *PoliciesClientDeletePollerResponse) Resume(ctx context.Context, client *PoliciesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PoliciesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PoliciesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PoliciesClientDeleteResponse contains the response from method PoliciesClient.Delete.
type PoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// PoliciesClientGetResponse contains the response from method PoliciesClient.Get.
type PoliciesClientGetResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientListResponse contains the response from method PoliciesClient.List.
type PoliciesClientListResponse struct {
	WebApplicationFirewallPolicyList
}

// PreconfiguredEndpointsClientListResponse contains the response from method PreconfiguredEndpointsClient.List.
type PreconfiguredEndpointsClientListResponse struct {
	PreconfiguredEndpointList
}

// ReportsClientGetLatencyScorecardsResponse contains the response from method ReportsClient.GetLatencyScorecards.
type ReportsClientGetLatencyScorecardsResponse struct {
	LatencyScorecard
}

// ReportsClientGetTimeseriesResponse contains the response from method ReportsClient.GetTimeseries.
type ReportsClientGetTimeseriesResponse struct {
	Timeseries
}

// RulesEnginesClientCreateOrUpdatePollerResponse contains the response from method RulesEnginesClient.CreateOrUpdate.
type RulesEnginesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RulesEnginesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RulesEnginesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RulesEnginesClientCreateOrUpdateResponse, error) {
	respType := RulesEnginesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RulesEngine)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RulesEnginesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RulesEnginesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RulesEnginesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RulesEnginesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RulesEnginesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RulesEnginesClientCreateOrUpdateResponse contains the response from method RulesEnginesClient.CreateOrUpdate.
type RulesEnginesClientCreateOrUpdateResponse struct {
	RulesEngine
}

// RulesEnginesClientDeletePollerResponse contains the response from method RulesEnginesClient.Delete.
type RulesEnginesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RulesEnginesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RulesEnginesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RulesEnginesClientDeleteResponse, error) {
	respType := RulesEnginesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RulesEnginesClientDeletePollerResponse from the provided client and resume token.
func (l *RulesEnginesClientDeletePollerResponse) Resume(ctx context.Context, client *RulesEnginesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RulesEnginesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RulesEnginesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RulesEnginesClientDeleteResponse contains the response from method RulesEnginesClient.Delete.
type RulesEnginesClientDeleteResponse struct {
	// placeholder for future response values
}

// RulesEnginesClientGetResponse contains the response from method RulesEnginesClient.Get.
type RulesEnginesClientGetResponse struct {
	RulesEngine
}

// RulesEnginesClientListByFrontDoorResponse contains the response from method RulesEnginesClient.ListByFrontDoor.
type RulesEnginesClientListByFrontDoorResponse struct {
	RulesEngineListResult
}
