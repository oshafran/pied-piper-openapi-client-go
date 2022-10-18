# DeviceHealthDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdSessions** | Pointer to **int32** |  | [optional] 
**BfdSessionsUp** | Pointer to **int32** |  | [optional] 
**BoardSerialNumber** | Pointer to **string** |  | [optional] 
**ChassisNumber** | Pointer to **string** |  | [optional] 
**ConnectedVmanages** | Pointer to **[]string** |  | [optional] 
**ControlConnectionsToVsmat** | Pointer to [**DeviceHealthDetails**](DeviceHealthDetails.md) |  | [optional] 
**ControlConnections** | Pointer to **int32** |  | [optional] 
**ControlConnectionsUp** | Pointer to **int32** |  | [optional] 
**CpuLoad** | Pointer to **float64** |  | [optional] 
**DeviceGroups** | Pointer to **[]string** |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**ExpectedVsmartConnections** | Pointer to **int32** |  | [optional] 
**HasGeoData** | Pointer to **bool** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **string** |  | [optional] 
**LocalSystemIp** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **string** |  | [optional] 
**MemoryUtilization** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OmpPeers** | Pointer to **int64** |  | [optional] 
**OmpPeersUp** | Pointer to **int64** |  | [optional] 
**Personality** | Pointer to **string** |  | [optional] 
**Qoe** | Pointer to **int32** |  | [optional] 
**Reachability** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**SystemIp** | Pointer to **string** |  | [optional] 
**UptimeDate** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**VpnIds** | Pointer to **[]string** |  | [optional] 
**VsmartControlConnections** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeviceHealthDetails

`func NewDeviceHealthDetails() *DeviceHealthDetails`

NewDeviceHealthDetails instantiates a new DeviceHealthDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceHealthDetailsWithDefaults

`func NewDeviceHealthDetailsWithDefaults() *DeviceHealthDetails`

NewDeviceHealthDetailsWithDefaults instantiates a new DeviceHealthDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdSessions

`func (o *DeviceHealthDetails) GetBfdSessions() int32`

GetBfdSessions returns the BfdSessions field if non-nil, zero value otherwise.

### GetBfdSessionsOk

`func (o *DeviceHealthDetails) GetBfdSessionsOk() (*int32, bool)`

GetBfdSessionsOk returns a tuple with the BfdSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdSessions

`func (o *DeviceHealthDetails) SetBfdSessions(v int32)`

SetBfdSessions sets BfdSessions field to given value.

### HasBfdSessions

`func (o *DeviceHealthDetails) HasBfdSessions() bool`

HasBfdSessions returns a boolean if a field has been set.

### GetBfdSessionsUp

`func (o *DeviceHealthDetails) GetBfdSessionsUp() int32`

GetBfdSessionsUp returns the BfdSessionsUp field if non-nil, zero value otherwise.

### GetBfdSessionsUpOk

`func (o *DeviceHealthDetails) GetBfdSessionsUpOk() (*int32, bool)`

GetBfdSessionsUpOk returns a tuple with the BfdSessionsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdSessionsUp

`func (o *DeviceHealthDetails) SetBfdSessionsUp(v int32)`

SetBfdSessionsUp sets BfdSessionsUp field to given value.

### HasBfdSessionsUp

`func (o *DeviceHealthDetails) HasBfdSessionsUp() bool`

HasBfdSessionsUp returns a boolean if a field has been set.

### GetBoardSerialNumber

`func (o *DeviceHealthDetails) GetBoardSerialNumber() string`

GetBoardSerialNumber returns the BoardSerialNumber field if non-nil, zero value otherwise.

### GetBoardSerialNumberOk

`func (o *DeviceHealthDetails) GetBoardSerialNumberOk() (*string, bool)`

GetBoardSerialNumberOk returns a tuple with the BoardSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardSerialNumber

`func (o *DeviceHealthDetails) SetBoardSerialNumber(v string)`

SetBoardSerialNumber sets BoardSerialNumber field to given value.

### HasBoardSerialNumber

`func (o *DeviceHealthDetails) HasBoardSerialNumber() bool`

HasBoardSerialNumber returns a boolean if a field has been set.

### GetChassisNumber

`func (o *DeviceHealthDetails) GetChassisNumber() string`

GetChassisNumber returns the ChassisNumber field if non-nil, zero value otherwise.

### GetChassisNumberOk

`func (o *DeviceHealthDetails) GetChassisNumberOk() (*string, bool)`

