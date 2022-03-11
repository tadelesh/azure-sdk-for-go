# Release History

## 0.4.0 (2022-03-10)
### Breaking Changes

- Function `*LoadTestsClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(LoadTestsClientListBySubscriptionResponse, error)`
- Function `*LoadTestsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(LoadTestsClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*LoadTestsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*LoadTestsClientListBySubscriptionPager.Err` has been removed
- Function `*LoadTestsClientListBySubscriptionPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*LoadTestsClientListByResourceGroupPager.Err` has been removed
- Struct `LoadTestsClientCreateOrUpdateResult` has been removed
- Struct `LoadTestsClientGetResult` has been removed
- Struct `LoadTestsClientListByResourceGroupResult` has been removed
- Struct `LoadTestsClientListBySubscriptionResult` has been removed
- Struct `LoadTestsClientUpdateResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `RawResponse` of struct `LoadTestsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientDeletePollerResponse` has been removed
- Field `LoadTestsClientListByResourceGroupResult` of struct `LoadTestsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientListByResourceGroupResponse` has been removed
- Field `LoadTestsClientUpdateResult` of struct `LoadTestsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientUpdateResponse` has been removed
- Field `LoadTestsClientListBySubscriptionResult` of struct `LoadTestsClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientListBySubscriptionResponse` has been removed
- Field `LoadTestsClientGetResult` of struct `LoadTestsClientGetResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientGetResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `LoadTestsClientCreateOrUpdateResult` of struct `LoadTestsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `LoadTestsClientCreateOrUpdateResponse` has been removed

### Features Added

- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New function `*OperationsClientListPager.More() bool`
- New function `*LoadTestsClientListByResourceGroupPager.More() bool`
- New function `*LoadTestsClientListBySubscriptionPager.More() bool`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `LoadTestResourcePageList` in struct `LoadTestsClientListByResourceGroupResponse`
- New anonymous field `LoadTestResource` in struct `LoadTestsClientUpdateResponse`
- New anonymous field `LoadTestResource` in struct `LoadTestsClientGetResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `LoadTestResource` in struct `LoadTestsClientCreateOrUpdateResponse`
- New anonymous field `LoadTestResourcePageList` in struct `LoadTestsClientListBySubscriptionResponse`


## 0.3.0 (2022-03-10)
### Breaking Changes

- Type of `LoadTestResourcePatchRequestBody.Tags` has been changed from `map[string]interface{}` to `interface{}`
- Function `ErrorDetail.MarshalJSON` has been removed
- Struct `ErrorAdditionalInfo` has been removed
- Struct `ErrorDetail` has been removed
- Struct `ErrorResponse` has been removed

### Features Added



## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*LoadTestsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *LoadTestsListByResourceGroupOptions)` to `(string, *LoadTestsClientListByResourceGroupOptions)`
- Function `*LoadTestsClient.ListByResourceGroup` return value(s) have been changed from `(*LoadTestsListByResourceGroupPager)` to `(*LoadTestsClientListByResourceGroupPager)`
- Function `*LoadTestsClient.ListBySubscription` parameter(s) have been changed from `(*LoadTestsListBySubscriptionOptions)` to `(*LoadTestsClientListBySubscriptionOptions)`
- Function `*LoadTestsClient.ListBySubscription` return value(s) have been changed from `(*LoadTestsListBySubscriptionPager)` to `(*LoadTestsClientListBySubscriptionPager)`
- Function `*LoadTestsClient.Update` parameter(s) have been changed from `(context.Context, string, string, LoadTestResourcePatchRequestBody, *LoadTestsUpdateOptions)` to `(context.Context, string, string, LoadTestResourcePatchRequestBody, *LoadTestsClientUpdateOptions)`
- Function `*LoadTestsClient.Update` return value(s) have been changed from `(LoadTestsUpdateResponse, error)` to `(LoadTestsClientUpdateResponse, error)`
- Function `*LoadTestsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *LoadTestsGetOptions)` to `(context.Context, string, string, *LoadTestsClientGetOptions)`
- Function `*LoadTestsClient.Get` return value(s) have been changed from `(LoadTestsGetResponse, error)` to `(LoadTestsClientGetResponse, error)`
- Function `*LoadTestsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *LoadTestsBeginDeleteOptions)` to `(context.Context, string, string, *LoadTestsClientBeginDeleteOptions)`
- Function `*LoadTestsClient.BeginDelete` return value(s) have been changed from `(LoadTestsDeletePollerResponse, error)` to `(LoadTestsClientDeletePollerResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*LoadTestsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, LoadTestResource, *LoadTestsCreateOrUpdateOptions)` to `(context.Context, string, string, LoadTestResource, *LoadTestsClientCreateOrUpdateOptions)`
- Function `*LoadTestsClient.CreateOrUpdate` return value(s) have been changed from `(LoadTestsCreateOrUpdateResponse, error)` to `(LoadTestsClientCreateOrUpdateResponse, error)`
- Function `*LoadTestsListByResourceGroupPager.Err` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `*LoadTestsListBySubscriptionPager.NextPage` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `*LoadTestsListByResourceGroupPager.NextPage` has been removed
- Function `*LoadTestsListByResourceGroupPager.PageResponse` has been removed
- Function `*LoadTestsDeletePollerResponse.Resume` has been removed
- Function `*LoadTestsDeletePoller.Done` has been removed
- Function `*LoadTestsListBySubscriptionPager.Err` has been removed
- Function `*LoadTestsListBySubscriptionPager.PageResponse` has been removed
- Function `*LoadTestsDeletePoller.FinalResponse` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `LoadTestsDeletePollerResponse.PollUntilDone` has been removed
- Function `*LoadTestsDeletePoller.Poll` has been removed
- Function `*LoadTestsDeletePoller.ResumeToken` has been removed
- Struct `LoadTestsBeginDeleteOptions` has been removed
- Struct `LoadTestsCreateOrUpdateOptions` has been removed
- Struct `LoadTestsCreateOrUpdateResponse` has been removed
- Struct `LoadTestsCreateOrUpdateResult` has been removed
- Struct `LoadTestsDeletePoller` has been removed
- Struct `LoadTestsDeletePollerResponse` has been removed
- Struct `LoadTestsDeleteResponse` has been removed
- Struct `LoadTestsGetOptions` has been removed
- Struct `LoadTestsGetResponse` has been removed
- Struct `LoadTestsGetResult` has been removed
- Struct `LoadTestsListByResourceGroupOptions` has been removed
- Struct `LoadTestsListByResourceGroupPager` has been removed
- Struct `LoadTestsListByResourceGroupResponse` has been removed
- Struct `LoadTestsListByResourceGroupResult` has been removed
- Struct `LoadTestsListBySubscriptionOptions` has been removed
- Struct `LoadTestsListBySubscriptionPager` has been removed
- Struct `LoadTestsListBySubscriptionResponse` has been removed
- Struct `LoadTestsListBySubscriptionResult` has been removed
- Struct `LoadTestsUpdateOptions` has been removed
- Struct `LoadTestsUpdateResponse` has been removed
- Struct `LoadTestsUpdateResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `TrackedResource` of struct `LoadTestResource` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed
- Field `Resource` of struct `TrackedResource` has been removed

### Features Added

- New function `LoadTestsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (LoadTestsClientDeleteResponse, error)`
- New function `*LoadTestsClientListByResourceGroupPager.Err() error`
- New function `*LoadTestsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*LoadTestsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*LoadTestsClientListBySubscriptionPager.Err() error`
- New function `*LoadTestsClientDeletePoller.Done() bool`
- New function `*LoadTestsClientListByResourceGroupPager.PageResponse() LoadTestsClientListByResourceGroupResponse`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*LoadTestsClientDeletePoller.ResumeToken() (string, error)`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*LoadTestsClientDeletePoller.FinalResponse(context.Context) (LoadTestsClientDeleteResponse, error)`
- New function `*LoadTestsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*LoadTestsClientDeletePollerResponse.Resume(context.Context, *LoadTestsClient, string) error`
- New function `*LoadTestsClientListBySubscriptionPager.PageResponse() LoadTestsClientListBySubscriptionResponse`
- New function `*OperationsClientListPager.Err() error`
- New struct `LoadTestsClientBeginDeleteOptions`
- New struct `LoadTestsClientCreateOrUpdateOptions`
- New struct `LoadTestsClientCreateOrUpdateResponse`
- New struct `LoadTestsClientCreateOrUpdateResult`
- New struct `LoadTestsClientDeletePoller`
- New struct `LoadTestsClientDeletePollerResponse`
- New struct `LoadTestsClientDeleteResponse`
- New struct `LoadTestsClientGetOptions`
- New struct `LoadTestsClientGetResponse`
- New struct `LoadTestsClientGetResult`
- New struct `LoadTestsClientListByResourceGroupOptions`
- New struct `LoadTestsClientListByResourceGroupPager`
- New struct `LoadTestsClientListByResourceGroupResponse`
- New struct `LoadTestsClientListByResourceGroupResult`
- New struct `LoadTestsClientListBySubscriptionOptions`
- New struct `LoadTestsClientListBySubscriptionPager`
- New struct `LoadTestsClientListBySubscriptionResponse`
- New struct `LoadTestsClientListBySubscriptionResult`
- New struct `LoadTestsClientUpdateOptions`
- New struct `LoadTestsClientUpdateResponse`
- New struct `LoadTestsClientUpdateResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `TrackedResource`
- New field `SystemData` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `Error` in struct `ErrorResponse`
- New field `Tags` in struct `LoadTestResource`
- New field `ID` in struct `LoadTestResource`
- New field `Name` in struct `LoadTestResource`
- New field `SystemData` in struct `LoadTestResource`
- New field `Type` in struct `LoadTestResource`
- New field `Location` in struct `LoadTestResource`


## 0.1.0 (2021-12-07)

- Init release.
