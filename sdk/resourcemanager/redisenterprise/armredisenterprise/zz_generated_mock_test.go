//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredisenterprise

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

func TestOperationsStatus_Get(t *testing.T) {
	// From example OperationsStatusGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewOperationsStatusClient(con,
		"subid")
	res, err := client.Get(ctx,
		"West US",
		"testoperationid",
		&OperationsStatusGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.OperationStatus.ID == nil {
		t.Fatal("OperationStatus.ID should not be nil!")
	}
}

func TestRedisEnterprise_Create(t *testing.T) {
	// From example RedisEnterpriseCreate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cache1",
		Cluster{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("West US"),
				Tags: map[string]*string{
					"tag1": to.StringPtr("value1"),
				},
			},
			Properties: &ClusterProperties{
				MinimumTLSVersion: TLSVersionOne2.ToPtr(),
			},
			SKU: &SKU{
				Name:     SKUNameEnterpriseFlashF300.ToPtr(),
				Capacity: to.Int32Ptr(3),
			},
			Zones: []*string{
				to.StringPtr("1"),
				to.StringPtr("2"),
				to.StringPtr("3")},
		},
		&RedisEnterpriseBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.Cluster.ID == nil {
		t.Fatal("Cluster.ID should not be nil!")
	}
}

func TestRedisEnterprise_Update(t *testing.T) {
	// From example RedisEnterpriseUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"cache1",
		ClusterUpdate{
			Properties: &ClusterProperties{
				MinimumTLSVersion: TLSVersionOne2.ToPtr(),
			},
			SKU: &SKU{
				Name:     SKUNameEnterpriseFlashF300.ToPtr(),
				Capacity: to.Int32Ptr(9),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
			},
		},
		&RedisEnterpriseBeginUpdateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.Cluster.ID == nil {
		t.Fatal("Cluster.ID should not be nil!")
	}
}

func TestRedisEnterprise_Delete(t *testing.T) {
	// From example RedisEnterpriseDelete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"cache1",
		&RedisEnterpriseBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisEnterprise_Get(t *testing.T) {
	// From example RedisEnterpriseGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"cache1",
		&RedisEnterpriseGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Cluster.ID == nil {
		t.Fatal("Cluster.ID should not be nil!")
	}
}

func TestRedisEnterprise_ListByResourceGroup(t *testing.T) {
	// From example RedisEnterpriseListByResourceGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	pager := client.ListByResourceGroup("rg1",
		&RedisEnterpriseListByResourceGroupOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("Cluster.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("Cluster.ID should not be nil!")
			}
		}
	}
}

func TestRedisEnterprise_List(t *testing.T) {
	// From example RedisEnterpriseList
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRedisEnterpriseClient(con,
		"subid")
	pager := client.List(&RedisEnterpriseListOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("Cluster.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("Cluster.ID should not be nil!")
			}
		}
	}
}

func TestDatabases_ListByCluster(t *testing.T) {
	// From example RedisEnterpriseDatabasesListByCluster
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	pager := client.ListByCluster("rg1",
		"cache1",
		&DatabasesListByClusterOptions{})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			fmt.Printf("Database.ID: %s\n", *v.ID)
			if v.ID == nil {
				t.Fatal("Database.ID should not be nil!")
			}
		}
	}
}

func TestDatabases_Create(t *testing.T) {
	// From example RedisEnterpriseDatabasesCreate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cache1",
		"default",
		Database{
			Properties: &DatabaseProperties{
				ClientProtocol:   ProtocolEncrypted.ToPtr(),
				ClusteringPolicy: ClusteringPolicyEnterpriseCluster.ToPtr(),
				EvictionPolicy:   EvictionPolicyAllKeysLRU.ToPtr(),
				Modules: []*Module{
					{
						Name: to.StringPtr("RedisBloom"),
						Args: to.StringPtr("ERROR_RATE 0.00 INITIAL_SIZE 400"),
					},
					{
						Name: to.StringPtr("RedisTimeSeries"),
						Args: to.StringPtr("RETENTION_POLICY 20"),
					},
					{
						Name: to.StringPtr("RediSearch"),
					}},
				Persistence: &Persistence{
					AofEnabled:   to.BoolPtr(true),
					AofFrequency: AofFrequencyOneS.ToPtr(),
				},
				Port: to.Int32Ptr(10000),
			},
		},
		&DatabasesBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.Database.ID == nil {
		t.Fatal("Database.ID should not be nil!")
	}
}

