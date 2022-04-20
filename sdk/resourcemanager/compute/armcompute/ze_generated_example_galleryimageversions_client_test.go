//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImageVersionWithVMAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	galleryImageVersionsClient, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	galleryImageVersionsClientCreateOrUpdateResponsePoller, err := galleryImageVersionsClient.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersion{
			Location: to.Ptr("<resource-location>"),
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name: to.Ptr("<the-name-of-the-region.>"),
							Encryption: &armcompute.EncryptionImages{
								DataDiskImages: []*armcompute.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &armcompute.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name: to.Ptr("<the-name-of-the-region.>"),
							Encryption: &armcompute.EncryptionImages{
								DataDiskImages: []*armcompute.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &armcompute.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<a-relative-uri-containing-the-resource-id-of-the-disk-encryption-set.>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.Ptr("<the-id-of-the-gallery-artifact-version-source.-can-specify-a-disk-uri,-snapshot-uri,-user-image-or-storage-account-resource.>"),
					},
				},
			},
		},
		&armcompute.GalleryImageVersionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	galleryImageVersionsClientCreateOrUpdateResponse, err := galleryImageVersionsClientCreateOrUpdateResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = galleryImageVersionsClientCreateOrUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/UpdateASimpleGalleryImageVersion.json
func ExampleGalleryImageVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	galleryImageVersionsClient, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	galleryImageVersionsClientUpdateResponsePoller, err := galleryImageVersionsClient.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersionUpdate{
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name:                 to.Ptr("<the-name-of-the-region.>"),
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name:                 to.Ptr("<the-name-of-the-region.>"),
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.Ptr("<the-id-of-the-gallery-artifact-version-source.-can-specify-a-disk-uri,-snapshot-uri,-user-image-or-storage-account-resource.>"),
					},
				},
			},
		},
		&armcompute.GalleryImageVersionsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	galleryImageVersionsClientUpdateResponse, err := galleryImageVersionsClientUpdateResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = galleryImageVersionsClientUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryImageVersionWithReplicationStatus.json
func ExampleGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	galleryImageVersionsClient, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	galleryImageVersionsClientGetResponse, err := galleryImageVersionsClient.Get(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&armcompute.GalleryImageVersionsClientGetOptions{Expand: to.Ptr(armcompute.ReplicationStatusTypesReplicationStatus)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = galleryImageVersionsClientGetResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/DeleteAGalleryImageVersion.json
func ExampleGalleryImageVersionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	galleryImageVersionsClient, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	galleryImageVersionsClientDeleteResponsePoller, err := galleryImageVersionsClient.BeginDelete(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&armcompute.GalleryImageVersionsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = galleryImageVersionsClientDeleteResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/ListGalleryImageVersionsInAGalleryImage.json
func ExampleGalleryImageVersionsClient_NewListByGalleryImagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	galleryImageVersionsClient, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	galleryImageVersionsClientNewListByGalleryImagePager := galleryImageVersionsClient.NewListByGalleryImagePager("<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		nil)
	for galleryImageVersionsClientNewListByGalleryImagePager.More() {
		nextResult, err := galleryImageVersionsClientNewListByGalleryImagePager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
