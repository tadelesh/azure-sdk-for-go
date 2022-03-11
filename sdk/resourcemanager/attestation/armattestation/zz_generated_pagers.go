//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armattestation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// PrivateEndpointConnectionsClientListPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsClientListPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsClientListResponse
	requester func(context.Context) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *PrivateEndpointConnectionsClientListPager) More() bool {
	return reflect.ValueOf(p.current).IsZero()
}

// NextPage advances the pager to the next page.
func (p *PrivateEndpointConnectionsClientListPager) NextPage(ctx context.Context) (PrivateEndpointConnectionsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !p.More() {
		return PrivateEndpointConnectionsClientListResponse{}, errors.New("no more pages")
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return PrivateEndpointConnectionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
