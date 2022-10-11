# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**CloudTenantId** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureCredentials

`func NewAzureCredentials() *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureCredentials) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCloudTenantId

`func (o *AzureCredentials) GetCloudTenantId() string`

GetCloudTenantId returns the CloudTenantId field if non-nil, zero value otherwise.

### GetCloudTenantIdOk

`func (o *AzureCredentials) GetCloudTenantIdOk() (*string, bool)`

GetCloudTenantIdOk returns a tuple with the CloudTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTenantId

`func (o *AzureCredentials) SetCloudTenantId(v string)`

SetCloudTenantId sets CloudTenantId field to given value.

### HasCloudTenantId

`func (o *AzureCredentials) HasCloudTenantId() bool`

HasCloudTenantId returns a boolean if a field has been set.

### GetSecretKey

`func (o *AzureCredentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AzureCredentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AzureCredentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AzureCredentials) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureCredentials) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureCredentials) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureCredentials) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureCredentials) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


