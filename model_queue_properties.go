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

// QueueProperties struct for QueueProperties
type QueueProperties struct {
	CurrentSize *int32 `json:"current_size,omitempty"`
	MaxSize *int32 `json:"max_size,omitempty"`
	TenantCurrentSize *int32 `json:"tenant_current_size,omitempty"`
	TenantMaxSize *int32 `json:"tenant_max_size,omitempty"`
}

// NewQueueProperties instantiates a new QueueProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueProperties() *QueueProperties {
	this := QueueProperties{}
	return &this
}

// NewQueuePropertiesWithDefaults instantiates a new QueueProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueuePropertiesWithDefaults() *QueueProperties {
	this := QueueProperties{}
	return &this
}

// GetCurrentSize returns the CurrentSize field value if set, zero value otherwise.
func (o *QueueProperties) GetCurrentSize() int32 {
	if o == nil || o.CurrentSize == nil {
		var ret int32
		return ret
	}
	return *o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueProperties) GetCurrentSizeOk() (*int32, bool) {
	if o == nil || o.CurrentSize == nil {
		return nil, false
	}
	return o.CurrentSize, true
}

// HasCurrentSize returns a boolean if a field has been set.
func (o *QueueProperties) HasCurrentSize() bool {
	if o != nil && o.CurrentSize != nil {
		return true
	}

	return false
}

// SetCurrentSize gets a reference to the given int32 and assigns it to the CurrentSize field.
func (o *QueueProperties) SetCurrentSize(v int32) {
	o.CurrentSize = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *QueueProperties) GetMaxSize() int32 {
	if o == nil || o.MaxSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueProperties) GetMaxSizeOk() (*int32, bool) {
	if o == nil || o.MaxSize == nil {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *QueueProperties) HasMaxSize() bool {
	if o != nil && o.MaxSize != nil {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *QueueProperties) SetMaxSize(v int32) {
	o.MaxSize = &v
}

// GetTenantCurrentSize returns the TenantCurrentSize field value if set, zero value otherwise.
func (o *QueueProperties) GetTenantCurrentSize() int32 {
	if o == nil || o.TenantCurrentSize == nil {
		var ret int32
		return ret
	}
	return *o.TenantCurrentSize
}

// GetTenantCurrentSizeOk returns a tuple with the TenantCurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueProperties) GetTenantCurrentSizeOk() (*int32, bool) {
	if o == nil || o.TenantCurrentSize == nil {
		return nil, false
	}
	return o.TenantCurrentSize, true
}

// HasTenantCurrentSize returns a boolean if a field has been set.
func (o *QueueProperties) HasTenantCurrentSize() bool {
	if o != nil && o.TenantCurrentSize != nil {
		return true
	}

	return false
}

// SetTenantCurrentSize gets a reference to the given int32 and assigns it to the TenantCurrentSize field.
func (o *QueueProperties) SetTenantCurrentSize(v int32) {
	o.TenantCurrentSize = &v
}

// GetTenantMaxSize returns the TenantMaxSize field value if set, zero value otherwise.
func (o *QueueProperties) GetTenantMaxSize() int32 {
	if o == nil || o.TenantMaxSize == nil {
		var ret int32
		return ret
	}
	return *o.TenantMaxSize
}

// GetTenantMaxSizeOk returns a tuple with the TenantMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueProperties) GetTenantMaxSizeOk() (*int32, bool) {
	if o == nil || o.TenantMaxSize == nil {
		return nil, false
	}
	return o.TenantMaxSize, true
}

// HasTenantMaxSize returns a boolean if a field has been set.
func (o *QueueProperties) HasTenantMaxSize() bool {
	if o != nil && o.TenantMaxSize != nil {
		return true
	}

	return false
}

// SetTenantMaxSize gets a reference to the given int32 and assigns it to the TenantMaxSize field.
func (o *QueueProperties) SetTenantMaxSize(v int32) {
	o.TenantMaxSize = &v
}

func (o QueueProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentSize != nil {
		toSerialize["current_size"] = o.CurrentSize
	}
	if o.MaxSize != nil {
		toSerialize["max_size"] = o.MaxSize
	}
	if o.TenantCurrentSize != nil {
		toSerialize["tenant_current_size"] = o.TenantCurrentSize
	}
	if o.TenantMaxSize != nil {
		toSerialize["tenant_max_size"] = o.TenantMaxSize
	}
	return json.Marshal(toSerialize)
}

type NullableQueueProperties struct {
	value *QueueProperties
	isSet bool
}

func (v NullableQueueProperties) Get() *QueueProperties {
	return v.value
}

func (v *NullableQueueProperties) Set(val *QueueProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueProperties(val *QueueProperties) *NullableQueueProperties {
	return &NullableQueueProperties{value: val, isSet: true}
}

func (v NullableQueueProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


