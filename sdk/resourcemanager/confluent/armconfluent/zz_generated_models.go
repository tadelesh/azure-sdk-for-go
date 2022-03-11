//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import "time"

// AgreementProperties - Terms properties for Marketplace and Confluent.
type AgreementProperties struct {
	// If any version of the terms have been accepted, otherwise false.
	Accepted *bool `json:"accepted,omitempty"`

	// Link to HTML with Microsoft and Publisher terms.
	LicenseTextLink *string `json:"licenseTextLink,omitempty"`

	// Plan identifier string.
	Plan *string `json:"plan,omitempty"`

	// Link to the privacy policy of the publisher.
	PrivacyPolicyLink *string `json:"privacyPolicyLink,omitempty"`

	// Product identifier string.
	Product *string `json:"product,omitempty"`

	// Publisher identifier string.
	Publisher *string `json:"publisher,omitempty"`

	// Date and time in UTC of when the terms were accepted. This is empty if Accepted is false.
	RetrieveDatetime *time.Time `json:"retrieveDatetime,omitempty"`

	// Terms signature.
	Signature *string `json:"signature,omitempty"`
}

// AgreementResource - Agreement Terms definition
type AgreementResource struct {
	// Represents the properties of the resource.
	Properties *AgreementProperties `json:"properties,omitempty"`

	// READ-ONLY; The ARM id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the agreement.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the agreement.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AgreementResourceListResponse - Response of a list operation.
type AgreementResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*AgreementResource `json:"value,omitempty"`
}

// ErrorResponseBody - Response body of Error
type ErrorResponseBody struct {
	// READ-ONLY; Error code
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error detail
	Details []*ErrorResponseBody `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Error message
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Error target
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarketplaceAgreementsClientCreateOptions contains the optional parameters for the MarketplaceAgreementsClient.Create method.
type MarketplaceAgreementsClientCreateOptions struct {
	// Confluent Marketplace Agreement resource
	Body *AgreementResource
}

// MarketplaceAgreementsClientListOptions contains the optional parameters for the MarketplaceAgreementsClient.List method.
type MarketplaceAgreementsClientListOptions struct {
	// placeholder for future optional parameters
}

// OfferDetail - Confluent Offer detail
type OfferDetail struct {
	// REQUIRED; Offer Id
	ID *string `json:"id,omitempty"`

	// REQUIRED; Offer Plan Id
	PlanID *string `json:"planId,omitempty"`

	// REQUIRED; Offer Plan Name
	PlanName *string `json:"planName,omitempty"`

	// REQUIRED; Publisher Id
	PublisherID *string `json:"publisherId,omitempty"`

	// REQUIRED; Offer Plan Term unit
	TermUnit *string `json:"termUnit,omitempty"`

	// READ-ONLY; SaaS Offer Status
	Status *SaaSOfferStatus `json:"status,omitempty" azure:"ro"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write confluent'.
	Description *string `json:"description,omitempty"`

	// Operation type, e.g., read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Confluent
	Provider *string `json:"provider,omitempty"`

	// Type on which the operation is performed, e.g., 'clusters'.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of GET request to list Confluent operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Confluent operations supported by the Microsoft.Confluent provider.
	Value []*OperationResult `json:"value,omitempty"`
}

// OperationResult - An Confluent REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OrganizationClientBeginCreateOptions contains the optional parameters for the OrganizationClient.BeginCreate method.
type OrganizationClientBeginCreateOptions struct {
	// Organization resource model
	Body *OrganizationResource
}

// OrganizationClientBeginDeleteOptions contains the optional parameters for the OrganizationClient.BeginDelete method.
type OrganizationClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// OrganizationClientGetOptions contains the optional parameters for the OrganizationClient.Get method.
type OrganizationClientGetOptions struct {
	// placeholder for future optional parameters
}

// OrganizationClientListByResourceGroupOptions contains the optional parameters for the OrganizationClient.ListByResourceGroup
// method.
type OrganizationClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// OrganizationClientListBySubscriptionOptions contains the optional parameters for the OrganizationClient.ListBySubscription
// method.
type OrganizationClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// OrganizationClientUpdateOptions contains the optional parameters for the OrganizationClient.Update method.
type OrganizationClientUpdateOptions struct {
	// Updated Organization resource
	Body *OrganizationResourceUpdate
}

// OrganizationOperationsClientListOptions contains the optional parameters for the OrganizationOperationsClient.List method.
type OrganizationOperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrganizationResource - Organization resource.
type OrganizationResource struct {
	// REQUIRED; Organization resource properties
	Properties *OrganizationResourceProperties `json:"properties,omitempty"`

	// Location of Organization resource
	Location *string `json:"location,omitempty"`

	// Organization resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The ARM id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OrganizationResourceListResult - The response of a list operation.
type OrganizationResourceListResult struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Result of a list operation.
	Value []*OrganizationResource `json:"value,omitempty"`
}

// OrganizationResourceProperties - Organization resource property
type OrganizationResourceProperties struct {
	// REQUIRED; Confluent offer detail
	OfferDetail *OfferDetail `json:"offerDetail,omitempty"`

	// REQUIRED; Subscriber detail
	UserDetail *UserDetail `json:"userDetail,omitempty"`

	// READ-ONLY; The creation time of the resource.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Id of the Confluent organization.
	OrganizationID *string `json:"organizationId,omitempty" azure:"ro"`

	// READ-ONLY; Provision states for confluent RP
	ProvisioningState *ProvisionState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; SSO url for the Confluent organization.
	SsoURL *string `json:"ssoUrl,omitempty" azure:"ro"`
}

// OrganizationResourceUpdate - Organization Resource update
type OrganizationResourceUpdate struct {
	// ARM resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// ResourceProviderDefaultErrorResponse - Default error response for resource provider
type ResourceProviderDefaultErrorResponse struct {
	// READ-ONLY; Response body of Error
	Error *ErrorResponseBody `json:"error,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// UserDetail - Subscriber detail
type UserDetail struct {
	// REQUIRED; Email address
	EmailAddress *string `json:"emailAddress,omitempty"`

	// First name
	FirstName *string `json:"firstName,omitempty"`

	// Last name
	LastName *string `json:"lastName,omitempty"`
}

// ValidationsClientValidateOrganizationOptions contains the optional parameters for the ValidationsClient.ValidateOrganization
// method.
type ValidationsClientValidateOrganizationOptions struct {
	// placeholder for future optional parameters
}
