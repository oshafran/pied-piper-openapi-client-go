# ThreadPoolsDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadPools** | Pointer to [**[]ThreadPoolDefinition**](ThreadPoolDefinition.md) |  | [optional] 
**TotalThreadCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewThreadPoolsDefinition

`func NewThreadPoolsDefinition() *ThreadPoolsDefinition`

NewThreadPoolsDefinition instantiates a new ThreadPoolsDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadPoolsDefinitionWithDefaults

`func NewThreadPoolsDefinitionWithDefaults() *ThreadPoolsDefinition`

NewThreadPoolsDefinitionWithDefaults instantiates a new ThreadPoolsDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadPools

`func (o *ThreadPoolsDefinition) GetThreadPools() []ThreadPoolDefinition`

GetThreadPools returns the ThreadPools field if non-nil, zero value otherwise.

### GetThreadPoolsOk

`func (o *ThreadPoolsDefinition) GetThreadPoolsOk() (*[]ThreadPoolDefinition, bool)`

GetThreadPoolsOk returns a tuple with the ThreadPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadPools

`func (o *ThreadPoolsDefinition) SetThreadPools(v []ThreadPoolDefinition)`

SetThreadPools sets ThreadPools field to given value.

### HasThreadPools

`func (o *ThreadPoolsDefinition) HasThreadPools() bool`

HasThreadPools returns a boolean if a field has been set.

### GetTotalThreadCount

`func (o *ThreadPoolsDefinition) GetTotalThreadCount() int32`

GetTotalThreadCount returns the TotalThreadCount field if non-nil, zero value otherwise.

### GetTotalThreadCountOk

`func (o *ThreadPoolsDefinition) GetTotalThreadCountOk() (*int32, bool)`

GetTotalThreadCountOk returns a tuple with the TotalThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThreadCount

`func (o *ThreadPoolsDefinition) SetTotalThreadCount(v int32)`

SetTotalThreadCount sets TotalThreadCount field to given value.

### HasTotalThreadCount

`func (o *ThreadPoolsDefinition) HasTotalThreadCount() bool`

HasTotalThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


