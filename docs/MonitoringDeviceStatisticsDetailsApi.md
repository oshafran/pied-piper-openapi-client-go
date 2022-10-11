# \MonitoringDeviceStatisticsDetailsApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateDeviceStatisticsData**](MonitoringDeviceStatisticsDetailsApi.md#GenerateDeviceStatisticsData) | **Get** /data/device/statistics/{state_data_type} | 
[**GetActiveAlarms**](MonitoringDeviceStatisticsDetailsApi.md#GetActiveAlarms) | **Get** /data/device/statistics/alarm/active | 
[**GetCountWithStateDataType**](MonitoringDeviceStatisticsDetailsApi.md#GetCountWithStateDataType) | **Get** /data/device/statistics/{state_data_type}/doccount | 
[**GetStatDataFieldsByStateDataType**](MonitoringDeviceStatisticsDetailsApi.md#GetStatDataFieldsByStateDataType) | **Get** /data/device/statistics/{state_data_type}/fields | 
[**GetStatisticsType**](MonitoringDeviceStatisticsDetailsApi.md#GetStatisticsType) | **Get** /data/device/statistics | 



## GenerateDeviceStatisticsData

> []map[string]interface{} GenerateDeviceStatisticsData(ctx, stateDataType).ScrollId(scrollId).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Count(count).Execute()





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
    scrollId := "scrollId_example" // string | Scroll Id
    startDate := "startDate_example" // string | Start date
    endDate := "endDate_example" // string | End date
    timeZone := "timeZone_example" // string | Time zone
    count := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GenerateDeviceStatisticsData(context.Background(), stateDataType).ScrollId(scrollId).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceStatisticsDetailsApi.GenerateDeviceStatisticsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceStatisticsData`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceStatisticsDetailsApi.GenerateDeviceStatisticsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceStatisticsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scrollId** | **string** | Scroll Id | 
 **startDate** | **string** | Start date | 
 **endDate** | **string** | End date | 
 **timeZone** | **string** | Time zone | 
 **count** | **int32** |  | 

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


## GetActiveAlarms

> []map[string]interface{} GetActiveAlarms(ctx).ScrollId(scrollId).StartDate(startDate).EndDate(endDate).Count(count).TimeZone(timeZone).Execute()





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
    scrollId := "scrollId_example" // string | SrollId
    startDate := "startDate_example" // string | Start date
    endDate := "endDate_example" // string | End date
    count := int64(789) // int64 | count
    timeZone := "timeZone_example" // string | Time zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetActiveAlarms(context.Background()).ScrollId(scrollId).StartDate(startDate).EndDate(endDate).Count(count).TimeZone(timeZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceStatisticsDetailsApi.GetActiveAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveAlarms`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceStatisticsDetailsApi.GetActiveAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scrollId** | **string** | SrollId | 
 **startDate** | **string** | Start date | 
 **endDate** | **string** | End date | 
 **count** | **int64** | count | 
 **timeZone** | **string** | Time zone | 

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


## GetCountWithStateDataType

> map[string]interface{} GetCountWithStateDataType(ctx, stateDataType).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Execute()





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
    startDate := "startDate_example" // string | Start date
    endDate := "endDate_example" // string | End date
    timeZone := "timeZone_example" // string | Time zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetCountWithStateDataType(context.Background(), stateDataType).StartDate(startDate).EndDate(endDate).TimeZone(timeZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceStatisticsDetailsApi.GetCountWithStateDataType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountWithStateDataType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceStatisticsDetailsApi.GetCountWithStateDataType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountWithStateDataTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Start date | 
 **endDate** | **string** | End date | 
 **timeZone** | **string** | Time zone | 

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


## GetStatDataFieldsByStateDataType

> []map[string]interface{} GetStatDataFieldsByStateDataType(ctx, stateDataType).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetStatDataFieldsByStateDataType(context.Background(), stateDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceStatisticsDetailsApi.GetStatDataFieldsByStateDataType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFieldsByStateDataType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceStatisticsDetailsApi.GetStatDataFieldsByStateDataType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateDataType** | **string** | State data type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFieldsByStateDataTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetStatisticsType

> []map[string]interface{} GetStatisticsType(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetStatisticsType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDeviceStatisticsDetailsApi.GetStatisticsType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatisticsType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDeviceStatisticsDetailsApi.GetStatisticsType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsTypeRequest struct via the builder pattern


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

