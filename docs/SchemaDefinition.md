# SchemaDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountOfUniqueuVertexAndEdgeObjectsCreated** | Pointer to **int32** |  | [optional] 
**CountOfVertexAndEdgesRegisteredWithDeviceDataManager** | Pointer to **int32** |  | [optional] 
**VertexDefinitionsForDataCollection** | Pointer to **[]string** |  | [optional] 
**VertexDefinitionsFoundInDBWithData** | Pointer to [**[]VertexDefinition**](VertexDefinition.md) |  | [optional] 
**VertexDefinitionsFoundInDBWithoutData** | Pointer to **[]string** |  | [optional] 
**VertexDefinitionsNotFoundInDB** | Pointer to **[]string** |  | [optional] 
**VertexDefinitionsNotRegisteredWithDeviceDataManager** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSchemaDefinition

`func NewSchemaDefinition() *SchemaDefinition`

NewSchemaDefinition instantiates a new SchemaDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDefinitionWithDefaults

`func NewSchemaDefinitionWithDefaults() *SchemaDefinition`

NewSchemaDefinitionWithDefaults instantiates a new SchemaDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountOfUniqueuVertexAndEdgeObjectsCreated

`func (o *SchemaDefinition) GetCountOfUniqueuVertexAndEdgeObjectsCreated() int32`

GetCountOfUniqueuVertexAndEdgeObjectsCreated returns the CountOfUniqueuVertexAndEdgeObjectsCreated field if non-nil, zero value otherwise.

### GetCountOfUniqueuVertexAndEdgeObjectsCreatedOk

`func (o *SchemaDefinition) GetCountOfUniqueuVertexAndEdgeObjectsCreatedOk() (*int32, bool)`

GetCountOfUniqueuVertexAndEdgeObjectsCreatedOk returns a tuple with the CountOfUniqueuVertexAndEdgeObjectsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOfUniqueuVertexAndEdgeObjectsCreated

`func (o *SchemaDefinition) SetCountOfUniqueuVertexAndEdgeObjectsCreated(v int32)`

SetCountOfUniqueuVertexAndEdgeObjectsCreated sets CountOfUniqueuVertexAndEdgeObjectsCreated field to given value.

### HasCountOfUniqueuVertexAndEdgeObjectsCreated

`func (o *SchemaDefinition) HasCountOfUniqueuVertexAndEdgeObjectsCreated() bool`

HasCountOfUniqueuVertexAndEdgeObjectsCreated returns a boolean if a field has been set.

### GetCountOfVertexAndEdgesRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) GetCountOfVertexAndEdgesRegisteredWithDeviceDataManager() int32`

GetCountOfVertexAndEdgesRegisteredWithDeviceDataManager returns the CountOfVertexAndEdgesRegisteredWithDeviceDataManager field if non-nil, zero value otherwise.

### GetCountOfVertexAndEdgesRegisteredWithDeviceDataManagerOk

`func (o *SchemaDefinition) GetCountOfVertexAndEdgesRegisteredWithDeviceDataManagerOk() (*int32, bool)`

GetCountOfVertexAndEdgesRegisteredWithDeviceDataManagerOk returns a tuple with the CountOfVertexAndEdgesRegisteredWithDeviceDataManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOfVertexAndEdgesRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) SetCountOfVertexAndEdgesRegisteredWithDeviceDataManager(v int32)`

SetCountOfVertexAndEdgesRegisteredWithDeviceDataManager sets CountOfVertexAndEdgesRegisteredWithDeviceDataManager field to given value.

### HasCountOfVertexAndEdgesRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) HasCountOfVertexAndEdgesRegisteredWithDeviceDataManager() bool`

HasCountOfVertexAndEdgesRegisteredWithDeviceDataManager returns a boolean if a field has been set.

### GetVertexDefinitionsForDataCollection

`func (o *SchemaDefinition) GetVertexDefinitionsForDataCollection() []string`

GetVertexDefinitionsForDataCollection returns the VertexDefinitionsForDataCollection field if non-nil, zero value otherwise.

### GetVertexDefinitionsForDataCollectionOk

`func (o *SchemaDefinition) GetVertexDefinitionsForDataCollectionOk() (*[]string, bool)`

GetVertexDefinitionsForDataCollectionOk returns a tuple with the VertexDefinitionsForDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexDefinitionsForDataCollection

`func (o *SchemaDefinition) SetVertexDefinitionsForDataCollection(v []string)`

SetVertexDefinitionsForDataCollection sets VertexDefinitionsForDataCollection field to given value.

### HasVertexDefinitionsForDataCollection

`func (o *SchemaDefinition) HasVertexDefinitionsForDataCollection() bool`

