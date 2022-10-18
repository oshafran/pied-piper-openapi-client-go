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

// DeviceIP This is the valid DeviceIP
type DeviceIP struct {
	DeviceIp *string `json:"deviceIp,omitempty"`
}

// NewDeviceIP instantiates a new DeviceIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIP() *DeviceIP {
	this := DeviceIP{}
	return &this
}

// NewDeviceIPWithDefaults instantiates a new DeviceIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIPWithDefaults() *DeviceIP {
	this := DeviceIP{}
	return &this
}

// GetDeviceIp returns the DeviceIp field value if set, zero value otherwise.
func (o *DeviceIP) GetDeviceIp() string {
	if o == nil || o.DeviceIp == nil {
		var ret string
		return ret
	}
	return *o.DeviceIp
}

// GetDeviceIpOk returns a tuple with the DeviceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIP) GetDeviceIpOk() (*string, bool) {
	if o == nil || o.DeviceIp == nil {
		return nil, false
	}
	return o.DeviceIp, true
}

// HasDeviceIp returns a boolean if a field has been set.
func (o *DeviceIP) HasDeviceIp() bool {
	if o != nil && o.DeviceIp != nil {
		return true
	}

	return false
}

// SetDeviceIp gets a reference to the given string and assigns it to the DeviceIp field.
func (o *DeviceIP) SetDeviceIp(v string) {
	o.DeviceIp = &v
}

func (o DeviceIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceIp != nil {
		toSerialize["deviceIp"] = o.DeviceIp
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceIP struct {
	value *DeviceIP
	isSet bool
}

func (v NullableDeviceIP) Get() *DeviceIP {
	return v.value
}

func (v *NullableDeviceIP) Set(val *DeviceIP) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIP) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIP(val *DeviceIP) *NullableDeviceIP {
	return &NullableDeviceIP{value: val, isSet: true}
}

func (v NullableDeviceIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


