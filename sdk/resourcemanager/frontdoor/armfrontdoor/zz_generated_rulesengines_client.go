//go:build go1.18
// +build go1.18

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

// RulesEnginesClient contains the methods for the RulesEngines group.
// Don't use this type directly, use NewRulesEnginesClient() instead.
type RulesEnginesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRulesEnginesClient creates a new instance of RulesEnginesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRulesEnginesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RulesEnginesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RulesEnginesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// rulesEngineName - Name of the Rules Engine which is unique within the Front Door.
// rulesEngineParameters - Rules Engine Configuration properties needed to create a new Rules Engine Configuration.
// options - RulesEnginesClientBeginCreateOrUpdateOptions contains the optional parameters for the RulesEnginesClient.BeginCreateOrUpdate
// method.
func (client *RulesEnginesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[RulesEnginesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, frontDoorName, rulesEngineName, rulesEngineParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[RulesEnginesClientCreateOrUpdateResponse]("RulesEnginesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[RulesEnginesClientCreateOrUpdateResponse]("RulesEnginesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RulesEnginesClient) createOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, rulesEngineParameters, options)
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
func (client *RulesEnginesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine, options *RulesEnginesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, rulesEngineParameters)
}

// BeginDelete - Deletes an existing Rules Engine Configuration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// rulesEngineName - Name of the Rules Engine which is unique within the Front Door.
// options - RulesEnginesClientBeginDeleteOptions contains the optional parameters for the RulesEnginesClient.BeginDelete
// method.
func (client *RulesEnginesClient) BeginDelete(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesClientBeginDeleteOptions) (*armruntime.Poller[RulesEnginesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[RulesEnginesClientDeleteResponse]("RulesEnginesClient.Delete", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[RulesEnginesClientDeleteResponse]("RulesEnginesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an existing Rules Engine Configuration with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RulesEnginesClient) deleteOperation(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RulesEnginesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a Rules Engine Configuration with the specified name within the specified Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// rulesEngineName - Name of the Rules Engine which is unique within the Front Door.
// options - RulesEnginesClientGetOptions contains the optional parameters for the RulesEnginesClient.Get method.
func (client *RulesEnginesClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesClientGetOptions) (RulesEnginesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, frontDoorName, rulesEngineName, options)
	if err != nil {
		return RulesEnginesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesEnginesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RulesEnginesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RulesEnginesClient) getCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, options *RulesEnginesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	if rulesEngineName == "" {
		return nil, errors.New("parameter rulesEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rulesEngineName}", url.PathEscape(rulesEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RulesEnginesClient) getHandleResponse(resp *http.Response) (RulesEnginesClientGetResponse, error) {
	result := RulesEnginesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RulesEngine); err != nil {
		return RulesEnginesClientGetResponse{}, err
	}
	return result, nil
}

// ListByFrontDoor - Lists all of the Rules Engine Configurations within a Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// options - RulesEnginesClientListByFrontDoorOptions contains the optional parameters for the RulesEnginesClient.ListByFrontDoor
// method.
func (client *RulesEnginesClient) ListByFrontDoor(resourceGroupName string, frontDoorName string, options *RulesEnginesClientListByFrontDoorOptions) *runtime.Pager[RulesEnginesClientListByFrontDoorResponse] {
	return runtime.NewPager(runtime.PageProcessor[RulesEnginesClientListByFrontDoorResponse]{
		More: func(page RulesEnginesClientListByFrontDoorResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RulesEnginesClientListByFrontDoorResponse) (RulesEnginesClientListByFrontDoorResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFrontDoorCreateRequest(ctx, resourceGroupName, frontDoorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RulesEnginesClientListByFrontDoorResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RulesEnginesClientListByFrontDoorResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RulesEnginesClientListByFrontDoorResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFrontDoorHandleResponse(resp)
		},
	})
}

// listByFrontDoorCreateRequest creates the ListByFrontDoor request.
func (client *RulesEnginesClient) listByFrontDoorCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *RulesEnginesClientListByFrontDoorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFrontDoorHandleResponse handles the ListByFrontDoor response.
func (client *RulesEnginesClient) listByFrontDoorHandleResponse(resp *http.Response) (RulesEnginesClientListByFrontDoorResponse, error) {
	result := RulesEnginesClientListByFrontDoorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RulesEngineListResult); err != nil {
		return RulesEnginesClientListByFrontDoorResponse{}, err
	}
	return result, nil
}
