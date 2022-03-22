//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

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

// AttachedDatabaseConfigurationsClient contains the methods for the AttachedDatabaseConfigurations group.
// Don't use this type directly, use NewAttachedDatabaseConfigurationsClient() instead.
type AttachedDatabaseConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAttachedDatabaseConfigurationsClient creates a new instance of AttachedDatabaseConfigurationsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAttachedDatabaseConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AttachedDatabaseConfigurationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AttachedDatabaseConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Checks that the attached database configuration resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// resourceName - The name of the resource.
// options - AttachedDatabaseConfigurationsClientCheckNameAvailabilityOptions contains the optional parameters for the AttachedDatabaseConfigurationsClient.CheckNameAvailability
// method.
func (client *AttachedDatabaseConfigurationsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName AttachedDatabaseConfigurationsCheckNameRequest, options *AttachedDatabaseConfigurationsClientCheckNameAvailabilityOptions) (AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *AttachedDatabaseConfigurationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName AttachedDatabaseConfigurationsCheckNameRequest, options *AttachedDatabaseConfigurationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurationCheckNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *AttachedDatabaseConfigurationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse, error) {
	result := AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return AttachedDatabaseConfigurationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates an attached database configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// attachedDatabaseConfigurationName - The name of the attached database configuration.
// parameters - The database parameters supplied to the CreateOrUpdate operation.
// options - AttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the AttachedDatabaseConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *AttachedDatabaseConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, parameters AttachedDatabaseConfiguration, options *AttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[AttachedDatabaseConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, attachedDatabaseConfigurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AttachedDatabaseConfigurationsClientCreateOrUpdateResponse]("AttachedDatabaseConfigurationsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AttachedDatabaseConfigurationsClientCreateOrUpdateResponse]("AttachedDatabaseConfigurationsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an attached database configuration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AttachedDatabaseConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, parameters AttachedDatabaseConfiguration, options *AttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, attachedDatabaseConfigurationName, parameters, options)
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
func (client *AttachedDatabaseConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, parameters AttachedDatabaseConfiguration, options *AttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if attachedDatabaseConfigurationName == "" {
		return nil, errors.New("parameter attachedDatabaseConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDatabaseConfigurationName}", url.PathEscape(attachedDatabaseConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the attached database configuration with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// attachedDatabaseConfigurationName - The name of the attached database configuration.
// options - AttachedDatabaseConfigurationsClientBeginDeleteOptions contains the optional parameters for the AttachedDatabaseConfigurationsClient.BeginDelete
// method.
func (client *AttachedDatabaseConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, options *AttachedDatabaseConfigurationsClientBeginDeleteOptions) (*armruntime.Poller[AttachedDatabaseConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, attachedDatabaseConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AttachedDatabaseConfigurationsClientDeleteResponse]("AttachedDatabaseConfigurationsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AttachedDatabaseConfigurationsClientDeleteResponse]("AttachedDatabaseConfigurationsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the attached database configuration with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AttachedDatabaseConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, options *AttachedDatabaseConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, attachedDatabaseConfigurationName, options)
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
func (client *AttachedDatabaseConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, options *AttachedDatabaseConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if attachedDatabaseConfigurationName == "" {
		return nil, errors.New("parameter attachedDatabaseConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDatabaseConfigurationName}", url.PathEscape(attachedDatabaseConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns an attached database configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// attachedDatabaseConfigurationName - The name of the attached database configuration.
// options - AttachedDatabaseConfigurationsClientGetOptions contains the optional parameters for the AttachedDatabaseConfigurationsClient.Get
// method.
func (client *AttachedDatabaseConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, options *AttachedDatabaseConfigurationsClientGetOptions) (AttachedDatabaseConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, attachedDatabaseConfigurationName, options)
	if err != nil {
		return AttachedDatabaseConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedDatabaseConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedDatabaseConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AttachedDatabaseConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, options *AttachedDatabaseConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if attachedDatabaseConfigurationName == "" {
		return nil, errors.New("parameter attachedDatabaseConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDatabaseConfigurationName}", url.PathEscape(attachedDatabaseConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AttachedDatabaseConfigurationsClient) getHandleResponse(resp *http.Response) (AttachedDatabaseConfigurationsClientGetResponse, error) {
	result := AttachedDatabaseConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDatabaseConfiguration); err != nil {
		return AttachedDatabaseConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// ListByCluster - Returns the list of attached database configurations of the given Kusto cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// options - AttachedDatabaseConfigurationsClientListByClusterOptions contains the optional parameters for the AttachedDatabaseConfigurationsClient.ListByCluster
// method.
func (client *AttachedDatabaseConfigurationsClient) ListByCluster(resourceGroupName string, clusterName string, options *AttachedDatabaseConfigurationsClientListByClusterOptions) *runtime.Pager[AttachedDatabaseConfigurationsClientListByClusterResponse] {
	return runtime.NewPager(runtime.PageProcessor[AttachedDatabaseConfigurationsClientListByClusterResponse]{
		More: func(page AttachedDatabaseConfigurationsClientListByClusterResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AttachedDatabaseConfigurationsClientListByClusterResponse) (AttachedDatabaseConfigurationsClientListByClusterResponse, error) {
			req, err := client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			if err != nil {
				return AttachedDatabaseConfigurationsClientListByClusterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttachedDatabaseConfigurationsClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttachedDatabaseConfigurationsClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *AttachedDatabaseConfigurationsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *AttachedDatabaseConfigurationsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *AttachedDatabaseConfigurationsClient) listByClusterHandleResponse(resp *http.Response) (AttachedDatabaseConfigurationsClientListByClusterResponse, error) {
	result := AttachedDatabaseConfigurationsClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDatabaseConfigurationListResult); err != nil {
		return AttachedDatabaseConfigurationsClientListByClusterResponse{}, err
	}
	return result, nil
}
