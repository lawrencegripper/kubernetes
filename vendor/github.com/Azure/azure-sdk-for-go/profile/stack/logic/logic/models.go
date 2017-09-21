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

package logic

import original "github.com/lawrencegripper/azure-sdk-for-go/service/logic/management/2016-06-01/logic"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type WorkflowTriggerHistoriesClient = original.WorkflowTriggerHistoriesClient
type CertificatesClient = original.CertificatesClient
type MapsClient = original.MapsClient
type SchemasClient = original.SchemasClient
type SessionsClient = original.SessionsClient
type WorkflowRunActionsClient = original.WorkflowRunActionsClient
type IntegrationAccountsClient = original.IntegrationAccountsClient
type AgreementType = original.AgreementType

const (
	AS2		AgreementType	= original.AS2
	Edifact		AgreementType	= original.Edifact
	NotSpecified	AgreementType	= original.NotSpecified
	X12		AgreementType	= original.X12
)

type DayOfWeek = original.DayOfWeek

const (
	Friday		DayOfWeek	= original.Friday
	Monday		DayOfWeek	= original.Monday
	Saturday	DayOfWeek	= original.Saturday
	Sunday		DayOfWeek	= original.Sunday
	Thursday	DayOfWeek	= original.Thursday
	Tuesday		DayOfWeek	= original.Tuesday
	Wednesday	DayOfWeek	= original.Wednesday
)

type DaysOfWeek = original.DaysOfWeek

const (
	DaysOfWeekFriday	DaysOfWeek	= original.DaysOfWeekFriday
	DaysOfWeekMonday	DaysOfWeek	= original.DaysOfWeekMonday
	DaysOfWeekSaturday	DaysOfWeek	= original.DaysOfWeekSaturday
	DaysOfWeekSunday	DaysOfWeek	= original.DaysOfWeekSunday
	DaysOfWeekThursday	DaysOfWeek	= original.DaysOfWeekThursday
	DaysOfWeekTuesday	DaysOfWeek	= original.DaysOfWeekTuesday
	DaysOfWeekWednesday	DaysOfWeek	= original.DaysOfWeekWednesday
)

type EdifactCharacterSet = original.EdifactCharacterSet

const (
	EdifactCharacterSetKECA		EdifactCharacterSet	= original.EdifactCharacterSetKECA
	EdifactCharacterSetNotSpecified	EdifactCharacterSet	= original.EdifactCharacterSetNotSpecified
	EdifactCharacterSetUNOA		EdifactCharacterSet	= original.EdifactCharacterSetUNOA
	EdifactCharacterSetUNOB		EdifactCharacterSet	= original.EdifactCharacterSetUNOB
	EdifactCharacterSetUNOC		EdifactCharacterSet	= original.EdifactCharacterSetUNOC
	EdifactCharacterSetUNOD		EdifactCharacterSet	= original.EdifactCharacterSetUNOD
	EdifactCharacterSetUNOE		EdifactCharacterSet	= original.EdifactCharacterSetUNOE
	EdifactCharacterSetUNOF		EdifactCharacterSet	= original.EdifactCharacterSetUNOF
	EdifactCharacterSetUNOG		EdifactCharacterSet	= original.EdifactCharacterSetUNOG
	EdifactCharacterSetUNOH		EdifactCharacterSet	= original.EdifactCharacterSetUNOH
	EdifactCharacterSetUNOI		EdifactCharacterSet	= original.EdifactCharacterSetUNOI
	EdifactCharacterSetUNOJ		EdifactCharacterSet	= original.EdifactCharacterSetUNOJ
	EdifactCharacterSetUNOK		EdifactCharacterSet	= original.EdifactCharacterSetUNOK
	EdifactCharacterSetUNOX		EdifactCharacterSet	= original.EdifactCharacterSetUNOX
	EdifactCharacterSetUNOY		EdifactCharacterSet	= original.EdifactCharacterSetUNOY
)

type EdifactDecimalIndicator = original.EdifactDecimalIndicator

const (
	EdifactDecimalIndicatorComma		EdifactDecimalIndicator	= original.EdifactDecimalIndicatorComma
	EdifactDecimalIndicatorDecimal		EdifactDecimalIndicator	= original.EdifactDecimalIndicatorDecimal
	EdifactDecimalIndicatorNotSpecified	EdifactDecimalIndicator	= original.EdifactDecimalIndicatorNotSpecified
)

