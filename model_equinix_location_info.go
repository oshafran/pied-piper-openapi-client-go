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

// EquinixLocationInfo struct for EquinixLocationInfo
type EquinixLocationInfo struct {
	EqBillingAccountInfoList []MultiCloudEdgeBillingAccountInfo `json:"eqBillingAccountInfoList,omitempty"`
	EqNEInfo *EquinixNEInfo `json:"eqNEInfo,omitempty"`
	MetroCode *string `json:"metroCode,omitempty"`
	MetroName *string `json:"metroName,omitempty"`
	NetworkRegion *string `json:"networkRegion,omitempty"`
	SiteCode *string `json:"siteCode,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewEquinixLocationInfo instantiates a new EquinixLocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquinixLocationInfo() *EquinixLocationInfo {
	this := EquinixLocationInfo{}
	return &this
}

// NewEquinixLocationInfoWithDefaults instantiates a new EquinixLocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquinixLocationInfoWithDefaults() *EquinixLocationInfo {
	this := EquinixLocationInfo{}
	return &this
}

// GetEqBillingAccountInfoList returns the EqBillingAccountInfoList field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetEqBillingAccountInfoList() []MultiCloudEdgeBillingAccountInfo {
	if o == nil || isNil(o.EqBillingAccountInfoList) {
		var ret []MultiCloudEdgeBillingAccountInfo
		return ret
	}
	return o.EqBillingAccountInfoList
}

// GetEqBillingAccountInfoListOk returns a tuple with the EqBillingAccountInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetEqBillingAccountInfoListOk() ([]MultiCloudEdgeBillingAccountInfo, bool) {
	if o == nil || isNil(o.EqBillingAccountInfoList) {
    return nil, false
	}
	return o.EqBillingAccountInfoList, true
}

// HasEqBillingAccountInfoList returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasEqBillingAccountInfoList() bool {
	if o != nil && !isNil(o.EqBillingAccountInfoList) {
		return true
	}

	return false
}

// SetEqBillingAccountInfoList gets a reference to the given []MultiCloudEdgeBillingAccountInfo and assigns it to the EqBillingAccountInfoList field.
func (o *EquinixLocationInfo) SetEqBillingAccountInfoList(v []MultiCloudEdgeBillingAccountInfo) {
	o.EqBillingAccountInfoList = v
}

// GetEqNEInfo returns the EqNEInfo field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetEqNEInfo() EquinixNEInfo {
	if o == nil || isNil(o.EqNEInfo) {
		var ret EquinixNEInfo
		return ret
	}
	return *o.EqNEInfo
}

// GetEqNEInfoOk returns a tuple with the EqNEInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetEqNEInfoOk() (*EquinixNEInfo, bool) {
	if o == nil || isNil(o.EqNEInfo) {
    return nil, false
	}
	return o.EqNEInfo, true
}

// HasEqNEInfo returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasEqNEInfo() bool {
	if o != nil && !isNil(o.EqNEInfo) {
		return true
	}

	return false
}

// SetEqNEInfo gets a reference to the given EquinixNEInfo and assigns it to the EqNEInfo field.
func (o *EquinixLocationInfo) SetEqNEInfo(v EquinixNEInfo) {
	o.EqNEInfo = &v
}

// GetMetroCode returns the MetroCode field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetMetroCode() string {
	if o == nil || isNil(o.MetroCode) {
		var ret string
		return ret
	}
	return *o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetMetroCodeOk() (*string, bool) {
	if o == nil || isNil(o.MetroCode) {
    return nil, false
	}
	return o.MetroCode, true
}

// HasMetroCode returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasMetroCode() bool {
	if o != nil && !isNil(o.MetroCode) {
		return true
	}

	return false
}

// SetMetroCode gets a reference to the given string and assigns it to the MetroCode field.
func (o *EquinixLocationInfo) SetMetroCode(v string) {
	o.MetroCode = &v
}

// GetMetroName returns the MetroName field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetMetroName() string {
	if o == nil || isNil(o.MetroName) {
		var ret string
		return ret
	}
	return *o.MetroName
}

// GetMetroNameOk returns a tuple with the MetroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetMetroNameOk() (*string, bool) {
	if o == nil || isNil(o.MetroName) {
    return nil, false
	}
	return o.MetroName, true
}

// HasMetroName returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasMetroName() bool {
	if o != nil && !isNil(o.MetroName) {
		return true
	}

	return false
}

// SetMetroName gets a reference to the given string and assigns it to the MetroName field.
func (o *EquinixLocationInfo) SetMetroName(v string) {
	o.MetroName = &v
}

// GetNetworkRegion returns the NetworkRegion field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetNetworkRegion() string {
	if o == nil || isNil(o.NetworkRegion) {
		var ret string
		return ret
	}
	return *o.NetworkRegion
}

// GetNetworkRegionOk returns a tuple with the NetworkRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetNetworkRegionOk() (*string, bool) {
	if o == nil || isNil(o.NetworkRegion) {
    return nil, false
	}
	return o.NetworkRegion, true
}

// HasNetworkRegion returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasNetworkRegion() bool {
	if o != nil && !isNil(o.NetworkRegion) {
		return true
	}

	return false
}

// SetNetworkRegion gets a reference to the given string and assigns it to the NetworkRegion field.
func (o *EquinixLocationInfo) SetNetworkRegion(v string) {
	o.NetworkRegion = &v
}

// GetSiteCode returns the SiteCode field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetSiteCode() string {
	if o == nil || isNil(o.SiteCode) {
		var ret string
		return ret
	}
	return *o.SiteCode
}

// GetSiteCodeOk returns a tuple with the SiteCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetSiteCodeOk() (*string, bool) {
	if o == nil || isNil(o.SiteCode) {
    return nil, false
	}
	return o.SiteCode, true
}

// HasSiteCode returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasSiteCode() bool {
	if o != nil && !isNil(o.SiteCode) {
		return true
	}

	return false
}

// SetSiteCode gets a reference to the given string and assigns it to the SiteCode field.
func (o *EquinixLocationInfo) SetSiteCode(v string) {
	o.SiteCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EquinixLocationInfo) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EquinixLocationInfo) SetStatus(v string) {
	o.Status = &v
}

func (o EquinixLocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EqBillingAccountInfoList) {
		toSerialize["eqBillingAccountInfoList"] = o.EqBillingAccountInfoList
	}
	if !isNil(o.EqNEInfo) {
		toSerialize["eqNEInfo"] = o.EqNEInfo
	}
	if !isNil(o.MetroCode) {
		toSerialize["metroCode"] = o.MetroCode
	}
	if !isNil(o.MetroName) {
		toSerialize["metroName"] = o.MetroName
	}
	if !isNil(o.NetworkRegion) {
		toSerialize["networkRegion"] = o.NetworkRegion
	}
	if !isNil(o.SiteCode) {
		toSerialize["siteCode"] = o.SiteCode
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEquinixLocationInfo struct {
	value *EquinixLocationInfo
	isSet bool
}

func (v NullableEquinixLocationInfo) Get() *EquinixLocationInfo {
	return v.value
}

func (v *NullableEquinixLocationInfo) Set(val *EquinixLocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEquinixLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEquinixLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquinixLocationInfo(val *EquinixLocationInfo) *NullableEquinixLocationInfo {
	return &NullableEquinixLocationInfo{value: val, isSet: true}
}

func (v NullableEquinixLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquinixLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


