# Release History

## 0.4.0 (2022-03-10)
### Breaking Changes

- Function `*JobsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(JobsClientListByResourceGroupResponse, error)`
- Function `*LocationsClient.List` parameter(s) have been changed from `(context.Context, *LocationsClientListOptions)` to `(*LocationsClientListOptions)`
- Function `*LocationsClient.List` return value(s) have been changed from `(LocationsClientListResponse, error)` to `(*LocationsClientListPager)`
- Function `*JobsClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(JobsClientListBySubscriptionResponse, error)`
- Function `*BitLockerKeysClient.List` parameter(s) have been changed from `(context.Context, string, string, *BitLockerKeysClientListOptions)` to `(string, string, *BitLockerKeysClientListOptions)`
- Function `*BitLockerKeysClient.List` return value(s) have been changed from `(BitLockerKeysClientListResponse, error)` to `(*BitLockerKeysClientListPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*OperationsClientListPager)`
- Function `*JobsClientListByResourceGroupPager.Err` has been removed
- Function `*JobsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*JobsClientListBySubscriptionPager.Err` has been removed
- Function `*JobsClientListBySubscriptionPager.PageResponse` has been removed
- Struct `BitLockerKeysClientListResult` has been removed
- Struct `JobsClientCreateResult` has been removed
- Struct `JobsClientGetResult` has been removed
- Struct `JobsClientListByResourceGroupResult` has been removed
- Struct `JobsClientListBySubscriptionResult` has been removed
- Struct `JobsClientUpdateResult` has been removed
- Struct `LocationsClientGetResult` has been removed
- Struct `LocationsClientListResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `JobsClientCreateResult` of struct `JobsClientCreateResponse` has been removed
- Field `RawResponse` of struct `JobsClientCreateResponse` has been removed
- Field `LocationsClientListResult` of struct `LocationsClientListResponse` has been removed
- Field `RawResponse` of struct `LocationsClientListResponse` has been removed
- Field `JobsClientListBySubscriptionResult` of struct `JobsClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `JobsClientListBySubscriptionResponse` has been removed
- Field `BitLockerKeysClientListResult` of struct `BitLockerKeysClientListResponse` has been removed
- Field `RawResponse` of struct `BitLockerKeysClientListResponse` has been removed
- Field `JobsClientListByResourceGroupResult` of struct `JobsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `JobsClientListByResourceGroupResponse` has been removed
- Field `JobsClientUpdateResult` of struct `JobsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `JobsClientUpdateResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `JobsClientDeleteResponse` has been removed
- Field `JobsClientGetResult` of struct `JobsClientGetResponse` has been removed
- Field `RawResponse` of struct `JobsClientGetResponse` has been removed
- Field `LocationsClientGetResult` of struct `LocationsClientGetResponse` has been removed
- Field `RawResponse` of struct `LocationsClientGetResponse` has been removed

### Features Added

- New function `*JobsClientListBySubscriptionPager.More() bool`
- New function `*LocationsClientListPager.NextPage(context.Context) (LocationsClientListResponse, error)`
- New function `*BitLockerKeysClientListPager.More() bool`
- New function `*BitLockerKeysClientListPager.NextPage(context.Context) (BitLockerKeysClientListResponse, error)`
- New function `ErrorResponseError.MarshalJSON() ([]byte, error)`
- New function `*JobsClientListByResourceGroupPager.More() bool`
- New function `*OperationsClientListPager.NextPage(context.Context) (OperationsClientListResponse, error)`
- New function `*LocationsClientListPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New struct `BitLockerKeysClientListPager`
- New struct `ErrorResponse`
- New struct `ErrorResponseError`
- New struct `ErrorResponseErrorDetailsItem`
- New struct `LocationsClientListPager`
- New struct `OperationsClientListPager`
- New anonymous field `ListOperationsResponse` in struct `OperationsClientListResponse`
- New anonymous field `ListJobsResponse` in struct `JobsClientListBySubscriptionResponse`
- New anonymous field `JobResponse` in struct `JobsClientGetResponse`
- New anonymous field `LocationsResponse` in struct `LocationsClientListResponse`
- New anonymous field `JobResponse` in struct `JobsClientUpdateResponse`
- New anonymous field `GetBitLockerKeysResponse` in struct `BitLockerKeysClientListResponse`
- New anonymous field `ListJobsResponse` in struct `JobsClientListByResourceGroupResponse`
- New anonymous field `Location` in struct `LocationsClientGetResponse`
- New anonymous field `JobResponse` in struct `JobsClientCreateResponse`


