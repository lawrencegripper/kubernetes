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

package appserviceplans

import original "github.com/lawrencegripper/azure-sdk-for-go/service/web/management/2016-09-01/appServicePlans"

type GroupClient = original.GroupClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AutoHealActionType = original.AutoHealActionType

const (
	CustomAction	AutoHealActionType	= original.CustomAction
	LogEvent	AutoHealActionType	= original.LogEvent
	Recycle		AutoHealActionType	= original.Recycle
)

type ConnectionStringType = original.ConnectionStringType

const (
	APIHub		ConnectionStringType	= original.APIHub
	Custom		ConnectionStringType	= original.Custom
	DocDb		ConnectionStringType	= original.DocDb
	EventHub	ConnectionStringType	= original.EventHub
	MySQL		ConnectionStringType	= original.MySQL
	NotificationHub	ConnectionStringType	= original.NotificationHub
	PostgreSQL	ConnectionStringType	= original.PostgreSQL
	RedisCache	ConnectionStringType	= original.RedisCache
	ServiceBus	ConnectionStringType	= original.ServiceBus
	SQLAzure	ConnectionStringType	= original.SQLAzure
	SQLServer	ConnectionStringType	= original.SQLServer
)

type HostType = original.HostType

const (
	Repository	HostType	= original.Repository
	Standard	HostType	= original.Standard
)

type ManagedPipelineMode = original.ManagedPipelineMode

const (
	Classic		ManagedPipelineMode	= original.Classic
	Integrated	ManagedPipelineMode	= original.Integrated
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled	ProvisioningState	= original.Canceled
	Deleting	ProvisioningState	= original.Deleting
	Failed		ProvisioningState	= original.Failed
	InProgress	ProvisioningState	= original.InProgress
	Succeeded	ProvisioningState	= original.Succeeded
)

type RouteType = original.RouteType

const (
	DEFAULT		RouteType	= original.DEFAULT
	INHERITED	RouteType	= original.INHERITED
	STATIC		RouteType	= original.STATIC
)

type ScmType = original.ScmType

const (
	BitbucketGit	ScmType	= original.BitbucketGit
	BitbucketHg	ScmType	= original.BitbucketHg
	CodePlexGit	ScmType	= original.CodePlexGit
	CodePlexHg	ScmType	= original.CodePlexHg
	Dropbox		ScmType	= original.Dropbox
	ExternalGit	ScmType	= original.ExternalGit
	ExternalHg	ScmType	= original.ExternalHg
	GitHub		ScmType	= original.GitHub
	LocalGit	ScmType	= original.LocalGit
	None		ScmType	= original.None
	OneDrive	ScmType	= original.OneDrive
	Tfs		ScmType	= original.Tfs
	VSO		ScmType	= original.VSO
)

type SiteAvailabilityState = original.SiteAvailabilityState

const (
	DisasterRecoveryMode	SiteAvailabilityState	= original.DisasterRecoveryMode
	Limited			SiteAvailabilityState	= original.Limited
	Normal			SiteAvailabilityState	= original.Normal
)

type SiteLoadBalancing = original.SiteLoadBalancing

const (
	LeastRequests		SiteLoadBalancing	= original.LeastRequests
	LeastResponseTime	SiteLoadBalancing	= original.LeastResponseTime
	RequestHash		SiteLoadBalancing	= original.RequestHash
	WeightedRoundRobin	SiteLoadBalancing	= original.WeightedRoundRobin
	WeightedTotalTraffic	SiteLoadBalancing	= original.WeightedTotalTraffic
)

type SslState = original.SslState

const (
	Disabled	SslState	= original.Disabled
	IPBasedEnabled	SslState	= original.IPBasedEnabled
	SniEnabled	SslState	= original.SniEnabled
)

type StatusOptions = original.StatusOptions

const (
	Pending	StatusOptions	= original.Pending
	Ready	StatusOptions	= original.Ready
)

type UsageState = original.UsageState

const (
	UsageStateExceeded	UsageState	= original.UsageStateExceeded
	UsageStateNormal	UsageState	= original.UsageStateNormal
)

type APIDefinitionInfo = original.APIDefinitionInfo
type AppServicePlan = original.AppServicePlan
type AppServicePlanProperties = original.AppServicePlanProperties
type AutoHealActions = original.AutoHealActions
type AutoHealCustomAction = original.AutoHealCustomAction
type AutoHealRules = original.AutoHealRules
type AutoHealTriggers = original.AutoHealTriggers
type Capability = original.Capability
type CloningInfo = original.CloningInfo
type Collection = original.Collection
type ConnStringInfo = original.ConnStringInfo
type CorsSettings = original.CorsSettings
type Experiments = original.Experiments
type HandlerMapping = original.HandlerMapping
type HostingEnvironmentProfile = original.HostingEnvironmentProfile
type HostNameSslState = original.HostNameSslState
type HybridConnection = original.HybridConnection
type HybridConnectionProperties = original.HybridConnectionProperties
type HybridConnectionCollection = original.HybridConnectionCollection
type HybridConnectionKey = original.HybridConnectionKey
type HybridConnectionKeyProperties = original.HybridConnectionKeyProperties
type HybridConnectionLimits = original.HybridConnectionLimits
type HybridConnectionLimitsProperties = original.HybridConnectionLimitsProperties
type IPSecurityRestriction = original.IPSecurityRestriction
type ListCapability = original.ListCapability
type ListVnetInfo = original.ListVnetInfo
type ListVnetRoute = original.ListVnetRoute
type NameValuePair = original.NameValuePair
type PushSettings = original.PushSettings
type RampUpRule = original.RampUpRule
type RequestsBasedTrigger = original.RequestsBasedTrigger
type Resource = original.Resource
type ResourceCollection = original.ResourceCollection
type ResourceMetric = original.ResourceMetric
type ResourceMetricAvailability = original.ResourceMetricAvailability
type ResourceMetricCollection = original.ResourceMetricCollection
type ResourceMetricDefinition = original.ResourceMetricDefinition
type ResourceMetricDefinitionProperties = original.ResourceMetricDefinitionProperties
type ResourceMetricDefinitionCollection = original.ResourceMetricDefinitionCollection
type ResourceMetricName = original.ResourceMetricName
type ResourceMetricProperty = original.ResourceMetricProperty
type ResourceMetricValue = original.ResourceMetricValue
type Site = original.Site
type SiteProperties = original.SiteProperties
type SiteConfig = original.SiteConfig
type SiteLimits = original.SiteLimits
type SiteMachineKey = original.SiteMachineKey
type SkuCapacity = original.SkuCapacity
type SkuDescription = original.SkuDescription
type SlotSwapStatus = original.SlotSwapStatus
type SlowRequestsBasedTrigger = original.SlowRequestsBasedTrigger
type StatusCodesBasedTrigger = original.StatusCodesBasedTrigger
type VirtualApplication = original.VirtualApplication
type VirtualDirectory = original.VirtualDirectory
type VnetGateway = original.VnetGateway
type VnetGatewayProperties = original.VnetGatewayProperties
type VnetInfo = original.VnetInfo
type VnetRoute = original.VnetRoute
type VnetRouteProperties = original.VnetRouteProperties
type WebAppCollection = original.WebAppCollection

func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
