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

// EquinixNEInfo struct for EquinixNEInfo
type EquinixNEInfo struct {
	ImageNameList []string `json:"imageNameList,omitempty"`
	ProductSizeList []string `json:"productSizeList,omitempty"`
}

// NewEquinixNEInfo instantiates a new EquinixNEInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquinixNEInfo() *EquinixNEInfo {
	this := EquinixNEInfo{}
	return &this
}

// NewEquinixNEInfoWithDefaults instantiates a new EquinixNEInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquinixNEInfoWithDefaults() *EquinixNEInfo {
	this := EquinixNEInfo{}
	return &this
}

// GetImageNameList returns the ImageNameList field value if set, zero value otherwise.
func (o *EquinixNEInfo) GetImageNameList() []string {
	if o == nil || o.ImageNameList == nil {
		var ret []string
		return ret
	}
	return o.ImageNameList
}

// GetImageNameListOk returns a tuple with the ImageNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixNEInfo) GetImageNameListOk() ([]string, bool) {
	if o == nil || o.ImageNameList == nil {
		return nil, false
	}
	return o.ImageNameList, true
}

// HasImageNameList returns a boolean if a field has been set.
func (o *EquinixNEInfo) HasImageNameList() bool {
	if o != nil && o.ImageNameList != nil {
		return true
	}

	return false
}

// SetImageNameList gets a reference to the given []string and assigns it to the ImageNameList field.
func (o *EquinixNEInfo) SetImageNameList(v []string) {
	o.ImageNameList = v
}

// GetProductSizeList returns the ProductSizeList field value if set, zero value otherwise.
func (o *EquinixNEInfo) GetProductSizeList() []string {
	if o == nil || o.ProductSizeList == nil {
		var ret []string
		return ret
	}
	return o.ProductSizeList
}

// GetProductSizeListOk returns a tuple with the ProductSizeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixNEInfo) GetProductSizeListOk() ([]string, bool) {
	if o == nil || o.ProductSizeList == nil {
		return nil, false
	}
	return o.ProductSizeList, true
}

// HasProductSizeList returns a boolean if a field has been set.
func (o *EquinixNEInfo) HasProductSizeList() bool {
	if o != nil && o.ProductSizeList != nil {
		return true
	}

	return false
}

// SetProductSizeList gets a reference to the given []string and assigns it to the ProductSizeList field.
func (o *EquinixNEInfo) SetProductSizeList(v []string) {
	o.ProductSizeList = v
}

func (o EquinixNEInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageNameList != nil {
		toSerialize["imageNameList"] = o.ImageNameList
	}
	if o.ProductSizeList != nil {
		toSerialize["productSizeList"] = o.ProductSizeList
	}
	return json.Marshal(toSerialize)
}

type NullableEquinixNEInfo struct {
	value *EquinixNEInfo
	isSet bool
}

func (v NullableEquinixNEInfo) Get() *EquinixNEInfo {
	return v.value
}

func (v *NullableEquinixNEInfo) Set(val *EquinixNEInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEquinixNEInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEquinixNEInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquinixNEInfo(val *EquinixNEInfo) *NullableEquinixNEInfo {
	return &NullableEquinixNEInfo{value: val, isSet: true}
}

func (v NullableEquinixNEInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquinixNEInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


