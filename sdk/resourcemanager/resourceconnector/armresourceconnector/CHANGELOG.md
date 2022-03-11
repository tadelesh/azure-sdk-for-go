# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*AppliancesClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(AppliancesClientListByResourceGroupResponse, error)`
- Function `*AppliancesClientListOperationsPager.NextPage` return value(s) have been changed from `(bool)` to `(AppliancesClientListOperationsResponse, error)`
- Function `*AppliancesClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(AppliancesClientListBySubscriptionResponse, error)`
- Function `*AppliancesClientListOperationsPager.PageResponse` has been removed
- Function `*AppliancesClientListBySubscriptionPager.Err` has been removed
- Function `*AppliancesClientListOperationsPager.Err` has been removed
- Function `*AppliancesClientListByResourceGroupPager.PageResponse` has been removed
- Function `*AppliancesClientListByResourceGroupPager.Err` has been removed
- Function `*AppliancesClientListBySubscriptionPager.PageResponse` has been removed
- Struct `AppliancesClientCreateOrUpdateResult` has been removed
- Struct `AppliancesClientGetResult` has been removed
- Struct `AppliancesClientListByResourceGroupResult` has been removed
- Struct `AppliancesClientListBySubscriptionResult` has been removed
- Struct `AppliancesClientListClusterUserCredentialResult` has been removed
- Struct `AppliancesClientListOperationsResult` has been removed
- Struct `AppliancesClientUpdateResult` has been removed
- Field `AppliancesClientListClusterUserCredentialResult` of struct `AppliancesClientListClusterUserCredentialResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientListClusterUserCredentialResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientDeletePollerResponse` has been removed
- Field `AppliancesClientGetResult` of struct `AppliancesClientGetResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientGetResponse` has been removed
- Field `AppliancesClientUpdateResult` of struct `AppliancesClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientUpdateResponse` has been removed
- Field `AppliancesClientListBySubscriptionResult` of struct `AppliancesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientDeleteResponse` has been removed
- Field `AppliancesClientListByResourceGroupResult` of struct `AppliancesClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientListByResourceGroupResponse` has been removed
- Field `AppliancesClientListOperationsResult` of struct `AppliancesClientListOperationsResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientListOperationsResponse` has been removed
- Field `AppliancesClientCreateOrUpdateResult` of struct `AppliancesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AppliancesClientCreateOrUpdatePollerResponse` has been removed

### Features Added

- New function `*AppliancesClientListBySubscriptionPager.More() bool`
- New function `*AppliancesClientListOperationsPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New function `*AppliancesClientListByResourceGroupPager.More() bool`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `ApplianceListResult` in struct `AppliancesClientListBySubscriptionResponse`
- New anonymous field `ApplianceListCredentialResults` in struct `AppliancesClientListClusterUserCredentialResponse`
- New anonymous field `ApplianceListResult` in struct `AppliancesClientListByResourceGroupResponse`
- New anonymous field `Appliance` in struct `AppliancesClientCreateOrUpdateResponse`
- New anonymous field `Appliance` in struct `AppliancesClientGetResponse`
- New anonymous field `Appliance` in struct `AppliancesClientUpdateResponse`
- New anonymous field `ApplianceOperationsList` in struct `AppliancesClientListOperationsResponse`


## 0.1.0 (2022-03-10)

- Init release.