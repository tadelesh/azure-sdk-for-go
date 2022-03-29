//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListRevisions.json
func ExampleContainerAppsRevisionsClient_ListRevisions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsRevisionsClient("<subscription-id>", cred, nil)
	pager := client.ListRevisions("<resource-group-name>",
		"<container-app-name>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetRevision.json
func ExampleContainerAppsRevisionsClient_GetRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsRevisionsClient("<subscription-id>", cred, nil)
	res, err := client.GetRevision(ctx,
		"<resource-group-name>",
		"<container-app-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContainerAppsRevisionsClientGetRevisionResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ActivateRevision.json
func ExampleContainerAppsRevisionsClient_ActivateRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsRevisionsClient("<subscription-id>", cred, nil)
	_, err = client.ActivateRevision(ctx,
		"<resource-group-name>",
		"<container-app-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeactivateRevision.json
func ExampleContainerAppsRevisionsClient_DeactivateRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsRevisionsClient("<subscription-id>", cred, nil)
	_, err = client.DeactivateRevision(ctx,
		"<resource-group-name>",
		"<container-app-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/RestartRevision.json
func ExampleContainerAppsRevisionsClient_RestartRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsRevisionsClient("<subscription-id>", cred, nil)
	_, err = client.RestartRevision(ctx,
		"<resource-group-name>",
		"<container-app-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}