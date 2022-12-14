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

// FeatureProfile List of devices UUIDs associated with this config group
type FeatureProfile struct {
	// User who last created this.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Timestamp of creation
	CreatedOn *int64 `json:"createdOn,omitempty"`
	// Description of the feature Profile.
	Description *string `json:"description,omitempty"`
	// System generated unique identifier of the feature profile in UUID format.
	Id *string `json:"id,omitempty"`
	// User who last updated this.
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// Timestamp of last update
	LastUpdatedOn *int64 `json:"lastUpdatedOn,omitempty"`
	// Name of the feature Profile. Must be unique.
	Name string `json:"name"`
	// Solution of the feature Profile.
	Solution string `json:"solution"`
	// Type of the feature Profile.
	Type string `json:"type"`
}

// NewFeatureProfile instantiates a new FeatureProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureProfile(name string, solution string, type_ string) *FeatureProfile {
	this := FeatureProfile{}
	this.Name = name
	this.Solution = solution
	this.Type = type_
	return &this
}

// NewFeatureProfileWithDefaults instantiates a new FeatureProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureProfileWithDefaults() *FeatureProfile {
	this := FeatureProfile{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FeatureProfile) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetCreatedByOk() (*string, bool) {
	if o == nil || isNil(o.CreatedBy) {
    return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FeatureProfile) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *FeatureProfile) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *FeatureProfile) GetCreatedOn() int64 {
	if o == nil || isNil(o.CreatedOn) {
		var ret int64
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetCreatedOnOk() (*int64, bool) {
	if o == nil || isNil(o.CreatedOn) {
    return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *FeatureProfile) HasCreatedOn() bool {
	if o != nil && !isNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given int64 and assigns it to the CreatedOn field.
func (o *FeatureProfile) SetCreatedOn(v int64) {
	o.CreatedOn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureProfile) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureProfile) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureProfile) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeatureProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeatureProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeatureProfile) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *FeatureProfile) GetLastUpdatedBy() string {
	if o == nil || isNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedBy) {
    return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *FeatureProfile) HasLastUpdatedBy() bool {
	if o != nil && !isNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *FeatureProfile) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetLastUpdatedOn returns the LastUpdatedOn field value if set, zero value otherwise.
func (o *FeatureProfile) GetLastUpdatedOn() int64 {
	if o == nil || isNil(o.LastUpdatedOn) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedOn
}

// GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetLastUpdatedOnOk() (*int64, bool) {
	if o == nil || isNil(o.LastUpdatedOn) {
    return nil, false
	}
	return o.LastUpdatedOn, true
}

// HasLastUpdatedOn returns a boolean if a field has been set.
func (o *FeatureProfile) HasLastUpdatedOn() bool {
	if o != nil && !isNil(o.LastUpdatedOn) {
		return true
	}

	return false
}

// SetLastUpdatedOn gets a reference to the given int64 and assigns it to the LastUpdatedOn field.
func (o *FeatureProfile) SetLastUpdatedOn(v int64) {
	o.LastUpdatedOn = &v
}

// GetName returns the Name field value
func (o *FeatureProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FeatureProfile) SetName(v string) {
	o.Name = v
}

// GetSolution returns the Solution field value
func (o *FeatureProfile) GetSolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetSolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Solution, true
}

// SetSolution sets field value
func (o *FeatureProfile) SetSolution(v string) {
	o.Solution = v
}

// GetType returns the Type field value
func (o *FeatureProfile) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FeatureProfile) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FeatureProfile) SetType(v string) {
	o.Type = v
}

func (o FeatureProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !isNil(o.LastUpdatedOn) {
		toSerialize["lastUpdatedOn"] = o.LastUpdatedOn
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["solution"] = o.Solution
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureProfile struct {
	value *FeatureProfile
	isSet bool
}

func (v NullableFeatureProfile) Get() *FeatureProfile {
	return v.value
}

func (v *NullableFeatureProfile) Set(val *FeatureProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureProfile(val *FeatureProfile) *NullableFeatureProfile {
	return &NullableFeatureProfile{value: val, isSet: true}
}

func (v NullableFeatureProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


