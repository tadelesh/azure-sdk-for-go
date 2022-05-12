//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armresources_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func TestDelete(t *testing.T) {
	subscriptionID := os.Getenv("AZURE_TENANT_ID")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Errorf("failed to obtain a credential: %v", err)
	}
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	ctx := context.Background()
	resp, err := rgClient.Delete(ctx, "test123", nil)
	if err != nil {
		t.Errorf("failed to delete resource group: %v", err)
	}
	t.Logf("%v", resp)
}

func TestBeginDelete(t *testing.T) {
	subscriptionID := os.Getenv("AZURE_TENANT_ID")
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Errorf("failed to obtain a credential: %v", err)
	}
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	ctx := context.Background()
	resp, err := armresources.NewOperationWaiter(rgClient.BeginDelete(ctx, "test123", nil)).Wait(ctx, time.Second)
	if err != nil {
		t.Errorf("failed to delete resource group: %v", err)
	}
	t.Logf("%v", resp)
}

func TestList(t *testing.T) {
	subscriptionID := os.Getenv("AZURE_TENANT_ID")
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Errorf("failed to obtain a credential: %v", err)
	}
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	ctx := context.Background()
	iter := armresources.NewIterator[armresources.ResourceGroup](rgClient.NewListPager(nil))
	for iter.More() {
		rg, err := iter.NextItem(ctx)
		if err != nil {
			t.Errorf("failed to fatch item: %v", err)
		}
		t.Logf("%v", rg)
	}
}
