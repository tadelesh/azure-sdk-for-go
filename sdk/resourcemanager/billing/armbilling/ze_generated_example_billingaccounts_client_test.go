//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsList.json
func ExampleAccountsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewAccountsClient(cred, nil)
	pager := client.List(&armbilling.AccountsClientListOptions{Expand: nil})
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

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountWithExpand.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewAccountsClient(cred, nil)
	res, err := client.Get(ctx,
		"<billing-account-name>",
		&armbilling.AccountsClientGetOptions{Expand: to.StringPtr("<expand>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientGetResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingAccount.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewAccountsClient(cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<billing-account-name>",
		armbilling.AccountUpdateRequest{
			Properties: &armbilling.AccountProperties{
				DisplayName: to.StringPtr("<display-name>"),
				SoldTo: &armbilling.AddressDetails{
					AddressLine1: to.StringPtr("<address-line1>"),
					City:         to.StringPtr("<city>"),
					CompanyName:  to.StringPtr("<company-name>"),
					Country:      to.StringPtr("<country>"),
					FirstName:    to.StringPtr("<first-name>"),
					LastName:     to.StringPtr("<last-name>"),
					PostalCode:   to.StringPtr("<postal-code>"),
					Region:       to.StringPtr("<region>"),
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
	log.Printf("Response result: %#v\n", res.AccountsClientUpdateResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionsListWithCreateSubPermission.json
func ExampleAccountsClient_ListInvoiceSectionsByCreateSubscriptionPermission() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewAccountsClient(cred, nil)
	pager := client.ListInvoiceSectionsByCreateSubscriptionPermission("<billing-account-name>",
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