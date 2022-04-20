//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CloudServiceRoleInstancesClient contains the methods for the CloudServiceRoleInstances group.
// Don't use this type directly, use NewCloudServiceRoleInstancesClient() instead.
type CloudServiceRoleInstancesClient struct {
	host string
	subscriptionID string
	pl runtime.Pipeline
}

// NewCloudServiceRoleInstancesClient creates a new instance of CloudServiceRoleInstancesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudServiceRoleInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudServiceRoleInstancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CloudServiceRoleInstancesClient{
		subscriptionID: subscriptionID,
		host: ep,
pl: pl,
	}
	return client, nil
}

// BeginDelete - Deletes a role instance from a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientBeginDeleteOptions contains the optional parameters for the CloudServiceRoleInstancesClient.BeginDelete
// method.
func (client *CloudServiceRoleInstancesClient) BeginDelete(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginDeleteOptions) (*armruntime.Poller[CloudServiceRoleInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CloudServiceRoleInstancesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CloudServiceRoleInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a role instance from a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CloudServiceRoleInstancesClient) deleteOperation(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
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
func (client *CloudServiceRoleInstancesClient) deleteCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a role instance from a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientGetOptions contains the optional parameters for the CloudServiceRoleInstancesClient.Get
// method.
func (client *CloudServiceRoleInstancesClient) Get(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetOptions) (CloudServiceRoleInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return CloudServiceRoleInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceRoleInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceRoleInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudServiceRoleInstancesClient) getCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudServiceRoleInstancesClient) getHandleResponse(resp *http.Response) (CloudServiceRoleInstancesClientGetResponse, error) {
	result := CloudServiceRoleInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleInstance); err != nil {
		return CloudServiceRoleInstancesClientGetResponse{}, err
	}
	return result, nil
}

// GetInstanceView - Retrieves information about the run-time state of a role instance in a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientGetInstanceViewOptions contains the optional parameters for the CloudServiceRoleInstancesClient.GetInstanceView
// method.
func (client *CloudServiceRoleInstancesClient) GetInstanceView(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetInstanceViewOptions) (CloudServiceRoleInstancesClientGetInstanceViewResponse, error) {
	req, err := client.getInstanceViewCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return CloudServiceRoleInstancesClientGetInstanceViewResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceRoleInstancesClientGetInstanceViewResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceRoleInstancesClientGetInstanceViewResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInstanceViewHandleResponse(resp)
}

// getInstanceViewCreateRequest creates the GetInstanceView request.
func (client *CloudServiceRoleInstancesClient) getInstanceViewCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetInstanceViewOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/instanceView"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInstanceViewHandleResponse handles the GetInstanceView response.
func (client *CloudServiceRoleInstancesClient) getInstanceViewHandleResponse(resp *http.Response) (CloudServiceRoleInstancesClientGetInstanceViewResponse, error) {
	result := CloudServiceRoleInstancesClientGetInstanceViewResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleInstanceView); err != nil {
		return CloudServiceRoleInstancesClientGetInstanceViewResponse{}, err
	}
	return result, nil
}

// GetRemoteDesktopFile - Gets a remote desktop file for a role instance in a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientGetRemoteDesktopFileOptions contains the optional parameters for the CloudServiceRoleInstancesClient.GetRemoteDesktopFile
// method.
func (client *CloudServiceRoleInstancesClient) GetRemoteDesktopFile(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetRemoteDesktopFileOptions) (CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse, error) {
	req, err := client.getRemoteDesktopFileCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse{}, runtime.NewResponseError(resp)
	}
	return CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse{Body: resp.Body}, nil
}

// getRemoteDesktopFileCreateRequest creates the GetRemoteDesktopFile request.
func (client *CloudServiceRoleInstancesClient) getRemoteDesktopFileCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientGetRemoteDesktopFileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/remoteDesktopFile"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/x-rdp")
	return req, nil
}

