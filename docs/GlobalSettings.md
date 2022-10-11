# GlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicName** | Pointer to **string** |  | [optional] 
**BasicDescription** | Pointer to **string** |  | [optional] 
**NtpServer** | Pointer to [**[]ConnectToNtpServer**](ConnectToNtpServer.md) |  | [optional] 
**Systems** | Pointer to [**Systems**](Systems.md) |  | [optional] 
**Banner** | Pointer to [**Banner**](Banner.md) |  | [optional] 
**LoginAccessToRouter** | Pointer to [**LoginAccessToRouter**](LoginAccessToRouter.md) |  | [optional] 
**Bfd** | Pointer to [**Bfd**](Bfd.md) |  | [optional] 
**Omp** | Pointer to [**OMP**](OMP.md) |  | [optional] 
**IpSecSecurity** | Pointer to [**IpSecSecurity**](IpSecSecurity.md) |  | [optional] 
**LoggingSystemMessages** | Pointer to [**LoggingSystemMessages**](LoggingSystemMessages.md) |  | [optional] 

## Methods

### NewGlobalSettings

`func NewGlobalSettings() *GlobalSettings`

NewGlobalSettings instantiates a new GlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSettingsWithDefaults

`func NewGlobalSettingsWithDefaults() *GlobalSettings`

NewGlobalSettingsWithDefaults instantiates a new GlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicName

`func (o *GlobalSettings) GetBasicName() string`

GetBasicName returns the BasicName field if non-nil, zero value otherwise.

### GetBasicNameOk

`func (o *GlobalSettings) GetBasicNameOk() (*string, bool)`

GetBasicNameOk returns a tuple with the BasicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicName

`func (o *GlobalSettings) SetBasicName(v string)`

SetBasicName sets BasicName field to given value.

### HasBasicName

`func (o *GlobalSettings) HasBasicName() bool`

HasBasicName returns a boolean if a field has been set.

### GetBasicDescription

`func (o *GlobalSettings) GetBasicDescription() string`

GetBasicDescription returns the BasicDescription field if non-nil, zero value otherwise.

### GetBasicDescriptionOk

`func (o *GlobalSettings) GetBasicDescriptionOk() (*string, bool)`

GetBasicDescriptionOk returns a tuple with the BasicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicDescription

`func (o *GlobalSettings) SetBasicDescription(v string)`

SetBasicDescription sets BasicDescription field to given value.

### HasBasicDescription

`func (o *GlobalSettings) HasBasicDescription() bool`

HasBasicDescription returns a boolean if a field has been set.

### GetNtpServer

`func (o *GlobalSettings) GetNtpServer() []ConnectToNtpServer`

GetNtpServer returns the NtpServer field if non-nil, zero value otherwise.

### GetNtpServerOk

`func (o *GlobalSettings) GetNtpServerOk() (*[]ConnectToNtpServer, bool)`

GetNtpServerOk returns a tuple with the NtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServer

`func (o *GlobalSettings) SetNtpServer(v []ConnectToNtpServer)`

SetNtpServer sets NtpServer field to given value.

### HasNtpServer

`func (o *GlobalSettings) HasNtpServer() bool`

HasNtpServer returns a boolean if a field has been set.

### GetSystems

`func (o *GlobalSettings) GetSystems() Systems`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *GlobalSettings) GetSystemsOk() (*Systems, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *GlobalSettings) SetSystems(v Systems)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *GlobalSettings) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetBanner

`func (o *GlobalSettings) GetBanner() Banner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GlobalSettings) GetBannerOk() (*Banner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GlobalSettings) SetBanner(v Banner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GlobalSettings) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetLoginAccessToRouter

`func (o *GlobalSettings) GetLoginAccessToRouter() LoginAccessToRouter`

GetLoginAccessToRouter returns the LoginAccessToRouter field if non-nil, zero value otherwise.

### GetLoginAccessToRouterOk

`func (o *GlobalSettings) GetLoginAccessToRouterOk() (*LoginAccessToRouter, bool)`

GetLoginAccessToRouterOk returns a tuple with the LoginAccessToRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccessToRouter

`func (o *GlobalSettings) SetLoginAccessToRouter(v LoginAccessToRouter)`

SetLoginAccessToRouter sets LoginAccessToRouter field to given value.

### HasLoginAccessToRouter

`func (o *GlobalSettings) HasLoginAccessToRouter() bool`

HasLoginAccessToRouter returns a boolean if a field has been set.

### GetBfd

`func (o *GlobalSettings) GetBfd() Bfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *GlobalSettings) GetBfdOk() (*Bfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *GlobalSettings) SetBfd(v Bfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *GlobalSettings) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetOmp

`func (o *GlobalSettings) GetOmp() OMP`

GetOmp returns the Omp field if non-nil, zero value otherwise.

### GetOmpOk

`func (o *GlobalSettings) GetOmpOk() (*OMP, bool)`

GetOmpOk returns a tuple with the Omp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmp

`func (o *GlobalSettings) SetOmp(v OMP)`

SetOmp sets Omp field to given value.

### HasOmp

`func (o *GlobalSettings) HasOmp() bool`

HasOmp returns a boolean if a field has been set.

### GetIpSecSecurity

`func (o *GlobalSettings) GetIpSecSecurity() IpSecSecurity`

GetIpSecSecurity returns the IpSecSecurity field if non-nil, zero value otherwise.

### GetIpSecSecurityOk

`func (o *GlobalSettings) GetIpSecSecurityOk() (*IpSecSecurity, bool)`

GetIpSecSecurityOk returns a tuple with the IpSecSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecSecurity

`func (o *GlobalSettings) SetIpSecSecurity(v IpSecSecurity)`

SetIpSecSecurity sets IpSecSecurity field to given value.

### HasIpSecSecurity

`func (o *GlobalSettings) HasIpSecSecurity() bool`

HasIpSecSecurity returns a boolean if a field has been set.

### GetLoggingSystemMessages

`func (o *GlobalSettings) GetLoggingSystemMessages() LoggingSystemMessages`

GetLoggingSystemMessages returns the LoggingSystemMessages field if non-nil, zero value otherwise.

### GetLoggingSystemMessagesOk

`func (o *GlobalSettings) GetLoggingSystemMessagesOk() (*LoggingSystemMessages, bool)`

GetLoggingSystemMessagesOk returns a tuple with the LoggingSystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingSystemMessages

`func (o *GlobalSettings) SetLoggingSystemMessages(v LoggingSystemMessages)`

SetLoggingSystemMessages sets LoggingSystemMessages field to given value.

### HasLoggingSystemMessages

`func (o *GlobalSettings) HasLoggingSystemMessages() bool`

HasLoggingSystemMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


