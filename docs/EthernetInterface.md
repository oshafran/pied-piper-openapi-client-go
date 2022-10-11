# EthernetInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** |  | [optional] 
**CorporateLan** | Pointer to **bool** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**IpAssignment** | Pointer to **string** |  | [optional] 
**PortType** | Pointer to **string** |  | [optional] 
**StaticIpAddress** | Pointer to **string** |  | [optional] 
**StaticIpAddressSubnetMask** | Pointer to **string** |  | [optional] 
**StaticRouteIp** | Pointer to **string** |  | [optional] 
**WanConfiguration** | Pointer to **string** |  | [optional] 

## Methods

### NewEthernetInterface

`func NewEthernetInterface() *EthernetInterface`

NewEthernetInterface instantiates a new EthernetInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetInterfaceWithDefaults

`func NewEthernetInterfaceWithDefaults() *EthernetInterface`

NewEthernetInterfaceWithDefaults instantiates a new EthernetInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *EthernetInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EthernetInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EthernetInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EthernetInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetCorporateLan

`func (o *EthernetInterface) GetCorporateLan() bool`

GetCorporateLan returns the CorporateLan field if non-nil, zero value otherwise.

### GetCorporateLanOk

`func (o *EthernetInterface) GetCorporateLanOk() (*bool, bool)`

GetCorporateLanOk returns a tuple with the CorporateLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateLan

`func (o *EthernetInterface) SetCorporateLan(v bool)`

SetCorporateLan sets CorporateLan field to given value.

### HasCorporateLan

`func (o *EthernetInterface) HasCorporateLan() bool`

HasCorporateLan returns a boolean if a field has been set.

### GetInterfaceName

`func (o *EthernetInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *EthernetInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *EthernetInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *EthernetInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpAssignment

`func (o *EthernetInterface) GetIpAssignment() string`

GetIpAssignment returns the IpAssignment field if non-nil, zero value otherwise.

### GetIpAssignmentOk

`func (o *EthernetInterface) GetIpAssignmentOk() (*string, bool)`

GetIpAssignmentOk returns a tuple with the IpAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignment

`func (o *EthernetInterface) SetIpAssignment(v string)`

SetIpAssignment sets IpAssignment field to given value.

### HasIpAssignment

`func (o *EthernetInterface) HasIpAssignment() bool`

HasIpAssignment returns a boolean if a field has been set.

### GetPortType

`func (o *EthernetInterface) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EthernetInterface) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EthernetInterface) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EthernetInterface) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetStaticIpAddress

`func (o *EthernetInterface) GetStaticIpAddress() string`

GetStaticIpAddress returns the StaticIpAddress field if non-nil, zero value otherwise.

### GetStaticIpAddressOk

`func (o *EthernetInterface) GetStaticIpAddressOk() (*string, bool)`

GetStaticIpAddressOk returns a tuple with the StaticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpAddress

`func (o *EthernetInterface) SetStaticIpAddress(v string)`

SetStaticIpAddress sets StaticIpAddress field to given value.

### HasStaticIpAddress

`func (o *EthernetInterface) HasStaticIpAddress() bool`

HasStaticIpAddress returns a boolean if a field has been set.

### GetStaticIpAddressSubnetMask

`func (o *EthernetInterface) GetStaticIpAddressSubnetMask() string`

GetStaticIpAddressSubnetMask returns the StaticIpAddressSubnetMask field if non-nil, zero value otherwise.

### GetStaticIpAddressSubnetMaskOk

`func (o *EthernetInterface) GetStaticIpAddressSubnetMaskOk() (*string, bool)`

GetStaticIpAddressSubnetMaskOk returns a tuple with the StaticIpAddressSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpAddressSubnetMask

`func (o *EthernetInterface) SetStaticIpAddressSubnetMask(v string)`

SetStaticIpAddressSubnetMask sets StaticIpAddressSubnetMask field to given value.

### HasStaticIpAddressSubnetMask

`func (o *EthernetInterface) HasStaticIpAddressSubnetMask() bool`

HasStaticIpAddressSubnetMask returns a boolean if a field has been set.

### GetStaticRouteIp

`func (o *EthernetInterface) GetStaticRouteIp() string`

GetStaticRouteIp returns the StaticRouteIp field if non-nil, zero value otherwise.

### GetStaticRouteIpOk

`func (o *EthernetInterface) GetStaticRouteIpOk() (*string, bool)`

GetStaticRouteIpOk returns a tuple with the StaticRouteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouteIp

`func (o *EthernetInterface) SetStaticRouteIp(v string)`

SetStaticRouteIp sets StaticRouteIp field to given value.

### HasStaticRouteIp

`func (o *EthernetInterface) HasStaticRouteIp() bool`

HasStaticRouteIp returns a boolean if a field has been set.

### GetWanConfiguration

`func (o *EthernetInterface) GetWanConfiguration() string`

GetWanConfiguration returns the WanConfiguration field if non-nil, zero value otherwise.

### GetWanConfigurationOk

`func (o *EthernetInterface) GetWanConfigurationOk() (*string, bool)`

GetWanConfigurationOk returns a tuple with the WanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanConfiguration

`func (o *EthernetInterface) SetWanConfiguration(v string)`

SetWanConfiguration sets WanConfiguration field to given value.

### HasWanConfiguration

`func (o *EthernetInterface) HasWanConfiguration() bool`

HasWanConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


