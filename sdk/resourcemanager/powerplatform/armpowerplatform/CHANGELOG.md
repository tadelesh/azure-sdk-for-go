# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*EnterprisePoliciesClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(EnterprisePoliciesClientListByResourceGroupResponse, error)`
- Function `*PrivateLinkResourcesClient.ListByEnterprisePolicy` parameter(s) have been changed from `(context.Context, string, string, *PrivateLinkResourcesClientListByEnterprisePolicyOptions)` to `(string, string, *PrivateLinkResourcesClientListByEnterprisePolicyOptions)`
- Function `*PrivateLinkResourcesClient.ListByEnterprisePolicy` return value(s) have been changed from `(PrivateLinkResourcesClientListByEnterprisePolicyResponse, error)` to `(*PrivateLinkResourcesClientListByEnterprisePolicyPager)`
- Function `*PrivateEndpointConnectionsClient.ListByEnterprisePolicy` parameter(s) have been changed from `(context.Context, string, string, *PrivateEndpointConnectionsClientListByEnterprisePolicyOptions)` to `(string, string, *PrivateEndpointConnectionsClientListByEnterprisePolicyOptions)`
- Function `*PrivateEndpointConnectionsClient.ListByEnterprisePolicy` return value(s) have been changed from `(PrivateEndpointConnectionsClientListByEnterprisePolicyResponse, error)` to `(*PrivateEndpointConnectionsClientListByEnterprisePolicyPager)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*EnterprisePoliciesClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(EnterprisePoliciesClientListBySubscriptionResponse, error)`
- Function `*AccountsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListByResourceGroupResponse, error)`
- Function `*AccountsClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListBySubscriptionResponse, error)`
- Function `*AccountsClientListBySubscriptionPager.Err` has been removed
- Function `*AccountsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*AccountsClientListBySubscriptionPager.PageResponse` has been removed
- Function `*EnterprisePoliciesClientListBySubscriptionPager.PageResponse` has been removed
- Function `*EnterprisePoliciesClientListBySubscriptionPager.Err` has been removed
- Function `*AccountsClientListByResourceGroupPager.Err` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*EnterprisePoliciesClientListByResourceGroupPager.Err` has been removed
- Function `*EnterprisePoliciesClientListByResourceGroupPager.PageResponse` has been removed
- Struct `AccountsClientCreateOrUpdateResult` has been removed
- Struct `AccountsClientGetResult` has been removed
- Struct `AccountsClientListByResourceGroupResult` has been removed
- Struct `AccountsClientListBySubscriptionResult` has been removed
- Struct `AccountsClientUpdateResult` has been removed
- Struct `EnterprisePoliciesClientCreateOrUpdateResult` has been removed
- Struct `EnterprisePoliciesClientGetResult` has been removed
- Struct `EnterprisePoliciesClientListByResourceGroupResult` has been removed
- Struct `EnterprisePoliciesClientListBySubscriptionResult` has been removed
- Struct `EnterprisePoliciesClientUpdateResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `PrivateEndpointConnectionsClientCreateOrUpdateResult` has been removed
- Struct `PrivateEndpointConnectionsClientGetResult` has been removed
- Struct `PrivateEndpointConnectionsClientListByEnterprisePolicyResult` has been removed
- Struct `PrivateLinkResourcesClientGetResult` has been removed
- Struct `PrivateLinkResourcesClientListByEnterprisePolicyResult` has been removed
- Field `AccountsClientGetResult` of struct `AccountsClientGetResponse` has been removed
- Field `RawResponse` of struct `AccountsClientGetResponse` has been removed
- Field `AccountsClientListByResourceGroupResult` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `EnterprisePoliciesClientGetResult` of struct `EnterprisePoliciesClientGetResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientGetResponse` has been removed
- Field `EnterprisePoliciesClientUpdateResult` of struct `EnterprisePoliciesClientUpdateResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientUpdateResponse` has been removed
- Field `PrivateEndpointConnectionsClientCreateOrUpdateResult` of struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientDeleteResponse` has been removed
- Field `PrivateLinkResourcesClientListByEnterprisePolicyResult` of struct `PrivateLinkResourcesClientListByEnterprisePolicyResponse` has been removed
- Field `RawResponse` of struct `PrivateLinkResourcesClientListByEnterprisePolicyResponse` has been removed
- Field `PrivateEndpointConnectionsClientGetResult` of struct `PrivateEndpointConnectionsClientGetResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientGetResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `PrivateLinkResourcesClientGetResult` of struct `PrivateLinkResourcesClientGetResponse` has been removed
- Field `RawResponse` of struct `PrivateLinkResourcesClientGetResponse` has been removed
- Field `EnterprisePoliciesClientListBySubscriptionResult` of struct `EnterprisePoliciesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientDeletePollerResponse` has been removed
- Field `PrivateEndpointConnectionsClientListByEnterprisePolicyResult` of struct `PrivateEndpointConnectionsClientListByEnterprisePolicyResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientListByEnterprisePolicyResponse` has been removed
- Field `AccountsClientCreateOrUpdateResult` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse` has been removed
- Field `EnterprisePoliciesClientCreateOrUpdateResult` of struct `EnterprisePoliciesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientCreateOrUpdateResponse` has been removed
- Field `EnterprisePoliciesClientListByResourceGroupResult` of struct `EnterprisePoliciesClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `EnterprisePoliciesClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AccountsClientDeleteResponse` has been removed
- Field `AccountsClientUpdateResult` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientUpdateResponse` has been removed
- Field `AccountsClientListBySubscriptionResult` of struct `AccountsClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListBySubscriptionResponse` has been removed

### Features Added

- New function `*EnterprisePoliciesClientListBySubscriptionPager.More() bool`
- New function `*PrivateEndpointConnectionsClientListByEnterprisePolicyPager.NextPage(context.Context) (PrivateEndpointConnectionsClientListByEnterprisePolicyResponse, error)`
- New function `*PrivateLinkResourcesClientListByEnterprisePolicyPager.NextPage(context.Context) (PrivateLinkResourcesClientListByEnterprisePolicyResponse, error)`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New function `*AccountsClientListByResourceGroupPager.More() bool`
- New function `*EnterprisePoliciesClientListByResourceGroupPager.More() bool`
- New function `*PrivateEndpointConnectionsClientListByEnterprisePolicyPager.More() bool`
- New function `*AccountsClientListBySubscriptionPager.More() bool`
- New function `*PrivateLinkResourcesClientListByEnterprisePolicyPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `PrivateEndpointConnectionsClientListByEnterprisePolicyPager`
- New struct `PrivateLinkResourcesClientListByEnterprisePolicyPager`
- New anonymous field `Account` in struct `AccountsClientGetResponse`
- New anonymous field `EnterprisePolicy` in struct `EnterprisePoliciesClientUpdateResponse`
- New anonymous field `Account` in struct `AccountsClientCreateOrUpdateResponse`
- New anonymous field `PrivateEndpointConnection` in struct `PrivateEndpointConnectionsClientGetResponse`
- New anonymous field `EnterprisePolicy` in struct `EnterprisePoliciesClientGetResponse`
- New anonymous field `AccountList` in struct `AccountsClientListBySubscriptionResponse`
- New anonymous field `EnterprisePolicy` in struct `EnterprisePoliciesClientCreateOrUpdateResponse`
- New anonymous field `PrivateLinkResource` in struct `PrivateLinkResourcesClientGetResponse`
- New anonymous field `AccountList` in struct `AccountsClientListByResourceGroupResponse`
- New anonymous field `Account` in struct `AccountsClientUpdateResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `PrivateEndpointConnection` in struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New anonymous field `EnterprisePolicyList` in struct `EnterprisePoliciesClientListBySubscriptionResponse`
- New anonymous field `EnterprisePolicyList` in struct `EnterprisePoliciesClientListByResourceGroupResponse`
- New anonymous field `PrivateEndpointConnectionListResult` in struct `PrivateEndpointConnectionsClientListByEnterprisePolicyResponse`
- New anonymous field `PrivateLinkResourceListResult` in struct `PrivateLinkResourcesClientListByEnterprisePolicyResponse`


## 0.1.0 (2022-03-10)

- Init release.