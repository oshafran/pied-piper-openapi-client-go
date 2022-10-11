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

// ThreadPoolDefinition struct for ThreadPoolDefinition
type ThreadPoolDefinition struct {
	ConsumerClass *string `json:"consumerClass,omitempty"`
	ConsumerMethod *string `json:"consumerMethod,omitempty"`
	ThreadPoolName *string `json:"threadPoolName,omitempty"`
	ThreadPoolSize *int32 `json:"threadPoolSize,omitempty"`
}

// NewThreadPoolDefinition instantiates a new ThreadPoolDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadPoolDefinition() *ThreadPoolDefinition {
	this := ThreadPoolDefinition{}
	return &this
}

// NewThreadPoolDefinitionWithDefaults instantiates a new ThreadPoolDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadPoolDefinitionWithDefaults() *ThreadPoolDefinition {
	this := ThreadPoolDefinition{}
	return &this
}

// GetConsumerClass returns the ConsumerClass field value if set, zero value otherwise.
func (o *ThreadPoolDefinition) GetConsumerClass() string {
	if o == nil || o.ConsumerClass == nil {
		var ret string
		return ret
	}
	return *o.ConsumerClass
}

// GetConsumerClassOk returns a tuple with the ConsumerClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolDefinition) GetConsumerClassOk() (*string, bool) {
	if o == nil || o.ConsumerClass == nil {
		return nil, false
	}
	return o.ConsumerClass, true
}

// HasConsumerClass returns a boolean if a field has been set.
func (o *ThreadPoolDefinition) HasConsumerClass() bool {
	if o != nil && o.ConsumerClass != nil {
		return true
	}

	return false
}

// SetConsumerClass gets a reference to the given string and assigns it to the ConsumerClass field.
func (o *ThreadPoolDefinition) SetConsumerClass(v string) {
	o.ConsumerClass = &v
}

// GetConsumerMethod returns the ConsumerMethod field value if set, zero value otherwise.
func (o *ThreadPoolDefinition) GetConsumerMethod() string {
	if o == nil || o.ConsumerMethod == nil {
		var ret string
		return ret
	}
	return *o.ConsumerMethod
}

// GetConsumerMethodOk returns a tuple with the ConsumerMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolDefinition) GetConsumerMethodOk() (*string, bool) {
	if o == nil || o.ConsumerMethod == nil {
		return nil, false
	}
	return o.ConsumerMethod, true
}

// HasConsumerMethod returns a boolean if a field has been set.
func (o *ThreadPoolDefinition) HasConsumerMethod() bool {
	if o != nil && o.ConsumerMethod != nil {
		return true
	}

	return false
}

// SetConsumerMethod gets a reference to the given string and assigns it to the ConsumerMethod field.
func (o *ThreadPoolDefinition) SetConsumerMethod(v string) {
	o.ConsumerMethod = &v
}

// GetThreadPoolName returns the ThreadPoolName field value if set, zero value otherwise.
func (o *ThreadPoolDefinition) GetThreadPoolName() string {
	if o == nil || o.ThreadPoolName == nil {
		var ret string
		return ret
	}
	return *o.ThreadPoolName
}

// GetThreadPoolNameOk returns a tuple with the ThreadPoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolDefinition) GetThreadPoolNameOk() (*string, bool) {
	if o == nil || o.ThreadPoolName == nil {
		return nil, false
	}
	return o.ThreadPoolName, true
}

// HasThreadPoolName returns a boolean if a field has been set.
func (o *ThreadPoolDefinition) HasThreadPoolName() bool {
	if o != nil && o.ThreadPoolName != nil {
		return true
	}

	return false
}

// SetThreadPoolName gets a reference to the given string and assigns it to the ThreadPoolName field.
func (o *ThreadPoolDefinition) SetThreadPoolName(v string) {
	o.ThreadPoolName = &v
}

// GetThreadPoolSize returns the ThreadPoolSize field value if set, zero value otherwise.
func (o *ThreadPoolDefinition) GetThreadPoolSize() int32 {
	if o == nil || o.ThreadPoolSize == nil {
		var ret int32
		return ret
	}
	return *o.ThreadPoolSize
}

// GetThreadPoolSizeOk returns a tuple with the ThreadPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolDefinition) GetThreadPoolSizeOk() (*int32, bool) {
	if o == nil || o.ThreadPoolSize == nil {
		return nil, false
	}
	return o.ThreadPoolSize, true
}

// HasThreadPoolSize returns a boolean if a field has been set.
func (o *ThreadPoolDefinition) HasThreadPoolSize() bool {
	if o != nil && o.ThreadPoolSize != nil {
		return true
	}

	return false
}

// SetThreadPoolSize gets a reference to the given int32 and assigns it to the ThreadPoolSize field.
func (o *ThreadPoolDefinition) SetThreadPoolSize(v int32) {
	o.ThreadPoolSize = &v
}

func (o ThreadPoolDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumerClass != nil {
		toSerialize["consumerClass"] = o.ConsumerClass
	}
	if o.ConsumerMethod != nil {
		toSerialize["consumerMethod"] = o.ConsumerMethod
	}
	if o.ThreadPoolName != nil {
		toSerialize["threadPoolName"] = o.ThreadPoolName
	}
	if o.ThreadPoolSize != nil {
		toSerialize["threadPoolSize"] = o.ThreadPoolSize
	}
	return json.Marshal(toSerialize)
}

type NullableThreadPoolDefinition struct {
	value *ThreadPoolDefinition
	isSet bool
}

func (v NullableThreadPoolDefinition) Get() *ThreadPoolDefinition {
	return v.value
}

func (v *NullableThreadPoolDefinition) Set(val *ThreadPoolDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadPoolDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadPoolDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadPoolDefinition(val *ThreadPoolDefinition) *NullableThreadPoolDefinition {
	return &NullableThreadPoolDefinition{value: val, isSet: true}
}

func (v NullableThreadPoolDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadPoolDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

