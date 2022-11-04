# DomainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**ResolvedIp** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDomainDetail

`func NewDomainDetail() *DomainDetail`

NewDomainDetail instantiates a new DomainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailWithDefaults

`func NewDomainDetailWithDefaults() *DomainDetail`

NewDomainDetailWithDefaults instantiates a new DomainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainDetail) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetResolvedIp

`func (o *DomainDetail) GetResolvedIp() []string`

GetResolvedIp returns the ResolvedIp field if non-nil, zero value otherwise.

### GetResolvedIpOk

`func (o *DomainDetail) GetResolvedIpOk() (*[]string, bool)`

GetResolvedIpOk returns a tuple with the ResolvedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIp

`func (o *DomainDetail) SetResolvedIp(v []string)`

SetResolvedIp sets ResolvedIp field to given value.

### HasResolvedIp

`func (o *DomainDetail) HasResolvedIp() bool`

HasResolvedIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


