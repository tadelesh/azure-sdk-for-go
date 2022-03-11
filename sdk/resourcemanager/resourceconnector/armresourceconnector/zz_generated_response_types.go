//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armresourceconnector

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"time"
)

// AppliancesClientCreateOrUpdatePollerResponse contains the response from method AppliancesClient.CreateOrUpdate.
type AppliancesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AppliancesClientCreateOrUpdatePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AppliancesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AppliancesClientCreateOrUpdateResponse, error) {
	respType := AppliancesClientCreateOrUpdateResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Appliance)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a AppliancesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *AppliancesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *AppliancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AppliancesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AppliancesClientCreateOrUpdatePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// AppliancesClientCreateOrUpdateResponse contains the response from method AppliancesClient.CreateOrUpdate.
type AppliancesClientCreateOrUpdateResponse struct {
	Appliance
}

// AppliancesClientDeletePollerResponse contains the response from method AppliancesClient.Delete.
type AppliancesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AppliancesClientDeletePoller
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AppliancesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AppliancesClientDeleteResponse, error) {
	respType := AppliancesClientDeleteResponse{}
	_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	return respType, nil
}

// Resume rehydrates a AppliancesClientDeletePollerResponse from the provided client and resume token.
func (l *AppliancesClientDeletePollerResponse) Resume(ctx context.Context, client *AppliancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AppliancesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AppliancesClientDeletePoller{
		pt: pt,
	}
	_, err = poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	return nil
}

// AppliancesClientDeleteResponse contains the response from method AppliancesClient.Delete.
type AppliancesClientDeleteResponse struct {
	// placeholder for future response values
}

// AppliancesClientGetResponse contains the response from method AppliancesClient.Get.
type AppliancesClientGetResponse struct {
	Appliance
}

// AppliancesClientListByResourceGroupResponse contains the response from method AppliancesClient.ListByResourceGroup.
type AppliancesClientListByResourceGroupResponse struct {
	ApplianceListResult
}

// AppliancesClientListBySubscriptionResponse contains the response from method AppliancesClient.ListBySubscription.
type AppliancesClientListBySubscriptionResponse struct {
	ApplianceListResult
}

// AppliancesClientListClusterUserCredentialResponse contains the response from method AppliancesClient.ListClusterUserCredential.
type AppliancesClientListClusterUserCredentialResponse struct {
	ApplianceListCredentialResults
}

// AppliancesClientListOperationsResponse contains the response from method AppliancesClient.ListOperations.
type AppliancesClientListOperationsResponse struct {
	ApplianceOperationsList
}

// AppliancesClientUpdateResponse contains the response from method AppliancesClient.Update.
type AppliancesClientUpdateResponse struct {
	Appliance
}
