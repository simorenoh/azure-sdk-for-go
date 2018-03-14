package trafficmanager

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// EndpointMonitorStatus enumerates the values for endpoint monitor status.
type EndpointMonitorStatus string

const (
	// CheckingEndpoint ...
	CheckingEndpoint EndpointMonitorStatus = "CheckingEndpoint"
	// Degraded ...
	Degraded EndpointMonitorStatus = "Degraded"
	// Disabled ...
	Disabled EndpointMonitorStatus = "Disabled"
	// Inactive ...
	Inactive EndpointMonitorStatus = "Inactive"
	// Online ...
	Online EndpointMonitorStatus = "Online"
	// Stopped ...
	Stopped EndpointMonitorStatus = "Stopped"
)

// PossibleEndpointMonitorStatusValues returns an array of possible values for the EndpointMonitorStatus const type.
func PossibleEndpointMonitorStatusValues() [6]EndpointMonitorStatus {
	return [6]EndpointMonitorStatus{CheckingEndpoint, Degraded, Disabled, Inactive, Online, Stopped}
}

// EndpointStatus enumerates the values for endpoint status.
type EndpointStatus string

const (
	// EndpointStatusDisabled ...
	EndpointStatusDisabled EndpointStatus = "Disabled"
	// EndpointStatusEnabled ...
	EndpointStatusEnabled EndpointStatus = "Enabled"
)

// PossibleEndpointStatusValues returns an array of possible values for the EndpointStatus const type.
func PossibleEndpointStatusValues() [2]EndpointStatus {
	return [2]EndpointStatus{EndpointStatusDisabled, EndpointStatusEnabled}
}

// MonitorProtocol enumerates the values for monitor protocol.
type MonitorProtocol string

const (
	// HTTP ...
	HTTP MonitorProtocol = "HTTP"
	// HTTPS ...
	HTTPS MonitorProtocol = "HTTPS"
	// TCP ...
	TCP MonitorProtocol = "TCP"
)

// PossibleMonitorProtocolValues returns an array of possible values for the MonitorProtocol const type.
func PossibleMonitorProtocolValues() [3]MonitorProtocol {
	return [3]MonitorProtocol{HTTP, HTTPS, TCP}
}

// ProfileMonitorStatus enumerates the values for profile monitor status.
type ProfileMonitorStatus string

const (
	// ProfileMonitorStatusCheckingEndpoints ...
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = "CheckingEndpoints"
	// ProfileMonitorStatusDegraded ...
	ProfileMonitorStatusDegraded ProfileMonitorStatus = "Degraded"
	// ProfileMonitorStatusDisabled ...
	ProfileMonitorStatusDisabled ProfileMonitorStatus = "Disabled"
	// ProfileMonitorStatusInactive ...
	ProfileMonitorStatusInactive ProfileMonitorStatus = "Inactive"
	// ProfileMonitorStatusOnline ...
	ProfileMonitorStatusOnline ProfileMonitorStatus = "Online"
)

// PossibleProfileMonitorStatusValues returns an array of possible values for the ProfileMonitorStatus const type.
func PossibleProfileMonitorStatusValues() [5]ProfileMonitorStatus {
	return [5]ProfileMonitorStatus{ProfileMonitorStatusCheckingEndpoints, ProfileMonitorStatusDegraded, ProfileMonitorStatusDisabled, ProfileMonitorStatusInactive, ProfileMonitorStatusOnline}
}

// ProfileStatus enumerates the values for profile status.
type ProfileStatus string

const (
	// ProfileStatusDisabled ...
	ProfileStatusDisabled ProfileStatus = "Disabled"
	// ProfileStatusEnabled ...
	ProfileStatusEnabled ProfileStatus = "Enabled"
)

// PossibleProfileStatusValues returns an array of possible values for the ProfileStatus const type.
func PossibleProfileStatusValues() [2]ProfileStatus {
	return [2]ProfileStatus{ProfileStatusDisabled, ProfileStatusEnabled}
}

// TrafficRoutingMethod enumerates the values for traffic routing method.
type TrafficRoutingMethod string

const (
	// Geographic ...
	Geographic TrafficRoutingMethod = "Geographic"
	// Performance ...
	Performance TrafficRoutingMethod = "Performance"
	// Priority ...
	Priority TrafficRoutingMethod = "Priority"
	// Weighted ...
	Weighted TrafficRoutingMethod = "Weighted"
)

// PossibleTrafficRoutingMethodValues returns an array of possible values for the TrafficRoutingMethod const type.
func PossibleTrafficRoutingMethodValues() [4]TrafficRoutingMethod {
	return [4]TrafficRoutingMethod{Geographic, Performance, Priority, Weighted}
}

