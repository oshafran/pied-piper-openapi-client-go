# SecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**PolicyName** | Pointer to **string** |  | [optional] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**PolicyRules** | Pointer to [**[]PolicyRule**](PolicyRule.md) |  | [optional] 

## Methods

### NewSecurityPolicy

`func NewSecurityPolicy(name string, type_ string, ) *SecurityPolicy`

NewSecurityPolicy instantiates a new SecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPolicyWithDefaults

`func NewSecurityPolicyWithDefaults() *SecurityPolicy`

NewSecurityPolicyWithDefaults instantiates a new SecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SecurityPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetPolicyName

`func (o *SecurityPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *SecurityPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *SecurityPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *SecurityPolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetDefaultAction

`func (o *SecurityPolicy) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *SecurityPolicy) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *SecurityPolicy) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *SecurityPolicy) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetPolicyRules

`func (o *SecurityPolicy) GetPolicyRules() []PolicyRule`

GetPolicyRules returns the PolicyRules field if non-nil, zero value otherwise.

### GetPolicyRulesOk

`func (o *SecurityPolicy) GetPolicyRulesOk() (*[]PolicyRule, bool)`

GetPolicyRulesOk returns a tuple with the PolicyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRules

`func (o *SecurityPolicy) SetPolicyRules(v []PolicyRule)`

SetPolicyRules sets PolicyRules field to given value.

### HasPolicyRules

`func (o *SecurityPolicy) HasPolicyRules() bool`

HasPolicyRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


