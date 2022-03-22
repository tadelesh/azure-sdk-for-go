//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// SourceControlSyncJobStreamsClient contains the methods for the SourceControlSyncJobStreams group.
// Don't use this type directly, use NewSourceControlSyncJobStreamsClient() instead.
type SourceControlSyncJobStreamsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSourceControlSyncJobStreamsClient creates a new instance of SourceControlSyncJobStreamsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSourceControlSyncJobStreamsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SourceControlSyncJobStreamsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SourceControlSyncJobStreamsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Retrieve a sync job stream identified by stream id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The source control name.
// sourceControlSyncJobID - The source control sync job id.
// streamID - The id of the sync job stream.
// options - SourceControlSyncJobStreamsClientGetOptions contains the optional parameters for the SourceControlSyncJobStreamsClient.Get
// method.
func (client *SourceControlSyncJobStreamsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, streamID string, options *SourceControlSyncJobStreamsClientGetOptions) (SourceControlSyncJobStreamsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, streamID, options)
	if err != nil {
		return SourceControlSyncJobStreamsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlSyncJobStreamsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SourceControlSyncJobStreamsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SourceControlSyncJobStreamsClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, streamID string, options *SourceControlSyncJobStreamsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}/streams/{streamId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlSyncJobId}", url.PathEscape(sourceControlSyncJobID))
	if streamID == "" {
		return nil, errors.New("parameter streamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{streamId}", url.PathEscape(streamID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SourceControlSyncJobStreamsClient) getHandleResponse(resp *http.Response) (SourceControlSyncJobStreamsClientGetResponse, error) {
	result := SourceControlSyncJobStreamsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlSyncJobStreamByID); err != nil {
		return SourceControlSyncJobStreamsClientGetResponse{}, err
	}
	return result, nil
}

// ListBySyncJob - Retrieve a list of sync job streams identified by sync job id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// sourceControlName - The source control name.
// sourceControlSyncJobID - The source control sync job id.
// options - SourceControlSyncJobStreamsClientListBySyncJobOptions contains the optional parameters for the SourceControlSyncJobStreamsClient.ListBySyncJob
// method.
func (client *SourceControlSyncJobStreamsClient) ListBySyncJob(resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, options *SourceControlSyncJobStreamsClientListBySyncJobOptions) *runtime.Pager[SourceControlSyncJobStreamsClientListBySyncJobResponse] {
	return runtime.NewPager(runtime.PageProcessor[SourceControlSyncJobStreamsClientListBySyncJobResponse]{
		More: func(page SourceControlSyncJobStreamsClientListBySyncJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlSyncJobStreamsClientListBySyncJobResponse) (SourceControlSyncJobStreamsClientListBySyncJobResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySyncJobCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SourceControlSyncJobStreamsClientListBySyncJobResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SourceControlSyncJobStreamsClientListBySyncJobResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SourceControlSyncJobStreamsClientListBySyncJobResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySyncJobHandleResponse(resp)
		},
	})
}

// listBySyncJobCreateRequest creates the ListBySyncJob request.
func (client *SourceControlSyncJobStreamsClient) listBySyncJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, options *SourceControlSyncJobStreamsClientListBySyncJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}/streams"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlSyncJobId}", url.PathEscape(sourceControlSyncJobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySyncJobHandleResponse handles the ListBySyncJob response.
func (client *SourceControlSyncJobStreamsClient) listBySyncJobHandleResponse(resp *http.Response) (SourceControlSyncJobStreamsClientListBySyncJobResponse, error) {
	result := SourceControlSyncJobStreamsClientListBySyncJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlSyncJobStreamsListBySyncJob); err != nil {
		return SourceControlSyncJobStreamsClientListBySyncJobResponse{}, err
	}
	return result, nil
}
