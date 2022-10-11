# DHCPPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseTimeDay** | **int32** |  | 
**LeaseTimeHour** | **int32** |  | 
**LeaseTimeMin** | **int32** |  | 
**PoolNetwork** | **string** |  | 

## Methods

### NewDHCPPool

`func NewDHCPPool(leaseTimeDay int32, leaseTimeHour int32, leaseTimeMin int32, poolNetwork string, ) *DHCPPool`

NewDHCPPool instantiates a new DHCPPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPPoolWithDefaults

`func NewDHCPPoolWithDefaults() *DHCPPool`

NewDHCPPoolWithDefaults instantiates a new DHCPPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseTimeDay

`func (o *DHCPPool) GetLeaseTimeDay() int32`

GetLeaseTimeDay returns the LeaseTimeDay field if non-nil, zero value otherwise.

### GetLeaseTimeDayOk

`func (o *DHCPPool) GetLeaseTimeDayOk() (*int32, bool)`

GetLeaseTimeDayOk returns a tuple with the LeaseTimeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeDay

`func (o *DHCPPool) SetLeaseTimeDay(v int32)`

SetLeaseTimeDay sets LeaseTimeDay field to given value.


### GetLeaseTimeHour

`func (o *DHCPPool) GetLeaseTimeHour() int32`

GetLeaseTimeHour returns the LeaseTimeHour field if non-nil, zero value otherwise.

### GetLeaseTimeHourOk

`func (o *DHCPPool) GetLeaseTimeHourOk() (*int32, bool)`

GetLeaseTimeHourOk returns a tuple with the LeaseTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeHour

`func (o *DHCPPool) SetLeaseTimeHour(v int32)`

SetLeaseTimeHour sets LeaseTimeHour field to given value.


### GetLeaseTimeMin

`func (o *DHCPPool) GetLeaseTimeMin() int32`

GetLeaseTimeMin returns the LeaseTimeMin field if non-nil, zero value otherwise.

### GetLeaseTimeMinOk

`func (o *DHCPPool) GetLeaseTimeMinOk() (*int32, bool)`

GetLeaseTimeMinOk returns a tuple with the LeaseTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeMin

`func (o *DHCPPool) SetLeaseTimeMin(v int32)`

SetLeaseTimeMin sets LeaseTimeMin field to given value.


### GetPoolNetwork

`func (o *DHCPPool) GetPoolNetwork() string`

GetPoolNetwork returns the PoolNetwork field if non-nil, zero value otherwise.

### GetPoolNetworkOk

`func (o *DHCPPool) GetPoolNetworkOk() (*string, bool)`

GetPoolNetworkOk returns a tuple with the PoolNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolNetwork

`func (o *DHCPPool) SetPoolNetwork(v string)`

SetPoolNetwork sets PoolNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