type EncryptionAlgorithm = original.EncryptionAlgorithm

const (
	EncryptionAlgorithmAES128	EncryptionAlgorithm	= original.EncryptionAlgorithmAES128
	EncryptionAlgorithmAES192	EncryptionAlgorithm	= original.EncryptionAlgorithmAES192
	EncryptionAlgorithmAES256	EncryptionAlgorithm	= original.EncryptionAlgorithmAES256
	EncryptionAlgorithmDES3		EncryptionAlgorithm	= original.EncryptionAlgorithmDES3
	EncryptionAlgorithmNone		EncryptionAlgorithm	= original.EncryptionAlgorithmNone
	EncryptionAlgorithmNotSpecified	EncryptionAlgorithm	= original.EncryptionAlgorithmNotSpecified
	EncryptionAlgorithmRC2		EncryptionAlgorithm	= original.EncryptionAlgorithmRC2
)

type HashingAlgorithm = original.HashingAlgorithm

const (
	HashingAlgorithmMD5		HashingAlgorithm	= original.HashingAlgorithmMD5
	HashingAlgorithmNone		HashingAlgorithm	= original.HashingAlgorithmNone
	HashingAlgorithmNotSpecified	HashingAlgorithm	= original.HashingAlgorithmNotSpecified
	HashingAlgorithmSHA1		HashingAlgorithm	= original.HashingAlgorithmSHA1
	HashingAlgorithmSHA2256		HashingAlgorithm	= original.HashingAlgorithmSHA2256
	HashingAlgorithmSHA2384		HashingAlgorithm	= original.HashingAlgorithmSHA2384
	HashingAlgorithmSHA2512		HashingAlgorithm	= original.HashingAlgorithmSHA2512
)

type IntegrationAccountSkuName = original.IntegrationAccountSkuName

const (
	IntegrationAccountSkuNameFree		IntegrationAccountSkuName	= original.IntegrationAccountSkuNameFree
	IntegrationAccountSkuNameNotSpecified	IntegrationAccountSkuName	= original.IntegrationAccountSkuNameNotSpecified
	IntegrationAccountSkuNameStandard	IntegrationAccountSkuName	= original.IntegrationAccountSkuNameStandard
)

type KeyType = original.KeyType

const (
	KeyTypeNotSpecified	KeyType	= original.KeyTypeNotSpecified
	KeyTypePrimary		KeyType	= original.KeyTypePrimary
	KeyTypeSecondary	KeyType	= original.KeyTypeSecondary
)

type MapType = original.MapType

const (
	MapTypeNotSpecified	MapType	= original.MapTypeNotSpecified
	MapTypeXslt		MapType	= original.MapTypeXslt
)

type MessageFilterType = original.MessageFilterType

const (
	MessageFilterTypeExclude	MessageFilterType	= original.MessageFilterTypeExclude
	MessageFilterTypeInclude	MessageFilterType	= original.MessageFilterTypeInclude
	MessageFilterTypeNotSpecified	MessageFilterType	= original.MessageFilterTypeNotSpecified
)

type ParameterType = original.ParameterType

const (
	ParameterTypeArray		ParameterType	= original.ParameterTypeArray
	ParameterTypeBool		ParameterType	= original.ParameterTypeBool
	ParameterTypeFloat		ParameterType	= original.ParameterTypeFloat
	ParameterTypeInt		ParameterType	= original.ParameterTypeInt
	ParameterTypeNotSpecified	ParameterType	= original.ParameterTypeNotSpecified
	ParameterTypeObject		ParameterType	= original.ParameterTypeObject
	ParameterTypeSecureObject	ParameterType	= original.ParameterTypeSecureObject
	ParameterTypeSecureString	ParameterType	= original.ParameterTypeSecureString
	ParameterTypeString		ParameterType	= original.ParameterTypeString
)

type PartnerType = original.PartnerType

const (
	PartnerTypeB2B		PartnerType	= original.PartnerTypeB2B
	PartnerTypeNotSpecified	PartnerType	= original.PartnerTypeNotSpecified
)

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	RecurrenceFrequencyDay		RecurrenceFrequency	= original.RecurrenceFrequencyDay
	RecurrenceFrequencyHour		RecurrenceFrequency	= original.RecurrenceFrequencyHour
	RecurrenceFrequencyMinute	RecurrenceFrequency	= original.RecurrenceFrequencyMinute
	RecurrenceFrequencyMonth	RecurrenceFrequency	= original.RecurrenceFrequencyMonth
	RecurrenceFrequencyNotSpecified	RecurrenceFrequency	= original.RecurrenceFrequencyNotSpecified
	RecurrenceFrequencySecond	RecurrenceFrequency	= original.RecurrenceFrequencySecond
	RecurrenceFrequencyWeek		RecurrenceFrequency	= original.RecurrenceFrequencyWeek
	RecurrenceFrequencyYear		RecurrenceFrequency	= original.RecurrenceFrequencyYear
)

