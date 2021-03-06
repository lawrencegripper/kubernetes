package topleveldomains

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
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AgreementOption is options for retrieving the list of top level domain legal
// agreements.
type AgreementOption struct {
	IncludePrivacy *bool `json:"includePrivacy,omitempty"`
	ForTransfer    *bool `json:"forTransfer,omitempty"`
}

// Collection is collection of Top-level domains.
type Collection struct {
	autorest.Response `json:"-"`
	Value             *[]TopLevelDomain `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// CollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client Collection) CollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is azure resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Kind     *string             `json:"kind,omitempty"`
	Location *string             `json:"location,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// TldLegalAgreement is legal agreement for a top level domain.
type TldLegalAgreement struct {
	AgreementKey *string `json:"agreementKey,omitempty"`
	Title        *string `json:"title,omitempty"`
	Content      *string `json:"content,omitempty"`
	URL          *string `json:"url,omitempty"`
}

// TldLegalAgreementCollection is collection of top-level domain legal
// agreements.
type TldLegalAgreementCollection struct {
	autorest.Response `json:"-"`
	Value             *[]TldLegalAgreement `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// TldLegalAgreementCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client TldLegalAgreementCollection) TldLegalAgreementCollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// TopLevelDomain is a top level domain object.
type TopLevelDomain struct {
	autorest.Response         `json:"-"`
	ID                        *string             `json:"id,omitempty"`
	Name                      *string             `json:"name,omitempty"`
	Kind                      *string             `json:"kind,omitempty"`
	Location                  *string             `json:"location,omitempty"`
	Type                      *string             `json:"type,omitempty"`
	Tags                      *map[string]*string `json:"tags,omitempty"`
	*TopLevelDomainProperties `json:"properties,omitempty"`
}

// TopLevelDomainProperties is topLevelDomain resource specific properties
type TopLevelDomainProperties struct {
	DomainName *string `json:"name,omitempty"`
	Privacy    *bool   `json:"privacy,omitempty"`
}
