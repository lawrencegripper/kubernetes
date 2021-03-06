package logicappsmanagementclient

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
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ConnectionParameterType enumerates the values for connection parameter type.
type ConnectionParameterType string

const (
	// ConnectionParameterTypeArray specifies the connection parameter type
	// array state for connection parameter type.
	ConnectionParameterTypeArray ConnectionParameterType = "array"
	// ConnectionParameterTypeBool specifies the connection parameter type bool
	// state for connection parameter type.
	ConnectionParameterTypeBool ConnectionParameterType = "bool"
	// ConnectionParameterTypeConnection specifies the connection parameter
	// type connection state for connection parameter type.
	ConnectionParameterTypeConnection ConnectionParameterType = "connection"
	// ConnectionParameterTypeInt specifies the connection parameter type int
	// state for connection parameter type.
	ConnectionParameterTypeInt ConnectionParameterType = "int"
	// ConnectionParameterTypeOauthSetting specifies the connection parameter
	// type oauth setting state for connection parameter type.
	ConnectionParameterTypeOauthSetting ConnectionParameterType = "oauthSetting"
	// ConnectionParameterTypeObject specifies the connection parameter type
	// object state for connection parameter type.
	ConnectionParameterTypeObject ConnectionParameterType = "object"
	// ConnectionParameterTypeSecureobject specifies the connection parameter
	// type secureobject state for connection parameter type.
	ConnectionParameterTypeSecureobject ConnectionParameterType = "secureobject"
	// ConnectionParameterTypeSecurestring specifies the connection parameter
	// type securestring state for connection parameter type.
	ConnectionParameterTypeSecurestring ConnectionParameterType = "securestring"
	// ConnectionParameterTypeString specifies the connection parameter type
	// string state for connection parameter type.
	ConnectionParameterTypeString ConnectionParameterType = "string"
)

// LinkState enumerates the values for link state.
type LinkState string

const (
	// Authenticated specifies the authenticated state for link state.
	Authenticated LinkState = "Authenticated"
	// Error specifies the error state for link state.
	Error LinkState = "Error"
	// Unauthenticated specifies the unauthenticated state for link state.
	Unauthenticated LinkState = "Unauthenticated"
)

// PrincipalType enumerates the values for principal type.
type PrincipalType string

const (
	// PrincipalTypeActiveDirectory specifies the principal type active
	// directory state for principal type.
	PrincipalTypeActiveDirectory PrincipalType = "ActiveDirectory"
	// PrincipalTypeConnection specifies the principal type connection state
	// for principal type.
	PrincipalTypeConnection PrincipalType = "Connection"
	// PrincipalTypeMicrosoftAccount specifies the principal type microsoft
	// account state for principal type.
	PrincipalTypeMicrosoftAccount PrincipalType = "MicrosoftAccount"
)

// APIEntity is aPI Management
type APIEntity struct {
	autorest.Response    `json:"-"`
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Kind                 *string             `json:"kind,omitempty"`
	Location             *string             `json:"location,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	*APIEntityProperties `json:"properties,omitempty"`
}

// APIEntityProperties is
type APIEntityProperties struct {
	Name                 *string                          `json:"name,omitempty"`
	GeneralInformation   *GeneralAPIInformation           `json:"generalInformation,omitempty"`
	Path                 *string                          `json:"path,omitempty"`
	RuntimeUrls          *[]string                        `json:"runtimeUrls,omitempty"`
	Protocols            *[]string                        `json:"protocols,omitempty"`
	Policies             *APIPolicies                     `json:"policies,omitempty"`
	BackendService       *BackendServiceDefinition        `json:"backendService,omitempty"`
	APIDefinitionURL     *string                          `json:"apiDefinitionUrl,omitempty"`
	Metadata             *map[string]interface{}          `json:"metadata,omitempty"`
	Capabilities         *[]string                        `json:"capabilities,omitempty"`
	ConnectionParameters *map[string]*ConnectionParameter `json:"connectionParameters,omitempty"`
	CreatedTime          *date.Time                       `json:"createdTime,omitempty"`
	ChangedTime          *date.Time                       `json:"changedTime,omitempty"`
}

// APIOAuthSettings is oAuth settings for the conenction provider
type APIOAuthSettings struct {
	IdentityProvider *string                                `json:"identityProvider,omitempty"`
	ClientID         *string                                `json:"clientId,omitempty"`
	ClientSecret     *string                                `json:"clientSecret,omitempty"`
	Scopes           *[]string                              `json:"scopes,omitempty"`
	RedirectURL      *string                                `json:"redirectUrl,omitempty"`
	Properties       *map[string]interface{}                `json:"properties,omitempty"`
	CustomParameters *map[string]*APIOAuthSettingsParameter `json:"customParameters,omitempty"`
}

// APIOAuthSettingsParameter is oAuth Settings Parameter
type APIOAuthSettingsParameter struct {
	Value        *string                 `json:"value,omitempty"`
	Options      *map[string]interface{} `json:"options,omitempty"`
	UIDefinition *map[string]interface{} `json:"uiDefinition,omitempty"`
}

// APIPolicies is api policies
type APIPolicies struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Kind                   *string             `json:"kind,omitempty"`
	Location               *string             `json:"location,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*APIPoliciesProperties `json:"properties,omitempty"`
}

