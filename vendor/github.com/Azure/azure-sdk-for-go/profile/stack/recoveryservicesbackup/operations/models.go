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

package operations

import original "github.com/lawrencegripper/azure-sdk-for-go/service/recoveryservicesbackup/management/2016-08-10/operations"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ClientDiscoveryDisplay = original.ClientDiscoveryDisplay
type ClientDiscoveryForLogSpecification = original.ClientDiscoveryForLogSpecification
type ClientDiscoveryForProperties = original.ClientDiscoveryForProperties
type ClientDiscoveryForServiceSpecification = original.ClientDiscoveryForServiceSpecification
type ClientDiscoveryResponse = original.ClientDiscoveryResponse
type ClientDiscoveryValueForSingleAPI = original.ClientDiscoveryValueForSingleAPI
type GroupClient = original.GroupClient

func New() ManagementClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) ManagementClient {
	return original.NewWithBaseURI(baseURI)
}
func NewGroupClient() GroupClient {
	return original.NewGroupClient()
}
func NewGroupClientWithBaseURI(baseURI string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
