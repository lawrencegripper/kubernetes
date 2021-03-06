// Package catalog implements the Azure ARM Catalog service API version
// 2016-11-01.
//
// Creates an Azure Data Lake Analytics catalog client.
package catalog

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

const (
	// DefaultAdlaCatalogDNSSuffix is the default value for adla catalog dns suffix
	DefaultAdlaCatalogDNSSuffix = "azuredatalakeanalytics.net"
)

// ManagementClient is the base client for Catalog.
type ManagementClient struct {
	autorest.Client
	AdlaCatalogDNSSuffix string
}

// New creates an instance of the ManagementClient client.
func New() ManagementClient {
	return NewWithoutDefaults(DefaultAdlaCatalogDNSSuffix)
}

// NewWithoutDefaults creates an instance of the ManagementClient client.
func NewWithoutDefaults(adlaCatalogDNSSuffix string) ManagementClient {
	return ManagementClient{
		Client:               autorest.NewClientWithUserAgent(UserAgent()),
		AdlaCatalogDNSSuffix: adlaCatalogDNSSuffix,
	}
}
