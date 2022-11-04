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

// ApiReKeyDeviceList struct for ApiReKeyDeviceList
type ApiReKeyDeviceList struct {
	DeviceList []ApiReKeyDevice `json:"deviceList,omitempty"`
}

// NewApiReKeyDeviceList instantiates a new ApiReKeyDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReKeyDeviceList() *ApiReKeyDeviceList {
	this := ApiReKeyDeviceList{}
	return &this
}

// NewApiReKeyDeviceListWithDefaults instantiates a new ApiReKeyDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReKeyDeviceListWithDefaults() *ApiReKeyDeviceList {
	this := ApiReKeyDeviceList{}
	return &this
}

// GetDeviceList returns the DeviceList field value if set, zero value otherwise.
func (o *ApiReKeyDeviceList) GetDeviceList() []ApiReKeyDevice {
	if o == nil || isNil(o.DeviceList) {
		var ret []ApiReKeyDevice
		return ret
	}
	return o.DeviceList
}

// GetDeviceListOk returns a tuple with the DeviceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReKeyDeviceList) GetDeviceListOk() ([]ApiReKeyDevice, bool) {
	if o == nil || isNil(o.DeviceList) {
    return nil, false
	}
	return o.DeviceList, true
}

// HasDeviceList returns a boolean if a field has been set.
func (o *ApiReKeyDeviceList) HasDeviceList() bool {
	if o != nil && !isNil(o.DeviceList) {
		return true
	}

	return false
}

// SetDeviceList gets a reference to the given []ApiReKeyDevice and assigns it to the DeviceList field.
func (o *ApiReKeyDeviceList) SetDeviceList(v []ApiReKeyDevice) {
	o.DeviceList = v
}

func (o ApiReKeyDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceList) {
		toSerialize["deviceList"] = o.DeviceList
	}
	return json.Marshal(toSerialize)
}

type NullableApiReKeyDeviceList struct {
	value *ApiReKeyDeviceList
	isSet bool
}

func (v NullableApiReKeyDeviceList) Get() *ApiReKeyDeviceList {
	return v.value
}

func (v *NullableApiReKeyDeviceList) Set(val *ApiReKeyDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReKeyDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReKeyDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReKeyDeviceList(val *ApiReKeyDeviceList) *NullableApiReKeyDeviceList {
	return &NullableApiReKeyDeviceList{value: val, isSet: true}
}

func (v NullableApiReKeyDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReKeyDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


