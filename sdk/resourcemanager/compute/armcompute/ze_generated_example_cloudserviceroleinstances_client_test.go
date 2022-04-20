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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/DeleteCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientDeleteResponsePoller, err := cloudServiceRoleInstancesClient.BeginDelete(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = cloudServiceRoleInstancesClientDeleteResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientGetResponse, err := cloudServiceRoleInstancesClient.Get(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = cloudServiceRoleInstancesClientGetResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetInstanceViewOfCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientGetInstanceViewResponse, err := cloudServiceRoleInstancesClient.GetInstanceView(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = cloudServiceRoleInstancesClientGetInstanceViewResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceRolesInstances.json
func ExampleCloudServiceRoleInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientNewListPager := cloudServiceRoleInstancesClient.NewListPager("<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientListOptions{Expand: nil})
	for cloudServiceRoleInstancesClientNewListPager.More() {
		nextResult, err := cloudServiceRoleInstancesClientNewListPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/RestartCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientRestartResponsePoller, err := cloudServiceRoleInstancesClient.BeginRestart(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientBeginRestartOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = cloudServiceRoleInstancesClientRestartResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ReimageCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_BeginReimage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientReimageResponsePoller, err := cloudServiceRoleInstancesClient.BeginReimage(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientBeginReimageOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = cloudServiceRoleInstancesClientReimageResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/RebuildCloudServiceRoleInstance.json
func ExampleCloudServiceRoleInstancesClient_BeginRebuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	cloudServiceRoleInstancesClient, err := armcompute.NewCloudServiceRoleInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	cloudServiceRoleInstancesClientRebuildResponsePoller, err := cloudServiceRoleInstancesClient.BeginRebuild(ctx,
		"<role-instance-name>",
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServiceRoleInstancesClientBeginRebuildOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = cloudServiceRoleInstancesClientRebuildResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
