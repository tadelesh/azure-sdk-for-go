//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityandcompliance

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

// PrivateLinkServicesForSCCPowershellClient contains the methods for the PrivateLinkServicesForSCCPowershell group.
// Don't use this type directly, use NewPrivateLinkServicesForSCCPowershellClient() instead.
type PrivateLinkServicesForSCCPowershellClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkServicesForSCCPowershellClient creates a new instance of PrivateLinkServicesForSCCPowershellClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkServicesForSCCPowershellClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkServicesForSCCPowershellClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PrivateLinkServicesForSCCPowershellClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateLinkServicesForSCCPowershellDescription - The service instance metadata.
// options - PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions contains the optional parameters for the
// PrivateLinkServicesForSCCPowershellClient.BeginCreateOrUpdate method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*armruntime.Poller[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateLinkServicesForSCCPowershellDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse]("PrivateLinkServicesForSCCPowershellClient.CreateOrUpdate", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse]("PrivateLinkServicesForSCCPowershellClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForSCCPowershellClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateLinkServicesForSCCPowershellDescription, options)
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
func (client *PrivateLinkServicesForSCCPowershellClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateLinkServicesForSCCPowershellDescription)
}

// BeginDelete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.BeginDelete
// method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*armruntime.Poller[PrivateLinkServicesForSCCPowershellClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateLinkServicesForSCCPowershellClientDeleteResponse]("PrivateLinkServicesForSCCPowershellClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForSCCPowershellClientDeleteResponse]("PrivateLinkServicesForSCCPowershellClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForSCCPowershellClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
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
func (client *PrivateLinkServicesForSCCPowershellClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the metadata of a privateLinkServicesForSCCPowershell resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForSCCPowershellClientGetOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.Get
// method.
func (client *PrivateLinkServicesForSCCPowershellClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientGetOptions) (PrivateLinkServicesForSCCPowershellClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServicesForSCCPowershellClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkServicesForSCCPowershellClient) getHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientGetResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescription); err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all the privateLinkServicesForSCCPowershell instances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateLinkServicesForSCCPowershellClientListOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.List
// method.
func (client *PrivateLinkServicesForSCCPowershellClient) List(options *PrivateLinkServicesForSCCPowershellClientListOptions) *runtime.Pager[PrivateLinkServicesForSCCPowershellClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateLinkServicesForSCCPowershellClientListResponse]{
		More: func(page PrivateLinkServicesForSCCPowershellClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForSCCPowershellClientListResponse) (PrivateLinkServicesForSCCPowershellClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForSCCPowershellClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkServicesForSCCPowershellClient) listCreateRequest(ctx context.Context, options *PrivateLinkServicesForSCCPowershellClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkServicesForSCCPowershellClient) listHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientListResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescriptionListResult); err != nil {
		return PrivateLinkServicesForSCCPowershellClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all the service instances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// options - PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions contains the optional parameters for the
// PrivateLinkServicesForSCCPowershellClient.ListByResourceGroup method.
func (client *PrivateLinkServicesForSCCPowershellClient) ListByResourceGroup(resourceGroupName string, options *PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse]{
		More: func(page PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse) (PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkServicesForSCCPowershellClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell"
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
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkServicesForSCCPowershellClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescriptionListResult); err != nil {
		return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// servicePatchDescription - The service instance metadata and security metadata.
// options - PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.BeginUpdate
// method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*armruntime.Poller[PrivateLinkServicesForSCCPowershellClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateLinkServicesForSCCPowershellClientUpdateResponse]("PrivateLinkServicesForSCCPowershellClient.Update", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForSCCPowershellClientUpdateResponse]("PrivateLinkServicesForSCCPowershellClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForSCCPowershellClient) update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
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
func (client *PrivateLinkServicesForSCCPowershellClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, servicePatchDescription)
}