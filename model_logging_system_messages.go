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

// LoggingSystemMessages struct for LoggingSystemMessages
type LoggingSystemMessages struct {
	MaxFileSize *int32 `json:"maxFileSize,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Rotations *int32 `json:"rotations,omitempty"`
}

// NewLoggingSystemMessages instantiates a new LoggingSystemMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSystemMessages() *LoggingSystemMessages {
	this := LoggingSystemMessages{}
	return &this
}

// NewLoggingSystemMessagesWithDefaults instantiates a new LoggingSystemMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSystemMessagesWithDefaults() *LoggingSystemMessages {
	this := LoggingSystemMessages{}
	return &this
}

// GetMaxFileSize returns the MaxFileSize field value if set, zero value otherwise.
func (o *LoggingSystemMessages) GetMaxFileSize() int32 {
	if o == nil || isNil(o.MaxFileSize) {
		var ret int32
		return ret
	}
	return *o.MaxFileSize
}

// GetMaxFileSizeOk returns a tuple with the MaxFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSystemMessages) GetMaxFileSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxFileSize) {
    return nil, false
	}
	return o.MaxFileSize, true
}

// HasMaxFileSize returns a boolean if a field has been set.
func (o *LoggingSystemMessages) HasMaxFileSize() bool {
	if o != nil && !isNil(o.MaxFileSize) {
		return true
	}

	return false
}

// SetMaxFileSize gets a reference to the given int32 and assigns it to the MaxFileSize field.
func (o *LoggingSystemMessages) SetMaxFileSize(v int32) {
	o.MaxFileSize = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LoggingSystemMessages) GetPriority() string {
	if o == nil || isNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSystemMessages) GetPriorityOk() (*string, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LoggingSystemMessages) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *LoggingSystemMessages) SetPriority(v string) {
	o.Priority = &v
}

// GetRotations returns the Rotations field value if set, zero value otherwise.
func (o *LoggingSystemMessages) GetRotations() int32 {
	if o == nil || isNil(o.Rotations) {
		var ret int32
		return ret
	}
	return *o.Rotations
}

// GetRotationsOk returns a tuple with the Rotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSystemMessages) GetRotationsOk() (*int32, bool) {
	if o == nil || isNil(o.Rotations) {
    return nil, false
	}
	return o.Rotations, true
}

// HasRotations returns a boolean if a field has been set.
func (o *LoggingSystemMessages) HasRotations() bool {
	if o != nil && !isNil(o.Rotations) {
		return true
	}

	return false
}

// SetRotations gets a reference to the given int32 and assigns it to the Rotations field.
func (o *LoggingSystemMessages) SetRotations(v int32) {
	o.Rotations = &v
}

func (o LoggingSystemMessages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxFileSize) {
		toSerialize["maxFileSize"] = o.MaxFileSize
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.Rotations) {
		toSerialize["rotations"] = o.Rotations
	}
	return json.Marshal(toSerialize)
}

type NullableLoggingSystemMessages struct {
	value *LoggingSystemMessages
	isSet bool
}

func (v NullableLoggingSystemMessages) Get() *LoggingSystemMessages {
	return v.value
}

func (v *NullableLoggingSystemMessages) Set(val *LoggingSystemMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingSystemMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingSystemMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingSystemMessages(val *LoggingSystemMessages) *NullableLoggingSystemMessages {
	return &NullableLoggingSystemMessages{value: val, isSet: true}
}

func (v NullableLoggingSystemMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingSystemMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


