//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/listLabPlans.json
func ExampleLabPlansClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armlabservices.LabPlansClientListBySubscriptionOptions{Filter: nil})
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

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/listResourceGroupLabPlans.json
func ExampleLabPlansClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/getLabPlan.json
func ExampleLabPlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LabPlansClientGetResult)
}

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/putLabPlan.json
func ExampleLabPlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		armlabservices.LabPlan{
			Location: to.StringPtr("<location>"),
			Properties: &armlabservices.LabPlanProperties{
				DefaultAutoShutdownProfile: &armlabservices.AutoShutdownProfile{
					DisconnectDelay:          to.StringPtr("<disconnect-delay>"),
					IdleDelay:                to.StringPtr("<idle-delay>"),
					NoConnectDelay:           to.StringPtr("<no-connect-delay>"),
					ShutdownOnDisconnect:     armlabservices.EnableStateEnabled.ToPtr(),
					ShutdownOnIdle:           armlabservices.ShutdownOnIdleModeUserAbsence.ToPtr(),
					ShutdownWhenNotConnected: armlabservices.EnableStateEnabled.ToPtr(),
				},
				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					ClientSSHAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					WebRdpAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
					WebSSHAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
				},
				DefaultNetworkProfile: &armlabservices.LabPlanNetworkProfile{
					SubnetID: to.StringPtr("<subnet-id>"),
				},
				SharedGalleryID: to.StringPtr("<shared-gallery-id>"),
				SupportInfo: &armlabservices.SupportInfo{
					Email:        to.StringPtr("<email>"),
					Instructions: to.StringPtr("<instructions>"),
					Phone:        to.StringPtr("<phone>"),
					URL:          to.StringPtr("<url>"),
				},
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
	log.Printf("Response result: %#v\n", res.LabPlansClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/patchLabPlan.json
func ExampleLabPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		armlabservices.LabPlanUpdate{
			Properties: &armlabservices.LabPlanUpdateProperties{
				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					ClientSSHAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					WebRdpAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
					WebSSHAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
				},
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
	log.Printf("Response result: %#v\n", res.LabPlansClientUpdateResult)
}

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/deleteLabPlan.json
func ExampleLabPlansClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/saveImageVirtualMachine.json
func ExampleLabPlansClient_BeginSaveImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginSaveImage(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		armlabservices.SaveImageBody{
			Name:                to.StringPtr("<name>"),
			LabVirtualMachineID: to.StringPtr("<lab-virtual-machine-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}