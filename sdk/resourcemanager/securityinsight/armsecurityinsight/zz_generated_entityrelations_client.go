//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsight

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

// EntityRelationsClient contains the methods for the EntityRelations group.
// Don't use this type directly, use NewEntityRelationsClient() instead.
type EntityRelationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEntityRelationsClient creates a new instance of EntityRelationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEntityRelationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EntityRelationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &EntityRelationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetRelation - Gets an entity relation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// entityID - entity ID
// relationName - Relation Name
// options - EntityRelationsClientGetRelationOptions contains the optional parameters for the EntityRelationsClient.GetRelation
// method.
func (client *EntityRelationsClient) GetRelation(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string, options *EntityRelationsClientGetRelationOptions) (EntityRelationsClientGetRelationResponse, error) {
	req, err := client.getRelationCreateRequest(ctx, resourceGroupName, workspaceName, entityID, relationName, options)
	if err != nil {
		return EntityRelationsClientGetRelationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EntityRelationsClientGetRelationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EntityRelationsClientGetRelationResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRelationHandleResponse(resp)
}

// getRelationCreateRequest creates the GetRelation request.
func (client *EntityRelationsClient) getRelationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string, options *EntityRelationsClientGetRelationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getRelationHandleResponse handles the GetRelation response.
func (client *EntityRelationsClient) getRelationHandleResponse(resp *http.Response) (EntityRelationsClientGetRelationResponse, error) {
	result := EntityRelationsClientGetRelationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Relation); err != nil {
		return EntityRelationsClientGetRelationResponse{}, err
	}
	return result, nil
}
