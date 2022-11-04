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

// SmartAccountModel struct for SmartAccountModel
type SmartAccountModel struct {
	Env *string `json:"env,omitempty"`
	OrganizationName *string `json:"organization_name,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
	ValidityString *string `json:"validity_string,omitempty"`
}

// NewSmartAccountModel instantiates a new SmartAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartAccountModel() *SmartAccountModel {
	this := SmartAccountModel{}
	return &this
}

// NewSmartAccountModelWithDefaults instantiates a new SmartAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartAccountModelWithDefaults() *SmartAccountModel {
	this := SmartAccountModel{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *SmartAccountModel) GetEnv() string {
	if o == nil || isNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountModel) GetEnvOk() (*string, bool) {
	if o == nil || isNil(o.Env) {
    return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *SmartAccountModel) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *SmartAccountModel) SetEnv(v string) {
	o.Env = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *SmartAccountModel) GetOrganizationName() string {
	if o == nil || isNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountModel) GetOrganizationNameOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationName) {
    return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *SmartAccountModel) HasOrganizationName() bool {
	if o != nil && !isNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *SmartAccountModel) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SmartAccountModel) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountModel) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SmartAccountModel) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SmartAccountModel) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SmartAccountModel) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountModel) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SmartAccountModel) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SmartAccountModel) SetUsername(v string) {
	o.Username = &v
}

// GetValidityString returns the ValidityString field value if set, zero value otherwise.
func (o *SmartAccountModel) GetValidityString() string {
	if o == nil || isNil(o.ValidityString) {
		var ret string
		return ret
	}
	return *o.ValidityString
}

// GetValidityStringOk returns a tuple with the ValidityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountModel) GetValidityStringOk() (*string, bool) {
	if o == nil || isNil(o.ValidityString) {
    return nil, false
	}
	return o.ValidityString, true
}

// HasValidityString returns a boolean if a field has been set.
func (o *SmartAccountModel) HasValidityString() bool {
	if o != nil && !isNil(o.ValidityString) {
		return true
	}

	return false
}

// SetValidityString gets a reference to the given string and assigns it to the ValidityString field.
func (o *SmartAccountModel) SetValidityString(v string) {
	o.ValidityString = &v
}

func (o SmartAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.OrganizationName) {
		toSerialize["organization_name"] = o.OrganizationName
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.ValidityString) {
		toSerialize["validity_string"] = o.ValidityString
	}
	return json.Marshal(toSerialize)
}

type NullableSmartAccountModel struct {
	value *SmartAccountModel
	isSet bool
}

func (v NullableSmartAccountModel) Get() *SmartAccountModel {
	return v.value
}

func (v *NullableSmartAccountModel) Set(val *SmartAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartAccountModel(val *SmartAccountModel) *NullableSmartAccountModel {
	return &NullableSmartAccountModel{value: val, isSet: true}
}

func (v NullableSmartAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


