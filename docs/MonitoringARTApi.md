# \MonitoringARTApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationDataByQuery3**](MonitoringARTApi.md#GetAggregationDataByQuery3) | **Get** /statistics/art/aggregation | 
[**GetCount5**](MonitoringARTApi.md#GetCount5) | **Get** /statistics/art/doccount | 
[**GetCountPost5**](MonitoringARTApi.md#GetCountPost5) | **Post** /statistics/art/doccount | 
[**GetPostAggregationAppDataByQuery3**](MonitoringARTApi.md#GetPostAggregationAppDataByQuery3) | **Post** /statistics/art/app-agg/aggregation | 
[**GetPostAggregationDataByQuery3**](MonitoringARTApi.md#GetPostAggregationDataByQuery3) | **Post** /statistics/art/aggregation | 
[**GetPostStatBulkRawData3**](MonitoringARTApi.md#GetPostStatBulkRawData3) | **Post** /statistics/art/page | 
[**GetStatBulkRawData3**](MonitoringARTApi.md#GetStatBulkRawData3) | **Get** /statistics/art/page | 
[**GetStatDataFields5**](MonitoringARTApi.md#GetStatDataFields5) | **Get** /statistics/art/fields | 
[**GetStatDataRawData3**](MonitoringARTApi.md#GetStatDataRawData3) | **Get** /statistics/art | 
[**GetStatDataRawDataAsCSV3**](MonitoringARTApi.md#GetStatDataRawDataAsCSV3) | **Get** /statistics/art/csv | 
[**GetStatQueryFields5**](MonitoringARTApi.md#GetStatQueryFields5) | **Get** /statistics/art/query/fields | 
[**GetStatsRawData3**](MonitoringARTApi.md#GetStatsRawData3) | **Post** /statistics/art | 



## GetAggregationDataByQuery3

> map[string]interface{} GetAggregationDataByQuery3(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetAggregationDataByQuery3(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetAggregationDataByQuery3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetAggregationDataByQuery3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery3Request struct via the builder pattern


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


## GetCount5

> map[string]interface{} GetCount5(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetCount5(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetCount5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount5`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetCount5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount5Request struct via the builder pattern


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


## GetCountPost5

> map[string]interface{} GetCountPost5(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetCountPost5(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetCountPost5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost5`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetCountPost5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost5Request struct via the builder pattern


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


## GetPostAggregationAppDataByQuery3

> map[string]interface{} GetPostAggregationAppDataByQuery3(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetPostAggregationAppDataByQuery3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetPostAggregationAppDataByQuery3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetPostAggregationAppDataByQuery3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery3Request struct via the builder pattern


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


## GetPostAggregationDataByQuery3

> map[string]interface{} GetPostAggregationDataByQuery3(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetPostAggregationDataByQuery3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetPostAggregationDataByQuery3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetPostAggregationDataByQuery3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery3Request struct via the builder pattern


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


## GetPostStatBulkRawData3

> map[string]interface{} GetPostStatBulkRawData3(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetPostStatBulkRawData3(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetPostStatBulkRawData3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetPostStatBulkRawData3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData3Request struct via the builder pattern


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


## GetStatBulkRawData3

> map[string]interface{} GetStatBulkRawData3(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatBulkRawData3(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatBulkRawData3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatBulkRawData3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData3Request struct via the builder pattern


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


## GetStatDataFields5

> map[string]interface{} GetStatDataFields5(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatDataFields5(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatDataFields5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields5`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatDataFields5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields5Request struct via the builder pattern


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


## GetStatDataRawData3

> map[string]interface{} GetStatDataRawData3(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatDataRawData3(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatDataRawData3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatDataRawData3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData3Request struct via the builder pattern


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


## GetStatDataRawDataAsCSV3

> string GetStatDataRawDataAsCSV3(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatDataRawDataAsCSV3(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatDataRawDataAsCSV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV3`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatDataRawDataAsCSV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV3Request struct via the builder pattern


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


## GetStatQueryFields5

> map[string]interface{} GetStatQueryFields5(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatQueryFields5(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatQueryFields5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields5`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatQueryFields5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields5Request struct via the builder pattern


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


## GetStatsRawData3

> map[string]interface{} GetStatsRawData3(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringARTApi.GetStatsRawData3(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringARTApi.GetStatsRawData3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringARTApi.GetStatsRawData3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData3Request struct via the builder pattern


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

