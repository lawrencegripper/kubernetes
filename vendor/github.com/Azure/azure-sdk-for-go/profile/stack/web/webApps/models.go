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

package webapps

import original "github.com/lawrencegripper/azure-sdk-for-go/service/web/management/2016-08-01/webApps"

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

type AzureResourceType = original.AzureResourceType

const (
	TrafficManager	AzureResourceType	= original.TrafficManager
	Website		AzureResourceType	= original.Website
)

type BackupItemStatus = original.BackupItemStatus

const (
	Created			BackupItemStatus	= original.Created
	Deleted			BackupItemStatus	= original.Deleted
	DeleteFailed		BackupItemStatus	= original.DeleteFailed
	DeleteInProgress	BackupItemStatus	= original.DeleteInProgress
	Failed			BackupItemStatus	= original.Failed
	InProgress		BackupItemStatus	= original.InProgress
	PartiallySucceeded	BackupItemStatus	= original.PartiallySucceeded
	Skipped			BackupItemStatus	= original.Skipped
	Succeeded		BackupItemStatus	= original.Succeeded
	TimedOut		BackupItemStatus	= original.TimedOut
)

type BackupRestoreOperationType = original.BackupRestoreOperationType

const (
	Clone		BackupRestoreOperationType	= original.Clone
	Default		BackupRestoreOperationType	= original.Default
	Relocation	BackupRestoreOperationType	= original.Relocation
)

type BuiltInAuthenticationProvider = original.BuiltInAuthenticationProvider

const (
	AzureActiveDirectory	BuiltInAuthenticationProvider	= original.AzureActiveDirectory
	Facebook		BuiltInAuthenticationProvider	= original.Facebook
	Google			BuiltInAuthenticationProvider	= original.Google
	MicrosoftAccount	BuiltInAuthenticationProvider	= original.MicrosoftAccount
	Twitter			BuiltInAuthenticationProvider	= original.Twitter
)

type CloneAbilityResult = original.CloneAbilityResult

