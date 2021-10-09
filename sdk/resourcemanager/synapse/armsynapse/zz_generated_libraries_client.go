//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LibrariesClient contains the methods for the Libraries group.
// Don't use this type directly, use NewLibrariesClient() instead.
type LibrariesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLibrariesClient creates a new instance of LibrariesClient with the specified values.
func NewLibrariesClient(con *arm.Connection, subscriptionID string) *LibrariesClient {
	return &LibrariesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListByWorkspace - List libraries in a workspace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LibrariesClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *LibrariesListByWorkspaceOptions) *LibrariesListByWorkspacePager {
	return &LibrariesListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp LibrariesListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LibraryListResponse.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LibrariesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LibrariesListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/libraries"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LibrariesClient) listByWorkspaceHandleResponse(resp *http.Response) (LibrariesListByWorkspaceResponse, error) {
	result := LibrariesListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryListResponse); err != nil {
		return LibrariesListByWorkspaceResponse{}, err
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *LibrariesClient) listByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
