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

// TagRestfulResourceAllGroups struct for TagRestfulResourceAllGroups
type TagRestfulResourceAllGroups struct {
	Empty *bool `json:"empty,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
}

// NewTagRestfulResourceAllGroups instantiates a new TagRestfulResourceAllGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRestfulResourceAllGroups() *TagRestfulResourceAllGroups {
	this := TagRestfulResourceAllGroups{}
	return &this
}

// NewTagRestfulResourceAllGroupsWithDefaults instantiates a new TagRestfulResourceAllGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagRestfulResourceAllGroupsWithDefaults() *TagRestfulResourceAllGroups {
	this := TagRestfulResourceAllGroups{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *TagRestfulResourceAllGroups) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRestfulResourceAllGroups) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *TagRestfulResourceAllGroups) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *TagRestfulResourceAllGroups) SetEmpty(v bool) {
	o.Empty = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *TagRestfulResourceAllGroups) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRestfulResourceAllGroups) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *TagRestfulResourceAllGroups) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *TagRestfulResourceAllGroups) SetValueType(v string) {
	o.ValueType = &v
}

func (o TagRestfulResourceAllGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableTagRestfulResourceAllGroups struct {
	value *TagRestfulResourceAllGroups
	isSet bool
}

func (v NullableTagRestfulResourceAllGroups) Get() *TagRestfulResourceAllGroups {
	return v.value
}

func (v *NullableTagRestfulResourceAllGroups) Set(val *TagRestfulResourceAllGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRestfulResourceAllGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRestfulResourceAllGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRestfulResourceAllGroups(val *TagRestfulResourceAllGroups) *NullableTagRestfulResourceAllGroups {
	return &NullableTagRestfulResourceAllGroups{value: val, isSet: true}
}

func (v NullableTagRestfulResourceAllGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRestfulResourceAllGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


