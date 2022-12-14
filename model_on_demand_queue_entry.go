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

// OnDemandQueueEntry struct for OnDemandQueueEntry
type OnDemandQueueEntry struct {
	Complete *bool `json:"complete,omitempty"`
	CompletionTime *int64 `json:"completionTime,omitempty"`
	CreationTime *int64 `json:"creationTime,omitempty"`
	DataType *string `json:"data_type,omitempty"`
	DeviceId *string `json:"device_id,omitempty"`
	EndTime *int64 `json:"end_time,omitempty"`
	Id *string `json:"id,omitempty"`
	StartProcessingTime *int64 `json:"startProcessingTime,omitempty"`
	StartTime *int64 `json:"start_time,omitempty"`
	Status *string `json:"status,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	TimePeriod *string `json:"time_period,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewOnDemandQueueEntry instantiates a new OnDemandQueueEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandQueueEntry() *OnDemandQueueEntry {
	this := OnDemandQueueEntry{}
	return &this
}

// NewOnDemandQueueEntryWithDefaults instantiates a new OnDemandQueueEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandQueueEntryWithDefaults() *OnDemandQueueEntry {
	this := OnDemandQueueEntry{}
	return &this
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetComplete() bool {
	if o == nil || isNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetCompleteOk() (*bool, bool) {
	if o == nil || isNil(o.Complete) {
    return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasComplete() bool {
	if o != nil && !isNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *OnDemandQueueEntry) SetComplete(v bool) {
	o.Complete = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetCompletionTime() int64 {
	if o == nil || isNil(o.CompletionTime) {
		var ret int64
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetCompletionTimeOk() (*int64, bool) {
	if o == nil || isNil(o.CompletionTime) {
    return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasCompletionTime() bool {
	if o != nil && !isNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given int64 and assigns it to the CompletionTime field.
func (o *OnDemandQueueEntry) SetCompletionTime(v int64) {
	o.CompletionTime = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetCreationTime() int64 {
	if o == nil || isNil(o.CreationTime) {
		var ret int64
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetCreationTimeOk() (*int64, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int64 and assigns it to the CreationTime field.
func (o *OnDemandQueueEntry) SetCreationTime(v int64) {
	o.CreationTime = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetDataType() string {
	if o == nil || isNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetDataTypeOk() (*string, bool) {
	if o == nil || isNil(o.DataType) {
    return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasDataType() bool {
	if o != nil && !isNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *OnDemandQueueEntry) SetDataType(v string) {
	o.DataType = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *OnDemandQueueEntry) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetEndTime() int64 {
	if o == nil || isNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetEndTimeOk() (*int64, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *OnDemandQueueEntry) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OnDemandQueueEntry) SetId(v string) {
	o.Id = &v
}

// GetStartProcessingTime returns the StartProcessingTime field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetStartProcessingTime() int64 {
	if o == nil || isNil(o.StartProcessingTime) {
		var ret int64
		return ret
	}
	return *o.StartProcessingTime
}

// GetStartProcessingTimeOk returns a tuple with the StartProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetStartProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.StartProcessingTime) {
    return nil, false
	}
	return o.StartProcessingTime, true
}

// HasStartProcessingTime returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasStartProcessingTime() bool {
	if o != nil && !isNil(o.StartProcessingTime) {
		return true
	}

	return false
}

// SetStartProcessingTime gets a reference to the given int64 and assigns it to the StartProcessingTime field.
func (o *OnDemandQueueEntry) SetStartProcessingTime(v int64) {
	o.StartProcessingTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetStartTime() int64 {
	if o == nil || isNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetStartTimeOk() (*int64, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *OnDemandQueueEntry) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OnDemandQueueEntry) SetStatus(v string) {
	o.Status = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OnDemandQueueEntry) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetTimePeriod() string {
	if o == nil || isNil(o.TimePeriod) {
		var ret string
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetTimePeriodOk() (*string, bool) {
	if o == nil || isNil(o.TimePeriod) {
    return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasTimePeriod() bool {
	if o != nil && !isNil(o.TimePeriod) {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given string and assigns it to the TimePeriod field.
func (o *OnDemandQueueEntry) SetTimePeriod(v string) {
	o.TimePeriod = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OnDemandQueueEntry) GetValue() int32 {
	if o == nil || isNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandQueueEntry) GetValueOk() (*int32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OnDemandQueueEntry) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *OnDemandQueueEntry) SetValue(v int32) {
	o.Value = &v
}

func (o OnDemandQueueEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !isNil(o.CompletionTime) {
		toSerialize["completionTime"] = o.CompletionTime
	}
	if !isNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !isNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !isNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !isNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.StartProcessingTime) {
		toSerialize["startProcessingTime"] = o.StartProcessingTime
	}
	if !isNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !isNil(o.TimePeriod) {
		toSerialize["time_period"] = o.TimePeriod
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOnDemandQueueEntry struct {
	value *OnDemandQueueEntry
	isSet bool
}

func (v NullableOnDemandQueueEntry) Get() *OnDemandQueueEntry {
	return v.value
}

func (v *NullableOnDemandQueueEntry) Set(val *OnDemandQueueEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandQueueEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandQueueEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandQueueEntry(val *OnDemandQueueEntry) *NullableOnDemandQueueEntry {
	return &NullableOnDemandQueueEntry{value: val, isSet: true}
}

func (v NullableOnDemandQueueEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandQueueEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


