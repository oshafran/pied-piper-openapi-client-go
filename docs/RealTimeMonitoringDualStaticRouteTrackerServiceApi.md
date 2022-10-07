# \RealTimeMonitoringDualStaticRouteTrackerServiceApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDualStaticRouteTrackerInfo**](RealTimeMonitoringDualStaticRouteTrackerServiceApi.md#GetDualStaticRouteTrackerInfo) | **Get** /device/dualStaticRouteTracker | 



## GetDualStaticRouteTrackerInfo

> map[string]interface{} GetDualStaticRouteTrackerInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringDualStaticRouteTrackerServiceApi.GetDualStaticRouteTrackerInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDualStaticRouteTrackerServiceApi.GetDualStaticRouteTrackerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDualStaticRouteTrackerInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDualStaticRouteTrackerServiceApi.GetDualStaticRouteTrackerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDualStaticRouteTrackerInfoRequest struct via the builder pattern


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

