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

// GlobalSettings struct for GlobalSettings
type GlobalSettings struct {
	BasicName *string `json:"basicName,omitempty"`
	BasicDescription *string `json:"basicDescription,omitempty"`
	NtpServer []ConnectToNtpServer `json:"ntpServer,omitempty"`
	Systems *Systems `json:"systems,omitempty"`
	Banner *Banner `json:"banner,omitempty"`
	LoginAccessToRouter *LoginAccessToRouter `json:"loginAccessToRouter,omitempty"`
	Bfd *Bfd `json:"bfd,omitempty"`
	Omp *OMP `json:"omp,omitempty"`
	IpSecSecurity *IpSecSecurity `json:"ipSecSecurity,omitempty"`
	LoggingSystemMessages *LoggingSystemMessages `json:"loggingSystemMessages,omitempty"`
}

// NewGlobalSettings instantiates a new GlobalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSettings(name string, type_ string) *GlobalSettings {
	this := GlobalSettings{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewGlobalSettingsWithDefaults instantiates a new GlobalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSettingsWithDefaults() *GlobalSettings {
	this := GlobalSettings{}
	return &this
}

// GetBasicName returns the BasicName field value if set, zero value otherwise.
func (o *GlobalSettings) GetBasicName() string {
	if o == nil || o.BasicName == nil {
		var ret string
		return ret
	}
	return *o.BasicName
}

// GetBasicNameOk returns a tuple with the BasicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBasicNameOk() (*string, bool) {
	if o == nil || o.BasicName == nil {
		return nil, false
	}
	return o.BasicName, true
}

// HasBasicName returns a boolean if a field has been set.
func (o *GlobalSettings) HasBasicName() bool {
	if o != nil && o.BasicName != nil {
		return true
	}

	return false
}

// SetBasicName gets a reference to the given string and assigns it to the BasicName field.
func (o *GlobalSettings) SetBasicName(v string) {
	o.BasicName = &v
}

// GetBasicDescription returns the BasicDescription field value if set, zero value otherwise.
func (o *GlobalSettings) GetBasicDescription() string {
	if o == nil || o.BasicDescription == nil {
		var ret string
		return ret
	}
	return *o.BasicDescription
}

// GetBasicDescriptionOk returns a tuple with the BasicDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBasicDescriptionOk() (*string, bool) {
	if o == nil || o.BasicDescription == nil {
		return nil, false
	}
	return o.BasicDescription, true
}

// HasBasicDescription returns a boolean if a field has been set.
func (o *GlobalSettings) HasBasicDescription() bool {
	if o != nil && o.BasicDescription != nil {
		return true
	}

	return false
}

// SetBasicDescription gets a reference to the given string and assigns it to the BasicDescription field.
func (o *GlobalSettings) SetBasicDescription(v string) {
	o.BasicDescription = &v
}

// GetNtpServer returns the NtpServer field value if set, zero value otherwise.
func (o *GlobalSettings) GetNtpServer() []ConnectToNtpServer {
	if o == nil || o.NtpServer == nil {
		var ret []ConnectToNtpServer
		return ret
	}
	return o.NtpServer
}

// GetNtpServerOk returns a tuple with the NtpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetNtpServerOk() ([]ConnectToNtpServer, bool) {
	if o == nil || o.NtpServer == nil {
		return nil, false
	}
	return o.NtpServer, true
}

// HasNtpServer returns a boolean if a field has been set.
func (o *GlobalSettings) HasNtpServer() bool {
	if o != nil && o.NtpServer != nil {
		return true
	}

	return false
}

// SetNtpServer gets a reference to the given []ConnectToNtpServer and assigns it to the NtpServer field.
func (o *GlobalSettings) SetNtpServer(v []ConnectToNtpServer) {
	o.NtpServer = v
}

// GetSystems returns the Systems field value if set, zero value otherwise.
func (o *GlobalSettings) GetSystems() Systems {
	if o == nil || o.Systems == nil {
		var ret Systems
		return ret
	}
	return *o.Systems
}

// GetSystemsOk returns a tuple with the Systems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetSystemsOk() (*Systems, bool) {
	if o == nil || o.Systems == nil {
		return nil, false
	}
	return o.Systems, true
}

// HasSystems returns a boolean if a field has been set.
func (o *GlobalSettings) HasSystems() bool {
	if o != nil && o.Systems != nil {
		return true
	}

	return false
}

