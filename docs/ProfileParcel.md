# ProfileParcel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | User who last created this. | [optional] [readonly] 
**CreatedOn** | Pointer to **int64** | Timestamp of creation | [optional] [readonly] 
**Id** | Pointer to **string** | System generated unique identifier of the Profile Parcel in UUID format. | [optional] 
**LastUpdatedBy** | Pointer to **string** | User who last updated this. | [optional] [readonly] 
**LastUpdatedOn** | Pointer to **int64** | Timestamp of last update | [optional] [readonly] 
**Name** | **string** | Name of the Profile Parcel. Must be unique. | 
**Type** | **string** | type | 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 

## Methods

### NewProfileParcel

`func NewProfileParcel(name string, type_ string, ) *ProfileParcel`

NewProfileParcel instantiates a new ProfileParcel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileParcelWithDefaults

`func NewProfileParcelWithDefaults() *ProfileParcel`

NewProfileParcelWithDefaults instantiates a new ProfileParcel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ProfileParcel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProfileParcel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProfileParcel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProfileParcel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ProfileParcel) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ProfileParcel) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ProfileParcel) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ProfileParcel) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetId

`func (o *ProfileParcel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileParcel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileParcel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileParcel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ProfileParcel) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ProfileParcel) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ProfileParcel) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ProfileParcel) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *ProfileParcel) GetLastUpdatedOn() int64`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *ProfileParcel) GetLastUpdatedOnOk() (*int64, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *ProfileParcel) SetLastUpdatedOn(v int64)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *ProfileParcel) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetName

`func (o *ProfileParcel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileParcel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileParcel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ProfileParcel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileParcel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileParcel) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *ProfileParcel) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProfileParcel) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProfileParcel) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProfileParcel) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


