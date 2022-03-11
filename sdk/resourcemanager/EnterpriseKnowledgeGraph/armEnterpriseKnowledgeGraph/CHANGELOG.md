# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*ClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(ClientListResponse, error)`
- Function `*ClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(ClientListByResourceGroupResponse, error)`
- Function `*ClientListPager.PageResponse` has been removed
- Function `*ClientListByResourceGroupPager.Err` has been removed
- Function `*ClientListByResourceGroupPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*ClientListPager.Err` has been removed
- Struct `ClientCreateResult` has been removed
- Struct `ClientGetResult` has been removed
- Struct `ClientListByResourceGroupResult` has been removed
- Struct `ClientListResult` has been removed
- Struct `ClientUpdateResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `ClientListByResourceGroupResult` of struct `ClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `ClientListByResourceGroupResponse` has been removed
- Field `ClientListResult` of struct `ClientListResponse` has been removed
- Field `RawResponse` of struct `ClientListResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `ClientDeleteResponse` has been removed
- Field `ClientGetResult` of struct `ClientGetResponse` has been removed
- Field `RawResponse` of struct `ClientGetResponse` has been removed
- Field `ClientUpdateResult` of struct `ClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ClientUpdateResponse` has been removed
- Field `ClientCreateResult` of struct `ClientCreateResponse` has been removed
- Field `RawResponse` of struct `ClientCreateResponse` has been removed

### Features Added

- New function `*ClientListPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*ClientListByResourceGroupPager.More() bool`
- New struct `Error`
- New struct `ErrorBody`
- New anonymous field `ResponseList` in struct `ClientListByResourceGroupResponse`
- New anonymous field `EnterpriseKnowledgeGraph` in struct `ClientCreateResponse`
- New anonymous field `EnterpriseKnowledgeGraph` in struct `ClientUpdateResponse`
- New anonymous field `OperationEntityListResult` in struct `OperationsClientListResponse`
- New anonymous field `ResponseList` in struct `ClientListResponse`
- New anonymous field `EnterpriseKnowledgeGraph` in struct `ClientGetResponse`


## 0.1.0 (2022-03-10)

- Init release.