# \RealTimeMonitoringDPIApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDPICollectorList**](RealTimeMonitoringDPIApi.md#CreateDPICollectorList) | **Get** /device/dpi/applications | 
[**CreateDPIFlowsList**](RealTimeMonitoringDPIApi.md#CreateDPIFlowsList) | **Get** /device/dpi/flows | 
[**CreateDPIStatistics**](RealTimeMonitoringDPIApi.md#CreateDPIStatistics) | **Get** /device/dpi/supported-applications | 
[**CreateDPISummaryRealTime**](RealTimeMonitoringDPIApi.md#CreateDPISummaryRealTime) | **Get** /device/dpi/summary | 
[**GetCommonApplicationList**](RealTimeMonitoringDPIApi.md#GetCommonApplicationList) | **Get** /device/dpi/common/applications | 
[**GetDPIDeviceDetailsFieldJSON**](RealTimeMonitoringDPIApi.md#GetDPIDeviceDetailsFieldJSON) | **Get** /device/dpi/devicedetails/fields | 
[**GetDPIDeviceFieldJSON**](RealTimeMonitoringDPIApi.md#GetDPIDeviceFieldJSON) | **Get** /device/dpi/application/fields | 
[**GetDPIFieldJSON**](RealTimeMonitoringDPIApi.md#GetDPIFieldJSON) | **Get** /device/dpi/device/fields | 
[**GetQosmosApplicationList**](RealTimeMonitoringDPIApi.md#GetQosmosApplicationList) | **Get** /device/dpi/qosmos/applications | 



## CreateDPICollectorList

> []map[string]interface{} CreateDPICollectorList(ctx).DeviceId(deviceId).VpnId(vpnId).Application(application).Family(family).Execute()





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
    application := "application_example" // string | Application (optional)
    family := "family_example" // string | Family (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.CreateDPICollectorList(context.Background()).DeviceId(deviceId).VpnId(vpnId).Application(application).Family(family).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.CreateDPICollectorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDPICollectorList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.CreateDPICollectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDPICollectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **application** | **string** | Application | 
 **family** | **string** | Family | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDPIFlowsList

> []map[string]interface{} CreateDPIFlowsList(ctx).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).Application(application).Family(family).Execute()





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
    srcIp := "srcIp_example" // string | Source IP (optional)
    application := "application_example" // string | Application (optional)
    family := "family_example" // string | Family (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.CreateDPIFlowsList(context.Background()).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).Application(application).Family(family).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.CreateDPIFlowsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDPIFlowsList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.CreateDPIFlowsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDPIFlowsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **srcIp** | **string** | Source IP | 
 **application** | **string** | Application | 
 **family** | **string** | Family | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDPIStatistics

> []map[string]interface{} CreateDPIStatistics(ctx).DeviceId(deviceId).Application(application).Family(family).Execute()





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
    application := "application_example" // string | Application (optional)
    family := "family_example" // string | Family (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.CreateDPIStatistics(context.Background()).DeviceId(deviceId).Application(application).Family(family).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.CreateDPIStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDPIStatistics`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.CreateDPIStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDPIStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **application** | **string** | Application | 
 **family** | **string** | Family | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDPISummaryRealTime

> map[string]interface{} CreateDPISummaryRealTime(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.CreateDPISummaryRealTime(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.CreateDPISummaryRealTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDPISummaryRealTime`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.CreateDPISummaryRealTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDPISummaryRealTimeRequest struct via the builder pattern


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


## GetCommonApplicationList

> []map[string]interface{} GetCommonApplicationList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.GetCommonApplicationList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.GetCommonApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonApplicationList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.GetCommonApplicationList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonApplicationListRequest struct via the builder pattern


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDPIDeviceDetailsFieldJSON

> map[string]interface{} GetDPIDeviceDetailsFieldJSON(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.GetDPIDeviceDetailsFieldJSON(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.GetDPIDeviceDetailsFieldJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDPIDeviceDetailsFieldJSON`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.GetDPIDeviceDetailsFieldJSON`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDPIDeviceDetailsFieldJSONRequest struct via the builder pattern


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


## GetDPIDeviceFieldJSON

> map[string]interface{} GetDPIDeviceFieldJSON(ctx).IsDeviceDashBoard(isDeviceDashBoard).Execute()





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
    isDeviceDashBoard := true // bool | Flag whether is device dashboard request (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.GetDPIDeviceFieldJSON(context.Background()).IsDeviceDashBoard(isDeviceDashBoard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.GetDPIDeviceFieldJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDPIDeviceFieldJSON`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.GetDPIDeviceFieldJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDPIDeviceFieldJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeviceDashBoard** | **bool** | Flag whether is device dashboard request | [default to false]

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


## GetDPIFieldJSON

> map[string]interface{} GetDPIFieldJSON(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.GetDPIFieldJSON(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.GetDPIFieldJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDPIFieldJSON`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.GetDPIFieldJSON`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDPIFieldJSONRequest struct via the builder pattern


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


## GetQosmosApplicationList

> []map[string]interface{} GetQosmosApplicationList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringDPIApi.GetQosmosApplicationList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringDPIApi.GetQosmosApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQosmosApplicationList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringDPIApi.GetQosmosApplicationList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQosmosApplicationListRequest struct via the builder pattern


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