HasVertexDefinitionsForDataCollection returns a boolean if a field has been set.

### GetVertexDefinitionsFoundInDBWithData

`func (o *SchemaDefinition) GetVertexDefinitionsFoundInDBWithData() []VertexDefinition`

GetVertexDefinitionsFoundInDBWithData returns the VertexDefinitionsFoundInDBWithData field if non-nil, zero value otherwise.

### GetVertexDefinitionsFoundInDBWithDataOk

`func (o *SchemaDefinition) GetVertexDefinitionsFoundInDBWithDataOk() (*[]VertexDefinition, bool)`

GetVertexDefinitionsFoundInDBWithDataOk returns a tuple with the VertexDefinitionsFoundInDBWithData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexDefinitionsFoundInDBWithData

`func (o *SchemaDefinition) SetVertexDefinitionsFoundInDBWithData(v []VertexDefinition)`

SetVertexDefinitionsFoundInDBWithData sets VertexDefinitionsFoundInDBWithData field to given value.

### HasVertexDefinitionsFoundInDBWithData

`func (o *SchemaDefinition) HasVertexDefinitionsFoundInDBWithData() bool`

HasVertexDefinitionsFoundInDBWithData returns a boolean if a field has been set.

### GetVertexDefinitionsFoundInDBWithoutData

`func (o *SchemaDefinition) GetVertexDefinitionsFoundInDBWithoutData() []string`

GetVertexDefinitionsFoundInDBWithoutData returns the VertexDefinitionsFoundInDBWithoutData field if non-nil, zero value otherwise.

### GetVertexDefinitionsFoundInDBWithoutDataOk

`func (o *SchemaDefinition) GetVertexDefinitionsFoundInDBWithoutDataOk() (*[]string, bool)`

GetVertexDefinitionsFoundInDBWithoutDataOk returns a tuple with the VertexDefinitionsFoundInDBWithoutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexDefinitionsFoundInDBWithoutData

`func (o *SchemaDefinition) SetVertexDefinitionsFoundInDBWithoutData(v []string)`

SetVertexDefinitionsFoundInDBWithoutData sets VertexDefinitionsFoundInDBWithoutData field to given value.

### HasVertexDefinitionsFoundInDBWithoutData

`func (o *SchemaDefinition) HasVertexDefinitionsFoundInDBWithoutData() bool`

HasVertexDefinitionsFoundInDBWithoutData returns a boolean if a field has been set.

### GetVertexDefinitionsNotFoundInDB

`func (o *SchemaDefinition) GetVertexDefinitionsNotFoundInDB() []string`

GetVertexDefinitionsNotFoundInDB returns the VertexDefinitionsNotFoundInDB field if non-nil, zero value otherwise.

### GetVertexDefinitionsNotFoundInDBOk

`func (o *SchemaDefinition) GetVertexDefinitionsNotFoundInDBOk() (*[]string, bool)`

GetVertexDefinitionsNotFoundInDBOk returns a tuple with the VertexDefinitionsNotFoundInDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexDefinitionsNotFoundInDB

`func (o *SchemaDefinition) SetVertexDefinitionsNotFoundInDB(v []string)`

SetVertexDefinitionsNotFoundInDB sets VertexDefinitionsNotFoundInDB field to given value.

### HasVertexDefinitionsNotFoundInDB

`func (o *SchemaDefinition) HasVertexDefinitionsNotFoundInDB() bool`

HasVertexDefinitionsNotFoundInDB returns a boolean if a field has been set.

### GetVertexDefinitionsNotRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) GetVertexDefinitionsNotRegisteredWithDeviceDataManager() []string`

GetVertexDefinitionsNotRegisteredWithDeviceDataManager returns the VertexDefinitionsNotRegisteredWithDeviceDataManager field if non-nil, zero value otherwise.

### GetVertexDefinitionsNotRegisteredWithDeviceDataManagerOk

`func (o *SchemaDefinition) GetVertexDefinitionsNotRegisteredWithDeviceDataManagerOk() (*[]string, bool)`

GetVertexDefinitionsNotRegisteredWithDeviceDataManagerOk returns a tuple with the VertexDefinitionsNotRegisteredWithDeviceDataManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexDefinitionsNotRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) SetVertexDefinitionsNotRegisteredWithDeviceDataManager(v []string)`

SetVertexDefinitionsNotRegisteredWithDeviceDataManager sets VertexDefinitionsNotRegisteredWithDeviceDataManager field to given value.

### HasVertexDefinitionsNotRegisteredWithDeviceDataManager

`func (o *SchemaDefinition) HasVertexDefinitionsNotRegisteredWithDeviceDataManager() bool`

HasVertexDefinitionsNotRegisteredWithDeviceDataManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


