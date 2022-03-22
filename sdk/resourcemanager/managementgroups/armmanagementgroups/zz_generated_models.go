//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import "time"

// APIClientCheckNameAvailabilityOptions contains the optional parameters for the APIClient.CheckNameAvailability method.
type APIClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// APIClientStartTenantBackfillOptions contains the optional parameters for the APIClient.StartTenantBackfill method.
type APIClientStartTenantBackfillOptions struct {
	// placeholder for future optional parameters
}

// APIClientTenantBackfillStatusOptions contains the optional parameters for the APIClient.TenantBackfillStatus method.
type APIClientTenantBackfillStatusOptions struct {
	// placeholder for future optional parameters
}

// AzureAsyncOperationResults - The results of Azure-AsyncOperation.
type AzureAsyncOperationResults struct {
	// The generic properties of a management group.
	Properties *ManagementGroupInfoProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The current status of the asynchronous operation performed . For example, Running, Succeeded, Failed
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CheckNameAvailabilityRequest - Management group name availability check parameters.
type CheckNameAvailabilityRequest struct {
	// the name to check for availability
	Name *string `json:"name,omitempty"`

	// fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResult - Describes the result of the request to check management group name availability.
type CheckNameAvailabilityResult struct {
	// READ-ONLY; Required if nameAvailable == false. Localized. If reason == invalid, provide the user with the reason why the
	// given name is invalid, and provide the resource naming requirements so that the user can
	// select a valid name. If reason == AlreadyExists, explain that is already in use, and direct them to select a different
	// name.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Required. True indicates name is valid and available. False indicates the name is invalid, unavailable, or both.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; Required if nameAvailable == false. Invalid indicates the name provided does not match the resource provider's
	// naming requirements (incorrect length, unsupported characters, etc.) AlreadyExists
	// indicates that the name is already in use and is therefore unavailable.
	Reason *Reason `json:"reason,omitempty" azure:"ro"`
}

// ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
type ClientBeginCreateOrUpdateOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
type ClientBeginDeleteOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientGetDescendantsOptions contains the optional parameters for the Client.GetDescendants method.
type ClientGetDescendantsOptions struct {
	// Page continuation token is only used if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a token
	// parameter that specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// Number of elements to return when retrieving results. Passing this in will override $skipToken.
	Top *int32
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
	// The $expand=children query string parameter allows clients to request inclusion of children in the response payload. $expand=path
	// includes the path from the root group to the current group.
	// $expand=ancestors includes the ancestor Ids of the current group.
	Expand *Enum0
	// A filter which allows the exclusion of subscriptions from results (i.e. '$filter=children.childType ne Subscription')
	Filter *string
	// The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload.
	// Note that $expand=children must be passed up if $recurse is set to true.
	Recurse *bool
}

// ClientListOptions contains the optional parameters for the Client.List method.
type ClientListOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
	// Page continuation token is only used if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a token
	// parameter that specifies a starting point to use for subsequent calls.
	Skiptoken *string
}

// ClientUpdateOptions contains the optional parameters for the Client.Update method.
type ClientUpdateOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
}

// CreateManagementGroupChildInfo - The child information of a management group used during creation.
type CreateManagementGroupChildInfo struct {
	// READ-ONLY; The list of children.
	Children []*CreateManagementGroupChildInfo `json:"children,omitempty" azure:"ro"`

	// READ-ONLY; The friendly name of the child resource.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The fully qualified ID for the child resource (management group or subscription). For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the child entity.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The fully qualified resource type which includes provider namespace (e.g. Microsoft.Management/managementGroups)
	Type *ManagementGroupChildType `json:"type,omitempty" azure:"ro"`
}

// CreateManagementGroupDetails - The details of a management group used during creation.
type CreateManagementGroupDetails struct {
	// (Optional) The ID of the parent management group used during creation.
	Parent *CreateParentGroupInfo `json:"parent,omitempty"`

	// READ-ONLY; The identity of the principal or process that updated the object.
	UpdatedBy *string `json:"updatedBy,omitempty" azure:"ro"`

	// READ-ONLY; The date and time when this object was last updated.
	UpdatedTime *time.Time `json:"updatedTime,omitempty" azure:"ro"`

	// READ-ONLY; The version number of the object.
	Version *int32 `json:"version,omitempty" azure:"ro"`
}

