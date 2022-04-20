//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/CreateOrUpdateVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	virtualMachineScaleSetVMRunCommandsClient, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponsePoller, err := virtualMachineScaleSetVMRunCommandsClient.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		armcompute.VirtualMachineRunCommand{
			Location: to.Ptr("<resource-location>"),
			Properties: &armcompute.VirtualMachineRunCommandProperties{
				AsyncExecution: to.Ptr(false),
				Parameters: []*armcompute.RunCommandInputParameter{
					{
						Name:  to.Ptr("<the-run-command-parameter-name.>"),
						Value: to.Ptr("<the-run-command-parameter-value.>"),
					},
					{
						Name:  to.Ptr("<the-run-command-parameter-name.>"),
						Value: to.Ptr("<the-run-command-parameter-value.>"),
					}},
				RunAsPassword: to.Ptr("<specifies-the-user-account-password-on-the-vm-when-executing-the-run-command.>"),
				RunAsUser:     to.Ptr("<specifies-the-user-account-on-the-vm-when-executing-the-run-command.>"),
				Source: &armcompute.VirtualMachineRunCommandScriptSource{
					Script: to.Ptr("<specifies-the-script-content-to-be-executed-on-the-vm.>"),
				},
				TimeoutInSeconds: to.Ptr[int32](3600),
			},
		},
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse, err := virtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = virtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/UpdateVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	virtualMachineScaleSetVMRunCommandsClient, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientUpdateResponsePoller, err := virtualMachineScaleSetVMRunCommandsClient.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		armcompute.VirtualMachineRunCommandUpdate{
			Properties: &armcompute.VirtualMachineRunCommandProperties{
				Source: &armcompute.VirtualMachineRunCommandScriptSource{
					Script: to.Ptr("<specifies-the-script-content-to-be-executed-on-the-vm.>"),
				},
			},
		},
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientUpdateResponse, err := virtualMachineScaleSetVMRunCommandsClientUpdateResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = virtualMachineScaleSetVMRunCommandsClientUpdateResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/DeleteVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	virtualMachineScaleSetVMRunCommandsClient, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientDeleteResponsePoller, err := virtualMachineScaleSetVMRunCommandsClient.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = virtualMachineScaleSetVMRunCommandsClientDeleteResponsePoller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/GetVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	virtualMachineScaleSetVMRunCommandsClient, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientGetResponse, err := virtualMachineScaleSetVMRunCommandsClient.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = virtualMachineScaleSetVMRunCommandsClientGetResponse
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/ListVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	virtualMachineScaleSetVMRunCommandsClient, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	virtualMachineScaleSetVMRunCommandsClientNewListPager := virtualMachineScaleSetVMRunCommandsClient.NewListPager("<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions{Expand: nil})
	for virtualMachineScaleSetVMRunCommandsClientNewListPager.More() {
		nextResult, err := virtualMachineScaleSetVMRunCommandsClientNewListPager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
