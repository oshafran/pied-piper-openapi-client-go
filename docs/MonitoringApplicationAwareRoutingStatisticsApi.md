# \MonitoringApplicationAwareRoutingStatisticsApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransportHealth**](MonitoringApplicationAwareRoutingStatisticsApi.md#GetTransportHealth) | **Get** /statistics/approute/transport/{type} | 
[**GetTransportHealthSummary**](MonitoringApplicationAwareRoutingStatisticsApi.md#GetTransportHealthSummary) | **Get** /statistics/approute/transport/summary/{type} | 



## GetTransportHealth

> map[string]interface{} GetTransportHealth(ctx, type_).Limit(limit).Query(query).Execute()





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
    type_ := "type__example" // string | Type
    limit := "limit_example" // string | Query filter
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealth(context.Background(), type_).Limit(limit).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransportHealth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | Query filter | 
 **query** | **string** |  | 

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


## GetTransportHealthSummary

> map[string]interface{} GetTransportHealthSummary(ctx, type_).Limit(limit).Query(query).Execute()





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
    type_ := "type__example" // string | Type
    limit := int64(789) // int64 | Query result size (optional) (default to 10)
    query := "{"size":25,"query":{"condition":"AND","rules":[{"field":"summary_name","type":"string","value":["loss_percentage"],"operator":"equal"},{"field":"summary_interval","type":"string","value":["last_24_hours"],"operator":"equal"}]},"plotData":[],"sort":[{"field":"loss_percentage","type":"double","order":"desc"}],"additionalProperties":{},"fields":["name","jitter","rx_octets","loss_percentage","latency","tx_octets","summary_time"]}" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealthSummary(context.Background(), type_).Limit(limit).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealthSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransportHealthSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApplicationAwareRoutingStatisticsApi.GetTransportHealthSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportHealthSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Query result size | [default to 10]
 **query** | **string** | Query filter | 

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

