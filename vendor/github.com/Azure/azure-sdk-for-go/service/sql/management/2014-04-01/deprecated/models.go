package deprecated

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/uuid"
)

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// Copy specifies the copy state for create mode.
	Copy CreateMode = "Copy"
	// Default specifies the default state for create mode.
	Default CreateMode = "Default"
	// NonReadableSecondary specifies the non readable secondary state for
	// create mode.
	NonReadableSecondary CreateMode = "NonReadableSecondary"
	// OnlineSecondary specifies the online secondary state for create mode.
	OnlineSecondary CreateMode = "OnlineSecondary"
	// PointInTimeRestore specifies the point in time restore state for create
	// mode.
	PointInTimeRestore CreateMode = "PointInTimeRestore"
	// Recovery specifies the recovery state for create mode.
	Recovery CreateMode = "Recovery"
	// Restore specifies the restore state for create mode.
	Restore CreateMode = "Restore"
	// RestoreLongTermRetentionBackup specifies the restore long term retention
	// backup state for create mode.
	RestoreLongTermRetentionBackup CreateMode = "RestoreLongTermRetentionBackup"
)

// DatabaseEdition enumerates the values for database edition.
type DatabaseEdition string

const (
	// Basic specifies the basic state for database edition.
	Basic DatabaseEdition = "Basic"
	// Business specifies the business state for database edition.
	Business DatabaseEdition = "Business"
	// DataWarehouse specifies the data warehouse state for database edition.
	DataWarehouse DatabaseEdition = "DataWarehouse"
	// Free specifies the free state for database edition.
	Free DatabaseEdition = "Free"
	// Premium specifies the premium state for database edition.
	Premium DatabaseEdition = "Premium"
	// Standard specifies the standard state for database edition.
	Standard DatabaseEdition = "Standard"
	// Stretch specifies the stretch state for database edition.
	Stretch DatabaseEdition = "Stretch"
	// System specifies the system state for database edition.
	System DatabaseEdition = "System"
	// System2 specifies the system 2 state for database edition.
	System2 DatabaseEdition = "System2"
	// Web specifies the web state for database edition.
	Web DatabaseEdition = "Web"
)

// ElasticPoolEdition enumerates the values for elastic pool edition.
type ElasticPoolEdition string

const (
	// ElasticPoolEditionBasic specifies the elastic pool edition basic state
	// for elastic pool edition.
	ElasticPoolEditionBasic ElasticPoolEdition = "Basic"
	// ElasticPoolEditionPremium specifies the elastic pool edition premium
	// state for elastic pool edition.
	ElasticPoolEditionPremium ElasticPoolEdition = "Premium"
	// ElasticPoolEditionStandard specifies the elastic pool edition standard
	// state for elastic pool edition.
	ElasticPoolEditionStandard ElasticPoolEdition = "Standard"
)

// ElasticPoolState enumerates the values for elastic pool state.
type ElasticPoolState string

const (
	// Creating specifies the creating state for elastic pool state.
	Creating ElasticPoolState = "Creating"
	// Disabled specifies the disabled state for elastic pool state.
	Disabled ElasticPoolState = "Disabled"
	// Ready specifies the ready state for elastic pool state.
	Ready ElasticPoolState = "Ready"
)

// ReadScale enumerates the values for read scale.
type ReadScale string

const (
	// ReadScaleDisabled specifies the read scale disabled state for read
	// scale.
	ReadScaleDisabled ReadScale = "Disabled"
	// ReadScaleEnabled specifies the read scale enabled state for read scale.
	ReadScaleEnabled ReadScale = "Enabled"
)

// RecommendedIndexAction enumerates the values for recommended index action.
type RecommendedIndexAction string

const (
	// Create specifies the create state for recommended index action.
	Create RecommendedIndexAction = "Create"
	// Drop specifies the drop state for recommended index action.
	Drop RecommendedIndexAction = "Drop"
	// Rebuild specifies the rebuild state for recommended index action.
	Rebuild RecommendedIndexAction = "Rebuild"
)

// RecommendedIndexState enumerates the values for recommended index state.
type RecommendedIndexState string

