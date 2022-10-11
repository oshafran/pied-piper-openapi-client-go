# \RealTimeMonitoringGeoFencingApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeofenceStatus**](RealTimeMonitoringGeoFencingApi.md#GetGeofenceStatus) | **Get** /device/geofence/status | 



## GetGeofenceStatus

> map[string]interface{} GetGeofenceStatus(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringGeoFencingApi.GetGeofenceStatus(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringGeoFencingApi.GetGeofenceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeofenceStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringGeoFencingApi.GetGeofenceStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGeofenceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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