GetChassisNumberOk returns a tuple with the ChassisNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisNumber

`func (o *DeviceHealthDetails) SetChassisNumber(v string)`

SetChassisNumber sets ChassisNumber field to given value.

### HasChassisNumber

`func (o *DeviceHealthDetails) HasChassisNumber() bool`

HasChassisNumber returns a boolean if a field has been set.

### GetConnectedVmanages

`func (o *DeviceHealthDetails) GetConnectedVmanages() []string`

GetConnectedVmanages returns the ConnectedVmanages field if non-nil, zero value otherwise.

### GetConnectedVmanagesOk

`func (o *DeviceHealthDetails) GetConnectedVmanagesOk() (*[]string, bool)`

GetConnectedVmanagesOk returns a tuple with the ConnectedVmanages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVmanages

`func (o *DeviceHealthDetails) SetConnectedVmanages(v []string)`

SetConnectedVmanages sets ConnectedVmanages field to given value.

### HasConnectedVmanages

`func (o *DeviceHealthDetails) HasConnectedVmanages() bool`

HasConnectedVmanages returns a boolean if a field has been set.

### GetControlConnectionsToVsmat

`func (o *DeviceHealthDetails) GetControlConnectionsToVsmat() DeviceHealthDetails`

GetControlConnectionsToVsmat returns the ControlConnectionsToVsmat field if non-nil, zero value otherwise.

### GetControlConnectionsToVsmatOk

`func (o *DeviceHealthDetails) GetControlConnectionsToVsmatOk() (*DeviceHealthDetails, bool)`

GetControlConnectionsToVsmatOk returns a tuple with the ControlConnectionsToVsmat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlConnectionsToVsmat

`func (o *DeviceHealthDetails) SetControlConnectionsToVsmat(v DeviceHealthDetails)`

SetControlConnectionsToVsmat sets ControlConnectionsToVsmat field to given value.

### HasControlConnectionsToVsmat

`func (o *DeviceHealthDetails) HasControlConnectionsToVsmat() bool`

HasControlConnectionsToVsmat returns a boolean if a field has been set.

### GetControlConnections

`func (o *DeviceHealthDetails) GetControlConnections() int32`

GetControlConnections returns the ControlConnections field if non-nil, zero value otherwise.

### GetControlConnectionsOk

`func (o *DeviceHealthDetails) GetControlConnectionsOk() (*int32, bool)`

GetControlConnectionsOk returns a tuple with the ControlConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlConnections

`func (o *DeviceHealthDetails) SetControlConnections(v int32)`

SetControlConnections sets ControlConnections field to given value.

### HasControlConnections

`func (o *DeviceHealthDetails) HasControlConnections() bool`

HasControlConnections returns a boolean if a field has been set.

### GetControlConnectionsUp

`func (o *DeviceHealthDetails) GetControlConnectionsUp() int32`

GetControlConnectionsUp returns the ControlConnectionsUp field if non-nil, zero value otherwise.

### GetControlConnectionsUpOk

`func (o *DeviceHealthDetails) GetControlConnectionsUpOk() (*int32, bool)`

GetControlConnectionsUpOk returns a tuple with the ControlConnectionsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlConnectionsUp

`func (o *DeviceHealthDetails) SetControlConnectionsUp(v int32)`

SetControlConnectionsUp sets ControlConnectionsUp field to given value.

### HasControlConnectionsUp

`func (o *DeviceHealthDetails) HasControlConnectionsUp() bool`

HasControlConnectionsUp returns a boolean if a field has been set.

### GetCpuLoad

`func (o *DeviceHealthDetails) GetCpuLoad() float64`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *DeviceHealthDetails) GetCpuLoadOk() (*float64, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *DeviceHealthDetails) SetCpuLoad(v float64)`

SetCpuLoad sets CpuLoad field to given value.

### HasCpuLoad

`func (o *DeviceHealthDetails) HasCpuLoad() bool`

HasCpuLoad returns a boolean if a field has been set.

### GetDeviceGroups

`func (o *DeviceHealthDetails) GetDeviceGroups() []string`

GetDeviceGroups returns the DeviceGroups field if non-nil, zero value otherwise.

### GetDeviceGroupsOk

`func (o *DeviceHealthDetails) GetDeviceGroupsOk() (*[]string, bool)`

GetDeviceGroupsOk returns a tuple with the DeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroups

`func (o *DeviceHealthDetails) SetDeviceGroups(v []string)`

SetDeviceGroups sets DeviceGroups field to given value.

### HasDeviceGroups

`func (o *DeviceHealthDetails) HasDeviceGroups() bool`

HasDeviceGroups returns a boolean if a field has been set.

### GetDeviceModel

`func (o *DeviceHealthDetails) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DeviceHealthDetails) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DeviceHealthDetails) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DeviceHealthDetails) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceHealthDetails) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceHealthDetails) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceHealthDetails) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceHealthDetails) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetExpectedVsmartConnections

