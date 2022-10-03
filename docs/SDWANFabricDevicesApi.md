# \SDWANFabricDevicesApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataserviceDeviceCountersGet**](SDWANFabricDevicesApi.md#DataserviceDeviceCountersGet) | **Get** /dataservice/device/counters | Device Counters
[**DataserviceDeviceGet**](SDWANFabricDevicesApi.md#DataserviceDeviceGet) | **Get** /dataservice/device | Fabric Devices
[**DataserviceDeviceMonitorGet**](SDWANFabricDevicesApi.md#DataserviceDeviceMonitorGet) | **Get** /dataservice/device/monitor | Devices Status
[**DataserviceStatisticsInterfaceGet**](SDWANFabricDevicesApi.md#DataserviceStatisticsInterfaceGet) | **Get** /dataservice/statistics/interface | Interface statistics



## DataserviceDeviceCountersGet

> string DataserviceDeviceCountersGet(ctx).XXSRFTOKEN(xXSRFTOKEN).Execute()

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.DataserviceDeviceCountersGet(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.DataserviceDeviceCountersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceDeviceCountersGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.DataserviceDeviceCountersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceDeviceCountersGetRequest struct via the builder pattern


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


## DataserviceDeviceGet

> string DataserviceDeviceGet(ctx).XXSRFTOKEN(xXSRFTOKEN).Execute()

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.DataserviceDeviceGet(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.DataserviceDeviceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceDeviceGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.DataserviceDeviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceDeviceGetRequest struct via the builder pattern


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


## DataserviceDeviceMonitorGet

> string DataserviceDeviceMonitorGet(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.DataserviceDeviceMonitorGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.DataserviceDeviceMonitorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceDeviceMonitorGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.DataserviceDeviceMonitorGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceDeviceMonitorGetRequest struct via the builder pattern


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


## DataserviceStatisticsInterfaceGet

> string DataserviceStatisticsInterfaceGet(ctx).XXSRFTOKEN(xXSRFTOKEN).Execute()

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANFabricDevicesApi.DataserviceStatisticsInterfaceGet(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANFabricDevicesApi.DataserviceStatisticsInterfaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceStatisticsInterfaceGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANFabricDevicesApi.DataserviceStatisticsInterfaceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceStatisticsInterfaceGetRequest struct via the builder pattern


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

