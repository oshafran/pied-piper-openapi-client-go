# MultiCloudAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCredentials** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**AwsCloudGatewayWithTvpcEnabled** | Pointer to **bool** |  | [optional] 
**AwsIamCredentials** | Pointer to [**AwsIamCredentials**](AwsIamCredentials.md) |  | [optional] 
**AwsKeyCredentials** | Pointer to [**AwsKeyCredentials**](AwsKeyCredentials.md) |  | [optional] 
**AzO365Enabled** | Pointer to **bool** |  | [optional] 
**AzureCredentials** | Pointer to [**AzureCredentials**](AzureCredentials.md) |  | [optional] 
**ClientEmail** | Pointer to **string** |  | [optional] 
**CloudGatewayEnabled** | Pointer to **bool** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**CredType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GcpBillingId** | Pointer to **string** |  | [optional] 
**GcpCredentials** | Pointer to [**GcpCredentials**](GcpCredentials.md) |  | [optional] 
**HostVpcEnabled** | Pointer to **bool** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**PrivateKeyId** | Pointer to **string** |  | [optional] 
**RegionList** | Pointer to **[]string** |  | [optional] 
**ServiceDiscoveryEnabled** | Pointer to **bool** |  | [optional] 
**VnetEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMultiCloudAccountInfo

`func NewMultiCloudAccountInfo() *MultiCloudAccountInfo`

NewMultiCloudAccountInfo instantiates a new MultiCloudAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudAccountInfoWithDefaults

`func NewMultiCloudAccountInfoWithDefaults() *MultiCloudAccountInfo`

NewMultiCloudAccountInfoWithDefaults instantiates a new MultiCloudAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCredentials

`func (o *MultiCloudAccountInfo) GetAccountCredentials() string`

GetAccountCredentials returns the AccountCredentials field if non-nil, zero value otherwise.

### GetAccountCredentialsOk

`func (o *MultiCloudAccountInfo) GetAccountCredentialsOk() (*string, bool)`

GetAccountCredentialsOk returns a tuple with the AccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCredentials

`func (o *MultiCloudAccountInfo) SetAccountCredentials(v string)`

SetAccountCredentials sets AccountCredentials field to given value.

### HasAccountCredentials

`func (o *MultiCloudAccountInfo) HasAccountCredentials() bool`

HasAccountCredentials returns a boolean if a field has been set.

### GetAccountId

