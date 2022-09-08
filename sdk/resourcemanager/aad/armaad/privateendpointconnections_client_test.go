//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armaad_test

import (
	"context"
	"log"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

func TestPrivateEndpointConnectionsClientDo(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	request, err := runtime.NewRequest(ctx, "GET", runtime.JoinPaths("host", "url"))
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}
	request.Raw().Header["Accept"] = []string{"application/json"} // do anything with request
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	_ = response // handle response
}
