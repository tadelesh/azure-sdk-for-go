// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This file enables 'go generate' to regenerate this specific SDK
//go:generate pwsh.exe ../../../../eng/scripts/build.ps1 -skipBuild -cleanGenerated -format -tidy -generate resourcemanager/devspaces/armdevspaces

package armdevspaces
