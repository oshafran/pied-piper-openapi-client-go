# \RealTimeMonitoringRebootHistoryApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRebootHistoryList**](RealTimeMonitoringRebootHistoryApi.md#CreateRebootHistoryList) | **Get** /device/reboothistory | 
[**CreateSyncedRebootHistoryList**](RealTimeMonitoringRebootHistoryApi.md#CreateSyncedRebootHistoryList) | **Get** /device/reboothistory/synced | 
[**GetRebootHistoryDetails**](RealTimeMonitoringRebootHistoryApi.md#GetRebootHistoryDetails) | **Get** /device/reboothistory/details | 



## CreateRebootHistoryList

> map[string]interface{} CreateRebootHistoryList(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringRebootHistoryApi.CreateRebootHistoryList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringRebootHistoryApi.CreateRebootHistoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRebootHistoryList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringRebootHistoryApi.CreateRebootHistoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRebootHistoryListRequest struct via the builder pattern


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


## CreateSyncedRebootHistoryList

> map[string]interface{} CreateSyncedRebootHistoryList(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringRebootHistoryApi.CreateSyncedRebootHistoryList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringRebootHistoryApi.CreateSyncedRebootHistoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncedRebootHistoryList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringRebootHistoryApi.CreateSyncedRebootHistoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyncedRebootHistoryListRequest struct via the builder pattern


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


## GetRebootHistoryDetails

> map[string]interface{} GetRebootHistoryDetails(ctx).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringRebootHistoryApi.GetRebootHistoryDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringRebootHistoryApi.GetRebootHistoryDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRebootHistoryDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringRebootHistoryApi.GetRebootHistoryDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootHistoryDetailsRequest struct via the builder pattern


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

