# Header

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]HeaderElement**](HeaderElement.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewHeader

`func NewHeader() *Header`

NewHeader instantiates a new Header object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderWithDefaults

`func NewHeaderWithDefaults() *Header`

NewHeaderWithDefaults instantiates a new Header object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *Header) GetElements() []HeaderElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Header) GetElementsOk() (*[]HeaderElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Header) SetElements(v []HeaderElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *Header) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetName

`func (o *Header) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Header) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Header) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Header) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *Header) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Header) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Header) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Header) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


