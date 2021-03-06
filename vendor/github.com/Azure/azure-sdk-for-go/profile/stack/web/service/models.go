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

package service

import original "github.com/lawrencegripper/azure-sdk-for-go/service/web/management/2015-08-01/service"

type CertificateOrdersClient = original.CertificateOrdersClient
type ClassicMobileServicesClient = original.ClassicMobileServicesClient
type GlobalClient = original.GlobalClient
type HostingEnvironmentsClient = original.HostingEnvironmentsClient
type UsageClient = original.UsageClient
type TopLevelDomainsClient = original.TopLevelDomainsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type GlobalCertificateOrderClient = original.GlobalCertificateOrderClient
type ManagedHostingEnvironmentsClient = original.ManagedHostingEnvironmentsClient
type ServerFarmsClient = original.ServerFarmsClient
type SitesClient = original.SitesClient
type GlobalDomainRegistrationClient = original.GlobalDomainRegistrationClient
type AccessControlEntryAction = original.AccessControlEntryAction

const (
	Deny	AccessControlEntryAction	= original.Deny
	Permit	AccessControlEntryAction	= original.Permit
)

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

type CertificateOrderActionType = original.CertificateOrderActionType

const (
	CertificateIssued		CertificateOrderActionType	= original.CertificateIssued
	CertificateOrderCanceled	CertificateOrderActionType	= original.CertificateOrderCanceled
	CertificateOrderCreated		CertificateOrderActionType	= original.CertificateOrderCreated
	CertificateRevoked		CertificateOrderActionType	= original.CertificateRevoked
	DomainValidationComplete	CertificateOrderActionType	= original.DomainValidationComplete
	FraudDetected			CertificateOrderActionType	= original.FraudDetected
	OrgNameChange			CertificateOrderActionType	= original.OrgNameChange
	OrgValidationComplete		CertificateOrderActionType	= original.OrgValidationComplete
	SanDrop				CertificateOrderActionType	= original.SanDrop
)

type CertificateOrderStatus = original.CertificateOrderStatus

const (
	Canceled		CertificateOrderStatus	= original.Canceled
	Denied			CertificateOrderStatus	= original.Denied
	Expired			CertificateOrderStatus	= original.Expired
	Issued			CertificateOrderStatus	= original.Issued
	NotSubmitted		CertificateOrderStatus	= original.NotSubmitted
	Pendingissuance		CertificateOrderStatus	= original.Pendingissuance
	PendingRekey		CertificateOrderStatus	= original.PendingRekey
	Pendingrevocation	CertificateOrderStatus	= original.Pendingrevocation
	Revoked			CertificateOrderStatus	= original.Revoked
	Unused			CertificateOrderStatus	= original.Unused
)

type CertificateProductType = original.CertificateProductType

const (
	StandardDomainValidatedSsl		CertificateProductType	= original.StandardDomainValidatedSsl
	StandardDomainValidatedWildCardSsl	CertificateProductType	= original.StandardDomainValidatedWildCardSsl
)

type Channels = original.Channels

const (
	All		Channels	= original.All
	API		Channels	= original.API
	Email		Channels	= original.Email
	Notification	Channels	= original.Notification
)

type CloneAbilityResult = original.CloneAbilityResult

const (
	Cloneable		CloneAbilityResult	= original.Cloneable
	NotCloneable		CloneAbilityResult	= original.NotCloneable
	PartiallyCloneable	CloneAbilityResult	= original.PartiallyCloneable
)

type ComputeModeOptions = original.ComputeModeOptions

const (
	Dedicated	ComputeModeOptions	= original.Dedicated
	Dynamic		ComputeModeOptions	= original.Dynamic
	Shared		ComputeModeOptions	= original.Shared
)

type CustomHostNameDNSRecordType = original.CustomHostNameDNSRecordType

const (
	A	CustomHostNameDNSRecordType	= original.A
	CName	CustomHostNameDNSRecordType	= original.CName
)

type DatabaseServerType = original.DatabaseServerType

const (
	Custom		DatabaseServerType	= original.Custom
	MySQL		DatabaseServerType	= original.MySQL
	SQLAzure	DatabaseServerType	= original.SQLAzure
	SQLServer	DatabaseServerType	= original.SQLServer
)

type DomainStatus = original.DomainStatus