type SchemaType = original.SchemaType

const (
	SchemaTypeNotSpecified	SchemaType	= original.SchemaTypeNotSpecified
	SchemaTypeXML		SchemaType	= original.SchemaTypeXML
)

type SegmentTerminatorSuffix = original.SegmentTerminatorSuffix

const (
	SegmentTerminatorSuffixCR		SegmentTerminatorSuffix	= original.SegmentTerminatorSuffixCR
	SegmentTerminatorSuffixCRLF		SegmentTerminatorSuffix	= original.SegmentTerminatorSuffixCRLF
	SegmentTerminatorSuffixLF		SegmentTerminatorSuffix	= original.SegmentTerminatorSuffixLF
	SegmentTerminatorSuffixNone		SegmentTerminatorSuffix	= original.SegmentTerminatorSuffixNone
	SegmentTerminatorSuffixNotSpecified	SegmentTerminatorSuffix	= original.SegmentTerminatorSuffixNotSpecified
)

type SigningAlgorithm = original.SigningAlgorithm

const (
	SigningAlgorithmDefault		SigningAlgorithm	= original.SigningAlgorithmDefault
	SigningAlgorithmNotSpecified	SigningAlgorithm	= original.SigningAlgorithmNotSpecified
	SigningAlgorithmSHA1		SigningAlgorithm	= original.SigningAlgorithmSHA1
	SigningAlgorithmSHA2256		SigningAlgorithm	= original.SigningAlgorithmSHA2256
	SigningAlgorithmSHA2384		SigningAlgorithm	= original.SigningAlgorithmSHA2384
	SigningAlgorithmSHA2512		SigningAlgorithm	= original.SigningAlgorithmSHA2512
)

type SkuName = original.SkuName

const (
	SkuNameBasic		SkuName	= original.SkuNameBasic
	SkuNameFree		SkuName	= original.SkuNameFree
	SkuNameNotSpecified	SkuName	= original.SkuNameNotSpecified
	SkuNamePremium		SkuName	= original.SkuNamePremium
	SkuNameShared		SkuName	= original.SkuNameShared
	SkuNameStandard		SkuName	= original.SkuNameStandard
)

type TrailingSeparatorPolicy = original.TrailingSeparatorPolicy

const (
	TrailingSeparatorPolicyMandatory	TrailingSeparatorPolicy	= original.TrailingSeparatorPolicyMandatory
	TrailingSeparatorPolicyNotAllowed	TrailingSeparatorPolicy	= original.TrailingSeparatorPolicyNotAllowed
	TrailingSeparatorPolicyNotSpecified	TrailingSeparatorPolicy	= original.TrailingSeparatorPolicyNotSpecified
	TrailingSeparatorPolicyOptional		TrailingSeparatorPolicy	= original.TrailingSeparatorPolicyOptional
)

type UsageIndicator = original.UsageIndicator

const (
	UsageIndicatorInformation	UsageIndicator	= original.UsageIndicatorInformation
	UsageIndicatorNotSpecified	UsageIndicator	= original.UsageIndicatorNotSpecified
	UsageIndicatorProduction	UsageIndicator	= original.UsageIndicatorProduction
	UsageIndicatorTest		UsageIndicator	= original.UsageIndicatorTest
)

type WorkflowProvisioningState = original.WorkflowProvisioningState

