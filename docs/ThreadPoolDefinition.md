# ThreadPoolDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerClass** | Pointer to **string** |  | [optional] 
**ConsumerMethod** | Pointer to **string** |  | [optional] 
**ThreadPoolName** | Pointer to **string** |  | [optional] 
**ThreadPoolSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewThreadPoolDefinition

`func NewThreadPoolDefinition() *ThreadPoolDefinition`

NewThreadPoolDefinition instantiates a new ThreadPoolDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadPoolDefinitionWithDefaults

`func NewThreadPoolDefinitionWithDefaults() *ThreadPoolDefinition`

NewThreadPoolDefinitionWithDefaults instantiates a new ThreadPoolDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerClass

`func (o *ThreadPoolDefinition) GetConsumerClass() string`

GetConsumerClass returns the ConsumerClass field if non-nil, zero value otherwise.

### GetConsumerClassOk

`func (o *ThreadPoolDefinition) GetConsumerClassOk() (*string, bool)`

GetConsumerClassOk returns a tuple with the ConsumerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerClass

`func (o *ThreadPoolDefinition) SetConsumerClass(v string)`

SetConsumerClass sets ConsumerClass field to given value.

### HasConsumerClass

`func (o *ThreadPoolDefinition) HasConsumerClass() bool`

HasConsumerClass returns a boolean if a field has been set.

### GetConsumerMethod

`func (o *ThreadPoolDefinition) GetConsumerMethod() string`

GetConsumerMethod returns the ConsumerMethod field if non-nil, zero value otherwise.

### GetConsumerMethodOk

`func (o *ThreadPoolDefinition) GetConsumerMethodOk() (*string, bool)`

GetConsumerMethodOk returns a tuple with the ConsumerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerMethod

`func (o *ThreadPoolDefinition) SetConsumerMethod(v string)`

SetConsumerMethod sets ConsumerMethod field to given value.

### HasConsumerMethod

`func (o *ThreadPoolDefinition) HasConsumerMethod() bool`

HasConsumerMethod returns a boolean if a field has been set.

### GetThreadPoolName

`func (o *ThreadPoolDefinition) GetThreadPoolName() string`

GetThreadPoolName returns the ThreadPoolName field if non-nil, zero value otherwise.

### GetThreadPoolNameOk

`func (o *ThreadPoolDefinition) GetThreadPoolNameOk() (*string, bool)`

GetThreadPoolNameOk returns a tuple with the ThreadPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadPoolName

`func (o *ThreadPoolDefinition) SetThreadPoolName(v string)`

SetThreadPoolName sets ThreadPoolName field to given value.

### HasThreadPoolName

`func (o *ThreadPoolDefinition) HasThreadPoolName() bool`

HasThreadPoolName returns a boolean if a field has been set.

### GetThreadPoolSize

`func (o *ThreadPoolDefinition) GetThreadPoolSize() int32`

GetThreadPoolSize returns the ThreadPoolSize field if non-nil, zero value otherwise.

### GetThreadPoolSizeOk

`func (o *ThreadPoolDefinition) GetThreadPoolSizeOk() (*int32, bool)`

GetThreadPoolSizeOk returns a tuple with the ThreadPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadPoolSize

`func (o *ThreadPoolDefinition) SetThreadPoolSize(v int32)`

SetThreadPoolSize sets ThreadPoolSize field to given value.

### HasThreadPoolSize

`func (o *ThreadPoolDefinition) HasThreadPoolSize() bool`

HasThreadPoolSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


