# \MonitoringStatusApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisabledDeviceList**](MonitoringStatusApi.md#GetDisabledDeviceList) | **Get** /statistics/settings/disable/devicelist/{indexName} | 
[**GetEnabledIndexForDevice**](MonitoringStatusApi.md#GetEnabledIndexForDevice) | **Get** /statistics/settings/status/device | 
[**GetStatisticsSettings**](MonitoringStatusApi.md#GetStatisticsSettings) | **Get** /statistics/settings/status | 
[**UpdateStatisticsDeviceList**](MonitoringStatusApi.md#UpdateStatisticsDeviceList) | **Put** /statistics/settings/disable/devicelist/{indexName} | 
[**UpdateStatisticsSettings**](MonitoringStatusApi.md#UpdateStatisticsSettings) | **Put** /statistics/settings/status | 



## GetDisabledDeviceList

> map[string]interface{} GetDisabledDeviceList(ctx, indexName).Execute()





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
    indexName := "indexName_example" // string | Index name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatusApi.GetDisabledDeviceList(context.Background(), indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatusApi.GetDisabledDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDisabledDeviceList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatusApi.GetDisabledDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** | Index name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisabledDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetEnabledIndexForDevice

> map[string]interface{} GetEnabledIndexForDevice(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.MonitoringStatusApi.GetEnabledIndexForDevice(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatusApi.GetEnabledIndexForDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnabledIndexForDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatusApi.GetEnabledIndexForDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledIndexForDeviceRequest struct via the builder pattern


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


## GetStatisticsSettings

> map[string]interface{} GetStatisticsSettings(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatusApi.GetStatisticsSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatusApi.GetStatisticsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatisticsSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatusApi.GetStatisticsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsSettingsRequest struct via the builder pattern


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


## UpdateStatisticsDeviceList

> map[string]interface{} UpdateStatisticsDeviceList(ctx, indexName).Body(body).Execute()





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
    indexName := "indexName_example" // string | Index name
    body := map[string]interface{}{ ... } // map[string]interface{} | Disabled device (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatusApi.UpdateStatisticsDeviceList(context.Background(), indexName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatusApi.UpdateStatisticsDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStatisticsDeviceList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatusApi.UpdateStatisticsDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** | Index name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatisticsDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Disabled device | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatisticsSettings

> UpdateStatisticsSettings(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatusApi.UpdateStatisticsSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatusApi.UpdateStatisticsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatisticsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats setting | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

