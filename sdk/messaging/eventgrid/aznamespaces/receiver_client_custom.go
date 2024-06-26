// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/eventgrid/aznamespaces/internal"
)

type receiverData struct {
	topic        string
	subscription string
}

// NewReceiverClient creates a [ReceiverClient] which uses an azcore.TokenCredential for authentication.
//   - topicName - Topic Name.
//   - subscriptionName - Event Subscription Name.
func NewReceiverClient(endpoint string, topic string, subscription string, cred azcore.TokenCredential, options *ClientOptions) (*ReceiverClient, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	azc, err := azcore.NewClient(internal.ModuleName+".Client", internal.ModuleVersion, runtime.PipelineOptions{
		PerRetry: []policy.Policy{
			runtime.NewBearerTokenPolicy(cred, []string{authScope}, nil),
		},
	}, &options.ClientOptions)

	if err != nil {
		return nil, err
	}

	return &ReceiverClient{
		internal: azc,
		endpoint: endpoint,
		data: receiverData{
			topic:        topic,
			subscription: subscription,
		},
	}, nil
}

// NewReceiverClientWithSharedKeyCredential creates a [ReceiverClient] using a shared key.
//   - topicName - Topic Name.
//   - subscriptionName - Event Subscription Name.
func NewReceiverClientWithSharedKeyCredential(endpoint string, topic string, subscription string, keyCred *azcore.KeyCredential, options *ClientOptions) (*ReceiverClient, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	azc, err := azcore.NewClient(internal.ModuleName+".Client", internal.ModuleVersion, runtime.PipelineOptions{
		PerRetry: []policy.Policy{
			runtime.NewKeyCredentialPolicy(keyCred, "Authorization", &runtime.KeyCredentialPolicyOptions{
				Prefix: "SharedAccessKey ",
			}),
		},
	}, &options.ClientOptions)

	if err != nil {
		return nil, err
	}

	return &ReceiverClient{
		internal: azc,
		endpoint: endpoint,
		data: receiverData{
			topic:        topic,
			subscription: subscription,
		},
	}, nil
}

// Reject rejects a batch of Cloud Events. The server responds with an HTTP 200 status code if the request is successfully
// accepted. The response body will include the set of successfully rejected lockTokens,
// along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - lockTokens - slice of lock tokens.
//   - options - RejectOptions contains the optional parameters for the ReceiverClient.Reject method.
func (client *ReceiverClient) Reject(ctx context.Context, lockTokens []string, options *RejectOptions) (RejectResponse, error) {
	return client.reject(ctx, client.data.topic, client.data.subscription, rejectOptions{LockTokens: lockTokens}, options)
}

// Acknowledge acknowledges a batch of Cloud Events. The server responds with an HTTP 200 status code if the request
// is successfully accepted. The response body will include the set of successfully acknowledged
// lockTokens, along with other failed lockTokens with their corresponding error information. Successfully acknowledged events
// will no longer be available to any consumer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - lockTokens - slice of lock tokens.
//   - options - AcknowledgeOptions contains the optional parameters for the ReceiverClient.Acknowledge method.
func (client *ReceiverClient) Acknowledge(ctx context.Context, lockTokens []string, options *AcknowledgeOptions) (AcknowledgeResponse, error) {
	return client.acknowledge(ctx, client.data.topic, client.data.subscription, acknowledgeOptions{LockTokens: lockTokens}, options)
}

// Release releases a batch of Cloud Events. The server responds with an HTTP 200 status code if the request is
// successfully accepted. The response body will include the set of successfully released lockTokens,
// along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - lockTokens - slice of lock tokens.
//   - options - ReleaseOptions contains the optional parameters for the ReceiverClient.Release method.
func (client *ReceiverClient) Release(ctx context.Context, lockTokens []string, options *ReleaseOptions) (ReleaseResponse, error) {
	return client.release(ctx, client.data.topic, client.data.subscription, releaseOptions{LockTokens: lockTokens}, options)
}

// RenewLocks renews locks for batch of Cloud Events. The server responds with an HTTP 200 status code if the request
// is successfully accepted. The response body will include the set of successfully renewed
// lockTokens, along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - lockTokens - slice of lock tokens.
//   - options - RenewLocksOptions contains the optional parameters for the ReceiverClient.RenewLocks method.
func (client *ReceiverClient) RenewLocks(ctx context.Context, lockTokens []string, options *RenewLocksOptions) (RenewLocksResponse, error) {
	return client.renewLock(ctx, client.data.topic, client.data.subscription, renewLockOptions{LockTokens: lockTokens}, options)
}

// Receive a batch of Cloud Events from a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - options - ReceiveOptions contains the optional parameters for the ReceiverClient.Receive method.
func (client *ReceiverClient) Receive(ctx context.Context, options *ReceiveOptions) (ReceiveResponse, error) {
	return client.receive(ctx, client.data.topic, client.data.subscription, options)
}
