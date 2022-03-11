//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// ManagedRuleSetsClient contains the methods for the ManagedRuleSets group.
// Don't use this type directly, use NewManagedRuleSetsClient() instead.
type ManagedRuleSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedRuleSetsClient creates a new instance of ManagedRuleSetsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedRuleSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedRuleSetsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedRuleSetsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists all available managed rule sets.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ManagedRuleSetsClientListOptions contains the optional parameters for the ManagedRuleSetsClient.List method.
func (client *ManagedRuleSetsClient) List(options *ManagedRuleSetsClientListOptions) *ManagedRuleSetsClientListPager {
	return &ManagedRuleSetsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ManagedRuleSetsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedRuleSetDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ManagedRuleSetsClient) listCreateRequest(ctx context.Context, options *ManagedRuleSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedRuleSetsClient) listHandleResponse(resp *http.Response) (ManagedRuleSetsClientListResponse, error) {
	result := ManagedRuleSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedRuleSetDefinitionList); err != nil {
		return ManagedRuleSetsClientListResponse{}, err
	}
	return result, nil
}
