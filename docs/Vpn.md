# Vpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteToSiteVpn** | [**SiteToSiteVpn**](SiteToSiteVpn.md) |  | 
**IpSecPolicy** | Pointer to [**IpSecPolicy**](IpSecPolicy.md) |  | [optional] 

## Methods

### NewVpn

`func NewVpn(siteToSiteVpn SiteToSiteVpn, ) *Vpn`

NewVpn instantiates a new Vpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnWithDefaults

`func NewVpnWithDefaults() *Vpn`

NewVpnWithDefaults instantiates a new Vpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteToSiteVpn

`func (o *Vpn) GetSiteToSiteVpn() SiteToSiteVpn`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *Vpn) GetSiteToSiteVpnOk() (*SiteToSiteVpn, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *Vpn) SetSiteToSiteVpn(v SiteToSiteVpn)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.


### GetIpSecPolicy

`func (o *Vpn) GetIpSecPolicy() IpSecPolicy`

GetIpSecPolicy returns the IpSecPolicy field if non-nil, zero value otherwise.

### GetIpSecPolicyOk

`func (o *Vpn) GetIpSecPolicyOk() (*IpSecPolicy, bool)`

GetIpSecPolicyOk returns a tuple with the IpSecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecPolicy

`func (o *Vpn) SetIpSecPolicy(v IpSecPolicy)`

SetIpSecPolicy sets IpSecPolicy field to given value.

### HasIpSecPolicy

`func (o *Vpn) HasIpSecPolicy() bool`

HasIpSecPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


