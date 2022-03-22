//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

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

// ReportsClient contains the methods for the Reports group.
// Don't use this type directly, use NewReportsClient() instead.
type ReportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewReportsClient creates a new instance of ReportsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReportsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ReportsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get information about a report associated with a configuration profile assignment run
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - The configuration profile assignment name.
// reportName - The report name.
// vmName - The name of the virtual machine.
// options - ReportsClientGetOptions contains the optional parameters for the ReportsClient.Get method.
func (client *ReportsClient) Get(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, reportName string, vmName string, options *ReportsClientGetOptions) (ReportsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, reportName, vmName, options)
	if err != nil {
		return ReportsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReportsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, reportName string, vmName string, options *ReportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports/{reportName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReportsClient) getHandleResponse(resp *http.Response) (ReportsClientGetResponse, error) {
	result := ReportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Report); err != nil {
		return ReportsClientGetResponse{}, err
	}
	return result, nil
}

// ListByConfigurationProfileAssignments - Retrieve a list of reports within a given configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - The configuration profile assignment name.
// vmName - The name of the virtual machine.
// options - ReportsClientListByConfigurationProfileAssignmentsOptions contains the optional parameters for the ReportsClient.ListByConfigurationProfileAssignments
// method.
func (client *ReportsClient) ListByConfigurationProfileAssignments(resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ReportsClientListByConfigurationProfileAssignmentsOptions) *runtime.Pager[ReportsClientListByConfigurationProfileAssignmentsResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByConfigurationProfileAssignmentsResponse]{
		More: func(page ReportsClientListByConfigurationProfileAssignmentsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByConfigurationProfileAssignmentsResponse) (ReportsClientListByConfigurationProfileAssignmentsResponse, error) {
			req, err := client.listByConfigurationProfileAssignmentsCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, vmName, options)
			if err != nil {
				return ReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByConfigurationProfileAssignmentsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationProfileAssignmentsHandleResponse(resp)
		},
	})
}

// listByConfigurationProfileAssignmentsCreateRequest creates the ListByConfigurationProfileAssignments request.
func (client *ReportsClient) listByConfigurationProfileAssignmentsCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ReportsClientListByConfigurationProfileAssignmentsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByConfigurationProfileAssignmentsHandleResponse handles the ListByConfigurationProfileAssignments response.
func (client *ReportsClient) listByConfigurationProfileAssignmentsHandleResponse(resp *http.Response) (ReportsClientListByConfigurationProfileAssignmentsResponse, error) {
	result := ReportsClientListByConfigurationProfileAssignmentsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportList); err != nil {
		return ReportsClientListByConfigurationProfileAssignmentsResponse{}, err
	}
	return result, nil
}
