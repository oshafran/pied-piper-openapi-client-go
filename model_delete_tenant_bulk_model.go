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

// DeleteTenantBulkModel struct for DeleteTenantBulkModel
type DeleteTenantBulkModel struct {
	Password *string `json:"password,omitempty"`
	TenantIdList []string `json:"tenantIdList,omitempty"`
}

// NewDeleteTenantBulkModel instantiates a new DeleteTenantBulkModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTenantBulkModel() *DeleteTenantBulkModel {
	this := DeleteTenantBulkModel{}
	return &this
}

// NewDeleteTenantBulkModelWithDefaults instantiates a new DeleteTenantBulkModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTenantBulkModelWithDefaults() *DeleteTenantBulkModel {
	this := DeleteTenantBulkModel{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DeleteTenantBulkModel) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTenantBulkModel) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DeleteTenantBulkModel) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DeleteTenantBulkModel) SetPassword(v string) {
	o.Password = &v
}

// GetTenantIdList returns the TenantIdList field value if set, zero value otherwise.
func (o *DeleteTenantBulkModel) GetTenantIdList() []string {
	if o == nil || isNil(o.TenantIdList) {
		var ret []string
		return ret
	}
	return o.TenantIdList
}

// GetTenantIdListOk returns a tuple with the TenantIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTenantBulkModel) GetTenantIdListOk() ([]string, bool) {
	if o == nil || isNil(o.TenantIdList) {
    return nil, false
	}
	return o.TenantIdList, true
}

// HasTenantIdList returns a boolean if a field has been set.
func (o *DeleteTenantBulkModel) HasTenantIdList() bool {
	if o != nil && !isNil(o.TenantIdList) {
		return true
	}

	return false
}

// SetTenantIdList gets a reference to the given []string and assigns it to the TenantIdList field.
func (o *DeleteTenantBulkModel) SetTenantIdList(v []string) {
	o.TenantIdList = v
}

func (o DeleteTenantBulkModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.TenantIdList) {
		toSerialize["tenantIdList"] = o.TenantIdList
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteTenantBulkModel struct {
	value *DeleteTenantBulkModel
	isSet bool
}

func (v NullableDeleteTenantBulkModel) Get() *DeleteTenantBulkModel {
	return v.value
}

func (v *NullableDeleteTenantBulkModel) Set(val *DeleteTenantBulkModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTenantBulkModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTenantBulkModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTenantBulkModel(val *DeleteTenantBulkModel) *NullableDeleteTenantBulkModel {
	return &NullableDeleteTenantBulkModel{value: val, isSet: true}
}

func (v NullableDeleteTenantBulkModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTenantBulkModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