// CheckTrafficManagerRelativeDNSNameAvailabilityParameters parameters supplied to check Traffic Manager name
// operation.
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters struct {
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource.
	Type *string `json:"type,omitempty"`
}

// CloudError an error returned by the Azure Resource Manager
type CloudError struct {
	// Error - The content of the error.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody the content of an error returned by the Azure Resource Manager
type CloudErrorBody struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message
	Message *string `json:"message,omitempty"`
	// Target - Error target
	Target *string `json:"target,omitempty"`
	// Details - Error details
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// DeleteOperationResult the result of the request or operation.
type DeleteOperationResult struct {
	autorest.Response `json:"-"`
	// OperationResult - The result of the operation or request.
	OperationResult *bool `json:"boolean,omitempty"`
}

// DNSConfig class containing DNS settings in a Traffic Manager profile.
type DNSConfig struct {
	// RelativeName - The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain name used by Azure Traffic Manager to form the fully-qualified domain name (FQDN) of the profile.
	RelativeName *string `json:"relativeName,omitempty"`
	// Fqdn - The fully-qualified domain name (FQDN) of the Traffic Manager profile. This is formed from the concatenation of the RelativeName with the DNS domain used by Azure Traffic Manager.
	Fqdn *string `json:"fqdn,omitempty"`
	// TTL - The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS responses provided by this Traffic Manager profile.
	TTL *int64 `json:"ttl,omitempty"`
}

// Endpoint class representing a Traffic Manager endpoint.
type Endpoint struct {
	autorest.Response `json:"-"`
	// EndpointProperties - The properties of the Traffic Manager endpoint.
	*EndpointProperties `json:"properties,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Endpoint.
func (e Endpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if e.EndpointProperties != nil {
		objectMap["properties"] = e.EndpointProperties
	}
	if e.ID != nil {
		objectMap["id"] = e.ID
	}
	if e.Name != nil {
		objectMap["name"] = e.Name
	}
	if e.Type != nil {
		objectMap["type"] = e.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Endpoint struct.
func (e *Endpoint) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var endpointProperties EndpointProperties
				err = json.Unmarshal(*v, &endpointProperties)
				if err != nil {
					return err
				}
				e.EndpointProperties = &endpointProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				e.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				e.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				e.Type = &typeVar
			}
		}
	}

	return nil
}

// EndpointProperties class representing a Traffic Manager endpoint properties.
type EndpointProperties struct {
	// TargetResourceID - The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceID *string `json:"targetResourceId,omitempty"`
	// Target - The fully-qualified DNS name of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `json:"target,omitempty"`
	// EndpointStatus - The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method. Possible values include: 'EndpointStatusEnabled', 'EndpointStatusDisabled'
	EndpointStatus EndpointStatus `json:"endpointStatus,omitempty"`
	// Weight - The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *int64 `json:"weight,omitempty"`
	// Priority - The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority *int64 `json:"priority,omitempty"`
	// EndpointLocation - Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty"`
	// EndpointMonitorStatus - The monitoring status of the endpoint. Possible values include: 'CheckingEndpoint', 'Online', 'Degraded', 'Disabled', 'Inactive', 'Stopped'
	EndpointMonitorStatus EndpointMonitorStatus `json:"endpointMonitorStatus,omitempty"`
	// MinChildEndpoints - The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *int64 `json:"minChildEndpoints,omitempty"`
	// GeoMapping - The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping *[]string `json:"geoMapping,omitempty"`
}

// GeographicHierarchy class representing the Geographic hierarchy used with the Geographic traffic routing method.
type GeographicHierarchy struct {
	autorest.Response `json:"-"`
	// GeographicHierarchyProperties - The properties of the Geographic Hierarchy resource.
	*GeographicHierarchyProperties `json:"properties,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for GeographicHierarchy.
func (gh GeographicHierarchy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gh.GeographicHierarchyProperties != nil {
		objectMap["properties"] = gh.GeographicHierarchyProperties
	}
	if gh.ID != nil {
		objectMap["id"] = gh.ID
	}
	if gh.Name != nil {
		objectMap["name"] = gh.Name
	}
	if gh.Type != nil {
		objectMap["type"] = gh.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for GeographicHierarchy struct.
func (gh *GeographicHierarchy) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var geographicHierarchyProperties GeographicHierarchyProperties
				err = json.Unmarshal(*v, &geographicHierarchyProperties)
				if err != nil {
					return err
				}
				gh.GeographicHierarchyProperties = &geographicHierarchyProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				gh.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				gh.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				gh.Type = &typeVar
			}
		}
	}

	return nil
}

// GeographicHierarchyProperties class representing the properties of the Geographic hierarchy used with the
// Geographic traffic routing method.
type GeographicHierarchyProperties struct {
	// GeographicHierarchy - The region at the root of the hierarchy from all the regions in the hierarchy can be retrieved.
	GeographicHierarchy *Region `json:"geographicHierarchy,omitempty"`
}

// MonitorConfig class containing endpoint monitoring settings in a Traffic Manager profile.
type MonitorConfig struct {
	// ProfileMonitorStatus - The profile-level monitoring status of the Traffic Manager profile. Possible values include: 'ProfileMonitorStatusCheckingEndpoints', 'ProfileMonitorStatusOnline', 'ProfileMonitorStatusDegraded', 'ProfileMonitorStatusDisabled', 'ProfileMonitorStatusInactive'
	ProfileMonitorStatus ProfileMonitorStatus `json:"profileMonitorStatus,omitempty"`
	// Protocol - The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health. Possible values include: 'HTTP', 'HTTPS', 'TCP'
	Protocol MonitorProtocol `json:"protocol,omitempty"`
	// Port - The TCP port used to probe for endpoint health.
	Port *int64 `json:"port,omitempty"`
	// Path - The path relative to the endpoint domain name used to probe for endpoint health.
	Path *string `json:"path,omitempty"`
	// IntervalInSeconds - The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager will check the health of each endpoint in this profile.
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty"`
	// TimeoutInSeconds - The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows endpoints in this profile to response to the health check.
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty"`
	// ToleratedNumberOfFailures - The number of consecutive failed health check that Traffic Manager tolerates before declaring an endpoint in this profile Degraded after the next failed health check.
	ToleratedNumberOfFailures *int64 `json:"toleratedNumberOfFailures,omitempty"`
}

// NameAvailability class representing a Traffic Manager Name Availability response.
type NameAvailability struct {
	autorest.Response `json:"-"`
	// Name - The relative name.
	Name *string `json:"name,omitempty"`
	// Type - Traffic Manager profile resource type.
	Type *string `json:"type,omitempty"`
	// NameAvailable - Describes whether the relative name is available or not.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason why the name is not available, when applicable.
	Reason *string `json:"reason,omitempty"`
	// Message - Descriptive message that explains why the name is not available, when applicable.
	Message *string `json:"message,omitempty"`
}

// Profile class representing a Traffic Manager profile.
type Profile struct {
	autorest.Response `json:"-"`
	// ProfileProperties - The properties of the Traffic Manager profile.
	*ProfileProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Profile.
func (p Profile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ProfileProperties != nil {
		objectMap["properties"] = p.ProfileProperties
	}
	if p.Tags != nil {
		objectMap["tags"] = p.Tags
	}
	if p.Location != nil {
		objectMap["location"] = p.Location
	}
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Type != nil {
		objectMap["type"] = p.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Profile struct.
func (p *Profile) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var profileProperties ProfileProperties
				err = json.Unmarshal(*v, &profileProperties)
				if err != nil {
					return err
				}
				p.ProfileProperties = &profileProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				p.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				p.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				p.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				p.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				p.Type = &typeVar
			}
		}
	}

	return nil
}

