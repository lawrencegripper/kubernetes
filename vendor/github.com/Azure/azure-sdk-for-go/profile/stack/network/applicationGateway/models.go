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

package applicationgateway

import original "github.com/lawrencegripper/azure-sdk-for-go/service/network/management/2015-06-15/applicationGateway"

type ApplicationGatewaysClient = original.ApplicationGatewaysClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type CookieBasedAffinity = original.CookieBasedAffinity

const (
	Disabled	CookieBasedAffinity	= original.Disabled
	Enabled		CookieBasedAffinity	= original.Enabled
)

type IPAllocationMethod = original.IPAllocationMethod

const (
	Dynamic	IPAllocationMethod	= original.Dynamic
	Static	IPAllocationMethod	= original.Static
)

type OperationalState = original.OperationalState

const (
	Running		OperationalState	= original.Running
	Starting	OperationalState	= original.Starting
	Stopped		OperationalState	= original.Stopped
	Stopping	OperationalState	= original.Stopping
)

type Protocol = original.Protocol

const (
	HTTP	Protocol	= original.HTTP
	HTTPS	Protocol	= original.HTTPS
)

type RequestRoutingRuleType = original.RequestRoutingRuleType

const (
	Basic			RequestRoutingRuleType	= original.Basic
	PathBasedRouting	RequestRoutingRuleType	= original.PathBasedRouting
)

type RouteNextHopType = original.RouteNextHopType

const (
	Internet		RouteNextHopType	= original.Internet
	None			RouteNextHopType	= original.None
	VirtualAppliance	RouteNextHopType	= original.VirtualAppliance
	VirtualNetworkGateway	RouteNextHopType	= original.VirtualNetworkGateway
	VnetLocal		RouteNextHopType	= original.VnetLocal
)

type SecurityRuleAccess = original.SecurityRuleAccess

const (
	Allow	SecurityRuleAccess	= original.Allow
	Deny	SecurityRuleAccess	= original.Deny
)

type SecurityRuleDirection = original.SecurityRuleDirection

const (
	Inbound		SecurityRuleDirection	= original.Inbound
	Outbound	SecurityRuleDirection	= original.Outbound
)

type SecurityRuleProtocol = original.SecurityRuleProtocol

const (
	Asterisk	SecurityRuleProtocol	= original.Asterisk
	TCP		SecurityRuleProtocol	= original.TCP
	UDP		SecurityRuleProtocol	= original.UDP
)

type SkuName = original.SkuName

const (
	StandardLarge	SkuName	= original.StandardLarge
	StandardMedium	SkuName	= original.StandardMedium
	StandardSmall	SkuName	= original.StandardSmall
)

type Tier = original.Tier

const (
	Standard Tier = original.Standard
)

type TransportProtocol = original.TransportProtocol

const (
	TransportProtocolTCP	TransportProtocol	= original.TransportProtocolTCP
	TransportProtocolUDP	TransportProtocol	= original.TransportProtocolUDP
)

type BackendAddress = original.BackendAddress
type BackendAddressPool = original.BackendAddressPool
type BackendAddressPoolPropertiesFormat = original.BackendAddressPoolPropertiesFormat
type BackendAddressPoolPropertiesFormatType = original.BackendAddressPoolPropertiesFormatType
type BackendAddressPoolType = original.BackendAddressPoolType
type BackendHTTPSettings = original.BackendHTTPSettings
type BackendHTTPSettingsPropertiesFormat = original.BackendHTTPSettingsPropertiesFormat
type FrontendIPConfiguration = original.FrontendIPConfiguration
type FrontendIPConfigurationPropertiesFormat = original.FrontendIPConfigurationPropertiesFormat
type FrontendPort = original.FrontendPort
type FrontendPortPropertiesFormat = original.FrontendPortPropertiesFormat
type HTTPListener = original.HTTPListener
type HTTPListenerPropertiesFormat = original.HTTPListenerPropertiesFormat
type InboundNatRule = original.InboundNatRule
type InboundNatRulePropertiesFormat = original.InboundNatRulePropertiesFormat
type IPConfiguration = original.IPConfiguration
type IPConfigurationPropertiesFormat = original.IPConfigurationPropertiesFormat
type IPConfigurationPropertiesFormatType = original.IPConfigurationPropertiesFormatType
type IPConfigurationType = original.IPConfigurationType
type ListResult = original.ListResult
type Model = original.Model
type NetworkInterface = original.NetworkInterface
type NetworkInterfaceDNSSettings = original.NetworkInterfaceDNSSettings
type NetworkInterfaceIPConfiguration = original.NetworkInterfaceIPConfiguration
type NetworkInterfaceIPConfigurationPropertiesFormat = original.NetworkInterfaceIPConfigurationPropertiesFormat
type NetworkInterfacePropertiesFormat = original.NetworkInterfacePropertiesFormat
type NetworkSecurityGroup = original.NetworkSecurityGroup
type NetworkSecurityGroupPropertiesFormat = original.NetworkSecurityGroupPropertiesFormat
type PathRule = original.PathRule
type PathRulePropertiesFormat = original.PathRulePropertiesFormat
type Probe = original.Probe
type ProbePropertiesFormat = original.ProbePropertiesFormat
type PropertiesFormat = original.PropertiesFormat
type PublicIPAddress = original.PublicIPAddress
type PublicIPAddressDNSSettings = original.PublicIPAddressDNSSettings
type PublicIPAddressPropertiesFormat = original.PublicIPAddressPropertiesFormat
type RequestRoutingRule = original.RequestRoutingRule
type RequestRoutingRulePropertiesFormat = original.RequestRoutingRulePropertiesFormat
type Resource = original.Resource
type Route = original.Route
type RoutePropertiesFormat = original.RoutePropertiesFormat
type RouteTable = original.RouteTable
type RouteTablePropertiesFormat = original.RouteTablePropertiesFormat
type SecurityRule = original.SecurityRule
type SecurityRulePropertiesFormat = original.SecurityRulePropertiesFormat
type Sku = original.Sku
type SslCertificate = original.SslCertificate
type SslCertificatePropertiesFormat = original.SslCertificatePropertiesFormat
type Subnet = original.Subnet
type SubnetPropertiesFormat = original.SubnetPropertiesFormat
type SubResource = original.SubResource
type URLPathMap = original.URLPathMap
type URLPathMapPropertiesFormat = original.URLPathMapPropertiesFormat

func NewApplicationGatewaysClient(subscriptionID string) ApplicationGatewaysClient {
	return original.NewApplicationGatewaysClient(subscriptionID)
}
func NewApplicationGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewaysClient {
	return original.NewApplicationGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
