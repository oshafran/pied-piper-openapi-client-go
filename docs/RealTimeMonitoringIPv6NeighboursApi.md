# \RealTimeMonitoringIPv6NeighboursApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpv6Interface**](RealTimeMonitoringIPv6NeighboursApi.md#GetIpv6Interface) | **Get** /device/ndv6 | 



## GetIpv6Interface

> map[string]interface{} GetIpv6Interface(ctx).DeviceId(deviceId).VpnId(vpnId).IfName(ifName).Mac(mac).Execute()





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
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifName := "ifName_example" // string | Interface name (optional)
    mac := "mac_example" // string | Mac address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPv6NeighboursApi.GetIpv6Interface(context.Background()).DeviceId(deviceId).VpnId(vpnId).IfName(ifName).Mac(mac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPv6NeighboursApi.GetIpv6Interface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpv6Interface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPv6NeighboursApi.GetIpv6Interface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpv6InterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **vpnId** | **string** | VPN Id | 
 **ifName** | **string** | Interface name | 
 **mac** | **string** | Mac address | 

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

