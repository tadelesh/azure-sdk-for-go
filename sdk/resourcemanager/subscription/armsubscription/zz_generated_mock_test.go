//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"testing"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"golang.org/x/net/http2"
)

var (
	ctx            context.Context
	subscriptionId string
	cred           azcore.TokenCredential
	err            error
	con            *arm.Connection
	mockHost       string
)

func TestSubscriptions_ListLocations(t *testing.T) {
	// From example listLocations
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionsClient(con)
	_, err := client.ListLocations(ctx,
		"83aa47df-e3e9-49ff-877b-94304bf3d3ad",
		&SubscriptionsListLocationsOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscriptions_Get(t *testing.T) {
	// From example getSubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionsClient(con)
	res, err := client.Get(ctx,
		"83aa47df-e3e9-49ff-877b-94304bf3d3ad",
		&SubscriptionsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Subscription.ID == nil {
		t.Fatal("Subscription.ID should not be nil!")
	}
}

func TestSubscriptions_List(t *testing.T) {
	// From example listSubscriptions
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionsClient(con)
	pager := client.List(&SubscriptionsListOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("Subscription.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("Subscription.ID should not be nil!")
			}
		}
	}
}

func TestTenants_List(t *testing.T) {
	// From example listTenants
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewTenantsClient(con)
	pager := client.List(&TenantsListOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("TenantIDDescription.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("TenantIDDescription.ID should not be nil!")
			}
		}
	}
}

func TestSubscription_Cancel(t *testing.T) {
	// From example cancelSubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionClient(con)
	_, err := client.Cancel(ctx,
		"83aa47df-e3e9-49ff-877b-94304bf3d3ad",
		&SubscriptionCancelOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscription_Rename(t *testing.T) {
	// From example renameSubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionClient(con)
	_, err := client.Rename(ctx,
		"83aa47df-e3e9-49ff-877b-94304bf3d3ad",
		SubscriptionName{
			SubscriptionName: to.StringPtr("Test Sub"),
		},
		&SubscriptionRenameOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscription_Enable(t *testing.T) {
	// From example enableSubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSubscriptionClient(con)
	_, err := client.Enable(ctx,
		"7948bcee-488c-47ce-941c-38e20ede803d",
		&SubscriptionEnableOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestOperations_List(t *testing.T) {
	// From example getOperations
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewOperationsClient(con)
	_, err := client.List(ctx,
		&OperationsListOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlias_Create(t *testing.T) {
	// From example CreateAlias
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAliasClient(con)
	poller, err := client.BeginCreate(ctx,
		"aliasForNewSub",
		PutAliasRequest{
			Properties: &PutAliasRequestProperties{
				BillingScope: to.StringPtr("/providers/Microsoft.Billing/billingAccounts/e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31/billingProfiles/PE2Q-NOIT-BG7-TGB/invoiceSections/MTT4-OBS7-PJA-TGB"),
				DisplayName:  to.StringPtr("Contoso MCA subscription"),
				Workload:     WorkloadProduction.ToPtr(),
			},
		},
		&AliasBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.PutAliasResponse.ID == nil {
		t.Fatal("PutAliasResponse.ID should not be nil!")
	}
}

func TestAlias_Get(t *testing.T) {
	// From example GetAlias
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAliasClient(con)
	res, err := client.Get(ctx,
		"aliasForNewSub",
		&AliasGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PutAliasResponse.ID == nil {
		t.Fatal("PutAliasResponse.ID should not be nil!")
	}
}

func TestAlias_Delete(t *testing.T) {
	// From example DeleteAlias
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAliasClient(con)
	_, err := client.Delete(ctx,
		"aliasForNewSub",
		&AliasDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlias_List(t *testing.T) {
	// From example GetAlias
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAliasClient(con)
	_, err := client.List(ctx,
		&AliasListOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

// TestMain will exec each test
func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run() // exec test and this returns an exit code to pass to os
	tearDown()
	os.Exit(retCode)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func setUp() {
	ctx = context.Background()
	subscriptionId = getEnv("SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	mockHost = getEnv("AZURE_VIRTUAL_SERVER_HOST", "https://localhost:8443")

	tr := &http.Transport{}
	if err := http2.ConfigureTransport(tr); err != nil {
		fmt.Printf("Failed to configure http2 transport: %v", err)
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}
	cred, err = azidentity.NewEnvironmentCredential(&azidentity.EnvironmentCredentialOptions{AuthorityHost: mockHost, HTTPClient: client})
	if err != nil {
		panic(err)
	}

	con = arm.NewConnection(mockHost, cred, &arm.ConnectionOptions{
		Logging: policy.LogOptions{
			IncludeBody: true,
		},
		HTTPClient: client,
	})
}

func tearDown() {

}
