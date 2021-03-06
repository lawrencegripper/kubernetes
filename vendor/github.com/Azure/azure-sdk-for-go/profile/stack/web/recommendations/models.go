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

package recommendations

import original "github.com/lawrencegripper/azure-sdk-for-go/service/web/management/2016-03-01/recommendations"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type Channels = original.Channels

const (
	All		Channels	= original.All
	API		Channels	= original.API
	Email		Channels	= original.Email
	Notification	Channels	= original.Notification
	Webhook		Channels	= original.Webhook
)

type NotificationLevel = original.NotificationLevel

const (
	Critical		NotificationLevel	= original.Critical
	Information		NotificationLevel	= original.Information
	NonUrgentSuggestion	NotificationLevel	= original.NonUrgentSuggestion
	Warning			NotificationLevel	= original.Warning
)

type ResourceScopeType = original.ResourceScopeType

const (
	ServerFarm	ResourceScopeType	= original.ServerFarm
	Subscription	ResourceScopeType	= original.Subscription
	WebSite		ResourceScopeType	= original.WebSite
)

type ListRecommendation = original.ListRecommendation
type Recommendation = original.Recommendation
type Rule = original.Rule
type GroupClient = original.GroupClient

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
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
