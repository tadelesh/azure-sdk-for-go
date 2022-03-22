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
	"strings"
)

// DatabaseMigrationsSQLMiClient contains the methods for the DatabaseMigrationsSQLMi group.
// Don't use this type directly, use NewDatabaseMigrationsSQLMiClient() instead.
type DatabaseMigrationsSQLMiClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDatabaseMigrationsSQLMiClient creates a new instance of DatabaseMigrationsSQLMiClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDatabaseMigrationsSQLMiClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabaseMigrationsSQLMiClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DatabaseMigrationsSQLMiClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCancel - Stop migrations in progress for the database
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// targetDbName - The name of the target database.
// parameters - Required migration operation ID for which cancel will be initiated.
// options - DatabaseMigrationsSQLMiClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCancel
// method.
func (client *DatabaseMigrationsSQLMiClient) BeginCancel(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*armruntime.Poller[DatabaseMigrationsSQLMiClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DatabaseMigrationsSQLMiClientCancelResponse]("DatabaseMigrationsSQLMiClient.Cancel", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DatabaseMigrationsSQLMiClientCancelResponse]("DatabaseMigrationsSQLMiClient.Cancel", options.ResumeToken, client.pl, nil)
	}
}

// Cancel - Stop migrations in progress for the database
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabaseMigrationsSQLMiClient) cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
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

// cancelCreateRequest creates the Cancel request.
func (client *DatabaseMigrationsSQLMiClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdate - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// targetDbName - The name of the target database.
// parameters - Details of SqlMigrationService resource.
// options - DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate
// method.
func (client *DatabaseMigrationsSQLMiClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, options *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]("DatabaseMigrationsSQLMiClient.CreateOrUpdate", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]("DatabaseMigrationsSQLMiClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabaseMigrationsSQLMiClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, options *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseMigrationsSQLMiClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, options *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// BeginCutover - Initiate cutover for online migration in progress for the database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// targetDbName - The name of the target database.
// parameters - Required migration operation ID for which cutover will be initiated.
// options - DatabaseMigrationsSQLMiClientBeginCutoverOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCutover
// method.
func (client *DatabaseMigrationsSQLMiClient) BeginCutover(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*armruntime.Poller[DatabaseMigrationsSQLMiClientCutoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cutover(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DatabaseMigrationsSQLMiClientCutoverResponse]("DatabaseMigrationsSQLMiClient.Cutover", "", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DatabaseMigrationsSQLMiClientCutoverResponse]("DatabaseMigrationsSQLMiClient.Cutover", options.ResumeToken, client.pl, nil)
	}
}

// Cutover - Initiate cutover for online migration in progress for the database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabaseMigrationsSQLMiClient) cutover(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*http.Response, error) {
	req, err := client.cutoverCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
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

// cutoverCreateRequest creates the Cutover request.
func (client *DatabaseMigrationsSQLMiClient) cutoverCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cutover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Retrieve the Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// targetDbName - The name of the target database.
// options - DatabaseMigrationsSQLMiClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.Get
// method.
func (client *DatabaseMigrationsSQLMiClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, options *DatabaseMigrationsSQLMiClientGetOptions) (DatabaseMigrationsSQLMiClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, options)
	if err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseMigrationsSQLMiClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseMigrationsSQLMiClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, options *DatabaseMigrationsSQLMiClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MigrationOperationID != nil {
		reqQP.Set("migrationOperationId", *options.MigrationOperationID)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseMigrationsSQLMiClient) getHandleResponse(resp *http.Response) (DatabaseMigrationsSQLMiClientGetResponse, error) {
	result := DatabaseMigrationsSQLMiClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationSQLMi); err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	return result, nil
}
