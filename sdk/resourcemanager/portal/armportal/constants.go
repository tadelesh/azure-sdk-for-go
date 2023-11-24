//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	moduleVersion = "v0.7.0"
)

type ConfigurationName string

const (
	ConfigurationNameDefault ConfigurationName = "default"
)

// PossibleConfigurationNameValues returns the possible values for the ConfigurationName const type.
func PossibleConfigurationNameValues() []ConfigurationName {
	return []ConfigurationName{
		ConfigurationNameDefault,
	}
}
