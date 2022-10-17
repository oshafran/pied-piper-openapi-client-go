# AwsKeyCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsKeyCredentials

`func NewAwsKeyCredentials() *AwsKeyCredentials`

NewAwsKeyCredentials instantiates a new AwsKeyCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKeyCredentialsWithDefaults

`func NewAwsKeyCredentialsWithDefaults() *AwsKeyCredentials`

NewAwsKeyCredentialsWithDefaults instantiates a new AwsKeyCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *AwsKeyCredentials) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AwsKeyCredentials) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AwsKeyCredentials) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AwsKeyCredentials) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AwsKeyCredentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsKeyCredentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsKeyCredentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AwsKeyCredentials) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


