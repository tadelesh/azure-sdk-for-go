//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// TopLevelDomainsClient contains the methods for the TopLevelDomains group.
// Don't use this type directly, use NewTopLevelDomainsClient() instead.
type TopLevelDomainsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTopLevelDomainsClient creates a new instance of TopLevelDomainsClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTopLevelDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TopLevelDomainsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TopLevelDomainsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Description for Get details of a top-level domain.
// If the operation fails it returns an *azcore.ResponseError type.
// name - Name of the top-level domain.
// options - TopLevelDomainsClientGetOptions contains the optional parameters for the TopLevelDomainsClient.Get method.
func (client *TopLevelDomainsClient) Get(ctx context.Context, name string, options *TopLevelDomainsClientGetOptions) (TopLevelDomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TopLevelDomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TopLevelDomainsClient) getCreateRequest(ctx context.Context, name string, options *TopLevelDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopLevelDomainsClient) getHandleResponse(resp *http.Response) (TopLevelDomainsClientGetResponse, error) {
	result := TopLevelDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelDomain); err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	return result, nil
}

// List - Description for Get all top-level domains supported for registration.
// If the operation fails it returns an *azcore.ResponseError type.
// options - TopLevelDomainsClientListOptions contains the optional parameters for the TopLevelDomainsClient.List method.
func (client *TopLevelDomainsClient) List(options *TopLevelDomainsClientListOptions) *runtime.Pager[TopLevelDomainsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TopLevelDomainsClientListResponse]{
		More: func(page TopLevelDomainsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelDomainsClientListResponse) (TopLevelDomainsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TopLevelDomainsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TopLevelDomainsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopLevelDomainsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TopLevelDomainsClient) listCreateRequest(ctx context.Context, options *TopLevelDomainsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TopLevelDomainsClient) listHandleResponse(resp *http.Response) (TopLevelDomainsClientListResponse, error) {
	result := TopLevelDomainsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelDomainCollection); err != nil {
		return TopLevelDomainsClientListResponse{}, err
	}
	return result, nil
}

// ListAgreements - Description for Gets all legal agreements that user needs to accept before purchasing a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// name - Name of the top-level domain.
// agreementOption - Domain agreement options.
// options - TopLevelDomainsClientListAgreementsOptions contains the optional parameters for the TopLevelDomainsClient.ListAgreements
// method.
func (client *TopLevelDomainsClient) ListAgreements(name string, agreementOption TopLevelDomainAgreementOption, options *TopLevelDomainsClientListAgreementsOptions) *runtime.Pager[TopLevelDomainsClientListAgreementsResponse] {
	return runtime.NewPager(runtime.PageProcessor[TopLevelDomainsClientListAgreementsResponse]{
		More: func(page TopLevelDomainsClientListAgreementsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelDomainsClientListAgreementsResponse) (TopLevelDomainsClientListAgreementsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAgreementsCreateRequest(ctx, name, agreementOption, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TopLevelDomainsClientListAgreementsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TopLevelDomainsClientListAgreementsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopLevelDomainsClientListAgreementsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAgreementsHandleResponse(resp)
		},
	})
}

// listAgreementsCreateRequest creates the ListAgreements request.
func (client *TopLevelDomainsClient) listAgreementsCreateRequest(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption, options *TopLevelDomainsClientListAgreementsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}/listAgreements"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agreementOption)
}

// listAgreementsHandleResponse handles the ListAgreements response.
func (client *TopLevelDomainsClient) listAgreementsHandleResponse(resp *http.Response) (TopLevelDomainsClientListAgreementsResponse, error) {
	result := TopLevelDomainsClientListAgreementsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TldLegalAgreementCollection); err != nil {
		return TopLevelDomainsClientListAgreementsResponse{}, err
	}
	return result, nil
}
