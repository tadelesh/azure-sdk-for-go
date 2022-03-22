//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServicesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckChildrenNameAvailability - This method checks whether a proposed nested resource name is valid and available.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// parameters - Requested name to validate
// options - ServicesClientCheckChildrenNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckChildrenNameAvailability
// method.
func (client *ServicesClient) CheckChildrenNameAvailability(ctx context.Context, groupName string, serviceName string, parameters NameAvailabilityRequest, options *ServicesClientCheckChildrenNameAvailabilityOptions) (ServicesClientCheckChildrenNameAvailabilityResponse, error) {
	req, err := client.checkChildrenNameAvailabilityCreateRequest(ctx, groupName, serviceName, parameters, options)
	if err != nil {
		return ServicesClientCheckChildrenNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCheckChildrenNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckChildrenNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkChildrenNameAvailabilityHandleResponse(resp)
}

// checkChildrenNameAvailabilityCreateRequest creates the CheckChildrenNameAvailability request.
func (client *ServicesClient) checkChildrenNameAvailabilityCreateRequest(ctx context.Context, groupName string, serviceName string, parameters NameAvailabilityRequest, options *ServicesClientCheckChildrenNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkChildrenNameAvailabilityHandleResponse handles the CheckChildrenNameAvailability response.
func (client *ServicesClient) checkChildrenNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckChildrenNameAvailabilityResponse, error) {
	result := ServicesClientCheckChildrenNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityResponse); err != nil {
		return ServicesClientCheckChildrenNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - This method checks whether a proposed top-level resource name is valid and available.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The Azure region of the operation
// parameters - Requested name to validate
// options - ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
func (client *ServicesClient) CheckNameAvailability(ctx context.Context, location string, parameters NameAvailabilityRequest, options *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, parameters NameAvailabilityRequest, options *ServicesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckNameAvailabilityResponse, error) {
	result := ServicesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityResponse); err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckStatus - The services resource is the top-level resource that represents the Database Migration Service. This action
// performs a health check and returns the status of the service and virtual machine size.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientCheckStatusOptions contains the optional parameters for the ServicesClient.CheckStatus method.
func (client *ServicesClient) CheckStatus(ctx context.Context, groupName string, serviceName string, options *ServicesClientCheckStatusOptions) (ServicesClientCheckStatusResponse, error) {
	req, err := client.checkStatusCreateRequest(ctx, groupName, serviceName, options)
	if err != nil {
		return ServicesClientCheckStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCheckStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkStatusHandleResponse(resp)
}

// checkStatusCreateRequest creates the CheckStatus request.
func (client *ServicesClient) checkStatusCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientCheckStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/checkStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkStatusHandleResponse handles the CheckStatus response.
func (client *ServicesClient) checkStatusHandleResponse(resp *http.Response) (ServicesClientCheckStatusResponse, error) {
	result := ServicesClientCheckStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceStatusResponse); err != nil {
		return ServicesClientCheckStatusResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - The services resource is the top-level resource that represents the Database Migration Service. The
// PUT method creates a new service or updates an existing one. When a service is updated, existing
// child resources (i.e. tasks) are unaffected. Services currently support a single kind, "vm", which refers to a VM-based
// service, although other kinds may be added in the future. This method can change
// the kind, SKU, and network of the service, but if tasks are currently running (i.e. the service is busy), this will fail
// with 400 Bad Request ("ServiceIsBusy"). The provider will reply when successful
// with 200 OK or 201 Created. Long-running operations use the provisioningState property.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// parameters - Information about the service
// options - ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate
// method.
func (client *ServicesClient) BeginCreateOrUpdate(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, groupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServicesClientCreateOrUpdateResponse]("ServicesClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServicesClientCreateOrUpdateResponse]("ServicesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - The services resource is the top-level resource that represents the Database Migration Service. The PUT
// method creates a new service or updates an existing one. When a service is updated, existing
// child resources (i.e. tasks) are unaffected. Services currently support a single kind, "vm", which refers to a VM-based
// service, although other kinds may be added in the future. This method can change
// the kind, SKU, and network of the service, but if tasks are currently running (i.e. the service is busy), this will fail
// with 400 Bad Request ("ServiceIsBusy"). The provider will reply when successful
// with 200 OK or 201 Created. Long-running operations use the provisioningState property.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) createOrUpdate(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - The services resource is the top-level resource that represents the Database Migration Service. The DELETE
// method deletes a service. Any running tasks will be canceled.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
func (client *ServicesClient) BeginDelete(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*armruntime.Poller[ServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, groupName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServicesClientDeleteResponse]("ServicesClient.Delete", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServicesClientDeleteResponse]("ServicesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - The services resource is the top-level resource that represents the Database Migration Service. The DELETE method
// deletes a service. Any running tasks will be canceled.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) deleteOperation(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	if options != nil && options.DeleteRunningTasks != nil {
		reqQP.Set("deleteRunningTasks", strconv.FormatBool(*options.DeleteRunningTasks))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves
// information about a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, groupName string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupName, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// List - The services resource is the top-level resource that represents the Database Migration Service. This method returns
// a list of service resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ServicesClientListOptions contains the optional parameters for the ServicesClient.List method.
func (client *ServicesClient) List(options *ServicesClientListOptions) *runtime.Pager[ServicesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServicesClientListResponse]{
		More: func(page ServicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListResponse) (ServicesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, options *ServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/services"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesClientListResponse, error) {
	result := ServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceList); err != nil {
		return ServicesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - The Services resource is the top-level resource that represents the Database Migration Service. This
// method returns a list of service resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// options - ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.ListByResourceGroup
// method.
func (client *ServicesClient) ListByResourceGroup(groupName string, options *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServicesClientListByResourceGroupResponse]{
		More: func(page ServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByResourceGroupResponse) (ServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, groupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, groupName string, options *ServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesClientListByResourceGroupResponse, error) {
	result := ServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceList); err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListSKUs - The services resource is the top-level resource that represents the Database Migration Service. The skus action
// returns the list of SKUs that a service resource can be updated to.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientListSKUsOptions contains the optional parameters for the ServicesClient.ListSKUs method.
func (client *ServicesClient) ListSKUs(groupName string, serviceName string, options *ServicesClientListSKUsOptions) *runtime.Pager[ServicesClientListSKUsResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServicesClientListSKUsResponse]{
		More: func(page ServicesClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListSKUsResponse) (ServicesClientListSKUsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSKUsCreateRequest(ctx, groupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListSKUsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListSKUsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListSKUsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSKUsHandleResponse(resp)
		},
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *ServicesClient) listSKUsCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *ServicesClient) listSKUsHandleResponse(resp *http.Response) (ServicesClientListSKUsResponse, error) {
	result := ServicesClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceSKUList); err != nil {
		return ServicesClientListSKUsResponse{}, err
	}
	return result, nil
}

// BeginStart - The services resource is the top-level resource that represents the Database Migration Service. This action
// starts the service and the service can be used for data migration.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientBeginStartOptions contains the optional parameters for the ServicesClient.BeginStart method.
func (client *ServicesClient) BeginStart(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStartOptions) (*armruntime.Poller[ServicesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, groupName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServicesClientStartResponse]("ServicesClient.Start", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServicesClientStartResponse]("ServicesClient.Start", options.ResumeToken, client.pl, nil)
	}
}

// Start - The services resource is the top-level resource that represents the Database Migration Service. This action starts
// the service and the service can be used for data migration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) start(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, groupName, serviceName, options)
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

// startCreateRequest creates the Start request.
func (client *ServicesClient) startCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - The services resource is the top-level resource that represents the Database Migration Service. This action
// stops the service and the service cannot be used for data migration. The service owner won't
// be billed when the service is stopped.
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// options - ServicesClientBeginStopOptions contains the optional parameters for the ServicesClient.BeginStop method.
func (client *ServicesClient) BeginStop(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStopOptions) (*armruntime.Poller[ServicesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, groupName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServicesClientStopResponse]("ServicesClient.Stop", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServicesClientStopResponse]("ServicesClient.Stop", options.ResumeToken, client.pl, nil)
	}
}

// Stop - The services resource is the top-level resource that represents the Database Migration Service. This action stops
// the service and the service cannot be used for data migration. The service owner won't
// be billed when the service is stopped.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) stop(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, groupName, serviceName, options)
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

// stopCreateRequest creates the Stop request.
func (client *ServicesClient) stopCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServicesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - The services resource is the top-level resource that represents the Database Migration Service. The PATCH
// method updates an existing service. This method can change the kind, SKU, and network of the
// service, but if tasks are currently running (i.e. the service is busy), this will fail with 400 Bad Request ("ServiceIsBusy").
// If the operation fails it returns an *azcore.ResponseError type.
// groupName - Name of the resource group
// serviceName - Name of the service
// parameters - Information about the service
// options - ServicesClientBeginUpdateOptions contains the optional parameters for the ServicesClient.BeginUpdate method.
func (client *ServicesClient) BeginUpdate(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginUpdateOptions) (*armruntime.Poller[ServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, groupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServicesClientUpdateResponse]("ServicesClient.Update", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServicesClientUpdateResponse]("ServicesClient.Update", options.ResumeToken, client.pl, nil)
	}
}

// Update - The services resource is the top-level resource that represents the Database Migration Service. The PATCH method
// updates an existing service. This method can change the kind, SKU, and network of the
// service, but if tasks are currently running (i.e. the service is busy), this will fail with 400 Bad Request ("ServiceIsBusy").
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) update(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, parameters Service, options *ServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
