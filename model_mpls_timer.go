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

// MplsTimer struct for MplsTimer
type MplsTimer struct {
	Dscp *int32 `json:"dscp,omitempty"`
	HelloInterval *int32 `json:"helloInterval,omitempty"`
	Multiplier *int32 `json:"multiplier,omitempty"`
	PathMtu *string `json:"pathMtu,omitempty"`
}

// NewMplsTimer instantiates a new MplsTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMplsTimer() *MplsTimer {
	this := MplsTimer{}
	return &this
}

// NewMplsTimerWithDefaults instantiates a new MplsTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMplsTimerWithDefaults() *MplsTimer {
	this := MplsTimer{}
	return &this
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *MplsTimer) GetDscp() int32 {
	if o == nil || o.Dscp == nil {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MplsTimer) GetDscpOk() (*int32, bool) {
	if o == nil || o.Dscp == nil {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *MplsTimer) HasDscp() bool {
	if o != nil && o.Dscp != nil {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *MplsTimer) SetDscp(v int32) {
	o.Dscp = &v
}

// GetHelloInterval returns the HelloInterval field value if set, zero value otherwise.
func (o *MplsTimer) GetHelloInterval() int32 {
	if o == nil || o.HelloInterval == nil {
		var ret int32
		return ret
	}
	return *o.HelloInterval
}

// GetHelloIntervalOk returns a tuple with the HelloInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MplsTimer) GetHelloIntervalOk() (*int32, bool) {
	if o == nil || o.HelloInterval == nil {
		return nil, false
	}
	return o.HelloInterval, true
}

// HasHelloInterval returns a boolean if a field has been set.
func (o *MplsTimer) HasHelloInterval() bool {
	if o != nil && o.HelloInterval != nil {
		return true
	}

	return false
}

// SetHelloInterval gets a reference to the given int32 and assigns it to the HelloInterval field.
func (o *MplsTimer) SetHelloInterval(v int32) {
	o.HelloInterval = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *MplsTimer) GetMultiplier() int32 {
	if o == nil || o.Multiplier == nil {
		var ret int32
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MplsTimer) GetMultiplierOk() (*int32, bool) {
	if o == nil || o.Multiplier == nil {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *MplsTimer) HasMultiplier() bool {
	if o != nil && o.Multiplier != nil {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given int32 and assigns it to the Multiplier field.
func (o *MplsTimer) SetMultiplier(v int32) {
	o.Multiplier = &v
}

// GetPathMtu returns the PathMtu field value if set, zero value otherwise.
func (o *MplsTimer) GetPathMtu() string {
	if o == nil || o.PathMtu == nil {
		var ret string
		return ret
	}
	return *o.PathMtu
}

// GetPathMtuOk returns a tuple with the PathMtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MplsTimer) GetPathMtuOk() (*string, bool) {
	if o == nil || o.PathMtu == nil {
		return nil, false
	}
	return o.PathMtu, true
}

// HasPathMtu returns a boolean if a field has been set.
func (o *MplsTimer) HasPathMtu() bool {
	if o != nil && o.PathMtu != nil {
		return true
	}

	return false
}

// SetPathMtu gets a reference to the given string and assigns it to the PathMtu field.
func (o *MplsTimer) SetPathMtu(v string) {
	o.PathMtu = &v
}

func (o MplsTimer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dscp != nil {
		toSerialize["dscp"] = o.Dscp
	}
	if o.HelloInterval != nil {
		toSerialize["helloInterval"] = o.HelloInterval
	}
	if o.Multiplier != nil {
		toSerialize["multiplier"] = o.Multiplier
	}
	if o.PathMtu != nil {
		toSerialize["pathMtu"] = o.PathMtu
	}
	return json.Marshal(toSerialize)
}

type NullableMplsTimer struct {
	value *MplsTimer
	isSet bool
}

func (v NullableMplsTimer) Get() *MplsTimer {
	return v.value
}

func (v *NullableMplsTimer) Set(val *MplsTimer) {
	v.value = val
	v.isSet = true
}

func (v NullableMplsTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableMplsTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMplsTimer(val *MplsTimer) *NullableMplsTimer {
	return &NullableMplsTimer{value: val, isSet: true}
}

func (v NullableMplsTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMplsTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

