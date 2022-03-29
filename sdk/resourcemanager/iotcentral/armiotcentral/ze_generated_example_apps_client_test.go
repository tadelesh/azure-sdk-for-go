//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotcentral_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Get.json
func ExampleAppsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AppsClientGetResult)
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiotcentral.App{
			Location: to.StringPtr("<location>"),
			Identity: &armiotcentral.SystemAssignedServiceIdentity{
				Type: armiotcentral.SystemAssignedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armiotcentral.AppProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Subdomain:   to.StringPtr("<subdomain>"),
				Template:    to.StringPtr("<template>"),
			},
			SKU: &armiotcentral.AppSKUInfo{
				Name: armiotcentral.AppSKU("ST2").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AppsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiotcentral.AppPatch{
			Identity: &armiotcentral.SystemAssignedServiceIdentity{
				Type: armiotcentral.SystemAssignedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armiotcentral.AppProperties{
				DisplayName: to.StringPtr("<display-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AppsClientUpdateResult)
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Delete.json
func ExampleAppsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_ListBySubscription.json
func ExampleAppsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_ListByResourceGroup.json
func ExampleAppsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CheckNameAvailability.json
func ExampleAppsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		armiotcentral.OperationInputs{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AppsClientCheckNameAvailabilityResult)
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CheckSubdomainAvailability.json
func ExampleAppsClient_CheckSubdomainAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	res, err := client.CheckSubdomainAvailability(ctx,
		armiotcentral.OperationInputs{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AppsClientCheckSubdomainAvailabilityResult)
}

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Templates.json
func ExampleAppsClient_ListTemplates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	pager := client.ListTemplates(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}