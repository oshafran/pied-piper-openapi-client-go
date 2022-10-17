# \MonitoringURLFilteringApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationDataByQuery22**](MonitoringURLFilteringApi.md#GetAggregationDataByQuery22) | **Get** /statistics/urlf/aggregation | 
[**GetCount24**](MonitoringURLFilteringApi.md#GetCount24) | **Get** /statistics/urlf/doccount | 
[**GetCountPost24**](MonitoringURLFilteringApi.md#GetCountPost24) | **Post** /statistics/urlf/doccount | 
[**GetPostAggregationAppDataByQuery21**](MonitoringURLFilteringApi.md#GetPostAggregationAppDataByQuery21) | **Post** /statistics/urlf/app-agg/aggregation | 
[**GetPostAggregationDataByQuery21**](MonitoringURLFilteringApi.md#GetPostAggregationDataByQuery21) | **Post** /statistics/urlf/aggregation | 
[**GetPostStatBulkRawData22**](MonitoringURLFilteringApi.md#GetPostStatBulkRawData22) | **Post** /statistics/urlf/page | 
[**GetStatBulkRawData22**](MonitoringURLFilteringApi.md#GetStatBulkRawData22) | **Get** /statistics/urlf/page | 
[**GetStatDataFields24**](MonitoringURLFilteringApi.md#GetStatDataFields24) | **Get** /statistics/urlf/fields | 
[**GetStatDataRawData21**](MonitoringURLFilteringApi.md#GetStatDataRawData21) | **Get** /statistics/urlf | 
[**GetStatDataRawDataAsCSV22**](MonitoringURLFilteringApi.md#GetStatDataRawDataAsCSV22) | **Get** /statistics/urlf/csv | 
[**GetStatQueryFields24**](MonitoringURLFilteringApi.md#GetStatQueryFields24) | **Get** /statistics/urlf/query/fields | 
[**GetStatsRawData22**](MonitoringURLFilteringApi.md#GetStatsRawData22) | **Post** /statistics/urlf | 



## GetAggregationDataByQuery22

> map[string]interface{} GetAggregationDataByQuery22(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetAggregationDataByQuery22(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetAggregationDataByQuery22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery22`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetAggregationDataByQuery22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery22Request struct via the builder pattern


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


## GetCount24

> map[string]interface{} GetCount24(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetCount24(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetCount24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetCount24`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount24Request struct via the builder pattern


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


## GetCountPost24

> map[string]interface{} GetCountPost24(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetCountPost24(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetCountPost24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetCountPost24`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost24Request struct via the builder pattern


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


## GetPostAggregationAppDataByQuery21

> map[string]interface{} GetPostAggregationAppDataByQuery21(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetPostAggregationAppDataByQuery21(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetPostAggregationAppDataByQuery21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetPostAggregationAppDataByQuery21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery21Request struct via the builder pattern


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


## GetPostAggregationDataByQuery21

> map[string]interface{} GetPostAggregationDataByQuery21(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetPostAggregationDataByQuery21(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetPostAggregationDataByQuery21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetPostAggregationDataByQuery21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery21Request struct via the builder pattern


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


## GetPostStatBulkRawData22

> map[string]interface{} GetPostStatBulkRawData22(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetPostStatBulkRawData22(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetPostStatBulkRawData22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData22`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetPostStatBulkRawData22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData22Request struct via the builder pattern


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


## GetStatBulkRawData22

> map[string]interface{} GetStatBulkRawData22(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatBulkRawData22(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatBulkRawData22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData22`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatBulkRawData22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData22Request struct via the builder pattern


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


## GetStatDataFields24

> map[string]interface{} GetStatDataFields24(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatDataFields24(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatDataFields24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatDataFields24`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields24Request struct via the builder pattern


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


## GetStatDataRawData21

> map[string]interface{} GetStatDataRawData21(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatDataRawData21(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatDataRawData21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData21`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatDataRawData21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData21Request struct via the builder pattern


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


## GetStatDataRawDataAsCSV22

> string GetStatDataRawDataAsCSV22(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatDataRawDataAsCSV22(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatDataRawDataAsCSV22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV22`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatDataRawDataAsCSV22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV22Request struct via the builder pattern


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


## GetStatQueryFields24

> map[string]interface{} GetStatQueryFields24(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatQueryFields24(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatQueryFields24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatQueryFields24`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields24Request struct via the builder pattern


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


## GetStatsRawData22

> map[string]interface{} GetStatsRawData22(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringURLFilteringApi.GetStatsRawData22(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringURLFilteringApi.GetStatsRawData22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData22`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringURLFilteringApi.GetStatsRawData22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData22Request struct via the builder pattern


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

