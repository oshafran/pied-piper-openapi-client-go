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

// NwpiDomainMonitor struct for NwpiDomainMonitor
type NwpiDomainMonitor struct {
	ClientIp *string `json:"clientIp,omitempty"`
	DeviceToDomainId []UuidToDomainId `json:"deviceToDomainId,omitempty"`
	DomainAppGrp *string `json:"domainAppGrp,omitempty"`
	DomainAppVis *string `json:"domainAppVis,omitempty"`
	DomainList []DomainDetail `json:"domainList,omitempty"`
	TraceId *string `json:"traceId,omitempty"`
}

// NewNwpiDomainMonitor instantiates a new NwpiDomainMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwpiDomainMonitor() *NwpiDomainMonitor {
	this := NwpiDomainMonitor{}
	return &this
}

// NewNwpiDomainMonitorWithDefaults instantiates a new NwpiDomainMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwpiDomainMonitorWithDefaults() *NwpiDomainMonitor {
	this := NwpiDomainMonitor{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetClientIp() string {
	if o == nil || isNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetClientIpOk() (*string, bool) {
	if o == nil || isNil(o.ClientIp) {
    return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasClientIp() bool {
	if o != nil && !isNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *NwpiDomainMonitor) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetDeviceToDomainId returns the DeviceToDomainId field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetDeviceToDomainId() []UuidToDomainId {
	if o == nil || isNil(o.DeviceToDomainId) {
		var ret []UuidToDomainId
		return ret
	}
	return o.DeviceToDomainId
}

// GetDeviceToDomainIdOk returns a tuple with the DeviceToDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetDeviceToDomainIdOk() ([]UuidToDomainId, bool) {
	if o == nil || isNil(o.DeviceToDomainId) {
    return nil, false
	}
	return o.DeviceToDomainId, true
}

// HasDeviceToDomainId returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasDeviceToDomainId() bool {
	if o != nil && !isNil(o.DeviceToDomainId) {
		return true
	}

	return false
}

// SetDeviceToDomainId gets a reference to the given []UuidToDomainId and assigns it to the DeviceToDomainId field.
func (o *NwpiDomainMonitor) SetDeviceToDomainId(v []UuidToDomainId) {
	o.DeviceToDomainId = v
}

// GetDomainAppGrp returns the DomainAppGrp field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetDomainAppGrp() string {
	if o == nil || isNil(o.DomainAppGrp) {
		var ret string
		return ret
	}
	return *o.DomainAppGrp
}

// GetDomainAppGrpOk returns a tuple with the DomainAppGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetDomainAppGrpOk() (*string, bool) {
	if o == nil || isNil(o.DomainAppGrp) {
    return nil, false
	}
	return o.DomainAppGrp, true
}

// HasDomainAppGrp returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasDomainAppGrp() bool {
	if o != nil && !isNil(o.DomainAppGrp) {
		return true
	}

	return false
}

// SetDomainAppGrp gets a reference to the given string and assigns it to the DomainAppGrp field.
func (o *NwpiDomainMonitor) SetDomainAppGrp(v string) {
	o.DomainAppGrp = &v
}

// GetDomainAppVis returns the DomainAppVis field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetDomainAppVis() string {
	if o == nil || isNil(o.DomainAppVis) {
		var ret string
		return ret
	}
	return *o.DomainAppVis
}

// GetDomainAppVisOk returns a tuple with the DomainAppVis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetDomainAppVisOk() (*string, bool) {
	if o == nil || isNil(o.DomainAppVis) {
    return nil, false
	}
	return o.DomainAppVis, true
}

// HasDomainAppVis returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasDomainAppVis() bool {
	if o != nil && !isNil(o.DomainAppVis) {
		return true
	}

	return false
}

// SetDomainAppVis gets a reference to the given string and assigns it to the DomainAppVis field.
func (o *NwpiDomainMonitor) SetDomainAppVis(v string) {
	o.DomainAppVis = &v
}

// GetDomainList returns the DomainList field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetDomainList() []DomainDetail {
	if o == nil || isNil(o.DomainList) {
		var ret []DomainDetail
		return ret
	}
	return o.DomainList
}

// GetDomainListOk returns a tuple with the DomainList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetDomainListOk() ([]DomainDetail, bool) {
	if o == nil || isNil(o.DomainList) {
    return nil, false
	}
	return o.DomainList, true
}

// HasDomainList returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasDomainList() bool {
	if o != nil && !isNil(o.DomainList) {
		return true
	}

	return false
}

// SetDomainList gets a reference to the given []DomainDetail and assigns it to the DomainList field.
func (o *NwpiDomainMonitor) SetDomainList(v []DomainDetail) {
	o.DomainList = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *NwpiDomainMonitor) GetTraceId() string {
	if o == nil || isNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwpiDomainMonitor) GetTraceIdOk() (*string, bool) {
	if o == nil || isNil(o.TraceId) {
    return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *NwpiDomainMonitor) HasTraceId() bool {
	if o != nil && !isNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *NwpiDomainMonitor) SetTraceId(v string) {
	o.TraceId = &v
}

func (o NwpiDomainMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientIp) {
		toSerialize["clientIp"] = o.ClientIp
	}
	if !isNil(o.DeviceToDomainId) {
		toSerialize["deviceToDomainId"] = o.DeviceToDomainId
	}
	if !isNil(o.DomainAppGrp) {
		toSerialize["domainAppGrp"] = o.DomainAppGrp
	}
	if !isNil(o.DomainAppVis) {
		toSerialize["domainAppVis"] = o.DomainAppVis
	}
	if !isNil(o.DomainList) {
		toSerialize["domainList"] = o.DomainList
	}
	if !isNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	return json.Marshal(toSerialize)
}

type NullableNwpiDomainMonitor struct {
	value *NwpiDomainMonitor
	isSet bool
}

func (v NullableNwpiDomainMonitor) Get() *NwpiDomainMonitor {
	return v.value
}

func (v *NullableNwpiDomainMonitor) Set(val *NwpiDomainMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableNwpiDomainMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableNwpiDomainMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwpiDomainMonitor(val *NwpiDomainMonitor) *NullableNwpiDomainMonitor {
	return &NullableNwpiDomainMonitor{value: val, isSet: true}
}

func (v NullableNwpiDomainMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwpiDomainMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


