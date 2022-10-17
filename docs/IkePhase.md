# IkePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherSuite** | **string** |  | 
**DiffeHellmanGroup** | Pointer to **string** |  | [optional] 
**IkeVersion** | Pointer to **int32** |  | [optional] 
**RekeyTimer** | Pointer to **int32** |  | [optional] 

## Methods

### NewIkePhase

`func NewIkePhase(cipherSuite string, ) *IkePhase`

NewIkePhase instantiates a new IkePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIkePhaseWithDefaults

`func NewIkePhaseWithDefaults() *IkePhase`

NewIkePhaseWithDefaults instantiates a new IkePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherSuite

`func (o *IkePhase) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *IkePhase) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *IkePhase) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.


### GetDiffeHellmanGroup

`func (o *IkePhase) GetDiffeHellmanGroup() string`

GetDiffeHellmanGroup returns the DiffeHellmanGroup field if non-nil, zero value otherwise.

### GetDiffeHellmanGroupOk

`func (o *IkePhase) GetDiffeHellmanGroupOk() (*string, bool)`

GetDiffeHellmanGroupOk returns a tuple with the DiffeHellmanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffeHellmanGroup

`func (o *IkePhase) SetDiffeHellmanGroup(v string)`

SetDiffeHellmanGroup sets DiffeHellmanGroup field to given value.

### HasDiffeHellmanGroup

`func (o *IkePhase) HasDiffeHellmanGroup() bool`

HasDiffeHellmanGroup returns a boolean if a field has been set.

### GetIkeVersion

`func (o *IkePhase) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *IkePhase) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *IkePhase) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *IkePhase) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetRekeyTimer

`func (o *IkePhase) GetRekeyTimer() int32`

GetRekeyTimer returns the RekeyTimer field if non-nil, zero value otherwise.

### GetRekeyTimerOk

`func (o *IkePhase) GetRekeyTimerOk() (*int32, bool)`

GetRekeyTimerOk returns a tuple with the RekeyTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyTimer

`func (o *IkePhase) SetRekeyTimer(v int32)`

SetRekeyTimer sets RekeyTimer field to given value.

### HasRekeyTimer

`func (o *IkePhase) HasRekeyTimer() bool`

HasRekeyTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


