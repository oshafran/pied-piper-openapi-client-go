# HeaderElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ParameterCount** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewHeaderElement

`func NewHeaderElement() *HeaderElement`

NewHeaderElement instantiates a new HeaderElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderElementWithDefaults

`func NewHeaderElementWithDefaults() *HeaderElement`

NewHeaderElementWithDefaults instantiates a new HeaderElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HeaderElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeaderElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeaderElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeaderElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterCount

`func (o *HeaderElement) GetParameterCount() int32`

GetParameterCount returns the ParameterCount field if non-nil, zero value otherwise.

### GetParameterCountOk

`func (o *HeaderElement) GetParameterCountOk() (*int32, bool)`

GetParameterCountOk returns a tuple with the ParameterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterCount

`func (o *HeaderElement) SetParameterCount(v int32)`

SetParameterCount sets ParameterCount field to given value.

### HasParameterCount

`func (o *HeaderElement) HasParameterCount() bool`

HasParameterCount returns a boolean if a field has been set.

### GetParameters

`func (o *HeaderElement) GetParameters() []NameValuePair`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *HeaderElement) GetParametersOk() (*[]NameValuePair, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *HeaderElement) SetParameters(v []NameValuePair)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *HeaderElement) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetValue

`func (o *HeaderElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HeaderElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HeaderElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HeaderElement) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


