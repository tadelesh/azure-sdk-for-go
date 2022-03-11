# Release History

## 0.3.0 (2022-03-10)
### Breaking Changes

- Function `*EnergyServicesClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(EnergyServicesClientListBySubscriptionResponse, error)`
- Function `*EnergyServicesClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(EnergyServicesClientListByResourceGroupResponse, error)`
- Function `*EnergyServicesClientListByResourceGroupPager.PageResponse` has been removed
- Function `*EnergyServicesClientListBySubscriptionPager.Err` has been removed
- Function `*EnergyServicesClientListByResourceGroupPager.Err` has been removed
- Function `*EnergyServicesClientListBySubscriptionPager.PageResponse` has been removed
- Struct `EnergyServicesClientCreateResult` has been removed
- Struct `EnergyServicesClientGetResult` has been removed
- Struct `EnergyServicesClientListByResourceGroupResult` has been removed
- Struct `EnergyServicesClientListBySubscriptionResult` has been removed
- Struct `EnergyServicesClientUpdateResult` has been removed
- Struct `LocationsClientCheckNameAvailabilityResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientCreatePollerResponse` has been removed
- Field `EnergyServicesClientGetResult` of struct `EnergyServicesClientGetResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientGetResponse` has been removed
- Field `EnergyServicesClientUpdateResult` of struct `EnergyServicesClientUpdateResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientUpdateResponse` has been removed
- Field `LocationsClientCheckNameAvailabilityResult` of struct `LocationsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `LocationsClientCheckNameAvailabilityResponse` has been removed
- Field `EnergyServicesClientListByResourceGroupResult` of struct `EnergyServicesClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientListByResourceGroupResponse` has been removed
- Field `EnergyServicesClientCreateResult` of struct `EnergyServicesClientCreateResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientCreateResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientDeletePollerResponse` has been removed
- Field `EnergyServicesClientListBySubscriptionResult` of struct `EnergyServicesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `EnergyServicesClientDeleteResponse` has been removed

### Features Added

- New function `*EnergyServicesClientListBySubscriptionPager.More() bool`
- New function `*EnergyServicesClientListByResourceGroupPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `CheckNameAvailabilityResponse` in struct `LocationsClientCheckNameAvailabilityResponse`
- New anonymous field `EnergyServiceList` in struct `EnergyServicesClientListBySubscriptionResponse`
- New anonymous field `EnergyServiceList` in struct `EnergyServicesClientListByResourceGroupResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `EnergyService` in struct `EnergyServicesClientGetResponse`
- New anonymous field `EnergyService` in struct `EnergyServicesClientUpdateResponse`
- New anonymous field `EnergyService` in struct `EnergyServicesClientCreateResponse`


## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `ErrorDetail.MarshalJSON` has been removed
- Struct `ErrorAdditionalInfo` has been removed
- Struct `ErrorDetail` has been removed
- Struct `ErrorResponse` has been removed

### Features Added



## 0.1.0 (2022-01-14)

- Init release.