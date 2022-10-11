# \RealTimeMonitoringVMApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudDockVMLifecycleNics**](RealTimeMonitoringVMApi.md#GetCloudDockVMLifecycleNics) | **Get** /device/vm/notifications | 
[**GetVMLifeCycleState**](RealTimeMonitoringVMApi.md#GetVMLifeCycleState) | **Get** /device/vm/state | 
[**GetVbranchVMLifecycle**](RealTimeMonitoringVMApi.md#GetVbranchVMLifecycle) | **Get** /device/vm/oper/state | 
[**GetVbranchVMLifecycleNics**](RealTimeMonitoringVMApi.md#GetVbranchVMLifecycleNics) | **Get** /device/vm/nics | 



## GetCloudDockVMLifecycleNics

> map[string]interface{} GetCloudDockVMLifecycleNics(ctx).UserGroup(userGroup).Execute()





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
    userGroup := "userGroup_example" // string | userGroup Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringVMApi.GetCloudDockVMLifecycleNics(context.Background()).UserGroup(userGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringVMApi.GetCloudDockVMLifecycleNics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDockVMLifecycleNics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringVMApi.GetCloudDockVMLifecycleNics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDockVMLifecycleNicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userGroup** | **string** | userGroup Name | 

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


## GetVMLifeCycleState

> map[string]interface{} GetVMLifeCycleState(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringVMApi.GetVMLifeCycleState(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringVMApi.GetVMLifeCycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMLifeCycleState`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringVMApi.GetVMLifeCycleState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMLifeCycleStateRequest struct via the builder pattern


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


## GetVbranchVMLifecycle

> map[string]interface{} GetVbranchVMLifecycle(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringVMApi.GetVbranchVMLifecycle(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringVMApi.GetVbranchVMLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVbranchVMLifecycle`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringVMApi.GetVbranchVMLifecycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVbranchVMLifecycleRequest struct via the builder pattern


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


## GetVbranchVMLifecycleNics

> map[string]interface{} GetVbranchVMLifecycleNics(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringVMApi.GetVbranchVMLifecycleNics(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringVMApi.GetVbranchVMLifecycleNics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVbranchVMLifecycleNics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringVMApi.GetVbranchVMLifecycleNics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVbranchVMLifecycleNicsRequest struct via the builder pattern


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

