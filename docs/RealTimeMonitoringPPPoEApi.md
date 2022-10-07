# \RealTimeMonitoringPPPoEApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePPPoEInterfaceList**](RealTimeMonitoringPPPoEApi.md#CreatePPPoEInterfaceList) | **Get** /device/pppoe/session | 
[**CreatePPPoENeighborList**](RealTimeMonitoringPPPoEApi.md#CreatePPPoENeighborList) | **Get** /device/pppoe/statistic | 



## CreatePPPoEInterfaceList

> map[string]interface{} CreatePPPoEInterfaceList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := map[string][]openapiclient.DeviceIP{ ... } // DeviceIP | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringPPPoEApi.CreatePPPoEInterfaceList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringPPPoEApi.CreatePPPoEInterfaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePPPoEInterfaceList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringPPPoEApi.CreatePPPoEInterfaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPPoEInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | [**DeviceIP**](DeviceIP.md) | Device IP | 

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


## CreatePPPoENeighborList

> map[string]interface{} CreatePPPoENeighborList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := map[string][]openapiclient.DeviceIP{ ... } // DeviceIP | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringPPPoEApi.CreatePPPoENeighborList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringPPPoEApi.CreatePPPoENeighborList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePPPoENeighborList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringPPPoEApi.CreatePPPoENeighborList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPPoENeighborListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | [**DeviceIP**](DeviceIP.md) | Device IP | 

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

