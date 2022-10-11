# \MonitoringFlowlogApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationDataByQuery26**](MonitoringFlowlogApi.md#GetAggregationDataByQuery26) | **Get** /statistics/flowlog/aggregation | 
[**GetCount28**](MonitoringFlowlogApi.md#GetCount28) | **Get** /statistics/flowlog/doccount | 
[**GetCountPost28**](MonitoringFlowlogApi.md#GetCountPost28) | **Post** /statistics/flowlog/doccount | 
[**GetPostAggregationAppDataByQuery25**](MonitoringFlowlogApi.md#GetPostAggregationAppDataByQuery25) | **Post** /statistics/flowlog/app-agg/aggregation | 
[**GetPostAggregationDataByQuery26**](MonitoringFlowlogApi.md#GetPostAggregationDataByQuery26) | **Post** /statistics/flowlog/aggregation | 
[**GetPostStatBulkRawData26**](MonitoringFlowlogApi.md#GetPostStatBulkRawData26) | **Post** /statistics/flowlog/page | 
[**GetStatBulkRawData26**](MonitoringFlowlogApi.md#GetStatBulkRawData26) | **Get** /statistics/flowlog/page | 
[**GetStatDataFields28**](MonitoringFlowlogApi.md#GetStatDataFields28) | **Get** /statistics/flowlog/fields | 
[**GetStatDataRawData25**](MonitoringFlowlogApi.md#GetStatDataRawData25) | **Get** /statistics/flowlog | 
[**GetStatDataRawDataAsCSV26**](MonitoringFlowlogApi.md#GetStatDataRawDataAsCSV26) | **Get** /statistics/flowlog/csv | 
[**GetStatQueryFields28**](MonitoringFlowlogApi.md#GetStatQueryFields28) | **Get** /statistics/flowlog/query/fields | 
[**GetStatsRawData26**](MonitoringFlowlogApi.md#GetStatsRawData26) | **Post** /statistics/flowlog | 



## GetAggregationDataByQuery26

> map[string]interface{} GetAggregationDataByQuery26(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetAggregationDataByQuery26(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetAggregationDataByQuery26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetAggregationDataByQuery26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery26Request struct via the builder pattern


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


## GetCount28

> map[string]interface{} GetCount28(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetCount28(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetCount28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetCount28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount28Request struct via the builder pattern


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


## GetCountPost28

> map[string]interface{} GetCountPost28(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetCountPost28(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetCountPost28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetCountPost28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost28Request struct via the builder pattern


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


## GetPostAggregationAppDataByQuery25

> map[string]interface{} GetPostAggregationAppDataByQuery25(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetPostAggregationAppDataByQuery25(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetPostAggregationAppDataByQuery25``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery25`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetPostAggregationAppDataByQuery25`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery25Request struct via the builder pattern


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


## GetPostAggregationDataByQuery26

> map[string]interface{} GetPostAggregationDataByQuery26(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetPostAggregationDataByQuery26(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetPostAggregationDataByQuery26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetPostAggregationDataByQuery26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery26Request struct via the builder pattern


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


## GetPostStatBulkRawData26

> map[string]interface{} GetPostStatBulkRawData26(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetPostStatBulkRawData26(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetPostStatBulkRawData26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetPostStatBulkRawData26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData26Request struct via the builder pattern


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


## GetStatBulkRawData26

> map[string]interface{} GetStatBulkRawData26(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatBulkRawData26(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatBulkRawData26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatBulkRawData26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData26Request struct via the builder pattern


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


## GetStatDataFields28

> map[string]interface{} GetStatDataFields28(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatDataFields28(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatDataFields28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatDataFields28`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields28Request struct via the builder pattern


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


## GetStatDataRawData25

> map[string]interface{} GetStatDataRawData25(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatDataRawData25(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatDataRawData25``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData25`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatDataRawData25`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData25Request struct via the builder pattern


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


## GetStatDataRawDataAsCSV26

> string GetStatDataRawDataAsCSV26(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatDataRawDataAsCSV26(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatDataRawDataAsCSV26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV26`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatDataRawDataAsCSV26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV26Request struct via the builder pattern


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


## GetStatQueryFields28

> map[string]interface{} GetStatQueryFields28(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatQueryFields28(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatQueryFields28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatQueryFields28`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields28Request struct via the builder pattern


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


## GetStatsRawData26

> map[string]interface{} GetStatsRawData26(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringFlowlogApi.GetStatsRawData26(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringFlowlogApi.GetStatsRawData26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringFlowlogApi.GetStatsRawData26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData26Request struct via the builder pattern


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

