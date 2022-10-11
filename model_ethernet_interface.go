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

// EthernetInterface struct for EthernetInterface
type EthernetInterface struct {
	AdminState *string `json:"adminState,omitempty"`
	CorporateLan *bool `json:"corporateLan,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty"`
	IpAssignment *string `json:"ipAssignment,omitempty"`
	PortType *string `json:"portType,omitempty"`
	StaticIpAddress *string `json:"staticIpAddress,omitempty"`
	StaticIpAddressSubnetMask *string `json:"staticIpAddressSubnetMask,omitempty"`
	StaticRouteIp *string `json:"staticRouteIp,omitempty"`
	WanConfiguration *string `json:"wanConfiguration,omitempty"`
}

// NewEthernetInterface instantiates a new EthernetInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthernetInterface() *EthernetInterface {
	this := EthernetInterface{}
	return &this
}

// NewEthernetInterfaceWithDefaults instantiates a new EthernetInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthernetInterfaceWithDefaults() *EthernetInterface {
	this := EthernetInterface{}
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *EthernetInterface) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *EthernetInterface) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *EthernetInterface) SetAdminState(v string) {
	o.AdminState = &v
}

// GetCorporateLan returns the CorporateLan field value if set, zero value otherwise.
func (o *EthernetInterface) GetCorporateLan() bool {
	if o == nil || o.CorporateLan == nil {
		var ret bool
		return ret
	}
	return *o.CorporateLan
}

// GetCorporateLanOk returns a tuple with the CorporateLan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetCorporateLanOk() (*bool, bool) {
	if o == nil || o.CorporateLan == nil {
		return nil, false
	}
	return o.CorporateLan, true
}

// HasCorporateLan returns a boolean if a field has been set.
func (o *EthernetInterface) HasCorporateLan() bool {
	if o != nil && o.CorporateLan != nil {
		return true
	}

	return false
}

// SetCorporateLan gets a reference to the given bool and assigns it to the CorporateLan field.
func (o *EthernetInterface) SetCorporateLan(v bool) {
	o.CorporateLan = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *EthernetInterface) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *EthernetInterface) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *EthernetInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIpAssignment returns the IpAssignment field value if set, zero value otherwise.
func (o *EthernetInterface) GetIpAssignment() string {
	if o == nil || o.IpAssignment == nil {
		var ret string
		return ret
	}
	return *o.IpAssignment
}

// GetIpAssignmentOk returns a tuple with the IpAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetIpAssignmentOk() (*string, bool) {
	if o == nil || o.IpAssignment == nil {
		return nil, false
	}
	return o.IpAssignment, true
}

// HasIpAssignment returns a boolean if a field has been set.
func (o *EthernetInterface) HasIpAssignment() bool {
	if o != nil && o.IpAssignment != nil {
		return true
	}

	return false
}