const (
	WorkflowProvisioningStateAccepted	WorkflowProvisioningState	= original.WorkflowProvisioningStateAccepted
	WorkflowProvisioningStateCanceled	WorkflowProvisioningState	= original.WorkflowProvisioningStateCanceled
	WorkflowProvisioningStateCompleted	WorkflowProvisioningState	= original.WorkflowProvisioningStateCompleted
	WorkflowProvisioningStateCreated	WorkflowProvisioningState	= original.WorkflowProvisioningStateCreated
	WorkflowProvisioningStateCreating	WorkflowProvisioningState	= original.WorkflowProvisioningStateCreating
	WorkflowProvisioningStateDeleted	WorkflowProvisioningState	= original.WorkflowProvisioningStateDeleted
	WorkflowProvisioningStateDeleting	WorkflowProvisioningState	= original.WorkflowProvisioningStateDeleting
	WorkflowProvisioningStateFailed		WorkflowProvisioningState	= original.WorkflowProvisioningStateFailed
	WorkflowProvisioningStateMoving		WorkflowProvisioningState	= original.WorkflowProvisioningStateMoving
	WorkflowProvisioningStateNotSpecified	WorkflowProvisioningState	= original.WorkflowProvisioningStateNotSpecified
	WorkflowProvisioningStateReady		WorkflowProvisioningState	= original.WorkflowProvisioningStateReady
	WorkflowProvisioningStateRegistered	WorkflowProvisioningState	= original.WorkflowProvisioningStateRegistered
	WorkflowProvisioningStateRegistering	WorkflowProvisioningState	= original.WorkflowProvisioningStateRegistering
	WorkflowProvisioningStateRunning	WorkflowProvisioningState	= original.WorkflowProvisioningStateRunning
	WorkflowProvisioningStateSucceeded	WorkflowProvisioningState	= original.WorkflowProvisioningStateSucceeded
	WorkflowProvisioningStateUnregistered	WorkflowProvisioningState	= original.WorkflowProvisioningStateUnregistered
	WorkflowProvisioningStateUnregistering	WorkflowProvisioningState	= original.WorkflowProvisioningStateUnregistering
	WorkflowProvisioningStateUpdating	WorkflowProvisioningState	= original.WorkflowProvisioningStateUpdating
)

type WorkflowState = original.WorkflowState

const (
	WorkflowStateCompleted		WorkflowState	= original.WorkflowStateCompleted
	WorkflowStateDeleted		WorkflowState	= original.WorkflowStateDeleted
	WorkflowStateDisabled		WorkflowState	= original.WorkflowStateDisabled
	WorkflowStateEnabled		WorkflowState	= original.WorkflowStateEnabled
	WorkflowStateNotSpecified	WorkflowState	= original.WorkflowStateNotSpecified
	WorkflowStateSuspended		WorkflowState	= original.WorkflowStateSuspended
)

type WorkflowStatus = original.WorkflowStatus

const (
	WorkflowStatusAborted		WorkflowStatus	= original.WorkflowStatusAborted
	WorkflowStatusCancelled		WorkflowStatus	= original.WorkflowStatusCancelled
	WorkflowStatusFailed		WorkflowStatus	= original.WorkflowStatusFailed
	WorkflowStatusFaulted		WorkflowStatus	= original.WorkflowStatusFaulted
	WorkflowStatusIgnored		WorkflowStatus	= original.WorkflowStatusIgnored
	WorkflowStatusNotSpecified	WorkflowStatus	= original.WorkflowStatusNotSpecified
	WorkflowStatusPaused		WorkflowStatus	= original.WorkflowStatusPaused
	WorkflowStatusRunning		WorkflowStatus	= original.WorkflowStatusRunning
	WorkflowStatusSkipped		WorkflowStatus	= original.WorkflowStatusSkipped
	WorkflowStatusSucceeded		WorkflowStatus	= original.WorkflowStatusSucceeded
	WorkflowStatusSuspended		WorkflowStatus	= original.WorkflowStatusSuspended
	WorkflowStatusTimedOut		WorkflowStatus	= original.WorkflowStatusTimedOut
	WorkflowStatusWaiting		WorkflowStatus	= original.WorkflowStatusWaiting
)

type WorkflowTriggerProvisioningState = original.WorkflowTriggerProvisioningState

