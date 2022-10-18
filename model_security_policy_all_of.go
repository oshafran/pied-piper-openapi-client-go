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

// SecurityPolicyAllOf struct for SecurityPolicyAllOf
type SecurityPolicyAllOf struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	DefaultAction *string `json:"defaultAction,omitempty"`
	PolicyRules []PolicyRule `json:"policyRules,omitempty"`
}

// NewSecurityPolicyAllOf instantiates a new SecurityPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyAllOf() *SecurityPolicyAllOf {
	this := SecurityPolicyAllOf{}
	return &this
}

// NewSecurityPolicyAllOfWithDefaults instantiates a new SecurityPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyAllOfWithDefaults() *SecurityPolicyAllOf {
	this := SecurityPolicyAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityPolicyAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityPolicyAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityPolicyAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityPolicyAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityPolicyAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityPolicyAllOf) SetType(v string) {
	o.Type = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *SecurityPolicyAllOf) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyAllOf) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *SecurityPolicyAllOf) HasPolicyName() bool {
	if o != nil && o.PolicyName != nil {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *SecurityPolicyAllOf) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetDefaultAction returns the DefaultAction field value if set, zero value otherwise.
func (o *SecurityPolicyAllOf) GetDefaultAction() string {
	if o == nil || o.DefaultAction == nil {
		var ret string
		return ret
	}
	return *o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyAllOf) GetDefaultActionOk() (*string, bool) {
	if o == nil || o.DefaultAction == nil {
		return nil, false
	}
	return o.DefaultAction, true
}

// HasDefaultAction returns a boolean if a field has been set.
func (o *SecurityPolicyAllOf) HasDefaultAction() bool {
	if o != nil && o.DefaultAction != nil {
		return true
	}

	return false
}

// SetDefaultAction gets a reference to the given string and assigns it to the DefaultAction field.
func (o *SecurityPolicyAllOf) SetDefaultAction(v string) {
	o.DefaultAction = &v
}

// GetPolicyRules returns the PolicyRules field value if set, zero value otherwise.
func (o *SecurityPolicyAllOf) GetPolicyRules() []PolicyRule {
	if o == nil || o.PolicyRules == nil {
		var ret []PolicyRule
		return ret
	}
	return o.PolicyRules
}

// GetPolicyRulesOk returns a tuple with the PolicyRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPolicyAllOf) GetPolicyRulesOk() ([]PolicyRule, bool) {
	if o == nil || o.PolicyRules == nil {
		return nil, false
	}
	return o.PolicyRules, true
}

// HasPolicyRules returns a boolean if a field has been set.
func (o *SecurityPolicyAllOf) HasPolicyRules() bool {
	if o != nil && o.PolicyRules != nil {
		return true
	}

	return false
}

// SetPolicyRules gets a reference to the given []PolicyRule and assigns it to the PolicyRules field.
func (o *SecurityPolicyAllOf) SetPolicyRules(v []PolicyRule) {
	o.PolicyRules = v
}

func (o SecurityPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.PolicyName != nil {
		toSerialize["policyName"] = o.PolicyName
	}
	if o.DefaultAction != nil {
		toSerialize["defaultAction"] = o.DefaultAction
	}
	if o.PolicyRules != nil {
		toSerialize["policyRules"] = o.PolicyRules
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityPolicyAllOf struct {
	value *SecurityPolicyAllOf
	isSet bool
}

func (v NullableSecurityPolicyAllOf) Get() *SecurityPolicyAllOf {
	return v.value
}

func (v *NullableSecurityPolicyAllOf) Set(val *SecurityPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPolicyAllOf(val *SecurityPolicyAllOf) *NullableSecurityPolicyAllOf {
	return &NullableSecurityPolicyAllOf{value: val, isSet: true}
}

func (v NullableSecurityPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


