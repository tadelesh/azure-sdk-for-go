//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// ExperimentsClient contains the methods for the Experiments group.
// Don't use this type directly, use NewExperimentsClient() instead.
type ExperimentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExperimentsClient creates a new instance of ExperimentsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExperimentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExperimentsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExperimentsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// parameters - The Experiment resource
// options - ExperimentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.BeginCreateOrUpdate
// method.
func (client *ExperimentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (ExperimentsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, profileName, experimentName, parameters, options)
	if err != nil {
		return ExperimentsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ExperimentsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ExperimentsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ExperimentsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExperimentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExperimentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, profileName, experimentName, parameters, options)
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
func (client *ExperimentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// options - ExperimentsClientBeginDeleteOptions contains the optional parameters for the ExperimentsClient.BeginDelete method.
func (client *ExperimentsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (ExperimentsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, experimentName, options)
	if err != nil {
		return ExperimentsClientDeletePollerResponse{}, err
	}
	result := ExperimentsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("ExperimentsClient.Delete", "", resp, client.pl)
	if err != nil {
		return ExperimentsClientDeletePollerResponse{}, err
	}
	result.Poller = &ExperimentsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExperimentsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, experimentName, options)
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
func (client *ExperimentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an Experiment by ExperimentName
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// options - ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
func (client *ExperimentsClient) Get(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientGetOptions) (ExperimentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, experimentName, options)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExperimentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExperimentsClient) getHandleResponse(resp *http.Response) (ExperimentsClientGetResponse, error) {
	result := ExperimentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	return result, nil
}

// ListByProfile - Gets a list of Experiments
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// options - ExperimentsClientListByProfileOptions contains the optional parameters for the ExperimentsClient.ListByProfile
// method.
func (client *ExperimentsClient) ListByProfile(resourceGroupName string, profileName string, options *ExperimentsClientListByProfileOptions) *ExperimentsClientListByProfilePager {
	return &ExperimentsClientListByProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
		},
		advancer: func(ctx context.Context, resp ExperimentsClientListByProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExperimentList.NextLink)
		},
	}
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *ExperimentsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ExperimentsClientListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *ExperimentsClient) listByProfileHandleResponse(resp *http.Response) (ExperimentsClientListByProfileResponse, error) {
	result := ExperimentsClientListByProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentList); err != nil {
		return ExperimentsClientListByProfileResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// parameters - The Experiment Update Model
// options - ExperimentsClientBeginUpdateOptions contains the optional parameters for the ExperimentsClient.BeginUpdate method.
func (client *ExperimentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (ExperimentsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, experimentName, parameters, options)
	if err != nil {
		return ExperimentsClientUpdatePollerResponse{}, err
	}
	result := ExperimentsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("ExperimentsClient.Update", "", resp, client.pl)
	if err != nil {
		return ExperimentsClientUpdatePollerResponse{}, err
	}
	result.Poller = &ExperimentsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExperimentsClient) update(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, experimentName, parameters, options)
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
func (client *ExperimentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
