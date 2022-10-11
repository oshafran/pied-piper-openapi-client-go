# \RealTimeMonitoringCloudExpressApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationsDetailList**](RealTimeMonitoringCloudExpressApi.md#CreateApplicationsDetailList) | **Get** /device/cloudx/application/detail | 
[**CreateApplicationsList**](RealTimeMonitoringCloudExpressApi.md#CreateApplicationsList) | **Get** /device/cloudx/applications | 
[**CreateGatewayExitsList**](RealTimeMonitoringCloudExpressApi.md#CreateGatewayExitsList) | **Get** /device/cloudx/gatewayexits | 
[**CreateLbApplicationsList**](RealTimeMonitoringCloudExpressApi.md#CreateLbApplicationsList) | **Get** /device/cloudx/loadbalance | 
[**CreateLocalExitsList**](RealTimeMonitoringCloudExpressApi.md#CreateLocalExitsList) | **Get** /device/cloudx/localexits | 



## CreateApplicationsDetailList

> map[string]interface{} CreateApplicationsDetailList(ctx).VpnId(vpnId).Application(application).Query(query).Execute()





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
    vpnId := "vpnId_example" // string | VPN Id (optional)
    application := "application_example" // string | Application (optional)
    query := "query_example" // string | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCloudExpressApi.CreateApplicationsDetailList(context.Background()).VpnId(vpnId).Application(application).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCloudExpressApi.CreateApplicationsDetailList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationsDetailList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCloudExpressApi.CreateApplicationsDetailList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationsDetailListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 
 **query** | **string** | Query | 

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


## CreateApplicationsList

> map[string]interface{} CreateApplicationsList(ctx).VpnId(vpnId).Application(application).Query(query).Execute()





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
    vpnId := "vpnId_example" // string | VPN Id (optional)
    application := "application_example" // string | Application (optional)
    query := "query_example" // string | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCloudExpressApi.CreateApplicationsList(context.Background()).VpnId(vpnId).Application(application).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCloudExpressApi.CreateApplicationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCloudExpressApi.CreateApplicationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 
 **query** | **string** | Query | 

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


## CreateGatewayExitsList

> map[string]interface{} CreateGatewayExitsList(ctx).DeviceId(deviceId).VpnId(vpnId).Application(application).Execute()





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
    application := "application_example" // string | Application (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCloudExpressApi.CreateGatewayExitsList(context.Background()).DeviceId(deviceId).VpnId(vpnId).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCloudExpressApi.CreateGatewayExitsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayExitsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCloudExpressApi.CreateGatewayExitsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayExitsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 

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


## CreateLbApplicationsList

> map[string]interface{} CreateLbApplicationsList(ctx).VpnId(vpnId).Application(application).Query(query).Execute()





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
    vpnId := "vpnId_example" // string | VPN Id (optional)
    application := "application_example" // string | Application (optional)
    query := "query_example" // string | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCloudExpressApi.CreateLbApplicationsList(context.Background()).VpnId(vpnId).Application(application).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCloudExpressApi.CreateLbApplicationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLbApplicationsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCloudExpressApi.CreateLbApplicationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLbApplicationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 
 **query** | **string** | Query | 

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


## CreateLocalExitsList

> map[string]interface{} CreateLocalExitsList(ctx).DeviceId(deviceId).VpnId(vpnId).Application(application).Execute()





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
    application := "application_example" // string | Application (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCloudExpressApi.CreateLocalExitsList(context.Background()).DeviceId(deviceId).VpnId(vpnId).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCloudExpressApi.CreateLocalExitsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocalExitsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCloudExpressApi.CreateLocalExitsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalExitsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 

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

