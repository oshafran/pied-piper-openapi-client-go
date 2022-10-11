# \MonitoringStatsProcessingApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableStatisticsDemoMode**](MonitoringStatsProcessingApi.md#EnableStatisticsDemoMode) | **Get** /statistics/demomode | 
[**GenerateStatsCollectThreadReport**](MonitoringStatsProcessingApi.md#GenerateStatsCollectThreadReport) | **Get** /statistics/collect/thread/status | 
[**GenerateStatsProcessReport**](MonitoringStatsProcessingApi.md#GenerateStatsProcessReport) | **Get** /statistics/process/status | 
[**GenerateStatsProcessThreadReport**](MonitoringStatsProcessingApi.md#GenerateStatsProcessThreadReport) | **Get** /statistics/process/thread/status | 
[**GetStatisticType**](MonitoringStatsProcessingApi.md#GetStatisticType) | **Get** /statistics | 
[**GetStatisticsProcessingCounters**](MonitoringStatsProcessingApi.md#GetStatisticsProcessingCounters) | **Get** /statistics/process/counters | 
[**ProcessStatisticsData**](MonitoringStatsProcessingApi.md#ProcessStatisticsData) | **Get** /statistics/process | 
[**ResetStatsCollection**](MonitoringStatsProcessingApi.md#ResetStatsCollection) | **Get** /statistics/collection/reset/{processQueue} | 
[**StartStatsCollection**](MonitoringStatsProcessingApi.md#StartStatsCollection) | **Get** /statistics/collect | 



## EnableStatisticsDemoMode

> map[string]interface{} EnableStatisticsDemoMode(ctx).Enable(enable).Execute()





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
    enable := true // bool | Demo mode flag (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsProcessingApi.EnableStatisticsDemoMode(context.Background()).Enable(enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.EnableStatisticsDemoMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableStatisticsDemoMode`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.EnableStatisticsDemoMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableStatisticsDemoModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enable** | **bool** | Demo mode flag | [default to true]

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


## GenerateStatsCollectThreadReport

> []map[string]interface{} GenerateStatsCollectThreadReport(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.GenerateStatsCollectThreadReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.GenerateStatsCollectThreadReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateStatsCollectThreadReport`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.GenerateStatsCollectThreadReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStatsCollectThreadReportRequest struct via the builder pattern


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


## GenerateStatsProcessReport

> []map[string]interface{} GenerateStatsProcessReport(ctx).ProcessQueue(processQueue).Execute()





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
    processQueue := int64(789) // int64 | Process queue (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsProcessingApi.GenerateStatsProcessReport(context.Background()).ProcessQueue(processQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.GenerateStatsProcessReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateStatsProcessReport`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.GenerateStatsProcessReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStatsProcessReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processQueue** | **int64** | Process queue | 

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


## GenerateStatsProcessThreadReport

> []map[string]interface{} GenerateStatsProcessThreadReport(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.GenerateStatsProcessThreadReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.GenerateStatsProcessThreadReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateStatsProcessThreadReport`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.GenerateStatsProcessThreadReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStatsProcessThreadReportRequest struct via the builder pattern


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


## GetStatisticType

> []map[string]interface{} GetStatisticType(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.GetStatisticType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.GetStatisticType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatisticType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.GetStatisticType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticTypeRequest struct via the builder pattern


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


## GetStatisticsProcessingCounters

> []map[string]interface{} GetStatisticsProcessingCounters(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.GetStatisticsProcessingCounters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.GetStatisticsProcessingCounters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatisticsProcessingCounters`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.GetStatisticsProcessingCounters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsProcessingCountersRequest struct via the builder pattern


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


## ProcessStatisticsData

> map[string]interface{} ProcessStatisticsData(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.ProcessStatisticsData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.ProcessStatisticsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessStatisticsData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.ProcessStatisticsData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProcessStatisticsDataRequest struct via the builder pattern


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


## ResetStatsCollection

> map[string]interface{} ResetStatsCollection(ctx, processQueue).Execute()





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
    processQueue := int64(789) // int64 | Process queue (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsProcessingApi.ResetStatsCollection(context.Background(), processQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.ResetStatsCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetStatsCollection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.ResetStatsCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processQueue** | **int64** | Process queue | [default to -1]

### Other Parameters

Other parameters are passed through a pointer to a apiResetStatsCollectionRequest struct via the builder pattern


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


## StartStatsCollection

> map[string]interface{} StartStatsCollection(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringStatsProcessingApi.StartStatsCollection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsProcessingApi.StartStatsCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartStatsCollection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsProcessingApi.StartStatsCollection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartStatsCollectionRequest struct via the builder pattern


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

