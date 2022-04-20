//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateOrUpdateARestorePointCollectionForCrossRegionCopy.json
func ExampleRestorePointCollectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientCreateOrUpdateResponse, err := restorePointCollectionsClient.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		armcompute.RestorePointCollection{
			Location: to.Ptr("<resource-location>"),
			Tags: map[string]*string{
				"myTag1": to.Ptr("tagValue1"),
			},
			Properties: &armcompute.RestorePointCollectionProperties{
				Source: &armcompute.RestorePointCollectionSourceProperties{
					ID: to.Ptr("<resource-id-of-the-source-resource-used-to-create-this-restore-point-collection>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = restorePointCollectionsClientCreateOrUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RestorePointCollections_Update_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientUpdateResponse, err := restorePointCollectionsClient.Update(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		armcompute.RestorePointCollectionUpdate{
			Tags: map[string]*string{
				"key8536": to.Ptr("aaaaaaaaaaaaaaaaaaa"),
			},
			Properties: &armcompute.RestorePointCollectionProperties{
				Source: &armcompute.RestorePointCollectionSourceProperties{
					ID: to.Ptr("<resource-id-of-the-source-resource-used-to-create-this-restore-point-collection>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = restorePointCollectionsClientUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RestorePointCollections_Delete_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientDeleteResponsePoller, err := restorePointCollectionsClient.BeginDelete(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		&armcompute.RestorePointCollectionsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = restorePointCollectionsClientDeleteResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointCollection.json
func ExampleRestorePointCollectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientGetResponse, err := restorePointCollectionsClient.Get(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		&armcompute.RestorePointCollectionsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = restorePointCollectionsClientGetResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointCollectionsInAResourceGroup.json
func ExampleRestorePointCollectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientNewListPager := restorePointCollectionsClient.NewListPager("<resource-group-name>",
		nil)
	for restorePointCollectionsClientNewListPager.More() {
		nextResult, err := restorePointCollectionsClientNewListPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointCollectionsInASubscription.json
func ExampleRestorePointCollectionsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	restorePointCollectionsClient, err := armcompute.NewRestorePointCollectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	restorePointCollectionsClientNewListAllPager := restorePointCollectionsClient.NewListAllPager(nil)
	for restorePointCollectionsClientNewListAllPager.More() {
		nextResult, err := restorePointCollectionsClientNewListAllPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
