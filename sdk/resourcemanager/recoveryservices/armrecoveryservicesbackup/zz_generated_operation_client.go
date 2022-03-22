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

// OperationClient contains the methods for the Operation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Validate - Validate operation for specified backed up item. This is a synchronous operation.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - resource validate operation request
// options - OperationClientValidateOptions contains the optional parameters for the OperationClient.Validate method.
func (client *OperationClient) Validate(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateOperationRequestClassification, options *OperationClientValidateOptions) (OperationClientValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return OperationClientValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationClientValidateResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *OperationClient) validateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateOperationRequestClassification, options *OperationClientValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupValidateOperation"
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
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateHandleResponse handles the Validate response.
func (client *OperationClient) validateHandleResponse(resp *http.Response) (OperationClientValidateResponse, error) {
	result := OperationClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateOperationsResponse); err != nil {
		return OperationClientValidateResponse{}, err
	}
	return result, nil
}
