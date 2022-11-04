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

// DataCenter struct for DataCenter
type DataCenter struct {
	DcPersonality *string `json:"dcPersonality,omitempty"`
	Members []Node `json:"members,omitempty"`
	MgmtIPAddress *string `json:"mgmtIPAddress,omitempty"`
	Name *string `json:"name,omitempty"`
	NmsPersonality *string `json:"nmsPersonality,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewDataCenter instantiates a new DataCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataCenter() *DataCenter {
	this := DataCenter{}
	return &this
}

// NewDataCenterWithDefaults instantiates a new DataCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataCenterWithDefaults() *DataCenter {
	this := DataCenter{}
	return &this
}

// GetDcPersonality returns the DcPersonality field value if set, zero value otherwise.
func (o *DataCenter) GetDcPersonality() string {
	if o == nil || isNil(o.DcPersonality) {
		var ret string
		return ret
	}
	return *o.DcPersonality
}

// GetDcPersonalityOk returns a tuple with the DcPersonality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetDcPersonalityOk() (*string, bool) {
	if o == nil || isNil(o.DcPersonality) {
    return nil, false
	}
	return o.DcPersonality, true
}

// HasDcPersonality returns a boolean if a field has been set.
func (o *DataCenter) HasDcPersonality() bool {
	if o != nil && !isNil(o.DcPersonality) {
		return true
	}

	return false
}

// SetDcPersonality gets a reference to the given string and assigns it to the DcPersonality field.
func (o *DataCenter) SetDcPersonality(v string) {
	o.DcPersonality = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DataCenter) GetMembers() []Node {
	if o == nil || isNil(o.Members) {
		var ret []Node
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetMembersOk() ([]Node, bool) {
	if o == nil || isNil(o.Members) {
    return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DataCenter) HasMembers() bool {
	if o != nil && !isNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Node and assigns it to the Members field.
func (o *DataCenter) SetMembers(v []Node) {
	o.Members = v
}

// GetMgmtIPAddress returns the MgmtIPAddress field value if set, zero value otherwise.
func (o *DataCenter) GetMgmtIPAddress() string {
	if o == nil || isNil(o.MgmtIPAddress) {
		var ret string
		return ret
	}
	return *o.MgmtIPAddress
}

// GetMgmtIPAddressOk returns a tuple with the MgmtIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetMgmtIPAddressOk() (*string, bool) {
	if o == nil || isNil(o.MgmtIPAddress) {
    return nil, false
	}
	return o.MgmtIPAddress, true
}

// HasMgmtIPAddress returns a boolean if a field has been set.
func (o *DataCenter) HasMgmtIPAddress() bool {
	if o != nil && !isNil(o.MgmtIPAddress) {
		return true
	}

	return false
}

// SetMgmtIPAddress gets a reference to the given string and assigns it to the MgmtIPAddress field.
func (o *DataCenter) SetMgmtIPAddress(v string) {
	o.MgmtIPAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataCenter) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataCenter) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataCenter) SetName(v string) {
	o.Name = &v
}

// GetNmsPersonality returns the NmsPersonality field value if set, zero value otherwise.
func (o *DataCenter) GetNmsPersonality() string {
	if o == nil || isNil(o.NmsPersonality) {
		var ret string
		return ret
	}
	return *o.NmsPersonality
}

// GetNmsPersonalityOk returns a tuple with the NmsPersonality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetNmsPersonalityOk() (*string, bool) {
	if o == nil || isNil(o.NmsPersonality) {
    return nil, false
	}
	return o.NmsPersonality, true
}

// HasNmsPersonality returns a boolean if a field has been set.
func (o *DataCenter) HasNmsPersonality() bool {
	if o != nil && !isNil(o.NmsPersonality) {
		return true
	}

	return false
}

// SetNmsPersonality gets a reference to the given string and assigns it to the NmsPersonality field.
func (o *DataCenter) SetNmsPersonality(v string) {
	o.NmsPersonality = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DataCenter) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DataCenter) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DataCenter) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DataCenter) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DataCenter) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DataCenter) SetUsername(v string) {
	o.Username = &v
}

func (o DataCenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DcPersonality) {
		toSerialize["dcPersonality"] = o.DcPersonality
	}
	if !isNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !isNil(o.MgmtIPAddress) {
		toSerialize["mgmtIPAddress"] = o.MgmtIPAddress
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NmsPersonality) {
		toSerialize["nmsPersonality"] = o.NmsPersonality
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableDataCenter struct {
	value *DataCenter
	isSet bool
}

func (v NullableDataCenter) Get() *DataCenter {
	return v.value
}

func (v *NullableDataCenter) Set(val *DataCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCenter(val *DataCenter) *NullableDataCenter {
	return &NullableDataCenter{value: val, isSet: true}
}

func (v NullableDataCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


