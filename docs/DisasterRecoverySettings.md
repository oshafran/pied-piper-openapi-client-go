# DisasterRecoverySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayThreshold** | Pointer to **int32** |  | [optional] 
**DrPaused** | Pointer to **bool** |  | [optional] 
**IgnoredIndexes** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**IntervalModified** | Pointer to **bool** |  | [optional] 
**PauseDR** | Pointer to [**DisasterRecoverySettings**](DisasterRecoverySettings.md) |  | [optional] 
**PauseReplication** | Pointer to **bool** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDisasterRecoverySettings

`func NewDisasterRecoverySettings() *DisasterRecoverySettings`

NewDisasterRecoverySettings instantiates a new DisasterRecoverySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisasterRecoverySettingsWithDefaults

`func NewDisasterRecoverySettingsWithDefaults() *DisasterRecoverySettings`

NewDisasterRecoverySettingsWithDefaults instantiates a new DisasterRecoverySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayThreshold

`func (o *DisasterRecoverySettings) GetDelayThreshold() int32`

GetDelayThreshold returns the DelayThreshold field if non-nil, zero value otherwise.

### GetDelayThresholdOk

`func (o *DisasterRecoverySettings) GetDelayThresholdOk() (*int32, bool)`

GetDelayThresholdOk returns a tuple with the DelayThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayThreshold

`func (o *DisasterRecoverySettings) SetDelayThreshold(v int32)`

SetDelayThreshold sets DelayThreshold field to given value.

### HasDelayThreshold

`func (o *DisasterRecoverySettings) HasDelayThreshold() bool`

HasDelayThreshold returns a boolean if a field has been set.

### GetDrPaused

`func (o *DisasterRecoverySettings) GetDrPaused() bool`

GetDrPaused returns the DrPaused field if non-nil, zero value otherwise.

### GetDrPausedOk

`func (o *DisasterRecoverySettings) GetDrPausedOk() (*bool, bool)`

GetDrPausedOk returns a tuple with the DrPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrPaused

`func (o *DisasterRecoverySettings) SetDrPaused(v bool)`

SetDrPaused sets DrPaused field to given value.

### HasDrPaused

`func (o *DisasterRecoverySettings) HasDrPaused() bool`

HasDrPaused returns a boolean if a field has been set.

### GetIgnoredIndexes

`func (o *DisasterRecoverySettings) GetIgnoredIndexes() []string`

GetIgnoredIndexes returns the IgnoredIndexes field if non-nil, zero value otherwise.

### GetIgnoredIndexesOk

`func (o *DisasterRecoverySettings) GetIgnoredIndexesOk() (*[]string, bool)`

GetIgnoredIndexesOk returns a tuple with the IgnoredIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredIndexes

`func (o *DisasterRecoverySettings) SetIgnoredIndexes(v []string)`

SetIgnoredIndexes sets IgnoredIndexes field to given value.

### HasIgnoredIndexes

`func (o *DisasterRecoverySettings) HasIgnoredIndexes() bool`

HasIgnoredIndexes returns a boolean if a field has been set.

### GetInterval

`func (o *DisasterRecoverySettings) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DisasterRecoverySettings) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DisasterRecoverySettings) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DisasterRecoverySettings) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalModified

`func (o *DisasterRecoverySettings) GetIntervalModified() bool`

GetIntervalModified returns the IntervalModified field if non-nil, zero value otherwise.

### GetIntervalModifiedOk

`func (o *DisasterRecoverySettings) GetIntervalModifiedOk() (*bool, bool)`

GetIntervalModifiedOk returns a tuple with the IntervalModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalModified

`func (o *DisasterRecoverySettings) SetIntervalModified(v bool)`

SetIntervalModified sets IntervalModified field to given value.

### HasIntervalModified

`func (o *DisasterRecoverySettings) HasIntervalModified() bool`

HasIntervalModified returns a boolean if a field has been set.

### GetPauseDR

`func (o *DisasterRecoverySettings) GetPauseDR() DisasterRecoverySettings`

GetPauseDR returns the PauseDR field if non-nil, zero value otherwise.

### GetPauseDROk

`func (o *DisasterRecoverySettings) GetPauseDROk() (*DisasterRecoverySettings, bool)`

GetPauseDROk returns a tuple with the PauseDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDR

`func (o *DisasterRecoverySettings) SetPauseDR(v DisasterRecoverySettings)`

SetPauseDR sets PauseDR field to given value.

### HasPauseDR

`func (o *DisasterRecoverySettings) HasPauseDR() bool`

HasPauseDR returns a boolean if a field has been set.

### GetPauseReplication

`func (o *DisasterRecoverySettings) GetPauseReplication() bool`

GetPauseReplication returns the PauseReplication field if non-nil, zero value otherwise.

### GetPauseReplicationOk

`func (o *DisasterRecoverySettings) GetPauseReplicationOk() (*bool, bool)`

GetPauseReplicationOk returns a tuple with the PauseReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReplication

`func (o *DisasterRecoverySettings) SetPauseReplication(v bool)`

SetPauseReplication sets PauseReplication field to given value.

### HasPauseReplication

`func (o *DisasterRecoverySettings) HasPauseReplication() bool`

HasPauseReplication returns a boolean if a field has been set.

### GetStartTime

`func (o *DisasterRecoverySettings) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DisasterRecoverySettings) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DisasterRecoverySettings) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DisasterRecoverySettings) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


