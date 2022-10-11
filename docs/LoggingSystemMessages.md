# LoggingSystemMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFileSize** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Rotations** | Pointer to **int32** |  | [optional] 

## Methods

### NewLoggingSystemMessages

`func NewLoggingSystemMessages() *LoggingSystemMessages`

NewLoggingSystemMessages instantiates a new LoggingSystemMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSystemMessagesWithDefaults

`func NewLoggingSystemMessagesWithDefaults() *LoggingSystemMessages`

NewLoggingSystemMessagesWithDefaults instantiates a new LoggingSystemMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFileSize

`func (o *LoggingSystemMessages) GetMaxFileSize() int32`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *LoggingSystemMessages) GetMaxFileSizeOk() (*int32, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *LoggingSystemMessages) SetMaxFileSize(v int32)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *LoggingSystemMessages) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetPriority

`func (o *LoggingSystemMessages) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LoggingSystemMessages) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LoggingSystemMessages) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LoggingSystemMessages) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRotations

`func (o *LoggingSystemMessages) GetRotations() int32`

GetRotations returns the Rotations field if non-nil, zero value otherwise.

### GetRotationsOk

`func (o *LoggingSystemMessages) GetRotationsOk() (*int32, bool)`

GetRotationsOk returns a tuple with the Rotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotations

`func (o *LoggingSystemMessages) SetRotations(v int32)`

SetRotations sets Rotations field to given value.

### HasRotations

`func (o *LoggingSystemMessages) HasRotations() bool`

HasRotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


