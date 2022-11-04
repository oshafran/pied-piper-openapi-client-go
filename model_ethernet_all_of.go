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

// EthernetAllOf struct for EthernetAllOf
type EthernetAllOf struct {
	EthernetInterfaceList []EthernetInterface `json:"ethernetInterfaceList,omitempty"`
}

// NewEthernetAllOf instantiates a new EthernetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthernetAllOf() *EthernetAllOf {
	this := EthernetAllOf{}
	return &this
}

// NewEthernetAllOfWithDefaults instantiates a new EthernetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthernetAllOfWithDefaults() *EthernetAllOf {
	this := EthernetAllOf{}
	return &this
}

// GetEthernetInterfaceList returns the EthernetInterfaceList field value if set, zero value otherwise.
func (o *EthernetAllOf) GetEthernetInterfaceList() []EthernetInterface {
	if o == nil || isNil(o.EthernetInterfaceList) {
		var ret []EthernetInterface
		return ret
	}
	return o.EthernetInterfaceList
}

// GetEthernetInterfaceListOk returns a tuple with the EthernetInterfaceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetAllOf) GetEthernetInterfaceListOk() ([]EthernetInterface, bool) {
	if o == nil || isNil(o.EthernetInterfaceList) {
    return nil, false
	}
	return o.EthernetInterfaceList, true
}

// HasEthernetInterfaceList returns a boolean if a field has been set.
func (o *EthernetAllOf) HasEthernetInterfaceList() bool {
	if o != nil && !isNil(o.EthernetInterfaceList) {
		return true
	}

	return false
}

// SetEthernetInterfaceList gets a reference to the given []EthernetInterface and assigns it to the EthernetInterfaceList field.
func (o *EthernetAllOf) SetEthernetInterfaceList(v []EthernetInterface) {
	o.EthernetInterfaceList = v
}

func (o EthernetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EthernetInterfaceList) {
		toSerialize["ethernetInterfaceList"] = o.EthernetInterfaceList
	}
	return json.Marshal(toSerialize)
}

type NullableEthernetAllOf struct {
	value *EthernetAllOf
	isSet bool
}

func (v NullableEthernetAllOf) Get() *EthernetAllOf {
	return v.value
}

func (v *NullableEthernetAllOf) Set(val *EthernetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEthernetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEthernetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthernetAllOf(val *EthernetAllOf) *NullableEthernetAllOf {
	return &NullableEthernetAllOf{value: val, isSet: true}
}

func (v NullableEthernetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthernetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


