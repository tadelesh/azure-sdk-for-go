//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorSimple1200Series

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

// IscsiDisksClient contains the methods for the IscsiDisks group.
// Don't use this type directly, use NewIscsiDisksClient() instead.
type IscsiDisksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIscsiDisksClient creates a new instance of IscsiDisksClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIscsiDisksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IscsiDisksClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IscsiDisksClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the iSCSI disk.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// diskName - The disk name.
// resourceGroupName - The resource group name
// managerName - The manager name
// iscsiDisk - The iSCSI disk.
// options - IscsiDisksClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiDisksClient.BeginCreateOrUpdate
// method.
func (client *IscsiDisksClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, iscsiDisk ISCSIDisk, options *IscsiDisksClientBeginCreateOrUpdateOptions) (*armruntime.Poller[IscsiDisksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, iscsiDisk, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[IscsiDisksClientCreateOrUpdateResponse]("IscsiDisksClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[IscsiDisksClientCreateOrUpdateResponse]("IscsiDisksClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the iSCSI disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiDisksClient) createOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, iscsiDisk ISCSIDisk, options *IscsiDisksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, iscsiDisk, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IscsiDisksClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, iscsiDisk ISCSIDisk, options *IscsiDisksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iscsiDisk)
}

// BeginDelete - Deletes the iSCSI disk.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// diskName - The disk name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientBeginDeleteOptions contains the optional parameters for the IscsiDisksClient.BeginDelete method.
func (client *IscsiDisksClient) BeginDelete(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientBeginDeleteOptions) (*armruntime.Poller[IscsiDisksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[IscsiDisksClientDeleteResponse]("IscsiDisksClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[IscsiDisksClientDeleteResponse]("IscsiDisksClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the iSCSI disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IscsiDisksClient) deleteOperation(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IscsiDisksClient) deleteCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of the specified iSCSI disk name.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// diskName - The disk name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientGetOptions contains the optional parameters for the IscsiDisksClient.Get method.
func (client *IscsiDisksClient) Get(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientGetOptions) (IscsiDisksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, options)
	if err != nil {
		return IscsiDisksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IscsiDisksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IscsiDisksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IscsiDisksClient) getCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IscsiDisksClient) getHandleResponse(resp *http.Response) (IscsiDisksClientGetResponse, error) {
	result := IscsiDisksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIDisk); err != nil {
		return IscsiDisksClientGetResponse{}, err
	}
	return result, nil
}

// ListByDevice - Retrieves all the iSCSI disks in a device.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientListByDeviceOptions contains the optional parameters for the IscsiDisksClient.ListByDevice method.
func (client *IscsiDisksClient) ListByDevice(deviceName string, resourceGroupName string, managerName string, options *IscsiDisksClientListByDeviceOptions) *runtime.Pager[IscsiDisksClientListByDeviceResponse] {
	return runtime.NewPager(runtime.PageProcessor[IscsiDisksClientListByDeviceResponse]{
		More: func(page IscsiDisksClientListByDeviceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IscsiDisksClientListByDeviceResponse) (IscsiDisksClientListByDeviceResponse, error) {
			req, err := client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
			if err != nil {
				return IscsiDisksClientListByDeviceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IscsiDisksClientListByDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IscsiDisksClientListByDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDeviceHandleResponse(resp)
		},
	})
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *IscsiDisksClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *IscsiDisksClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/disks"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *IscsiDisksClient) listByDeviceHandleResponse(resp *http.Response) (IscsiDisksClientListByDeviceResponse, error) {
	result := IscsiDisksClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIDiskList); err != nil {
		return IscsiDisksClientListByDeviceResponse{}, err
	}
	return result, nil
}

