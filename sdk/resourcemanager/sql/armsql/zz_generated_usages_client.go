//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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
	"strconv"
	"strings"
)

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UsagesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListByInstancePool - Gets all instance pool usage metrics
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// instancePoolName - The name of the instance pool to be retrieved.
// options - UsagesClientListByInstancePoolOptions contains the optional parameters for the UsagesClient.ListByInstancePool
// method.
func (client *UsagesClient) ListByInstancePool(resourceGroupName string, instancePoolName string, options *UsagesClientListByInstancePoolOptions) *UsagesClientListByInstancePoolPager {
	return &UsagesClientListByInstancePoolPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInstancePoolCreateRequest(ctx, resourceGroupName, instancePoolName, options)
		},
		advancer: func(ctx context.Context, resp UsagesClientListByInstancePoolResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UsageListResult.NextLink)
		},
	}
}

// listByInstancePoolCreateRequest creates the ListByInstancePool request.
func (client *UsagesClient) listByInstancePoolCreateRequest(ctx context.Context, resourceGroupName string, instancePoolName string, options *UsagesClientListByInstancePoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instancePoolName == "" {
		return nil, errors.New("parameter instancePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instancePoolName}", url.PathEscape(instancePoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ExpandChildren != nil {
		reqQP.Set("expandChildren", strconv.FormatBool(*options.ExpandChildren))
	}
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstancePoolHandleResponse handles the ListByInstancePool response.
func (client *UsagesClient) listByInstancePoolHandleResponse(resp *http.Response) (UsagesClientListByInstancePoolResponse, error) {
	result := UsagesClientListByInstancePoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return UsagesClientListByInstancePoolResponse{}, err
	}
	return result, nil
}
