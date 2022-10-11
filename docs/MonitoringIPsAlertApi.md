# \MonitoringIPsAlertApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationDataByQuery21**](MonitoringIPsAlertApi.md#GetAggregationDataByQuery21) | **Get** /statistics/ipsalert/aggregation | 
[**GetCount23**](MonitoringIPsAlertApi.md#GetCount23) | **Get** /statistics/ipsalert/doccount | 
[**GetCountPost23**](MonitoringIPsAlertApi.md#GetCountPost23) | **Post** /statistics/ipsalert/doccount | 
[**GetPostAggregationAppDataByQuery20**](MonitoringIPsAlertApi.md#GetPostAggregationAppDataByQuery20) | **Post** /statistics/ipsalert/app-agg/aggregation | 
[**GetPostAggregationDataByQuery20**](MonitoringIPsAlertApi.md#GetPostAggregationDataByQuery20) | **Post** /statistics/ipsalert/aggregation | 
[**GetPostStatBulkRawData21**](MonitoringIPsAlertApi.md#GetPostStatBulkRawData21) | **Post** /statistics/ipsalert/page | 
[**GetStatBulkRawData21**](MonitoringIPsAlertApi.md#GetStatBulkRawData21) | **Get** /statistics/ipsalert/page | 
[**GetStatDataFields23**](MonitoringIPsAlertApi.md#GetStatDataFields23) | **Get** /statistics/ipsalert/fields | 
[**GetStatDataRawData20**](MonitoringIPsAlertApi.md#GetStatDataRawData20) | **Get** /statistics/ipsalert | 
[**GetStatDataRawDataAsCSV21**](MonitoringIPsAlertApi.md#GetStatDataRawDataAsCSV21) | **Get** /statistics/ipsalert/csv | 
[**GetStatQueryFields23**](MonitoringIPsAlertApi.md#GetStatQueryFields23) | **Get** /statistics/ipsalert/query/fields | 
[**GetStatsRawData21**](MonitoringIPsAlertApi.md#GetStatsRawData21) | **Post** /statistics/ipsalert | 



## GetAggregationDataByQuery21

> map[string]interface{} GetAggregationDataByQuery21(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetAggregationDataByQuery21(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetAggregationDataByQuery21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetAggregationDataByQuery21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## GetCount23

> map[string]interface{} GetCount23(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"]}" // string | Query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetCount23(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetCount23``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount23`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetCount23`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

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


## GetCountPost23

> map[string]interface{} GetCountPost23(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetCountPost23(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetCountPost23``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost23`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetCountPost23`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Query | 

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


## GetPostAggregationAppDataByQuery20

> map[string]interface{} GetPostAggregationAppDataByQuery20(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetPostAggregationAppDataByQuery20(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetPostAggregationAppDataByQuery20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetPostAggregationAppDataByQuery20`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostAggregationDataByQuery20

> map[string]interface{} GetPostAggregationDataByQuery20(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetPostAggregationDataByQuery20(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetPostAggregationDataByQuery20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetPostAggregationDataByQuery20`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostStatBulkRawData21

> map[string]interface{} GetPostStatBulkRawData21(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetPostStatBulkRawData21(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetPostStatBulkRawData21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetPostStatBulkRawData21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetStatBulkRawData21

> map[string]interface{} GetStatBulkRawData21(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"]}" // string | Query string (optional)
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatBulkRawData21(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatBulkRawData21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatBulkRawData21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 

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


## GetStatDataFields23

> map[string]interface{} GetStatDataFields23(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatDataFields23(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatDataFields23``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields23`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatDataFields23`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields23Request struct via the builder pattern


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


## GetStatDataRawData20

> map[string]interface{} GetStatDataRawData20(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["2020-05-10T01:00:00 UTC","2020-05-10T01:30:00 UTC"],"field":"entry_time","type":"date","operator":"between"}]},"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatDataRawData20(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatDataRawData20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatDataRawData20`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 

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


## GetStatDataRawDataAsCSV21

> string GetStatDataRawDataAsCSV21(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1000,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatDataRawDataAsCSV21(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatDataRawDataAsCSV21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV21`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatDataRawDataAsCSV21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatQueryFields23

> map[string]interface{} GetStatQueryFields23(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatQueryFields23(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatQueryFields23``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields23`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatQueryFields23`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields23Request struct via the builder pattern


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


## GetStatsRawData21

> map[string]interface{} GetStatsRawData21(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringIPsAlertApi.GetStatsRawData21(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringIPsAlertApi.GetStatsRawData21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringIPsAlertApi.GetStatsRawData21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 
 **body** | **map[string]interface{}** | Stats query string | 

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

