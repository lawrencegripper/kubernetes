package apimanagement

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

// PolicyScopeContract enumerates the values for policy scope contract.
type PolicyScopeContract string

const (
	// All specifies the all state for policy scope contract.
	All PolicyScopeContract = "All"
	// API specifies the api state for policy scope contract.
	API PolicyScopeContract = "Api"
	// Operation specifies the operation state for policy scope contract.
	Operation PolicyScopeContract = "Operation"
	// Product specifies the product state for policy scope contract.
	Product PolicyScopeContract = "Product"
	// Tenant specifies the tenant state for policy scope contract.
	Tenant PolicyScopeContract = "Tenant"
)

// ErrorBodyContract is error Body contract.
type ErrorBodyContract struct {
	Code    *string               `json:"code,omitempty"`
	Message *string               `json:"message,omitempty"`
	Details *[]ErrorFieldContract `json:"details,omitempty"`
}

// ErrorFieldContract is error Field contract.
type ErrorFieldContract struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// PolicySnippetContract is policy snippet.
type PolicySnippetContract struct {
	Name    *string             `json:"name,omitempty"`
	Content *string             `json:"content,omitempty"`
	ToolTip *string             `json:"toolTip,omitempty"`
	Scope   PolicyScopeContract `json:"scope,omitempty"`
}

// PolicySnippetsCollection is the response of the list policy snippets
// operation.
type PolicySnippetsCollection struct {
	autorest.Response `json:"-"`
	Value             *[]PolicySnippetContract `json:"value,omitempty"`
}

// RegionContract is region profile.
type RegionContract struct {
	Name           *string `json:"name,omitempty"`
	IsMasterRegion *bool   `json:"isMasterRegion,omitempty"`
}

// RegionListResult is lists Regions operation response details.
type RegionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RegionContract `json:"value,omitempty"`
}
