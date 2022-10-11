# \RealTimeMonitoringStaticRouteTrackerServiceApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStaticRouteTrackerInfo**](RealTimeMonitoringStaticRouteTrackerServiceApi.md#GetStaticRouteTrackerInfo) | **Get** /device/staticRouteTracker | 



## GetStaticRouteTrackerInfo

> map[string]interface{} GetStaticRouteTrackerInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringStaticRouteTrackerServiceApi.GetStaticRouteTrackerInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringStaticRouteTrackerServiceApi.GetStaticRouteTrackerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaticRouteTrackerInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringStaticRouteTrackerServiceApi.GetStaticRouteTrackerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRouteTrackerInfoRequest struct via the builder pattern


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