// CreateManagementGroupProperties - The generic properties of a management group used during creation.
type CreateManagementGroupProperties struct {
	// The details of a management group used during creation.
	Details *CreateManagementGroupDetails `json:"details,omitempty"`

	// The friendly name of the management group. If no value is passed then this field will be set to the groupId.
	DisplayName *string `json:"displayName,omitempty"`

	// READ-ONLY; The list of children.
	Children []*CreateManagementGroupChildInfo `json:"children,omitempty" azure:"ro"`

	// READ-ONLY; The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// CreateManagementGroupRequest - Management group creation parameters.
type CreateManagementGroupRequest struct {
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty"`

	// The generic properties of a management group used during creation.
	Properties *CreateManagementGroupProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CreateOrUpdateSettingsProperties - The properties of the request to create or update Management Group settings
type CreateOrUpdateSettingsProperties struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `json:"defaultManagementGroup,omitempty"`

	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will
	// require Microsoft.Management/managementGroups/write action on the root
	// Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating
	// new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `json:"requireAuthorizationForGroupCreation,omitempty"`
}

// CreateOrUpdateSettingsRequest - Parameters for creating or updating Management Group settings
type CreateOrUpdateSettingsRequest struct {
	// The properties of the request to create or update Management Group settings
	Properties *CreateOrUpdateSettingsProperties `json:"properties,omitempty"`
}

// CreateParentGroupInfo - (Optional) The ID of the parent management group used during creation.
type CreateParentGroupInfo struct {
	// The fully qualified ID for the parent management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`

	// READ-ONLY; The friendly name of the parent management group.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The name of the parent management group
	Name *string `json:"name,omitempty" azure:"ro"`
}

// DescendantInfo - The descendant.
type DescendantInfo struct {
	// The generic properties of an descendant.
	Properties *DescendantInfoProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the descendant. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	// or /subscriptions/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the descendant. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups or /subscriptions
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DescendantInfoProperties - The generic properties of an descendant.
type DescendantInfoProperties struct {
	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The ID of the parent management group.
	Parent *DescendantParentGroupInfo `json:"parent,omitempty"`
}

// DescendantListResult - Describes the result of the request to view descendants.
type DescendantListResult struct {
	// The list of descendants.
	Value []*DescendantInfo `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// DescendantParentGroupInfo - The ID of the parent management group.
type DescendantParentGroupInfo struct {
	// The fully qualified ID for the parent management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
}

// EntitiesClientListOptions contains the optional parameters for the EntitiesClient.List method.
type EntitiesClientListOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
	// The filter parameter allows you to filter on the the name or display name fields. You can check for equality on the name
	// field (e.g. name eq '{entityName}') and you can check for substrings on either
	// the name or display name fields(e.g. contains(name, '{substringToSearch}'), contains(displayName, '{substringToSearch')).
	// Note that the '{entityName}' and '{substringToSearch}' fields are checked case
	// insensitively.
	Filter *string
	// A filter which allows the get entities call to focus on a particular group (i.e. "$filter=name eq 'groupName'")
	GroupName *string
	// The $search parameter is used in conjunction with the $filter parameter to return three different outputs depending on
	// the parameter passed in. With $search=AllowedParents the API will return the
	// entity info of all groups that the requested entity will be able to reparent to as determined by the user's permissions.
	// With $search=AllowedChildren the API will return the entity info of all
	// entities that can be added as children of the requested entity. With $search=ParentAndFirstLevelChildren the API will return
	// the parent and first level of children that the user has either direct
	// access to or indirect access via one of their descendants. With $search=ParentOnly the API will return only the group if
	// the user has access to at least one of the descendants of the group. With
	// $search=ChildrenOnly the API will return only the first level of children of the group entity info specified in $filter.
	// The user must have direct access to the children entities or one of it's
	// descendants for it to show up in the results.
	Search *Enum2
	// This parameter specifies the fields to include in the response. Can include any combination of Name,DisplayName,Type,ParentDisplayNameChain,ParentChain,
	// e.g.
	// '$select=Name,DisplayName,Type,ParentDisplayNameChain,ParentNameChain'. When specified the $select parameter can override
	// select in $skipToken.
	Select *string
	// Number of entities to skip over when retrieving results. Passing this in will override $skipToken.
	Skip *int32
	// Page continuation token is only used if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a token
	// parameter that specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// Number of elements to return when retrieving results. Passing this in will override $skipToken.
	Top *int32
	// The view parameter allows clients to filter the type of data that is returned by the getEntities call.
	View *Enum3
}

// EntityHierarchyItem - The management group details for the hierarchy view.
type EntityHierarchyItem struct {
	// The generic properties of a management group.
	Properties *EntityHierarchyItemProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// EntityHierarchyItemProperties - The generic properties of a management group.
type EntityHierarchyItemProperties struct {
	// The list of children.
	Children []*EntityHierarchyItem `json:"children,omitempty"`

	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The users specific permissions to this item.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// EntityInfo - The entity.
type EntityInfo struct {
	// The generic properties of an entity.
	Properties *EntityInfoProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the entity. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the entity. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// EntityInfoProperties - The generic properties of an entity.
type EntityInfoProperties struct {
	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The users specific permissions to this item.
	InheritedPermissions *Permissions `json:"inheritedPermissions,omitempty"`

	// Number of children is the number of Groups that are exactly one level underneath the current Group.
	NumberOfChildGroups *int32 `json:"numberOfChildGroups,omitempty"`

	// Number of children is the number of Groups and Subscriptions that are exactly one level underneath the current Group.
	NumberOfChildren *int32 `json:"numberOfChildren,omitempty"`

	// Number of Descendants
	NumberOfDescendants *int32 `json:"numberOfDescendants,omitempty"`

	// (Optional) The ID of the parent management group.
	Parent *EntityParentGroupInfo `json:"parent,omitempty"`

	// The parent display name chain from the root group to the immediate parent
	ParentDisplayNameChain []*string `json:"parentDisplayNameChain,omitempty"`

	// The parent name chain from the root group to the immediate parent
	ParentNameChain []*string `json:"parentNameChain,omitempty"`

	// The users specific permissions to this item.
	Permissions *Permissions `json:"permissions,omitempty"`

	// The AAD Tenant ID associated with the entity. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty"`
}

// EntityListResult - Describes the result of the request to view entities.
type EntityListResult struct {
	// The list of entities.
	Value []*EntityInfo `json:"value,omitempty"`

	// READ-ONLY; Total count of records that match the filter
	Count *int32 `json:"count,omitempty" azure:"ro"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// EntityParentGroupInfo - (Optional) The ID of the parent management group.
type EntityParentGroupInfo struct {
	// The fully qualified ID for the parent management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
}

// ErrorDetails - The details of the error.
type ErrorDetails struct {
	// One of a server-defined set of error codes.
	Code *string `json:"code,omitempty"`

	// A human-readable representation of the error's details.
	Details *string `json:"details,omitempty"`

	// A human-readable representation of the error.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse - The error object.
type ErrorResponse struct {
	// The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// HierarchySettings - Settings defined at the Management Group scope.
type HierarchySettings struct {
	// The generic properties of hierarchy settings.
	Properties *HierarchySettingsProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the settings object. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/settings/default.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the object. In this case, default.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups/settings.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// HierarchySettingsClientCreateOrUpdateOptions contains the optional parameters for the HierarchySettingsClient.CreateOrUpdate
// method.
type HierarchySettingsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// HierarchySettingsClientDeleteOptions contains the optional parameters for the HierarchySettingsClient.Delete method.
type HierarchySettingsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// HierarchySettingsClientGetOptions contains the optional parameters for the HierarchySettingsClient.Get method.
type HierarchySettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// HierarchySettingsClientListOptions contains the optional parameters for the HierarchySettingsClient.List method.
type HierarchySettingsClientListOptions struct {
	// placeholder for future optional parameters
}

// HierarchySettingsClientUpdateOptions contains the optional parameters for the HierarchySettingsClient.Update method.
type HierarchySettingsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// HierarchySettingsInfo - The hierarchy settings resource.
type HierarchySettingsInfo struct {
	// The generic properties of hierarchy settings.
	Properties *HierarchySettingsProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the settings object. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/settings/default.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the object. In this case, default.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups/settings.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// HierarchySettingsList - Lists all hierarchy settings.
type HierarchySettingsList struct {
	// The list of hierarchy settings.
	Value []*HierarchySettingsInfo `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// HierarchySettingsProperties - The generic properties of hierarchy settings.
type HierarchySettingsProperties struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `json:"defaultManagementGroup,omitempty"`

	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will
	// require Microsoft.Management/managementGroups/write action on the root
	// Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating
	// new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `json:"requireAuthorizationForGroupCreation,omitempty"`

	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty"`
}

// ListSubscriptionUnderManagementGroup - The details of all subscriptions under management group.
type ListSubscriptionUnderManagementGroup struct {
	// The list of subscriptions.
	Value []*SubscriptionUnderManagementGroup `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ManagementGroup - The management group details.
type ManagementGroup struct {
	// The generic properties of a management group.
	Properties *ManagementGroupProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ManagementGroupChildInfo - The child information of a management group.
type ManagementGroupChildInfo struct {
	// The list of children.
	Children []*ManagementGroupChildInfo `json:"children,omitempty"`

	// The friendly name of the child resource.
	DisplayName *string `json:"displayName,omitempty"`

	// The fully qualified ID for the child resource (management group or subscription). For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`

	// The name of the child entity.
	Name *string `json:"name,omitempty"`

	// The fully qualified resource type which includes provider namespace (e.g. Microsoft.Management/managementGroups)
	Type *ManagementGroupChildType `json:"type,omitempty"`
}

// ManagementGroupDetails - The details of a management group.
type ManagementGroupDetails struct {
	// The ancestors of the management group.
	ManagementGroupAncestors []*string `json:"managementGroupAncestors,omitempty"`

	// The ancestors of the management group displayed in reversed order, from immediate parent to the root.
	ManagementGroupAncestorsChain []*ManagementGroupPathElement `json:"managementGroupAncestorsChain,omitempty"`

	// (Optional) The ID of the parent management group.
	Parent *ParentGroupInfo `json:"parent,omitempty"`

	// The path from the root to the current group.
	Path []*ManagementGroupPathElement `json:"path,omitempty"`

	// The identity of the principal or process that updated the object.
	UpdatedBy *string `json:"updatedBy,omitempty"`

	// The date and time when this object was last updated.
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`

	// The version number of the object.
	Version *int32 `json:"version,omitempty"`
}

// ManagementGroupInfo - The management group resource.
type ManagementGroupInfo struct {
	// The generic properties of a management group.
	Properties *ManagementGroupInfoProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ManagementGroupInfoProperties - The generic properties of a management group.
type ManagementGroupInfoProperties struct {
	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty"`
}

// ManagementGroupListResult - Describes the result of the request to list management groups.
type ManagementGroupListResult struct {
	// The list of management groups.
	Value []*ManagementGroupInfo `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ManagementGroupPathElement - A path element of a management group ancestors.
type ManagementGroupPathElement struct {
	// The friendly name of the group.
	DisplayName *string `json:"displayName,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty"`
}

// ManagementGroupProperties - The generic properties of a management group.
type ManagementGroupProperties struct {
	// The list of children.
	Children []*ManagementGroupChildInfo `json:"children,omitempty"`

	// The details of a management group.
	Details *ManagementGroupDetails `json:"details,omitempty"`

	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty"`
}

// ManagementGroupSubscriptionsClientCreateOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Create
// method.
type ManagementGroupSubscriptionsClientCreateOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
}

// ManagementGroupSubscriptionsClientDeleteOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Delete
// method.
type ManagementGroupSubscriptionsClientDeleteOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
}

