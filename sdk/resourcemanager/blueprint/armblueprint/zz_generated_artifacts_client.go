//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// ArtifactsClient contains the methods for the Artifacts group.
// Don't use this type directly, use NewArtifactsClient() instead.
type ArtifactsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewArtifactsClient creates a new instance of ArtifactsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewArtifactsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ArtifactsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ArtifactsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update blueprint artifact.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// artifactName - Name of the blueprint artifact.
// artifact - Blueprint artifact to create or update.
// options - ArtifactsClientCreateOrUpdateOptions contains the optional parameters for the ArtifactsClient.CreateOrUpdate
// method.
func (client *ArtifactsClient) CreateOrUpdate(ctx context.Context, resourceScope string, blueprintName string, artifactName string, artifact ArtifactClassification, options *ArtifactsClientCreateOrUpdateOptions) (ArtifactsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceScope, blueprintName, artifactName, artifact, options)
	if err != nil {
		return ArtifactsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ArtifactsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ArtifactsClient) createOrUpdateCreateRequest(ctx context.Context, resourceScope string, blueprintName string, artifactName string, artifact ArtifactClassification, options *ArtifactsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if artifactName == "" {
		return nil, errors.New("parameter artifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactName}", url.PathEscape(artifactName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, artifact)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ArtifactsClient) createOrUpdateHandleResponse(resp *http.Response) (ArtifactsClientCreateOrUpdateResponse, error) {
	result := ArtifactsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ArtifactsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a blueprint artifact.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// artifactName - Name of the blueprint artifact.
// options - ArtifactsClientDeleteOptions contains the optional parameters for the ArtifactsClient.Delete method.
func (client *ArtifactsClient) Delete(ctx context.Context, resourceScope string, blueprintName string, artifactName string, options *ArtifactsClientDeleteOptions) (ArtifactsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceScope, blueprintName, artifactName, options)
	if err != nil {
		return ArtifactsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ArtifactsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ArtifactsClient) deleteCreateRequest(ctx context.Context, resourceScope string, blueprintName string, artifactName string, options *ArtifactsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if artifactName == "" {
		return nil, errors.New("parameter artifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactName}", url.PathEscape(artifactName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ArtifactsClient) deleteHandleResponse(resp *http.Response) (ArtifactsClientDeleteResponse, error) {
	result := ArtifactsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ArtifactsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get a blueprint artifact.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// artifactName - Name of the blueprint artifact.
// options - ArtifactsClientGetOptions contains the optional parameters for the ArtifactsClient.Get method.
func (client *ArtifactsClient) Get(ctx context.Context, resourceScope string, blueprintName string, artifactName string, options *ArtifactsClientGetOptions) (ArtifactsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceScope, blueprintName, artifactName, options)
	if err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArtifactsClient) getCreateRequest(ctx context.Context, resourceScope string, blueprintName string, artifactName string, options *ArtifactsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if artifactName == "" {
		return nil, errors.New("parameter artifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactName}", url.PathEscape(artifactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArtifactsClient) getHandleResponse(resp *http.Response) (ArtifactsClientGetResponse, error) {
	result := ArtifactsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	return result, nil
}

// List - List artifacts for a given blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// options - ArtifactsClientListOptions contains the optional parameters for the ArtifactsClient.List method.
func (client *ArtifactsClient) List(resourceScope string, blueprintName string, options *ArtifactsClientListOptions) *runtime.Pager[ArtifactsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ArtifactsClientListResponse]{
		More: func(page ArtifactsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArtifactsClientListResponse) (ArtifactsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceScope, blueprintName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArtifactsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ArtifactsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArtifactsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ArtifactsClient) listCreateRequest(ctx context.Context, resourceScope string, blueprintName string, options *ArtifactsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArtifactsClient) listHandleResponse(resp *http.Response) (ArtifactsClientListResponse, error) {
	result := ArtifactsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactList); err != nil {
		return ArtifactsClientListResponse{}, err
	}
	return result, nil
}
