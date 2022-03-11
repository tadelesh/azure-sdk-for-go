//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// InstructionsClient contains the methods for the Instructions group.
// Don't use this type directly, use NewInstructionsClient() instead.
type InstructionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewInstructionsClient creates a new instance of InstructionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInstructionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *InstructionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &InstructionsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get the instruction by name. These are custom billing instructions and are only applicable for certain customers.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// instructionName - Instruction Name.
// options - InstructionsClientGetOptions contains the optional parameters for the InstructionsClient.Get method.
func (client *InstructionsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string, options *InstructionsClientGetOptions) (InstructionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, instructionName, options)
	if err != nil {
		return InstructionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstructionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstructionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InstructionsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string, options *InstructionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/{instructionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if instructionName == "" {
		return nil, errors.New("parameter instructionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instructionName}", url.PathEscape(instructionName))
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
func (client *InstructionsClient) getHandleResponse(resp *http.Response) (InstructionsClientGetResponse, error) {
	result := InstructionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instruction); err != nil {
		return InstructionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByBillingProfile - Lists the instructions by billing profile id.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// options - InstructionsClientListByBillingProfileOptions contains the optional parameters for the InstructionsClient.ListByBillingProfile
// method.
func (client *InstructionsClient) ListByBillingProfile(billingAccountName string, billingProfileName string, options *InstructionsClientListByBillingProfileOptions) *InstructionsClientListByBillingProfilePager {
	return &InstructionsClientListByBillingProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
		},
		advancer: func(ctx context.Context, resp InstructionsClientListByBillingProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InstructionListResult.NextLink)
		},
	}
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *InstructionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *InstructionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
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

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *InstructionsClient) listByBillingProfileHandleResponse(resp *http.Response) (InstructionsClientListByBillingProfileResponse, error) {
	result := InstructionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstructionListResult); err != nil {
		return InstructionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// Put - Creates or updates an instruction. These are custom billing instructions and are only applicable for certain customers.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// instructionName - Instruction Name.
// parameters - The new instruction.
// options - InstructionsClientPutOptions contains the optional parameters for the InstructionsClient.Put method.
func (client *InstructionsClient) Put(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string, parameters Instruction, options *InstructionsClientPutOptions) (InstructionsClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, billingAccountName, billingProfileName, instructionName, parameters, options)
	if err != nil {
		return InstructionsClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstructionsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstructionsClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *InstructionsClient) putCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string, parameters Instruction, options *InstructionsClientPutOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/{instructionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if instructionName == "" {
		return nil, errors.New("parameter instructionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instructionName}", url.PathEscape(instructionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// putHandleResponse handles the Put response.
func (client *InstructionsClient) putHandleResponse(resp *http.Response) (InstructionsClientPutResponse, error) {
	result := InstructionsClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instruction); err != nil {
		return InstructionsClientPutResponse{}, err
	}
	return result, nil
}
