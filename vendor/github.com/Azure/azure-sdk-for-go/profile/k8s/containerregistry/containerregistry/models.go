// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder

package containerregistry

import original "github.com/lawrencegripper/azure-sdk-for-go/service/containerregistry/management/2017-03-01/containerregistry"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type PasswordName = original.PasswordName

const (
	Password	PasswordName	= original.Password
	Password2	PasswordName	= original.Password2
)

type ProvisioningState = original.ProvisioningState

const (
	Creating	ProvisioningState	= original.Creating
	Succeeded	ProvisioningState	= original.Succeeded
)

type SkuTier = original.SkuTier

const (
	Basic SkuTier = original.Basic
)

type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type Registry = original.Registry
type RegistryCreateParameters = original.RegistryCreateParameters
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesCreateParameters = original.RegistryPropertiesCreateParameters
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type Resource = original.Resource
type Sku = original.Sku
type StorageAccountParameters = original.StorageAccountParameters
type StorageAccountProperties = original.StorageAccountProperties
type OperationsClient = original.OperationsClient
type RegistriesClient = original.RegistriesClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
