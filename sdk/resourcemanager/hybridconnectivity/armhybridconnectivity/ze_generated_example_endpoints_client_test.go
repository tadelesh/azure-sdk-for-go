//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsList.json
func ExampleEndpointsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	pager := client.List("<resource-uri>",
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

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsGetCustom.json
func ExampleEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	res, err := client.Get(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EndpointsClientGetResult)
}

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsPutCustom.json
func ExampleEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		armhybridconnectivity.EndpointResource{
			Properties: &armhybridconnectivity.EndpointProperties{
				Type:       armhybridconnectivity.Type("custom").ToPtr(),
				ResourceID: to.StringPtr("<resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EndpointsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsPatchDefault.json
func ExampleEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	res, err := client.Update(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		armhybridconnectivity.EndpointResource{
			Properties: &armhybridconnectivity.EndpointProperties{
				Type: armhybridconnectivity.Type("default").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EndpointsClientUpdateResult)
}

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsDeleteDefault.json
func ExampleEndpointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	_, err = client.Delete(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsPostListCredentials.json
func ExampleEndpointsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	res, err := client.ListCredentials(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		&armhybridconnectivity.EndpointsClientListCredentialsOptions{Expiresin: to.Int64Ptr(10800)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EndpointsClientListCredentialsResult)
}