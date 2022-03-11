# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*HybridUseBenefitClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(HybridUseBenefitClientListResponse, error)`
- Function `*HybridUseBenefitRevisionClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(HybridUseBenefitRevisionClientListResponse, error)`
- Function `*OperationsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(OperationsClientListResponse, error)`
- Function `*HybridUseBenefitClientListPager.PageResponse` has been removed
- Function `*HybridUseBenefitClientListPager.Err` has been removed
- Function `*HybridUseBenefitRevisionClientListPager.Err` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*HybridUseBenefitRevisionClientListPager.PageResponse` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Struct `HybridUseBenefitClientCreateResult` has been removed
- Struct `HybridUseBenefitClientGetResult` has been removed
- Struct `HybridUseBenefitClientListResult` has been removed
- Struct `HybridUseBenefitClientUpdateResult` has been removed
- Struct `HybridUseBenefitRevisionClientListResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `RawResponse` of struct `HybridUseBenefitClientDeleteResponse` has been removed
- Field `HybridUseBenefitRevisionClientListResult` of struct `HybridUseBenefitRevisionClientListResponse` has been removed
- Field `RawResponse` of struct `HybridUseBenefitRevisionClientListResponse` has been removed
- Field `HybridUseBenefitClientListResult` of struct `HybridUseBenefitClientListResponse` has been removed
- Field `RawResponse` of struct `HybridUseBenefitClientListResponse` has been removed
- Field `HybridUseBenefitClientUpdateResult` of struct `HybridUseBenefitClientUpdateResponse` has been removed
- Field `RawResponse` of struct `HybridUseBenefitClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ClientRegisterResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `HybridUseBenefitClientGetResult` of struct `HybridUseBenefitClientGetResponse` has been removed
- Field `RawResponse` of struct `HybridUseBenefitClientGetResponse` has been removed
- Field `HybridUseBenefitClientCreateResult` of struct `HybridUseBenefitClientCreateResponse` has been removed
- Field `RawResponse` of struct `HybridUseBenefitClientCreateResponse` has been removed

### Features Added

- New function `*HybridUseBenefitClientListPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*HybridUseBenefitRevisionClientListPager.More() bool`
- New struct `Error`
- New anonymous field `HybridUseBenefitModel` in struct `HybridUseBenefitClientUpdateResponse`
- New anonymous field `HybridUseBenefitModel` in struct `HybridUseBenefitClientCreateResponse`
- New anonymous field `HybridUseBenefitModel` in struct `HybridUseBenefitClientGetResponse`
- New anonymous field `HybridUseBenefitListResult` in struct `HybridUseBenefitRevisionClientListResponse`
- New anonymous field `OperationList` in struct `OperationsClientListResponse`
- New anonymous field `HybridUseBenefitListResult` in struct `HybridUseBenefitClientListResponse`


## 0.1.0 (2022-03-10)

- Init release.