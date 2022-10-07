# \RealTimeMonitoringIPv6FIBApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPv6FibList**](RealTimeMonitoringIPv6FIBApi.md#CreateIPv6FibList) | **Get** /device/ip/v6fib | 



## CreateIPv6FibList

> map[string]interface{} CreateIPv6FibList(ctx).DeviceId(deviceId).VpnId(vpnId).Prefix(prefix).Tloc(tloc).Color(color).Execute()





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
    vpnId := "vpnId_example" // string | VPN Id (optional)
    prefix := "prefix_example" // string | IP prefix (optional)
    tloc := "tloc_example" // string | tloc IP (optional)
    color := "color_example" // string | tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPv6FIBApi.CreateIPv6FibList(context.Background()).DeviceId(deviceId).VpnId(vpnId).Prefix(prefix).Tloc(tloc).Color(color).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPv6FIBApi.CreateIPv6FibList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPv6FibList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPv6FIBApi.CreateIPv6FibList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPv6FibListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **prefix** | **string** | IP prefix | 
 **tloc** | **string** | tloc IP | 
 **color** | **string** | tloc color | 

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

