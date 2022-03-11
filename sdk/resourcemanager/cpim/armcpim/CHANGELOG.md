# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*B2CTenantsClient.ListBySubscription` parameter(s) have been changed from `(context.Context, *B2CTenantsClientListBySubscriptionOptions)` to `(*B2CTenantsClientListBySubscriptionOptions)`
- Function `*B2CTenantsClient.ListBySubscription` return value(s) have been changed from `(B2CTenantsClientListBySubscriptionResponse, error)` to `(*B2CTenantsClientListBySubscriptionPager)`
- Function `*B2CTenantsClient.ListByResourceGroup` parameter(s) have been changed from `(context.Context, string, *B2CTenantsClientListByResourceGroupOptions)` to `(string, *B2CTenantsClientListByResourceGroupOptions)`
- Function `*B2CTenantsClient.ListByResourceGroup` return value(s) have been changed from `(B2CTenantsClientListByResourceGroupResponse, error)` to `(*B2CTenantsClientListByResourceGroupPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*OperationsClientListPager)`
- Struct `B2CTenantsClientCheckNameAvailabilityResult` has been removed
- Struct `B2CTenantsClientCreateResult` has been removed
- Struct `B2CTenantsClientGetResult` has been removed
- Struct `B2CTenantsClientListByResourceGroupResult` has been removed
- Struct `B2CTenantsClientListBySubscriptionResult` has been removed
- Struct `B2CTenantsClientUpdateResult` has been removed
- Struct `OperationsClientGetAsyncStatusResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `RawResponse` of struct `B2CTenantsClientCreatePollerResponse` has been removed
- Field `B2CTenantsClientListBySubscriptionResult` of struct `B2CTenantsClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientListBySubscriptionResponse` has been removed
- Field `B2CTenantsClientUpdateResult` of struct `B2CTenantsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientUpdateResponse` has been removed
- Field `B2CTenantsClientCheckNameAvailabilityResult` of struct `B2CTenantsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientCheckNameAvailabilityResponse` has been removed
- Field `OperationsClientGetAsyncStatusResult` of struct `OperationsClientGetAsyncStatusResponse` has been removed
- Field `RawResponse` of struct `OperationsClientGetAsyncStatusResponse` has been removed
- Field `B2CTenantsClientCreateResult` of struct `B2CTenantsClientCreateResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientCreateResponse` has been removed
- Field `B2CTenantsClientListByResourceGroupResult` of struct `B2CTenantsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientListByResourceGroupResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `B2CTenantsClientGetResult` of struct `B2CTenantsClientGetResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientGetResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `B2CTenantsClientDeleteResponse` has been removed

### Features Added

- New function `*B2CTenantsClientListByResourceGroupPager.NextPage(context.Context) (B2CTenantsClientListByResourceGroupResponse, error)`
- New function `ErrorResponse.MarshalJSON() ([]byte, error)`
- New function `*OperationsClientListPager.NextPage(context.Context) (OperationsClientListResponse, error)`
- New function `*OperationsClientListPager.More() bool`
- New function `*B2CTenantsClientListByResourceGroupPager.More() bool`
- New function `*B2CTenantsClientListBySubscriptionPager.NextPage(context.Context) (B2CTenantsClientListBySubscriptionResponse, error)`
- New function `*B2CTenantsClientListBySubscriptionPager.More() bool`
- New struct `B2CTenantsClientListByResourceGroupPager`
- New struct `B2CTenantsClientListBySubscriptionPager`
- New struct `CloudError`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorResponse`
- New struct `OperationsClientListPager`
- New anonymous field `NameAvailabilityResponse` in struct `B2CTenantsClientCheckNameAvailabilityResponse`
- New anonymous field `B2CTenantResource` in struct `B2CTenantsClientGetResponse`
- New anonymous field `AsyncOperationStatus` in struct `OperationsClientGetAsyncStatusResponse`
- New anonymous field `B2CTenantResource` in struct `B2CTenantsClientUpdateResponse`
- New anonymous field `B2CTenantResourceList` in struct `B2CTenantsClientListByResourceGroupResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `B2CTenantResourceList` in struct `B2CTenantsClientListBySubscriptionResponse`
- New anonymous field `B2CTenantResource` in struct `B2CTenantsClientCreateResponse`


## 0.1.0 (2022-03-10)

- Init release.