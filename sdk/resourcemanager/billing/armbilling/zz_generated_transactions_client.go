//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// TransactionsClient contains the methods for the Transactions group.
// Don't use this type directly, use NewTransactionsClient() instead.
type TransactionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewTransactionsClient creates a new instance of TransactionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTransactionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TransactionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TransactionsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// ListByInvoice - Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage charges.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - The ID that uniquely identifies a billing account.
// invoiceName - The ID that uniquely identifies an invoice.
// options - TransactionsClientListByInvoiceOptions contains the optional parameters for the TransactionsClient.ListByInvoice
// method.
func (client *TransactionsClient) ListByInvoice(billingAccountName string, invoiceName string, options *TransactionsClientListByInvoiceOptions) *TransactionsClientListByInvoicePager {
	return &TransactionsClientListByInvoicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInvoiceCreateRequest(ctx, billingAccountName, invoiceName, options)
		},
		advancer: func(ctx context.Context, resp TransactionsClientListByInvoiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TransactionListResult.NextLink)
		},
	}
}

// listByInvoiceCreateRequest creates the ListByInvoice request.
func (client *TransactionsClient) listByInvoiceCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientListByInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/transactions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInvoiceHandleResponse handles the ListByInvoice response.
func (client *TransactionsClient) listByInvoiceHandleResponse(resp *http.Response) (TransactionsClientListByInvoiceResponse, error) {
	result := TransactionsClientListByInvoiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionListResult); err != nil {
		return TransactionsClientListByInvoiceResponse{}, err
	}
	return result, nil
}
