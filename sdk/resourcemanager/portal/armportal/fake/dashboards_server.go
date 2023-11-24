//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"net/http"
	"net/url"
	"regexp"
)

// DashboardsServer is a fake server for instances of the armportal.DashboardsClient type.
type DashboardsServer struct {
	// CreateOrUpdate is the fake for method DashboardsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, dashboardName string, dashboard armportal.Dashboard, options *armportal.DashboardsClientCreateOrUpdateOptions) (resp azfake.Responder[armportal.DashboardsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DashboardsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, dashboardName string, options *armportal.DashboardsClientDeleteOptions) (resp azfake.Responder[armportal.DashboardsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DashboardsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	Get func(ctx context.Context, resourceGroupName string, dashboardName string, options *armportal.DashboardsClientGetOptions) (resp azfake.Responder[armportal.DashboardsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method DashboardsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armportal.DashboardsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armportal.DashboardsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method DashboardsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armportal.DashboardsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armportal.DashboardsClientListBySubscriptionResponse])

	// Update is the fake for method DashboardsClient.Update
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	Update func(ctx context.Context, resourceGroupName string, dashboardName string, dashboard armportal.PatchableDashboard, options *armportal.DashboardsClientUpdateOptions) (resp azfake.Responder[armportal.DashboardsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDashboardsServerTransport creates a new instance of DashboardsServerTransport with the provided implementation.
// The returned DashboardsServerTransport instance is connected to an instance of armportal.DashboardsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDashboardsServerTransport(srv *DashboardsServer) *DashboardsServerTransport {
	return &DashboardsServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armportal.DashboardsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armportal.DashboardsClientListBySubscriptionResponse]](),
	}
}

// DashboardsServerTransport connects instances of armportal.DashboardsClient to instances of DashboardsServer.
// Don't use this type directly, use NewDashboardsServerTransport instead.
type DashboardsServerTransport struct {
	srv                         *DashboardsServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armportal.DashboardsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armportal.DashboardsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for DashboardsServerTransport.
func (d *DashboardsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DashboardsClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DashboardsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DashboardsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DashboardsClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DashboardsClient.NewListBySubscriptionPager":
		resp, err = d.dispatchNewListBySubscriptionPager(req)
	case "DashboardsClient.Update":
		resp, err = d.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DashboardsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards/(?P<dashboardName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armportal.Dashboard](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dashboardNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dashboardName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, dashboardNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Dashboard, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DashboardsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards/(?P<dashboardName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dashboardNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dashboardName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, dashboardNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DashboardsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards/(?P<dashboardName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dashboardNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dashboardName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, dashboardNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Dashboard, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DashboardsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := d.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		d.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armportal.DashboardsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		d.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DashboardsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := d.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		d.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armportal.DashboardsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		d.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (d *DashboardsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Portal/dashboards/(?P<dashboardName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armportal.PatchableDashboard](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dashboardNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dashboardName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameParam, dashboardNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Dashboard, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
