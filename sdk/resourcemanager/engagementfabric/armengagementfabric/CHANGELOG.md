# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*ChannelsClient.ListByAccount` parameter(s) have been changed from `(context.Context, string, string, *ChannelsClientListByAccountOptions)` to `(string, string, *ChannelsClientListByAccountOptions)`
- Function `*ChannelsClient.ListByAccount` return value(s) have been changed from `(ChannelsClientListByAccountResponse, error)` to `(*ChannelsClientListByAccountPager)`
- Function `*AccountsClient.List` parameter(s) have been changed from `(context.Context, *AccountsClientListOptions)` to `(*AccountsClientListOptions)`
- Function `*AccountsClient.List` return value(s) have been changed from `(AccountsClientListResponse, error)` to `(*AccountsClientListPager)`
- Function `*AccountsClient.ListByResourceGroup` parameter(s) have been changed from `(context.Context, string, *AccountsClientListByResourceGroupOptions)` to `(string, *AccountsClientListByResourceGroupOptions)`
- Function `*AccountsClient.ListByResourceGroup` return value(s) have been changed from `(AccountsClientListByResourceGroupResponse, error)` to `(*AccountsClientListByResourceGroupPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*OperationsClientListPager)`
- Function `*AccountsClient.ListKeys` parameter(s) have been changed from `(context.Context, string, string, *AccountsClientListKeysOptions)` to `(string, string, *AccountsClientListKeysOptions)`
- Function `*AccountsClient.ListKeys` return value(s) have been changed from `(AccountsClientListKeysResponse, error)` to `(*AccountsClientListKeysPager)`
- Function `*SKUsClient.List` parameter(s) have been changed from `(context.Context, *SKUsClientListOptions)` to `(*SKUsClientListOptions)`
- Function `*SKUsClient.List` return value(s) have been changed from `(SKUsClientListResponse, error)` to `(*SKUsClientListPager)`
- Struct `AccountsClientCreateOrUpdateResult` has been removed
- Struct `AccountsClientGetResult` has been removed
- Struct `AccountsClientListByResourceGroupResult` has been removed
- Struct `AccountsClientListChannelTypesResult` has been removed
- Struct `AccountsClientListKeysResult` has been removed
- Struct `AccountsClientListResult` has been removed
- Struct `AccountsClientRegenerateKeyResult` has been removed
- Struct `AccountsClientUpdateResult` has been removed
- Struct `ChannelsClientCreateOrUpdateResult` has been removed
- Struct `ChannelsClientGetResult` has been removed
- Struct `ChannelsClientListByAccountResult` has been removed
- Struct `ClientCheckNameAvailabilityResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `SKUsClientListResult` has been removed
- Field `ChannelsClientCreateOrUpdateResult` of struct `ChannelsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ChannelsClientCreateOrUpdateResponse` has been removed
- Field `AccountsClientCreateOrUpdateResult` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `AccountsClientListResult` of struct `AccountsClientListResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListResponse` has been removed
- Field `ChannelsClientGetResult` of struct `ChannelsClientGetResponse` has been removed
- Field `RawResponse` of struct `ChannelsClientGetResponse` has been removed
- Field `ClientCheckNameAvailabilityResult` of struct `ClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `ClientCheckNameAvailabilityResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `AccountsClientGetResult` of struct `AccountsClientGetResponse` has been removed
- Field `RawResponse` of struct `AccountsClientGetResponse` has been removed
- Field `AccountsClientListKeysResult` of struct `AccountsClientListKeysResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListKeysResponse` has been removed
- Field `AccountsClientRegenerateKeyResult` of struct `AccountsClientRegenerateKeyResponse` has been removed
- Field `RawResponse` of struct `AccountsClientRegenerateKeyResponse` has been removed
- Field `RawResponse` of struct `ChannelsClientDeleteResponse` has been removed
- Field `AccountsClientListByResourceGroupResult` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `ChannelsClientListByAccountResult` of struct `ChannelsClientListByAccountResponse` has been removed
- Field `RawResponse` of struct `ChannelsClientListByAccountResponse` has been removed
- Field `AccountsClientUpdateResult` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientDeleteResponse` has been removed
- Field `SKUsClientListResult` of struct `SKUsClientListResponse` has been removed
- Field `RawResponse` of struct `SKUsClientListResponse` has been removed
- Field `AccountsClientListChannelTypesResult` of struct `AccountsClientListChannelTypesResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListChannelTypesResponse` has been removed

### Features Added

- New function `*AccountsClientListKeysPager.More() bool`
- New function `*SKUsClientListPager.More() bool`
- New function `*ChannelsClientListByAccountPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*AccountsClientListByResourceGroupPager.More() bool`
- New function `*AccountsClientListKeysPager.NextPage(context.Context) (AccountsClientListKeysResponse, error)`
- New function `*ChannelsClientListByAccountPager.NextPage(context.Context) (ChannelsClientListByAccountResponse, error)`
- New function `CloudErrorBody.MarshalJSON() ([]byte, error)`
- New function `*AccountsClientListPager.More() bool`
- New function `*SKUsClientListPager.NextPage(context.Context) (SKUsClientListResponse, error)`
- New function `*AccountsClientListPager.NextPage(context.Context) (AccountsClientListResponse, error)`
- New function `*OperationsClientListPager.NextPage(context.Context) (OperationsClientListResponse, error)`
- New function `*AccountsClientListByResourceGroupPager.NextPage(context.Context) (AccountsClientListByResourceGroupResponse, error)`
- New struct `AccountsClientListByResourceGroupPager`
- New struct `AccountsClientListKeysPager`
- New struct `AccountsClientListPager`
- New struct `ChannelsClientListByAccountPager`
- New struct `CloudError`
- New struct `CloudErrorBody`
- New struct `OperationsClientListPager`
- New struct `SKUsClientListPager`
- New anonymous field `Account` in struct `AccountsClientCreateOrUpdateResponse`
- New anonymous field `OperationList` in struct `OperationsClientListResponse`
- New anonymous field `Channel` in struct `ChannelsClientCreateOrUpdateResponse`
- New anonymous field `Account` in struct `AccountsClientUpdateResponse`
- New anonymous field `KeyDescription` in struct `AccountsClientRegenerateKeyResponse`
- New anonymous field `ChannelTypeDescriptionList` in struct `AccountsClientListChannelTypesResponse`
- New anonymous field `KeyDescriptionList` in struct `AccountsClientListKeysResponse`
- New anonymous field `Account` in struct `AccountsClientGetResponse`
- New anonymous field `ChannelList` in struct `ChannelsClientListByAccountResponse`
- New anonymous field `SKUDescriptionList` in struct `SKUsClientListResponse`
- New anonymous field `AccountList` in struct `AccountsClientListResponse`
- New anonymous field `Channel` in struct `ChannelsClientGetResponse`
- New anonymous field `AccountList` in struct `AccountsClientListByResourceGroupResponse`
- New anonymous field `CheckNameAvailabilityResult` in struct `ClientCheckNameAvailabilityResponse`


## 0.1.0 (2022-03-10)

- Init release.