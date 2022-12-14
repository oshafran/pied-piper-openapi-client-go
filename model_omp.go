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

// OMP struct for OMP
type OMP struct {
	AdvertisementInterval *int32 `json:"advertisementInterval,omitempty"`
	Advertisements *string `json:"advertisements,omitempty"`
	EcmpLimit *int32 `json:"ecmpLimit,omitempty"`
	EorTimer *int32 `json:"eorTimer,omitempty"`
	GracefulRestart *int32 `json:"gracefulRestart,omitempty"`
	HoldTime *int32 `json:"holdTime,omitempty"`
	PathsAdvertisedPerPrefix *int32 `json:"pathsAdvertisedPerPrefix,omitempty"`
}

// NewOMP instantiates a new OMP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOMP() *OMP {
	this := OMP{}
	return &this
}

// NewOMPWithDefaults instantiates a new OMP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOMPWithDefaults() *OMP {
	this := OMP{}
	return &this
}

// GetAdvertisementInterval returns the AdvertisementInterval field value if set, zero value otherwise.
func (o *OMP) GetAdvertisementInterval() int32 {
	if o == nil || isNil(o.AdvertisementInterval) {
		var ret int32
		return ret
	}
	return *o.AdvertisementInterval
}

// GetAdvertisementIntervalOk returns a tuple with the AdvertisementInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetAdvertisementIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.AdvertisementInterval) {
    return nil, false
	}
	return o.AdvertisementInterval, true
}

// HasAdvertisementInterval returns a boolean if a field has been set.
func (o *OMP) HasAdvertisementInterval() bool {
	if o != nil && !isNil(o.AdvertisementInterval) {
		return true
	}

	return false
}

// SetAdvertisementInterval gets a reference to the given int32 and assigns it to the AdvertisementInterval field.
func (o *OMP) SetAdvertisementInterval(v int32) {
	o.AdvertisementInterval = &v
}

// GetAdvertisements returns the Advertisements field value if set, zero value otherwise.
func (o *OMP) GetAdvertisements() string {
	if o == nil || isNil(o.Advertisements) {
		var ret string
		return ret
	}
	return *o.Advertisements
}

// GetAdvertisementsOk returns a tuple with the Advertisements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetAdvertisementsOk() (*string, bool) {
	if o == nil || isNil(o.Advertisements) {
    return nil, false
	}
	return o.Advertisements, true
}

// HasAdvertisements returns a boolean if a field has been set.
func (o *OMP) HasAdvertisements() bool {
	if o != nil && !isNil(o.Advertisements) {
		return true
	}

	return false
}

// SetAdvertisements gets a reference to the given string and assigns it to the Advertisements field.
func (o *OMP) SetAdvertisements(v string) {
	o.Advertisements = &v
}

// GetEcmpLimit returns the EcmpLimit field value if set, zero value otherwise.
func (o *OMP) GetEcmpLimit() int32 {
	if o == nil || isNil(o.EcmpLimit) {
		var ret int32
		return ret
	}
	return *o.EcmpLimit
}

// GetEcmpLimitOk returns a tuple with the EcmpLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetEcmpLimitOk() (*int32, bool) {
	if o == nil || isNil(o.EcmpLimit) {
    return nil, false
	}
	return o.EcmpLimit, true
}

// HasEcmpLimit returns a boolean if a field has been set.
func (o *OMP) HasEcmpLimit() bool {
	if o != nil && !isNil(o.EcmpLimit) {
		return true
	}

	return false
}

// SetEcmpLimit gets a reference to the given int32 and assigns it to the EcmpLimit field.
func (o *OMP) SetEcmpLimit(v int32) {
	o.EcmpLimit = &v
}

// GetEorTimer returns the EorTimer field value if set, zero value otherwise.
func (o *OMP) GetEorTimer() int32 {
	if o == nil || isNil(o.EorTimer) {
		var ret int32
		return ret
	}
	return *o.EorTimer
}

// GetEorTimerOk returns a tuple with the EorTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetEorTimerOk() (*int32, bool) {
	if o == nil || isNil(o.EorTimer) {
    return nil, false
	}
	return o.EorTimer, true
}

// HasEorTimer returns a boolean if a field has been set.
func (o *OMP) HasEorTimer() bool {
	if o != nil && !isNil(o.EorTimer) {
		return true
	}

	return false
}

// SetEorTimer gets a reference to the given int32 and assigns it to the EorTimer field.
func (o *OMP) SetEorTimer(v int32) {
	o.EorTimer = &v
}

