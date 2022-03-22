//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

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

// ConfigurationsClient contains the methods for the Configurations group.
// Don't use this type directly, use NewConfigurationsClient() instead.
type ConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateInResourceGroup - Create/Overwrite Azure Advisor configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// configurationName - Advisor configuration name. Value must be 'default'
// resourceGroup - The name of the Azure resource group.
// configContract - The Azure Advisor configuration data structure.
// options - ConfigurationsClientCreateInResourceGroupOptions contains the optional parameters for the ConfigurationsClient.CreateInResourceGroup
// method.
func (client *ConfigurationsClient) CreateInResourceGroup(ctx context.Context, configurationName ConfigurationName, resourceGroup string, configContract ConfigData, options *ConfigurationsClientCreateInResourceGroupOptions) (ConfigurationsClientCreateInResourceGroupResponse, error) {
	req, err := client.createInResourceGroupCreateRequest(ctx, configurationName, resourceGroup, configContract, options)
	if err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationsClientCreateInResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.createInResourceGroupHandleResponse(resp)
}

// createInResourceGroupCreateRequest creates the CreateInResourceGroup request.
func (client *ConfigurationsClient) createInResourceGroupCreateRequest(ctx context.Context, configurationName ConfigurationName, resourceGroup string, configContract ConfigData, options *ConfigurationsClientCreateInResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configContract)
}

// createInResourceGroupHandleResponse handles the CreateInResourceGroup response.
func (client *ConfigurationsClient) createInResourceGroupHandleResponse(resp *http.Response) (ConfigurationsClientCreateInResourceGroupResponse, error) {
	result := ConfigurationsClientCreateInResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigData); err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	return result, nil
}

// CreateInSubscription - Create/Overwrite Azure Advisor configuration and also delete all configurations of contained resource
// groups.
// If the operation fails it returns an *azcore.ResponseError type.
// configurationName - Advisor configuration name. Value must be 'default'
// configContract - The Azure Advisor configuration data structure.
// options - ConfigurationsClientCreateInSubscriptionOptions contains the optional parameters for the ConfigurationsClient.CreateInSubscription
// method.
func (client *ConfigurationsClient) CreateInSubscription(ctx context.Context, configurationName ConfigurationName, configContract ConfigData, options *ConfigurationsClientCreateInSubscriptionOptions) (ConfigurationsClientCreateInSubscriptionResponse, error) {
	req, err := client.createInSubscriptionCreateRequest(ctx, configurationName, configContract, options)
	if err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationsClientCreateInSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.createInSubscriptionHandleResponse(resp)
}

// createInSubscriptionCreateRequest creates the CreateInSubscription request.
func (client *ConfigurationsClient) createInSubscriptionCreateRequest(ctx context.Context, configurationName ConfigurationName, configContract ConfigData, options *ConfigurationsClientCreateInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configContract)
}

// createInSubscriptionHandleResponse handles the CreateInSubscription response.
func (client *ConfigurationsClient) createInSubscriptionHandleResponse(resp *http.Response) (ConfigurationsClientCreateInSubscriptionResponse, error) {
	result := ConfigurationsClientCreateInSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigData); err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieve Azure Advisor configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - The name of the Azure resource group.
// options - ConfigurationsClientListByResourceGroupOptions contains the optional parameters for the ConfigurationsClient.ListByResourceGroup
// method.
func (client *ConfigurationsClient) ListByResourceGroup(resourceGroup string, options *ConfigurationsClientListByResourceGroupOptions) *runtime.Pager[ConfigurationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConfigurationsClientListByResourceGroupResponse]{
		More: func(page ConfigurationsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListByResourceGroupResponse) (ConfigurationsClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
			if err != nil {
				return ConfigurationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *ConfigurationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationsClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationsClientListByResourceGroupResponse, error) {
	result := ConfigurationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationListResult); err != nil {
		return ConfigurationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieve Azure Advisor configurations and also retrieve configurations of contained resource groups.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConfigurationsClientListBySubscriptionOptions contains the optional parameters for the ConfigurationsClient.ListBySubscription
// method.
func (client *ConfigurationsClient) ListBySubscription(options *ConfigurationsClientListBySubscriptionOptions) *runtime.Pager[ConfigurationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConfigurationsClientListBySubscriptionResponse]{
		More: func(page ConfigurationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListBySubscriptionResponse) (ConfigurationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationsClientListBySubscriptionResponse, error) {
	result := ConfigurationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationListResult); err != nil {
		return ConfigurationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
