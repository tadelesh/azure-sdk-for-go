//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
func NewJobsClient(con *arm.Connection, subscriptionID string) *JobsClient {
	return &JobsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BookShipmentPickUp - Book shipment pick up.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) BookShipmentPickUp(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest ShipmentPickUpRequest, options *JobsBookShipmentPickUpOptions) (JobsBookShipmentPickUpResponse, error) {
	req, err := client.bookShipmentPickUpCreateRequest(ctx, resourceGroupName, jobName, shipmentPickUpRequest, options)
	if err != nil {
		return JobsBookShipmentPickUpResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsBookShipmentPickUpResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsBookShipmentPickUpResponse{}, client.bookShipmentPickUpHandleError(resp)
	}
	return client.bookShipmentPickUpHandleResponse(resp)
}

// bookShipmentPickUpCreateRequest creates the BookShipmentPickUp request.
func (client *JobsClient) bookShipmentPickUpCreateRequest(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest ShipmentPickUpRequest, options *JobsBookShipmentPickUpOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/bookShipmentPickUp"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, shipmentPickUpRequest)
}

// bookShipmentPickUpHandleResponse handles the BookShipmentPickUp response.
func (client *JobsClient) bookShipmentPickUpHandleResponse(resp *http.Response) (JobsBookShipmentPickUpResponse, error) {
	result := JobsBookShipmentPickUpResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShipmentPickUpResponse); err != nil {
		return JobsBookShipmentPickUpResponse{}, err
	}
	return result, nil
}

// bookShipmentPickUpHandleError handles the BookShipmentPickUp error response.
func (client *JobsClient) bookShipmentPickUpHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Cancel - CancelJob.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) Cancel(ctx context.Context, resourceGroupName string, jobName string, cancellationReason CancellationReason, options *JobsCancelOptions) (JobsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, jobName, cancellationReason, options)
	if err != nil {
		return JobsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return JobsCancelResponse{}, client.cancelHandleError(resp)
	}
	return JobsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *JobsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, jobName string, cancellationReason CancellationReason, options *JobsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, cancellationReason)
}

// cancelHandleError handles the Cancel error response.
func (client *JobsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreate - Creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update
// job API.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) BeginCreate(ctx context.Context, resourceGroupName string, jobName string, jobResource JobResource, options *JobsBeginCreateOptions) (JobsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, jobName, jobResource, options)
	if err != nil {
		return JobsCreatePollerResponse{}, err
	}
	result := JobsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return JobsCreatePollerResponse{}, err
	}
	result.Poller = &JobsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update
// job API.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) create(ctx context.Context, resourceGroupName string, jobName string, jobResource JobResource, options *JobsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, jobName, jobResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *JobsClient) createCreateRequest(ctx context.Context, resourceGroupName string, jobName string, jobResource JobResource, options *JobsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, jobResource)
}

// createHandleError handles the Create error response.
func (client *JobsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) BeginDelete(ctx context.Context, resourceGroupName string, jobName string, options *JobsBeginDeleteOptions) (JobsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return JobsDeletePollerResponse{}, err
	}
	result := JobsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return JobsDeletePollerResponse{}, err
	}
	result.Poller = &JobsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) deleteOperation(ctx context.Context, resourceGroupName string, jobName string, options *JobsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *JobsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *JobsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets information about the specified job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) Get(ctx context.Context, resourceGroupName string, jobName string, options *JobsGetOptions) (JobsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return JobsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *JobsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsGetResponse, error) {
	result := JobsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResource); err != nil {
		return JobsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all the jobs available under the subscription.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) List(options *JobsListOptions) *JobsListPager {
	return &JobsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp JobsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobResourceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *JobsClient) listCreateRequest(ctx context.Context, options *JobsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JobsClient) listHandleResponse(resp *http.Response) (JobsListResponse, error) {
	result := JobsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResourceList); err != nil {
		return JobsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *JobsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all the jobs available under the given resource group.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) ListByResourceGroup(resourceGroupName string, options *JobsListByResourceGroupOptions) *JobsListByResourceGroupPager {
	return &JobsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp JobsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobResourceList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *JobsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *JobsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *JobsClient) listByResourceGroupHandleResponse(resp *http.Response) (JobsListByResourceGroupResponse, error) {
	result := JobsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResourceList); err != nil {
		return JobsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *JobsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListCredentials - This method gets the unencrypted secrets related to the job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) ListCredentials(ctx context.Context, resourceGroupName string, jobName string, options *JobsListCredentialsOptions) (JobsListCredentialsResponse, error) {
	req, err := client.listCredentialsCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return JobsListCredentialsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsListCredentialsResponse{}, client.listCredentialsHandleError(resp)
	}
	return client.listCredentialsHandleResponse(resp)
}

// listCredentialsCreateRequest creates the ListCredentials request.
func (client *JobsClient) listCredentialsCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *JobsListCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/listCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCredentialsHandleResponse handles the ListCredentials response.
func (client *JobsClient) listCredentialsHandleResponse(resp *http.Response) (JobsListCredentialsResponse, error) {
	result := JobsListCredentialsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnencryptedCredentialsList); err != nil {
		return JobsListCredentialsResponse{}, err
	}
	return result, nil
}

// listCredentialsHandleError handles the ListCredentials error response.
func (client *JobsClient) listCredentialsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// MarkDevicesShipped - Request to mark devices for a given job as shipped
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) MarkDevicesShipped(ctx context.Context, jobName string, resourceGroupName string, markDevicesShippedRequest MarkDevicesShippedRequest, options *JobsMarkDevicesShippedOptions) (JobsMarkDevicesShippedResponse, error) {
	req, err := client.markDevicesShippedCreateRequest(ctx, jobName, resourceGroupName, markDevicesShippedRequest, options)
	if err != nil {
		return JobsMarkDevicesShippedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsMarkDevicesShippedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return JobsMarkDevicesShippedResponse{}, client.markDevicesShippedHandleError(resp)
	}
	return JobsMarkDevicesShippedResponse{RawResponse: resp}, nil
}

// markDevicesShippedCreateRequest creates the MarkDevicesShipped request.
func (client *JobsClient) markDevicesShippedCreateRequest(ctx context.Context, jobName string, resourceGroupName string, markDevicesShippedRequest MarkDevicesShippedRequest, options *JobsMarkDevicesShippedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/markDevicesShipped"
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, markDevicesShippedRequest)
}

// markDevicesShippedHandleError handles the MarkDevicesShipped error response.
func (client *JobsClient) markDevicesShippedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates the properties of an existing job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) BeginUpdate(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter JobResourceUpdateParameter, options *JobsBeginUpdateOptions) (JobsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, jobName, jobResourceUpdateParameter, options)
	if err != nil {
		return JobsUpdatePollerResponse{}, err
	}
	result := JobsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return JobsUpdatePollerResponse{}, err
	}
	result.Poller = &JobsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the properties of an existing job.
// If the operation fails it returns the *APIError error type.
func (client *JobsClient) update(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter JobResourceUpdateParameter, options *JobsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, jobResourceUpdateParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *JobsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter JobResourceUpdateParameter, options *JobsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, jobResourceUpdateParameter)
}

// updateHandleError handles the Update error response.
func (client *JobsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := APIError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
