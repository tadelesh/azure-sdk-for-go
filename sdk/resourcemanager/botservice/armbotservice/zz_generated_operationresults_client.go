//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

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

// OperationResultsClient contains the methods for the OperationResults group.
// Don't use this type directly, use NewOperationResultsClient() instead.
type OperationResultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationResultsClient creates a new instance of OperationResultsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationResultsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OperationResultsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginGet - Get the operation result for a long running operation.
// If the operation fails it returns an *azcore.ResponseError type.
// operationResultID - The ID of the operation result to get.
// options - OperationResultsClientBeginGetOptions contains the optional parameters for the OperationResultsClient.BeginGet
// method.
func (client *OperationResultsClient) BeginGet(ctx context.Context, operationResultID string, options *OperationResultsClientBeginGetOptions) (*armruntime.Poller[OperationResultsClientGetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.get(ctx, operationResultID, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OperationResultsClientGetResponse]("OperationResultsClient.Get", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OperationResultsClientGetResponse]("OperationResultsClient.Get", options.ResumeToken, client.pl, nil)
	}
}

// Get - Get the operation result for a long running operation.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OperationResultsClient) get(ctx context.Context, operationResultID string, options *OperationResultsClientBeginGetOptions) (*http.Response, error) {
	req, err := client.getCreateRequest(ctx, operationResultID, options)
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

// getCreateRequest creates the Get request.
func (client *OperationResultsClient) getCreateRequest(ctx context.Context, operationResultID string, options *OperationResultsClientBeginGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/operationresults/{operationResultId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if operationResultID == "" {
		return nil, errors.New("parameter operationResultID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationResultId}", url.PathEscape(operationResultID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
