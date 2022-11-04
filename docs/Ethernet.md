# Ethernet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EthernetInterfaceList** | Pointer to [**[]EthernetInterface**](EthernetInterface.md) |  | [optional] 

## Methods

### NewEthernet

`func NewEthernet(name string, type_ string, ) *Ethernet`

NewEthernet instantiates a new Ethernet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetWithDefaults

`func NewEthernetWithDefaults() *Ethernet`

NewEthernetWithDefaults instantiates a new Ethernet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Ethernet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ethernet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ethernet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Ethernet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ethernet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ethernet) SetType(v string)`

SetType sets Type field to given value.


### GetEthernetInterfaceList

`func (o *Ethernet) GetEthernetInterfaceList() []EthernetInterface`

GetEthernetInterfaceList returns the EthernetInterfaceList field if non-nil, zero value otherwise.

### GetEthernetInterfaceListOk

`func (o *Ethernet) GetEthernetInterfaceListOk() (*[]EthernetInterface, bool)`

GetEthernetInterfaceListOk returns a tuple with the EthernetInterfaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetInterfaceList

`func (o *Ethernet) SetEthernetInterfaceList(v []EthernetInterface)`

SetEthernetInterfaceList sets EthernetInterfaceList field to given value.

### HasEthernetInterfaceList

`func (o *Ethernet) HasEthernetInterfaceList() bool`

HasEthernetInterfaceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


