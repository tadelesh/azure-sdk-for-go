//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// RecommendedActionsClient contains the methods for the RecommendedActions group.
// Don't use this type directly, use NewRecommendedActionsClient() instead.
type RecommendedActionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecommendedActionsClient creates a new instance of RecommendedActionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecommendedActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RecommendedActionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RecommendedActionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Retrieve recommended actions from the advisor.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// advisorName - The advisor name for recommendation action.
// recommendedActionName - The recommended action name.
// options - RecommendedActionsClientGetOptions contains the optional parameters for the RecommendedActionsClient.Get method.
func (client *RecommendedActionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string, recommendedActionName string, options *RecommendedActionsClientGetOptions) (RecommendedActionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, advisorName, recommendedActionName, options)
	if err != nil {
		return RecommendedActionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendedActionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendedActionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecommendedActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, recommendedActionName string, options *RecommendedActionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}"
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
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if recommendedActionName == "" {
		return nil, errors.New("parameter recommendedActionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendedActionName}", url.PathEscape(recommendedActionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecommendedActionsClient) getHandleResponse(resp *http.Response) (RecommendedActionsClientGetResponse, error) {
	result := RecommendedActionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationAction); err != nil {
		return RecommendedActionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Retrieve recommended actions from the advisor.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// advisorName - The advisor name for recommendation action.
// options - RecommendedActionsClientListByServerOptions contains the optional parameters for the RecommendedActionsClient.ListByServer
// method.
func (client *RecommendedActionsClient) ListByServer(resourceGroupName string, serverName string, advisorName string, options *RecommendedActionsClientListByServerOptions) *runtime.Pager[RecommendedActionsClientListByServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[RecommendedActionsClientListByServerResponse]{
		More: func(page RecommendedActionsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecommendedActionsClientListByServerResponse) (RecommendedActionsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, advisorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RecommendedActionsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RecommendedActionsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RecommendedActionsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *RecommendedActionsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *RecommendedActionsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors/{advisorName}/recommendedActions"
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
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	if options != nil && options.SessionID != nil {
		reqQP.Set("sessionId", *options.SessionID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *RecommendedActionsClient) listByServerHandleResponse(resp *http.Response) (RecommendedActionsClientListByServerResponse, error) {
	result := RecommendedActionsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationActionsResultList); err != nil {
		return RecommendedActionsClientListByServerResponse{}, err
	}
	return result, nil
}
