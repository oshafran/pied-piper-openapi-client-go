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

// IpSecPolicy struct for IpSecPolicy
type IpSecPolicy struct {
	IkePhase1 IkePhase `json:"ikePhase1"`
	IkePhase2CipherSuite string `json:"ikePhase2CipherSuite"`
	Preset *string `json:"preset,omitempty"`
}

// NewIpSecPolicy instantiates a new IpSecPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpSecPolicy(ikePhase1 IkePhase, ikePhase2CipherSuite string) *IpSecPolicy {
	this := IpSecPolicy{}
	this.IkePhase1 = ikePhase1
	this.IkePhase2CipherSuite = ikePhase2CipherSuite
	return &this
}

// NewIpSecPolicyWithDefaults instantiates a new IpSecPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpSecPolicyWithDefaults() *IpSecPolicy {
	this := IpSecPolicy{}
	return &this
}

// GetIkePhase1 returns the IkePhase1 field value
func (o *IpSecPolicy) GetIkePhase1() IkePhase {
	if o == nil {
		var ret IkePhase
		return ret
	}

	return o.IkePhase1
}

// GetIkePhase1Ok returns a tuple with the IkePhase1 field value
// and a boolean to check if the value has been set.
func (o *IpSecPolicy) GetIkePhase1Ok() (*IkePhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IkePhase1, true
}

// SetIkePhase1 sets field value
func (o *IpSecPolicy) SetIkePhase1(v IkePhase) {
	o.IkePhase1 = v
}

// GetIkePhase2CipherSuite returns the IkePhase2CipherSuite field value
func (o *IpSecPolicy) GetIkePhase2CipherSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IkePhase2CipherSuite
}

// GetIkePhase2CipherSuiteOk returns a tuple with the IkePhase2CipherSuite field value
// and a boolean to check if the value has been set.
func (o *IpSecPolicy) GetIkePhase2CipherSuiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IkePhase2CipherSuite, true
}

// SetIkePhase2CipherSuite sets field value
func (o *IpSecPolicy) SetIkePhase2CipherSuite(v string) {
	o.IkePhase2CipherSuite = v
}

// GetPreset returns the Preset field value if set, zero value otherwise.
func (o *IpSecPolicy) GetPreset() string {
	if o == nil || o.Preset == nil {
		var ret string
		return ret
	}
	return *o.Preset
}

// GetPresetOk returns a tuple with the Preset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpSecPolicy) GetPresetOk() (*string, bool) {
	if o == nil || o.Preset == nil {
		return nil, false
	}
	return o.Preset, true
}

// HasPreset returns a boolean if a field has been set.
func (o *IpSecPolicy) HasPreset() bool {
	if o != nil && o.Preset != nil {
		return true
	}

	return false
}

// SetPreset gets a reference to the given string and assigns it to the Preset field.
func (o *IpSecPolicy) SetPreset(v string) {
	o.Preset = &v
}

func (o IpSecPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ikePhase1"] = o.IkePhase1
	}
	if true {
		toSerialize["ikePhase2CipherSuite"] = o.IkePhase2CipherSuite
	}
	if o.Preset != nil {
		toSerialize["preset"] = o.Preset
	}
	return json.Marshal(toSerialize)
}

type NullableIpSecPolicy struct {
	value *IpSecPolicy
	isSet bool
}

func (v NullableIpSecPolicy) Get() *IpSecPolicy {
	return v.value
}

func (v *NullableIpSecPolicy) Set(val *IpSecPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIpSecPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIpSecPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpSecPolicy(val *IpSecPolicy) *NullableIpSecPolicy {
	return &NullableIpSecPolicy{value: val, isSet: true}
}

func (v NullableIpSecPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpSecPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

