//go:build go1.18
// +build go1.18

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

// ElasticPoolsClient contains the methods for the ElasticPools group.
// Don't use this type directly, use NewElasticPoolsClient() instead.
type ElasticPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewElasticPoolsClient creates a new instance of ElasticPoolsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewElasticPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ElasticPoolsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ElasticPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// parameters - The elastic pool parameters.
// options - ElasticPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the ElasticPoolsClient.BeginCreateOrUpdate
// method.
func (client *ElasticPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ElasticPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ElasticPoolsClientCreateOrUpdateResponse]("ElasticPoolsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ElasticPoolsClientCreateOrUpdateResponse]("ElasticPoolsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ElasticPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientBeginDeleteOptions contains the optional parameters for the ElasticPoolsClient.BeginDelete
// method.
func (client *ElasticPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (*armruntime.Poller[ElasticPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, elasticPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ElasticPoolsClientDeleteResponse]("ElasticPoolsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ElasticPoolsClientDeleteResponse]("ElasticPoolsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ElasticPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Failovers an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool to failover.
// options - ElasticPoolsClientBeginFailoverOptions contains the optional parameters for the ElasticPoolsClient.BeginFailover
// method.
func (client *ElasticPoolsClient) BeginFailover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (*armruntime.Poller[ElasticPoolsClientFailoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.failover(ctx, resourceGroupName, serverName, elasticPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ElasticPoolsClientFailoverResponse]("ElasticPoolsClient.Failover", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ElasticPoolsClientFailoverResponse]("ElasticPoolsClient.Failover", options.ResumeToken, client.pl, nil)
	}
}

// Failover - Failovers an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) failover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
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

// failoverCreateRequest creates the Failover request.
func (client *ElasticPoolsClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/failover"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientGetOptions contains the optional parameters for the ElasticPoolsClient.Get method.
func (client *ElasticPoolsClient) Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientGetOptions) (ElasticPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ElasticPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ElasticPoolsClient) getHandleResponse(resp *http.Response) (ElasticPoolsClientGetResponse, error) {
	result := ElasticPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPool); err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets all elastic pools in a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ElasticPoolsClientListByServerOptions contains the optional parameters for the ElasticPoolsClient.ListByServer
// method.
func (client *ElasticPoolsClient) ListByServer(resourceGroupName string, serverName string, options *ElasticPoolsClientListByServerOptions) *runtime.Pager[ElasticPoolsClientListByServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[ElasticPoolsClientListByServerResponse]{
		More: func(page ElasticPoolsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ElasticPoolsClientListByServerResponse) (ElasticPoolsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ElasticPoolsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticPoolsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticPoolsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ElasticPoolsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ElasticPoolsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(*options.Skip, 10))
	}
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ElasticPoolsClient) listByServerHandleResponse(resp *http.Response) (ElasticPoolsClientListByServerResponse, error) {
	result := ElasticPoolsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolListResult); err != nil {
		return ElasticPoolsClientListByServerResponse{}, err
	}
	return result, nil
}

// ListMetricDefinitions - Returns elastic pool metric definitions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientListMetricDefinitionsOptions contains the optional parameters for the ElasticPoolsClient.ListMetricDefinitions
// method.
func (client *ElasticPoolsClient) ListMetricDefinitions(resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientListMetricDefinitionsOptions) *runtime.Pager[ElasticPoolsClientListMetricDefinitionsResponse] {
	return runtime.NewPager(runtime.PageProcessor[ElasticPoolsClientListMetricDefinitionsResponse]{
		More: func(page ElasticPoolsClientListMetricDefinitionsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ElasticPoolsClientListMetricDefinitionsResponse) (ElasticPoolsClientListMetricDefinitionsResponse, error) {
			req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
			if err != nil {
				return ElasticPoolsClientListMetricDefinitionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticPoolsClientListMetricDefinitionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticPoolsClientListMetricDefinitionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricDefinitionsHandleResponse(resp)
		},
	})
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *ElasticPoolsClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metricDefinitions"
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

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *ElasticPoolsClient) listMetricDefinitionsHandleResponse(resp *http.Response) (ElasticPoolsClientListMetricDefinitionsResponse, error) {
	result := ElasticPoolsClientListMetricDefinitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionListResult); err != nil {
		return ElasticPoolsClientListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// ListMetrics - Returns elastic pool metrics.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// filter - An OData filter expression that describes a subset of metrics to return.
// options - ElasticPoolsClientListMetricsOptions contains the optional parameters for the ElasticPoolsClient.ListMetrics
// method.
func (client *ElasticPoolsClient) ListMetrics(resourceGroupName string, serverName string, elasticPoolName string, filter string, options *ElasticPoolsClientListMetricsOptions) *runtime.Pager[ElasticPoolsClientListMetricsResponse] {
	return runtime.NewPager(runtime.PageProcessor[ElasticPoolsClientListMetricsResponse]{
		More: func(page ElasticPoolsClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ElasticPoolsClientListMetricsResponse) (ElasticPoolsClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, filter, options)
			if err != nil {
				return ElasticPoolsClientListMetricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticPoolsClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticPoolsClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *ElasticPoolsClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string, options *ElasticPoolsClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metrics"
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
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *ElasticPoolsClient) listMetricsHandleResponse(resp *http.Response) (ElasticPoolsClientListMetricsResponse, error) {
	result := ElasticPoolsClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return ElasticPoolsClientListMetricsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// parameters - The elastic pool update parameters.
// options - ElasticPoolsClientBeginUpdateOptions contains the optional parameters for the ElasticPoolsClient.BeginUpdate
// method.
func (client *ElasticPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (*armruntime.Poller[ElasticPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ElasticPoolsClientUpdateResponse]("ElasticPoolsClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ElasticPoolsClientUpdateResponse]("ElasticPoolsClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ElasticPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
