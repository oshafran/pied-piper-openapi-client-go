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

// CertificateValidity This is Certificate Validity 
type CertificateValidity struct {
	CertificateValidity *string `json:"certificateValidity,omitempty"`
}

// NewCertificateValidity instantiates a new CertificateValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateValidity() *CertificateValidity {
	this := CertificateValidity{}
	return &this
}

// NewCertificateValidityWithDefaults instantiates a new CertificateValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateValidityWithDefaults() *CertificateValidity {
	this := CertificateValidity{}
	return &this
}

// GetCertificateValidity returns the CertificateValidity field value if set, zero value otherwise.
func (o *CertificateValidity) GetCertificateValidity() string {
	if o == nil || o.CertificateValidity == nil {
		var ret string
		return ret
	}
	return *o.CertificateValidity
}

// GetCertificateValidityOk returns a tuple with the CertificateValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateValidity) GetCertificateValidityOk() (*string, bool) {
	if o == nil || o.CertificateValidity == nil {
		return nil, false
	}
	return o.CertificateValidity, true
}

// HasCertificateValidity returns a boolean if a field has been set.
func (o *CertificateValidity) HasCertificateValidity() bool {
	if o != nil && o.CertificateValidity != nil {
		return true
	}

	return false
}

// SetCertificateValidity gets a reference to the given string and assigns it to the CertificateValidity field.
func (o *CertificateValidity) SetCertificateValidity(v string) {
	o.CertificateValidity = &v
}

func (o CertificateValidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateValidity != nil {
		toSerialize["certificateValidity"] = o.CertificateValidity
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateValidity struct {
	value *CertificateValidity
	isSet bool
}

func (v NullableCertificateValidity) Get() *CertificateValidity {
	return v.value
}

func (v *NullableCertificateValidity) Set(val *CertificateValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateValidity(val *CertificateValidity) *NullableCertificateValidity {
	return &NullableCertificateValidity{value: val, isSet: true}
}

func (v NullableCertificateValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

