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

// Bfd struct for Bfd
type Bfd struct {
	BfdTimerOnTransportTunnels *BfdTimerOnTransportTunnels `json:"bfdTimerOnTransportTunnels,omitempty"`
	DscpForBfdPackets *int32 `json:"dscpForBfdPackets,omitempty"`
	Multiplier *int32 `json:"multiplier,omitempty"`
	PollInterval *int32 `json:"pollInterval,omitempty"`
}

// NewBfd instantiates a new Bfd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBfd() *Bfd {
	this := Bfd{}
	return &this
}

// NewBfdWithDefaults instantiates a new Bfd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBfdWithDefaults() *Bfd {
	this := Bfd{}
	return &this
}

// GetBfdTimerOnTransportTunnels returns the BfdTimerOnTransportTunnels field value if set, zero value otherwise.
func (o *Bfd) GetBfdTimerOnTransportTunnels() BfdTimerOnTransportTunnels {
	if o == nil || o.BfdTimerOnTransportTunnels == nil {
		var ret BfdTimerOnTransportTunnels
		return ret
	}
	return *o.BfdTimerOnTransportTunnels
}

// GetBfdTimerOnTransportTunnelsOk returns a tuple with the BfdTimerOnTransportTunnels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bfd) GetBfdTimerOnTransportTunnelsOk() (*BfdTimerOnTransportTunnels, bool) {
	if o == nil || o.BfdTimerOnTransportTunnels == nil {
		return nil, false
	}
	return o.BfdTimerOnTransportTunnels, true
}

// HasBfdTimerOnTransportTunnels returns a boolean if a field has been set.
func (o *Bfd) HasBfdTimerOnTransportTunnels() bool {
	if o != nil && o.BfdTimerOnTransportTunnels != nil {
		return true
	}

	return false
}

// SetBfdTimerOnTransportTunnels gets a reference to the given BfdTimerOnTransportTunnels and assigns it to the BfdTimerOnTransportTunnels field.
func (o *Bfd) SetBfdTimerOnTransportTunnels(v BfdTimerOnTransportTunnels) {
	o.BfdTimerOnTransportTunnels = &v
}

// GetDscpForBfdPackets returns the DscpForBfdPackets field value if set, zero value otherwise.
func (o *Bfd) GetDscpForBfdPackets() int32 {
	if o == nil || o.DscpForBfdPackets == nil {
		var ret int32
		return ret
	}
	return *o.DscpForBfdPackets
}

// GetDscpForBfdPacketsOk returns a tuple with the DscpForBfdPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bfd) GetDscpForBfdPacketsOk() (*int32, bool) {
	if o == nil || o.DscpForBfdPackets == nil {
		return nil, false
	}
	return o.DscpForBfdPackets, true
}

// HasDscpForBfdPackets returns a boolean if a field has been set.
func (o *Bfd) HasDscpForBfdPackets() bool {
	if o != nil && o.DscpForBfdPackets != nil {
		return true
	}

	return false
}

// SetDscpForBfdPackets gets a reference to the given int32 and assigns it to the DscpForBfdPackets field.
func (o *Bfd) SetDscpForBfdPackets(v int32) {
	o.DscpForBfdPackets = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *Bfd) GetMultiplier() int32 {
	if o == nil || o.Multiplier == nil {
		var ret int32
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bfd) GetMultiplierOk() (*int32, bool) {
	if o == nil || o.Multiplier == nil {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *Bfd) HasMultiplier() bool {
	if o != nil && o.Multiplier != nil {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given int32 and assigns it to the Multiplier field.
func (o *Bfd) SetMultiplier(v int32) {
	o.Multiplier = &v
}

// GetPollInterval returns the PollInterval field value if set, zero value otherwise.
func (o *Bfd) GetPollInterval() int32 {
	if o == nil || o.PollInterval == nil {
		var ret int32
		return ret
	}
	return *o.PollInterval
}

// GetPollIntervalOk returns a tuple with the PollInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bfd) GetPollIntervalOk() (*int32, bool) {
	if o == nil || o.PollInterval == nil {
		return nil, false
	}
	return o.PollInterval, true
}

// HasPollInterval returns a boolean if a field has been set.
func (o *Bfd) HasPollInterval() bool {
	if o != nil && o.PollInterval != nil {
		return true
	}

	return false
}

// SetPollInterval gets a reference to the given int32 and assigns it to the PollInterval field.
func (o *Bfd) SetPollInterval(v int32) {
	o.PollInterval = &v
}

func (o Bfd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BfdTimerOnTransportTunnels != nil {
		toSerialize["bfdTimerOnTransportTunnels"] = o.BfdTimerOnTransportTunnels
	}
	if o.DscpForBfdPackets != nil {
		toSerialize["dscpForBfdPackets"] = o.DscpForBfdPackets
	}
	if o.Multiplier != nil {
		toSerialize["multiplier"] = o.Multiplier
	}
	if o.PollInterval != nil {
		toSerialize["pollInterval"] = o.PollInterval
	}
	return json.Marshal(toSerialize)
}

type NullableBfd struct {
	value *Bfd
	isSet bool
}

func (v NullableBfd) Get() *Bfd {
	return v.value
}

func (v *NullableBfd) Set(val *Bfd) {
	v.value = val
	v.isSet = true
}

func (v NullableBfd) IsSet() bool {
	return v.isSet
}

func (v *NullableBfd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBfd(val *Bfd) *NullableBfd {
	return &NullableBfd{value: val, isSet: true}
}

func (v NullableBfd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBfd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


