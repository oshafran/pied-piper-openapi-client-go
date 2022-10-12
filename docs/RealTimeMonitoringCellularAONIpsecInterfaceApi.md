# \RealTimeMonitoringCellularAONIpsecInterfaceApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAONIpsecInterfaceCountersInfo**](RealTimeMonitoringCellularAONIpsecInterfaceApi.md#GetAONIpsecInterfaceCountersInfo) | **Get** /device/cellularEiolte/ipsec/interface/counters | 
[**GetAONIpsecInterfaceSessionnfo**](RealTimeMonitoringCellularAONIpsecInterfaceApi.md#GetAONIpsecInterfaceSessionnfo) | **Get** /device/cellularEiolte/ipsec/interface/session | 



## GetAONIpsecInterfaceCountersInfo

> map[string]interface{} GetAONIpsecInterfaceCountersInfo(ctx).DeviceId(deviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceCountersInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceCountersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAONIpsecInterfaceCountersInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceCountersInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAONIpsecInterfaceCountersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAONIpsecInterfaceSessionnfo

> map[string]interface{} GetAONIpsecInterfaceSessionnfo(ctx).DeviceId(deviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceSessionnfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceSessionnfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAONIpsecInterfaceSessionnfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCellularAONIpsecInterfaceApi.GetAONIpsecInterfaceSessionnfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAONIpsecInterfaceSessionnfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

