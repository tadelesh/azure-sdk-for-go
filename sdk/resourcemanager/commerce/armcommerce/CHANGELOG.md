# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*UsageAggregatesClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(UsageAggregatesClientListResponse, error)`
- Function `*UsageAggregatesClientListPager.PageResponse` has been removed
- Function `*UsageAggregatesClientListPager.Err` has been removed
- Struct `RateCardClientGetResult` has been removed
- Struct `UsageAggregatesClientListResult` has been removed
- Field `UsageAggregatesClientListResult` of struct `UsageAggregatesClientListResponse` has been removed
- Field `RawResponse` of struct `UsageAggregatesClientListResponse` has been removed
- Field `RateCardClientGetResult` of struct `RateCardClientGetResponse` has been removed
- Field `RawResponse` of struct `RateCardClientGetResponse` has been removed

### Features Added

- New function `*UsageAggregatesClientListPager.More() bool`
- New struct `ErrorResponse`
- New struct `RateCardQueryParameters`
- New anonymous field `UsageAggregationListResult` in struct `UsageAggregatesClientListResponse`
- New anonymous field `ResourceRateCardInfo` in struct `RateCardClientGetResponse`


## 0.1.0 (2022-03-10)

- Init release.