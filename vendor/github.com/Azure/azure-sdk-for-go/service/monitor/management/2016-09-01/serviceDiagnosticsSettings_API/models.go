package servicediagnosticssettingsapi

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
)

// ErrorResponse is describes the format of Error response.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// LogSettings is part of MultiTenantDiagnosticSettings. Specifies the settings
// for a particular log.
type LogSettings struct {
	Category        *string          `json:"category,omitempty"`
	Enabled         *bool            `json:"enabled,omitempty"`
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// MetricSettings is part of MultiTenantDiagnosticSettings. Specifies the
// settings for a particular metric.
type MetricSettings struct {
	TimeGrain       *string          `json:"timeGrain,omitempty"`
	Enabled         *bool            `json:"enabled,omitempty"`
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// Resource is an azure resource object
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// RetentionPolicy is specifies the retention policy for the log.
type RetentionPolicy struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Days    *int32 `json:"days,omitempty"`
}

// ServiceDiagnosticSettings is the diagnostic settings for service.
type ServiceDiagnosticSettings struct {
	StorageAccountID            *string           `json:"storageAccountId,omitempty"`
	ServiceBusRuleID            *string           `json:"serviceBusRuleId,omitempty"`
	EventHubAuthorizationRuleID *string           `json:"eventHubAuthorizationRuleId,omitempty"`
	Metrics                     *[]MetricSettings `json:"metrics,omitempty"`
	Logs                        *[]LogSettings    `json:"logs,omitempty"`
	WorkspaceID                 *string           `json:"workspaceId,omitempty"`
}

// ServiceDiagnosticSettingsResource is description of a service diagnostic
// setting
type ServiceDiagnosticSettingsResource struct {
	autorest.Response          `json:"-"`
	ID                         *string             `json:"id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	Type                       *string             `json:"type,omitempty"`
	Location                   *string             `json:"location,omitempty"`
	Tags                       *map[string]*string `json:"tags,omitempty"`
	*ServiceDiagnosticSettings `json:"properties,omitempty"`
}

// ServiceDiagnosticSettingsResourcePatch is service diagnostic setting
// resource for patch operations
type ServiceDiagnosticSettingsResourcePatch struct {
	Tags                       *map[string]*string `json:"tags,omitempty"`
	*ServiceDiagnosticSettings `json:"properties,omitempty"`
}
