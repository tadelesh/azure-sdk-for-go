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

// SecurityPINsClient contains the methods for the SecurityPINs group.
// Don't use this type directly, use NewSecurityPINsClient() instead.
type SecurityPINsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityPINsClient creates a new instance of SecurityPINsClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityPINsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecurityPINsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SecurityPINsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get the security PIN.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - SecurityPINsClientGetOptions contains the optional parameters for the SecurityPINsClient.Get method.
func (client *SecurityPINsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, options *SecurityPINsClientGetOptions) (SecurityPINsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, options)
	if err != nil {
		return SecurityPINsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityPINsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityPINsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityPINsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *SecurityPINsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupSecurityPIN"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityPINsClient) getHandleResponse(resp *http.Response) (SecurityPINsClientGetResponse, error) {
	result := SecurityPINsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TokenInformation); err != nil {
		return SecurityPINsClientGetResponse{}, err
	}
	return result, nil
}
