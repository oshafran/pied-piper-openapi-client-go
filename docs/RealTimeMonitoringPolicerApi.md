# \RealTimeMonitoringPolicerApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicedInterface**](RealTimeMonitoringPolicerApi.md#GetPolicedInterface) | **Get** /device/policer | 



## GetPolicedInterface

> map[string]interface{} GetPolicedInterface(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringPolicerApi.GetPolicedInterface(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringPolicerApi.GetPolicedInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicedInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringPolicerApi.GetPolicedInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicedInterfaceRequest struct via the builder pattern


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

