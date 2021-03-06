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

package databasesecurityalertpolicies

import original "github.com/lawrencegripper/azure-sdk-for-go/service/sql/management/2014-04-01/databaseSecurityAlertPolicies"

type DatabaseThreatDetectionPoliciesClient = original.DatabaseThreatDetectionPoliciesClient
type SecurityAlertPolicyEmailAccountAdmins = original.SecurityAlertPolicyEmailAccountAdmins

const (
	Disabled	SecurityAlertPolicyEmailAccountAdmins	= original.Disabled
	Enabled		SecurityAlertPolicyEmailAccountAdmins	= original.Enabled
)

type SecurityAlertPolicyState = original.SecurityAlertPolicyState

const (
	SecurityAlertPolicyStateDisabled	SecurityAlertPolicyState	= original.SecurityAlertPolicyStateDisabled
	SecurityAlertPolicyStateEnabled		SecurityAlertPolicyState	= original.SecurityAlertPolicyStateEnabled
	SecurityAlertPolicyStateNew		SecurityAlertPolicyState	= original.SecurityAlertPolicyStateNew
)

type SecurityAlertPolicyUseServerDefault = original.SecurityAlertPolicyUseServerDefault

const (
	SecurityAlertPolicyUseServerDefaultDisabled	SecurityAlertPolicyUseServerDefault	= original.SecurityAlertPolicyUseServerDefaultDisabled
	SecurityAlertPolicyUseServerDefaultEnabled	SecurityAlertPolicyUseServerDefault	= original.SecurityAlertPolicyUseServerDefaultEnabled
)

type DatabaseSecurityAlertPolicy = original.DatabaseSecurityAlertPolicy
type DatabaseSecurityAlertPolicyProperties = original.DatabaseSecurityAlertPolicyProperties
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

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
func NewDatabaseThreatDetectionPoliciesClient(subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClient(subscriptionID)
}
func NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
