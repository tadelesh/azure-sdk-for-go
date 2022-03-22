//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

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

// WorkspacesClient contains the methods for the Workspaces group.
// Don't use this type directly, use NewWorkspacesClient() instead.
type WorkspacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspacesClient creates a new instance of WorkspacesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspacesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &WorkspacesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a workspace resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// workspaceName - The name of the quantum workspace resource.
// quantumWorkspace - Workspace details.
// options - WorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.BeginCreateOrUpdate
// method.
func (client *WorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, quantumWorkspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[WorkspacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, quantumWorkspace, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[WorkspacesClientCreateOrUpdateResponse]("WorkspacesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[WorkspacesClientCreateOrUpdateResponse]("WorkspacesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a workspace resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, quantumWorkspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, quantumWorkspace, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, quantumWorkspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Quantum/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, quantumWorkspace)
}

// BeginDelete - Deletes a Workspace resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// workspaceName - The name of the quantum workspace resource.
// options - WorkspacesClientBeginDeleteOptions contains the optional parameters for the WorkspacesClient.BeginDelete method.
func (client *WorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*armruntime.Poller[WorkspacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[WorkspacesClientDeleteResponse]("WorkspacesClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[WorkspacesClientDeleteResponse]("WorkspacesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Workspace resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, options)
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
func (client *WorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Quantum/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the Workspace resource associated with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// workspaceName - The name of the quantum workspace resource.
// options - WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
func (client *WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (WorkspacesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Quantum/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspacesClient) getHandleResponse(resp *http.Response) (WorkspacesClientGetResponse, error) {
	result := WorkspacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets the list of Workspaces within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.ListByResourceGroup
// method.
func (client *WorkspacesClient) ListByResourceGroup(resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) *runtime.Pager[WorkspacesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[WorkspacesClientListByResourceGroupResponse]{
		More: func(page WorkspacesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListByResourceGroupResponse) (WorkspacesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspacesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspacesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspacesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Quantum/workspaces"
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
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkspacesClientListByResourceGroupResponse, error) {
	result := WorkspacesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return WorkspacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets the list of Workspaces within a Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - WorkspacesClientListBySubscriptionOptions contains the optional parameters for the WorkspacesClient.ListBySubscription
// method.
func (client *WorkspacesClient) ListBySubscription(options *WorkspacesClientListBySubscriptionOptions) *runtime.Pager[WorkspacesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[WorkspacesClientListBySubscriptionResponse]{
		More: func(page WorkspacesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListBySubscriptionResponse) (WorkspacesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspacesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspacesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspacesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WorkspacesClient) listBySubscriptionCreateRequest(ctx context.Context, options *WorkspacesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Quantum/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WorkspacesClient) listBySubscriptionHandleResponse(resp *http.Response) (WorkspacesClientListBySubscriptionResponse, error) {
	result := WorkspacesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return WorkspacesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an existing workspace's tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// workspaceName - The name of the quantum workspace resource.
// workspaceTags - Parameters supplied to update tags.
// options - WorkspacesClientUpdateTagsOptions contains the optional parameters for the WorkspacesClient.UpdateTags method.
func (client *WorkspacesClient) UpdateTags(ctx context.Context, resourceGroupName string, workspaceName string, workspaceTags TagsObject, options *WorkspacesClientUpdateTagsOptions) (WorkspacesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, workspaceName, workspaceTags, options)
	if err != nil {
		return WorkspacesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *WorkspacesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceTags TagsObject, options *WorkspacesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Quantum/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspaceTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *WorkspacesClient) updateTagsHandleResponse(resp *http.Response) (WorkspacesClientUpdateTagsResponse, error) {
	result := WorkspacesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
