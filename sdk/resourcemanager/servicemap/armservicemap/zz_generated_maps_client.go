//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicemap

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

// MapsClient contains the methods for the Maps group.
// Don't use this type directly, use NewMapsClient() instead.
type MapsClient struct {
	host string
	subscriptionID string
	pl runtime.Pipeline
}

// NewMapsClient creates a new instance of MapsClient with the specified values.
// subscriptionID - Azure subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MapsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &MapsClient{
		subscriptionID: subscriptionID,
		host: string(ep),
		pl: armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Generate - Generates the specified map.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// request - Request options.
// options - MapsClientGenerateOptions contains the optional parameters for the MapsClient.Generate method.
func (client *MapsClient) Generate(ctx context.Context, resourceGroupName string, workspaceName string, request MapRequestClassification, options *MapsClientGenerateOptions) (MapsClientGenerateResponse, error) {
	req, err := client.generateCreateRequest(ctx, resourceGroupName, workspaceName, request, options)
	if err != nil {
		return MapsClientGenerateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MapsClientGenerateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MapsClientGenerateResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateHandleResponse(resp)
}

// generateCreateRequest creates the Generate request.
func (client *MapsClient) generateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, request MapRequestClassification, options *MapsClientGenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/generateMap"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// generateHandleResponse handles the Generate response.
func (client *MapsClient) generateHandleResponse(resp *http.Response) (MapsClientGenerateResponse, error) {
	result := MapsClientGenerateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MapResponse); err != nil {
		return MapsClientGenerateResponse{}, err
	}
	return result, nil
}
