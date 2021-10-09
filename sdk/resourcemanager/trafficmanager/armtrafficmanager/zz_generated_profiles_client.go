//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

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

// ProfilesClient contains the methods for the Profiles group.
// Don't use this type directly, use NewProfilesClient() instead.
type ProfilesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProfilesClient creates a new instance of ProfilesClient with the specified values.
func NewProfilesClient(con *arm.Connection, subscriptionID string) *ProfilesClient {
	return &ProfilesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckTrafficManagerRelativeDNSNameAvailability - Checks the availability of a Traffic Manager Relative DNS name.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailability(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters, options *ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityOptions) (ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse, error) {
	req, err := client.checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, client.checkTrafficManagerRelativeDNSNameAvailabilityHandleError(resp)
	}
	return client.checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse(resp)
}

// checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest creates the CheckTrafficManagerRelativeDNSNameAvailability request.
func (client *ProfilesClient) checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters, options *ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Network/checkTrafficManagerNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse handles the CheckTrafficManagerRelativeDNSNameAvailability response.
func (client *ProfilesClient) checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse(resp *http.Response) (ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse, error) {
	result := ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficManagerNameAvailability); err != nil {
		return ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkTrafficManagerRelativeDNSNameAvailabilityHandleError handles the CheckTrafficManagerRelativeDNSNameAvailability error response.
func (client *ProfilesClient) checkTrafficManagerRelativeDNSNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CreateOrUpdate - Create or update a Traffic Manager profile.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesCreateOrUpdateOptions) (ProfilesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, profileName, parameters, options)
	if err != nil {
		return ProfilesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProfilesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (ProfilesCreateOrUpdateResponse, error) {
	result := ProfilesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ProfilesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a Traffic Manager profile.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) Delete(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesDeleteOptions) (ProfilesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return ProfilesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProfilesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ProfilesClient) deleteHandleResponse(resp *http.Response) (ProfilesDeleteResponse, error) {
	result := ProfilesDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteOperationResult); err != nil {
		return ProfilesDeleteResponse{}, err
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *ProfilesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a Traffic Manager profile.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) Get(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesGetOptions) (ProfilesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return ProfilesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProfilesClient) getHandleResponse(resp *http.Response) (ProfilesGetResponse, error) {
	result := ProfilesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ProfilesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all Traffic Manager profiles within a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ProfilesListByResourceGroupOptions) (ProfilesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ProfilesListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProfilesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (ProfilesListByResourceGroupResponse, error) {
	result := ProfilesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ProfilesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists all Traffic Manager profiles within a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) ListBySubscription(ctx context.Context, options *ProfilesListBySubscriptionOptions) (ProfilesListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ProfilesListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProfilesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (ProfilesListBySubscriptionResponse, error) {
	result := ProfilesListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ProfilesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update a Traffic Manager profile.
// If the operation fails it returns the *CloudError error type.
func (client *ProfilesClient) Update(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesUpdateOptions) (ProfilesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, parameters, options)
	if err != nil {
		return ProfilesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProfilesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ProfilesClient) updateHandleResponse(resp *http.Response) (ProfilesUpdateResponse, error) {
	result := ProfilesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ProfilesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
