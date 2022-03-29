//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/createOrUpdateGuestConfigurationAssignment.json
func ExampleAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<guest-configuration-assignment-name>",
		"<resource-group-name>",
		"<vm-name>",
		armguestconfiguration.Assignment{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armguestconfiguration.AssignmentProperties{
				Context: to.StringPtr("<context>"),
				GuestConfiguration: &armguestconfiguration.Navigation{
					Name:           to.StringPtr("<name>"),
					AssignmentType: armguestconfiguration.AssignmentType("ApplyAndAutoCorrect").ToPtr(),
					ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						}},
					ContentHash: to.StringPtr("<content-hash>"),
					ContentURI:  to.StringPtr("<content-uri>"),
					Version:     to.StringPtr("<version>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssignmentsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/getGuestConfigurationAssignment.json
func ExampleAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<guest-configuration-assignment-name>",
		"<vm-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssignmentsClientGetResult)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/deleteGuestConfigurationAssignment.json
func ExampleAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<guest-configuration-assignment-name>",
		"<vm-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/listSubGuestConfigurationAssignments.json
func ExampleAssignmentsClient_SubscriptionList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.SubscriptionList(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssignmentsClientSubscriptionListResult)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/listRGGuestConfigurationAssignments.json
func ExampleAssignmentsClient_RGList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.RGList(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssignmentsClientRGListResult)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/listGuestConfigurationAssignments.json
func ExampleAssignmentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<vm-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssignmentsClientListResult)
}