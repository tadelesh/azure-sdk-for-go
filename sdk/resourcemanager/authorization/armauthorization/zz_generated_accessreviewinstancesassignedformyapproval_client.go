//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstancesAssignedForMyApprovalClient contains the methods for the AccessReviewInstancesAssignedForMyApproval group.
// Don't use this type directly, use NewAccessReviewInstancesAssignedForMyApprovalClient() instead.
type AccessReviewInstancesAssignedForMyApprovalClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAccessReviewInstancesAssignedForMyApprovalClient creates a new instance of AccessReviewInstancesAssignedForMyApprovalClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewInstancesAssignedForMyApprovalClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AccessReviewInstancesAssignedForMyApprovalClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AccessReviewInstancesAssignedForMyApprovalClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetByID - Get single access review instance assigned for my approval.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions contains the optional parameters for the AccessReviewInstancesAssignedForMyApprovalClient.GetByID
// method.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) GetByID(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions) (AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	return result, nil
}

// List - Get access review instances assigned for my approval.
// If the operation fails it returns an *azcore.ResponseError type.
// scheduleDefinitionID - The id of the access review schedule definition.
// options - AccessReviewInstancesAssignedForMyApprovalClientListOptions contains the optional parameters for the AccessReviewInstancesAssignedForMyApprovalClient.List
// method.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) List(scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalClientListOptions) *AccessReviewInstancesAssignedForMyApprovalClientListPager {
	return &AccessReviewInstancesAssignedForMyApprovalClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scheduleDefinitionID, options)
		},
		advancer: func(ctx context.Context, resp AccessReviewInstancesAssignedForMyApprovalClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccessReviewInstanceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalClientListResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstanceListResult); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientListResponse{}, err
	}
	return result, nil
}