// ManagementGroupSubscriptionsClientGetSubscriptionOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.GetSubscription
// method.
type ManagementGroupSubscriptionsClientGetSubscriptionOptions struct {
	// Indicates whether the request should utilize any caches. Populate the header with 'no-cache' value to bypass existing caches.
	CacheControl *string
}

// ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions contains the optional parameters for the
// ManagementGroupSubscriptionsClient.GetSubscriptionsUnderManagementGroup method.
type ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions struct {
	// Page continuation token is only used if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a token
	// parameter that specifies a starting point to use for subsequent calls.
	Skiptoken *string
}

// Operation supported by the Microsoft.Management resource provider.
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplayProperties `json:"display,omitempty"`

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationDisplayProperties - The object that represents the operation.
type OperationDisplayProperties struct {
	// READ-ONLY; Operation description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The operation that can be performed.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The name of the provider.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - Describes the result of the request to list Microsoft.Management operations.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the Microsoft.Management resource provider.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationResults - The results of an asynchronous operation.
type OperationResults struct {
	// The generic properties of a management group.
	Properties *ManagementGroupInfoProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ParentGroupInfo - (Optional) The ID of the parent management group.
type ParentGroupInfo struct {
	// The friendly name of the parent management group.
	DisplayName *string `json:"displayName,omitempty"`

	// The fully qualified ID for the parent management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`

	// The name of the parent management group
	Name *string `json:"name,omitempty"`
}

// PatchManagementGroupRequest - Management group patch parameters.
type PatchManagementGroupRequest struct {
	// The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`

	// (Optional) The fully qualified ID for the parent management group. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	ParentGroupID *string `json:"parentGroupId,omitempty"`
}

// SubscriptionUnderManagementGroup - The details of subscription under management group.
type SubscriptionUnderManagementGroup struct {
	// The generic properties of subscription under a management group.
	Properties *SubscriptionUnderManagementGroupProperties `json:"properties,omitempty"`

	// READ-ONLY; The fully qualified ID for the subscription. For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/subscriptions/0000000-0000-0000-0000-000000000001
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. For example, Microsoft.Management/managementGroups/subscriptions
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubscriptionUnderManagementGroupProperties - The generic properties of subscription under a management group.
type SubscriptionUnderManagementGroupProperties struct {
	// The friendly name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`

	// The ID of the parent management group.
	Parent *DescendantParentGroupInfo `json:"parent,omitempty"`

	// The state of the subscription.
	State *string `json:"state,omitempty"`

	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant *string `json:"tenant,omitempty"`
}

// TenantBackfillStatusResult - The tenant backfill status
type TenantBackfillStatusResult struct {
	// READ-ONLY; The status of the Tenant Backfill
	Status *Status `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}
