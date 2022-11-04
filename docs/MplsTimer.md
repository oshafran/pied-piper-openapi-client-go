# MplsTimer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dscp** | Pointer to **int32** |  | [optional] 
**HelloInterval** | Pointer to **int32** |  | [optional] 
**Multiplier** | Pointer to **int32** |  | [optional] 
**PathMtu** | Pointer to **string** |  | [optional] 

## Methods

### NewMplsTimer

`func NewMplsTimer() *MplsTimer`

NewMplsTimer instantiates a new MplsTimer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMplsTimerWithDefaults

`func NewMplsTimerWithDefaults() *MplsTimer`

NewMplsTimerWithDefaults instantiates a new MplsTimer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscp

`func (o *MplsTimer) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *MplsTimer) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *MplsTimer) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *MplsTimer) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetHelloInterval

`func (o *MplsTimer) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *MplsTimer) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *MplsTimer) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *MplsTimer) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetMultiplier

`func (o *MplsTimer) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *MplsTimer) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *MplsTimer) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *MplsTimer) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetPathMtu

`func (o *MplsTimer) GetPathMtu() string`

GetPathMtu returns the PathMtu field if non-nil, zero value otherwise.

### GetPathMtuOk

`func (o *MplsTimer) GetPathMtuOk() (*string, bool)`

GetPathMtuOk returns a tuple with the PathMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathMtu

`func (o *MplsTimer) SetPathMtu(v string)`

SetPathMtu sets PathMtu field to given value.

### HasPathMtu

`func (o *MplsTimer) HasPathMtu() bool`

HasPathMtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


