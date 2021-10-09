//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

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

func TestStorageSyncServices_CheckNameAvailability(t *testing.T) {
	// From example StorageSyncServiceCheckNameAvailability_AlreadyExists
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"5c6bc8e1-1eaf-4192-94d8-58ce463ac86c")
	_, err := client.CheckNameAvailability(ctx,
		"westus",
		CheckNameAvailabilityParameters{
			Name: to.StringPtr("newstoragesyncservicename"),
			Type: to.StringPtr("Microsoft.StorageSync/storageSyncServices"),
		},
		&StorageSyncServicesCheckNameAvailabilityOptions{})
	if err != nil {
		t.Fatal(err)
	}

	// From example StorageSyncServiceCheckNameAvailability_Available
	_, err = client.CheckNameAvailability(ctx,
		"westus",
		CheckNameAvailabilityParameters{
			Name: to.StringPtr("newstoragesyncservicename"),
			Type: to.StringPtr("Microsoft.StorageSync/storageSyncServices"),
		},
		&StorageSyncServicesCheckNameAvailabilityOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorageSyncServices_Create(t *testing.T) {
	// From example StorageSyncServices_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		StorageSyncServiceCreateParameters{
			Location: to.StringPtr("WestUS"),
			Properties: &StorageSyncServiceCreateParametersProperties{
				IncomingTrafficPolicy: IncomingTrafficPolicyAllowAllTraffic.ToPtr(),
			},
			Tags: map[string]*string{},
		},
		&StorageSyncServicesBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.StorageSyncService.ID == nil {
		t.Fatal("StorageSyncService.ID should not be nil!")
	}
}

func TestStorageSyncServices_Get(t *testing.T) {
	// From example StorageSyncServices_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&StorageSyncServicesGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.StorageSyncService.ID == nil {
		t.Fatal("StorageSyncService.ID should not be nil!")
	}
}

func TestStorageSyncServices_Update(t *testing.T) {
	// From example StorageSyncServices_Update
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginUpdate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&StorageSyncServicesBeginUpdateOptions{Parameters: &StorageSyncServiceUpdateParameters{
			Properties: &StorageSyncServiceUpdateProperties{
				IncomingTrafficPolicy: IncomingTrafficPolicyAllowAllTraffic.ToPtr(),
			},
			Tags: map[string]*string{
				"Dept":        to.StringPtr("IT"),
				"Environment": to.StringPtr("Test"),
			},
		},
		})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.StorageSyncService.ID == nil {
		t.Fatal("StorageSyncService.ID should not be nil!")
	}
}

