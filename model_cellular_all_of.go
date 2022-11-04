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

// CellularAllOf struct for CellularAllOf
type CellularAllOf struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	SimSlot0 *SimSlotConfig `json:"simSlot0,omitempty"`
	SimSlot1 *SimSlotConfig `json:"simSlot1,omitempty"`
	PrimarySlot *int32 `json:"primarySlot,omitempty"`
	WanConfig *string `json:"wanConfig,omitempty"`
}

// NewCellularAllOf instantiates a new CellularAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellularAllOf() *CellularAllOf {
	this := CellularAllOf{}
	return &this
}

// NewCellularAllOfWithDefaults instantiates a new CellularAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellularAllOfWithDefaults() *CellularAllOf {
	this := CellularAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CellularAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CellularAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CellularAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CellularAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CellularAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CellularAllOf) SetType(v string) {
	o.Type = &v
}

// GetSimSlot0 returns the SimSlot0 field value if set, zero value otherwise.
func (o *CellularAllOf) GetSimSlot0() SimSlotConfig {
	if o == nil || o.SimSlot0 == nil {
		var ret SimSlotConfig
		return ret
	}
	return *o.SimSlot0
}

// GetSimSlot0Ok returns a tuple with the SimSlot0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetSimSlot0Ok() (*SimSlotConfig, bool) {
	if o == nil || o.SimSlot0 == nil {
		return nil, false
	}
	return o.SimSlot0, true
}

// HasSimSlot0 returns a boolean if a field has been set.
func (o *CellularAllOf) HasSimSlot0() bool {
	if o != nil && o.SimSlot0 != nil {
		return true
	}

	return false
}

// SetSimSlot0 gets a reference to the given SimSlotConfig and assigns it to the SimSlot0 field.
func (o *CellularAllOf) SetSimSlot0(v SimSlotConfig) {
	o.SimSlot0 = &v
}

// GetSimSlot1 returns the SimSlot1 field value if set, zero value otherwise.
func (o *CellularAllOf) GetSimSlot1() SimSlotConfig {
	if o == nil || o.SimSlot1 == nil {
		var ret SimSlotConfig
		return ret
	}
	return *o.SimSlot1
}

// GetSimSlot1Ok returns a tuple with the SimSlot1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetSimSlot1Ok() (*SimSlotConfig, bool) {
	if o == nil || o.SimSlot1 == nil {
		return nil, false
	}
	return o.SimSlot1, true
}

// HasSimSlot1 returns a boolean if a field has been set.
func (o *CellularAllOf) HasSimSlot1() bool {
	if o != nil && o.SimSlot1 != nil {
		return true
	}

	return false
}

// SetSimSlot1 gets a reference to the given SimSlotConfig and assigns it to the SimSlot1 field.
func (o *CellularAllOf) SetSimSlot1(v SimSlotConfig) {
	o.SimSlot1 = &v
}

// GetPrimarySlot returns the PrimarySlot field value if set, zero value otherwise.
func (o *CellularAllOf) GetPrimarySlot() int32 {
	if o == nil || o.PrimarySlot == nil {
		var ret int32
		return ret
	}
	return *o.PrimarySlot
}

// GetPrimarySlotOk returns a tuple with the PrimarySlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetPrimarySlotOk() (*int32, bool) {
	if o == nil || o.PrimarySlot == nil {
		return nil, false
	}
	return o.PrimarySlot, true
}

// HasPrimarySlot returns a boolean if a field has been set.
func (o *CellularAllOf) HasPrimarySlot() bool {
	if o != nil && o.PrimarySlot != nil {
		return true
	}

	return false
}

// SetPrimarySlot gets a reference to the given int32 and assigns it to the PrimarySlot field.
func (o *CellularAllOf) SetPrimarySlot(v int32) {
	o.PrimarySlot = &v
}

// GetWanConfig returns the WanConfig field value if set, zero value otherwise.
func (o *CellularAllOf) GetWanConfig() string {
	if o == nil || o.WanConfig == nil {
		var ret string
		return ret
	}
	return *o.WanConfig
}

// GetWanConfigOk returns a tuple with the WanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularAllOf) GetWanConfigOk() (*string, bool) {
	if o == nil || o.WanConfig == nil {
		return nil, false
	}
	return o.WanConfig, true
}

// HasWanConfig returns a boolean if a field has been set.
func (o *CellularAllOf) HasWanConfig() bool {
	if o != nil && o.WanConfig != nil {
		return true
	}

	return false
}

// SetWanConfig gets a reference to the given string and assigns it to the WanConfig field.
func (o *CellularAllOf) SetWanConfig(v string) {
	o.WanConfig = &v
}

func (o CellularAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SimSlot0 != nil {
		toSerialize["simSlot0"] = o.SimSlot0
	}
	if o.SimSlot1 != nil {
		toSerialize["simSlot1"] = o.SimSlot1
	}
	if o.PrimarySlot != nil {
		toSerialize["primarySlot"] = o.PrimarySlot
	}
	if o.WanConfig != nil {
		toSerialize["wanConfig"] = o.WanConfig
	}
	return json.Marshal(toSerialize)
}

type NullableCellularAllOf struct {
	value *CellularAllOf
	isSet bool
}

func (v NullableCellularAllOf) Get() *CellularAllOf {
	return v.value
}

func (v *NullableCellularAllOf) Set(val *CellularAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCellularAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCellularAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellularAllOf(val *CellularAllOf) *NullableCellularAllOf {
	return &NullableCellularAllOf{value: val, isSet: true}
}

func (v NullableCellularAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellularAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