// APIPoliciesProperties is
type APIPoliciesProperties struct {
	Content *string `json:"content,omitempty"`
}

// ApisCollection is collection of Apis
type ApisCollection struct {
	autorest.Response `json:"-"`
	Value             *[]APIEntity `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// ApisCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApisCollection) ApisCollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ArmPlan is the plan object in an ARM, represents a marketplace plan
type ArmPlan struct {
	Name          *string `json:"name,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Version       *string `json:"version,omitempty"`
}

// BackendServiceDefinition is aPI definitions with backend urls
type BackendServiceDefinition struct {
	ID                                  *string             `json:"id,omitempty"`
	Name                                *string             `json:"name,omitempty"`
	Kind                                *string             `json:"kind,omitempty"`
	Location                            *string             `json:"location,omitempty"`
	Type                                *string             `json:"type,omitempty"`
	Tags                                *map[string]*string `json:"tags,omitempty"`
	*BackendServiceDefinitionProperties `json:"properties,omitempty"`
}

// BackendServiceDefinitionProperties is
type BackendServiceDefinitionProperties struct {
	ServiceURL                    *string                                  `json:"serviceUrl,omitempty"`
	HostingEnvironmentServiceUrls *[]HostingEnvironmentServiceDescriptions `json:"hostingEnvironmentServiceUrls,omitempty"`
}

// ConfirmConsentCodeInput is confirm Consent Code Input payload
type ConfirmConsentCodeInput struct {
	ID                                 *string             `json:"id,omitempty"`
	Name                               *string             `json:"name,omitempty"`
	Kind                               *string             `json:"kind,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	Type                               *string             `json:"type,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*ConfirmConsentCodeInputProperties `json:"properties,omitempty"`
}

// ConfirmConsentCodeInputProperties is
type ConfirmConsentCodeInputProperties struct {
	PrincipalType PrincipalType `json:"principalType,omitempty"`
	TenantID      *string       `json:"tenantId,omitempty"`
	ObjectID      *string       `json:"objectId,omitempty"`
	Code          *string       `json:"code,omitempty"`
}

// Connection is api Connection
type Connection struct {
	autorest.Response     `json:"-"`
	ID                    *string             `json:"id,omitempty"`
	Name                  *string             `json:"name,omitempty"`
	Kind                  *string             `json:"kind,omitempty"`
	Location              *string             `json:"location,omitempty"`
	Type                  *string             `json:"type,omitempty"`
	Tags                  *map[string]*string `json:"tags,omitempty"`
	*ConnectionProperties `json:"properties,omitempty"`
}

// ConnectionProperties is
type ConnectionProperties struct {
	Name                     *string                                        `json:"name,omitempty"`
	DisplayName              *string                                        `json:"displayName,omitempty"`
	Statuses                 *[]ConnectionStatus                            `json:"statuses,omitempty"`
	CustomParameterValues    *map[string]*ParameterCustomLoginSettingValues `json:"customParameterValues,omitempty"`
	TenantID                 *string                                        `json:"tenantId,omitempty"`
	ParameterValues          *map[string]*map[string]interface{}            `json:"parameterValues,omitempty"`
	NonSecretParameterValues *map[string]*map[string]interface{}            `json:"nonSecretParameterValues,omitempty"`
	Metadata                 *map[string]interface{}                        `json:"metadata,omitempty"`
	FirstExpirationTime      *date.Time                                     `json:"firstExpirationTime,omitempty"`
	Keywords                 *[]string                                      `json:"keywords,omitempty"`
	CreatedTime              *date.Time                                     `json:"createdTime,omitempty"`
	ChangedTime              *date.Time                                     `json:"changedTime,omitempty"`
	API                      *ExpandedParentAPIEntity                       `json:"api,omitempty"`
}

