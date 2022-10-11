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

// IkePhase struct for IkePhase
type IkePhase struct {
	CipherSuite string `json:"cipherSuite"`
	DiffeHellmanGroup *string `json:"diffeHellmanGroup,omitempty"`
	IkeVersion *int32 `json:"ikeVersion,omitempty"`
	RekeyTimer *int32 `json:"rekeyTimer,omitempty"`
}

// NewIkePhase instantiates a new IkePhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIkePhase(cipherSuite string) *IkePhase {
	this := IkePhase{}
	this.CipherSuite = cipherSuite
	return &this
}

// NewIkePhaseWithDefaults instantiates a new IkePhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIkePhaseWithDefaults() *IkePhase {
	this := IkePhase{}
	return &this
}

// GetCipherSuite returns the CipherSuite field value
func (o *IkePhase) GetCipherSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CipherSuite
}

// GetCipherSuiteOk returns a tuple with the CipherSuite field value
// and a boolean to check if the value has been set.
func (o *IkePhase) GetCipherSuiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CipherSuite, true
}

// SetCipherSuite sets field value
func (o *IkePhase) SetCipherSuite(v string) {
	o.CipherSuite = v
}

// GetDiffeHellmanGroup returns the DiffeHellmanGroup field value if set, zero value otherwise.
func (o *IkePhase) GetDiffeHellmanGroup() string {
	if o == nil || o.DiffeHellmanGroup == nil {
		var ret string
		return ret
	}
	return *o.DiffeHellmanGroup
}

// GetDiffeHellmanGroupOk returns a tuple with the DiffeHellmanGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IkePhase) GetDiffeHellmanGroupOk() (*string, bool) {
	if o == nil || o.DiffeHellmanGroup == nil {
		return nil, false
	}
	return o.DiffeHellmanGroup, true
}

// HasDiffeHellmanGroup returns a boolean if a field has been set.
func (o *IkePhase) HasDiffeHellmanGroup() bool {
	if o != nil && o.DiffeHellmanGroup != nil {
		return true
	}

	return false
}

// SetDiffeHellmanGroup gets a reference to the given string and assigns it to the DiffeHellmanGroup field.
func (o *IkePhase) SetDiffeHellmanGroup(v string) {
	o.DiffeHellmanGroup = &v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *IkePhase) GetIkeVersion() int32 {
	if o == nil || o.IkeVersion == nil {
		var ret int32
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IkePhase) GetIkeVersionOk() (*int32, bool) {
	if o == nil || o.IkeVersion == nil {
		return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *IkePhase) HasIkeVersion() bool {
	if o != nil && o.IkeVersion != nil {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given int32 and assigns it to the IkeVersion field.
func (o *IkePhase) SetIkeVersion(v int32) {
	o.IkeVersion = &v
}

// GetRekeyTimer returns the RekeyTimer field value if set, zero value otherwise.
func (o *IkePhase) GetRekeyTimer() int32 {
	if o == nil || o.RekeyTimer == nil {
		var ret int32
		return ret
	}
	return *o.RekeyTimer
}

// GetRekeyTimerOk returns a tuple with the RekeyTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IkePhase) GetRekeyTimerOk() (*int32, bool) {
	if o == nil || o.RekeyTimer == nil {
		return nil, false
	}
	return o.RekeyTimer, true
}

// HasRekeyTimer returns a boolean if a field has been set.
func (o *IkePhase) HasRekeyTimer() bool {
	if o != nil && o.RekeyTimer != nil {
		return true
	}

	return false
}

// SetRekeyTimer gets a reference to the given int32 and assigns it to the RekeyTimer field.
func (o *IkePhase) SetRekeyTimer(v int32) {
	o.RekeyTimer = &v
}

func (o IkePhase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cipherSuite"] = o.CipherSuite
	}
	if o.DiffeHellmanGroup != nil {
		toSerialize["diffeHellmanGroup"] = o.DiffeHellmanGroup
	}
	if o.IkeVersion != nil {
		toSerialize["ikeVersion"] = o.IkeVersion
	}
	if o.RekeyTimer != nil {
		toSerialize["rekeyTimer"] = o.RekeyTimer
	}
	return json.Marshal(toSerialize)
}

type NullableIkePhase struct {
	value *IkePhase
	isSet bool
}

func (v NullableIkePhase) Get() *IkePhase {
	return v.value
}

func (v *NullableIkePhase) Set(val *IkePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableIkePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableIkePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIkePhase(val *IkePhase) *NullableIkePhase {
	return &NullableIkePhase{value: val, isSet: true}
}

func (v NullableIkePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIkePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