const (
	Cloneable		CloneAbilityResult	= original.Cloneable
	NotCloneable		CloneAbilityResult	= original.NotCloneable
	PartiallyCloneable	CloneAbilityResult	= original.PartiallyCloneable
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

type CustomHostNameDNSRecordType = original.CustomHostNameDNSRecordType

const (
	A	CustomHostNameDNSRecordType	= original.A
	CName	CustomHostNameDNSRecordType	= original.CName
)

type DatabaseType = original.DatabaseType

const (
	DatabaseTypeLocalMySQL	DatabaseType	= original.DatabaseTypeLocalMySQL
	DatabaseTypeMySQL	DatabaseType	= original.DatabaseTypeMySQL
	DatabaseTypePostgreSQL	DatabaseType	= original.DatabaseTypePostgreSQL
	DatabaseTypeSQLAzure	DatabaseType	= original.DatabaseTypeSQLAzure
)

type DNSVerificationTestResult = original.DNSVerificationTestResult

const (
	DNSVerificationTestResultFailed		DNSVerificationTestResult	= original.DNSVerificationTestResultFailed
	DNSVerificationTestResultPassed		DNSVerificationTestResult	= original.DNSVerificationTestResultPassed
	DNSVerificationTestResultSkipped	DNSVerificationTestResult	= original.DNSVerificationTestResultSkipped
)

type FrequencyUnit = original.FrequencyUnit

const (
	Day	FrequencyUnit	= original.Day
	Hour	FrequencyUnit	= original.Hour
)

type HostNameType = original.HostNameType

const (
	Managed		HostNameType	= original.Managed
	Verified	HostNameType	= original.Verified
)

type HostType = original.HostType

const (
	Repository	HostType	= original.Repository
	Standard	HostType	= original.Standard
)

type LogLevel = original.LogLevel

const (
	Error		LogLevel	= original.Error
	Information	LogLevel	= original.Information
	Off		LogLevel	= original.Off
	Verbose		LogLevel	= original.Verbose
	Warning		LogLevel	= original.Warning
)

type ManagedPipelineMode = original.ManagedPipelineMode

const (
	Classic		ManagedPipelineMode	= original.Classic
	Integrated	ManagedPipelineMode	= original.Integrated
)

type OperationStatus = original.OperationStatus

const (
	OperationStatusCreated		OperationStatus	= original.OperationStatusCreated
	OperationStatusFailed		OperationStatus	= original.OperationStatusFailed
	OperationStatusInProgress	OperationStatus	= original.OperationStatusInProgress
	OperationStatusSucceeded	OperationStatus	= original.OperationStatusSucceeded
	OperationStatusTimedOut		OperationStatus	= original.OperationStatusTimedOut
)

type PublishingProfileFormat = original.PublishingProfileFormat

const (
	FileZilla3	PublishingProfileFormat	= original.FileZilla3
	Ftp		PublishingProfileFormat	= original.Ftp
	WebDeploy	PublishingProfileFormat	= original.WebDeploy
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

type UnauthenticatedClientAction = original.UnauthenticatedClientAction

const (
	AllowAnonymous		UnauthenticatedClientAction	= original.AllowAnonymous
	RedirectToLoginPage	UnauthenticatedClientAction	= original.RedirectToLoginPage
)

type UsageState = original.UsageState

const (
	UsageStateExceeded	UsageState	= original.UsageStateExceeded
	UsageStateNormal	UsageState	= original.UsageStateNormal
)

type APIDefinitionInfo = original.APIDefinitionInfo
type ApplicationLogsConfig = original.ApplicationLogsConfig
type AutoHealActions = original.AutoHealActions
type AutoHealCustomAction = original.AutoHealCustomAction
type AutoHealRules = original.AutoHealRules
type AutoHealTriggers = original.AutoHealTriggers
type AzureBlobStorageApplicationLogsConfig = original.AzureBlobStorageApplicationLogsConfig
type AzureBlobStorageHTTPLogsConfig = original.AzureBlobStorageHTTPLogsConfig
type AzureTableStorageApplicationLogsConfig = original.AzureTableStorageApplicationLogsConfig
type BackupItem = original.BackupItem
type BackupItemProperties = original.BackupItemProperties
type BackupItemCollection = original.BackupItemCollection
type BackupRequest = original.BackupRequest
type BackupRequestProperties = original.BackupRequestProperties
type BackupSchedule = original.BackupSchedule
type CloningInfo = original.CloningInfo
type Collection = original.Collection
type ConnectionStringDictionary = original.ConnectionStringDictionary
type ConnStringInfo = original.ConnStringInfo
type ConnStringValueTypePair = original.ConnStringValueTypePair
type CorsSettings = original.CorsSettings
type CsmPublishingProfileOptions = original.CsmPublishingProfileOptions
type CsmSiteRecoveryEntity = original.CsmSiteRecoveryEntity
type CsmSlotEntity = original.CsmSlotEntity
type CsmUsageQuota = original.CsmUsageQuota
type CsmUsageQuotaCollection = original.CsmUsageQuotaCollection
type CustomHostnameAnalysisResult = original.CustomHostnameAnalysisResult
type CustomHostnameAnalysisResultProperties = original.CustomHostnameAnalysisResultProperties
type DatabaseBackupSetting = original.DatabaseBackupSetting
type Deployment = original.Deployment
type DeploymentProperties = original.DeploymentProperties
type DeploymentCollection = original.DeploymentCollection
type EnabledConfig = original.EnabledConfig
type ErrorEntity = original.ErrorEntity
type Experiments = original.Experiments
type FileSystemApplicationLogsConfig = original.FileSystemApplicationLogsConfig
type FileSystemHTTPLogsConfig = original.FileSystemHTTPLogsConfig
type HandlerMapping = original.HandlerMapping
type HostingEnvironmentProfile = original.HostingEnvironmentProfile
type HostNameBinding = original.HostNameBinding
type HostNameBindingProperties = original.HostNameBindingProperties
type HostNameBindingCollection = original.HostNameBindingCollection
type HostNameSslState = original.HostNameSslState
type HTTPLogsConfig = original.HTTPLogsConfig
type HybridConnection = original.HybridConnection
type HybridConnectionProperties = original.HybridConnectionProperties
type HybridConnectionKey = original.HybridConnectionKey
type HybridConnectionKeyProperties = original.HybridConnectionKeyProperties
type Identifier = original.Identifier
type IdentifierProperties = original.IdentifierProperties
type IdentifierCollection = original.IdentifierCollection
type InstanceCollection = original.InstanceCollection
type IPSecurityRestriction = original.IPSecurityRestriction
type ListSiteConfigurationSnapshotInfo = original.ListSiteConfigurationSnapshotInfo
type ListVnetInfo = original.ListVnetInfo
type LocalizableString = original.LocalizableString
type MigrateMySQLRequest = original.MigrateMySQLRequest
type MigrateMySQLRequestProperties = original.MigrateMySQLRequestProperties
type MigrateMySQLStatus = original.MigrateMySQLStatus
type MigrateMySQLStatusProperties = original.MigrateMySQLStatusProperties
type NameValuePair = original.NameValuePair
type NetworkFeatures = original.NetworkFeatures
type NetworkFeaturesProperties = original.NetworkFeaturesProperties
type Operation = original.Operation
type PerfMonCounterCollection = original.PerfMonCounterCollection
type PerfMonResponse = original.PerfMonResponse
type PerfMonSample = original.PerfMonSample
type PerfMonSet = original.PerfMonSet
type PremierAddOn = original.PremierAddOn
type PremierAddOnProperties = original.PremierAddOnProperties
type PushSettings = original.PushSettings
type RampUpRule = original.RampUpRule
type ReadCloser = original.ReadCloser
type RecoverResponse = original.RecoverResponse
type RecoverResponseProperties = original.RecoverResponseProperties
type RelayServiceConnectionEntity = original.RelayServiceConnectionEntity
type RelayServiceConnectionEntityProperties = original.RelayServiceConnectionEntityProperties
type RequestsBasedTrigger = original.RequestsBasedTrigger
type Resource = original.Resource
type ResourceHealthMetadata = original.ResourceHealthMetadata
type ResourceHealthMetadataProperties = original.ResourceHealthMetadataProperties
type ResourceMetric = original.ResourceMetric
type ResourceMetricAvailability = original.ResourceMetricAvailability
type ResourceMetricCollection = original.ResourceMetricCollection
type ResourceMetricDefinition = original.ResourceMetricDefinition
type ResourceMetricDefinitionProperties = original.ResourceMetricDefinitionProperties
type ResourceMetricDefinitionCollection = original.ResourceMetricDefinitionCollection
type ResourceMetricName = original.ResourceMetricName
type ResourceMetricProperty = original.ResourceMetricProperty
type ResourceMetricValue = original.ResourceMetricValue
type RestoreRequest = original.RestoreRequest
type RestoreRequestProperties = original.RestoreRequestProperties
type RestoreResponse = original.RestoreResponse
type RestoreResponseProperties = original.RestoreResponseProperties
type Site = original.Site
type SiteProperties = original.SiteProperties
type SiteAuthSettings = original.SiteAuthSettings
type SiteAuthSettingsProperties = original.SiteAuthSettingsProperties
type SiteCloneability = original.SiteCloneability
type SiteCloneabilityCriterion = original.SiteCloneabilityCriterion
type SiteConfig = original.SiteConfig
type SiteConfigResource = original.SiteConfigResource
type SiteConfigResourceCollection = original.SiteConfigResourceCollection
type SiteConfigurationSnapshotInfo = original.SiteConfigurationSnapshotInfo
type SiteConfigurationSnapshotInfoProperties = original.SiteConfigurationSnapshotInfoProperties
type SiteInstance = original.SiteInstance
type SiteInstanceProperties = original.SiteInstanceProperties
type SiteLimits = original.SiteLimits
type SiteLogsConfig = original.SiteLogsConfig
type SiteLogsConfigProperties = original.SiteLogsConfigProperties
type SiteMachineKey = original.SiteMachineKey
type SitePhpErrorLogFlag = original.SitePhpErrorLogFlag
type SitePhpErrorLogFlagProperties = original.SitePhpErrorLogFlagProperties
type SiteSourceControl = original.SiteSourceControl
type SiteSourceControlProperties = original.SiteSourceControlProperties
type SlotConfigNames = original.SlotConfigNames
type SlotConfigNamesResource = original.SlotConfigNamesResource
type SlotDifference = original.SlotDifference
type SlotDifferenceProperties = original.SlotDifferenceProperties
type SlotDifferenceCollection = original.SlotDifferenceCollection
type SlotSwapStatus = original.SlotSwapStatus
type SlowRequestsBasedTrigger = original.SlowRequestsBasedTrigger
type Snapshot = original.Snapshot
type SnapshotProperties = original.SnapshotProperties
type SnapshotCollection = original.SnapshotCollection
type StatusCodesBasedTrigger = original.StatusCodesBasedTrigger
type StorageMigrationOptions = original.StorageMigrationOptions
type StorageMigrationOptionsProperties = original.StorageMigrationOptionsProperties
type StorageMigrationResponse = original.StorageMigrationResponse
type StorageMigrationResponseProperties = original.StorageMigrationResponseProperties
type String = original.String
type StringDictionary = original.StringDictionary
type User = original.User
type UserProperties = original.UserProperties
type VirtualApplication = original.VirtualApplication
type VirtualDirectory = original.VirtualDirectory
type VnetGateway = original.VnetGateway
type VnetGatewayProperties = original.VnetGatewayProperties
type VnetInfo = original.VnetInfo
type VnetRoute = original.VnetRoute
type VnetRouteProperties = original.VnetRouteProperties
type GroupClient = original.GroupClient

func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
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
