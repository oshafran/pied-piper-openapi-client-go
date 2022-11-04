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

// SiteHealth struct for SiteHealth
type SiteHealth struct {
	FullConnectivity *int32 `json:"fullConnectivity,omitempty"`
	NoConnectivity *int32 `json:"noConnectivity,omitempty"`
	PartialConnectivity *int32 `json:"partialConnectivity,omitempty"`
}

// NewSiteHealth instantiates a new SiteHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteHealth() *SiteHealth {
	this := SiteHealth{}
	return &this
}

// NewSiteHealthWithDefaults instantiates a new SiteHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteHealthWithDefaults() *SiteHealth {
	this := SiteHealth{}
	return &this
}

// GetFullConnectivity returns the FullConnectivity field value if set, zero value otherwise.
func (o *SiteHealth) GetFullConnectivity() int32 {
	if o == nil || isNil(o.FullConnectivity) {
		var ret int32
		return ret
	}
	return *o.FullConnectivity
}

// GetFullConnectivityOk returns a tuple with the FullConnectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteHealth) GetFullConnectivityOk() (*int32, bool) {
	if o == nil || isNil(o.FullConnectivity) {
    return nil, false
	}
	return o.FullConnectivity, true
}

// HasFullConnectivity returns a boolean if a field has been set.
func (o *SiteHealth) HasFullConnectivity() bool {
	if o != nil && !isNil(o.FullConnectivity) {
		return true
	}

	return false
}

// SetFullConnectivity gets a reference to the given int32 and assigns it to the FullConnectivity field.
func (o *SiteHealth) SetFullConnectivity(v int32) {
	o.FullConnectivity = &v
}

// GetNoConnectivity returns the NoConnectivity field value if set, zero value otherwise.
func (o *SiteHealth) GetNoConnectivity() int32 {
	if o == nil || isNil(o.NoConnectivity) {
		var ret int32
		return ret
	}
	return *o.NoConnectivity
}

// GetNoConnectivityOk returns a tuple with the NoConnectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteHealth) GetNoConnectivityOk() (*int32, bool) {
	if o == nil || isNil(o.NoConnectivity) {
    return nil, false
	}
	return o.NoConnectivity, true
}

// HasNoConnectivity returns a boolean if a field has been set.
func (o *SiteHealth) HasNoConnectivity() bool {
	if o != nil && !isNil(o.NoConnectivity) {
		return true
	}

	return false
}

// SetNoConnectivity gets a reference to the given int32 and assigns it to the NoConnectivity field.
func (o *SiteHealth) SetNoConnectivity(v int32) {
	o.NoConnectivity = &v
}

// GetPartialConnectivity returns the PartialConnectivity field value if set, zero value otherwise.
func (o *SiteHealth) GetPartialConnectivity() int32 {
	if o == nil || isNil(o.PartialConnectivity) {
		var ret int32
		return ret
	}
	return *o.PartialConnectivity
}

// GetPartialConnectivityOk returns a tuple with the PartialConnectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteHealth) GetPartialConnectivityOk() (*int32, bool) {
	if o == nil || isNil(o.PartialConnectivity) {
    return nil, false
	}
	return o.PartialConnectivity, true
}

// HasPartialConnectivity returns a boolean if a field has been set.
func (o *SiteHealth) HasPartialConnectivity() bool {
	if o != nil && !isNil(o.PartialConnectivity) {
		return true
	}

	return false
}

// SetPartialConnectivity gets a reference to the given int32 and assigns it to the PartialConnectivity field.
func (o *SiteHealth) SetPartialConnectivity(v int32) {
	o.PartialConnectivity = &v
}

func (o SiteHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FullConnectivity) {
		toSerialize["fullConnectivity"] = o.FullConnectivity
	}
	if !isNil(o.NoConnectivity) {
		toSerialize["noConnectivity"] = o.NoConnectivity
	}
	if !isNil(o.PartialConnectivity) {
		toSerialize["partialConnectivity"] = o.PartialConnectivity
	}
	return json.Marshal(toSerialize)
}

type NullableSiteHealth struct {
	value *SiteHealth
	isSet bool
}

func (v NullableSiteHealth) Get() *SiteHealth {
	return v.value
}

func (v *NullableSiteHealth) Set(val *SiteHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteHealth(val *SiteHealth) *NullableSiteHealth {
	return &NullableSiteHealth{value: val, isSet: true}
}

func (v NullableSiteHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