`func (o *DeviceHealthDetails) GetExpectedVsmartConnections() int32`

GetExpectedVsmartConnections returns the ExpectedVsmartConnections field if non-nil, zero value otherwise.

### GetExpectedVsmartConnectionsOk

`func (o *DeviceHealthDetails) GetExpectedVsmartConnectionsOk() (*int32, bool)`

GetExpectedVsmartConnectionsOk returns a tuple with the ExpectedVsmartConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVsmartConnections

`func (o *DeviceHealthDetails) SetExpectedVsmartConnections(v int32)`

SetExpectedVsmartConnections sets ExpectedVsmartConnections field to given value.

### HasExpectedVsmartConnections

`func (o *DeviceHealthDetails) HasExpectedVsmartConnections() bool`

HasExpectedVsmartConnections returns a boolean if a field has been set.

### GetHasGeoData

`func (o *DeviceHealthDetails) GetHasGeoData() bool`

GetHasGeoData returns the HasGeoData field if non-nil, zero value otherwise.

### GetHasGeoDataOk

`func (o *DeviceHealthDetails) GetHasGeoDataOk() (*bool, bool)`

GetHasGeoDataOk returns a tuple with the HasGeoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGeoData

`func (o *DeviceHealthDetails) SetHasGeoData(v bool)`

SetHasGeoData sets HasGeoData field to given value.

### HasHasGeoData

`func (o *DeviceHealthDetails) HasHasGeoData() bool`

HasHasGeoData returns a boolean if a field has been set.

### GetHealth

`func (o *DeviceHealthDetails) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DeviceHealthDetails) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DeviceHealthDetails) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DeviceHealthDetails) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLatitude

`func (o *DeviceHealthDetails) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DeviceHealthDetails) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DeviceHealthDetails) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *DeviceHealthDetails) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLocalSystemIp

`func (o *DeviceHealthDetails) GetLocalSystemIp() string`

GetLocalSystemIp returns the LocalSystemIp field if non-nil, zero value otherwise.

### GetLocalSystemIpOk

`func (o *DeviceHealthDetails) GetLocalSystemIpOk() (*string, bool)`

GetLocalSystemIpOk returns a tuple with the LocalSystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSystemIp

`func (o *DeviceHealthDetails) SetLocalSystemIp(v string)`

SetLocalSystemIp sets LocalSystemIp field to given value.

### HasLocalSystemIp

`func (o *DeviceHealthDetails) HasLocalSystemIp() bool`

HasLocalSystemIp returns a boolean if a field has been set.

### GetLocation

`func (o *DeviceHealthDetails) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeviceHealthDetails) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeviceHealthDetails) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeviceHealthDetails) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLongitude

`func (o *DeviceHealthDetails) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DeviceHealthDetails) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DeviceHealthDetails) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *DeviceHealthDetails) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMemoryUtilization

`func (o *DeviceHealthDetails) GetMemoryUtilization() float64`

GetMemoryUtilization returns the MemoryUtilization field if non-nil, zero value otherwise.

### GetMemoryUtilizationOk

`func (o *DeviceHealthDetails) GetMemoryUtilizationOk() (*float64, bool)`

GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtilization

`func (o *DeviceHealthDetails) SetMemoryUtilization(v float64)`

SetMemoryUtilization sets MemoryUtilization field to given value.

### HasMemoryUtilization

`func (o *DeviceHealthDetails) HasMemoryUtilization() bool`

HasMemoryUtilization returns a boolean if a field has been set.

### GetName

