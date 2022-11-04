# NATRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**InPort** | **int32** |  | 
**InsideIp** | **string** |  | 
**Interface** | **string** |  | 
**OutPort** | **int32** |  | 
**Protocol** | **string** |  | 

## Methods

### NewNATRule

`func NewNATRule(description string, inPort int32, insideIp string, interface_ string, outPort int32, protocol string, ) *NATRule`

NewNATRule instantiates a new NATRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNATRuleWithDefaults

`func NewNATRuleWithDefaults() *NATRule`

NewNATRuleWithDefaults instantiates a new NATRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NATRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NATRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NATRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInPort

`func (o *NATRule) GetInPort() int32`

GetInPort returns the InPort field if non-nil, zero value otherwise.

### GetInPortOk

`func (o *NATRule) GetInPortOk() (*int32, bool)`

GetInPortOk returns a tuple with the InPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPort

`func (o *NATRule) SetInPort(v int32)`

SetInPort sets InPort field to given value.


### GetInsideIp

`func (o *NATRule) GetInsideIp() string`

GetInsideIp returns the InsideIp field if non-nil, zero value otherwise.

### GetInsideIpOk

`func (o *NATRule) GetInsideIpOk() (*string, bool)`

GetInsideIpOk returns a tuple with the InsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideIp

`func (o *NATRule) SetInsideIp(v string)`

SetInsideIp sets InsideIp field to given value.


### GetInterface

`func (o *NATRule) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NATRule) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NATRule) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetOutPort

`func (o *NATRule) GetOutPort() int32`

GetOutPort returns the OutPort field if non-nil, zero value otherwise.

### GetOutPortOk

`func (o *NATRule) GetOutPortOk() (*int32, bool)`

GetOutPortOk returns a tuple with the OutPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutPort

`func (o *NATRule) SetOutPort(v int32)`

SetOutPort sets OutPort field to given value.


### GetProtocol

`func (o *NATRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NATRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NATRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


