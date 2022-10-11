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

// DevicesHealthOverview struct for DevicesHealthOverview
type DevicesHealthOverview struct {
	Fair *int32 `json:"fair,omitempty"`
	Good *int32 `json:"good,omitempty"`
	Poor *int32 `json:"poor,omitempty"`
}

// NewDevicesHealthOverview instantiates a new DevicesHealthOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesHealthOverview() *DevicesHealthOverview {
	this := DevicesHealthOverview{}
	return &this
}

// NewDevicesHealthOverviewWithDefaults instantiates a new DevicesHealthOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesHealthOverviewWithDefaults() *DevicesHealthOverview {
	this := DevicesHealthOverview{}
	return &this
}

// GetFair returns the Fair field value if set, zero value otherwise.
func (o *DevicesHealthOverview) GetFair() int32 {
	if o == nil || o.Fair == nil {
		var ret int32
		return ret
	}
	return *o.Fair
}

// GetFairOk returns a tuple with the Fair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesHealthOverview) GetFairOk() (*int32, bool) {
	if o == nil || o.Fair == nil {
		return nil, false
	}
	return o.Fair, true
}

// HasFair returns a boolean if a field has been set.
func (o *DevicesHealthOverview) HasFair() bool {
	if o != nil && o.Fair != nil {
		return true
	}

	return false
}

// SetFair gets a reference to the given int32 and assigns it to the Fair field.
func (o *DevicesHealthOverview) SetFair(v int32) {
	o.Fair = &v
}

// GetGood returns the Good field value if set, zero value otherwise.
func (o *DevicesHealthOverview) GetGood() int32 {
	if o == nil || o.Good == nil {
		var ret int32
		return ret
	}
	return *o.Good
}

// GetGoodOk returns a tuple with the Good field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesHealthOverview) GetGoodOk() (*int32, bool) {
	if o == nil || o.Good == nil {
		return nil, false
	}
	return o.Good, true
}

// HasGood returns a boolean if a field has been set.
func (o *DevicesHealthOverview) HasGood() bool {
	if o != nil && o.Good != nil {
		return true
	}

	return false
}

// SetGood gets a reference to the given int32 and assigns it to the Good field.
func (o *DevicesHealthOverview) SetGood(v int32) {
	o.Good = &v
}

// GetPoor returns the Poor field value if set, zero value otherwise.
func (o *DevicesHealthOverview) GetPoor() int32 {
	if o == nil || o.Poor == nil {
		var ret int32
		return ret
	}
	return *o.Poor
}

// GetPoorOk returns a tuple with the Poor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesHealthOverview) GetPoorOk() (*int32, bool) {
	if o == nil || o.Poor == nil {
		return nil, false
	}
	return o.Poor, true
}

// HasPoor returns a boolean if a field has been set.
func (o *DevicesHealthOverview) HasPoor() bool {
	if o != nil && o.Poor != nil {
		return true
	}

	return false
}

// SetPoor gets a reference to the given int32 and assigns it to the Poor field.
func (o *DevicesHealthOverview) SetPoor(v int32) {
	o.Poor = &v
}

func (o DevicesHealthOverview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fair != nil {
		toSerialize["fair"] = o.Fair
	}
	if o.Good != nil {
		toSerialize["good"] = o.Good
	}
	if o.Poor != nil {
		toSerialize["poor"] = o.Poor
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesHealthOverview struct {
	value *DevicesHealthOverview
	isSet bool
}

func (v NullableDevicesHealthOverview) Get() *DevicesHealthOverview {
	return v.value
}

func (v *NullableDevicesHealthOverview) Set(val *DevicesHealthOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesHealthOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesHealthOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesHealthOverview(val *DevicesHealthOverview) *NullableDevicesHealthOverview {
	return &NullableDevicesHealthOverview{value: val, isSet: true}
}

func (v NullableDevicesHealthOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesHealthOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

