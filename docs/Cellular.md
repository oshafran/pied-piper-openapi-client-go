# Cellular

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimSlot0** | [**SimSlotConfig**](SimSlotConfig.md) |  | 
**SimSlot1** | Pointer to [**SimSlotConfig**](SimSlotConfig.md) |  | [optional] 
**PrimarySlot** | Pointer to **int32** |  | [optional] 
**WanConfig** | Pointer to **string** |  | [optional] 

## Methods

### NewCellular

`func NewCellular(simSlot0 SimSlotConfig, ) *Cellular`

NewCellular instantiates a new Cellular object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellularWithDefaults

`func NewCellularWithDefaults() *Cellular`

NewCellularWithDefaults instantiates a new Cellular object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimSlot0

`func (o *Cellular) GetSimSlot0() SimSlotConfig`

GetSimSlot0 returns the SimSlot0 field if non-nil, zero value otherwise.

### GetSimSlot0Ok

`func (o *Cellular) GetSimSlot0Ok() (*SimSlotConfig, bool)`

GetSimSlot0Ok returns a tuple with the SimSlot0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSlot0

`func (o *Cellular) SetSimSlot0(v SimSlotConfig)`

SetSimSlot0 sets SimSlot0 field to given value.


### GetSimSlot1

`func (o *Cellular) GetSimSlot1() SimSlotConfig`

GetSimSlot1 returns the SimSlot1 field if non-nil, zero value otherwise.

### GetSimSlot1Ok

`func (o *Cellular) GetSimSlot1Ok() (*SimSlotConfig, bool)`

GetSimSlot1Ok returns a tuple with the SimSlot1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSlot1

`func (o *Cellular) SetSimSlot1(v SimSlotConfig)`

SetSimSlot1 sets SimSlot1 field to given value.

### HasSimSlot1

`func (o *Cellular) HasSimSlot1() bool`

HasSimSlot1 returns a boolean if a field has been set.

### GetPrimarySlot

`func (o *Cellular) GetPrimarySlot() int32`

GetPrimarySlot returns the PrimarySlot field if non-nil, zero value otherwise.

### GetPrimarySlotOk

`func (o *Cellular) GetPrimarySlotOk() (*int32, bool)`

GetPrimarySlotOk returns a tuple with the PrimarySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySlot

`func (o *Cellular) SetPrimarySlot(v int32)`

SetPrimarySlot sets PrimarySlot field to given value.

### HasPrimarySlot

`func (o *Cellular) HasPrimarySlot() bool`

HasPrimarySlot returns a boolean if a field has been set.

### GetWanConfig

`func (o *Cellular) GetWanConfig() string`

GetWanConfig returns the WanConfig field if non-nil, zero value otherwise.

### GetWanConfigOk

`func (o *Cellular) GetWanConfigOk() (*string, bool)`

GetWanConfigOk returns a tuple with the WanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanConfig

`func (o *Cellular) SetWanConfig(v string)`

SetWanConfig sets WanConfig field to given value.

### HasWanConfig

`func (o *Cellular) HasWanConfig() bool`

HasWanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


