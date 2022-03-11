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
	"strings"
)

// ElasticPoolActivitiesClient contains the methods for the ElasticPoolActivities group.
// Don't use this type directly, use NewElasticPoolActivitiesClient() instead.
type ElasticPoolActivitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewElasticPoolActivitiesClient creates a new instance of ElasticPoolActivitiesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewElasticPoolActivitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ElasticPoolActivitiesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ElasticPoolActivitiesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListByElasticPool - Returns elastic pool activities.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool for which to get the current activity.
// options - ElasticPoolActivitiesClientListByElasticPoolOptions contains the optional parameters for the ElasticPoolActivitiesClient.ListByElasticPool
// method.
func (client *ElasticPoolActivitiesClient) ListByElasticPool(resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolActivitiesClientListByElasticPoolOptions) *ElasticPoolActivitiesClientListByElasticPoolPager {
	return &ElasticPoolActivitiesClientListByElasticPoolPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByElasticPoolCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
		},
	}
}

// listByElasticPoolCreateRequest creates the ListByElasticPool request.
func (client *ElasticPoolActivitiesClient) listByElasticPoolCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolActivitiesClientListByElasticPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolActivity"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByElasticPoolHandleResponse handles the ListByElasticPool response.
func (client *ElasticPoolActivitiesClient) listByElasticPoolHandleResponse(resp *http.Response) (ElasticPoolActivitiesClientListByElasticPoolResponse, error) {
	result := ElasticPoolActivitiesClientListByElasticPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolActivityListResult); err != nil {
		return ElasticPoolActivitiesClientListByElasticPoolResponse{}, err
	}
	return result, nil
}
