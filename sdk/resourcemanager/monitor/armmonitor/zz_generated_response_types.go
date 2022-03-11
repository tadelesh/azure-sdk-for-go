//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// ActionGroupsClientCreateOrUpdateResponse contains the response from method ActionGroupsClient.CreateOrUpdate.
type ActionGroupsClientCreateOrUpdateResponse struct {
	ActionGroupResource
}

// ActionGroupsClientDeleteResponse contains the response from method ActionGroupsClient.Delete.
type ActionGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientEnableReceiverResponse contains the response from method ActionGroupsClient.EnableReceiver.
type ActionGroupsClientEnableReceiverResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientGetResponse contains the response from method ActionGroupsClient.Get.
type ActionGroupsClientGetResponse struct {
	ActionGroupResource
}

// ActionGroupsClientGetTestNotificationsResponse contains the response from method ActionGroupsClient.GetTestNotifications.
type ActionGroupsClientGetTestNotificationsResponse struct {
	TestNotificationDetailsResponse
}

// ActionGroupsClientListByResourceGroupResponse contains the response from method ActionGroupsClient.ListByResourceGroup.
type ActionGroupsClientListByResourceGroupResponse struct {
	ActionGroupList
}

// ActionGroupsClientListBySubscriptionIDResponse contains the response from method ActionGroupsClient.ListBySubscriptionID.
type ActionGroupsClientListBySubscriptionIDResponse struct {
	ActionGroupList
}

// ActionGroupsClientPostTestNotificationsPollerResponse contains the response from method ActionGroupsClient.PostTestNotifications.
type ActionGroupsClientPostTestNotificationsPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ActionGroupsClientPostTestNotificationsPoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ActionGroupsClientPostTestNotificationsPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ActionGroupsClientPostTestNotificationsResponse, error) {
	respType := ActionGroupsClientPostTestNotificationsResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.TestNotificationResponse)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ActionGroupsClientPostTestNotificationsPollerResponse from the provided client and resume token.