const (
	// Active specifies the active state for recommended index state.
	Active RecommendedIndexState = "Active"
	// Blocked specifies the blocked state for recommended index state.
	Blocked RecommendedIndexState = "Blocked"
	// Executing specifies the executing state for recommended index state.
	Executing RecommendedIndexState = "Executing"
	// Expired specifies the expired state for recommended index state.
	Expired RecommendedIndexState = "Expired"
	// Ignored specifies the ignored state for recommended index state.
	Ignored RecommendedIndexState = "Ignored"
	// Pending specifies the pending state for recommended index state.
	Pending RecommendedIndexState = "Pending"
	// PendingRevert specifies the pending revert state for recommended index
	// state.
	PendingRevert RecommendedIndexState = "Pending Revert"
	// Reverted specifies the reverted state for recommended index state.
	Reverted RecommendedIndexState = "Reverted"
	// Reverting specifies the reverting state for recommended index state.
	Reverting RecommendedIndexState = "Reverting"
	// Success specifies the success state for recommended index state.
	Success RecommendedIndexState = "Success"
	// Verifying specifies the verifying state for recommended index state.
	Verifying RecommendedIndexState = "Verifying"
)

// RecommendedIndexType enumerates the values for recommended index type.
type RecommendedIndexType string

const (
	// CLUSTERED specifies the clustered state for recommended index type.
	CLUSTERED RecommendedIndexType = "CLUSTERED"
	// CLUSTEREDCOLUMNSTORE specifies the clusteredcolumnstore state for
	// recommended index type.
	CLUSTEREDCOLUMNSTORE RecommendedIndexType = "CLUSTERED COLUMNSTORE"
	// COLUMNSTORE specifies the columnstore state for recommended index type.
	COLUMNSTORE RecommendedIndexType = "COLUMNSTORE"
	// NONCLUSTERED specifies the nonclustered state for recommended index
	// type.
	NONCLUSTERED RecommendedIndexType = "NONCLUSTERED"
)

// SampleName enumerates the values for sample name.
type SampleName string

const (
	// AdventureWorksLT specifies the adventure works lt state for sample name.
	AdventureWorksLT SampleName = "AdventureWorksLT"
)

// ServiceObjectiveName enumerates the values for service objective name.
type ServiceObjectiveName string

const (
	// ServiceObjectiveNameBasic specifies the service objective name basic
	// state for service objective name.
	ServiceObjectiveNameBasic ServiceObjectiveName = "Basic"
	// ServiceObjectiveNameElasticPool specifies the service objective name
	// elastic pool state for service objective name.
	ServiceObjectiveNameElasticPool ServiceObjectiveName = "ElasticPool"
	// ServiceObjectiveNameP1 specifies the service objective name p1 state for
	// service objective name.
	ServiceObjectiveNameP1 ServiceObjectiveName = "P1"
	// ServiceObjectiveNameP11 specifies the service objective name p11 state
	// for service objective name.
	ServiceObjectiveNameP11 ServiceObjectiveName = "P11"
	// ServiceObjectiveNameP15 specifies the service objective name p15 state
	// for service objective name.
	ServiceObjectiveNameP15 ServiceObjectiveName = "P15"
	// ServiceObjectiveNameP2 specifies the service objective name p2 state for
	// service objective name.
	ServiceObjectiveNameP2 ServiceObjectiveName = "P2"
	// ServiceObjectiveNameP3 specifies the service objective name p3 state for
	// service objective name.
	ServiceObjectiveNameP3 ServiceObjectiveName = "P3"
	// ServiceObjectiveNameP4 specifies the service objective name p4 state for
	// service objective name.
	ServiceObjectiveNameP4 ServiceObjectiveName = "P4"
	// ServiceObjectiveNameP6 specifies the service objective name p6 state for
	// service objective name.
	ServiceObjectiveNameP6 ServiceObjectiveName = "P6"
	// ServiceObjectiveNameS0 specifies the service objective name s0 state for
	// service objective name.
	ServiceObjectiveNameS0 ServiceObjectiveName = "S0"
	// ServiceObjectiveNameS1 specifies the service objective name s1 state for
	// service objective name.
	ServiceObjectiveNameS1 ServiceObjectiveName = "S1"
	// ServiceObjectiveNameS2 specifies the service objective name s2 state for
	// service objective name.
	ServiceObjectiveNameS2 ServiceObjectiveName = "S2"
	// ServiceObjectiveNameS3 specifies the service objective name s3 state for
	// service objective name.
	ServiceObjectiveNameS3 ServiceObjectiveName = "S3"
	// ServiceObjectiveNameSystem specifies the service objective name system
	// state for service objective name.
	ServiceObjectiveNameSystem ServiceObjectiveName = "System"
	// ServiceObjectiveNameSystem2 specifies the service objective name system
	// 2 state for service objective name.
	ServiceObjectiveNameSystem2 ServiceObjectiveName = "System2"
)

