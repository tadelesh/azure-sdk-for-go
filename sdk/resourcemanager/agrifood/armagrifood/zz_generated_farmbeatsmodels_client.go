//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

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

// FarmBeatsModelsClient contains the methods for the FarmBeatsModels group.
// Don't use this type directly, use NewFarmBeatsModelsClient() instead.
type FarmBeatsModelsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFarmBeatsModelsClient creates a new instance of FarmBeatsModelsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFarmBeatsModelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FarmBeatsModelsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &FarmBeatsModelsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update FarmBeats resource.
// If the operation fails it returns an *azcore.ResponseError type.
// farmBeatsResourceName - FarmBeats resource name.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// body - FarmBeats resource create or update request object.
// options - FarmBeatsModelsClientCreateOrUpdateOptions contains the optional parameters for the FarmBeatsModelsClient.CreateOrUpdate
// method.
func (client *FarmBeatsModelsClient) CreateOrUpdate(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsClientCreateOrUpdateOptions) (FarmBeatsModelsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsModelsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FarmBeatsModelsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FarmBeatsModelsClient) createOrUpdateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FarmBeatsModelsClient) createOrUpdateHandleResponse(resp *http.Response) (FarmBeatsModelsClientCreateOrUpdateResponse, error) {
	result := FarmBeatsModelsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a FarmBeats resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// farmBeatsResourceName - FarmBeats resource name.
// options - FarmBeatsModelsClientDeleteOptions contains the optional parameters for the FarmBeatsModelsClient.Delete method.
func (client *FarmBeatsModelsClient) Delete(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsClientDeleteOptions) (FarmBeatsModelsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return FarmBeatsModelsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FarmBeatsModelsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FarmBeatsModelsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FarmBeatsModelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get FarmBeats resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// farmBeatsResourceName - FarmBeats resource name.
// options - FarmBeatsModelsClientGetOptions contains the optional parameters for the FarmBeatsModelsClient.Get method.
func (client *FarmBeatsModelsClient) Get(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsClientGetOptions) (FarmBeatsModelsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return FarmBeatsModelsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FarmBeatsModelsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FarmBeatsModelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FarmBeatsModelsClient) getHandleResponse(resp *http.Response) (FarmBeatsModelsClientGetResponse, error) {
	result := FarmBeatsModelsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists the FarmBeats instances for a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - FarmBeatsModelsClientListByResourceGroupOptions contains the optional parameters for the FarmBeatsModelsClient.ListByResourceGroup
// method.
func (client *FarmBeatsModelsClient) ListByResourceGroup(resourceGroupName string, options *FarmBeatsModelsClientListByResourceGroupOptions) *runtime.Pager[FarmBeatsModelsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[FarmBeatsModelsClientListByResourceGroupResponse]{
		More: func(page FarmBeatsModelsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FarmBeatsModelsClientListByResourceGroupResponse) (FarmBeatsModelsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FarmBeatsModelsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FarmBeatsModelsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FarmBeatsModelsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FarmBeatsModelsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FarmBeatsModelsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FarmBeatsModelsClient) listByResourceGroupHandleResponse(resp *http.Response) (FarmBeatsModelsClientListByResourceGroupResponse, error) {
	result := FarmBeatsModelsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsListResponse); err != nil {
		return FarmBeatsModelsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists the FarmBeats instances for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - FarmBeatsModelsClientListBySubscriptionOptions contains the optional parameters for the FarmBeatsModelsClient.ListBySubscription
// method.
func (client *FarmBeatsModelsClient) ListBySubscription(options *FarmBeatsModelsClientListBySubscriptionOptions) *runtime.Pager[FarmBeatsModelsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[FarmBeatsModelsClientListBySubscriptionResponse]{
		More: func(page FarmBeatsModelsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FarmBeatsModelsClientListBySubscriptionResponse) (FarmBeatsModelsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FarmBeatsModelsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FarmBeatsModelsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FarmBeatsModelsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *FarmBeatsModelsClient) listBySubscriptionCreateRequest(ctx context.Context, options *FarmBeatsModelsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *FarmBeatsModelsClient) listBySubscriptionHandleResponse(resp *http.Response) (FarmBeatsModelsClientListBySubscriptionResponse, error) {
	result := FarmBeatsModelsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsListResponse); err != nil {
		return FarmBeatsModelsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a FarmBeats resource.
// If the operation fails it returns an *azcore.ResponseError type.
// farmBeatsResourceName - FarmBeats resource name.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// body - Request object.
// options - FarmBeatsModelsClientUpdateOptions contains the optional parameters for the FarmBeatsModelsClient.Update method.
func (client *FarmBeatsModelsClient) Update(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsClientUpdateOptions) (FarmBeatsModelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsModelsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FarmBeatsModelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FarmBeatsModelsClient) updateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *FarmBeatsModelsClient) updateHandleResponse(resp *http.Response) (FarmBeatsModelsClientUpdateResponse, error) {
	result := FarmBeatsModelsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsClientUpdateResponse{}, err
	}
	return result, nil
}