func TestStorageSyncServices_Delete(t *testing.T) {
	// From example StorageSyncServices_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginDelete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&StorageSyncServicesBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorageSyncServices_ListByResourceGroup(t *testing.T) {
	// From example StorageSyncServices_ListByResourceGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListByResourceGroup(ctx,
		"SampleResourceGroup_1",
		&StorageSyncServicesListByResourceGroupOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorageSyncServices_ListBySubscription(t *testing.T) {
	// From example StorageSyncServices_ListBySubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewStorageSyncServicesClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListBySubscription(ctx,
		&StorageSyncServicesListBySubscriptionOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateLinkResources_ListByStorageSyncService(t *testing.T) {
	// From example PrivateLinkResources_List
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateLinkResourcesClient(con,
		"{subscription-id}")
	_, err := client.ListByStorageSyncService(ctx,
		"res6977",
		"sss2527",
		&PrivateLinkResourcesListByStorageSyncServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateEndpointConnections_Get(t *testing.T) {
	// From example PrivateEndpointConnections_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"{subscription-id}")
	res, err := client.Get(ctx,
		"res6977",
		"sss2527",
		"{privateEndpointConnectionName}",
		&PrivateEndpointConnectionsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.PrivateEndpointConnection.ID == nil {
		t.Fatal("PrivateEndpointConnection.ID should not be nil!")
	}
}

func TestPrivateEndpointConnections_Create(t *testing.T) {
	// From example PrivateEndpointConnections_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreate(ctx,
		"res7687",
		"sss2527",
		"{privateEndpointConnectionName}",
		PrivateEndpointConnection{
			Properties: &PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &PrivateLinkServiceConnectionState{
					Description: to.StringPtr("Auto-Approved"),
					Status:      PrivateEndpointServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		&PrivateEndpointConnectionsBeginCreateOptions{})
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
	// From example PrivateEndpointConnections_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"res6977",
		"sss2527",
		"{privateEndpointConnectionName}",
		&PrivateEndpointConnectionsBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateEndpointConnections_ListByStorageSyncService(t *testing.T) {
	// From example PrivateEndpointConnections_ListByStorageSyncService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewPrivateEndpointConnectionsClient(con,
		"{subscription-id}")
	_, err := client.ListByStorageSyncService(ctx,
		"res6977",
		"sss2527",
		&PrivateEndpointConnectionsListByStorageSyncServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSyncGroups_ListByStorageSyncService(t *testing.T) {
	// From example SyncGroups_ListByStorageSyncService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSyncGroupsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListByStorageSyncService(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&SyncGroupsListByStorageSyncServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSyncGroups_Create(t *testing.T) {
	// From example SyncGroups_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSyncGroupsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Create(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		SyncGroupCreateParameters{
			Properties: map[string]interface{}{},
		},
		&SyncGroupsCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SyncGroup.ID == nil {
		t.Fatal("SyncGroup.ID should not be nil!")
	}
}

func TestSyncGroups_Get(t *testing.T) {
	// From example SyncGroups_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSyncGroupsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		&SyncGroupsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.SyncGroup.ID == nil {
		t.Fatal("SyncGroup.ID should not be nil!")
	}
}

func TestSyncGroups_Delete(t *testing.T) {
	// From example SyncGroups_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSyncGroupsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.Delete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		&SyncGroupsDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_Create(t *testing.T) {
	// From example CloudEndpoints_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		CloudEndpointCreateParameters{
			Properties: &CloudEndpointCreateParametersProperties{
				AzureFileShareName:       to.StringPtr("cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4"),
				FriendlyName:             to.StringPtr("ankushbsubscriptionmgmtmab"),
				StorageAccountResourceID: to.StringPtr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage"),
				StorageAccountTenantID:   to.StringPtr("\"72f988bf-86f1-41af-91ab-2d7cd011db47\""),
			},
		},
		&CloudEndpointsBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.CloudEndpoint.ID == nil {
		t.Fatal("CloudEndpoint.ID should not be nil!")
	}
}

func TestCloudEndpoints_Get(t *testing.T) {
	// From example CloudEndpoints_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		&CloudEndpointsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.CloudEndpoint.ID == nil {
		t.Fatal("CloudEndpoint.ID should not be nil!")
	}
}

func TestCloudEndpoints_Delete(t *testing.T) {
	// From example CloudEndpoints_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginDelete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		&CloudEndpointsBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_ListBySyncGroup(t *testing.T) {
	// From example CloudEndpoints_ListBySyncGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListBySyncGroup(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		&CloudEndpointsListBySyncGroupOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_PreBackup(t *testing.T) {
	// From example CloudEndpoints_PreBackup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginPreBackup(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		BackupRequest{
			AzureFileShare: to.StringPtr("https://sampleserver.file.core.test-cint.azure-test.net/sampleFileShare"),
		},
		&CloudEndpointsBeginPreBackupOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_PostBackup(t *testing.T) {
	// From example CloudEndpoints_PostBackup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginPostBackup(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		BackupRequest{
			AzureFileShare: to.StringPtr("https://sampleserver.file.core.test-cint.azure-test.net/sampleFileShare"),
		},
		&CloudEndpointsBeginPostBackupOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_PreRestore(t *testing.T) {
	// From example CloudEndpoints_PreRestore
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginPreRestore(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		PreRestoreRequest{
			AzureFileShareURI: to.StringPtr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
			RestoreFileSpec: []*RestoreFileSpec{
				{
					Path:  to.StringPtr("text1.txt"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("MyDir"),
					Isdir: to.BoolPtr(true),
				},
				{
					Path:  to.StringPtr("MyDir/SubDir"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("MyDir/SubDir/File1.pdf"),
					Isdir: to.BoolPtr(false),
				}},
		},
		&CloudEndpointsBeginPreRestoreOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_Restoreheartbeat(t *testing.T) {
	// From example CloudEndpoints_restoreheartbeat
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.Restoreheartbeat(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		&CloudEndpointsRestoreheartbeatOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_PostRestore(t *testing.T) {
	// From example CloudEndpoints_PostRestore
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginPostRestore(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		PostRestoreRequest{
			AzureFileShareURI: to.StringPtr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
			RestoreFileSpec: []*RestoreFileSpec{
				{
					Path:  to.StringPtr("text1.txt"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("MyDir"),
					Isdir: to.BoolPtr(true),
				},
				{
					Path:  to.StringPtr("MyDir/SubDir"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("MyDir/SubDir/File1.pdf"),
					Isdir: to.BoolPtr(false),
				}},
			SourceAzureFileShareURI: to.StringPtr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
			Status:                  to.StringPtr("Succeeded"),
		},
		&CloudEndpointsBeginPostRestoreOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloudEndpoints_TriggerChangeDetection(t *testing.T) {
	// From example CloudEndpoints_TriggerChangeDetection
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginTriggerChangeDetection(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		TriggerChangeDetectionParameters{
			ChangeDetectionMode: ChangeDetectionModeRecursive.ToPtr(),
			DirectoryPath:       to.StringPtr("NewDirectory"),
		},
		&CloudEndpointsBeginTriggerChangeDetectionOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServerEndpoints_Create(t *testing.T) {
	// From example ServerEndpoints_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		ServerEndpointCreateParameters{
			Properties: &ServerEndpointCreateParametersProperties{
				CloudTiering:                 FeatureStatusOff.ToPtr(),
				InitialDownloadPolicy:        InitialDownloadPolicyNamespaceThenModifiedFiles.ToPtr(),
				InitialUploadPolicy:          InitialUploadPolicyServerAuthoritative.ToPtr(),
				LocalCacheMode:               LocalCacheModeUpdateLocallyCachedFiles.ToPtr(),
				OfflineDataTransfer:          FeatureStatusOn.ToPtr(),
				OfflineDataTransferShareName: to.StringPtr("myfileshare"),
				ServerLocalPath:              to.StringPtr("D:\\SampleServerEndpoint_1"),
				ServerResourceID:             to.StringPtr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
				TierFilesOlderThanDays:       to.Int32Ptr(0),
				VolumeFreeSpacePercent:       to.Int32Ptr(100),
			},
		},
		&ServerEndpointsBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.ServerEndpoint.ID == nil {
		t.Fatal("ServerEndpoint.ID should not be nil!")
	}
}

func TestServerEndpoints_Update(t *testing.T) {
	// From example ServerEndpoints_Update
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginUpdate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		&ServerEndpointsBeginUpdateOptions{Parameters: &ServerEndpointUpdateParameters{
			Properties: &ServerEndpointUpdateProperties{
				CloudTiering:           FeatureStatusOff.ToPtr(),
				LocalCacheMode:         LocalCacheModeUpdateLocallyCachedFiles.ToPtr(),
				OfflineDataTransfer:    FeatureStatusOff.ToPtr(),
				TierFilesOlderThanDays: to.Int32Ptr(0),
				VolumeFreeSpacePercent: to.Int32Ptr(100),
			},
		},
		})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.ServerEndpoint.ID == nil {
		t.Fatal("ServerEndpoint.ID should not be nil!")
	}
}

func TestServerEndpoints_Get(t *testing.T) {
	// From example ServerEndpoints_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		&ServerEndpointsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.ServerEndpoint.ID == nil {
		t.Fatal("ServerEndpoint.ID should not be nil!")
	}
}

func TestServerEndpoints_Delete(t *testing.T) {
	// From example ServerEndpoints_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginDelete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		&ServerEndpointsBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServerEndpoints_ListBySyncGroup(t *testing.T) {
	// From example ServerEndpoints_ListBySyncGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListBySyncGroup(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		&ServerEndpointsListBySyncGroupOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestServerEndpoints_RecallAction(t *testing.T) {
	// From example ServerEndpoints_recallAction
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewServerEndpointsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginRecallAction(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		RecallActionParameters{
			Pattern:    to.StringPtr(""),
			RecallPath: to.StringPtr(""),
		},
		&ServerEndpointsBeginRecallActionOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisteredServers_ListByStorageSyncService(t *testing.T) {
	// From example RegisteredServers_ListByStorageSyncService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRegisteredServersClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListByStorageSyncService(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&RegisteredServersListByStorageSyncServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisteredServers_Get(t *testing.T) {
	// From example RegisteredServers_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRegisteredServersClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"080d4133-bdb5-40a0-96a0-71a6057bfe9a",
		&RegisteredServersGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.RegisteredServer.ID == nil {
		t.Fatal("RegisteredServer.ID should not be nil!")
	}
}

func TestRegisteredServers_Create(t *testing.T) {
	// From example RegisteredServers_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRegisteredServersClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"080d4133-bdb5-40a0-96a0-71a6057bfe9a",
		RegisteredServerCreateParameters{
			Properties: &RegisteredServerCreateParametersProperties{
				AgentVersion:      to.StringPtr("1.0.277.0"),
				FriendlyName:      to.StringPtr("afscv-2304-139"),
				ServerCertificate: to.StringPtr("MIIDFjCCAf6gAwIBAgIQQS+DS8uhc4VNzUkTw7wbRjANBgkqhkiG9w0BAQ0FADAzMTEwLwYDVQQDEyhhbmt1c2hiLXByb2QzLnJlZG1vbmQuY29ycC5taWNyb3NvZnQuY29tMB4XDTE3MDgwMzE3MDQyNFoXDTE4MDgwNDE3MDQyNFowMzExMC8GA1UEAxMoYW5rdXNoYi1wcm9kMy5yZWRtb25kLmNvcnAubWljcm9zb2Z0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALDRvV4gmsIy6jGDPiHsXmvgVP749NNP7DopdlbHaNhjFmYINHl0uWylyaZmgJrROt2mnxN/zEyJtGnqYHlzUr4xvGq/qV5pqgdB9tag/sw9i22gfe9PRZ0FmSOZnXMbLYgLiDFqLtut5gHcOuWMj03YnkfoBEKlFBxWbagvW2yxz/Sxi9OVSJOKCaXra0RpcIHrO/KFl6ho2eE1/7Ykmfa8hZvSdoPd5gHdLiQcMB/pxq+mWp1fI6c8vFZoDu7Atn+NXTzYPKUxKzaisF12TsaKpohUsJpbB3Wocb0F5frn614D2pg14ERB5otjAMWw1m65csQWPI6dP8KIYe0+QPkCAwEAAaMmMCQwIgYDVR0lAQH/BBgwFgYIKwYBBQUHAwIGCisGAQQBgjcKAwwwDQYJKoZIhvcNAQENBQADggEBAA4RhVIBkw34M1RwakJgHvtjsOFxF1tVQA941NtLokx1l2Z8+GFQkcG4xpZSt+UN6wLerdCbnNhtkCErWUDeaT0jxk4g71Ofex7iM04crT4iHJr8mi96/XnhnkTUs+GDk12VgdeeNEczMZz+8Mxw9dJ5NCnYgTwO0SzGlclRsDvjzkLo8rh2ZG6n/jKrEyNXXo+hOqhupij0QbRP2Tvexdfw201kgN1jdZify8XzJ8Oi0bTS0KpJf2pNPOlooK2bjMUei9ANtEdXwwfVZGWvVh6tJjdv6k14wWWJ1L7zhA1IIVb1J+sQUzJji5iX0DrezjTz1Fg+gAzITaA/WsuujlM="),
				ServerID:          to.StringPtr("080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
				ServerOSVersion:   to.StringPtr("10.0.14393.0"),
				ServerRole:        to.StringPtr("Standalone"),
			},
		},
		&RegisteredServersBeginCreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if res.RegisteredServer.ID == nil {
		t.Fatal("RegisteredServer.ID should not be nil!")
	}
}

func TestRegisteredServers_Delete(t *testing.T) {
	// From example RegisteredServers_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRegisteredServersClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginDelete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"41166691-ab03-43e9-ab3e-0330eda162ac",
		&RegisteredServersBeginDeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisteredServers_TriggerRollover(t *testing.T) {
	// From example RegisteredServers_triggerRollover
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRegisteredServersClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	poller, err := client.BeginTriggerRollover(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"d166ca76-dad2-49df-b409-12345642d730",
		TriggerRolloverRequest{
			ServerCertificate: to.StringPtr("\"MIIDFjCCAf6gAwIBAgIQQS+DS8uhc4VNzUkTw7wbRjANBgkqhkiG9w0BAQ0FADAzMTEwLwYDVQQDEyhhbmt1c2hiLXByb2QzLnJlZG1vbmQuY29ycC5taWNyb3NvZnQuY29tMB4XDTE3MDgwMzE3MDQyNFoXDTE4MDgwNDE3MDQyNFowMzExMC8GA1UEAxMoYW5rdXNoYi1wcm9kMy5yZWRtb25kLmNvcnAubWljcm9zb2Z0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALDRvV4gmsIy6jGDPiHsXmvgVP749NNP7DopdlbHaNhjFmYINHl0uWylyaZmgJrROt2mnxN/zEyJtGnqYHlzUr4xvGq/qV5pqgdB9tag/sw9i22gfe9PRZ0FmSOZnXMbLYgLiDFqLtut5gHcOuWMj03YnkfoBEKlFBxWbagvW2yxz/Sxi9OVSJOKCaXra0RpcIHrO/KFl6ho2eE1/7Ykmfa8hZvSdoPd5gHdLiQcMB/pxq+mWp1fI6c8vFZoDu7Atn+NXTzYPKUxKzaisF12TsaKpohUsJpbB3Wocb0F5frn614D2pg14ERB5otjAMWw1m65csQWPI6dP8KIYe0+QPkCAwEAAaMmMCQwIgYDVR0lAQH/BBgwFgYIKwYBBQUHAwIGCisGAQQBgjcKAwwwDQYJKoZIhvcNAQENBQADggEBAA4RhVIBkw34M1RwakJgHvtjsOFxF1tVQA941NtLokx1l2Z8+GFQkcG4xpZSt+UN6wLerdCbnNhtkCErWUDeaT0jxk4g71Ofex7iM04crT4iHJr8mi96/XnhnkTUs+GDk12VgdeeNEczMZz+8Mxw9dJ5NCnYgTwO0SzGlclRsDvjzkLo8rh2ZG6n/jKrEyNXXo+hOqhupij0QbRP2Tvexdfw201kgN1jdZify8XzJ8Oi0bTS0KpJf2pNPOlooK2bjMUei9ANtEdXwwfVZGWvVh6tJjdv6k14wWWJ1L7zhA1IIVb1J+sQUzJji5iX0DrezjTz1Fg+gAzITaA/WsuujlM=\""),
		},
		&RegisteredServersBeginTriggerRolloverOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWorkflows_ListByStorageSyncService(t *testing.T) {
	// From example Workflows_ListByStorageSyncService
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewWorkflowsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.ListByStorageSyncService(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		&WorkflowsListByStorageSyncServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestWorkflows_Get(t *testing.T) {
	// From example Workflows_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewWorkflowsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	res, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"828219ea-083e-48b5-89ea-8fd9991b2e75",
		&WorkflowsGetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Workflow.ID == nil {
		t.Fatal("Workflow.ID should not be nil!")
	}
}

func TestWorkflows_Abort(t *testing.T) {
	// From example Workflows_Abort
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewWorkflowsClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.Abort(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"7ffd50b3-5574-478d-9ff2-9371bc42ce68",
		&WorkflowsAbortOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestOperationStatus_Get(t *testing.T) {
	// From example Workflows_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewOperationStatusClient(con,
		"52b8da2f-61e0-4a1f-8dde-336911f367fb")
	_, err := client.Get(ctx,
		"SampleResourceGroup_1",
		"westus",
		"828219ea-083e-48b5-89ea-8fd9991b2e75",
		"14b50e24-f68d-4b29-a882-38be9dfb8bd1",
		&OperationStatusGetOptions{})
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