const (
	WorkflowTriggerProvisioningStateAccepted	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateAccepted
	WorkflowTriggerProvisioningStateCanceled	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateCanceled
	WorkflowTriggerProvisioningStateCompleted	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateCompleted
	WorkflowTriggerProvisioningStateCreated		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateCreated
	WorkflowTriggerProvisioningStateCreating	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateCreating
	WorkflowTriggerProvisioningStateDeleted		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateDeleted
	WorkflowTriggerProvisioningStateDeleting	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateDeleting
	WorkflowTriggerProvisioningStateFailed		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateFailed
	WorkflowTriggerProvisioningStateMoving		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateMoving
	WorkflowTriggerProvisioningStateNotSpecified	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateNotSpecified
	WorkflowTriggerProvisioningStateReady		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateReady
	WorkflowTriggerProvisioningStateRegistered	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateRegistered
	WorkflowTriggerProvisioningStateRegistering	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateRegistering
	WorkflowTriggerProvisioningStateRunning		WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateRunning
	WorkflowTriggerProvisioningStateSucceeded	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateSucceeded
	WorkflowTriggerProvisioningStateUnregistered	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateUnregistered
	WorkflowTriggerProvisioningStateUnregistering	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateUnregistering
	WorkflowTriggerProvisioningStateUpdating	WorkflowTriggerProvisioningState	= original.WorkflowTriggerProvisioningStateUpdating
)

type X12CharacterSet = original.X12CharacterSet

const (
	X12CharacterSetBasic		X12CharacterSet	= original.X12CharacterSetBasic
	X12CharacterSetExtended		X12CharacterSet	= original.X12CharacterSetExtended
	X12CharacterSetNotSpecified	X12CharacterSet	= original.X12CharacterSetNotSpecified
	X12CharacterSetUTF8		X12CharacterSet	= original.X12CharacterSetUTF8
)

type X12DateFormat = original.X12DateFormat

const (
	X12DateFormatCCYYMMDD		X12DateFormat	= original.X12DateFormatCCYYMMDD
	X12DateFormatNotSpecified	X12DateFormat	= original.X12DateFormatNotSpecified
	X12DateFormatYYMMDD		X12DateFormat	= original.X12DateFormatYYMMDD
)

type X12TimeFormat = original.X12TimeFormat

const (
	X12TimeFormatHHMM		X12TimeFormat	= original.X12TimeFormatHHMM
	X12TimeFormatHHMMSS		X12TimeFormat	= original.X12TimeFormatHHMMSS
	X12TimeFormatHHMMSSd		X12TimeFormat	= original.X12TimeFormatHHMMSSd
	X12TimeFormatHHMMSSdd		X12TimeFormat	= original.X12TimeFormatHHMMSSdd
	X12TimeFormatNotSpecified	X12TimeFormat	= original.X12TimeFormatNotSpecified
)

