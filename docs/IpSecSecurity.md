# IpSecSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** |  | [optional] 
**IpSecPairwiseKeying** | Pointer to **string** |  | [optional] 
**RekeyTime** | Pointer to **int32** |  | [optional] 
**ReplayWindow** | Pointer to **int32** |  | [optional] 

## Methods

### NewIpSecSecurity

`func NewIpSecSecurity() *IpSecSecurity`

NewIpSecSecurity instantiates a new IpSecSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSecSecurityWithDefaults

`func NewIpSecSecurityWithDefaults() *IpSecSecurity`

NewIpSecSecurityWithDefaults instantiates a new IpSecSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *IpSecSecurity) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *IpSecSecurity) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *IpSecSecurity) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *IpSecSecurity) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetIpSecPairwiseKeying

`func (o *IpSecSecurity) GetIpSecPairwiseKeying() string`

GetIpSecPairwiseKeying returns the IpSecPairwiseKeying field if non-nil, zero value otherwise.

### GetIpSecPairwiseKeyingOk

`func (o *IpSecSecurity) GetIpSecPairwiseKeyingOk() (*string, bool)`

GetIpSecPairwiseKeyingOk returns a tuple with the IpSecPairwiseKeying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecPairwiseKeying

`func (o *IpSecSecurity) SetIpSecPairwiseKeying(v string)`

SetIpSecPairwiseKeying sets IpSecPairwiseKeying field to given value.

### HasIpSecPairwiseKeying

`func (o *IpSecSecurity) HasIpSecPairwiseKeying() bool`

HasIpSecPairwiseKeying returns a boolean if a field has been set.

### GetRekeyTime

`func (o *IpSecSecurity) GetRekeyTime() int32`

GetRekeyTime returns the RekeyTime field if non-nil, zero value otherwise.

### GetRekeyTimeOk

`func (o *IpSecSecurity) GetRekeyTimeOk() (*int32, bool)`

GetRekeyTimeOk returns a tuple with the RekeyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyTime

`func (o *IpSecSecurity) SetRekeyTime(v int32)`

SetRekeyTime sets RekeyTime field to given value.

### HasRekeyTime

`func (o *IpSecSecurity) HasRekeyTime() bool`

HasRekeyTime returns a boolean if a field has been set.

### GetReplayWindow

`func (o *IpSecSecurity) GetReplayWindow() int32`

GetReplayWindow returns the ReplayWindow field if non-nil, zero value otherwise.

### GetReplayWindowOk

`func (o *IpSecSecurity) GetReplayWindowOk() (*int32, bool)`

GetReplayWindowOk returns a tuple with the ReplayWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayWindow

`func (o *IpSecSecurity) SetReplayWindow(v int32)`

SetReplayWindow sets ReplayWindow field to given value.

### HasReplayWindow

`func (o *IpSecSecurity) HasReplayWindow() bool`

HasReplayWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


