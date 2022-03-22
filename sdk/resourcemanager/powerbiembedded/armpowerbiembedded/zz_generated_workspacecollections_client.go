//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiembedded

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

// WorkspaceCollectionsClient contains the methods for the WorkspaceCollections group.
// Don't use this type directly, use NewWorkspaceCollectionsClient() instead.
type WorkspaceCollectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceCollectionsClient creates a new instance of WorkspaceCollectionsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify a Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceCollectionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkspaceCollectionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Verify the specified Power BI Workspace Collection name is valid and not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Azure location
// body - Check name availability request
// options - WorkspaceCollectionsClientCheckNameAvailabilityOptions contains the optional parameters for the WorkspaceCollectionsClient.CheckNameAvailability
// method.
func (client *WorkspaceCollectionsClient) CheckNameAvailability(ctx context.Context, location string, body CheckNameRequest, options *WorkspaceCollectionsClientCheckNameAvailabilityOptions) (WorkspaceCollectionsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return WorkspaceCollectionsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *WorkspaceCollectionsClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, body CheckNameRequest, options *WorkspaceCollectionsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *WorkspaceCollectionsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (WorkspaceCollectionsClientCheckNameAvailabilityResponse, error) {
	result := WorkspaceCollectionsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResponse); err != nil {
		return WorkspaceCollectionsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// Create - Creates a new Power BI Workspace Collection with the specified properties. A Power BI Workspace Collection contains
// one or more workspaces, and can be used to provision keys that provide API access to
// those workspaces.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// body - Create workspace collection request
// options - WorkspaceCollectionsClientCreateOptions contains the optional parameters for the WorkspaceCollectionsClient.Create
// method.
func (client *WorkspaceCollectionsClient) Create(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientCreateOptions) (WorkspaceCollectionsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceCollectionName, body, options)
	if err != nil {
		return WorkspaceCollectionsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *WorkspaceCollectionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *WorkspaceCollectionsClient) createHandleResponse(resp *http.Response) (WorkspaceCollectionsClientCreateResponse, error) {
	result := WorkspaceCollectionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollection); err != nil {
		return WorkspaceCollectionsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a Power BI Workspace Collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// options - WorkspaceCollectionsClientBeginDeleteOptions contains the optional parameters for the WorkspaceCollectionsClient.BeginDelete
// method.
func (client *WorkspaceCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientBeginDeleteOptions) (*armruntime.Poller[WorkspaceCollectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceCollectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[WorkspaceCollectionsClientDeleteResponse]("WorkspaceCollectionsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[WorkspaceCollectionsClientDeleteResponse]("WorkspaceCollectionsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Power BI Workspace Collection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspaceCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceCollectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetAccessKeys - Retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// options - WorkspaceCollectionsClientGetAccessKeysOptions contains the optional parameters for the WorkspaceCollectionsClient.GetAccessKeys
// method.
func (client *WorkspaceCollectionsClient) GetAccessKeys(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientGetAccessKeysOptions) (WorkspaceCollectionsClientGetAccessKeysResponse, error) {
	req, err := client.getAccessKeysCreateRequest(ctx, resourceGroupName, workspaceCollectionName, options)
	if err != nil {
		return WorkspaceCollectionsClientGetAccessKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientGetAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientGetAccessKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAccessKeysHandleResponse(resp)
}

// getAccessKeysCreateRequest creates the GetAccessKeys request.
func (client *WorkspaceCollectionsClient) getAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientGetAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAccessKeysHandleResponse handles the GetAccessKeys response.
func (client *WorkspaceCollectionsClient) getAccessKeysHandleResponse(resp *http.Response) (WorkspaceCollectionsClientGetAccessKeysResponse, error) {
	result := WorkspaceCollectionsClientGetAccessKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollectionAccessKeys); err != nil {
		return WorkspaceCollectionsClientGetAccessKeysResponse{}, err
	}
	return result, nil
}

// GetByName - Retrieves an existing Power BI Workspace Collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// options - WorkspaceCollectionsClientGetByNameOptions contains the optional parameters for the WorkspaceCollectionsClient.GetByName
// method.
func (client *WorkspaceCollectionsClient) GetByName(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientGetByNameOptions) (WorkspaceCollectionsClientGetByNameResponse, error) {
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, workspaceCollectionName, options)
	if err != nil {
		return WorkspaceCollectionsClientGetByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientGetByNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByNameHandleResponse(resp)
}

// getByNameCreateRequest creates the GetByName request.
func (client *WorkspaceCollectionsClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, options *WorkspaceCollectionsClientGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *WorkspaceCollectionsClient) getByNameHandleResponse(resp *http.Response) (WorkspaceCollectionsClientGetByNameResponse, error) {
	result := WorkspaceCollectionsClientGetByNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollection); err != nil {
		return WorkspaceCollectionsClientGetByNameResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieves all existing Power BI workspace collections in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// options - WorkspaceCollectionsClientListByResourceGroupOptions contains the optional parameters for the WorkspaceCollectionsClient.ListByResourceGroup
// method.
func (client *WorkspaceCollectionsClient) ListByResourceGroup(resourceGroupName string, options *WorkspaceCollectionsClientListByResourceGroupOptions) *runtime.Pager[WorkspaceCollectionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[WorkspaceCollectionsClientListByResourceGroupResponse]{
		More: func(page WorkspaceCollectionsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkspaceCollectionsClientListByResourceGroupResponse) (WorkspaceCollectionsClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return WorkspaceCollectionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceCollectionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceCollectionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkspaceCollectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkspaceCollectionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkspaceCollectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkspaceCollectionsClientListByResourceGroupResponse, error) {
	result := WorkspaceCollectionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollectionList); err != nil {
		return WorkspaceCollectionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieves all existing Power BI workspace collections in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - WorkspaceCollectionsClientListBySubscriptionOptions contains the optional parameters for the WorkspaceCollectionsClient.ListBySubscription
// method.
func (client *WorkspaceCollectionsClient) ListBySubscription(options *WorkspaceCollectionsClientListBySubscriptionOptions) *runtime.Pager[WorkspaceCollectionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[WorkspaceCollectionsClientListBySubscriptionResponse]{
		More: func(page WorkspaceCollectionsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkspaceCollectionsClientListBySubscriptionResponse) (WorkspaceCollectionsClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return WorkspaceCollectionsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceCollectionsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceCollectionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WorkspaceCollectionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *WorkspaceCollectionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WorkspaceCollectionsClient) listBySubscriptionHandleResponse(resp *http.Response) (WorkspaceCollectionsClientListBySubscriptionResponse, error) {
	result := WorkspaceCollectionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollectionList); err != nil {
		return WorkspaceCollectionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Migrate - Migrates an existing Power BI Workspace Collection to a different resource group and/or subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// body - Workspace migration request
// options - WorkspaceCollectionsClientMigrateOptions contains the optional parameters for the WorkspaceCollectionsClient.Migrate
// method.
func (client *WorkspaceCollectionsClient) Migrate(ctx context.Context, resourceGroupName string, body MigrateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientMigrateOptions) (WorkspaceCollectionsClientMigrateResponse, error) {
	req, err := client.migrateCreateRequest(ctx, resourceGroupName, body, options)
	if err != nil {
		return WorkspaceCollectionsClientMigrateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientMigrateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientMigrateResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceCollectionsClientMigrateResponse{}, nil
}

// migrateCreateRequest creates the Migrate request.
func (client *WorkspaceCollectionsClient) migrateCreateRequest(ctx context.Context, resourceGroupName string, body MigrateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientMigrateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// RegenerateKey - Regenerates the primary or secondary access key for the specified Power BI Workspace Collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// body - Access key to regenerate
// options - WorkspaceCollectionsClientRegenerateKeyOptions contains the optional parameters for the WorkspaceCollectionsClient.RegenerateKey
// method.
func (client *WorkspaceCollectionsClient) RegenerateKey(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey, options *WorkspaceCollectionsClientRegenerateKeyOptions) (WorkspaceCollectionsClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, workspaceCollectionName, body, options)
	if err != nil {
		return WorkspaceCollectionsClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *WorkspaceCollectionsClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey, options *WorkspaceCollectionsClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *WorkspaceCollectionsClient) regenerateKeyHandleResponse(resp *http.Response) (WorkspaceCollectionsClientRegenerateKeyResponse, error) {
	result := WorkspaceCollectionsClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollectionAccessKeys); err != nil {
		return WorkspaceCollectionsClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// Update - Update an existing Power BI Workspace Collection with the specified properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group
// workspaceCollectionName - Power BI Embedded Workspace Collection name
// body - Update workspace collection request
// options - WorkspaceCollectionsClientUpdateOptions contains the optional parameters for the WorkspaceCollectionsClient.Update
// method.
func (client *WorkspaceCollectionsClient) Update(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientUpdateOptions) (WorkspaceCollectionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceCollectionName, body, options)
	if err != nil {
		return WorkspaceCollectionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceCollectionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceCollectionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *WorkspaceCollectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest, options *WorkspaceCollectionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceCollectionName == "" {
		return nil, errors.New("parameter workspaceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceCollectionName}", url.PathEscape(workspaceCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-01-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *WorkspaceCollectionsClient) updateHandleResponse(resp *http.Response) (WorkspaceCollectionsClientUpdateResponse, error) {
	result := WorkspaceCollectionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceCollection); err != nil {
		return WorkspaceCollectionsClientUpdateResponse{}, err
	}
	return result, nil
}
