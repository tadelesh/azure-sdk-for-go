//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_Get.json
func ExampleAppsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		&armappplatform.AppsClientGetOptions{SyncStatus: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res.AppsClientGetResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		armappplatform.AppResource{
			Location: to.StringPtr("<location>"),
			Properties: &armappplatform.AppResourceProperties{
				ActiveDeploymentName: to.StringPtr("<active-deployment-name>"),
				EnableEndToEndTLS:    to.BoolPtr(false),
				Fqdn:                 to.StringPtr("<fqdn>"),
				HTTPSOnly:            to.BoolPtr(false),
				PersistentDisk: &armappplatform.PersistentDisk{
					MountPath: to.StringPtr("<mount-path>"),
					SizeInGB:  to.Int32Ptr(2),
				},
				Public: to.BoolPtr(true),
				TemporaryDisk: &armappplatform.TemporaryDisk{
					MountPath: to.StringPtr("<mount-path>"),
					SizeInGB:  to.Int32Ptr(2),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res.AppsClientCreateOrUpdateResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_Delete.json
func ExampleAppsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		armappplatform.AppResource{
			Identity: &armappplatform.ManagedIdentityProperties{
				Type: armappplatform.ManagedIdentityTypeSystemAssigned.ToPtr(),
			},
			Location: to.StringPtr("<location>"),
			Properties: &armappplatform.AppResourceProperties{
				ActiveDeploymentName: to.StringPtr("<active-deployment-name>"),
				EnableEndToEndTLS:    to.BoolPtr(false),
				Fqdn:                 to.StringPtr("<fqdn>"),
				HTTPSOnly:            to.BoolPtr(false),
				PersistentDisk: &armappplatform.PersistentDisk{
					MountPath: to.StringPtr("<mount-path>"),
					SizeInGB:  to.Int32Ptr(2),
				},
				Public: to.BoolPtr(true),
				TemporaryDisk: &armappplatform.TemporaryDisk{
					MountPath: to.StringPtr("<mount-path>"),
					SizeInGB:  to.Int32Ptr(2),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res.AppsClientUpdateResult
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_List.json
func ExampleAppsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<service-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Apps_ValidateDomain.json
func ExampleAppsClient_ValidateDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}

	ctx := context.Background()
	client := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	res, err := client.ValidateDomain(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		armappplatform.CustomDomainValidatePayload{
			Name: to.StringPtr("<name>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res.AppsClientValidateDomainResult
}
