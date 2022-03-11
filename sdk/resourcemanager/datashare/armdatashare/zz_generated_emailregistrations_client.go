//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// EmailRegistrationsClient contains the methods for the EmailRegistrations group.
// Don't use this type directly, use NewEmailRegistrationsClient() instead.
type EmailRegistrationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewEmailRegistrationsClient creates a new instance of EmailRegistrationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEmailRegistrationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *EmailRegistrationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &EmailRegistrationsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ActivateEmail - Activate the email registration for the current tenant
// If the operation fails it returns an *azcore.ResponseError type.
// location - Location of the activation.
// emailRegistration - The payload for tenant domain activation.
// options - EmailRegistrationsClientActivateEmailOptions contains the optional parameters for the EmailRegistrationsClient.ActivateEmail
// method.
func (client *EmailRegistrationsClient) ActivateEmail(ctx context.Context, location string, emailRegistration EmailRegistration, options *EmailRegistrationsClientActivateEmailOptions) (EmailRegistrationsClientActivateEmailResponse, error) {
	req, err := client.activateEmailCreateRequest(ctx, location, emailRegistration, options)
	if err != nil {
		return EmailRegistrationsClientActivateEmailResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmailRegistrationsClientActivateEmailResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailRegistrationsClientActivateEmailResponse{}, runtime.NewResponseError(resp)
	}
	return client.activateEmailHandleResponse(resp)
}

// activateEmailCreateRequest creates the ActivateEmail request.
func (client *EmailRegistrationsClient) activateEmailCreateRequest(ctx context.Context, location string, emailRegistration EmailRegistration, options *EmailRegistrationsClientActivateEmailOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DataShare/locations/{location}/activateEmail"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, emailRegistration)
}

// activateEmailHandleResponse handles the ActivateEmail response.
func (client *EmailRegistrationsClient) activateEmailHandleResponse(resp *http.Response) (EmailRegistrationsClientActivateEmailResponse, error) {
	result := EmailRegistrationsClientActivateEmailResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailRegistration); err != nil {
		return EmailRegistrationsClientActivateEmailResponse{}, err
	}
	return result, nil
}

// RegisterEmail - Register an email for the current tenant
// If the operation fails it returns an *azcore.ResponseError type.
// location - Location of the registration
// options - EmailRegistrationsClientRegisterEmailOptions contains the optional parameters for the EmailRegistrationsClient.RegisterEmail
// method.
func (client *EmailRegistrationsClient) RegisterEmail(ctx context.Context, location string, options *EmailRegistrationsClientRegisterEmailOptions) (EmailRegistrationsClientRegisterEmailResponse, error) {
	req, err := client.registerEmailCreateRequest(ctx, location, options)
	if err != nil {
		return EmailRegistrationsClientRegisterEmailResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmailRegistrationsClientRegisterEmailResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailRegistrationsClientRegisterEmailResponse{}, runtime.NewResponseError(resp)
	}
	return client.registerEmailHandleResponse(resp)
}

// registerEmailCreateRequest creates the RegisterEmail request.
func (client *EmailRegistrationsClient) registerEmailCreateRequest(ctx context.Context, location string, options *EmailRegistrationsClientRegisterEmailOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DataShare/locations/{location}/registerEmail"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// registerEmailHandleResponse handles the RegisterEmail response.
func (client *EmailRegistrationsClient) registerEmailHandleResponse(resp *http.Response) (EmailRegistrationsClientRegisterEmailResponse, error) {
	result := EmailRegistrationsClientRegisterEmailResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailRegistration); err != nil {
		return EmailRegistrationsClientRegisterEmailResponse{}, err
	}
	return result, nil
}
