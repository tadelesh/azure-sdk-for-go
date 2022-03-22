//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

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

// Client contains the methods for the SaaS group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host string
	pl   runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) *Client {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &Client{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateResource - Creates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
// parameters - Parameters supplied to the create saas operation.
// options - ClientBeginCreateResourceOptions contains the optional parameters for the Client.BeginCreateResource method.
func (client *Client) BeginCreateResource(ctx context.Context, parameters ResourceCreation, options *ClientBeginCreateResourceOptions) (*armruntime.Poller[ClientCreateResourceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createResource(ctx, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientCreateResourceResponse]("Client.CreateResource", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientCreateResourceResponse]("Client.CreateResource", options.ResumeToken, client.pl, nil)
	}
}

// CreateResource - Creates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) createResource(ctx context.Context, parameters ResourceCreation, options *ClientBeginCreateResourceOptions) (*http.Response, error) {
	req, err := client.createResourceCreateRequest(ctx, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createResourceCreateRequest creates the CreateResource request.
func (client *Client) createResourceCreateRequest(ctx context.Context, parameters ResourceCreation, options *ClientBeginCreateResourceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// parameters - Parameters supplied to delete saas operation.
// options - ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
func (client *Client) BeginDelete(ctx context.Context, resourceID string, parameters DeleteOptions, options *ClientBeginDeleteOptions) (*armruntime.Poller[ClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceID, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientDeleteResponse]("Client.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientDeleteResponse]("Client.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) deleteOperation(ctx context.Context, resourceID string, parameters DeleteOptions, options *ClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, resourceID string, parameters DeleteOptions, options *ClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources/{resourceId}"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// GetResource - Gets information about the specified SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// options - ClientGetResourceOptions contains the optional parameters for the Client.GetResource method.
func (client *Client) GetResource(ctx context.Context, resourceID string, options *ClientGetResourceOptions) (ClientGetResourceResponse, error) {
	req, err := client.getResourceCreateRequest(ctx, resourceID, options)
	if err != nil {
		return ClientGetResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getResourceHandleResponse(resp)
}

// getResourceCreateRequest creates the GetResource request.
func (client *Client) getResourceCreateRequest(ctx context.Context, resourceID string, options *ClientGetResourceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources/{resourceId}"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getResourceHandleResponse handles the GetResource response.
func (client *Client) getResourceHandleResponse(resp *http.Response) (ClientGetResourceResponse, error) {
	result := ClientGetResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Resource); err != nil {
		return ClientGetResourceResponse{}, err
	}
	return result, nil
}

// BeginUpdateResource - Updates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// parameters - Parameters supplied to the update saas operation.
// options - ClientBeginUpdateResourceOptions contains the optional parameters for the Client.BeginUpdateResource method.
func (client *Client) BeginUpdateResource(ctx context.Context, resourceID string, parameters ResourceCreation, options *ClientBeginUpdateResourceOptions) (*armruntime.Poller[ClientUpdateResourceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateResource(ctx, resourceID, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientUpdateResourceResponse]("Client.UpdateResource", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientUpdateResourceResponse]("Client.UpdateResource", options.ResumeToken, client.pl, nil)
	}
}

// UpdateResource - Updates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) updateResource(ctx context.Context, resourceID string, parameters ResourceCreation, options *ClientBeginUpdateResourceOptions) (*http.Response, error) {
	req, err := client.updateResourceCreateRequest(ctx, resourceID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateResourceCreateRequest creates the UpdateResource request.
func (client *Client) updateResourceCreateRequest(ctx context.Context, resourceID string, parameters ResourceCreation, options *ClientBeginUpdateResourceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources/{resourceId}"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
