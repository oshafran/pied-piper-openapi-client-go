# \RealTimeMonitoringUcseApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUcseStats**](RealTimeMonitoringUcseApi.md#CreateUcseStats) | **Get** /device/ucse/stats | 



## CreateUcseStats

> map[string]interface{} CreateUcseStats(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringUcseApi.CreateUcseStats(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringUcseApi.CreateUcseStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUcseStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringUcseApi.CreateUcseStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUcseStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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

