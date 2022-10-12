# PropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPropertyDefinition

`func NewPropertyDefinition() *PropertyDefinition`

NewPropertyDefinition instantiates a new PropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDefinitionWithDefaults

`func NewPropertyDefinitionWithDefaults() *PropertyDefinition`

NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSamples

`func (o *PropertyDefinition) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *PropertyDefinition) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *PropertyDefinition) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *PropertyDefinition) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetType

`func (o *PropertyDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


