# \RealTimeMonitoringSFPApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDetail**](RealTimeMonitoringSFPApi.md#GetDetail) | **Get** /device/sfp/detail | 
[**GetDiagnostic**](RealTimeMonitoringSFPApi.md#GetDiagnostic) | **Get** /device/sfp/diagnostic | 
[**GetDiagnosticMeasurementAlarm**](RealTimeMonitoringSFPApi.md#GetDiagnosticMeasurementAlarm) | **Get** /device/sfp/diagnosticMeasurementAlarm | 
[**GetDiagnosticMeasurementValue**](RealTimeMonitoringSFPApi.md#GetDiagnosticMeasurementValue) | **Get** /device/sfp/diagnosticMeasurementValue | 



## GetDetail

> map[string]interface{} GetDetail(ctx).DeviceId(deviceId).Ifname(ifname).Execute()





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
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringSFPApi.GetDetail(context.Background()).DeviceId(deviceId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringSFPApi.GetDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetail`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringSFPApi.GetDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **ifname** | **string** | IF Name | 

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


## GetDiagnostic

> map[string]interface{} GetDiagnostic(ctx).DeviceId(deviceId).Ifname(ifname).Execute()





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
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringSFPApi.GetDiagnostic(context.Background()).DeviceId(deviceId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringSFPApi.GetDiagnostic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnostic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringSFPApi.GetDiagnostic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **ifname** | **string** | IF Name | 

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


## GetDiagnosticMeasurementAlarm

> map[string]interface{} GetDiagnosticMeasurementAlarm(ctx).DeviceId(deviceId).Ifname(ifname).Execute()





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
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringSFPApi.GetDiagnosticMeasurementAlarm(context.Background()).DeviceId(deviceId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringSFPApi.GetDiagnosticMeasurementAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnosticMeasurementAlarm`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringSFPApi.GetDiagnosticMeasurementAlarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticMeasurementAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **ifname** | **string** | IF Name | 

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


## GetDiagnosticMeasurementValue

> map[string]interface{} GetDiagnosticMeasurementValue(ctx).DeviceId(deviceId).Ifname(ifname).Execute()





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
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringSFPApi.GetDiagnosticMeasurementValue(context.Background()).DeviceId(deviceId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringSFPApi.GetDiagnosticMeasurementValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnosticMeasurementValue`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringSFPApi.GetDiagnosticMeasurementValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticMeasurementValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **ifname** | **string** | IF Name | 

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