// TransparentDataEncryptionActivityStatus enumerates the values for
// transparent data encryption activity status.
type TransparentDataEncryptionActivityStatus string

const (
	// Decrypting specifies the decrypting state for transparent data
	// encryption activity status.
	Decrypting TransparentDataEncryptionActivityStatus = "Decrypting"
	// Encrypting specifies the encrypting state for transparent data
	// encryption activity status.
	Encrypting TransparentDataEncryptionActivityStatus = "Encrypting"
)

// TransparentDataEncryptionStatus enumerates the values for transparent data
// encryption status.
type TransparentDataEncryptionStatus string

const (
	// TransparentDataEncryptionStatusDisabled specifies the transparent data
	// encryption status disabled state for transparent data encryption status.
	TransparentDataEncryptionStatusDisabled TransparentDataEncryptionStatus = "Disabled"
	// TransparentDataEncryptionStatusEnabled specifies the transparent data
	// encryption status enabled state for transparent data encryption status.
	TransparentDataEncryptionStatusEnabled TransparentDataEncryptionStatus = "Enabled"
)

// Database is represents a database.
type Database struct {
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Kind                *string             `json:"kind,omitempty"`
	*DatabaseProperties `json:"properties,omitempty"`
}

// DatabaseProperties is represents the properties of a database.
type DatabaseProperties struct {
	Collation                               *string                      `json:"collation,omitempty"`
	CreationDate                            *date.Time                   `json:"creationDate,omitempty"`
	ContainmentState                        *int64                       `json:"containmentState,omitempty"`
	CurrentServiceObjectiveID               *uuid.UUID                   `json:"currentServiceObjectiveId,omitempty"`
	DatabaseID                              *uuid.UUID                   `json:"databaseId,omitempty"`
	EarliestRestoreDate                     *date.Time                   `json:"earliestRestoreDate,omitempty"`
	CreateMode                              CreateMode                   `json:"createMode,omitempty"`
	SourceDatabaseID                        *string                      `json:"sourceDatabaseId,omitempty"`
	SourceDatabaseDeletionDate              *date.Time                   `json:"sourceDatabaseDeletionDate,omitempty"`
	RestorePointInTime                      *date.Time                   `json:"restorePointInTime,omitempty"`
	RecoveryServicesRecoveryPointResourceID *string                      `json:"recoveryServicesRecoveryPointResourceId,omitempty"`
	Edition                                 DatabaseEdition              `json:"edition,omitempty"`
	MaxSizeBytes                            *string                      `json:"maxSizeBytes,omitempty"`
	RequestedServiceObjectiveID             *uuid.UUID                   `json:"requestedServiceObjectiveId,omitempty"`
	RequestedServiceObjectiveName           ServiceObjectiveName         `json:"requestedServiceObjectiveName,omitempty"`
	ServiceLevelObjective                   ServiceObjectiveName         `json:"serviceLevelObjective,omitempty"`
	Status                                  *string                      `json:"status,omitempty"`
	ElasticPoolName                         *string                      `json:"elasticPoolName,omitempty"`
	DefaultSecondaryLocation                *string                      `json:"defaultSecondaryLocation,omitempty"`
	ServiceTierAdvisors                     *[]ServiceTierAdvisor        `json:"serviceTierAdvisors,omitempty"`
	TransparentDataEncryption               *[]TransparentDataEncryption `json:"transparentDataEncryption,omitempty"`
	RecommendedIndex                        *[]RecommendedIndex          `json:"recommendedIndex,omitempty"`
	FailoverGroupID                         *string                      `json:"failoverGroupId,omitempty"`
	ReadScale                               ReadScale                    `json:"readScale,omitempty"`
	SampleName                              SampleName                   `json:"sampleName,omitempty"`
}

// DatabaseUpdate is represents a database update.
type DatabaseUpdate struct {
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	*DatabaseProperties `json:"properties,omitempty"`
}

// ElasticPool is represents a database elastic pool.
type ElasticPool struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	Location               *string             `json:"location,omitempty"`
	*ElasticPoolProperties `json:"properties,omitempty"`
	Kind                   *string `json:"kind,omitempty"`
}

