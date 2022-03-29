//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/RecoveryPointsRecommendedForMove_List.json
func ExampleRecoveryPointsRecommendedForMoveClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewRecoveryPointsRecommendedForMoveClient("<subscription-id>", cred, nil)
	pager := client.List("<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		"<protected-item-name>",
		armrecoveryservicesbackup.ListRecoveryPointsRecommendedForMoveRequest{
			ExcludedRPList: []*string{
				to.StringPtr("348916168024334"),
				to.StringPtr("348916168024335")},
			ObjectType: to.StringPtr("<object-type>"),
		},
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