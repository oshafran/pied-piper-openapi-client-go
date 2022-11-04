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

// DomainDetail struct for DomainDetail
type DomainDetail struct {
	Domain *string `json:"domain,omitempty"`
	ResolvedIp []string `json:"resolvedIp,omitempty"`
}

// NewDomainDetail instantiates a new DomainDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainDetail() *DomainDetail {
	this := DomainDetail{}
	return &this
}

// NewDomainDetailWithDefaults instantiates a new DomainDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainDetailWithDefaults() *DomainDetail {
	this := DomainDetail{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DomainDetail) GetDomain() string {
	if o == nil || isNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDomainOk() (*string, bool) {
	if o == nil || isNil(o.Domain) {
    return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainDetail) HasDomain() bool {
	if o != nil && !isNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DomainDetail) SetDomain(v string) {
	o.Domain = &v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *DomainDetail) GetResolvedIp() []string {
	if o == nil || isNil(o.ResolvedIp) {
		var ret []string
		return ret
	}
	return o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetResolvedIpOk() ([]string, bool) {
	if o == nil || isNil(o.ResolvedIp) {
    return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *DomainDetail) HasResolvedIp() bool {
	if o != nil && !isNil(o.ResolvedIp) {
		return true
	}

	return false
}

// SetResolvedIp gets a reference to the given []string and assigns it to the ResolvedIp field.
func (o *DomainDetail) SetResolvedIp(v []string) {
	o.ResolvedIp = v
}

func (o DomainDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !isNil(o.ResolvedIp) {
		toSerialize["resolvedIp"] = o.ResolvedIp
	}
	return json.Marshal(toSerialize)
}

type NullableDomainDetail struct {
	value *DomainDetail
	isSet bool
}

func (v NullableDomainDetail) Get() *DomainDetail {
	return v.value
}

func (v *NullableDomainDetail) Set(val *DomainDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainDetail(val *DomainDetail) *NullableDomainDetail {
	return &NullableDomainDetail{value: val, isSet: true}
}

func (v NullableDomainDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


