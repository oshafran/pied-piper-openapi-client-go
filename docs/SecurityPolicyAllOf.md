# SecurityPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**PolicyRules** | Pointer to [**[]PolicyRule**](PolicyRule.md) |  | [optional] 

## Methods

### NewSecurityPolicyAllOf

`func NewSecurityPolicyAllOf() *SecurityPolicyAllOf`

NewSecurityPolicyAllOf instantiates a new SecurityPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPolicyAllOfWithDefaults

`func NewSecurityPolicyAllOfWithDefaults() *SecurityPolicyAllOf`

NewSecurityPolicyAllOfWithDefaults instantiates a new SecurityPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityPolicyAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityPolicyAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityPolicyAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityPolicyAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SecurityPolicyAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityPolicyAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityPolicyAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityPolicyAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicyName

`func (o *SecurityPolicyAllOf) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *SecurityPolicyAllOf) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *SecurityPolicyAllOf) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *SecurityPolicyAllOf) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetDefaultAction

`func (o *SecurityPolicyAllOf) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *SecurityPolicyAllOf) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *SecurityPolicyAllOf) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *SecurityPolicyAllOf) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetPolicyRules

`func (o *SecurityPolicyAllOf) GetPolicyRules() []PolicyRule`

GetPolicyRules returns the PolicyRules field if non-nil, zero value otherwise.

### GetPolicyRulesOk

`func (o *SecurityPolicyAllOf) GetPolicyRulesOk() (*[]PolicyRule, bool)`

GetPolicyRulesOk returns a tuple with the PolicyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRules

`func (o *SecurityPolicyAllOf) SetPolicyRules(v []PolicyRule)`

SetPolicyRules sets PolicyRules field to given value.

### HasPolicyRules

`func (o *SecurityPolicyAllOf) HasPolicyRules() bool`

HasPolicyRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


