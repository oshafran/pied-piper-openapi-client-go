# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonPath** | **string** |  | 
**VarName** | **string** |  | 

## Methods

### NewVariable

`func NewVariable(jsonPath string, varName string, ) *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonPath

`func (o *Variable) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *Variable) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *Variable) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.


### GetVarName

`func (o *Variable) GetVarName() string`

GetVarName returns the VarName field if non-nil, zero value otherwise.

### GetVarNameOk

`func (o *Variable) GetVarNameOk() (*string, bool)`

GetVarNameOk returns a tuple with the VarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarName

`func (o *Variable) SetVarName(v string)`

SetVarName sets VarName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


