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

package registeredidentities

import original "github.com/lawrencegripper/azure-sdk-for-go/service/recoveryservices/management/2016-06-01/registeredidentities"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AuthType = original.AuthType

const (
	AAD			AuthType	= original.AAD
	AccessControlService	AuthType	= original.AccessControlService
	ACS			AuthType	= original.ACS
	AzureActiveDirectory	AuthType	= original.AzureActiveDirectory
	Invalid			AuthType	= original.Invalid
)

type CertificateRequest = original.CertificateRequest
type RawCertificateData = original.RawCertificateData
type ResourceCertificateAndAadDetails = original.ResourceCertificateAndAadDetails
type ResourceCertificateAndAcsDetails = original.ResourceCertificateAndAcsDetails
type ResourceCertificateDetails = original.ResourceCertificateDetails
type VaultCertificateResponse = original.VaultCertificateResponse
type GroupClient = original.GroupClient
type VaultCertificatesClient = original.VaultCertificatesClient

func NewVaultCertificatesClient(subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClient(subscriptionID)
}
func NewVaultCertificatesClientWithBaseURI(baseURI string, subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
