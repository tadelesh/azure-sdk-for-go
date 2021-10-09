//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ProblemClassificationsClient contains the methods for the ProblemClassifications group.
// Don't use this type directly, use NewProblemClassificationsClient() instead.
type ProblemClassificationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewProblemClassificationsClient creates a new instance of ProblemClassificationsClient with the specified values.
func NewProblemClassificationsClient(con *arm.Connection) *ProblemClassificationsClient {
	return &ProblemClassificationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get problem classification details for a specific Azure service.
// If the operation fails it returns the *ExceptionResponse error type.
func (client *ProblemClassificationsClient) Get(ctx context.Context, serviceName string, problemClassificationName string, options *ProblemClassificationsGetOptions) (ProblemClassificationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, serviceName, problemClassificationName, options)
	if err != nil {
		return ProblemClassificationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProblemClassificationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProblemClassificationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProblemClassificationsClient) getCreateRequest(ctx context.Context, serviceName string, problemClassificationName string, options *ProblemClassificationsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services/{serviceName}/problemClassifications/{problemClassificationName}"
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if problemClassificationName == "" {
		return nil, errors.New("parameter problemClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{problemClassificationName}", url.PathEscape(problemClassificationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProblemClassificationsClient) getHandleResponse(resp *http.Response) (ProblemClassificationsGetResponse, error) {
	result := ProblemClassificationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProblemClassification); err != nil {
		return ProblemClassificationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ProblemClassificationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ExceptionResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all the problem classifications (categories) available for a specific Azure service. Always use the service and problem classifications
// obtained programmatically. This practice ensures that you
// always have the most recent set of service and problem classification Ids.
// If the operation fails it returns the *ExceptionResponse error type.
func (client *ProblemClassificationsClient) List(ctx context.Context, serviceName string, options *ProblemClassificationsListOptions) (ProblemClassificationsListResponse, error) {
	req, err := client.listCreateRequest(ctx, serviceName, options)
	if err != nil {
		return ProblemClassificationsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProblemClassificationsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProblemClassificationsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProblemClassificationsClient) listCreateRequest(ctx context.Context, serviceName string, options *ProblemClassificationsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services/{serviceName}/problemClassifications"
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProblemClassificationsClient) listHandleResponse(resp *http.Response) (ProblemClassificationsListResponse, error) {
	result := ProblemClassificationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProblemClassificationsListResult); err != nil {
		return ProblemClassificationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ProblemClassificationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ExceptionResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