type AgreementContent = original.AgreementContent
type AS2AcknowledgementConnectionSettings = original.AS2AcknowledgementConnectionSettings
type AS2AgreementContent = original.AS2AgreementContent
type AS2EnvelopeSettings = original.AS2EnvelopeSettings
type AS2ErrorSettings = original.AS2ErrorSettings
type AS2MdnSettings = original.AS2MdnSettings
type AS2MessageConnectionSettings = original.AS2MessageConnectionSettings
type AS2OneWayAgreement = original.AS2OneWayAgreement
type AS2ProtocolSettings = original.AS2ProtocolSettings
type AS2SecuritySettings = original.AS2SecuritySettings
type AS2ValidationSettings = original.AS2ValidationSettings
type B2BPartnerContent = original.B2BPartnerContent
type BusinessIdentity = original.BusinessIdentity
type CallbackURL = original.CallbackURL
type ContentHash = original.ContentHash
type ContentLink = original.ContentLink
type Correlation = original.Correlation
type EdifactAcknowledgementSettings = original.EdifactAcknowledgementSettings
type EdifactAgreementContent = original.EdifactAgreementContent
type EdifactDelimiterOverride = original.EdifactDelimiterOverride
type EdifactEnvelopeOverride = original.EdifactEnvelopeOverride
type EdifactEnvelopeSettings = original.EdifactEnvelopeSettings
type EdifactFramingSettings = original.EdifactFramingSettings
type EdifactMessageFilter = original.EdifactMessageFilter
type EdifactMessageIdentifier = original.EdifactMessageIdentifier
type EdifactOneWayAgreement = original.EdifactOneWayAgreement
type EdifactProcessingSettings = original.EdifactProcessingSettings
type EdifactProtocolSettings = original.EdifactProtocolSettings
type EdifactSchemaReference = original.EdifactSchemaReference
type EdifactValidationOverride = original.EdifactValidationOverride
type EdifactValidationSettings = original.EdifactValidationSettings
type ErrorProperties = original.ErrorProperties
type ErrorResponse = original.ErrorResponse
type GenerateUpgradedDefinitionParameters = original.GenerateUpgradedDefinitionParameters
type GetCallbackURLParameters = original.GetCallbackURLParameters
type IntegrationAccount = original.IntegrationAccount
type IntegrationAccountAgreement = original.IntegrationAccountAgreement
type IntegrationAccountAgreementFilter = original.IntegrationAccountAgreementFilter
type IntegrationAccountAgreementListResult = original.IntegrationAccountAgreementListResult
type IntegrationAccountAgreementProperties = original.IntegrationAccountAgreementProperties
type IntegrationAccountCertificate = original.IntegrationAccountCertificate
type IntegrationAccountCertificateListResult = original.IntegrationAccountCertificateListResult
type IntegrationAccountCertificateProperties = original.IntegrationAccountCertificateProperties
type IntegrationAccountListResult = original.IntegrationAccountListResult
type IntegrationAccountMap = original.IntegrationAccountMap
type IntegrationAccountMapFilter = original.IntegrationAccountMapFilter
type IntegrationAccountMapListResult = original.IntegrationAccountMapListResult
type IntegrationAccountMapProperties = original.IntegrationAccountMapProperties
type IntegrationAccountMapPropertiesParametersSchema = original.IntegrationAccountMapPropertiesParametersSchema
type IntegrationAccountPartner = original.IntegrationAccountPartner
type IntegrationAccountPartnerFilter = original.IntegrationAccountPartnerFilter
type IntegrationAccountPartnerListResult = original.IntegrationAccountPartnerListResult
type IntegrationAccountPartnerProperties = original.IntegrationAccountPartnerProperties
type IntegrationAccountSchema = original.IntegrationAccountSchema
type IntegrationAccountSchemaFilter = original.IntegrationAccountSchemaFilter
type IntegrationAccountSchemaListResult = original.IntegrationAccountSchemaListResult
type IntegrationAccountSchemaProperties = original.IntegrationAccountSchemaProperties
type IntegrationAccountSession = original.IntegrationAccountSession
type IntegrationAccountSessionFilter = original.IntegrationAccountSessionFilter
type IntegrationAccountSessionListResult = original.IntegrationAccountSessionListResult
type IntegrationAccountSessionProperties = original.IntegrationAccountSessionProperties
type IntegrationAccountSku = original.IntegrationAccountSku
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultKeyReferenceKeyVault = original.KeyVaultKeyReferenceKeyVault
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type PartnerContent = original.PartnerContent
type RecurrenceSchedule = original.RecurrenceSchedule
type RecurrenceScheduleOccurrence = original.RecurrenceScheduleOccurrence
type RegenerateActionParameter = original.RegenerateActionParameter
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type RetryHistory = original.RetryHistory
type SetObject = original.SetObject
type Sku = original.Sku
type SubResource = original.SubResource
type Workflow = original.Workflow
type WorkflowFilter = original.WorkflowFilter
type WorkflowListResult = original.WorkflowListResult
type WorkflowOutputParameter = original.WorkflowOutputParameter
type WorkflowParameter = original.WorkflowParameter
type WorkflowProperties = original.WorkflowProperties
type WorkflowRun = original.WorkflowRun
type WorkflowRunAction = original.WorkflowRunAction
type WorkflowRunActionFilter = original.WorkflowRunActionFilter
type WorkflowRunActionListResult = original.WorkflowRunActionListResult
type WorkflowRunActionProperties = original.WorkflowRunActionProperties
type WorkflowRunFilter = original.WorkflowRunFilter
type WorkflowRunListResult = original.WorkflowRunListResult
type WorkflowRunProperties = original.WorkflowRunProperties
type WorkflowRunTrigger = original.WorkflowRunTrigger
type WorkflowTrigger = original.WorkflowTrigger
type WorkflowTriggerCallbackURL = original.WorkflowTriggerCallbackURL
type WorkflowTriggerFilter = original.WorkflowTriggerFilter
type WorkflowTriggerHistory = original.WorkflowTriggerHistory
type WorkflowTriggerHistoryFilter = original.WorkflowTriggerHistoryFilter
type WorkflowTriggerHistoryListResult = original.WorkflowTriggerHistoryListResult
type WorkflowTriggerHistoryProperties = original.WorkflowTriggerHistoryProperties
type WorkflowTriggerListCallbackURLQueries = original.WorkflowTriggerListCallbackURLQueries
type WorkflowTriggerListResult = original.WorkflowTriggerListResult
type WorkflowTriggerProperties = original.WorkflowTriggerProperties
type WorkflowTriggerRecurrence = original.WorkflowTriggerRecurrence
type WorkflowVersion = original.WorkflowVersion
type WorkflowVersionListResult = original.WorkflowVersionListResult
type WorkflowVersionProperties = original.WorkflowVersionProperties
type X12AcknowledgementSettings = original.X12AcknowledgementSettings
type X12AgreementContent = original.X12AgreementContent
type X12DelimiterOverrides = original.X12DelimiterOverrides
type X12EnvelopeOverride = original.X12EnvelopeOverride
type X12EnvelopeSettings = original.X12EnvelopeSettings
type X12FramingSettings = original.X12FramingSettings
type X12MessageFilter = original.X12MessageFilter
type X12MessageIdentifier = original.X12MessageIdentifier
type X12OneWayAgreement = original.X12OneWayAgreement
type X12ProcessingSettings = original.X12ProcessingSettings
type X12ProtocolSettings = original.X12ProtocolSettings
type X12SchemaReference = original.X12SchemaReference
type X12SecuritySettings = original.X12SecuritySettings
type X12ValidationOverride = original.X12ValidationOverride
type X12ValidationSettings = original.X12ValidationSettings
type WorkflowsClient = original.WorkflowsClient
type WorkflowTriggersClient = original.WorkflowTriggersClient
type WorkflowVersionsClient = original.WorkflowVersionsClient
type AgreementsClient = original.AgreementsClient
type PartnersClient = original.PartnersClient
type WorkflowRunsClient = original.WorkflowRunsClient