// SetSystems gets a reference to the given Systems and assigns it to the Systems field.
func (o *GlobalSettings) SetSystems(v Systems) {
	o.Systems = &v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *GlobalSettings) GetBanner() Banner {
	if o == nil || o.Banner == nil {
		var ret Banner
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBannerOk() (*Banner, bool) {
	if o == nil || o.Banner == nil {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *GlobalSettings) HasBanner() bool {
	if o != nil && o.Banner != nil {
		return true
	}

	return false
}

// SetBanner gets a reference to the given Banner and assigns it to the Banner field.
func (o *GlobalSettings) SetBanner(v Banner) {
	o.Banner = &v
}

// GetLoginAccessToRouter returns the LoginAccessToRouter field value if set, zero value otherwise.
func (o *GlobalSettings) GetLoginAccessToRouter() LoginAccessToRouter {
	if o == nil || o.LoginAccessToRouter == nil {
		var ret LoginAccessToRouter
		return ret
	}
	return *o.LoginAccessToRouter
}

// GetLoginAccessToRouterOk returns a tuple with the LoginAccessToRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetLoginAccessToRouterOk() (*LoginAccessToRouter, bool) {
	if o == nil || o.LoginAccessToRouter == nil {
		return nil, false
	}
	return o.LoginAccessToRouter, true
}

// HasLoginAccessToRouter returns a boolean if a field has been set.
func (o *GlobalSettings) HasLoginAccessToRouter() bool {
	if o != nil && o.LoginAccessToRouter != nil {
		return true
	}

	return false
}

// SetLoginAccessToRouter gets a reference to the given LoginAccessToRouter and assigns it to the LoginAccessToRouter field.
func (o *GlobalSettings) SetLoginAccessToRouter(v LoginAccessToRouter) {
	o.LoginAccessToRouter = &v
}

// GetBfd returns the Bfd field value if set, zero value otherwise.
func (o *GlobalSettings) GetBfd() Bfd {
	if o == nil || o.Bfd == nil {
		var ret Bfd
		return ret
	}
	return *o.Bfd
}

// GetBfdOk returns a tuple with the Bfd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBfdOk() (*Bfd, bool) {
	if o == nil || o.Bfd == nil {
		return nil, false
	}
	return o.Bfd, true
}

// HasBfd returns a boolean if a field has been set.
func (o *GlobalSettings) HasBfd() bool {
	if o != nil && o.Bfd != nil {
		return true
	}

	return false
}

// SetBfd gets a reference to the given Bfd and assigns it to the Bfd field.
func (o *GlobalSettings) SetBfd(v Bfd) {
	o.Bfd = &v
}

// GetOmp returns the Omp field value if set, zero value otherwise.
func (o *GlobalSettings) GetOmp() OMP {
	if o == nil || o.Omp == nil {
		var ret OMP
		return ret
	}
	return *o.Omp
}

// GetOmpOk returns a tuple with the Omp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetOmpOk() (*OMP, bool) {
	if o == nil || o.Omp == nil {
		return nil, false
	}
	return o.Omp, true
}

// HasOmp returns a boolean if a field has been set.
func (o *GlobalSettings) HasOmp() bool {
	if o != nil && o.Omp != nil {
		return true
	}

	return false
}

// SetOmp gets a reference to the given OMP and assigns it to the Omp field.
func (o *GlobalSettings) SetOmp(v OMP) {
	o.Omp = &v
}

// GetIpSecSecurity returns the IpSecSecurity field value if set, zero value otherwise.
func (o *GlobalSettings) GetIpSecSecurity() IpSecSecurity {
	if o == nil || o.IpSecSecurity == nil {
		var ret IpSecSecurity
		return ret
	}
	return *o.IpSecSecurity
}

// GetIpSecSecurityOk returns a tuple with the IpSecSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetIpSecSecurityOk() (*IpSecSecurity, bool) {
	if o == nil || o.IpSecSecurity == nil {
		return nil, false
	}
	return o.IpSecSecurity, true
}

// HasIpSecSecurity returns a boolean if a field has been set.
func (o *GlobalSettings) HasIpSecSecurity() bool {
	if o != nil && o.IpSecSecurity != nil {
		return true
	}

	return false
}

// SetIpSecSecurity gets a reference to the given IpSecSecurity and assigns it to the IpSecSecurity field.
func (o *GlobalSettings) SetIpSecSecurity(v IpSecSecurity) {
	o.IpSecSecurity = &v
}

// GetLoggingSystemMessages returns the LoggingSystemMessages field value if set, zero value otherwise.
func (o *GlobalSettings) GetLoggingSystemMessages() LoggingSystemMessages {
	if o == nil || o.LoggingSystemMessages == nil {
		var ret LoggingSystemMessages
		return ret
	}
	return *o.LoggingSystemMessages
}

// GetLoggingSystemMessagesOk returns a tuple with the LoggingSystemMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetLoggingSystemMessagesOk() (*LoggingSystemMessages, bool) {
	if o == nil || o.LoggingSystemMessages == nil {
		return nil, false
	}
	return o.LoggingSystemMessages, true
}

// HasLoggingSystemMessages returns a boolean if a field has been set.
func (o *GlobalSettings) HasLoggingSystemMessages() bool {
	if o != nil && o.LoggingSystemMessages != nil {
		return true
	}

	return false
}

// SetLoggingSystemMessages gets a reference to the given LoggingSystemMessages and assigns it to the LoggingSystemMessages field.
func (o *GlobalSettings) SetLoggingSystemMessages(v LoggingSystemMessages) {
	o.LoggingSystemMessages = &v
}

func (o GlobalSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BasicName != nil {
		toSerialize["basicName"] = o.BasicName
	}
	if o.BasicDescription != nil {
		toSerialize["basicDescription"] = o.BasicDescription
	}
	if o.NtpServer != nil {
		toSerialize["ntpServer"] = o.NtpServer
	}
	if o.Systems != nil {
		toSerialize["systems"] = o.Systems
	}
	if o.Banner != nil {
		toSerialize["banner"] = o.Banner
	}
	if o.LoginAccessToRouter != nil {
		toSerialize["loginAccessToRouter"] = o.LoginAccessToRouter
	}
	if o.Bfd != nil {
		toSerialize["bfd"] = o.Bfd
	}
	if o.Omp != nil {
		toSerialize["omp"] = o.Omp
	}
	if o.IpSecSecurity != nil {
		toSerialize["ipSecSecurity"] = o.IpSecSecurity
	}
	if o.LoggingSystemMessages != nil {
		toSerialize["loggingSystemMessages"] = o.LoggingSystemMessages
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalSettings struct {
	value *GlobalSettings
	isSet bool
}

func (v NullableGlobalSettings) Get() *GlobalSettings {
	return v.value
}

func (v *NullableGlobalSettings) Set(val *GlobalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSettings(val *GlobalSettings) *NullableGlobalSettings {
	return &NullableGlobalSettings{value: val, isSet: true}
}

func (v NullableGlobalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

