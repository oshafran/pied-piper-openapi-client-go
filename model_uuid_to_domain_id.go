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

// UuidToDomainId struct for UuidToDomainId
type UuidToDomainId struct {
	Domain *string `json:"domain,omitempty"`
	Mapping []UuidToDomainIdMapping `json:"mapping,omitempty"`
}

// NewUuidToDomainId instantiates a new UuidToDomainId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidToDomainId() *UuidToDomainId {
	this := UuidToDomainId{}
	return &this
}

// NewUuidToDomainIdWithDefaults instantiates a new UuidToDomainId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidToDomainIdWithDefaults() *UuidToDomainId {
	this := UuidToDomainId{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *UuidToDomainId) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidToDomainId) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *UuidToDomainId) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *UuidToDomainId) SetDomain(v string) {
	o.Domain = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *UuidToDomainId) GetMapping() []UuidToDomainIdMapping {
	if o == nil || o.Mapping == nil {
		var ret []UuidToDomainIdMapping
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidToDomainId) GetMappingOk() ([]UuidToDomainIdMapping, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *UuidToDomainId) HasMapping() bool {
	if o != nil && o.Mapping != nil {
		return true
	}

	return false
}

// SetMapping gets a reference to the given []UuidToDomainIdMapping and assigns it to the Mapping field.
func (o *UuidToDomainId) SetMapping(v []UuidToDomainIdMapping) {
	o.Mapping = v
}

func (o UuidToDomainId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	return json.Marshal(toSerialize)
}

type NullableUuidToDomainId struct {
	value *UuidToDomainId
	isSet bool
}

func (v NullableUuidToDomainId) Get() *UuidToDomainId {
	return v.value
}

func (v *NullableUuidToDomainId) Set(val *UuidToDomainId) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidToDomainId) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidToDomainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidToDomainId(val *UuidToDomainId) *NullableUuidToDomainId {
	return &NullableUuidToDomainId{value: val, isSet: true}
}

func (v NullableUuidToDomainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidToDomainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

