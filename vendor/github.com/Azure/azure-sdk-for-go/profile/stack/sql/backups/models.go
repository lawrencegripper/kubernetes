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

package backups

import original "github.com/lawrencegripper/azure-sdk-for-go/service/sql/management/2014-04-01/backups"

type RecoverableDatabasesClient = original.RecoverableDatabasesClient
type RestorableDroppedDatabasesClient = original.RestorableDroppedDatabasesClient
type RestorePointsClient = original.RestorePointsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type RestorePointType = original.RestorePointType

const (
	CONTINUOUS	RestorePointType	= original.CONTINUOUS
	DISCRETE	RestorePointType	= original.DISCRETE
)

type ProxyResource = original.ProxyResource
type RecoverableDatabase = original.RecoverableDatabase
type RecoverableDatabaseListResult = original.RecoverableDatabaseListResult
type RecoverableDatabaseProperties = original.RecoverableDatabaseProperties
type Resource = original.Resource
type RestorableDroppedDatabase = original.RestorableDroppedDatabase
type RestorableDroppedDatabaseListResult = original.RestorableDroppedDatabaseListResult
type RestorableDroppedDatabaseProperties = original.RestorableDroppedDatabaseProperties
type RestorePoint = original.RestorePoint
type RestorePointListResult = original.RestorePointListResult
type RestorePointProperties = original.RestorePointProperties
type TrackedResource = original.TrackedResource

func NewRecoverableDatabasesClient(subscriptionID string) RecoverableDatabasesClient {
	return original.NewRecoverableDatabasesClient(subscriptionID)
}
func NewRecoverableDatabasesClientWithBaseURI(baseURI string, subscriptionID string) RecoverableDatabasesClient {
	return original.NewRecoverableDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRestorableDroppedDatabasesClient(subscriptionID string) RestorableDroppedDatabasesClient {
	return original.NewRestorableDroppedDatabasesClient(subscriptionID)
}
func NewRestorableDroppedDatabasesClientWithBaseURI(baseURI string, subscriptionID string) RestorableDroppedDatabasesClient {
	return original.NewRestorableDroppedDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRestorePointsClient(subscriptionID string) RestorePointsClient {
	return original.NewRestorePointsClient(subscriptionID)
}
func NewRestorePointsClientWithBaseURI(baseURI string, subscriptionID string) RestorePointsClient {
	return original.NewRestorePointsClientWithBaseURI(baseURI, subscriptionID)
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
