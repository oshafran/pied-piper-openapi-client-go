# Banner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginMessage** | Pointer to **string** |  | [optional] 
**Motd** | Pointer to **string** |  | [optional] 

## Methods

### NewBanner

`func NewBanner() *Banner`

NewBanner instantiates a new Banner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannerWithDefaults

`func NewBannerWithDefaults() *Banner`

NewBannerWithDefaults instantiates a new Banner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginMessage

`func (o *Banner) GetLoginMessage() string`

GetLoginMessage returns the LoginMessage field if non-nil, zero value otherwise.

### GetLoginMessageOk

`func (o *Banner) GetLoginMessageOk() (*string, bool)`

GetLoginMessageOk returns a tuple with the LoginMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMessage

`func (o *Banner) SetLoginMessage(v string)`

SetLoginMessage sets LoginMessage field to given value.

### HasLoginMessage

`func (o *Banner) HasLoginMessage() bool`

HasLoginMessage returns a boolean if a field has been set.

### GetMotd

`func (o *Banner) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *Banner) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *Banner) SetMotd(v string)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *Banner) HasMotd() bool`

HasMotd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


