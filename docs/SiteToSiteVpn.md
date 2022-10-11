# SiteToSiteVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalInterface** | **string** |  | 
**LocalPrivateSubnet** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**PreSharedSecret** | **string** |  | 
**RemotePrivateSubnets** | **string** |  | 
**RemotePublicIp** | **string** |  | 
**TunnelDnsAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteToSiteVpn

`func NewSiteToSiteVpn(localInterface string, localPrivateSubnet string, preSharedSecret string, remotePrivateSubnets string, remotePublicIp string, ) *SiteToSiteVpn`

NewSiteToSiteVpn instantiates a new SiteToSiteVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteToSiteVpnWithDefaults

`func NewSiteToSiteVpnWithDefaults() *SiteToSiteVpn`

NewSiteToSiteVpnWithDefaults instantiates a new SiteToSiteVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalInterface

`func (o *SiteToSiteVpn) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *SiteToSiteVpn) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *SiteToSiteVpn) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.


### GetLocalPrivateSubnet

`func (o *SiteToSiteVpn) GetLocalPrivateSubnet() string`

GetLocalPrivateSubnet returns the LocalPrivateSubnet field if non-nil, zero value otherwise.

### GetLocalPrivateSubnetOk

`func (o *SiteToSiteVpn) GetLocalPrivateSubnetOk() (*string, bool)`

GetLocalPrivateSubnetOk returns a tuple with the LocalPrivateSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPrivateSubnet

`func (o *SiteToSiteVpn) SetLocalPrivateSubnet(v string)`

SetLocalPrivateSubnet sets LocalPrivateSubnet field to given value.


### GetName

`func (o *SiteToSiteVpn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteToSiteVpn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteToSiteVpn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteToSiteVpn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreSharedSecret

`func (o *SiteToSiteVpn) GetPreSharedSecret() string`

GetPreSharedSecret returns the PreSharedSecret field if non-nil, zero value otherwise.

### GetPreSharedSecretOk

`func (o *SiteToSiteVpn) GetPreSharedSecretOk() (*string, bool)`

GetPreSharedSecretOk returns a tuple with the PreSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedSecret

`func (o *SiteToSiteVpn) SetPreSharedSecret(v string)`

SetPreSharedSecret sets PreSharedSecret field to given value.


### GetRemotePrivateSubnets

`func (o *SiteToSiteVpn) GetRemotePrivateSubnets() string`

GetRemotePrivateSubnets returns the RemotePrivateSubnets field if non-nil, zero value otherwise.

### GetRemotePrivateSubnetsOk

`func (o *SiteToSiteVpn) GetRemotePrivateSubnetsOk() (*string, bool)`

GetRemotePrivateSubnetsOk returns a tuple with the RemotePrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePrivateSubnets

`func (o *SiteToSiteVpn) SetRemotePrivateSubnets(v string)`

SetRemotePrivateSubnets sets RemotePrivateSubnets field to given value.


### GetRemotePublicIp

`func (o *SiteToSiteVpn) GetRemotePublicIp() string`

GetRemotePublicIp returns the RemotePublicIp field if non-nil, zero value otherwise.

### GetRemotePublicIpOk

`func (o *SiteToSiteVpn) GetRemotePublicIpOk() (*string, bool)`

GetRemotePublicIpOk returns a tuple with the RemotePublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIp

`func (o *SiteToSiteVpn) SetRemotePublicIp(v string)`

SetRemotePublicIp sets RemotePublicIp field to given value.


### GetTunnelDnsAddress

`func (o *SiteToSiteVpn) GetTunnelDnsAddress() string`

GetTunnelDnsAddress returns the TunnelDnsAddress field if non-nil, zero value otherwise.

### GetTunnelDnsAddressOk

`func (o *SiteToSiteVpn) GetTunnelDnsAddressOk() (*string, bool)`

GetTunnelDnsAddressOk returns a tuple with the TunnelDnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelDnsAddress

`func (o *SiteToSiteVpn) SetTunnelDnsAddress(v string)`

SetTunnelDnsAddress sets TunnelDnsAddress field to given value.

### HasTunnelDnsAddress

`func (o *SiteToSiteVpn) HasTunnelDnsAddress() bool`

HasTunnelDnsAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