const (
	DomainStatusActive		DomainStatus	= original.DomainStatusActive
	DomainStatusAwaiting		DomainStatus	= original.DomainStatusAwaiting
	DomainStatusCancelled		DomainStatus	= original.DomainStatusCancelled
	DomainStatusConfiscated		DomainStatus	= original.DomainStatusConfiscated
	DomainStatusDisabled		DomainStatus	= original.DomainStatusDisabled
	DomainStatusExcluded		DomainStatus	= original.DomainStatusExcluded
	DomainStatusExpired		DomainStatus	= original.DomainStatusExpired
	DomainStatusFailed		DomainStatus	= original.DomainStatusFailed
	DomainStatusHeld		DomainStatus	= original.DomainStatusHeld
	DomainStatusJSONConverterFailed	DomainStatus	= original.DomainStatusJSONConverterFailed
	DomainStatusLocked		DomainStatus	= original.DomainStatusLocked
	DomainStatusParked		DomainStatus	= original.DomainStatusParked
	DomainStatusPending		DomainStatus	= original.DomainStatusPending
	DomainStatusReserved		DomainStatus	= original.DomainStatusReserved
	DomainStatusReverted		DomainStatus	= original.DomainStatusReverted
	DomainStatusSuspended		DomainStatus	= original.DomainStatusSuspended
	DomainStatusTransferred		DomainStatus	= original.DomainStatusTransferred
	DomainStatusUnknown		DomainStatus	= original.DomainStatusUnknown
	DomainStatusUnlocked		DomainStatus	= original.DomainStatusUnlocked
	DomainStatusUnparked		DomainStatus	= original.DomainStatusUnparked
	DomainStatusUpdated		DomainStatus	= original.DomainStatusUpdated
)

type DomainType = original.DomainType

const (
	Regular		DomainType	= original.Regular
	SoftDeleted	DomainType	= original.SoftDeleted
)

type FrequencyUnit = original.FrequencyUnit

const (
	Day	FrequencyUnit	= original.Day
	Hour	FrequencyUnit	= original.Hour
)

type HostingEnvironmentStatus = original.HostingEnvironmentStatus

const (
	Deleting	HostingEnvironmentStatus	= original.Deleting
	Preparing	HostingEnvironmentStatus	= original.Preparing
	Ready		HostingEnvironmentStatus	= original.Ready
	Scaling		HostingEnvironmentStatus	= original.Scaling
)

type HostNameType = original.HostNameType

const (
	Managed		HostNameType	= original.Managed
	Verified	HostNameType	= original.Verified
)

type InternalLoadBalancingMode = original.InternalLoadBalancingMode

const (
	None		InternalLoadBalancingMode	= original.None
	Publishing	InternalLoadBalancingMode	= original.Publishing
	Web		InternalLoadBalancingMode	= original.Web
)

type KeyVaultSecretStatus = original.KeyVaultSecretStatus

const (
	KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault	KeyVaultSecretStatus	= original.KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault
	KeyVaultSecretStatusCertificateOrderFailed			KeyVaultSecretStatus	= original.KeyVaultSecretStatusCertificateOrderFailed
	KeyVaultSecretStatusInitialized					KeyVaultSecretStatus	= original.KeyVaultSecretStatusInitialized
	KeyVaultSecretStatusKeyVaultDoesNotExist			KeyVaultSecretStatus	= original.KeyVaultSecretStatusKeyVaultDoesNotExist
	KeyVaultSecretStatusKeyVaultSecretDoesNotExist			KeyVaultSecretStatus	= original.KeyVaultSecretStatusKeyVaultSecretDoesNotExist
	KeyVaultSecretStatusOperationNotPermittedOnKeyVault		KeyVaultSecretStatus	= original.KeyVaultSecretStatusOperationNotPermittedOnKeyVault
	KeyVaultSecretStatusSucceeded					KeyVaultSecretStatus	= original.KeyVaultSecretStatusSucceeded
	KeyVaultSecretStatusUnknown					KeyVaultSecretStatus	= original.KeyVaultSecretStatusUnknown
	KeyVaultSecretStatusUnknownError				KeyVaultSecretStatus	= original.KeyVaultSecretStatusUnknownError
	KeyVaultSecretStatusWaitingOnCertificateOrder			KeyVaultSecretStatus	= original.KeyVaultSecretStatusWaitingOnCertificateOrder
)

type LogLevel = original.LogLevel

