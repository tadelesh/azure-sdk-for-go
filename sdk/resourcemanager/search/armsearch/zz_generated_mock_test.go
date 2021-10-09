//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

func TestOperations_List(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestAdminKeys_Get(t *testing.T) {
	// From example SearchGetAdminKeys
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAdminKeysClient(con,
		"subid")
	_, err := client.Get(ctx,
		"rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdminKeys_Regenerate(t *testing.T) {
	// From example SearchRegenerateAdminKey
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAdminKeysClient(con,
		"subid")
	_, err := client.Regenerate(ctx,
		"rg1",
		"mysearchservice",
		AdminKeyKindPrimary,
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryKeys_Create(t *testing.T) {
	// From example SearchCreateQueryKey
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewQueryKeysClient(con,
		"subid")
	_, err := client.Create(ctx,
		"rg1",
		"mysearchservice",
		"Query key for browser-based clients",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryKeys_ListBySearchService(t *testing.T) {
	// From example SearchListQueryKeysBySearchService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewQueryKeysClient(con,
		"subid")
	pager := client.ListBySearchService("rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
	}
}

func TestQueryKeys_Delete(t *testing.T) {
	// From example SearchDeleteQueryKey
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewQueryKeysClient(con,
		"subid")
	_, err := client.Delete(ctx,
		"rg1",
		"mysearchservice",
		"<a query API key>",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestServices_CreateOrUpdate(t *testing.T) {
	// From example SearchCreateOrUpdateService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		SearchService{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"app-name": to.StringPtr("My e-commerce app"),
				},
			},
			Properties: &SearchServiceProperties{
				HostingMode:    HostingModeDefault.ToPtr(),
				PartitionCount: to.Int32Ptr(1),
				ReplicaCount:   to.Int32Ptr(3),
			},
			SKU: &SKU{
				Name: SKUNameStandard.ToPtr(),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example SearchCreateOrUpdateServiceToAllowAccessFromPrivateEndpoints
	poller, err = client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		SearchService{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"app-name": to.StringPtr("My e-commerce app"),
				},
			},
			Properties: &SearchServiceProperties{
				HostingMode:         HostingModeDefault.ToPtr(),
				PartitionCount:      to.Int32Ptr(1),
				PublicNetworkAccess: PublicNetworkAccessDisabled.ToPtr(),
				ReplicaCount:        to.Int32Ptr(3),
			},
			SKU: &SKU{
				Name: SKUNameStandard.ToPtr(),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example SearchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPs
	poller, err = client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		SearchService{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"app-name": to.StringPtr("My e-commerce app"),
				},
			},
			Properties: &SearchServiceProperties{
				HostingMode: HostingModeDefault.ToPtr(),
				NetworkRuleSet: &NetworkRuleSet{
					IPRules: []*IPRule{
						{
							Value: to.StringPtr("123.4.5.6"),
						},
						{
							Value: to.StringPtr("123.4.6.0/18"),
						}},
				},
				PartitionCount: to.Int32Ptr(1),
				ReplicaCount:   to.Int32Ptr(1),
			},
			SKU: &SKU{
				Name: SKUNameStandard.ToPtr(),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example SearchCreateOrUpdateServiceWithIdentity
	poller, err = client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		SearchService{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"app-name": to.StringPtr("My e-commerce app"),
				},
			},
			Identity: &Identity{
				Type: IdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &SearchServiceProperties{
				HostingMode:    HostingModeDefault.ToPtr(),
				PartitionCount: to.Int32Ptr(1),
				ReplicaCount:   to.Int32Ptr(3),
			},
			SKU: &SKU{
				Name: SKUNameStandard.ToPtr(),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}
}

func TestServices_Update(t *testing.T) {
	// From example SearchUpdateService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	res, err := client.Update(ctx,
		"rg1",
		"mysearchservice",
		SearchServiceUpdate{
			Properties: &SearchServiceProperties{
				ReplicaCount: to.Int32Ptr(2),
			},
			Tags: map[string]*string{
				"app-name": to.StringPtr("My e-commerce app"),
				"new-tag":  to.StringPtr("Adding a new tag"),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example SearchUpdateServiceToRemoveIdentity
	res, err = client.Update(ctx,
		"rg1",
		"mysearchservice",
		SearchServiceUpdate{
			Identity: &Identity{
				Type: IdentityTypeNone.ToPtr(),
			},
			SKU: &SKU{
				Name: SKUNameStandard.ToPtr(),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example searchUpdateServiceToAllowAccessFromPrivateEndpoints
	res, err = client.Update(ctx,
		"rg1",
		"mysearchservice",
		SearchServiceUpdate{
			Properties: &SearchServiceProperties{
				PartitionCount:      to.Int32Ptr(1),
				PublicNetworkAccess: PublicNetworkAccessDisabled.ToPtr(),
				ReplicaCount:        to.Int32Ptr(1),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}

	// From example searchUpdateServiceToAllowAccessFromPublicCustomIPs
	res, err = client.Update(ctx,
		"rg1",
		"mysearchservice",
		SearchServiceUpdate{
			Properties: &SearchServiceProperties{
				NetworkRuleSet: &NetworkRuleSet{
					IPRules: []*IPRule{
						{
							Value: to.StringPtr("123.4.5.6"),
						},
						{
							Value: to.StringPtr("123.4.6.0/18"),
						}},
				},
				PartitionCount:      to.Int32Ptr(1),
				PublicNetworkAccess: PublicNetworkAccessEnabled.ToPtr(),
				ReplicaCount:        to.Int32Ptr(3),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}
}

func TestServices_Get(t *testing.T) {
	// From example SearchGetService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SearchService.ID == nil {
		t.Fatal("SearchService.ID should not be nil!")
	}
}

func TestServices_Delete(t *testing.T) {
	// From example SearchDeleteService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	_, err := client.Delete(ctx,
		"rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestServices_ListByResourceGroup(t *testing.T) {
	// From example SearchListServicesByResourceGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	pager := client.ListByResourceGroup("rg1",
		&SearchManagementRequestOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("SearchService.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("SearchService.ID should not be nil!")
			}
		}
	}
}

func TestServices_ListBySubscription(t *testing.T) {
	// From example SearchListServicesBySubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	pager := client.ListBySubscription(&SearchManagementRequestOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("SearchService.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("SearchService.ID should not be nil!")
			}
		}
	}
}

func TestServices_CheckNameAvailability(t *testing.T) {
	// From example SearchCheckNameAvailability
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServicesClient(con,
		"subid")
	_, err := client.CheckNameAvailability(ctx,
		CheckNameAvailabilityInput{
			Name: to.StringPtr("mysearchservice"),
			Type: to.StringPtr("searchServices"),
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateLinkResources_ListSupported(t *testing.T) {
	// From example ListSupportedPrivateLinkResources
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateLinkResourcesClient(con,
		"subid")
	_, err := client.ListSupported(ctx,
		"rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateEndpointConnections_Update(t *testing.T) {
	// From example PrivateEndpointConnectionUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	res, err := client.Update(ctx,
		"rg1",
		"mysearchservice",
		"testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546",
		PrivateEndpointConnection{
			Properties: &PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState{
					Description: to.StringPtr("Rejected for some reason"),
					Status:      PrivateLinkServiceConnectionStatusRejected.ToPtr(),
				},
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_Get(t *testing.T) {
	// From example PrivateEndpointConnectionGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"mysearchservice",
		"testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_Delete(t *testing.T) {
	// From example PrivateEndpointConnectionDelete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	res, err := client.Delete(ctx,
		"rg1",
		"mysearchservice",
		"testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_ListByService(t *testing.T) {
	// From example ListPrivateEndpointConnectionsByService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	pager := client.ListByService("rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("PrivateEndpointConnection.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("PrivateEndpointConnection.ID should not be nil!")
			}
		}
	}
}

func TestSharedPrivateLinkResources_CreateOrUpdate(t *testing.T) {
	// From example SharedPrivateLinkResourceCreateOrUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedPrivateLinkResourcesClient(con,
		"subid")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"mysearchservice",
		"testResource",
		SharedPrivateLinkResource{
			Properties: &SharedPrivateLinkResourceProperties{
				GroupID:               to.StringPtr("blob"),
				PrivateLinkResourceID: to.StringPtr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
				RequestMessage:        to.StringPtr("please approve"),
				ResourceRegion:        to.StringPtr("null"),
			},
		},
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.SharedPrivateLinkResource.ID == nil {
		t.Fatal("SharedPrivateLinkResource.ID should not be nil!")
	}
}

func TestSharedPrivateLinkResources_Get(t *testing.T) {
	// From example SharedPrivateLinkResourceGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedPrivateLinkResourcesClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"mysearchservice",
		"testResource",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SharedPrivateLinkResource.ID == nil {
		t.Fatal("SharedPrivateLinkResource.ID should not be nil!")
	}
}

func TestSharedPrivateLinkResources_Delete(t *testing.T) {
	// From example SharedPrivateLinkResourceDelete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedPrivateLinkResourcesClient(con,
		"subid")
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"mysearchservice",
		"testResource",
		&SearchManagementRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSharedPrivateLinkResources_ListByService(t *testing.T) {
	// From example ListSharedPrivateLinkResourcesByService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedPrivateLinkResourcesClient(con,
		"subid")
	pager := client.ListByService("rg1",
		"mysearchservice",
		&SearchManagementRequestOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("SharedPrivateLinkResource.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("SharedPrivateLinkResource.ID should not be nil!")
			}
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
