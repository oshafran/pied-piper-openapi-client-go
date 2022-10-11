# \RealTimeMonitoringDREApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDreAutoBypassStats**](RealTimeMonitoringDREApi.md#GetDreAutoBypassStats) | **Get** /device/dre/auto-bypass-stats | 
[**GetDrePeerStats**](RealTimeMonitoringDREApi.md#GetDrePeerStats) | **Get** /device/dre/peer-stats | 
[**GetDreStats**](RealTimeMonitoringDREApi.md#GetDreStats) | **Get** /device/dre/dre-stats | 
[**GetDreStatus**](RealTimeMonitoringDREApi.md#GetDreStatus) | **Get** /device/dre/dre-status | 



## GetDreAutoBypassStats

> map[string]interface{} GetDreAutoBypassStats(ctx).DeviceId(deviceId).AppqoeDreAutoBypassServerIp(appqoeDreAutoBypassServerIp).AppqoeDreAutoBypassPort(appqoeDreAutoBypassPort).Execute()





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
    appqoeDreAutoBypassServerIp := "appqoeDreAutoBypassServerIp_example" // string | Server IP (optional)
    appqoeDreAutoBypassPort := float32(8.14) // float32 | Port (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDREApi.GetDreAutoBypassStats(context.Background()).DeviceId(deviceId).AppqoeDreAutoBypassServerIp(appqoeDreAutoBypassServerIp).AppqoeDreAutoBypassPort(appqoeDreAutoBypassPort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDREApi.GetDreAutoBypassStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDreAutoBypassStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDREApi.GetDreAutoBypassStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDreAutoBypassStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **appqoeDreAutoBypassServerIp** | **string** | Server IP | 
 **appqoeDreAutoBypassPort** | **float32** | Port | 

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


## GetDrePeerStats

> map[string]interface{} GetDrePeerStats(ctx).DeviceId(deviceId).AppqoeDreStatsPeerSystemIp(appqoeDreStatsPeerSystemIp).AppqoeDreStatsPeerPeerNo(appqoeDreStatsPeerPeerNo).Execute()





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
    appqoeDreStatsPeerSystemIp := "appqoeDreStatsPeerSystemIp_example" // string | System IP (optional)
    appqoeDreStatsPeerPeerNo := float32(8.14) // float32 | Peer Number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDREApi.GetDrePeerStats(context.Background()).DeviceId(deviceId).AppqoeDreStatsPeerSystemIp(appqoeDreStatsPeerSystemIp).AppqoeDreStatsPeerPeerNo(appqoeDreStatsPeerPeerNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDREApi.GetDrePeerStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrePeerStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDREApi.GetDrePeerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDrePeerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **appqoeDreStatsPeerSystemIp** | **string** | System IP | 
 **appqoeDreStatsPeerPeerNo** | **float32** | Peer Number | 

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


## GetDreStats

> map[string]interface{} GetDreStats(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringDREApi.GetDreStats(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDREApi.GetDreStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDreStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDREApi.GetDreStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDreStatsRequest struct via the builder pattern


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


## GetDreStatus

> map[string]interface{} GetDreStatus(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringDREApi.GetDreStatus(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDREApi.GetDreStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDreStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDREApi.GetDreStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDreStatusRequest struct via the builder pattern


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

