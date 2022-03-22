//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BackupProtectionContainersClient contains the methods for the BackupProtectionContainers group.
// Don't use this type directly, use NewBackupProtectionContainersClient() instead.
type BackupProtectionContainersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupProtectionContainersClient creates a new instance of BackupProtectionContainersClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupProtectionContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupProtectionContainersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BackupProtectionContainersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists the containers registered to Recovery Services Vault.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - BackupProtectionContainersClientListOptions contains the optional parameters for the BackupProtectionContainersClient.List
// method.
func (client *BackupProtectionContainersClient) List(vaultName string, resourceGroupName string, options *BackupProtectionContainersClientListOptions) *runtime.Pager[BackupProtectionContainersClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[BackupProtectionContainersClientListResponse]{
		More: func(page BackupProtectionContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupProtectionContainersClientListResponse) (BackupProtectionContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupProtectionContainersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupProtectionContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupProtectionContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupProtectionContainersClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupProtectionContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionContainers"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupProtectionContainersClient) listHandleResponse(resp *http.Response) (BackupProtectionContainersClientListResponse, error) {
	result := BackupProtectionContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerResourceList); err != nil {
		return BackupProtectionContainersClientListResponse{}, err
	}
	return result, nil
}
