//go:build go1.16
// +build go1.16

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

// JobCancellationsClient contains the methods for the JobCancellations group.
// Don't use this type directly, use NewJobCancellationsClient() instead.
type JobCancellationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobCancellationsClient creates a new instance of JobCancellationsClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobCancellationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobCancellationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &JobCancellationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Trigger - Cancels a job. This is an asynchronous operation. To know the status of the cancellation, call GetCancelOperationResult
// API.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// jobName - Name of the job to cancel.
// options - JobCancellationsClientTriggerOptions contains the optional parameters for the JobCancellationsClient.Trigger
// method.
func (client *JobCancellationsClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string, options *JobCancellationsClientTriggerOptions) (JobCancellationsClientTriggerResponse, error) {
	req, err := client.triggerCreateRequest(ctx, vaultName, resourceGroupName, jobName, options)
	if err != nil {
		return JobCancellationsClientTriggerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobCancellationsClientTriggerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return JobCancellationsClientTriggerResponse{}, runtime.NewResponseError(resp)
	}
	return JobCancellationsClientTriggerResponse{}, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *JobCancellationsClient) triggerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, jobName string, options *JobCancellationsClientTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}/cancel"
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
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
