//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListSubscriptionSqlServerInstance.json
func ExampleSQLServerInstancesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListByResourceGroupSqlServerInstance.json
func ExampleSQLServerInstancesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/GetSqlServerInstance.json
func ExampleSQLServerInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLServerInstancesClientGetResult)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateSqlServerInstance.json
func ExampleSQLServerInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		armazurearcdata.SQLServerInstance{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
			Properties: &armazurearcdata.SQLServerInstanceProperties{
				AzureDefenderStatus:            armazurearcdata.DefenderStatus("Protected").ToPtr(),
				AzureDefenderStatusLastUpdated: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t }()),
				Collation:                      to.StringPtr("<collation>"),
				ContainerResourceID:            to.StringPtr("<container-resource-id>"),
				CurrentVersion:                 to.StringPtr("<current-version>"),
				Edition:                        armazurearcdata.EditionType("Developer").ToPtr(),
				InstanceName:                   to.StringPtr("<instance-name>"),
				LicenseType:                    armazurearcdata.ArcSQLServerLicenseType("Free").ToPtr(),
				PatchLevel:                     to.StringPtr("<patch-level>"),
				ProductID:                      to.StringPtr("<product-id>"),
				Status:                         armazurearcdata.ConnectionStatus("Connected").ToPtr(),
				TCPDynamicPorts:                to.StringPtr("<tcpdynamic-ports>"),
				TCPStaticPorts:                 to.StringPtr("<tcpstatic-ports>"),
				VCore:                          to.StringPtr("<vcore>"),
				Version:                        armazurearcdata.SQLVersion("SQL Server 2017").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.SQLServerInstancesClientCreateResult)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/DeleteSqlServerInstance.json
func ExampleSQLServerInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/UpdateSqlServerInstance.json
func ExampleSQLServerInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLServerInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<sql-server-instance-name>",
		armazurearcdata.SQLServerInstanceUpdate{
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLServerInstancesClientUpdateResult)
}