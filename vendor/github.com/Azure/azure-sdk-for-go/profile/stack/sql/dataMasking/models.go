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

package datamasking

import original "github.com/lawrencegripper/azure-sdk-for-go/service/sql/management/2014-04-01/dataMasking"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type Function = original.Function

const (
	CCN	Function	= original.CCN
	Default	Function	= original.Default
	Email	Function	= original.Email
	Number	Function	= original.Number
	SSN	Function	= original.SSN
	Text	Function	= original.Text
)

type RuleState = original.RuleState

const (
	Disabled	RuleState	= original.Disabled
	Enabled		RuleState	= original.Enabled
)

type State = original.State

const (
	StateDisabled	State	= original.StateDisabled
	StateEnabled	State	= original.StateEnabled
)

type Policy = original.Policy
type PolicyProperties = original.PolicyProperties
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type Rule = original.Rule
type RuleListResult = original.RuleListResult
type RuleProperties = original.RuleProperties
type PoliciesClient = original.PoliciesClient
type RulesClient = original.RulesClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRulesClient(subscriptionID string) RulesClient {
	return original.NewRulesClient(subscriptionID)
}
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return original.NewRulesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
