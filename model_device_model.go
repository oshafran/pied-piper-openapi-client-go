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

// DeviceModel This is the valid DeviceModel 
type DeviceModel struct {
	DeviceModel *string `json:"deviceModel,omitempty"`
}

// NewDeviceModel instantiates a new DeviceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceModel() *DeviceModel {
	this := DeviceModel{}
	return &this
}

// NewDeviceModelWithDefaults instantiates a new DeviceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceModelWithDefaults() *DeviceModel {
	this := DeviceModel{}
	return &this
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *DeviceModel) GetDeviceModel() string {
	if o == nil || isNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModel) GetDeviceModelOk() (*string, bool) {
	if o == nil || isNil(o.DeviceModel) {
    return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *DeviceModel) HasDeviceModel() bool {
	if o != nil && !isNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *DeviceModel) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

func (o DeviceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceModel struct {
	value *DeviceModel
	isSet bool
}

func (v NullableDeviceModel) Get() *DeviceModel {
	return v.value
}

func (v *NullableDeviceModel) Set(val *DeviceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceModel(val *DeviceModel) *NullableDeviceModel {
	return &NullableDeviceModel{value: val, isSet: true}
}

func (v NullableDeviceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