// ListByIscsiServer - Retrieves all the disks in a iSCSI server.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientListByIscsiServerOptions contains the optional parameters for the IscsiDisksClient.ListByIscsiServer
// method.
func (client *IscsiDisksClient) ListByIscsiServer(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiDisksClientListByIscsiServerOptions) *runtime.Pager[IscsiDisksClientListByIscsiServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[IscsiDisksClientListByIscsiServerResponse]{
		More: func(page IscsiDisksClientListByIscsiServerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IscsiDisksClientListByIscsiServerResponse) (IscsiDisksClientListByIscsiServerResponse, error) {
			req, err := client.listByIscsiServerCreateRequest(ctx, deviceName, iscsiServerName, resourceGroupName, managerName, options)
			if err != nil {
				return IscsiDisksClientListByIscsiServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IscsiDisksClientListByIscsiServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IscsiDisksClientListByIscsiServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByIscsiServerHandleResponse(resp)
		},
	})
}

// listByIscsiServerCreateRequest creates the ListByIscsiServer request.
func (client *IscsiDisksClient) listByIscsiServerCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *IscsiDisksClientListByIscsiServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByIscsiServerHandleResponse handles the ListByIscsiServer response.
func (client *IscsiDisksClient) listByIscsiServerHandleResponse(resp *http.Response) (IscsiDisksClientListByIscsiServerResponse, error) {
	result := IscsiDisksClientListByIscsiServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISCSIDiskList); err != nil {
		return IscsiDisksClientListByIscsiServerResponse{}, err
	}
	return result, nil
}

// ListMetricDefinition - Retrieves metric definitions for all metric aggregated at the iSCSI disk.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// diskName - The iSCSI disk name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientListMetricDefinitionOptions contains the optional parameters for the IscsiDisksClient.ListMetricDefinition
// method.
func (client *IscsiDisksClient) ListMetricDefinition(deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientListMetricDefinitionOptions) *runtime.Pager[IscsiDisksClientListMetricDefinitionResponse] {
	return runtime.NewPager(runtime.PageProcessor[IscsiDisksClientListMetricDefinitionResponse]{
		More: func(page IscsiDisksClientListMetricDefinitionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IscsiDisksClientListMetricDefinitionResponse) (IscsiDisksClientListMetricDefinitionResponse, error) {
			req, err := client.listMetricDefinitionCreateRequest(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, options)
			if err != nil {
				return IscsiDisksClientListMetricDefinitionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IscsiDisksClientListMetricDefinitionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IscsiDisksClientListMetricDefinitionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricDefinitionHandleResponse(resp)
		},
	})
}

// listMetricDefinitionCreateRequest creates the ListMetricDefinition request.
func (client *IscsiDisksClient) listMetricDefinitionCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientListMetricDefinitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metricsDefinitions"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricDefinitionHandleResponse handles the ListMetricDefinition response.
func (client *IscsiDisksClient) listMetricDefinitionHandleResponse(resp *http.Response) (IscsiDisksClientListMetricDefinitionResponse, error) {
	result := IscsiDisksClientListMetricDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionList); err != nil {
		return IscsiDisksClientListMetricDefinitionResponse{}, err
	}
	return result, nil
}

// ListMetrics - Gets the iSCSI disk metrics
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// iscsiServerName - The iSCSI server name.
// diskName - The iSCSI disk name.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - IscsiDisksClientListMetricsOptions contains the optional parameters for the IscsiDisksClient.ListMetrics method.
func (client *IscsiDisksClient) ListMetrics(deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientListMetricsOptions) *runtime.Pager[IscsiDisksClientListMetricsResponse] {
	return runtime.NewPager(runtime.PageProcessor[IscsiDisksClientListMetricsResponse]{
		More: func(page IscsiDisksClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IscsiDisksClientListMetricsResponse) (IscsiDisksClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, deviceName, iscsiServerName, diskName, resourceGroupName, managerName, options)
			if err != nil {
				return IscsiDisksClientListMetricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IscsiDisksClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IscsiDisksClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *IscsiDisksClient) listMetricsCreateRequest(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *IscsiDisksClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metrics"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if iscsiServerName == "" {
		return nil, errors.New("parameter iscsiServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiServerName}", url.PathEscape(iscsiServerName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *IscsiDisksClient) listMetricsHandleResponse(resp *http.Response) (IscsiDisksClientListMetricsResponse, error) {
	result := IscsiDisksClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricList); err != nil {
		return IscsiDisksClientListMetricsResponse{}, err
	}
	return result, nil
}