func (l *ActionGroupsClientPostTestNotificationsPollerResponse) Resume(ctx context.Context, client *ActionGroupsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ActionGroupsClient.PostTestNotifications", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ActionGroupsClientPostTestNotificationsPoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ActionGroupsClientPostTestNotificationsResponse contains the response from method ActionGroupsClient.PostTestNotifications.
type ActionGroupsClientPostTestNotificationsResponse struct {
	TestNotificationResponse
}

// ActionGroupsClientUpdateResponse contains the response from method ActionGroupsClient.Update.
type ActionGroupsClientUpdateResponse struct {
	ActionGroupResource
}

// ActivityLogAlertsClientCreateOrUpdateResponse contains the response from method ActivityLogAlertsClient.CreateOrUpdate.
type ActivityLogAlertsClientCreateOrUpdateResponse struct {
	ActivityLogAlertResource
}

// ActivityLogAlertsClientDeleteResponse contains the response from method ActivityLogAlertsClient.Delete.
type ActivityLogAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActivityLogAlertsClientGetResponse contains the response from method ActivityLogAlertsClient.Get.
type ActivityLogAlertsClientGetResponse struct {
	ActivityLogAlertResource
}

// ActivityLogAlertsClientListByResourceGroupResponse contains the response from method ActivityLogAlertsClient.ListByResourceGroup.
type ActivityLogAlertsClientListByResourceGroupResponse struct {
	AlertRuleList
}

// ActivityLogAlertsClientListBySubscriptionIDResponse contains the response from method ActivityLogAlertsClient.ListBySubscriptionID.
type ActivityLogAlertsClientListBySubscriptionIDResponse struct {
	AlertRuleList
}

// ActivityLogAlertsClientUpdateResponse contains the response from method ActivityLogAlertsClient.Update.
type ActivityLogAlertsClientUpdateResponse struct {
	ActivityLogAlertResource
}

// ActivityLogsClientListResponse contains the response from method ActivityLogsClient.List.
type ActivityLogsClientListResponse struct {
	EventDataCollection
}

// AlertRuleIncidentsClientGetResponse contains the response from method AlertRuleIncidentsClient.Get.
type AlertRuleIncidentsClientGetResponse struct {
	Incident
}

// AlertRuleIncidentsClientListByAlertRuleResponse contains the response from method AlertRuleIncidentsClient.ListByAlertRule.
type AlertRuleIncidentsClientListByAlertRuleResponse struct {
	IncidentListResult
}

// AlertRulesClientCreateOrUpdateResponse contains the response from method AlertRulesClient.CreateOrUpdate.
type AlertRulesClientCreateOrUpdateResponse struct {
	AlertRuleResource
}

// AlertRulesClientDeleteResponse contains the response from method AlertRulesClient.Delete.
type AlertRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// AlertRulesClientGetResponse contains the response from method AlertRulesClient.Get.
type AlertRulesClientGetResponse struct {
	AlertRuleResource
}

// AlertRulesClientListByResourceGroupResponse contains the response from method AlertRulesClient.ListByResourceGroup.
type AlertRulesClientListByResourceGroupResponse struct {
	AlertRuleResourceCollection
}

// AlertRulesClientListBySubscriptionResponse contains the response from method AlertRulesClient.ListBySubscription.
type AlertRulesClientListBySubscriptionResponse struct {
	AlertRuleResourceCollection
}

// AlertRulesClientUpdateResponse contains the response from method AlertRulesClient.Update.
type AlertRulesClientUpdateResponse struct {
	AlertRuleResource
}

// AutoscaleSettingsClientCreateOrUpdateResponse contains the response from method AutoscaleSettingsClient.CreateOrUpdate.
type AutoscaleSettingsClientCreateOrUpdateResponse struct {
	AutoscaleSettingResource
}

// AutoscaleSettingsClientDeleteResponse contains the response from method AutoscaleSettingsClient.Delete.
type AutoscaleSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// AutoscaleSettingsClientGetResponse contains the response from method AutoscaleSettingsClient.Get.
type AutoscaleSettingsClientGetResponse struct {
	AutoscaleSettingResource
}

// AutoscaleSettingsClientListByResourceGroupResponse contains the response from method AutoscaleSettingsClient.ListByResourceGroup.
type AutoscaleSettingsClientListByResourceGroupResponse struct {
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientListBySubscriptionResponse contains the response from method AutoscaleSettingsClient.ListBySubscription.
type AutoscaleSettingsClientListBySubscriptionResponse struct {
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientUpdateResponse contains the response from method AutoscaleSettingsClient.Update.
type AutoscaleSettingsClientUpdateResponse struct {
	AutoscaleSettingResource
}

// BaselinesClientListResponse contains the response from method BaselinesClient.List.
type BaselinesClientListResponse struct {
	MetricBaselinesResponse
}

// DataCollectionEndpointsClientCreateResponse contains the response from method DataCollectionEndpointsClient.Create.
type DataCollectionEndpointsClientCreateResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientDeleteResponse contains the response from method DataCollectionEndpointsClient.Delete.
type DataCollectionEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionEndpointsClientGetResponse contains the response from method DataCollectionEndpointsClient.Get.
type DataCollectionEndpointsClientGetResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientListByResourceGroupResponse contains the response from method DataCollectionEndpointsClient.ListByResourceGroup.
type DataCollectionEndpointsClientListByResourceGroupResponse struct {
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientListBySubscriptionResponse contains the response from method DataCollectionEndpointsClient.ListBySubscription.
type DataCollectionEndpointsClientListBySubscriptionResponse struct {
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientUpdateResponse contains the response from method DataCollectionEndpointsClient.Update.
type DataCollectionEndpointsClientUpdateResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionRuleAssociationsClientCreateResponse contains the response from method DataCollectionRuleAssociationsClient.Create.
type DataCollectionRuleAssociationsClientCreateResponse struct {
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientDeleteResponse contains the response from method DataCollectionRuleAssociationsClient.Delete.
type DataCollectionRuleAssociationsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRuleAssociationsClientGetResponse contains the response from method DataCollectionRuleAssociationsClient.Get.
type DataCollectionRuleAssociationsClientGetResponse struct {
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientListByResourceResponse contains the response from method DataCollectionRuleAssociationsClient.ListByResource.
type DataCollectionRuleAssociationsClientListByResourceResponse struct {
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRuleAssociationsClientListByRuleResponse contains the response from method DataCollectionRuleAssociationsClient.ListByRule.
type DataCollectionRuleAssociationsClientListByRuleResponse struct {
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRulesClientCreateResponse contains the response from method DataCollectionRulesClient.Create.
type DataCollectionRulesClientCreateResponse struct {
	DataCollectionRuleResource
}

// DataCollectionRulesClientDeleteResponse contains the response from method DataCollectionRulesClient.Delete.
type DataCollectionRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRulesClientGetResponse contains the response from method DataCollectionRulesClient.Get.
type DataCollectionRulesClientGetResponse struct {
	DataCollectionRuleResource
}

// DataCollectionRulesClientListByResourceGroupResponse contains the response from method DataCollectionRulesClient.ListByResourceGroup.
type DataCollectionRulesClientListByResourceGroupResponse struct {
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientListBySubscriptionResponse contains the response from method DataCollectionRulesClient.ListBySubscription.
type DataCollectionRulesClientListBySubscriptionResponse struct {
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientUpdateResponse contains the response from method DataCollectionRulesClient.Update.
type DataCollectionRulesClientUpdateResponse struct {
	DataCollectionRuleResource
}

// DiagnosticSettingsCategoryClientGetResponse contains the response from method DiagnosticSettingsCategoryClient.Get.
type DiagnosticSettingsCategoryClientGetResponse struct {
	DiagnosticSettingsCategoryResource
}

// DiagnosticSettingsCategoryClientListResponse contains the response from method DiagnosticSettingsCategoryClient.List.
type DiagnosticSettingsCategoryClientListResponse struct {
	DiagnosticSettingsCategoryResourceCollection
}

// DiagnosticSettingsClientCreateOrUpdateResponse contains the response from method DiagnosticSettingsClient.CreateOrUpdate.
type DiagnosticSettingsClientCreateOrUpdateResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientDeleteResponse contains the response from method DiagnosticSettingsClient.Delete.
type DiagnosticSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiagnosticSettingsClientGetResponse contains the response from method DiagnosticSettingsClient.Get.
type DiagnosticSettingsClientGetResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientListResponse contains the response from method DiagnosticSettingsClient.List.
type DiagnosticSettingsClientListResponse struct {
	DiagnosticSettingsResourceCollection
}

// EventCategoriesClientListResponse contains the response from method EventCategoriesClient.List.
type EventCategoriesClientListResponse struct {
	EventCategoryCollection
}

// LogProfilesClientCreateOrUpdateResponse contains the response from method LogProfilesClient.CreateOrUpdate.
type LogProfilesClientCreateOrUpdateResponse struct {
	LogProfileResource
}

// LogProfilesClientDeleteResponse contains the response from method LogProfilesClient.Delete.
type LogProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// LogProfilesClientGetResponse contains the response from method LogProfilesClient.Get.
type LogProfilesClientGetResponse struct {
	LogProfileResource
}

// LogProfilesClientListResponse contains the response from method LogProfilesClient.List.
type LogProfilesClientListResponse struct {
	LogProfileCollection
}

// LogProfilesClientUpdateResponse contains the response from method LogProfilesClient.Update.
type LogProfilesClientUpdateResponse struct {
	LogProfileResource
}

// MetricAlertsClientCreateOrUpdateResponse contains the response from method MetricAlertsClient.CreateOrUpdate.
type MetricAlertsClientCreateOrUpdateResponse struct {
	MetricAlertResource
}

// MetricAlertsClientDeleteResponse contains the response from method MetricAlertsClient.Delete.
type MetricAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// MetricAlertsClientGetResponse contains the response from method MetricAlertsClient.Get.
type MetricAlertsClientGetResponse struct {
	MetricAlertResource
}

// MetricAlertsClientListByResourceGroupResponse contains the response from method MetricAlertsClient.ListByResourceGroup.
type MetricAlertsClientListByResourceGroupResponse struct {
	MetricAlertResourceCollection
}

// MetricAlertsClientListBySubscriptionResponse contains the response from method MetricAlertsClient.ListBySubscription.
type MetricAlertsClientListBySubscriptionResponse struct {
	MetricAlertResourceCollection
}

// MetricAlertsClientUpdateResponse contains the response from method MetricAlertsClient.Update.
type MetricAlertsClientUpdateResponse struct {
	MetricAlertResource
}

// MetricAlertsStatusClientListByNameResponse contains the response from method MetricAlertsStatusClient.ListByName.
type MetricAlertsStatusClientListByNameResponse struct {
	MetricAlertStatusCollection
}

// MetricAlertsStatusClientListResponse contains the response from method MetricAlertsStatusClient.List.
type MetricAlertsStatusClientListResponse struct {
	MetricAlertStatusCollection
}

// MetricDefinitionsClientListResponse contains the response from method MetricDefinitionsClient.List.
type MetricDefinitionsClientListResponse struct {
	MetricDefinitionCollection
}

// MetricNamespacesClientListResponse contains the response from method MetricNamespacesClient.List.
type MetricNamespacesClientListResponse struct {
	MetricNamespaceCollection
}

// MetricsClientListResponse contains the response from method MetricsClient.List.
type MetricsClientListResponse struct {
	Response
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeletePollerResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.ListByPrivateLinkScope.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.ListByPrivateLinkScope.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	PrivateLinkResourceListResult
}

// PrivateLinkScopeOperationStatusClientGetResponse contains the response from method PrivateLinkScopeOperationStatusClient.Get.
type PrivateLinkScopeOperationStatusClientGetResponse struct {
	OperationStatus
}

// PrivateLinkScopedResourcesClientCreateOrUpdatePollerResponse contains the response from method PrivateLinkScopedResourcesClient.CreateOrUpdate.
type PrivateLinkScopedResourcesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateLinkScopedResourcesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateLinkScopedResourcesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateLinkScopedResourcesClientCreateOrUpdateResponse, error) {
	respType := PrivateLinkScopedResourcesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ScopedResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PrivateLinkScopedResourcesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateLinkScopedResourcesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateLinkScopedResourcesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateLinkScopedResourcesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateLinkScopedResourcesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PrivateLinkScopedResourcesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopedResourcesClient.CreateOrUpdate.
type PrivateLinkScopedResourcesClientCreateOrUpdateResponse struct {
	ScopedResource
}

// PrivateLinkScopedResourcesClientDeletePollerResponse contains the response from method PrivateLinkScopedResourcesClient.Delete.
type PrivateLinkScopedResourcesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateLinkScopedResourcesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateLinkScopedResourcesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateLinkScopedResourcesClientDeleteResponse, error) {
	respType := PrivateLinkScopedResourcesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PrivateLinkScopedResourcesClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateLinkScopedResourcesClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateLinkScopedResourcesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateLinkScopedResourcesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateLinkScopedResourcesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PrivateLinkScopedResourcesClientDeleteResponse contains the response from method PrivateLinkScopedResourcesClient.Delete.
type PrivateLinkScopedResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopedResourcesClientGetResponse contains the response from method PrivateLinkScopedResourcesClient.Get.
type PrivateLinkScopedResourcesClientGetResponse struct {
	ScopedResource
}

// PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkScopedResourcesClient.ListByPrivateLinkScope.
type PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse struct {
	ScopedResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientDeletePollerResponse contains the response from method PrivateLinkScopesClient.Delete.
type PrivateLinkScopesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateLinkScopesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateLinkScopesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateLinkScopesClientDeleteResponse, error) {
	respType := PrivateLinkScopesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a PrivateLinkScopesClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateLinkScopesClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateLinkScopesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateLinkScopesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateLinkScopesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.Delete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientListByResourceGroupResponse contains the response from method PrivateLinkScopesClient.ListByResourceGroup.
type PrivateLinkScopesClientListByResourceGroupResponse struct {
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.List.
type PrivateLinkScopesClientListResponse struct {
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	AzureMonitorPrivateLinkScope
}

// ScheduledQueryRulesClientCreateOrUpdateResponse contains the response from method ScheduledQueryRulesClient.CreateOrUpdate.
type ScheduledQueryRulesClientCreateOrUpdateResponse struct {
	LogSearchRuleResource
}

// ScheduledQueryRulesClientDeleteResponse contains the response from method ScheduledQueryRulesClient.Delete.
type ScheduledQueryRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// ScheduledQueryRulesClientGetResponse contains the response from method ScheduledQueryRulesClient.Get.
type ScheduledQueryRulesClientGetResponse struct {
	LogSearchRuleResource
}

// ScheduledQueryRulesClientListByResourceGroupResponse contains the response from method ScheduledQueryRulesClient.ListByResourceGroup.
type ScheduledQueryRulesClientListByResourceGroupResponse struct {
	LogSearchRuleResourceCollection
}

// ScheduledQueryRulesClientListBySubscriptionResponse contains the response from method ScheduledQueryRulesClient.ListBySubscription.
type ScheduledQueryRulesClientListBySubscriptionResponse struct {
	LogSearchRuleResourceCollection
}

// ScheduledQueryRulesClientUpdateResponse contains the response from method ScheduledQueryRulesClient.Update.
type ScheduledQueryRulesClientUpdateResponse struct {
	LogSearchRuleResource
}

// TenantActivityLogsClientListResponse contains the response from method TenantActivityLogsClient.List.
type TenantActivityLogsClientListResponse struct {
	EventDataCollection
}

// VMInsightsClientGetOnboardingStatusResponse contains the response from method VMInsightsClient.GetOnboardingStatus.
type VMInsightsClientGetOnboardingStatusResponse struct {
	VMInsightsOnboardingStatus
}
