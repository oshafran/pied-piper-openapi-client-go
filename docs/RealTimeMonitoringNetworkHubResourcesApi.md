# \RealTimeMonitoringNetworkHubResourcesApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllocationInfo**](RealTimeMonitoringNetworkHubResourcesApi.md#GetAllocationInfo) | **Get** /device/csp/resources/cpu-info/allocation | 
[**GetCPUInfo**](RealTimeMonitoringNetworkHubResourcesApi.md#GetCPUInfo) | **Get** /device/csp/resources/cpu-info/cpus | 
[**GetVNFInfo**](RealTimeMonitoringNetworkHubResourcesApi.md#GetVNFInfo) | **Get** /device/csp/resources/cpu-info/vnfs | 



## GetAllocationInfo

> map[string]interface{} GetAllocationInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringNetworkHubResourcesApi.GetAllocationInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringNetworkHubResourcesApi.GetAllocationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllocationInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringNetworkHubResourcesApi.GetAllocationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllocationInfoRequest struct via the builder pattern


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


## GetCPUInfo

> map[string]interface{} GetCPUInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringNetworkHubResourcesApi.GetCPUInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringNetworkHubResourcesApi.GetCPUInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCPUInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringNetworkHubResourcesApi.GetCPUInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCPUInfoRequest struct via the builder pattern


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


## GetVNFInfo

> map[string]interface{} GetVNFInfo(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringNetworkHubResourcesApi.GetVNFInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringNetworkHubResourcesApi.GetVNFInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVNFInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringNetworkHubResourcesApi.GetVNFInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVNFInfoRequest struct via the builder pattern


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

