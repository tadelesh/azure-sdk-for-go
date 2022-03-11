# Release History

## 0.2.0 (2022-03-10)
### Breaking Changes

- Function `*AccountsClientListPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*OperationsClientListPager)`
- Function `*WorkspacesClientListByAccountsPager.NextPage` return value(s) have been changed from `(bool)` to `(WorkspacesClientListByAccountsResponse, error)`
- Function `*AccountsClientListByResourceGroupPager.NextPage` return value(s) have been changed from `(bool)` to `(AccountsClientListByResourceGroupResponse, error)`
- Function `*ProjectsClientListByWorkspacePager.NextPage` return value(s) have been changed from `(bool)` to `(ProjectsClientListByWorkspaceResponse, error)`
- Function `*WorkspacesClientListByAccountsPager.PageResponse` has been removed
- Function `*AccountsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*AccountsClientListByResourceGroupPager.Err` has been removed
- Function `*AccountsClientListPager.PageResponse` has been removed
- Function `*ProjectsClientListByWorkspacePager.Err` has been removed
- Function `*AccountsClientListPager.Err` has been removed
- Function `*ProjectsClientListByWorkspacePager.PageResponse` has been removed
- Function `*WorkspacesClientListByAccountsPager.Err` has been removed
- Struct `AccountsClientCreateOrUpdateResult` has been removed
- Struct `AccountsClientGetResult` has been removed
- Struct `AccountsClientListByResourceGroupResult` has been removed
- Struct `AccountsClientListResult` has been removed
- Struct `AccountsClientUpdateResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `ProjectsClientCreateOrUpdateResult` has been removed
- Struct `ProjectsClientGetResult` has been removed
- Struct `ProjectsClientListByWorkspaceResult` has been removed
- Struct `ProjectsClientUpdateResult` has been removed
- Struct `WorkspacesClientCreateOrUpdateResult` has been removed
- Struct `WorkspacesClientGetResult` has been removed
- Struct `WorkspacesClientListByAccountsResult` has been removed
- Struct `WorkspacesClientUpdateResult` has been removed
- Field `WorkspacesClientCreateOrUpdateResult` of struct `WorkspacesClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `WorkspacesClientCreateOrUpdateResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `ProjectsClientCreateOrUpdateResult` of struct `ProjectsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `ProjectsClientCreateOrUpdateResponse` has been removed
- Field `ProjectsClientUpdateResult` of struct `ProjectsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ProjectsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientDeleteResponse` has been removed
- Field `RawResponse` of struct `ProjectsClientDeleteResponse` has been removed
- Field `WorkspacesClientGetResult` of struct `WorkspacesClientGetResponse` has been removed
- Field `RawResponse` of struct `WorkspacesClientGetResponse` has been removed
- Field `ProjectsClientGetResult` of struct `ProjectsClientGetResponse` has been removed
- Field `RawResponse` of struct `ProjectsClientGetResponse` has been removed
- Field `WorkspacesClientUpdateResult` of struct `WorkspacesClientUpdateResponse` has been removed
- Field `RawResponse` of struct `WorkspacesClientUpdateResponse` has been removed
- Field `AccountsClientListByResourceGroupResult` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `WorkspacesClientDeleteResponse` has been removed
- Field `AccountsClientUpdateResult` of struct `AccountsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientUpdateResponse` has been removed
- Field `ProjectsClientListByWorkspaceResult` of struct `ProjectsClientListByWorkspaceResponse` has been removed
- Field `RawResponse` of struct `ProjectsClientListByWorkspaceResponse` has been removed
- Field `AccountsClientGetResult` of struct `AccountsClientGetResponse` has been removed
- Field `RawResponse` of struct `AccountsClientGetResponse` has been removed
- Field `AccountsClientCreateOrUpdateResult` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `AccountsClientCreateOrUpdateResponse` has been removed
- Field `AccountsClientListResult` of struct `AccountsClientListResponse` has been removed
- Field `RawResponse` of struct `AccountsClientListResponse` has been removed
- Field `WorkspacesClientListByAccountsResult` of struct `WorkspacesClientListByAccountsResponse` has been removed
- Field `RawResponse` of struct `WorkspacesClientListByAccountsResponse` has been removed

### Features Added

- New function `*AccountsClientListByResourceGroupPager.More() bool`
- New function `*OperationsClientListPager.More() bool`
- New function `*WorkspacesClientListByAccountsPager.More() bool`
- New function `*AccountsClientListPager.More() bool`
- New function `*ProjectsClientListByWorkspacePager.More() bool`
- New function `*OperationsClientListPager.NextPage(context.Context) (OperationsClientListResponse, error)`
- New struct `ErrorResponse`
- New struct `OperationsClientListPager`
- New anonymous field `WorkspaceListResult` in struct `WorkspacesClientListByAccountsResponse`
- New anonymous field `Account` in struct `AccountsClientGetResponse`
- New anonymous field `Account` in struct `AccountsClientCreateOrUpdateResponse`
- New anonymous field `AccountListResult` in struct `AccountsClientListByResourceGroupResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `Project` in struct `ProjectsClientGetResponse`
- New anonymous field `Workspace` in struct `WorkspacesClientCreateOrUpdateResponse`
- New anonymous field `Workspace` in struct `WorkspacesClientUpdateResponse`
- New anonymous field `AccountListResult` in struct `AccountsClientListResponse`
- New anonymous field `Project` in struct `ProjectsClientUpdateResponse`
- New anonymous field `ProjectListResult` in struct `ProjectsClientListByWorkspaceResponse`
- New anonymous field `Workspace` in struct `WorkspacesClientGetResponse`
- New anonymous field `Account` in struct `AccountsClientUpdateResponse`
- New anonymous field `Project` in struct `ProjectsClientCreateOrUpdateResponse`


## 0.1.0 (2022-03-10)

- Init release.