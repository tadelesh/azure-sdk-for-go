//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommerce

const (
	moduleName    = "armcommerce"
	moduleVersion = "v0.1.0"
)

type AggregationGranularity string

const (
	AggregationGranularityDaily  AggregationGranularity = "Daily"
	AggregationGranularityHourly AggregationGranularity = "Hourly"
)

// PossibleAggregationGranularityValues returns the possible values for the AggregationGranularity const type.
func PossibleAggregationGranularityValues() []AggregationGranularity {
	return []AggregationGranularity{
		AggregationGranularityDaily,
		AggregationGranularityHourly,
	}
}

// ToPtr returns a *AggregationGranularity pointing to the current value.
func (c AggregationGranularity) ToPtr() *AggregationGranularity {
	return &c
}

// OfferTermInfo - Name of the offer term
type OfferTermInfo string

const (
	OfferTermInfoRecurringCharge    OfferTermInfo = "Recurring Charge"
	OfferTermInfoMonetaryCommitment OfferTermInfo = "Monetary Commitment"
	OfferTermInfoMonetaryCredit     OfferTermInfo = "Monetary Credit"
)

// PossibleOfferTermInfoValues returns the possible values for the OfferTermInfo const type.
func PossibleOfferTermInfoValues() []OfferTermInfo {
	return []OfferTermInfo{
		OfferTermInfoRecurringCharge,
		OfferTermInfoMonetaryCommitment,
		OfferTermInfoMonetaryCredit,
	}
}

// ToPtr returns a *OfferTermInfo pointing to the current value.
func (c OfferTermInfo) ToPtr() *OfferTermInfo {
	return &c
}