// ConnectionCollection is collection of conenctions
type ConnectionCollection struct {
	autorest.Response `json:"-"`
	Value             *[]Connection `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// ConnectionCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ConnectionCollection) ConnectionCollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ConnectionError is connection error
type ConnectionError struct {
	ID                         *string             `json:"id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	Kind                       *string             `json:"kind,omitempty"`
	Location                   *string             `json:"location,omitempty"`
	Type                       *string             `json:"type,omitempty"`
	Tags                       *map[string]*string `json:"tags,omitempty"`
	*ConnectionErrorProperties `json:"properties,omitempty"`
}

// ConnectionErrorProperties is
type ConnectionErrorProperties struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ConnectionParameter is connection provider parameters
type ConnectionParameter struct {
	Type          ConnectionParameterType `json:"type,omitempty"`
	DefaultValue  *map[string]interface{} `json:"defaultValue,omitempty"`
	OAuthSettings *APIOAuthSettings       `json:"oAuthSettings,omitempty"`
	UIDefinition  *map[string]interface{} `json:"uiDefinition,omitempty"`
}

// ConnectionSecrets is
type ConnectionSecrets struct {
	autorest.Response `json:"-"`
	ParameterValues   *map[string]*map[string]interface{} `json:"parameterValues,omitempty"`
	ConnectionKey     *string                             `json:"connectionKey,omitempty"`
}

// ConnectionStatus is connection status
type ConnectionStatus struct {
	ID                          *string             `json:"id,omitempty"`
	Name                        *string             `json:"name,omitempty"`
	Kind                        *string             `json:"kind,omitempty"`
	Location                    *string             `json:"location,omitempty"`
	Type                        *string             `json:"type,omitempty"`
	Tags                        *map[string]*string `json:"tags,omitempty"`
	*ConnectionStatusProperties `json:"properties,omitempty"`
}

// ConnectionStatusProperties is
type ConnectionStatusProperties struct {
	Status *string          `json:"status,omitempty"`
	Target *string          `json:"target,omitempty"`
	Error  *ConnectionError `json:"error,omitempty"`
}

// ConsentLink is
type ConsentLink struct {
	Link               *string   `json:"link,omitempty"`
	FirstPartyLoginURI *string   `json:"firstPartyLoginUri,omitempty"`
	DisplayName        *string   `json:"displayName,omitempty"`
	Status             LinkState `json:"status,omitempty"`
}

// ConsentLinkInput is connection Constent Link payload
type ConsentLinkInput struct {
	ID                          *string             `json:"id,omitempty"`
	Name                        *string             `json:"name,omitempty"`
	Kind                        *string             `json:"kind,omitempty"`
	Location                    *string             `json:"location,omitempty"`
	Type                        *string             `json:"type,omitempty"`
	Tags                        *map[string]*string `json:"tags,omitempty"`
	*ConsentLinkInputProperties `json:"properties,omitempty"`
}

// ConsentLinkInputProperties is
type ConsentLinkInputProperties struct {
	Parameters *[]ConsentLinkInputParameter `json:"parameters,omitempty"`
}

// ConsentLinkInputParameter is
type ConsentLinkInputParameter struct {
	PrincipalType PrincipalType `json:"principalType,omitempty"`
	TenantID      *string       `json:"tenantId,omitempty"`
	ObjectID      *string       `json:"objectId,omitempty"`
	ParameterName *string       `json:"parameterName,omitempty"`
	RedirectURL   *string       `json:"redirectUrl,omitempty"`
}

// ConsentLinkPayload is collection of consent links
type ConsentLinkPayload struct {
	autorest.Response `json:"-"`
	Value             *[]ConsentLink `json:"value,omitempty"`
}

// CustomLoginSettingValue is custom logging setting value
type CustomLoginSettingValue struct {
	ID                                 *string             `json:"id,omitempty"`
	Name                               *string             `json:"name,omitempty"`
	Kind                               *string             `json:"kind,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	Type                               *string             `json:"type,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*CustomLoginSettingValueProperties `json:"properties,omitempty"`
}

