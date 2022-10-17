# \RealTimeMonitoringCellularEIOLTEHardwareServiceApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEiolteHardwareInfo**](RealTimeMonitoringCellularEIOLTEHardwareServiceApi.md#GetEiolteHardwareInfo) | **Get** /device/cellularEiolte/hardware | 



## GetEiolteHardwareInfo

> map[string]interface{} GetEiolteHardwareInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringCellularEIOLTEHardwareServiceApi.GetEiolteHardwareInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCellularEIOLTEHardwareServiceApi.GetEiolteHardwareInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEiolteHardwareInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCellularEIOLTEHardwareServiceApi.GetEiolteHardwareInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEiolteHardwareInfoRequest struct via the builder pattern


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

