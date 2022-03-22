//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedDatabaseSensitivityLabelsClient contains the methods for the ManagedDatabaseSensitivityLabels group.
// Don't use this type directly, use NewManagedDatabaseSensitivityLabelsClient() instead.
type ManagedDatabaseSensitivityLabelsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedDatabaseSensitivityLabelsClient creates a new instance of ManagedDatabaseSensitivityLabelsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedDatabaseSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedDatabaseSensitivityLabelsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedDatabaseSensitivityLabelsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// schemaName - The name of the schema.
// tableName - The name of the table.
// columnName - The name of the column.
// parameters - The column sensitivity label resource.
// options - ManagedDatabaseSensitivityLabelsClientCreateOrUpdateOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *ManagedDatabaseSensitivityLabelsClientCreateOrUpdateOptions) (ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, parameters, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedDatabaseSensitivityLabelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *ManagedDatabaseSensitivityLabelsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedDatabaseSensitivityLabelsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse, error) {
	result := ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// schemaName - The name of the schema.
// tableName - The name of the table.
// columnName - The name of the column.
// options - ManagedDatabaseSensitivityLabelsClientDeleteOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.Delete
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientDeleteOptions) (ManagedDatabaseSensitivityLabelsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSensitivityLabelsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedDatabaseSensitivityLabelsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedDatabaseSensitivityLabelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// DisableRecommendation - Disables sensitivity recommendations on a given column
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// schemaName - The name of the schema.
// tableName - The name of the table.
// columnName - The name of the column.
// options - ManagedDatabaseSensitivityLabelsClientDisableRecommendationOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.DisableRecommendation
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientDisableRecommendationOptions) (ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse, error) {
	req, err := client.disableRecommendationCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse{}, nil
}

// disableRecommendationCreateRequest creates the DisableRecommendation request.
func (client *ManagedDatabaseSensitivityLabelsClient) disableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientDisableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// EnableRecommendation - Enables sensitivity recommendations on a given column (recommendations are enabled by default on
// all columns)
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// schemaName - The name of the schema.
// tableName - The name of the table.
// columnName - The name of the column.
// options - ManagedDatabaseSensitivityLabelsClientEnableRecommendationOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.EnableRecommendation
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientEnableRecommendationOptions) (ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse, error) {
	req, err := client.enableRecommendationCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse{}, nil
}

// enableRecommendationCreateRequest creates the EnableRecommendation request.
func (client *ManagedDatabaseSensitivityLabelsClient) enableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseSensitivityLabelsClientEnableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// schemaName - The name of the schema.
// tableName - The name of the table.
// columnName - The name of the column.
// sensitivityLabelSource - The source of the sensitivity label.
// options - ManagedDatabaseSensitivityLabelsClientGetOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.Get
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *ManagedDatabaseSensitivityLabelsClientGetOptions) (ManagedDatabaseSensitivityLabelsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSensitivityLabelsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseSensitivityLabelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *ManagedDatabaseSensitivityLabelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	if sensitivityLabelSource == "" {
		return nil, errors.New("parameter sensitivityLabelSource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape(string(sensitivityLabelSource)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseSensitivityLabelsClient) getHandleResponse(resp *http.Response) (ManagedDatabaseSensitivityLabelsClientGetResponse, error) {
	result := ManagedDatabaseSensitivityLabelsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return ManagedDatabaseSensitivityLabelsClientGetResponse{}, err
	}
	return result, nil
}

// ListCurrentByDatabase - Gets the sensitivity labels of a given database
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// options - ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.ListCurrentByDatabase
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions) *runtime.Pager[ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse]{
		More: func(page ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse) (ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCurrentByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listCurrentByDatabaseHandleResponse(resp)
		},
	})
}

// listCurrentByDatabaseCreateRequest creates the ListCurrentByDatabase request.
func (client *ManagedDatabaseSensitivityLabelsClient) listCurrentByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/currentSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCurrentByDatabaseHandleResponse handles the ListCurrentByDatabase response.
func (client *ManagedDatabaseSensitivityLabelsClient) listCurrentByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse, error) {
	result := ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse{}, err
	}
	return result, nil
}

// ListRecommendedByDatabase - Gets the sensitivity labels of a given database
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// options - ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions contains the optional parameters for the
// ManagedDatabaseSensitivityLabelsClient.ListRecommendedByDatabase method.
func (client *ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions) *runtime.Pager[ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse]{
		More: func(page ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse) (ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRecommendedByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRecommendedByDatabaseHandleResponse(resp)
		},
	})
}

// listRecommendedByDatabaseCreateRequest creates the ListRecommendedByDatabase request.
func (client *ManagedDatabaseSensitivityLabelsClient) listRecommendedByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/recommendedSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.IncludeDisabledRecommendations != nil {
		reqQP.Set("includeDisabledRecommendations", strconv.FormatBool(*options.IncludeDisabledRecommendations))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRecommendedByDatabaseHandleResponse handles the ListRecommendedByDatabase response.
func (client *ManagedDatabaseSensitivityLabelsClient) listRecommendedByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse, error) {
	result := ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
	}
	return result, nil
}

// Update - Update sensitivity labels of a given database using an operations batch.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the database.
// options - ManagedDatabaseSensitivityLabelsClientUpdateOptions contains the optional parameters for the ManagedDatabaseSensitivityLabelsClient.Update
// method.
func (client *ManagedDatabaseSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters SensitivityLabelUpdateList, options *ManagedDatabaseSensitivityLabelsClientUpdateOptions) (ManagedDatabaseSensitivityLabelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSensitivityLabelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSensitivityLabelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedDatabaseSensitivityLabelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedDatabaseSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters SensitivityLabelUpdateList, options *ManagedDatabaseSensitivityLabelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/currentSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}
