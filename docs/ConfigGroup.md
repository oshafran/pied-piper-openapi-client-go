# ConfigGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | User who last created this. | [optional] [readonly] 
**CreatedOn** | Pointer to **int64** | Timestamp of creation | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the Config Group. | [optional] 
**Devices** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | System generated unique identifier of the Config Group in UUID format. | [optional] 
**LastUpdatedBy** | Pointer to **string** | User who last updated this. | [optional] [readonly] 
**LastUpdatedOn** | Pointer to **int64** | Timestamp of last update | [optional] [readonly] 
**Name** | **string** | Name of the Config Group. Must be unique. | 
**Profiles** | Pointer to [**[]FeatureProfile**](FeatureProfile.md) | List of devices UUIDs associated with this config group | [optional] 
**Solution** | **string** | Specify one of the device platform solution | 
**Source** | Pointer to **string** | Source of config-group | [optional] 
**State** | **string** | Config Group Deployment state | 

## Methods

### NewConfigGroup

`func NewConfigGroup(name string, solution string, state string, ) *ConfigGroup`

NewConfigGroup instantiates a new ConfigGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigGroupWithDefaults

`func NewConfigGroupWithDefaults() *ConfigGroup`

NewConfigGroupWithDefaults instantiates a new ConfigGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ConfigGroup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ConfigGroup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ConfigGroup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ConfigGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ConfigGroup) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ConfigGroup) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ConfigGroup) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ConfigGroup) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevices

`func (o *ConfigGroup) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ConfigGroup) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ConfigGroup) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ConfigGroup) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *ConfigGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ConfigGroup) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ConfigGroup) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ConfigGroup) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ConfigGroup) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *ConfigGroup) GetLastUpdatedOn() int64`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *ConfigGroup) GetLastUpdatedOnOk() (*int64, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *ConfigGroup) SetLastUpdatedOn(v int64)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *ConfigGroup) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetName

`func (o *ConfigGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProfiles

`func (o *ConfigGroup) GetProfiles() []FeatureProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ConfigGroup) GetProfilesOk() (*[]FeatureProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ConfigGroup) SetProfiles(v []FeatureProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ConfigGroup) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetSolution

`func (o *ConfigGroup) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ConfigGroup) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ConfigGroup) SetSolution(v string)`

SetSolution sets Solution field to given value.


### GetSource

`func (o *ConfigGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConfigGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *ConfigGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConfigGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConfigGroup) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


