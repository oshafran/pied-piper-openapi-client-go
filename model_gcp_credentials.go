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

// GcpCredentials struct for GcpCredentials
type GcpCredentials struct {
	ClientEmail *string `json:"client_email,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	CloudTenantId *string `json:"cloudTenantId,omitempty"`
	Name *string `json:"name,omitempty"`
	PrivateKeyData *string `json:"privateKeyData,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
	PrivateKeyId *string `json:"private_key_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
}

// NewGcpCredentials instantiates a new GcpCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpCredentials() *GcpCredentials {
	this := GcpCredentials{}
	return &this
}

// NewGcpCredentialsWithDefaults instantiates a new GcpCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpCredentialsWithDefaults() *GcpCredentials {
	this := GcpCredentials{}
	return &this
}

// GetClientEmail returns the ClientEmail field value if set, zero value otherwise.
func (o *GcpCredentials) GetClientEmail() string {
	if o == nil || o.ClientEmail == nil {
		var ret string
		return ret
	}
	return *o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetClientEmailOk() (*string, bool) {
	if o == nil || o.ClientEmail == nil {
		return nil, false
	}
	return o.ClientEmail, true
}

// HasClientEmail returns a boolean if a field has been set.
func (o *GcpCredentials) HasClientEmail() bool {
	if o != nil && o.ClientEmail != nil {
		return true
	}

	return false
}

// SetClientEmail gets a reference to the given string and assigns it to the ClientEmail field.
func (o *GcpCredentials) SetClientEmail(v string) {
	o.ClientEmail = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GcpCredentials) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GcpCredentials) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GcpCredentials) SetClientId(v string) {
	o.ClientId = &v
}

// GetCloudTenantId returns the CloudTenantId field value if set, zero value otherwise.
func (o *GcpCredentials) GetCloudTenantId() string {
	if o == nil || o.CloudTenantId == nil {
		var ret string
		return ret
	}
	return *o.CloudTenantId
}

// GetCloudTenantIdOk returns a tuple with the CloudTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetCloudTenantIdOk() (*string, bool) {
	if o == nil || o.CloudTenantId == nil {
		return nil, false
	}
	return o.CloudTenantId, true
}

// HasCloudTenantId returns a boolean if a field has been set.
func (o *GcpCredentials) HasCloudTenantId() bool {
	if o != nil && o.CloudTenantId != nil {
		return true
	}

	return false
}

// SetCloudTenantId gets a reference to the given string and assigns it to the CloudTenantId field.
func (o *GcpCredentials) SetCloudTenantId(v string) {
	o.CloudTenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GcpCredentials) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GcpCredentials) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GcpCredentials) SetName(v string) {
	o.Name = &v
}

// GetPrivateKeyData returns the PrivateKeyData field value if set, zero value otherwise.
func (o *GcpCredentials) GetPrivateKeyData() string {
	if o == nil || o.PrivateKeyData == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyData
}

// GetPrivateKeyDataOk returns a tuple with the PrivateKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetPrivateKeyDataOk() (*string, bool) {
	if o == nil || o.PrivateKeyData == nil {
		return nil, false
	}
	return o.PrivateKeyData, true
}

// HasPrivateKeyData returns a boolean if a field has been set.
func (o *GcpCredentials) HasPrivateKeyData() bool {
	if o != nil && o.PrivateKeyData != nil {
		return true
	}

	return false
}

// SetPrivateKeyData gets a reference to the given string and assigns it to the PrivateKeyData field.
func (o *GcpCredentials) SetPrivateKeyData(v string) {
	o.PrivateKeyData = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GcpCredentials) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GcpCredentials) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *GcpCredentials) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyId returns the PrivateKeyId field value if set, zero value otherwise.
func (o *GcpCredentials) GetPrivateKeyId() string {
	if o == nil || o.PrivateKeyId == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyId
}

// GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetPrivateKeyIdOk() (*string, bool) {
	if o == nil || o.PrivateKeyId == nil {
		return nil, false
	}
	return o.PrivateKeyId, true
}

// HasPrivateKeyId returns a boolean if a field has been set.
func (o *GcpCredentials) HasPrivateKeyId() bool {
	if o != nil && o.PrivateKeyId != nil {
		return true
	}

	return false
}

// SetPrivateKeyId gets a reference to the given string and assigns it to the PrivateKeyId field.
func (o *GcpCredentials) SetPrivateKeyId(v string) {
	o.PrivateKeyId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GcpCredentials) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpCredentials) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GcpCredentials) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GcpCredentials) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o GcpCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientEmail != nil {
		toSerialize["client_email"] = o.ClientEmail
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CloudTenantId != nil {
		toSerialize["cloudTenantId"] = o.CloudTenantId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PrivateKeyData != nil {
		toSerialize["privateKeyData"] = o.PrivateKeyData
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.PrivateKeyId != nil {
		toSerialize["private_key_id"] = o.PrivateKeyId
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableGcpCredentials struct {
	value *GcpCredentials
	isSet bool
}

func (v NullableGcpCredentials) Get() *GcpCredentials {
	return v.value
}

func (v *NullableGcpCredentials) Set(val *GcpCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpCredentials(val *GcpCredentials) *NullableGcpCredentials {
	return &NullableGcpCredentials{value: val, isSet: true}
}

func (v NullableGcpCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

