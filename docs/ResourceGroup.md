# ResourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** |  | [optional] 
**DeviceIPs** | Pointer to **[]string** |  | [optional] 
**DeviceIps** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **map[string]interface{}** |  | [optional] 
**MgmtSytemIpsMap** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SiteIds** | Pointer to **[]int64** |  | [optional] 
**UuidSytemIpsMap** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewResourceGroup

`func NewResourceGroup() *ResourceGroup`

NewResourceGroup instantiates a new ResourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupWithDefaults

`func NewResourceGroupWithDefaults() *ResourceGroup`

NewResourceGroupWithDefaults instantiates a new ResourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *ResourceGroup) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ResourceGroup) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ResourceGroup) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ResourceGroup) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDeviceIPs

`func (o *ResourceGroup) GetDeviceIPs() []string`

GetDeviceIPs returns the DeviceIPs field if non-nil, zero value otherwise.

### GetDeviceIPsOk

`func (o *ResourceGroup) GetDeviceIPsOk() (*[]string, bool)`

GetDeviceIPsOk returns a tuple with the DeviceIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIPs

`func (o *ResourceGroup) SetDeviceIPs(v []string)`

SetDeviceIPs sets DeviceIPs field to given value.

### HasDeviceIPs

`func (o *ResourceGroup) HasDeviceIPs() bool`

HasDeviceIPs returns a boolean if a field has been set.

### GetDeviceIps

`func (o *ResourceGroup) GetDeviceIps() []string`

GetDeviceIps returns the DeviceIps field if non-nil, zero value otherwise.

### GetDeviceIpsOk

`func (o *ResourceGroup) GetDeviceIpsOk() (*[]string, bool)`

GetDeviceIpsOk returns a tuple with the DeviceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIps

`func (o *ResourceGroup) SetDeviceIps(v []string)`

SetDeviceIps sets DeviceIps field to given value.

### HasDeviceIps

`func (o *ResourceGroup) HasDeviceIps() bool`

HasDeviceIps returns a boolean if a field has been set.

### GetId

`func (o *ResourceGroup) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceGroup) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceGroup) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *ResourceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMgmtSytemIpsMap

`func (o *ResourceGroup) GetMgmtSytemIpsMap() map[string]string`

GetMgmtSytemIpsMap returns the MgmtSytemIpsMap field if non-nil, zero value otherwise.

### GetMgmtSytemIpsMapOk

`func (o *ResourceGroup) GetMgmtSytemIpsMapOk() (*map[string]string, bool)`

GetMgmtSytemIpsMapOk returns a tuple with the MgmtSytemIpsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtSytemIpsMap

`func (o *ResourceGroup) SetMgmtSytemIpsMap(v map[string]string)`

SetMgmtSytemIpsMap sets MgmtSytemIpsMap field to given value.

### HasMgmtSytemIpsMap

`func (o *ResourceGroup) HasMgmtSytemIpsMap() bool`

HasMgmtSytemIpsMap returns a boolean if a field has been set.

### GetName

`func (o *ResourceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteIds

`func (o *ResourceGroup) GetSiteIds() []int64`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *ResourceGroup) GetSiteIdsOk() (*[]int64, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *ResourceGroup) SetSiteIds(v []int64)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *ResourceGroup) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetUuidSytemIpsMap

`func (o *ResourceGroup) GetUuidSytemIpsMap() map[string]string`

GetUuidSytemIpsMap returns the UuidSytemIpsMap field if non-nil, zero value otherwise.

### GetUuidSytemIpsMapOk

`func (o *ResourceGroup) GetUuidSytemIpsMapOk() (*map[string]string, bool)`

GetUuidSytemIpsMapOk returns a tuple with the UuidSytemIpsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSytemIpsMap

`func (o *ResourceGroup) SetUuidSytemIpsMap(v map[string]string)`

SetUuidSytemIpsMap sets UuidSytemIpsMap field to given value.

### HasUuidSytemIpsMap

`func (o *ResourceGroup) HasUuidSytemIpsMap() bool`

HasUuidSytemIpsMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


