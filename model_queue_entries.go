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

// QueueEntries struct for QueueEntries
type QueueEntries struct {
	Entries []OnDemandQueueEntry `json:"entries,omitempty"`
}

// NewQueueEntries instantiates a new QueueEntries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueEntries() *QueueEntries {
	this := QueueEntries{}
	return &this
}

// NewQueueEntriesWithDefaults instantiates a new QueueEntries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueEntriesWithDefaults() *QueueEntries {
	this := QueueEntries{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *QueueEntries) GetEntries() []OnDemandQueueEntry {
	if o == nil || isNil(o.Entries) {
		var ret []OnDemandQueueEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueEntries) GetEntriesOk() ([]OnDemandQueueEntry, bool) {
	if o == nil || isNil(o.Entries) {
    return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *QueueEntries) HasEntries() bool {
	if o != nil && !isNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []OnDemandQueueEntry and assigns it to the Entries field.
func (o *QueueEntries) SetEntries(v []OnDemandQueueEntry) {
	o.Entries = v
}

func (o QueueEntries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableQueueEntries struct {
	value *QueueEntries
	isSet bool
}

func (v NullableQueueEntries) Get() *QueueEntries {
	return v.value
}

func (v *NullableQueueEntries) Set(val *QueueEntries) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueEntries) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueEntries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueEntries(val *QueueEntries) *NullableQueueEntries {
	return &NullableQueueEntries{value: val, isSet: true}
}

func (v NullableQueueEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


