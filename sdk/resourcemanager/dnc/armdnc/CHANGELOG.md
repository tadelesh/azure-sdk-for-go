# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*DelegatedSubnetServiceClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(DelegatedSubnetServiceClientListBySubscriptionResponse, error)`
- Function `*DelegatedNetworkClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(DelegatedNetworkClientListBySubscriptionResponse, error)`
- Function `*OrchestratorInstanceServiceClientListBySubscriptionPager.NextPage` return value(s) have been changed from `(bool)` to `(OrchestratorInstanceServiceClientListBySubscriptionResponse, error)`
- Function `*DelegatedNetworkClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(DelegatedNetworkClientListByResourceGroupResponse, error)`
- Function `*OrchestratorInstanceServiceClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(OrchestratorInstanceServiceClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*DelegatedSubnetServiceClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(DelegatedSubnetServiceClientListByResourceGroupResponse, error)`
- Function `*OperationsClientListPager.Err` has been removed
- Function `*DelegatedSubnetServiceClientListBySubscriptionPager.PageResponse` has been removed
- Function `*DelegatedSubnetServiceClientListByResourceGroupPager.Err` has been removed
- Function `*OrchestratorInstanceServiceClientListBySubscriptionPager.PageResponse` has been removed
- Function `*OrchestratorInstanceServiceClientListByResourceGroupPager.Err` has been removed
- Function `*OrchestratorInstanceServiceClientListByResourceGroupPager.PageResponse` has been removed
- Function `*DelegatedNetworkClientListBySubscriptionPager.PageResponse` has been removed
- Function `*DelegatedSubnetServiceClientListBySubscriptionPager.Err` has been removed
- Function `*DelegatedNetworkClientListByResourceGroupPager.Err` has been removed
- Function `*OrchestratorInstanceServiceClientListBySubscriptionPager.Err` has been removed
- Function `*DelegatedNetworkClientListBySubscriptionPager.Err` has been removed
- Function `*DelegatedNetworkClientListByResourceGroupPager.PageResponse` has been removed
- Function `*DelegatedSubnetServiceClientListByResourceGroupPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Struct `ControllerClientCreateResult` has been removed
- Struct `ControllerClientGetDetailsResult` has been removed
- Struct `ControllerClientPatchResult` has been removed
- Struct `DelegatedNetworkClientListByResourceGroupResult` has been removed
- Struct `DelegatedNetworkClientListBySubscriptionResult` has been removed
- Struct `DelegatedSubnetServiceClientGetDetailsResult` has been removed
- Struct `DelegatedSubnetServiceClientListByResourceGroupResult` has been removed
- Struct `DelegatedSubnetServiceClientListBySubscriptionResult` has been removed
- Struct `DelegatedSubnetServiceClientPatchDetailsResult` has been removed
- Struct `DelegatedSubnetServiceClientPutDetailsResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `OrchestratorInstanceServiceClientCreateResult` has been removed
- Struct `OrchestratorInstanceServiceClientGetDetailsResult` has been removed
- Struct `OrchestratorInstanceServiceClientListByResourceGroupResult` has been removed
- Struct `OrchestratorInstanceServiceClientListBySubscriptionResult` has been removed
- Struct `OrchestratorInstanceServiceClientPatchResult` has been removed
- Field `DelegatedNetworkClientListByResourceGroupResult` of struct `DelegatedNetworkClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `DelegatedNetworkClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientPatchDetailsPollerResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientDeletePollerResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `DelegatedSubnetServiceClientListBySubscriptionResult` of struct `DelegatedSubnetServiceClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientListBySubscriptionResponse` has been removed
- Field `DelegatedSubnetServiceClientPutDetailsResult` of struct `DelegatedSubnetServiceClientPutDetailsResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientPutDetailsResponse` has been removed
- Field `DelegatedNetworkClientListBySubscriptionResult` of struct `DelegatedNetworkClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `DelegatedNetworkClientListBySubscriptionResponse` has been removed
- Field `ControllerClientGetDetailsResult` of struct `ControllerClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `ControllerClientGetDetailsResponse` has been removed
- Field `OrchestratorInstanceServiceClientPatchResult` of struct `OrchestratorInstanceServiceClientPatchResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientPatchResponse` has been removed
- Field `DelegatedSubnetServiceClientPatchDetailsResult` of struct `DelegatedSubnetServiceClientPatchDetailsResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientPatchDetailsResponse` has been removed
- Field `RawResponse` of struct `ControllerClientDeleteResponse` has been removed
- Field `OrchestratorInstanceServiceClientGetDetailsResult` of struct `OrchestratorInstanceServiceClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientGetDetailsResponse` has been removed
- Field `ControllerClientPatchResult` of struct `ControllerClientPatchResponse` has been removed
- Field `RawResponse` of struct `ControllerClientPatchResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientDeleteResponse` has been removed
- Field `OrchestratorInstanceServiceClientListByResourceGroupResult` of struct `OrchestratorInstanceServiceClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientListByResourceGroupResponse` has been removed
- Field `ControllerClientCreateResult` of struct `ControllerClientCreateResponse` has been removed
- Field `RawResponse` of struct `ControllerClientCreateResponse` has been removed
- Field `OrchestratorInstanceServiceClientCreateResult` of struct `OrchestratorInstanceServiceClientCreateResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientCreateResponse` has been removed
- Field `DelegatedSubnetServiceClientListByResourceGroupResult` of struct `DelegatedSubnetServiceClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientDeleteDetailsPollerResponse` has been removed
- Field `RawResponse` of struct `ControllerClientDeletePollerResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientDeleteDetailsResponse` has been removed
- Field `DelegatedSubnetServiceClientGetDetailsResult` of struct `DelegatedSubnetServiceClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientGetDetailsResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientCreatePollerResponse` has been removed
- Field `RawResponse` of struct `ControllerClientCreatePollerResponse` has been removed
- Field `OrchestratorInstanceServiceClientListBySubscriptionResult` of struct `OrchestratorInstanceServiceClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `OrchestratorInstanceServiceClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `DelegatedSubnetServiceClientPutDetailsPollerResponse` has been removed

