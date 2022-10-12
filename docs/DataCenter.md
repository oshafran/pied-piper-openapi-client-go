# DataCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DcPersonality** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]Node**](Node.md) |  | [optional] 
**MgmtIPAddress** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NmsPersonality** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDataCenter

`func NewDataCenter() *DataCenter`

NewDataCenter instantiates a new DataCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCenterWithDefaults

`func NewDataCenterWithDefaults() *DataCenter`

NewDataCenterWithDefaults instantiates a new DataCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDcPersonality

`func (o *DataCenter) GetDcPersonality() string`

GetDcPersonality returns the DcPersonality field if non-nil, zero value otherwise.

### GetDcPersonalityOk

`func (o *DataCenter) GetDcPersonalityOk() (*string, bool)`

GetDcPersonalityOk returns a tuple with the DcPersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcPersonality

`func (o *DataCenter) SetDcPersonality(v string)`

SetDcPersonality sets DcPersonality field to given value.

### HasDcPersonality

`func (o *DataCenter) HasDcPersonality() bool`

HasDcPersonality returns a boolean if a field has been set.

### GetMembers

`func (o *DataCenter) GetMembers() []Node`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DataCenter) GetMembersOk() (*[]Node, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DataCenter) SetMembers(v []Node)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DataCenter) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMgmtIPAddress

`func (o *DataCenter) GetMgmtIPAddress() string`

GetMgmtIPAddress returns the MgmtIPAddress field if non-nil, zero value otherwise.

### GetMgmtIPAddressOk

`func (o *DataCenter) GetMgmtIPAddressOk() (*string, bool)`

GetMgmtIPAddressOk returns a tuple with the MgmtIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIPAddress

`func (o *DataCenter) SetMgmtIPAddress(v string)`

SetMgmtIPAddress sets MgmtIPAddress field to given value.

### HasMgmtIPAddress

`func (o *DataCenter) HasMgmtIPAddress() bool`

HasMgmtIPAddress returns a boolean if a field has been set.

### GetName

`func (o *DataCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataCenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNmsPersonality

`func (o *DataCenter) GetNmsPersonality() string`

GetNmsPersonality returns the NmsPersonality field if non-nil, zero value otherwise.

### GetNmsPersonalityOk

`func (o *DataCenter) GetNmsPersonalityOk() (*string, bool)`

GetNmsPersonalityOk returns a tuple with the NmsPersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmsPersonality

`func (o *DataCenter) SetNmsPersonality(v string)`

SetNmsPersonality sets NmsPersonality field to given value.

### HasNmsPersonality

`func (o *DataCenter) HasNmsPersonality() bool`

HasNmsPersonality returns a boolean if a field has been set.

### GetPassword

`func (o *DataCenter) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DataCenter) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DataCenter) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DataCenter) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *DataCenter) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DataCenter) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DataCenter) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DataCenter) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


