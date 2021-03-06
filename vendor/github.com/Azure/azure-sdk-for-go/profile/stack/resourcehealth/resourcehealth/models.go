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

package resourcehealth

import original "github.com/lawrencegripper/azure-sdk-for-go/service/resourcehealth/management/2015-01-01/resourcehealth"

type AvailabilityStatusesClient = original.AvailabilityStatusesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AvailabilityStateValues = original.AvailabilityStateValues

const (
	Available	AvailabilityStateValues	= original.Available
	Unavailable	AvailabilityStateValues	= original.Unavailable
	Unknown		AvailabilityStateValues	= original.Unknown
)

type ReasonChronicityTypes = original.ReasonChronicityTypes

const (
	Persistent	ReasonChronicityTypes	= original.Persistent
	Transient	ReasonChronicityTypes	= original.Transient
)

type AvailabilityStatus = original.AvailabilityStatus
type AvailabilityStatusProperties = original.AvailabilityStatusProperties
type AvailabilityStatusPropertiesRecentlyResolvedState = original.AvailabilityStatusPropertiesRecentlyResolvedState
type AvailabilityStatusListResult = original.AvailabilityStatusListResult
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type RecommendedAction = original.RecommendedAction
type ServiceImpactingEvent = original.ServiceImpactingEvent
type ServiceImpactingEventIncidentProperties = original.ServiceImpactingEventIncidentProperties
type ServiceImpactingEventStatus = original.ServiceImpactingEventStatus
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string, resourceType string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, resourceType)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, resourceType string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, resourceType)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewAvailabilityStatusesClient(subscriptionID string, resourceType string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClient(subscriptionID, resourceType)
}
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string, resourceType string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClientWithBaseURI(baseURI, subscriptionID, resourceType)
}
func New(subscriptionID string, resourceType string) ManagementClient {
	return original.New(subscriptionID, resourceType)
}
func NewWithBaseURI(baseURI string, subscriptionID string, resourceType string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, resourceType)
}
