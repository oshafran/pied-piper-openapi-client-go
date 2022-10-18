# NetworkProtocolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DHCPPool** | Pointer to [**DHCPPool**](DHCPPool.md) |  | [optional] 
**DNSSettings** | Pointer to **string** |  | [optional] 
**NTPSettings** | Pointer to **[]string** |  | [optional] 
**NATRules** | Pointer to [**[]NATRule**](NATRule.md) |  | [optional] 
**NTPInherit** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkProtocolAllOf

`func NewNetworkProtocolAllOf() *NetworkProtocolAllOf`

NewNetworkProtocolAllOf instantiates a new NetworkProtocolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtocolAllOfWithDefaults

`func NewNetworkProtocolAllOfWithDefaults() *NetworkProtocolAllOf`

NewNetworkProtocolAllOfWithDefaults instantiates a new NetworkProtocolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtocolAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtocolAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtocolAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkProtocolAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NetworkProtocolAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkProtocolAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkProtocolAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkProtocolAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDHCPPool

`func (o *NetworkProtocolAllOf) GetDHCPPool() DHCPPool`

GetDHCPPool returns the DHCPPool field if non-nil, zero value otherwise.

### GetDHCPPoolOk

`func (o *NetworkProtocolAllOf) GetDHCPPoolOk() (*DHCPPool, bool)`

GetDHCPPoolOk returns a tuple with the DHCPPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDHCPPool

`func (o *NetworkProtocolAllOf) SetDHCPPool(v DHCPPool)`

SetDHCPPool sets DHCPPool field to given value.

### HasDHCPPool

`func (o *NetworkProtocolAllOf) HasDHCPPool() bool`

HasDHCPPool returns a boolean if a field has been set.

### GetDNSSettings

`func (o *NetworkProtocolAllOf) GetDNSSettings() string`

GetDNSSettings returns the DNSSettings field if non-nil, zero value otherwise.

### GetDNSSettingsOk

`func (o *NetworkProtocolAllOf) GetDNSSettingsOk() (*string, bool)`

GetDNSSettingsOk returns a tuple with the DNSSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSSettings

`func (o *NetworkProtocolAllOf) SetDNSSettings(v string)`

SetDNSSettings sets DNSSettings field to given value.

### HasDNSSettings

`func (o *NetworkProtocolAllOf) HasDNSSettings() bool`

HasDNSSettings returns a boolean if a field has been set.

### GetNTPSettings

`func (o *NetworkProtocolAllOf) GetNTPSettings() []string`

GetNTPSettings returns the NTPSettings field if non-nil, zero value otherwise.

### GetNTPSettingsOk

`func (o *NetworkProtocolAllOf) GetNTPSettingsOk() (*[]string, bool)`

GetNTPSettingsOk returns a tuple with the NTPSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPSettings

`func (o *NetworkProtocolAllOf) SetNTPSettings(v []string)`

SetNTPSettings sets NTPSettings field to given value.

### HasNTPSettings

`func (o *NetworkProtocolAllOf) HasNTPSettings() bool`

HasNTPSettings returns a boolean if a field has been set.

### GetNATRules

`func (o *NetworkProtocolAllOf) GetNATRules() []NATRule`

GetNATRules returns the NATRules field if non-nil, zero value otherwise.

### GetNATRulesOk

`func (o *NetworkProtocolAllOf) GetNATRulesOk() (*[]NATRule, bool)`

GetNATRulesOk returns a tuple with the NATRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNATRules

`func (o *NetworkProtocolAllOf) SetNATRules(v []NATRule)`

SetNATRules sets NATRules field to given value.

### HasNATRules

`func (o *NetworkProtocolAllOf) HasNATRules() bool`

HasNATRules returns a boolean if a field has been set.

### GetNTPInherit

`func (o *NetworkProtocolAllOf) GetNTPInherit() bool`

GetNTPInherit returns the NTPInherit field if non-nil, zero value otherwise.

### GetNTPInheritOk

`func (o *NetworkProtocolAllOf) GetNTPInheritOk() (*bool, bool)`

GetNTPInheritOk returns a tuple with the NTPInherit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPInherit

`func (o *NetworkProtocolAllOf) SetNTPInherit(v bool)`

SetNTPInherit sets NTPInherit field to given value.

### HasNTPInherit

`func (o *NetworkProtocolAllOf) HasNTPInherit() bool`

HasNTPInherit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


