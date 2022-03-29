//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_ListByJobDefinition-GET-example-91.json
func ExampleJobsClient_ListByJobDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.ListByJobDefinition("<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobsClientListByJobDefinitionOptions{Filter: nil})
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

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Get-GET-example-101.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<job-id>",
		"<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientGetResult)
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Cancel-POST-example-111.json
func ExampleJobsClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCancel(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<job-id>",
		"<resource-group-name>",
		"<data-manager-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Resume-POST-example-121.json
func ExampleJobsClient_BeginResume() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginResume(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<job-id>",
		"<resource-group-name>",
		"<data-manager-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_ListByDataService-GET-example-141.json
func ExampleJobsClient_ListByDataService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.ListByDataService("<data-service-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobsClientListByDataServiceOptions{Filter: nil})
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

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_ListByDataManager-GET-example-201.json
func ExampleJobsClient_ListByDataManager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.ListByDataManager("<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobsClientListByDataManagerOptions{Filter: nil})
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