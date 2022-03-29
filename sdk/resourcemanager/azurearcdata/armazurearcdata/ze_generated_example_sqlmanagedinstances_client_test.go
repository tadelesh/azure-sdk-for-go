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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListSubscriptionSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListByResourceGroupSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/GetSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLManagedInstancesClientGetResult)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		armazurearcdata.SQLManagedInstance{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.StringPtr("<name>"),
				Type: armazurearcdata.ExtendedLocationTypes("CustomLocation").ToPtr(),
			},
			Properties: &armazurearcdata.SQLManagedInstanceProperties{
				Admin: to.StringPtr("<admin>"),
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				ClusterID:   to.StringPtr("<cluster-id>"),
				EndTime:     to.StringPtr("<end-time>"),
				ExtensionID: to.StringPtr("<extension-id>"),
				K8SRaw: &armazurearcdata.SQLManagedInstanceK8SRaw{
					AdditionalProperties: map[string]map[string]interface{}{
						"additionalProperty": {},
					},
					Spec: &armazurearcdata.SQLManagedInstanceK8SSpec{
						Replicas: to.Int32Ptr(1),
						Scheduling: &armazurearcdata.K8SScheduling{
							Default: &armazurearcdata.K8SSchedulingOptions{
								Resources: &armazurearcdata.K8SResourceRequirements{
									Limits: map[string]*string{
										"additionalProperty": to.StringPtr("additionalValue"),
										"cpu":                to.StringPtr("1"),
										"memory":             to.StringPtr("8Gi"),
									},
									Requests: map[string]*string{
										"additionalProperty": to.StringPtr("additionalValue"),
										"cpu":                to.StringPtr("1"),
										"memory":             to.StringPtr("8Gi"),
									},
								},
							},
						},
					},
				},
				LicenseType: armazurearcdata.ArcSQLManagedInstanceLicenseType("LicenseIncluded").ToPtr(),
				StartTime:   to.StringPtr("<start-time>"),
			},
			SKU: &armazurearcdata.SQLManagedInstanceSKU{
				Name: armazurearcdata.SQLManagedInstanceSKUName("vCore").ToPtr(),
				Dev:  to.BoolPtr(true),
				Tier: armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.SQLManagedInstancesClientCreateResult)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/DeleteSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/UpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		armazurearcdata.SQLManagedInstanceUpdate{
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLManagedInstancesClientUpdateResult)
}