# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**DestIp** | Pointer to **string** |  | [optional] 
**DestPort** | Pointer to **int32** |  | [optional] 
**ProtocolType** | Pointer to **[]string** |  | [optional] 
**SourceIp** | Pointer to **string** |  | [optional] 
**SourcePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicyRule

`func NewPolicyRule() *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDestIp

`func (o *PolicyRule) GetDestIp() string`

GetDestIp returns the DestIp field if non-nil, zero value otherwise.

### GetDestIpOk

`func (o *PolicyRule) GetDestIpOk() (*string, bool)`

GetDestIpOk returns a tuple with the DestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIp

`func (o *PolicyRule) SetDestIp(v string)`

SetDestIp sets DestIp field to given value.

### HasDestIp

`func (o *PolicyRule) HasDestIp() bool`

HasDestIp returns a boolean if a field has been set.

### GetDestPort

`func (o *PolicyRule) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *PolicyRule) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *PolicyRule) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *PolicyRule) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetProtocolType

`func (o *PolicyRule) GetProtocolType() []string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *PolicyRule) GetProtocolTypeOk() (*[]string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *PolicyRule) SetProtocolType(v []string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *PolicyRule) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetSourceIp

`func (o *PolicyRule) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *PolicyRule) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *PolicyRule) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *PolicyRule) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetSourcePort

`func (o *PolicyRule) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *PolicyRule) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *PolicyRule) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *PolicyRule) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


