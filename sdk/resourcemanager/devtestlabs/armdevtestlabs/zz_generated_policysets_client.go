//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// PolicySetsClient contains the methods for the PolicySets group.
// Don't use this type directly, use NewPolicySetsClient() instead.
type PolicySetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPolicySetsClient creates a new instance of PolicySetsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPolicySetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PolicySetsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PolicySetsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// EvaluatePolicies - Evaluates lab policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the policy set.
// evaluatePoliciesRequest - Request body for evaluating a policy set.
// options - PolicySetsClientEvaluatePoliciesOptions contains the optional parameters for the PolicySetsClient.EvaluatePolicies
// method.
func (client *PolicySetsClient) EvaluatePolicies(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest, options *PolicySetsClientEvaluatePoliciesOptions) (PolicySetsClientEvaluatePoliciesResponse, error) {
	req, err := client.evaluatePoliciesCreateRequest(ctx, resourceGroupName, labName, name, evaluatePoliciesRequest, options)
	if err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicySetsClientEvaluatePoliciesResponse{}, runtime.NewResponseError(resp)
	}
	return client.evaluatePoliciesHandleResponse(resp)
}

// evaluatePoliciesCreateRequest creates the EvaluatePolicies request.
func (client *PolicySetsClient) evaluatePoliciesCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest, options *PolicySetsClientEvaluatePoliciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{name}/evaluatePolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, evaluatePoliciesRequest)
}

// evaluatePoliciesHandleResponse handles the EvaluatePolicies response.
func (client *PolicySetsClient) evaluatePoliciesHandleResponse(resp *http.Response) (PolicySetsClientEvaluatePoliciesResponse, error) {
	result := PolicySetsClientEvaluatePoliciesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatePoliciesResponse); err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	return result, nil
}
