# NetworkProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DHCPPool** | Pointer to [**DHCPPool**](DHCPPool.md) |  | [optional] 
**DNSSettings** | Pointer to **string** |  | [optional] 
**NTPSettings** | Pointer to **[]string** |  | [optional] 
**NATRules** | Pointer to [**[]NATRule**](NATRule.md) |  | [optional] 
**NTPInherit** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkProtocol

`func NewNetworkProtocol() *NetworkProtocol`

NewNetworkProtocol instantiates a new NetworkProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtocolWithDefaults

`func NewNetworkProtocolWithDefaults() *NetworkProtocol`

NewNetworkProtocolWithDefaults instantiates a new NetworkProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDHCPPool

`func (o *NetworkProtocol) GetDHCPPool() DHCPPool`

GetDHCPPool returns the DHCPPool field if non-nil, zero value otherwise.

### GetDHCPPoolOk

`func (o *NetworkProtocol) GetDHCPPoolOk() (*DHCPPool, bool)`

GetDHCPPoolOk returns a tuple with the DHCPPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDHCPPool

`func (o *NetworkProtocol) SetDHCPPool(v DHCPPool)`

SetDHCPPool sets DHCPPool field to given value.

### HasDHCPPool

`func (o *NetworkProtocol) HasDHCPPool() bool`

HasDHCPPool returns a boolean if a field has been set.

### GetDNSSettings

`func (o *NetworkProtocol) GetDNSSettings() string`

GetDNSSettings returns the DNSSettings field if non-nil, zero value otherwise.

### GetDNSSettingsOk

`func (o *NetworkProtocol) GetDNSSettingsOk() (*string, bool)`

GetDNSSettingsOk returns a tuple with the DNSSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSSettings

`func (o *NetworkProtocol) SetDNSSettings(v string)`

SetDNSSettings sets DNSSettings field to given value.

### HasDNSSettings

`func (o *NetworkProtocol) HasDNSSettings() bool`

HasDNSSettings returns a boolean if a field has been set.

### GetNTPSettings

`func (o *NetworkProtocol) GetNTPSettings() []string`

GetNTPSettings returns the NTPSettings field if non-nil, zero value otherwise.

### GetNTPSettingsOk

`func (o *NetworkProtocol) GetNTPSettingsOk() (*[]string, bool)`

GetNTPSettingsOk returns a tuple with the NTPSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPSettings

`func (o *NetworkProtocol) SetNTPSettings(v []string)`

SetNTPSettings sets NTPSettings field to given value.

### HasNTPSettings

`func (o *NetworkProtocol) HasNTPSettings() bool`

HasNTPSettings returns a boolean if a field has been set.

### GetNATRules

`func (o *NetworkProtocol) GetNATRules() []NATRule`

GetNATRules returns the NATRules field if non-nil, zero value otherwise.

### GetNATRulesOk

`func (o *NetworkProtocol) GetNATRulesOk() (*[]NATRule, bool)`

GetNATRulesOk returns a tuple with the NATRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNATRules

`func (o *NetworkProtocol) SetNATRules(v []NATRule)`

SetNATRules sets NATRules field to given value.

### HasNATRules

`func (o *NetworkProtocol) HasNATRules() bool`

HasNATRules returns a boolean if a field has been set.

### GetNTPInherit

`func (o *NetworkProtocol) GetNTPInherit() bool`

GetNTPInherit returns the NTPInherit field if non-nil, zero value otherwise.

### GetNTPInheritOk

`func (o *NetworkProtocol) GetNTPInheritOk() (*bool, bool)`

GetNTPInheritOk returns a tuple with the NTPInherit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPInherit

`func (o *NetworkProtocol) SetNTPInherit(v bool)`

SetNTPInherit sets NTPInherit field to given value.

### HasNTPInherit

`func (o *NetworkProtocol) HasNTPInherit() bool`

HasNTPInherit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


