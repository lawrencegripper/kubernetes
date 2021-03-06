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

package eventhub

import original "github.com/lawrencegripper/azure-sdk-for-go/service/eventhub/management/2015-08-01/eventHub"

type NamespacesClient = original.NamespacesClient
type OperationsClient = original.OperationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ConsumerGroupsClient = original.ConsumerGroupsClient
type EventHubsClient = original.EventHubsClient
type AccessRights = original.AccessRights

const (
	Listen	AccessRights	= original.Listen
	Manage	AccessRights	= original.Manage
	Send	AccessRights	= original.Send
)

type EntityStatus = original.EntityStatus

const (
	Active		EntityStatus	= original.Active
	Creating	EntityStatus	= original.Creating
	Deleting	EntityStatus	= original.Deleting
	Disabled	EntityStatus	= original.Disabled
	ReceiveDisabled	EntityStatus	= original.ReceiveDisabled
	Renaming	EntityStatus	= original.Renaming
	Restoring	EntityStatus	= original.Restoring
	SendDisabled	EntityStatus	= original.SendDisabled
	Unknown		EntityStatus	= original.Unknown
)

type NamespaceState = original.NamespaceState

const (
	NamespaceStateActivating	NamespaceState	= original.NamespaceStateActivating
	NamespaceStateActive		NamespaceState	= original.NamespaceStateActive
	NamespaceStateCreated		NamespaceState	= original.NamespaceStateCreated
	NamespaceStateCreating		NamespaceState	= original.NamespaceStateCreating
	NamespaceStateDisabled		NamespaceState	= original.NamespaceStateDisabled
	NamespaceStateDisabling		NamespaceState	= original.NamespaceStateDisabling
	NamespaceStateEnabling		NamespaceState	= original.NamespaceStateEnabling
	NamespaceStateFailed		NamespaceState	= original.NamespaceStateFailed
	NamespaceStateRemoved		NamespaceState	= original.NamespaceStateRemoved
	NamespaceStateRemoving		NamespaceState	= original.NamespaceStateRemoving
	NamespaceStateSoftDeleted	NamespaceState	= original.NamespaceStateSoftDeleted
	NamespaceStateSoftDeleting	NamespaceState	= original.NamespaceStateSoftDeleting
	NamespaceStateUnknown		NamespaceState	= original.NamespaceStateUnknown
)

type Policykey = original.Policykey

const (
	PrimaryKey	Policykey	= original.PrimaryKey
	SecondaryKey	Policykey	= original.SecondaryKey
)

type SkuName = original.SkuName

const (
	Basic		SkuName	= original.Basic
	Standard	SkuName	= original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic	SkuTier	= original.SkuTierBasic
	SkuTierPremium	SkuTier	= original.SkuTierPremium
	SkuTierStandard	SkuTier	= original.SkuTierStandard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName				UnavailableReason	= original.InvalidName
	NameInLockdown				UnavailableReason	= original.NameInLockdown
	NameInUse				UnavailableReason	= original.NameInUse
	None					UnavailableReason	= original.None
	SubscriptionIsDisabled			UnavailableReason	= original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription	UnavailableReason	= original.TooManyNamespaceInCurrentSubscription
)

type CheckNameAvailabilityParameter = original.CheckNameAvailabilityParameter
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ConsumerGroupCreateOrUpdateParameters = original.ConsumerGroupCreateOrUpdateParameters
type ConsumerGroupListResult = original.ConsumerGroupListResult
type ConsumerGroupProperties = original.ConsumerGroupProperties
type ConsumerGroupResource = original.ConsumerGroupResource
type CreateOrUpdateParameters = original.CreateOrUpdateParameters
type ListResult = original.ListResult
type NamespaceCreateOrUpdateParameters = original.NamespaceCreateOrUpdateParameters
type NamespaceListResult = original.NamespaceListResult
type NamespaceProperties = original.NamespaceProperties
type NamespaceResource = original.NamespaceResource
type NamespaceUpdateParameter = original.NamespaceUpdateParameter
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Properties = original.Properties
type RegenerateKeysParameters = original.RegenerateKeysParameters
type Resource = original.Resource
type ResourceListKeys = original.ResourceListKeys
type ResourceType = original.ResourceType
type SharedAccessAuthorizationRuleCreateOrUpdateParameters = original.SharedAccessAuthorizationRuleCreateOrUpdateParameters
type SharedAccessAuthorizationRuleListResult = original.SharedAccessAuthorizationRuleListResult
type SharedAccessAuthorizationRuleProperties = original.SharedAccessAuthorizationRuleProperties
type SharedAccessAuthorizationRuleResource = original.SharedAccessAuthorizationRuleResource
type Sku = original.Sku
type TrackedResource = original.TrackedResource

func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return original.NewEventHubsClient(subscriptionID)
}
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return original.NewEventHubsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
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
func NewConsumerGroupsClient(subscriptionID string) ConsumerGroupsClient {
	return original.NewConsumerGroupsClient(subscriptionID)
}
func NewConsumerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerGroupsClient {
	return original.NewConsumerGroupsClientWithBaseURI(baseURI, subscriptionID)
}
