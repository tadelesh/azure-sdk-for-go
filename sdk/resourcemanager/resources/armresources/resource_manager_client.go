//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armresources

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type ResourceManagerClient struct {
	subscriptionID string
	host           string
	pl             runtime.Pipeline

	resourcesClient             *Client
	deploymentOperationClient   *DeploymentOperationsClient
	deploymentsClient           *DeploymentsClient
	operationsClient            *OperationsClient
	providerResourceTypesClient *ProviderResourceTypesClient
	providersClient             *ProvidersClient
	resourceGroupsClient        *ResourceGroupsClient
}

func NewResourceManagerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceManagerClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ResourceManagerClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}, nil
}

func (mgmtClient *ResourceManagerClient) Resources() *Client {
	if mgmtClient.resourcesClient == nil {
		mgmtClient.resourcesClient = &Client{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.resourcesClient
}

func (mgmtClient *ResourceManagerClient) DeploymentOperations() *DeploymentOperationsClient {
	if mgmtClient.deploymentOperationClient == nil {
		mgmtClient.deploymentOperationClient = &DeploymentOperationsClient{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.deploymentOperationClient
}

func (mgmtClient *ResourceManagerClient) Deployments() *DeploymentsClient {
	if mgmtClient.deploymentsClient == nil {
		mgmtClient.deploymentsClient = &DeploymentsClient{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.deploymentsClient
}
func (mgmtClient *ResourceManagerClient) Operations() *OperationsClient {
	if mgmtClient.operationsClient == nil {
		mgmtClient.operationsClient = &OperationsClient{
			host: mgmtClient.host,
			pl:   mgmtClient.pl,
		}
	}
	return mgmtClient.operationsClient
}
func (mgmtClient *ResourceManagerClient) ProviderResourceTypes() *ProviderResourceTypesClient {
	if mgmtClient.providerResourceTypesClient == nil {
		mgmtClient.providerResourceTypesClient = &ProviderResourceTypesClient{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.providerResourceTypesClient
}
func (mgmtClient *ResourceManagerClient) Providers() *ProvidersClient {
	if mgmtClient.providersClient == nil {
		mgmtClient.providersClient = &ProvidersClient{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.providersClient
}
func (mgmtClient *ResourceManagerClient) ResourceGroups() *ResourceGroupsClient {
	if mgmtClient.resourceGroupsClient == nil {
		mgmtClient.resourceGroupsClient = &ResourceGroupsClient{
			subscriptionID: mgmtClient.subscriptionID,
			host:           mgmtClient.host,
			pl:             mgmtClient.pl,
		}
	}
	return mgmtClient.resourceGroupsClient
}
