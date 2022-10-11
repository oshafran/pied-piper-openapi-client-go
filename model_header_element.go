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

// HeaderElement struct for HeaderElement
type HeaderElement struct {
	Name *string `json:"name,omitempty"`
	ParameterCount *int32 `json:"parameterCount,omitempty"`
	Parameters []NameValuePair `json:"parameters,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewHeaderElement instantiates a new HeaderElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderElement() *HeaderElement {
	this := HeaderElement{}
	return &this
}

// NewHeaderElementWithDefaults instantiates a new HeaderElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderElementWithDefaults() *HeaderElement {
	this := HeaderElement{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HeaderElement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderElement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HeaderElement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HeaderElement) SetName(v string) {
	o.Name = &v
}

// GetParameterCount returns the ParameterCount field value if set, zero value otherwise.
func (o *HeaderElement) GetParameterCount() int32 {
	if o == nil || o.ParameterCount == nil {
		var ret int32
		return ret
	}
	return *o.ParameterCount
}

// GetParameterCountOk returns a tuple with the ParameterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderElement) GetParameterCountOk() (*int32, bool) {
	if o == nil || o.ParameterCount == nil {
		return nil, false
	}
	return o.ParameterCount, true
}

// HasParameterCount returns a boolean if a field has been set.
func (o *HeaderElement) HasParameterCount() bool {
	if o != nil && o.ParameterCount != nil {
		return true
	}

	return false
}

// SetParameterCount gets a reference to the given int32 and assigns it to the ParameterCount field.
func (o *HeaderElement) SetParameterCount(v int32) {
	o.ParameterCount = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *HeaderElement) GetParameters() []NameValuePair {
	if o == nil || o.Parameters == nil {
		var ret []NameValuePair
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderElement) GetParametersOk() ([]NameValuePair, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *HeaderElement) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []NameValuePair and assigns it to the Parameters field.
func (o *HeaderElement) SetParameters(v []NameValuePair) {
	o.Parameters = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HeaderElement) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderElement) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HeaderElement) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HeaderElement) SetValue(v string) {
	o.Value = &v
}

func (o HeaderElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParameterCount != nil {
		toSerialize["parameterCount"] = o.ParameterCount
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHeaderElement struct {
	value *HeaderElement
	isSet bool
}

func (v NullableHeaderElement) Get() *HeaderElement {
	return v.value
}

func (v *NullableHeaderElement) Set(val *HeaderElement) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderElement) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderElement(val *HeaderElement) *NullableHeaderElement {
	return &NullableHeaderElement{value: val, isSet: true}
}

func (v NullableHeaderElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