### Features Added

- New function `*OrchestratorInstanceServiceClientListByResourceGroupPager.More() bool`
- New function `*OrchestratorInstanceServiceClientListBySubscriptionPager.More() bool`
- New function `ErrorDetail.MarshalJSON() ([]byte, error)`
- New function `*DelegatedNetworkClientListBySubscriptionPager.More() bool`
- New function `*DelegatedNetworkClientListByResourceGroupPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*DelegatedSubnetServiceClientListBySubscriptionPager.More() bool`
- New function `*DelegatedSubnetServiceClientListByResourceGroupPager.More() bool`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New anonymous field `DelegatedController` in struct `ControllerClientCreateResponse`
- New anonymous field `DelegatedSubnets` in struct `DelegatedSubnetServiceClientListByResourceGroupResponse`
- New anonymous field `Orchestrator` in struct `OrchestratorInstanceServiceClientGetDetailsResponse`
- New anonymous field `DelegatedSubnet` in struct `DelegatedSubnetServiceClientGetDetailsResponse`
- New anonymous field `Orchestrators` in struct `OrchestratorInstanceServiceClientListBySubscriptionResponse`
- New anonymous field `DelegatedSubnet` in struct `DelegatedSubnetServiceClientPatchDetailsResponse`
- New anonymous field `Orchestrator` in struct `OrchestratorInstanceServiceClientCreateResponse`
- New anonymous field `DelegatedControllers` in struct `DelegatedNetworkClientListByResourceGroupResponse`
- New anonymous field `Orchestrators` in struct `OrchestratorInstanceServiceClientListByResourceGroupResponse`
- New anonymous field `DelegatedController` in struct `ControllerClientPatchResponse`
- New anonymous field `DelegatedSubnet` in struct `DelegatedSubnetServiceClientPutDetailsResponse`
- New anonymous field `DelegatedController` in struct `ControllerClientGetDetailsResponse`
- New anonymous field `DelegatedControllers` in struct `DelegatedNetworkClientListBySubscriptionResponse`
- New anonymous field `DelegatedSubnets` in struct `DelegatedSubnetServiceClientListBySubscriptionResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `Orchestrator` in struct `OrchestratorInstanceServiceClientPatchResponse`


## 0.1.0 (2022-03-10)

- Init release.