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

// Wifi struct for Wifi
type Wifi struct {
	SsidConfigList []SSIDConfig `json:"ssidConfigList,omitempty"`
	GuestWifi *GuestWifi `json:"guestWifi,omitempty"`
	CorporateWifi *CorporateWifi `json:"corporateWifi,omitempty"`
	AdvancedRadioSetting *AdvancedRadioSetting `json:"advancedRadioSetting,omitempty"`
}

// NewWifi instantiates a new Wifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWifi(name string, type_ string) *Wifi {
	this := Wifi{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewWifiWithDefaults instantiates a new Wifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWifiWithDefaults() *Wifi {
	this := Wifi{}
	return &this
}

// GetSsidConfigList returns the SsidConfigList field value if set, zero value otherwise.
func (o *Wifi) GetSsidConfigList() []SSIDConfig {
	if o == nil || isNil(o.SsidConfigList) {
		var ret []SSIDConfig
		return ret
	}
	return o.SsidConfigList
}

// GetSsidConfigListOk returns a tuple with the SsidConfigList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wifi) GetSsidConfigListOk() ([]SSIDConfig, bool) {
	if o == nil || isNil(o.SsidConfigList) {
    return nil, false
	}
	return o.SsidConfigList, true
}

// HasSsidConfigList returns a boolean if a field has been set.
func (o *Wifi) HasSsidConfigList() bool {
	if o != nil && !isNil(o.SsidConfigList) {
		return true
	}

	return false
}

// SetSsidConfigList gets a reference to the given []SSIDConfig and assigns it to the SsidConfigList field.
func (o *Wifi) SetSsidConfigList(v []SSIDConfig) {
	o.SsidConfigList = v
}

// GetGuestWifi returns the GuestWifi field value if set, zero value otherwise.
func (o *Wifi) GetGuestWifi() GuestWifi {
	if o == nil || isNil(o.GuestWifi) {
		var ret GuestWifi
		return ret
	}
	return *o.GuestWifi
}

// GetGuestWifiOk returns a tuple with the GuestWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wifi) GetGuestWifiOk() (*GuestWifi, bool) {
	if o == nil || isNil(o.GuestWifi) {
    return nil, false
	}
	return o.GuestWifi, true
}

// HasGuestWifi returns a boolean if a field has been set.
func (o *Wifi) HasGuestWifi() bool {
	if o != nil && !isNil(o.GuestWifi) {
		return true
	}

	return false
}

// SetGuestWifi gets a reference to the given GuestWifi and assigns it to the GuestWifi field.
func (o *Wifi) SetGuestWifi(v GuestWifi) {
	o.GuestWifi = &v
}

// GetCorporateWifi returns the CorporateWifi field value if set, zero value otherwise.
func (o *Wifi) GetCorporateWifi() CorporateWifi {
	if o == nil || isNil(o.CorporateWifi) {
		var ret CorporateWifi
		return ret
	}
	return *o.CorporateWifi
}

// GetCorporateWifiOk returns a tuple with the CorporateWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wifi) GetCorporateWifiOk() (*CorporateWifi, bool) {
	if o == nil || isNil(o.CorporateWifi) {
    return nil, false
	}
	return o.CorporateWifi, true
}

// HasCorporateWifi returns a boolean if a field has been set.
func (o *Wifi) HasCorporateWifi() bool {
	if o != nil && !isNil(o.CorporateWifi) {
		return true
	}

	return false
}

// SetCorporateWifi gets a reference to the given CorporateWifi and assigns it to the CorporateWifi field.
func (o *Wifi) SetCorporateWifi(v CorporateWifi) {
	o.CorporateWifi = &v
}

// GetAdvancedRadioSetting returns the AdvancedRadioSetting field value if set, zero value otherwise.
func (o *Wifi) GetAdvancedRadioSetting() AdvancedRadioSetting {
	if o == nil || isNil(o.AdvancedRadioSetting) {
		var ret AdvancedRadioSetting
		return ret
	}
	return *o.AdvancedRadioSetting
}

// GetAdvancedRadioSettingOk returns a tuple with the AdvancedRadioSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wifi) GetAdvancedRadioSettingOk() (*AdvancedRadioSetting, bool) {
	if o == nil || isNil(o.AdvancedRadioSetting) {
    return nil, false
	}
	return o.AdvancedRadioSetting, true
}

// HasAdvancedRadioSetting returns a boolean if a field has been set.
func (o *Wifi) HasAdvancedRadioSetting() bool {
	if o != nil && !isNil(o.AdvancedRadioSetting) {
		return true
	}

	return false
}

// SetAdvancedRadioSetting gets a reference to the given AdvancedRadioSetting and assigns it to the AdvancedRadioSetting field.
func (o *Wifi) SetAdvancedRadioSetting(v AdvancedRadioSetting) {
	o.AdvancedRadioSetting = &v
}

func (o Wifi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidConfigList) {
		toSerialize["ssidConfigList"] = o.SsidConfigList
	}
	if !isNil(o.GuestWifi) {
		toSerialize["guestWifi"] = o.GuestWifi
	}
	if !isNil(o.CorporateWifi) {
		toSerialize["corporateWifi"] = o.CorporateWifi
	}
	if !isNil(o.AdvancedRadioSetting) {
		toSerialize["advancedRadioSetting"] = o.AdvancedRadioSetting
	}
	return json.Marshal(toSerialize)
}

type NullableWifi struct {
	value *Wifi
	isSet bool
}

func (v NullableWifi) Get() *Wifi {
	return v.value
}

func (v *NullableWifi) Set(val *Wifi) {
	v.value = val
	v.isSet = true
}

func (v NullableWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWifi(val *Wifi) *NullableWifi {
	return &NullableWifi{value: val, isSet: true}
}

func (v NullableWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


