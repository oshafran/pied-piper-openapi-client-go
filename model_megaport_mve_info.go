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

// MegaportMVEInfo struct for MegaportMVEInfo
type MegaportMVEInfo struct {
	ImageNameList []string `json:"imageNameList,omitempty"`
	ProductSizeList []string `json:"productSizeList,omitempty"`
}

// NewMegaportMVEInfo instantiates a new MegaportMVEInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMegaportMVEInfo() *MegaportMVEInfo {
	this := MegaportMVEInfo{}
	return &this
}

// NewMegaportMVEInfoWithDefaults instantiates a new MegaportMVEInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMegaportMVEInfoWithDefaults() *MegaportMVEInfo {
	this := MegaportMVEInfo{}
	return &this
}

// GetImageNameList returns the ImageNameList field value if set, zero value otherwise.
func (o *MegaportMVEInfo) GetImageNameList() []string {
	if o == nil || o.ImageNameList == nil {
		var ret []string
		return ret
	}
	return o.ImageNameList
}

// GetImageNameListOk returns a tuple with the ImageNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MegaportMVEInfo) GetImageNameListOk() ([]string, bool) {
	if o == nil || o.ImageNameList == nil {
		return nil, false
	}
	return o.ImageNameList, true
}

// HasImageNameList returns a boolean if a field has been set.
func (o *MegaportMVEInfo) HasImageNameList() bool {
	if o != nil && o.ImageNameList != nil {
		return true
	}

	return false
}

// SetImageNameList gets a reference to the given []string and assigns it to the ImageNameList field.
func (o *MegaportMVEInfo) SetImageNameList(v []string) {
	o.ImageNameList = v
}

// GetProductSizeList returns the ProductSizeList field value if set, zero value otherwise.
func (o *MegaportMVEInfo) GetProductSizeList() []string {
	if o == nil || o.ProductSizeList == nil {
		var ret []string
		return ret
	}
	return o.ProductSizeList
}

// GetProductSizeListOk returns a tuple with the ProductSizeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MegaportMVEInfo) GetProductSizeListOk() ([]string, bool) {
	if o == nil || o.ProductSizeList == nil {
		return nil, false
	}
	return o.ProductSizeList, true
}

// HasProductSizeList returns a boolean if a field has been set.
func (o *MegaportMVEInfo) HasProductSizeList() bool {
	if o != nil && o.ProductSizeList != nil {
		return true
	}

	return false
}

// SetProductSizeList gets a reference to the given []string and assigns it to the ProductSizeList field.
func (o *MegaportMVEInfo) SetProductSizeList(v []string) {
	o.ProductSizeList = v
}

func (o MegaportMVEInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageNameList != nil {
		toSerialize["imageNameList"] = o.ImageNameList
	}
	if o.ProductSizeList != nil {
		toSerialize["productSizeList"] = o.ProductSizeList
	}
	return json.Marshal(toSerialize)
}

type NullableMegaportMVEInfo struct {
	value *MegaportMVEInfo
	isSet bool
}

func (v NullableMegaportMVEInfo) Get() *MegaportMVEInfo {
	return v.value
}

func (v *NullableMegaportMVEInfo) Set(val *MegaportMVEInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMegaportMVEInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMegaportMVEInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMegaportMVEInfo(val *MegaportMVEInfo) *NullableMegaportMVEInfo {
	return &NullableMegaportMVEInfo{value: val, isSet: true}
}

func (v NullableMegaportMVEInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMegaportMVEInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