// GetGracefulRestart returns the GracefulRestart field value if set, zero value otherwise.
func (o *OMP) GetGracefulRestart() int32 {
	if o == nil || isNil(o.GracefulRestart) {
		var ret int32
		return ret
	}
	return *o.GracefulRestart
}

// GetGracefulRestartOk returns a tuple with the GracefulRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetGracefulRestartOk() (*int32, bool) {
	if o == nil || isNil(o.GracefulRestart) {
    return nil, false
	}
	return o.GracefulRestart, true
}

// HasGracefulRestart returns a boolean if a field has been set.
func (o *OMP) HasGracefulRestart() bool {
	if o != nil && !isNil(o.GracefulRestart) {
		return true
	}

	return false
}

// SetGracefulRestart gets a reference to the given int32 and assigns it to the GracefulRestart field.
func (o *OMP) SetGracefulRestart(v int32) {
	o.GracefulRestart = &v
}

// GetHoldTime returns the HoldTime field value if set, zero value otherwise.
func (o *OMP) GetHoldTime() int32 {
	if o == nil || isNil(o.HoldTime) {
		var ret int32
		return ret
	}
	return *o.HoldTime
}

// GetHoldTimeOk returns a tuple with the HoldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetHoldTimeOk() (*int32, bool) {
	if o == nil || isNil(o.HoldTime) {
    return nil, false
	}
	return o.HoldTime, true
}

// HasHoldTime returns a boolean if a field has been set.
func (o *OMP) HasHoldTime() bool {
	if o != nil && !isNil(o.HoldTime) {
		return true
	}

	return false
}

// SetHoldTime gets a reference to the given int32 and assigns it to the HoldTime field.
func (o *OMP) SetHoldTime(v int32) {
	o.HoldTime = &v
}

// GetPathsAdvertisedPerPrefix returns the PathsAdvertisedPerPrefix field value if set, zero value otherwise.
func (o *OMP) GetPathsAdvertisedPerPrefix() int32 {
	if o == nil || isNil(o.PathsAdvertisedPerPrefix) {
		var ret int32
		return ret
	}
	return *o.PathsAdvertisedPerPrefix
}

// GetPathsAdvertisedPerPrefixOk returns a tuple with the PathsAdvertisedPerPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OMP) GetPathsAdvertisedPerPrefixOk() (*int32, bool) {
	if o == nil || isNil(o.PathsAdvertisedPerPrefix) {
    return nil, false
	}
	return o.PathsAdvertisedPerPrefix, true
}

// HasPathsAdvertisedPerPrefix returns a boolean if a field has been set.
func (o *OMP) HasPathsAdvertisedPerPrefix() bool {
	if o != nil && !isNil(o.PathsAdvertisedPerPrefix) {
		return true
	}

	return false
}

// SetPathsAdvertisedPerPrefix gets a reference to the given int32 and assigns it to the PathsAdvertisedPerPrefix field.
func (o *OMP) SetPathsAdvertisedPerPrefix(v int32) {
	o.PathsAdvertisedPerPrefix = &v
}

func (o OMP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdvertisementInterval) {
		toSerialize["advertisementInterval"] = o.AdvertisementInterval
	}
	if !isNil(o.Advertisements) {
		toSerialize["advertisements"] = o.Advertisements
	}
	if !isNil(o.EcmpLimit) {
		toSerialize["ecmpLimit"] = o.EcmpLimit
	}
	if !isNil(o.EorTimer) {
		toSerialize["eorTimer"] = o.EorTimer
	}
	if !isNil(o.GracefulRestart) {
		toSerialize["gracefulRestart"] = o.GracefulRestart
	}
	if !isNil(o.HoldTime) {
		toSerialize["holdTime"] = o.HoldTime
	}
	if !isNil(o.PathsAdvertisedPerPrefix) {
		toSerialize["pathsAdvertisedPerPrefix"] = o.PathsAdvertisedPerPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableOMP struct {
	value *OMP
	isSet bool
}

func (v NullableOMP) Get() *OMP {
	return v.value
}

func (v *NullableOMP) Set(val *OMP) {
	v.value = val
	v.isSet = true
}

func (v NullableOMP) IsSet() bool {
	return v.isSet
}

func (v *NullableOMP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOMP(val *OMP) *NullableOMP {
	return &NullableOMP{value: val, isSet: true}
}

func (v NullableOMP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOMP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


