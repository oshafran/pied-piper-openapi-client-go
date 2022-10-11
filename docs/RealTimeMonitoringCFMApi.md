# \RealTimeMonitoringCFMApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMpDatabase**](RealTimeMonitoringCFMApi.md#GetMpDatabase) | **Get** /device/cfm/mp/database | 
[**GetMpLocalMep**](RealTimeMonitoringCFMApi.md#GetMpLocalMep) | **Get** /device/cfm/mp/local/mep | 
[**GetMpLocalMip**](RealTimeMonitoringCFMApi.md#GetMpLocalMip) | **Get** /device/cfm/mp/local/mip | 
[**GetMpRemoteMep**](RealTimeMonitoringCFMApi.md#GetMpRemoteMep) | **Get** /device/cfm/mp/remotemep | 



## GetMpDatabase

> map[string]interface{} GetMpDatabase(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringCFMApi.GetMpDatabase(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCFMApi.GetMpDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMpDatabase`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCFMApi.GetMpDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMpDatabaseRequest struct via the builder pattern


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


## GetMpLocalMep

> map[string]interface{} GetMpLocalMep(ctx).DeviceId(deviceId).Domain(domain).Service(service).MepId(mepId).Execute()





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
    domain := "domain_example" // string | Domain Name (optional)
    service := "service_example" // string | Service Name (optional)
    mepId := float32(8.14) // float32 | MEP ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCFMApi.GetMpLocalMep(context.Background()).DeviceId(deviceId).Domain(domain).Service(service).MepId(mepId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCFMApi.GetMpLocalMep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMpLocalMep`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCFMApi.GetMpLocalMep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMpLocalMepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **domain** | **string** | Domain Name | 
 **service** | **string** | Service Name | 
 **mepId** | **float32** | MEP ID | 

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


## GetMpLocalMip

> map[string]interface{} GetMpLocalMip(ctx).DeviceId(deviceId).Level(level).Port(port).SvcInst(svcInst).Execute()





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
    level := float32(8.14) // float32 | Level (optional)
    port := "port_example" // string | Port (optional)
    svcInst := float32(8.14) // float32 | Service Instance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCFMApi.GetMpLocalMip(context.Background()).DeviceId(deviceId).Level(level).Port(port).SvcInst(svcInst).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCFMApi.GetMpLocalMip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMpLocalMip`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCFMApi.GetMpLocalMip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMpLocalMipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **level** | **float32** | Level | 
 **port** | **string** | Port | 
 **svcInst** | **float32** | Service Instance | 

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


## GetMpRemoteMep

> map[string]interface{} GetMpRemoteMep(ctx).DeviceId(deviceId).Domain(domain).Service(service).LocalMepId(localMepId).RemoteMepId(remoteMepId).Execute()





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
    domain := "domain_example" // string | Domain Name (optional)
    service := "service_example" // string | Service Name (optional)
    localMepId := float32(8.14) // float32 | Local MEP ID (optional)
    remoteMepId := float32(8.14) // float32 | Remote MEP ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCFMApi.GetMpRemoteMep(context.Background()).DeviceId(deviceId).Domain(domain).Service(service).LocalMepId(localMepId).RemoteMepId(remoteMepId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCFMApi.GetMpRemoteMep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMpRemoteMep`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCFMApi.GetMpRemoteMep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMpRemoteMepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **domain** | **string** | Domain Name | 
 **service** | **string** | Service Name | 
 **localMepId** | **float32** | Local MEP ID | 
 **remoteMepId** | **float32** | Remote MEP ID | 

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

