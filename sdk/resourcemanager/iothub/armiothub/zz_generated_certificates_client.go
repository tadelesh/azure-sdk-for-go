//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

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

// CertificatesClient contains the methods for the Certificates group.
// Don't use this type directly, use NewCertificatesClient() instead.
type CertificatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCertificatesClient creates a new instance of CertificatesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CertificatesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CertificatesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Adds new or replaces existing certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// certificateName - The name of the certificate
// certificateDescription - The certificate body.
// options - CertificatesClientCreateOrUpdateOptions contains the optional parameters for the CertificatesClient.CreateOrUpdate
// method.
func (client *CertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, certificateDescription CertificateDescription, options *CertificatesClientCreateOrUpdateOptions) (CertificatesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, certificateName, certificateDescription, options)
	if err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CertificatesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, certificateDescription CertificateDescription, options *CertificatesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, certificateDescription)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CertificatesClient) createOrUpdateHandleResponse(resp *http.Response) (CertificatesClientCreateOrUpdateResponse, error) {
	result := CertificatesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateDescription); err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing X509 certificate or does nothing if it does not exist.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// certificateName - The name of the certificate
// ifMatch - ETag of the Certificate.
// options - CertificatesClientDeleteOptions contains the optional parameters for the CertificatesClient.Delete method.
func (client *CertificatesClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, options *CertificatesClientDeleteOptions) (CertificatesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, certificateName, ifMatch, options)
	if err != nil {
		return CertificatesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CertificatesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CertificatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, options *CertificatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GenerateVerificationCode - Generates verification code for proof of possession flow. The verification code will be used
// to generate a leaf certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// certificateName - The name of the certificate
// ifMatch - ETag of the Certificate.
// options - CertificatesClientGenerateVerificationCodeOptions contains the optional parameters for the CertificatesClient.GenerateVerificationCode
// method.
func (client *CertificatesClient) GenerateVerificationCode(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, options *CertificatesClientGenerateVerificationCodeOptions) (CertificatesClientGenerateVerificationCodeResponse, error) {
	req, err := client.generateVerificationCodeCreateRequest(ctx, resourceGroupName, resourceName, certificateName, ifMatch, options)
	if err != nil {
		return CertificatesClientGenerateVerificationCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientGenerateVerificationCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificatesClientGenerateVerificationCodeResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateVerificationCodeHandleResponse(resp)
}

// generateVerificationCodeCreateRequest creates the GenerateVerificationCode request.
func (client *CertificatesClient) generateVerificationCodeCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, options *CertificatesClientGenerateVerificationCodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}/generateVerificationCode"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateVerificationCodeHandleResponse handles the GenerateVerificationCode response.
func (client *CertificatesClient) generateVerificationCodeHandleResponse(resp *http.Response) (CertificatesClientGenerateVerificationCodeResponse, error) {
	result := CertificatesClientGenerateVerificationCodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateWithNonceDescription); err != nil {
		return CertificatesClientGenerateVerificationCodeResponse{}, err
	}
	return result, nil
}

// Get - Returns the certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// certificateName - The name of the certificate
// options - CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
func (client *CertificatesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, options *CertificatesClientGetOptions) (CertificatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, certificateName, options)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, options *CertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificatesClient) getHandleResponse(resp *http.Response) (CertificatesClientGetResponse, error) {
	result := CertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateDescription); err != nil {
		return CertificatesClientGetResponse{}, err
	}
	return result, nil
}

// ListByIotHub - Returns the list of certificates.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// options - CertificatesClientListByIotHubOptions contains the optional parameters for the CertificatesClient.ListByIotHub
// method.
func (client *CertificatesClient) ListByIotHub(ctx context.Context, resourceGroupName string, resourceName string, options *CertificatesClientListByIotHubOptions) (CertificatesClientListByIotHubResponse, error) {
	req, err := client.listByIotHubCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CertificatesClientListByIotHubResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientListByIotHubResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificatesClientListByIotHubResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByIotHubHandleResponse(resp)
}

// listByIotHubCreateRequest creates the ListByIotHub request.
func (client *CertificatesClient) listByIotHubCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CertificatesClientListByIotHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByIotHubHandleResponse handles the ListByIotHub response.
func (client *CertificatesClient) listByIotHubHandleResponse(resp *http.Response) (CertificatesClientListByIotHubResponse, error) {
	result := CertificatesClientListByIotHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListDescription); err != nil {
		return CertificatesClientListByIotHubResponse{}, err
	}
	return result, nil
}

// Verify - Verifies the certificate's private key possession by providing the leaf cert issued by the verifying pre uploaded
// certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT hub.
// resourceName - The name of the IoT hub.
// certificateName - The name of the certificate
// ifMatch - ETag of the Certificate.
// certificateVerificationBody - The name of the certificate
// options - CertificatesClientVerifyOptions contains the optional parameters for the CertificatesClient.Verify method.
func (client *CertificatesClient) Verify(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, certificateVerificationBody CertificateVerificationDescription, options *CertificatesClientVerifyOptions) (CertificatesClientVerifyResponse, error) {
	req, err := client.verifyCreateRequest(ctx, resourceGroupName, resourceName, certificateName, ifMatch, certificateVerificationBody, options)
	if err != nil {
		return CertificatesClientVerifyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificatesClientVerifyResponse{}, runtime.NewResponseError(resp)
	}
	return client.verifyHandleResponse(resp)
}

// verifyCreateRequest creates the Verify request.
func (client *CertificatesClient) verifyCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string, certificateVerificationBody CertificateVerificationDescription, options *CertificatesClientVerifyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}/verify"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, certificateVerificationBody)
}

// verifyHandleResponse handles the Verify response.
func (client *CertificatesClient) verifyHandleResponse(resp *http.Response) (CertificatesClientVerifyResponse, error) {
	result := CertificatesClientVerifyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateDescription); err != nil {
		return CertificatesClientVerifyResponse{}, err
	}
	return result, nil
}
