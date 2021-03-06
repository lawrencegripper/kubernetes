package advisors

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
)

// AutoExecuteStatus enumerates the values for auto execute status.
type AutoExecuteStatus string

const (
	// Default specifies the default state for auto execute status.
	Default AutoExecuteStatus = "Default"
	// Disabled specifies the disabled state for auto execute status.
	Disabled AutoExecuteStatus = "Disabled"
	// Enabled specifies the enabled state for auto execute status.
	Enabled AutoExecuteStatus = "Enabled"
)

// AutoExecuteStatusInheritedFrom enumerates the values for auto execute status
// inherited from.
type AutoExecuteStatusInheritedFrom string

const (
	// AutoExecuteStatusInheritedFromDatabase specifies the auto execute status
	// inherited from database state for auto execute status inherited from.
	AutoExecuteStatusInheritedFromDatabase AutoExecuteStatusInheritedFrom = "Database"
	// AutoExecuteStatusInheritedFromDefault specifies the auto execute status
	// inherited from default state for auto execute status inherited from.
	AutoExecuteStatusInheritedFromDefault AutoExecuteStatusInheritedFrom = "Default"
	// AutoExecuteStatusInheritedFromElasticPool specifies the auto execute
	// status inherited from elastic pool state for auto execute status
	// inherited from.
	AutoExecuteStatusInheritedFromElasticPool AutoExecuteStatusInheritedFrom = "ElasticPool"
	// AutoExecuteStatusInheritedFromServer specifies the auto execute status
	// inherited from server state for auto execute status inherited from.
	AutoExecuteStatusInheritedFromServer AutoExecuteStatusInheritedFrom = "Server"
	// AutoExecuteStatusInheritedFromSubscription specifies the auto execute
	// status inherited from subscription state for auto execute status
	// inherited from.
	AutoExecuteStatusInheritedFromSubscription AutoExecuteStatusInheritedFrom = "Subscription"
)

// ImplementationMethod enumerates the values for implementation method.
type ImplementationMethod string

const (
	// AzurePowerShell specifies the azure power shell state for implementation
	// method.
	AzurePowerShell ImplementationMethod = "AzurePowerShell"
	// TSQL specifies the tsql state for implementation method.
	TSQL ImplementationMethod = "TSql"
)

// IsRetryable enumerates the values for is retryable.
type IsRetryable string

const (
	// No specifies the no state for is retryable.
	No IsRetryable = "No"
	// Yes specifies the yes state for is retryable.
	Yes IsRetryable = "Yes"
)

// RecommendedActionCurrentState enumerates the values for recommended action
// current state.
type RecommendedActionCurrentState string

const (
	// Active specifies the active state for recommended action current state.
	Active RecommendedActionCurrentState = "Active"
	// Error specifies the error state for recommended action current state.
	Error RecommendedActionCurrentState = "Error"
	// Executing specifies the executing state for recommended action current
	// state.
	Executing RecommendedActionCurrentState = "Executing"
	// Expired specifies the expired state for recommended action current
	// state.
	Expired RecommendedActionCurrentState = "Expired"
	// Ignored specifies the ignored state for recommended action current
	// state.
	Ignored RecommendedActionCurrentState = "Ignored"
	// Monitoring specifies the monitoring state for recommended action current
	// state.
	Monitoring RecommendedActionCurrentState = "Monitoring"
	// Pending specifies the pending state for recommended action current
	// state.
	Pending RecommendedActionCurrentState = "Pending"
	// PendingRevert specifies the pending revert state for recommended action
	// current state.
	PendingRevert RecommendedActionCurrentState = "PendingRevert"
	// Resolved specifies the resolved state for recommended action current
	// state.
	Resolved RecommendedActionCurrentState = "Resolved"
	// RevertCancelled specifies the revert cancelled state for recommended
	// action current state.
	RevertCancelled RecommendedActionCurrentState = "RevertCancelled"
	// Reverted specifies the reverted state for recommended action current
	// state.
	Reverted RecommendedActionCurrentState = "Reverted"
	// Reverting specifies the reverting state for recommended action current
	// state.
	Reverting RecommendedActionCurrentState = "Reverting"
	// Success specifies the success state for recommended action current
	// state.
	Success RecommendedActionCurrentState = "Success"
	// Verifying specifies the verifying state for recommended action current
	// state.
	Verifying RecommendedActionCurrentState = "Verifying"
)

// RecommendedActionInitiatedBy enumerates the values for recommended action
// initiated by.
type RecommendedActionInitiatedBy string

const (
	// System specifies the system state for recommended action initiated by.
	System RecommendedActionInitiatedBy = "System"
	// User specifies the user state for recommended action initiated by.
	User RecommendedActionInitiatedBy = "User"
)

// Status enumerates the values for status.
type Status string

const (
	// GA specifies the ga state for status.
	GA Status = "GA"
	// LimitedPublicPreview specifies the limited public preview state for
	// status.
	LimitedPublicPreview Status = "LimitedPublicPreview"
	// PrivatePreview specifies the private preview state for status.
	PrivatePreview Status = "PrivatePreview"
	// PublicPreview specifies the public preview state for status.
	PublicPreview Status = "PublicPreview"
)

// Advisor is database, Server or Elatic Pool Advisor.
type Advisor struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
	Kind              *string `json:"kind,omitempty"`
	Location          *string `json:"location,omitempty"`
	*Properties       `json:"properties,omitempty"`
}

// ListAdvisor is
type ListAdvisor struct {
	autorest.Response `json:"-"`
	Value             *[]Advisor `json:"value,omitempty"`
}

// ListRecommendedAction is
type ListRecommendedAction struct {
	autorest.Response `json:"-"`
	Value             *[]RecommendedAction `json:"value,omitempty"`
}

