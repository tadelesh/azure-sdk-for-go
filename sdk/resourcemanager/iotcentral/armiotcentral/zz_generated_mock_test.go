//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotcentral

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

func TestApps_Get(t *testing.T) {
	// From example Apps_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	res, err := client.Get(ctx,
		"resRg",
		"myIoTCentralApp",
		&AppsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.App.ID == nil {
		t.Fatal("App.ID should not be nil!")
	}
}

func TestApps_CreateOrUpdate(t *testing.T) {
	// From example Apps_CreateOrUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resRg",
		"myIoTCentralApp",
		nil,
		&AppsBeginCreateOrUpdateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.App.ID == nil {
		t.Fatal("App.ID should not be nil!")
	}
}

func TestApps_Update(t *testing.T) {
	// From example Apps_Update
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	poller, err := client.BeginUpdate(ctx,
		"resRg",
		"myIoTCentralApp",
		nil,
		&AppsBeginUpdateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.App.ID == nil {
		t.Fatal("App.ID should not be nil!")
	}
}

func TestApps_Delete(t *testing.T) {
	// From example Apps_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	poller, err := client.BeginDelete(ctx,
		"resRg",
		"myIoTCentralApp",
		&AppsBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApps_ListBySubscription(t *testing.T) {
	// From example Apps_ListBySubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	pager := client.ListBySubscription(&AppsListBySubscriptionOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("App.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("App.ID should not be nil!")
			}
		}
	}
}

func TestApps_ListByResourceGroup(t *testing.T) {
	// From example Apps_ListByResourceGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	pager := client.ListByResourceGroup("resRg",
		&AppsListByResourceGroupOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("App.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("App.ID should not be nil!")
			}
		}
	}
}

func TestApps_CheckNameAvailability(t *testing.T) {
	// From example Apps_CheckNameAvailability
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	_, err := client.CheckNameAvailability(ctx,
		OperationInputs{
			Name: to.StringPtr("myiotcentralapp"),
			Type: to.StringPtr("IoTApps"),
		},
		&AppsCheckNameAvailabilityOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestApps_CheckSubdomainAvailability(t *testing.T) {
	// From example Apps_SubdomainAvailability
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	_, err := client.CheckSubdomainAvailability(ctx,
		OperationInputs{
			Name: to.StringPtr("myiotcentralapp"),
			Type: to.StringPtr("IoTApps"),
		},
		&AppsCheckSubdomainAvailabilityOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestApps_ListTemplates(t *testing.T) {
	// From example Apps_ListTemplates
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAppsClient(con,
		"00000000-0000-0000-0000-000000000000")
	pager := client.ListTemplates(&AppsListTemplatesOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
	}
}

func TestOperations_List(t *testing.T) {
	// From example Operations_List
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewOperationsClient(con)
	pager := client.List(&OperationsListOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
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
