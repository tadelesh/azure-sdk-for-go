# Release History

## 0.4.0 (2022-03-10)
### Breaking Changes

- Function `*ChangesClientListChangesBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(ChangesClientListChangesBySubscriptionResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*ResourceChangesClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(ResourceChangesClientListResponse, error)`
- Function `*ChangesClientListChangesByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(ChangesClientListChangesByResourceGroupResponse, error)`
- Function `*ChangesClientListChangesByResourceGroupPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*ChangesClientListChangesByResourceGroupPager.Err` has been removed
- Function `*ChangesClientListChangesBySubscriptionPager.Err` has been removed
- Function `*ChangesClientListChangesBySubscriptionPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*ResourceChangesClientListPager.PageResponse` has been removed
- Function `*ResourceChangesClientListPager.Err` has been removed
- Struct `ChangesClientListChangesByResourceGroupResult` has been removed
- Struct `ChangesClientListChangesBySubscriptionResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `ResourceChangesClientListResult` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `ChangesClientListChangesByResourceGroupResult` of struct `ChangesClientListChangesByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `ChangesClientListChangesByResourceGroupResponse` has been removed
- Field `ChangesClientListChangesBySubscriptionResult` of struct `ChangesClientListChangesBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `ChangesClientListChangesBySubscriptionResponse` has been removed
- Field `ResourceChangesClientListResult` of struct `ResourceChangesClientListResponse` has been removed
- Field `RawResponse` of struct `ResourceChangesClientListResponse` has been removed

### Features Added

- New function `*OperationsClientListPager.More() bool`
- New function `*ChangesClientListChangesBySubscriptionPager.More() bool`
- New function `*ResourceChangesClientListPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New function `*ChangesClientListChangesByResourceGroupPager.More() bool`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `ChangeList` in struct `ChangesClientListChangesByResourceGroupResponse`
- New anonymous field `ChangeList` in struct `ResourceChangesClientListResponse`
- New anonymous field `ResourceProviderOperationList` in struct `OperationsClientListResponse`
- New anonymous field `ChangeList` in struct `ChangesClientListChangesBySubscriptionResponse`


## 0.3.0 (2022-03-10)
### Breaking Changes

- Function `ErrorDetail.MarshalJSON` has been removed
- Struct `ErrorAdditionalInfo` has been removed
- Struct `ErrorDetail` has been removed
- Struct `ErrorResponse` has been removed

### Features Added



## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ChangesClient.ListChangesByResourceGroup` parameter(s) have been changed from `(string, time.Time, time.Time, *ChangesListChangesByResourceGroupOptions)` to `(string, time.Time, time.Time, *ChangesClientListChangesByResourceGroupOptions)`
- Function `*ChangesClient.ListChangesByResourceGroup` return value(s) have been changed from `(*ChangesListChangesByResourceGroupPager)` to `(*ChangesClientListChangesByResourceGroupPager)`
- Function `*ResourceChangesClient.List` parameter(s) have been changed from `(string, time.Time, time.Time, *ResourceChangesListOptions)` to `(string, time.Time, time.Time, *ResourceChangesClientListOptions)`
- Function `*ResourceChangesClient.List` return value(s) have been changed from `(*ResourceChangesListPager)` to `(*ResourceChangesClientListPager)`
- Function `*ChangesClient.ListChangesBySubscription` parameter(s) have been changed from `(time.Time, time.Time, *ChangesListChangesBySubscriptionOptions)` to `(time.Time, time.Time, *ChangesClientListChangesBySubscriptionOptions)`
- Function `*ChangesClient.ListChangesBySubscription` return value(s) have been changed from `(*ChangesListChangesBySubscriptionPager)` to `(*ChangesClientListChangesBySubscriptionPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*ResourceChangesListPager.Err` has been removed
- Function `*ChangesListChangesByResourceGroupPager.NextPage` has been removed
- Function `*ChangesListChangesBySubscriptionPager.NextPage` has been removed
- Function `*ChangesListChangesBySubscriptionPager.PageResponse` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*ChangesListChangesBySubscriptionPager.Err` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*ChangesListChangesByResourceGroupPager.Err` has been removed
- Function `*ChangesListChangesByResourceGroupPager.PageResponse` has been removed
- Function `*ResourceChangesListPager.PageResponse` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `*ResourceChangesListPager.NextPage` has been removed
- Function `*OperationsListPager.Err` has been removed
- Struct `ChangesListChangesByResourceGroupOptions` has been removed
- Struct `ChangesListChangesByResourceGroupPager` has been removed
- Struct `ChangesListChangesByResourceGroupResponse` has been removed
- Struct `ChangesListChangesByResourceGroupResult` has been removed
- Struct `ChangesListChangesBySubscriptionOptions` has been removed
- Struct `ChangesListChangesBySubscriptionPager` has been removed
- Struct `ChangesListChangesBySubscriptionResponse` has been removed
- Struct `ChangesListChangesBySubscriptionResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `ResourceChangesListOptions` has been removed
- Struct `ResourceChangesListPager` has been removed
- Struct `ResourceChangesListResponse` has been removed
- Struct `ResourceChangesListResult` has been removed
- Field `ProxyResource` of struct `Change` has been removed
- Field `Resource` of struct `ProxyResource` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed

### Features Added

- New function `*ChangesClientListChangesByResourceGroupPager.Err() error`
- New function `*ChangesClientListChangesBySubscriptionPager.NextPage(context.Context) bool`
- New function `*ResourceChangesClientListPager.NextPage(context.Context) bool`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*ResourceChangesClientListPager.Err() error`
- New function `*ChangesClientListChangesByResourceGroupPager.NextPage(context.Context) bool`
- New function `*ChangesClientListChangesByResourceGroupPager.PageResponse() ChangesClientListChangesByResourceGroupResponse`
- New function `*ChangesClientListChangesBySubscriptionPager.Err() error`
- New function `*ResourceChangesClientListPager.PageResponse() ResourceChangesClientListResponse`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*OperationsClientListPager.Err() error`
- New function `*ChangesClientListChangesBySubscriptionPager.PageResponse() ChangesClientListChangesBySubscriptionResponse`
- New struct `ChangesClientListChangesByResourceGroupOptions`
- New struct `ChangesClientListChangesByResourceGroupPager`
- New struct `ChangesClientListChangesByResourceGroupResponse`
- New struct `ChangesClientListChangesByResourceGroupResult`
- New struct `ChangesClientListChangesBySubscriptionOptions`
- New struct `ChangesClientListChangesBySubscriptionPager`
- New struct `ChangesClientListChangesBySubscriptionResponse`
- New struct `ChangesClientListChangesBySubscriptionResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `ResourceChangesClientListOptions`
- New struct `ResourceChangesClientListPager`
- New struct `ResourceChangesClientListResponse`
- New struct `ResourceChangesClientListResult`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `ID` in struct `ProxyResource`
- New field `ID` in struct `Change`
- New field `Name` in struct `Change`
- New field `Type` in struct `Change`
- New field `Error` in struct `ErrorResponse`


## 0.1.0 (2021-12-01)

- Initial preview release.