// ElasticPoolActivity is represents the activity on an elastic pool.
type ElasticPoolActivity struct {
	ID                             *string `json:"id,omitempty"`
	Name                           *string `json:"name,omitempty"`
	Type                           *string `json:"type,omitempty"`
	Location                       *string `json:"location,omitempty"`
	*ElasticPoolActivityProperties `json:"properties,omitempty"`
}

// ElasticPoolActivityProperties is represents the properties of an elastic
// pool.
type ElasticPoolActivityProperties struct {
	EndTime                       *date.Time `json:"endTime,omitempty"`
	ErrorCode                     *int32     `json:"errorCode,omitempty"`
	ErrorMessage                  *string    `json:"errorMessage,omitempty"`
	ErrorSeverity                 *int32     `json:"errorSeverity,omitempty"`
	Operation                     *string    `json:"operation,omitempty"`
	OperationID                   *uuid.UUID `json:"operationId,omitempty"`
	PercentComplete               *int32     `json:"percentComplete,omitempty"`
	RequestedDatabaseDtuMax       *int32     `json:"requestedDatabaseDtuMax,omitempty"`
	RequestedDatabaseDtuMin       *int32     `json:"requestedDatabaseDtuMin,omitempty"`
	RequestedDtu                  *int32     `json:"requestedDtu,omitempty"`
	RequestedElasticPoolName      *string    `json:"requestedElasticPoolName,omitempty"`
	RequestedStorageLimitInGB     *int64     `json:"requestedStorageLimitInGB,omitempty"`
	ElasticPoolName               *string    `json:"elasticPoolName,omitempty"`
	ServerName                    *string    `json:"serverName,omitempty"`
	StartTime                     *date.Time `json:"startTime,omitempty"`
	State                         *string    `json:"state,omitempty"`
	RequestedStorageLimitInMB     *int32     `json:"requestedStorageLimitInMB,omitempty"`
	RequestedDatabaseDtuGuarantee *int32     `json:"requestedDatabaseDtuGuarantee,omitempty"`
	RequestedDatabaseDtuCap       *int32     `json:"requestedDatabaseDtuCap,omitempty"`
	RequestedDtuGuarantee         *int32     `json:"requestedDtuGuarantee,omitempty"`
}

// ElasticPoolDatabaseActivity is represents the activity on an elastic pool.
type ElasticPoolDatabaseActivity struct {
	ID                                     *string `json:"id,omitempty"`
	Name                                   *string `json:"name,omitempty"`
	Type                                   *string `json:"type,omitempty"`
	Location                               *string `json:"location,omitempty"`
	*ElasticPoolDatabaseActivityProperties `json:"properties,omitempty"`
}

// ElasticPoolDatabaseActivityProperties is represents the properties of an
// elastic pool database activity.
type ElasticPoolDatabaseActivityProperties struct {
	DatabaseName              *string    `json:"databaseName,omitempty"`
	EndTime                   *date.Time `json:"endTime,omitempty"`
	ErrorCode                 *int32     `json:"errorCode,omitempty"`
	ErrorMessage              *string    `json:"errorMessage,omitempty"`
	ErrorSeverity             *int32     `json:"errorSeverity,omitempty"`
	Operation                 *string    `json:"operation,omitempty"`
	OperationID               *uuid.UUID `json:"operationId,omitempty"`
	PercentComplete           *int32     `json:"percentComplete,omitempty"`
	RequestedElasticPoolName  *string    `json:"requestedElasticPoolName,omitempty"`
	CurrentElasticPoolName    *string    `json:"currentElasticPoolName,omitempty"`
	CurrentServiceObjective   *string    `json:"currentServiceObjective,omitempty"`
	RequestedServiceObjective *string    `json:"requestedServiceObjective,omitempty"`
	ServerName                *string    `json:"serverName,omitempty"`
	StartTime                 *date.Time `json:"startTime,omitempty"`
	State                     *string    `json:"state,omitempty"`
}

// ElasticPoolProperties is represents the properties of an elastic pool.
type ElasticPoolProperties struct {
	CreationDate   *date.Time         `json:"creationDate,omitempty"`
	State          ElasticPoolState   `json:"state,omitempty"`
	Edition        ElasticPoolEdition `json:"edition,omitempty"`
	Dtu            *int32             `json:"dtu,omitempty"`
	DatabaseDtuMax *int32             `json:"databaseDtuMax,omitempty"`
	DatabaseDtuMin *int32             `json:"databaseDtuMin,omitempty"`
	StorageMB      *int32             `json:"storageMB,omitempty"`
}

