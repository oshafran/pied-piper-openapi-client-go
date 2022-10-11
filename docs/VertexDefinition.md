# VertexDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InEdges** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutEdges** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**[]PropertyDefinition**](PropertyDefinition.md) |  | [optional] 

## Methods

### NewVertexDefinition

`func NewVertexDefinition() *VertexDefinition`

NewVertexDefinition instantiates a new VertexDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVertexDefinitionWithDefaults

`func NewVertexDefinitionWithDefaults() *VertexDefinition`

NewVertexDefinitionWithDefaults instantiates a new VertexDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInEdges

`func (o *VertexDefinition) GetInEdges() []string`

GetInEdges returns the InEdges field if non-nil, zero value otherwise.

### GetInEdgesOk

`func (o *VertexDefinition) GetInEdgesOk() (*[]string, bool)`

GetInEdgesOk returns a tuple with the InEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInEdges

`func (o *VertexDefinition) SetInEdges(v []string)`

SetInEdges sets InEdges field to given value.

### HasInEdges

`func (o *VertexDefinition) HasInEdges() bool`

HasInEdges returns a boolean if a field has been set.

### GetName

`func (o *VertexDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VertexDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VertexDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VertexDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutEdges

`func (o *VertexDefinition) GetOutEdges() []string`

GetOutEdges returns the OutEdges field if non-nil, zero value otherwise.

### GetOutEdgesOk

`func (o *VertexDefinition) GetOutEdgesOk() (*[]string, bool)`

GetOutEdgesOk returns a tuple with the OutEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutEdges

`func (o *VertexDefinition) SetOutEdges(v []string)`

SetOutEdges sets OutEdges field to given value.

### HasOutEdges

`func (o *VertexDefinition) HasOutEdges() bool`

HasOutEdges returns a boolean if a field has been set.

### GetProperties

`func (o *VertexDefinition) GetProperties() []PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VertexDefinition) GetPropertiesOk() (*[]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VertexDefinition) SetProperties(v []PropertyDefinition)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *VertexDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


