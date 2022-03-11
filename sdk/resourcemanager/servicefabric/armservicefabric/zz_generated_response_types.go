//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// ApplicationTypeVersionsClientCreateOrUpdatePollerResponse contains the response from method ApplicationTypeVersionsClient.CreateOrUpdate.
type ApplicationTypeVersionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationTypeVersionsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationTypeVersionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationTypeVersionsClientCreateOrUpdateResponse, error) {
	respType := ApplicationTypeVersionsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ApplicationTypeVersionResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationTypeVersionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ApplicationTypeVersionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ApplicationTypeVersionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationTypeVersionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationTypeVersionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationTypeVersionsClientCreateOrUpdateResponse contains the response from method ApplicationTypeVersionsClient.CreateOrUpdate.
type ApplicationTypeVersionsClientCreateOrUpdateResponse struct {
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientDeletePollerResponse contains the response from method ApplicationTypeVersionsClient.Delete.
type ApplicationTypeVersionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationTypeVersionsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationTypeVersionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationTypeVersionsClientDeleteResponse, error) {
	respType := ApplicationTypeVersionsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationTypeVersionsClientDeletePollerResponse from the provided client and resume token.
func (l *ApplicationTypeVersionsClientDeletePollerResponse) Resume(ctx context.Context, client *ApplicationTypeVersionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationTypeVersionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationTypeVersionsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationTypeVersionsClientDeleteResponse contains the response from method ApplicationTypeVersionsClient.Delete.
type ApplicationTypeVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypeVersionsClientGetResponse contains the response from method ApplicationTypeVersionsClient.Get.
type ApplicationTypeVersionsClientGetResponse struct {
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientListResponse contains the response from method ApplicationTypeVersionsClient.List.
type ApplicationTypeVersionsClientListResponse struct {
	ApplicationTypeVersionResourceList
}

// ApplicationTypesClientCreateOrUpdateResponse contains the response from method ApplicationTypesClient.CreateOrUpdate.
type ApplicationTypesClientCreateOrUpdateResponse struct {
	ApplicationTypeResource
}

// ApplicationTypesClientDeletePollerResponse contains the response from method ApplicationTypesClient.Delete.
type ApplicationTypesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationTypesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationTypesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationTypesClientDeleteResponse, error) {
	respType := ApplicationTypesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationTypesClientDeletePollerResponse from the provided client and resume token.
func (l *ApplicationTypesClientDeletePollerResponse) Resume(ctx context.Context, client *ApplicationTypesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationTypesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationTypesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationTypesClientDeleteResponse contains the response from method ApplicationTypesClient.Delete.
type ApplicationTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypesClientGetResponse contains the response from method ApplicationTypesClient.Get.
type ApplicationTypesClientGetResponse struct {
	ApplicationTypeResource
}

// ApplicationTypesClientListResponse contains the response from method ApplicationTypesClient.List.
type ApplicationTypesClientListResponse struct {
	ApplicationTypeResourceList
}

// ApplicationsClientCreateOrUpdatePollerResponse contains the response from method ApplicationsClient.CreateOrUpdate.
type ApplicationsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationsClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationsClientCreateOrUpdateResponse, error) {
	respType := ApplicationsClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ApplicationResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ApplicationsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ApplicationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationsClientCreateOrUpdateResponse contains the response from method ApplicationsClient.CreateOrUpdate.
type ApplicationsClientCreateOrUpdateResponse struct {
	ApplicationResource
}

// ApplicationsClientDeletePollerResponse contains the response from method ApplicationsClient.Delete.
type ApplicationsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationsClientDeleteResponse, error) {
	respType := ApplicationsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationsClientDeletePollerResponse from the provided client and resume token.
func (l *ApplicationsClientDeletePollerResponse) Resume(ctx context.Context, client *ApplicationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationsClientDeleteResponse contains the response from method ApplicationsClient.Delete.
type ApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationsClientGetResponse contains the response from method ApplicationsClient.Get.
type ApplicationsClientGetResponse struct {
	ApplicationResource
}

// ApplicationsClientListResponse contains the response from method ApplicationsClient.List.
type ApplicationsClientListResponse struct {
	ApplicationResourceList
}

// ApplicationsClientUpdatePollerResponse contains the response from method ApplicationsClient.Update.
type ApplicationsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ApplicationsClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ApplicationsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ApplicationsClientUpdateResponse, error) {
	respType := ApplicationsClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ApplicationResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ApplicationsClientUpdatePollerResponse from the provided client and resume token.
func (l *ApplicationsClientUpdatePollerResponse) Resume(ctx context.Context, client *ApplicationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ApplicationsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ApplicationsClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ApplicationsClientUpdateResponse contains the response from method ApplicationsClient.Update.
type ApplicationsClientUpdateResponse struct {
	ApplicationResource
}

// ClusterVersionsClientGetByEnvironmentResponse contains the response from method ClusterVersionsClient.GetByEnvironment.
type ClusterVersionsClientGetByEnvironmentResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientGetResponse contains the response from method ClusterVersionsClient.Get.
type ClusterVersionsClientGetResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListByEnvironmentResponse contains the response from method ClusterVersionsClient.ListByEnvironment.
type ClusterVersionsClientListByEnvironmentResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListResponse contains the response from method ClusterVersionsClient.List.
type ClusterVersionsClientListResponse struct {
	ClusterCodeVersionsListResult
}

// ClustersClientCreateOrUpdatePollerResponse contains the response from method ClustersClient.CreateOrUpdate.
type ClustersClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClustersClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersClientCreateOrUpdateResponse, error) {
	respType := ClustersClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ClustersClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ClustersClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClustersClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.CreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.Delete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.ListByResourceGroup.
type ClustersClientListByResourceGroupResponse struct {
	ClusterListResult
}

// ClustersClientListResponse contains the response from method ClustersClient.List.
type ClustersClientListResponse struct {
	ClusterListResult
}

// ClustersClientListUpgradableVersionsResponse contains the response from method ClustersClient.ListUpgradableVersions.
type ClustersClientListUpgradableVersionsResponse struct {
	UpgradableVersionPathResult
}

// ClustersClientUpdatePollerResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClustersClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersClientUpdateResponse, error) {
	respType := ClustersClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ClustersClientUpdatePollerResponse from the provided client and resume token.
func (l *ClustersClientUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClustersClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	Cluster
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ServicesClientCreateOrUpdatePollerResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServicesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServicesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServicesClientCreateOrUpdateResponse, error) {
	respType := ServicesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ServiceResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ServicesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ServicesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServicesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServicesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServiceResource
}

// ServicesClientDeletePollerResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServicesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServicesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServicesClientDeleteResponse, error) {
	respType := ServicesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ServicesClientDeletePollerResponse from the provided client and resume token.
func (l *ServicesClientDeletePollerResponse) Resume(ctx context.Context, client *ServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServicesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServicesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServiceResource
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServiceResourceList
}

// ServicesClientUpdatePollerResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServicesClientUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServicesClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServicesClientUpdateResponse, error) {
	respType := ServicesClientUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ServiceResource)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a ServicesClientUpdatePollerResponse from the provided client and resume token.
func (l *ServicesClientUpdatePollerResponse) Resume(ctx context.Context, client *ServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServicesClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServicesClientUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServiceResource
}
