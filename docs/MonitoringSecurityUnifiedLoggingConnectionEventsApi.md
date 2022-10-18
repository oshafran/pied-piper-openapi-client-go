# \MonitoringSecurityUnifiedLoggingConnectionEventsApi

All URIs are relative to *https://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationDataByQuery14**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetAggregationDataByQuery14) | **Get** /statistics/sul/connections/aggregation | 
[**GetCount16**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetCount16) | **Get** /statistics/sul/connections/doccount | 
[**GetCountPost16**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetCountPost16) | **Post** /statistics/sul/connections/doccount | 
[**GetFilterPolicyNameList**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetFilterPolicyNameList) | **Get** /statistics/sul/connections/filter/policy_name/{policyType} | 
[**GetPostAggregationAppDataByQuery13**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetPostAggregationAppDataByQuery13) | **Post** /statistics/sul/connections/app-agg/aggregation | 
[**GetPostAggregationDataByQuery13**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetPostAggregationDataByQuery13) | **Post** /statistics/sul/connections/aggregation | 
[**GetPostStatBulkRawData14**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetPostStatBulkRawData14) | **Post** /statistics/sul/connections/page | 
[**GetStatBulkRawData14**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetStatBulkRawData14) | **Get** /statistics/sul/connections/page | 
[**GetStatDataFields16**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetStatDataFields16) | **Get** /statistics/sul/connections/fields | 
[**GetStatDataRawDataAsCSV14**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetStatDataRawDataAsCSV14) | **Get** /statistics/sul/connections/csv | 
[**GetStatQueryFields16**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetStatQueryFields16) | **Get** /statistics/sul/connections/query/fields | 
[**GetStatsRawData14**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetStatsRawData14) | **Post** /statistics/sul/connections | 
[**GetSulStatDataRawData**](MonitoringSecurityUnifiedLoggingConnectionEventsApi.md#GetSulStatDataRawData) | **Get** /statistics/sul/connections | 



## GetAggregationDataByQuery14

> map[string]interface{} GetAggregationDataByQuery14(ctx).Query(query).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["24"],"field":"entry_time","type":"date","operator":"last_n_hours"},{"value":["172.16.255.15"],"field":"vdevice_name","type":"string","operator":"in"}]}}" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetAggregationDataByQuery14(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetAggregationDataByQuery14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery14`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetAggregationDataByQuery14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery14Request struct via the builder pattern


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


## GetCount16

> map[string]interface{} GetCount16(ctx).Query(query).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["24"],"field":"entry_time","type":"date","operator":"last_n_hours"},{"value":["172.16.255.15"],"field":"vdevice_name","type":"string","operator":"in"}]}}" // string | Query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCount16(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCount16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCount16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount16Request struct via the builder pattern


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


## GetCountPost16

> map[string]interface{} GetCountPost16(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCountPost16(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCountPost16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetCountPost16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost16Request struct via the builder pattern


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


## GetFilterPolicyNameList

> []map[string]interface{} GetFilterPolicyNameList(ctx, policyType).Query(query).Execute()





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
    policyType := "policyType_example" // string | Policy type
    query := "{"query":{"condition":"AND","rules":[{"field":"vdevice_name","type":"string","operator":"in","value":["172.16.255.15"]}]}}" // string | query string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetFilterPolicyNameList(context.Background(), policyType).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetFilterPolicyNameList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFilterPolicyNameList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetFilterPolicyNameList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyType** | **string** | Policy type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterPolicyNameListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | query string | 

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


## GetPostAggregationAppDataByQuery13

> map[string]interface{} GetPostAggregationAppDataByQuery13(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationAppDataByQuery13(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationAppDataByQuery13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationAppDataByQuery13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery13Request struct via the builder pattern


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


## GetPostAggregationDataByQuery13

> map[string]interface{} GetPostAggregationDataByQuery13(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationDataByQuery13(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationDataByQuery13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostAggregationDataByQuery13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery13Request struct via the builder pattern


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


## GetPostStatBulkRawData14

> map[string]interface{} GetPostStatBulkRawData14(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostStatBulkRawData14(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostStatBulkRawData14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData14`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetPostStatBulkRawData14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData14Request struct via the builder pattern


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


## GetStatBulkRawData14

> map[string]interface{} GetStatBulkRawData14(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["168"],"field":"entry_time","type":"date","operator":"last_n_hours"},{"value":["1"],"field":"vpn_id","type":"int","operator":"in"},{"value":["p1"],"field":"fw_policy","type":"string","operator":"in"},{"value":["172.16.255.15"],"field":"vdevice_name","type":"string","operator":"in"}]}}" // string | Query string (optional)
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatBulkRawData14(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatBulkRawData14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData14`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatBulkRawData14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData14Request struct via the builder pattern


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


## GetStatDataFields16

> map[string]interface{} GetStatDataFields16(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataFields16(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataFields16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataFields16`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields16Request struct via the builder pattern


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


## GetStatDataRawDataAsCSV14

> string GetStatDataRawDataAsCSV14(ctx).Query(query).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["168"],"field":"entry_time","type":"date","operator":"last_n_hours"},{"value":["1"],"field":"vpn_id","type":"int","operator":"in"},{"value":["p1"],"field":"fw_policy","type":"string","operator":"in"},{"value":["172.16.255.15"],"field":"vdevice_name","type":"string","operator":"in"}]}}" // string | Query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataRawDataAsCSV14(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataRawDataAsCSV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV14`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatDataRawDataAsCSV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV14Request struct via the builder pattern


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


## GetStatQueryFields16

> map[string]interface{} GetStatQueryFields16(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatQueryFields16(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatQueryFields16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatQueryFields16`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields16Request struct via the builder pattern


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


## GetStatsRawData14

> map[string]interface{} GetStatsRawData14(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatsRawData14(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatsRawData14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData14`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetStatsRawData14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData14Request struct via the builder pattern


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


## GetSulStatDataRawData

> map[string]interface{} GetSulStatDataRawData(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["24"],"field":"entry_time","type":"date","operator":"last_n_hours"},{"value":["172.16.255.15"],"field":"vdevice_name","type":"string","operator":"in"}]}}" // string | Query string (optional)
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetSulStatDataRawData(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetSulStatDataRawData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSulStatDataRawData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSecurityUnifiedLoggingConnectionEventsApi.GetSulStatDataRawData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSulStatDataRawDataRequest struct via the builder pattern


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

