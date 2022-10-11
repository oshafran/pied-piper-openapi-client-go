# \RealTimeMonitoringVPNApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVPNInstances**](RealTimeMonitoringVPNApi.md#GetVPNInstances) | **Get** /device/vpn | 



## GetVPNInstances

> map[string]interface{} GetVPNInstances(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringVPNApi.GetVPNInstances(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringVPNApi.GetVPNInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPNInstances`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringVPNApi.GetVPNInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVPNInstancesRequest struct via the builder pattern


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

