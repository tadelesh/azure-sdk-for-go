# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*ServersClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(ServersClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*ServersClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(ServersClientListBySubscriptionResponse, error)`
- Function `*ServersClientListBySubscriptionPager.PageResponse` has been removed
- Function `*ServersClientListBySubscriptionPager.Err` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*ServersClientListByResourceGroupPager.Err` has been removed
- Function `*ServersClientListByResourceGroupPager.PageResponse` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `ServersClientCreateOrUpdateResult` has been removed
- Struct `ServersClientDeleteResult` has been removed
- Struct `ServersClientGetKeysResult` has been removed
- Struct `ServersClientGetResult` has been removed
- Struct `ServersClientListByResourceGroupResult` has been removed
- Struct `ServersClientListBySubscriptionResult` has been removed
- Struct `ServersClientRegenerateKeyResult` has been removed
- Struct `ServersClientUpdateResult` has been removed
- Field `ServersClientListBySubscriptionResult` of struct `ServersClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `ServersClientListBySubscriptionResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `ServersClientListByResourceGroupResult` of struct `ServersClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `ServersClientListByResourceGroupResponse` has been removed
- Field `ServersClientGetResult` of struct `ServersClientGetResponse` has been removed
- Field `RawResponse` of struct `ServersClientGetResponse` has been removed
- Field `ServersClientRegenerateKeyResult` of struct `ServersClientRegenerateKeyResponse` has been removed
- Field `RawResponse` of struct `ServersClientRegenerateKeyResponse` has been removed
- Field `ServersClientGetKeysResult` of struct `ServersClientGetKeysResponse` has been removed
- Field `RawResponse` of struct `ServersClientGetKeysResponse` has been removed
- Field `ServersClientCreateOrUpdateResult` of struct `ServersClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ServersClientCreateOrUpdateResponse` has been removed
- Field `ServersClientUpdateResult` of struct `ServersClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ServersClientUpdateResponse` has been removed
- Field `ServersClientDeleteResult` of struct `ServersClientDeleteResponse` has been removed
- Field `RawResponse` of struct `ServersClientDeleteResponse` has been removed

### Features Added

- New function `*OperationsClientListPager.More() bool`
- New function `*ServersClientListBySubscriptionPager.More() bool`
- New function `*ServersClientListByResourceGroupPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `ServerKeys` in struct `ServersClientRegenerateKeyResponse`
- New field `XMSClientRequestID` in struct `ServersClientRegenerateKeyResponse`
- New field `XMSCorrelationRequestID` in struct `ServersClientRegenerateKeyResponse`
- New anonymous field `ServerKeys` in struct `ServersClientGetKeysResponse`
- New field `XMSClientRequestID` in struct `ServersClientGetKeysResponse`
- New field `XMSCorrelationRequestID` in struct `ServersClientGetKeysResponse`
- New anonymous field `Server` in struct `ServersClientUpdateResponse`
- New field `XMSClientRequestID` in struct `ServersClientUpdateResponse`
- New field `XMSCorrelationRequestID` in struct `ServersClientUpdateResponse`
- New anonymous field `Server` in struct `ServersClientGetResponse`
- New field `XMSClientRequestID` in struct `ServersClientDeleteResponse`
- New field `XMSCorrelationRequestID` in struct `ServersClientDeleteResponse`
- New anonymous field `ServerList` in struct `ServersClientListBySubscriptionResponse`
- New anonymous field `Server` in struct `ServersClientCreateOrUpdateResponse`
- New field `XMSClientRequestID` in struct `ServersClientCreateOrUpdateResponse`
- New field `XMSCorrelationRequestID` in struct `ServersClientCreateOrUpdateResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `ServerList` in struct `ServersClientListByResourceGroupResponse`


## 0.1.0 (2022-03-10)

- Init release.