const (
	Error		LogLevel	= original.Error
	Information	LogLevel	= original.Information
	Off		LogLevel	= original.Off
	Verbose		LogLevel	= original.Verbose
	Warning		LogLevel	= original.Warning
)

type ManagedHostingEnvironmentStatus = original.ManagedHostingEnvironmentStatus

const (
	ManagedHostingEnvironmentStatusDeleting		ManagedHostingEnvironmentStatus	= original.ManagedHostingEnvironmentStatusDeleting
	ManagedHostingEnvironmentStatusPreparing	ManagedHostingEnvironmentStatus	= original.ManagedHostingEnvironmentStatusPreparing
	ManagedHostingEnvironmentStatusReady		ManagedHostingEnvironmentStatus	= original.ManagedHostingEnvironmentStatusReady
)

type ManagedPipelineMode = original.ManagedPipelineMode

const (
	Classic		ManagedPipelineMode	= original.Classic
	Integrated	ManagedPipelineMode	= original.Integrated
)

type NotificationLevel = original.NotificationLevel

const (
	NotificationLevelCritical		NotificationLevel	= original.NotificationLevelCritical
	NotificationLevelInformation		NotificationLevel	= original.NotificationLevelInformation
	NotificationLevelNonUrgentSuggestion	NotificationLevel	= original.NotificationLevelNonUrgentSuggestion
	NotificationLevelWarning		NotificationLevel	= original.NotificationLevelWarning
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled	ProvisioningState	= original.ProvisioningStateCanceled
	ProvisioningStateDeleting	ProvisioningState	= original.ProvisioningStateDeleting
	ProvisioningStateFailed		ProvisioningState	= original.ProvisioningStateFailed
	ProvisioningStateInProgress	ProvisioningState	= original.ProvisioningStateInProgress
	ProvisioningStateSucceeded	ProvisioningState	= original.ProvisioningStateSucceeded
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
	StatusOptionsPending	StatusOptions	= original.StatusOptionsPending
	StatusOptionsReady	StatusOptions	= original.StatusOptionsReady
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

type WorkerSizeOptions = original.WorkerSizeOptions

const (
	WorkerSizeOptionsDefault	WorkerSizeOptions	= original.WorkerSizeOptionsDefault
	WorkerSizeOptionsLarge		WorkerSizeOptions	= original.WorkerSizeOptionsLarge
	WorkerSizeOptionsMedium		WorkerSizeOptions	= original.WorkerSizeOptionsMedium
	WorkerSizeOptionsSmall		WorkerSizeOptions	= original.WorkerSizeOptionsSmall
)

type Address = original.Address
type AddressResponse = original.AddressResponse
type APIDefinitionInfo = original.APIDefinitionInfo
type ApplicationLogsConfig = original.ApplicationLogsConfig
type ArmPlan = original.ArmPlan
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
type Certificate = original.Certificate
type CertificateProperties = original.CertificateProperties
type CertificateCollection = original.CertificateCollection
type CertificateDetails = original.CertificateDetails
type CertificateDetailsProperties = original.CertificateDetailsProperties
type CertificateEmail = original.CertificateEmail
type CertificateEmailProperties = original.CertificateEmailProperties
type CertificateOrder = original.CertificateOrder
type CertificateOrderProperties = original.CertificateOrderProperties
type CertificateOrderAction = original.CertificateOrderAction
type CertificateOrderActionProperties = original.CertificateOrderActionProperties
type CertificateOrderCertificate = original.CertificateOrderCertificate
type CertificateOrderCertificateProperties = original.CertificateOrderCertificateProperties
type CertificateOrderCertificateCollection = original.CertificateOrderCertificateCollection
type CertificateOrderCollection = original.CertificateOrderCollection
type ClassicMobileService = original.ClassicMobileService
type ClassicMobileServiceProperties = original.ClassicMobileServiceProperties
type ClassicMobileServiceCollection = original.ClassicMobileServiceCollection
type CloningInfo = original.CloningInfo
type ConnectionStringDictionary = original.ConnectionStringDictionary
type ConnStringInfo = original.ConnStringInfo
type ConnStringValueTypePair = original.ConnStringValueTypePair
type Contact = original.Contact
type CorsSettings = original.CorsSettings
type CsmMoveResourceEnvelope = original.CsmMoveResourceEnvelope
type CsmPublishingProfileOptions = original.CsmPublishingProfileOptions
type CsmSiteRecoveryEntity = original.CsmSiteRecoveryEntity
type CsmSlotEntity = original.CsmSlotEntity
type CsmUsageQuota = original.CsmUsageQuota
type CsmUsageQuotaCollection = original.CsmUsageQuotaCollection
type Csr = original.Csr
type CsrProperties = original.CsrProperties
type DatabaseBackupSetting = original.DatabaseBackupSetting
type DeletedSite = original.DeletedSite
type DeletedSiteProperties = original.DeletedSiteProperties
type DeletedSiteCollection = original.DeletedSiteCollection
type Deployment = original.Deployment
type DeploymentProperties = original.DeploymentProperties
type DeploymentCollection = original.DeploymentCollection
type Domain = original.Domain
type DomainProperties = original.DomainProperties
type DomainAvailablilityCheckResult = original.DomainAvailablilityCheckResult
type DomainCollection = original.DomainCollection
type DomainControlCenterSsoRequest = original.DomainControlCenterSsoRequest
type DomainPurchaseConsent = original.DomainPurchaseConsent
type DomainRecommendationSearchParameters = original.DomainRecommendationSearchParameters
type DomainRegistrationInput = original.DomainRegistrationInput
type DomainRegistrationInputProperties = original.DomainRegistrationInputProperties
type EnabledConfig = original.EnabledConfig
type Experiments = original.Experiments
type FileSystemApplicationLogsConfig = original.FileSystemApplicationLogsConfig
type FileSystemHTTPLogsConfig = original.FileSystemHTTPLogsConfig
type GeoRegion = original.GeoRegion
type GeoRegionProperties = original.GeoRegionProperties
type GeoRegionCollection = original.GeoRegionCollection
type HandlerMapping = original.HandlerMapping
type HostingEnvironment = original.HostingEnvironment
type HostingEnvironmentProperties = original.HostingEnvironmentProperties
type HostingEnvironmentCollection = original.HostingEnvironmentCollection
type HostingEnvironmentDiagnostics = original.HostingEnvironmentDiagnostics
type HostingEnvironmentProfile = original.HostingEnvironmentProfile
type HostName = original.HostName
type HostNameBinding = original.HostNameBinding
type HostNameBindingProperties = original.HostNameBindingProperties
type HostNameBindingCollection = original.HostNameBindingCollection
type HostNameSslState = original.HostNameSslState
type HTTPLogsConfig = original.HTTPLogsConfig
type IPSecurityRestriction = original.IPSecurityRestriction
type KeyValuePairStringString = original.KeyValuePairStringString
type ListCertificateEmail = original.ListCertificateEmail
type ListCertificateOrderAction = original.ListCertificateOrderAction
type ListCsr = original.ListCsr
type ListHostingEnvironmentDiagnostics = original.ListHostingEnvironmentDiagnostics
type ListRecommendation = original.ListRecommendation
type ListVnetInfo = original.ListVnetInfo
type ListVnetRoute = original.ListVnetRoute
type LocalizableString = original.LocalizableString
type ManagedHostingEnvironment = original.ManagedHostingEnvironment
type ManagedHostingEnvironmentProperties = original.ManagedHostingEnvironmentProperties
type ManagedHostingEnvironmentCollection = original.ManagedHostingEnvironmentCollection
type MetricAvailabilily = original.MetricAvailabilily
type MetricDefinition = original.MetricDefinition
type MetricDefinitionProperties = original.MetricDefinitionProperties
type MetricDefinitionCollection = original.MetricDefinitionCollection
type NameIdentifier = original.NameIdentifier
type NameIdentifierCollection = original.NameIdentifierCollection
type NameValuePair = original.NameValuePair
type NetworkAccessControlEntry = original.NetworkAccessControlEntry
type NetworkFeatures = original.NetworkFeatures
type NetworkFeaturesProperties = original.NetworkFeaturesProperties
type PremierAddOnRequest = original.PremierAddOnRequest
type RampUpRule = original.RampUpRule
type ReadCloser = original.ReadCloser
type Recommendation = original.Recommendation
type RecommendationRule = original.RecommendationRule
type ReissueCertificateOrderRequest = original.ReissueCertificateOrderRequest
type ReissueCertificateOrderRequestProperties = original.ReissueCertificateOrderRequestProperties
type RelayServiceConnectionEntity = original.RelayServiceConnectionEntity
type RelayServiceConnectionEntityProperties = original.RelayServiceConnectionEntityProperties
type RenewCertificateOrderRequest = original.RenewCertificateOrderRequest
type RenewCertificateOrderRequestProperties = original.RenewCertificateOrderRequestProperties
type RequestsBasedTrigger = original.RequestsBasedTrigger
type Resource = original.Resource
type ResourceMetric = original.ResourceMetric
type ResourceMetricCollection = original.ResourceMetricCollection
type ResourceMetricName = original.ResourceMetricName
type ResourceMetricValue = original.ResourceMetricValue
type ResourceNameAvailability = original.ResourceNameAvailability
type ResourceNameAvailabilityRequest = original.ResourceNameAvailabilityRequest
type RestoreRequest = original.RestoreRequest
type RestoreRequestProperties = original.RestoreRequestProperties
type RestoreResponse = original.RestoreResponse
type RestoreResponseProperties = original.RestoreResponseProperties
type RoutingRule = original.RoutingRule
type ServerFarmCollection = original.ServerFarmCollection
type ServerFarmWithRichSku = original.ServerFarmWithRichSku
type ServerFarmWithRichSkuProperties = original.ServerFarmWithRichSkuProperties
type SetObject = original.SetObject
type Site = original.Site
type SiteProperties = original.SiteProperties
type SiteAuthSettings = original.SiteAuthSettings
type SiteCloneability = original.SiteCloneability
type SiteCloneabilityCriterion = original.SiteCloneabilityCriterion
type SiteCollection = original.SiteCollection
type SiteConfig = original.SiteConfig
type SiteConfigProperties = original.SiteConfigProperties
type SiteInstance = original.SiteInstance
type SiteInstanceProperties = original.SiteInstanceProperties
type SiteInstanceCollection = original.SiteInstanceCollection
type SiteLimits = original.SiteLimits
type SiteLogsConfig = original.SiteLogsConfig
type SiteLogsConfigProperties = original.SiteLogsConfigProperties
type SitePropertiesModel = original.SitePropertiesModel
type SiteSourceControl = original.SiteSourceControl
type SiteSourceControlProperties = original.SiteSourceControlProperties
type SkuCapacity = original.SkuCapacity
type SkuDescription = original.SkuDescription
type SkuInfo = original.SkuInfo
type SkuInfoCollection = original.SkuInfoCollection
type SlotConfigNames = original.SlotConfigNames
type SlotConfigNamesResource = original.SlotConfigNamesResource
type SlotConfigNamesResourceProperties = original.SlotConfigNamesResourceProperties
type SlotDifference = original.SlotDifference
type SlotDifferenceProperties = original.SlotDifferenceProperties
type SlotDifferenceCollection = original.SlotDifferenceCollection
type SlowRequestsBasedTrigger = original.SlowRequestsBasedTrigger
type SourceControl = original.SourceControl
type SourceControlProperties = original.SourceControlProperties
type SourceControlCollection = original.SourceControlCollection
type StampCapacity = original.StampCapacity
type StampCapacityCollection = original.StampCapacityCollection
type StatusCodesBasedTrigger = original.StatusCodesBasedTrigger
type StringDictionary = original.StringDictionary
type TldLegalAgreement = original.TldLegalAgreement
type TldLegalAgreementCollection = original.TldLegalAgreementCollection
type TopLevelDomain = original.TopLevelDomain
type TopLevelDomainProperties = original.TopLevelDomainProperties
type TopLevelDomainAgreementOption = original.TopLevelDomainAgreementOption
type TopLevelDomainCollection = original.TopLevelDomainCollection
type Usage = original.Usage
type UsageProperties = original.UsageProperties
type UsageCollection = original.UsageCollection
type User = original.User
type UserProperties = original.UserProperties
type VirtualApplication = original.VirtualApplication
type VirtualDirectory = original.VirtualDirectory
type VirtualIPMapping = original.VirtualIPMapping
type VirtualNetworkProfile = original.VirtualNetworkProfile
type VnetGateway = original.VnetGateway
type VnetGatewayProperties = original.VnetGatewayProperties
type VnetInfo = original.VnetInfo
type VnetInfoProperties = original.VnetInfoProperties
type VnetRoute = original.VnetRoute
type VnetRouteProperties = original.VnetRouteProperties
type WorkerPool = original.WorkerPool
type WorkerPoolProperties = original.WorkerPoolProperties
type WorkerPoolCollection = original.WorkerPoolCollection
type ProviderClient = original.ProviderClient
type RecommendationsClient = original.RecommendationsClient
type CertificatesClient = original.CertificatesClient
type DomainsClient = original.DomainsClient
type GlobalResourceGroupsClient = original.GlobalResourceGroupsClient

func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDomainsClient(subscriptionID string) DomainsClient {
	return original.NewDomainsClient(subscriptionID)
}
func NewDomainsClientWithBaseURI(baseURI string, subscriptionID string) DomainsClient {
	return original.NewDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGlobalResourceGroupsClient(subscriptionID string) GlobalResourceGroupsClient {
	return original.NewGlobalResourceGroupsClient(subscriptionID)
}
func NewGlobalResourceGroupsClientWithBaseURI(baseURI string, subscriptionID string) GlobalResourceGroupsClient {
	return original.NewGlobalResourceGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewHostingEnvironmentsClient(subscriptionID string) HostingEnvironmentsClient {
	return original.NewHostingEnvironmentsClient(subscriptionID)
}
func NewHostingEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) HostingEnvironmentsClient {
	return original.NewHostingEnvironmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificateOrdersClient(subscriptionID string) CertificateOrdersClient {
	return original.NewCertificateOrdersClient(subscriptionID)
}
func NewCertificateOrdersClientWithBaseURI(baseURI string, subscriptionID string) CertificateOrdersClient {
	return original.NewCertificateOrdersClientWithBaseURI(baseURI, subscriptionID)
}
func NewClassicMobileServicesClient(subscriptionID string) ClassicMobileServicesClient {
	return original.NewClassicMobileServicesClient(subscriptionID)
}
func NewClassicMobileServicesClientWithBaseURI(baseURI string, subscriptionID string) ClassicMobileServicesClient {
	return original.NewClassicMobileServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewGlobalClient(subscriptionID string) GlobalClient {
	return original.NewGlobalClient(subscriptionID)
}
func NewGlobalClientWithBaseURI(baseURI string, subscriptionID string) GlobalClient {
	return original.NewGlobalClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerFarmsClient(subscriptionID string) ServerFarmsClient {
	return original.NewServerFarmsClient(subscriptionID)
}
func NewServerFarmsClientWithBaseURI(baseURI string, subscriptionID string) ServerFarmsClient {
	return original.NewServerFarmsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSitesClient(subscriptionID string) SitesClient {
	return original.NewSitesClient(subscriptionID)
}
func NewSitesClientWithBaseURI(baseURI string, subscriptionID string) SitesClient {
	return original.NewSitesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopLevelDomainsClient(subscriptionID string) TopLevelDomainsClient {
	return original.NewTopLevelDomainsClient(subscriptionID)
}
func NewTopLevelDomainsClientWithBaseURI(baseURI string, subscriptionID string) TopLevelDomainsClient {
	return original.NewTopLevelDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGlobalCertificateOrderClient(subscriptionID string) GlobalCertificateOrderClient {
	return original.NewGlobalCertificateOrderClient(subscriptionID)
}
func NewGlobalCertificateOrderClientWithBaseURI(baseURI string, subscriptionID string) GlobalCertificateOrderClient {
	return original.NewGlobalCertificateOrderClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedHostingEnvironmentsClient(subscriptionID string) ManagedHostingEnvironmentsClient {
	return original.NewManagedHostingEnvironmentsClient(subscriptionID)
}
func NewManagedHostingEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) ManagedHostingEnvironmentsClient {
	return original.NewManagedHostingEnvironmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClient(subscriptionID)
}
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGlobalDomainRegistrationClient(subscriptionID string) GlobalDomainRegistrationClient {
	return original.NewGlobalDomainRegistrationClient(subscriptionID)
}
func NewGlobalDomainRegistrationClientWithBaseURI(baseURI string, subscriptionID string) GlobalDomainRegistrationClient {
	return original.NewGlobalDomainRegistrationClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderClient(subscriptionID string) ProviderClient {
	return original.NewProviderClient(subscriptionID)
}
func NewProviderClientWithBaseURI(baseURI string, subscriptionID string) ProviderClient {
	return original.NewProviderClientWithBaseURI(baseURI, subscriptionID)
}
