# VpnAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SiteToSiteVpn** | Pointer to [**SiteToSiteVpn**](SiteToSiteVpn.md) |  | [optional] 
**IpSecPolicy** | Pointer to [**IpSecPolicy**](IpSecPolicy.md) |  | [optional] 

## Methods

### NewVpnAllOf

`func NewVpnAllOf() *VpnAllOf`

NewVpnAllOf instantiates a new VpnAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnAllOfWithDefaults

`func NewVpnAllOfWithDefaults() *VpnAllOf`

NewVpnAllOfWithDefaults instantiates a new VpnAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VpnAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpnAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpnAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VpnAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSiteToSiteVpn

`func (o *VpnAllOf) GetSiteToSiteVpn() SiteToSiteVpn`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *VpnAllOf) GetSiteToSiteVpnOk() (*SiteToSiteVpn, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *VpnAllOf) SetSiteToSiteVpn(v SiteToSiteVpn)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *VpnAllOf) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.

### GetIpSecPolicy

`func (o *VpnAllOf) GetIpSecPolicy() IpSecPolicy`

GetIpSecPolicy returns the IpSecPolicy field if non-nil, zero value otherwise.

### GetIpSecPolicyOk

`func (o *VpnAllOf) GetIpSecPolicyOk() (*IpSecPolicy, bool)`

GetIpSecPolicyOk returns a tuple with the IpSecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecPolicy

`func (o *VpnAllOf) SetIpSecPolicy(v IpSecPolicy)`

SetIpSecPolicy sets IpSecPolicy field to given value.

### HasIpSecPolicy

`func (o *VpnAllOf) HasIpSecPolicy() bool`

HasIpSecPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


