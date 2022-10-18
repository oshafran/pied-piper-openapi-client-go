# SimSlotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachProfileId** | Pointer to **int32** |  | [optional] 
**CarrierName** | Pointer to **string** |  | [optional] 
**DataProfileIdList** | Pointer to **[]int32** |  | [optional] 
**ProfileList** | Pointer to [**[]CellularProfile**](CellularProfile.md) |  | [optional] 
**SlotNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimSlotConfig

`func NewSimSlotConfig() *SimSlotConfig`

NewSimSlotConfig instantiates a new SimSlotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimSlotConfigWithDefaults

`func NewSimSlotConfigWithDefaults() *SimSlotConfig`

NewSimSlotConfigWithDefaults instantiates a new SimSlotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachProfileId

`func (o *SimSlotConfig) GetAttachProfileId() int32`

GetAttachProfileId returns the AttachProfileId field if non-nil, zero value otherwise.

### GetAttachProfileIdOk

`func (o *SimSlotConfig) GetAttachProfileIdOk() (*int32, bool)`

GetAttachProfileIdOk returns a tuple with the AttachProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachProfileId

`func (o *SimSlotConfig) SetAttachProfileId(v int32)`

SetAttachProfileId sets AttachProfileId field to given value.

### HasAttachProfileId

`func (o *SimSlotConfig) HasAttachProfileId() bool`

HasAttachProfileId returns a boolean if a field has been set.

### GetCarrierName

`func (o *SimSlotConfig) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *SimSlotConfig) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *SimSlotConfig) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *SimSlotConfig) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetDataProfileIdList

`func (o *SimSlotConfig) GetDataProfileIdList() []int32`

GetDataProfileIdList returns the DataProfileIdList field if non-nil, zero value otherwise.

### GetDataProfileIdListOk

`func (o *SimSlotConfig) GetDataProfileIdListOk() (*[]int32, bool)`

GetDataProfileIdListOk returns a tuple with the DataProfileIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProfileIdList

`func (o *SimSlotConfig) SetDataProfileIdList(v []int32)`

SetDataProfileIdList sets DataProfileIdList field to given value.

### HasDataProfileIdList

`func (o *SimSlotConfig) HasDataProfileIdList() bool`

HasDataProfileIdList returns a boolean if a field has been set.

### GetProfileList

`func (o *SimSlotConfig) GetProfileList() []CellularProfile`

GetProfileList returns the ProfileList field if non-nil, zero value otherwise.

### GetProfileListOk

`func (o *SimSlotConfig) GetProfileListOk() (*[]CellularProfile, bool)`

GetProfileListOk returns a tuple with the ProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileList

`func (o *SimSlotConfig) SetProfileList(v []CellularProfile)`

SetProfileList sets ProfileList field to given value.

### HasProfileList

`func (o *SimSlotConfig) HasProfileList() bool`

HasProfileList returns a boolean if a field has been set.

### GetSlotNumber

`func (o *SimSlotConfig) GetSlotNumber() int32`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *SimSlotConfig) GetSlotNumberOk() (*int32, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *SimSlotConfig) SetSlotNumber(v int32)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *SimSlotConfig) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


