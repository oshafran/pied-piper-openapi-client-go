# DevicesHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]DeviceHealthDetails**](DeviceHealthDetails.md) |  | [optional] 
**Header** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TotalDevices** | Pointer to **int32** |  | [optional] 

## Methods

### NewDevicesHealth

`func NewDevicesHealth() *DevicesHealth`

NewDevicesHealth instantiates a new DevicesHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesHealthWithDefaults

`func NewDevicesHealthWithDefaults() *DevicesHealth`

NewDevicesHealthWithDefaults instantiates a new DevicesHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *DevicesHealth) GetDevices() []DeviceHealthDetails`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DevicesHealth) GetDevicesOk() (*[]DeviceHealthDetails, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DevicesHealth) SetDevices(v []DeviceHealthDetails)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DevicesHealth) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHeader

`func (o *DevicesHealth) GetHeader() map[string]map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *DevicesHealth) GetHeaderOk() (*map[string]map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *DevicesHealth) SetHeader(v map[string]map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *DevicesHealth) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTotalDevices

`func (o *DevicesHealth) GetTotalDevices() int32`

GetTotalDevices returns the TotalDevices field if non-nil, zero value otherwise.

### GetTotalDevicesOk

`func (o *DevicesHealth) GetTotalDevicesOk() (*int32, bool)`

GetTotalDevicesOk returns a tuple with the TotalDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevices

`func (o *DevicesHealth) SetTotalDevices(v int32)`

SetTotalDevices sets TotalDevices field to given value.

### HasTotalDevices

`func (o *DevicesHealth) HasTotalDevices() bool`

HasTotalDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


