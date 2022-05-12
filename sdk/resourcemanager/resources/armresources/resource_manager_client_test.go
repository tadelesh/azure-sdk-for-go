//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armresources_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"os"
	"testing"
)

func TestNewManagementClient(t *testing.T) {
	subscriptionID := os.Getenv("AZURE_TENANT_ID")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Errorf("failed to obtain a credential: %v", err)
	}

	mgmtClient, err := armresources.NewResourceManagerClient(subscriptionID, cred, nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}

	ctx := context.Background()
	mgmtClient.ResourceGroups().Get(ctx, "test", nil)
}
