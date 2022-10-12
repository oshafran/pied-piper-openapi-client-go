# RadiusServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewRadiusServer

`func NewRadiusServer() *RadiusServer`

NewRadiusServer instantiates a new RadiusServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusServerWithDefaults

`func NewRadiusServerWithDefaults() *RadiusServer`

NewRadiusServerWithDefaults instantiates a new RadiusServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RadiusServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RadiusServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RadiusServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RadiusServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *RadiusServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadiusServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadiusServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RadiusServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *RadiusServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RadiusServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RadiusServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RadiusServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