// CustomLoginSettingValueProperties is
type CustomLoginSettingValueProperties struct {
	Option *string `json:"option,omitempty"`
}

// ExpandedParentAPIEntity is expanded parent object for expansion
type ExpandedParentAPIEntity struct {
	ID                                 *string             `json:"id,omitempty"`
	Name                               *string             `json:"name,omitempty"`
	Kind                               *string             `json:"kind,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	Type                               *string             `json:"type,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*ExpandedParentAPIEntityProperties `json:"properties,omitempty"`
}

// ExpandedParentAPIEntityProperties is
type ExpandedParentAPIEntityProperties struct {
	ID     *string                           `json:"id,omitempty"`
	Entity *ResponseMessageEnvelopeAPIEntity `json:"entity,omitempty"`
}

// GeneralAPIInformation is general API information
type GeneralAPIInformation struct {
	ID                               *string             `json:"id,omitempty"`
	Name                             *string             `json:"name,omitempty"`
	Kind                             *string             `json:"kind,omitempty"`
	Location                         *string             `json:"location,omitempty"`
	Type                             *string             `json:"type,omitempty"`
	Tags                             *map[string]*string `json:"tags,omitempty"`
	*GeneralAPIInformationProperties `json:"properties,omitempty"`
}

// GeneralAPIInformationProperties is
type GeneralAPIInformationProperties struct {
	IconURL               *string                 `json:"iconUrl,omitempty"`
	DisplayName           *string                 `json:"displayName,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	TermsOfUseURL         *string                 `json:"termsOfUseUrl,omitempty"`
	ConnectionDisplayName *string                 `json:"connectionDisplayName,omitempty"`
	ConnectionPortalURL   *map[string]interface{} `json:"connectionPortalUrl,omitempty"`
}

// HostingEnvironmentServiceDescriptions is back end service per ASE
type HostingEnvironmentServiceDescriptions struct {
	HostingEnvironmentID *string `json:"hostingEnvironmentId,omitempty"`
	HostID               *string `json:"hostId,omitempty"`
	ServiceURL           *string `json:"serviceUrl,omitempty"`
	UseInternalRouting   *bool   `json:"useInternalRouting,omitempty"`
}

// ListConnectionKeysInput is list Connection Keys Input payload
type ListConnectionKeysInput struct {
	ID                                 *string             `json:"id,omitempty"`
	Name                               *string             `json:"name,omitempty"`
	Kind                               *string             `json:"kind,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	Type                               *string             `json:"type,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*ListConnectionKeysInputProperties `json:"properties,omitempty"`
}

// ListConnectionKeysInputProperties is
type ListConnectionKeysInputProperties struct {
	ValidityTimeSpan *string `json:"validityTimeSpan,omitempty"`
}

// ParameterCustomLoginSettingValues is custom logging setting values
type ParameterCustomLoginSettingValues struct {
	ID                                           *string             `json:"id,omitempty"`
	Name                                         *string             `json:"name,omitempty"`
	Kind                                         *string             `json:"kind,omitempty"`
	Location                                     *string             `json:"location,omitempty"`
	Type                                         *string             `json:"type,omitempty"`
	Tags                                         *map[string]*string `json:"tags,omitempty"`
	*ParameterCustomLoginSettingValuesProperties `json:"properties,omitempty"`
}

// ParameterCustomLoginSettingValuesProperties is
type ParameterCustomLoginSettingValuesProperties struct {
	CustomParameters *map[string]*CustomLoginSettingValue `json:"customParameters,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Kind     *string             `json:"kind,omitempty"`
	Location *string             `json:"location,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResponseMessageEnvelopeAPIEntity is message envelope that contains the
// common Azure resource manager properties and the resource provider specific
// content
type ResponseMessageEnvelopeAPIEntity struct {
	ID         *string             `json:"id,omitempty"`
	Name       *string             `json:"name,omitempty"`
	Type       *string             `json:"type,omitempty"`
	Location   *string             `json:"location,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Plan       *ArmPlan            `json:"plan,omitempty"`
	Properties *APIEntity          `json:"properties,omitempty"`
	Sku        *SkuDescription     `json:"sku,omitempty"`
}

// SkuDescription is describes a sku for a scalable resource
type SkuDescription struct {
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
	Size     *string `json:"size,omitempty"`
	Family   *string `json:"family,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
}
