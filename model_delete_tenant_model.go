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

// DeleteTenantModel struct for DeleteTenantModel
type DeleteTenantModel struct {
	Password *string `json:"password,omitempty"`
}

// NewDeleteTenantModel instantiates a new DeleteTenantModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTenantModel() *DeleteTenantModel {
	this := DeleteTenantModel{}
	return &this
}

// NewDeleteTenantModelWithDefaults instantiates a new DeleteTenantModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTenantModelWithDefaults() *DeleteTenantModel {
	this := DeleteTenantModel{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DeleteTenantModel) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTenantModel) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DeleteTenantModel) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DeleteTenantModel) SetPassword(v string) {
	o.Password = &v
}

func (o DeleteTenantModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteTenantModel struct {
	value *DeleteTenantModel
	isSet bool
}

func (v NullableDeleteTenantModel) Get() *DeleteTenantModel {
	return v.value
}

func (v *NullableDeleteTenantModel) Set(val *DeleteTenantModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTenantModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTenantModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTenantModel(val *DeleteTenantModel) *NullableDeleteTenantModel {
	return &NullableDeleteTenantModel{value: val, isSet: true}
}

func (v NullableDeleteTenantModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTenantModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


