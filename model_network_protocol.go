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

// NetworkProtocol struct for NetworkProtocol
type NetworkProtocol struct {
	Name string `json:"name"`
	Type string `json:"type"`
	DHCPPool *DHCPPool `json:"DHCPPool,omitempty"`
	DNSSettings *string `json:"DNSSettings,omitempty"`
	NTPSettings []string `json:"NTPSettings,omitempty"`
	NATRules []NATRule `json:"NATRules,omitempty"`
	NTPInherit *bool `json:"NTPInherit,omitempty"`
}

// NewNetworkProtocol instantiates a new NetworkProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtocol(name string, type_ string) *NetworkProtocol {
	this := NetworkProtocol{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewNetworkProtocolWithDefaults instantiates a new NetworkProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtocolWithDefaults() *NetworkProtocol {
	this := NetworkProtocol{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtocol) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtocol) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *NetworkProtocol) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkProtocol) SetType(v string) {
	o.Type = v
}

// GetDHCPPool returns the DHCPPool field value if set, zero value otherwise.
func (o *NetworkProtocol) GetDHCPPool() DHCPPool {
	if o == nil || isNil(o.DHCPPool) {
		var ret DHCPPool
		return ret
	}
	return *o.DHCPPool
}

// GetDHCPPoolOk returns a tuple with the DHCPPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetDHCPPoolOk() (*DHCPPool, bool) {
	if o == nil || isNil(o.DHCPPool) {
    return nil, false
	}
	return o.DHCPPool, true
}

// HasDHCPPool returns a boolean if a field has been set.
func (o *NetworkProtocol) HasDHCPPool() bool {
	if o != nil && !isNil(o.DHCPPool) {
		return true
	}

	return false
}

// SetDHCPPool gets a reference to the given DHCPPool and assigns it to the DHCPPool field.
func (o *NetworkProtocol) SetDHCPPool(v DHCPPool) {
	o.DHCPPool = &v
}

// GetDNSSettings returns the DNSSettings field value if set, zero value otherwise.
func (o *NetworkProtocol) GetDNSSettings() string {
	if o == nil || isNil(o.DNSSettings) {
		var ret string
		return ret
	}
	return *o.DNSSettings
}

// GetDNSSettingsOk returns a tuple with the DNSSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetDNSSettingsOk() (*string, bool) {
	if o == nil || isNil(o.DNSSettings) {
    return nil, false
	}
	return o.DNSSettings, true
}

// HasDNSSettings returns a boolean if a field has been set.
func (o *NetworkProtocol) HasDNSSettings() bool {
	if o != nil && !isNil(o.DNSSettings) {
		return true
	}

	return false
}

// SetDNSSettings gets a reference to the given string and assigns it to the DNSSettings field.
func (o *NetworkProtocol) SetDNSSettings(v string) {
	o.DNSSettings = &v
}

// GetNTPSettings returns the NTPSettings field value if set, zero value otherwise.
func (o *NetworkProtocol) GetNTPSettings() []string {
	if o == nil || isNil(o.NTPSettings) {
		var ret []string
		return ret
	}
	return o.NTPSettings
}

// GetNTPSettingsOk returns a tuple with the NTPSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetNTPSettingsOk() ([]string, bool) {
	if o == nil || isNil(o.NTPSettings) {
    return nil, false
	}
	return o.NTPSettings, true
}

// HasNTPSettings returns a boolean if a field has been set.
func (o *NetworkProtocol) HasNTPSettings() bool {
	if o != nil && !isNil(o.NTPSettings) {
		return true
	}

	return false
}

// SetNTPSettings gets a reference to the given []string and assigns it to the NTPSettings field.
func (o *NetworkProtocol) SetNTPSettings(v []string) {
	o.NTPSettings = v
}

// GetNATRules returns the NATRules field value if set, zero value otherwise.
func (o *NetworkProtocol) GetNATRules() []NATRule {
	if o == nil || isNil(o.NATRules) {
		var ret []NATRule
		return ret
	}
	return o.NATRules
}

// GetNATRulesOk returns a tuple with the NATRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetNATRulesOk() ([]NATRule, bool) {
	if o == nil || isNil(o.NATRules) {
    return nil, false
	}
	return o.NATRules, true
}

// HasNATRules returns a boolean if a field has been set.
func (o *NetworkProtocol) HasNATRules() bool {
	if o != nil && !isNil(o.NATRules) {
		return true
	}

	return false
}

// SetNATRules gets a reference to the given []NATRule and assigns it to the NATRules field.
func (o *NetworkProtocol) SetNATRules(v []NATRule) {
	o.NATRules = v
}

// GetNTPInherit returns the NTPInherit field value if set, zero value otherwise.
func (o *NetworkProtocol) GetNTPInherit() bool {
	if o == nil || isNil(o.NTPInherit) {
		var ret bool
		return ret
	}
	return *o.NTPInherit
}

// GetNTPInheritOk returns a tuple with the NTPInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtocol) GetNTPInheritOk() (*bool, bool) {
	if o == nil || isNil(o.NTPInherit) {
    return nil, false
	}
	return o.NTPInherit, true
}

// HasNTPInherit returns a boolean if a field has been set.
func (o *NetworkProtocol) HasNTPInherit() bool {
	if o != nil && !isNil(o.NTPInherit) {
		return true
	}

	return false
}

// SetNTPInherit gets a reference to the given bool and assigns it to the NTPInherit field.
func (o *NetworkProtocol) SetNTPInherit(v bool) {
	o.NTPInherit = &v
}

func (o NetworkProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
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

type NullableNetworkProtocol struct {
	value *NetworkProtocol
	isSet bool
}

func (v NullableNetworkProtocol) Get() *NetworkProtocol {
	return v.value
}

func (v *NullableNetworkProtocol) Set(val *NetworkProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkProtocol(val *NetworkProtocol) *NullableNetworkProtocol {
	return &NullableNetworkProtocol{value: val, isSet: true}
}

func (v NullableNetworkProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


