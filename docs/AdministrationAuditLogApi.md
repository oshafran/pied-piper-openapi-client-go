# \AdministrationAuditLogApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAuditLog**](AdministrationAuditLogApi.md#GenerateAuditLog) | **Get** /auditlog/severity | 
[**GetAuditSeverityCustomHistogram**](AdministrationAuditLogApi.md#GetAuditSeverityCustomHistogram) | **Get** /auditlog/severity/summary | 
[**GetCount**](AdministrationAuditLogApi.md#GetCount) | **Get** /auditlog/doccount | 
[**GetCountPost**](AdministrationAuditLogApi.md#GetCountPost) | **Post** /auditlog/doccount | 
[**GetPostPropertyAggregationData**](AdministrationAuditLogApi.md#GetPostPropertyAggregationData) | **Post** /auditlog/aggregation | 
[**GetPostStatBulkRawPropertyData**](AdministrationAuditLogApi.md#GetPostStatBulkRawPropertyData) | **Post** /auditlog/page | 
[**GetPropertyAggregationData**](AdministrationAuditLogApi.md#GetPropertyAggregationData) | **Get** /auditlog/aggregation | 
[**GetRawPropertyData**](AdministrationAuditLogApi.md#GetRawPropertyData) | **Post** /auditlog | 
[**GetStatBulkRawPropertyData**](AdministrationAuditLogApi.md#GetStatBulkRawPropertyData) | **Get** /auditlog/page | 
[**GetStatDataFields**](AdministrationAuditLogApi.md#GetStatDataFields) | **Get** /auditlog/fields | 
[**GetStatDataRawAuditLogData**](AdministrationAuditLogApi.md#GetStatDataRawAuditLogData) | **Get** /auditlog | 
[**GetStatQueryFields**](AdministrationAuditLogApi.md#GetStatQueryFields) | **Get** /auditlog/query/fields | 



## GenerateAuditLog

> map[string]interface{} GenerateAuditLog(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    query := "query_example" // string | Query filter (optional)
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GenerateAuditLog(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GenerateAuditLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAuditLog`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GenerateAuditLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query filter | 
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


## GetAuditSeverityCustomHistogram

> map[string]interface{} GetAuditSeverityCustomHistogram(ctx).Query(query).Execute()





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
    query := "query_example" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetAuditSeverityCustomHistogram(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetAuditSeverityCustomHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditSeverityCustomHistogram`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetAuditSeverityCustomHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditSeverityCustomHistogramRequest struct via the builder pattern


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


## GetCount

> map[string]interface{} GetCount(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.AdministrationAuditLogApi.GetCount(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountRequest struct via the builder pattern


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


## GetCountPost

> map[string]interface{} GetCountPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.AdministrationAuditLogApi.GetCountPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetCountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetCountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPostRequest struct via the builder pattern


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


## GetPostPropertyAggregationData

> map[string]interface{} GetPostPropertyAggregationData(ctx).Body(body).Execute()





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
    body := "body_example" // string | Query filter for getting stat raw data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetPostPropertyAggregationData(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetPostPropertyAggregationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostPropertyAggregationData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetPostPropertyAggregationData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostPropertyAggregationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Query filter for getting stat raw data | 

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


## GetPostStatBulkRawPropertyData

> map[string]interface{} GetPostStatBulkRawPropertyData(ctx).Body(body).ScrollId(scrollId).Count(count).Execute()





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
    body := "body_example" // string | Query filter for getting stat raw data
    scrollId := "scrollId_example" // string | Offset of the query result (optional)
    count := "count_example" // string | Size of the query result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetPostStatBulkRawPropertyData(context.Background()).Body(body).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetPostStatBulkRawPropertyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawPropertyData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetPostStatBulkRawPropertyData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawPropertyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Query filter for getting stat raw data | 
 **scrollId** | **string** | Offset of the query result | 
 **count** | **string** | Size of the query result | 

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


## GetPropertyAggregationData

> map[string]interface{} GetPropertyAggregationData(ctx).InputQuery(inputQuery).Execute()





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
    inputQuery := "inputQuery_example" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetPropertyAggregationData(context.Background()).InputQuery(inputQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetPropertyAggregationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPropertyAggregationData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetPropertyAggregationData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyAggregationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputQuery** | **string** | Query filter | 

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


## GetRawPropertyData

> map[string]interface{} GetRawPropertyData(ctx).Body(body).Execute()





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
    body := "body_example" // string | Query filter for getting stat raw data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetRawPropertyData(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetRawPropertyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawPropertyData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetRawPropertyData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRawPropertyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Query filter for getting stat raw data | 

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


## GetStatBulkRawPropertyData

> map[string]interface{} GetStatBulkRawPropertyData(ctx).InputQuery(inputQuery).ScrollId(scrollId).Count(count).Execute()





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
    inputQuery := "inputQuery_example" // string | Query filter (optional)
    scrollId := "scrollId_example" // string | Offset of the query result (optional)
    count := "count_example" // string | size of the query result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetStatBulkRawPropertyData(context.Background()).InputQuery(inputQuery).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetStatBulkRawPropertyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawPropertyData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetStatBulkRawPropertyData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawPropertyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputQuery** | **string** | Query filter | 
 **scrollId** | **string** | Offset of the query result | 
 **count** | **string** | size of the query result | 

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


## GetStatDataFields

> map[string]interface{} GetStatDataFields(ctx).Execute()





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
    resp, r, err := apiClient.AdministrationAuditLogApi.GetStatDataFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetStatDataFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetStatDataFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFieldsRequest struct via the builder pattern


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


## GetStatDataRawAuditLogData

> map[string]interface{} GetStatDataRawAuditLogData(ctx).InputQuery(inputQuery).Execute()





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
    inputQuery := "inputQuery_example" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrationAuditLogApi.GetStatDataRawAuditLogData(context.Background()).InputQuery(inputQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetStatDataRawAuditLogData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawAuditLogData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetStatDataRawAuditLogData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawAuditLogDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputQuery** | **string** | Query filter | 

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


## GetStatQueryFields

> map[string]interface{} GetStatQueryFields(ctx).Execute()





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
    resp, r, err := apiClient.AdministrationAuditLogApi.GetStatQueryFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrationAuditLogApi.GetStatQueryFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdministrationAuditLogApi.GetStatQueryFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFieldsRequest struct via the builder pattern


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

