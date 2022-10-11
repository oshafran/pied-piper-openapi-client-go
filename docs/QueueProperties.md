# QueueProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentSize** | Pointer to **int32** |  | [optional] 
**MaxSize** | Pointer to **int32** |  | [optional] 
**TenantCurrentSize** | Pointer to **int32** |  | [optional] 
**TenantMaxSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewQueueProperties

`func NewQueueProperties() *QueueProperties`

NewQueueProperties instantiates a new QueueProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuePropertiesWithDefaults

`func NewQueuePropertiesWithDefaults() *QueueProperties`

NewQueuePropertiesWithDefaults instantiates a new QueueProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentSize

`func (o *QueueProperties) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *QueueProperties) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *QueueProperties) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *QueueProperties) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetMaxSize

`func (o *QueueProperties) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *QueueProperties) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *QueueProperties) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *QueueProperties) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetTenantCurrentSize

`func (o *QueueProperties) GetTenantCurrentSize() int32`

GetTenantCurrentSize returns the TenantCurrentSize field if non-nil, zero value otherwise.

### GetTenantCurrentSizeOk

`func (o *QueueProperties) GetTenantCurrentSizeOk() (*int32, bool)`

GetTenantCurrentSizeOk returns a tuple with the TenantCurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCurrentSize

`func (o *QueueProperties) SetTenantCurrentSize(v int32)`

SetTenantCurrentSize sets TenantCurrentSize field to given value.

### HasTenantCurrentSize

`func (o *QueueProperties) HasTenantCurrentSize() bool`

HasTenantCurrentSize returns a boolean if a field has been set.

### GetTenantMaxSize

`func (o *QueueProperties) GetTenantMaxSize() int32`

GetTenantMaxSize returns the TenantMaxSize field if non-nil, zero value otherwise.

### GetTenantMaxSizeOk

`func (o *QueueProperties) GetTenantMaxSizeOk() (*int32, bool)`

GetTenantMaxSizeOk returns a tuple with the TenantMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMaxSize

`func (o *QueueProperties) SetTenantMaxSize(v int32)`

SetTenantMaxSize sets TenantMaxSize field to given value.

### HasTenantMaxSize

`func (o *QueueProperties) HasTenantMaxSize() bool`

HasTenantMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