`func (o *MultiCloudAccountInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MultiCloudAccountInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MultiCloudAccountInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MultiCloudAccountInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *MultiCloudAccountInfo) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MultiCloudAccountInfo) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *MultiCloudAccountInfo) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *MultiCloudAccountInfo) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAwsCloudGatewayWithTvpcEnabled

`func (o *MultiCloudAccountInfo) GetAwsCloudGatewayWithTvpcEnabled() bool`

GetAwsCloudGatewayWithTvpcEnabled returns the AwsCloudGatewayWithTvpcEnabled field if non-nil, zero value otherwise.

### GetAwsCloudGatewayWithTvpcEnabledOk

`func (o *MultiCloudAccountInfo) GetAwsCloudGatewayWithTvpcEnabledOk() (*bool, bool)`

GetAwsCloudGatewayWithTvpcEnabledOk returns a tuple with the AwsCloudGatewayWithTvpcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudGatewayWithTvpcEnabled

`func (o *MultiCloudAccountInfo) SetAwsCloudGatewayWithTvpcEnabled(v bool)`

SetAwsCloudGatewayWithTvpcEnabled sets AwsCloudGatewayWithTvpcEnabled field to given value.

### HasAwsCloudGatewayWithTvpcEnabled

`func (o *MultiCloudAccountInfo) HasAwsCloudGatewayWithTvpcEnabled() bool`

HasAwsCloudGatewayWithTvpcEnabled returns a boolean if a field has been set.

### GetAwsIamCredentials

`func (o *MultiCloudAccountInfo) GetAwsIamCredentials() AwsIamCredentials`

GetAwsIamCredentials returns the AwsIamCredentials field if non-nil, zero value otherwise.

### GetAwsIamCredentialsOk

`func (o *MultiCloudAccountInfo) GetAwsIamCredentialsOk() (*AwsIamCredentials, bool)`

GetAwsIamCredentialsOk returns a tuple with the AwsIamCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamCredentials

`func (o *MultiCloudAccountInfo) SetAwsIamCredentials(v AwsIamCredentials)`

SetAwsIamCredentials sets AwsIamCredentials field to given value.

### HasAwsIamCredentials

`func (o *MultiCloudAccountInfo) HasAwsIamCredentials() bool`

HasAwsIamCredentials returns a boolean if a field has been set.

### GetAwsKeyCredentials

`func (o *MultiCloudAccountInfo) GetAwsKeyCredentials() AwsKeyCredentials`

GetAwsKeyCredentials returns the AwsKeyCredentials field if non-nil, zero value otherwise.

### GetAwsKeyCredentialsOk

`func (o *MultiCloudAccountInfo) GetAwsKeyCredentialsOk() (*AwsKeyCredentials, bool)`

GetAwsKeyCredentialsOk returns a tuple with the AwsKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyCredentials

`func (o *MultiCloudAccountInfo) SetAwsKeyCredentials(v AwsKeyCredentials)`

SetAwsKeyCredentials sets AwsKeyCredentials field to given value.

### HasAwsKeyCredentials

`func (o *MultiCloudAccountInfo) HasAwsKeyCredentials() bool`

HasAwsKeyCredentials returns a boolean if a field has been set.

### GetAzO365Enabled

`func (o *MultiCloudAccountInfo) GetAzO365Enabled() bool`

GetAzO365Enabled returns the AzO365Enabled field if non-nil, zero value otherwise.

### GetAzO365EnabledOk

`func (o *MultiCloudAccountInfo) GetAzO365EnabledOk() (*bool, bool)`

GetAzO365EnabledOk returns a tuple with the AzO365Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzO365Enabled

`func (o *MultiCloudAccountInfo) SetAzO365Enabled(v bool)`

SetAzO365Enabled sets AzO365Enabled field to given value.

### HasAzO365Enabled

`func (o *MultiCloudAccountInfo) HasAzO365Enabled() bool`

HasAzO365Enabled returns a boolean if a field has been set.

### GetAzureCredentials

`func (o *MultiCloudAccountInfo) GetAzureCredentials() AzureCredentials`

GetAzureCredentials returns the AzureCredentials field if non-nil, zero value otherwise.

### GetAzureCredentialsOk

`func (o *MultiCloudAccountInfo) GetAzureCredentialsOk() (*AzureCredentials, bool)`

GetAzureCredentialsOk returns a tuple with the AzureCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentials

`func (o *MultiCloudAccountInfo) SetAzureCredentials(v AzureCredentials)`

SetAzureCredentials sets AzureCredentials field to given value.

### HasAzureCredentials

`func (o *MultiCloudAccountInfo) HasAzureCredentials() bool`

HasAzureCredentials returns a boolean if a field has been set.

### GetClientEmail

`func (o *MultiCloudAccountInfo) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *MultiCloudAccountInfo) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *MultiCloudAccountInfo) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *MultiCloudAccountInfo) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetCloudGatewayEnabled

`func (o *MultiCloudAccountInfo) GetCloudGatewayEnabled() bool`

GetCloudGatewayEnabled returns the CloudGatewayEnabled field if non-nil, zero value otherwise.

### GetCloudGatewayEnabledOk

`func (o *MultiCloudAccountInfo) GetCloudGatewayEnabledOk() (*bool, bool)`

GetCloudGatewayEnabledOk returns a tuple with the CloudGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGatewayEnabled

`func (o *MultiCloudAccountInfo) SetCloudGatewayEnabled(v bool)`

SetCloudGatewayEnabled sets CloudGatewayEnabled field to given value.

### HasCloudGatewayEnabled

`func (o *MultiCloudAccountInfo) HasCloudGatewayEnabled() bool`

HasCloudGatewayEnabled returns a boolean if a field has been set.

### GetCloudType

