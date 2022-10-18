# OnDemandQueueEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**CompletionTime** | Pointer to **int64** |  | [optional] 
**CreationTime** | Pointer to **int64** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**StartProcessingTime** | Pointer to **int64** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TimePeriod** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] 

## Methods

### NewOnDemandQueueEntry

`func NewOnDemandQueueEntry() *OnDemandQueueEntry`

NewOnDemandQueueEntry instantiates a new OnDemandQueueEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnDemandQueueEntryWithDefaults

`func NewOnDemandQueueEntryWithDefaults() *OnDemandQueueEntry`

NewOnDemandQueueEntryWithDefaults instantiates a new OnDemandQueueEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *OnDemandQueueEntry) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *OnDemandQueueEntry) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *OnDemandQueueEntry) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *OnDemandQueueEntry) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCompletionTime

`func (o *OnDemandQueueEntry) GetCompletionTime() int64`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *OnDemandQueueEntry) GetCompletionTimeOk() (*int64, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *OnDemandQueueEntry) SetCompletionTime(v int64)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *OnDemandQueueEntry) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *OnDemandQueueEntry) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *OnDemandQueueEntry) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *OnDemandQueueEntry) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *OnDemandQueueEntry) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDataType

`func (o *OnDemandQueueEntry) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *OnDemandQueueEntry) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *OnDemandQueueEntry) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *OnDemandQueueEntry) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDeviceId

`func (o *OnDemandQueueEntry) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *OnDemandQueueEntry) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *OnDemandQueueEntry) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *OnDemandQueueEntry) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEndTime

`func (o *OnDemandQueueEntry) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *OnDemandQueueEntry) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *OnDemandQueueEntry) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *OnDemandQueueEntry) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *OnDemandQueueEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnDemandQueueEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnDemandQueueEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnDemandQueueEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartProcessingTime

`func (o *OnDemandQueueEntry) GetStartProcessingTime() int64`

GetStartProcessingTime returns the StartProcessingTime field if non-nil, zero value otherwise.

### GetStartProcessingTimeOk

`func (o *OnDemandQueueEntry) GetStartProcessingTimeOk() (*int64, bool)`

GetStartProcessingTimeOk returns a tuple with the StartProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartProcessingTime

`func (o *OnDemandQueueEntry) SetStartProcessingTime(v int64)`

SetStartProcessingTime sets StartProcessingTime field to given value.

### HasStartProcessingTime

`func (o *OnDemandQueueEntry) HasStartProcessingTime() bool`

HasStartProcessingTime returns a boolean if a field has been set.

### GetStartTime

`func (o *OnDemandQueueEntry) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OnDemandQueueEntry) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OnDemandQueueEntry) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *OnDemandQueueEntry) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *OnDemandQueueEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OnDemandQueueEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OnDemandQueueEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OnDemandQueueEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *OnDemandQueueEntry) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OnDemandQueueEntry) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OnDemandQueueEntry) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OnDemandQueueEntry) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTimePeriod

`func (o *OnDemandQueueEntry) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *OnDemandQueueEntry) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *OnDemandQueueEntry) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *OnDemandQueueEntry) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetValue

`func (o *OnDemandQueueEntry) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OnDemandQueueEntry) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OnDemandQueueEntry) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *OnDemandQueueEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


