# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*UserClassicAccountsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(UserClassicAccountsClientListResponse, error)`
- Function `*AccountsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListByResourceGroupResponse, error)`
- Function `*AccountsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*AccountsClientListPager.Err` has been removed
- Function `*AccountsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*UserClassicAccountsClientListPager.PageResponse` has been removed
- Function `*AccountsClientListByResourceGroupPager.Err` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*UserClassicAccountsClientListPager.Err` has been removed
- Function `*AccountsClientListPager.PageResponse` has been removed
- Struct `AccountsClientCheckNameAvailabilityResult` has been removed
- Struct `AccountsClientCreateOrUpdateResult` has been removed
- Struct `AccountsClientGetResult` has been removed
- Struct `AccountsClientListByResourceGroupResult` has been removed
- Struct `AccountsClientListResult` has been removed
- Struct `AccountsClientUpdateResult` has been removed
- Struct `ClassicAccountsClientGetDetailsResult` has been removed
- Struct `GenerateClientAccessTokenResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `UserClassicAccountsClientListResult` has been removed
- Field `AccountsClientCheckNameAvailabilityResult` of struct `AccountsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `AccountsClientCheckNameAvailabilityResponse` has been removed
- Field `AccountsClientCreateOrUpdateResult` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `AccountsClientListByResourceGroupResult` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `GenerateClientAccessTokenResult` of struct `GenerateClientAccessTokenResponse` has been removed
- Field `RawResponse` of struct `GenerateClientAccessTokenResponse` has been removed
- Field `UserClassicAccountsClientListResult` of struct `UserClassicAccountsClientListResponse` has been removed
- Field `RawResponse` of struct `UserClassicAccountsClientListResponse` has been removed
- Field `ClassicAccountsClientGetDetailsResult` of struct `ClassicAccountsClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `ClassicAccountsClientGetDetailsResponse` has been removed
- Field `AccountsClientUpdateResult` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientDeleteResponse` has been removed
- Field `AccountsClientGetResult` of struct `AccountsClientGetResponse` has been removed
- Field `RawResponse` of struct `AccountsClientGetResponse` has been removed
- Field `AccountsClientListResult` of struct `AccountsClientListResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed

### Features Added

- New function `*AccountsClientListPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*UserClassicAccountsClientListPager.More() bool`
- New function `*AccountsClientListByResourceGroupPager.More() bool`
- New function `ErrorDefinition.MarshalJSON() ([]byte, error)`
- New struct `ErrorDefinition`
- New struct `ErrorResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `Account` in struct `AccountsClientGetResponse`
- New anonymous field `CheckNameAvailabilityResult` in struct `AccountsClientCheckNameAvailabilityResponse`
- New anonymous field `AccountList` in struct `AccountsClientListResponse`
- New anonymous field `Account` in struct `AccountsClientUpdateResponse`
- New anonymous field `AccountList` in struct `AccountsClientListByResourceGroupResponse`
- New anonymous field `AccessToken` in struct `GenerateClientAccessTokenResponse`
- New anonymous field `Account` in struct `AccountsClientCreateOrUpdateResponse`
- New anonymous field `ClassicAccount` in struct `ClassicAccountsClientGetDetailsResponse`
- New anonymous field `UserClassicAccountList` in struct `UserClassicAccountsClientListResponse`


## 0.1.0 (2022-03-10)

- Init release.