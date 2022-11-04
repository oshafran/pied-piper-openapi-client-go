/*
Cisco SD-WAN vManage API

The vManage API exposes the functionality of operations maintaining devices and the overlay network

API version: 2.0.0
Contact: vmanage@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NetworkProtocolAllOf struct for NetworkProtocolAllOf
type NetworkProtocolAllOf struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	DHCPPool *DHCPPool `json:"DHCPPool,omitempty"`
	DNSSettings *string `json:"DNSSettings,omitempty"`
	NTPSettings []string `json:"NTPSettings,omitempty"`
	NATRules []NATRule `json:"NATRules,omitempty"`
	NTPInherit *bool `json:"NTPInherit,omitempty"`
}

// NewNetworkProtocolAllOf instantiates a new NetworkProtocolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtocolAllOf() *NetworkProtocolAllOf {
	this := NetworkProtocolAllOf{}
	return &this
}

// NewNetworkProtocolAllOfWithDefaults instantiates a new NetworkProtocolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtocolAllOfWithDefaults() *NetworkProtocolAllOf {
	this := NetworkProtocolAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkProtocolAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkProtocolAllOf) SetType(v string) {
	o.Type = &v
}

// GetDHCPPool returns the DHCPPool field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetDHCPPool() DHCPPool {
	if o == nil || isNil(o.DHCPPool) {
		var ret DHCPPool
		return ret
	}
	return *o.DHCPPool
}

// GetDHCPPoolOk returns a tuple with the DHCPPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetDHCPPoolOk() (*DHCPPool, bool) {
	if o == nil || isNil(o.DHCPPool) {
    return nil, false
	}
	return o.DHCPPool, true
}

// HasDHCPPool returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasDHCPPool() bool {
	if o != nil && !isNil(o.DHCPPool) {
		return true
	}

	return false
}

// SetDHCPPool gets a reference to the given DHCPPool and assigns it to the DHCPPool field.
func (o *NetworkProtocolAllOf) SetDHCPPool(v DHCPPool) {
	o.DHCPPool = &v
}

// GetDNSSettings returns the DNSSettings field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetDNSSettings() string {
	if o == nil || isNil(o.DNSSettings) {
		var ret string
		return ret
	}
	return *o.DNSSettings
}

// GetDNSSettingsOk returns a tuple with the DNSSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetDNSSettingsOk() (*string, bool) {
	if o == nil || isNil(o.DNSSettings) {
    return nil, false
	}
	return o.DNSSettings, true
}

// HasDNSSettings returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasDNSSettings() bool {
	if o != nil && !isNil(o.DNSSettings) {
		return true
	}

	return false
}

// SetDNSSettings gets a reference to the given string and assigns it to the DNSSettings field.
func (o *NetworkProtocolAllOf) SetDNSSettings(v string) {
	o.DNSSettings = &v
}

// GetNTPSettings returns the NTPSettings field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetNTPSettings() []string {
	if o == nil || isNil(o.NTPSettings) {
		var ret []string
		return ret
	}
	return o.NTPSettings
}

// GetNTPSettingsOk returns a tuple with the NTPSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetNTPSettingsOk() ([]string, bool) {
	if o == nil || isNil(o.NTPSettings) {
    return nil, false
	}
	return o.NTPSettings, true
}

// HasNTPSettings returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasNTPSettings() bool {
	if o != nil && !isNil(o.NTPSettings) {
		return true
	}

	return false
}

// SetNTPSettings gets a reference to the given []string and assigns it to the NTPSettings field.
func (o *NetworkProtocolAllOf) SetNTPSettings(v []string) {
	o.NTPSettings = v
}

// GetNATRules returns the NATRules field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetNATRules() []NATRule {
	if o == nil || isNil(o.NATRules) {
		var ret []NATRule
		return ret
	}
	return o.NATRules
}

// GetNATRulesOk returns a tuple with the NATRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetNATRulesOk() ([]NATRule, bool) {
	if o == nil || isNil(o.NATRules) {
    return nil, false
	}
	return o.NATRules, true
}

// HasNATRules returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasNATRules() bool {
	if o != nil && !isNil(o.NATRules) {
		return true
	}

	return false
}

// SetNATRules gets a reference to the given []NATRule and assigns it to the NATRules field.
func (o *NetworkProtocolAllOf) SetNATRules(v []NATRule) {
	o.NATRules = v
}

// GetNTPInherit returns the NTPInherit field value if set, zero value otherwise.
func (o *NetworkProtocolAllOf) GetNTPInherit() bool {
	if o == nil || isNil(o.NTPInherit) {
		var ret bool
		return ret
	}
	return *o.NTPInherit
}

// GetNTPInheritOk returns a tuple with the NTPInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocolAllOf) GetNTPInheritOk() (*bool, bool) {
	if o == nil || isNil(o.NTPInherit) {
    return nil, false
	}
	return o.NTPInherit, true
}

// HasNTPInherit returns a boolean if a field has been set.
func (o *NetworkProtocolAllOf) HasNTPInherit() bool {
	if o != nil && !isNil(o.NTPInherit) {
		return true
	}

	return false
}

// SetNTPInherit gets a reference to the given bool and assigns it to the NTPInherit field.
func (o *NetworkProtocolAllOf) SetNTPInherit(v bool) {
	o.NTPInherit = &v
}

func (o NetworkProtocolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.DHCPPool) {
		toSerialize["DHCPPool"] = o.DHCPPool
	}
	if !isNil(o.DNSSettings) {
		toSerialize["DNSSettings"] = o.DNSSettings
	}
	if !isNil(o.NTPSettings) {
		toSerialize["NTPSettings"] = o.NTPSettings
	}
	if !isNil(o.NATRules) {
		toSerialize["NATRules"] = o.NATRules
	}
	if !isNil(o.NTPInherit) {
		toSerialize["NTPInherit"] = o.NTPInherit
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkProtocolAllOf struct {
	value *NetworkProtocolAllOf
	isSet bool
}

func (v NullableNetworkProtocolAllOf) Get() *NetworkProtocolAllOf {
	return v.value
}

func (v *NullableNetworkProtocolAllOf) Set(val *NetworkProtocolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkProtocolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkProtocolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkProtocolAllOf(val *NetworkProtocolAllOf) *NullableNetworkProtocolAllOf {
	return &NullableNetworkProtocolAllOf{value: val, isSet: true}
}

func (v NullableNetworkProtocolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkProtocolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


