# Release History

## 0.4.0 (2022-03-10)
### Breaking Changes

- Struct `DiagnosticSettingsCategoryClientListResult` has been removed
- Struct `DiagnosticSettingsClientCreateOrUpdateResult` has been removed
- Struct `DiagnosticSettingsClientGetResult` has been removed
- Struct `DiagnosticSettingsClientListResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `DiagnosticSettingsClientGetResult` of struct `DiagnosticSettingsClientGetResponse` has been removed
- Field `RawResponse` of struct `DiagnosticSettingsClientGetResponse` has been removed
- Field `DiagnosticSettingsCategoryClientListResult` of struct `DiagnosticSettingsCategoryClientListResponse` has been removed
- Field `RawResponse` of struct `DiagnosticSettingsCategoryClientListResponse` has been removed
- Field `RawResponse` of struct `DiagnosticSettingsClientDeleteResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `DiagnosticSettingsClientCreateOrUpdateResult` of struct `DiagnosticSettingsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `DiagnosticSettingsClientCreateOrUpdateResponse` has been removed
- Field `DiagnosticSettingsClientListResult` of struct `DiagnosticSettingsClientListResponse` has been removed
- Field `RawResponse` of struct `DiagnosticSettingsClientListResponse` has been removed

### Features Added

- New function `ErrorDefinition.MarshalJSON() ([]byte, error)`
- New struct `ErrorDefinition`
- New struct `ErrorResponse`
- New anonymous field `DiagnosticSettingsCategoryResourceCollection` in struct `DiagnosticSettingsCategoryClientListResponse`
- New anonymous field `DiagnosticSettingsResource` in struct `DiagnosticSettingsClientCreateOrUpdateResponse`
- New anonymous field `OperationsDiscoveryCollection` in struct `OperationsClientListResponse`
- New anonymous field `DiagnosticSettingsResource` in struct `DiagnosticSettingsClientGetResponse`
- New anonymous field `DiagnosticSettingsResourceCollection` in struct `DiagnosticSettingsClientListResponse`


## 0.3.0 (2022-03-10)
### Breaking Changes

- Type of `OperationsDiscovery.Properties` has been changed from `map[string]interface{}` to `interface{}`
- Function `ErrorDefinition.MarshalJSON` has been removed
- Struct `ErrorDefinition` has been removed
- Struct `ErrorResponse` has been removed

### Features Added



## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*DiagnosticSettingsClient.Get` parameter(s) have been changed from `(context.Context, string, *DiagnosticSettingsGetOptions)` to `(context.Context, string, *DiagnosticSettingsClientGetOptions)`
- Function `*DiagnosticSettingsClient.Get` return value(s) have been changed from `(DiagnosticSettingsGetResponse, error)` to `(DiagnosticSettingsClientGetResponse, error)`
- Function `*DiagnosticSettingsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, DiagnosticSettingsResource, *DiagnosticSettingsCreateOrUpdateOptions)` to `(context.Context, string, DiagnosticSettingsResource, *DiagnosticSettingsClientCreateOrUpdateOptions)`
- Function `*DiagnosticSettingsClient.CreateOrUpdate` return value(s) have been changed from `(DiagnosticSettingsCreateOrUpdateResponse, error)` to `(DiagnosticSettingsClientCreateOrUpdateResponse, error)`
- Function `*DiagnosticSettingsCategoryClient.List` parameter(s) have been changed from `(context.Context, *DiagnosticSettingsCategoryListOptions)` to `(context.Context, *DiagnosticSettingsCategoryClientListOptions)`
- Function `*DiagnosticSettingsCategoryClient.List` return value(s) have been changed from `(DiagnosticSettingsCategoryListResponse, error)` to `(DiagnosticSettingsCategoryClientListResponse, error)`
- Function `*DiagnosticSettingsClient.Delete` parameter(s) have been changed from `(context.Context, string, *DiagnosticSettingsDeleteOptions)` to `(context.Context, string, *DiagnosticSettingsClientDeleteOptions)`
- Function `*DiagnosticSettingsClient.Delete` return value(s) have been changed from `(DiagnosticSettingsDeleteResponse, error)` to `(DiagnosticSettingsClientDeleteResponse, error)`
- Function `*DiagnosticSettingsClient.List` parameter(s) have been changed from `(context.Context, *DiagnosticSettingsListOptions)` to `(context.Context, *DiagnosticSettingsClientListOptions)`
- Function `*DiagnosticSettingsClient.List` return value(s) have been changed from `(DiagnosticSettingsListResponse, error)` to `(DiagnosticSettingsClientListResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `ErrorResponse.Error` has been removed
- Struct `DiagnosticSettingsCategoryListOptions` has been removed
- Struct `DiagnosticSettingsCategoryListResponse` has been removed
- Struct `DiagnosticSettingsCategoryListResult` has been removed
- Struct `DiagnosticSettingsCreateOrUpdateOptions` has been removed
- Struct `DiagnosticSettingsCreateOrUpdateResponse` has been removed
- Struct `DiagnosticSettingsCreateOrUpdateResult` has been removed
- Struct `DiagnosticSettingsDeleteOptions` has been removed
- Struct `DiagnosticSettingsDeleteResponse` has been removed
- Struct `DiagnosticSettingsGetOptions` has been removed
- Struct `DiagnosticSettingsGetResponse` has been removed
- Struct `DiagnosticSettingsGetResult` has been removed
- Struct `DiagnosticSettingsListOptions` has been removed
- Struct `DiagnosticSettingsListResponse` has been removed
- Struct `DiagnosticSettingsListResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `ProxyOnlyResource` of struct `DiagnosticSettingsCategoryResource` has been removed
- Field `ProxyOnlyResource` of struct `DiagnosticSettingsResource` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed

### Features Added

- New struct `DiagnosticSettingsCategoryClientListOptions`
- New struct `DiagnosticSettingsCategoryClientListResponse`
- New struct `DiagnosticSettingsCategoryClientListResult`
- New struct `DiagnosticSettingsClientCreateOrUpdateOptions`
- New struct `DiagnosticSettingsClientCreateOrUpdateResponse`
- New struct `DiagnosticSettingsClientCreateOrUpdateResult`
- New struct `DiagnosticSettingsClientDeleteOptions`
- New struct `DiagnosticSettingsClientDeleteResponse`
- New struct `DiagnosticSettingsClientGetOptions`
- New struct `DiagnosticSettingsClientGetResponse`
- New struct `DiagnosticSettingsClientGetResult`
- New struct `DiagnosticSettingsClientListOptions`
- New struct `DiagnosticSettingsClientListResponse`
- New struct `DiagnosticSettingsClientListResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New field `Error` in struct `ErrorResponse`
- New field `ID` in struct `DiagnosticSettingsCategoryResource`
- New field `Name` in struct `DiagnosticSettingsCategoryResource`
- New field `Type` in struct `DiagnosticSettingsCategoryResource`
- New field `Name` in struct `DiagnosticSettingsResource`
- New field `Type` in struct `DiagnosticSettingsResource`
- New field `ID` in struct `DiagnosticSettingsResource`


## 0.1.0 (2021-11-30)

- Initial preview release.
