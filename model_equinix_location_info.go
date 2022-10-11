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
	if o == nil || o.EqBillingAccountInfoList == nil {
		var ret []MultiCloudEdgeBillingAccountInfo
		return ret
	}
	return o.EqBillingAccountInfoList
}

// GetEqBillingAccountInfoListOk returns a tuple with the EqBillingAccountInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetEqBillingAccountInfoListOk() ([]MultiCloudEdgeBillingAccountInfo, bool) {
	if o == nil || o.EqBillingAccountInfoList == nil {
		return nil, false
	}
	return o.EqBillingAccountInfoList, true
}

// HasEqBillingAccountInfoList returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasEqBillingAccountInfoList() bool {
	if o != nil && o.EqBillingAccountInfoList != nil {
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
	if o == nil || o.EqNEInfo == nil {
		var ret EquinixNEInfo
		return ret
	}
	return *o.EqNEInfo
}

// GetEqNEInfoOk returns a tuple with the EqNEInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetEqNEInfoOk() (*EquinixNEInfo, bool) {
	if o == nil || o.EqNEInfo == nil {
		return nil, false
	}
	return o.EqNEInfo, true
}

// HasEqNEInfo returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasEqNEInfo() bool {
	if o != nil && o.EqNEInfo != nil {
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
	if o == nil || o.MetroCode == nil {
		var ret string
		return ret
	}
	return *o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetMetroCodeOk() (*string, bool) {
	if o == nil || o.MetroCode == nil {
		return nil, false
	}
	return o.MetroCode, true
}

// HasMetroCode returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasMetroCode() bool {
	if o != nil && o.MetroCode != nil {
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
	if o == nil || o.MetroName == nil {
		var ret string
		return ret
	}
	return *o.MetroName
}

// GetMetroNameOk returns a tuple with the MetroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetMetroNameOk() (*string, bool) {
	if o == nil || o.MetroName == nil {
		return nil, false
	}
	return o.MetroName, true
}

// HasMetroName returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasMetroName() bool {
	if o != nil && o.MetroName != nil {
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
	if o == nil || o.NetworkRegion == nil {
		var ret string
		return ret
	}
	return *o.NetworkRegion
}

// GetNetworkRegionOk returns a tuple with the NetworkRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetNetworkRegionOk() (*string, bool) {
	if o == nil || o.NetworkRegion == nil {
		return nil, false
	}
	return o.NetworkRegion, true
}

// HasNetworkRegion returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasNetworkRegion() bool {
	if o != nil && o.NetworkRegion != nil {
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
	if o == nil || o.SiteCode == nil {
		var ret string
		return ret
	}
	return *o.SiteCode
}

// GetSiteCodeOk returns a tuple with the SiteCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetSiteCodeOk() (*string, bool) {
	if o == nil || o.SiteCode == nil {
		return nil, false
	}
	return o.SiteCode, true
}

// HasSiteCode returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasSiteCode() bool {
	if o != nil && o.SiteCode != nil {
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
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixLocationInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EquinixLocationInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
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
	if o.EqBillingAccountInfoList != nil {
		toSerialize["eqBillingAccountInfoList"] = o.EqBillingAccountInfoList
	}
	if o.EqNEInfo != nil {
		toSerialize["eqNEInfo"] = o.EqNEInfo
	}
	if o.MetroCode != nil {
		toSerialize["metroCode"] = o.MetroCode
	}
	if o.MetroName != nil {
		toSerialize["metroName"] = o.MetroName
	}
	if o.NetworkRegion != nil {
		toSerialize["networkRegion"] = o.NetworkRegion
	}
	if o.SiteCode != nil {
		toSerialize["siteCode"] = o.SiteCode
	}
	if o.Status != nil {
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