`func (o *DeviceHealthDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceHealthDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceHealthDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceHealthDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmpPeers

`func (o *DeviceHealthDetails) GetOmpPeers() int64`

GetOmpPeers returns the OmpPeers field if non-nil, zero value otherwise.

### GetOmpPeersOk

`func (o *DeviceHealthDetails) GetOmpPeersOk() (*int64, bool)`

GetOmpPeersOk returns a tuple with the OmpPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmpPeers

`func (o *DeviceHealthDetails) SetOmpPeers(v int64)`

SetOmpPeers sets OmpPeers field to given value.

### HasOmpPeers

`func (o *DeviceHealthDetails) HasOmpPeers() bool`

HasOmpPeers returns a boolean if a field has been set.

### GetOmpPeersUp

`func (o *DeviceHealthDetails) GetOmpPeersUp() int64`

GetOmpPeersUp returns the OmpPeersUp field if non-nil, zero value otherwise.

### GetOmpPeersUpOk

`func (o *DeviceHealthDetails) GetOmpPeersUpOk() (*int64, bool)`

GetOmpPeersUpOk returns a tuple with the OmpPeersUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmpPeersUp

`func (o *DeviceHealthDetails) SetOmpPeersUp(v int64)`

SetOmpPeersUp sets OmpPeersUp field to given value.

### HasOmpPeersUp

`func (o *DeviceHealthDetails) HasOmpPeersUp() bool`

HasOmpPeersUp returns a boolean if a field has been set.

### GetPersonality

`func (o *DeviceHealthDetails) GetPersonality() string`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *DeviceHealthDetails) GetPersonalityOk() (*string, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *DeviceHealthDetails) SetPersonality(v string)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *DeviceHealthDetails) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.

### GetQoe

`func (o *DeviceHealthDetails) GetQoe() int32`

GetQoe returns the Qoe field if non-nil, zero value otherwise.

### GetQoeOk

`func (o *DeviceHealthDetails) GetQoeOk() (*int32, bool)`

GetQoeOk returns a tuple with the Qoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoe

`func (o *DeviceHealthDetails) SetQoe(v int32)`

SetQoe sets Qoe field to given value.

### HasQoe

`func (o *DeviceHealthDetails) HasQoe() bool`

HasQoe returns a boolean if a field has been set.

### GetReachability

`func (o *DeviceHealthDetails) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *DeviceHealthDetails) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *DeviceHealthDetails) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *DeviceHealthDetails) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetSiteId

`func (o *DeviceHealthDetails) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceHealthDetails) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceHealthDetails) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DeviceHealthDetails) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *DeviceHealthDetails) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *DeviceHealthDetails) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *DeviceHealthDetails) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *DeviceHealthDetails) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSystemIp

`func (o *DeviceHealthDetails) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *DeviceHealthDetails) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *DeviceHealthDetails) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *DeviceHealthDetails) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetUptimeDate

`func (o *DeviceHealthDetails) GetUptimeDate() int64`

GetUptimeDate returns the UptimeDate field if non-nil, zero value otherwise.

### GetUptimeDateOk

`func (o *DeviceHealthDetails) GetUptimeDateOk() (*int64, bool)`

GetUptimeDateOk returns a tuple with the UptimeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeDate

`func (o *DeviceHealthDetails) SetUptimeDate(v int64)`

SetUptimeDate sets UptimeDate field to given value.

### HasUptimeDate

`func (o *DeviceHealthDetails) HasUptimeDate() bool`

HasUptimeDate returns a boolean if a field has been set.

### GetUuid

`func (o *DeviceHealthDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceHealthDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceHealthDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceHealthDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVpnIds

`func (o *DeviceHealthDetails) GetVpnIds() []string`

GetVpnIds returns the VpnIds field if non-nil, zero value otherwise.

### GetVpnIdsOk

`func (o *DeviceHealthDetails) GetVpnIdsOk() (*[]string, bool)`

GetVpnIdsOk returns a tuple with the VpnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnIds

`func (o *DeviceHealthDetails) SetVpnIds(v []string)`

SetVpnIds sets VpnIds field to given value.

### HasVpnIds

`func (o *DeviceHealthDetails) HasVpnIds() bool`

HasVpnIds returns a boolean if a field has been set.

### GetVsmartControlConnections

`func (o *DeviceHealthDetails) GetVsmartControlConnections() int32`

GetVsmartControlConnections returns the VsmartControlConnections field if non-nil, zero value otherwise.

### GetVsmartControlConnectionsOk

`func (o *DeviceHealthDetails) GetVsmartControlConnectionsOk() (*int32, bool)`

GetVsmartControlConnectionsOk returns a tuple with the VsmartControlConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmartControlConnections

`func (o *DeviceHealthDetails) SetVsmartControlConnections(v int32)`

SetVsmartControlConnections sets VsmartControlConnections field to given value.

### HasVsmartControlConnections

`func (o *DeviceHealthDetails) HasVsmartControlConnections() bool`

HasVsmartControlConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


