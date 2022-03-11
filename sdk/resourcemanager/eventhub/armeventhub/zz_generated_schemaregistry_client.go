//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// SchemaRegistryClient contains the methods for the SchemaRegistry group.
// Don't use this type directly, use NewSchemaRegistryClient() instead.
type SchemaRegistryClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSchemaRegistryClient creates a new instance of SchemaRegistryClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSchemaRegistryClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SchemaRegistryClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SchemaRegistryClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate -
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// namespaceName - The Namespace name
// schemaGroupName - The Schema Group name
// parameters - Parameters supplied to create an Event Hub resource.
// options - SchemaRegistryClientCreateOrUpdateOptions contains the optional parameters for the SchemaRegistryClient.CreateOrUpdate
// method.
func (client *SchemaRegistryClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, parameters SchemaGroup, options *SchemaRegistryClientCreateOrUpdateOptions) (SchemaRegistryClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, schemaGroupName, parameters, options)
	if err != nil {
		return SchemaRegistryClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchemaRegistryClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchemaRegistryClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SchemaRegistryClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, parameters SchemaGroup, options *SchemaRegistryClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/schemagroups/{schemaGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if schemaGroupName == "" {
		return nil, errors.New("parameter schemaGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaGroupName}", url.PathEscape(schemaGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SchemaRegistryClient) createOrUpdateHandleResponse(resp *http.Response) (SchemaRegistryClientCreateOrUpdateResponse, error) {
	result := SchemaRegistryClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemaGroup); err != nil {
		return SchemaRegistryClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete -
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// namespaceName - The Namespace name
// schemaGroupName - The Schema Group name
// options - SchemaRegistryClientDeleteOptions contains the optional parameters for the SchemaRegistryClient.Delete method.
func (client *SchemaRegistryClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *SchemaRegistryClientDeleteOptions) (SchemaRegistryClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, schemaGroupName, options)
	if err != nil {
		return SchemaRegistryClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchemaRegistryClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SchemaRegistryClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SchemaRegistryClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SchemaRegistryClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *SchemaRegistryClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/schemagroups/{schemaGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if schemaGroupName == "" {
		return nil, errors.New("parameter schemaGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaGroupName}", url.PathEscape(schemaGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// namespaceName - The Namespace name
// schemaGroupName - The Schema Group name
// options - SchemaRegistryClientGetOptions contains the optional parameters for the SchemaRegistryClient.Get method.
func (client *SchemaRegistryClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *SchemaRegistryClientGetOptions) (SchemaRegistryClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, schemaGroupName, options)
	if err != nil {
		return SchemaRegistryClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchemaRegistryClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchemaRegistryClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SchemaRegistryClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *SchemaRegistryClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/schemagroups/{schemaGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if schemaGroupName == "" {
		return nil, errors.New("parameter schemaGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaGroupName}", url.PathEscape(schemaGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SchemaRegistryClient) getHandleResponse(resp *http.Response) (SchemaRegistryClientGetResponse, error) {
	result := SchemaRegistryClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemaGroup); err != nil {
		return SchemaRegistryClientGetResponse{}, err
	}
	return result, nil
}

// ListByNamespace - Gets all the Schema Groups in a Namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// namespaceName - The Namespace name
// options - SchemaRegistryClientListByNamespaceOptions contains the optional parameters for the SchemaRegistryClient.ListByNamespace
// method.
func (client *SchemaRegistryClient) ListByNamespace(resourceGroupName string, namespaceName string, options *SchemaRegistryClientListByNamespaceOptions) *SchemaRegistryClientListByNamespacePager {
	return &SchemaRegistryClientListByNamespacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp SchemaRegistryClientListByNamespaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SchemaGroupListResult.NextLink)
		},
	}
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *SchemaRegistryClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *SchemaRegistryClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/schemagroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *SchemaRegistryClient) listByNamespaceHandleResponse(resp *http.Response) (SchemaRegistryClientListByNamespaceResponse, error) {
	result := SchemaRegistryClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemaGroupListResult); err != nil {
		return SchemaRegistryClientListByNamespaceResponse{}, err
	}
	return result, nil
}
