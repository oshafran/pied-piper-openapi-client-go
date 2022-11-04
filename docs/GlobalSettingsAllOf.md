# GlobalSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
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

### NewGlobalSettingsAllOf

`func NewGlobalSettingsAllOf() *GlobalSettingsAllOf`

NewGlobalSettingsAllOf instantiates a new GlobalSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSettingsAllOfWithDefaults

`func NewGlobalSettingsAllOfWithDefaults() *GlobalSettingsAllOf`

NewGlobalSettingsAllOfWithDefaults instantiates a new GlobalSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlobalSettingsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalSettingsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalSettingsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalSettingsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GlobalSettingsAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalSettingsAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalSettingsAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalSettingsAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBasicName

`func (o *GlobalSettingsAllOf) GetBasicName() string`

GetBasicName returns the BasicName field if non-nil, zero value otherwise.

### GetBasicNameOk

`func (o *GlobalSettingsAllOf) GetBasicNameOk() (*string, bool)`

GetBasicNameOk returns a tuple with the BasicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicName

`func (o *GlobalSettingsAllOf) SetBasicName(v string)`

SetBasicName sets BasicName field to given value.

### HasBasicName

`func (o *GlobalSettingsAllOf) HasBasicName() bool`

HasBasicName returns a boolean if a field has been set.

### GetBasicDescription

`func (o *GlobalSettingsAllOf) GetBasicDescription() string`

GetBasicDescription returns the BasicDescription field if non-nil, zero value otherwise.

### GetBasicDescriptionOk

`func (o *GlobalSettingsAllOf) GetBasicDescriptionOk() (*string, bool)`

GetBasicDescriptionOk returns a tuple with the BasicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicDescription

`func (o *GlobalSettingsAllOf) SetBasicDescription(v string)`

SetBasicDescription sets BasicDescription field to given value.

### HasBasicDescription

`func (o *GlobalSettingsAllOf) HasBasicDescription() bool`

HasBasicDescription returns a boolean if a field has been set.

### GetNtpServer

`func (o *GlobalSettingsAllOf) GetNtpServer() []ConnectToNtpServer`

GetNtpServer returns the NtpServer field if non-nil, zero value otherwise.

### GetNtpServerOk

`func (o *GlobalSettingsAllOf) GetNtpServerOk() (*[]ConnectToNtpServer, bool)`

GetNtpServerOk returns a tuple with the NtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServer

`func (o *GlobalSettingsAllOf) SetNtpServer(v []ConnectToNtpServer)`

SetNtpServer sets NtpServer field to given value.

### HasNtpServer

`func (o *GlobalSettingsAllOf) HasNtpServer() bool`

HasNtpServer returns a boolean if a field has been set.

### GetSystems

`func (o *GlobalSettingsAllOf) GetSystems() Systems`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *GlobalSettingsAllOf) GetSystemsOk() (*Systems, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *GlobalSettingsAllOf) SetSystems(v Systems)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *GlobalSettingsAllOf) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetBanner

`func (o *GlobalSettingsAllOf) GetBanner() Banner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *GlobalSettingsAllOf) GetBannerOk() (*Banner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *GlobalSettingsAllOf) SetBanner(v Banner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *GlobalSettingsAllOf) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetLoginAccessToRouter

`func (o *GlobalSettingsAllOf) GetLoginAccessToRouter() LoginAccessToRouter`

GetLoginAccessToRouter returns the LoginAccessToRouter field if non-nil, zero value otherwise.

### GetLoginAccessToRouterOk

`func (o *GlobalSettingsAllOf) GetLoginAccessToRouterOk() (*LoginAccessToRouter, bool)`

GetLoginAccessToRouterOk returns a tuple with the LoginAccessToRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccessToRouter

`func (o *GlobalSettingsAllOf) SetLoginAccessToRouter(v LoginAccessToRouter)`

SetLoginAccessToRouter sets LoginAccessToRouter field to given value.

### HasLoginAccessToRouter

`func (o *GlobalSettingsAllOf) HasLoginAccessToRouter() bool`

HasLoginAccessToRouter returns a boolean if a field has been set.

### GetBfd

`func (o *GlobalSettingsAllOf) GetBfd() Bfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *GlobalSettingsAllOf) GetBfdOk() (*Bfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *GlobalSettingsAllOf) SetBfd(v Bfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *GlobalSettingsAllOf) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetOmp

`func (o *GlobalSettingsAllOf) GetOmp() OMP`

GetOmp returns the Omp field if non-nil, zero value otherwise.

### GetOmpOk

`func (o *GlobalSettingsAllOf) GetOmpOk() (*OMP, bool)`

GetOmpOk returns a tuple with the Omp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmp

`func (o *GlobalSettingsAllOf) SetOmp(v OMP)`

SetOmp sets Omp field to given value.

### HasOmp

`func (o *GlobalSettingsAllOf) HasOmp() bool`

HasOmp returns a boolean if a field has been set.

### GetIpSecSecurity

`func (o *GlobalSettingsAllOf) GetIpSecSecurity() IpSecSecurity`

GetIpSecSecurity returns the IpSecSecurity field if non-nil, zero value otherwise.

### GetIpSecSecurityOk

`func (o *GlobalSettingsAllOf) GetIpSecSecurityOk() (*IpSecSecurity, bool)`

GetIpSecSecurityOk returns a tuple with the IpSecSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecSecurity

`func (o *GlobalSettingsAllOf) SetIpSecSecurity(v IpSecSecurity)`

SetIpSecSecurity sets IpSecSecurity field to given value.

### HasIpSecSecurity

`func (o *GlobalSettingsAllOf) HasIpSecSecurity() bool`

HasIpSecSecurity returns a boolean if a field has been set.

### GetLoggingSystemMessages

`func (o *GlobalSettingsAllOf) GetLoggingSystemMessages() LoggingSystemMessages`

GetLoggingSystemMessages returns the LoggingSystemMessages field if non-nil, zero value otherwise.

### GetLoggingSystemMessagesOk

`func (o *GlobalSettingsAllOf) GetLoggingSystemMessagesOk() (*LoggingSystemMessages, bool)`

GetLoggingSystemMessagesOk returns a tuple with the LoggingSystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingSystemMessages

`func (o *GlobalSettingsAllOf) SetLoggingSystemMessages(v LoggingSystemMessages)`

SetLoggingSystemMessages sets LoggingSystemMessages field to given value.

### HasLoggingSystemMessages

`func (o *GlobalSettingsAllOf) HasLoggingSystemMessages() bool`

HasLoggingSystemMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