// SetIpAssignment gets a reference to the given string and assigns it to the IpAssignment field.
func (o *EthernetInterface) SetIpAssignment(v string) {
	o.IpAssignment = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *EthernetInterface) GetPortType() string {
	if o == nil || o.PortType == nil {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetPortTypeOk() (*string, bool) {
	if o == nil || o.PortType == nil {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *EthernetInterface) HasPortType() bool {
	if o != nil && o.PortType != nil {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *EthernetInterface) SetPortType(v string) {
	o.PortType = &v
}

// GetStaticIpAddress returns the StaticIpAddress field value if set, zero value otherwise.
func (o *EthernetInterface) GetStaticIpAddress() string {
	if o == nil || o.StaticIpAddress == nil {
		var ret string
		return ret
	}
	return *o.StaticIpAddress
}

// GetStaticIpAddressOk returns a tuple with the StaticIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetStaticIpAddressOk() (*string, bool) {
	if o == nil || o.StaticIpAddress == nil {
		return nil, false
	}
	return o.StaticIpAddress, true
}

// HasStaticIpAddress returns a boolean if a field has been set.
func (o *EthernetInterface) HasStaticIpAddress() bool {
	if o != nil && o.StaticIpAddress != nil {
		return true
	}

	return false
}

// SetStaticIpAddress gets a reference to the given string and assigns it to the StaticIpAddress field.
func (o *EthernetInterface) SetStaticIpAddress(v string) {
	o.StaticIpAddress = &v
}

// GetStaticIpAddressSubnetMask returns the StaticIpAddressSubnetMask field value if set, zero value otherwise.
func (o *EthernetInterface) GetStaticIpAddressSubnetMask() string {
	if o == nil || o.StaticIpAddressSubnetMask == nil {
		var ret string
		return ret
	}
	return *o.StaticIpAddressSubnetMask
}

// GetStaticIpAddressSubnetMaskOk returns a tuple with the StaticIpAddressSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetStaticIpAddressSubnetMaskOk() (*string, bool) {
	if o == nil || o.StaticIpAddressSubnetMask == nil {
		return nil, false
	}
	return o.StaticIpAddressSubnetMask, true
}

// HasStaticIpAddressSubnetMask returns a boolean if a field has been set.
func (o *EthernetInterface) HasStaticIpAddressSubnetMask() bool {
	if o != nil && o.StaticIpAddressSubnetMask != nil {
		return true
	}

	return false
}

// SetStaticIpAddressSubnetMask gets a reference to the given string and assigns it to the StaticIpAddressSubnetMask field.
func (o *EthernetInterface) SetStaticIpAddressSubnetMask(v string) {
	o.StaticIpAddressSubnetMask = &v
}

// GetStaticRouteIp returns the StaticRouteIp field value if set, zero value otherwise.
func (o *EthernetInterface) GetStaticRouteIp() string {
	if o == nil || o.StaticRouteIp == nil {
		var ret string
		return ret
	}
	return *o.StaticRouteIp
}

// GetStaticRouteIpOk returns a tuple with the StaticRouteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetStaticRouteIpOk() (*string, bool) {
	if o == nil || o.StaticRouteIp == nil {
		return nil, false
	}
	return o.StaticRouteIp, true
}

// HasStaticRouteIp returns a boolean if a field has been set.
func (o *EthernetInterface) HasStaticRouteIp() bool {
	if o != nil && o.StaticRouteIp != nil {
		return true
	}

	return false
}

// SetStaticRouteIp gets a reference to the given string and assigns it to the StaticRouteIp field.
func (o *EthernetInterface) SetStaticRouteIp(v string) {
	o.StaticRouteIp = &v
}

// GetWanConfiguration returns the WanConfiguration field value if set, zero value otherwise.
func (o *EthernetInterface) GetWanConfiguration() string {
	if o == nil || o.WanConfiguration == nil {
		var ret string
		return ret
	}
	return *o.WanConfiguration
}

// GetWanConfigurationOk returns a tuple with the WanConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterface) GetWanConfigurationOk() (*string, bool) {
	if o == nil || o.WanConfiguration == nil {
		return nil, false
	}
	return o.WanConfiguration, true
}

// HasWanConfiguration returns a boolean if a field has been set.
func (o *EthernetInterface) HasWanConfiguration() bool {
	if o != nil && o.WanConfiguration != nil {
		return true
	}

	return false
}

// SetWanConfiguration gets a reference to the given string and assigns it to the WanConfiguration field.
func (o *EthernetInterface) SetWanConfiguration(v string) {
	o.WanConfiguration = &v
}

func (o EthernetInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminState != nil {
		toSerialize["adminState"] = o.AdminState
	}
	if o.CorporateLan != nil {
		toSerialize["corporateLan"] = o.CorporateLan
	}
	if o.InterfaceName != nil {
		toSerialize["interfaceName"] = o.InterfaceName
	}
	if o.IpAssignment != nil {
		toSerialize["ipAssignment"] = o.IpAssignment
	}
	if o.PortType != nil {
		toSerialize["portType"] = o.PortType
	}
	if o.StaticIpAddress != nil {
		toSerialize["staticIpAddress"] = o.StaticIpAddress
	}
	if o.StaticIpAddressSubnetMask != nil {
		toSerialize["staticIpAddressSubnetMask"] = o.StaticIpAddressSubnetMask
	}
	if o.StaticRouteIp != nil {
		toSerialize["staticRouteIp"] = o.StaticRouteIp
	}
	if o.WanConfiguration != nil {
		toSerialize["wanConfiguration"] = o.WanConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableEthernetInterface struct {
	value *EthernetInterface
	isSet bool
}

func (v NullableEthernetInterface) Get() *EthernetInterface {
	return v.value
}

func (v *NullableEthernetInterface) Set(val *EthernetInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableEthernetInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableEthernetInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthernetInterface(val *EthernetInterface) *NullableEthernetInterface {
	return &NullableEthernetInterface{value: val, isSet: true}
}

func (v NullableEthernetInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthernetInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