`func (o *MultiCloudAccountInfo) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *MultiCloudAccountInfo) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *MultiCloudAccountInfo) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *MultiCloudAccountInfo) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCredType

`func (o *MultiCloudAccountInfo) GetCredType() string`

GetCredType returns the CredType field if non-nil, zero value otherwise.

### GetCredTypeOk

`func (o *MultiCloudAccountInfo) GetCredTypeOk() (*string, bool)`

GetCredTypeOk returns a tuple with the CredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredType

`func (o *MultiCloudAccountInfo) SetCredType(v string)`

SetCredType sets CredType field to given value.

### HasCredType

`func (o *MultiCloudAccountInfo) HasCredType() bool`

HasCredType returns a boolean if a field has been set.

### GetDescription

`func (o *MultiCloudAccountInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiCloudAccountInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiCloudAccountInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiCloudAccountInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcpBillingId

`func (o *MultiCloudAccountInfo) GetGcpBillingId() string`

GetGcpBillingId returns the GcpBillingId field if non-nil, zero value otherwise.

### GetGcpBillingIdOk

`func (o *MultiCloudAccountInfo) GetGcpBillingIdOk() (*string, bool)`

GetGcpBillingIdOk returns a tuple with the GcpBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBillingId

`func (o *MultiCloudAccountInfo) SetGcpBillingId(v string)`

SetGcpBillingId sets GcpBillingId field to given value.

### HasGcpBillingId

`func (o *MultiCloudAccountInfo) HasGcpBillingId() bool`

HasGcpBillingId returns a boolean if a field has been set.

### GetGcpCredentials

`func (o *MultiCloudAccountInfo) GetGcpCredentials() GcpCredentials`

GetGcpCredentials returns the GcpCredentials field if non-nil, zero value otherwise.

### GetGcpCredentialsOk

`func (o *MultiCloudAccountInfo) GetGcpCredentialsOk() (*GcpCredentials, bool)`

GetGcpCredentialsOk returns a tuple with the GcpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentials

`func (o *MultiCloudAccountInfo) SetGcpCredentials(v GcpCredentials)`

SetGcpCredentials sets GcpCredentials field to given value.

### HasGcpCredentials

`func (o *MultiCloudAccountInfo) HasGcpCredentials() bool`

HasGcpCredentials returns a boolean if a field has been set.

### GetHostVpcEnabled

`func (o *MultiCloudAccountInfo) GetHostVpcEnabled() bool`

GetHostVpcEnabled returns the HostVpcEnabled field if non-nil, zero value otherwise.

### GetHostVpcEnabledOk

`func (o *MultiCloudAccountInfo) GetHostVpcEnabledOk() (*bool, bool)`

GetHostVpcEnabledOk returns a tuple with the HostVpcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcEnabled

`func (o *MultiCloudAccountInfo) SetHostVpcEnabled(v bool)`

SetHostVpcEnabled sets HostVpcEnabled field to given value.

### HasHostVpcEnabled

`func (o *MultiCloudAccountInfo) HasHostVpcEnabled() bool`

HasHostVpcEnabled returns a boolean if a field has been set.

### GetOrgName

`func (o *MultiCloudAccountInfo) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MultiCloudAccountInfo) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MultiCloudAccountInfo) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MultiCloudAccountInfo) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPrivateKeyId

`func (o *MultiCloudAccountInfo) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *MultiCloudAccountInfo) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *MultiCloudAccountInfo) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *MultiCloudAccountInfo) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetRegionList

`func (o *MultiCloudAccountInfo) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *MultiCloudAccountInfo) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *MultiCloudAccountInfo) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *MultiCloudAccountInfo) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetServiceDiscoveryEnabled

`func (o *MultiCloudAccountInfo) GetServiceDiscoveryEnabled() bool`

GetServiceDiscoveryEnabled returns the ServiceDiscoveryEnabled field if non-nil, zero value otherwise.

### GetServiceDiscoveryEnabledOk

`func (o *MultiCloudAccountInfo) GetServiceDiscoveryEnabledOk() (*bool, bool)`

GetServiceDiscoveryEnabledOk returns a tuple with the ServiceDiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscoveryEnabled

`func (o *MultiCloudAccountInfo) SetServiceDiscoveryEnabled(v bool)`

SetServiceDiscoveryEnabled sets ServiceDiscoveryEnabled field to given value.

### HasServiceDiscoveryEnabled

`func (o *MultiCloudAccountInfo) HasServiceDiscoveryEnabled() bool`

HasServiceDiscoveryEnabled returns a boolean if a field has been set.

### GetVnetEnabled

`func (o *MultiCloudAccountInfo) GetVnetEnabled() bool`

GetVnetEnabled returns the VnetEnabled field if non-nil, zero value otherwise.

### GetVnetEnabledOk

`func (o *MultiCloudAccountInfo) GetVnetEnabledOk() (*bool, bool)`

GetVnetEnabledOk returns a tuple with the VnetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetEnabled

`func (o *MultiCloudAccountInfo) SetVnetEnabled(v bool)`

SetVnetEnabled sets VnetEnabled field to given value.

### HasVnetEnabled

`func (o *MultiCloudAccountInfo) HasVnetEnabled() bool`

HasVnetEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


