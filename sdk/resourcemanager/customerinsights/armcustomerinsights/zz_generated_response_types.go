//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// AuthorizationPoliciesClientCreateOrUpdateResponse contains the response from method AuthorizationPoliciesClient.CreateOrUpdate.
type AuthorizationPoliciesClientCreateOrUpdateResponse struct {
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientGetResponse contains the response from method AuthorizationPoliciesClient.Get.
type AuthorizationPoliciesClientGetResponse struct {
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientListByHubResponse contains the response from method AuthorizationPoliciesClient.ListByHub.
type AuthorizationPoliciesClientListByHubResponse struct {
	AuthorizationPolicyListResult
}

// AuthorizationPoliciesClientRegeneratePrimaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegeneratePrimaryKey.
type AuthorizationPoliciesClientRegeneratePrimaryKeyResponse struct {
	AuthorizationPolicy
}

// AuthorizationPoliciesClientRegenerateSecondaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegenerateSecondaryKey.
type AuthorizationPoliciesClientRegenerateSecondaryKeyResponse struct {
	AuthorizationPolicy
}

// ConnectorMappingsClientCreateOrUpdateResponse contains the response from method ConnectorMappingsClient.CreateOrUpdate.
type ConnectorMappingsClientCreateOrUpdateResponse struct {
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientDeleteResponse contains the response from method ConnectorMappingsClient.Delete.
type ConnectorMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorMappingsClientGetResponse contains the response from method ConnectorMappingsClient.Get.
type ConnectorMappingsClientGetResponse struct {
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientListByConnectorResponse contains the response from method ConnectorMappingsClient.ListByConnector.
type ConnectorMappingsClientListByConnectorResponse struct {
	ConnectorMappingListResult
}

// ConnectorsClientCreateOrUpdatePollerResponse contains the response from method ConnectorsClient.CreateOrUpdate.
type ConnectorsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConnectorsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConnectorsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConnectorsClientCreateOrUpdateResponse, error) {
	respType := ConnectorsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ConnectorResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ConnectorsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ConnectorsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ConnectorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConnectorsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ConnectorsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ConnectorsClientCreateOrUpdateResponse contains the response from method ConnectorsClient.CreateOrUpdate.
type ConnectorsClientCreateOrUpdateResponse struct {
	ConnectorResourceFormat
}

// ConnectorsClientDeletePollerResponse contains the response from method ConnectorsClient.Delete.
type ConnectorsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConnectorsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConnectorsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConnectorsClientDeleteResponse, error) {
	respType := ConnectorsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ConnectorsClientDeletePollerResponse from the provided client and resume token.
func (l *ConnectorsClientDeletePollerResponse) Resume(ctx context.Context, client *ConnectorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConnectorsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ConnectorsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ConnectorsClientDeleteResponse contains the response from method ConnectorsClient.Delete.
type ConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorsClientGetResponse contains the response from method ConnectorsClient.Get.
type ConnectorsClientGetResponse struct {
	ConnectorResourceFormat
}

// ConnectorsClientListByHubResponse contains the response from method ConnectorsClient.ListByHub.
type ConnectorsClientListByHubResponse struct {
	ConnectorListResult
}

// HubsClientCreateOrUpdateResponse contains the response from method HubsClient.CreateOrUpdate.
type HubsClientCreateOrUpdateResponse struct {
	Hub
}

// HubsClientDeletePollerResponse contains the response from method HubsClient.Delete.
type HubsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *HubsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l HubsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (HubsClientDeleteResponse, error) {
	respType := HubsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a HubsClientDeletePollerResponse from the provided client and resume token.
func (l *HubsClientDeletePollerResponse) Resume(ctx context.Context, client *HubsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("HubsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &HubsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// HubsClientDeleteResponse contains the response from method HubsClient.Delete.
type HubsClientDeleteResponse struct {
	// placeholder for future response values
}

// HubsClientGetResponse contains the response from method HubsClient.Get.
type HubsClientGetResponse struct {
	Hub
}

// HubsClientListByResourceGroupResponse contains the response from method HubsClient.ListByResourceGroup.
type HubsClientListByResourceGroupResponse struct {
	HubListResult
}

// HubsClientListResponse contains the response from method HubsClient.List.
type HubsClientListResponse struct {
	HubListResult
}

// HubsClientUpdateResponse contains the response from method HubsClient.Update.
type HubsClientUpdateResponse struct {
	Hub
}

// ImagesClientGetUploadURLForDataResponse contains the response from method ImagesClient.GetUploadURLForData.
type ImagesClientGetUploadURLForDataResponse struct {
	ImageDefinition
}

// ImagesClientGetUploadURLForEntityTypeResponse contains the response from method ImagesClient.GetUploadURLForEntityType.
type ImagesClientGetUploadURLForEntityTypeResponse struct {
	ImageDefinition
}

// InteractionsClientCreateOrUpdatePollerResponse contains the response from method InteractionsClient.CreateOrUpdate.
type InteractionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *InteractionsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l InteractionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (InteractionsClientCreateOrUpdateResponse, error) {
	respType := InteractionsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.InteractionResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a InteractionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *InteractionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *InteractionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("InteractionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &InteractionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// InteractionsClientCreateOrUpdateResponse contains the response from method InteractionsClient.CreateOrUpdate.
type InteractionsClientCreateOrUpdateResponse struct {
	InteractionResourceFormat
}

// InteractionsClientGetResponse contains the response from method InteractionsClient.Get.
type InteractionsClientGetResponse struct {
	InteractionResourceFormat
}

// InteractionsClientListByHubResponse contains the response from method InteractionsClient.ListByHub.
type InteractionsClientListByHubResponse struct {
	InteractionListResult
}

// InteractionsClientSuggestRelationshipLinksResponse contains the response from method InteractionsClient.SuggestRelationshipLinks.
type InteractionsClientSuggestRelationshipLinksResponse struct {
	SuggestRelationshipLinksResponse
}

// KpiClientCreateOrUpdatePollerResponse contains the response from method KpiClient.CreateOrUpdate.
type KpiClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *KpiClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l KpiClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (KpiClientCreateOrUpdateResponse, error) {
	respType := KpiClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.KpiResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a KpiClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *KpiClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *KpiClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("KpiClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &KpiClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// KpiClientCreateOrUpdateResponse contains the response from method KpiClient.CreateOrUpdate.
type KpiClientCreateOrUpdateResponse struct {
	KpiResourceFormat
}

// KpiClientDeletePollerResponse contains the response from method KpiClient.Delete.
type KpiClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *KpiClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l KpiClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (KpiClientDeleteResponse, error) {
	respType := KpiClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a KpiClientDeletePollerResponse from the provided client and resume token.
func (l *KpiClientDeletePollerResponse) Resume(ctx context.Context, client *KpiClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("KpiClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &KpiClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// KpiClientDeleteResponse contains the response from method KpiClient.Delete.
type KpiClientDeleteResponse struct {
	// placeholder for future response values
}

// KpiClientGetResponse contains the response from method KpiClient.Get.
type KpiClientGetResponse struct {
	KpiResourceFormat
}

// KpiClientListByHubResponse contains the response from method KpiClient.ListByHub.
type KpiClientListByHubResponse struct {
	KpiListResult
}

// KpiClientReprocessResponse contains the response from method KpiClient.Reprocess.
type KpiClientReprocessResponse struct {
	// placeholder for future response values
}

// LinksClientCreateOrUpdatePollerResponse contains the response from method LinksClient.CreateOrUpdate.
type LinksClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *LinksClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l LinksClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (LinksClientCreateOrUpdateResponse, error) {
	respType := LinksClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.LinkResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a LinksClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *LinksClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *LinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("LinksClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &LinksClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// LinksClientCreateOrUpdateResponse contains the response from method LinksClient.CreateOrUpdate.
type LinksClientCreateOrUpdateResponse struct {
	LinkResourceFormat
}

// LinksClientDeleteResponse contains the response from method LinksClient.Delete.
type LinksClientDeleteResponse struct {
	// placeholder for future response values
}

// LinksClientGetResponse contains the response from method LinksClient.Get.
type LinksClientGetResponse struct {
	LinkResourceFormat
}

// LinksClientListByHubResponse contains the response from method LinksClient.ListByHub.
type LinksClientListByHubResponse struct {
	LinkListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PredictionsClientCreateOrUpdatePollerResponse contains the response from method PredictionsClient.CreateOrUpdate.
type PredictionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PredictionsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PredictionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PredictionsClientCreateOrUpdateResponse, error) {
	respType := PredictionsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PredictionResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PredictionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PredictionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PredictionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PredictionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PredictionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PredictionsClientCreateOrUpdateResponse contains the response from method PredictionsClient.CreateOrUpdate.
type PredictionsClientCreateOrUpdateResponse struct {
	PredictionResourceFormat
}

// PredictionsClientDeletePollerResponse contains the response from method PredictionsClient.Delete.
type PredictionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PredictionsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PredictionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PredictionsClientDeleteResponse, error) {
	respType := PredictionsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PredictionsClientDeletePollerResponse from the provided client and resume token.
func (l *PredictionsClientDeletePollerResponse) Resume(ctx context.Context, client *PredictionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PredictionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PredictionsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PredictionsClientDeleteResponse contains the response from method PredictionsClient.Delete.
type PredictionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PredictionsClientGetModelStatusResponse contains the response from method PredictionsClient.GetModelStatus.
type PredictionsClientGetModelStatusResponse struct {
	PredictionModelStatus
}

// PredictionsClientGetResponse contains the response from method PredictionsClient.Get.
type PredictionsClientGetResponse struct {
	PredictionResourceFormat
}

// PredictionsClientGetTrainingResultsResponse contains the response from method PredictionsClient.GetTrainingResults.
type PredictionsClientGetTrainingResultsResponse struct {
	PredictionTrainingResults
}

// PredictionsClientListByHubResponse contains the response from method PredictionsClient.ListByHub.
type PredictionsClientListByHubResponse struct {
	PredictionListResult
}

// PredictionsClientModelStatusResponse contains the response from method PredictionsClient.ModelStatus.
type PredictionsClientModelStatusResponse struct {
	// placeholder for future response values
}

// ProfilesClientCreateOrUpdatePollerResponse contains the response from method ProfilesClient.CreateOrUpdate.
type ProfilesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ProfilesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ProfilesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ProfilesClientCreateOrUpdateResponse, error) {
	respType := ProfilesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ProfileResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ProfilesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ProfilesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ProfilesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ProfilesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ProfilesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ProfilesClientCreateOrUpdateResponse contains the response from method ProfilesClient.CreateOrUpdate.
type ProfilesClientCreateOrUpdateResponse struct {
	ProfileResourceFormat
}

// ProfilesClientDeletePollerResponse contains the response from method ProfilesClient.Delete.
type ProfilesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ProfilesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ProfilesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ProfilesClientDeleteResponse, error) {
	respType := ProfilesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ProfilesClientDeletePollerResponse from the provided client and resume token.
func (l *ProfilesClientDeletePollerResponse) Resume(ctx context.Context, client *ProfilesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ProfilesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ProfilesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ProfilesClientDeleteResponse contains the response from method ProfilesClient.Delete.
type ProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProfilesClientGetEnrichingKpisResponse contains the response from method ProfilesClient.GetEnrichingKpis.
type ProfilesClientGetEnrichingKpisResponse struct {
	// Array of KpiDefinition
	KpiDefinitionArray []*KpiDefinition
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	ProfileResourceFormat
}

// ProfilesClientListByHubResponse contains the response from method ProfilesClient.ListByHub.
type ProfilesClientListByHubResponse struct {
	ProfileListResult
}

// RelationshipLinksClientCreateOrUpdatePollerResponse contains the response from method RelationshipLinksClient.CreateOrUpdate.
type RelationshipLinksClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RelationshipLinksClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RelationshipLinksClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RelationshipLinksClientCreateOrUpdateResponse, error) {
	respType := RelationshipLinksClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RelationshipLinkResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RelationshipLinksClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RelationshipLinksClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RelationshipLinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RelationshipLinksClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RelationshipLinksClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RelationshipLinksClientCreateOrUpdateResponse contains the response from method RelationshipLinksClient.CreateOrUpdate.
type RelationshipLinksClientCreateOrUpdateResponse struct {
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientDeletePollerResponse contains the response from method RelationshipLinksClient.Delete.
type RelationshipLinksClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RelationshipLinksClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RelationshipLinksClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RelationshipLinksClientDeleteResponse, error) {
	respType := RelationshipLinksClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RelationshipLinksClientDeletePollerResponse from the provided client and resume token.
func (l *RelationshipLinksClientDeletePollerResponse) Resume(ctx context.Context, client *RelationshipLinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RelationshipLinksClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RelationshipLinksClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RelationshipLinksClientDeleteResponse contains the response from method RelationshipLinksClient.Delete.
type RelationshipLinksClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipLinksClientGetResponse contains the response from method RelationshipLinksClient.Get.
type RelationshipLinksClientGetResponse struct {
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientListByHubResponse contains the response from method RelationshipLinksClient.ListByHub.
type RelationshipLinksClientListByHubResponse struct {
	RelationshipLinkListResult
}

// RelationshipsClientCreateOrUpdatePollerResponse contains the response from method RelationshipsClient.CreateOrUpdate.
type RelationshipsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RelationshipsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RelationshipsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RelationshipsClientCreateOrUpdateResponse, error) {
	respType := RelationshipsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RelationshipResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RelationshipsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RelationshipsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RelationshipsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RelationshipsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RelationshipsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RelationshipsClientCreateOrUpdateResponse contains the response from method RelationshipsClient.CreateOrUpdate.
type RelationshipsClientCreateOrUpdateResponse struct {
	RelationshipResourceFormat
}

// RelationshipsClientDeletePollerResponse contains the response from method RelationshipsClient.Delete.
type RelationshipsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RelationshipsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RelationshipsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RelationshipsClientDeleteResponse, error) {
	respType := RelationshipsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RelationshipsClientDeletePollerResponse from the provided client and resume token.
func (l *RelationshipsClientDeletePollerResponse) Resume(ctx context.Context, client *RelationshipsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RelationshipsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RelationshipsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RelationshipsClientDeleteResponse contains the response from method RelationshipsClient.Delete.
type RelationshipsClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipsClientGetResponse contains the response from method RelationshipsClient.Get.
type RelationshipsClientGetResponse struct {
	RelationshipResourceFormat
}

// RelationshipsClientListByHubResponse contains the response from method RelationshipsClient.ListByHub.
type RelationshipsClientListByHubResponse struct {
	RelationshipListResult
}

// RoleAssignmentsClientCreateOrUpdatePollerResponse contains the response from method RoleAssignmentsClient.CreateOrUpdate.
type RoleAssignmentsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RoleAssignmentsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RoleAssignmentsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RoleAssignmentsClientCreateOrUpdateResponse, error) {
	respType := RoleAssignmentsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RoleAssignmentResourceFormat)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a RoleAssignmentsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RoleAssignmentsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RoleAssignmentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RoleAssignmentsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RoleAssignmentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// RoleAssignmentsClientCreateOrUpdateResponse contains the response from method RoleAssignmentsClient.CreateOrUpdate.
type RoleAssignmentsClientCreateOrUpdateResponse struct {
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientDeleteResponse contains the response from method RoleAssignmentsClient.Delete.
type RoleAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// RoleAssignmentsClientGetResponse contains the response from method RoleAssignmentsClient.Get.
type RoleAssignmentsClientGetResponse struct {
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientListByHubResponse contains the response from method RoleAssignmentsClient.ListByHub.
type RoleAssignmentsClientListByHubResponse struct {
	RoleAssignmentListResult
}

// RolesClientListByHubResponse contains the response from method RolesClient.ListByHub.
type RolesClientListByHubResponse struct {
	RoleListResult
}

// ViewsClientCreateOrUpdateResponse contains the response from method ViewsClient.CreateOrUpdate.
type ViewsClientCreateOrUpdateResponse struct {
	ViewResourceFormat
}

// ViewsClientDeleteResponse contains the response from method ViewsClient.Delete.
type ViewsClientDeleteResponse struct {
	// placeholder for future response values
}

// ViewsClientGetResponse contains the response from method ViewsClient.Get.
type ViewsClientGetResponse struct {
	ViewResourceFormat
}

// ViewsClientListByHubResponse contains the response from method ViewsClient.ListByHub.
type ViewsClientListByHubResponse struct {
	ViewListResult
}

// WidgetTypesClientGetResponse contains the response from method WidgetTypesClient.Get.
type WidgetTypesClientGetResponse struct {
	WidgetTypeResourceFormat
}

// WidgetTypesClientListByHubResponse contains the response from method WidgetTypesClient.ListByHub.
type WidgetTypesClientListByHubResponse struct {
	WidgetTypeListResult
}
