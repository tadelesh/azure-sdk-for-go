//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_ListByFactory.json
func ExampleTriggersClient_ListByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	pager := client.ListByFactory("<resource-group-name>",
		"<factory-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_QueryByFactory.json
func ExampleTriggersClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	res, err := client.QueryByFactory(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.TriggerFilterParameters{
			ParentTriggerName: to.StringPtr("<parent-trigger-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientQueryByFactoryResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
func ExampleTriggersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		armdatafactory.TriggerResource{
			Properties: &armdatafactory.ScheduleTrigger{
				Type: to.StringPtr("<type>"),
				Pipelines: []*armdatafactory.TriggerPipelineReference{
					{
						Parameters: map[string]interface{}{
							"OutputBlobNameList": []interface{}{
								"exampleoutput.csv",
							},
						},
						PipelineReference: &armdatafactory.PipelineReference{
							Type:          armdatafactory.PipelineReferenceType("PipelineReference").ToPtr(),
							ReferenceName: to.StringPtr("<reference-name>"),
						},
					}},
				TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
					Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
						EndTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.8441801Z"); return t }()),
						Frequency: armdatafactory.RecurrenceFrequency("Minute").ToPtr(),
						Interval:  to.Int32Ptr(4),
						StartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.8441801Z"); return t }()),
						TimeZone:  to.StringPtr("<time-zone>"),
					},
				},
			},
		},
		&armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Get.json
func ExampleTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		&armdatafactory.TriggersClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientGetResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Delete.json
func ExampleTriggersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_SubscribeToEvents.json
func ExampleTriggersClient_BeginSubscribeToEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginSubscribeToEvents(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientSubscribeToEventsResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_GetEventSubscriptionStatus.json
func ExampleTriggersClient_GetEventSubscriptionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	res, err := client.GetEventSubscriptionStatus(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientGetEventSubscriptionStatusResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_UnsubscribeFromEvents.json
func ExampleTriggersClient_BeginUnsubscribeFromEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUnsubscribeFromEvents(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientUnsubscribeFromEventsResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Start.json
func ExampleTriggersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Stop.json
func ExampleTriggersClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}