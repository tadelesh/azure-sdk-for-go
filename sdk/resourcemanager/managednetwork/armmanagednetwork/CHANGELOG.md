# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*ScopeAssignmentsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(ScopeAssignmentsClientListResponse, error)`
- Function `*PeeringPoliciesClientListByManagedNetworkPager.NextPage` return value(s) have been changed from `(bool)` to `(PeeringPoliciesClientListByManagedNetworkResponse, error)`
- Function `*ManagedNetworksClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(ManagedNetworksClientListBySubscriptionResponse, error)`
- Function `*ManagedNetworksClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(ManagedNetworksClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*GroupsClientListByManagedNetworkPager.NextPage` return value(s) have been changed from `(bool)` to `(GroupsClientListByManagedNetworkResponse, error)`
- Function `*ScopeAssignmentsClientListPager.Err` has been removed
- Function `*ManagedNetworksClientListByResourceGroupPager.Err` has been removed
- Function `*ScopeAssignmentsClientListPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*PeeringPoliciesClientListByManagedNetworkPager.Err` has been removed
- Function `*ManagedNetworksClientListBySubscriptionPager.Err` has been removed
- Function `*GroupsClientListByManagedNetworkPager.Err` has been removed
- Function `*ManagedNetworksClientListByResourceGroupPager.PageResponse` has been removed
- Function `*ManagedNetworksClientListBySubscriptionPager.PageResponse` has been removed
- Function `*PeeringPoliciesClientListByManagedNetworkPager.PageResponse` has been removed
- Function `*GroupsClientListByManagedNetworkPager.PageResponse` has been removed
- Struct `GroupsClientCreateOrUpdateResult` has been removed
- Struct `GroupsClientGetResult` has been removed
- Struct `GroupsClientListByManagedNetworkResult` has been removed
- Struct `ManagedNetworksClientCreateOrUpdateResult` has been removed
- Struct `ManagedNetworksClientGetResult` has been removed
- Struct `ManagedNetworksClientListByResourceGroupResult` has been removed
- Struct `ManagedNetworksClientListBySubscriptionResult` has been removed
- Struct `ManagedNetworksClientUpdateResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `PeeringPoliciesClientCreateOrUpdateResult` has been removed
- Struct `PeeringPoliciesClientGetResult` has been removed
- Struct `PeeringPoliciesClientListByManagedNetworkResult` has been removed
- Struct `ScopeAssignmentsClientCreateOrUpdateResult` has been removed
- Struct `ScopeAssignmentsClientGetResult` has been removed
- Struct `ScopeAssignmentsClientListResult` has been removed
- Field `GroupsClientListByManagedNetworkResult` of struct `GroupsClientListByManagedNetworkResponse` has been removed
- Field `RawResponse` of struct `GroupsClientListByManagedNetworkResponse` has been removed
- Field `ManagedNetworksClientListByResourceGroupResult` of struct `ManagedNetworksClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientListByResourceGroupResponse` has been removed
- Field `PeeringPoliciesClientListByManagedNetworkResult` of struct `PeeringPoliciesClientListByManagedNetworkResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientListByManagedNetworkResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientDeleteResponse` has been removed
- Field `PeeringPoliciesClientGetResult` of struct `PeeringPoliciesClientGetResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientGetResponse` has been removed
- Field `RawResponse` of struct `GroupsClientDeleteResponse` has been removed
- Field `ManagedNetworksClientListBySubscriptionResult` of struct `ManagedNetworksClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientListBySubscriptionResponse` has been removed
- Field `ScopeAssignmentsClientCreateOrUpdateResult` of struct `ScopeAssignmentsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ScopeAssignmentsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `GroupsClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientCreateOrUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `ScopeAssignmentsClientDeleteResponse` has been removed
- Field `ScopeAssignmentsClientGetResult` of struct `ScopeAssignmentsClientGetResponse` has been removed
- Field `RawResponse` of struct `ScopeAssignmentsClientGetResponse` has been removed
- Field `ManagedNetworksClientUpdateResult` of struct `ManagedNetworksClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientUpdateResponse` has been removed
- Field `ManagedNetworksClientCreateOrUpdateResult` of struct `ManagedNetworksClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientCreateOrUpdateResponse` has been removed
- Field `GroupsClientCreateOrUpdateResult` of struct `GroupsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `GroupsClientCreateOrUpdateResponse` has been removed
- Field `GroupsClientGetResult` of struct `GroupsClientGetResponse` has been removed
- Field `RawResponse` of struct `GroupsClientGetResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientDeleteResponse` has been removed
- Field `RawResponse` of struct `GroupsClientDeletePollerResponse` has been removed
- Field `PeeringPoliciesClientCreateOrUpdateResult` of struct `PeeringPoliciesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientCreateOrUpdateResponse` has been removed
- Field `ScopeAssignmentsClientListResult` of struct `ScopeAssignmentsClientListResponse` has been removed
- Field `RawResponse` of struct `ScopeAssignmentsClientListResponse` has been removed
- Field `ManagedNetworksClientGetResult` of struct `ManagedNetworksClientGetResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientGetResponse` has been removed
- Field `RawResponse` of struct `ManagedNetworksClientUpdatePollerResponse` has been removed
- Field `RawResponse` of struct `PeeringPoliciesClientDeletePollerResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed

### Features Added

- New function `*ManagedNetworksClientListBySubscriptionPager.More() bool`
- New function `*GroupsClientListByManagedNetworkPager.More() bool`
- New function `*ScopeAssignmentsClientListPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*PeeringPoliciesClientListByManagedNetworkPager.More() bool`
- New function `*ManagedNetworksClientListByResourceGroupPager.More() bool`
- New struct `ErrorResponse`
- New anonymous field `ListResult` in struct `ManagedNetworksClientListByResourceGroupResponse`
- New anonymous field `ScopeAssignment` in struct `ScopeAssignmentsClientCreateOrUpdateResponse`
- New anonymous field `PeeringPolicy` in struct `PeeringPoliciesClientCreateOrUpdateResponse`
- New anonymous field `PeeringPolicy` in struct `PeeringPoliciesClientGetResponse`
- New anonymous field `ManagedNetwork` in struct `ManagedNetworksClientGetResponse`
- New anonymous field `ScopeAssignment` in struct `ScopeAssignmentsClientGetResponse`
- New anonymous field `PeeringPolicyListResult` in struct `PeeringPoliciesClientListByManagedNetworkResponse`
- New anonymous field `ManagedNetwork` in struct `ManagedNetworksClientCreateOrUpdateResponse`
- New anonymous field `ListResult` in struct `ManagedNetworksClientListBySubscriptionResponse`
- New anonymous field `GroupListResult` in struct `GroupsClientListByManagedNetworkResponse`
- New anonymous field `Group` in struct `GroupsClientCreateOrUpdateResponse`
- New anonymous field `ScopeAssignmentListResult` in struct `ScopeAssignmentsClientListResponse`
- New anonymous field `Group` in struct `GroupsClientGetResponse`
- New anonymous field `ManagedNetwork` in struct `ManagedNetworksClientUpdateResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`


## 0.1.0 (2022-03-10)

- Init release.