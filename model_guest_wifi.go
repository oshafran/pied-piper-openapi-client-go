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

// GuestWifi struct for GuestWifi
type GuestWifi struct {
	SecurityAuthType *string `json:"securityAuthType,omitempty"`
	Ssid *string `json:"ssid,omitempty"`
	Visibility *bool `json:"visibility,omitempty"`
	WpaPskKey *string `json:"wpaPskKey,omitempty"`
}

// NewGuestWifi instantiates a new GuestWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestWifi() *GuestWifi {
	this := GuestWifi{}
	return &this
}

// NewGuestWifiWithDefaults instantiates a new GuestWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestWifiWithDefaults() *GuestWifi {
	this := GuestWifi{}
	return &this
}

// GetSecurityAuthType returns the SecurityAuthType field value if set, zero value otherwise.
func (o *GuestWifi) GetSecurityAuthType() string {
	if o == nil || isNil(o.SecurityAuthType) {
		var ret string
		return ret
	}
	return *o.SecurityAuthType
}

// GetSecurityAuthTypeOk returns a tuple with the SecurityAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestWifi) GetSecurityAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.SecurityAuthType) {
    return nil, false
	}
	return o.SecurityAuthType, true
}

// HasSecurityAuthType returns a boolean if a field has been set.
func (o *GuestWifi) HasSecurityAuthType() bool {
	if o != nil && !isNil(o.SecurityAuthType) {
		return true
	}

	return false
}

// SetSecurityAuthType gets a reference to the given string and assigns it to the SecurityAuthType field.
func (o *GuestWifi) SetSecurityAuthType(v string) {
	o.SecurityAuthType = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GuestWifi) GetSsid() string {
	if o == nil || isNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestWifi) GetSsidOk() (*string, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GuestWifi) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *GuestWifi) SetSsid(v string) {
	o.Ssid = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GuestWifi) GetVisibility() bool {
	if o == nil || isNil(o.Visibility) {
		var ret bool
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestWifi) GetVisibilityOk() (*bool, bool) {
	if o == nil || isNil(o.Visibility) {
    return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *GuestWifi) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given bool and assigns it to the Visibility field.
func (o *GuestWifi) SetVisibility(v bool) {
	o.Visibility = &v
}

// GetWpaPskKey returns the WpaPskKey field value if set, zero value otherwise.
func (o *GuestWifi) GetWpaPskKey() string {
	if o == nil || isNil(o.WpaPskKey) {
		var ret string
		return ret
	}
	return *o.WpaPskKey
}

// GetWpaPskKeyOk returns a tuple with the WpaPskKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestWifi) GetWpaPskKeyOk() (*string, bool) {
	if o == nil || isNil(o.WpaPskKey) {
    return nil, false
	}
	return o.WpaPskKey, true
}

// HasWpaPskKey returns a boolean if a field has been set.
func (o *GuestWifi) HasWpaPskKey() bool {
	if o != nil && !isNil(o.WpaPskKey) {
		return true
	}

	return false
}

// SetWpaPskKey gets a reference to the given string and assigns it to the WpaPskKey field.
func (o *GuestWifi) SetWpaPskKey(v string) {
	o.WpaPskKey = &v
}

func (o GuestWifi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SecurityAuthType) {
		toSerialize["securityAuthType"] = o.SecurityAuthType
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !isNil(o.WpaPskKey) {
		toSerialize["wpaPskKey"] = o.WpaPskKey
	}
	return json.Marshal(toSerialize)
}

type NullableGuestWifi struct {
	value *GuestWifi
	isSet bool
}

func (v NullableGuestWifi) Get() *GuestWifi {
	return v.value
}

func (v *NullableGuestWifi) Set(val *GuestWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestWifi(val *GuestWifi) *NullableGuestWifi {
	return &NullableGuestWifi{value: val, isSet: true}
}

func (v NullableGuestWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