// ProfileListResult the list Traffic Manager profiles operation response.
type ProfileListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets the list of Traffic manager profiles.
	Value *[]Profile `json:"value,omitempty"`
}

// ProfileProperties class representing the Traffic Manager profile properties.
type ProfileProperties struct {
	// ProfileStatus - The status of the Traffic Manager profile. Possible values include: 'ProfileStatusEnabled', 'ProfileStatusDisabled'
	ProfileStatus ProfileStatus `json:"profileStatus,omitempty"`
	// TrafficRoutingMethod - The traffic routing method of the Traffic Manager profile. Possible values include: 'Performance', 'Priority', 'Weighted', 'Geographic'
	TrafficRoutingMethod TrafficRoutingMethod `json:"trafficRoutingMethod,omitempty"`
	// DNSConfig - The DNS settings of the Traffic Manager profile.
	DNSConfig *DNSConfig `json:"dnsConfig,omitempty"`
	// MonitorConfig - The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig `json:"monitorConfig,omitempty"`
	// Endpoints - The list of endpoints in the Traffic Manager profile.
	Endpoints *[]Endpoint `json:"endpoints,omitempty"`
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// Region class representing a region in the Geographic hierarchy used with the Geographic traffic routing method.
type Region struct {
	// Code - The code of the region
	Code *string `json:"code,omitempty"`
	// Name - The name of the region
	Name *string `json:"name,omitempty"`
	// Regions - The list of Regions grouped under this Region in the Geographic Hierarchy.
	Regions *[]Region `json:"regions,omitempty"`
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Network/trafficmanagerProfiles.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.ID != nil {
		objectMap["id"] = tr.ID
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Type != nil {
		objectMap["type"] = tr.Type
	}
	return json.Marshal(objectMap)
}
