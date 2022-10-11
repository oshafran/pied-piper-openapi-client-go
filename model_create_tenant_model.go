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

// CreateTenantModel struct for CreateTenantModel
type CreateTenantModel struct {
	Desc *string `json:"desc,omitempty"`
	GetvBondAddress *string `json:"getvBondAddress,omitempty"`
	IdpMetadata *string `json:"idpMetadata,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
	OldIdpMetadata *string `json:"oldIdpMetadata,omitempty"`
	OrgName *string `json:"orgName,omitempty"`
	SpMetadata *string `json:"spMetadata,omitempty"`
	SubDomain *string `json:"subDomain,omitempty"`
	WanEdgeForecast *string `json:"wanEdgeForecast,omitempty"`
}

// NewCreateTenantModel instantiates a new CreateTenantModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenantModel() *CreateTenantModel {
	this := CreateTenantModel{}
	return &this
}

// NewCreateTenantModelWithDefaults instantiates a new CreateTenantModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenantModelWithDefaults() *CreateTenantModel {
	this := CreateTenantModel{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *CreateTenantModel) GetDesc() string {
	if o == nil || o.Desc == nil {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetDescOk() (*string, bool) {
	if o == nil || o.Desc == nil {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *CreateTenantModel) HasDesc() bool {
	if o != nil && o.Desc != nil {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *CreateTenantModel) SetDesc(v string) {
	o.Desc = &v
}

// GetGetvBondAddress returns the GetvBondAddress field value if set, zero value otherwise.
func (o *CreateTenantModel) GetGetvBondAddress() string {
	if o == nil || o.GetvBondAddress == nil {
		var ret string
		return ret
	}
	return *o.GetvBondAddress
}

// GetGetvBondAddressOk returns a tuple with the GetvBondAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetGetvBondAddressOk() (*string, bool) {
	if o == nil || o.GetvBondAddress == nil {
		return nil, false
	}
	return o.GetvBondAddress, true
}

// HasGetvBondAddress returns a boolean if a field has been set.
func (o *CreateTenantModel) HasGetvBondAddress() bool {
	if o != nil && o.GetvBondAddress != nil {
		return true
	}

	return false
}

// SetGetvBondAddress gets a reference to the given string and assigns it to the GetvBondAddress field.
func (o *CreateTenantModel) SetGetvBondAddress(v string) {
	o.GetvBondAddress = &v
}

// GetIdpMetadata returns the IdpMetadata field value if set, zero value otherwise.
func (o *CreateTenantModel) GetIdpMetadata() string {
	if o == nil || o.IdpMetadata == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadata
}

// GetIdpMetadataOk returns a tuple with the IdpMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetIdpMetadataOk() (*string, bool) {
	if o == nil || o.IdpMetadata == nil {
		return nil, false
	}
	return o.IdpMetadata, true
}

// HasIdpMetadata returns a boolean if a field has been set.
func (o *CreateTenantModel) HasIdpMetadata() bool {
	if o != nil && o.IdpMetadata != nil {
		return true
	}

	return false
}

// SetIdpMetadata gets a reference to the given string and assigns it to the IdpMetadata field.
func (o *CreateTenantModel) SetIdpMetadata(v string) {
	o.IdpMetadata = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateTenantModel) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateTenantModel) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateTenantModel) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateTenantModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateTenantModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateTenantModel) SetName(v string) {
	o.Name = &v
}

// GetOldIdpMetadata returns the OldIdpMetadata field value if set, zero value otherwise.
func (o *CreateTenantModel) GetOldIdpMetadata() string {
	if o == nil || o.OldIdpMetadata == nil {
		var ret string
		return ret
	}
	return *o.OldIdpMetadata
}

// GetOldIdpMetadataOk returns a tuple with the OldIdpMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetOldIdpMetadataOk() (*string, bool) {
	if o == nil || o.OldIdpMetadata == nil {
		return nil, false
	}
	return o.OldIdpMetadata, true
}

// HasOldIdpMetadata returns a boolean if a field has been set.
func (o *CreateTenantModel) HasOldIdpMetadata() bool {
	if o != nil && o.OldIdpMetadata != nil {
		return true
	}

	return false
}

// SetOldIdpMetadata gets a reference to the given string and assigns it to the OldIdpMetadata field.
func (o *CreateTenantModel) SetOldIdpMetadata(v string) {
	o.OldIdpMetadata = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *CreateTenantModel) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *CreateTenantModel) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *CreateTenantModel) SetOrgName(v string) {
	o.OrgName = &v
}

// GetSpMetadata returns the SpMetadata field value if set, zero value otherwise.
func (o *CreateTenantModel) GetSpMetadata() string {
	if o == nil || o.SpMetadata == nil {
		var ret string
		return ret
	}
	return *o.SpMetadata
}

// GetSpMetadataOk returns a tuple with the SpMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetSpMetadataOk() (*string, bool) {
	if o == nil || o.SpMetadata == nil {
		return nil, false
	}
	return o.SpMetadata, true
}

// HasSpMetadata returns a boolean if a field has been set.
func (o *CreateTenantModel) HasSpMetadata() bool {
	if o != nil && o.SpMetadata != nil {
		return true
	}

	return false
}

// SetSpMetadata gets a reference to the given string and assigns it to the SpMetadata field.
func (o *CreateTenantModel) SetSpMetadata(v string) {
	o.SpMetadata = &v
}

// GetSubDomain returns the SubDomain field value if set, zero value otherwise.
func (o *CreateTenantModel) GetSubDomain() string {
	if o == nil || o.SubDomain == nil {
		var ret string
		return ret
	}
	return *o.SubDomain
}

// GetSubDomainOk returns a tuple with the SubDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetSubDomainOk() (*string, bool) {
	if o == nil || o.SubDomain == nil {
		return nil, false
	}
	return o.SubDomain, true
}

// HasSubDomain returns a boolean if a field has been set.
func (o *CreateTenantModel) HasSubDomain() bool {
	if o != nil && o.SubDomain != nil {
		return true
	}

	return false
}

// SetSubDomain gets a reference to the given string and assigns it to the SubDomain field.
func (o *CreateTenantModel) SetSubDomain(v string) {
	o.SubDomain = &v
}

// GetWanEdgeForecast returns the WanEdgeForecast field value if set, zero value otherwise.
func (o *CreateTenantModel) GetWanEdgeForecast() string {
	if o == nil || o.WanEdgeForecast == nil {
		var ret string
		return ret
	}
	return *o.WanEdgeForecast
}

// GetWanEdgeForecastOk returns a tuple with the WanEdgeForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantModel) GetWanEdgeForecastOk() (*string, bool) {
	if o == nil || o.WanEdgeForecast == nil {
		return nil, false
	}
	return o.WanEdgeForecast, true
}

// HasWanEdgeForecast returns a boolean if a field has been set.
func (o *CreateTenantModel) HasWanEdgeForecast() bool {
	if o != nil && o.WanEdgeForecast != nil {
		return true
	}

	return false
}

// SetWanEdgeForecast gets a reference to the given string and assigns it to the WanEdgeForecast field.
func (o *CreateTenantModel) SetWanEdgeForecast(v string) {
	o.WanEdgeForecast = &v
}

func (o CreateTenantModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Desc != nil {
		toSerialize["desc"] = o.Desc
	}
	if o.GetvBondAddress != nil {
		toSerialize["getvBondAddress"] = o.GetvBondAddress
	}
	if o.IdpMetadata != nil {
		toSerialize["idpMetadata"] = o.IdpMetadata
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OldIdpMetadata != nil {
		toSerialize["oldIdpMetadata"] = o.OldIdpMetadata
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.SpMetadata != nil {
		toSerialize["spMetadata"] = o.SpMetadata
	}
	if o.SubDomain != nil {
		toSerialize["subDomain"] = o.SubDomain
	}
	if o.WanEdgeForecast != nil {
		toSerialize["wanEdgeForecast"] = o.WanEdgeForecast
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTenantModel struct {
	value *CreateTenantModel
	isSet bool
}

func (v NullableCreateTenantModel) Get() *CreateTenantModel {
	return v.value
}

func (v *NullableCreateTenantModel) Set(val *CreateTenantModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenantModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenantModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenantModel(val *CreateTenantModel) *NullableCreateTenantModel {
	return &NullableCreateTenantModel{value: val, isSet: true}
}

func (v NullableCreateTenantModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenantModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

