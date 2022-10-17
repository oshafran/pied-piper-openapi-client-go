# DataCenterRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterEncKey** | Pointer to **string** |  | [optional] 
**ClusterEnvKey** | Pointer to **string** |  | [optional] 
**DataCenters** | Pointer to [**[]DataCenter**](DataCenter.md) |  | [optional] 
**DisasterRecoverySettings** | Pointer to [**DisasterRecoverySettings**](DisasterRecoverySettings.md) |  | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Vbonds** | Pointer to [**[]Node**](Node.md) |  | [optional] 
**VmanageRootCA** | Pointer to **string** |  | [optional] 

## Methods

### NewDataCenterRegistration

`func NewDataCenterRegistration() *DataCenterRegistration`

NewDataCenterRegistration instantiates a new DataCenterRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCenterRegistrationWithDefaults

`func NewDataCenterRegistrationWithDefaults() *DataCenterRegistration`

NewDataCenterRegistrationWithDefaults instantiates a new DataCenterRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterEncKey

`func (o *DataCenterRegistration) GetClusterEncKey() string`

GetClusterEncKey returns the ClusterEncKey field if non-nil, zero value otherwise.

### GetClusterEncKeyOk

`func (o *DataCenterRegistration) GetClusterEncKeyOk() (*string, bool)`

GetClusterEncKeyOk returns a tuple with the ClusterEncKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEncKey

`func (o *DataCenterRegistration) SetClusterEncKey(v string)`

SetClusterEncKey sets ClusterEncKey field to given value.

### HasClusterEncKey

`func (o *DataCenterRegistration) HasClusterEncKey() bool`

HasClusterEncKey returns a boolean if a field has been set.

### GetClusterEnvKey

`func (o *DataCenterRegistration) GetClusterEnvKey() string`

GetClusterEnvKey returns the ClusterEnvKey field if non-nil, zero value otherwise.

### GetClusterEnvKeyOk

`func (o *DataCenterRegistration) GetClusterEnvKeyOk() (*string, bool)`

GetClusterEnvKeyOk returns a tuple with the ClusterEnvKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEnvKey

`func (o *DataCenterRegistration) SetClusterEnvKey(v string)`

SetClusterEnvKey sets ClusterEnvKey field to given value.

### HasClusterEnvKey

`func (o *DataCenterRegistration) HasClusterEnvKey() bool`

HasClusterEnvKey returns a boolean if a field has been set.

### GetDataCenters

`func (o *DataCenterRegistration) GetDataCenters() []DataCenter`

GetDataCenters returns the DataCenters field if non-nil, zero value otherwise.

### GetDataCentersOk

`func (o *DataCenterRegistration) GetDataCentersOk() (*[]DataCenter, bool)`

GetDataCentersOk returns a tuple with the DataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenters

`func (o *DataCenterRegistration) SetDataCenters(v []DataCenter)`

SetDataCenters sets DataCenters field to given value.

### HasDataCenters

`func (o *DataCenterRegistration) HasDataCenters() bool`

HasDataCenters returns a boolean if a field has been set.

### GetDisasterRecoverySettings

`func (o *DataCenterRegistration) GetDisasterRecoverySettings() DisasterRecoverySettings`

GetDisasterRecoverySettings returns the DisasterRecoverySettings field if non-nil, zero value otherwise.

### GetDisasterRecoverySettingsOk

`func (o *DataCenterRegistration) GetDisasterRecoverySettingsOk() (*DisasterRecoverySettings, bool)`

GetDisasterRecoverySettingsOk returns a tuple with the DisasterRecoverySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRecoverySettings

`func (o *DataCenterRegistration) SetDisasterRecoverySettings(v DisasterRecoverySettings)`

SetDisasterRecoverySettings sets DisasterRecoverySettings field to given value.

### HasDisasterRecoverySettings

`func (o *DataCenterRegistration) HasDisasterRecoverySettings() bool`

HasDisasterRecoverySettings returns a boolean if a field has been set.

### GetHost

`func (o *DataCenterRegistration) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DataCenterRegistration) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DataCenterRegistration) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *DataCenterRegistration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetVbonds

`func (o *DataCenterRegistration) GetVbonds() []Node`

GetVbonds returns the Vbonds field if non-nil, zero value otherwise.

### GetVbondsOk

`func (o *DataCenterRegistration) GetVbondsOk() (*[]Node, bool)`

GetVbondsOk returns a tuple with the Vbonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbonds

`func (o *DataCenterRegistration) SetVbonds(v []Node)`

SetVbonds sets Vbonds field to given value.

### HasVbonds

`func (o *DataCenterRegistration) HasVbonds() bool`

HasVbonds returns a boolean if a field has been set.

### GetVmanageRootCA

`func (o *DataCenterRegistration) GetVmanageRootCA() string`

GetVmanageRootCA returns the VmanageRootCA field if non-nil, zero value otherwise.

### GetVmanageRootCAOk

`func (o *DataCenterRegistration) GetVmanageRootCAOk() (*string, bool)`

GetVmanageRootCAOk returns a tuple with the VmanageRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmanageRootCA

`func (o *DataCenterRegistration) SetVmanageRootCA(v string)`

SetVmanageRootCA sets VmanageRootCA field to given value.

### HasVmanageRootCA

`func (o *DataCenterRegistration) HasVmanageRootCA() bool`

HasVmanageRootCA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


