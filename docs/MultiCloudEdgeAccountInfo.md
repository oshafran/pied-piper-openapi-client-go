# MultiCloudEdgeAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCredentials** | Pointer to **string** |  | [optional] 
**CredType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EdgeAccountId** | Pointer to **string** |  | [optional] 
**EdgeAccountName** | Pointer to **string** |  | [optional] 
**EdgeBillingAccountInfo** | Pointer to [**MultiCloudEdgeBillingAccountInfo**](MultiCloudEdgeBillingAccountInfo.md) |  | [optional] 
**EdgeLocationInfoList** | Pointer to [**[]MultiCloudEdgeLocationInfo**](MultiCloudEdgeLocationInfo.md) |  | [optional] 
**EdgePartnerPortsList** | Pointer to [**[]MultiCloudEdgePartnerPort**](MultiCloudEdgePartnerPort.md) |  | [optional] 
**EdgeType** | Pointer to **string** |  | [optional] 
**EquinixCredentials** | Pointer to [**EquinixCredentials**](EquinixCredentials.md) |  | [optional] 
**MegaportCredentials** | Pointer to [**MegaportCredentials**](MegaportCredentials.md) |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**RegionList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMultiCloudEdgeAccountInfo

`func NewMultiCloudEdgeAccountInfo() *MultiCloudEdgeAccountInfo`

NewMultiCloudEdgeAccountInfo instantiates a new MultiCloudEdgeAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudEdgeAccountInfoWithDefaults

`func NewMultiCloudEdgeAccountInfoWithDefaults() *MultiCloudEdgeAccountInfo`

NewMultiCloudEdgeAccountInfoWithDefaults instantiates a new MultiCloudEdgeAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCredentials

`func (o *MultiCloudEdgeAccountInfo) GetAccountCredentials() string`

GetAccountCredentials returns the AccountCredentials field if non-nil, zero value otherwise.

### GetAccountCredentialsOk

`func (o *MultiCloudEdgeAccountInfo) GetAccountCredentialsOk() (*string, bool)`

GetAccountCredentialsOk returns a tuple with the AccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCredentials

`func (o *MultiCloudEdgeAccountInfo) SetAccountCredentials(v string)`

SetAccountCredentials sets AccountCredentials field to given value.

### HasAccountCredentials

`func (o *MultiCloudEdgeAccountInfo) HasAccountCredentials() bool`

HasAccountCredentials returns a boolean if a field has been set.

### GetCredType

`func (o *MultiCloudEdgeAccountInfo) GetCredType() string`

GetCredType returns the CredType field if non-nil, zero value otherwise.

### GetCredTypeOk

`func (o *MultiCloudEdgeAccountInfo) GetCredTypeOk() (*string, bool)`

GetCredTypeOk returns a tuple with the CredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredType

`func (o *MultiCloudEdgeAccountInfo) SetCredType(v string)`

SetCredType sets CredType field to given value.

### HasCredType

`func (o *MultiCloudEdgeAccountInfo) HasCredType() bool`

HasCredType returns a boolean if a field has been set.

### GetDescription

`func (o *MultiCloudEdgeAccountInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiCloudEdgeAccountInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiCloudEdgeAccountInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiCloudEdgeAccountInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEdgeAccountId

`func (o *MultiCloudEdgeAccountInfo) GetEdgeAccountId() string`

GetEdgeAccountId returns the EdgeAccountId field if non-nil, zero value otherwise.

### GetEdgeAccountIdOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgeAccountIdOk() (*string, bool)`

GetEdgeAccountIdOk returns a tuple with the EdgeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccountId

`func (o *MultiCloudEdgeAccountInfo) SetEdgeAccountId(v string)`

SetEdgeAccountId sets EdgeAccountId field to given value.

### HasEdgeAccountId

`func (o *MultiCloudEdgeAccountInfo) HasEdgeAccountId() bool`

HasEdgeAccountId returns a boolean if a field has been set.

### GetEdgeAccountName

`func (o *MultiCloudEdgeAccountInfo) GetEdgeAccountName() string`

GetEdgeAccountName returns the EdgeAccountName field if non-nil, zero value otherwise.

### GetEdgeAccountNameOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgeAccountNameOk() (*string, bool)`

GetEdgeAccountNameOk returns a tuple with the EdgeAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccountName

`func (o *MultiCloudEdgeAccountInfo) SetEdgeAccountName(v string)`

SetEdgeAccountName sets EdgeAccountName field to given value.

### HasEdgeAccountName

`func (o *MultiCloudEdgeAccountInfo) HasEdgeAccountName() bool`

HasEdgeAccountName returns a boolean if a field has been set.

### GetEdgeBillingAccountInfo

`func (o *MultiCloudEdgeAccountInfo) GetEdgeBillingAccountInfo() MultiCloudEdgeBillingAccountInfo`

GetEdgeBillingAccountInfo returns the EdgeBillingAccountInfo field if non-nil, zero value otherwise.

### GetEdgeBillingAccountInfoOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgeBillingAccountInfoOk() (*MultiCloudEdgeBillingAccountInfo, bool)`

GetEdgeBillingAccountInfoOk returns a tuple with the EdgeBillingAccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeBillingAccountInfo

`func (o *MultiCloudEdgeAccountInfo) SetEdgeBillingAccountInfo(v MultiCloudEdgeBillingAccountInfo)`

SetEdgeBillingAccountInfo sets EdgeBillingAccountInfo field to given value.

### HasEdgeBillingAccountInfo

`func (o *MultiCloudEdgeAccountInfo) HasEdgeBillingAccountInfo() bool`

HasEdgeBillingAccountInfo returns a boolean if a field has been set.

### GetEdgeLocationInfoList

`func (o *MultiCloudEdgeAccountInfo) GetEdgeLocationInfoList() []MultiCloudEdgeLocationInfo`

GetEdgeLocationInfoList returns the EdgeLocationInfoList field if non-nil, zero value otherwise.

### GetEdgeLocationInfoListOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgeLocationInfoListOk() (*[]MultiCloudEdgeLocationInfo, bool)`

GetEdgeLocationInfoListOk returns a tuple with the EdgeLocationInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeLocationInfoList

`func (o *MultiCloudEdgeAccountInfo) SetEdgeLocationInfoList(v []MultiCloudEdgeLocationInfo)`

SetEdgeLocationInfoList sets EdgeLocationInfoList field to given value.

### HasEdgeLocationInfoList

`func (o *MultiCloudEdgeAccountInfo) HasEdgeLocationInfoList() bool`

HasEdgeLocationInfoList returns a boolean if a field has been set.

### GetEdgePartnerPortsList

`func (o *MultiCloudEdgeAccountInfo) GetEdgePartnerPortsList() []MultiCloudEdgePartnerPort`

GetEdgePartnerPortsList returns the EdgePartnerPortsList field if non-nil, zero value otherwise.

### GetEdgePartnerPortsListOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgePartnerPortsListOk() (*[]MultiCloudEdgePartnerPort, bool)`

GetEdgePartnerPortsListOk returns a tuple with the EdgePartnerPortsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgePartnerPortsList

`func (o *MultiCloudEdgeAccountInfo) SetEdgePartnerPortsList(v []MultiCloudEdgePartnerPort)`

SetEdgePartnerPortsList sets EdgePartnerPortsList field to given value.

### HasEdgePartnerPortsList

`func (o *MultiCloudEdgeAccountInfo) HasEdgePartnerPortsList() bool`

HasEdgePartnerPortsList returns a boolean if a field has been set.

### GetEdgeType

`func (o *MultiCloudEdgeAccountInfo) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *MultiCloudEdgeAccountInfo) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *MultiCloudEdgeAccountInfo) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *MultiCloudEdgeAccountInfo) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### GetEquinixCredentials

`func (o *MultiCloudEdgeAccountInfo) GetEquinixCredentials() EquinixCredentials`

GetEquinixCredentials returns the EquinixCredentials field if non-nil, zero value otherwise.

### GetEquinixCredentialsOk

`func (o *MultiCloudEdgeAccountInfo) GetEquinixCredentialsOk() (*EquinixCredentials, bool)`

GetEquinixCredentialsOk returns a tuple with the EquinixCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixCredentials

`func (o *MultiCloudEdgeAccountInfo) SetEquinixCredentials(v EquinixCredentials)`

SetEquinixCredentials sets EquinixCredentials field to given value.

### HasEquinixCredentials

`func (o *MultiCloudEdgeAccountInfo) HasEquinixCredentials() bool`

HasEquinixCredentials returns a boolean if a field has been set.

### GetMegaportCredentials

`func (o *MultiCloudEdgeAccountInfo) GetMegaportCredentials() MegaportCredentials`

GetMegaportCredentials returns the MegaportCredentials field if non-nil, zero value otherwise.

### GetMegaportCredentialsOk

`func (o *MultiCloudEdgeAccountInfo) GetMegaportCredentialsOk() (*MegaportCredentials, bool)`

GetMegaportCredentialsOk returns a tuple with the MegaportCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegaportCredentials

`func (o *MultiCloudEdgeAccountInfo) SetMegaportCredentials(v MegaportCredentials)`

SetMegaportCredentials sets MegaportCredentials field to given value.

### HasMegaportCredentials

`func (o *MultiCloudEdgeAccountInfo) HasMegaportCredentials() bool`

HasMegaportCredentials returns a boolean if a field has been set.

### GetOrgName

`func (o *MultiCloudEdgeAccountInfo) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MultiCloudEdgeAccountInfo) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MultiCloudEdgeAccountInfo) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MultiCloudEdgeAccountInfo) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetRegionList

`func (o *MultiCloudEdgeAccountInfo) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *MultiCloudEdgeAccountInfo) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *MultiCloudEdgeAccountInfo) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *MultiCloudEdgeAccountInfo) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