// NewListPager - Gets the list of all role instances in a cloud service. Use nextLink property in the response to get the
// next page of role instances. Do this till nextLink is null to fetch all the role instances.
// If the operation fails it returns an *azcore.ResponseError type.
// options - CloudServiceRoleInstancesClientListOptions contains the optional parameters for the CloudServiceRoleInstancesClient.List
// method.
func (client *CloudServiceRoleInstancesClient) NewListPager(resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientListOptions) (*runtime.Pager[CloudServiceRoleInstancesClientListResponse]) {
	return runtime.NewPager(runtime.PageProcessor[CloudServiceRoleInstancesClientListResponse]{
		More: func(page CloudServiceRoleInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudServiceRoleInstancesClientListResponse) (CloudServiceRoleInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CloudServiceRoleInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CloudServiceRoleInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CloudServiceRoleInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CloudServiceRoleInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudServiceRoleInstancesClient) listHandleResponse(resp *http.Response) (CloudServiceRoleInstancesClientListResponse, error) {
	result := CloudServiceRoleInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleInstanceListResult); err != nil {
		return CloudServiceRoleInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginRebuild - The Rebuild Role Instance asynchronous operation reinstalls the operating system on instances of web roles
// or worker roles and initializes the storage resources that are used by them. If you do not
// want to initialize storage resources, you can use Reimage Role Instance.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientBeginRebuildOptions contains the optional parameters for the CloudServiceRoleInstancesClient.BeginRebuild
// method.
func (client *CloudServiceRoleInstancesClient) BeginRebuild(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRebuildOptions) (*armruntime.Poller[CloudServiceRoleInstancesClientRebuildResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.rebuild(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CloudServiceRoleInstancesClientRebuildResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CloudServiceRoleInstancesClientRebuildResponse](options.ResumeToken, client.pl, nil)
	}
}

// Rebuild - The Rebuild Role Instance asynchronous operation reinstalls the operating system on instances of web roles or
// worker roles and initializes the storage resources that are used by them. If you do not
// want to initialize storage resources, you can use Reimage Role Instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CloudServiceRoleInstancesClient) rebuild(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRebuildOptions) (*http.Response, error) {
	req, err := client.rebuildCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
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

// rebuildCreateRequest creates the Rebuild request.
func (client *CloudServiceRoleInstancesClient) rebuildCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRebuildOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/rebuild"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginReimage - The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles
// or worker roles.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientBeginReimageOptions contains the optional parameters for the CloudServiceRoleInstancesClient.BeginReimage
// method.
func (client *CloudServiceRoleInstancesClient) BeginReimage(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginReimageOptions) (*armruntime.Poller[CloudServiceRoleInstancesClientReimageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reimage(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CloudServiceRoleInstancesClientReimageResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CloudServiceRoleInstancesClientReimageResponse](options.ResumeToken, client.pl, nil)
	}
}

// Reimage - The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles or
// worker roles.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CloudServiceRoleInstancesClient) reimage(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginReimageOptions) (*http.Response, error) {
	req, err := client.reimageCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
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

// reimageCreateRequest creates the Reimage request.
func (client *CloudServiceRoleInstancesClient) reimageCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginReimageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/reimage"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginRestart - The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// roleInstanceName - Name of the role instance.
// options - CloudServiceRoleInstancesClientBeginRestartOptions contains the optional parameters for the CloudServiceRoleInstancesClient.BeginRestart
// method.
func (client *CloudServiceRoleInstancesClient) BeginRestart(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRestartOptions) (*armruntime.Poller[CloudServiceRoleInstancesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[CloudServiceRoleInstancesClientRestartResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[CloudServiceRoleInstancesClientRestartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Restart - The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CloudServiceRoleInstancesClient) restart(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
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

// restartCreateRequest creates the Restart request.
func (client *CloudServiceRoleInstancesClient) restartCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/restart"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

