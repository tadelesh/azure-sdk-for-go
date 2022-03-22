//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

import "time"

// AcknowledgeOfferNotificationDetails - Notification update request payload details
type AcknowledgeOfferNotificationDetails struct {
	// Gets or sets a value indicating whether acknowledge action flag is enabled
	Acknowledge *bool `json:"acknowledge,omitempty"`

	// Gets or sets added plans
	AddPlans []*string `json:"addPlans,omitempty"`

	// Gets or sets a value indicating whether dismiss action flag is enabled
	Dismiss *bool `json:"dismiss,omitempty"`

	// Gets or sets a value indicating whether remove offer action flag is enabled
	RemoveOffer *bool `json:"removeOffer,omitempty"`

	// Gets or sets remove plans
	RemovePlans []*string `json:"removePlans,omitempty"`
}

// AcknowledgeOfferNotificationProperties - Notification update request payload
type AcknowledgeOfferNotificationProperties struct {
	// Notification update request payload details
	Properties *AcknowledgeOfferNotificationDetails `json:"properties,omitempty"`
}

// AdminRequestApprovalProperties - Admin approval request resource properties
type AdminRequestApprovalProperties struct {
	// Gets or sets admin action
	AdminAction *AdminAction `json:"adminAction,omitempty"`

	// Gets or sets admin details
	Administrator *string `json:"administrator,omitempty"`

	// Gets or sets Approved plans ids, empty in case of rejected
	ApprovedPlans []*string `json:"approvedPlans,omitempty"`

	// Gets or sets list of associated collection ids
	CollectionIDs []*string `json:"collectionIds,omitempty"`

	// Gets or sets admin comment
	Comment *string `json:"comment,omitempty"`

	// Gets or sets offer Id
	OfferID *string `json:"offerId,omitempty"`

	// Gets or sets publisher Id
	PublisherID *string `json:"publisherId,omitempty"`

	// READ-ONLY; Gets display name
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; Gets list of plans with requesters details
	Plans []*PlanRequesterDetails `json:"plans,omitempty" azure:"ro"`
}