func NewAgreementsClient(subscriptionID string) AgreementsClient {
	return original.NewAgreementsClient(subscriptionID)
}
func NewAgreementsClientWithBaseURI(baseURI string, subscriptionID string) AgreementsClient {
	return original.NewAgreementsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPartnersClient(subscriptionID string) PartnersClient {
	return original.NewPartnersClient(subscriptionID)
}
func NewPartnersClientWithBaseURI(baseURI string, subscriptionID string) PartnersClient {
	return original.NewPartnersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowRunsClient(subscriptionID string) WorkflowRunsClient {
	return original.NewWorkflowRunsClient(subscriptionID)
}
func NewWorkflowRunsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunsClient {
	return original.NewWorkflowRunsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowTriggerHistoriesClient(subscriptionID string) WorkflowTriggerHistoriesClient {
	return original.NewWorkflowTriggerHistoriesClient(subscriptionID)
}
func NewWorkflowTriggerHistoriesClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggerHistoriesClient {
	return original.NewWorkflowTriggerHistoriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMapsClient(subscriptionID string) MapsClient {
	return original.NewMapsClient(subscriptionID)
}
func NewMapsClientWithBaseURI(baseURI string, subscriptionID string) MapsClient {
	return original.NewMapsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSchemasClient(subscriptionID string) SchemasClient {
	return original.NewSchemasClient(subscriptionID)
}
func NewSchemasClientWithBaseURI(baseURI string, subscriptionID string) SchemasClient {
	return original.NewSchemasClientWithBaseURI(baseURI, subscriptionID)
}
func NewSessionsClient(subscriptionID string) SessionsClient {
	return original.NewSessionsClient(subscriptionID)
}
func NewSessionsClientWithBaseURI(baseURI string, subscriptionID string) SessionsClient {
	return original.NewSessionsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewWorkflowRunActionsClient(subscriptionID string) WorkflowRunActionsClient {
	return original.NewWorkflowRunActionsClient(subscriptionID)
}
func NewWorkflowRunActionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunActionsClient {
	return original.NewWorkflowRunActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountsClient(subscriptionID string) IntegrationAccountsClient {
	return original.NewIntegrationAccountsClient(subscriptionID)
}
func NewIntegrationAccountsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountsClient {
	return original.NewIntegrationAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClient(subscriptionID)
}
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowTriggersClient(subscriptionID string) WorkflowTriggersClient {
	return original.NewWorkflowTriggersClient(subscriptionID)
}
func NewWorkflowTriggersClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggersClient {
	return original.NewWorkflowTriggersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowVersionsClient(subscriptionID string) WorkflowVersionsClient {
	return original.NewWorkflowVersionsClient(subscriptionID)
}
func NewWorkflowVersionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowVersionsClient {
	return original.NewWorkflowVersionsClientWithBaseURI(baseURI, subscriptionID)
}
