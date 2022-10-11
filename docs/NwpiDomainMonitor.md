# NwpiDomainMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**DeviceToDomainId** | Pointer to [**[]UuidToDomainId**](UuidToDomainId.md) |  | [optional] 
**DomainAppGrp** | Pointer to **string** |  | [optional] 
**DomainAppVis** | Pointer to **string** |  | [optional] 
**DomainList** | Pointer to [**[]DomainDetail**](DomainDetail.md) |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 

## Methods

### NewNwpiDomainMonitor

`func NewNwpiDomainMonitor() *NwpiDomainMonitor`

NewNwpiDomainMonitor instantiates a new NwpiDomainMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwpiDomainMonitorWithDefaults

`func NewNwpiDomainMonitorWithDefaults() *NwpiDomainMonitor`

NewNwpiDomainMonitorWithDefaults instantiates a new NwpiDomainMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *NwpiDomainMonitor) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *NwpiDomainMonitor) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *NwpiDomainMonitor) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *NwpiDomainMonitor) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetDeviceToDomainId

`func (o *NwpiDomainMonitor) GetDeviceToDomainId() []UuidToDomainId`

GetDeviceToDomainId returns the DeviceToDomainId field if non-nil, zero value otherwise.

### GetDeviceToDomainIdOk

`func (o *NwpiDomainMonitor) GetDeviceToDomainIdOk() (*[]UuidToDomainId, bool)`

GetDeviceToDomainIdOk returns a tuple with the DeviceToDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceToDomainId

`func (o *NwpiDomainMonitor) SetDeviceToDomainId(v []UuidToDomainId)`

SetDeviceToDomainId sets DeviceToDomainId field to given value.

### HasDeviceToDomainId

`func (o *NwpiDomainMonitor) HasDeviceToDomainId() bool`

HasDeviceToDomainId returns a boolean if a field has been set.

### GetDomainAppGrp

`func (o *NwpiDomainMonitor) GetDomainAppGrp() string`

GetDomainAppGrp returns the DomainAppGrp field if non-nil, zero value otherwise.

### GetDomainAppGrpOk

`func (o *NwpiDomainMonitor) GetDomainAppGrpOk() (*string, bool)`

GetDomainAppGrpOk returns a tuple with the DomainAppGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAppGrp

`func (o *NwpiDomainMonitor) SetDomainAppGrp(v string)`

SetDomainAppGrp sets DomainAppGrp field to given value.

### HasDomainAppGrp

`func (o *NwpiDomainMonitor) HasDomainAppGrp() bool`

HasDomainAppGrp returns a boolean if a field has been set.

### GetDomainAppVis

`func (o *NwpiDomainMonitor) GetDomainAppVis() string`

GetDomainAppVis returns the DomainAppVis field if non-nil, zero value otherwise.

### GetDomainAppVisOk

`func (o *NwpiDomainMonitor) GetDomainAppVisOk() (*string, bool)`

GetDomainAppVisOk returns a tuple with the DomainAppVis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAppVis

`func (o *NwpiDomainMonitor) SetDomainAppVis(v string)`

SetDomainAppVis sets DomainAppVis field to given value.

### HasDomainAppVis

`func (o *NwpiDomainMonitor) HasDomainAppVis() bool`

HasDomainAppVis returns a boolean if a field has been set.

### GetDomainList

`func (o *NwpiDomainMonitor) GetDomainList() []DomainDetail`

GetDomainList returns the DomainList field if non-nil, zero value otherwise.

### GetDomainListOk

`func (o *NwpiDomainMonitor) GetDomainListOk() (*[]DomainDetail, bool)`

GetDomainListOk returns a tuple with the DomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainList

`func (o *NwpiDomainMonitor) SetDomainList(v []DomainDetail)`

SetDomainList sets DomainList field to given value.

### HasDomainList

`func (o *NwpiDomainMonitor) HasDomainList() bool`

HasDomainList returns a boolean if a field has been set.

### GetTraceId

`func (o *NwpiDomainMonitor) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *NwpiDomainMonitor) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *NwpiDomainMonitor) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *NwpiDomainMonitor) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