// AdminRequestApprovalsList - List of admin request approval resources
type AdminRequestApprovalsList struct {
	Value []*AdminRequestApprovalsResource `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of notifications list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// AdminRequestApprovalsResource - Admin request approval resource.
type AdminRequestApprovalsResource struct {
	// The privateStore admin Approval request data structure.
	Properties *AdminRequestApprovalProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// BillingAccountsResponse - Billing accounts response object
type BillingAccountsResponse struct {
	// Billing accounts list
	BillingAccounts []*string `json:"billingAccounts,omitempty"`
}

// BulkCollectionsDetails - Bulk collection details
type BulkCollectionsDetails struct {
	// Action to perform (For example: EnableCollections, DisableCollections)
	Action *string `json:"action,omitempty"`

	// collection ids list that the action is performed on
	CollectionIDs []*string `json:"collectionIds,omitempty"`
}

// BulkCollectionsPayload - Bulk collections action properties
type BulkCollectionsPayload struct {
	// bulk collections properties details
	Properties *BulkCollectionsDetails `json:"properties,omitempty"`
}

// BulkCollectionsResponse - The bulk collections response. The response contains two lists that indicate for each collection
// whether the operation succeeded or failed
type BulkCollectionsResponse struct {
	// Failed collections
	Failed []*CollectionsDetails `json:"failed,omitempty"`

	// Succeeded collections
	Succeeded []*CollectionsDetails `json:"succeeded,omitempty"`
}

// Collection - The Collection data structure.
type Collection struct {
	// The collection data structure.
	Properties *CollectionProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CollectionProperties - The collection details
type CollectionProperties struct {
	// Indicating whether all subscriptions are selected (=true) or not (=false).
	AllSubscriptions *bool `json:"allSubscriptions,omitempty"`

	// Gets or sets the association with Commercial's Billing Account.
	Claim *string `json:"claim,omitempty"`

	// Gets or sets collection name.
	CollectionName *string `json:"collectionName,omitempty"`

	// Indicating whether the collection is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Gets or sets subscription ids list. Empty list indicates all subscriptions are selected, null indicates no update is done,
	// explicit list indicates the explicit selected subscriptions. On insert, null
	// is considered as bad request
	SubscriptionsList []*string `json:"subscriptionsList,omitempty"`

	// READ-ONLY; Gets collection Id.
	CollectionID *string `json:"collectionId,omitempty" azure:"ro"`

	// READ-ONLY; Gets the number of offers associated with the collection.
	NumberOfOffers *int64 `json:"numberOfOffers,omitempty" azure:"ro"`
}

// CollectionsDetails - Collection name and id.
type CollectionsDetails struct {
	// Collection id.
	CollectionID *string `json:"collectionId,omitempty"`

	// Collection name.
	CollectionName *string `json:"collectionName,omitempty"`
}

type CollectionsList struct {
	// URL to get the next set of offer list results if there are any.
	NextLink *string       `json:"nextLink,omitempty"`
	Value    []*Collection `json:"value,omitempty"`
}

// CollectionsSubscriptionsMappingDetails - Collection name and related subscriptions list
type CollectionsSubscriptionsMappingDetails struct {
	// Collection name
	CollectionName *string `json:"collectionName,omitempty"`

	// Subscriptions ids list
	Subscriptions []*string `json:"subscriptions,omitempty"`
}

// CollectionsToSubscriptionsMappingPayload - The subscriptions list to get the related collections
type CollectionsToSubscriptionsMappingPayload struct {
	// Subscriptions ids list
	Properties *CollectionsToSubscriptionsMappingProperties `json:"properties,omitempty"`
}

// CollectionsToSubscriptionsMappingProperties - The subscriptions list to get the related collections
type CollectionsToSubscriptionsMappingProperties struct {
	// Subscriptions ids list
	SubscriptionIDs []*string `json:"subscriptionIds,omitempty"`
}

// CollectionsToSubscriptionsMappingResponse - A map of collections subscriptions details
type CollectionsToSubscriptionsMappingResponse struct {
	// The map of collections subscriptions
	Details map[string]*CollectionsSubscriptionsMappingDetails `json:"details,omitempty"`
}

// ErrorResponse - Error response indicates Microsoft.Marketplace service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// The details of the error.
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError - The details of the error.
type ErrorResponseError struct {
	// READ-ONLY; Error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// NewNotifications - New plans notification details
type NewNotifications struct {
	// Gets offer display name
	DisplayName *string `json:"displayName,omitempty"`

	// Gets or sets the icon url
	Icon *string `json:"icon,omitempty"`

	// Gets a value indicating whether future plans is enabled.
	IsFuturePlansEnabled *bool `json:"isFuturePlansEnabled,omitempty"`

	// Gets or sets the notification message id
	MessageCode *int64 `json:"messageCode,omitempty"`

	// Gets offer id
	OfferID *string `json:"offerId,omitempty"`

	// Gets or sets removed plans notifications
	Plans []*PlanNotificationDetails `json:"plans,omitempty"`
}

// NotificationsSettingsProperties - Describes the json payload for notifications settings
type NotificationsSettingsProperties struct {
	// Gets or sets list of notified recipients for new requests
	Recipients []*Recipient `json:"recipients,omitempty"`

	// Gets or sets whether to send email to all marketplace admins for new requests
	SendToAllMarketplaceAdmins *bool `json:"sendToAllMarketplaceAdmins,omitempty"`
}

// Offer - The privateStore offer data structure.
type Offer struct {
	// The privateStore offer data structure.
	Properties *OfferProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

type OfferListResponse struct {
	// URL to get the next set of offer list results if there are any.
	NextLink *string  `json:"nextLink,omitempty"`
	Value    []*Offer `json:"value,omitempty"`
}

type OfferProperties struct {
	// Identifier for purposes of race condition
	ETag *string `json:"eTag,omitempty"`

	// Icon File Uris
	IconFileUris map[string]*string `json:"iconFileUris,omitempty"`

	// Offer plans
	Plans []*Plan `json:"plans,omitempty"`

	// Plan ids limitation for this offer
	SpecificPlanIDsLimitation []*string `json:"specificPlanIdsLimitation,omitempty"`

	// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed
	// one in db, the offer would not be updated.
	UpdateSuppressedDueIdempotence *bool `json:"updateSuppressedDueIdempotence,omitempty"`

	// READ-ONLY; Private store offer creation date
	CreatedAt *string `json:"createdAt,omitempty" azure:"ro"`

	// READ-ONLY; Private store offer modification date
	ModifiedAt *string `json:"modifiedAt,omitempty" azure:"ro"`

	// READ-ONLY; It will be displayed prominently in the marketplace
	OfferDisplayName *string `json:"offerDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Private store unique id
	PrivateStoreID *string `json:"privateStoreId,omitempty" azure:"ro"`

	// READ-ONLY; Publisher name that will be displayed prominently in the marketplace
	PublisherDisplayName *string `json:"publisherDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Offers unique id
	UniqueOfferID *string `json:"uniqueOfferId,omitempty" azure:"ro"`
}

// OperationListResult - Result of the request to list Marketplace operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	// List of Microsoft.Marketplace operations supported by the Microsoft.Marketplace resource provider.
	Value []*SingleOperation `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

type Plan struct {
	// Plan accessibility
	Accessibility *Accessibility `json:"accessibility,omitempty"`

	// READ-ONLY; Alternative stack type
	AltStackReference *string `json:"altStackReference,omitempty" azure:"ro"`

	// READ-ONLY; Friendly name for the plan for display in the marketplace
	PlanDisplayName *string `json:"planDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Text identifier for this plan
	PlanID *string `json:"planId,omitempty" azure:"ro"`

	// READ-ONLY; Identifier for this plan
	SKUID *string `json:"skuId,omitempty" azure:"ro"`

	// READ-ONLY; Stack type (classic or arm)
	StackType *string `json:"stackType,omitempty" azure:"ro"`
}

// PlanDetails - Return plan with request details
type PlanDetails struct {
	// Gets or sets user's justification for the plan's request
	Justification *string `json:"justification,omitempty"`

	// Gets or sets Plan Id
	PlanID *string `json:"planId,omitempty"`

	// Gets or sets the subscription id that the user is requesting to add the plan to
	SubscriptionID *string `json:"subscriptionId,omitempty"`

	// Gets or sets the subscription name that the user is requesting to add the plan to
	SubscriptionName *string `json:"subscriptionName,omitempty"`

	// READ-ONLY; Gets request date
	RequestDate interface{} `json:"requestDate,omitempty" azure:"ro"`

	// READ-ONLY; Gets the plan status
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// PlanNotificationDetails - Plan notification details
type PlanNotificationDetails struct {
	// Gets or sets the plan display name
	PlanDisplayName *string `json:"planDisplayName,omitempty"`

	// Gets or sets the plan id
	PlanID *string `json:"planId,omitempty"`
}

// PlanRequesterDetails - Plan with requesters details
type PlanRequesterDetails struct {
	// READ-ONLY; Gets the plan display name
	PlanDisplayName *string `json:"planDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Gets the plan id
	PlanID *string `json:"planId,omitempty" azure:"ro"`

	// READ-ONLY; Gets requesters details list
	Requesters []*UserRequestDetails `json:"requesters,omitempty" azure:"ro"`
}

// PrivateStore - The PrivateStore data structure.
type PrivateStore struct {
	// The PrivateStore data structure.
	Properties *PrivateStoreProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateStoreClientAcknowledgeOfferNotificationOptions contains the optional parameters for the PrivateStoreClient.AcknowledgeOfferNotification
// method.
type PrivateStoreClientAcknowledgeOfferNotificationOptions struct {
	Payload *AcknowledgeOfferNotificationProperties
}

// PrivateStoreClientAdminRequestApprovalsListOptions contains the optional parameters for the PrivateStoreClient.AdminRequestApprovalsList
// method.
type PrivateStoreClientAdminRequestApprovalsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientBillingAccountsOptions contains the optional parameters for the PrivateStoreClient.BillingAccounts method.
type PrivateStoreClientBillingAccountsOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientBulkCollectionsActionOptions contains the optional parameters for the PrivateStoreClient.BulkCollectionsAction
// method.
type PrivateStoreClientBulkCollectionsActionOptions struct {
	Payload *BulkCollectionsPayload
}

// PrivateStoreClientCollectionsToSubscriptionsMappingOptions contains the optional parameters for the PrivateStoreClient.CollectionsToSubscriptionsMapping
// method.
type PrivateStoreClientCollectionsToSubscriptionsMappingOptions struct {
	Payload *CollectionsToSubscriptionsMappingPayload
}

// PrivateStoreClientCreateApprovalRequestOptions contains the optional parameters for the PrivateStoreClient.CreateApprovalRequest
// method.
type PrivateStoreClientCreateApprovalRequestOptions struct {
	Payload *RequestApprovalResource
}

// PrivateStoreClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreClient.CreateOrUpdate method.
type PrivateStoreClientCreateOrUpdateOptions struct {
	Payload *PrivateStore
}

// PrivateStoreClientDeleteOptions contains the optional parameters for the PrivateStoreClient.Delete method.
type PrivateStoreClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetAdminRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.GetAdminRequestApproval
// method.
type PrivateStoreClientGetAdminRequestApprovalOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetApprovalRequestsListOptions contains the optional parameters for the PrivateStoreClient.GetApprovalRequestsList
// method.
type PrivateStoreClientGetApprovalRequestsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetOptions contains the optional parameters for the PrivateStoreClient.Get method.
type PrivateStoreClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.GetRequestApproval
// method.
type PrivateStoreClientGetRequestApprovalOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientListOptions contains the optional parameters for the PrivateStoreClient.List method.
type PrivateStoreClientListOptions struct {
	// Determines if to use cache or DB for serving this request
	UseCache *string
}

// PrivateStoreClientQueryApprovedPlansOptions contains the optional parameters for the PrivateStoreClient.QueryApprovedPlans
// method.
type PrivateStoreClientQueryApprovedPlansOptions struct {
	Payload *QueryApprovedPlansPayload
}

// PrivateStoreClientQueryNotificationsStateOptions contains the optional parameters for the PrivateStoreClient.QueryNotificationsState
// method.
type PrivateStoreClientQueryNotificationsStateOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientQueryOffersOptions contains the optional parameters for the PrivateStoreClient.QueryOffers method.
type PrivateStoreClientQueryOffersOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientQueryRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.QueryRequestApproval
// method.
type PrivateStoreClientQueryRequestApprovalOptions struct {
	Payload *QueryRequestApprovalProperties
}

// PrivateStoreClientUpdateAdminRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.UpdateAdminRequestApproval
// method.
type PrivateStoreClientUpdateAdminRequestApprovalOptions struct {
	Payload *AdminRequestApprovalsResource
}

// PrivateStoreClientWithdrawPlanOptions contains the optional parameters for the PrivateStoreClient.WithdrawPlan method.
type PrivateStoreClientWithdrawPlanOptions struct {
	Payload *WithdrawProperties
}

// PrivateStoreCollectionClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreCollectionClient.CreateOrUpdate
// method.
type PrivateStoreCollectionClientCreateOrUpdateOptions struct {
	Payload *Collection
}

// PrivateStoreCollectionClientDeleteOptions contains the optional parameters for the PrivateStoreCollectionClient.Delete
// method.
type PrivateStoreCollectionClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientGetOptions contains the optional parameters for the PrivateStoreCollectionClient.Get method.
type PrivateStoreCollectionClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientListOptions contains the optional parameters for the PrivateStoreCollectionClient.List method.
type PrivateStoreCollectionClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientPostOptions contains the optional parameters for the PrivateStoreCollectionClient.Post method.
type PrivateStoreCollectionClientPostOptions struct {
	Payload *Operation
}

// PrivateStoreCollectionClientTransferOffersOptions contains the optional parameters for the PrivateStoreCollectionClient.TransferOffers
// method.
type PrivateStoreCollectionClientTransferOffersOptions struct {
	Payload *TransferOffersProperties
}

// PrivateStoreCollectionOfferClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.CreateOrUpdate
// method.
type PrivateStoreCollectionOfferClientCreateOrUpdateOptions struct {
	Payload *Offer
}

// PrivateStoreCollectionOfferClientDeleteOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Delete
// method.
type PrivateStoreCollectionOfferClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientGetOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Get
// method.
type PrivateStoreCollectionOfferClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientListOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.List
// method.
type PrivateStoreCollectionOfferClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientPostOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Post
// method.
type PrivateStoreCollectionOfferClientPostOptions struct {
	Payload *Operation
}

// PrivateStoreList - Describes the json payload for the list of available private stores (between zero and one, inclusive)
type PrivateStoreList struct {
	// URL to get the next set of PrivateStore list results if there are any.
	NextLink *string         `json:"nextLink,omitempty"`
	Value    []*PrivateStore `json:"value,omitempty"`
}

// PrivateStoreNotificationsState - Get private store notifications state
type PrivateStoreNotificationsState struct {
	ApprovalRequests      []*RequestApprovalsDetails `json:"approvalRequests,omitempty"`
	NewNotifications      []*NewNotifications        `json:"newNotifications,omitempty"`
	StopSellNotifications []*StopSellNotifications   `json:"stopSellNotifications,omitempty"`
}

// PrivateStoreProperties - Describes the json payload on whether or not the private store is enabled for a given tenant
type PrivateStoreProperties struct {
	// Indicates private store availability
	Availability *Availability `json:"availability,omitempty"`

	// Gets or sets list of branding characteristics
	Branding map[string]*string `json:"branding,omitempty"`

	// Identifier for purposes of race condition
	ETag *string `json:"eTag,omitempty"`

	// Is government
	IsGov *bool `json:"isGov,omitempty"`

	// Gets or sets notifications settings
	NotificationsSettings *NotificationsSettingsProperties `json:"notificationsSettings,omitempty"`

	// Private Store Name
	PrivateStoreName *string `json:"privateStoreName,omitempty"`

	// Tenant id
	TenantID *string `json:"tenantId,omitempty"`

	// READ-ONLY; Gets list of associated collection ids
	CollectionIDs []*string `json:"collectionIds,omitempty" azure:"ro"`

	// READ-ONLY; Private Store id
	PrivateStoreID *string `json:"privateStoreId,omitempty" azure:"ro"`
}

// QueryApprovedPlans - Query approved plans details
type QueryApprovedPlans struct {
	// Offer id
	OfferID *string `json:"offerId,omitempty"`

	// Offer plan ids
	PlanIDs []*string `json:"planIds,omitempty"`
}

// QueryApprovedPlansDetails - Query approved plans response
type QueryApprovedPlansDetails struct {
	// Indicates whether all subscriptions are approved for this plan
	AllSubscriptions *bool `json:"allSubscriptions,omitempty"`

	// Plan id
	PlanID *string `json:"planId,omitempty"`

	// Approved subscription ids list. In case all subscriptions are approved for a plan, allSubscriptions flag is true and list
	// is empty ( else flag is set to false). In case both subscriptions list is
	// empty and allSubscriptions flag is false, the plan is not approved for any subscription.
	SubscriptionIDs []*string `json:"subscriptionIds,omitempty"`
}

// QueryApprovedPlansPayload - Query approved plans payload
type QueryApprovedPlansPayload struct {
	// Query approved plans details
	Properties *QueryApprovedPlans `json:"properties,omitempty"`
}

// QueryApprovedPlansResponse - Query approved plans response
type QueryApprovedPlansResponse struct {
	// A list indicating for each plan which subscriptions are approved. Plan Id is unique
	Details []*QueryApprovedPlansDetails `json:"details,omitempty"`
}

// QueryOffers - List of offers
type QueryOffers struct {
	// URL to get the next set of PrivateStore list results if there are any.
	NextLink *string            `json:"nextLink,omitempty"`
	Value    []*OfferProperties `json:"value,omitempty"`
}

// QueryRequestApproval - Gets the request plans with indication on each plan whether is approved by the admin, has pending
// request or not requested yet
type QueryRequestApproval struct {
	// Gets or sets e-tag field
	Etag *string `json:"etag,omitempty"`

	// Gets or sets the notification message id
	MessageCode *int64 `json:"messageCode,omitempty"`

	// Gets or sets the plans details
	PlansDetails map[string]*PlanDetails `json:"plansDetails,omitempty"`

	// Gets or sets unique offer id.
	UniqueOfferID *string `json:"uniqueOfferId,omitempty"`
}

// QueryRequestApprovalProperties - The details to get the request plans statuses
type QueryRequestApprovalProperties struct {
	// The details to get the request plans statuses
	Properties *RequestDetails `json:"properties,omitempty"`
}

// Recipient - Describes the json payload for a notified recipient for new requests
type Recipient struct {
	// Principal ID
	PrincipalID *string `json:"principalId,omitempty"`

	// READ-ONLY; Display Name
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; Email Address
	EmailAddress *string `json:"emailAddress,omitempty" azure:"ro"`
}

// RequestApprovalProperties - Approval request resource properties
type RequestApprovalProperties struct {
	// Gets or sets the request approval message code
	MessageCode *int64 `json:"messageCode,omitempty"`

	// Gets or sets unique offer id.
	OfferID *string `json:"offerId,omitempty"`

	// Gets or sets the plans details
	PlansDetails []*PlanDetails `json:"plansDetails,omitempty"`

	// The offer's publisher id
	PublisherID *string `json:"publisherId,omitempty"`

	// READ-ONLY; Gets a value indicating whether the request is closed
	IsClosed *bool `json:"isClosed,omitempty" azure:"ro"`

	// READ-ONLY; Gets offer display name
	OfferDisplayName *string `json:"offerDisplayName,omitempty" azure:"ro"`
}

// RequestApprovalResource - Request approval resource.
type RequestApprovalResource struct {
	// The privateStore approval request data structure.
	Properties *RequestApprovalProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestApprovalsDetails - Request approvals details
type RequestApprovalsDetails struct {
	// Gets offer display name
	DisplayName *string `json:"displayName,omitempty"`

	// Gets or sets the icon url
	Icon *string `json:"icon,omitempty"`

	// Gets or sets the notification message id
	MessageCode *int64 `json:"messageCode,omitempty"`

	// Gets offer id
	OfferID *string `json:"offerId,omitempty"`

	// Gets or sets removed plans notifications
	Plans []*PlanNotificationDetails `json:"plans,omitempty"`

	// Gets or sets publisher id
	PublisherID *string `json:"publisherId,omitempty"`
}

// RequestApprovalsList - List of admin request approval resources
type RequestApprovalsList struct {
	Value []*RequestApprovalResource `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of notifications list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// RequestDetails - Request details needed to get the plans statuses
type RequestDetails struct {
	// Current plans list
	PlanIDs []*string `json:"planIds,omitempty"`

	// The offer's publisher id
	PublisherID *string `json:"publisherId,omitempty"`

	// Gets or sets the subscription id
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// Resource - An Azure resource.
type Resource struct {
	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SingleOperation - Microsoft.Marketplace REST API operation
type SingleOperation struct {
	// The object that represents the operation.
	Display *SingleOperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// Origin of the operation
	Origin *string `json:"origin,omitempty"`

	// Properties of the operation
	Properties interface{} `json:"properties,omitempty"`
}

// SingleOperationDisplay - The object that represents the operation.
type SingleOperationDisplay struct {
	// READ-ONLY; Friendly description for the operation,
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Operation type
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Service provider: Microsoft.Marketplace
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource on which the operation is performed
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// StopSellNotifications - Stop sell notification details
type StopSellNotifications struct {
	// Gets offer display name
	DisplayName *string `json:"displayName,omitempty"`

	// Gets or sets the icon url
	Icon *string `json:"icon,omitempty"`

	// Gets a value indicating whether entire offer is in stop sell or only few of its plans
	IsEntire *bool `json:"isEntire,omitempty"`

	// Gets or sets the notification message id
	MessageCode *int64 `json:"messageCode,omitempty"`

	// Gets offer id
	OfferID *string `json:"offerId,omitempty"`

	// Gets or sets removed plans notifications
	Plans []*PlanNotificationDetails `json:"plans,omitempty"`
}

// SystemData - Read only system data
type SystemData struct {
	// The timestamp of resource creation (UTC)
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource
	CreatedByType *IdentityType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource
	LastModifiedByType *IdentityType `json:"lastModifiedByType,omitempty"`
}

// TransferOffersDetails - Transfer offers response details
type TransferOffersDetails struct {
	// Offers ids list to transfer from source collection to target collection(s)
	OfferIDsList []*string `json:"offerIdsList,omitempty"`

	// Operation to perform (For example: Copy or Move)
	Operation *string `json:"operation,omitempty"`

	// Target collections ids
	TargetCollections []*string `json:"targetCollections,omitempty"`
}

// TransferOffersProperties - Transfer offers properties
type TransferOffersProperties struct {
	// transfer offers properties details
	Properties *TransferOffersDetails `json:"properties,omitempty"`
}

// TransferOffersResponse - The transfer items response. The response contains two lists that indicate for each collection
// whether the operation succeeded or failed
type TransferOffersResponse struct {
	// Failed collections
	Failed []*CollectionsDetails `json:"failed,omitempty"`

	// Succeeded collections
	Succeeded []*CollectionsDetails `json:"succeeded,omitempty"`
}

// UserRequestDetails - user request details
type UserRequestDetails struct {
	// Gets the subscription id that the user is requesting to add the plan to
	SubscriptionID *string `json:"subscriptionId,omitempty"`

	// Gets the subscription name that the user is requesting to add the plan to
	SubscriptionName *string `json:"subscriptionName,omitempty"`

	// READ-ONLY; Gets request date
	Date *string `json:"date,omitempty" azure:"ro"`

	// READ-ONLY; Gets justification
	Justification *string `json:"justification,omitempty" azure:"ro"`

	// READ-ONLY; Gets user id
	User *string `json:"user,omitempty" azure:"ro"`
}

// WithdrawDetails - Withdraw properties details
type WithdrawDetails struct {
	// Gets or sets Plan Id
	PlanID *string `json:"planId,omitempty"`

	// The offer's publisher id
	PublisherID *string `json:"publisherId,omitempty"`
}

// WithdrawProperties - Withdraw properties
type WithdrawProperties struct {
	// Withdraw properties details
	Properties *WithdrawDetails `json:"properties,omitempty"`
}
