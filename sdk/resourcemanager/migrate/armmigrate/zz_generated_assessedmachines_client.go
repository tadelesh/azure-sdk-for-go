//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// AssessedMachinesClient contains the methods for the AssessedMachines group.
// Don't use this type directly, use NewAssessedMachinesClient() instead.
type AssessedMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAssessedMachinesClient creates a new instance of AssessedMachinesClient with the specified values.
// subscriptionID - Azure Subscription Id in which project was created.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssessedMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AssessedMachinesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AssessedMachinesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get an assessed machine with its size & cost estimate that was evaluated in the specified assessment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// groupName - Unique name of a group within a project.
// assessmentName - Unique name of an assessment within a project.
// assessedMachineName - Unique name of an assessed machine evaluated as part of an assessment.
// options - AssessedMachinesClientGetOptions contains the optional parameters for the AssessedMachinesClient.Get method.
func (client *AssessedMachinesClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string, options *AssessedMachinesClientGetOptions) (AssessedMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedMachineName, options)
	if err != nil {
		return AssessedMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessedMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessedMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssessedMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string, options *AssessedMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assessedMachineName == "" {
		return nil, errors.New("parameter assessedMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedMachineName}", url.PathEscape(assessedMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessedMachinesClient) getHandleResponse(resp *http.Response) (AssessedMachinesClientGetResponse, error) {
	result := AssessedMachinesClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedMachine); err != nil {
		return AssessedMachinesClientGetResponse{}, err
	}
	return result, nil
}

// ListByAssessment - Get list of machines that assessed as part of the specified assessment. Returns a json array of objects
// of type 'assessedMachine' as specified in the Models section.
// Whenever an assessment is created or updated, it goes under computation. During this phase, the 'status' field of Assessment
// object reports 'Computing'. During the period when the assessment is under
// computation, the list of assessed machines is empty and no assessed machines are returned by this call.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// groupName - Unique name of a group within a project.
// assessmentName - Unique name of an assessment within a project.
// options - AssessedMachinesClientListByAssessmentOptions contains the optional parameters for the AssessedMachinesClient.ListByAssessment
// method.
func (client *AssessedMachinesClient) ListByAssessment(resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedMachinesClientListByAssessmentOptions) *AssessedMachinesClientListByAssessmentPager {
	return &AssessedMachinesClientListByAssessmentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAssessmentCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
		},
		advancer: func(ctx context.Context, resp AssessedMachinesClientListByAssessmentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssessedMachineResultList.NextLink)
		},
	}
}

// listByAssessmentCreateRequest creates the ListByAssessment request.
func (client *AssessedMachinesClient) listByAssessmentCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedMachinesClientListByAssessmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAssessmentHandleResponse handles the ListByAssessment response.
func (client *AssessedMachinesClient) listByAssessmentHandleResponse(resp *http.Response) (AssessedMachinesClientListByAssessmentResponse, error) {
	result := AssessedMachinesClientListByAssessmentResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedMachineResultList); err != nil {
		return AssessedMachinesClientListByAssessmentResponse{}, err
	}
	return result, nil
}