## 0.3.0 (2022-03-10)
### Breaking Changes

- Type of `JobResponse.Tags` has been changed from `map[string]interface{}` to `interface{}`
- Type of `UpdateJobParameters.Tags` has been changed from `map[string]interface{}` to `interface{}`
- Type of `PutJobParameters.Tags` has been changed from `map[string]interface{}` to `interface{}`
- Function `ErrorResponseError.MarshalJSON` has been removed
- Struct `ErrorResponse` has been removed
- Struct `ErrorResponseError` has been removed
- Struct `ErrorResponseErrorDetailsItem` has been removed

### Features Added



## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*JobsClient.Update` parameter(s) have been changed from `(context.Context, string, string, UpdateJobParameters, *JobsUpdateOptions)` to `(context.Context, string, string, UpdateJobParameters, *JobsClientUpdateOptions)`
- Function `*JobsClient.Update` return value(s) have been changed from `(JobsUpdateResponse, error)` to `(JobsClientUpdateResponse, error)`
- Function `*JobsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *JobsListByResourceGroupOptions)` to `(string, *JobsClientListByResourceGroupOptions)`
- Function `*JobsClient.ListByResourceGroup` return value(s) have been changed from `(*JobsListByResourceGroupPager)` to `(*JobsClientListByResourceGroupPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*JobsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *JobsGetOptions)` to `(context.Context, string, string, *JobsClientGetOptions)`
- Function `*JobsClient.Get` return value(s) have been changed from `(JobsGetResponse, error)` to `(JobsClientGetResponse, error)`
- Function `*LocationsClient.List` parameter(s) have been changed from `(context.Context, *LocationsListOptions)` to `(context.Context, *LocationsClientListOptions)`
- Function `*LocationsClient.List` return value(s) have been changed from `(LocationsListResponse, error)` to `(LocationsClientListResponse, error)`
- Function `*JobsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, *JobsDeleteOptions)` to `(context.Context, string, string, *JobsClientDeleteOptions)`
- Function `*JobsClient.Delete` return value(s) have been changed from `(JobsDeleteResponse, error)` to `(JobsClientDeleteResponse, error)`
- Function `*JobsClient.ListBySubscription` parameter(s) have been changed from `(*JobsListBySubscriptionOptions)` to `(*JobsClientListBySubscriptionOptions)`
- Function `*JobsClient.ListBySubscription` return value(s) have been changed from `(*JobsListBySubscriptionPager)` to `(*JobsClientListBySubscriptionPager)`
- Function `*JobsClient.Create` parameter(s) have been changed from `(context.Context, string, string, PutJobParameters, *JobsCreateOptions)` to `(context.Context, string, string, PutJobParameters, *JobsClientCreateOptions)`
- Function `*JobsClient.Create` return value(s) have been changed from `(JobsCreateResponse, error)` to `(JobsClientCreateResponse, error)`
- Function `*LocationsClient.Get` parameter(s) have been changed from `(context.Context, string, *LocationsGetOptions)` to `(context.Context, string, *LocationsClientGetOptions)`
- Function `*LocationsClient.Get` return value(s) have been changed from `(LocationsGetResponse, error)` to `(LocationsClientGetResponse, error)`
- Function `*BitLockerKeysClient.List` parameter(s) have been changed from `(context.Context, string, string, *BitLockerKeysListOptions)` to `(context.Context, string, string, *BitLockerKeysClientListOptions)`
- Function `*BitLockerKeysClient.List` return value(s) have been changed from `(BitLockerKeysListResponse, error)` to `(BitLockerKeysClientListResponse, error)`
- Function `*JobsListByResourceGroupPager.PageResponse` has been removed
- Function `*JobsListByResourceGroupPager.Err` has been removed
- Function `*JobsListByResourceGroupPager.NextPage` has been removed
- Function `*JobsListBySubscriptionPager.Err` has been removed
- Function `*JobsListBySubscriptionPager.NextPage` has been removed
- Function `*JobsListBySubscriptionPager.PageResponse` has been removed
- Function `ErrorResponse.Error` has been removed
- Struct `BitLockerKeysListOptions` has been removed
- Struct `BitLockerKeysListResponse` has been removed
- Struct `BitLockerKeysListResult` has been removed
- Struct `JobsCreateOptions` has been removed
- Struct `JobsCreateResponse` has been removed
- Struct `JobsCreateResult` has been removed
- Struct `JobsDeleteOptions` has been removed
- Struct `JobsDeleteResponse` has been removed
- Struct `JobsGetOptions` has been removed
- Struct `JobsGetResponse` has been removed
- Struct `JobsGetResult` has been removed
- Struct `JobsListByResourceGroupOptions` has been removed
- Struct `JobsListByResourceGroupPager` has been removed
- Struct `JobsListByResourceGroupResponse` has been removed
- Struct `JobsListByResourceGroupResult` has been removed
- Struct `JobsListBySubscriptionOptions` has been removed
- Struct `JobsListBySubscriptionPager` has been removed
- Struct `JobsListBySubscriptionResponse` has been removed
- Struct `JobsListBySubscriptionResult` has been removed
- Struct `JobsUpdateOptions` has been removed
- Struct `JobsUpdateResponse` has been removed
- Struct `JobsUpdateResult` has been removed
- Struct `LocationsGetOptions` has been removed
- Struct `LocationsGetResponse` has been removed
- Struct `LocationsGetResult` has been removed
- Struct `LocationsListOptions` has been removed
- Struct `LocationsListResponse` has been removed
- Struct `LocationsListResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed

### Features Added

- New function `*JobsClientListBySubscriptionPager.Err() error`
- New function `*JobsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*JobsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*JobsClientListByResourceGroupPager.PageResponse() JobsClientListByResourceGroupResponse`
- New function `*JobsClientListByResourceGroupPager.Err() error`
- New function `*JobsClientListBySubscriptionPager.PageResponse() JobsClientListBySubscriptionResponse`
- New struct `BitLockerKeysClientListOptions`
- New struct `BitLockerKeysClientListResponse`
- New struct `BitLockerKeysClientListResult`
- New struct `JobsClientCreateOptions`
- New struct `JobsClientCreateResponse`
- New struct `JobsClientCreateResult`
- New struct `JobsClientDeleteOptions`
- New struct `JobsClientDeleteResponse`
- New struct `JobsClientGetOptions`
- New struct `JobsClientGetResponse`
- New struct `JobsClientGetResult`
- New struct `JobsClientListByResourceGroupOptions`
- New struct `JobsClientListByResourceGroupPager`
- New struct `JobsClientListByResourceGroupResponse`
- New struct `JobsClientListByResourceGroupResult`
- New struct `JobsClientListBySubscriptionOptions`
- New struct `JobsClientListBySubscriptionPager`
- New struct `JobsClientListBySubscriptionResponse`
- New struct `JobsClientListBySubscriptionResult`
- New struct `JobsClientUpdateOptions`
- New struct `JobsClientUpdateResponse`
- New struct `JobsClientUpdateResult`
- New struct `LocationsClientGetOptions`
- New struct `LocationsClientGetResponse`
- New struct `LocationsClientGetResult`
- New struct `LocationsClientListOptions`
- New struct `LocationsClientListResponse`
- New struct `LocationsClientListResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New field `Error` in struct `ErrorResponse`


## 0.1.0 (2021-12-22)

- Init release.
