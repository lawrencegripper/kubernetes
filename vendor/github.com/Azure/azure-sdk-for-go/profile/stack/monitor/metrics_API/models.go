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

package metricsapi

import original "github.com/lawrencegripper/azure-sdk-for-go/service/monitor/2016-09-01/metrics_API"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type MetricsClient = original.MetricsClient
type Unit = original.Unit

const (
	Bytes		Unit	= original.Bytes
	BytesPerSecond	Unit	= original.BytesPerSecond
	Count		Unit	= original.Count
	CountPerSecond	Unit	= original.CountPerSecond
	MilliSeconds	Unit	= original.MilliSeconds
	Percent		Unit	= original.Percent
	Seconds		Unit	= original.Seconds
)

type ErrorResponse = original.ErrorResponse
type LocalizableString = original.LocalizableString
type Metric = original.Metric
type MetricCollection = original.MetricCollection
type MetricValue = original.MetricValue

func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func New() ManagementClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) ManagementClient {
	return original.NewWithBaseURI(baseURI)
}
func NewMetricsClient() MetricsClient {
	return original.NewMetricsClient()
}
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI)
}
