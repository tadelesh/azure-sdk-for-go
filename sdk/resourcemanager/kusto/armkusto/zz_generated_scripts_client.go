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

// ScriptsClient contains the methods for the Scripts group.
// Don't use this type directly, use NewScriptsClient() instead.
type ScriptsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewScriptsClient creates a new instance of ScriptsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewScriptsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ScriptsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ScriptsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckNameAvailability - Checks that the script name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// scriptName - The name of the script.
// options - ScriptsClientCheckNameAvailabilityOptions contains the optional parameters for the ScriptsClient.CheckNameAvailability
// method.
func (client *ScriptsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName ScriptCheckNameRequest, options *ScriptsClientCheckNameAvailabilityOptions) (ScriptsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, databaseName, scriptName, options)
	if err != nil {
		return ScriptsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ScriptsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName ScriptCheckNameRequest, options *ScriptsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scriptsCheckNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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
	return req, runtime.MarshalAsJSON(req, scriptName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ScriptsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ScriptsClientCheckNameAvailabilityResponse, error) {
	result := ScriptsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return ScriptsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates a Kusto database script.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// scriptName - The name of the Kusto database script.
// parameters - The Kusto Script parameters contains the KQL to run.
// options - ScriptsClientBeginCreateOrUpdateOptions contains the optional parameters for the ScriptsClient.BeginCreateOrUpdate
// method.
func (client *ScriptsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ScriptsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, databaseName, scriptName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ScriptsClientCreateOrUpdateResponse]("ScriptsClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ScriptsClientCreateOrUpdateResponse]("ScriptsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a Kusto database script.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ScriptsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, scriptName, parameters, options)
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
func (client *ScriptsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
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

// BeginDelete - Deletes a Kusto principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// scriptName - The name of the Kusto database script.
// options - ScriptsClientBeginDeleteOptions contains the optional parameters for the ScriptsClient.BeginDelete method.
func (client *ScriptsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, options *ScriptsClientBeginDeleteOptions) (*armruntime.Poller[ScriptsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, scriptName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ScriptsClientDeleteResponse]("ScriptsClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ScriptsClientDeleteResponse]("ScriptsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Kusto principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ScriptsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, options *ScriptsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, scriptName, options)
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
func (client *ScriptsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, options *ScriptsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
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

// Get - Gets a Kusto cluster database script.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// scriptName - The name of the Kusto database script.
// options - ScriptsClientGetOptions contains the optional parameters for the ScriptsClient.Get method.
func (client *ScriptsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, options *ScriptsClientGetOptions) (ScriptsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, scriptName, options)
	if err != nil {
		return ScriptsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScriptsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, options *ScriptsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
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
func (client *ScriptsClient) getHandleResponse(resp *http.Response) (ScriptsClientGetResponse, error) {
	result := ScriptsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Script); err != nil {
		return ScriptsClientGetResponse{}, err
	}
	return result, nil
}

// ListByDatabase - Returns the list of database scripts for given database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// options - ScriptsClientListByDatabaseOptions contains the optional parameters for the ScriptsClient.ListByDatabase method.
func (client *ScriptsClient) ListByDatabase(resourceGroupName string, clusterName string, databaseName string, options *ScriptsClientListByDatabaseOptions) *runtime.Pager[ScriptsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PageProcessor[ScriptsClientListByDatabaseResponse]{
		More: func(page ScriptsClientListByDatabaseResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ScriptsClientListByDatabaseResponse) (ScriptsClientListByDatabaseResponse, error) {
			req, err := client.listByDatabaseCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
			if err != nil {
				return ScriptsClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ScriptsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScriptsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ScriptsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *ScriptsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scripts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ScriptsClient) listByDatabaseHandleResponse(resp *http.Response) (ScriptsClientListByDatabaseResponse, error) {
	result := ScriptsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptListResult); err != nil {
		return ScriptsClientListByDatabaseResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a database script.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// scriptName - The name of the Kusto database script.
// parameters - The Kusto Script parameters contains to the KQL to run.
// options - ScriptsClientBeginUpdateOptions contains the optional parameters for the ScriptsClient.BeginUpdate method.
func (client *ScriptsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginUpdateOptions) (*armruntime.Poller[ScriptsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, scriptName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ScriptsClientUpdateResponse]("ScriptsClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ScriptsClientUpdateResponse]("ScriptsClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a database script.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ScriptsClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, scriptName, parameters, options)
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
func (client *ScriptsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, scriptName string, parameters Script, options *ScriptsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/scripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
