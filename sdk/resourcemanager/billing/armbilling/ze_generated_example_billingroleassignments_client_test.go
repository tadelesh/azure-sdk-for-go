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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignment.json
func ExampleRoleAssignmentsClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.GetByBillingAccount(ctx,
		"<billing-account-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientGetByBillingAccountResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignmentDelete.json
func ExampleRoleAssignmentsClient_DeleteByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.DeleteByBillingAccount(ctx,
		"<billing-account-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientDeleteByBillingAccountResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignment.json
func ExampleRoleAssignmentsClient_GetByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.GetByInvoiceSection(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<invoice-section-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientGetByInvoiceSectionResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentDelete.json
func ExampleRoleAssignmentsClient_DeleteByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.DeleteByInvoiceSection(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<invoice-section-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientDeleteByInvoiceSectionResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignment.json
func ExampleRoleAssignmentsClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.GetByBillingProfile(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientGetByBillingProfileResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignmentDelete.json
func ExampleRoleAssignmentsClient_DeleteByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	res, err := client.DeleteByBillingProfile(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RoleAssignmentsClientDeleteByBillingProfileResult)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignmentList.json
func ExampleRoleAssignmentsClient_ListByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	pager := client.ListByBillingAccount("<billing-account-name>",
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

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentList.json
func ExampleRoleAssignmentsClient_ListByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	pager := client.ListByInvoiceSection("<billing-account-name>",
		"<billing-profile-name>",
		"<invoice-section-name>",
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

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignmentList.json
func ExampleRoleAssignmentsClient_ListByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewRoleAssignmentsClient(cred, nil)
	pager := client.ListByBillingProfile("<billing-account-name>",
		"<billing-profile-name>",
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