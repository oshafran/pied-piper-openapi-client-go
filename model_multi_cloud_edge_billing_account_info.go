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

// MultiCloudEdgeBillingAccountInfo struct for MultiCloudEdgeBillingAccountInfo
type MultiCloudEdgeBillingAccountInfo struct {
	EdgeBillingAccountId *string `json:"edgeBillingAccountId,omitempty"`
	EdgeBillingAccountName *string `json:"edgeBillingAccountName,omitempty"`
	EdgeType *string `json:"edgeType,omitempty"`
}

// NewMultiCloudEdgeBillingAccountInfo instantiates a new MultiCloudEdgeBillingAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCloudEdgeBillingAccountInfo() *MultiCloudEdgeBillingAccountInfo {
	this := MultiCloudEdgeBillingAccountInfo{}
	return &this
}

// NewMultiCloudEdgeBillingAccountInfoWithDefaults instantiates a new MultiCloudEdgeBillingAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCloudEdgeBillingAccountInfoWithDefaults() *MultiCloudEdgeBillingAccountInfo {
	this := MultiCloudEdgeBillingAccountInfo{}
	return &this
}

// GetEdgeBillingAccountId returns the EdgeBillingAccountId field value if set, zero value otherwise.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeBillingAccountId() string {
	if o == nil || isNil(o.EdgeBillingAccountId) {
		var ret string
		return ret
	}
	return *o.EdgeBillingAccountId
}

// GetEdgeBillingAccountIdOk returns a tuple with the EdgeBillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeBillingAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.EdgeBillingAccountId) {
    return nil, false
	}
	return o.EdgeBillingAccountId, true
}

// HasEdgeBillingAccountId returns a boolean if a field has been set.
func (o *MultiCloudEdgeBillingAccountInfo) HasEdgeBillingAccountId() bool {
	if o != nil && !isNil(o.EdgeBillingAccountId) {
		return true
	}

	return false
}

// SetEdgeBillingAccountId gets a reference to the given string and assigns it to the EdgeBillingAccountId field.
func (o *MultiCloudEdgeBillingAccountInfo) SetEdgeBillingAccountId(v string) {
	o.EdgeBillingAccountId = &v
}

// GetEdgeBillingAccountName returns the EdgeBillingAccountName field value if set, zero value otherwise.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeBillingAccountName() string {
	if o == nil || isNil(o.EdgeBillingAccountName) {
		var ret string
		return ret
	}
	return *o.EdgeBillingAccountName
}

// GetEdgeBillingAccountNameOk returns a tuple with the EdgeBillingAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeBillingAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.EdgeBillingAccountName) {
    return nil, false
	}
	return o.EdgeBillingAccountName, true
}

// HasEdgeBillingAccountName returns a boolean if a field has been set.
func (o *MultiCloudEdgeBillingAccountInfo) HasEdgeBillingAccountName() bool {
	if o != nil && !isNil(o.EdgeBillingAccountName) {
		return true
	}

	return false
}

// SetEdgeBillingAccountName gets a reference to the given string and assigns it to the EdgeBillingAccountName field.
func (o *MultiCloudEdgeBillingAccountInfo) SetEdgeBillingAccountName(v string) {
	o.EdgeBillingAccountName = &v
}

// GetEdgeType returns the EdgeType field value if set, zero value otherwise.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeType() string {
	if o == nil || isNil(o.EdgeType) {
		var ret string
		return ret
	}
	return *o.EdgeType
}

// GetEdgeTypeOk returns a tuple with the EdgeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCloudEdgeBillingAccountInfo) GetEdgeTypeOk() (*string, bool) {
	if o == nil || isNil(o.EdgeType) {
    return nil, false
	}
	return o.EdgeType, true
}

// HasEdgeType returns a boolean if a field has been set.
func (o *MultiCloudEdgeBillingAccountInfo) HasEdgeType() bool {
	if o != nil && !isNil(o.EdgeType) {
		return true
	}

	return false
}

// SetEdgeType gets a reference to the given string and assigns it to the EdgeType field.
func (o *MultiCloudEdgeBillingAccountInfo) SetEdgeType(v string) {
	o.EdgeType = &v
}

func (o MultiCloudEdgeBillingAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EdgeBillingAccountId) {
		toSerialize["edgeBillingAccountId"] = o.EdgeBillingAccountId
	}
	if !isNil(o.EdgeBillingAccountName) {
		toSerialize["edgeBillingAccountName"] = o.EdgeBillingAccountName
	}
	if !isNil(o.EdgeType) {
		toSerialize["edgeType"] = o.EdgeType
	}
	return json.Marshal(toSerialize)
}

type NullableMultiCloudEdgeBillingAccountInfo struct {
	value *MultiCloudEdgeBillingAccountInfo
	isSet bool
}

func (v NullableMultiCloudEdgeBillingAccountInfo) Get() *MultiCloudEdgeBillingAccountInfo {
	return v.value
}

func (v *NullableMultiCloudEdgeBillingAccountInfo) Set(val *MultiCloudEdgeBillingAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCloudEdgeBillingAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCloudEdgeBillingAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCloudEdgeBillingAccountInfo(val *MultiCloudEdgeBillingAccountInfo) *NullableMultiCloudEdgeBillingAccountInfo {
	return &NullableMultiCloudEdgeBillingAccountInfo{value: val, isSet: true}
}

func (v NullableMultiCloudEdgeBillingAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiCloudEdgeBillingAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


