# \MonitoringDeviceDetailsApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableSDAVCOnDevice**](MonitoringDeviceDetailsApi.md#EnableSDAVCOnDevice) | **Post** /device/enableSDAVC/{deviceIP}/{enable} | 
[**GenerateDeviceStateData**](MonitoringDeviceDetailsApi.md#GenerateDeviceStateData) | **Get** /data/device/state/{state_data_type} | 
[**GenerateDeviceStateDataFields**](MonitoringDeviceDetailsApi.md#GenerateDeviceStateDataFields) | **Get** /data/device/state/{state_data_type}/fields | 
[**GenerateDeviceStateDataWithQueryString**](MonitoringDeviceDetailsApi.md#GenerateDeviceStateDataWithQueryString) | **Get** /data/device/state/{state_data_type}/query | 
[**GetAllDeviceStatus**](MonitoringDeviceDetailsApi.md#GetAllDeviceStatus) | **Get** /device/status | 
[**GetDeviceCounters**](MonitoringDeviceDetailsApi.md#GetDeviceCounters) | **Get** /device/counters | 
[**GetDeviceListAsKeyValue**](MonitoringDeviceDetailsApi.md#GetDeviceListAsKeyValue) | **Get** /device/keyvalue | 
[**GetDeviceModels**](MonitoringDeviceDetailsApi.md#GetDeviceModels) | **Get** /device/models/{uuid} | 
[**GetDeviceOnlyStatus**](MonitoringDeviceDetailsApi.md#GetDeviceOnlyStatus) | **Get** /device/devicestatus | 
[**GetDeviceRunningConfig**](MonitoringDeviceDetailsApi.md#GetDeviceRunningConfig) | **Get** /device/config | 
[**GetDeviceRunningConfigHTML**](MonitoringDeviceDetailsApi.md#GetDeviceRunningConfigHTML) | **Get** /device/config/html | 
[**GetDeviceTlocStatus**](MonitoringDeviceDetailsApi.md#GetDeviceTlocStatus) | **Get** /device/tloc | 
[**GetDeviceTlocUtil**](MonitoringDeviceDetailsApi.md#GetDeviceTlocUtil) | **Get** /device/tlocutil | 
[**GetDeviceTlocUtilDetails**](MonitoringDeviceDetailsApi.md#GetDeviceTlocUtilDetails) | **Get** /device/tlocutil/detail | 
[**GetHardwareHealthDetails**](MonitoringDeviceDetailsApi.md#GetHardwareHealthDetails) | **Get** /device/hardwarehealth/detail | 
[**GetHardwareHealthSummary**](MonitoringDeviceDetailsApi.md#GetHardwareHealthSummary) | **Get** /device/hardwarehealth/summary | 
[**GetStatsQueues**](MonitoringDeviceDetailsApi.md#GetStatsQueues) | **Get** /device/stats | 
[**GetSyncQueues**](MonitoringDeviceDetailsApi.md#GetSyncQueues) | **Get** /device/queues | 
[**GetUnconfigured**](MonitoringDeviceDetailsApi.md#GetUnconfigured) | **Get** /device/unconfigured | 
[**GetVManageSystemIP**](MonitoringDeviceDetailsApi.md#GetVManageSystemIP) | **Get** /device/vmanage | 
[**GetVedgeInventory**](MonitoringDeviceDetailsApi.md#GetVedgeInventory) | **Get** /device/vedgeinventory/detail | 
[**GetVedgeInventorySummary**](MonitoringDeviceDetailsApi.md#GetVedgeInventorySummary) | **Get** /device/vedgeinventory/summary | 
[**ListAllDeviceModels**](MonitoringDeviceDetailsApi.md#ListAllDeviceModels) | **Get** /device/models | 
[**ListAllDevices**](MonitoringDeviceDetailsApi.md#ListAllDevices) | **Get** /device | 
[**ListAllMonitorDetailsDevices**](MonitoringDeviceDetailsApi.md#ListAllMonitorDetailsDevices) | **Get** /device/monitor | 
[**ListCurrentlySyncingDevices**](MonitoringDeviceDetailsApi.md#ListCurrentlySyncingDevices) | **Get** /device/sync_status | 
[**ListReachableDevices**](MonitoringDeviceDetailsApi.md#ListReachableDevices) | **Get** /device/reachable | 
[**ListUnreachableDevices**](MonitoringDeviceDetailsApi.md#ListUnreachableDevices) | **Get** /device/unreachable | 
[**RemoveUnreachableDevice**](MonitoringDeviceDetailsApi.md#RemoveUnreachableDevice) | **Delete** /device/unreachable/{deviceIP} | 
[**SetBlockSync**](MonitoringDeviceDetailsApi.md#SetBlockSync) | **Post** /device/blockSync | 
[**SyncAllDevicesMemDB**](MonitoringDeviceDetailsApi.md#SyncAllDevicesMemDB) | **Post** /device/syncall/memorydb | 



## EnableSDAVCOnDevice

> EnableSDAVCOnDevice(ctx, deviceIP, enable).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    enable := true // bool | Enable/Disable flag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.EnableSDAVCOnDevice(context.Background(), deviceIP, enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.EnableSDAVCOnDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 
**enable** | **bool** | Enable/Disable flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSDAVCOnDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDeviceStateData

> map[string]interface{} GenerateDeviceStateData(ctx, stateDataType).StartId(startId).Count(count).Execute()





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
    stateDataType := "stateDataType_example" // string | State data type
    startId := "startId_example" // string | Start Id
    count := "count_example" // string | Count (optional) (default to "1000")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateData(context.Background(), stateDataType).StartId(startId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GenerateDeviceStateData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceStateData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GenerateDeviceStateData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceStateDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Start Id | 
 **count** | **string** | Count | [default to &quot;1000&quot;]

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


## GenerateDeviceStateDataFields

> map[string]interface{} GenerateDeviceStateDataFields(ctx, stateDataType).Execute()





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
    stateDataType := "stateDataType_example" // string | State data type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateDataFields(context.Background(), stateDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GenerateDeviceStateDataFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceStateDataFields`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GenerateDeviceStateDataFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceStateDataFieldsRequest struct via the builder pattern


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


## GenerateDeviceStateDataWithQueryString

> map[string]interface{} GenerateDeviceStateDataWithQueryString(ctx, stateDataType).Execute()





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
    stateDataType := "stateDataType_example" // string | State data type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateDataWithQueryString(context.Background(), stateDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GenerateDeviceStateDataWithQueryString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceStateDataWithQueryString`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GenerateDeviceStateDataWithQueryString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceStateDataWithQueryStringRequest struct via the builder pattern


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


## GetAllDeviceStatus

> []map[string]interface{} GetAllDeviceStatus(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetAllDeviceStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetAllDeviceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDeviceStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetAllDeviceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeviceStatusRequest struct via the builder pattern


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


## GetDeviceCounters

> map[string]interface{} GetDeviceCounters(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceCounters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceCounters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCounters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceCounters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCountersRequest struct via the builder pattern


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


## GetDeviceListAsKeyValue

> map[string]interface{} GetDeviceListAsKeyValue(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceListAsKeyValue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceListAsKeyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceListAsKeyValue`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceListAsKeyValue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceListAsKeyValueRequest struct via the builder pattern


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


## GetDeviceModels

> map[string]interface{} GetDeviceModels(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceModels(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceModels`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelsRequest struct via the builder pattern


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


## GetDeviceOnlyStatus

> []map[string]interface{} GetDeviceOnlyStatus(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceOnlyStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceOnlyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceOnlyStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceOnlyStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceOnlyStatusRequest struct via the builder pattern


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


## GetDeviceRunningConfig

> string GetDeviceRunningConfig(ctx).DeviceId(deviceId).Execute()





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
    deviceId := []string{"Inner_example"} // []string | Device Id list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceRunningConfig(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceRunningConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceRunningConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceRunningConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRunningConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **[]string** | Device Id list | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceRunningConfigHTML

> string GetDeviceRunningConfigHTML(ctx).DeviceId(deviceId).Execute()





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
    deviceId := []string{"Inner_example"} // []string | Device Id list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceRunningConfigHTML(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceRunningConfigHTML``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceRunningConfigHTML`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceRunningConfigHTML`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRunningConfigHTMLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **[]string** | Device Id list | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTlocStatus

> map[string]interface{} GetDeviceTlocStatus(ctx).DeviceId(deviceId).Color(color).Execute()





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
    deviceId := "deviceId_example" // string | Device Id (optional)
    color := "color_example" // string | Status color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocStatus(context.Background()).DeviceId(deviceId).Color(color).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceTlocStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTlocStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceTlocStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTlocStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **color** | **string** | Status color | 

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


## GetDeviceTlocUtil

> map[string]interface{} GetDeviceTlocUtil(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocUtil(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceTlocUtil``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTlocUtil`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceTlocUtil`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTlocUtilRequest struct via the builder pattern


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


## GetDeviceTlocUtilDetails

> map[string]interface{} GetDeviceTlocUtilDetails(ctx).Util(util).Execute()





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
    util := "util_example" // string | Tloc util (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocUtilDetails(context.Background()).Util(util).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetDeviceTlocUtilDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTlocUtilDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetDeviceTlocUtilDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTlocUtilDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **util** | **string** | Tloc util | 

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


## GetHardwareHealthDetails

> []map[string]interface{} GetHardwareHealthDetails(ctx).DeviceId(deviceId).State(state).Execute()





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
    deviceId := "deviceId_example" // string | Device Id (optional)
    state := "state_example" // string | Device state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetHardwareHealthDetails(context.Background()).DeviceId(deviceId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetHardwareHealthDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHardwareHealthDetails`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetHardwareHealthDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHardwareHealthDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **state** | **string** | Device state | 

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


## GetHardwareHealthSummary

> []map[string]interface{} GetHardwareHealthSummary(ctx).VpnId(vpnId).IsCached(isCached).Execute()





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
    vpnId := []string{"Inner_example"} // []string | VPN Id
    isCached := true // bool | Status cached (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetHardwareHealthSummary(context.Background()).VpnId(vpnId).IsCached(isCached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetHardwareHealthSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHardwareHealthSummary`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetHardwareHealthSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHardwareHealthSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | **[]string** | VPN Id | 
 **isCached** | **bool** | Status cached | [default to false]

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


## GetStatsQueues

> map[string]interface{} GetStatsQueues(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetStatsQueues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetStatsQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsQueues`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetStatsQueues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsQueuesRequest struct via the builder pattern


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


## GetSyncQueues

> map[string]interface{} GetSyncQueues(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetSyncQueues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetSyncQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyncQueues`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetSyncQueues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncQueuesRequest struct via the builder pattern


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


## GetUnconfigured

> []Device GetUnconfigured(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetUnconfigured(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetUnconfigured``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnconfigured`: []Device
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetUnconfigured`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnconfiguredRequest struct via the builder pattern


### Return type

[**[]Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVManageSystemIP

> map[string]interface{} GetVManageSystemIP(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetVManageSystemIP(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetVManageSystemIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVManageSystemIP`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetVManageSystemIP`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVManageSystemIPRequest struct via the builder pattern


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


## GetVedgeInventory

> map[string]interface{} GetVedgeInventory(ctx).Status(status).Execute()





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
    status := "status_example" // string | Device status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetVedgeInventory(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetVedgeInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVedgeInventory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetVedgeInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVedgeInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Device status | 

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


## GetVedgeInventorySummary

> map[string]interface{} GetVedgeInventorySummary(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.GetVedgeInventorySummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.GetVedgeInventorySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVedgeInventorySummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.GetVedgeInventorySummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVedgeInventorySummaryRequest struct via the builder pattern


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


## ListAllDeviceModels

> []map[string]interface{} ListAllDeviceModels(ctx).List(list).Execute()





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
    list := "list_example" // string | List type of device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListAllDeviceModels(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListAllDeviceModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDeviceModels`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListAllDeviceModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllDeviceModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | List type of device | 

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


## ListAllDevices

> []map[string]interface{} ListAllDevices(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListAllDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListAllDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListAllDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllDevicesRequest struct via the builder pattern


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


## ListAllMonitorDetailsDevices

> map[string]interface{} ListAllMonitorDetailsDevices(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListAllMonitorDetailsDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListAllMonitorDetailsDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllMonitorDetailsDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListAllMonitorDetailsDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllMonitorDetailsDevicesRequest struct via the builder pattern


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


## ListCurrentlySyncingDevices

> []map[string]interface{} ListCurrentlySyncingDevices(ctx).GroupId(groupId).Execute()





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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListCurrentlySyncingDevices(context.Background()).GroupId(groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListCurrentlySyncingDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCurrentlySyncingDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListCurrentlySyncingDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCurrentlySyncingDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Group Id | 

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


## ListReachableDevices

> []map[string]interface{} ListReachableDevices(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListReachableDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListReachableDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReachableDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListReachableDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListReachableDevicesRequest struct via the builder pattern


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


## ListUnreachableDevices

> []map[string]interface{} ListUnreachableDevices(ctx).Personality(personality).Execute()





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
    personality := "personality_example" // string | Device personality

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.ListUnreachableDevices(context.Background()).Personality(personality).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.ListUnreachableDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnreachableDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.ListUnreachableDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUnreachableDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personality** | **string** | Device personality | 

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


## RemoveUnreachableDevice

> RemoveUnreachableDevice(ctx, deviceIP).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.RemoveUnreachableDevice(context.Background(), deviceIP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.RemoveUnreachableDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUnreachableDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBlockSync

> map[string]interface{} SetBlockSync(ctx).BlockSync(blockSync).Execute()





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
    blockSync := "blockSync_example" // string | Block sync flag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.SetBlockSync(context.Background()).BlockSync(blockSync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.SetBlockSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBlockSync`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceDetailsApi.SetBlockSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBlockSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockSync** | **string** | Block sync flag | 

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


## SyncAllDevicesMemDB

> SyncAllDevicesMemDB(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceDetailsApi.SyncAllDevicesMemDB(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceDetailsApi.SyncAllDevicesMemDB``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAllDevicesMemDBRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