// ElasticPoolUpdate is represents an elastic pool update.
type ElasticPoolUpdate struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*ElasticPoolProperties `json:"properties,omitempty"`
}

// OperationImpact is the impact of an operation, both in absolute and relative
// terms.
type OperationImpact struct {
	Name                *string  `json:"name,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	ChangeValueAbsolute *float64 `json:"changeValueAbsolute,omitempty"`
	ChangeValueRelative *float64 `json:"changeValueRelative,omitempty"`
}

// ProxyResource is aRM proxy resource.
type ProxyResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// RecommendedElasticPool is represents a recommented elastic pool.
type RecommendedElasticPool struct {
	ID                                *string `json:"id,omitempty"`
	Name                              *string `json:"name,omitempty"`
	Type                              *string `json:"type,omitempty"`
	*RecommendedElasticPoolProperties `json:"properties,omitempty"`
}

// RecommendedElasticPoolMetric is represents recommended elastic pool metric.
type RecommendedElasticPoolMetric struct {
	DateTime *date.Time `json:"dateTime,omitempty"`
	Dtu      *float64   `json:"dtu,omitempty"`
	SizeGB   *float64   `json:"sizeGB,omitempty"`
}

// RecommendedElasticPoolProperties is represents the properties of a
// recommented elastic pool.
type RecommendedElasticPoolProperties struct {
	DatabaseEdition        ElasticPoolEdition              `json:"databaseEdition,omitempty"`
	Dtu                    *float64                        `json:"dtu,omitempty"`
	DatabaseDtuMin         *float64                        `json:"databaseDtuMin,omitempty"`
	DatabaseDtuMax         *float64                        `json:"databaseDtuMax,omitempty"`
	StorageMB              *float64                        `json:"storageMB,omitempty"`
	ObservationPeriodStart *date.Time                      `json:"observationPeriodStart,omitempty"`
	ObservationPeriodEnd   *date.Time                      `json:"observationPeriodEnd,omitempty"`
	MaxObservedDtu         *float64                        `json:"maxObservedDtu,omitempty"`
	MaxObservedStorageMB   *float64                        `json:"maxObservedStorageMB,omitempty"`
	Databases              *[]Database                     `json:"databases,omitempty"`
	Metrics                *[]RecommendedElasticPoolMetric `json:"metrics,omitempty"`
}

// RecommendedIndex is represents a database recommended index.
type RecommendedIndex struct {
	ID                          *string `json:"id,omitempty"`
	Name                        *string `json:"name,omitempty"`
	Type                        *string `json:"type,omitempty"`
	*RecommendedIndexProperties `json:"properties,omitempty"`
}

// RecommendedIndexProperties is represents the properties of a database
// recommended index.
type RecommendedIndexProperties struct {
	Action          RecommendedIndexAction `json:"action,omitempty"`
	State           RecommendedIndexState  `json:"state,omitempty"`
	Created         *date.Time             `json:"created,omitempty"`
	LastModified    *date.Time             `json:"lastModified,omitempty"`
	IndexType       RecommendedIndexType   `json:"indexType,omitempty"`
	Schema          *string                `json:"schema,omitempty"`
	Table           *string                `json:"table,omitempty"`
	Columns         *[]string              `json:"columns,omitempty"`
	IncludedColumns *[]string              `json:"includedColumns,omitempty"`
	IndexScript     *string                `json:"indexScript,omitempty"`
	EstimatedImpact *[]OperationImpact     `json:"estimatedImpact,omitempty"`
	ReportedImpact  *[]OperationImpact     `json:"reportedImpact,omitempty"`
}

// Resource is aRM resource.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ServiceTierAdvisor is represents a Service Tier Advisor.
type ServiceTierAdvisor struct {
	ID                            *string `json:"id,omitempty"`
	Name                          *string `json:"name,omitempty"`
	Type                          *string `json:"type,omitempty"`
	*ServiceTierAdvisorProperties `json:"properties,omitempty"`
}

// ServiceTierAdvisorProperties is represents the properties of a Service Tier
// Advisor.
type ServiceTierAdvisorProperties struct {
	ObservationPeriodStart                                 *date.Time        `json:"observationPeriodStart,omitempty"`
	ObservationPeriodEnd                                   *date.Time        `json:"observationPeriodEnd,omitempty"`
	ActiveTimeRatio                                        *float64          `json:"activeTimeRatio,omitempty"`
	MinDtu                                                 *float64          `json:"minDtu,omitempty"`
	AvgDtu                                                 *float64          `json:"avgDtu,omitempty"`
	MaxDtu                                                 *float64          `json:"maxDtu,omitempty"`
	MaxSizeInGB                                            *float64          `json:"maxSizeInGB,omitempty"`
	ServiceLevelObjectiveUsageMetrics                      *[]SloUsageMetric `json:"serviceLevelObjectiveUsageMetrics,omitempty"`
	CurrentServiceLevelObjective                           *string           `json:"currentServiceLevelObjective,omitempty"`
	CurrentServiceLevelObjectiveID                         *uuid.UUID        `json:"currentServiceLevelObjectiveId,omitempty"`
	UsageBasedRecommendationServiceLevelObjective          *string           `json:"usageBasedRecommendationServiceLevelObjective,omitempty"`
	UsageBasedRecommendationServiceLevelObjectiveID        *uuid.UUID        `json:"usageBasedRecommendationServiceLevelObjectiveId,omitempty"`
	DatabaseSizeBasedRecommendationServiceLevelObjective   *string           `json:"databaseSizeBasedRecommendationServiceLevelObjective,omitempty"`
	DatabaseSizeBasedRecommendationServiceLevelObjectiveID *uuid.UUID        `json:"databaseSizeBasedRecommendationServiceLevelObjectiveId,omitempty"`
	DisasterPlanBasedRecommendationServiceLevelObjective   *string           `json:"disasterPlanBasedRecommendationServiceLevelObjective,omitempty"`
	DisasterPlanBasedRecommendationServiceLevelObjectiveID *uuid.UUID        `json:"disasterPlanBasedRecommendationServiceLevelObjectiveId,omitempty"`
	OverallRecommendationServiceLevelObjective             *string           `json:"overallRecommendationServiceLevelObjective,omitempty"`
	OverallRecommendationServiceLevelObjectiveID           *uuid.UUID        `json:"overallRecommendationServiceLevelObjectiveId,omitempty"`
	Confidence                                             *float64          `json:"confidence,omitempty"`
}

// SloUsageMetric is a Slo Usage Metric.
type SloUsageMetric struct {
	ServiceLevelObjective   ServiceObjectiveName `json:"serviceLevelObjective,omitempty"`
	ServiceLevelObjectiveID *uuid.UUID           `json:"serviceLevelObjectiveId,omitempty"`
	InRangeTimeRatio        *float64             `json:"inRangeTimeRatio,omitempty"`
}

// TrackedResource is aRM tracked top level resource.
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Location *string             `json:"location,omitempty"`
}

// TransparentDataEncryption is represents a database transparent data
// encryption configuration.
type TransparentDataEncryption struct {
	ID                                   *string `json:"id,omitempty"`
	Name                                 *string `json:"name,omitempty"`
	Type                                 *string `json:"type,omitempty"`
	Location                             *string `json:"location,omitempty"`
	*TransparentDataEncryptionProperties `json:"properties,omitempty"`
}

// TransparentDataEncryptionActivity is represents a database transparent data
// encryption Scan.
type TransparentDataEncryptionActivity struct {
	ID                                           *string `json:"id,omitempty"`
	Name                                         *string `json:"name,omitempty"`
	Type                                         *string `json:"type,omitempty"`
	Location                                     *string `json:"location,omitempty"`
	*TransparentDataEncryptionActivityProperties `json:"properties,omitempty"`
}

// TransparentDataEncryptionActivityProperties is represents the properties of
// a database transparent data encryption Scan.
type TransparentDataEncryptionActivityProperties struct {
	Status          TransparentDataEncryptionActivityStatus `json:"status,omitempty"`
	PercentComplete *float64                                `json:"percentComplete,omitempty"`
}

// TransparentDataEncryptionListResult is represents the response to a list
// transparent data encryption configurations request.
type TransparentDataEncryptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]TransparentDataEncryption `json:"value,omitempty"`
}

// TransparentDataEncryptionProperties is represents the properties of a
// database transparent data encryption.
type TransparentDataEncryptionProperties struct {
	Status TransparentDataEncryptionStatus `json:"status,omitempty"`
}
