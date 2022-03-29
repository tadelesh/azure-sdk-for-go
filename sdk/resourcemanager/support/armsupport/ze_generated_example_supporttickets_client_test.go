//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CheckNameAvailabilityWithSubscription.json
func ExampleTicketsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		armsupport.CheckNameAvailabilityInput{
			Name: to.StringPtr("<name>"),
			Type: armsupport.TypeMicrosoftSupportSupportTickets.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TicketsClientCheckNameAvailabilityResult)
}

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListSupportTicketsCreatedOnOrAfterAndInOpenStateBySubscription.json
func ExampleTicketsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	pager := client.List(&armsupport.TicketsClientListOptions{Top: nil,
		Filter: to.StringPtr("<filter>"),
	})
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

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetSubscriptionSupportTicketDetails.json
func ExampleTicketsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<support-ticket-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TicketsClientGetResult)
}

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/UpdateContactDetailsOfSupportTicketForSubscription.json
func ExampleTicketsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<support-ticket-name>",
		armsupport.UpdateSupportTicket{
			ContactDetails: &armsupport.UpdateContactProfile{
				AdditionalEmailAddresses: []*string{
					to.StringPtr("tname@contoso.com"),
					to.StringPtr("teamtest@contoso.com")},
				Country:                  to.StringPtr("<country>"),
				FirstName:                to.StringPtr("<first-name>"),
				LastName:                 to.StringPtr("<last-name>"),
				PhoneNumber:              to.StringPtr("<phone-number>"),
				PreferredContactMethod:   armsupport.PreferredContactMethod("email").ToPtr(),
				PreferredSupportLanguage: to.StringPtr("<preferred-support-language>"),
				PreferredTimeZone:        to.StringPtr("<preferred-time-zone>"),
				PrimaryEmailAddress:      to.StringPtr("<primary-email-address>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TicketsClientUpdateResult)
}

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateBillingSupportTicketForSubscription.json
func ExampleTicketsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<support-ticket-name>",
		armsupport.TicketDetails{
			Properties: &armsupport.TicketDetailsProperties{
				Description: to.StringPtr("<description>"),
				ContactDetails: &armsupport.ContactProfile{
					Country:                  to.StringPtr("<country>"),
					FirstName:                to.StringPtr("<first-name>"),
					LastName:                 to.StringPtr("<last-name>"),
					PreferredContactMethod:   armsupport.PreferredContactMethod("email").ToPtr(),
					PreferredSupportLanguage: to.StringPtr("<preferred-support-language>"),
					PreferredTimeZone:        to.StringPtr("<preferred-time-zone>"),
					PrimaryEmailAddress:      to.StringPtr("<primary-email-address>"),
				},
				ProblemClassificationID: to.StringPtr("<problem-classification-id>"),
				ServiceID:               to.StringPtr("<service-id>"),
				Severity:                armsupport.SeverityLevel("moderate").ToPtr(),
				Title:                   to.StringPtr("<title>"),
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
	log.Printf("Response result: %#v\n", res.TicketsClientCreateResult)
}