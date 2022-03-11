# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*ClientListOperationsPager.NextPage` return value(s) have been changed from `(bool)` to `(ClientListOperationsResponse, error)`
- Function `*InstancesClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(InstancesClientListByResourceGroupResponse, error)`
- Function `*InstancesClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(InstancesClientListResponse, error)`
- Function `*InstancesClientListByResourceGroupPager.PageResponse` has been removed
- Function `*ClientListOperationsPager.Err` has been removed
- Function `*InstancesClientListPager.PageResponse` has been removed
- Function `*ClientListOperationsPager.PageResponse` has been removed
- Function `*InstancesClientListByResourceGroupPager.Err` has been removed
- Function `*InstancesClientListPager.Err` has been removed
- Struct `ClientListOperationsResult` has been removed
- Struct `InstancesClientCheckNameAvailabilityResult` has been removed
- Struct `InstancesClientCreateResult` has been removed
- Struct `InstancesClientGetDetailsResult` has been removed
- Struct `InstancesClientListByResourceGroupResult` has been removed
- Struct `InstancesClientListResult` has been removed
- Struct `InstancesClientUpdateResult` has been removed
- Field `InstancesClientGetDetailsResult` of struct `InstancesClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `InstancesClientGetDetailsResponse` has been removed
- Field `ClientListOperationsResult` of struct `ClientListOperationsResponse` has been removed
- Field `RawResponse` of struct `ClientListOperationsResponse` has been removed
- Field `RawResponse` of struct `InstancesClientUpdatePollerResponse` has been removed
- Field `InstancesClientCheckNameAvailabilityResult` of struct `InstancesClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `InstancesClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `InstancesClientDeleteResponse` has been removed
- Field `RawResponse` of struct `InstancesClientCreatePollerResponse` has been removed
- Field `RawResponse` of struct `InstancesClientDeletePollerResponse` has been removed
- Field `InstancesClientUpdateResult` of struct `InstancesClientUpdateResponse` has been removed
- Field `RawResponse` of struct `InstancesClientUpdateResponse` has been removed
- Field `InstancesClientListByResourceGroupResult` of struct `InstancesClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `InstancesClientListByResourceGroupResponse` has been removed
- Field `InstancesClientCreateResult` of struct `InstancesClientCreateResponse` has been removed
- Field `RawResponse` of struct `InstancesClientCreateResponse` has been removed
- Field `InstancesClientListResult` of struct `InstancesClientListResponse` has been removed
- Field `RawResponse` of struct `InstancesClientListResponse` has been removed

### Features Added

- New function `*InstancesClientListPager.More() bool`
- New function `*InstancesClientListByResourceGroupPager.More() bool`
- New function `*ClientListOperationsPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `OperationDisplayValue`
- New anonymous field `Instance` in struct `InstancesClientGetDetailsResponse`
- New anonymous field `Instances` in struct `InstancesClientListResponse`
- New anonymous field `OperationListResult` in struct `ClientListOperationsResponse`
- New anonymous field `Instances` in struct `InstancesClientListByResourceGroupResponse`
- New anonymous field `CheckInstanceNameAvailabilityResult` in struct `InstancesClientCheckNameAvailabilityResponse`
- New anonymous field `Instance` in struct `InstancesClientCreateResponse`
- New anonymous field `Instance` in struct `InstancesClientUpdateResponse`


## 0.1.0 (2022-03-10)

- Init release.