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

// VResourceGroupNameModel struct for VResourceGroupNameModel
type VResourceGroupNameModel struct {
	VResourceGroupName *string `json:"VResourceGroupName,omitempty"`
}

// NewVResourceGroupNameModel instantiates a new VResourceGroupNameModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVResourceGroupNameModel() *VResourceGroupNameModel {
	this := VResourceGroupNameModel{}
	return &this
}

// NewVResourceGroupNameModelWithDefaults instantiates a new VResourceGroupNameModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVResourceGroupNameModelWithDefaults() *VResourceGroupNameModel {
	this := VResourceGroupNameModel{}
	return &this
}

// GetVResourceGroupName returns the VResourceGroupName field value if set, zero value otherwise.
func (o *VResourceGroupNameModel) GetVResourceGroupName() string {
	if o == nil || isNil(o.VResourceGroupName) {
		var ret string
		return ret
	}
	return *o.VResourceGroupName
}

// GetVResourceGroupNameOk returns a tuple with the VResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VResourceGroupNameModel) GetVResourceGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.VResourceGroupName) {
    return nil, false
	}
	return o.VResourceGroupName, true
}

// HasVResourceGroupName returns a boolean if a field has been set.
func (o *VResourceGroupNameModel) HasVResourceGroupName() bool {
	if o != nil && !isNil(o.VResourceGroupName) {
		return true
	}

	return false
}

// SetVResourceGroupName gets a reference to the given string and assigns it to the VResourceGroupName field.
func (o *VResourceGroupNameModel) SetVResourceGroupName(v string) {
	o.VResourceGroupName = &v
}

func (o VResourceGroupNameModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VResourceGroupName) {
		toSerialize["VResourceGroupName"] = o.VResourceGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableVResourceGroupNameModel struct {
	value *VResourceGroupNameModel
	isSet bool
}

func (v NullableVResourceGroupNameModel) Get() *VResourceGroupNameModel {
	return v.value
}

func (v *NullableVResourceGroupNameModel) Set(val *VResourceGroupNameModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVResourceGroupNameModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVResourceGroupNameModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVResourceGroupNameModel(val *VResourceGroupNameModel) *NullableVResourceGroupNameModel {
	return &NullableVResourceGroupNameModel{value: val, isSet: true}
}

func (v NullableVResourceGroupNameModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVResourceGroupNameModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


