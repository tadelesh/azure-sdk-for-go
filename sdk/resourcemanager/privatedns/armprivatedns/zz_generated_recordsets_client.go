//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

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
	"strconv"
	"strings"
)

// RecordSetsClient contains the methods for the RecordSets group.
// Don't use this type directly, use NewRecordSetsClient() instead.
type RecordSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecordSetsClient creates a new instance of RecordSetsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecordSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RecordSetsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RecordSetsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates a record set within a Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// recordType - The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are
// created when the Private DNS zone is created).
// relativeRecordSetName - The name of the record set, relative to the name of the zone.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - RecordSetsClientCreateOrUpdateOptions contains the optional parameters for the RecordSetsClient.CreateOrUpdate
// method.
func (client *RecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsClientCreateOrUpdateOptions) (RecordSetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, options)
	if err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RecordSetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RecordSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RecordSetsClient) createOrUpdateHandleResponse(resp *http.Response) (RecordSetsClientCreateOrUpdateResponse, error) {
	result := RecordSetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a record set from a Private DNS zone. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// recordType - The type of DNS record in this record set. Record sets of type SOA cannot be deleted (they are deleted when
// the Private DNS zone is deleted).
// relativeRecordSetName - The name of the record set, relative to the name of the zone.
// options - RecordSetsClientDeleteOptions contains the optional parameters for the RecordSetsClient.Delete method.
func (client *RecordSetsClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsClientDeleteOptions) (RecordSetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, options)
	if err != nil {
		return RecordSetsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecordSetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RecordSetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RecordSetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RecordSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a record set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// recordType - The type of DNS record in this record set.
// relativeRecordSetName - The name of the record set, relative to the name of the zone.
// options - RecordSetsClientGetOptions contains the optional parameters for the RecordSetsClient.Get method.
func (client *RecordSetsClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsClientGetOptions) (RecordSetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, options)
	if err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecordSetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecordSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecordSetsClient) getHandleResponse(resp *http.Response) (RecordSetsClientGetResponse, error) {
	result := RecordSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all record sets in a Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// options - RecordSetsClientListOptions contains the optional parameters for the RecordSetsClient.List method.
func (client *RecordSetsClient) List(resourceGroupName string, privateZoneName string, options *RecordSetsClientListOptions) *RecordSetsClientListPager {
	return &RecordSetsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateZoneName, options)
		},
		advancer: func(ctx context.Context, resp RecordSetsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecordSetListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RecordSetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, options *RecordSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecordSetsClient) listHandleResponse(resp *http.Response) (RecordSetsClientListResponse, error) {
	result := RecordSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSetListResult); err != nil {
		return RecordSetsClientListResponse{}, err
	}
	return result, nil
}

// ListByType - Lists the record sets of a specified type in a Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// recordType - The type of record sets to enumerate.
// options - RecordSetsClientListByTypeOptions contains the optional parameters for the RecordSetsClient.ListByType method.
func (client *RecordSetsClient) ListByType(resourceGroupName string, privateZoneName string, recordType RecordType, options *RecordSetsClientListByTypeOptions) *RecordSetsClientListByTypePager {
	return &RecordSetsClientListByTypePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTypeCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, options)
		},
		advancer: func(ctx context.Context, resp RecordSetsClientListByTypeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecordSetListResult.NextLink)
		},
	}
}

// listByTypeCreateRequest creates the ListByType request.
func (client *RecordSetsClient) listByTypeCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, options *RecordSetsClientListByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTypeHandleResponse handles the ListByType response.
func (client *RecordSetsClient) listByTypeHandleResponse(resp *http.Response) (RecordSetsClientListByTypeResponse, error) {
	result := RecordSetsClientListByTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSetListResult); err != nil {
		return RecordSetsClientListByTypeResponse{}, err
	}
	return result, nil
}

// Update - Updates a record set within a Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// privateZoneName - The name of the Private DNS zone (without a terminating dot).
// recordType - The type of DNS record in this record set.
// relativeRecordSetName - The name of the record set, relative to the name of the zone.
// parameters - Parameters supplied to the Update operation.
// options - RecordSetsClientUpdateOptions contains the optional parameters for the RecordSetsClient.Update method.
func (client *RecordSetsClient) Update(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsClientUpdateOptions) (RecordSetsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, options)
	if err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecordSetsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RecordSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RecordSetsClient) updateHandleResponse(resp *http.Response) (RecordSetsClientUpdateResponse, error) {
	result := RecordSetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	return result, nil
}
