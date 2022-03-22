//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// ManagementClient contains the methods for the HybridComputeManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagementClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginUpgradeExtensions - The operation to Upgrade Machine Extensions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// machineName - The name of the hybrid machine.
// extensionUpgradeParameters - Parameters supplied to the Upgrade Extensions operation.
// options - ManagementClientBeginUpgradeExtensionsOptions contains the optional parameters for the ManagementClient.BeginUpgradeExtensions
// method.
func (client *ManagementClient) BeginUpgradeExtensions(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*armruntime.Poller[ManagementClientUpgradeExtensionsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgradeExtensions(ctx, resourceGroupName, machineName, extensionUpgradeParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ManagementClientUpgradeExtensionsResponse]("ManagementClient.UpgradeExtensions", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ManagementClientUpgradeExtensionsResponse]("ManagementClient.UpgradeExtensions", options.ResumeToken, client.pl, nil)
	}
}

// UpgradeExtensions - The operation to Upgrade Machine Extensions.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagementClient) upgradeExtensions(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*http.Response, error) {
	req, err := client.upgradeExtensionsCreateRequest(ctx, resourceGroupName, machineName, extensionUpgradeParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeExtensionsCreateRequest creates the UpgradeExtensions request.
func (client *ManagementClient) upgradeExtensionsCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/upgradeExtensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionUpgradeParameters)
}
