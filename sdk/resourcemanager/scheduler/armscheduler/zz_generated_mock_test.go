//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"golang.org/x/net/http2"
)

var (
	ctx            context.Context
	subscriptionId string
	cred           azcore.TokenCredential
	err            error
	con            *arm.Connection
	mockHost       string
)

func TestJobCollections_ListBySubscription(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_ListByResourceGroup(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_Get(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_CreateOrUpdate(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_Patch(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_Delete(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_Enable(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobCollections_Disable(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_Get(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_CreateOrUpdate(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_Patch(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_Delete(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_Run(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_List(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

func TestJobs_ListJobHistory(t *testing.T) {
	t.Skip("Warning: No test steps for this operation!")
}

// TestMain will exec each test
func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run() // exec test and this returns an exit code to pass to os
	tearDown()
	os.Exit(retCode)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func setUp() {
	ctx = context.Background()
	subscriptionId = getEnv("SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	mockHost = getEnv("AZURE_VIRTUAL_SERVER_HOST", "https://localhost:8443")

	tr := &http.Transport{}
	if err := http2.ConfigureTransport(tr); err != nil {
		fmt.Printf("Failed to configure http2 transport: %v", err)
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}
	cred, err = azidentity.NewEnvironmentCredential(&azidentity.EnvironmentCredentialOptions{AuthorityHost: mockHost, HTTPClient: client})
	if err != nil {
		panic(err)
	}

	con = arm.NewConnection(mockHost, cred, &arm.ConnectionOptions{
		Logging: policy.LogOptions{
			IncludeBody: true,
		},
		HTTPClient: client,
	})
}

func tearDown() {

}
