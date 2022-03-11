//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armelastic

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

// VMIngestionClient contains the methods for the VMIngestion group.
// Don't use this type directly, use NewVMIngestionClient() instead.
type VMIngestionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVMIngestionClient creates a new instance of VMIngestionClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVMIngestionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VMIngestionClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VMIngestionClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Details - List the vm ingestion details that will be monitored by the Elastic monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// options - VMIngestionClientDetailsOptions contains the optional parameters for the VMIngestionClient.Details method.
func (client *VMIngestionClient) Details(ctx context.Context, resourceGroupName string, monitorName string, options *VMIngestionClientDetailsOptions) (VMIngestionClientDetailsResponse, error) {
	req, err := client.detailsCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return VMIngestionClientDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMIngestionClientDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VMIngestionClientDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.detailsHandleResponse(resp)
}

// detailsCreateRequest creates the Details request.
func (client *VMIngestionClient) detailsCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *VMIngestionClientDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/vmIngestionDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// detailsHandleResponse handles the Details response.
func (client *VMIngestionClient) detailsHandleResponse(resp *http.Response) (VMIngestionClientDetailsResponse, error) {
	result := VMIngestionClientDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMIngestionDetailsResponse); err != nil {
		return VMIngestionClientDetailsResponse{}, err
	}
	return result, nil
}
