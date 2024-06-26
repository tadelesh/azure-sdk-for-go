// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/messaging"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SenderClient contains the methods for the Microsoft.EventGrid namespace.
// Don't use this type directly, use a constructor function instead.
type SenderClient struct {
	data senderData

	internal *azcore.Client
	endpoint string
}

// send - Publish a single Cloud Event to a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - topicName - Topic Name.
//   - event - Single Cloud Event being published.
//   - options - SendOptions contains the optional parameters for the SenderClient.send method.
func (client *SenderClient) send(ctx context.Context, topicName string, event *messaging.CloudEvent, options *SendOptions) (SendResponse, error) {
	var err error
	req, err := client.sendCreateRequest(ctx, topicName, event, options)
	if err != nil {
		return SendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SendResponse{}, err
	}
	resp, err := client.sendHandleResponse(httpResp)
	return resp, err
}

// sendCreateRequest creates the send request.
func (client *SenderClient) sendCreateRequest(ctx context.Context, topicName string, event *messaging.CloudEvent, _ *SendOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/topics/{topicName}:publish"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/cloudevents+json; charset=utf-8"}
	if err := runtime.MarshalAsJSON(req, event); err != nil {
		return nil, err
	}
	return req, nil
}

// sendHandleResponse handles the send response.
func (client *SenderClient) sendHandleResponse(resp *http.Response) (SendResponse, error) {
	result := SendResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishResult); err != nil {
		return SendResponse{}, err
	}
	return result, nil
}

// sendEvents - Publish a batch of Cloud Events to a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - topicName - Topic Name.
//   - events - Array of Cloud Events being published.
//   - options - senderClientsendEventsOptions contains the optional parameters for the SenderClient.sendEvents method.
func (client *SenderClient) sendEvents(ctx context.Context, topicName string, events []*messaging.CloudEvent, options *SendEventsOptions) (SendEventsResponse, error) {
	var err error
	req, err := client.sendEventsCreateRequest(ctx, topicName, events, options)
	if err != nil {
		return SendEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SendEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SendEventsResponse{}, err
	}
	resp, err := client.sendEventsHandleResponse(httpResp)
	return resp, err
}

// sendEventsCreateRequest creates the sendEvents request.
func (client *SenderClient) sendEventsCreateRequest(ctx context.Context, topicName string, events []*messaging.CloudEvent, _ *SendEventsOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/topics/{topicName}:publish"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/cloudevents-batch+json; charset=utf-8"}
	if err := runtime.MarshalAsJSON(req, events); err != nil {
		return nil, err
	}
	return req, nil
}

// sendEventsHandleResponse handles the sendEvents response.
func (client *SenderClient) sendEventsHandleResponse(resp *http.Response) (SendEventsResponse, error) {
	result := SendEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishResult); err != nil {
		return SendEventsResponse{}, err
	}
	return result, nil
}
