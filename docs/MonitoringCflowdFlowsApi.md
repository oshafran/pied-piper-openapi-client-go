# \MonitoringCflowdFlowsApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCflowCollectorList**](MonitoringCflowdFlowsApi.md#CreateCflowCollectorList) | **Get** /device/cflowd/flows | 
[**CreateCflowdCollectorList**](MonitoringCflowdFlowsApi.md#CreateCflowdCollectorList) | **Get** /device/cflowd/collector | 
[**CreateCflowdFlowsCountList**](MonitoringCflowdFlowsApi.md#CreateCflowdFlowsCountList) | **Get** /device/cflowd/flows-count | 
[**CreateCflowdStatistics**](MonitoringCflowdFlowsApi.md#CreateCflowdStatistics) | **Get** /device/cflowd/statistics | 
[**CreateCflowdTemplate**](MonitoringCflowdFlowsApi.md#CreateCflowdTemplate) | **Get** /device/cflowd/template | 
[**GetCflowdDPIDeviceFieldJSON**](MonitoringCflowdFlowsApi.md#GetCflowdDPIDeviceFieldJSON) | **Get** /device/cflowd/application/fields | 
[**GetCflowdDPIFieldJSON**](MonitoringCflowdFlowsApi.md#GetCflowdDPIFieldJSON) | **Get** /device/cflowd/device/fields | 
[**GetFnFCacheStats**](MonitoringCflowdFlowsApi.md#GetFnFCacheStats) | **Get** /device/cflowd/fnf/cache-stats | 
[**GetFnFExportClientStats**](MonitoringCflowdFlowsApi.md#GetFnFExportClientStats) | **Get** /device/cflowd/fnf/export-client-stats | 
[**GetFnFExportStats**](MonitoringCflowdFlowsApi.md#GetFnFExportStats) | **Get** /device/cflowd/fnf/export-stats | 
[**GetFnFMonitorStats**](MonitoringCflowdFlowsApi.md#GetFnFMonitorStats) | **Get** /device/cflowd/fnf/monitor-stats | 
[**GetFnf**](MonitoringCflowdFlowsApi.md#GetFnf) | **Get** /device/cflowd/fnf/flow-monitor | 



## CreateCflowCollectorList

> []map[string]interface{} CreateCflowCollectorList(ctx).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).DestIp(destIp).Execute()





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
    srcIp := "srcIp_example" // string | Source IP (optional)
    destIp := "destIp_example" // string | Destination IP (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.CreateCflowCollectorList(context.Background()).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).DestIp(destIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.CreateCflowCollectorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCflowCollectorList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.CreateCflowCollectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCflowCollectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **vpnId** | **string** | VPN Id | 
 **srcIp** | **string** | Source IP | 
 **destIp** | **string** | Destination IP | 

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


## CreateCflowdCollectorList

> []map[string]interface{} CreateCflowdCollectorList(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.CreateCflowdCollectorList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.CreateCflowdCollectorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCflowdCollectorList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.CreateCflowdCollectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCflowdCollectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateCflowdFlowsCountList

> map[string]interface{} CreateCflowdFlowsCountList(ctx).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).DestIp(destIp).Execute()





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
    srcIp := "srcIp_example" // string | Source IP (optional)
    destIp := "destIp_example" // string | Destination IP (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.CreateCflowdFlowsCountList(context.Background()).DeviceId(deviceId).VpnId(vpnId).SrcIp(srcIp).DestIp(destIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.CreateCflowdFlowsCountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCflowdFlowsCountList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.CreateCflowdFlowsCountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCflowdFlowsCountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **vpnId** | **string** | VPN Id | 
 **srcIp** | **string** | Source IP | 
 **destIp** | **string** | Destination IP | 

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


## CreateCflowdStatistics

> map[string]interface{} CreateCflowdStatistics(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.CreateCflowdStatistics(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.CreateCflowdStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCflowdStatistics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.CreateCflowdStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCflowdStatisticsRequest struct via the builder pattern


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


## CreateCflowdTemplate

> map[string]interface{} CreateCflowdTemplate(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.CreateCflowdTemplate(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.CreateCflowdTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCflowdTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.CreateCflowdTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCflowdTemplateRequest struct via the builder pattern


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


## GetCflowdDPIDeviceFieldJSON

> map[string]interface{} GetCflowdDPIDeviceFieldJSON(ctx).IsDeviceDashBoard(isDeviceDashBoard).Execute()





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
    isDeviceDashBoard := true // bool | Flag whether it is device dashboard request (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetCflowdDPIDeviceFieldJSON(context.Background()).IsDeviceDashBoard(isDeviceDashBoard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetCflowdDPIDeviceFieldJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCflowdDPIDeviceFieldJSON`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetCflowdDPIDeviceFieldJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCflowdDPIDeviceFieldJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeviceDashBoard** | **bool** | Flag whether it is device dashboard request | [default to false]

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


## GetCflowdDPIFieldJSON

> map[string]interface{} GetCflowdDPIFieldJSON(ctx).IsDeviceDashBoard(isDeviceDashBoard).Execute()





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
    isDeviceDashBoard := true // bool | Flag whether it is device dashboard request (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetCflowdDPIFieldJSON(context.Background()).IsDeviceDashBoard(isDeviceDashBoard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetCflowdDPIFieldJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCflowdDPIFieldJSON`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetCflowdDPIFieldJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCflowdDPIFieldJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeviceDashBoard** | **bool** | Flag whether it is device dashboard request | [default to false]

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


## GetFnFCacheStats

> map[string]interface{} GetFnFCacheStats(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetFnFCacheStats(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetFnFCacheStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFnFCacheStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetFnFCacheStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFnFCacheStatsRequest struct via the builder pattern


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


## GetFnFExportClientStats

> map[string]interface{} GetFnFExportClientStats(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetFnFExportClientStats(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetFnFExportClientStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFnFExportClientStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetFnFExportClientStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFnFExportClientStatsRequest struct via the builder pattern


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


## GetFnFExportStats

> map[string]interface{} GetFnFExportStats(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetFnFExportStats(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetFnFExportStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFnFExportStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetFnFExportStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFnFExportStatsRequest struct via the builder pattern


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


## GetFnFMonitorStats

> map[string]interface{} GetFnFMonitorStats(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetFnFMonitorStats(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetFnFMonitorStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFnFMonitorStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetFnFMonitorStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFnFMonitorStatsRequest struct via the builder pattern


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


## GetFnf

> map[string]interface{} GetFnf(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringCflowdFlowsApi.GetFnf(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringCflowdFlowsApi.GetFnf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFnf`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringCflowdFlowsApi.GetFnf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFnfRequest struct via the builder pattern


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

