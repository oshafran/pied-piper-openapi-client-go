# DeleteTenantBulkModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**TenantIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteTenantBulkModel

`func NewDeleteTenantBulkModel() *DeleteTenantBulkModel`

NewDeleteTenantBulkModel instantiates a new DeleteTenantBulkModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTenantBulkModelWithDefaults

`func NewDeleteTenantBulkModelWithDefaults() *DeleteTenantBulkModel`

NewDeleteTenantBulkModelWithDefaults instantiates a new DeleteTenantBulkModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *DeleteTenantBulkModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeleteTenantBulkModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeleteTenantBulkModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeleteTenantBulkModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTenantIdList

`func (o *DeleteTenantBulkModel) GetTenantIdList() []string`

GetTenantIdList returns the TenantIdList field if non-nil, zero value otherwise.

### GetTenantIdListOk

`func (o *DeleteTenantBulkModel) GetTenantIdListOk() (*[]string, bool)`

GetTenantIdListOk returns a tuple with the TenantIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdList

`func (o *DeleteTenantBulkModel) SetTenantIdList(v []string)`

SetTenantIdList sets TenantIdList field to given value.

### HasTenantIdList

`func (o *DeleteTenantBulkModel) HasTenantIdList() bool`

HasTenantIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


