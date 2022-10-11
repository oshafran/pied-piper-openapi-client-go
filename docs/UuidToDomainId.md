# UuidToDomainId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Mapping** | Pointer to [**[]UuidToDomainIdMapping**](UuidToDomainIdMapping.md) |  | [optional] 

## Methods

### NewUuidToDomainId

`func NewUuidToDomainId() *UuidToDomainId`

NewUuidToDomainId instantiates a new UuidToDomainId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidToDomainIdWithDefaults

`func NewUuidToDomainIdWithDefaults() *UuidToDomainId`

NewUuidToDomainIdWithDefaults instantiates a new UuidToDomainId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UuidToDomainId) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UuidToDomainId) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UuidToDomainId) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UuidToDomainId) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMapping

`func (o *UuidToDomainId) GetMapping() []UuidToDomainIdMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *UuidToDomainId) GetMappingOk() (*[]UuidToDomainIdMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *UuidToDomainId) SetMapping(v []UuidToDomainIdMapping)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *UuidToDomainId) HasMapping() bool`

HasMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


