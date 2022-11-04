/*
Cisco SD-WAN vManage API

The vManage API exposes the functionality of operations maintaining devices and the overlay network

API version: 2.0.0
Contact: vmanage@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DeviceHealthDetails struct for DeviceHealthDetails
type DeviceHealthDetails struct {
	BfdSessions *int32 `json:"bfd_sessions,omitempty"`
	BfdSessionsUp *int32 `json:"bfd_sessions_up,omitempty"`
	BoardSerialNumber *string `json:"board_serial_number,omitempty"`
	ChassisNumber *string `json:"chassis_number,omitempty"`
	ConnectedVmanages []string `json:"connected_vmanages,omitempty"`
	ControlConnectionsToVsmat *DeviceHealthDetails `json:"controlConnectionsToVsmat,omitempty"`
	ControlConnections *int32 `json:"control_connections,omitempty"`
	ControlConnectionsUp *int32 `json:"control_connections_up,omitempty"`
	CpuLoad *float64 `json:"cpu_load,omitempty"`
	DeviceGroups []string `json:"device_groups,omitempty"`
	DeviceModel *string `json:"device_model,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	ExpectedVsmartConnections *int32 `json:"expected_vsmart_connections,omitempty"`
	HasGeoData *bool `json:"has_geo_data,omitempty"`
	Health *string `json:"health,omitempty"`
	Latitude *string `json:"latitude,omitempty"`
	LocalSystemIp *string `json:"local_system_ip,omitempty"`
	Location *string `json:"location,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	MemoryUtilization *float64 `json:"memory_utilization,omitempty"`
	Name *string `json:"name,omitempty"`
	OmpPeers *int64 `json:"omp_peers,omitempty"`
	OmpPeersUp *int64 `json:"omp_peers_up,omitempty"`
	Personality *string `json:"personality,omitempty"`
	Qoe *int32 `json:"qoe,omitempty"`
	Reachability *string `json:"reachability,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SoftwareVersion *string `json:"software_version,omitempty"`
	SystemIp *string `json:"system_ip,omitempty"`
	UptimeDate *int64 `json:"uptime_date,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	VpnIds []string `json:"vpn_ids,omitempty"`
	VsmartControlConnections *int32 `json:"vsmart_control_connections,omitempty"`
}

// NewDeviceHealthDetails instantiates a new DeviceHealthDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceHealthDetails() *DeviceHealthDetails {
	this := DeviceHealthDetails{}
	return &this
}

// NewDeviceHealthDetailsWithDefaults instantiates a new DeviceHealthDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceHealthDetailsWithDefaults() *DeviceHealthDetails {
	this := DeviceHealthDetails{}
	return &this
}

// GetBfdSessions returns the BfdSessions field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetBfdSessions() int32 {
	if o == nil || isNil(o.BfdSessions) {
		var ret int32
		return ret
	}
	return *o.BfdSessions
}

// GetBfdSessionsOk returns a tuple with the BfdSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetBfdSessionsOk() (*int32, bool) {
	if o == nil || isNil(o.BfdSessions) {
    return nil, false
	}
	return o.BfdSessions, true
}

// HasBfdSessions returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasBfdSessions() bool {
	if o != nil && !isNil(o.BfdSessions) {
		return true
	}

	return false
}

// SetBfdSessions gets a reference to the given int32 and assigns it to the BfdSessions field.
func (o *DeviceHealthDetails) SetBfdSessions(v int32) {
	o.BfdSessions = &v
}

// GetBfdSessionsUp returns the BfdSessionsUp field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetBfdSessionsUp() int32 {
	if o == nil || isNil(o.BfdSessionsUp) {
		var ret int32
		return ret
	}
	return *o.BfdSessionsUp
}

// GetBfdSessionsUpOk returns a tuple with the BfdSessionsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetBfdSessionsUpOk() (*int32, bool) {
	if o == nil || isNil(o.BfdSessionsUp) {
    return nil, false
	}
	return o.BfdSessionsUp, true
}

// HasBfdSessionsUp returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasBfdSessionsUp() bool {
	if o != nil && !isNil(o.BfdSessionsUp) {
		return true
	}

	return false
}

// SetBfdSessionsUp gets a reference to the given int32 and assigns it to the BfdSessionsUp field.
func (o *DeviceHealthDetails) SetBfdSessionsUp(v int32) {
	o.BfdSessionsUp = &v
}

// GetBoardSerialNumber returns the BoardSerialNumber field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetBoardSerialNumber() string {
	if o == nil || isNil(o.BoardSerialNumber) {
		var ret string
		return ret
	}
	return *o.BoardSerialNumber
}

// GetBoardSerialNumberOk returns a tuple with the BoardSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetBoardSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.BoardSerialNumber) {
    return nil, false
	}
	return o.BoardSerialNumber, true
}

// HasBoardSerialNumber returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasBoardSerialNumber() bool {
	if o != nil && !isNil(o.BoardSerialNumber) {
		return true
	}

	return false
}

// SetBoardSerialNumber gets a reference to the given string and assigns it to the BoardSerialNumber field.
func (o *DeviceHealthDetails) SetBoardSerialNumber(v string) {
	o.BoardSerialNumber = &v
}

// GetChassisNumber returns the ChassisNumber field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetChassisNumber() string {
	if o == nil || isNil(o.ChassisNumber) {
		var ret string
		return ret
	}
	return *o.ChassisNumber
}

// GetChassisNumberOk returns a tuple with the ChassisNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetChassisNumberOk() (*string, bool) {
	if o == nil || isNil(o.ChassisNumber) {
    return nil, false
	}
	return o.ChassisNumber, true
}

// HasChassisNumber returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasChassisNumber() bool {
	if o != nil && !isNil(o.ChassisNumber) {
		return true
	}

	return false
}

// SetChassisNumber gets a reference to the given string and assigns it to the ChassisNumber field.
func (o *DeviceHealthDetails) SetChassisNumber(v string) {
	o.ChassisNumber = &v
}

// GetConnectedVmanages returns the ConnectedVmanages field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetConnectedVmanages() []string {
	if o == nil || isNil(o.ConnectedVmanages) {
		var ret []string
		return ret
	}
	return o.ConnectedVmanages
}

// GetConnectedVmanagesOk returns a tuple with the ConnectedVmanages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetConnectedVmanagesOk() ([]string, bool) {
	if o == nil || isNil(o.ConnectedVmanages) {
    return nil, false
	}
	return o.ConnectedVmanages, true
}

// HasConnectedVmanages returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasConnectedVmanages() bool {
	if o != nil && !isNil(o.ConnectedVmanages) {
		return true
	}

	return false
}

// SetConnectedVmanages gets a reference to the given []string and assigns it to the ConnectedVmanages field.
func (o *DeviceHealthDetails) SetConnectedVmanages(v []string) {
	o.ConnectedVmanages = v
}

// GetControlConnectionsToVsmat returns the ControlConnectionsToVsmat field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetControlConnectionsToVsmat() DeviceHealthDetails {
	if o == nil || isNil(o.ControlConnectionsToVsmat) {
		var ret DeviceHealthDetails
		return ret
	}
	return *o.ControlConnectionsToVsmat
}

// GetControlConnectionsToVsmatOk returns a tuple with the ControlConnectionsToVsmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetControlConnectionsToVsmatOk() (*DeviceHealthDetails, bool) {
	if o == nil || isNil(o.ControlConnectionsToVsmat) {
    return nil, false
	}
	return o.ControlConnectionsToVsmat, true
}

// HasControlConnectionsToVsmat returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasControlConnectionsToVsmat() bool {
	if o != nil && !isNil(o.ControlConnectionsToVsmat) {
		return true
	}

	return false
}

// SetControlConnectionsToVsmat gets a reference to the given DeviceHealthDetails and assigns it to the ControlConnectionsToVsmat field.
func (o *DeviceHealthDetails) SetControlConnectionsToVsmat(v DeviceHealthDetails) {
	o.ControlConnectionsToVsmat = &v
}

// GetControlConnections returns the ControlConnections field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetControlConnections() int32 {
	if o == nil || isNil(o.ControlConnections) {
		var ret int32
		return ret
	}
	return *o.ControlConnections
}

// GetControlConnectionsOk returns a tuple with the ControlConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetControlConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.ControlConnections) {
    return nil, false
	}
	return o.ControlConnections, true
}

// HasControlConnections returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasControlConnections() bool {
	if o != nil && !isNil(o.ControlConnections) {
		return true
	}

	return false
}

// SetControlConnections gets a reference to the given int32 and assigns it to the ControlConnections field.
func (o *DeviceHealthDetails) SetControlConnections(v int32) {
	o.ControlConnections = &v
}

// GetControlConnectionsUp returns the ControlConnectionsUp field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetControlConnectionsUp() int32 {
	if o == nil || isNil(o.ControlConnectionsUp) {
		var ret int32
		return ret
	}
	return *o.ControlConnectionsUp
}

// GetControlConnectionsUpOk returns a tuple with the ControlConnectionsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetControlConnectionsUpOk() (*int32, bool) {
	if o == nil || isNil(o.ControlConnectionsUp) {
    return nil, false
	}
	return o.ControlConnectionsUp, true
}

// HasControlConnectionsUp returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasControlConnectionsUp() bool {
	if o != nil && !isNil(o.ControlConnectionsUp) {
		return true
	}

	return false
}

// SetControlConnectionsUp gets a reference to the given int32 and assigns it to the ControlConnectionsUp field.
func (o *DeviceHealthDetails) SetControlConnectionsUp(v int32) {
	o.ControlConnectionsUp = &v
}

// GetCpuLoad returns the CpuLoad field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetCpuLoad() float64 {
	if o == nil || isNil(o.CpuLoad) {
		var ret float64
		return ret
	}
	return *o.CpuLoad
}

// GetCpuLoadOk returns a tuple with the CpuLoad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetCpuLoadOk() (*float64, bool) {
	if o == nil || isNil(o.CpuLoad) {
    return nil, false
	}
	return o.CpuLoad, true
}

// HasCpuLoad returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasCpuLoad() bool {
	if o != nil && !isNil(o.CpuLoad) {
		return true
	}

	return false
}

// SetCpuLoad gets a reference to the given float64 and assigns it to the CpuLoad field.
func (o *DeviceHealthDetails) SetCpuLoad(v float64) {
	o.CpuLoad = &v
}

// GetDeviceGroups returns the DeviceGroups field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetDeviceGroups() []string {
	if o == nil || isNil(o.DeviceGroups) {
		var ret []string
		return ret
	}
	return o.DeviceGroups
}

// GetDeviceGroupsOk returns a tuple with the DeviceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetDeviceGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.DeviceGroups) {
    return nil, false
	}
	return o.DeviceGroups, true
}

// HasDeviceGroups returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasDeviceGroups() bool {
	if o != nil && !isNil(o.DeviceGroups) {
		return true
	}

	return false
}

// SetDeviceGroups gets a reference to the given []string and assigns it to the DeviceGroups field.
func (o *DeviceHealthDetails) SetDeviceGroups(v []string) {
	o.DeviceGroups = v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetDeviceModel() string {
	if o == nil || isNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetDeviceModelOk() (*string, bool) {
	if o == nil || isNil(o.DeviceModel) {
    return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasDeviceModel() bool {
	if o != nil && !isNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *DeviceHealthDetails) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetDeviceType() string {
	if o == nil || isNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetDeviceTypeOk() (*string, bool) {
	if o == nil || isNil(o.DeviceType) {
    return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasDeviceType() bool {
	if o != nil && !isNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *DeviceHealthDetails) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetExpectedVsmartConnections returns the ExpectedVsmartConnections field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetExpectedVsmartConnections() int32 {
	if o == nil || isNil(o.ExpectedVsmartConnections) {
		var ret int32
		return ret
	}
	return *o.ExpectedVsmartConnections
}

// GetExpectedVsmartConnectionsOk returns a tuple with the ExpectedVsmartConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetExpectedVsmartConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.ExpectedVsmartConnections) {
    return nil, false
	}
	return o.ExpectedVsmartConnections, true
}

// HasExpectedVsmartConnections returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasExpectedVsmartConnections() bool {
	if o != nil && !isNil(o.ExpectedVsmartConnections) {
		return true
	}

	return false
}

// SetExpectedVsmartConnections gets a reference to the given int32 and assigns it to the ExpectedVsmartConnections field.
func (o *DeviceHealthDetails) SetExpectedVsmartConnections(v int32) {
	o.ExpectedVsmartConnections = &v
}

// GetHasGeoData returns the HasGeoData field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetHasGeoData() bool {
	if o == nil || isNil(o.HasGeoData) {
		var ret bool
		return ret
	}
	return *o.HasGeoData
}

// GetHasGeoDataOk returns a tuple with the HasGeoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetHasGeoDataOk() (*bool, bool) {
	if o == nil || isNil(o.HasGeoData) {
    return nil, false
	}
	return o.HasGeoData, true
}

// HasHasGeoData returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasHasGeoData() bool {
	if o != nil && !isNil(o.HasGeoData) {
		return true
	}

	return false
}

// SetHasGeoData gets a reference to the given bool and assigns it to the HasGeoData field.
func (o *DeviceHealthDetails) SetHasGeoData(v bool) {
	o.HasGeoData = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetHealth() string {
	if o == nil || isNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetHealthOk() (*string, bool) {
	if o == nil || isNil(o.Health) {
    return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasHealth() bool {
	if o != nil && !isNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *DeviceHealthDetails) SetHealth(v string) {
	o.Health = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetLatitude() string {
	if o == nil || isNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetLatitudeOk() (*string, bool) {
	if o == nil || isNil(o.Latitude) {
    return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *DeviceHealthDetails) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLocalSystemIp returns the LocalSystemIp field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetLocalSystemIp() string {
	if o == nil || isNil(o.LocalSystemIp) {
		var ret string
		return ret
	}
	return *o.LocalSystemIp
}

// GetLocalSystemIpOk returns a tuple with the LocalSystemIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetLocalSystemIpOk() (*string, bool) {
	if o == nil || isNil(o.LocalSystemIp) {
    return nil, false
	}
	return o.LocalSystemIp, true
}

// HasLocalSystemIp returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasLocalSystemIp() bool {
	if o != nil && !isNil(o.LocalSystemIp) {
		return true
	}

	return false
}

// SetLocalSystemIp gets a reference to the given string and assigns it to the LocalSystemIp field.
func (o *DeviceHealthDetails) SetLocalSystemIp(v string) {
	o.LocalSystemIp = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DeviceHealthDetails) SetLocation(v string) {
	o.Location = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetLongitude() string {
	if o == nil || isNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetLongitudeOk() (*string, bool) {
	if o == nil || isNil(o.Longitude) {
    return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *DeviceHealthDetails) SetLongitude(v string) {
	o.Longitude = &v
}

// GetMemoryUtilization returns the MemoryUtilization field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetMemoryUtilization() float64 {
	if o == nil || isNil(o.MemoryUtilization) {
		var ret float64
		return ret
	}
	return *o.MemoryUtilization
}

// GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetMemoryUtilizationOk() (*float64, bool) {
	if o == nil || isNil(o.MemoryUtilization) {
    return nil, false
	}
	return o.MemoryUtilization, true
}

// HasMemoryUtilization returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasMemoryUtilization() bool {
	if o != nil && !isNil(o.MemoryUtilization) {
		return true
	}

	return false
}

// SetMemoryUtilization gets a reference to the given float64 and assigns it to the MemoryUtilization field.
func (o *DeviceHealthDetails) SetMemoryUtilization(v float64) {
	o.MemoryUtilization = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceHealthDetails) SetName(v string) {
	o.Name = &v
}

// GetOmpPeers returns the OmpPeers field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetOmpPeers() int64 {
	if o == nil || isNil(o.OmpPeers) {
		var ret int64
		return ret
	}
	return *o.OmpPeers
}

// GetOmpPeersOk returns a tuple with the OmpPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetOmpPeersOk() (*int64, bool) {
	if o == nil || isNil(o.OmpPeers) {
    return nil, false
	}
	return o.OmpPeers, true
}

// HasOmpPeers returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasOmpPeers() bool {
	if o != nil && !isNil(o.OmpPeers) {
		return true
	}

	return false
}

// SetOmpPeers gets a reference to the given int64 and assigns it to the OmpPeers field.
func (o *DeviceHealthDetails) SetOmpPeers(v int64) {
	o.OmpPeers = &v
}

// GetOmpPeersUp returns the OmpPeersUp field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetOmpPeersUp() int64 {
	if o == nil || isNil(o.OmpPeersUp) {
		var ret int64
		return ret
	}
	return *o.OmpPeersUp
}

// GetOmpPeersUpOk returns a tuple with the OmpPeersUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetOmpPeersUpOk() (*int64, bool) {
	if o == nil || isNil(o.OmpPeersUp) {
    return nil, false
	}
	return o.OmpPeersUp, true
}

// HasOmpPeersUp returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasOmpPeersUp() bool {
	if o != nil && !isNil(o.OmpPeersUp) {
		return true
	}

	return false
}

// SetOmpPeersUp gets a reference to the given int64 and assigns it to the OmpPeersUp field.
func (o *DeviceHealthDetails) SetOmpPeersUp(v int64) {
	o.OmpPeersUp = &v
}

// GetPersonality returns the Personality field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetPersonality() string {
	if o == nil || isNil(o.Personality) {
		var ret string
		return ret
	}
	return *o.Personality
}

// GetPersonalityOk returns a tuple with the Personality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetPersonalityOk() (*string, bool) {
	if o == nil || isNil(o.Personality) {
    return nil, false
	}
	return o.Personality, true
}

// HasPersonality returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasPersonality() bool {
	if o != nil && !isNil(o.Personality) {
		return true
	}

	return false
}

// SetPersonality gets a reference to the given string and assigns it to the Personality field.
func (o *DeviceHealthDetails) SetPersonality(v string) {
	o.Personality = &v
}

// GetQoe returns the Qoe field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetQoe() int32 {
	if o == nil || isNil(o.Qoe) {
		var ret int32
		return ret
	}
	return *o.Qoe
}

// GetQoeOk returns a tuple with the Qoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetQoeOk() (*int32, bool) {
	if o == nil || isNil(o.Qoe) {
    return nil, false
	}
	return o.Qoe, true
}

// HasQoe returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasQoe() bool {
	if o != nil && !isNil(o.Qoe) {
		return true
	}

	return false
}

// SetQoe gets a reference to the given int32 and assigns it to the Qoe field.
func (o *DeviceHealthDetails) SetQoe(v int32) {
	o.Qoe = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetReachability() string {
	if o == nil || isNil(o.Reachability) {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetReachabilityOk() (*string, bool) {
	if o == nil || isNil(o.Reachability) {
    return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasReachability() bool {
	if o != nil && !isNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *DeviceHealthDetails) SetReachability(v string) {
	o.Reachability = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetSiteId() string {
	if o == nil || isNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetSiteIdOk() (*string, bool) {
	if o == nil || isNil(o.SiteId) {
    return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasSiteId() bool {
	if o != nil && !isNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DeviceHealthDetails) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetSoftwareVersion() string {
	if o == nil || isNil(o.SoftwareVersion) {
		var ret string
		return ret
	}
	return *o.SoftwareVersion
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetSoftwareVersionOk() (*string, bool) {
	if o == nil || isNil(o.SoftwareVersion) {
    return nil, false
	}
	return o.SoftwareVersion, true
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasSoftwareVersion() bool {
	if o != nil && !isNil(o.SoftwareVersion) {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given string and assigns it to the SoftwareVersion field.
func (o *DeviceHealthDetails) SetSoftwareVersion(v string) {
	o.SoftwareVersion = &v
}

// GetSystemIp returns the SystemIp field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetSystemIp() string {
	if o == nil || isNil(o.SystemIp) {
		var ret string
		return ret
	}
	return *o.SystemIp
}

// GetSystemIpOk returns a tuple with the SystemIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetSystemIpOk() (*string, bool) {
	if o == nil || isNil(o.SystemIp) {
    return nil, false
	}
	return o.SystemIp, true
}

// HasSystemIp returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasSystemIp() bool {
	if o != nil && !isNil(o.SystemIp) {
		return true
	}

	return false
}

// SetSystemIp gets a reference to the given string and assigns it to the SystemIp field.
func (o *DeviceHealthDetails) SetSystemIp(v string) {
	o.SystemIp = &v
}

// GetUptimeDate returns the UptimeDate field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetUptimeDate() int64 {
	if o == nil || isNil(o.UptimeDate) {
		var ret int64
		return ret
	}
	return *o.UptimeDate
}

// GetUptimeDateOk returns a tuple with the UptimeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetUptimeDateOk() (*int64, bool) {
	if o == nil || isNil(o.UptimeDate) {
    return nil, false
	}
	return o.UptimeDate, true
}

// HasUptimeDate returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasUptimeDate() bool {
	if o != nil && !isNil(o.UptimeDate) {
		return true
	}

	return false
}

// SetUptimeDate gets a reference to the given int64 and assigns it to the UptimeDate field.
func (o *DeviceHealthDetails) SetUptimeDate(v int64) {
	o.UptimeDate = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeviceHealthDetails) SetUuid(v string) {
	o.Uuid = &v
}

// GetVpnIds returns the VpnIds field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetVpnIds() []string {
	if o == nil || isNil(o.VpnIds) {
		var ret []string
		return ret
	}
	return o.VpnIds
}

// GetVpnIdsOk returns a tuple with the VpnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetVpnIdsOk() ([]string, bool) {
	if o == nil || isNil(o.VpnIds) {
    return nil, false
	}
	return o.VpnIds, true
}

// HasVpnIds returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasVpnIds() bool {
	if o != nil && !isNil(o.VpnIds) {
		return true
	}

	return false
}

// SetVpnIds gets a reference to the given []string and assigns it to the VpnIds field.
func (o *DeviceHealthDetails) SetVpnIds(v []string) {
	o.VpnIds = v
}

// GetVsmartControlConnections returns the VsmartControlConnections field value if set, zero value otherwise.
func (o *DeviceHealthDetails) GetVsmartControlConnections() int32 {
	if o == nil || isNil(o.VsmartControlConnections) {
		var ret int32
		return ret
	}
	return *o.VsmartControlConnections
}

// GetVsmartControlConnectionsOk returns a tuple with the VsmartControlConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthDetails) GetVsmartControlConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.VsmartControlConnections) {
    return nil, false
	}
	return o.VsmartControlConnections, true
}

// HasVsmartControlConnections returns a boolean if a field has been set.
func (o *DeviceHealthDetails) HasVsmartControlConnections() bool {
	if o != nil && !isNil(o.VsmartControlConnections) {
		return true
	}

	return false
}

// SetVsmartControlConnections gets a reference to the given int32 and assigns it to the VsmartControlConnections field.
func (o *DeviceHealthDetails) SetVsmartControlConnections(v int32) {
	o.VsmartControlConnections = &v
}

func (o DeviceHealthDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BfdSessions) {
		toSerialize["bfd_sessions"] = o.BfdSessions
	}
	if !isNil(o.BfdSessionsUp) {
		toSerialize["bfd_sessions_up"] = o.BfdSessionsUp
	}
	if !isNil(o.BoardSerialNumber) {
		toSerialize["board_serial_number"] = o.BoardSerialNumber
	}
	if !isNil(o.ChassisNumber) {
		toSerialize["chassis_number"] = o.ChassisNumber
	}
	if !isNil(o.ConnectedVmanages) {
		toSerialize["connected_vmanages"] = o.ConnectedVmanages
	}
	if !isNil(o.ControlConnectionsToVsmat) {
		toSerialize["controlConnectionsToVsmat"] = o.ControlConnectionsToVsmat
	}
	if !isNil(o.ControlConnections) {
		toSerialize["control_connections"] = o.ControlConnections
	}
	if !isNil(o.ControlConnectionsUp) {
		toSerialize["control_connections_up"] = o.ControlConnectionsUp
	}
	if !isNil(o.CpuLoad) {
		toSerialize["cpu_load"] = o.CpuLoad
	}
	if !isNil(o.DeviceGroups) {
		toSerialize["device_groups"] = o.DeviceGroups
	}
	if !isNil(o.DeviceModel) {
		toSerialize["device_model"] = o.DeviceModel
	}
	if !isNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !isNil(o.ExpectedVsmartConnections) {
		toSerialize["expected_vsmart_connections"] = o.ExpectedVsmartConnections
	}
	if !isNil(o.HasGeoData) {
		toSerialize["has_geo_data"] = o.HasGeoData
	}
	if !isNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !isNil(o.LocalSystemIp) {
		toSerialize["local_system_ip"] = o.LocalSystemIp
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !isNil(o.MemoryUtilization) {
		toSerialize["memory_utilization"] = o.MemoryUtilization
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OmpPeers) {
		toSerialize["omp_peers"] = o.OmpPeers
	}
	if !isNil(o.OmpPeersUp) {
		toSerialize["omp_peers_up"] = o.OmpPeersUp
	}
	if !isNil(o.Personality) {
		toSerialize["personality"] = o.Personality
	}
	if !isNil(o.Qoe) {
		toSerialize["qoe"] = o.Qoe
	}
	if !isNil(o.Reachability) {
		toSerialize["reachability"] = o.Reachability
	}
	if !isNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !isNil(o.SoftwareVersion) {
		toSerialize["software_version"] = o.SoftwareVersion
	}
	if !isNil(o.SystemIp) {
		toSerialize["system_ip"] = o.SystemIp
	}
	if !isNil(o.UptimeDate) {
		toSerialize["uptime_date"] = o.UptimeDate
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.VpnIds) {
		toSerialize["vpn_ids"] = o.VpnIds
	}
	if !isNil(o.VsmartControlConnections) {
		toSerialize["vsmart_control_connections"] = o.VsmartControlConnections
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceHealthDetails struct {
	value *DeviceHealthDetails
	isSet bool
}

func (v NullableDeviceHealthDetails) Get() *DeviceHealthDetails {
	return v.value
}

func (v *NullableDeviceHealthDetails) Set(val *DeviceHealthDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceHealthDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceHealthDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceHealthDetails(val *DeviceHealthDetails) *NullableDeviceHealthDetails {
	return &NullableDeviceHealthDetails{value: val, isSet: true}
}

func (v NullableDeviceHealthDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceHealthDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


