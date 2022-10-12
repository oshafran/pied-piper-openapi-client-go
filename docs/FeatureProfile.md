# FeatureProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | User who last created this. | [optional] [readonly] 
**CreatedOn** | Pointer to **int64** | Timestamp of creation | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the feature Profile. | [optional] 
**Id** | Pointer to **string** | System generated unique identifier of the feature profile in UUID format. | [optional] 
**LastUpdatedBy** | Pointer to **string** | User who last updated this. | [optional] [readonly] 
**LastUpdatedOn** | Pointer to **int64** | Timestamp of last update | [optional] [readonly] 
**Name** | **string** | Name of the feature Profile. Must be unique. | 
**Solution** | **string** | Solution of the feature Profile. | 
**Type** | **string** | Type of the feature Profile. | 

## Methods

### NewFeatureProfile

`func NewFeatureProfile(name string, solution string, type_ string, ) *FeatureProfile`

NewFeatureProfile instantiates a new FeatureProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureProfileWithDefaults

`func NewFeatureProfileWithDefaults() *FeatureProfile`

NewFeatureProfileWithDefaults instantiates a new FeatureProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *FeatureProfile) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FeatureProfile) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FeatureProfile) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FeatureProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *FeatureProfile) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *FeatureProfile) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *FeatureProfile) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *FeatureProfile) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDescription

`func (o *FeatureProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FeatureProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *FeatureProfile) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *FeatureProfile) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *FeatureProfile) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *FeatureProfile) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *FeatureProfile) GetLastUpdatedOn() int64`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *FeatureProfile) GetLastUpdatedOnOk() (*int64, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *FeatureProfile) SetLastUpdatedOn(v int64)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *FeatureProfile) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetName

`func (o *FeatureProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureProfile) SetName(v string)`

SetName sets Name field to given value.


### GetSolution

`func (o *FeatureProfile) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *FeatureProfile) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *FeatureProfile) SetSolution(v string)`

SetSolution sets Solution field to given value.


### GetType

`func (o *FeatureProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureProfile) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


