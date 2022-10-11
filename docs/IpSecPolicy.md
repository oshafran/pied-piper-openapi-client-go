# IpSecPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IkePhase1** | [**IkePhase**](IkePhase.md) |  | 
**IkePhase2CipherSuite** | **string** |  | 
**Preset** | Pointer to **string** |  | [optional] 

## Methods

### NewIpSecPolicy

`func NewIpSecPolicy(ikePhase1 IkePhase, ikePhase2CipherSuite string, ) *IpSecPolicy`

NewIpSecPolicy instantiates a new IpSecPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSecPolicyWithDefaults

`func NewIpSecPolicyWithDefaults() *IpSecPolicy`

NewIpSecPolicyWithDefaults instantiates a new IpSecPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIkePhase1

`func (o *IpSecPolicy) GetIkePhase1() IkePhase`

GetIkePhase1 returns the IkePhase1 field if non-nil, zero value otherwise.

### GetIkePhase1Ok

`func (o *IpSecPolicy) GetIkePhase1Ok() (*IkePhase, bool)`

GetIkePhase1Ok returns a tuple with the IkePhase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePhase1

`func (o *IpSecPolicy) SetIkePhase1(v IkePhase)`

SetIkePhase1 sets IkePhase1 field to given value.


### GetIkePhase2CipherSuite

`func (o *IpSecPolicy) GetIkePhase2CipherSuite() string`

GetIkePhase2CipherSuite returns the IkePhase2CipherSuite field if non-nil, zero value otherwise.

### GetIkePhase2CipherSuiteOk

`func (o *IpSecPolicy) GetIkePhase2CipherSuiteOk() (*string, bool)`

GetIkePhase2CipherSuiteOk returns a tuple with the IkePhase2CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePhase2CipherSuite

`func (o *IpSecPolicy) SetIkePhase2CipherSuite(v string)`

SetIkePhase2CipherSuite sets IkePhase2CipherSuite field to given value.


### GetPreset

`func (o *IpSecPolicy) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *IpSecPolicy) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *IpSecPolicy) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *IpSecPolicy) HasPreset() bool`

HasPreset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


