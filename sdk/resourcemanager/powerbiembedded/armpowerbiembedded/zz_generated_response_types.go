//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiembedded

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// ManagementClientGetAvailableOperationsResponse contains the response from method ManagementClient.GetAvailableOperations.
type ManagementClientGetAvailableOperationsResponse struct {
	OperationList
}

// WorkspaceCollectionsClientCheckNameAvailabilityResponse contains the response from method WorkspaceCollectionsClient.CheckNameAvailability.
type WorkspaceCollectionsClientCheckNameAvailabilityResponse struct {
	CheckNameResponse
}

// WorkspaceCollectionsClientCreateResponse contains the response from method WorkspaceCollectionsClient.Create.
type WorkspaceCollectionsClientCreateResponse struct {
	WorkspaceCollection
}

// WorkspaceCollectionsClientDeletePollerResponse contains the response from method WorkspaceCollectionsClient.Delete.
type WorkspaceCollectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspaceCollectionsClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspaceCollectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspaceCollectionsClientDeleteResponse, error) {
	respType := WorkspaceCollectionsClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a WorkspaceCollectionsClientDeletePollerResponse from the provided client and resume token.
func (l *WorkspaceCollectionsClientDeletePollerResponse) Resume(ctx context.Context, client *WorkspaceCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspaceCollectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &WorkspaceCollectionsClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// WorkspaceCollectionsClientDeleteResponse contains the response from method WorkspaceCollectionsClient.Delete.
type WorkspaceCollectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspaceCollectionsClientGetAccessKeysResponse contains the response from method WorkspaceCollectionsClient.GetAccessKeys.
type WorkspaceCollectionsClientGetAccessKeysResponse struct {
	WorkspaceCollectionAccessKeys
}

// WorkspaceCollectionsClientGetByNameResponse contains the response from method WorkspaceCollectionsClient.GetByName.
type WorkspaceCollectionsClientGetByNameResponse struct {
	WorkspaceCollection
}

// WorkspaceCollectionsClientListByResourceGroupResponse contains the response from method WorkspaceCollectionsClient.ListByResourceGroup.
type WorkspaceCollectionsClientListByResourceGroupResponse struct {
	WorkspaceCollectionList
}

// WorkspaceCollectionsClientListBySubscriptionResponse contains the response from method WorkspaceCollectionsClient.ListBySubscription.
type WorkspaceCollectionsClientListBySubscriptionResponse struct {
	WorkspaceCollectionList
}

// WorkspaceCollectionsClientMigrateResponse contains the response from method WorkspaceCollectionsClient.Migrate.
type WorkspaceCollectionsClientMigrateResponse struct {
	// placeholder for future response values
}

// WorkspaceCollectionsClientRegenerateKeyResponse contains the response from method WorkspaceCollectionsClient.RegenerateKey.
type WorkspaceCollectionsClientRegenerateKeyResponse struct {
	WorkspaceCollectionAccessKeys
}

// WorkspaceCollectionsClientUpdateResponse contains the response from method WorkspaceCollectionsClient.Update.
type WorkspaceCollectionsClientUpdateResponse struct {
	WorkspaceCollection
}

// WorkspacesClientListResponse contains the response from method WorkspacesClient.List.
type WorkspacesClientListResponse struct {
	WorkspaceList
}