func TestDatabases_Update(t *testing.T) {
	// From example RedisEnterpriseDatabasesUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"cache1",
		"default",
		DatabaseUpdate{
			Properties: &DatabaseProperties{
				ClientProtocol: ProtocolEncrypted.ToPtr(),
				EvictionPolicy: EvictionPolicyAllKeysLRU.ToPtr(),
				Persistence: &Persistence{
					RdbEnabled:   to.BoolPtr(true),
					RdbFrequency: RdbFrequencyTwelveH.ToPtr(),
				},
			},
		},
		&DatabasesBeginUpdateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.Database.ID == nil {
		t.Fatal("Database.ID should not be nil!")
	}
}

func TestDatabases_Get(t *testing.T) {
	// From example RedisEnterpriseDatabasesGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"cache1",
		"default",
		&DatabasesGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Database.ID == nil {
		t.Fatal("Database.ID should not be nil!")
	}
}

func TestDatabases_Delete(t *testing.T) {
	// From example RedisEnterpriseDatabasesDelete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"cache1",
		"db1",
		&DatabasesBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabases_ListKeys(t *testing.T) {
	// From example RedisEnterpriseDatabasesListKeys
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	_, err := client.ListKeys(ctx,
		"rg1",
		"cache1",
		"default",
		&DatabasesListKeysOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabases_RegenerateKey(t *testing.T) {
	// From example RedisEnterpriseDatabasesRegenerateKey
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginRegenerateKey(ctx,
		"rg1",
		"cache1",
		"default",
		RegenerateKeyParameters{
			KeyType: AccessKeyTypePrimary.ToPtr(),
		},
		&DatabasesBeginRegenerateKeyOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabases_Import(t *testing.T) {
	// From example RedisEnterpriseDatabasesImport
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginImport(ctx,
		"rg1",
		"cache1",
		"default",
		ImportClusterParameters{
			SasUris: []*string{
				to.StringPtr("https://contosostorage.blob.core.window.net/urltoBlobFile1?sasKeyParameters"),
				to.StringPtr("https://contosostorage.blob.core.window.net/urltoBlobFile2?sasKeyParameters")},
		},
		&DatabasesBeginImportOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabases_Export(t *testing.T) {
	// From example RedisEnterpriseDatabasesExport
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDatabasesClient(con,
		"subid")
	poller, err := client.BeginExport(ctx,
		"rg1",
		"cache1",
		"default",
		ExportClusterParameters{
			SasURI: to.StringPtr("https://contosostorage.blob.core.window.net/urlToBlobContainer?sasKeyParameters"),
		},
		&DatabasesBeginExportOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateEndpointConnections_List(t *testing.T) {
	// From example RedisEnterpriseListPrivateEndpointConnections
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	_, err := client.List(ctx,
		"rg1",
		"cache1",
		&PrivateEndpointConnectionsListOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateEndpointConnections_Get(t *testing.T) {
	// From example RedisEnterpriseGetPrivateEndpointConnection
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	res, err := client.Get(ctx,
		"rg1",
		"cache1",
		"pectest01",
		&PrivateEndpointConnectionsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_Put(t *testing.T) {
	// From example RedisEnterprisePutPrivateEndpointConnection
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	poller, err := client.BeginPut(ctx,
		"rg1",
		"cache1",
		"pectest01",
		PrivateEndpointConnection{
			Properties: &PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &PrivateLinkServiceConnectionState{
					Description: to.StringPtr("Auto-Approved"),
					Status:      PrivateEndpointServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		&PrivateEndpointConnectionsBeginPutOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_Delete(t *testing.T) {
	// From example RedisEnterpriseDeletePrivateEndpointConnection
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"subid")
	_, err := client.Delete(ctx,
		"rg1",
		"cache1",
		"pectest01",
		&PrivateEndpointConnectionsDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateLinkResources_ListByCluster(t *testing.T) {
	// From example RedisEnterpriseListPrivateLinkResources
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateLinkResourcesClient(con,
		"subid")
	_, err := client.ListByCluster(ctx,
		"rg1",
		"cache1",
		&PrivateLinkResourcesListByClusterOptions{})
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
