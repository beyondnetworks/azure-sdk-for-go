package azurestack

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Category enumerates the values for category.
type Category string

const (
	// ADFS ...
	ADFS Category = "ADFS"
	// AzureAD ...
	AzureAD Category = "AzureAD"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{ADFS, AzureAD}
}

// CompatibilityIssue enumerates the values for compatibility issue.
type CompatibilityIssue string

const (
	// ADFSIdentitySystemRequired ...
	ADFSIdentitySystemRequired CompatibilityIssue = "ADFSIdentitySystemRequired"
	// AzureADIdentitySystemRequired ...
	AzureADIdentitySystemRequired CompatibilityIssue = "AzureADIdentitySystemRequired"
	// CapacityBillingModelRequired ...
	CapacityBillingModelRequired CompatibilityIssue = "CapacityBillingModelRequired"
	// ConnectionToAzureRequired ...
	ConnectionToAzureRequired CompatibilityIssue = "ConnectionToAzureRequired"
	// ConnectionToInternetRequired ...
	ConnectionToInternetRequired CompatibilityIssue = "ConnectionToInternetRequired"
	// DevelopmentBillingModelRequired ...
	DevelopmentBillingModelRequired CompatibilityIssue = "DevelopmentBillingModelRequired"
	// DisconnectedEnvironmentRequired ...
	DisconnectedEnvironmentRequired CompatibilityIssue = "DisconnectedEnvironmentRequired"
	// HigherDeviceVersionRequired ...
	HigherDeviceVersionRequired CompatibilityIssue = "HigherDeviceVersionRequired"
	// LowerDeviceVersionRequired ...
	LowerDeviceVersionRequired CompatibilityIssue = "LowerDeviceVersionRequired"
	// PayAsYouGoBillingModelRequired ...
	PayAsYouGoBillingModelRequired CompatibilityIssue = "PayAsYouGoBillingModelRequired"
)

// PossibleCompatibilityIssueValues returns an array of possible values for the CompatibilityIssue const type.
func PossibleCompatibilityIssueValues() []CompatibilityIssue {
	return []CompatibilityIssue{ADFSIdentitySystemRequired, AzureADIdentitySystemRequired, CapacityBillingModelRequired, ConnectionToAzureRequired, ConnectionToInternetRequired, DevelopmentBillingModelRequired, DisconnectedEnvironmentRequired, HigherDeviceVersionRequired, LowerDeviceVersionRequired, PayAsYouGoBillingModelRequired}
}

// ComputeRole enumerates the values for compute role.
type ComputeRole string

const (
	// IaaS ...
	IaaS ComputeRole = "IaaS"
	// None ...
	None ComputeRole = "None"
	// PaaS ...
	PaaS ComputeRole = "PaaS"
)

// PossibleComputeRoleValues returns an array of possible values for the ComputeRole const type.
func PossibleComputeRoleValues() []ComputeRole {
	return []ComputeRole{IaaS, None, PaaS}
}

// OperatingSystem enumerates the values for operating system.
type OperatingSystem string

const (
	// OperatingSystemLinux ...
	OperatingSystemLinux OperatingSystem = "Linux"
	// OperatingSystemNone ...
	OperatingSystemNone OperatingSystem = "None"
	// OperatingSystemWindows ...
	OperatingSystemWindows OperatingSystem = "Windows"
)

// PossibleOperatingSystemValues returns an array of possible values for the OperatingSystem const type.
func PossibleOperatingSystemValues() []OperatingSystem {
	return []OperatingSystem{OperatingSystemLinux, OperatingSystemNone, OperatingSystemWindows}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Failed, Succeeded}
}
