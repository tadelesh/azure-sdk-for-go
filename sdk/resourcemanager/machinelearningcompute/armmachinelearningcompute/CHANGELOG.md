# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*OperationalizationClustersClientListBySubscriptionIDPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationalizationClustersClientListBySubscriptionIDResponse, error)`
- Function `*OperationalizationClustersClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationalizationClustersClientListByResourceGroupResponse, error)`
- Function `*OperationalizationClustersClientListByResourceGroupPager.Err` has been removed
- Function `*OperationalizationClustersClientListByResourceGroupPager.PageResponse` has been removed
- Function `*OperationalizationClustersClientListBySubscriptionIDPager.PageResponse` has been removed
- Function `*OperationalizationClustersClientListBySubscriptionIDPager.Err` has been removed
- Struct `ClientListAvailableOperationsResult` has been removed
- Struct `OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResult` has been removed
- Struct `OperationalizationClustersClientCreateOrUpdateResult` has been removed
- Struct `OperationalizationClustersClientGetResult` has been removed
- Struct `OperationalizationClustersClientListByResourceGroupResult` has been removed
- Struct `OperationalizationClustersClientListBySubscriptionIDResult` has been removed
- Struct `OperationalizationClustersClientListKeysResult` has been removed
- Struct `OperationalizationClustersClientUpdateResult` has been removed
- Struct `OperationalizationClustersClientUpdateSystemServicesResult` has been removed
- Field `ClientListAvailableOperationsResult` of struct `ClientListAvailableOperationsResponse` has been removed
- Field `RawResponse` of struct `ClientListAvailableOperationsResponse` has been removed
- Field `OperationalizationClustersClientUpdateSystemServicesResult` of struct `OperationalizationClustersClientUpdateSystemServicesResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientUpdateSystemServicesResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientDeletePollerResponse` has been removed
- Field `OperationalizationClustersClientListByResourceGroupResult` of struct `OperationalizationClustersClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientListByResourceGroupResponse` has been removed
- Field `OperationalizationClustersClientListKeysResult` of struct `OperationalizationClustersClientListKeysResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientListKeysResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientUpdateSystemServicesPollerResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientCreateOrUpdatePollerResponse` has been removed
- Field `OperationalizationClustersClientCreateOrUpdateResult` of struct `OperationalizationClustersClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientCreateOrUpdateResponse` has been removed
- Field `OperationalizationClustersClientGetResult` of struct `OperationalizationClustersClientGetResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientGetResponse` has been removed
- Field `OperationalizationClustersClientListBySubscriptionIDResult` of struct `OperationalizationClustersClientListBySubscriptionIDResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientListBySubscriptionIDResponse` has been removed
- Field `OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResult` of struct `OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientDeleteResponse` has been removed
- Field `OperationalizationClustersClientUpdateResult` of struct `OperationalizationClustersClientUpdateResponse` has been removed
- Field `RawResponse` of struct `OperationalizationClustersClientUpdateResponse` has been removed

### Features Added

- New function `*OperationalizationClustersClientListBySubscriptionIDPager.More() bool`
- New function `*OperationalizationClustersClientListByResourceGroupPager.More() bool`
- New anonymous field `OperationalizationCluster` in struct `OperationalizationClustersClientUpdateResponse`
- New anonymous field `CheckSystemServicesUpdatesAvailableResponse` in struct `OperationalizationClustersClientCheckSystemServicesUpdatesAvailableResponse`
- New anonymous field `OperationalizationCluster` in struct `OperationalizationClustersClientGetResponse`
- New anonymous field `OperationalizationCluster` in struct `OperationalizationClustersClientCreateOrUpdateResponse`
- New anonymous field `PaginatedOperationalizationClustersList` in struct `OperationalizationClustersClientListBySubscriptionIDResponse`
- New anonymous field `PaginatedOperationalizationClustersList` in struct `OperationalizationClustersClientListByResourceGroupResponse`
- New anonymous field `AvailableOperations` in struct `ClientListAvailableOperationsResponse`
- New anonymous field `OperationalizationClusterCredentials` in struct `OperationalizationClustersClientListKeysResponse`
- New anonymous field `UpdateSystemServicesResponse` in struct `OperationalizationClustersClientUpdateSystemServicesResponse`


## 0.1.0 (2022-03-10)

- Init release.