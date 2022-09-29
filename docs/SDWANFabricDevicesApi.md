# \SDWANFabricDevicesApi

All URIs are relative to *https://https*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleComportDataserviceDeviceCountersGet**](SDWANFabricDevicesApi.md#ExampleComportDataserviceDeviceCountersGet) | **Get** //example.com:{port}/dataservice/device/counters | Device Counters
[**ExampleComportDataserviceDeviceGet**](SDWANFabricDevicesApi.md#ExampleComportDataserviceDeviceGet) | **Get** //example.com:{port}/dataservice/device | Fabric Devices
[**ExampleComportDataserviceDeviceMonitorGet**](SDWANFabricDevicesApi.md#ExampleComportDataserviceDeviceMonitorGet) | **Get** //example.com:{port}/dataservice/device/monitor | Devices Status
[**ExampleComportDataserviceStatisticsInterfaceGet**](SDWANFabricDevicesApi.md#ExampleComportDataserviceStatisticsInterfaceGet) | **Get** //example.com:{port}/dataservice/statistics/interface | Interface statistics



## ExampleComportDataserviceDeviceCountersGet

> string ExampleComportDataserviceDeviceCountersGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Device Counters

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
    port := "port_example" // string | 
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.ExampleComportDataserviceDeviceCountersGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceCountersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceDeviceCountersGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceCountersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceDeviceCountersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExampleComportDataserviceDeviceGet

> string ExampleComportDataserviceDeviceGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Fabric Devices

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
    port := "port_example" // string | 
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.ExampleComportDataserviceDeviceGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceDeviceGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceDeviceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExampleComportDataserviceDeviceMonitorGet

> string ExampleComportDataserviceDeviceMonitorGet(ctx, port).Execute()

Devices Status

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
    port := "port_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.ExampleComportDataserviceDeviceMonitorGet(context.Background(), port).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceMonitorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceDeviceMonitorGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.ExampleComportDataserviceDeviceMonitorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceDeviceMonitorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExampleComportDataserviceStatisticsInterfaceGet

> string ExampleComportDataserviceStatisticsInterfaceGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Interface statistics

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
    port := "port_example" // string | 
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.ExampleComportDataserviceStatisticsInterfaceGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.ExampleComportDataserviceStatisticsInterfaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceStatisticsInterfaceGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.ExampleComportDataserviceStatisticsInterfaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceStatisticsInterfaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