// Properties is properties for a Database, Server or Elastic Pool Advisor.
type Properties struct {
	AdvisorStatus                  Status                         `json:"advisorStatus,omitempty"`
	AutoExecuteStatus              AutoExecuteStatus              `json:"autoExecuteStatus,omitempty"`
	AutoExecuteStatusInheritedFrom AutoExecuteStatusInheritedFrom `json:"autoExecuteStatusInheritedFrom,omitempty"`
	RecommendationsStatus          *string                        `json:"recommendationsStatus,omitempty"`
	LastChecked                    *date.Time                     `json:"lastChecked,omitempty"`
	RecommendedActions             *[]RecommendedAction           `json:"recommendedActions,omitempty"`
}

// ProxyResource is aRM proxy resource.
type ProxyResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// RecommendedAction is database, Server or Elastic Pool Recommended Action.
type RecommendedAction struct {
	autorest.Response            `json:"-"`
	ID                           *string `json:"id,omitempty"`
	Name                         *string `json:"name,omitempty"`
	Type                         *string `json:"type,omitempty"`
	Kind                         *string `json:"kind,omitempty"`
	Location                     *string `json:"location,omitempty"`
	*RecommendedActionProperties `json:"properties,omitempty"`
}

// RecommendedActionErrorInfo is contains error information for an Azure SQL
// Database, Server or Elastic Pool Recommended Action.
type RecommendedActionErrorInfo struct {
	ErrorCode   *string     `json:"errorCode,omitempty"`
	IsRetryable IsRetryable `json:"isRetryable,omitempty"`
}

// RecommendedActionImpactRecord is contains information of estimated or
// observed impact on various metrics for an Azure SQL Database, Server or
// Elastic Pool Recommended Action.
type RecommendedActionImpactRecord struct {
	DimensionName       *string  `json:"dimensionName,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	AbsoluteValue       *float64 `json:"absoluteValue,omitempty"`
	ChangeValueAbsolute *float64 `json:"changeValueAbsolute,omitempty"`
	ChangeValueRelative *float64 `json:"changeValueRelative,omitempty"`
}

// RecommendedActionImplementationInfo is contains information for manual
// implementation for an Azure SQL Database, Server or Elastic Pool Recommended
// Action.
type RecommendedActionImplementationInfo struct {
	Method ImplementationMethod `json:"method,omitempty"`
	Script *string              `json:"script,omitempty"`
}

// RecommendedActionMetricInfo is contains time series of various impacted
// metrics for an Azure SQL Database, Server or Elastic Pool Recommended
// Action.
type RecommendedActionMetricInfo struct {
	MetricName *string    `json:"metricName,omitempty"`
	Unit       *string    `json:"unit,omitempty"`
	TimeGrain  *string    `json:"timeGrain,omitempty"`
	StartTime  *date.Time `json:"startTime,omitempty"`
	Value      *float64   `json:"value,omitempty"`
}

// RecommendedActionProperties is properties for a Database, Server or Elastic
// Pool Recommended Action.
type RecommendedActionProperties struct {
	RecommendationReason       *string                              `json:"recommendationReason,omitempty"`
	ValidSince                 *date.Time                           `json:"validSince,omitempty"`
	LastRefresh                *date.Time                           `json:"lastRefresh,omitempty"`
	State                      *RecommendedActionStateInfo          `json:"state,omitempty"`
	IsExecutableAction         *bool                                `json:"isExecutableAction,omitempty"`
	IsRevertableAction         *bool                                `json:"isRevertableAction,omitempty"`
	IsArchivedAction           *bool                                `json:"isArchivedAction,omitempty"`
	ExecuteActionStartTime     *date.Time                           `json:"executeActionStartTime,omitempty"`
	ExecuteActionDuration      *string                              `json:"executeActionDuration,omitempty"`
	RevertActionStartTime      *date.Time                           `json:"revertActionStartTime,omitempty"`
	RevertActionDuration       *string                              `json:"revertActionDuration,omitempty"`
	ExecuteActionInitiatedBy   RecommendedActionInitiatedBy         `json:"executeActionInitiatedBy,omitempty"`
	ExecuteActionInitiatedTime *date.Time                           `json:"executeActionInitiatedTime,omitempty"`
	RevertActionInitiatedBy    RecommendedActionInitiatedBy         `json:"revertActionInitiatedBy,omitempty"`
	RevertActionInitiatedTime  *date.Time                           `json:"revertActionInitiatedTime,omitempty"`
	Score                      *int32                               `json:"score,omitempty"`
	ImplementationDetails      *RecommendedActionImplementationInfo `json:"implementationDetails,omitempty"`
	ErrorDetails               *RecommendedActionErrorInfo          `json:"errorDetails,omitempty"`
	EstimatedImpact            *[]RecommendedActionImpactRecord     `json:"estimatedImpact,omitempty"`
	ObservedImpact             *[]RecommendedActionImpactRecord     `json:"observedImpact,omitempty"`
	TimeSeries                 *[]RecommendedActionMetricInfo       `json:"timeSeries,omitempty"`
	LinkedObjects              *[]string                            `json:"linkedObjects,omitempty"`
	Details                    *map[string]*map[string]interface{}  `json:"details,omitempty"`
}

// RecommendedActionStateInfo is contains information of current state for an
// Azure SQL Database, Server or Elastic Pool Recommended Action.
type RecommendedActionStateInfo struct {
	CurrentValue      RecommendedActionCurrentState `json:"currentValue,omitempty"`
	ActionInitiatedBy RecommendedActionInitiatedBy  `json:"actionInitiatedBy,omitempty"`
	LastModified      *date.Time                    `json:"lastModified,omitempty"`
}

// Resource is aRM resource.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}
