# TenantStatusListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TenantStatus**](TenantStatus.md) |  | [optional] 
**Header** | Pointer to [**Header**](Header.md) |  | [optional] 

## Methods

### NewTenantStatusListModel

`func NewTenantStatusListModel() *TenantStatusListModel`

NewTenantStatusListModel instantiates a new TenantStatusListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantStatusListModelWithDefaults

`func NewTenantStatusListModelWithDefaults() *TenantStatusListModel`

NewTenantStatusListModelWithDefaults instantiates a new TenantStatusListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TenantStatusListModel) GetData() []TenantStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TenantStatusListModel) GetDataOk() (*[]TenantStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TenantStatusListModel) SetData(v []TenantStatus)`

SetData sets Data field to given value.

### HasData

`func (o *TenantStatusListModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeader

`func (o *TenantStatusListModel) GetHeader() Header`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TenantStatusListModel) GetHeaderOk() (*Header, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TenantStatusListModel) SetHeader(v Header)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TenantStatusListModel) HasHeader() bool`

HasHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


